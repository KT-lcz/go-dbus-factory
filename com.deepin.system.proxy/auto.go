// Code generated by "./generator ./com.deepin.system.proxy"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package proxy

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type App interface {
	app // interface com.deepin.system.proxy.App
	proxy.Object
}

type objectApp struct {
	interfaceApp // interface com.deepin.system.proxy.App
	proxy.ImplObject
}

func NewApp(conn *dbus.Conn) App {
	obj := new(objectApp)
	obj.ImplObject.Init_(conn, "com.deepin.system.proxy", "/com/deepin/system/proxy/App")
	return obj
}

type app interface {
	GoAddProxy(flags dbus.Flags, ch chan *dbus.Call, proto string, name string, proxy []uint8) *dbus.Call
	AddProxy(flags dbus.Flags, proto string, name string, proxy []uint8) error
	GoAddProxyApps(flags dbus.Flags, ch chan *dbus.Call, app []string) *dbus.Call
	AddProxyApps(flags dbus.Flags, app []string) error
	GoClearProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ClearProxy(flags dbus.Flags) error
	GoDelProxyApps(flags dbus.Flags, ch chan *dbus.Call, app []string) *dbus.Call
	DelProxyApps(flags dbus.Flags, app []string) error
	GoGetProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetProxy(flags dbus.Flags) (string, error)
	GoSetProxies(flags dbus.Flags, ch chan *dbus.Call, proxies []interface{}) *dbus.Call
	SetProxies(flags dbus.Flags, proxies []interface{}) error
	GoGetCGroups(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetCGroups(flags dbus.Flags) (string, error)
	GoAddProc(flags dbus.Flags, ch chan *dbus.Call, pid int32) *dbus.Call
	AddProc(flags dbus.Flags, pid int32) error
	GoStartProxy(flags dbus.Flags, ch chan *dbus.Call, proto string, name string, udp bool) *dbus.Call
	StartProxy(flags dbus.Flags, proto string, name string, udp bool) error
	GoStopProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	StopProxy(flags dbus.Flags) error
	ConnectProxy(cb func(proxy []interface{})) (dbusutil.SignalHandlerId, error)
}

type interfaceApp struct{}

func (v *interfaceApp) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceApp) GetInterfaceName_() string {
	return "com.deepin.system.proxy.App"
}

// method AddProxy

func (v *interfaceApp) GoAddProxy(flags dbus.Flags, ch chan *dbus.Call, proto string, name string, proxy []uint8) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddProxy", flags, ch, proto, name, proxy)
}

func (v *interfaceApp) AddProxy(flags dbus.Flags, proto string, name string, proxy []uint8) error {
	return (<-v.GoAddProxy(flags, make(chan *dbus.Call, 1), proto, name, proxy).Done).Err
}

// method AddProxyApps

func (v *interfaceApp) GoAddProxyApps(flags dbus.Flags, ch chan *dbus.Call, app []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddProxyApps", flags, ch, app)
}

func (v *interfaceApp) AddProxyApps(flags dbus.Flags, app []string) error {
	return (<-v.GoAddProxyApps(flags, make(chan *dbus.Call, 1), app).Done).Err
}

// method ClearProxy

func (v *interfaceApp) GoClearProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ClearProxy", flags, ch)
}

