// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package audio

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/audio/bindings"
)

type StreamType uint32

const (
	_ StreamType = iota

	StreamType_INPUT
	StreamType_OUTPUT
)

func (StreamType) FromRef(str js.Ref) StreamType {
	return StreamType(bindings.ConstOfStreamType(str))
}

func (x StreamType) String() (string, bool) {
	switch x {
	case StreamType_INPUT:
		return "INPUT", true
	case StreamType_OUTPUT:
		return "OUTPUT", true
	default:
		return "", false
	}
}

type DeviceType uint32

const (
	_ DeviceType = iota

	DeviceType_HEADPHONE
	DeviceType_MIC
	DeviceType_USB
	DeviceType_BLUETOOTH
	DeviceType_HDMI
	DeviceType_INTERNAL_SPEAKER
	DeviceType_INTERNAL_MIC
	DeviceType_FRONT_MIC
	DeviceType_REAR_MIC
	DeviceType_KEYBOARD_MIC
	DeviceType_HOTWORD
	DeviceType_LINEOUT
	DeviceType_POST_MIX_LOOPBACK
	DeviceType_POST_DSP_LOOPBACK
	DeviceType_ALSA_LOOPBACK
	DeviceType_OTHER
)

func (DeviceType) FromRef(str js.Ref) DeviceType {
	return DeviceType(bindings.ConstOfDeviceType(str))
}

func (x DeviceType) String() (string, bool) {
	switch x {
	case DeviceType_HEADPHONE:
		return "HEADPHONE", true
	case DeviceType_MIC:
		return "MIC", true
	case DeviceType_USB:
		return "USB", true
	case DeviceType_BLUETOOTH:
		return "BLUETOOTH", true
	case DeviceType_HDMI:
		return "HDMI", true
	case DeviceType_INTERNAL_SPEAKER:
		return "INTERNAL_SPEAKER", true
	case DeviceType_INTERNAL_MIC:
		return "INTERNAL_MIC", true
	case DeviceType_FRONT_MIC:
		return "FRONT_MIC", true
	case DeviceType_REAR_MIC:
		return "REAR_MIC", true
	case DeviceType_KEYBOARD_MIC:
		return "KEYBOARD_MIC", true
	case DeviceType_HOTWORD:
		return "HOTWORD", true
	case DeviceType_LINEOUT:
		return "LINEOUT", true
	case DeviceType_POST_MIX_LOOPBACK:
		return "POST_MIX_LOOPBACK", true
	case DeviceType_POST_DSP_LOOPBACK:
		return "POST_DSP_LOOPBACK", true
	case DeviceType_ALSA_LOOPBACK:
		return "ALSA_LOOPBACK", true
	case DeviceType_OTHER:
		return "OTHER", true
	default:
		return "", false
	}
}

