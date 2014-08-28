//    Copyright 2014, Raphael Estrada
//    Author email:  <galaktor@gmx.de>
//    Project home:  <https://github.com/galaktor/gostwriter>
//    Licensed under The GPL v3 License (see README and LICENSE files)
package gostwriter

// a simple virtual keyboard for linux which uses /dev/uinput to emulate keyboard input

import (
	"github.com/galaktor/gostwriter/key"
	"github.com/galaktor/gostwriter/uinput"
)

// represents the virtual keyboard. holds a reference to a device
// on /dev/uinput where it's keys will send events to.
type Keyboard struct {
	device uinput.D       // the underlying uinput device instance
	keys map[key.Code]*K  // keys are added to this map for re-use
}

// factory function to create a uinput device.
// replace this in tests to inject fake UinputDevice.
var getUinput uinput.Factory = getUinputProper

// the default uinput device factory method. creates an actual
// 'proper' uinput device.
func getUinputProper(devicePath, deviceName string, keys ...key.Code) (uinput.D, error) {
	return uinput.New(devicePath, deviceName, keys...)
}

// constructs and returns a new virtual keyboard. the name provided is used
// to name the underlying uinput device in the operating system. 
// registers *all* available key codes with uinput for simplicity. only
// actually creates virtual keys on first request to save some memory.
func New(name string) (*Keyboard, error) {
	dev, err := getUinput("/dev/uinput", name, key.ALL_CODES[0:]...)
	if err != nil {
		return nil, err
	}

	vk := &Keyboard{}
	vk.device = dev
	vk.keys = make(map[key.Code]*K)

	return vk, nil
}

// destroys an existing virtual keyboard, i.e. it unregisters
// the uinput device on the kernel. the keyboard will not work
// after this function was called. make sure to call it before
// your application exits, e.g. using 'defer'
func (kb *Keyboard) Destroy() error {
	return kb.device.Destroy()
	return nil
}

// gets a key on the keyboard. since the constructor will
// register all known key codes, you can use them all here.
// once a key is created the first time, it will be re-used
// in subsequent Get() calls.
func (kb *Keyboard) Get(c key.Code) (*K, error) {
	if k, ok := kb.keys[c]; ok {
		return k, nil
	} else {
		k, err := newK(c, kb.device)

		if err != nil {
			return nil, err
		}

		kb.keys[c] = k
		return k, nil
	}
}
