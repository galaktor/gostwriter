/*  Copyright 2014, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gostwriter>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package gostwriter

import (
	"github.com/galaktor/gostwriter/key"
	"github.com/galaktor/gostwriter/uinput"
)

type Key struct {

}

type Keyboard struct {
	device uinput.D
	//	keys map[string]Key
}

/* replace this in tests to inject fake UinputDevice */
var getUinput uinput.Factory = getUinputProper

func getUinputProper(devicePath, deviceName string, keys ...key.Code) (uinput.D, error) {
	return uinput.New(devicePath, deviceName, keys...)
}

/* register all codes for now */
func New(name string) (*Keyboard, error) {
	dev, err := getUinput("/dev/uinput", name, key.ALL_CODES[0:]...)
	if err != nil {
		return nil, err
	}

	vk := &Keyboard{}
	vk.device = dev

	return vk, nil
}

//func (k *Keyboard) Get(c key.Code) *







