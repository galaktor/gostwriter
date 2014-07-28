/*  Copyright 2014, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gostwriter>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package gostwriter

import "github.com/galaktor/gostwriter/uinput"

type VirtualKeyboard struct {
	device UinputDevice
//	keys map[string]Key 
}

func New(devName string, keys ...Key) (*VirtualKeyboard, error) {
	dev, err := Create(devName)
	if err != nil {
		return err
	}

	vk := &VirtualKeyboard{}
	vk.device = dev
	vk.register()
	vk.defKeys()

	return vk, nil	
}
