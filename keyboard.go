/*  Copyright 2014, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gostwriter>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package gostwriter

import (
	"github.com/galaktor/gostwriter/uinput"
)

type VirtualKeyboard struct {
	device uinput.D
	//	keys map[string]Key
}

/* register all codes for now */
func New() (*VirtualKeyboard, error) {
	dev, err := uinput.New("/dev/uinput", "gostwriter", uinput.ALL_CODES[0:]...)
	if err != nil {
		return nil, err
	}

	vk := &VirtualKeyboard{}
	vk.device = dev

	return vk, nil
}
