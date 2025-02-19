// Code generated by "./generator ./com.deepin.daemon.authenticate.fingerprint"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package fingerprint

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
	"github.com/stretchr/testify/mock"
)

type MockDevice struct {
	MockInterfaceDevice // interface com.deepin.daemon.Authenticate.Fingerprint.Device
	proxy.MockObject
}

type MockInterfaceDevice struct {
	mock.Mock
}

func (v *MockInterfaceDevice) SetInterfaceName_(string) {
}

// method Claim

func (v *MockInterfaceDevice) GoClaim(flags dbus.Flags, ch chan *dbus.Call, userId string, claimed bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, userId, claimed)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) Claim(flags dbus.Flags, userId string, claimed bool) error {
	mockArgs := v.Called(flags, userId, claimed)

	return mockArgs.Error(0)
}

// method DeleteAllFingers

func (v *MockInterfaceDevice) GoDeleteAllFingers(flags dbus.Flags, ch chan *dbus.Call, userId string) *dbus.Call {
	mockArgs := v.Called(flags, ch, userId)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) DeleteAllFingers(flags dbus.Flags, userId string) error {
	mockArgs := v.Called(flags, userId)

	return mockArgs.Error(0)
}

// method DeleteFinger

func (v *MockInterfaceDevice) GoDeleteFinger(flags dbus.Flags, ch chan *dbus.Call, userId string, finger string) *dbus.Call {
	mockArgs := v.Called(flags, ch, userId, finger)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) DeleteFinger(flags dbus.Flags, userId string, finger string) error {
	mockArgs := v.Called(flags, userId, finger)

	return mockArgs.Error(0)
}

// method Enroll

func (v *MockInterfaceDevice) GoEnroll(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	mockArgs := v.Called(flags, ch, finger)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) Enroll(flags dbus.Flags, finger string) error {
	mockArgs := v.Called(flags, finger)

	return mockArgs.Error(0)
}

// method ListFingers

func (v *MockInterfaceDevice) GoListFingers(flags dbus.Flags, ch chan *dbus.Call, userId string) *dbus.Call {
	mockArgs := v.Called(flags, ch, userId)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) ListFingers(flags dbus.Flags, userId string) ([]string, error) {
	mockArgs := v.Called(flags, userId)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method RenameFinger

func (v *MockInterfaceDevice) GoRenameFinger(flags dbus.Flags, ch chan *dbus.Call, userId string, finger string, newName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, userId, finger, newName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) RenameFinger(flags dbus.Flags, userId string, finger string, newName string) error {
	mockArgs := v.Called(flags, userId, finger, newName)

	return mockArgs.Error(0)
}

// method StopEnroll

func (v *MockInterfaceDevice) GoStopEnroll(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) StopEnroll(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method StopVerify

func (v *MockInterfaceDevice) GoStopVerify(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) StopVerify(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method Verify

func (v *MockInterfaceDevice) GoVerify(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	mockArgs := v.Called(flags, ch, finger)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDevice) Verify(flags dbus.Flags, finger string) error {
	mockArgs := v.Called(flags, finger)

	return mockArgs.Error(0)
}

// signal EnrollStatus

func (v *MockInterfaceDevice) ConnectEnrollStatus(cb func(userId string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal VerifyStatus

func (v *MockInterfaceDevice) ConnectVerifyStatus(cb func(userId string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal Touch

func (v *MockInterfaceDevice) ConnectTouch(cb func(userId string, pressed bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property Name s

func (v *MockInterfaceDevice) Name() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property State i

func (v *MockInterfaceDevice) State() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Type i

func (v *MockInterfaceDevice) Type() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Capability i

func (v *MockInterfaceDevice) Capability() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockCommonDevice struct {
	MockInterfaceCommonDevice // interface com.deepin.daemon.Authenticate.Fingerprint.CommonDevice
	proxy.MockObject
}

type MockInterfaceCommonDevice struct {
	mock.Mock
}

func (v *MockInterfaceCommonDevice) SetInterfaceName_(string) {
}

// method Claim

func (v *MockInterfaceCommonDevice) GoClaim(flags dbus.Flags, ch chan *dbus.Call, userId string, claimed bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, userId, claimed)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceCommonDevice) Claim(flags dbus.Flags, userId string, claimed bool) error {
	mockArgs := v.Called(flags, userId, claimed)

	return mockArgs.Error(0)
}

// method DeleteAllFingers

func (v *MockInterfaceCommonDevice) GoDeleteAllFingers(flags dbus.Flags, ch chan *dbus.Call, userId string) *dbus.Call {
	mockArgs := v.Called(flags, ch, userId)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceCommonDevice) DeleteAllFingers(flags dbus.Flags, userId string) error {
	mockArgs := v.Called(flags, userId)

	return mockArgs.Error(0)
}

// method DeleteFinger

func (v *MockInterfaceCommonDevice) GoDeleteFinger(flags dbus.Flags, ch chan *dbus.Call, userId string, finger string) *dbus.Call {
	mockArgs := v.Called(flags, ch, userId, finger)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceCommonDevice) DeleteFinger(flags dbus.Flags, userId string, finger string) error {
	mockArgs := v.Called(flags, userId, finger)

	return mockArgs.Error(0)
}

// method Enroll

func (v *MockInterfaceCommonDevice) GoEnroll(flags dbus.Flags, ch chan *dbus.Call, username string, finger string) *dbus.Call {
	mockArgs := v.Called(flags, ch, username, finger)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceCommonDevice) Enroll(flags dbus.Flags, username string, finger string) error {
	mockArgs := v.Called(flags, username, finger)

	return mockArgs.Error(0)
}

// method ListFingers

func (v *MockInterfaceCommonDevice) GoListFingers(flags dbus.Flags, ch chan *dbus.Call, userId string) *dbus.Call {
	mockArgs := v.Called(flags, ch, userId)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceCommonDevice) ListFingers(flags dbus.Flags, userId string) ([]string, error) {
	mockArgs := v.Called(flags, userId)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method RenameFinger

func (v *MockInterfaceCommonDevice) GoRenameFinger(flags dbus.Flags, ch chan *dbus.Call, userId string, finger string, newName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, userId, finger, newName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceCommonDevice) RenameFinger(flags dbus.Flags, userId string, finger string, newName string) error {
	mockArgs := v.Called(flags, userId, finger, newName)

	return mockArgs.Error(0)
}

// method StopEnroll

func (v *MockInterfaceCommonDevice) GoStopEnroll(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceCommonDevice) StopEnroll(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method StopVerify

func (v *MockInterfaceCommonDevice) GoStopVerify(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceCommonDevice) StopVerify(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method Verify

func (v *MockInterfaceCommonDevice) GoVerify(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	mockArgs := v.Called(flags, ch, finger)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceCommonDevice) Verify(flags dbus.Flags, finger string) error {
	mockArgs := v.Called(flags, finger)

	return mockArgs.Error(0)
}

// signal EnrollStatus

func (v *MockInterfaceCommonDevice) ConnectEnrollStatus(cb func(userId string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal VerifyStatus

func (v *MockInterfaceCommonDevice) ConnectVerifyStatus(cb func(userId string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal Touch

func (v *MockInterfaceCommonDevice) ConnectTouch(cb func(userId string, pressed bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property Name s

func (v *MockInterfaceCommonDevice) Name() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property State i

func (v *MockInterfaceCommonDevice) State() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Type i

func (v *MockInterfaceCommonDevice) Type() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Capability i

func (v *MockInterfaceCommonDevice) Capability() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
