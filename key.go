// Copyright 2014, Raphael Estrada
// Author email:  <galaktor@gmx.de>
// Project home:  <https://github.com/galaktor/gostwriter>
// Licensed under The GPL v3 License (see README and LICENSE files)
package gostwriter

import (
	"github.com/galaktor/gostwriter/key"
	"github.com/galaktor/gostwriter/uinput"
)

// represents a key on the virtual keyboard
type K struct {
	code  key.Code   // the code that the kernel associates with this key
	dev   uinput.D   // the /dev/uinput device that the key events will be sent to
	state State      // what state is the button in, i.e. is it pressed or not
}

// represents the state a virtual key can be in. Bool because there are 
// only two states.
type State bool

const (
	NOT_PRESSED State = State(false)  // the key is not pressed, aka released
	PRESSED           = State(true)   // the key is pressed, and not (yet) released
)

// internal constructor to create a virtual key. this is typically done by
// the virtual keyboard, not by the user. the keyboard represents the single
// device that events go through into the kernel, so the keys share that 
// device in form of the virtual keyboard
func newK(c key.Code, dev uinput.D) (*K, error) {
	k := &K{c, dev, PRESSED} // assume PRESSED before reset...
	err := k.Release()       // ...force release to reset into known state

	if err != nil {
		return nil, err
	}

	return k, nil
}

// getter which returns the keycode associated with the key
func (k *K) Code() key.Code {
	return k.code
}

// presses the key then subsequently releases it. for convenience when
// typical 'typing'-style key presses are desired. for key combinations,
// you probably should be using Press() and Release() directly.
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

// getter which returns the state that the key is in
func (k *K) State() State {
	return k.state
}

// returns whether or not the key is currently pressed
func (k *K) IsPressed() bool {
	return k.state == PRESSED
}

// switches the key from it's current state into the other.
// not sure what this would be good for, but as they say:
//   
//    "we do what we must, because we can." 
//                    -aperture science laboratories
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

// changes the key state into PRESSED if it isn't already.
// you can repeatedly call this but the state will switch
// at most once .
func (k *K) Press() (err error) {
	if k.state == PRESSED {
		k.dev.Sync()
		return nil
	}

	err = k.dev.Press(k.code)

	if err != nil {
		return err
	}

	// success, update state
	k.state = PRESSED

	err = k.dev.Sync()

	return err
}

// changes the key state into NOT_PRESSED if it isn't already
// you can repeatedly call this but the state will switch
// at most once.
func (k *K) Release() (err error) {
	if k.state == NOT_PRESSED {
		k.dev.Sync()
		return nil
	}

	err = k.dev.Release(k.code)

	if err != nil {
		return err
	}

	// success, update state
	k.state = NOT_PRESSED

	err = k.dev.Sync()

	return err
}
