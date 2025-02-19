// Code generated by "./generator ./org.ayatana.bamf"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package bamf

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
	"github.com/stretchr/testify/mock"
)

type MockControl struct {
	MockInterfaceControl // interface org.ayatana.bamf.control
	proxy.MockObject
}

type MockInterfaceControl struct {
	mock.Mock
}

// method Quit

func (v *MockInterfaceControl) GoQuit(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceControl) Quit(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method InsertDesktopFile

func (v *MockInterfaceControl) GoInsertDesktopFile(flags dbus.Flags, ch chan *dbus.Call, desktop_path string) *dbus.Call {
	mockArgs := v.Called(flags, ch, desktop_path)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceControl) InsertDesktopFile(flags dbus.Flags, desktop_path string) error {
	mockArgs := v.Called(flags, desktop_path)

	return mockArgs.Error(0)
}

// method RegisterApplicationForPid

func (v *MockInterfaceControl) GoRegisterApplicationForPid(flags dbus.Flags, ch chan *dbus.Call, application string, pid int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, application, pid)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceControl) RegisterApplicationForPid(flags dbus.Flags, application string, pid int32) error {
	mockArgs := v.Called(flags, application, pid)

	return mockArgs.Error(0)
}

// method CreateLocalDesktopFile

func (v *MockInterfaceControl) GoCreateLocalDesktopFile(flags dbus.Flags, ch chan *dbus.Call, application string) *dbus.Call {
	mockArgs := v.Called(flags, ch, application)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceControl) CreateLocalDesktopFile(flags dbus.Flags, application string) error {
	mockArgs := v.Called(flags, application)

	return mockArgs.Error(0)
}

// method OmNomNomDesktopFile

func (v *MockInterfaceControl) GoOmNomNomDesktopFile(flags dbus.Flags, ch chan *dbus.Call, desktop_path string) *dbus.Call {
	mockArgs := v.Called(flags, ch, desktop_path)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceControl) OmNomNomDesktopFile(flags dbus.Flags, desktop_path string) error {
	mockArgs := v.Called(flags, desktop_path)

	return mockArgs.Error(0)
}

type MockMatcher struct {
	MockInterfaceMatcher // interface org.ayatana.bamf.matcher
	proxy.MockObject
}

type MockInterfaceMatcher struct {
	mock.Mock
}

// method XidsForApplication

