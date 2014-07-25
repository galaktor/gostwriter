#Travis build status
[![Build Status](https://travis-ci.org/galaktor/gostwriter.svg)](https://travis-ci.org/galaktor/gostwriter)

#What is it?
It's a simple virtual keyboard for Go that uses /dev/uinput to inject key events, as if from a real keyboard.

#Why?
To programatically emulate user keystrokes, i.e. an emulated gamepad or other device that can map to keyboard inputs.

#Limitations
It currently only does keyboard events, and only the ones I needed so far. Can be extended in the scope of uinput, i.e. mouse events, but I didn't do that yet. Feel free to fork!

#Copyright
Licensed under the GPL v3. See LICENSE and COPYRIGHT files.