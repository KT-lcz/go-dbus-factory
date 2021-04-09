// Code generated by "./generator ./org.freedesktop.dbus"; DO NOT EDIT.

package dbus

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type DBus interface {
	dbusIfc // interface org.freedesktop.DBus
	proxy.Object
}

type objectDBus struct {
	interfaceDbusIfc // interface org.freedesktop.DBus
	proxy.ImplObject
}

func NewDBus(conn *dbus.Conn) DBus {
	obj := new(objectDBus)
	obj.ImplObject.Init_(conn, "org.freedesktop.DBus", "/org/freedesktop/DBus")
	return obj
}

type dbusIfc interface {
	GoHello(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Hello(flags dbus.Flags) (string, error)
	GoRequestName(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32) *dbus.Call
	RequestName(flags dbus.Flags, arg0 string, arg1 uint32) (uint32, error)
	GoReleaseName(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	ReleaseName(flags dbus.Flags, arg0 string) (uint32, error)
	GoStartServiceByName(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32) *dbus.Call
	StartServiceByName(flags dbus.Flags, arg0 string, arg1 uint32) (uint32, error)
	GoUpdateActivationEnvironment(flags dbus.Flags, ch chan *dbus.Call, arg0 map[string]string) *dbus.Call
	UpdateActivationEnvironment(flags dbus.Flags, arg0 map[string]string) error
	GoNameHasOwner(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	NameHasOwner(flags dbus.Flags, arg0 string) (bool, error)
	GoListNames(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ListNames(flags dbus.Flags) ([]string, error)
	GoListActivatableNames(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ListActivatableNames(flags dbus.Flags) ([]string, error)
	GoAddMatch(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	AddMatch(flags dbus.Flags, arg0 string) error
	GoRemoveMatch(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	RemoveMatch(flags dbus.Flags, arg0 string) error
	GoGetNameOwner(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	GetNameOwner(flags dbus.Flags, arg0 string) (string, error)
	GoListQueuedOwners(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	ListQueuedOwners(flags dbus.Flags, arg0 string) ([]string, error)
	GoGetConnectionUnixUser(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	GetConnectionUnixUser(flags dbus.Flags, arg0 string) (uint32, error)
	GoGetConnectionUnixProcessID(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	GetConnectionUnixProcessID(flags dbus.Flags, arg0 string) (uint32, error)
	GoGetAdtAuditSessionData(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	GetAdtAuditSessionData(flags dbus.Flags, arg0 string) ([]uint8, error)
	GoGetConnectionSELinuxSecurityContext(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	GetConnectionSELinuxSecurityContext(flags dbus.Flags, arg0 string) ([]uint8, error)
	GoReloadConfig(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ReloadConfig(flags dbus.Flags) error
	GoGetId(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetId(flags dbus.Flags) (string, error)
	GoGetConnectionCredentials(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	GetConnectionCredentials(flags dbus.Flags, arg0 string) (map[string]dbus.Variant, error)
	ConnectNameOwnerChanged(cb func(arg0 string, arg1 string, arg2 string)) (dbusutil.SignalHandlerId, error)
	ConnectNameLost(cb func(arg0 string)) (dbusutil.SignalHandlerId, error)
	ConnectNameAcquired(cb func(arg0 string)) (dbusutil.SignalHandlerId, error)
}

type interfaceDbusIfc struct{}

func (v *interfaceDbusIfc) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceDbusIfc) GetInterfaceName_() string {
	return "org.freedesktop.DBus"
}

// method Hello

func (v *interfaceDbusIfc) GoHello(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Hello", flags, ch)
}

func (*interfaceDbusIfc) StoreHello(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceDbusIfc) Hello(flags dbus.Flags) (string, error) {
	return v.StoreHello(
		<-v.GoHello(flags, make(chan *dbus.Call, 1)).Done)
}

// method RequestName

func (v *interfaceDbusIfc) GoRequestName(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestName", flags, ch, arg0, arg1)
}

func (*interfaceDbusIfc) StoreRequestName(call *dbus.Call) (arg2 uint32, err error) {
	err = call.Store(&arg2)
	return
}

func (v *interfaceDbusIfc) RequestName(flags dbus.Flags, arg0 string, arg1 uint32) (uint32, error) {
	return v.StoreRequestName(
		<-v.GoRequestName(flags, make(chan *dbus.Call, 1), arg0, arg1).Done)
}

// method ReleaseName

func (v *interfaceDbusIfc) GoReleaseName(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReleaseName", flags, ch, arg0)
}

func (*interfaceDbusIfc) StoreReleaseName(call *dbus.Call) (arg1 uint32, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceDbusIfc) ReleaseName(flags dbus.Flags, arg0 string) (uint32, error) {
	return v.StoreReleaseName(
		<-v.GoReleaseName(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method StartServiceByName

func (v *interfaceDbusIfc) GoStartServiceByName(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StartServiceByName", flags, ch, arg0, arg1)
}

func (*interfaceDbusIfc) StoreStartServiceByName(call *dbus.Call) (arg2 uint32, err error) {
	err = call.Store(&arg2)
	return
}

func (v *interfaceDbusIfc) StartServiceByName(flags dbus.Flags, arg0 string, arg1 uint32) (uint32, error) {
	return v.StoreStartServiceByName(
		<-v.GoStartServiceByName(flags, make(chan *dbus.Call, 1), arg0, arg1).Done)
}

// method UpdateActivationEnvironment

func (v *interfaceDbusIfc) GoUpdateActivationEnvironment(flags dbus.Flags, ch chan *dbus.Call, arg0 map[string]string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UpdateActivationEnvironment", flags, ch, arg0)
}

func (v *interfaceDbusIfc) UpdateActivationEnvironment(flags dbus.Flags, arg0 map[string]string) error {
	return (<-v.GoUpdateActivationEnvironment(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method NameHasOwner

func (v *interfaceDbusIfc) GoNameHasOwner(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".NameHasOwner", flags, ch, arg0)
}

func (*interfaceDbusIfc) StoreNameHasOwner(call *dbus.Call) (arg1 bool, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceDbusIfc) NameHasOwner(flags dbus.Flags, arg0 string) (bool, error) {
	return v.StoreNameHasOwner(
		<-v.GoNameHasOwner(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method ListNames

func (v *interfaceDbusIfc) GoListNames(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListNames", flags, ch)
}

func (*interfaceDbusIfc) StoreListNames(call *dbus.Call) (arg0 []string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceDbusIfc) ListNames(flags dbus.Flags) ([]string, error) {
	return v.StoreListNames(
		<-v.GoListNames(flags, make(chan *dbus.Call, 1)).Done)
}

// method ListActivatableNames

func (v *interfaceDbusIfc) GoListActivatableNames(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListActivatableNames", flags, ch)
}

func (*interfaceDbusIfc) StoreListActivatableNames(call *dbus.Call) (arg0 []string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceDbusIfc) ListActivatableNames(flags dbus.Flags) ([]string, error) {
	return v.StoreListActivatableNames(
		<-v.GoListActivatableNames(flags, make(chan *dbus.Call, 1)).Done)
}

// method AddMatch

func (v *interfaceDbusIfc) GoAddMatch(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddMatch", flags, ch, arg0)
}

func (v *interfaceDbusIfc) AddMatch(flags dbus.Flags, arg0 string) error {
	return (<-v.GoAddMatch(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method RemoveMatch

func (v *interfaceDbusIfc) GoRemoveMatch(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveMatch", flags, ch, arg0)
}

func (v *interfaceDbusIfc) RemoveMatch(flags dbus.Flags, arg0 string) error {
	return (<-v.GoRemoveMatch(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method GetNameOwner

func (v *interfaceDbusIfc) GoGetNameOwner(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetNameOwner", flags, ch, arg0)
}

func (*interfaceDbusIfc) StoreGetNameOwner(call *dbus.Call) (arg1 string, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceDbusIfc) GetNameOwner(flags dbus.Flags, arg0 string) (string, error) {
	return v.StoreGetNameOwner(
		<-v.GoGetNameOwner(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method ListQueuedOwners

func (v *interfaceDbusIfc) GoListQueuedOwners(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListQueuedOwners", flags, ch, arg0)
}

func (*interfaceDbusIfc) StoreListQueuedOwners(call *dbus.Call) (arg1 []string, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceDbusIfc) ListQueuedOwners(flags dbus.Flags, arg0 string) ([]string, error) {
	return v.StoreListQueuedOwners(
		<-v.GoListQueuedOwners(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method GetConnectionUnixUser

func (v *interfaceDbusIfc) GoGetConnectionUnixUser(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetConnectionUnixUser", flags, ch, arg0)
}

func (*interfaceDbusIfc) StoreGetConnectionUnixUser(call *dbus.Call) (arg1 uint32, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceDbusIfc) GetConnectionUnixUser(flags dbus.Flags, arg0 string) (uint32, error) {
	return v.StoreGetConnectionUnixUser(
		<-v.GoGetConnectionUnixUser(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method GetConnectionUnixProcessID

func (v *interfaceDbusIfc) GoGetConnectionUnixProcessID(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetConnectionUnixProcessID", flags, ch, arg0)
}

func (*interfaceDbusIfc) StoreGetConnectionUnixProcessID(call *dbus.Call) (arg1 uint32, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceDbusIfc) GetConnectionUnixProcessID(flags dbus.Flags, arg0 string) (uint32, error) {
	return v.StoreGetConnectionUnixProcessID(
		<-v.GoGetConnectionUnixProcessID(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method GetAdtAuditSessionData

func (v *interfaceDbusIfc) GoGetAdtAuditSessionData(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAdtAuditSessionData", flags, ch, arg0)
}

func (*interfaceDbusIfc) StoreGetAdtAuditSessionData(call *dbus.Call) (arg1 []uint8, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceDbusIfc) GetAdtAuditSessionData(flags dbus.Flags, arg0 string) ([]uint8, error) {
	return v.StoreGetAdtAuditSessionData(
		<-v.GoGetAdtAuditSessionData(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method GetConnectionSELinuxSecurityContext

func (v *interfaceDbusIfc) GoGetConnectionSELinuxSecurityContext(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetConnectionSELinuxSecurityContext", flags, ch, arg0)
}

func (*interfaceDbusIfc) StoreGetConnectionSELinuxSecurityContext(call *dbus.Call) (arg1 []uint8, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceDbusIfc) GetConnectionSELinuxSecurityContext(flags dbus.Flags, arg0 string) ([]uint8, error) {
	return v.StoreGetConnectionSELinuxSecurityContext(
		<-v.GoGetConnectionSELinuxSecurityContext(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method ReloadConfig

func (v *interfaceDbusIfc) GoReloadConfig(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReloadConfig", flags, ch)
}

func (v *interfaceDbusIfc) ReloadConfig(flags dbus.Flags) error {
	return (<-v.GoReloadConfig(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method GetId

func (v *interfaceDbusIfc) GoGetId(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetId", flags, ch)
}

func (*interfaceDbusIfc) StoreGetId(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceDbusIfc) GetId(flags dbus.Flags) (string, error) {
	return v.StoreGetId(
		<-v.GoGetId(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetConnectionCredentials

func (v *interfaceDbusIfc) GoGetConnectionCredentials(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetConnectionCredentials", flags, ch, arg0)
}

func (*interfaceDbusIfc) StoreGetConnectionCredentials(call *dbus.Call) (arg1 map[string]dbus.Variant, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceDbusIfc) GetConnectionCredentials(flags dbus.Flags, arg0 string) (map[string]dbus.Variant, error) {
	return v.StoreGetConnectionCredentials(
		<-v.GoGetConnectionCredentials(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// signal NameOwnerChanged

func (v *interfaceDbusIfc) ConnectNameOwnerChanged(cb func(arg0 string, arg1 string, arg2 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NameOwnerChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NameOwnerChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		var arg1 string
		var arg2 string
		err := dbus.Store(sig.Body, &arg0, &arg1, &arg2)
		if err == nil {
			cb(arg0, arg1, arg2)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal NameLost

func (v *interfaceDbusIfc) ConnectNameLost(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NameLost", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NameLost",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		err := dbus.Store(sig.Body, &arg0)
		if err == nil {
			cb(arg0)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal NameAcquired

func (v *interfaceDbusIfc) ConnectNameAcquired(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NameAcquired", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NameAcquired",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		err := dbus.Store(sig.Body, &arg0)
		if err == nil {
			cb(arg0)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
