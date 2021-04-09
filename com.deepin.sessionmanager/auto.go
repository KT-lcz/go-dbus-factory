// Code generated by "./generator ./com.deepin.sessionmanager"; DO NOT EDIT.

package sessionmanager

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type StartManager interface {
	startManager // interface com.deepin.StartManager
	proxy.Object
}

type objectStartManager struct {
	interfaceStartManager // interface com.deepin.StartManager
	proxy.ImplObject
}

func NewStartManager(conn *dbus.Conn) StartManager {
	obj := new(objectStartManager)
	obj.ImplObject.Init_(conn, "com.deepin.SessionManager", "/com/deepin/StartManager")
	return obj
}

type startManager interface {
	GoAddAutostart(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	AddAutostart(flags dbus.Flags, arg0 string) (bool, error)
	GoAutostartList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	AutostartList(flags dbus.Flags) ([]string, error)
	GoDumpMemRecord(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	DumpMemRecord(flags dbus.Flags) (string, error)
	GoGetApps(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetApps(flags dbus.Flags) (map[uint32]string, error)
	GoIsAutostart(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	IsAutostart(flags dbus.Flags, arg0 string) (bool, error)
	GoIsMemSufficient(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	IsMemSufficient(flags dbus.Flags) (bool, error)
	GoLaunch(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	Launch(flags dbus.Flags, arg0 string) (bool, error)
	GoLaunchApp(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32, arg2 []string) *dbus.Call
	LaunchApp(flags dbus.Flags, arg0 string, arg1 uint32, arg2 []string) error
	GoLaunchAppAction(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string, arg2 uint32) *dbus.Call
	LaunchAppAction(flags dbus.Flags, arg0 string, arg1 string, arg2 uint32) error
	GoLaunchAppWithOptions(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32, arg2 []string, arg3 map[string]dbus.Variant) *dbus.Call
	LaunchAppWithOptions(flags dbus.Flags, arg0 string, arg1 uint32, arg2 []string, arg3 map[string]dbus.Variant) error
	GoLaunchWithTimestamp(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32) *dbus.Call
	LaunchWithTimestamp(flags dbus.Flags, arg0 string, arg1 uint32) (bool, error)
	GoRemoveAutostart(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	RemoveAutostart(flags dbus.Flags, arg0 string) (bool, error)
	GoRunCommand(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 []string) *dbus.Call
	RunCommand(flags dbus.Flags, arg0 string, arg1 []string) error
	GoTryAgain(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call
	TryAgain(flags dbus.Flags, arg0 bool) error
	ConnectAutostartChanged(cb func(arg0 string, arg1 string)) (dbusutil.SignalHandlerId, error)
	NeededMemory() proxy.PropUint64
}

type interfaceStartManager struct{}

func (v *interfaceStartManager) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceStartManager) GetInterfaceName_() string {
	return "com.deepin.StartManager"
}

// method AddAutostart

func (v *interfaceStartManager) GoAddAutostart(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddAutostart", flags, ch, arg0)
}

func (*interfaceStartManager) StoreAddAutostart(call *dbus.Call) (arg1 bool, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceStartManager) AddAutostart(flags dbus.Flags, arg0 string) (bool, error) {
	return v.StoreAddAutostart(
		<-v.GoAddAutostart(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method AutostartList

func (v *interfaceStartManager) GoAutostartList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AutostartList", flags, ch)
}

func (*interfaceStartManager) StoreAutostartList(call *dbus.Call) (arg0 []string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceStartManager) AutostartList(flags dbus.Flags) ([]string, error) {
	return v.StoreAutostartList(
		<-v.GoAutostartList(flags, make(chan *dbus.Call, 1)).Done)
}

// method DumpMemRecord

func (v *interfaceStartManager) GoDumpMemRecord(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DumpMemRecord", flags, ch)
}

func (*interfaceStartManager) StoreDumpMemRecord(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceStartManager) DumpMemRecord(flags dbus.Flags) (string, error) {
	return v.StoreDumpMemRecord(
		<-v.GoDumpMemRecord(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetApps

func (v *interfaceStartManager) GoGetApps(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetApps", flags, ch)
}

func (*interfaceStartManager) StoreGetApps(call *dbus.Call) (arg0 map[uint32]string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceStartManager) GetApps(flags dbus.Flags) (map[uint32]string, error) {
	return v.StoreGetApps(
		<-v.GoGetApps(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsAutostart

func (v *interfaceStartManager) GoIsAutostart(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsAutostart", flags, ch, arg0)
}

func (*interfaceStartManager) StoreIsAutostart(call *dbus.Call) (arg1 bool, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceStartManager) IsAutostart(flags dbus.Flags, arg0 string) (bool, error) {
	return v.StoreIsAutostart(
		<-v.GoIsAutostart(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method IsMemSufficient

func (v *interfaceStartManager) GoIsMemSufficient(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsMemSufficient", flags, ch)
}

func (*interfaceStartManager) StoreIsMemSufficient(call *dbus.Call) (arg0 bool, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceStartManager) IsMemSufficient(flags dbus.Flags) (bool, error) {
	return v.StoreIsMemSufficient(
		<-v.GoIsMemSufficient(flags, make(chan *dbus.Call, 1)).Done)
}

// method Launch

func (v *interfaceStartManager) GoLaunch(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Launch", flags, ch, arg0)
}

func (*interfaceStartManager) StoreLaunch(call *dbus.Call) (arg1 bool, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceStartManager) Launch(flags dbus.Flags, arg0 string) (bool, error) {
	return v.StoreLaunch(
		<-v.GoLaunch(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method LaunchApp

func (v *interfaceStartManager) GoLaunchApp(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32, arg2 []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LaunchApp", flags, ch, arg0, arg1, arg2)
}

func (v *interfaceStartManager) LaunchApp(flags dbus.Flags, arg0 string, arg1 uint32, arg2 []string) error {
	return (<-v.GoLaunchApp(flags, make(chan *dbus.Call, 1), arg0, arg1, arg2).Done).Err
}

// method LaunchAppAction

func (v *interfaceStartManager) GoLaunchAppAction(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string, arg2 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LaunchAppAction", flags, ch, arg0, arg1, arg2)
}

func (v *interfaceStartManager) LaunchAppAction(flags dbus.Flags, arg0 string, arg1 string, arg2 uint32) error {
	return (<-v.GoLaunchAppAction(flags, make(chan *dbus.Call, 1), arg0, arg1, arg2).Done).Err
}

// method LaunchAppWithOptions

func (v *interfaceStartManager) GoLaunchAppWithOptions(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32, arg2 []string, arg3 map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LaunchAppWithOptions", flags, ch, arg0, arg1, arg2, arg3)
}

func (v *interfaceStartManager) LaunchAppWithOptions(flags dbus.Flags, arg0 string, arg1 uint32, arg2 []string, arg3 map[string]dbus.Variant) error {
	return (<-v.GoLaunchAppWithOptions(flags, make(chan *dbus.Call, 1), arg0, arg1, arg2, arg3).Done).Err
}

// method LaunchWithTimestamp

func (v *interfaceStartManager) GoLaunchWithTimestamp(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LaunchWithTimestamp", flags, ch, arg0, arg1)
}

func (*interfaceStartManager) StoreLaunchWithTimestamp(call *dbus.Call) (arg2 bool, err error) {
	err = call.Store(&arg2)
	return
}

func (v *interfaceStartManager) LaunchWithTimestamp(flags dbus.Flags, arg0 string, arg1 uint32) (bool, error) {
	return v.StoreLaunchWithTimestamp(
		<-v.GoLaunchWithTimestamp(flags, make(chan *dbus.Call, 1), arg0, arg1).Done)
}

// method RemoveAutostart

func (v *interfaceStartManager) GoRemoveAutostart(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveAutostart", flags, ch, arg0)
}

func (*interfaceStartManager) StoreRemoveAutostart(call *dbus.Call) (arg1 bool, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceStartManager) RemoveAutostart(flags dbus.Flags, arg0 string) (bool, error) {
	return v.StoreRemoveAutostart(
		<-v.GoRemoveAutostart(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method RunCommand

func (v *interfaceStartManager) GoRunCommand(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RunCommand", flags, ch, arg0, arg1)
}

func (v *interfaceStartManager) RunCommand(flags dbus.Flags, arg0 string, arg1 []string) error {
	return (<-v.GoRunCommand(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method TryAgain

func (v *interfaceStartManager) GoTryAgain(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".TryAgain", flags, ch, arg0)
}

func (v *interfaceStartManager) TryAgain(flags dbus.Flags, arg0 bool) error {
	return (<-v.GoTryAgain(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// signal AutostartChanged

func (v *interfaceStartManager) ConnectAutostartChanged(cb func(arg0 string, arg1 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "AutostartChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".AutostartChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		var arg1 string
		err := dbus.Store(sig.Body, &arg0, &arg1)
		if err == nil {
			cb(arg0, arg1)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property NeededMemory t

func (v *interfaceStartManager) NeededMemory() proxy.PropUint64 {
	return &proxy.ImplPropUint64{
		Impl: v,
		Name: "NeededMemory",
	}
}

type SessionManager interface {
	sessionManager // interface com.deepin.SessionManager
	proxy.Object
}

type objectSessionManager struct {
	interfaceSessionManager // interface com.deepin.SessionManager
	proxy.ImplObject
}

func NewSessionManager(conn *dbus.Conn) SessionManager {
	obj := new(objectSessionManager)
	obj.ImplObject.Init_(conn, "com.deepin.SessionManager", "/com/deepin/SessionManager")
	return obj
}

type sessionManager interface {
	GoAllowSessionDaemonRun(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	AllowSessionDaemonRun(flags dbus.Flags) (bool, error)
	GoCanHibernate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	CanHibernate(flags dbus.Flags) (bool, error)
	GoCanLogout(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	CanLogout(flags dbus.Flags) (bool, error)
	GoCanReboot(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	CanReboot(flags dbus.Flags) (bool, error)
	GoCanShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	CanShutdown(flags dbus.Flags) (bool, error)
	GoCanSuspend(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	CanSuspend(flags dbus.Flags) (bool, error)
	GoForceLogout(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ForceLogout(flags dbus.Flags) error
	GoForceReboot(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ForceReboot(flags dbus.Flags) error
	GoForceShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ForceShutdown(flags dbus.Flags) error
	GoLogout(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Logout(flags dbus.Flags) error
	GoPowerOffChoose(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	PowerOffChoose(flags dbus.Flags) error
	GoReboot(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Reboot(flags dbus.Flags) error
	GoRegister(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	Register(flags dbus.Flags, arg0 string) (bool, error)
	GoRequestHibernate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	RequestHibernate(flags dbus.Flags) error
	GoRequestLock(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	RequestLock(flags dbus.Flags) error
	GoRequestLogout(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	RequestLogout(flags dbus.Flags) error
	GoRequestReboot(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	RequestReboot(flags dbus.Flags) error
	GoRequestShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	RequestShutdown(flags dbus.Flags) error
	GoRequestSuspend(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	RequestSuspend(flags dbus.Flags) error
	GoSetLocked(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call
	SetLocked(flags dbus.Flags, arg0 bool) error
	GoShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Shutdown(flags dbus.Flags) error
	GoToggleDebug(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ToggleDebug(flags dbus.Flags) error
	ConnectUnlock(cb func()) (dbusutil.SignalHandlerId, error)
	Locked() proxy.PropBool
	CurrentUid() proxy.PropString
	Stage() proxy.PropInt32
	CurrentSessionPath() proxy.PropObjectPath
}

type interfaceSessionManager struct{}

func (v *interfaceSessionManager) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceSessionManager) GetInterfaceName_() string {
	return "com.deepin.SessionManager"
}

// method AllowSessionDaemonRun

func (v *interfaceSessionManager) GoAllowSessionDaemonRun(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AllowSessionDaemonRun", flags, ch)
}

func (*interfaceSessionManager) StoreAllowSessionDaemonRun(call *dbus.Call) (arg0 bool, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceSessionManager) AllowSessionDaemonRun(flags dbus.Flags) (bool, error) {
	return v.StoreAllowSessionDaemonRun(
		<-v.GoAllowSessionDaemonRun(flags, make(chan *dbus.Call, 1)).Done)
}

// method CanHibernate

func (v *interfaceSessionManager) GoCanHibernate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanHibernate", flags, ch)
}

func (*interfaceSessionManager) StoreCanHibernate(call *dbus.Call) (arg0 bool, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceSessionManager) CanHibernate(flags dbus.Flags) (bool, error) {
	return v.StoreCanHibernate(
		<-v.GoCanHibernate(flags, make(chan *dbus.Call, 1)).Done)
}

// method CanLogout

func (v *interfaceSessionManager) GoCanLogout(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanLogout", flags, ch)
}

func (*interfaceSessionManager) StoreCanLogout(call *dbus.Call) (arg0 bool, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceSessionManager) CanLogout(flags dbus.Flags) (bool, error) {
	return v.StoreCanLogout(
		<-v.GoCanLogout(flags, make(chan *dbus.Call, 1)).Done)
}

// method CanReboot

func (v *interfaceSessionManager) GoCanReboot(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanReboot", flags, ch)
}

func (*interfaceSessionManager) StoreCanReboot(call *dbus.Call) (arg0 bool, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceSessionManager) CanReboot(flags dbus.Flags) (bool, error) {
	return v.StoreCanReboot(
		<-v.GoCanReboot(flags, make(chan *dbus.Call, 1)).Done)
}

// method CanShutdown

func (v *interfaceSessionManager) GoCanShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanShutdown", flags, ch)
}

func (*interfaceSessionManager) StoreCanShutdown(call *dbus.Call) (arg0 bool, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceSessionManager) CanShutdown(flags dbus.Flags) (bool, error) {
	return v.StoreCanShutdown(
		<-v.GoCanShutdown(flags, make(chan *dbus.Call, 1)).Done)
}

// method CanSuspend

func (v *interfaceSessionManager) GoCanSuspend(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanSuspend", flags, ch)
}

func (*interfaceSessionManager) StoreCanSuspend(call *dbus.Call) (arg0 bool, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceSessionManager) CanSuspend(flags dbus.Flags) (bool, error) {
	return v.StoreCanSuspend(
		<-v.GoCanSuspend(flags, make(chan *dbus.Call, 1)).Done)
}

// method ForceLogout

func (v *interfaceSessionManager) GoForceLogout(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ForceLogout", flags, ch)
}

func (v *interfaceSessionManager) ForceLogout(flags dbus.Flags) error {
	return (<-v.GoForceLogout(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ForceReboot

func (v *interfaceSessionManager) GoForceReboot(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ForceReboot", flags, ch)
}

func (v *interfaceSessionManager) ForceReboot(flags dbus.Flags) error {
	return (<-v.GoForceReboot(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ForceShutdown

func (v *interfaceSessionManager) GoForceShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ForceShutdown", flags, ch)
}

func (v *interfaceSessionManager) ForceShutdown(flags dbus.Flags) error {
	return (<-v.GoForceShutdown(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Logout

func (v *interfaceSessionManager) GoLogout(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Logout", flags, ch)
}

func (v *interfaceSessionManager) Logout(flags dbus.Flags) error {
	return (<-v.GoLogout(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method PowerOffChoose

func (v *interfaceSessionManager) GoPowerOffChoose(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PowerOffChoose", flags, ch)
}

func (v *interfaceSessionManager) PowerOffChoose(flags dbus.Flags) error {
	return (<-v.GoPowerOffChoose(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Reboot

func (v *interfaceSessionManager) GoReboot(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reboot", flags, ch)
}

func (v *interfaceSessionManager) Reboot(flags dbus.Flags) error {
	return (<-v.GoReboot(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Register

func (v *interfaceSessionManager) GoRegister(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Register", flags, ch, arg0)
}

func (*interfaceSessionManager) StoreRegister(call *dbus.Call) (arg1 bool, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceSessionManager) Register(flags dbus.Flags, arg0 string) (bool, error) {
	return v.StoreRegister(
		<-v.GoRegister(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method RequestHibernate

func (v *interfaceSessionManager) GoRequestHibernate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestHibernate", flags, ch)
}

func (v *interfaceSessionManager) RequestHibernate(flags dbus.Flags) error {
	return (<-v.GoRequestHibernate(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RequestLock

func (v *interfaceSessionManager) GoRequestLock(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestLock", flags, ch)
}

func (v *interfaceSessionManager) RequestLock(flags dbus.Flags) error {
	return (<-v.GoRequestLock(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RequestLogout

func (v *interfaceSessionManager) GoRequestLogout(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestLogout", flags, ch)
}

func (v *interfaceSessionManager) RequestLogout(flags dbus.Flags) error {
	return (<-v.GoRequestLogout(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RequestReboot

func (v *interfaceSessionManager) GoRequestReboot(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestReboot", flags, ch)
}

func (v *interfaceSessionManager) RequestReboot(flags dbus.Flags) error {
	return (<-v.GoRequestReboot(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RequestShutdown

func (v *interfaceSessionManager) GoRequestShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestShutdown", flags, ch)
}

func (v *interfaceSessionManager) RequestShutdown(flags dbus.Flags) error {
	return (<-v.GoRequestShutdown(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RequestSuspend

func (v *interfaceSessionManager) GoRequestSuspend(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestSuspend", flags, ch)
}

func (v *interfaceSessionManager) RequestSuspend(flags dbus.Flags) error {
	return (<-v.GoRequestSuspend(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method SetLocked

func (v *interfaceSessionManager) GoSetLocked(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLocked", flags, ch, arg0)
}

func (v *interfaceSessionManager) SetLocked(flags dbus.Flags, arg0 bool) error {
	return (<-v.GoSetLocked(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method Shutdown

func (v *interfaceSessionManager) GoShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Shutdown", flags, ch)
}

func (v *interfaceSessionManager) Shutdown(flags dbus.Flags) error {
	return (<-v.GoShutdown(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ToggleDebug

func (v *interfaceSessionManager) GoToggleDebug(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ToggleDebug", flags, ch)
}

func (v *interfaceSessionManager) ToggleDebug(flags dbus.Flags) error {
	return (<-v.GoToggleDebug(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal Unlock

func (v *interfaceSessionManager) ConnectUnlock(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Unlock", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Unlock",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Locked b

func (v *interfaceSessionManager) Locked() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "Locked",
	}
}

// property CurrentUid s

func (v *interfaceSessionManager) CurrentUid() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "CurrentUid",
	}
}

// property Stage i

func (v *interfaceSessionManager) Stage() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "Stage",
	}
}

// property CurrentSessionPath o

func (v *interfaceSessionManager) CurrentSessionPath() proxy.PropObjectPath {
	return &proxy.ImplPropObjectPath{
		Impl: v,
		Name: "CurrentSessionPath",
	}
}

type WMSwitcher interface {
	wmSwitcher // interface com.deepin.WMSwitcher
	proxy.Object
}

type objectWMSwitcher struct {
	interfaceWmSwitcher // interface com.deepin.WMSwitcher
	proxy.ImplObject
}

func NewWMSwitcher(conn *dbus.Conn) WMSwitcher {
	obj := new(objectWMSwitcher)
	obj.ImplObject.Init_(conn, "com.deepin.SessionManager", "/com/deepin/WMSwitcher")
	return obj
}

type wmSwitcher interface {
	GoCurrentWM(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	CurrentWM(flags dbus.Flags) (string, error)
	GoRequestSwitchWM(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	RequestSwitchWM(flags dbus.Flags) error
	GoRestartWM(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	RestartWM(flags dbus.Flags) error
	GoStart2DWM(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Start2DWM(flags dbus.Flags) error
	ConnectWMChanged(cb func(arg0 string)) (dbusutil.SignalHandlerId, error)
}

type interfaceWmSwitcher struct{}

func (v *interfaceWmSwitcher) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceWmSwitcher) GetInterfaceName_() string {
	return "com.deepin.WMSwitcher"
}

// method CurrentWM

func (v *interfaceWmSwitcher) GoCurrentWM(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CurrentWM", flags, ch)
}

func (*interfaceWmSwitcher) StoreCurrentWM(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceWmSwitcher) CurrentWM(flags dbus.Flags) (string, error) {
	return v.StoreCurrentWM(
		<-v.GoCurrentWM(flags, make(chan *dbus.Call, 1)).Done)
}

// method RequestSwitchWM

func (v *interfaceWmSwitcher) GoRequestSwitchWM(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestSwitchWM", flags, ch)
}

func (v *interfaceWmSwitcher) RequestSwitchWM(flags dbus.Flags) error {
	return (<-v.GoRequestSwitchWM(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RestartWM

func (v *interfaceWmSwitcher) GoRestartWM(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RestartWM", flags, ch)
}

func (v *interfaceWmSwitcher) RestartWM(flags dbus.Flags) error {
	return (<-v.GoRestartWM(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Start2DWM

func (v *interfaceWmSwitcher) GoStart2DWM(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Start2DWM", flags, ch)
}

func (v *interfaceWmSwitcher) Start2DWM(flags dbus.Flags) error {
	return (<-v.GoStart2DWM(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal WMChanged

func (v *interfaceWmSwitcher) ConnectWMChanged(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "WMChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".WMChanged",
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

type XSettings interface {
	xSettings // interface com.deepin.XSettings
	proxy.Object
}

type objectXSettings struct {
	interfaceXSettings // interface com.deepin.XSettings
	proxy.ImplObject
}

func NewXSettings(conn *dbus.Conn) XSettings {
	obj := new(objectXSettings)
	obj.ImplObject.Init_(conn, "com.deepin.SessionManager", "/com/deepin/XSettings")
	return obj
}

type xSettings interface {
	GoGetColor(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	GetColor(flags dbus.Flags, arg0 string) ([]int16, error)
	GoGetInteger(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	GetInteger(flags dbus.Flags, arg0 string) (int32, error)
	GoGetScaleFactor(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetScaleFactor(flags dbus.Flags) (float64, error)
	GoGetScreenScaleFactors(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetScreenScaleFactors(flags dbus.Flags) (map[string]float64, error)
	GoGetString(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	GetString(flags dbus.Flags, arg0 string) (string, error)
	GoListProps(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ListProps(flags dbus.Flags) (string, error)
	GoSetColor(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 []int16) *dbus.Call
	SetColor(flags dbus.Flags, arg0 string, arg1 []int16) error
	GoSetInteger(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 int32) *dbus.Call
	SetInteger(flags dbus.Flags, arg0 string, arg1 int32) error
	GoSetScaleFactor(flags dbus.Flags, ch chan *dbus.Call, arg0 float64) *dbus.Call
	SetScaleFactor(flags dbus.Flags, arg0 float64) error
	GoSetScreenScaleFactors(flags dbus.Flags, ch chan *dbus.Call, arg0 map[string]float64) *dbus.Call
	SetScreenScaleFactors(flags dbus.Flags, arg0 map[string]float64) error
	GoSetString(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call
	SetString(flags dbus.Flags, arg0 string, arg1 string) error
	ConnectSetScaleFactorStarted(cb func()) (dbusutil.SignalHandlerId, error)
	ConnectSetScaleFactorDone(cb func()) (dbusutil.SignalHandlerId, error)
}

type interfaceXSettings struct{}

func (v *interfaceXSettings) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceXSettings) GetInterfaceName_() string {
	return "com.deepin.XSettings"
}

// method GetColor

func (v *interfaceXSettings) GoGetColor(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetColor", flags, ch, arg0)
}

func (*interfaceXSettings) StoreGetColor(call *dbus.Call) (arg1 []int16, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceXSettings) GetColor(flags dbus.Flags, arg0 string) ([]int16, error) {
	return v.StoreGetColor(
		<-v.GoGetColor(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method GetInteger

func (v *interfaceXSettings) GoGetInteger(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetInteger", flags, ch, arg0)
}

func (*interfaceXSettings) StoreGetInteger(call *dbus.Call) (arg1 int32, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceXSettings) GetInteger(flags dbus.Flags, arg0 string) (int32, error) {
	return v.StoreGetInteger(
		<-v.GoGetInteger(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method GetScaleFactor

func (v *interfaceXSettings) GoGetScaleFactor(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetScaleFactor", flags, ch)
}

func (*interfaceXSettings) StoreGetScaleFactor(call *dbus.Call) (arg0 float64, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceXSettings) GetScaleFactor(flags dbus.Flags) (float64, error) {
	return v.StoreGetScaleFactor(
		<-v.GoGetScaleFactor(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetScreenScaleFactors

func (v *interfaceXSettings) GoGetScreenScaleFactors(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetScreenScaleFactors", flags, ch)
}

func (*interfaceXSettings) StoreGetScreenScaleFactors(call *dbus.Call) (arg0 map[string]float64, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceXSettings) GetScreenScaleFactors(flags dbus.Flags) (map[string]float64, error) {
	return v.StoreGetScreenScaleFactors(
		<-v.GoGetScreenScaleFactors(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetString

func (v *interfaceXSettings) GoGetString(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetString", flags, ch, arg0)
}

func (*interfaceXSettings) StoreGetString(call *dbus.Call) (arg1 string, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceXSettings) GetString(flags dbus.Flags, arg0 string) (string, error) {
	return v.StoreGetString(
		<-v.GoGetString(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method ListProps

func (v *interfaceXSettings) GoListProps(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListProps", flags, ch)
}

func (*interfaceXSettings) StoreListProps(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceXSettings) ListProps(flags dbus.Flags) (string, error) {
	return v.StoreListProps(
		<-v.GoListProps(flags, make(chan *dbus.Call, 1)).Done)
}

// method SetColor

func (v *interfaceXSettings) GoSetColor(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 []int16) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetColor", flags, ch, arg0, arg1)
}

func (v *interfaceXSettings) SetColor(flags dbus.Flags, arg0 string, arg1 []int16) error {
	return (<-v.GoSetColor(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetInteger

func (v *interfaceXSettings) GoSetInteger(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetInteger", flags, ch, arg0, arg1)
}

func (v *interfaceXSettings) SetInteger(flags dbus.Flags, arg0 string, arg1 int32) error {
	return (<-v.GoSetInteger(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetScaleFactor

func (v *interfaceXSettings) GoSetScaleFactor(flags dbus.Flags, ch chan *dbus.Call, arg0 float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetScaleFactor", flags, ch, arg0)
}

func (v *interfaceXSettings) SetScaleFactor(flags dbus.Flags, arg0 float64) error {
	return (<-v.GoSetScaleFactor(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SetScreenScaleFactors

func (v *interfaceXSettings) GoSetScreenScaleFactors(flags dbus.Flags, ch chan *dbus.Call, arg0 map[string]float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetScreenScaleFactors", flags, ch, arg0)
}

func (v *interfaceXSettings) SetScreenScaleFactors(flags dbus.Flags, arg0 map[string]float64) error {
	return (<-v.GoSetScreenScaleFactors(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SetString

func (v *interfaceXSettings) GoSetString(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetString", flags, ch, arg0, arg1)
}

func (v *interfaceXSettings) SetString(flags dbus.Flags, arg0 string, arg1 string) error {
	return (<-v.GoSetString(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// signal SetScaleFactorStarted

func (v *interfaceXSettings) ConnectSetScaleFactorStarted(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SetScaleFactorStarted", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SetScaleFactorStarted",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal SetScaleFactorDone

func (v *interfaceXSettings) ConnectSetScaleFactorDone(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SetScaleFactorDone", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SetScaleFactorDone",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
