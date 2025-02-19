// Code generated by "./generator ./com.deepin.daemon.timedate"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package timedate

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type Timedate interface {
	timedate // interface com.deepin.daemon.Timedate
	proxy.Object
}

type objectTimedate struct {
	interfaceTimedate // interface com.deepin.daemon.Timedate
	proxy.ImplObject
}

func NewTimedate(conn *dbus.Conn) Timedate {
	obj := new(objectTimedate)
	obj.ImplObject.Init_(conn, "com.deepin.daemon.Timedate", "/com/deepin/daemon/Timedate")
	return obj
}

type timedate interface {
	ConnectTimeUpdate(cb func()) (dbusutil.SignalHandlerId, error)
}

type interfaceTimedate struct{}

func (v *interfaceTimedate) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceTimedate) GetInterfaceName_() string {
	return "com.deepin.daemon.Timedate"
}

// signal TimeUpdate

func (v *interfaceTimedate) ConnectTimeUpdate(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "TimeUpdate", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".TimeUpdate",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
