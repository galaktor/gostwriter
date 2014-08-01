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

	OnPress   func(k key.Code) error
	OnRelease func(k key.Code) error
	OnSync    func() error
}

func (f *Fake)  New(devicePath string, deviceName string, keys ...key.Code) (D, error) {
	f.DevicePath = devicePath
	f.DeviceName = deviceName
	f.Keys = keys
	return f, nil
}

func (f *Fake) Press(k key.Code) error {
	if f.OnPress != nil {
		return f.OnPress(k)
	}
	return nil
}

func (f *Fake) Release(k key.Code) error {
	if f.OnRelease != nil {
		return f.OnRelease(k)
	}
	return nil
}

func (f *Fake) Sync() error {
	if f.OnSync != nil {
		return f.OnSync()
	}
	return nil
}

func (f *Fake) Destroy() error {
	return nil
}


