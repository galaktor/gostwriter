/*  Copyright 2014, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gostwriter>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package gostwriter

import (
	"testing"

	"github.com/galaktor/gostwriter/uinput"
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

func TestNew_Always_SendsReleaseAndSyncToUinputDevice(t *testing.T) {
	t.Error("todo")
	// this MIGHT not be necessary!
	// but might serve to set into known state on creation
	// could compete with other parallel keys, though!
}

func TestPress_AlreadyPressed_RemainsPressed(t *testing.T) {
	t.Error("todo")
}

func TestRelease_AlreadyReleased_RemainsReleased(t *testing.T) {
	t.Error("todo")
}

func TestPress_NotPressed_UinputPressReturnsError_RemainsNotPressed(t *testing.T) {
	t.Error("todo")
}

func TestRelease_Pressed_UinputPressReturnsError_RemainsPressed(t *testing.T) {
	t.Error("todo")
}

func TestPress_NotPressed_SendsPressAndSyncToUinputDevice(t *testing.T) {
	t.Error("todo")
}

func TestPress_AlreadyPressed_SendsSyncToUinputDevice(t *testing.T) {
	t.Errorf("todo")
}

func TestRelease_NotPressed_SendsSyncToUinputDevice(t *testing.T) {
	t.Errorf("todo")
}

func TestRelease_Pressed_SendsReleaseAndSyncToUinputDevice(t *testing.T) {
	t.Errorf("todo")
}
