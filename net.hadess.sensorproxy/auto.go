// Code generated by "./generator ./net.hadess.sensorproxy"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package sensorproxy

import (
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type SensorProxy interface {
	sensorProxy // interface net.hadess.SensorProxy
	proxy.Object
}

type objectSensorProxy struct {
	interfaceSensorProxy // interface net.hadess.SensorProxy
	proxy.ImplObject
}

func NewSensorProxy(conn *dbus.Conn) SensorProxy {
	obj := new(objectSensorProxy)
	obj.ImplObject.Init_(conn, "net.hadess.SensorProxy", "/net/hadess/SensorProxy")
	return obj
}

type sensorProxy interface {
	GoClaimAccelerometer(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ClaimAccelerometer(flags dbus.Flags) error
	GoReleaseAccelerometer(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ReleaseAccelerometer(flags dbus.Flags) error
	GoClaimLight(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ClaimLight(flags dbus.Flags) error
	GoReleaseLight(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ReleaseLight(flags dbus.Flags) error
	HasAccelerometer() proxy.PropBool
	AccelerometerOrientation() proxy.PropString
	HasAmbientLight() proxy.PropBool
	LightLevelUnit() proxy.PropString
	LightLevel() proxy.PropDouble
}

type interfaceSensorProxy struct{}

func (v *interfaceSensorProxy) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceSensorProxy) GetInterfaceName_() string {
	return "net.hadess.SensorProxy"
}

// method ClaimAccelerometer

func (v *interfaceSensorProxy) GoClaimAccelerometer(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ClaimAccelerometer", flags, ch)
}

func (v *interfaceSensorProxy) ClaimAccelerometer(flags dbus.Flags) error {
	return (<-v.GoClaimAccelerometer(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ReleaseAccelerometer

func (v *interfaceSensorProxy) GoReleaseAccelerometer(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReleaseAccelerometer", flags, ch)
}

func (v *interfaceSensorProxy) ReleaseAccelerometer(flags dbus.Flags) error {
	return (<-v.GoReleaseAccelerometer(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ClaimLight

func (v *interfaceSensorProxy) GoClaimLight(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ClaimLight", flags, ch)
}

func (v *interfaceSensorProxy) ClaimLight(flags dbus.Flags) error {
	return (<-v.GoClaimLight(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ReleaseLight

func (v *interfaceSensorProxy) GoReleaseLight(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReleaseLight", flags, ch)
}

func (v *interfaceSensorProxy) ReleaseLight(flags dbus.Flags) error {
	return (<-v.GoReleaseLight(flags, make(chan *dbus.Call, 1)).Done).Err
}

// property HasAccelerometer b

func (v *interfaceSensorProxy) HasAccelerometer() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "HasAccelerometer",
	}
}

// property AccelerometerOrientation s

func (v *interfaceSensorProxy) AccelerometerOrientation() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "AccelerometerOrientation",
	}
}

// property HasAmbientLight b

func (v *interfaceSensorProxy) HasAmbientLight() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "HasAmbientLight",
	}
}

// property LightLevelUnit s

func (v *interfaceSensorProxy) LightLevelUnit() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "LightLevelUnit",
	}
}

// property LightLevel d

func (v *interfaceSensorProxy) LightLevel() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "LightLevel",
	}
}
