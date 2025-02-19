// Code generated by "./generator ./org.freedesktop.hostname1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package hostname1

import (
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type Hostname interface {
	hostname // interface org.freedesktop.hostname1
	proxy.Object
}

type objectHostname struct {
	interfaceHostname // interface org.freedesktop.hostname1
	proxy.ImplObject
}

func NewHostname(conn *dbus.Conn) Hostname {
	obj := new(objectHostname)
	obj.ImplObject.Init_(conn, "org.freedesktop.hostname1", "/org/freedesktop/hostname1")
	return obj
}

type hostname interface {
	GoSetHostname(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call
	SetHostname(flags dbus.Flags, arg0 string, arg1 bool) error
	GoSetStaticHostname(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call
	SetStaticHostname(flags dbus.Flags, arg0 string, arg1 bool) error
	GoSetPrettyHostname(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call
	SetPrettyHostname(flags dbus.Flags, arg0 string, arg1 bool) error
	GoSetIconName(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call
	SetIconName(flags dbus.Flags, arg0 string, arg1 bool) error
	GoSetChassis(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call
	SetChassis(flags dbus.Flags, arg0 string, arg1 bool) error
	GoSetDeployment(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call
	SetDeployment(flags dbus.Flags, arg0 string, arg1 bool) error
	GoSetLocation(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call
	SetLocation(flags dbus.Flags, arg0 string, arg1 bool) error
	GoGetProductUUID(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call
	GetProductUUID(flags dbus.Flags, arg0 bool) ([]uint8, error)
	Hostname() proxy.PropString
	StaticHostname() proxy.PropString
	PrettyHostname() proxy.PropString
	IconName() proxy.PropString
	Chassis() proxy.PropString
	Deployment() proxy.PropString
	Location() proxy.PropString
	KernelName() proxy.PropString
	KernelRelease() proxy.PropString
	KernelVersion() proxy.PropString
	OperatingSystemPrettyName() proxy.PropString
	OperatingSystemCPEName() proxy.PropString
	HomeURL() proxy.PropString
}

type interfaceHostname struct{}

func (v *interfaceHostname) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceHostname) GetInterfaceName_() string {
	return "org.freedesktop.hostname1"
}

// method SetHostname

func (v *interfaceHostname) GoSetHostname(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetHostname", flags, ch, arg0, arg1)
}

func (v *interfaceHostname) SetHostname(flags dbus.Flags, arg0 string, arg1 bool) error {
	return (<-v.GoSetHostname(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetStaticHostname

func (v *interfaceHostname) GoSetStaticHostname(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetStaticHostname", flags, ch, arg0, arg1)
}

func (v *interfaceHostname) SetStaticHostname(flags dbus.Flags, arg0 string, arg1 bool) error {
	return (<-v.GoSetStaticHostname(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetPrettyHostname

func (v *interfaceHostname) GoSetPrettyHostname(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetPrettyHostname", flags, ch, arg0, arg1)
}

func (v *interfaceHostname) SetPrettyHostname(flags dbus.Flags, arg0 string, arg1 bool) error {
	return (<-v.GoSetPrettyHostname(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetIconName

func (v *interfaceHostname) GoSetIconName(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetIconName", flags, ch, arg0, arg1)
}

func (v *interfaceHostname) SetIconName(flags dbus.Flags, arg0 string, arg1 bool) error {
	return (<-v.GoSetIconName(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetChassis

func (v *interfaceHostname) GoSetChassis(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetChassis", flags, ch, arg0, arg1)
}

func (v *interfaceHostname) SetChassis(flags dbus.Flags, arg0 string, arg1 bool) error {
	return (<-v.GoSetChassis(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetDeployment

func (v *interfaceHostname) GoSetDeployment(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDeployment", flags, ch, arg0, arg1)
}

func (v *interfaceHostname) SetDeployment(flags dbus.Flags, arg0 string, arg1 bool) error {
	return (<-v.GoSetDeployment(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetLocation

func (v *interfaceHostname) GoSetLocation(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLocation", flags, ch, arg0, arg1)
}

func (v *interfaceHostname) SetLocation(flags dbus.Flags, arg0 string, arg1 bool) error {
	return (<-v.GoSetLocation(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method GetProductUUID

func (v *interfaceHostname) GoGetProductUUID(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetProductUUID", flags, ch, arg0)
}

func (*interfaceHostname) StoreGetProductUUID(call *dbus.Call) (arg1 []uint8, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceHostname) GetProductUUID(flags dbus.Flags, arg0 bool) ([]uint8, error) {
	return v.StoreGetProductUUID(
		<-v.GoGetProductUUID(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// property Hostname s

func (v *interfaceHostname) Hostname() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Hostname",
	}
}

// property StaticHostname s

func (v *interfaceHostname) StaticHostname() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "StaticHostname",
	}
}

// property PrettyHostname s

func (v *interfaceHostname) PrettyHostname() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "PrettyHostname",
	}
}

// property IconName s

func (v *interfaceHostname) IconName() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "IconName",
	}
}

// property Chassis s

func (v *interfaceHostname) Chassis() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Chassis",
	}
}

// property Deployment s

func (v *interfaceHostname) Deployment() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Deployment",
	}
}

// property Location s

func (v *interfaceHostname) Location() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Location",
	}
}

// property KernelName s

func (v *interfaceHostname) KernelName() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "KernelName",
	}
}

// property KernelRelease s

func (v *interfaceHostname) KernelRelease() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "KernelRelease",
	}
}

// property KernelVersion s

func (v *interfaceHostname) KernelVersion() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "KernelVersion",
	}
}

// property OperatingSystemPrettyName s

func (v *interfaceHostname) OperatingSystemPrettyName() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "OperatingSystemPrettyName",
	}
}

// property OperatingSystemCPEName s

func (v *interfaceHostname) OperatingSystemCPEName() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "OperatingSystemCPEName",
	}
}

// property HomeURL s

func (v *interfaceHostname) HomeURL() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "HomeURL",
	}
}
