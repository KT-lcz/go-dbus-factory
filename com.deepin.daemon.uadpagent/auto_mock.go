// Code generated by "./generator ./com.deepin.daemon.uadpagent"; DO NOT EDIT.

package uadpagent

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/stretchr/testify/mock"
)

type MockUadpAgent struct {
	mockInterfaceUadpagent // interface com.deepin.daemon.UadpAgent
}

type mockInterfaceUadpagent struct {
	mock.Mock
}

// method SetDataKey

func (v *mockInterfaceUadpagent) GoSetDataKey(flags dbus.Flags, ch chan *dbus.Call, keyName string, dataKey string) *dbus.Call {
	mockArgs := v.Called(flags, ch, keyName, dataKey)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUadpagent) SetDataKey(flags dbus.Flags, keyName string, dataKey string) error {
	mockArgs := v.Called(flags, keyName, dataKey)

	return mockArgs.Error(0)
}

// method GetDataKey

func (v *mockInterfaceUadpagent) GoGetDataKey(flags dbus.Flags, ch chan *dbus.Call, keyName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, keyName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceUadpagent) GetDataKey(flags dbus.Flags, keyName string) (string, error) {
	mockArgs := v.Called(flags, keyName)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}
