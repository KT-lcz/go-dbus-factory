// Code generated by "./generator ./com.deepin.daemon.imageeffect"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package imageeffect

import (
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type ImageEffect interface {
	imageEffect // interface com.deepin.daemon.ImageEffect
	proxy.Object
}

type objectImageEffect struct {
	interfaceImageEffect // interface com.deepin.daemon.ImageEffect
	proxy.ImplObject
}

func NewImageEffect(conn *dbus.Conn) ImageEffect {
	obj := new(objectImageEffect)
	obj.ImplObject.Init_(conn, "com.deepin.daemon.ImageEffect", "/com/deepin/daemon/ImageEffect")
	return obj
}

type imageEffect interface {
	GoDelete(flags dbus.Flags, ch chan *dbus.Call, effect string, filename string) *dbus.Call
	Delete(flags dbus.Flags, effect string, filename string) error
	GoGet(flags dbus.Flags, ch chan *dbus.Call, effect string, filename string) *dbus.Call
	Get(flags dbus.Flags, effect string, filename string) (string, error)
}

type interfaceImageEffect struct{}

func (v *interfaceImageEffect) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceImageEffect) GetInterfaceName_() string {
	return "com.deepin.daemon.ImageEffect"
}

// method Delete

func (v *interfaceImageEffect) GoDelete(flags dbus.Flags, ch chan *dbus.Call, effect string, filename string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Delete", flags, ch, effect, filename)
}

func (v *interfaceImageEffect) Delete(flags dbus.Flags, effect string, filename string) error {
	return (<-v.GoDelete(flags, make(chan *dbus.Call, 1), effect, filename).Done).Err
}

// method Get

func (v *interfaceImageEffect) GoGet(flags dbus.Flags, ch chan *dbus.Call, effect string, filename string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Get", flags, ch, effect, filename)
}

func (*interfaceImageEffect) StoreGet(call *dbus.Call) (outputFile string, err error) {
	err = call.Store(&outputFile)
	return
}

func (v *interfaceImageEffect) Get(flags dbus.Flags, effect string, filename string) (string, error) {
	return v.StoreGet(
		<-v.GoGet(flags, make(chan *dbus.Call, 1), effect, filename).Done)
}