type AudioDeviceInfo struct {
	// Id is "AudioDeviceInfo.id"
	//
	// Optional
	Id js.String
	// StreamType is "AudioDeviceInfo.streamType"
	//
	// Optional
	StreamType StreamType
	// DeviceType is "AudioDeviceInfo.deviceType"
	//
	// Optional
	DeviceType DeviceType
	// DisplayName is "AudioDeviceInfo.displayName"
	//
	// Optional
	DisplayName js.String
	// DeviceName is "AudioDeviceInfo.deviceName"
	//
	// Optional
	DeviceName js.String
	// IsActive is "AudioDeviceInfo.isActive"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsActive MUST be set to true to make this field effective.
	IsActive bool
	// Level is "AudioDeviceInfo.level"
	//
	// Optional
	//
	// NOTE: FFI_USE_Level MUST be set to true to make this field effective.
	Level int32
	// StableDeviceId is "AudioDeviceInfo.stableDeviceId"
	//
	// Optional
	StableDeviceId js.String

	FFI_USE_IsActive bool // for IsActive.
	FFI_USE_Level    bool // for Level.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioDeviceInfo with all fields set.
func (p AudioDeviceInfo) FromRef(ref js.Ref) AudioDeviceInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioDeviceInfo in the application heap.
func (p AudioDeviceInfo) New() js.Ref {
	return bindings.AudioDeviceInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AudioDeviceInfo) UpdateFrom(ref js.Ref) {
	bindings.AudioDeviceInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioDeviceInfo) Update(ref js.Ref) {
	bindings.AudioDeviceInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioDeviceInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.DisplayName.Ref(),
		p.DeviceName.Ref(),
		p.StableDeviceId.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.DisplayName = p.DisplayName.FromRef(js.Undefined)
	p.DeviceName = p.DeviceName.FromRef(js.Undefined)
	p.StableDeviceId = p.StableDeviceId.FromRef(js.Undefined)
}

type DeviceFilter struct {
	// StreamTypes is "DeviceFilter.streamTypes"
	//
	// Optional
	StreamTypes js.Array[StreamType]
	// IsActive is "DeviceFilter.isActive"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsActive MUST be set to true to make this field effective.
	IsActive bool

	FFI_USE_IsActive bool // for IsActive.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DeviceFilter with all fields set.
func (p DeviceFilter) FromRef(ref js.Ref) DeviceFilter {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DeviceFilter in the application heap.
func (p DeviceFilter) New() js.Ref {
	return bindings.DeviceFilterJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DeviceFilter) UpdateFrom(ref js.Ref) {
	bindings.DeviceFilterJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DeviceFilter) Update(ref js.Ref) {
	bindings.DeviceFilterJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DeviceFilter) FreeMembers(recursive bool) {
	js.Free(
		p.StreamTypes.Ref(),
	)
	p.StreamTypes = p.StreamTypes.FromRef(js.Undefined)
}

type DeviceIdLists struct {
	// Input is "DeviceIdLists.input"
	//
	// Optional
	Input js.Array[js.String]
	// Output is "DeviceIdLists.output"
	//
	// Optional
	Output js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DeviceIdLists with all fields set.
func (p DeviceIdLists) FromRef(ref js.Ref) DeviceIdLists {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DeviceIdLists in the application heap.
func (p DeviceIdLists) New() js.Ref {
	return bindings.DeviceIdListsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DeviceIdLists) UpdateFrom(ref js.Ref) {
	bindings.DeviceIdListsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DeviceIdLists) Update(ref js.Ref) {
	bindings.DeviceIdListsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DeviceIdLists) FreeMembers(recursive bool) {
	js.Free(
		p.Input.Ref(),
		p.Output.Ref(),
	)
	p.Input = p.Input.FromRef(js.Undefined)
	p.Output = p.Output.FromRef(js.Undefined)
}

type DeviceProperties struct {
	// Level is "DeviceProperties.level"
	//
	// Optional
	//
	// NOTE: FFI_USE_Level MUST be set to true to make this field effective.
	Level int32

	FFI_USE_Level bool // for Level.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DeviceProperties with all fields set.
func (p DeviceProperties) FromRef(ref js.Ref) DeviceProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DeviceProperties in the application heap.
func (p DeviceProperties) New() js.Ref {
	return bindings.DevicePropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DeviceProperties) UpdateFrom(ref js.Ref) {
	bindings.DevicePropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DeviceProperties) Update(ref js.Ref) {
	bindings.DevicePropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DeviceProperties) FreeMembers(recursive bool) {
}

type EmptyCallbackFunc func(this js.Ref) js.Ref

func (fn EmptyCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn EmptyCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type EmptyCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *EmptyCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *EmptyCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDevicesCallbackFunc func(this js.Ref, devices js.Array[AudioDeviceInfo]) js.Ref

func (fn GetDevicesCallbackFunc) Register() js.Func[func(devices js.Array[AudioDeviceInfo])] {
	return js.RegisterCallback[func(devices js.Array[AudioDeviceInfo])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDevicesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[AudioDeviceInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDevicesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, devices js.Array[AudioDeviceInfo]) js.Ref
	Arg T
}

func (cb *GetDevicesCallback[T]) Register() js.Func[func(devices js.Array[AudioDeviceInfo])] {
	return js.RegisterCallback[func(devices js.Array[AudioDeviceInfo])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDevicesCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Array[AudioDeviceInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetMuteCallbackFunc func(this js.Ref, value bool) js.Ref

func (fn GetMuteCallbackFunc) Register() js.Func[func(value bool)] {
	return js.RegisterCallback[func(value bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetMuteCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetMuteCallback[T any] struct {
	Fn  func(arg T, this js.Ref, value bool) js.Ref
	Arg T
}

func (cb *GetMuteCallback[T]) Register() js.Func[func(value bool)] {
	return js.RegisterCallback[func(value bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetMuteCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type LevelChangedEvent struct {
	// DeviceId is "LevelChangedEvent.deviceId"
	//
	// Optional
	DeviceId js.String
	// Level is "LevelChangedEvent.level"
	//
	// Optional
	//
	// NOTE: FFI_USE_Level MUST be set to true to make this field effective.
	Level int32

	FFI_USE_Level bool // for Level.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LevelChangedEvent with all fields set.
func (p LevelChangedEvent) FromRef(ref js.Ref) LevelChangedEvent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new LevelChangedEvent in the application heap.
func (p LevelChangedEvent) New() js.Ref {
	return bindings.LevelChangedEventJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *LevelChangedEvent) UpdateFrom(ref js.Ref) {
	bindings.LevelChangedEventJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *LevelChangedEvent) Update(ref js.Ref) {
	bindings.LevelChangedEventJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *LevelChangedEvent) FreeMembers(recursive bool) {
	js.Free(
		p.DeviceId.Ref(),
	)
	p.DeviceId = p.DeviceId.FromRef(js.Undefined)
}

type MuteChangedEvent struct {
	// StreamType is "MuteChangedEvent.streamType"
	//
	// Optional
	StreamType StreamType
	// IsMuted is "MuteChangedEvent.isMuted"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsMuted MUST be set to true to make this field effective.
	IsMuted bool

	FFI_USE_IsMuted bool // for IsMuted.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MuteChangedEvent with all fields set.
func (p MuteChangedEvent) FromRef(ref js.Ref) MuteChangedEvent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MuteChangedEvent in the application heap.
func (p MuteChangedEvent) New() js.Ref {
	return bindings.MuteChangedEventJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MuteChangedEvent) UpdateFrom(ref js.Ref) {
	bindings.MuteChangedEventJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MuteChangedEvent) Update(ref js.Ref) {
	bindings.MuteChangedEventJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MuteChangedEvent) FreeMembers(recursive bool) {
}

// HasFuncGetDevices returns true if the function "WEBEXT.audio.getDevices" exists.
func HasFuncGetDevices() bool {
	return js.True == bindings.HasFuncGetDevices()
}

// FuncGetDevices returns the function "WEBEXT.audio.getDevices".
func FuncGetDevices() (fn js.Func[func(filter DeviceFilter) js.Promise[js.Array[AudioDeviceInfo]]]) {
	bindings.FuncGetDevices(
		js.Pointer(&fn),
	)
	return
}

// GetDevices calls the function "WEBEXT.audio.getDevices" directly.
func GetDevices(filter DeviceFilter) (ret js.Promise[js.Array[AudioDeviceInfo]]) {
	bindings.CallGetDevices(
		js.Pointer(&ret),
		js.Pointer(&filter),
	)

	return
}

// TryGetDevices calls the function "WEBEXT.audio.getDevices"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDevices(filter DeviceFilter) (ret js.Promise[js.Array[AudioDeviceInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDevices(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&filter),
	)

	return
}

// HasFuncGetMute returns true if the function "WEBEXT.audio.getMute" exists.
func HasFuncGetMute() bool {
	return js.True == bindings.HasFuncGetMute()
}

// FuncGetMute returns the function "WEBEXT.audio.getMute".
func FuncGetMute() (fn js.Func[func(streamType StreamType) js.Promise[js.Boolean]]) {
	bindings.FuncGetMute(
		js.Pointer(&fn),
	)
	return
}

// GetMute calls the function "WEBEXT.audio.getMute" directly.
func GetMute(streamType StreamType) (ret js.Promise[js.Boolean]) {
	bindings.CallGetMute(
		js.Pointer(&ret),
		uint32(streamType),
	)

	return
}

// TryGetMute calls the function "WEBEXT.audio.getMute"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetMute(streamType StreamType) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetMute(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(streamType),
	)

	return
}

type OnDeviceListChangedEventCallbackFunc func(this js.Ref, devices js.Array[AudioDeviceInfo]) js.Ref

func (fn OnDeviceListChangedEventCallbackFunc) Register() js.Func[func(devices js.Array[AudioDeviceInfo])] {
	return js.RegisterCallback[func(devices js.Array[AudioDeviceInfo])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDeviceListChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[AudioDeviceInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnDeviceListChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, devices js.Array[AudioDeviceInfo]) js.Ref
	Arg T
}

func (cb *OnDeviceListChangedEventCallback[T]) Register() js.Func[func(devices js.Array[AudioDeviceInfo])] {
	return js.RegisterCallback[func(devices js.Array[AudioDeviceInfo])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDeviceListChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Array[AudioDeviceInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnDeviceListChanged returns true if the function "WEBEXT.audio.onDeviceListChanged.addListener" exists.
func HasFuncOnDeviceListChanged() bool {
	return js.True == bindings.HasFuncOnDeviceListChanged()
}

// FuncOnDeviceListChanged returns the function "WEBEXT.audio.onDeviceListChanged.addListener".
func FuncOnDeviceListChanged() (fn js.Func[func(callback js.Func[func(devices js.Array[AudioDeviceInfo])])]) {
	bindings.FuncOnDeviceListChanged(
		js.Pointer(&fn),
	)
	return
}

// OnDeviceListChanged calls the function "WEBEXT.audio.onDeviceListChanged.addListener" directly.
func OnDeviceListChanged(callback js.Func[func(devices js.Array[AudioDeviceInfo])]) (ret js.Void) {
	bindings.CallOnDeviceListChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDeviceListChanged calls the function "WEBEXT.audio.onDeviceListChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDeviceListChanged(callback js.Func[func(devices js.Array[AudioDeviceInfo])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDeviceListChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDeviceListChanged returns true if the function "WEBEXT.audio.onDeviceListChanged.removeListener" exists.
func HasFuncOffDeviceListChanged() bool {
	return js.True == bindings.HasFuncOffDeviceListChanged()
}

// FuncOffDeviceListChanged returns the function "WEBEXT.audio.onDeviceListChanged.removeListener".
func FuncOffDeviceListChanged() (fn js.Func[func(callback js.Func[func(devices js.Array[AudioDeviceInfo])])]) {
	bindings.FuncOffDeviceListChanged(
		js.Pointer(&fn),
	)
	return
}

// OffDeviceListChanged calls the function "WEBEXT.audio.onDeviceListChanged.removeListener" directly.
func OffDeviceListChanged(callback js.Func[func(devices js.Array[AudioDeviceInfo])]) (ret js.Void) {
	bindings.CallOffDeviceListChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDeviceListChanged calls the function "WEBEXT.audio.onDeviceListChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDeviceListChanged(callback js.Func[func(devices js.Array[AudioDeviceInfo])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDeviceListChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDeviceListChanged returns true if the function "WEBEXT.audio.onDeviceListChanged.hasListener" exists.
func HasFuncHasOnDeviceListChanged() bool {
	return js.True == bindings.HasFuncHasOnDeviceListChanged()
}

// FuncHasOnDeviceListChanged returns the function "WEBEXT.audio.onDeviceListChanged.hasListener".
func FuncHasOnDeviceListChanged() (fn js.Func[func(callback js.Func[func(devices js.Array[AudioDeviceInfo])]) bool]) {
	bindings.FuncHasOnDeviceListChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnDeviceListChanged calls the function "WEBEXT.audio.onDeviceListChanged.hasListener" directly.
func HasOnDeviceListChanged(callback js.Func[func(devices js.Array[AudioDeviceInfo])]) (ret bool) {
	bindings.CallHasOnDeviceListChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDeviceListChanged calls the function "WEBEXT.audio.onDeviceListChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDeviceListChanged(callback js.Func[func(devices js.Array[AudioDeviceInfo])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDeviceListChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnLevelChangedEventCallbackFunc func(this js.Ref, event *LevelChangedEvent) js.Ref

func (fn OnLevelChangedEventCallbackFunc) Register() js.Func[func(event *LevelChangedEvent)] {
	return js.RegisterCallback[func(event *LevelChangedEvent)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnLevelChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 LevelChangedEvent
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnLevelChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event *LevelChangedEvent) js.Ref
	Arg T
}

func (cb *OnLevelChangedEventCallback[T]) Register() js.Func[func(event *LevelChangedEvent)] {
	return js.RegisterCallback[func(event *LevelChangedEvent)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnLevelChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 LevelChangedEvent
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnLevelChanged returns true if the function "WEBEXT.audio.onLevelChanged.addListener" exists.
func HasFuncOnLevelChanged() bool {
	return js.True == bindings.HasFuncOnLevelChanged()
}

// FuncOnLevelChanged returns the function "WEBEXT.audio.onLevelChanged.addListener".
func FuncOnLevelChanged() (fn js.Func[func(callback js.Func[func(event *LevelChangedEvent)])]) {
	bindings.FuncOnLevelChanged(
		js.Pointer(&fn),
	)
	return
}

// OnLevelChanged calls the function "WEBEXT.audio.onLevelChanged.addListener" directly.
func OnLevelChanged(callback js.Func[func(event *LevelChangedEvent)]) (ret js.Void) {
	bindings.CallOnLevelChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnLevelChanged calls the function "WEBEXT.audio.onLevelChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnLevelChanged(callback js.Func[func(event *LevelChangedEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnLevelChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffLevelChanged returns true if the function "WEBEXT.audio.onLevelChanged.removeListener" exists.
func HasFuncOffLevelChanged() bool {
	return js.True == bindings.HasFuncOffLevelChanged()
}

// FuncOffLevelChanged returns the function "WEBEXT.audio.onLevelChanged.removeListener".
func FuncOffLevelChanged() (fn js.Func[func(callback js.Func[func(event *LevelChangedEvent)])]) {
	bindings.FuncOffLevelChanged(
		js.Pointer(&fn),
	)
	return
}

// OffLevelChanged calls the function "WEBEXT.audio.onLevelChanged.removeListener" directly.
func OffLevelChanged(callback js.Func[func(event *LevelChangedEvent)]) (ret js.Void) {
	bindings.CallOffLevelChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffLevelChanged calls the function "WEBEXT.audio.onLevelChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffLevelChanged(callback js.Func[func(event *LevelChangedEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffLevelChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnLevelChanged returns true if the function "WEBEXT.audio.onLevelChanged.hasListener" exists.
func HasFuncHasOnLevelChanged() bool {
	return js.True == bindings.HasFuncHasOnLevelChanged()
}

// FuncHasOnLevelChanged returns the function "WEBEXT.audio.onLevelChanged.hasListener".
func FuncHasOnLevelChanged() (fn js.Func[func(callback js.Func[func(event *LevelChangedEvent)]) bool]) {
	bindings.FuncHasOnLevelChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnLevelChanged calls the function "WEBEXT.audio.onLevelChanged.hasListener" directly.
func HasOnLevelChanged(callback js.Func[func(event *LevelChangedEvent)]) (ret bool) {
	bindings.CallHasOnLevelChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnLevelChanged calls the function "WEBEXT.audio.onLevelChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnLevelChanged(callback js.Func[func(event *LevelChangedEvent)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnLevelChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnMuteChangedEventCallbackFunc func(this js.Ref, event *MuteChangedEvent) js.Ref

func (fn OnMuteChangedEventCallbackFunc) Register() js.Func[func(event *MuteChangedEvent)] {
	return js.RegisterCallback[func(event *MuteChangedEvent)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnMuteChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 MuteChangedEvent
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnMuteChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event *MuteChangedEvent) js.Ref
	Arg T
}

func (cb *OnMuteChangedEventCallback[T]) Register() js.Func[func(event *MuteChangedEvent)] {
	return js.RegisterCallback[func(event *MuteChangedEvent)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnMuteChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 MuteChangedEvent
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnMuteChanged returns true if the function "WEBEXT.audio.onMuteChanged.addListener" exists.
func HasFuncOnMuteChanged() bool {
	return js.True == bindings.HasFuncOnMuteChanged()
}

// FuncOnMuteChanged returns the function "WEBEXT.audio.onMuteChanged.addListener".
func FuncOnMuteChanged() (fn js.Func[func(callback js.Func[func(event *MuteChangedEvent)])]) {
	bindings.FuncOnMuteChanged(
		js.Pointer(&fn),
	)
	return
}

// OnMuteChanged calls the function "WEBEXT.audio.onMuteChanged.addListener" directly.
func OnMuteChanged(callback js.Func[func(event *MuteChangedEvent)]) (ret js.Void) {
	bindings.CallOnMuteChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMuteChanged calls the function "WEBEXT.audio.onMuteChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMuteChanged(callback js.Func[func(event *MuteChangedEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMuteChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMuteChanged returns true if the function "WEBEXT.audio.onMuteChanged.removeListener" exists.
func HasFuncOffMuteChanged() bool {
	return js.True == bindings.HasFuncOffMuteChanged()
}

// FuncOffMuteChanged returns the function "WEBEXT.audio.onMuteChanged.removeListener".
func FuncOffMuteChanged() (fn js.Func[func(callback js.Func[func(event *MuteChangedEvent)])]) {
	bindings.FuncOffMuteChanged(
		js.Pointer(&fn),
	)
	return
}

// OffMuteChanged calls the function "WEBEXT.audio.onMuteChanged.removeListener" directly.
func OffMuteChanged(callback js.Func[func(event *MuteChangedEvent)]) (ret js.Void) {
	bindings.CallOffMuteChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMuteChanged calls the function "WEBEXT.audio.onMuteChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMuteChanged(callback js.Func[func(event *MuteChangedEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMuteChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMuteChanged returns true if the function "WEBEXT.audio.onMuteChanged.hasListener" exists.
func HasFuncHasOnMuteChanged() bool {
	return js.True == bindings.HasFuncHasOnMuteChanged()
}

// FuncHasOnMuteChanged returns the function "WEBEXT.audio.onMuteChanged.hasListener".
func FuncHasOnMuteChanged() (fn js.Func[func(callback js.Func[func(event *MuteChangedEvent)]) bool]) {
	bindings.FuncHasOnMuteChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnMuteChanged calls the function "WEBEXT.audio.onMuteChanged.hasListener" directly.
func HasOnMuteChanged(callback js.Func[func(event *MuteChangedEvent)]) (ret bool) {
	bindings.CallHasOnMuteChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMuteChanged calls the function "WEBEXT.audio.onMuteChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnMuteChanged(callback js.Func[func(event *MuteChangedEvent)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnMuteChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSetActiveDevices returns true if the function "WEBEXT.audio.setActiveDevices" exists.
func HasFuncSetActiveDevices() bool {
	return js.True == bindings.HasFuncSetActiveDevices()
}

// FuncSetActiveDevices returns the function "WEBEXT.audio.setActiveDevices".
func FuncSetActiveDevices() (fn js.Func[func(ids DeviceIdLists) js.Promise[js.Void]]) {
	bindings.FuncSetActiveDevices(
		js.Pointer(&fn),
	)
	return
}

// SetActiveDevices calls the function "WEBEXT.audio.setActiveDevices" directly.
func SetActiveDevices(ids DeviceIdLists) (ret js.Promise[js.Void]) {
	bindings.CallSetActiveDevices(
		js.Pointer(&ret),
		js.Pointer(&ids),
	)

	return
}

// TrySetActiveDevices calls the function "WEBEXT.audio.setActiveDevices"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetActiveDevices(ids DeviceIdLists) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetActiveDevices(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&ids),
	)

	return
}

// HasFuncSetMute returns true if the function "WEBEXT.audio.setMute" exists.
func HasFuncSetMute() bool {
	return js.True == bindings.HasFuncSetMute()
}

// FuncSetMute returns the function "WEBEXT.audio.setMute".
func FuncSetMute() (fn js.Func[func(streamType StreamType, isMuted bool) js.Promise[js.Void]]) {
	bindings.FuncSetMute(
		js.Pointer(&fn),
	)
	return
}

// SetMute calls the function "WEBEXT.audio.setMute" directly.
func SetMute(streamType StreamType, isMuted bool) (ret js.Promise[js.Void]) {
	bindings.CallSetMute(
		js.Pointer(&ret),
		uint32(streamType),
		js.Bool(bool(isMuted)),
	)

	return
}

// TrySetMute calls the function "WEBEXT.audio.setMute"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetMute(streamType StreamType, isMuted bool) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetMute(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(streamType),
		js.Bool(bool(isMuted)),
	)

	return
}

// HasFuncSetProperties returns true if the function "WEBEXT.audio.setProperties" exists.
func HasFuncSetProperties() bool {
	return js.True == bindings.HasFuncSetProperties()
}

// FuncSetProperties returns the function "WEBEXT.audio.setProperties".
func FuncSetProperties() (fn js.Func[func(id js.String, properties DeviceProperties) js.Promise[js.Void]]) {
	bindings.FuncSetProperties(
		js.Pointer(&fn),
	)
	return
}

// SetProperties calls the function "WEBEXT.audio.setProperties" directly.
func SetProperties(id js.String, properties DeviceProperties) (ret js.Promise[js.Void]) {
	bindings.CallSetProperties(
		js.Pointer(&ret),
		id.Ref(),
		js.Pointer(&properties),
	)

	return
}

// TrySetProperties calls the function "WEBEXT.audio.setProperties"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetProperties(id js.String, properties DeviceProperties) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetProperties(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		js.Pointer(&properties),
	)

	return
}
