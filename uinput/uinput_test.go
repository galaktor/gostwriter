/*  Copyright 2014, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gostwriter>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package uinput

import "testing"

func TestCreate_DevUinput_ReturnsDevice(t *testing.T) {
	d, err := New("/dev/uinput")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if d == nil {
		t.Error("expected device but found nil")
	}
}

/*func TestCreate_InvalidDevice_ReturnsError(t *testing.T) {
	t.Error("todo")
}*/

func TestTODO_UINPUT_INTEGRATION_TESTS(t *testing.T) {
	t.Error("todo")
}


