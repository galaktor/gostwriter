/*  Copyright 2014, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gostwriter>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package uinput

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
)

const MAX_NAME_SIZE = 80

type D interface {
	Press(k KeyCode) error
	Release(k KeyCode) error
	Sync() error

	Destroy() error
}

type U struct {
	f *os.File

	// this is being a bit wasteful with memory, but ok for first shot
	registeredKeys map[KeyCode]bool
}

func New(devicePath, deviceName string, keys ...KeyCode) (D, error) {
	/* open device */
	f, err := openDeviceFile(devicePath)
	if err != nil {
		return nil, err
	}

	dev := &U{f, make(map[KeyCode]bool,len(keys))}

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

func openDeviceFile(devicePath string) (*os.File, error) {
	f, err := os.OpenFile(devicePath, os.O_WRONLY|syscall.O_NONBLOCK, os.ModeDevice)
	if err != nil {
		return nil, err
	}
	return f, nil
}

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



// TODO: consider re-using structs? on stack anyway, so no worries I think
func (u *U) Press(k KeyCode) error {
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

func (u *U) Release(k KeyCode) error {
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

// TODO: consider auto-sync with timer
func (u *U) Sync() error {
	evt := C.struct_input_event{}
	evt._type = C.EV_SYN
	evt.code = 0  // don't care
	evt.value = 0 // don't care

	return binary.Write(u.f, binary.LittleEndian, &evt)
}

// for now, EV_KEY and EV_SYN will suffice
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

func (u *U) registerKey(k KeyCode) error {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(u.f.Fd()), uintptr(C.UI_SET_KEYBIT), uintptr(k))
	if errno != 0 {
		return errno
	}

	u.registeredKeys[k] = true

	return nil
}

func (u *U) isRegistered(k KeyCode) bool {
	_, present := u.registeredKeys[k]
	return present
}

func (u *U) Destroy() error {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(u.f.Fd()), uintptr(C.UI_DEV_DESTROY), 0)
	if errno != 0 {
		return errno
	}

	return nil
}
