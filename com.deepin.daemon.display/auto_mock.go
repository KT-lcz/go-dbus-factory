// Code generated by "./generator ./com.deepin.daemon.display"; DO NOT EDIT.

package display

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/stretchr/testify/mock"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type MockDisplay struct {
	MockInterfaceDisplay // interface com.deepin.daemon.Display
	proxy.MockObject
}

type MockInterfaceDisplay struct {
	mock.Mock
}

// method ApplyChanges

func (v *MockInterfaceDisplay) GoApplyChanges(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDisplay) ApplyChanges(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method AssociateTouch

func (v *MockInterfaceDisplay) GoAssociateTouch(flags dbus.Flags, ch chan *dbus.Call, outputName string, touchSerial string) *dbus.Call {
	mockArgs := v.Called(flags, ch, outputName, touchSerial)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDisplay) AssociateTouch(flags dbus.Flags, outputName string, touchSerial string) error {
	mockArgs := v.Called(flags, outputName, touchSerial)

	return mockArgs.Error(0)
}

// method CanRotate

func (v *MockInterfaceDisplay) GoCanRotate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDisplay) CanRotate(flags dbus.Flags) (bool, error) {
	mockArgs := v.Called(flags)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method CanSetBrightness

func (v *MockInterfaceDisplay) GoCanSetBrightness(flags dbus.Flags, ch chan *dbus.Call, outputName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, outputName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDisplay) CanSetBrightness(flags dbus.Flags, outputName string) (bool, error) {
	mockArgs := v.Called(flags, outputName)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method ChangeBrightness

func (v *MockInterfaceDisplay) GoChangeBrightness(flags dbus.Flags, ch chan *dbus.Call, raised bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, raised)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDisplay) ChangeBrightness(flags dbus.Flags, raised bool) error {
	mockArgs := v.Called(flags, raised)

	return mockArgs.Error(0)
}

// method DeleteCustomMode

func (v *MockInterfaceDisplay) GoDeleteCustomMode(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	mockArgs := v.Called(flags, ch, name)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDisplay) DeleteCustomMode(flags dbus.Flags, name string) error {
	mockArgs := v.Called(flags, name)

	return mockArgs.Error(0)
}

// method GetBrightness

func (v *MockInterfaceDisplay) GoGetBrightness(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDisplay) GetBrightness(flags dbus.Flags) (map[string]float64, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(map[string]float64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetBuiltinMonitor

func (v *MockInterfaceDisplay) GoGetBuiltinMonitor(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDisplay) GetBuiltinMonitor(flags dbus.Flags) (string, dbus.ObjectPath, error) {
	mockArgs := v.Called(flags)

	ret1, ok := mockArgs.Get(1).(dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 1, mockArgs.Get(1)))
	}

	return mockArgs.String(0), ret1, mockArgs.Error(2)
}

// method GetRealDisplayMode

func (v *MockInterfaceDisplay) GoGetRealDisplayMode(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDisplay) GetRealDisplayMode(flags dbus.Flags) (uint8, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(uint8)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method ListOutputNames

func (v *MockInterfaceDisplay) GoListOutputNames(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDisplay) ListOutputNames(flags dbus.Flags) ([]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method ListOutputsCommonModes

func (v *MockInterfaceDisplay) GoListOutputsCommonModes(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDisplay) ListOutputsCommonModes(flags dbus.Flags) ([][]interface{}, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([][]interface{})
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method ModifyConfigName

func (v *MockInterfaceDisplay) GoModifyConfigName(flags dbus.Flags, ch chan *dbus.Call, name string, newName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, name, newName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDisplay) ModifyConfigName(flags dbus.Flags, name string, newName string) error {
	mockArgs := v.Called(flags, name, newName)

	return mockArgs.Error(0)
}

// method RefreshBrightness

func (v *MockInterfaceDisplay) GoRefreshBrightness(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDisplay) RefreshBrightness(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method Reset

func (v *MockInterfaceDisplay) GoReset(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDisplay) Reset(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method ResetChanges

func (v *MockInterfaceDisplay) GoResetChanges(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDisplay) ResetChanges(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method Save

func (v *MockInterfaceDisplay) GoSave(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDisplay) Save(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method SetAndSaveBrightness

func (v *MockInterfaceDisplay) GoSetAndSaveBrightness(flags dbus.Flags, ch chan *dbus.Call, outputName string, value float64) *dbus.Call {
	mockArgs := v.Called(flags, ch, outputName, value)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDisplay) SetAndSaveBrightness(flags dbus.Flags, outputName string, value float64) error {
	mockArgs := v.Called(flags, outputName, value)

	return mockArgs.Error(0)
}

// method SetBrightness

func (v *MockInterfaceDisplay) GoSetBrightness(flags dbus.Flags, ch chan *dbus.Call, outputName string, value float64) *dbus.Call {
	mockArgs := v.Called(flags, ch, outputName, value)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDisplay) SetBrightness(flags dbus.Flags, outputName string, value float64) error {
	mockArgs := v.Called(flags, outputName, value)

	return mockArgs.Error(0)
}

// method SetColorTemperature

func (v *MockInterfaceDisplay) GoSetColorTemperature(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, value)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDisplay) SetColorTemperature(flags dbus.Flags, value int32) error {
	mockArgs := v.Called(flags, value)

	return mockArgs.Error(0)
}

// method SetMethodAdjustCCT

func (v *MockInterfaceDisplay) GoSetMethodAdjustCCT(flags dbus.Flags, ch chan *dbus.Call, adjustMethod int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, adjustMethod)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDisplay) SetMethodAdjustCCT(flags dbus.Flags, adjustMethod int32) error {
	mockArgs := v.Called(flags, adjustMethod)

	return mockArgs.Error(0)
}

// method SetPrimary

func (v *MockInterfaceDisplay) GoSetPrimary(flags dbus.Flags, ch chan *dbus.Call, outputName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, outputName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDisplay) SetPrimary(flags dbus.Flags, outputName string) error {
	mockArgs := v.Called(flags, outputName)

	return mockArgs.Error(0)
}

// method SwitchMode

func (v *MockInterfaceDisplay) GoSwitchMode(flags dbus.Flags, ch chan *dbus.Call, mode uint8, name string) *dbus.Call {
	mockArgs := v.Called(flags, ch, mode, name)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDisplay) SwitchMode(flags dbus.Flags, mode uint8, name string) error {
	mockArgs := v.Called(flags, mode, name)

	return mockArgs.Error(0)
}

// property ScreenWidth q

func (v *MockInterfaceDisplay) ScreenWidth() proxy.PropUint16 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint16)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property MaxBacklightBrightness u

func (v *MockInterfaceDisplay) MaxBacklightBrightness() proxy.PropUint32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockPropDisplayTouchMap struct {
	mock.Mock
}

func (p MockPropDisplayTouchMap) Get(flags dbus.Flags) (value map[string]string, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(map[string]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p MockPropDisplayTouchMap) Set(flags dbus.Flags, value map[string]string) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p MockPropDisplayTouchMap) ConnectChanged(cb func(hasValue bool, value map[string]string)) error {
	args := p.Called(cb)

	return args.Error(0)
}

// property TouchMap a{ss}

func (v *MockInterfaceDisplay) TouchMap() PropDisplayTouchMap {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropDisplayTouchMap)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property CurrentCustomId s

func (v *MockInterfaceDisplay) CurrentCustomId() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property ScreenHeight q

func (v *MockInterfaceDisplay) ScreenHeight() proxy.PropUint16 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint16)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property ColorTemperatureManual i

func (v *MockInterfaceDisplay) ColorTemperatureManual() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property DisplayMode y

func (v *MockInterfaceDisplay) DisplayMode() proxy.PropByte {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropByte)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockPropDisplayBrightness struct {
	mock.Mock
}

func (p MockPropDisplayBrightness) Get(flags dbus.Flags) (value map[string]float64, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(map[string]float64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p MockPropDisplayBrightness) Set(flags dbus.Flags, value map[string]float64) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p MockPropDisplayBrightness) ConnectChanged(cb func(hasValue bool, value map[string]float64)) error {
	args := p.Called(cb)

	return args.Error(0)
}

// property Brightness a{sd}

func (v *MockInterfaceDisplay) Brightness() PropDisplayBrightness {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropDisplayBrightness)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property HasChanged b

func (v *MockInterfaceDisplay) HasChanged() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockPropTouchscreens struct {
	mock.Mock
}

func (p MockPropTouchscreens) Get(flags dbus.Flags) (value []Touchscreen, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).([]Touchscreen)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p MockPropTouchscreens) Set(flags dbus.Flags, value []Touchscreen) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p MockPropTouchscreens) ConnectChanged(cb func(hasValue bool, value []Touchscreen)) error {
	args := p.Called(cb)

	return args.Error(0)
}

// property Touchscreens a(isss)

func (v *MockInterfaceDisplay) Touchscreens() PropTouchscreens {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropTouchscreens)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Primary s

func (v *MockInterfaceDisplay) Primary() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Monitors ao

func (v *MockInterfaceDisplay) Monitors() proxy.PropObjectPathArray {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropObjectPathArray)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property CustomIdList as

func (v *MockInterfaceDisplay) CustomIdList() proxy.PropStringArray {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropStringArray)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockPropDisplayPrimaryRect struct {
	mock.Mock
}

func (p MockPropDisplayPrimaryRect) Get(flags dbus.Flags) (value Rectangle, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(Rectangle)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p MockPropDisplayPrimaryRect) Set(flags dbus.Flags, value Rectangle) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p MockPropDisplayPrimaryRect) ConnectChanged(cb func(hasValue bool, value Rectangle)) error {
	args := p.Called(cb)

	return args.Error(0)
}

// property PrimaryRect (nnqq)

func (v *MockInterfaceDisplay) PrimaryRect() PropDisplayPrimaryRect {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropDisplayPrimaryRect)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property ColorTemperatureMode i

func (v *MockInterfaceDisplay) ColorTemperatureMode() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockMonitor struct {
	MockInterfaceMonitor // interface com.deepin.daemon.Display.Monitor
	proxy.MockObject
}

type MockInterfaceMonitor struct {
	mock.Mock
}

// method Enable

func (v *MockInterfaceMonitor) GoEnable(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, enabled)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceMonitor) Enable(flags dbus.Flags, enabled bool) error {
	mockArgs := v.Called(flags, enabled)

	return mockArgs.Error(0)
}

