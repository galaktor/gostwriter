/*  Copyright 2014, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gostwriter>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package gostwriter

import(
	"testing"

	"github.com/galaktor/gostwriter/uinput"
)

func TestNew_Always_CreatesUinputDevice(t *testing.T) {
	fake := &uinput.Fake{}
	getUinput = fake.New

	k, err := New("")

	actual := k.device
	if actual != fake {
		t.Errorf("expected fake device '%v' but found '%v'", fake, actual)
	}

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestNew_UinputDeviceName_IsAsProvided(t *testing.T) {
	expected := "abc"
	fake := &uinput.Fake{}
	getUinput = fake.New

	_, err := New(expected)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	actual := fake.DeviceName
	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestNew_UinputDevicePath_IsDevUinput(t *testing.T) {
	expected := "/dev/uinput"
	fake := &uinput.Fake{}
	getUinput = fake.New

	_, err := New(expected)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	actual := fake.DevicePath
	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestNew_UinputDeviceKeyCodes_IsGivenAllCodes(t *testing.T) {
	expected := uinput.ALL_CODES[0:]
	fake := &uinput.Fake{}
	getUinput = fake.New

	_, err := New("")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	actual := fake.Keys
	if !AreEqual(expected, actual) {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestGet_DefinedKey_ReturnsThatKey(t *testing.T) {
	t.Error("todo")
}

func TestGet_UndefinedKey_ReturnsError(t *testing.T) {
	t.Error("todo")
}

func TestDestroy_Always_DestroysUinputDevice(t *testing.T) {
	t.Error("todo")
}


func AreEqual(a, b []uinput.KeyCode) bool {
	if len(a) != len(b) {
        return false
	}

	for i := range a {
		if a[i] != b[i] {
            return false
		}
	}

    return true
}









