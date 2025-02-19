// Code generated by "./generator ./com.deepin.dde.launcher"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package launcher

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type Launcher interface {
	launcher // interface com.deepin.dde.Launcher
	proxy.Object
}

type objectLauncher struct {
	interfaceLauncher // interface com.deepin.dde.Launcher
	proxy.ImplObject
}

func NewLauncher(conn *dbus.Conn) Launcher {
	obj := new(objectLauncher)
	obj.ImplObject.Init_(conn, "com.deepin.dde.Launcher", "/com/deepin/dde/Launcher")
	return obj
}

type launcher interface {
	GoExit(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Exit(flags dbus.Flags) error
	GoHide(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Hide(flags dbus.Flags) error
	GoShow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Show(flags dbus.Flags) error
	GoShowByMode(flags dbus.Flags, ch chan *dbus.Call, in0 int64) *dbus.Call
	ShowByMode(flags dbus.Flags, in0 int64) error
	GoUninstallApp(flags dbus.Flags, ch chan *dbus.Call, appKey string) *dbus.Call
	UninstallApp(flags dbus.Flags, appKey string) error
	GoToggle(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Toggle(flags dbus.Flags) error
	GoIsVisible(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	IsVisible(flags dbus.Flags) (bool, error)
	ConnectClosed(cb func()) (dbusutil.SignalHandlerId, error)
	ConnectShown(cb func()) (dbusutil.SignalHandlerId, error)
	ConnectVisibleChanged(cb func(visible bool)) (dbusutil.SignalHandlerId, error)
	Visible() proxy.PropBool
}

type interfaceLauncher struct{}

func (v *interfaceLauncher) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceLauncher) GetInterfaceName_() string {
	return "com.deepin.dde.Launcher"
}

// method Exit

func (v *interfaceLauncher) GoExit(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Exit", flags, ch)
}

func (v *interfaceLauncher) Exit(flags dbus.Flags) error {
	return (<-v.GoExit(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Hide

func (v *interfaceLauncher) GoHide(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Hide", flags, ch)
}

func (v *interfaceLauncher) Hide(flags dbus.Flags) error {
	return (<-v.GoHide(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Show

func (v *interfaceLauncher) GoShow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Show", flags, ch)
}

func (v *interfaceLauncher) Show(flags dbus.Flags) error {
	return (<-v.GoShow(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ShowByMode

func (v *interfaceLauncher) GoShowByMode(flags dbus.Flags, ch chan *dbus.Call, in0 int64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ShowByMode", flags, ch, in0)
}

func (v *interfaceLauncher) ShowByMode(flags dbus.Flags, in0 int64) error {
	return (<-v.GoShowByMode(flags, make(chan *dbus.Call, 1), in0).Done).Err
}

// method UninstallApp

func (v *interfaceLauncher) GoUninstallApp(flags dbus.Flags, ch chan *dbus.Call, appKey string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UninstallApp", flags, ch, appKey)
}

func (v *interfaceLauncher) UninstallApp(flags dbus.Flags, appKey string) error {
	return (<-v.GoUninstallApp(flags, make(chan *dbus.Call, 1), appKey).Done).Err
}

// method Toggle

func (v *interfaceLauncher) GoToggle(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Toggle", flags, ch)
}

func (v *interfaceLauncher) Toggle(flags dbus.Flags) error {
	return (<-v.GoToggle(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method IsVisible

func (v *interfaceLauncher) GoIsVisible(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsVisible", flags, ch)
}

func (*interfaceLauncher) StoreIsVisible(call *dbus.Call) (arg0 bool, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceLauncher) IsVisible(flags dbus.Flags) (bool, error) {
	return v.StoreIsVisible(
		<-v.GoIsVisible(flags, make(chan *dbus.Call, 1)).Done)
}

// signal Closed

func (v *interfaceLauncher) ConnectClosed(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Closed", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Closed",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Shown

func (v *interfaceLauncher) ConnectShown(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Shown", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Shown",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal VisibleChanged

func (v *interfaceLauncher) ConnectVisibleChanged(cb func(visible bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "VisibleChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".VisibleChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var visible bool
		err := dbus.Store(sig.Body, &visible)
		if err == nil {
			cb(visible)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Visible b

func (v *interfaceLauncher) Visible() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "Visible",
	}
}
