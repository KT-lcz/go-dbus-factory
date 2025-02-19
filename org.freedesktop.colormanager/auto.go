// Code generated by "./generator ./org.freedesktop.colormanager"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package colormanager

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type Manager interface {
	manager // interface org.freedesktop.ColorManager
	proxy.Object
}

type objectManager struct {
	interfaceManager // interface org.freedesktop.ColorManager
	proxy.ImplObject
}

func NewManager(conn *dbus.Conn) Manager {
	obj := new(objectManager)
	obj.ImplObject.Init_(conn, "org.freedesktop.ColorManager", "/org/freedesktop/ColorManager")
	return obj
}

type manager interface {
	GoGetDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetDevices(flags dbus.Flags) ([]dbus.ObjectPath, error)
	GoGetDevicesByKind(flags dbus.Flags, ch chan *dbus.Call, kind string) *dbus.Call
	GetDevicesByKind(flags dbus.Flags, kind string) ([]dbus.ObjectPath, error)
	GoFindDeviceById(flags dbus.Flags, ch chan *dbus.Call, device_id string) *dbus.Call
	FindDeviceById(flags dbus.Flags, device_id string) (dbus.ObjectPath, error)
	GoFindSensorById(flags dbus.Flags, ch chan *dbus.Call, sensor_id string) *dbus.Call
	FindSensorById(flags dbus.Flags, sensor_id string) (dbus.ObjectPath, error)
	GoFindDeviceByProperty(flags dbus.Flags, ch chan *dbus.Call, key string, value string) *dbus.Call
	FindDeviceByProperty(flags dbus.Flags, key string, value string) (dbus.ObjectPath, error)
	GoFindProfileById(flags dbus.Flags, ch chan *dbus.Call, profile_id string) *dbus.Call
	FindProfileById(flags dbus.Flags, profile_id string) (dbus.ObjectPath, error)
	GoFindProfileByProperty(flags dbus.Flags, ch chan *dbus.Call, key string, value string) *dbus.Call
	FindProfileByProperty(flags dbus.Flags, key string, value string) (dbus.ObjectPath, error)
	GoFindProfileByFilename(flags dbus.Flags, ch chan *dbus.Call, filename string) *dbus.Call
	FindProfileByFilename(flags dbus.Flags, filename string) (dbus.ObjectPath, error)
	GoGetStandardSpace(flags dbus.Flags, ch chan *dbus.Call, standard_space string) *dbus.Call
	GetStandardSpace(flags dbus.Flags, standard_space string) (dbus.ObjectPath, error)
	GoGetProfiles(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetProfiles(flags dbus.Flags) ([]dbus.ObjectPath, error)
	GoGetSensors(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetSensors(flags dbus.Flags) ([]dbus.ObjectPath, error)
	GoGetProfilesByKind(flags dbus.Flags, ch chan *dbus.Call, kind string) *dbus.Call
	GetProfilesByKind(flags dbus.Flags, kind string) ([]dbus.ObjectPath, error)
	GoCreateProfileWithFd(flags dbus.Flags, ch chan *dbus.Call, profile_id string, scope string, handle dbus.UnixFD, properties map[string]string) *dbus.Call
	CreateProfileWithFd(flags dbus.Flags, profile_id string, scope string, handle dbus.UnixFD, properties map[string]string) (dbus.ObjectPath, error)
	GoCreateProfile(flags dbus.Flags, ch chan *dbus.Call, profile_id string, scope string, properties map[string]string) *dbus.Call
	CreateProfile(flags dbus.Flags, profile_id string, scope string, properties map[string]string) (dbus.ObjectPath, error)
	GoCreateDevice(flags dbus.Flags, ch chan *dbus.Call, device_id string, scope string, properties map[string]string) *dbus.Call
	CreateDevice(flags dbus.Flags, device_id string, scope string, properties map[string]string) (dbus.ObjectPath, error)
	GoDeleteDevice(flags dbus.Flags, ch chan *dbus.Call, object_path dbus.ObjectPath) *dbus.Call
	DeleteDevice(flags dbus.Flags, object_path dbus.ObjectPath) error
	GoDeleteProfile(flags dbus.Flags, ch chan *dbus.Call, object_path dbus.ObjectPath) *dbus.Call
	DeleteProfile(flags dbus.Flags, object_path dbus.ObjectPath) error
	ConnectChanged(cb func()) (dbusutil.SignalHandlerId, error)
	ConnectDeviceAdded(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error)
	ConnectDeviceRemoved(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error)
	ConnectDeviceChanged(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error)
	ConnectProfileAdded(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error)
	ConnectProfileRemoved(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error)
	ConnectSensorAdded(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error)
	ConnectSensorRemoved(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error)
	ConnectProfileChanged(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error)
	DaemonVersion() proxy.PropString
	SystemVendor() proxy.PropString
	SystemModel() proxy.PropString
}

type interfaceManager struct{}

func (v *interfaceManager) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceManager) GetInterfaceName_() string {
	return "org.freedesktop.ColorManager"
}

// method GetDevices

func (v *interfaceManager) GoGetDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDevices", flags, ch)
}

