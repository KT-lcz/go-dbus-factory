// Code generated by "./generator ./com.deepin.dde.osd"; DO NOT EDIT.

package osd

import (
	"unsafe"

	"github.com/godbus/dbus"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type OSD interface {
	osd // interface com.deepin.dde.osd
	proxy.Object
}

type objectOSD struct {
	interfaceOsd // interface com.deepin.dde.osd
	proxy.ImplObject
}

func NewOSD(conn *dbus.Conn) OSD {
	obj := new(objectOSD)
	obj.ImplObject.Init_(conn, "com.deepin.dde.osd", "/")
	return obj
}

type osd interface {
	GoShowOSD(flags dbus.Flags, ch chan *dbus.Call, osd string) *dbus.Call
	ShowOSD(flags dbus.Flags, osd string) error
}

type interfaceOsd struct{}

func (v *interfaceOsd) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceOsd) GetInterfaceName_() string {
	return "com.deepin.dde.osd"
}

// method ShowOSD

func (v *interfaceOsd) GoShowOSD(flags dbus.Flags, ch chan *dbus.Call, osd string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ShowOSD", flags, ch, osd)
}

func (v *interfaceOsd) ShowOSD(flags dbus.Flags, osd string) error {
	return (<-v.GoShowOSD(flags, make(chan *dbus.Call, 1), osd).Done).Err
}
