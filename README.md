[![Travis](https://travis-ci.org/galaktor/gostwriter.svg)](https://travis-ci.org/galaktor/gostwriter)
[![GoDoc](https://godoc.org/github.com/galaktor/gostwriter?status.png)](https://godoc.org/github.com/galaktor/gostwriter)

## What is it?

It's a simple virtual keyboard for Go that uses /dev/uinput to inject key events, as if from a real keyboard.

### Why?

To programatically emulate user keystrokes, i.e. an emulated gamepad or other device that can map to keyboard inputs.

### Limitations

It currently only does keyboard events, and only the ones I needed so far. Can be extended in the scope of uinput, i.e. mouse events, but I didn't do that yet. Feel free to fork!

## Installation

Do the usual go get:

    $> go get github.com/galaktor/gostwriter

### Key codes

Since gostwriter uses uinput, you might have to regenerate the key codes it uses to match the kernel you are running on. The key codes are defined as constants in "key/codes.go", and are basically just integers which are sent to the kernel to tell it what key was pressed.

As keys are added and moved around between kernels, you might want to either add or remove keys to the codes.go file to suit your purposes. gostwriter comes with a handy script to auto-generate a complete set of keycodes from any given "linux/input.h" kernel header file.

In the simplest case, you will get gostwriter and it will just build and work. If the keycodes it comes with don't match your kernel, you will get something like this:

    # github.com/galaktor/gostwriter/input
    38: error: 'KEY_BRIGHTNESS_AUTO' undeclared (first use in this function)
    38: error: 'KEY_JOURNAL' undeclared (first use in this function)
    38: error: 'KEY_VOICECOMMAND' undeclared (first use in this function)
    38: error: 'KEY_BRIGHTNESS_TOGGLE' undeclared (first use in this function)
    39: error: 'KEY_CONTROLPANEL' undeclared (first use in this function)
    39: error: 'KEY_APPSELECT' undeclared (first use in this function)
    39: error: 'KEY_BRIGHTNESS_MIN' undeclared (first use in this function)
    39: error: 'KEY_BRIGHTNESS_MAX' undeclared (first use in this function)
    39: error: 'KEY_TASKMANAGER' undeclared (first use in this function)
    39: error: 'KEY_SCREENSAVER' undeclared (first use in this function)
    39: error: 'KEY_BUTTONCONFIG' undeclared (first use in this function)

You can either remove the offending codes from codes.go by hand, or run the provided script. It is located at "key/get_keycodes.sh" and by default will try to get your kernel's header file from "/usr/include/linux/input.h". It accepts an argument to point it to another "input.h" file (if say you want to cross-compile to another header version).

I recommend to do this for a simple and quick way to have good key codes, which will suit most cases I suppose:

    $gostwriter>       cd key
    $gostwriter/key>   ./get_keycodes.sh
    $gostwriter/key>   cd ..
    $gostwriter>       go build
    $gostwriter/build> go install github.com/galaktor/gostwriter

## Documentation

The source code and it's comments should be more than enough to get what's going on within.

### godoc

It renders reasonably well in [GoDoc](http://godoc.org/github.com/galaktor/gostwriter):

http://godoc.org/github.com/galaktor/gostwriter

### examples

I put a little demo together, it's very straight-forward. Find it at "demo/demo.go".

## Copyright

Licensed under the GPL v3. See LICENSE and COPYRIGHT files.
