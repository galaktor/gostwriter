/*  Copyright 2014, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gostwriter>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package uinput

import(
	"github.com/galaktor/gostwriter/input"
)

type Fake struct{
	DevicePath string
	DeviceName string
	Keys []input.KeyCode
}

func (f *Fake)  New(devicePath string, deviceName string, keys ...input.KeyCode) (D, error) {
	f.DevicePath = devicePath
	f.DeviceName = deviceName
	f.Keys = keys
	return f, nil
}

func (f *Fake) Press(k input.KeyCode) error {
	return nil
}

func (f *Fake) Release(k input.KeyCode) error {
	return nil
}

func (f *Fake) Sync() error {
	return nil
}

func (f *Fake) Destroy() error {
	return nil
}


