//    Copyright 2014, Raphael Estrada
//    Author email:  <galaktor@gmx.de>
//    Project home:  <https://github.com/galaktor/gostwriter>
//    Licensed under The GPL v3 License (see README and LICENSE files)
package main

// a simple demo tool showing how gostwriter can be used.
// because scripting is too easy, let's automate some commands
// through the virtual keyboad instead!
// the demo opens nano, types stuff, saves, exits and prints the
// result to the terminal using 'cat'. run async from terminal,
// using '&' operator, like so
//    gostwriter/demo> go build nano.go
//    gostwriter/demo> sudo ./nano &

import (
	"log"
	"time"

	"github.com/galaktor/gostwriter"
	"github.com/galaktor/gostwriter/key"
)

func main() {
	// you could script all this in bash
	// but where's the fun in that?
	// let's gostwrite it instead :-D

	kb, err := gostwriter.New("foo")
	guard(err)

	removeDemoFile(kb)
	wait(500)
	
	typeNano(kb)
	push(kb, key.CODE_ENTER)
	wait(500)

	typeText(kb)
	wait(500)
	exitNano(kb)
	wait(500)
	typeNanoDemoTxt(kb)
	push(kb, key.CODE_ENTER)
	wait(500)

	listAndCat(kb)
		
	err = kb.Destroy()
	guard(err)
}

func removeDemoFile(kb *gostwriter.Keyboard) {
	push(kb, key.CODE_R)
	push(kb, key.CODE_M)
	push(kb, key.CODE_SPACE)
	typeNanoDemoTxt(kb)
	push(kb, key.CODE_ENTER)
}

func listAndCat(kb *gostwriter.Keyboard) {
	push(kb, key.CODE_L)
	push(kb, key.CODE_S)
	push(kb, key.CODE_ENTER)
	wait(500)

	push(kb, key.CODE_C)
	push(kb, key.CODE_A)
	push(kb, key.CODE_T)
	push(kb, key.CODE_SPACE)
	
	push(kb,    key.CODE_N)
	push(kb,    key.CODE_A)
	push(kb,    key.CODE_N)
	push(kb,    key.CODE_O)
	press(kb,   key.CODE_LEFTSHIFT)
	push(kb,    key.CODE_MINUS)
	release(kb, key.CODE_LEFTSHIFT)
	push(kb,    key.CODE_D)
	push(kb,    key.CODE_TAB) // autocomplete!
	push(kb,    key.CODE_ENTER)
}

func typeNano(kb *gostwriter.Keyboard) {
	push(kb, key.CODE_N)
	push(kb, key.CODE_A)
	push(kb, key.CODE_N)
	push(kb, key.CODE_O)
}

func typeNanoDemoTxt(kb *gostwriter.Keyboard) {
	push(kb,    key.CODE_N)
	push(kb,    key.CODE_A)
	push(kb,    key.CODE_N)
	push(kb,    key.CODE_O)
	press(kb,   key.CODE_LEFTSHIFT)
	push(kb,    key.CODE_MINUS)
	release(kb, key.CODE_LEFTSHIFT)
	push(kb,    key.CODE_D)
	push(kb,    key.CODE_E)
	push(kb,    key.CODE_M)
	push(kb,    key.CODE_O)
	push(kb,    key.CODE_DOT)
	push(kb,    key.CODE_T)
	push(kb,    key.CODE_X)
	push(kb,    key.CODE_T)
}

func exitNano(kb *gostwriter.Keyboard) {
	press(kb,   key.CODE_LEFTCTRL)
	push(kb,    key.CODE_X)
	release(kb, key.CODE_LEFTCTRL)
	push(kb,    key.CODE_Y)
}

func typeText(kb *gostwriter.Keyboard) {
	press(kb,   key.CODE_LEFTSHIFT)
	push(kb,    key.CODE_H)
	release(kb, key.CODE_LEFTSHIFT)
	push(kb,    key.CODE_E)
	push(kb,    key.CODE_L)
	push(kb,    key.CODE_L)
	push(kb,    key.CODE_O)
	push(kb,    key.CODE_COMMA)
	push(kb,    key.CODE_SPACE)
	press(kb,   key.CODE_LEFTSHIFT)
	push(kb,    key.CODE_W)
	release(kb, key.CODE_LEFTSHIFT)
	push(kb,    key.CODE_O)
	push(kb,    key.CODE_R)
	push(kb,    key.CODE_L)
	push(kb,    key.CODE_D)
	press(kb,   key.CODE_LEFTSHIFT)
	push(kb,    key.CODE_1)
	push(kb,    key.CODE_1)
	push(kb,    key.CODE_1)
	release(kb, key.CODE_LEFTSHIFT)
	push(kb,    key.CODE_1)
}


func wait(ms uint) {
	<-time.After(time.Millisecond*time.Duration(ms))
}

// presses and subsequently releases a key
func push(kb *gostwriter.Keyboard, c key.Code) {
	// pretend to be a slow, puny human
	// allow for some time between keystrokes
	wait(200)
	key, err := kb.Get(c); 	guard(err);
	err = key.Push(); guard(err);
}

// presses a key, if not already pressed. does not release
func press(kb *gostwriter.Keyboard, c key.Code) {
	// pretend to be a slow, puny human
	// allow for some time between keystrokes
	wait(200)
	key, err := kb.Get(c); 	guard(err);
	err = key.Press(); guard(err);
}

// releases a key, if not aready released.
func release(kb *gostwriter.Keyboard, c key.Code) {
	// pretend to be a slow, puny human
	// allow for some time between keystrokes
	wait(200)
	key, err := kb.Get(c); 	guard(err);
	err = key.Release(); guard(err);
}

// contains boilerplate error check. if error is present,
// prints it then exits the app
func guard(e error) {
	if e != nil {
		log.Fatal(e)
	}
}







