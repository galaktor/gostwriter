//    Copyright 2014, Raphael Estrada
//    Author email:  <galaktor@gmx.de>
//    Project home:  <https://github.com/galaktor/gostwriter>
//    Licensed under The GPL v3 License (see README and LICENSE files)
package gostwriter

import(
	"testing"
	"errors"

	"github.com/galaktor/gostwriter/key"
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
	expected := key.ALL_CODES[0:]
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

func TestGet_DefinedKey_ReturnsCorrectKey(t *testing.T) {
	fake := &uinput.Fake{}
	getUinput = fake.New

	kb, err := New("")

	if err != nil {
		t.Errorf("unexpected error in New(): %v", err)
	}

	k, err := kb.Get(key.CODE_C)

	if err != nil {
		t.Errorf("unexpected error in Get(): %v", err)
	}

	actual := k.Code()
	if actual != key.CODE_C {
		t.Errorf("expected key code '%v' but found '%v'", key.CODE_C, actual)
	}
}

func TestGet_DefinedKey_CalledTwice_ReturnsSameInstance(t *testing.T) {
	fake := &uinput.Fake{}
	getUinput = fake.New

	kb, err := New("")

	if err != nil {
		t.Errorf("unexpected error in New(): %v", err)
	}

	one, err := kb.Get(key.CODE_C)

	if err != nil {
		t.Errorf("unexpected error in Get(): %v", err)
	}

	two, err := kb.Get(key.CODE_C)

	if err != nil {
		t.Errorf("unexpected error in Get(): %v", err)
	}

	if one != two {
		t.Errorf("expected same as '%x' but found different instance '%x'", one, two)
	}
}

func TestDestroy_Always_DestroysUinputDevice(t *testing.T) {
	fake := &uinput.Fake{}
	getUinput = fake.New
	destroyed := false
	fake.OnDestroy = func() error { destroyed = true; return nil; }

	kb, err := New("")

	if err != nil {
		t.Errorf("unexpected error in New(): %v", err)
	}

	err = kb.Destroy()

	if destroyed != true {
		t.Errorf("expected call to Destroy() but found none")
	}
}

func TestDestroy_UinputDestroyReturnsError_ReturnsThatError(t *testing.T) {
	fake := &uinput.Fake{}
	getUinput = fake.New
	expected := errors.New("fake error")
	fake.OnDestroy = func() error { return expected; }

	kb, err := New("")

	if err != nil {
		t.Errorf("unexpected error in New(): %v", err)
	}

	actual := kb.Destroy()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func AreEqual(a, b []key.Code) bool {
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
