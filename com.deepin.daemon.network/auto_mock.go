// Code generated by "./generator ./com.deepin.daemon.network"; DO NOT EDIT.

package network

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/stretchr/testify/mock"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type MockNetwork struct {
	mockInterfaceNetwork // interface com.deepin.daemon.Network
}

type mockInterfaceNetwork struct {
	mock.Mock
}

// method ActivateAccessPoint

func (v *mockInterfaceNetwork) GoActivateAccessPoint(flags dbus.Flags, ch chan *dbus.Call, uuid string, apPath dbus.ObjectPath, devPath dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, uuid, apPath, devPath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) ActivateAccessPoint(flags dbus.Flags, uuid string, apPath dbus.ObjectPath, devPath dbus.ObjectPath) (dbus.ObjectPath, error) {
	mockArgs := v.Called(flags, uuid, apPath, devPath)

	ret0, ok := mockArgs.Get(0).(dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method ActivateConnection

func (v *mockInterfaceNetwork) GoActivateConnection(flags dbus.Flags, ch chan *dbus.Call, uuid string, devPath dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, uuid, devPath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) ActivateConnection(flags dbus.Flags, uuid string, devPath dbus.ObjectPath) (dbus.ObjectPath, error) {
	mockArgs := v.Called(flags, uuid, devPath)

	ret0, ok := mockArgs.Get(0).(dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method DeactivateConnection

func (v *mockInterfaceNetwork) GoDeactivateConnection(flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call {
	mockArgs := v.Called(flags, ch, uuid)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) DeactivateConnection(flags dbus.Flags, uuid string) error {
	mockArgs := v.Called(flags, uuid)

	return mockArgs.Error(0)
}

// method DeleteConnection

func (v *mockInterfaceNetwork) GoDeleteConnection(flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call {
	mockArgs := v.Called(flags, ch, uuid)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) DeleteConnection(flags dbus.Flags, uuid string) error {
	mockArgs := v.Called(flags, uuid)

	return mockArgs.Error(0)
}

// method DisableWirelessHotspotMode

func (v *mockInterfaceNetwork) GoDisableWirelessHotspotMode(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, devPath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) DisableWirelessHotspotMode(flags dbus.Flags, devPath dbus.ObjectPath) error {
	mockArgs := v.Called(flags, devPath)

	return mockArgs.Error(0)
}

// method DisconnectDevice

func (v *mockInterfaceNetwork) GoDisconnectDevice(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, devPath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) DisconnectDevice(flags dbus.Flags, devPath dbus.ObjectPath) error {
	mockArgs := v.Called(flags, devPath)

	return mockArgs.Error(0)
}

// method EnableDevice

func (v *mockInterfaceNetwork) GoEnableDevice(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath, enabled bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, devPath, enabled)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) EnableDevice(flags dbus.Flags, devPath dbus.ObjectPath, enabled bool) (dbus.ObjectPath, error) {
	mockArgs := v.Called(flags, devPath, enabled)

	ret0, ok := mockArgs.Get(0).(dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method EnableWirelessHotspotMode

func (v *mockInterfaceNetwork) GoEnableWirelessHotspotMode(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, devPath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) EnableWirelessHotspotMode(flags dbus.Flags, devPath dbus.ObjectPath) error {
	mockArgs := v.Called(flags, devPath)

	return mockArgs.Error(0)
}

// method GetAccessPoints

func (v *mockInterfaceNetwork) GoGetAccessPoints(flags dbus.Flags, ch chan *dbus.Call, path dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, path)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) GetAccessPoints(flags dbus.Flags, path dbus.ObjectPath) (string, error) {
	mockArgs := v.Called(flags, path)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetActiveConnectionInfo

func (v *mockInterfaceNetwork) GoGetActiveConnectionInfo(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) GetActiveConnectionInfo(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetAutoProxy

func (v *mockInterfaceNetwork) GoGetAutoProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) GetAutoProxy(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetProxy

func (v *mockInterfaceNetwork) GoGetProxy(flags dbus.Flags, ch chan *dbus.Call, proxyType string) *dbus.Call {
	mockArgs := v.Called(flags, ch, proxyType)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) GetProxy(flags dbus.Flags, proxyType string) (string, string, error) {
	mockArgs := v.Called(flags, proxyType)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	ret1, ok := mockArgs.Get(1).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 1, mockArgs.Get(1)))
	}

	return ret0, ret1, mockArgs.Error(2)
}

// method GetProxyIgnoreHosts

func (v *mockInterfaceNetwork) GoGetProxyIgnoreHosts(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) GetProxyIgnoreHosts(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetProxyMethod

func (v *mockInterfaceNetwork) GoGetProxyMethod(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) GetProxyMethod(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetSupportedConnectionTypes

func (v *mockInterfaceNetwork) GoGetSupportedConnectionTypes(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) GetSupportedConnectionTypes(flags dbus.Flags) ([]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method IsDeviceEnabled

func (v *mockInterfaceNetwork) GoIsDeviceEnabled(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, devPath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) IsDeviceEnabled(flags dbus.Flags, devPath dbus.ObjectPath) (bool, error) {
	mockArgs := v.Called(flags, devPath)

	ret0, ok := mockArgs.Get(0).(bool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method IsWirelessHotspotModeEnabled

func (v *mockInterfaceNetwork) GoIsWirelessHotspotModeEnabled(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, devPath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) IsWirelessHotspotModeEnabled(flags dbus.Flags, devPath dbus.ObjectPath) (bool, error) {
	mockArgs := v.Called(flags, devPath)

	ret0, ok := mockArgs.Get(0).(bool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method ListDeviceConnections

func (v *mockInterfaceNetwork) GoListDeviceConnections(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, devPath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) ListDeviceConnections(flags dbus.Flags, devPath dbus.ObjectPath) ([]dbus.ObjectPath, error) {
	mockArgs := v.Called(flags, devPath)

	ret0, ok := mockArgs.Get(0).([]dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method RequestWirelessScan

func (v *mockInterfaceNetwork) GoRequestWirelessScan(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) RequestWirelessScan(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method SetAutoProxy

func (v *mockInterfaceNetwork) GoSetAutoProxy(flags dbus.Flags, ch chan *dbus.Call, proxyAuto string) *dbus.Call {
	mockArgs := v.Called(flags, ch, proxyAuto)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) SetAutoProxy(flags dbus.Flags, proxyAuto string) error {
	mockArgs := v.Called(flags, proxyAuto)

	return mockArgs.Error(0)
}

// method SetDeviceManaged

func (v *mockInterfaceNetwork) GoSetDeviceManaged(flags dbus.Flags, ch chan *dbus.Call, devPathOrIfc string, managed bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, devPathOrIfc, managed)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) SetDeviceManaged(flags dbus.Flags, devPathOrIfc string, managed bool) error {
	mockArgs := v.Called(flags, devPathOrIfc, managed)

	return mockArgs.Error(0)
}

// method SetProxy

func (v *mockInterfaceNetwork) GoSetProxy(flags dbus.Flags, ch chan *dbus.Call, proxyType string, host string, port string) *dbus.Call {
	mockArgs := v.Called(flags, ch, proxyType, host, port)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) SetProxy(flags dbus.Flags, proxyType string, host string, port string) error {
	mockArgs := v.Called(flags, proxyType, host, port)

	return mockArgs.Error(0)
}

// method SetProxyIgnoreHosts

func (v *mockInterfaceNetwork) GoSetProxyIgnoreHosts(flags dbus.Flags, ch chan *dbus.Call, ignoreHosts string) *dbus.Call {
	mockArgs := v.Called(flags, ch, ignoreHosts)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) SetProxyIgnoreHosts(flags dbus.Flags, ignoreHosts string) error {
	mockArgs := v.Called(flags, ignoreHosts)

	return mockArgs.Error(0)
}

// method SetProxyMethod

func (v *mockInterfaceNetwork) GoSetProxyMethod(flags dbus.Flags, ch chan *dbus.Call, proxyMode string) *dbus.Call {
	mockArgs := v.Called(flags, ch, proxyMode)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceNetwork) SetProxyMethod(flags dbus.Flags, proxyMode string) error {
	mockArgs := v.Called(flags, proxyMode)

	return mockArgs.Error(0)
}

// signal AccessPointAdded

func (v *mockInterfaceNetwork) ConnectAccessPointAdded(cb func(devPath string, apJSON string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal AccessPointRemoved

func (v *mockInterfaceNetwork) ConnectAccessPointRemoved(cb func(devPath string, apJSON string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal AccessPointPropertiesChanged

func (v *mockInterfaceNetwork) ConnectAccessPointPropertiesChanged(cb func(devPath string, apJSON string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal DeviceEnabled

func (v *mockInterfaceNetwork) ConnectDeviceEnabled(cb func(devPath string, enabled bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property Connectivity u

func (v *mockInterfaceNetwork) Connectivity() proxy.PropUint32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property NetworkingEnabled b

func (v *mockInterfaceNetwork) NetworkingEnabled() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property VpnEnabled b

func (v *mockInterfaceNetwork) VpnEnabled() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Devices s

func (v *mockInterfaceNetwork) Devices() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Connections s

func (v *mockInterfaceNetwork) Connections() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property ActiveConnections s

func (v *mockInterfaceNetwork) ActiveConnections() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property State u

func (v *mockInterfaceNetwork) State() proxy.PropUint32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