// method SetMode

func (v *MockInterfaceMonitor) GoSetMode(flags dbus.Flags, ch chan *dbus.Call, mode uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, mode)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceMonitor) SetMode(flags dbus.Flags, mode uint32) error {
	mockArgs := v.Called(flags, mode)

	return mockArgs.Error(0)
}

// method SetModeBySize

func (v *MockInterfaceMonitor) GoSetModeBySize(flags dbus.Flags, ch chan *dbus.Call, width uint16, height uint16) *dbus.Call {
	mockArgs := v.Called(flags, ch, width, height)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceMonitor) SetModeBySize(flags dbus.Flags, width uint16, height uint16) error {
	mockArgs := v.Called(flags, width, height)

	return mockArgs.Error(0)
}

// method SetPosition

func (v *MockInterfaceMonitor) GoSetPosition(flags dbus.Flags, ch chan *dbus.Call, X int16, y int16) *dbus.Call {
	mockArgs := v.Called(flags, ch, X, y)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceMonitor) SetPosition(flags dbus.Flags, X int16, y int16) error {
	mockArgs := v.Called(flags, X, y)

	return mockArgs.Error(0)
}

// method SetReflect

func (v *MockInterfaceMonitor) GoSetReflect(flags dbus.Flags, ch chan *dbus.Call, value uint16) *dbus.Call {
	mockArgs := v.Called(flags, ch, value)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceMonitor) SetReflect(flags dbus.Flags, value uint16) error {
	mockArgs := v.Called(flags, value)

	return mockArgs.Error(0)
}

