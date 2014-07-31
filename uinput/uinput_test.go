/*  Copyright 2014, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gostwriter>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package uinput

import (
	"fmt"
	"testing"
	"time"

	"github.com/galaktor/gostwriter/key"
)

/***  NOTE ON INTEGRATION TESTS ***
  These tests will fail if /dev/uinput not available, i.e.
  'uinput' kernel module not loaded or test process does
  not have permissions to access /dev/input.

  If you wish to test this uinput package as 'sudo', run the
  following to compile the tests and run the test binary after:

     you@gostwriter/uinput$>  go test -c
     you@gostwriter/uinput$>  sudo ./uinput.test -test.v=true
*/

const UINPUT_DEV_PATH = "/dev/uinput"

func TestNew_DevUinput_ReturnsDevice(t *testing.T) {
	d, err := New(UINPUT_DEV_PATH, "abc")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if d == nil {
		t.Error("expected device but found nil")
	}
}

func TestNew_InvalidDevice_ReturnsError(t *testing.T) {
	_, err := New("/something/awful", "abc")

	if err == nil {
		t.Error("expected error, but found nil")
	}
}

func TestNew_DeviceNameLongerThan80Bytes_ReturnsError(t *testing.T) {
	tooLong := fmt.Sprintf("%s", make([]byte, MAX_NAME_SIZE+1))

	_, err := New(UINPUT_DEV_PATH, tooLong)

	if err == nil {
		t.Errorf("expected error, but found nil", err, tooLong)
	}
}

func TestNew_RegisterAlLCodes_NoErrors(t *testing.T) {
	_, err := New(UINPUT_DEV_PATH, "abc", key.ALL_CODES[0:]...)
	
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestPress_UnregisteredKey_ReturnsError(t *testing.T) {
	d, err := New(UINPUT_DEV_PATH, "abc")
	defer d.Destroy()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	err = d.Press(key.CODE_C)

	if err == nil {
		t.Error("expected error, but found nil")
	}
}

func TestRelease_UnregisteredKey_ReturnsError(t *testing.T) {
	d, err := New(UINPUT_DEV_PATH, "abc")
	defer d.Destroy()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	err = d.Release(key.CODE_C)

	if err == nil {
		t.Error("expected error, but found nil")
	}
}

func TestPressThenRelease_RegisteredKey_WritesThatKeyToStdIn(t *testing.T) {
	expected := "c"
	d, err := New(UINPUT_DEV_PATH, "abc", key.CODE_C, key.CODE_ENTER)
	defer d.Destroy()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	c := make(chan string)
	go func() {
		var o string
		fmt.Scan(&o)
		c <- o
	}()

	go func() {
		k := key.CODE_C
		for {
			<-time.After(time.Second)
			t.Logf("injecting key: %v", k)

			err = d.Press(k)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			err = d.Release(k)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			err = d.Press(key.CODE_ENTER)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			err = d.Release(key.CODE_ENTER)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			err = d.Sync()
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		}
	}()

	var actual string
	select {
	case in := <-c: actual = in
	case <-time.After(time.Second * 3): t.Errorf("timed out!")
	}
	

	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestPressThenRelease_MultipleRegisteredKeys_WritesThemKeyToStdIn(t *testing.T) {
	expected := "cd"
	d, err := New(UINPUT_DEV_PATH, "abc", key.CODE_C, key.CODE_D, key.CODE_ENTER)
	defer d.Destroy()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	c := make(chan string)
	go func() {
		var o string
		fmt.Scan(&o)
		c <- o
	}()

	go func() {
		for {
			<-time.After(time.Second)
			t.Logf("injecting keys: %v and %v", key.CODE_C, key.CODE_D)

			err = d.Press(key.CODE_C)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			err = d.Release(key.CODE_C)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			err = d.Press(key.CODE_D)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			err = d.Release(key.CODE_D)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			err = d.Press(key.CODE_ENTER)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			err = d.Release(key.CODE_ENTER)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			err = d.Sync()
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		}
	}()

	var actual string
	select {
	case in := <-c: actual = in
	case <-time.After(time.Second * 3): t.Errorf("timed out!")
	}
	

	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}
