// Code generated by "./generator ./org.mpris.mediaplayer2"; DO NOT EDIT.

package mediaplayer2

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type MediaPlayer interface {
	MediaPlayer2() mediaPlayer // interface org.mpris.MediaPlayer2
	Player() player            // interface org.mpris.MediaPlayer2.Player
	proxy.Object
}

type objectMediaPlayer struct {
	interfaceMediaPlayer // interface org.mpris.MediaPlayer2
	interfacePlayer      // interface org.mpris.MediaPlayer2.Player
	proxy.ImplObject
}

func NewMediaPlayer(conn *dbus.Conn, serviceName string) MediaPlayer {
	obj := new(objectMediaPlayer)
	obj.ImplObject.Init_(conn, serviceName, "/org/mpris/MediaPlayer2")
	return obj
}

func (obj *objectMediaPlayer) MediaPlayer2() mediaPlayer {
	return &obj.interfaceMediaPlayer
}

type mediaPlayer interface {
	GoQuit(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Quit(flags dbus.Flags) error
	GoRaise(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Raise(flags dbus.Flags) error
	CanQuit() proxy.PropBool
	CanRaise() proxy.PropBool
	DesktopEntry() proxy.PropString
	Identity() proxy.PropString
}

type interfaceMediaPlayer struct{}

func (v *interfaceMediaPlayer) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceMediaPlayer) GetInterfaceName_() string {
	return "org.mpris.MediaPlayer2"
}

// method Quit

func (v *interfaceMediaPlayer) GoQuit(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Quit", flags, ch)
}

func (v *interfaceMediaPlayer) Quit(flags dbus.Flags) error {
	return (<-v.GoQuit(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Raise

func (v *interfaceMediaPlayer) GoRaise(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Raise", flags, ch)
}

func (v *interfaceMediaPlayer) Raise(flags dbus.Flags) error {
	return (<-v.GoRaise(flags, make(chan *dbus.Call, 1)).Done).Err
}

// property CanQuit b

func (v *interfaceMediaPlayer) CanQuit() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "CanQuit",
	}
}

// property CanRaise b

func (v *interfaceMediaPlayer) CanRaise() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "CanRaise",
	}
}

// property DesktopEntry s

func (v *interfaceMediaPlayer) DesktopEntry() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "DesktopEntry",
	}
}

// property Identity s

func (v *interfaceMediaPlayer) Identity() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Identity",
	}
}

func (obj *objectMediaPlayer) Player() player {
	return &obj.interfacePlayer
}