func (v *interfaceApp) ClearProxy(flags dbus.Flags) error {
	return (<-v.GoClearProxy(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method DelProxyApps

func (v *interfaceApp) GoDelProxyApps(flags dbus.Flags, ch chan *dbus.Call, app []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DelProxyApps", flags, ch, app)
}

func (v *interfaceApp) DelProxyApps(flags dbus.Flags, app []string) error {
	return (<-v.GoDelProxyApps(flags, make(chan *dbus.Call, 1), app).Done).Err
}

// method GetProxy

func (v *interfaceApp) GoGetProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetProxy", flags, ch)
}

func (*interfaceApp) StoreGetProxy(call *dbus.Call) (proxy string, err error) {
	err = call.Store(&proxy)
	return
}

func (v *interfaceApp) GetProxy(flags dbus.Flags) (string, error) {
	return v.StoreGetProxy(
		<-v.GoGetProxy(flags, make(chan *dbus.Call, 1)).Done)
}

// method SetProxies

func (v *interfaceApp) GoSetProxies(flags dbus.Flags, ch chan *dbus.Call, proxies []interface{}) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetProxies", flags, ch, proxies)
}

func (v *interfaceApp) SetProxies(flags dbus.Flags, proxies []interface{}) error {
	return (<-v.GoSetProxies(flags, make(chan *dbus.Call, 1), proxies).Done).Err
}

// method GetCGroups

func (v *interfaceApp) GoGetCGroups(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetCGroups", flags, ch)
}

func (*interfaceApp) StoreGetCGroups(call *dbus.Call) (cgroups string, err error) {
	err = call.Store(&cgroups)
	return
}

func (v *interfaceApp) GetCGroups(flags dbus.Flags) (string, error) {
	return v.StoreGetCGroups(
		<-v.GoGetCGroups(flags, make(chan *dbus.Call, 1)).Done)
}

// method AddProc

func (v *interfaceApp) GoAddProc(flags dbus.Flags, ch chan *dbus.Call, pid int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddProc", flags, ch, pid)
}

func (v *interfaceApp) AddProc(flags dbus.Flags, pid int32) error {
	return (<-v.GoAddProc(flags, make(chan *dbus.Call, 1), pid).Done).Err
}

// method StartProxy

func (v *interfaceApp) GoStartProxy(flags dbus.Flags, ch chan *dbus.Call, proto string, name string, udp bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StartProxy", flags, ch, proto, name, udp)
}

func (v *interfaceApp) StartProxy(flags dbus.Flags, proto string, name string, udp bool) error {
	return (<-v.GoStartProxy(flags, make(chan *dbus.Call, 1), proto, name, udp).Done).Err
}

// method StopProxy

func (v *interfaceApp) GoStopProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopProxy", flags, ch)
}

