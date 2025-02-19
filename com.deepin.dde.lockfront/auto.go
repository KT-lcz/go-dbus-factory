// Code generated by "./generator ./com.deepin.dde.lockfront"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package lockfront

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type LockFront interface {
	lockfront // interface com.deepin.dde.lockFront
	proxy.Object
}

type objectLockFront struct {
	interfaceLockfront // interface com.deepin.dde.lockFront
	proxy.ImplObject
}

func NewLockFront(conn *dbus.Conn) LockFront {
	obj := new(objectLockFront)
	obj.ImplObject.Init_(conn, "com.deepin.dde.lockFront", "/com/deepin/dde/lockFront")
	return obj
}

type lockfront interface {
	GoShow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Show(flags dbus.Flags) error
	GoShowUserList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ShowUserList(flags dbus.Flags) error
	GoShowAuth(flags dbus.Flags, ch chan *dbus.Call, active bool) *dbus.Call
	ShowAuth(flags dbus.Flags, active bool) error
	GoSuspend(flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call
	Suspend(flags dbus.Flags, enable bool) error
	GoHibernate(flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call
	Hibernate(flags dbus.Flags, enable bool) error
	ConnectChangKey(cb func(keyEvent string)) (dbusutil.SignalHandlerId, error)
}

type interfaceLockfront struct{}

func (v *interfaceLockfront) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceLockfront) GetInterfaceName_() string {
	return "com.deepin.dde.lockFront"
}

// method Show

func (v *interfaceLockfront) GoShow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Show", flags, ch)
}

func (v *interfaceLockfront) Show(flags dbus.Flags) error {
	return (<-v.GoShow(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ShowUserList

func (v *interfaceLockfront) GoShowUserList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ShowUserList", flags, ch)
}

func (v *interfaceLockfront) ShowUserList(flags dbus.Flags) error {
	return (<-v.GoShowUserList(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ShowAuth

func (v *interfaceLockfront) GoShowAuth(flags dbus.Flags, ch chan *dbus.Call, active bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ShowAuth", flags, ch, active)
}

func (v *interfaceLockfront) ShowAuth(flags dbus.Flags, active bool) error {
	return (<-v.GoShowAuth(flags, make(chan *dbus.Call, 1), active).Done).Err
}

// method Suspend

func (v *interfaceLockfront) GoSuspend(flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Suspend", flags, ch, enable)
}

func (v *interfaceLockfront) Suspend(flags dbus.Flags, enable bool) error {
	return (<-v.GoSuspend(flags, make(chan *dbus.Call, 1), enable).Done).Err
}

// method Hibernate

func (v *interfaceLockfront) GoHibernate(flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Hibernate", flags, ch, enable)
}

func (v *interfaceLockfront) Hibernate(flags dbus.Flags, enable bool) error {
	return (<-v.GoHibernate(flags, make(chan *dbus.Call, 1), enable).Done).Err
}

// signal ChangKey

func (v *interfaceLockfront) ConnectChangKey(cb func(keyEvent string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ChangKey", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ChangKey",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var keyEvent string
		err := dbus.Store(sig.Body, &keyEvent)
		if err == nil {
			cb(keyEvent)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