// method SetRefreshRate

func (v *MockInterfaceMonitor) GoSetRefreshRate(flags dbus.Flags, ch chan *dbus.Call, value float64) *dbus.Call {
	mockArgs := v.Called(flags, ch, value)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceMonitor) SetRefreshRate(flags dbus.Flags, value float64) error {
	mockArgs := v.Called(flags, value)

	return mockArgs.Error(0)
}

// method SetRotation

func (v *MockInterfaceMonitor) GoSetRotation(flags dbus.Flags, ch chan *dbus.Call, value uint16) *dbus.Call {
	mockArgs := v.Called(flags, ch, value)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceMonitor) SetRotation(flags dbus.Flags, value uint16) error {
	mockArgs := v.Called(flags, value)

	return mockArgs.Error(0)
}

// property Reflect q

func (v *MockInterfaceMonitor) Reflect() proxy.PropUint16 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint16)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Brightness d

func (v *MockInterfaceMonitor) Brightness() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Model s

func (v *MockInterfaceMonitor) Model() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Reflects aq

func (v *MockInterfaceMonitor) Reflects() proxy.PropUint16Array {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint16Array)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property BestMode (uqqd)

func (v *MockInterfaceMonitor) BestMode() PropModeInfo {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropModeInfo)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property PreferredModes a(uqqd)

