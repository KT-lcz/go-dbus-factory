// Code generated by "./generator ./com.deepin.lastore"; DO NOT EDIT.

package lastore

import (
	"errors"
	"unsafe"

	"github.com/godbus/dbus"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type Lastore interface {
	Manager() manager // interface com.deepin.lastore.Manager
	Updater() updater // interface com.deepin.lastore.Updater
	proxy.Object
}

type objectLastore struct {
	interfaceManager // interface com.deepin.lastore.Manager
	interfaceUpdater // interface com.deepin.lastore.Updater
	proxy.ImplObject
}

func NewLastore(conn *dbus.Conn) Lastore {
	obj := new(objectLastore)
	obj.ImplObject.Init_(conn, "com.deepin.lastore", "/com/deepin/lastore")
	return obj
}

func (obj *objectLastore) Manager() manager {
	return &obj.interfaceManager
}

type manager interface {
	GoCleanArchives(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	CleanArchives(flags dbus.Flags) (dbus.ObjectPath, error)
	GoCleanJob(flags dbus.Flags, ch chan *dbus.Call, jobId string) *dbus.Call
	CleanJob(flags dbus.Flags, jobId string) error
	GoDistUpgrade(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	DistUpgrade(flags dbus.Flags) (dbus.ObjectPath, error)
	GoFixError(flags dbus.Flags, ch chan *dbus.Call, errType string) *dbus.Call
	FixError(flags dbus.Flags, errType string) (dbus.ObjectPath, error)
	GoInstallPackage(flags dbus.Flags, ch chan *dbus.Call, jobName string, packages string) *dbus.Call
	InstallPackage(flags dbus.Flags, jobName string, packages string) (dbus.ObjectPath, error)
	GoPackageDesktopPath(flags dbus.Flags, ch chan *dbus.Call, pkgId string) *dbus.Call
	PackageDesktopPath(flags dbus.Flags, pkgId string) (string, error)
	GoPackageExists(flags dbus.Flags, ch chan *dbus.Call, pkgId string) *dbus.Call
	PackageExists(flags dbus.Flags, pkgId string) (bool, error)
	GoPackageInstallable(flags dbus.Flags, ch chan *dbus.Call, pkgId string) *dbus.Call
	PackageInstallable(flags dbus.Flags, pkgId string) (bool, error)
	GoPackagesDownloadSize(flags dbus.Flags, ch chan *dbus.Call, packages []string) *dbus.Call
	PackagesDownloadSize(flags dbus.Flags, packages []string) (int64, error)
	GoPauseJob(flags dbus.Flags, ch chan *dbus.Call, jobId string) *dbus.Call
	PauseJob(flags dbus.Flags, jobId string) error
	GoPrepareDistUpgrade(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	PrepareDistUpgrade(flags dbus.Flags) (dbus.ObjectPath, error)
	GoRemovePackage(flags dbus.Flags, ch chan *dbus.Call, jobName string, packages string) *dbus.Call
	RemovePackage(flags dbus.Flags, jobName string, packages string) (dbus.ObjectPath, error)
	GoSetAutoClean(flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call
	SetAutoClean(flags dbus.Flags, enable bool) error
	GoSetLogger(flags dbus.Flags, ch chan *dbus.Call, levels string, format string, output string) *dbus.Call
	SetLogger(flags dbus.Flags, levels string, format string, output string) error
	GoSetRegion(flags dbus.Flags, ch chan *dbus.Call, region string) *dbus.Call
	SetRegion(flags dbus.Flags, region string) error
	GoStartJob(flags dbus.Flags, ch chan *dbus.Call, jobId string) *dbus.Call
	StartJob(flags dbus.Flags, jobId string) error
	GoUpdatePackage(flags dbus.Flags, ch chan *dbus.Call, jobName string, packages string) *dbus.Call
	UpdatePackage(flags dbus.Flags, jobName string, packages string) (dbus.ObjectPath, error)
	GoUpdateSource(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	UpdateSource(flags dbus.Flags) (dbus.ObjectPath, error)
	JobList() proxy.PropObjectPathArray
	SystemArchitectures() proxy.PropStringArray
	UpgradableApps() proxy.PropStringArray
	SystemOnChanging() proxy.PropBool
	AutoClean() proxy.PropBool
}

type interfaceManager struct{}

func (v *interfaceManager) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceManager) GetInterfaceName_() string {
	return "com.deepin.lastore.Manager"
}

// method CleanArchives

func (v *interfaceManager) GoCleanArchives(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CleanArchives", flags, ch)
}

func (*interfaceManager) StoreCleanArchives(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *interfaceManager) CleanArchives(flags dbus.Flags) (dbus.ObjectPath, error) {
	return v.StoreCleanArchives(
		<-v.GoCleanArchives(flags, make(chan *dbus.Call, 1)).Done)
}

// method CleanJob

func (v *interfaceManager) GoCleanJob(flags dbus.Flags, ch chan *dbus.Call, jobId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CleanJob", flags, ch, jobId)
}

func (v *interfaceManager) CleanJob(flags dbus.Flags, jobId string) error {
	return (<-v.GoCleanJob(flags, make(chan *dbus.Call, 1), jobId).Done).Err
}

// method DistUpgrade

func (v *interfaceManager) GoDistUpgrade(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DistUpgrade", flags, ch)
}

func (*interfaceManager) StoreDistUpgrade(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *interfaceManager) DistUpgrade(flags dbus.Flags) (dbus.ObjectPath, error) {
	return v.StoreDistUpgrade(
		<-v.GoDistUpgrade(flags, make(chan *dbus.Call, 1)).Done)
}

// method FixError

func (v *interfaceManager) GoFixError(flags dbus.Flags, ch chan *dbus.Call, errType string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FixError", flags, ch, errType)
}

func (*interfaceManager) StoreFixError(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *interfaceManager) FixError(flags dbus.Flags, errType string) (dbus.ObjectPath, error) {
	return v.StoreFixError(
		<-v.GoFixError(flags, make(chan *dbus.Call, 1), errType).Done)
}

// method InstallPackage

func (v *interfaceManager) GoInstallPackage(flags dbus.Flags, ch chan *dbus.Call, jobName string, packages string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".InstallPackage", flags, ch, jobName, packages)
}

func (*interfaceManager) StoreInstallPackage(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *interfaceManager) InstallPackage(flags dbus.Flags, jobName string, packages string) (dbus.ObjectPath, error) {
	return v.StoreInstallPackage(
		<-v.GoInstallPackage(flags, make(chan *dbus.Call, 1), jobName, packages).Done)
}

// method PackageDesktopPath

func (v *interfaceManager) GoPackageDesktopPath(flags dbus.Flags, ch chan *dbus.Call, pkgId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PackageDesktopPath", flags, ch, pkgId)
}

func (*interfaceManager) StorePackageDesktopPath(call *dbus.Call) (desktopPath string, err error) {
	err = call.Store(&desktopPath)
	return
}

func (v *interfaceManager) PackageDesktopPath(flags dbus.Flags, pkgId string) (string, error) {
	return v.StorePackageDesktopPath(
		<-v.GoPackageDesktopPath(flags, make(chan *dbus.Call, 1), pkgId).Done)
}

// method PackageExists

func (v *interfaceManager) GoPackageExists(flags dbus.Flags, ch chan *dbus.Call, pkgId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PackageExists", flags, ch, pkgId)
}

func (*interfaceManager) StorePackageExists(call *dbus.Call) (exist bool, err error) {
	err = call.Store(&exist)
	return
}

func (v *interfaceManager) PackageExists(flags dbus.Flags, pkgId string) (bool, error) {
	return v.StorePackageExists(
		<-v.GoPackageExists(flags, make(chan *dbus.Call, 1), pkgId).Done)
}

// method PackageInstallable

func (v *interfaceManager) GoPackageInstallable(flags dbus.Flags, ch chan *dbus.Call, pkgId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PackageInstallable", flags, ch, pkgId)
}

func (*interfaceManager) StorePackageInstallable(call *dbus.Call) (installable bool, err error) {
	err = call.Store(&installable)
	return
}

func (v *interfaceManager) PackageInstallable(flags dbus.Flags, pkgId string) (bool, error) {
	return v.StorePackageInstallable(
		<-v.GoPackageInstallable(flags, make(chan *dbus.Call, 1), pkgId).Done)
}

// method PackagesDownloadSize

func (v *interfaceManager) GoPackagesDownloadSize(flags dbus.Flags, ch chan *dbus.Call, packages []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PackagesDownloadSize", flags, ch, packages)
}

func (*interfaceManager) StorePackagesDownloadSize(call *dbus.Call) (size int64, err error) {
	err = call.Store(&size)
	return
}

func (v *interfaceManager) PackagesDownloadSize(flags dbus.Flags, packages []string) (int64, error) {
	return v.StorePackagesDownloadSize(
		<-v.GoPackagesDownloadSize(flags, make(chan *dbus.Call, 1), packages).Done)
}

// method PauseJob

func (v *interfaceManager) GoPauseJob(flags dbus.Flags, ch chan *dbus.Call, jobId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PauseJob", flags, ch, jobId)
}

func (v *interfaceManager) PauseJob(flags dbus.Flags, jobId string) error {
	return (<-v.GoPauseJob(flags, make(chan *dbus.Call, 1), jobId).Done).Err
}

// method PrepareDistUpgrade

func (v *interfaceManager) GoPrepareDistUpgrade(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PrepareDistUpgrade", flags, ch)
}

func (*interfaceManager) StorePrepareDistUpgrade(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *interfaceManager) PrepareDistUpgrade(flags dbus.Flags) (dbus.ObjectPath, error) {
	return v.StorePrepareDistUpgrade(
		<-v.GoPrepareDistUpgrade(flags, make(chan *dbus.Call, 1)).Done)
}

// method RemovePackage

func (v *interfaceManager) GoRemovePackage(flags dbus.Flags, ch chan *dbus.Call, jobName string, packages string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemovePackage", flags, ch, jobName, packages)
}

func (*interfaceManager) StoreRemovePackage(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *interfaceManager) RemovePackage(flags dbus.Flags, jobName string, packages string) (dbus.ObjectPath, error) {
	return v.StoreRemovePackage(
		<-v.GoRemovePackage(flags, make(chan *dbus.Call, 1), jobName, packages).Done)
}

// method SetAutoClean

func (v *interfaceManager) GoSetAutoClean(flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAutoClean", flags, ch, enable)
}

func (v *interfaceManager) SetAutoClean(flags dbus.Flags, enable bool) error {
	return (<-v.GoSetAutoClean(flags, make(chan *dbus.Call, 1), enable).Done).Err
}

// method SetLogger

func (v *interfaceManager) GoSetLogger(flags dbus.Flags, ch chan *dbus.Call, levels string, format string, output string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLogger", flags, ch, levels, format, output)
}

func (v *interfaceManager) SetLogger(flags dbus.Flags, levels string, format string, output string) error {
	return (<-v.GoSetLogger(flags, make(chan *dbus.Call, 1), levels, format, output).Done).Err
}

// method SetRegion

func (v *interfaceManager) GoSetRegion(flags dbus.Flags, ch chan *dbus.Call, region string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetRegion", flags, ch, region)
}

func (v *interfaceManager) SetRegion(flags dbus.Flags, region string) error {
	return (<-v.GoSetRegion(flags, make(chan *dbus.Call, 1), region).Done).Err
}

// method StartJob

func (v *interfaceManager) GoStartJob(flags dbus.Flags, ch chan *dbus.Call, jobId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StartJob", flags, ch, jobId)
}

func (v *interfaceManager) StartJob(flags dbus.Flags, jobId string) error {
	return (<-v.GoStartJob(flags, make(chan *dbus.Call, 1), jobId).Done).Err
}

// method UpdatePackage

func (v *interfaceManager) GoUpdatePackage(flags dbus.Flags, ch chan *dbus.Call, jobName string, packages string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UpdatePackage", flags, ch, jobName, packages)
}

func (*interfaceManager) StoreUpdatePackage(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *interfaceManager) UpdatePackage(flags dbus.Flags, jobName string, packages string) (dbus.ObjectPath, error) {
	return v.StoreUpdatePackage(
		<-v.GoUpdatePackage(flags, make(chan *dbus.Call, 1), jobName, packages).Done)
}

// method UpdateSource

func (v *interfaceManager) GoUpdateSource(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UpdateSource", flags, ch)
}

func (*interfaceManager) StoreUpdateSource(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *interfaceManager) UpdateSource(flags dbus.Flags) (dbus.ObjectPath, error) {
	return v.StoreUpdateSource(
		<-v.GoUpdateSource(flags, make(chan *dbus.Call, 1)).Done)
}

// property JobList ao

func (v *interfaceManager) JobList() proxy.PropObjectPathArray {
	return &proxy.ImplPropObjectPathArray{
		Impl: v,
		Name: "JobList",
	}
}

// property SystemArchitectures as

func (v *interfaceManager) SystemArchitectures() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "SystemArchitectures",
	}
}

// property UpgradableApps as

func (v *interfaceManager) UpgradableApps() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "UpgradableApps",
	}
}

