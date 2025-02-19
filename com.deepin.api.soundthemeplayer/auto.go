// Code generated by "./generator ./com.deepin.api.soundthemeplayer"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package soundthemeplayer

import (
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type SoundThemePlayer interface {
	soundThemePlayer // interface com.deepin.api.SoundThemePlayer
	proxy.Object
}

type objectSoundThemePlayer struct {
	interfaceSoundThemePlayer // interface com.deepin.api.SoundThemePlayer
	proxy.ImplObject
}

func NewSoundThemePlayer(conn *dbus.Conn) SoundThemePlayer {
	obj := new(objectSoundThemePlayer)
	obj.ImplObject.Init_(conn, "com.deepin.api.SoundThemePlayer", "/com/deepin/api/SoundThemePlayer")
	return obj
}

type soundThemePlayer interface {
	GoEnableSound(flags dbus.Flags, ch chan *dbus.Call, name string, enabled bool) *dbus.Call
	EnableSound(flags dbus.Flags, name string, enabled bool) error
	GoEnableSoundDesktopLogin(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call
	EnableSoundDesktopLogin(flags dbus.Flags, enabled bool) error
	GoPlay(flags dbus.Flags, ch chan *dbus.Call, theme string, event string, device string) *dbus.Call
	Play(flags dbus.Flags, theme string, event string, device string) error
	GoPlaySoundDesktopLogin(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	PlaySoundDesktopLogin(flags dbus.Flags) error
	GoPrepareShutdownSound(flags dbus.Flags, ch chan *dbus.Call, uid int32) *dbus.Call
	PrepareShutdownSound(flags dbus.Flags, uid int32) error
	GoSaveAudioState(flags dbus.Flags, ch chan *dbus.Call, activePlayback map[string]dbus.Variant) *dbus.Call
	SaveAudioState(flags dbus.Flags, activePlayback map[string]dbus.Variant) error
	GoSetSoundTheme(flags dbus.Flags, ch chan *dbus.Call, theme string) *dbus.Call
	SetSoundTheme(flags dbus.Flags, theme string) error
}

type interfaceSoundThemePlayer struct{}

func (v *interfaceSoundThemePlayer) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceSoundThemePlayer) GetInterfaceName_() string {
	return "com.deepin.api.SoundThemePlayer"
}

// method EnableSound

func (v *interfaceSoundThemePlayer) GoEnableSound(flags dbus.Flags, ch chan *dbus.Call, name string, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableSound", flags, ch, name, enabled)
}

func (v *interfaceSoundThemePlayer) EnableSound(flags dbus.Flags, name string, enabled bool) error {
	return (<-v.GoEnableSound(flags, make(chan *dbus.Call, 1), name, enabled).Done).Err
}

// method EnableSoundDesktopLogin

func (v *interfaceSoundThemePlayer) GoEnableSoundDesktopLogin(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableSoundDesktopLogin", flags, ch, enabled)
}

func (v *interfaceSoundThemePlayer) EnableSoundDesktopLogin(flags dbus.Flags, enabled bool) error {
	return (<-v.GoEnableSoundDesktopLogin(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

// method Play

func (v *interfaceSoundThemePlayer) GoPlay(flags dbus.Flags, ch chan *dbus.Call, theme string, event string, device string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Play", flags, ch, theme, event, device)
}

func (v *interfaceSoundThemePlayer) Play(flags dbus.Flags, theme string, event string, device string) error {
	return (<-v.GoPlay(flags, make(chan *dbus.Call, 1), theme, event, device).Done).Err
}

// method PlaySoundDesktopLogin

func (v *interfaceSoundThemePlayer) GoPlaySoundDesktopLogin(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PlaySoundDesktopLogin", flags, ch)
}

func (v *interfaceSoundThemePlayer) PlaySoundDesktopLogin(flags dbus.Flags) error {
	return (<-v.GoPlaySoundDesktopLogin(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method PrepareShutdownSound

func (v *interfaceSoundThemePlayer) GoPrepareShutdownSound(flags dbus.Flags, ch chan *dbus.Call, uid int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PrepareShutdownSound", flags, ch, uid)
}

func (v *interfaceSoundThemePlayer) PrepareShutdownSound(flags dbus.Flags, uid int32) error {
	return (<-v.GoPrepareShutdownSound(flags, make(chan *dbus.Call, 1), uid).Done).Err
}

// method SaveAudioState

func (v *interfaceSoundThemePlayer) GoSaveAudioState(flags dbus.Flags, ch chan *dbus.Call, activePlayback map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SaveAudioState", flags, ch, activePlayback)
}

func (v *interfaceSoundThemePlayer) SaveAudioState(flags dbus.Flags, activePlayback map[string]dbus.Variant) error {
	return (<-v.GoSaveAudioState(flags, make(chan *dbus.Call, 1), activePlayback).Done).Err
}

// method SetSoundTheme

func (v *interfaceSoundThemePlayer) GoSetSoundTheme(flags dbus.Flags, ch chan *dbus.Call, theme string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetSoundTheme", flags, ch, theme)
}

func (v *interfaceSoundThemePlayer) SetSoundTheme(flags dbus.Flags, theme string) error {
	return (<-v.GoSetSoundTheme(flags, make(chan *dbus.Call, 1), theme).Done).Err
}
