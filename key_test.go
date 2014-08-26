/*  Copyright 2014, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gostwriter>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package gostwriter

import (
	"testing"
	"errors"

	"github.com/galaktor/gostwriter/uinput"
	"github.com/galaktor/gostwriter/key"
)

func TestState_Pressed_ReturnsPressedState(t *testing.T) {
	expected := PRESSED
	k := K{0, nil, expected}

	actual := k.State()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestState_NotPressed_ReturnsNotPressedState(t *testing.T) {
	expected := NOT_PRESSED
	k := K{0, nil, expected}

	actual := k.State()

	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestPress_NotPressed_ChangesStateToPressed(t *testing.T) {
	expected := PRESSED
	k := K{0, &uinput.Fake{}, NOT_PRESSED}

	err := k.Press()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	actual := k.state
	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestNewK_Always_SendsReleaseAndSyncToUinputDevice(t *testing.T) {
	ui := &uinput.Fake{}
	released := false
	synced := false
	ui.OnRelease = func(k key.Code) error { released = true; return nil }
	ui.OnSync  = func() error { synced = true; return nil }

	_, err := newK(0, ui)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if released != true {
		t.Errorf("uinput release call did not happen")
	}

	if synced != true {
		t.Errorf("uinput sync call did not happen")
	}
}

func TestNewK_Always_ReturnsKWithStateNOT_PRESSED(t *testing.T) {
	ui := &uinput.Fake{}

	k, err := newK(0, ui)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if k.state != NOT_PRESSED {
		t.Errorf("expected state '%v' but found '%v'", NOT_PRESSED, k.state)
	}
}

func TestPress_AlreadyPressed_RemainsPressed(t *testing.T) {
	expected := PRESSED
	k := K{0, &uinput.Fake{}, PRESSED}

	err := k.Press()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	actual := k.state
	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestRelease_NotPressed_RemainsNotPressed(t *testing.T) {
	expected := NOT_PRESSED
	k := K{0, &uinput.Fake{}, NOT_PRESSED}

	err := k.Release()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	actual := k.state
	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestPress_NotPressed_UinputPressReturnsError_RemainsNotPressed(t *testing.T) {
	expected := NOT_PRESSED
	ui := &uinput.Fake{}
	ui.OnPress = func(k key.Code) error { return errors.New("fake error") }
	k := K{0, ui, NOT_PRESSED}

	err := k.Press()

	if err == nil {
		t.Error("expected error but found nil")
	}

	actual := k.state
	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestRelease_Pressed_UinputPressReturnsError_RemainsPressed(t *testing.T) {
	expected := PRESSED
	ui := &uinput.Fake{}
	ui.OnRelease = func(k key.Code) error { return errors.New("fake error") }
	k := K{0, ui, PRESSED}

	err := k.Release()

	if err == nil {
		t.Error("expected error but found nil")
	}

	actual := k.state
	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestPress_NotPressed_SendsPressAndSyncToUinputDevice(t *testing.T) {
	ui := &uinput.Fake{}
	pressed := false
	synced := false
	ui.OnPress = func(k key.Code) error { pressed = true; return nil }
	ui.OnSync  = func() error { synced = true; return nil }
	k := K{0, ui, NOT_PRESSED}

	err := k.Press()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if pressed != true {
		t.Errorf("uinput press call did not happen")
	}

	if synced != true {
		t.Errorf("uinput sync call did not happen")
	}
}

func TestPress_AlreadyPressed_DoesNotSendPressToUinputDevice(t *testing.T) {
	ui := &uinput.Fake{}
	pressed := false
	ui.OnPress = func(k key.Code) error { pressed = true; return nil }
	k := K{0, ui, PRESSED}

	err := k.Press()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if pressed != false {
		t.Errorf("uinput press call should not have happened")
	}
}

func TestPress_AlreadyPressed_SendsSyncToUinputDevice(t *testing.T) {
	ui := &uinput.Fake{}
	synced := false
	ui.OnSync = func() error { synced = true; return nil }
	k := K{0, ui, PRESSED}

	err := k.Press()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if synced != true {
		t.Errorf("uinput sync call should have happened")
	}
}

func TestRelease_NotPressed_SendsSyncToUinputDevice(t *testing.T) {
	ui := &uinput.Fake{}
	synced := false
	ui.OnSync = func() error { synced = true; return nil }
	k := K{0, ui, NOT_PRESSED}

	err := k.Release()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if synced != true {
		t.Errorf("uinput sync call should have happened")
	}
}

func TestRelease_Pressed_SendsReleaseAndSyncToUinputDevice(t *testing.T) {
	ui := &uinput.Fake{}
	released := false
	synced := false
	ui.OnRelease = func(k key.Code) error { released = true; return nil }
	ui.OnSync  = func() error { synced = true; return nil }
	k := K{0, ui, PRESSED}

	err := k.Release()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if released != true {
		t.Errorf("uinput release call did not happen")
	}

	if synced != true {
		t.Errorf("uinput sync call did not happen")
	}
}

func TestRelease_AlreadyReleased_DoesNotSendReleaseToUinputDevice(t *testing.T) {
	ui := &uinput.Fake{}
	released := false
	ui.OnRelease = func(k key.Code) error { released = true; return nil }
	k := K{0, ui, NOT_PRESSED}

	err := k.Release()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if released != false {
		t.Errorf("uinput release call should not have happened")
	}
}
