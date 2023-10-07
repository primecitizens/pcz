// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package filesystem

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/filesystem/bindings"
)

type AcceptOption struct {
	// Description is "AcceptOption.description"
	//
	// Optional
	Description js.String
	// MimeTypes is "AcceptOption.mimeTypes"
	//
	// Optional
	MimeTypes js.Array[js.String]
	// Extensions is "AcceptOption.extensions"
	//
	// Optional
	Extensions js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AcceptOption with all fields set.
func (p AcceptOption) FromRef(ref js.Ref) AcceptOption {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AcceptOption in the application heap.
func (p AcceptOption) New() js.Ref {
	return bindings.AcceptOptionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AcceptOption) UpdateFrom(ref js.Ref) {
	bindings.AcceptOptionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AcceptOption) Update(ref js.Ref) {
	bindings.AcceptOptionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AcceptOption) FreeMembers(recursive bool) {
	js.Free(
		p.Description.Ref(),
		p.MimeTypes.Ref(),
		p.Extensions.Ref(),
	)
	p.Description = p.Description.FromRef(js.Undefined)
	p.MimeTypes = p.MimeTypes.FromRef(js.Undefined)
	p.Extensions = p.Extensions.FromRef(js.Undefined)
}

type ChooseEntryType uint32

const (
	_ ChooseEntryType = iota

	ChooseEntryType_OPEN_FILE
	ChooseEntryType_OPEN_WRITABLE_FILE
	ChooseEntryType_SAVE_FILE
	ChooseEntryType_OPEN_DIRECTORY
)

func (ChooseEntryType) FromRef(str js.Ref) ChooseEntryType {
	return ChooseEntryType(bindings.ConstOfChooseEntryType(str))
}

func (x ChooseEntryType) String() (string, bool) {
	switch x {
	case ChooseEntryType_OPEN_FILE:
		return "openFile", true
	case ChooseEntryType_OPEN_WRITABLE_FILE:
		return "openWritableFile", true
	case ChooseEntryType_SAVE_FILE:
		return "saveFile", true
	case ChooseEntryType_OPEN_DIRECTORY:
		return "openDirectory", true
	default:
		return "", false
	}
}

type ChooseEntryOptions struct {
	// Type is "ChooseEntryOptions.type"
	//
	// Optional
	Type ChooseEntryType
	// SuggestedName is "ChooseEntryOptions.suggestedName"
	//
	// Optional
	SuggestedName js.String
	// Accepts is "ChooseEntryOptions.accepts"
	//
	// Optional
	Accepts js.Array[AcceptOption]
	// AcceptsAllTypes is "ChooseEntryOptions.acceptsAllTypes"
	//
	// Optional
	//
	// NOTE: FFI_USE_AcceptsAllTypes MUST be set to true to make this field effective.
	AcceptsAllTypes bool
	// AcceptsMultiple is "ChooseEntryOptions.acceptsMultiple"
	//
	// Optional
	//
	// NOTE: FFI_USE_AcceptsMultiple MUST be set to true to make this field effective.
	AcceptsMultiple bool

	FFI_USE_AcceptsAllTypes bool // for AcceptsAllTypes.
	FFI_USE_AcceptsMultiple bool // for AcceptsMultiple.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ChooseEntryOptions with all fields set.
func (p ChooseEntryOptions) FromRef(ref js.Ref) ChooseEntryOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ChooseEntryOptions in the application heap.
func (p ChooseEntryOptions) New() js.Ref {
	return bindings.ChooseEntryOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ChooseEntryOptions) UpdateFrom(ref js.Ref) {
	bindings.ChooseEntryOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ChooseEntryOptions) Update(ref js.Ref) {
	bindings.ChooseEntryOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ChooseEntryOptions) FreeMembers(recursive bool) {
	js.Free(
		p.SuggestedName.Ref(),
		p.Accepts.Ref(),
	)
	p.SuggestedName = p.SuggestedName.FromRef(js.Undefined)
	p.Accepts = p.Accepts.FromRef(js.Undefined)
}

type EntriesCallbackFunc func(this js.Ref, entry js.Object, fileEntries js.Array[js.Object]) js.Ref

func (fn EntriesCallbackFunc) Register() js.Func[func(entry js.Object, fileEntries js.Array[js.Object])] {
	return js.RegisterCallback[func(entry js.Object, fileEntries js.Array[js.Object])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn EntriesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Object{}.FromRef(args[0+1]),
		js.Array[js.Object]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type EntriesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entry js.Object, fileEntries js.Array[js.Object]) js.Ref
	Arg T
}

func (cb *EntriesCallback[T]) Register() js.Func[func(entry js.Object, fileEntries js.Array[js.Object])] {
	return js.RegisterCallback[func(entry js.Object, fileEntries js.Array[js.Object])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *EntriesCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Object{}.FromRef(args[0+1]),
		js.Array[js.Object]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type EntryCallbackFunc func(this js.Ref, entry js.Object) js.Ref

func (fn EntryCallbackFunc) Register() js.Func[func(entry js.Object)] {
	return js.RegisterCallback[func(entry js.Object)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn EntryCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Object{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type EntryCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entry js.Object) js.Ref
	Arg T
}

func (cb *EntryCallback[T]) Register() js.Func[func(entry js.Object)] {
	return js.RegisterCallback[func(entry js.Object)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *EntryCallback[T]) DispatchCallback(
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

		js.Object{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDisplayPathCallbackFunc func(this js.Ref, displayPath js.String) js.Ref

func (fn GetDisplayPathCallbackFunc) Register() js.Func[func(displayPath js.String)] {
	return js.RegisterCallback[func(displayPath js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDisplayPathCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDisplayPathCallback[T any] struct {
	Fn  func(arg T, this js.Ref, displayPath js.String) js.Ref
	Arg T
}

func (cb *GetDisplayPathCallback[T]) Register() js.Func[func(displayPath js.String)] {
	return js.RegisterCallback[func(displayPath js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDisplayPathCallback[T]) DispatchCallback(
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

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetVolumeListCallbackFunc func(this js.Ref, volumes js.Array[Volume]) js.Ref

func (fn GetVolumeListCallbackFunc) Register() js.Func[func(volumes js.Array[Volume])] {
	return js.RegisterCallback[func(volumes js.Array[Volume])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetVolumeListCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[Volume]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetVolumeListCallback[T any] struct {
	Fn  func(arg T, this js.Ref, volumes js.Array[Volume]) js.Ref
	Arg T
}

func (cb *GetVolumeListCallback[T]) Register() js.Func[func(volumes js.Array[Volume])] {
	return js.RegisterCallback[func(volumes js.Array[Volume])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetVolumeListCallback[T]) DispatchCallback(
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

		js.Array[Volume]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type Volume struct {
	// VolumeId is "Volume.volumeId"
	//
	// Optional
	VolumeId js.String
	// Writable is "Volume.writable"
	//
	// Optional
	//
	// NOTE: FFI_USE_Writable MUST be set to true to make this field effective.
	Writable bool

	FFI_USE_Writable bool // for Writable.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Volume with all fields set.
func (p Volume) FromRef(ref js.Ref) Volume {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Volume in the application heap.
func (p Volume) New() js.Ref {
	return bindings.VolumeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Volume) UpdateFrom(ref js.Ref) {
	bindings.VolumeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Volume) Update(ref js.Ref) {
	bindings.VolumeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Volume) FreeMembers(recursive bool) {
	js.Free(
		p.VolumeId.Ref(),
	)
	p.VolumeId = p.VolumeId.FromRef(js.Undefined)
}

type IsRestorableCallbackFunc func(this js.Ref, isRestorable bool) js.Ref

func (fn IsRestorableCallbackFunc) Register() js.Func[func(isRestorable bool)] {
	return js.RegisterCallback[func(isRestorable bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn IsRestorableCallbackFunc) DispatchCallback(
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

type IsRestorableCallback[T any] struct {
	Fn  func(arg T, this js.Ref, isRestorable bool) js.Ref
	Arg T
}

func (cb *IsRestorableCallback[T]) Register() js.Func[func(isRestorable bool)] {
	return js.RegisterCallback[func(isRestorable bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *IsRestorableCallback[T]) DispatchCallback(
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

type IsWritableCallbackFunc func(this js.Ref, isWritable bool) js.Ref

func (fn IsWritableCallbackFunc) Register() js.Func[func(isWritable bool)] {
	return js.RegisterCallback[func(isWritable bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn IsWritableCallbackFunc) DispatchCallback(
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

type IsWritableCallback[T any] struct {
	Fn  func(arg T, this js.Ref, isWritable bool) js.Ref
	Arg T
}

func (cb *IsWritableCallback[T]) Register() js.Func[func(isWritable bool)] {
	return js.RegisterCallback[func(isWritable bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *IsWritableCallback[T]) DispatchCallback(
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

type RequestFileSystemCallbackFunc func(this js.Ref, fileSystem js.Object) js.Ref

func (fn RequestFileSystemCallbackFunc) Register() js.Func[func(fileSystem js.Object)] {
	return js.RegisterCallback[func(fileSystem js.Object)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RequestFileSystemCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Object{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type RequestFileSystemCallback[T any] struct {
	Fn  func(arg T, this js.Ref, fileSystem js.Object) js.Ref
	Arg T
}

func (cb *RequestFileSystemCallback[T]) Register() js.Func[func(fileSystem js.Object)] {
	return js.RegisterCallback[func(fileSystem js.Object)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RequestFileSystemCallback[T]) DispatchCallback(
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

		js.Object{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type RequestFileSystemOptions struct {
	// VolumeId is "RequestFileSystemOptions.volumeId"
	//
	// Optional
	VolumeId js.String
	// Writable is "RequestFileSystemOptions.writable"
	//
	// Optional
	//
	// NOTE: FFI_USE_Writable MUST be set to true to make this field effective.
	Writable bool

	FFI_USE_Writable bool // for Writable.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RequestFileSystemOptions with all fields set.
func (p RequestFileSystemOptions) FromRef(ref js.Ref) RequestFileSystemOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RequestFileSystemOptions in the application heap.
func (p RequestFileSystemOptions) New() js.Ref {
	return bindings.RequestFileSystemOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RequestFileSystemOptions) UpdateFrom(ref js.Ref) {
	bindings.RequestFileSystemOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RequestFileSystemOptions) Update(ref js.Ref) {
	bindings.RequestFileSystemOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RequestFileSystemOptions) FreeMembers(recursive bool) {
	js.Free(
		p.VolumeId.Ref(),
	)
	p.VolumeId = p.VolumeId.FromRef(js.Undefined)
}

type VolumeListChangedEvent struct {
	// Volumes is "VolumeListChangedEvent.volumes"
	//
	// Optional
	Volumes js.Array[Volume]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VolumeListChangedEvent with all fields set.
func (p VolumeListChangedEvent) FromRef(ref js.Ref) VolumeListChangedEvent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VolumeListChangedEvent in the application heap.
func (p VolumeListChangedEvent) New() js.Ref {
	return bindings.VolumeListChangedEventJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *VolumeListChangedEvent) UpdateFrom(ref js.Ref) {
	bindings.VolumeListChangedEventJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *VolumeListChangedEvent) Update(ref js.Ref) {
	bindings.VolumeListChangedEventJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *VolumeListChangedEvent) FreeMembers(recursive bool) {
	js.Free(
		p.Volumes.Ref(),
	)
	p.Volumes = p.Volumes.FromRef(js.Undefined)
}

// HasFuncChooseEntry returns true if the function "WEBEXT.fileSystem.chooseEntry" exists.
func HasFuncChooseEntry() bool {
	return js.True == bindings.HasFuncChooseEntry()
}

// FuncChooseEntry returns the function "WEBEXT.fileSystem.chooseEntry".
func FuncChooseEntry() (fn js.Func[func(options ChooseEntryOptions, callback js.Func[func(entry js.Object, fileEntries js.Array[js.Object])])]) {
	bindings.FuncChooseEntry(
		js.Pointer(&fn),
	)
	return
}

// ChooseEntry calls the function "WEBEXT.fileSystem.chooseEntry" directly.
func ChooseEntry(options ChooseEntryOptions, callback js.Func[func(entry js.Object, fileEntries js.Array[js.Object])]) (ret js.Void) {
	bindings.CallChooseEntry(
		js.Pointer(&ret),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// TryChooseEntry calls the function "WEBEXT.fileSystem.chooseEntry"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryChooseEntry(options ChooseEntryOptions, callback js.Func[func(entry js.Object, fileEntries js.Array[js.Object])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryChooseEntry(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// HasFuncGetDisplayPath returns true if the function "WEBEXT.fileSystem.getDisplayPath" exists.
func HasFuncGetDisplayPath() bool {
	return js.True == bindings.HasFuncGetDisplayPath()
}

// FuncGetDisplayPath returns the function "WEBEXT.fileSystem.getDisplayPath".
func FuncGetDisplayPath() (fn js.Func[func(entry js.Object) js.Promise[js.String]]) {
	bindings.FuncGetDisplayPath(
		js.Pointer(&fn),
	)
	return
}

// GetDisplayPath calls the function "WEBEXT.fileSystem.getDisplayPath" directly.
func GetDisplayPath(entry js.Object) (ret js.Promise[js.String]) {
	bindings.CallGetDisplayPath(
		js.Pointer(&ret),
		entry.Ref(),
	)

	return
}

// TryGetDisplayPath calls the function "WEBEXT.fileSystem.getDisplayPath"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDisplayPath(entry js.Object) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDisplayPath(
		js.Pointer(&ret), js.Pointer(&exception),
		entry.Ref(),
	)

	return
}

// HasFuncGetVolumeList returns true if the function "WEBEXT.fileSystem.getVolumeList" exists.
func HasFuncGetVolumeList() bool {
	return js.True == bindings.HasFuncGetVolumeList()
}

// FuncGetVolumeList returns the function "WEBEXT.fileSystem.getVolumeList".
func FuncGetVolumeList() (fn js.Func[func() js.Promise[js.Array[Volume]]]) {
	bindings.FuncGetVolumeList(
		js.Pointer(&fn),
	)
	return
}

// GetVolumeList calls the function "WEBEXT.fileSystem.getVolumeList" directly.
func GetVolumeList() (ret js.Promise[js.Array[Volume]]) {
	bindings.CallGetVolumeList(
		js.Pointer(&ret),
	)

	return
}

// TryGetVolumeList calls the function "WEBEXT.fileSystem.getVolumeList"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetVolumeList() (ret js.Promise[js.Array[Volume]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetVolumeList(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetWritableEntry returns true if the function "WEBEXT.fileSystem.getWritableEntry" exists.
func HasFuncGetWritableEntry() bool {
	return js.True == bindings.HasFuncGetWritableEntry()
}

// FuncGetWritableEntry returns the function "WEBEXT.fileSystem.getWritableEntry".
func FuncGetWritableEntry() (fn js.Func[func(entry js.Object) js.Promise[js.Object]]) {
	bindings.FuncGetWritableEntry(
		js.Pointer(&fn),
	)
	return
}

// GetWritableEntry calls the function "WEBEXT.fileSystem.getWritableEntry" directly.
func GetWritableEntry(entry js.Object) (ret js.Promise[js.Object]) {
	bindings.CallGetWritableEntry(
		js.Pointer(&ret),
		entry.Ref(),
	)

	return
}

// TryGetWritableEntry calls the function "WEBEXT.fileSystem.getWritableEntry"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetWritableEntry(entry js.Object) (ret js.Promise[js.Object], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetWritableEntry(
		js.Pointer(&ret), js.Pointer(&exception),
		entry.Ref(),
	)

	return
}

// HasFuncIsRestorable returns true if the function "WEBEXT.fileSystem.isRestorable" exists.
func HasFuncIsRestorable() bool {
	return js.True == bindings.HasFuncIsRestorable()
}

// FuncIsRestorable returns the function "WEBEXT.fileSystem.isRestorable".
func FuncIsRestorable() (fn js.Func[func(id js.String) js.Promise[js.Boolean]]) {
	bindings.FuncIsRestorable(
		js.Pointer(&fn),
	)
	return
}

// IsRestorable calls the function "WEBEXT.fileSystem.isRestorable" directly.
func IsRestorable(id js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallIsRestorable(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryIsRestorable calls the function "WEBEXT.fileSystem.isRestorable"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsRestorable(id js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsRestorable(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncIsWritableEntry returns true if the function "WEBEXT.fileSystem.isWritableEntry" exists.
func HasFuncIsWritableEntry() bool {
	return js.True == bindings.HasFuncIsWritableEntry()
}

// FuncIsWritableEntry returns the function "WEBEXT.fileSystem.isWritableEntry".
func FuncIsWritableEntry() (fn js.Func[func(entry js.Object) js.Promise[js.Boolean]]) {
	bindings.FuncIsWritableEntry(
		js.Pointer(&fn),
	)
	return
}

// IsWritableEntry calls the function "WEBEXT.fileSystem.isWritableEntry" directly.
func IsWritableEntry(entry js.Object) (ret js.Promise[js.Boolean]) {
	bindings.CallIsWritableEntry(
		js.Pointer(&ret),
		entry.Ref(),
	)

	return
}

// TryIsWritableEntry calls the function "WEBEXT.fileSystem.isWritableEntry"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsWritableEntry(entry js.Object) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsWritableEntry(
		js.Pointer(&ret), js.Pointer(&exception),
		entry.Ref(),
	)

	return
}

type OnVolumeListChangedEventCallbackFunc func(this js.Ref, event *VolumeListChangedEvent) js.Ref

func (fn OnVolumeListChangedEventCallbackFunc) Register() js.Func[func(event *VolumeListChangedEvent)] {
	return js.RegisterCallback[func(event *VolumeListChangedEvent)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnVolumeListChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 VolumeListChangedEvent
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

type OnVolumeListChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event *VolumeListChangedEvent) js.Ref
	Arg T
}

func (cb *OnVolumeListChangedEventCallback[T]) Register() js.Func[func(event *VolumeListChangedEvent)] {
	return js.RegisterCallback[func(event *VolumeListChangedEvent)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnVolumeListChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 VolumeListChangedEvent
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

// HasFuncOnVolumeListChanged returns true if the function "WEBEXT.fileSystem.onVolumeListChanged.addListener" exists.
func HasFuncOnVolumeListChanged() bool {
	return js.True == bindings.HasFuncOnVolumeListChanged()
}

// FuncOnVolumeListChanged returns the function "WEBEXT.fileSystem.onVolumeListChanged.addListener".
func FuncOnVolumeListChanged() (fn js.Func[func(callback js.Func[func(event *VolumeListChangedEvent)])]) {
	bindings.FuncOnVolumeListChanged(
		js.Pointer(&fn),
	)
	return
}

// OnVolumeListChanged calls the function "WEBEXT.fileSystem.onVolumeListChanged.addListener" directly.
func OnVolumeListChanged(callback js.Func[func(event *VolumeListChangedEvent)]) (ret js.Void) {
	bindings.CallOnVolumeListChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnVolumeListChanged calls the function "WEBEXT.fileSystem.onVolumeListChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnVolumeListChanged(callback js.Func[func(event *VolumeListChangedEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnVolumeListChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffVolumeListChanged returns true if the function "WEBEXT.fileSystem.onVolumeListChanged.removeListener" exists.
func HasFuncOffVolumeListChanged() bool {
	return js.True == bindings.HasFuncOffVolumeListChanged()
}

// FuncOffVolumeListChanged returns the function "WEBEXT.fileSystem.onVolumeListChanged.removeListener".
func FuncOffVolumeListChanged() (fn js.Func[func(callback js.Func[func(event *VolumeListChangedEvent)])]) {
	bindings.FuncOffVolumeListChanged(
		js.Pointer(&fn),
	)
	return
}

// OffVolumeListChanged calls the function "WEBEXT.fileSystem.onVolumeListChanged.removeListener" directly.
func OffVolumeListChanged(callback js.Func[func(event *VolumeListChangedEvent)]) (ret js.Void) {
	bindings.CallOffVolumeListChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffVolumeListChanged calls the function "WEBEXT.fileSystem.onVolumeListChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffVolumeListChanged(callback js.Func[func(event *VolumeListChangedEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffVolumeListChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnVolumeListChanged returns true if the function "WEBEXT.fileSystem.onVolumeListChanged.hasListener" exists.
func HasFuncHasOnVolumeListChanged() bool {
	return js.True == bindings.HasFuncHasOnVolumeListChanged()
}

// FuncHasOnVolumeListChanged returns the function "WEBEXT.fileSystem.onVolumeListChanged.hasListener".
func FuncHasOnVolumeListChanged() (fn js.Func[func(callback js.Func[func(event *VolumeListChangedEvent)]) bool]) {
	bindings.FuncHasOnVolumeListChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnVolumeListChanged calls the function "WEBEXT.fileSystem.onVolumeListChanged.hasListener" directly.
func HasOnVolumeListChanged(callback js.Func[func(event *VolumeListChangedEvent)]) (ret bool) {
	bindings.CallHasOnVolumeListChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnVolumeListChanged calls the function "WEBEXT.fileSystem.onVolumeListChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnVolumeListChanged(callback js.Func[func(event *VolumeListChangedEvent)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnVolumeListChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncRequestFileSystem returns true if the function "WEBEXT.fileSystem.requestFileSystem" exists.
func HasFuncRequestFileSystem() bool {
	return js.True == bindings.HasFuncRequestFileSystem()
}

// FuncRequestFileSystem returns the function "WEBEXT.fileSystem.requestFileSystem".
func FuncRequestFileSystem() (fn js.Func[func(options RequestFileSystemOptions) js.Promise[js.Object]]) {
	bindings.FuncRequestFileSystem(
		js.Pointer(&fn),
	)
	return
}

// RequestFileSystem calls the function "WEBEXT.fileSystem.requestFileSystem" directly.
func RequestFileSystem(options RequestFileSystemOptions) (ret js.Promise[js.Object]) {
	bindings.CallRequestFileSystem(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRequestFileSystem calls the function "WEBEXT.fileSystem.requestFileSystem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRequestFileSystem(options RequestFileSystemOptions) (ret js.Promise[js.Object], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequestFileSystem(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncRestoreEntry returns true if the function "WEBEXT.fileSystem.restoreEntry" exists.
func HasFuncRestoreEntry() bool {
	return js.True == bindings.HasFuncRestoreEntry()
}

// FuncRestoreEntry returns the function "WEBEXT.fileSystem.restoreEntry".
func FuncRestoreEntry() (fn js.Func[func(id js.String, callback js.Func[func(entry js.Object)])]) {
	bindings.FuncRestoreEntry(
		js.Pointer(&fn),
	)
	return
}

// RestoreEntry calls the function "WEBEXT.fileSystem.restoreEntry" directly.
func RestoreEntry(id js.String, callback js.Func[func(entry js.Object)]) (ret js.Void) {
	bindings.CallRestoreEntry(
		js.Pointer(&ret),
		id.Ref(),
		callback.Ref(),
	)

	return
}

// TryRestoreEntry calls the function "WEBEXT.fileSystem.restoreEntry"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRestoreEntry(id js.String, callback js.Func[func(entry js.Object)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRestoreEntry(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncRetainEntry returns true if the function "WEBEXT.fileSystem.retainEntry" exists.
func HasFuncRetainEntry() bool {
	return js.True == bindings.HasFuncRetainEntry()
}

// FuncRetainEntry returns the function "WEBEXT.fileSystem.retainEntry".
func FuncRetainEntry() (fn js.Func[func(entry js.Object) js.String]) {
	bindings.FuncRetainEntry(
		js.Pointer(&fn),
	)
	return
}

// RetainEntry calls the function "WEBEXT.fileSystem.retainEntry" directly.
func RetainEntry(entry js.Object) (ret js.String) {
	bindings.CallRetainEntry(
		js.Pointer(&ret),
		entry.Ref(),
	)

	return
}

// TryRetainEntry calls the function "WEBEXT.fileSystem.retainEntry"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRetainEntry(entry js.Object) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRetainEntry(
		js.Pointer(&ret), js.Pointer(&exception),
		entry.Ref(),
	)

	return
}
