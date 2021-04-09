// Code generated by "./generator ./org.freedesktop.geoclue2"; DO NOT EDIT.

package geoclue2

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type Manager interface {
	manager // interface org.freedesktop.GeoClue2.Manager
	proxy.Object
}

type objectManager struct {
	interfaceManager // interface org.freedesktop.GeoClue2.Manager
	proxy.ImplObject
}

func NewManager(conn *dbus.Conn) Manager {
	obj := new(objectManager)
	obj.ImplObject.Init_(conn, "org.freedesktop.GeoClue2", "/org/freedesktop/GeoClue2/Manager")
	return obj
}

type manager interface {
	GoGetClient(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetClient(flags dbus.Flags) (dbus.ObjectPath, error)
	GoAddAgent(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call
	AddAgent(flags dbus.Flags, id string) error
	InUse() proxy.PropBool
	AvailableAccuracyLevel() proxy.PropUint32
}

type interfaceManager struct{}

func (v *interfaceManager) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceManager) GetInterfaceName_() string {
	return "org.freedesktop.GeoClue2.Manager"
}

// method GetClient

func (v *interfaceManager) GoGetClient(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetClient", flags, ch)
}

func (*interfaceManager) StoreGetClient(call *dbus.Call) (client dbus.ObjectPath, err error) {
	err = call.Store(&client)
	return
}

func (v *interfaceManager) GetClient(flags dbus.Flags) (dbus.ObjectPath, error) {
	return v.StoreGetClient(
		<-v.GoGetClient(flags, make(chan *dbus.Call, 1)).Done)
}

// method AddAgent

func (v *interfaceManager) GoAddAgent(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddAgent", flags, ch, id)
}

func (v *interfaceManager) AddAgent(flags dbus.Flags, id string) error {
	return (<-v.GoAddAgent(flags, make(chan *dbus.Call, 1), id).Done).Err
}

// property InUse b

func (v *interfaceManager) InUse() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "InUse",
	}
}

// property AvailableAccuracyLevel u

func (v *interfaceManager) AvailableAccuracyLevel() proxy.PropUint32 {
	return &proxy.ImplPropUint32{
		Impl: v,
		Name: "AvailableAccuracyLevel",
	}
}

type Client interface {
	client // interface org.freedesktop.GeoClue2.Client
	proxy.Object
}

type objectClient struct {
	interfaceClient // interface org.freedesktop.GeoClue2.Client
	proxy.ImplObject
}

func NewClient(conn *dbus.Conn, path dbus.ObjectPath) (Client, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectClient)
	obj.ImplObject.Init_(conn, "org.freedesktop.GeoClue2", path)
	return obj, nil
}

type client interface {
	GoStart(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Start(flags dbus.Flags) error
	GoStop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Stop(flags dbus.Flags) error
	ConnectLocationUpdated(cb func(old dbus.ObjectPath, new dbus.ObjectPath)) (dbusutil.SignalHandlerId, error)
	Location() proxy.PropObjectPath
	DistanceThreshold() proxy.PropUint32
	TimeThreshold() proxy.PropUint32
	DesktopId() proxy.PropString
	RequestedAccuracyLevel() proxy.PropUint32
	Active() proxy.PropBool
}

type interfaceClient struct{}

func (v *interfaceClient) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceClient) GetInterfaceName_() string {
	return "org.freedesktop.GeoClue2.Client"
}

// method Start

func (v *interfaceClient) GoStart(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Start", flags, ch)
}

func (v *interfaceClient) Start(flags dbus.Flags) error {
	return (<-v.GoStart(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Stop

func (v *interfaceClient) GoStop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Stop", flags, ch)
}

func (v *interfaceClient) Stop(flags dbus.Flags) error {
	return (<-v.GoStop(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal LocationUpdated

func (v *interfaceClient) ConnectLocationUpdated(cb func(old dbus.ObjectPath, new dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "LocationUpdated", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".LocationUpdated",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var old dbus.ObjectPath
		var new dbus.ObjectPath
		err := dbus.Store(sig.Body, &old, &new)
		if err == nil {
			cb(old, new)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Location o

func (v *interfaceClient) Location() proxy.PropObjectPath {
	return &proxy.ImplPropObjectPath{
		Impl: v,
		Name: "Location",
	}
}

// property DistanceThreshold u

func (v *interfaceClient) DistanceThreshold() proxy.PropUint32 {
	return &proxy.ImplPropUint32{
		Impl: v,
		Name: "DistanceThreshold",
	}
}

// property TimeThreshold u

func (v *interfaceClient) TimeThreshold() proxy.PropUint32 {
	return &proxy.ImplPropUint32{
		Impl: v,
		Name: "TimeThreshold",
	}
}

// property DesktopId s

func (v *interfaceClient) DesktopId() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "DesktopId",
	}
}

// property RequestedAccuracyLevel u

func (v *interfaceClient) RequestedAccuracyLevel() proxy.PropUint32 {
	return &proxy.ImplPropUint32{
		Impl: v,
		Name: "RequestedAccuracyLevel",
	}
}

// property Active b

func (v *interfaceClient) Active() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "Active",
	}
}

type Location interface {
	location // interface org.freedesktop.GeoClue2.Location
	proxy.Object
}

type objectLocation struct {
	interfaceLocation // interface org.freedesktop.GeoClue2.Location
	proxy.ImplObject
}

func NewLocation(conn *dbus.Conn, path dbus.ObjectPath) (Location, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectLocation)
	obj.ImplObject.Init_(conn, "org.freedesktop.GeoClue2", path)
	return obj, nil
}

type location interface {
	Latitude() proxy.PropDouble
	Longitude() proxy.PropDouble
	Accuracy() proxy.PropDouble
	Altitude() proxy.PropDouble
	Speed() proxy.PropDouble
	Heading() proxy.PropDouble
	Description() proxy.PropString
	Timestamp() PropTimestamp
}

type interfaceLocation struct{}

func (v *interfaceLocation) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceLocation) GetInterfaceName_() string {
	return "org.freedesktop.GeoClue2.Location"
}

// property Latitude d

func (v *interfaceLocation) Latitude() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "Latitude",
	}
}

// property Longitude d

func (v *interfaceLocation) Longitude() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "Longitude",
	}
}

// property Accuracy d

func (v *interfaceLocation) Accuracy() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "Accuracy",
	}
}

// property Altitude d

func (v *interfaceLocation) Altitude() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "Altitude",
	}
}

// property Speed d

func (v *interfaceLocation) Speed() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "Speed",
	}
}

// property Heading d

func (v *interfaceLocation) Heading() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "Heading",
	}
}

// property Description s

func (v *interfaceLocation) Description() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Description",
	}
}

type PropTimestamp interface {
	Get(flags dbus.Flags) (value Timestamp, err error)
	Set(flags dbus.Flags, value Timestamp) error
	ConnectChanged(cb func(hasValue bool, value Timestamp)) error
}

type implPropTimestamp struct {
	Impl proxy.Implementer
	Name string
}

func (p implPropTimestamp) Get(flags dbus.Flags) (value Timestamp, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p implPropTimestamp) Set(flags dbus.Flags, value Timestamp) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p implPropTimestamp) ConnectChanged(cb func(hasValue bool, value Timestamp)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v Timestamp
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, Timestamp{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}

// property Timestamp (tt)

func (v *interfaceLocation) Timestamp() PropTimestamp {
	return &implPropTimestamp{
		Impl: v,
		Name: "Timestamp",
	}
}
