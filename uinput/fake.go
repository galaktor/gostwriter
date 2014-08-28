//    Copyright 2014, Raphael Estrada
//    Author email:  <galaktor@gmx.de>
//    Project home:  <https://github.com/galaktor/gostwriter>
//    Licensed under The GPL v3 License (see README and LICENSE files)
package uinput

import(
	"github.com/galaktor/gostwriter/key"
)

// a fake uinput implementation to be used in tests
// provides callback functions to inject behaviour
type Fake struct{
	DevicePath string
	DeviceName string
	Keys []key.Code

	OnPress   func(k key.Code) error   // set to inject behaviour on calls to Press()
	OnRelease func(k key.Code) error   // set to inject behaviour on calls to Release()
	OnSync    func() error             // set to inject behaviour on calls to Sync()
	OnDestroy func() error             // set to inject behaviour on calls to Destroy()
}

// constructs a fake uinput device. can be used as a replacement for the factory
// function on gostwriter virtual keyboard to force it to create a fake uinput
// instead of the real thing.
func (f *Fake)  New(devicePath string, deviceName string, keys ...key.Code) (D, error) {
	f.DevicePath = devicePath
	f.DeviceName = deviceName
	f.Keys = keys
	return f, nil
}

// if set, will invoke the 'OnPress()' callback. otherwise,
// does nothing. by default returns no error, pretends
// everything is working fine.
func (f *Fake) Press(k key.Code) error {
	if f.OnPress != nil {
		return f.OnPress(k)
	}
	return nil
}

// if set, will invoke the 'OnRelease()' callback. otherwise,
// does nothing. by default returns no error, pretends
// everything is working fine.
func (f *Fake) Release(k key.Code) error {
	if f.OnRelease != nil {
		return f.OnRelease(k)
	}
	return nil
}

// if set, will invoke the 'OnSync()' callback. otherwise,
// does nothing. by default returns no error, pretends
// everything is working fine.
func (f *Fake) Sync() error {
	if f.OnSync != nil {
		return f.OnSync()
	}
	return nil
}

// if set, will invoke the 'OnDestroy()' callback. otherwise,
// does nothing. by default returns no error, pretends
// everything is working fine.
func (f *Fake) Destroy() error {
	if f.OnDestroy != nil {
		return f.OnDestroy()
	}
	return nil
}
