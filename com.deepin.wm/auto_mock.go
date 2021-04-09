// Code generated by "./generator ./com.deepin.wm"; DO NOT EDIT.

package wm

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/stretchr/testify/mock"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type MockWm struct {
	mockInterfaceWm // interface com.deepin.wm
}

type mockInterfaceWm struct {
	mock.Mock
}

// method SwitchApplication

func (v *mockInterfaceWm) GoSwitchApplication(flags dbus.Flags, ch chan *dbus.Call, backward bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, backward)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) SwitchApplication(flags dbus.Flags, backward bool) error {
	mockArgs := v.Called(flags, backward)

	return mockArgs.Error(0)
}

// method TileActiveWindow

func (v *mockInterfaceWm) GoTileActiveWindow(flags dbus.Flags, ch chan *dbus.Call, side uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, side)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) TileActiveWindow(flags dbus.Flags, side uint32) error {
	mockArgs := v.Called(flags, side)

	return mockArgs.Error(0)
}

// method BeginToMoveActiveWindow

func (v *mockInterfaceWm) GoBeginToMoveActiveWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) BeginToMoveActiveWindow(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method ToggleActiveWindowMaximize

func (v *mockInterfaceWm) GoToggleActiveWindowMaximize(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) ToggleActiveWindowMaximize(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method MinimizeActiveWindow

func (v *mockInterfaceWm) GoMinimizeActiveWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) MinimizeActiveWindow(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method ShowWorkspace

func (v *mockInterfaceWm) GoShowWorkspace(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) ShowWorkspace(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method ShowWindow

func (v *mockInterfaceWm) GoShowWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) ShowWindow(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method ShowAllWindow

func (v *mockInterfaceWm) GoShowAllWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) ShowAllWindow(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method ClearMoveStatus

func (v *mockInterfaceWm) GoClearMoveStatus(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) ClearMoveStatus(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method TouchToMove

func (v *mockInterfaceWm) GoTouchToMove(flags dbus.Flags, ch chan *dbus.Call, x int32, y int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, x, y)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) TouchToMove(flags dbus.Flags, x int32, y int32) error {
	mockArgs := v.Called(flags, x, y)

	return mockArgs.Error(0)
}

// method PerformAction

func (v *mockInterfaceWm) GoPerformAction(flags dbus.Flags, ch chan *dbus.Call, type0 int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, type0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) PerformAction(flags dbus.Flags, type0 int32) error {
	mockArgs := v.Called(flags, type0)

	return mockArgs.Error(0)
}

// method PreviewWindow

func (v *mockInterfaceWm) GoPreviewWindow(flags dbus.Flags, ch chan *dbus.Call, xid uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, xid)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) PreviewWindow(flags dbus.Flags, xid uint32) error {
	mockArgs := v.Called(flags, xid)

	return mockArgs.Error(0)
}

// method CancelPreviewWindow

func (v *mockInterfaceWm) GoCancelPreviewWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) CancelPreviewWindow(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method GetCurrentWorkspaceBackground

func (v *mockInterfaceWm) GoGetCurrentWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) GetCurrentWorkspaceBackground(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method SetCurrentWorkspaceBackground

func (v *mockInterfaceWm) GoSetCurrentWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, uri string) *dbus.Call {
	mockArgs := v.Called(flags, ch, uri)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) SetCurrentWorkspaceBackground(flags dbus.Flags, uri string) error {
	mockArgs := v.Called(flags, uri)

	return mockArgs.Error(0)
}

// method GetWorkspaceBackground

func (v *mockInterfaceWm) GoGetWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, index int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, index)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) GetWorkspaceBackground(flags dbus.Flags, index int32) (string, error) {
	mockArgs := v.Called(flags, index)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method SetWorkspaceBackground

func (v *mockInterfaceWm) GoSetWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, index int32, uri string) *dbus.Call {
	mockArgs := v.Called(flags, ch, index, uri)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) SetWorkspaceBackground(flags dbus.Flags, index int32, uri string) error {
	mockArgs := v.Called(flags, index, uri)

	return mockArgs.Error(0)
}

// method SetTransientBackground