func (v *MockInterfaceMatcher) GoXidsForApplication(flags dbus.Flags, ch chan *dbus.Call, desktop_file string) *dbus.Call {
	mockArgs := v.Called(flags, ch, desktop_file)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceMatcher) XidsForApplication(flags dbus.Flags, desktop_file string) ([]uint32, error) {
	mockArgs := v.Called(flags, desktop_file)

	ret0, ok := mockArgs.Get(0).([]uint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method TabPaths

func (v *MockInterfaceMatcher) GoTabPaths(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceMatcher) TabPaths(flags dbus.Flags) ([]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method RunningApplications

func (v *MockInterfaceMatcher) GoRunningApplications(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceMatcher) RunningApplications(flags dbus.Flags) ([]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method RunningApplicationsDesktopFiles

func (v *MockInterfaceMatcher) GoRunningApplicationsDesktopFiles(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceMatcher) RunningApplicationsDesktopFiles(flags dbus.Flags) ([]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method RegisterFavorites

func (v *MockInterfaceMatcher) GoRegisterFavorites(flags dbus.Flags, ch chan *dbus.Call, favorites []string) *dbus.Call {
	mockArgs := v.Called(flags, ch, favorites)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceMatcher) RegisterFavorites(flags dbus.Flags, favorites []string) error {
	mockArgs := v.Called(flags, favorites)

	return mockArgs.Error(0)
}

// method PathForApplication

func (v *MockInterfaceMatcher) GoPathForApplication(flags dbus.Flags, ch chan *dbus.Call, desktop_file string) *dbus.Call {
	mockArgs := v.Called(flags, ch, desktop_file)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceMatcher) PathForApplication(flags dbus.Flags, desktop_file string) (string, error) {
	mockArgs := v.Called(flags, desktop_file)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method WindowPaths

func (v *MockInterfaceMatcher) GoWindowPaths(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceMatcher) WindowPaths(flags dbus.Flags) ([]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method ApplicationPaths

func (v *MockInterfaceMatcher) GoApplicationPaths(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceMatcher) ApplicationPaths(flags dbus.Flags) ([]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method ApplicationIsRunning

func (v *MockInterfaceMatcher) GoApplicationIsRunning(flags dbus.Flags, ch chan *dbus.Call, desktop_file string) *dbus.Call {
	mockArgs := v.Called(flags, ch, desktop_file)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceMatcher) ApplicationIsRunning(flags dbus.Flags, desktop_file string) (bool, error) {
	mockArgs := v.Called(flags, desktop_file)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method ApplicationForXid

func (v *MockInterfaceMatcher) GoApplicationForXid(flags dbus.Flags, ch chan *dbus.Call, xid uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, xid)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceMatcher) ApplicationForXid(flags dbus.Flags, xid uint32) (string, error) {
	mockArgs := v.Called(flags, xid)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method ActiveWindow

func (v *MockInterfaceMatcher) GoActiveWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceMatcher) ActiveWindow(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method ActiveApplication

func (v *MockInterfaceMatcher) GoActiveApplication(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceMatcher) ActiveApplication(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method WindowStackForMonitor

func (v *MockInterfaceMatcher) GoWindowStackForMonitor(flags dbus.Flags, ch chan *dbus.Call, monitor_id int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, monitor_id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceMatcher) WindowStackForMonitor(flags dbus.Flags, monitor_id int32) ([]string, error) {
	mockArgs := v.Called(flags, monitor_id)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal ActiveApplicationChanged

func (v *MockInterfaceMatcher) ConnectActiveApplicationChanged(cb func(old_app string, new_app string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal ActiveWindowChanged

func (v *MockInterfaceMatcher) ConnectActiveWindowChanged(cb func(old_win string, new_win string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal ViewClosed

func (v *MockInterfaceMatcher) ConnectViewClosed(cb func(path string, type0 string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal ViewOpened

func (v *MockInterfaceMatcher) ConnectViewOpened(cb func(path string, type0 string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal StackingOrderChanged

func (v *MockInterfaceMatcher) ConnectStackingOrderChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal RunningApplicationsChanged

func (v *MockInterfaceMatcher) ConnectRunningApplicationsChanged(cb func(opened_desktop_files []string, closed_desktop_files []string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

type MockApplication struct {
	MockInterfaceApplication // interface org.ayatana.bamf.application
	MockInterfaceView        // interface org.ayatana.bamf.view
	proxy.MockObject
}

type MockInterfaceApplication struct {
	mock.Mock
}

// method ShowStubs

func (v *MockInterfaceApplication) GoShowStubs(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceApplication) ShowStubs(flags dbus.Flags) (bool, error) {
	mockArgs := v.Called(flags)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method Xids

func (v *MockInterfaceApplication) GoXids(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceApplication) Xids(flags dbus.Flags) ([]uint32, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]uint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method DesktopFile

func (v *MockInterfaceApplication) GoDesktopFile(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceApplication) DesktopFile(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method SupportedMimeTypes

func (v *MockInterfaceApplication) GoSupportedMimeTypes(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceApplication) SupportedMimeTypes(flags dbus.Flags) ([]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method ApplicationType

func (v *MockInterfaceApplication) GoApplicationType(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceApplication) ApplicationType(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method ApplicationMenu

func (v *MockInterfaceApplication) GoApplicationMenu(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceApplication) ApplicationMenu(flags dbus.Flags) (string, string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.String(1), mockArgs.Error(2)
}

// method FocusableChild

func (v *MockInterfaceApplication) GoFocusableChild(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceApplication) FocusableChild(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.Error(1)
}

// signal WindowRemoved

func (v *MockInterfaceApplication) ConnectWindowRemoved(cb func(path string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal WindowAdded

func (v *MockInterfaceApplication) ConnectWindowAdded(cb func(path string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal SupportedMimeTypesChanged

func (v *MockInterfaceApplication) ConnectSupportedMimeTypesChanged(cb func(dnd_mimes []string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal DesktopFileUpdated

func (v *MockInterfaceApplication) ConnectDesktopFileUpdated(cb func(desktop_file string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

type MockInterfaceView struct {
	mock.Mock
}

// method ViewType

func (v *MockInterfaceView) GoViewType(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceView) ViewType(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method Icon

func (v *MockInterfaceView) GoIcon(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceView) Icon(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method Name

func (v *MockInterfaceView) GoName(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceView) Name(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method UserVisible

func (v *MockInterfaceView) GoUserVisible(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceView) UserVisible(flags dbus.Flags) (bool, error) {
	mockArgs := v.Called(flags)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method IsUrgent

func (v *MockInterfaceView) GoIsUrgent(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceView) IsUrgent(flags dbus.Flags) (bool, error) {
	mockArgs := v.Called(flags)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method IsRunning

func (v *MockInterfaceView) GoIsRunning(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceView) IsRunning(flags dbus.Flags) (bool, error) {
	mockArgs := v.Called(flags)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method IsActive

func (v *MockInterfaceView) GoIsActive(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceView) IsActive(flags dbus.Flags) (bool, error) {
	mockArgs := v.Called(flags)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method Parents

func (v *MockInterfaceView) GoParents(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceView) Parents(flags dbus.Flags) ([]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method Children

func (v *MockInterfaceView) GoChildren(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceView) Children(flags dbus.Flags) ([]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal NameChanged

func (v *MockInterfaceView) ConnectNameChanged(cb func(old_name string, new_name string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal UserVisibleChanged

func (v *MockInterfaceView) ConnectUserVisibleChanged(cb func(user_visible bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal UrgentChanged

func (v *MockInterfaceView) ConnectUrgentChanged(cb func(is_urgent bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal RunningChanged

func (v *MockInterfaceView) ConnectRunningChanged(cb func(is_running bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal ActiveChanged

func (v *MockInterfaceView) ConnectActiveChanged(cb func(is_active bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal ChildRemoved

func (v *MockInterfaceView) ConnectChildRemoved(cb func(path string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal ChildAdded

func (v *MockInterfaceView) ConnectChildAdded(cb func(path string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal Closed

func (v *MockInterfaceView) ConnectClosed(cb func()) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property Name s

func (v *MockInterfaceView) PropName() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Icon s

func (v *MockInterfaceView) PropIcon() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property UserVisible b

func (v *MockInterfaceView) PropUserVisible() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Running b

func (v *MockInterfaceView) Running() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Starting b

func (v *MockInterfaceView) Starting() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Urgent b

func (v *MockInterfaceView) Urgent() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Active b

func (v *MockInterfaceView) Active() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockWindow struct {
	MockInterfaceWindow // interface org.ayatana.bamf.window
	MockInterfaceView   // interface org.ayatana.bamf.view
	proxy.MockObject
}

type MockInterfaceWindow struct {
	mock.Mock
}

// method GetXid

func (v *MockInterfaceWindow) GoGetXid(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWindow) GetXid(flags dbus.Flags) (uint32, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(uint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetPid

func (v *MockInterfaceWindow) GoGetPid(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWindow) GetPid(flags dbus.Flags) (uint32, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(uint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method Transient

func (v *MockInterfaceWindow) GoTransient(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWindow) Transient(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method WindowType

func (v *MockInterfaceWindow) GoWindowType(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWindow) WindowType(flags dbus.Flags) (uint32, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(uint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method Xprop

func (v *MockInterfaceWindow) GoXprop(flags dbus.Flags, ch chan *dbus.Call, xprop string) *dbus.Call {
	mockArgs := v.Called(flags, ch, xprop)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWindow) Xprop(flags dbus.Flags, xprop string) (string, error) {
	mockArgs := v.Called(flags, xprop)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method Monitor

func (v *MockInterfaceWindow) GoMonitor(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWindow) Monitor(flags dbus.Flags) (int32, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(int32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method Maximized

func (v *MockInterfaceWindow) GoMaximized(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWindow) Maximized(flags dbus.Flags) (int32, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(int32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal MonitorChanged

func (v *MockInterfaceWindow) ConnectMonitorChanged(cb func(old int32, new int32)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal MaximizedChanged

func (v *MockInterfaceWindow) ConnectMaximizedChanged(cb func(old int32, new int32)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}
