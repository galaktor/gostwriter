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
	"os"
	"syscall"
)

type KeyCode C.__u16

type UinputDevice interface {
	Press(k KeyCode) error
	Release(k KeyCode) error
	Sync() error

	Destroy() error
}

type U struct {
	f *os.File
}

func New(deviceName string, keys ...KeyCode) (UinputDevice, error) {
	/* open device */
	f, err := openDeviceFile(deviceName)
	if err != nil {
		return nil, err
	}

	dev := &U{f}

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
	err = dev.create()
	if err != nil {
		return nil, err
	}

	return dev, nil
}

func openDeviceFile(deviceName string) (*os.File, error) {
	f, err := os.OpenFile(deviceName, os.O_WRONLY|syscall.O_NONBLOCK, os.ModeDevice)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// TODO: pass in device metadata, name etc
func (u *U) create() error {
	dev := C.struct_uinput_user_dev{}
	dev.name[0] = C.char('T')
	dev.name[1] = C.char('E')
	dev.name[2] = C.char('S')
	dev.name[3] = C.char('T')
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

// press then releases in one action
func (u *U) Push(k KeyCode) error {
	return nil
}

// TODO: consider re-using structs? on stack anyway, so no worries I think
func (u *U) Press(k KeyCode) error {
	evt := C.struct_input_event{}
	evt._type = C.EV_KEY
	evt.code = C.__u16(k)
	evt.value = 1 // 1 press, 0 release

	return binary.Write(u.f, binary.LittleEndian, &evt)
}

func (u *U) Release(k KeyCode) error {
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

	return nil
}

func (u *U) Destroy() error {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(u.f.Fd()), uintptr(C.UI_DEV_DESTROY), 0)
	if errno != 0 {
		return errno
	}

	return nil
}
