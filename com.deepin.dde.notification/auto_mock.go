// Code generated by "./generator ./com.deepin.dde.notification"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package notification

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
	"github.com/stretchr/testify/mock"
)

type MockNotification struct {
	MockInterfaceNotification // interface com.deepin.dde.Notification
	proxy.MockObject
}

type MockInterfaceNotification struct {
	mock.Mock
}

// method CloseNotification

func (v *MockInterfaceNotification) GoCloseNotification(flags dbus.Flags, ch chan *dbus.Call, arg0 uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotification) CloseNotification(flags dbus.Flags, arg0 uint32) error {
	mockArgs := v.Called(flags, arg0)

	return mockArgs.Error(0)
}

// method GetCapbilities

func (v *MockInterfaceNotification) GoGetCapbilities(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotification) GetCapbilities(flags dbus.Flags) ([]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetServerInformation

func (v *MockInterfaceNotification) GoGetServerInformation(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotification) GetServerInformation(flags dbus.Flags) (string, string, string, string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.String(1), mockArgs.String(2), mockArgs.String(3), mockArgs.Error(4)
}

// method Notify

func (v *MockInterfaceNotification) GoNotify(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32, arg2 string, arg3 string, arg4 string, arg5 []string, arg6 map[string]dbus.Variant, arg7 int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotification) Notify(flags dbus.Flags, arg0 string, arg1 uint32, arg2 string, arg3 string, arg4 string, arg5 []string, arg6 map[string]dbus.Variant, arg7 int32) (uint32, error) {
	mockArgs := v.Called(flags, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)

	ret0, ok := mockArgs.Get(0).(uint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetAllRecords

func (v *MockInterfaceNotification) GoGetAllRecords(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotification) GetAllRecords(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method GetRecordById

func (v *MockInterfaceNotification) GoGetRecordById(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotification) GetRecordById(flags dbus.Flags, arg0 string) (string, error) {
	mockArgs := v.Called(flags, arg0)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method GetRecordsFromId

func (v *MockInterfaceNotification) GoGetRecordsFromId(flags dbus.Flags, ch chan *dbus.Call, arg0 int32, arg1 string) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotification) GetRecordsFromId(flags dbus.Flags, arg0 int32, arg1 string) (string, error) {
	mockArgs := v.Called(flags, arg0, arg1)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method RemoveRecord

func (v *MockInterfaceNotification) GoRemoveRecord(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotification) RemoveRecord(flags dbus.Flags, arg0 string) error {
	mockArgs := v.Called(flags, arg0)

	return mockArgs.Error(0)
}

// method ClearRecords

func (v *MockInterfaceNotification) GoClearRecords(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotification) ClearRecords(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method getAppSetting

func (v *MockInterfaceNotification) GoGetAppSetting(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotification) GetAppSetting(flags dbus.Flags, arg0 string) (string, error) {
	mockArgs := v.Called(flags, arg0)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method Toggle

func (v *MockInterfaceNotification) GoToggle(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotification) Toggle(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method Show

func (v *MockInterfaceNotification) GoShow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotification) Show(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method Hide

func (v *MockInterfaceNotification) GoHide(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotification) Hide(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method recordCount

func (v *MockInterfaceNotification) GoRecordCount(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotification) RecordCount(flags dbus.Flags) (uint32, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(uint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetAppList

func (v *MockInterfaceNotification) GoGetAppList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotification) GetAppList(flags dbus.Flags) ([]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetAppInfo

func (v *MockInterfaceNotification) GoGetAppInfo(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotification) GetAppInfo(flags dbus.Flags, arg0 string, arg1 uint32) (dbus.Variant, error) {
	mockArgs := v.Called(flags, arg0, arg1)

	ret0, ok := mockArgs.Get(0).(dbus.Variant)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetSystemInfo

func (v *MockInterfaceNotification) GoGetSystemInfo(flags dbus.Flags, ch chan *dbus.Call, arg0 uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotification) GetSystemInfo(flags dbus.Flags, arg0 uint32) (dbus.Variant, error) {
	mockArgs := v.Called(flags, arg0)

	ret0, ok := mockArgs.Get(0).(dbus.Variant)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method SetAppInfo

func (v *MockInterfaceNotification) GoSetAppInfo(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32, arg2 dbus.Variant) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1, arg2)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotification) SetAppInfo(flags dbus.Flags, arg0 string, arg1 uint32, arg2 dbus.Variant) error {
	mockArgs := v.Called(flags, arg0, arg1, arg2)

	return mockArgs.Error(0)
}

// method SetSystemInfo

func (v *MockInterfaceNotification) GoSetSystemInfo(flags dbus.Flags, ch chan *dbus.Call, arg0 uint32, arg1 dbus.Variant) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotification) SetSystemInfo(flags dbus.Flags, arg0 uint32, arg1 dbus.Variant) error {
	mockArgs := v.Called(flags, arg0, arg1)

	return mockArgs.Error(0)
}

// method setAppSetting

func (v *MockInterfaceNotification) GoSetAppSetting(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotification) SetAppSetting(flags dbus.Flags, arg0 string) error {
	mockArgs := v.Called(flags, arg0)

	return mockArgs.Error(0)
}

// signal NotificationClosed

func (v *MockInterfaceNotification) ConnectNotificationClosed(cb func(arg0 uint32, arg1 uint32)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal ActionInvoked

func (v *MockInterfaceNotification) ConnectActionInvoked(cb func(arg0 uint32, arg1 string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal RecordAdded

func (v *MockInterfaceNotification) ConnectRecordAdded(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal AppInfoChanged

func (v *MockInterfaceNotification) ConnectAppInfoChanged(cb func(arg0 string, arg1 uint32, arg2 dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal SystemInfoChanged

func (v *MockInterfaceNotification) ConnectSystemInfoChanged(cb func(arg0 uint32, arg1 dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal AppAddedSignal

func (v *MockInterfaceNotification) ConnectAppAddedSignal(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal AppRemovedSignal

func (v *MockInterfaceNotification) ConnectAppRemovedSignal(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal appRemoved

func (v *MockInterfaceNotification) ConnectAppRemoved(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal appAdded

func (v *MockInterfaceNotification) ConnectAppAdded(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal appSettingChanged

func (v *MockInterfaceNotification) ConnectAppSettingChanged(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal systemSettingChanged

func (v *MockInterfaceNotification) ConnectSystemSettingChanged(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property allSetting s

func (v *MockInterfaceNotification) AllSetting() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property systemSetting s

func (v *MockInterfaceNotification) SystemSetting() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
