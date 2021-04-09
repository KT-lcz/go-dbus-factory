// Code generated by "./generator ./com.deepin.daemon.authenticate.fingerprint"; DO NOT EDIT.

package fingerprint

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type Device interface {
	device // interface com.deepin.daemon.Authenticate.Fingerprint.Device
	proxy.Object
}

type objectDevice struct {
	interfaceDevice // interface com.deepin.daemon.Authenticate.Fingerprint.Device
	proxy.ImplObject
}

func NewDevice(conn *dbus.Conn, serviceName string, path dbus.ObjectPath) (Device, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectDevice)
	obj.ImplObject.Init_(conn, serviceName, path)
	return obj, nil
}

type device interface {
	SetInterfaceName_(name string)
	GoClaim(flags dbus.Flags, ch chan *dbus.Call, userId string, claimed bool) *dbus.Call
	Claim(flags dbus.Flags, userId string, claimed bool) error
	GoDeleteAllFingers(flags dbus.Flags, ch chan *dbus.Call, userId string) *dbus.Call
	DeleteAllFingers(flags dbus.Flags, userId string) error
	GoDeleteFinger(flags dbus.Flags, ch chan *dbus.Call, userId string, finger string) *dbus.Call
	DeleteFinger(flags dbus.Flags, userId string, finger string) error
	GoEnroll(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call
	Enroll(flags dbus.Flags, finger string) error
	GoListFingers(flags dbus.Flags, ch chan *dbus.Call, userId string) *dbus.Call
	ListFingers(flags dbus.Flags, userId string) ([]string, error)
	GoRenameFinger(flags dbus.Flags, ch chan *dbus.Call, userId string, finger string, newName string) *dbus.Call
	RenameFinger(flags dbus.Flags, userId string, finger string, newName string) error
	GoStopEnroll(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	StopEnroll(flags dbus.Flags) error
	GoStopVerify(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	StopVerify(flags dbus.Flags) error
	GoVerify(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call
	Verify(flags dbus.Flags, finger string) error
	ConnectEnrollStatus(cb func(userId string, code int32, msg string)) (dbusutil.SignalHandlerId, error)
	ConnectVerifyStatus(cb func(userId string, code int32, msg string)) (dbusutil.SignalHandlerId, error)
	ConnectTouch(cb func(userId string, pressed bool)) (dbusutil.SignalHandlerId, error)
	Name() proxy.PropString
	State() proxy.PropInt32
	Type() proxy.PropInt32
	Capability() proxy.PropInt32
}

type interfaceDevice struct{}

func (v *interfaceDevice) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (v *interfaceDevice) SetInterfaceName_(name string) {
	v.GetObject_().SetExtra("customIfc", name)
}

func (v *interfaceDevice) GetInterfaceName_() string {
	ifcName, _ := v.GetObject_().GetExtra("customIfc")
	ifcNameStr, _ := ifcName.(string)
	return ifcNameStr
}

// method Claim

func (v *interfaceDevice) GoClaim(flags dbus.Flags, ch chan *dbus.Call, userId string, claimed bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Claim", flags, ch, userId, claimed)
}

func (v *interfaceDevice) Claim(flags dbus.Flags, userId string, claimed bool) error {
	return (<-v.GoClaim(flags, make(chan *dbus.Call, 1), userId, claimed).Done).Err
}

// method DeleteAllFingers

func (v *interfaceDevice) GoDeleteAllFingers(flags dbus.Flags, ch chan *dbus.Call, userId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteAllFingers", flags, ch, userId)
}

func (v *interfaceDevice) DeleteAllFingers(flags dbus.Flags, userId string) error {
	return (<-v.GoDeleteAllFingers(flags, make(chan *dbus.Call, 1), userId).Done).Err
}

// method DeleteFinger

func (v *interfaceDevice) GoDeleteFinger(flags dbus.Flags, ch chan *dbus.Call, userId string, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteFinger", flags, ch, userId, finger)
}

func (v *interfaceDevice) DeleteFinger(flags dbus.Flags, userId string, finger string) error {
	return (<-v.GoDeleteFinger(flags, make(chan *dbus.Call, 1), userId, finger).Done).Err
}

// method Enroll

func (v *interfaceDevice) GoEnroll(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Enroll", flags, ch, finger)
}

func (v *interfaceDevice) Enroll(flags dbus.Flags, finger string) error {
	return (<-v.GoEnroll(flags, make(chan *dbus.Call, 1), finger).Done).Err
}

// method ListFingers

func (v *interfaceDevice) GoListFingers(flags dbus.Flags, ch chan *dbus.Call, userId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListFingers", flags, ch, userId)
}

func (*interfaceDevice) StoreListFingers(call *dbus.Call) (fingers []string, err error) {
	err = call.Store(&fingers)
	return
}

func (v *interfaceDevice) ListFingers(flags dbus.Flags, userId string) ([]string, error) {
	return v.StoreListFingers(
		<-v.GoListFingers(flags, make(chan *dbus.Call, 1), userId).Done)
}

// method RenameFinger

func (v *interfaceDevice) GoRenameFinger(flags dbus.Flags, ch chan *dbus.Call, userId string, finger string, newName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RenameFinger", flags, ch, userId, finger, newName)
}

func (v *interfaceDevice) RenameFinger(flags dbus.Flags, userId string, finger string, newName string) error {
	return (<-v.GoRenameFinger(flags, make(chan *dbus.Call, 1), userId, finger, newName).Done).Err
}

// method StopEnroll

func (v *interfaceDevice) GoStopEnroll(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopEnroll", flags, ch)
}

func (v *interfaceDevice) StopEnroll(flags dbus.Flags) error {
	return (<-v.GoStopEnroll(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method StopVerify

func (v *interfaceDevice) GoStopVerify(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopVerify", flags, ch)
}

func (v *interfaceDevice) StopVerify(flags dbus.Flags) error {
	return (<-v.GoStopVerify(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Verify

func (v *interfaceDevice) GoVerify(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Verify", flags, ch, finger)
}

func (v *interfaceDevice) Verify(flags dbus.Flags, finger string) error {
	return (<-v.GoVerify(flags, make(chan *dbus.Call, 1), finger).Done).Err
}

// signal EnrollStatus

func (v *interfaceDevice) ConnectEnrollStatus(cb func(userId string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
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
		var userId string
		var code int32
		var msg string
		err := dbus.Store(sig.Body, &userId, &code, &msg)
		if err == nil {
			cb(userId, code, msg)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal VerifyStatus

func (v *interfaceDevice) ConnectVerifyStatus(cb func(userId string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
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
		var userId string
		var code int32
		var msg string
		err := dbus.Store(sig.Body, &userId, &code, &msg)
		if err == nil {
			cb(userId, code, msg)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Touch

func (v *interfaceDevice) ConnectTouch(cb func(userId string, pressed bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Touch", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Touch",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var userId string
		var pressed bool
		err := dbus.Store(sig.Body, &userId, &pressed)
		if err == nil {
			cb(userId, pressed)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Name s

func (v *interfaceDevice) Name() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Name",
	}
}

// property State i

func (v *interfaceDevice) State() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "State",
	}
}

// property Type i

func (v *interfaceDevice) Type() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "Type",
	}
}

// property Capability i

func (v *interfaceDevice) Capability() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "Capability",
	}
}

type CommonDevice interface {
	commonDevice // interface com.deepin.daemon.Authenticate.Fingerprint.CommonDevice
	proxy.Object
}

type objectCommonDevice struct {
	interfaceCommonDevice // interface com.deepin.daemon.Authenticate.Fingerprint.CommonDevice
	proxy.ImplObject
}

func NewCommonDevice(conn *dbus.Conn, serviceName string, path dbus.ObjectPath) (CommonDevice, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectCommonDevice)
	obj.ImplObject.Init_(conn, serviceName, path)
	return obj, nil
}

type commonDevice interface {
	SetInterfaceName_(name string)
	GoClaim(flags dbus.Flags, ch chan *dbus.Call, userId string, claimed bool) *dbus.Call
	Claim(flags dbus.Flags, userId string, claimed bool) error
	GoDeleteAllFingers(flags dbus.Flags, ch chan *dbus.Call, userId string) *dbus.Call
	DeleteAllFingers(flags dbus.Flags, userId string) error
	GoDeleteFinger(flags dbus.Flags, ch chan *dbus.Call, userId string, finger string) *dbus.Call
	DeleteFinger(flags dbus.Flags, userId string, finger string) error
	GoEnroll(flags dbus.Flags, ch chan *dbus.Call, username string, finger string) *dbus.Call
	Enroll(flags dbus.Flags, username string, finger string) error
	GoListFingers(flags dbus.Flags, ch chan *dbus.Call, userId string) *dbus.Call
	ListFingers(flags dbus.Flags, userId string) ([]string, error)
	GoRenameFinger(flags dbus.Flags, ch chan *dbus.Call, userId string, finger string, newName string) *dbus.Call
	RenameFinger(flags dbus.Flags, userId string, finger string, newName string) error
	GoStopEnroll(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	StopEnroll(flags dbus.Flags) error
	GoStopVerify(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	StopVerify(flags dbus.Flags) error
	GoVerify(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call
	Verify(flags dbus.Flags, finger string) error
	ConnectEnrollStatus(cb func(userId string, code int32, msg string)) (dbusutil.SignalHandlerId, error)
	ConnectVerifyStatus(cb func(userId string, code int32, msg string)) (dbusutil.SignalHandlerId, error)
	ConnectTouch(cb func(userId string, pressed bool)) (dbusutil.SignalHandlerId, error)
	Name() proxy.PropString
	State() proxy.PropInt32
	Type() proxy.PropInt32
	Capability() proxy.PropInt32
}

type interfaceCommonDevice struct{}

func (v *interfaceCommonDevice) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (v *interfaceCommonDevice) SetInterfaceName_(name string) {
	v.GetObject_().SetExtra("customIfc", name)
}

func (v *interfaceCommonDevice) GetInterfaceName_() string {
	ifcName, _ := v.GetObject_().GetExtra("customIfc")
	ifcNameStr, _ := ifcName.(string)
	return ifcNameStr
}

// method Claim

func (v *interfaceCommonDevice) GoClaim(flags dbus.Flags, ch chan *dbus.Call, userId string, claimed bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Claim", flags, ch, userId, claimed)
}

func (v *interfaceCommonDevice) Claim(flags dbus.Flags, userId string, claimed bool) error {
	return (<-v.GoClaim(flags, make(chan *dbus.Call, 1), userId, claimed).Done).Err
}

// method DeleteAllFingers

func (v *interfaceCommonDevice) GoDeleteAllFingers(flags dbus.Flags, ch chan *dbus.Call, userId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteAllFingers", flags, ch, userId)
}

func (v *interfaceCommonDevice) DeleteAllFingers(flags dbus.Flags, userId string) error {
	return (<-v.GoDeleteAllFingers(flags, make(chan *dbus.Call, 1), userId).Done).Err
}

// method DeleteFinger

func (v *interfaceCommonDevice) GoDeleteFinger(flags dbus.Flags, ch chan *dbus.Call, userId string, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteFinger", flags, ch, userId, finger)
}

func (v *interfaceCommonDevice) DeleteFinger(flags dbus.Flags, userId string, finger string) error {
	return (<-v.GoDeleteFinger(flags, make(chan *dbus.Call, 1), userId, finger).Done).Err
}

// method Enroll

func (v *interfaceCommonDevice) GoEnroll(flags dbus.Flags, ch chan *dbus.Call, username string, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Enroll", flags, ch, username, finger)
}

func (v *interfaceCommonDevice) Enroll(flags dbus.Flags, username string, finger string) error {
	return (<-v.GoEnroll(flags, make(chan *dbus.Call, 1), username, finger).Done).Err
}

// method ListFingers

func (v *interfaceCommonDevice) GoListFingers(flags dbus.Flags, ch chan *dbus.Call, userId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListFingers", flags, ch, userId)
}

func (*interfaceCommonDevice) StoreListFingers(call *dbus.Call) (fingers []string, err error) {
	err = call.Store(&fingers)
	return
}

func (v *interfaceCommonDevice) ListFingers(flags dbus.Flags, userId string) ([]string, error) {
	return v.StoreListFingers(
		<-v.GoListFingers(flags, make(chan *dbus.Call, 1), userId).Done)
}

// method RenameFinger

func (v *interfaceCommonDevice) GoRenameFinger(flags dbus.Flags, ch chan *dbus.Call, userId string, finger string, newName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RenameFinger", flags, ch, userId, finger, newName)
}

func (v *interfaceCommonDevice) RenameFinger(flags dbus.Flags, userId string, finger string, newName string) error {
	return (<-v.GoRenameFinger(flags, make(chan *dbus.Call, 1), userId, finger, newName).Done).Err
}

// method StopEnroll

func (v *interfaceCommonDevice) GoStopEnroll(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopEnroll", flags, ch)
}

func (v *interfaceCommonDevice) StopEnroll(flags dbus.Flags) error {
	return (<-v.GoStopEnroll(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method StopVerify

func (v *interfaceCommonDevice) GoStopVerify(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopVerify", flags, ch)
}

func (v *interfaceCommonDevice) StopVerify(flags dbus.Flags) error {
	return (<-v.GoStopVerify(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Verify

func (v *interfaceCommonDevice) GoVerify(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Verify", flags, ch, finger)
}

func (v *interfaceCommonDevice) Verify(flags dbus.Flags, finger string) error {
	return (<-v.GoVerify(flags, make(chan *dbus.Call, 1), finger).Done).Err
}

// signal EnrollStatus

func (v *interfaceCommonDevice) ConnectEnrollStatus(cb func(userId string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
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
		var userId string
		var code int32
		var msg string
		err := dbus.Store(sig.Body, &userId, &code, &msg)
		if err == nil {
			cb(userId, code, msg)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal VerifyStatus

func (v *interfaceCommonDevice) ConnectVerifyStatus(cb func(userId string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
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
		var userId string
		var code int32
		var msg string
		err := dbus.Store(sig.Body, &userId, &code, &msg)
		if err == nil {
			cb(userId, code, msg)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Touch

func (v *interfaceCommonDevice) ConnectTouch(cb func(userId string, pressed bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Touch", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Touch",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var userId string
		var pressed bool
		err := dbus.Store(sig.Body, &userId, &pressed)
		if err == nil {
			cb(userId, pressed)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Name s

func (v *interfaceCommonDevice) Name() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Name",
	}
}

// property State i

func (v *interfaceCommonDevice) State() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "State",
	}
}

// property Type i

func (v *interfaceCommonDevice) Type() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "Type",
	}
}

// property Capability i

func (v *interfaceCommonDevice) Capability() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "Capability",
	}
}
