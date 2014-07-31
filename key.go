/*  Copyright 2014, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gostwriter>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package gostwriter

import(
	"github.com/galaktor/gostwriter/input"
	"github.com/galaktor/gostwriter/uinput"
)

type K struct {
	code input.KeyCode
	dev uinput.D // should a key go to uinput directly, or should the keyboard?
	isPressed bool
}

func newKey(k input.KeyCode) (*K) {
	return &K{k,nil,false}
}

/* press the key then release it in sequence */
func (k *K) Push() error {
	err := k.Press()
	if err != nil {
		return err
	}

	err = k.Release()
	if err != nil {
		return err
	}

	return nil
}

func (k *K) IsPressed() bool {
	return k.IsPressed()
}

func (k *K) Press() error {
	return k.dev.Press(k.code)
}

func (k *K) Release() error {
	return k.dev.Release(k.code)
}