func (v *mockInterfaceWm) GoSetTransientBackground(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) SetTransientBackground(flags dbus.Flags, arg0 string) error {
	mockArgs := v.Called(flags, arg0)

	return mockArgs.Error(0)
}

// method GetCurrentWorkspaceBackgroundForMonitor

func (v *mockInterfaceWm) GoGetCurrentWorkspaceBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, strMonitorName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, strMonitorName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) GetCurrentWorkspaceBackgroundForMonitor(flags dbus.Flags, strMonitorName string) (string, error) {
	mockArgs := v.Called(flags, strMonitorName)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method SetCurrentWorkspaceBackgroundForMonitor

func (v *mockInterfaceWm) GoSetCurrentWorkspaceBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, uri string, strMonitorName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, uri, strMonitorName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) SetCurrentWorkspaceBackgroundForMonitor(flags dbus.Flags, uri string, strMonitorName string) error {
	mockArgs := v.Called(flags, uri, strMonitorName)

	return mockArgs.Error(0)
}

// method GetWorkspaceBackgroundForMonitor

func (v *mockInterfaceWm) GoGetWorkspaceBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, index int32, strMonitorName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, index, strMonitorName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) GetWorkspaceBackgroundForMonitor(flags dbus.Flags, index int32, strMonitorName string) (string, error) {
	mockArgs := v.Called(flags, index, strMonitorName)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method SetWorkspaceBackgroundForMonitor

func (v *mockInterfaceWm) GoSetWorkspaceBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, index int32, strMonitorName string, uri string) *dbus.Call {
	mockArgs := v.Called(flags, ch, index, strMonitorName, uri)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) SetWorkspaceBackgroundForMonitor(flags dbus.Flags, index int32, strMonitorName string, uri string) error {
	mockArgs := v.Called(flags, index, strMonitorName, uri)

	return mockArgs.Error(0)
}

// method SetTransientBackgroundForMonitor

func (v *mockInterfaceWm) GoSetTransientBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, uri string, strMonitorName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, uri, strMonitorName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) SetTransientBackgroundForMonitor(flags dbus.Flags, uri string, strMonitorName string) error {
	mockArgs := v.Called(flags, uri, strMonitorName)

	return mockArgs.Error(0)
}

// method GetCurrentWorkspace

