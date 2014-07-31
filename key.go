/*  Copyright 2014, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gostwriter>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package gostwriter

import (
	"github.com/galaktor/gostwriter/key"
	"github.com/galaktor/gostwriter/uinput"
)

type K struct {
	code  key.Code
	dev   uinput.D
	state State
}

type State bool

const (
	NOT_PRESSED State = State(false)
	PRESSED           = State(true)
)

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

func (k *K) State() State {
	return k.state
}

func (k *K) IsPressed() bool {
	return k.IsPressed()
}

func (k *K) Toggle() (result State, err error) {
	switch k.state {
	case PRESSED:
		err = k.Release()
	case NOT_PRESSED:
		err = k.Press()
	}

	result = k.state
	return
}

func (k *K) Press() (err error) {
	err = k.dev.Press(k.code)
	if err == nil {
		// success, update state
		k.state = PRESSED
	}
}

func (k *K) Release() error {
	err = k.dev.Release(k.code)
	if err == nil {
		// success, update state
		k.state = NOT_PRESSED
	}
}