func (*interfaceManager) StoreGetDevices(call *dbus.Call) (devices []dbus.ObjectPath, err error) {
	err = call.Store(&devices)
	return
}

func (v *interfaceManager) GetDevices(flags dbus.Flags) ([]dbus.ObjectPath, error) {
	return v.StoreGetDevices(
		<-v.GoGetDevices(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetDevicesByKind

func (v *interfaceManager) GoGetDevicesByKind(flags dbus.Flags, ch chan *dbus.Call, kind string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDevicesByKind", flags, ch, kind)
}

func (*interfaceManager) StoreGetDevicesByKind(call *dbus.Call) (devices []dbus.ObjectPath, err error) {
	err = call.Store(&devices)
	return
}

func (v *interfaceManager) GetDevicesByKind(flags dbus.Flags, kind string) ([]dbus.ObjectPath, error) {
	return v.StoreGetDevicesByKind(
		<-v.GoGetDevicesByKind(flags, make(chan *dbus.Call, 1), kind).Done)
}

// method FindDeviceById

func (v *interfaceManager) GoFindDeviceById(flags dbus.Flags, ch chan *dbus.Call, device_id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FindDeviceById", flags, ch, device_id)
}

func (*interfaceManager) StoreFindDeviceById(call *dbus.Call) (object_path dbus.ObjectPath, err error) {
	err = call.Store(&object_path)
	return
}

func (v *interfaceManager) FindDeviceById(flags dbus.Flags, device_id string) (dbus.ObjectPath, error) {
	return v.StoreFindDeviceById(
		<-v.GoFindDeviceById(flags, make(chan *dbus.Call, 1), device_id).Done)
}

// method FindSensorById

func (v *interfaceManager) GoFindSensorById(flags dbus.Flags, ch chan *dbus.Call, sensor_id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FindSensorById", flags, ch, sensor_id)
}

func (*interfaceManager) StoreFindSensorById(call *dbus.Call) (object_path dbus.ObjectPath, err error) {
	err = call.Store(&object_path)
	return
}

func (v *interfaceManager) FindSensorById(flags dbus.Flags, sensor_id string) (dbus.ObjectPath, error) {
	return v.StoreFindSensorById(
		<-v.GoFindSensorById(flags, make(chan *dbus.Call, 1), sensor_id).Done)
}

// method FindDeviceByProperty

func (v *interfaceManager) GoFindDeviceByProperty(flags dbus.Flags, ch chan *dbus.Call, key string, value string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FindDeviceByProperty", flags, ch, key, value)
}

func (*interfaceManager) StoreFindDeviceByProperty(call *dbus.Call) (object_path dbus.ObjectPath, err error) {
	err = call.Store(&object_path)
	return
}

func (v *interfaceManager) FindDeviceByProperty(flags dbus.Flags, key string, value string) (dbus.ObjectPath, error) {
	return v.StoreFindDeviceByProperty(
		<-v.GoFindDeviceByProperty(flags, make(chan *dbus.Call, 1), key, value).Done)
}

// method FindProfileById

func (v *interfaceManager) GoFindProfileById(flags dbus.Flags, ch chan *dbus.Call, profile_id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FindProfileById", flags, ch, profile_id)
}

func (*interfaceManager) StoreFindProfileById(call *dbus.Call) (object_path dbus.ObjectPath, err error) {
	err = call.Store(&object_path)
	return
}

func (v *interfaceManager) FindProfileById(flags dbus.Flags, profile_id string) (dbus.ObjectPath, error) {
	return v.StoreFindProfileById(
		<-v.GoFindProfileById(flags, make(chan *dbus.Call, 1), profile_id).Done)
}

// method FindProfileByProperty

func (v *interfaceManager) GoFindProfileByProperty(flags dbus.Flags, ch chan *dbus.Call, key string, value string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FindProfileByProperty", flags, ch, key, value)
}

func (*interfaceManager) StoreFindProfileByProperty(call *dbus.Call) (object_path dbus.ObjectPath, err error) {
	err = call.Store(&object_path)
	return
}

func (v *interfaceManager) FindProfileByProperty(flags dbus.Flags, key string, value string) (dbus.ObjectPath, error) {
	return v.StoreFindProfileByProperty(
		<-v.GoFindProfileByProperty(flags, make(chan *dbus.Call, 1), key, value).Done)
}

// method FindProfileByFilename

func (v *interfaceManager) GoFindProfileByFilename(flags dbus.Flags, ch chan *dbus.Call, filename string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FindProfileByFilename", flags, ch, filename)
}

func (*interfaceManager) StoreFindProfileByFilename(call *dbus.Call) (object_path dbus.ObjectPath, err error) {
	err = call.Store(&object_path)
	return
}

func (v *interfaceManager) FindProfileByFilename(flags dbus.Flags, filename string) (dbus.ObjectPath, error) {
	return v.StoreFindProfileByFilename(
		<-v.GoFindProfileByFilename(flags, make(chan *dbus.Call, 1), filename).Done)
}

// method GetStandardSpace

func (v *interfaceManager) GoGetStandardSpace(flags dbus.Flags, ch chan *dbus.Call, standard_space string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetStandardSpace", flags, ch, standard_space)
}

func (*interfaceManager) StoreGetStandardSpace(call *dbus.Call) (object_path dbus.ObjectPath, err error) {
	err = call.Store(&object_path)
	return
}

func (v *interfaceManager) GetStandardSpace(flags dbus.Flags, standard_space string) (dbus.ObjectPath, error) {
	return v.StoreGetStandardSpace(
		<-v.GoGetStandardSpace(flags, make(chan *dbus.Call, 1), standard_space).Done)
}

// method GetProfiles

func (v *interfaceManager) GoGetProfiles(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetProfiles", flags, ch)
}

func (*interfaceManager) StoreGetProfiles(call *dbus.Call) (devices []dbus.ObjectPath, err error) {
	err = call.Store(&devices)
	return
}

func (v *interfaceManager) GetProfiles(flags dbus.Flags) ([]dbus.ObjectPath, error) {
	return v.StoreGetProfiles(
		<-v.GoGetProfiles(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetSensors

func (v *interfaceManager) GoGetSensors(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSensors", flags, ch)
}

func (*interfaceManager) StoreGetSensors(call *dbus.Call) (devices []dbus.ObjectPath, err error) {
	err = call.Store(&devices)
	return
}

func (v *interfaceManager) GetSensors(flags dbus.Flags) ([]dbus.ObjectPath, error) {
	return v.StoreGetSensors(
		<-v.GoGetSensors(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetProfilesByKind

func (v *interfaceManager) GoGetProfilesByKind(flags dbus.Flags, ch chan *dbus.Call, kind string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetProfilesByKind", flags, ch, kind)
}

func (*interfaceManager) StoreGetProfilesByKind(call *dbus.Call) (devices []dbus.ObjectPath, err error) {
	err = call.Store(&devices)
	return
}

func (v *interfaceManager) GetProfilesByKind(flags dbus.Flags, kind string) ([]dbus.ObjectPath, error) {
	return v.StoreGetProfilesByKind(
		<-v.GoGetProfilesByKind(flags, make(chan *dbus.Call, 1), kind).Done)
}

// method CreateProfileWithFd

func (v *interfaceManager) GoCreateProfileWithFd(flags dbus.Flags, ch chan *dbus.Call, profile_id string, scope string, handle dbus.UnixFD, properties map[string]string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateProfileWithFd", flags, ch, profile_id, scope, handle, properties)
}

func (*interfaceManager) StoreCreateProfileWithFd(call *dbus.Call) (object_path dbus.ObjectPath, err error) {
	err = call.Store(&object_path)
	return
}

func (v *interfaceManager) CreateProfileWithFd(flags dbus.Flags, profile_id string, scope string, handle dbus.UnixFD, properties map[string]string) (dbus.ObjectPath, error) {
	return v.StoreCreateProfileWithFd(
		<-v.GoCreateProfileWithFd(flags, make(chan *dbus.Call, 1), profile_id, scope, handle, properties).Done)
}

// method CreateProfile

func (v *interfaceManager) GoCreateProfile(flags dbus.Flags, ch chan *dbus.Call, profile_id string, scope string, properties map[string]string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateProfile", flags, ch, profile_id, scope, properties)
}

func (*interfaceManager) StoreCreateProfile(call *dbus.Call) (object_path dbus.ObjectPath, err error) {
	err = call.Store(&object_path)
	return
}

func (v *interfaceManager) CreateProfile(flags dbus.Flags, profile_id string, scope string, properties map[string]string) (dbus.ObjectPath, error) {
	return v.StoreCreateProfile(
		<-v.GoCreateProfile(flags, make(chan *dbus.Call, 1), profile_id, scope, properties).Done)
}

// method CreateDevice

func (v *interfaceManager) GoCreateDevice(flags dbus.Flags, ch chan *dbus.Call, device_id string, scope string, properties map[string]string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateDevice", flags, ch, device_id, scope, properties)
}

func (*interfaceManager) StoreCreateDevice(call *dbus.Call) (object_path dbus.ObjectPath, err error) {
	err = call.Store(&object_path)
	return
}

func (v *interfaceManager) CreateDevice(flags dbus.Flags, device_id string, scope string, properties map[string]string) (dbus.ObjectPath, error) {
	return v.StoreCreateDevice(
		<-v.GoCreateDevice(flags, make(chan *dbus.Call, 1), device_id, scope, properties).Done)
}

// method DeleteDevice

func (v *interfaceManager) GoDeleteDevice(flags dbus.Flags, ch chan *dbus.Call, object_path dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteDevice", flags, ch, object_path)
}

func (v *interfaceManager) DeleteDevice(flags dbus.Flags, object_path dbus.ObjectPath) error {
	return (<-v.GoDeleteDevice(flags, make(chan *dbus.Call, 1), object_path).Done).Err
}

// method DeleteProfile

func (v *interfaceManager) GoDeleteProfile(flags dbus.Flags, ch chan *dbus.Call, object_path dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteProfile", flags, ch, object_path)
}

func (v *interfaceManager) DeleteProfile(flags dbus.Flags, object_path dbus.ObjectPath) error {
	return (<-v.GoDeleteProfile(flags, make(chan *dbus.Call, 1), object_path).Done).Err
}

// signal Changed

func (v *interfaceManager) ConnectChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Changed", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Changed",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal DeviceAdded

func (v *interfaceManager) ConnectDeviceAdded(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "DeviceAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".DeviceAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var object_path dbus.ObjectPath
		err := dbus.Store(sig.Body, &object_path)
		if err == nil {
			cb(object_path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal DeviceRemoved

func (v *interfaceManager) ConnectDeviceRemoved(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "DeviceRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".DeviceRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var object_path dbus.ObjectPath
		err := dbus.Store(sig.Body, &object_path)
		if err == nil {
			cb(object_path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal DeviceChanged

func (v *interfaceManager) ConnectDeviceChanged(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "DeviceChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".DeviceChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var object_path dbus.ObjectPath
		err := dbus.Store(sig.Body, &object_path)
		if err == nil {
			cb(object_path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ProfileAdded

func (v *interfaceManager) ConnectProfileAdded(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ProfileAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ProfileAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var object_path dbus.ObjectPath
		err := dbus.Store(sig.Body, &object_path)
		if err == nil {
			cb(object_path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ProfileRemoved

func (v *interfaceManager) ConnectProfileRemoved(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ProfileRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ProfileRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var object_path dbus.ObjectPath
		err := dbus.Store(sig.Body, &object_path)
		if err == nil {
			cb(object_path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal SensorAdded

func (v *interfaceManager) ConnectSensorAdded(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SensorAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SensorAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var object_path dbus.ObjectPath
		err := dbus.Store(sig.Body, &object_path)
		if err == nil {
			cb(object_path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal SensorRemoved

func (v *interfaceManager) ConnectSensorRemoved(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SensorRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SensorRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var object_path dbus.ObjectPath
		err := dbus.Store(sig.Body, &object_path)
		if err == nil {
			cb(object_path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ProfileChanged

func (v *interfaceManager) ConnectProfileChanged(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ProfileChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ProfileChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var object_path dbus.ObjectPath
		err := dbus.Store(sig.Body, &object_path)
		if err == nil {
			cb(object_path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property DaemonVersion s

func (v *interfaceManager) DaemonVersion() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "DaemonVersion",
	}
}

// property SystemVendor s

func (v *interfaceManager) SystemVendor() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "SystemVendor",
	}
}

// property SystemModel s

func (v *interfaceManager) SystemModel() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "SystemModel",
	}
}

type Profile interface {
	profile // interface org.freedesktop.ColorManager.Profile
	proxy.Object
}

type objectProfile struct {
	interfaceProfile // interface org.freedesktop.ColorManager.Profile
	proxy.ImplObject
}

func NewProfile(conn *dbus.Conn, path dbus.ObjectPath) (Profile, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectProfile)
	obj.ImplObject.Init_(conn, "org.freedesktop.ColorManager", path)
	return obj, nil
}

type profile interface {
	GoSetProperty(flags dbus.Flags, ch chan *dbus.Call, property_name string, property_value string) *dbus.Call
	SetProperty(flags dbus.Flags, property_name string, property_value string) error
	GoInstallSystemWide(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	InstallSystemWide(flags dbus.Flags) error
	ConnectChanged(cb func()) (dbusutil.SignalHandlerId, error)
	ProfileId() proxy.PropString
	Title() proxy.PropString
	Metadata() PropProfileMetadata
	Qualifier() proxy.PropString
	Format() proxy.PropString
	Kind() proxy.PropString
	Colorspace() proxy.PropString
	HasVcgt() proxy.PropBool
	IsSystemWide() proxy.PropBool
	Filename() proxy.PropString
	Created() proxy.PropInt64
	Scope() proxy.PropString
	Owner() proxy.PropUint32
	Warnings() proxy.PropStringArray
}

type interfaceProfile struct{}

func (v *interfaceProfile) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceProfile) GetInterfaceName_() string {
	return "org.freedesktop.ColorManager.Profile"
}

// method SetProperty

func (v *interfaceProfile) GoSetProperty(flags dbus.Flags, ch chan *dbus.Call, property_name string, property_value string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetProperty", flags, ch, property_name, property_value)
}

func (v *interfaceProfile) SetProperty(flags dbus.Flags, property_name string, property_value string) error {
	return (<-v.GoSetProperty(flags, make(chan *dbus.Call, 1), property_name, property_value).Done).Err
}

// method InstallSystemWide

func (v *interfaceProfile) GoInstallSystemWide(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".InstallSystemWide", flags, ch)
}

func (v *interfaceProfile) InstallSystemWide(flags dbus.Flags) error {
	return (<-v.GoInstallSystemWide(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal Changed

func (v *interfaceProfile) ConnectChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Changed", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Changed",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property ProfileId s

func (v *interfaceProfile) ProfileId() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "ProfileId",
	}
}

// property Title s

func (v *interfaceProfile) Title() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Title",
	}
}

type PropProfileMetadata interface {
	Get(flags dbus.Flags) (value map[string]string, err error)
	Set(flags dbus.Flags, value map[string]string) error
	ConnectChanged(cb func(hasValue bool, value map[string]string)) error
}

type implPropProfileMetadata struct {
	Impl proxy.Implementer
	Name string
}

func (p implPropProfileMetadata) Get(flags dbus.Flags) (value map[string]string, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p implPropProfileMetadata) Set(flags dbus.Flags, value map[string]string) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p implPropProfileMetadata) ConnectChanged(cb func(hasValue bool, value map[string]string)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v map[string]string
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}

// property Metadata a{ss}

func (v *interfaceProfile) Metadata() PropProfileMetadata {
	return &implPropProfileMetadata{
		Impl: v,
		Name: "Metadata",
	}
}

// property Qualifier s

func (v *interfaceProfile) Qualifier() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Qualifier",
	}
}

// property Format s

func (v *interfaceProfile) Format() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Format",
	}
}

// property Kind s

func (v *interfaceProfile) Kind() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Kind",
	}
}

// property Colorspace s

func (v *interfaceProfile) Colorspace() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Colorspace",
	}
}

// property HasVcgt b

func (v *interfaceProfile) HasVcgt() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "HasVcgt",
	}
}

// property IsSystemWide b

func (v *interfaceProfile) IsSystemWide() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "IsSystemWide",
	}
}

// property Filename s

func (v *interfaceProfile) Filename() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Filename",
	}
}

// property Created x

func (v *interfaceProfile) Created() proxy.PropInt64 {
	return &proxy.ImplPropInt64{
		Impl: v,
		Name: "Created",
	}
}

// property Scope s

func (v *interfaceProfile) Scope() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Scope",
	}
}

// property Owner u

func (v *interfaceProfile) Owner() proxy.PropUint32 {
	return &proxy.ImplPropUint32{
		Impl: v,
		Name: "Owner",
	}
}

// property Warnings as

func (v *interfaceProfile) Warnings() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "Warnings",
	}
}
