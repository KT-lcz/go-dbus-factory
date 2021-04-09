// Code generated by "./generator ./com.deepin.system.inputdevices"; DO NOT EDIT.

package inputdevices

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type InputDevices interface {
	inputDevices // interface com.deepin.system.InputDevices
	proxy.Object
}

type objectInputDevices struct {
	interfaceInputDevices // interface com.deepin.system.InputDevices
	proxy.ImplObject
}

func NewInputDevices(conn *dbus.Conn) InputDevices {
	obj := new(objectInputDevices)
	obj.ImplObject.Init_(conn, "com.deepin.system.InputDevices", "/com/deepin/system/InputDevices")
	return obj
}

type inputDevices interface {
	ConnectTouchscreenAdded(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error)
	ConnectTouchscreenRemoved(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error)
	Touchscreens() proxy.PropObjectPathArray
}

type interfaceInputDevices struct{}

func (v *interfaceInputDevices) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceInputDevices) GetInterfaceName_() string {
	return "com.deepin.system.InputDevices"
}

// signal TouchscreenAdded

func (v *interfaceInputDevices) ConnectTouchscreenAdded(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "TouchscreenAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".TouchscreenAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		err := dbus.Store(sig.Body, &path)
		if err == nil {
			cb(path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal TouchscreenRemoved

func (v *interfaceInputDevices) ConnectTouchscreenRemoved(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "TouchscreenRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".TouchscreenRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		err := dbus.Store(sig.Body, &path)
		if err == nil {
			cb(path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Touchscreens ao

func (v *interfaceInputDevices) Touchscreens() proxy.PropObjectPathArray {
	return &proxy.ImplPropObjectPathArray{
		Impl: v,
		Name: "Touchscreens",
	}
}

type Touchscreen interface {
	touchscreen // interface com.deepin.system.InputDevices.Touchscreen
	proxy.Object
}

type objectTouchscreen struct {
	interfaceTouchscreen // interface com.deepin.system.InputDevices.Touchscreen
	proxy.ImplObject
}

func NewTouchscreen(conn *dbus.Conn, path dbus.ObjectPath) (Touchscreen, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectTouchscreen)
	obj.ImplObject.Init_(conn, "com.deepin.system.InputDevices", path)
	return obj, nil
}

type touchscreen interface {
	DevNode() proxy.PropString
	BusType() proxy.PropString
	UUID() proxy.PropString
	Phys() proxy.PropString
	OutputName() proxy.PropString
	Width() proxy.PropDouble
	Height() proxy.PropDouble
	Name() proxy.PropString
	Serial() proxy.PropString
}

type interfaceTouchscreen struct{}

func (v *interfaceTouchscreen) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceTouchscreen) GetInterfaceName_() string {
	return "com.deepin.system.InputDevices.Touchscreen"
}

// property DevNode s

func (v *interfaceTouchscreen) DevNode() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "DevNode",
	}
}

// property BusType s

func (v *interfaceTouchscreen) BusType() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "BusType",
	}
}

// property UUID s

func (v *interfaceTouchscreen) UUID() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "UUID",
	}
}

// property Phys s

func (v *interfaceTouchscreen) Phys() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Phys",
	}
}

// property OutputName s

func (v *interfaceTouchscreen) OutputName() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "OutputName",
	}
}

// property Width d

func (v *interfaceTouchscreen) Width() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "Width",
	}
}

// property Height d

func (v *interfaceTouchscreen) Height() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "Height",
	}
}

// property Name s

func (v *interfaceTouchscreen) Name() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Name",
	}
}

// property Serial s

func (v *interfaceTouchscreen) Serial() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Serial",
	}
}
