/*  Copyright 2014, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gostwriter>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package uinput

import(
	"github.com/galaktor/gostwriter/key"
)

type Fake struct{
	DevicePath string
	DeviceName string
	Keys []key.Code
}

func (f *Fake)  New(devicePath string, deviceName string, keys ...key.Code) (D, error) {
	f.DevicePath = devicePath
	f.DeviceName = deviceName
	f.Keys = keys
	return f, nil
}

func (f *Fake) Press(k key.Code) error {
	return nil
}

func (f *Fake) Release(k key.Code) error {
	return nil
}

func (f *Fake) Sync() error {
	return nil
}

func (f *Fake) Destroy() error {
	return nil
}