func (v *interfaceApp) StopProxy(flags dbus.Flags) error {
	return (<-v.GoStopProxy(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal Proxy

func (v *interfaceApp) ConnectProxy(cb func(proxy []interface{})) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Proxy", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Proxy",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var proxy []interface{}
		err := dbus.Store(sig.Body, &proxy)
		if err == nil {
			cb(proxy)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

type Global interface {
	global // interface com.deepin.system.proxy.Global
	proxy.Object
}

type objectGlobal struct {
	interfaceGlobal // interface com.deepin.system.proxy.Global
	proxy.ImplObject
}

func NewGlobal(conn *dbus.Conn) Global {
	obj := new(objectGlobal)
	obj.ImplObject.Init_(conn, "com.deepin.system.proxy", "/com/deepin/system/proxy/Global")
	return obj
}

type global interface {
	GoAddProxy(flags dbus.Flags, ch chan *dbus.Call, proto string, name string, proxy []uint8) *dbus.Call
	AddProxy(flags dbus.Flags, proto string, name string, proxy []uint8) error
	GoClearProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ClearProxy(flags dbus.Flags) error
	GoGetProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetProxy(flags dbus.Flags) (string, error)
	GoIgnoreProxyApps(flags dbus.Flags, ch chan *dbus.Call, app []string) *dbus.Call
	IgnoreProxyApps(flags dbus.Flags, app []string) error
	GoSetProxies(flags dbus.Flags, ch chan *dbus.Call, proxies []interface{}) *dbus.Call
	SetProxies(flags dbus.Flags, proxies []interface{}) error
	GoStartProxy(flags dbus.Flags, ch chan *dbus.Call, proto string, name string, udp bool) *dbus.Call
	StartProxy(flags dbus.Flags, proto string, name string, udp bool) error
	GoStopProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	StopProxy(flags dbus.Flags) error
	GoUnIgnoreProxyApps(flags dbus.Flags, ch chan *dbus.Call, app []string) *dbus.Call
	UnIgnoreProxyApps(flags dbus.Flags, app []string) error
	ConnectProxy(cb func(proxy []interface{})) (dbusutil.SignalHandlerId, error)
	IgnoreApp() proxy.PropStringArray
}

type interfaceGlobal struct{}

func (v *interfaceGlobal) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceGlobal) GetInterfaceName_() string {
	return "com.deepin.system.proxy.Global"
}

// method AddProxy

func (v *interfaceGlobal) GoAddProxy(flags dbus.Flags, ch chan *dbus.Call, proto string, name string, proxy []uint8) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddProxy", flags, ch, proto, name, proxy)
}

func (v *interfaceGlobal) AddProxy(flags dbus.Flags, proto string, name string, proxy []uint8) error {
	return (<-v.GoAddProxy(flags, make(chan *dbus.Call, 1), proto, name, proxy).Done).Err
}

// method ClearProxy

func (v *interfaceGlobal) GoClearProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ClearProxy", flags, ch)
}

func (v *interfaceGlobal) ClearProxy(flags dbus.Flags) error {
	return (<-v.GoClearProxy(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method GetProxy

func (v *interfaceGlobal) GoGetProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetProxy", flags, ch)
}

func (*interfaceGlobal) StoreGetProxy(call *dbus.Call) (proxy string, err error) {
	err = call.Store(&proxy)
	return
}

func (v *interfaceGlobal) GetProxy(flags dbus.Flags) (string, error) {
	return v.StoreGetProxy(
		<-v.GoGetProxy(flags, make(chan *dbus.Call, 1)).Done)
}

// method IgnoreProxyApps

func (v *interfaceGlobal) GoIgnoreProxyApps(flags dbus.Flags, ch chan *dbus.Call, app []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IgnoreProxyApps", flags, ch, app)
}

func (v *interfaceGlobal) IgnoreProxyApps(flags dbus.Flags, app []string) error {
	return (<-v.GoIgnoreProxyApps(flags, make(chan *dbus.Call, 1), app).Done).Err
}

// method SetProxies

func (v *interfaceGlobal) GoSetProxies(flags dbus.Flags, ch chan *dbus.Call, proxies []interface{}) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetProxies", flags, ch, proxies)
}

func (v *interfaceGlobal) SetProxies(flags dbus.Flags, proxies []interface{}) error {
	return (<-v.GoSetProxies(flags, make(chan *dbus.Call, 1), proxies).Done).Err
}

// method StartProxy

func (v *interfaceGlobal) GoStartProxy(flags dbus.Flags, ch chan *dbus.Call, proto string, name string, udp bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StartProxy", flags, ch, proto, name, udp)
}

func (v *interfaceGlobal) StartProxy(flags dbus.Flags, proto string, name string, udp bool) error {
	return (<-v.GoStartProxy(flags, make(chan *dbus.Call, 1), proto, name, udp).Done).Err
}

// method StopProxy

func (v *interfaceGlobal) GoStopProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopProxy", flags, ch)
}

func (v *interfaceGlobal) StopProxy(flags dbus.Flags) error {
	return (<-v.GoStopProxy(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method UnIgnoreProxyApps

func (v *interfaceGlobal) GoUnIgnoreProxyApps(flags dbus.Flags, ch chan *dbus.Call, app []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UnIgnoreProxyApps", flags, ch, app)
}

func (v *interfaceGlobal) UnIgnoreProxyApps(flags dbus.Flags, app []string) error {
	return (<-v.GoUnIgnoreProxyApps(flags, make(chan *dbus.Call, 1), app).Done).Err
}

// signal Proxy

func (v *interfaceGlobal) ConnectProxy(cb func(proxy []interface{})) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Proxy", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Proxy",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var proxy []interface{}
		err := dbus.Store(sig.Body, &proxy)
		if err == nil {
			cb(proxy)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property IgnoreApp as

func (v *interfaceGlobal) IgnoreApp() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "IgnoreApp",
	}
}
