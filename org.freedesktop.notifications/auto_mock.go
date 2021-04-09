// Code generated by "./generator ./org.freedesktop.notifications"; DO NOT EDIT.

package notifications

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/stretchr/testify/mock"
	"pkg.deepin.io/lib/dbusutil"
)

type MockNotifications struct {
	mockInterfaceNotifications // interface org.freedesktop.Notifications
}

type mockInterfaceNotifications struct {
	mock.Mock
}

// method GetCapabilities

func (v *mockInterfaceNotifications) GoGetCapabilities(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNotifications) GetCapabilities(flags dbus.Flags) ([]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method Notify

func (v *mockInterfaceNotifications) GoNotify(flags dbus.Flags, ch chan *dbus.Call, app_name string, replaces_id uint32, app_icon string, summary string, body string, actions []string, hints map[string]dbus.Variant, expire_timeout int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, app_name, replaces_id, app_icon, summary, body, actions, hints, expire_timeout)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNotifications) Notify(flags dbus.Flags, app_name string, replaces_id uint32, app_icon string, summary string, body string, actions []string, hints map[string]dbus.Variant, expire_timeout int32) (uint32, error) {
	mockArgs := v.Called(flags, app_name, replaces_id, app_icon, summary, body, actions, hints, expire_timeout)

	ret0, ok := mockArgs.Get(0).(uint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method CloseNotification

func (v *mockInterfaceNotifications) GoCloseNotification(flags dbus.Flags, ch chan *dbus.Call, id uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNotifications) CloseNotification(flags dbus.Flags, id uint32) error {
	mockArgs := v.Called(flags, id)

	return mockArgs.Error(0)
}

// method GetServerInformation

func (v *mockInterfaceNotifications) GoGetServerInformation(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNotifications) GetServerInformation(flags dbus.Flags) (string, string, string, string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	ret1, ok := mockArgs.Get(1).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 1, mockArgs.Get(1)))
	}

	ret2, ok := mockArgs.Get(2).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 2, mockArgs.Get(2)))
	}

	ret3, ok := mockArgs.Get(3).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 3, mockArgs.Get(3)))
	}

	return ret0, ret1, ret2, ret3, mockArgs.Error(4)
}

// signal NotificationClosed

func (v *mockInterfaceNotifications) ConnectNotificationClosed(cb func(id uint32, reason uint32)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal ActionInvoked

func (v *mockInterfaceNotifications) ConnectActionInvoked(cb func(id uint32, action_key string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}