func (v *MockInterfaceMonitor) PreferredModes() PropModeInfoSlice {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropModeInfoSlice)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property MmHeight u

func (v *MockInterfaceMonitor) MmHeight() proxy.PropUint32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Enabled b

func (v *MockInterfaceMonitor) Enabled() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Connected b

func (v *MockInterfaceMonitor) Connected() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Manufacturer s

func (v *MockInterfaceMonitor) Manufacturer() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Modes a(uqqd)

func (v *MockInterfaceMonitor) Modes() PropModeInfoSlice {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropModeInfoSlice)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Width q

func (v *MockInterfaceMonitor) Width() proxy.PropUint16 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint16)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Rotation q

func (v *MockInterfaceMonitor) Rotation() proxy.PropUint16 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint16)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property RefreshRate d

func (v *MockInterfaceMonitor) RefreshRate() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Name s

func (v *MockInterfaceMonitor) Name() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Rotations aq

func (v *MockInterfaceMonitor) Rotations() proxy.PropUint16Array {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint16Array)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property AvailableFillMode as

func (v *MockInterfaceMonitor) AvailableFillMode() proxy.PropStringArray {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropStringArray)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property SupportFillMode b

func (v *MockInterfaceMonitor) SupportFillMode() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property CurrentFillMode (qqs)

func (v *MockInterfaceMonitor) CurrentFillMode() PropCurrentFillMode {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropCurrentFillMode)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property ID u

func (v *MockInterfaceMonitor) ID() proxy.PropUint32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property MmWidth u

func (v *MockInterfaceMonitor) MmWidth() proxy.PropUint32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property X n

func (v *MockInterfaceMonitor) X() proxy.PropInt16 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt16)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Y n

func (v *MockInterfaceMonitor) Y() proxy.PropInt16 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt16)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Height q

func (v *MockInterfaceMonitor) Height() proxy.PropUint16 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint16)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property CurrentMode (uqqd)

func (v *MockInterfaceMonitor) CurrentMode() PropModeInfo {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropModeInfo)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockPropModeInfo struct {
	mock.Mock
}

func (p MockPropModeInfo) Get(flags dbus.Flags) (value ModeInfo, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(ModeInfo)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p MockPropModeInfo) Set(flags dbus.Flags, value ModeInfo) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p MockPropModeInfo) ConnectChanged(cb func(hasValue bool, value ModeInfo)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type MockPropModeInfoSlice struct {
	mock.Mock
}

func (p MockPropModeInfoSlice) Get(flags dbus.Flags) (value []ModeInfo, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).([]ModeInfo)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p MockPropModeInfoSlice) Set(flags dbus.Flags, value []ModeInfo) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p MockPropModeInfoSlice) ConnectChanged(cb func(hasValue bool, value []ModeInfo)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type MockPropCurrentFillMode struct {
	mock.Mock
}

func (p MockPropCurrentFillMode) Get(flags dbus.Flags) (value FillModeInfo, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(FillModeInfo)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p MockPropCurrentFillMode) Set(flags dbus.Flags, value FillModeInfo) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p MockPropCurrentFillMode) ConnectChanged(cb func(hasValue bool, value FillModeInfo)) error {
	args := p.Called(cb)

	return args.Error(0)
}
