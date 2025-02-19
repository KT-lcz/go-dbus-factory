// Code generated by "./generator ./com.deepin.api.cursorhelper"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package cursorhelper

import (
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type CursorHelper interface {
	cursorHelper // interface com.deepin.api.CursorHelper
	proxy.Object
}

type objectCursorHelper struct {
	interfaceCursorHelper // interface com.deepin.api.CursorHelper
	proxy.ImplObject
}

func NewCursorHelper(conn *dbus.Conn) CursorHelper {
	obj := new(objectCursorHelper)
	obj.ImplObject.Init_(conn, "com.deepin.api.CursorHelper", "/com/deepin/api/CursorHelper")
	return obj
}

type cursorHelper interface {
	GoSet(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call
	Set(flags dbus.Flags, name string) error
}

type interfaceCursorHelper struct{}

func (v *interfaceCursorHelper) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceCursorHelper) GetInterfaceName_() string {
	return "com.deepin.api.CursorHelper"
}

// method Set

func (v *interfaceCursorHelper) GoSet(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Set", flags, ch, name)
}

func (v *interfaceCursorHelper) Set(flags dbus.Flags, name string) error {
	return (<-v.GoSet(flags, make(chan *dbus.Call, 1), name).Done).Err
}