// property SystemOnChanging b

func (v *interfaceManager) SystemOnChanging() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "SystemOnChanging",
	}
}

// property AutoClean b

func (v *interfaceManager) AutoClean() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "AutoClean",
	}
}

func (obj *objectLastore) Updater() updater {
	return &obj.interfaceUpdater
}

type updater interface {
	GoApplicationUpdateInfos(flags dbus.Flags, ch chan *dbus.Call, lang string) *dbus.Call
	ApplicationUpdateInfos(flags dbus.Flags, lang string) ([][]interface{}, error)
	GoListMirrorSources(flags dbus.Flags, ch chan *dbus.Call, lang string) *dbus.Call
	ListMirrorSources(flags dbus.Flags, lang string) ([][]interface{}, error)
	GoRestoreSystemSource(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	RestoreSystemSource(flags dbus.Flags) error
	GoSetAutoCheckUpdates(flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call
	SetAutoCheckUpdates(flags dbus.Flags, enable bool) error
	GoSetAutoDownloadUpdates(flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call
	SetAutoDownloadUpdates(flags dbus.Flags, enable bool) error
	GoSetMirrorSource(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call
	SetMirrorSource(flags dbus.Flags, id string) error
	AutoCheckUpdates() proxy.PropBool
	AutoDownloadUpdates() proxy.PropBool
	MirrorSource() proxy.PropString
	UpdatableApps() proxy.PropStringArray
	UpdatablePackages() proxy.PropStringArray
}

type interfaceUpdater struct{}

func (v *interfaceUpdater) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceUpdater) GetInterfaceName_() string {
	return "com.deepin.lastore.Updater"
}

// method ApplicationUpdateInfos

func (v *interfaceUpdater) GoApplicationUpdateInfos(flags dbus.Flags, ch chan *dbus.Call, lang string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ApplicationUpdateInfos", flags, ch, lang)
}

func (*interfaceUpdater) StoreApplicationUpdateInfos(call *dbus.Call) (updateInfos [][]interface{}, err error) {
	err = call.Store(&updateInfos)
	return
}

func (v *interfaceUpdater) ApplicationUpdateInfos(flags dbus.Flags, lang string) ([][]interface{}, error) {
	return v.StoreApplicationUpdateInfos(
		<-v.GoApplicationUpdateInfos(flags, make(chan *dbus.Call, 1), lang).Done)
}

// method ListMirrorSources

func (v *interfaceUpdater) GoListMirrorSources(flags dbus.Flags, ch chan *dbus.Call, lang string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListMirrorSources", flags, ch, lang)
}

func (*interfaceUpdater) StoreListMirrorSources(call *dbus.Call) (mirrorSources [][]interface{}, err error) {
	err = call.Store(&mirrorSources)
	return
}

func (v *interfaceUpdater) ListMirrorSources(flags dbus.Flags, lang string) ([][]interface{}, error) {
	return v.StoreListMirrorSources(
		<-v.GoListMirrorSources(flags, make(chan *dbus.Call, 1), lang).Done)
}

// method RestoreSystemSource

func (v *interfaceUpdater) GoRestoreSystemSource(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RestoreSystemSource", flags, ch)
}

func (v *interfaceUpdater) RestoreSystemSource(flags dbus.Flags) error {
	return (<-v.GoRestoreSystemSource(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method SetAutoCheckUpdates

func (v *interfaceUpdater) GoSetAutoCheckUpdates(flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAutoCheckUpdates", flags, ch, enable)
}

func (v *interfaceUpdater) SetAutoCheckUpdates(flags dbus.Flags, enable bool) error {
	return (<-v.GoSetAutoCheckUpdates(flags, make(chan *dbus.Call, 1), enable).Done).Err
}

// method SetAutoDownloadUpdates

func (v *interfaceUpdater) GoSetAutoDownloadUpdates(flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAutoDownloadUpdates", flags, ch, enable)
}

func (v *interfaceUpdater) SetAutoDownloadUpdates(flags dbus.Flags, enable bool) error {
	return (<-v.GoSetAutoDownloadUpdates(flags, make(chan *dbus.Call, 1), enable).Done).Err
}

// method SetMirrorSource

func (v *interfaceUpdater) GoSetMirrorSource(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetMirrorSource", flags, ch, id)
}

func (v *interfaceUpdater) SetMirrorSource(flags dbus.Flags, id string) error {
	return (<-v.GoSetMirrorSource(flags, make(chan *dbus.Call, 1), id).Done).Err
}

// property AutoCheckUpdates b

func (v *interfaceUpdater) AutoCheckUpdates() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "AutoCheckUpdates",
	}
}

// property AutoDownloadUpdates b

func (v *interfaceUpdater) AutoDownloadUpdates() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "AutoDownloadUpdates",
	}
}

// property MirrorSource s

func (v *interfaceUpdater) MirrorSource() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "MirrorSource",
	}
}

