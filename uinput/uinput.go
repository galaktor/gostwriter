//    Copyright 2014, Raphael Estrada
//    Author email:  <galaktor@gmx.de>
//    Project home:  <https://github.com/galaktor/gostwriter>
//    Licensed under The GPL v3 License (see README and LICENSE files)
package uinput

// provides access to the uinput linux user-space virtual input API

/*
  #include <linux/input.h>
  #include <linux/uinput.h>
*/
import "C"

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"syscall"

	"github.com/galaktor/gostwriter/key"
)

// maximum length in bytes of a uinput device name as defined
// in the kernel API
const MAX_NAME_SIZE = 80

// interface to a uinput device. can be implemented by fake uinput for testing
type D interface {
	Press(k key.Code) error
	Release(k key.Code) error
	Sync() error

	Destroy() error
}

// implementation of uinput device interface
type U struct {
	f *os.File                       // file handle to the uinput device in the OS
	registeredKeys map[key.Code]bool // map legal keys to use, as registered when device is created
}

// defines the signature of the constructor. purely for more convenient
// function definition when faking the constructor. needs to be kept
// up-to-date with signature of 'New()'.
type Factory func(devicePath, deviceName string, keys ...key.Code) (D, error)

// constructs an instance of uinput via syscalls to the uinput API in the kernel.
// registers provided key codes with the device. later usage of unregistered keys
// will trigger errors - a luxury you do not get when using the raw API.
// *devicePath* is the path to uinput (typically "/dev/uinput"
// *deviceName* is the name that the device will receive in the OS
// *keys* is a list of key codes to register with the device
func New(devicePath, deviceName string, keys ...key.Code) (D, error) {
	/* open device */
	f, err := openDeviceFile(devicePath)
	if err != nil {
		return nil, err
	}

	dev := &U{f, make(map[key.Code]bool,len(keys))}

	/* set event types */
	err = dev.registerEventTypes()
	if err != nil {
		return nil, err
	}

	/* set keys */
	for _, k := range keys {
		err = dev.registerKey(k)
		if err != nil {
			return nil, err
		}
	}

	/* create device */
	err = dev.create(deviceName)
	if err != nil {
		return nil, err
	}

	return dev, nil
}

// opens a device file for use with uinput
func openDeviceFile(devicePath string) (*os.File, error) {
	f, err := os.OpenFile(devicePath, os.O_WRONLY|syscall.O_NONBLOCK, os.ModeDevice)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// creates the underlying uinput device using syscalls to the kernel API.
// TODO: set appropriate device id/vendor/product/version
func (u *U) create(deviceName string) error {
	nameB := []byte(deviceName)
	if len(nameB) > MAX_NAME_SIZE {
		msg := fmt.Sprintf("device name '%v' too long (%v bytes); cannot be longer than %v bytes", deviceName, len(nameB), MAX_NAME_SIZE)
		return errors.New(msg)
	}

	dev := C.struct_uinput_user_dev{}
	for i, c := range nameB {
		dev.name[i] = C.char(c)
	}
	dev.id.bustype = C.BUS_USB
	dev.id.vendor = 0x1234
	dev.id.product = 0xfedc
	dev.id.version = 1

	err := binary.Write(u.f, binary.LittleEndian, &dev)
	if err != nil {
		return err
	}

	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(u.f.Fd()), uintptr(C.UI_DEV_CREATE), 0)
	if errno != 0 {
		return errno
	}

	return nil
}


// sends a press event for the key code into the uinpt device. if the key
// code was not registered with the uinput api, will return an error
// TODO: consider re-using structs? on stack anyway, so no worries I think
func (u *U) Press(k key.Code) error {
	if !u.isRegistered(k) {
		msg := fmt.Sprintf("cannot press key that wasn't registered. key code: %v", k)
		return errors.New(msg)
	}

	evt := C.struct_input_event{}
	evt._type = C.EV_KEY
	evt.code = C.__u16(k)
	evt.value = 1 // 1 press, 0 release

	return binary.Write(u.f, binary.LittleEndian, &evt)
}

// sends a release event for the key code into the uinpt device. if the key
// code was not registered with the uinput api, will return an error
func (u *U) Release(k key.Code) error {
	if !u.isRegistered(k) {
		msg := fmt.Sprintf("cannot release key that wasn't registered. key code: %v", k)
		return errors.New(msg)
	}

	evt := C.struct_input_event{}
	evt._type = C.EV_KEY
	evt.code = C.__u16(k)
	evt.value = 0 // 1 press, 0 release

	return binary.Write(u.f, binary.LittleEndian, &evt)
}

// sends a sync event into the uinput device. flushes buffered
// events into the OS. use this if you don't see the events
// take effect or notice a delay.
// TODO: consider auto-sync with timer
func (u *U) Sync() error {
	evt := C.struct_input_event{}
	evt._type = C.EV_SYN
	evt.code = 0  // don't care
	evt.value = 0 // don't care

	return binary.Write(u.f, binary.LittleEndian, &evt)
}

// registers the types of events we will use on the uinput API.
// for now, EV_KEY and EV_SYN will suffice as we only emulate
// key strokes. others like relative axis events are not
// currently supported.
func (u *U) registerEventTypes() error {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(u.f.Fd()), uintptr(C.UI_SET_EVBIT), uintptr(C.EV_KEY))
	if errno != 0 {
		return errno
	}

	_, _, errno = syscall.Syscall(syscall.SYS_IOCTL, uintptr(u.f.Fd()), uintptr(C.UI_SET_EVBIT), uintptr(C.EV_SYN))
	if errno != 0 {
		return errno
	}

	return nil
}

// register a key code for use in the uinput API. add to the internal
// map of legal codes that can be used.
func (u *U) registerKey(k key.Code) error {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(u.f.Fd()), uintptr(C.UI_SET_KEYBIT), uintptr(k))
	if errno != 0 {
		return errno
	}

	u.registeredKeys[k] = true

	return nil
}

// check if a key code is registered and can be used on the uinput device
func (u *U) isRegistered(k key.Code) bool {
	_, present := u.registeredKeys[k]
	return present
}

// destroy the uinput device in the kernel using the uinput API.
func (u *U) Destroy() error {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(u.f.Fd()), uintptr(C.UI_DEV_DESTROY), 0)
	if errno != 0 {
		return errno
	}

	return nil
}
