// Code generated by "./generator ./com.deepin.system.power"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package power

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
	"github.com/stretchr/testify/mock"
)

type MockPower struct {
	MockInterfacePower // interface com.deepin.system.Power
	proxy.MockObject
}

type MockInterfacePower struct {
	mock.Mock
}

// method GetBatteries

func (v *MockInterfacePower) GoGetBatteries(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfacePower) GetBatteries(flags dbus.Flags) ([]dbus.ObjectPath, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method Refresh

func (v *MockInterfacePower) GoRefresh(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfacePower) Refresh(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method RefreshBatteries

func (v *MockInterfacePower) GoRefreshBatteries(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfacePower) RefreshBatteries(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method RefreshMains

func (v *MockInterfacePower) GoRefreshMains(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfacePower) RefreshMains(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method SetCpuGovernor

func (v *MockInterfacePower) GoSetCpuGovernor(flags dbus.Flags, ch chan *dbus.Call, governor string) *dbus.Call {
	mockArgs := v.Called(flags, ch, governor)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfacePower) SetCpuGovernor(flags dbus.Flags, governor string) error {
	mockArgs := v.Called(flags, governor)

	return mockArgs.Error(0)
}

// method SetCpuBoost

func (v *MockInterfacePower) GoSetCpuBoost(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, enabled)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfacePower) SetCpuBoost(flags dbus.Flags, enabled bool) error {
	mockArgs := v.Called(flags, enabled)

	return mockArgs.Error(0)
}

// method SetMode

func (v *MockInterfacePower) GoSetMode(flags dbus.Flags, ch chan *dbus.Call, mode string) *dbus.Call {
	mockArgs := v.Called(flags, ch, mode)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfacePower) SetMode(flags dbus.Flags, mode string) error {
	mockArgs := v.Called(flags, mode)

	return mockArgs.Error(0)
}

// method LockCpuFreq

func (v *MockInterfacePower) GoLockCpuFreq(flags dbus.Flags, ch chan *dbus.Call, governor string, lockTime int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, governor, lockTime)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfacePower) LockCpuFreq(flags dbus.Flags, governor string, lockTime int32) error {
	mockArgs := v.Called(flags, governor, lockTime)

	return mockArgs.Error(0)
}

// signal BatteryDisplayUpdate

func (v *MockInterfacePower) ConnectBatteryDisplayUpdate(cb func(timestamp int64)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal BatteryAdded

func (v *MockInterfacePower) ConnectBatteryAdded(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal BatteryRemoved

func (v *MockInterfacePower) ConnectBatteryRemoved(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal LidClosed

func (v *MockInterfacePower) ConnectLidClosed(cb func()) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal LidOpened

func (v *MockInterfacePower) ConnectLidOpened(cb func()) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal PowerActionCode

func (v *MockInterfacePower) ConnectPowerActionCode(cb func(actionCode int32)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property PowerSavingModeAuto b

func (v *MockInterfacePower) PowerSavingModeAuto() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property OnBattery b

func (v *MockInterfacePower) OnBattery() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property HasLidSwitch b

func (v *MockInterfacePower) HasLidSwitch() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property BatteryPercentage d

func (v *MockInterfacePower) BatteryPercentage() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property BatteryTimeToEmpty t

func (v *MockInterfacePower) BatteryTimeToEmpty() proxy.PropUint64 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property HasBattery b

func (v *MockInterfacePower) HasBattery() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property BatteryStatus u

func (v *MockInterfacePower) BatteryStatus() proxy.PropUint32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property BatteryTimeToFull t

func (v *MockInterfacePower) BatteryTimeToFull() proxy.PropUint64 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property BatteryCapacity d

func (v *MockInterfacePower) BatteryCapacity() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property PowerSavingModeEnabled b

func (v *MockInterfacePower) PowerSavingModeEnabled() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property PowerSavingModeAutoWhenBatteryLow b

func (v *MockInterfacePower) PowerSavingModeAutoWhenBatteryLow() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property PowerSavingModeBrightnessDropPercent u

func (v *MockInterfacePower) PowerSavingModeBrightnessDropPercent() proxy.PropUint32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property PowerSavingModeBrightnessData s

func (v *MockInterfacePower) PowerSavingModeBrightnessData() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property CpuGovernor s

func (v *MockInterfacePower) CpuGovernor() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property CpuBoost b

func (v *MockInterfacePower) CpuBoost() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property IsBoostSupported b

func (v *MockInterfacePower) IsBoostSupported() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property IsHighPerformanceSupported b

func (v *MockInterfacePower) IsHighPerformanceSupported() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property IsBalanceSupported b

func (v *MockInterfacePower) IsBalanceSupported() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property IsPowerSaveSupported b

func (v *MockInterfacePower) IsPowerSaveSupported() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Mode s

func (v *MockInterfacePower) Mode() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockBattery struct {
	MockInterfaceBattery // interface com.deepin.system.Power.Battery
	proxy.MockObject
}

type MockInterfaceBattery struct {
	mock.Mock
}

// property EnergyFullDesign d

func (v *MockInterfaceBattery) EnergyFullDesign() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Capacity d

func (v *MockInterfaceBattery) Capacity() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Technology s

func (v *MockInterfaceBattery) Technology() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Energy d

func (v *MockInterfaceBattery) Energy() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property EnergyFull d

func (v *MockInterfaceBattery) EnergyFull() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Manufacturer s

func (v *MockInterfaceBattery) Manufacturer() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property ModelName s

func (v *MockInterfaceBattery) ModelName() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property TimeToEmpty t

func (v *MockInterfaceBattery) TimeToEmpty() proxy.PropUint64 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property IsPresent b

func (v *MockInterfaceBattery) IsPresent() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Status u

func (v *MockInterfaceBattery) Status() proxy.PropUint32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property EnergyRate d

func (v *MockInterfaceBattery) EnergyRate() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Voltage d

func (v *MockInterfaceBattery) Voltage() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Percentage d

func (v *MockInterfaceBattery) Percentage() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property TimeToFull t

func (v *MockInterfaceBattery) TimeToFull() proxy.PropUint64 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property UpdateTime x

func (v *MockInterfaceBattery) UpdateTime() proxy.PropInt64 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property SysfsPath s

func (v *MockInterfaceBattery) SysfsPath() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property SerialNumber s

func (v *MockInterfaceBattery) SerialNumber() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Name s

func (v *MockInterfaceBattery) Name() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