// property UpdatableApps as

func (v *interfaceUpdater) UpdatableApps() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "UpdatableApps",
	}
}

// property UpdatablePackages as

func (v *interfaceUpdater) UpdatablePackages() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "UpdatablePackages",
	}
}

type Job interface {
	job // interface com.deepin.lastore.Job
	proxy.Object
}

type objectJob struct {
	interfaceJob // interface com.deepin.lastore.Job
	proxy.ImplObject
}

func NewJob(conn *dbus.Conn, path dbus.ObjectPath) (Job, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectJob)
	obj.ImplObject.Init_(conn, "com.deepin.lastore", path)
	return obj, nil
}

type job interface {
	GoString(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	String(flags dbus.Flags) (string, error)
	Id() proxy.PropString
	Name() proxy.PropString
	Packages() proxy.PropStringArray
	CreateTime() proxy.PropInt64
	Type() proxy.PropString
	Status() proxy.PropString
	Progress() proxy.PropDouble
	Description() proxy.PropString
	Speed() proxy.PropInt64
	DownloadSize() proxy.PropInt64
	Cancelable() proxy.PropBool
}

type interfaceJob struct{}

func (v *interfaceJob) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceJob) GetInterfaceName_() string {
	return "com.deepin.lastore.Job"
}

// method String

