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

type K struct {
	keyCode int
	dev UinputDevice
	isPressed bool
}

var (
	A = K{C.KEY_A}
	B = K{C.KEY_B}
	C = K{C.KEY_C}
	D = K{C.KEY_D}
	E = K{C.KEY_E}

)

func newKey(keyCode int, ) (*Key) {
	
}

func (k *K) IsPressed() bool {
	return k.IsPressed
}

func (k *K) Press() error {
	return dev.Press(k.keyCode)
}










