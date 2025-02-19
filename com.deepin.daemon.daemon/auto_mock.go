// Code generated by "./generator ./com.deepin.daemon.daemon"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package daemon

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
	"github.com/stretchr/testify/mock"
)

type MockDaemon struct {
	MockInterfaceDaemon // interface com.deepin.daemon.Daemon
	proxy.MockObject
}

type MockInterfaceDaemon struct {
	mock.Mock
}

// method BluetoothGetDeviceTechnologies

func (v *MockInterfaceDaemon) GoBluetoothGetDeviceTechnologies(flags dbus.Flags, ch chan *dbus.Call, adapter string, device string) *dbus.Call {
	mockArgs := v.Called(flags, ch, adapter, device)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDaemon) BluetoothGetDeviceTechnologies(flags dbus.Flags, adapter string, device string) ([]string, error) {
	mockArgs := v.Called(flags, adapter, device)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method IsPidVirtualMachine

func (v *MockInterfaceDaemon) GoIsPidVirtualMachine(flags dbus.Flags, ch chan *dbus.Call, pid uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, pid)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDaemon) IsPidVirtualMachine(flags dbus.Flags, pid uint32) (bool, error) {
	mockArgs := v.Called(flags, pid)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method ClearTtys

func (v *MockInterfaceDaemon) GoClearTtys(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDaemon) ClearTtys(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method ClearTty

func (v *MockInterfaceDaemon) GoClearTty(flags dbus.Flags, ch chan *dbus.Call, num uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, num)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDaemon) ClearTty(flags dbus.Flags, num uint32) error {
	mockArgs := v.Called(flags, num)

	return mockArgs.Error(0)
}

// method NetworkGetConnections

func (v *MockInterfaceDaemon) GoNetworkGetConnections(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDaemon) NetworkGetConnections(flags dbus.Flags) ([]uint8, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]uint8)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method NetworkSetConnections

func (v *MockInterfaceDaemon) GoNetworkSetConnections(flags dbus.Flags, ch chan *dbus.Call, data []uint8) *dbus.Call {
	mockArgs := v.Called(flags, ch, data)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDaemon) NetworkSetConnections(flags dbus.Flags, data []uint8) error {
	mockArgs := v.Called(flags, data)

	return mockArgs.Error(0)
}

// method ScalePlymouth

func (v *MockInterfaceDaemon) GoScalePlymouth(flags dbus.Flags, ch chan *dbus.Call, scale uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, scale)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDaemon) ScalePlymouth(flags dbus.Flags, scale uint32) error {
	mockArgs := v.Called(flags, scale)

	return mockArgs.Error(0)
}

// method SetLongPressDuration

func (v *MockInterfaceDaemon) GoSetLongPressDuration(flags dbus.Flags, ch chan *dbus.Call, duration uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, duration)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDaemon) SetLongPressDuration(flags dbus.Flags, duration uint32) error {
	mockArgs := v.Called(flags, duration)

	return mockArgs.Error(0)
}

// method SaveCustomWallPaper

func (v *MockInterfaceDaemon) GoSaveCustomWallPaper(flags dbus.Flags, ch chan *dbus.Call, username string, file string) *dbus.Call {
	mockArgs := v.Called(flags, ch, username, file)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDaemon) SaveCustomWallPaper(flags dbus.Flags, username string, file string) (string, error) {
	mockArgs := v.Called(flags, username, file)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method GetCustomWallPapers

func (v *MockInterfaceDaemon) GoGetCustomWallPapers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	mockArgs := v.Called(flags, ch, username)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDaemon) GetCustomWallPapers(flags dbus.Flags, username string) ([]string, error) {
	mockArgs := v.Called(flags, username)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method DeleteCustomWallPaper

func (v *MockInterfaceDaemon) GoDeleteCustomWallPaper(flags dbus.Flags, ch chan *dbus.Call, username string, file string) *dbus.Call {
	mockArgs := v.Called(flags, ch, username, file)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDaemon) DeleteCustomWallPaper(flags dbus.Flags, username string, file string) error {
	mockArgs := v.Called(flags, username, file)

	return mockArgs.Error(0)
}

// signal HandleForSleep

func (v *MockInterfaceDaemon) ConnectHandleForSleep(cb func(start bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}
