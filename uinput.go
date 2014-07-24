/*  Copyright 2014, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gostwriter>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package gostwriter

/*
  #include <linux/input.h>
  #include <linux/uinput.h>
*/
import "C"

import(
	"os"
	"syscall"
//	"encoding/binary"
)

// maybe a struct later?
type UinputDevice *os.File

// maybe use a struct per key instead?
type Key int
const(
	D Key  = C.KEY_D
	E      = C.KEY_E

)

func Create(deviceName string) (UinputDevice, error) {
	f, err := os.OpenFile(deviceName, os.O_WRONLY | syscall.O_NONBLOCK, os.ModeDevice)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// for now, EV_KEY and EV_SYN will suffice
func (u UinputDevice) registerEventTypes() error {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(f.Fd()), uintptr(C.UI_SET_EVBIT), uintptr(C.EV_KEY))
	if errno != 0 {
		return errno
	}

	_, _, errno = syscall.Syscall(syscall.SYS_IOCTL, uintptr(f.Fd()), uintptr(C.UI_SET_EVBIT), uintptr(C.EV_SYN))
	if errno != 0 {
		return errno
	}

	return nil
}

func (u UinputDevice) register(k Key) error {
	_, _, errno = syscall.Syscall(syscall.SYS_IOCTL, uintptr(f.Fd()), uintptr(C.UI_SET_KEYBIT), uintptr(C.KEY_D))
	if errno != 0 {
		return errno
	}	

	return nil
}

func (u UinputDevice) defKeys() {
	
}

func (u UinputDevice) Destroy() {

}
