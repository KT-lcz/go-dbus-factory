// Code generated by "./generator ./com.deepin.system.ipwatchd"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package ipwatchd

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type IPWatchD interface {
	ipwatchd // interface com.deepin.system.IPWatchD
	proxy.Object
}

type objectIPWatchD struct {
	interfaceIpwatchd // interface com.deepin.system.IPWatchD
	proxy.ImplObject
}

func NewIPWatchD(conn *dbus.Conn) IPWatchD {
	obj := new(objectIPWatchD)
	obj.ImplObject.Init_(conn, "com.deepin.system.IPWatchD", "/com/deepin/system/IPWatchD")
	return obj
}

type ipwatchd interface {
	GoRequestIPConflictCheck(flags dbus.Flags, ch chan *dbus.Call, ip string, ifc string) *dbus.Call
	RequestIPConflictCheck(flags dbus.Flags, ip string, ifc string) (string, error)
	ConnectIPConflict(cb func(ip string, smac string, dmac string)) (dbusutil.SignalHandlerId, error)
}

type interfaceIpwatchd struct{}

func (v *interfaceIpwatchd) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceIpwatchd) GetInterfaceName_() string {
	return "com.deepin.system.IPWatchD"
}

// method RequestIPConflictCheck

func (v *interfaceIpwatchd) GoRequestIPConflictCheck(flags dbus.Flags, ch chan *dbus.Call, ip string, ifc string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestIPConflictCheck", flags, ch, ip, ifc)
}

func (*interfaceIpwatchd) StoreRequestIPConflictCheck(call *dbus.Call) (mac string, err error) {
	err = call.Store(&mac)
	return
}

func (v *interfaceIpwatchd) RequestIPConflictCheck(flags dbus.Flags, ip string, ifc string) (string, error) {
	return v.StoreRequestIPConflictCheck(
		<-v.GoRequestIPConflictCheck(flags, make(chan *dbus.Call, 1), ip, ifc).Done)
}

// signal IPConflict

func (v *interfaceIpwatchd) ConnectIPConflict(cb func(ip string, smac string, dmac string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "IPConflict", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".IPConflict",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var ip string
		var smac string
		var dmac string
		err := dbus.Store(sig.Body, &ip, &smac, &dmac)
		if err == nil {
			cb(ip, smac, dmac)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