func (v *mockInterfaceWm) GoGetCurrentWorkspace(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) GetCurrentWorkspace(flags dbus.Flags) (int32, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(int32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method WorkspaceCount

func (v *mockInterfaceWm) GoWorkspaceCount(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) WorkspaceCount(flags dbus.Flags) (int32, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(int32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method SetCurrentWorkspace

func (v *mockInterfaceWm) GoSetCurrentWorkspace(flags dbus.Flags, ch chan *dbus.Call, index int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, index)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) SetCurrentWorkspace(flags dbus.Flags, index int32) error {
	mockArgs := v.Called(flags, index)

	return mockArgs.Error(0)
}

// method PreviousWorkspace

func (v *mockInterfaceWm) GoPreviousWorkspace(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) PreviousWorkspace(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method NextWorkspace

func (v *mockInterfaceWm) GoNextWorkspace(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) NextWorkspace(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method GetAllAccels

func (v *mockInterfaceWm) GoGetAllAccels(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) GetAllAccels(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetAccel

func (v *mockInterfaceWm) GoGetAccel(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) GetAccel(flags dbus.Flags, id string) ([]string, error) {
	mockArgs := v.Called(flags, id)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetDefaultAccel

func (v *mockInterfaceWm) GoGetDefaultAccel(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) GetDefaultAccel(flags dbus.Flags, id string) ([]string, error) {
	mockArgs := v.Called(flags, id)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method SetAccel

func (v *mockInterfaceWm) GoSetAccel(flags dbus.Flags, ch chan *dbus.Call, data string) *dbus.Call {
	mockArgs := v.Called(flags, ch, data)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) SetAccel(flags dbus.Flags, data string) (bool, error) {
	mockArgs := v.Called(flags, data)

	ret0, ok := mockArgs.Get(0).(bool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method RemoveAccel

func (v *mockInterfaceWm) GoRemoveAccel(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) RemoveAccel(flags dbus.Flags, id string) error {
	mockArgs := v.Called(flags, id)

	return mockArgs.Error(0)
}

// method SetDecorationTheme

func (v *mockInterfaceWm) GoSetDecorationTheme(flags dbus.Flags, ch chan *dbus.Call, themeType string, themeName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, themeType, themeName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) SetDecorationTheme(flags dbus.Flags, themeType string, themeName string) error {
	mockArgs := v.Called(flags, themeType, themeName)

	return mockArgs.Error(0)
}

// method SetDecorationDeepinTheme

func (v *mockInterfaceWm) GoSetDecorationDeepinTheme(flags dbus.Flags, ch chan *dbus.Call, deepinThemeName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, deepinThemeName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) SetDecorationDeepinTheme(flags dbus.Flags, deepinThemeName string) error {
	mockArgs := v.Called(flags, deepinThemeName)

	return mockArgs.Error(0)
}

// method ChangeCurrentWorkspaceBackground

func (v *mockInterfaceWm) GoChangeCurrentWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, uri string) *dbus.Call {
	mockArgs := v.Called(flags, ch, uri)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) ChangeCurrentWorkspaceBackground(flags dbus.Flags, uri string) error {
	mockArgs := v.Called(flags, uri)

	return mockArgs.Error(0)
}

// method SwitchToWorkspace

func (v *mockInterfaceWm) GoSwitchToWorkspace(flags dbus.Flags, ch chan *dbus.Call, backward bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, backward)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) SwitchToWorkspace(flags dbus.Flags, backward bool) error {
	mockArgs := v.Called(flags, backward)

	return mockArgs.Error(0)
}

// method PresentWindows

func (v *mockInterfaceWm) GoPresentWindows(flags dbus.Flags, ch chan *dbus.Call, xids []uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, xids)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) PresentWindows(flags dbus.Flags, xids []uint32) error {
	mockArgs := v.Called(flags, xids)

	return mockArgs.Error(0)
}

// method EnableZoneDetected

func (v *mockInterfaceWm) GoEnableZoneDetected(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, enabled)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) EnableZoneDetected(flags dbus.Flags, enabled bool) error {
	mockArgs := v.Called(flags, enabled)

	return mockArgs.Error(0)
}

// method GetIsShowDesktop

func (v *mockInterfaceWm) GoGetIsShowDesktop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) GetIsShowDesktop(flags dbus.Flags) (bool, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(bool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetMultiTaskingStatus

func (v *mockInterfaceWm) GoGetMultiTaskingStatus(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *mockInterfaceWm) GetMultiTaskingStatus(flags dbus.Flags) (bool, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(bool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal WorkspaceBackgroundChanged

func (v *mockInterfaceWm) ConnectWorkspaceBackgroundChanged(cb func(index int32, newUri string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal WorkspaceBackgroundChangedForMonitor

func (v *mockInterfaceWm) ConnectWorkspaceBackgroundChangedForMonitor(cb func(index int32, monitor string, newUri string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal compositingEnabledChanged

func (v *mockInterfaceWm) ConnectCompositingEnabledChanged(cb func(enabled bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal wmCompositingEnabledChanged

func (v *mockInterfaceWm) ConnectWmCompositingEnabledChanged(cb func(enabled bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal workspaceCountChanged

func (v *mockInterfaceWm) ConnectWorkspaceCountChanged(cb func(count int32)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal WorkspaceSwitched

func (v *mockInterfaceWm) ConnectWorkspaceSwitched(cb func(from int32, to int32)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property compositingEnabled b

func (v *mockInterfaceWm) CompositingEnabled() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property compositingPossible b

func (v *mockInterfaceWm) CompositingPossible() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property compositingAllowSwitch b

func (v *mockInterfaceWm) CompositingAllowSwitch() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property zoneEnabled b

func (v *mockInterfaceWm) ZoneEnabled() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property cursorTheme s

func (v *mockInterfaceWm) CursorTheme() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property cursorSize i

func (v *mockInterfaceWm) CursorSize() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
