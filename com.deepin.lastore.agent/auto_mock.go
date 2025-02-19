// Code generated by "./generator com.deepin.lastore.agent"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package agent

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
	"github.com/stretchr/testify/mock"
)

type MockAgent struct {
	MockInterfaceAgent // interface com.deepin.lastore.Agent
	proxy.MockObject
}

type MockInterfaceAgent struct {
	mock.Mock
}

// method GetManualProxy

func (v *MockInterfaceAgent) GoGetManualProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceAgent) GetManualProxy(flags dbus.Flags) (map[string]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(map[string]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method SendNotify

func (v *MockInterfaceAgent) GoSendNotify(flags dbus.Flags, ch chan *dbus.Call, app_name string, replaces_id uint32, app_icon string, summary string, body string, actions []string, hints map[string]dbus.Variant, expire_timeout int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, app_name, replaces_id, app_icon, summary, body, actions, hints, expire_timeout)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceAgent) SendNotify(flags dbus.Flags, app_name string, replaces_id uint32, app_icon string, summary string, body string, actions []string, hints map[string]dbus.Variant, expire_timeout int32) (uint32, error) {
	mockArgs := v.Called(flags, app_name, replaces_id, app_icon, summary, body, actions, hints, expire_timeout)

	ret0, ok := mockArgs.Get(0).(uint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method CloseNotification

func (v *MockInterfaceAgent) GoCloseNotification(flags dbus.Flags, ch chan *dbus.Call, id uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceAgent) CloseNotification(flags dbus.Flags, id uint32) error {
	mockArgs := v.Called(flags, id)

	return mockArgs.Error(0)
}
