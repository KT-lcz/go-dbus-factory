// Code generated by "./generator ./com.deepin.daemon.uadp"; DO NOT EDIT.

package uadp

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/stretchr/testify/mock"
)

type MockUadp struct {
	mockInterfaceUadp // interface com.deepin.daemon.Uadp
}

type mockInterfaceUadp struct {
	mock.Mock
}

// method SetDataKey

func (v *mockInterfaceUadp) GoSetDataKey(flags dbus.Flags, ch chan *dbus.Call, exePath string, keyName string, dataKey string, keyringKey string) *dbus.Call {
	mockArgs := v.Called(flags, ch, exePath, keyName, dataKey, keyringKey)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUadp) SetDataKey(flags dbus.Flags, exePath string, keyName string, dataKey string, keyringKey string) error {
	mockArgs := v.Called(flags, exePath, keyName, dataKey, keyringKey)

	return mockArgs.Error(0)
}

// method GetDataKey

func (v *mockInterfaceUadp) GoGetDataKey(flags dbus.Flags, ch chan *dbus.Call, exePath string, keyName string, keyringKey string) *dbus.Call {
	mockArgs := v.Called(flags, ch, exePath, keyName, keyringKey)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUadp) GetDataKey(flags dbus.Flags, exePath string, keyName string, keyringKey string) (string, error) {
	mockArgs := v.Called(flags, exePath, keyName, keyringKey)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}
