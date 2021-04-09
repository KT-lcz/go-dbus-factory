// Code generated by "./generator ./com.huawei.fingerprint"; DO NOT EDIT.

package fingerprint

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type Fingerprint interface {
	fingerprint // interface com.huawei.Fingerprint
	proxy.Object
}

type objectFingerprint struct {
	interfaceFingerprint // interface com.huawei.Fingerprint
	proxy.ImplObject
}

func NewFingerprint(conn *dbus.Conn) Fingerprint {
	obj := new(objectFingerprint)
	obj.ImplObject.Init_(conn, "com.huawei.Fingerprint", "/com/huawei/Fingerprint")
	return obj
}

type fingerprint interface {
	GoSearchDevice(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	SearchDevice(flags dbus.Flags) (bool, error)
	GoIdentify(flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call
	Identify(flags dbus.Flags, uuid string) error
	GoIdentifyWithMultipleUser(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	IdentifyWithMultipleUser(flags dbus.Flags) error
	GoGetStatus(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetStatus(flags dbus.Flags) (int32, error)
	GoEnroll(flags dbus.Flags, ch chan *dbus.Call, filePath string, uuid string) *dbus.Call
	Enroll(flags dbus.Flags, filePath string, uuid string) error
	GoClose(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Close(flags dbus.Flags) (int32, error)
	GoReload(flags dbus.Flags, ch chan *dbus.Call, deleteType int32) *dbus.Call
	Reload(flags dbus.Flags, deleteType int32) (int32, error)
	GoClearPovImage(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ClearPovImage(flags dbus.Flags) (int32, error)
	ConnectEnrollStatus(cb func(progress int32, result int32)) (dbusutil.SignalHandlerId, error)
	ConnectIdentifyStatus(cb func(result int32)) (dbusutil.SignalHandlerId, error)
	ConnectIdentifyNoAccount(cb func(result int32, userName string)) (dbusutil.SignalHandlerId, error)
	ConnectVerifyStatus(cb func(result int32)) (dbusutil.SignalHandlerId, error)
}

type interfaceFingerprint struct{}

func (v *interfaceFingerprint) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceFingerprint) GetInterfaceName_() string {
	return "com.huawei.Fingerprint"
}

// method SearchDevice

func (v *interfaceFingerprint) GoSearchDevice(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SearchDevice", flags, ch)
}

func (*interfaceFingerprint) StoreSearchDevice(call *dbus.Call) (result bool, err error) {
	err = call.Store(&result)
	return
}

func (v *interfaceFingerprint) SearchDevice(flags dbus.Flags) (bool, error) {
	return v.StoreSearchDevice(
		<-v.GoSearchDevice(flags, make(chan *dbus.Call, 1)).Done)
}

// method Identify

func (v *interfaceFingerprint) GoIdentify(flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Identify", flags, ch, uuid)
}

func (v *interfaceFingerprint) Identify(flags dbus.Flags, uuid string) error {
	return (<-v.GoIdentify(flags, make(chan *dbus.Call, 1), uuid).Done).Err
}

// method IdentifyWithMultipleUser

func (v *interfaceFingerprint) GoIdentifyWithMultipleUser(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IdentifyWithMultipleUser", flags, ch)
}

func (v *interfaceFingerprint) IdentifyWithMultipleUser(flags dbus.Flags) error {
	return (<-v.GoIdentifyWithMultipleUser(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method GetStatus

func (v *interfaceFingerprint) GoGetStatus(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetStatus", flags, ch)
}

func (*interfaceFingerprint) StoreGetStatus(call *dbus.Call) (result int32, err error) {
	err = call.Store(&result)
	return
}

func (v *interfaceFingerprint) GetStatus(flags dbus.Flags) (int32, error) {
	return v.StoreGetStatus(
		<-v.GoGetStatus(flags, make(chan *dbus.Call, 1)).Done)
}

// method Enroll

func (v *interfaceFingerprint) GoEnroll(flags dbus.Flags, ch chan *dbus.Call, filePath string, uuid string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Enroll", flags, ch, filePath, uuid)
}

func (v *interfaceFingerprint) Enroll(flags dbus.Flags, filePath string, uuid string) error {
	return (<-v.GoEnroll(flags, make(chan *dbus.Call, 1), filePath, uuid).Done).Err
}

// method Close

func (v *interfaceFingerprint) GoClose(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Close", flags, ch)
}

func (*interfaceFingerprint) StoreClose(call *dbus.Call) (result int32, err error) {
	err = call.Store(&result)
	return
}

func (v *interfaceFingerprint) Close(flags dbus.Flags) (int32, error) {
	return v.StoreClose(
		<-v.GoClose(flags, make(chan *dbus.Call, 1)).Done)
}

// method Reload

func (v *interfaceFingerprint) GoReload(flags dbus.Flags, ch chan *dbus.Call, deleteType int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reload", flags, ch, deleteType)
}

func (*interfaceFingerprint) StoreReload(call *dbus.Call) (result int32, err error) {
	err = call.Store(&result)
	return
}

func (v *interfaceFingerprint) Reload(flags dbus.Flags, deleteType int32) (int32, error) {
	return v.StoreReload(
		<-v.GoReload(flags, make(chan *dbus.Call, 1), deleteType).Done)
}

// method ClearPovImage

func (v *interfaceFingerprint) GoClearPovImage(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ClearPovImage", flags, ch)
}

func (*interfaceFingerprint) StoreClearPovImage(call *dbus.Call) (result int32, err error) {
	err = call.Store(&result)
	return
}

func (v *interfaceFingerprint) ClearPovImage(flags dbus.Flags) (int32, error) {
	return v.StoreClearPovImage(
		<-v.GoClearPovImage(flags, make(chan *dbus.Call, 1)).Done)
}

// signal EnrollStatus

func (v *interfaceFingerprint) ConnectEnrollStatus(cb func(progress int32, result int32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "EnrollStatus", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".EnrollStatus",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var progress int32
		var result int32
		err := dbus.Store(sig.Body, &progress, &result)
		if err == nil {
			cb(progress, result)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal IdentifyStatus

func (v *interfaceFingerprint) ConnectIdentifyStatus(cb func(result int32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "IdentifyStatus", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".IdentifyStatus",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var result int32
		err := dbus.Store(sig.Body, &result)
		if err == nil {
			cb(result)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal IdentifyNoAccount

func (v *interfaceFingerprint) ConnectIdentifyNoAccount(cb func(result int32, userName string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "IdentifyNoAccount", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".IdentifyNoAccount",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var result int32
		var userName string
		err := dbus.Store(sig.Body, &result, &userName)
		if err == nil {
			cb(result, userName)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal VerifyStatus

func (v *interfaceFingerprint) ConnectVerifyStatus(cb func(result int32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "VerifyStatus", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".VerifyStatus",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var result int32
		err := dbus.Store(sig.Body, &result)
		if err == nil {
			cb(result)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