type player interface {
	GoNext(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Next(flags dbus.Flags) error
	GoPause(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Pause(flags dbus.Flags) error
	GoPlay(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Play(flags dbus.Flags) error
	GoPlayPause(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	PlayPause(flags dbus.Flags) error
	GoPrevious(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Previous(flags dbus.Flags) error
	GoSeek(flags dbus.Flags, ch chan *dbus.Call, Offset int64) *dbus.Call
	Seek(flags dbus.Flags, Offset int64) error
	GoSetPosition(flags dbus.Flags, ch chan *dbus.Call, TrackId dbus.ObjectPath, Position int64) *dbus.Call
	SetPosition(flags dbus.Flags, TrackId dbus.ObjectPath, Position int64) error
	GoStop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Stop(flags dbus.Flags) error
	ConnectSeeked(cb func(Position int64)) (dbusutil.SignalHandlerId, error)
	CanControl() proxy.PropBool
	CanGoNext() proxy.PropBool
	CanGoPrevious() proxy.PropBool
	CanPause() proxy.PropBool
	CanPlay() proxy.PropBool
	CanSeek() proxy.PropBool
	Metadata() PropPlayerMetadata
	PlaybackStatus() proxy.PropString
	Position() proxy.PropInt64
	Volume() proxy.PropDouble
}

type interfacePlayer struct{}

func (v *interfacePlayer) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfacePlayer) GetInterfaceName_() string {
	return "org.mpris.MediaPlayer2.Player"
}

// method Next

func (v *interfacePlayer) GoNext(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Next", flags, ch)
}

func (v *interfacePlayer) Next(flags dbus.Flags) error {
	return (<-v.GoNext(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Pause

func (v *interfacePlayer) GoPause(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Pause", flags, ch)
}

func (v *interfacePlayer) Pause(flags dbus.Flags) error {
	return (<-v.GoPause(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Play

func (v *interfacePlayer) GoPlay(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Play", flags, ch)
}

func (v *interfacePlayer) Play(flags dbus.Flags) error {
	return (<-v.GoPlay(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method PlayPause

func (v *interfacePlayer) GoPlayPause(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PlayPause", flags, ch)
}

func (v *interfacePlayer) PlayPause(flags dbus.Flags) error {
	return (<-v.GoPlayPause(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Previous

func (v *interfacePlayer) GoPrevious(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Previous", flags, ch)
}

func (v *interfacePlayer) Previous(flags dbus.Flags) error {
	return (<-v.GoPrevious(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Seek

func (v *interfacePlayer) GoSeek(flags dbus.Flags, ch chan *dbus.Call, Offset int64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Seek", flags, ch, Offset)
}

func (v *interfacePlayer) Seek(flags dbus.Flags, Offset int64) error {
	return (<-v.GoSeek(flags, make(chan *dbus.Call, 1), Offset).Done).Err
}

// method SetPosition

func (v *interfacePlayer) GoSetPosition(flags dbus.Flags, ch chan *dbus.Call, TrackId dbus.ObjectPath, Position int64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetPosition", flags, ch, TrackId, Position)
}

func (v *interfacePlayer) SetPosition(flags dbus.Flags, TrackId dbus.ObjectPath, Position int64) error {
	return (<-v.GoSetPosition(flags, make(chan *dbus.Call, 1), TrackId, Position).Done).Err
}

// method Stop

func (v *interfacePlayer) GoStop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Stop", flags, ch)
}

func (v *interfacePlayer) Stop(flags dbus.Flags) error {
	return (<-v.GoStop(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal Seeked

func (v *interfacePlayer) ConnectSeeked(cb func(Position int64)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Seeked", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Seeked",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var Position int64
		err := dbus.Store(sig.Body, &Position)
		if err == nil {
			cb(Position)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property CanControl b

func (v *interfacePlayer) CanControl() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "CanControl",
	}
}

// property CanGoNext b

func (v *interfacePlayer) CanGoNext() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "CanGoNext",
	}
}

// property CanGoPrevious b

func (v *interfacePlayer) CanGoPrevious() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "CanGoPrevious",
	}
}

// property CanPause b

func (v *interfacePlayer) CanPause() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "CanPause",
	}
}

// property CanPlay b

func (v *interfacePlayer) CanPlay() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "CanPlay",
	}
}

// property CanSeek b

func (v *interfacePlayer) CanSeek() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "CanSeek",
	}
}

type PropPlayerMetadata interface {
	Get(flags dbus.Flags) (value map[string]dbus.Variant, err error)
	Set(flags dbus.Flags, value map[string]dbus.Variant) error
	ConnectChanged(cb func(hasValue bool, value map[string]dbus.Variant)) error
}

type implPropPlayerMetadata struct {
	Impl proxy.Implementer
	Name string
}

func (p implPropPlayerMetadata) Get(flags dbus.Flags) (value map[string]dbus.Variant, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p implPropPlayerMetadata) Set(flags dbus.Flags, value map[string]dbus.Variant) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p implPropPlayerMetadata) ConnectChanged(cb func(hasValue bool, value map[string]dbus.Variant)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v map[string]dbus.Variant
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}

// property Metadata a{sv}

func (v *interfacePlayer) Metadata() PropPlayerMetadata {
	return &implPropPlayerMetadata{
		Impl: v,
		Name: "Metadata",
	}
}

// property PlaybackStatus s

func (v *interfacePlayer) PlaybackStatus() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "PlaybackStatus",
	}
}

// property Position x

func (v *interfacePlayer) Position() proxy.PropInt64 {
	return &proxy.ImplPropInt64{
		Impl: v,
		Name: "Position",
	}
}

// property Volume d

func (v *interfacePlayer) Volume() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "Volume",
	}
}
