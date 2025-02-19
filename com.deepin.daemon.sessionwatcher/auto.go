// Code generated by "./generator ./com.deepin.daemon.sessionwatcher"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package sessionwatcher

import (
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type SessionWatcher interface {
	sessionWatcher // interface com.deepin.daemon.SessionWatcher
	proxy.Object
}

type objectSessionWatcher struct {
	interfaceSessionWatcher // interface com.deepin.daemon.SessionWatcher
	proxy.ImplObject
}

func NewSessionWatcher(conn *dbus.Conn) SessionWatcher {
	obj := new(objectSessionWatcher)
	obj.ImplObject.Init_(conn, "com.deepin.daemon.SessionWatcher", "/com/deepin/daemon/SessionWatcher")
	return obj
}

type sessionWatcher interface {
	GoGetSessions(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetSessions(flags dbus.Flags) ([]dbus.ObjectPath, error)
	GoIsX11SessionActive(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	IsX11SessionActive(flags dbus.Flags) (bool, error)
	IsActive() proxy.PropBool
}

type interfaceSessionWatcher struct{}

func (v *interfaceSessionWatcher) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceSessionWatcher) GetInterfaceName_() string {
	return "com.deepin.daemon.SessionWatcher"
}

// method GetSessions

func (v *interfaceSessionWatcher) GoGetSessions(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSessions", flags, ch)
}

func (*interfaceSessionWatcher) StoreGetSessions(call *dbus.Call) (sessions []dbus.ObjectPath, err error) {
	err = call.Store(&sessions)
	return
}

func (v *interfaceSessionWatcher) GetSessions(flags dbus.Flags) ([]dbus.ObjectPath, error) {
	return v.StoreGetSessions(
		<-v.GoGetSessions(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsX11SessionActive

func (v *interfaceSessionWatcher) GoIsX11SessionActive(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsX11SessionActive", flags, ch)
}

func (*interfaceSessionWatcher) StoreIsX11SessionActive(call *dbus.Call) (is_active bool, err error) {
	err = call.Store(&is_active)
	return
}

func (v *interfaceSessionWatcher) IsX11SessionActive(flags dbus.Flags) (bool, error) {
	return v.StoreIsX11SessionActive(
		<-v.GoIsX11SessionActive(flags, make(chan *dbus.Call, 1)).Done)
}

// property IsActive b

func (v *interfaceSessionWatcher) IsActive() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "IsActive",
	}
}