func (v *interfaceJob) GoString(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".String", flags, ch)
}

func (*interfaceJob) StoreString(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceJob) String(flags dbus.Flags) (string, error) {
	return v.StoreString(
		<-v.GoString(flags, make(chan *dbus.Call, 1)).Done)
}

// property Id s

func (v *interfaceJob) Id() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Id",
	}
}

// property Name s

func (v *interfaceJob) Name() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Name",
	}
}

// property Packages as

func (v *interfaceJob) Packages() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "Packages",
	}
}

// property CreateTime x

func (v *interfaceJob) CreateTime() proxy.PropInt64 {
	return &proxy.ImplPropInt64{
		Impl: v,
		Name: "CreateTime",
	}
}

// property Type s

func (v *interfaceJob) Type() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Type",
	}
}

// property Status s

func (v *interfaceJob) Status() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Status",
	}
}

// property Progress d

func (v *interfaceJob) Progress() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "Progress",
	}
}

// property Description s

func (v *interfaceJob) Description() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Description",
	}
}

// property Speed x

func (v *interfaceJob) Speed() proxy.PropInt64 {
	return &proxy.ImplPropInt64{
		Impl: v,
		Name: "Speed",
	}
}

// property DownloadSize x

func (v *interfaceJob) DownloadSize() proxy.PropInt64 {
	return &proxy.ImplPropInt64{
		Impl: v,
		Name: "DownloadSize",
	}
}

// property Cancelable b

func (v *interfaceJob) Cancelable() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "Cancelable",
	}
}
