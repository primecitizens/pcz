// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package filemanagerprivateinternal

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/filemanagerprivate"
	"github.com/primecitizens/pcz/std/plat/js/webext/filemanagerprivateinternal/bindings"
	"github.com/primecitizens/pcz/std/plat/js/webext/filesystemprovider"
)

type AddFileWatchCallbackFunc func(this js.Ref, success bool) js.Ref

func (fn AddFileWatchCallbackFunc) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn AddFileWatchCallbackFunc) DispatchCallback(
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

type AddFileWatchCallback[T any] struct {
	Fn  func(arg T, this js.Ref, success bool) js.Ref
	Arg T
}

func (cb *AddFileWatchCallback[T]) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *AddFileWatchCallback[T]) DispatchCallback(
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

type BooleanCallbackFunc func(this js.Ref, result bool) js.Ref

func (fn BooleanCallbackFunc) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn BooleanCallbackFunc) DispatchCallback(
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

type BooleanCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result bool) js.Ref
	Arg T
}

func (cb *BooleanCallback[T]) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *BooleanCallback[T]) DispatchCallback(
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

type ComputeChecksumCallbackFunc func(this js.Ref, checksum js.String) js.Ref

func (fn ComputeChecksumCallbackFunc) Register() js.Func[func(checksum js.String)] {
	return js.RegisterCallback[func(checksum js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ComputeChecksumCallbackFunc) DispatchCallback(
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

type ComputeChecksumCallback[T any] struct {
	Fn  func(arg T, this js.Ref, checksum js.String) js.Ref
	Arg T
}

func (cb *ComputeChecksumCallback[T]) Register() js.Func[func(checksum js.String)] {
	return js.RegisterCallback[func(checksum js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ComputeChecksumCallback[T]) DispatchCallback(
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

type EntryDescription struct {
	// FileSystemName is "EntryDescription.fileSystemName"
	//
	// Optional
	FileSystemName js.String
	// FileSystemRoot is "EntryDescription.fileSystemRoot"
	//
	// Optional
	FileSystemRoot js.String
	// FileFullPath is "EntryDescription.fileFullPath"
	//
	// Optional
	FileFullPath js.String
	// FileIsDirectory is "EntryDescription.fileIsDirectory"
	//
	// Optional
	//
	// NOTE: FFI_USE_FileIsDirectory MUST be set to true to make this field effective.
	FileIsDirectory bool

	FFI_USE_FileIsDirectory bool // for FileIsDirectory.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EntryDescription with all fields set.
func (p EntryDescription) FromRef(ref js.Ref) EntryDescription {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EntryDescription in the application heap.
func (p EntryDescription) New() js.Ref {
	return bindings.EntryDescriptionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *EntryDescription) UpdateFrom(ref js.Ref) {
	bindings.EntryDescriptionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EntryDescription) Update(ref js.Ref) {
	bindings.EntryDescriptionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EntryDescription) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemName.Ref(),
		p.FileSystemRoot.Ref(),
		p.FileFullPath.Ref(),
	)
	p.FileSystemName = p.FileSystemName.FromRef(js.Undefined)
	p.FileSystemRoot = p.FileSystemRoot.FromRef(js.Undefined)
	p.FileFullPath = p.FileFullPath.FromRef(js.Undefined)
}

type ExecuteTaskCallbackFunc func(this js.Ref, result filemanagerprivate.TaskResult) js.Ref

func (fn ExecuteTaskCallbackFunc) Register() js.Func[func(result filemanagerprivate.TaskResult)] {
	return js.RegisterCallback[func(result filemanagerprivate.TaskResult)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ExecuteTaskCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		filemanagerprivate.TaskResult(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ExecuteTaskCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result filemanagerprivate.TaskResult) js.Ref
	Arg T
}

func (cb *ExecuteTaskCallback[T]) Register() js.Func[func(result filemanagerprivate.TaskResult)] {
	return js.RegisterCallback[func(result filemanagerprivate.TaskResult)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ExecuteTaskCallback[T]) DispatchCallback(
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

		filemanagerprivate.TaskResult(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetContentMetadataCallbackFunc func(this js.Ref, result *filemanagerprivate.MediaMetadata) js.Ref

func (fn GetContentMetadataCallbackFunc) Register() js.Func[func(result *filemanagerprivate.MediaMetadata)] {
	return js.RegisterCallback[func(result *filemanagerprivate.MediaMetadata)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetContentMetadataCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 filemanagerprivate.MediaMetadata
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

type GetContentMetadataCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *filemanagerprivate.MediaMetadata) js.Ref
	Arg T
}

func (cb *GetContentMetadataCallback[T]) Register() js.Func[func(result *filemanagerprivate.MediaMetadata)] {
	return js.RegisterCallback[func(result *filemanagerprivate.MediaMetadata)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetContentMetadataCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 filemanagerprivate.MediaMetadata
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

type GetContentMimeTypeCallbackFunc func(this js.Ref, result js.String) js.Ref

func (fn GetContentMimeTypeCallbackFunc) Register() js.Func[func(result js.String)] {
	return js.RegisterCallback[func(result js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetContentMimeTypeCallbackFunc) DispatchCallback(
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

type GetContentMimeTypeCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result js.String) js.Ref
	Arg T
}

func (cb *GetContentMimeTypeCallback[T]) Register() js.Func[func(result js.String)] {
	return js.RegisterCallback[func(result js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetContentMimeTypeCallback[T]) DispatchCallback(
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

type GetCrostiniSharedPathsCallbackFunc func(this js.Ref, entries js.Array[EntryDescription], firstForSession bool) js.Ref

func (fn GetCrostiniSharedPathsCallbackFunc) Register() js.Func[func(entries js.Array[EntryDescription], firstForSession bool)] {
	return js.RegisterCallback[func(entries js.Array[EntryDescription], firstForSession bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetCrostiniSharedPathsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[EntryDescription]{}.FromRef(args[0+1]),
		args[1+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetCrostiniSharedPathsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[EntryDescription], firstForSession bool) js.Ref
	Arg T
}

func (cb *GetCrostiniSharedPathsCallback[T]) Register() js.Func[func(entries js.Array[EntryDescription], firstForSession bool)] {
	return js.RegisterCallback[func(entries js.Array[EntryDescription], firstForSession bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetCrostiniSharedPathsCallback[T]) DispatchCallback(
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

		js.Array[EntryDescription]{}.FromRef(args[0+1]),
		args[1+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetCustomActionsCallbackFunc func(this js.Ref, actions js.Array[filesystemprovider.Action]) js.Ref

func (fn GetCustomActionsCallbackFunc) Register() js.Func[func(actions js.Array[filesystemprovider.Action])] {
	return js.RegisterCallback[func(actions js.Array[filesystemprovider.Action])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetCustomActionsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[filesystemprovider.Action]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetCustomActionsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, actions js.Array[filesystemprovider.Action]) js.Ref
	Arg T
}

func (cb *GetCustomActionsCallback[T]) Register() js.Func[func(actions js.Array[filesystemprovider.Action])] {
	return js.RegisterCallback[func(actions js.Array[filesystemprovider.Action])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetCustomActionsCallback[T]) DispatchCallback(
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

		js.Array[filesystemprovider.Action]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDirectorySizeCallbackFunc func(this js.Ref, size float64) js.Ref

func (fn GetDirectorySizeCallbackFunc) Register() js.Func[func(size float64)] {
	return js.RegisterCallback[func(size float64)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDirectorySizeCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[float64]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDirectorySizeCallback[T any] struct {
	Fn  func(arg T, this js.Ref, size float64) js.Ref
	Arg T
}

func (cb *GetDirectorySizeCallback[T]) Register() js.Func[func(size float64)] {
	return js.RegisterCallback[func(size float64)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDirectorySizeCallback[T]) DispatchCallback(
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

		js.Number[float64]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDisallowedTransfersCallbackFunc func(this js.Ref, entries js.Array[EntryDescription]) js.Ref

func (fn GetDisallowedTransfersCallbackFunc) Register() js.Func[func(entries js.Array[EntryDescription])] {
	return js.RegisterCallback[func(entries js.Array[EntryDescription])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDisallowedTransfersCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[EntryDescription]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDisallowedTransfersCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[EntryDescription]) js.Ref
	Arg T
}

func (cb *GetDisallowedTransfersCallback[T]) Register() js.Func[func(entries js.Array[EntryDescription])] {
	return js.RegisterCallback[func(entries js.Array[EntryDescription])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDisallowedTransfersCallback[T]) DispatchCallback(
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

		js.Array[EntryDescription]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDlpMetadataCallbackFunc func(this js.Ref, entries js.Array[filemanagerprivate.DlpMetadata]) js.Ref

func (fn GetDlpMetadataCallbackFunc) Register() js.Func[func(entries js.Array[filemanagerprivate.DlpMetadata])] {
	return js.RegisterCallback[func(entries js.Array[filemanagerprivate.DlpMetadata])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDlpMetadataCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[filemanagerprivate.DlpMetadata]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDlpMetadataCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[filemanagerprivate.DlpMetadata]) js.Ref
	Arg T
}

func (cb *GetDlpMetadataCallback[T]) Register() js.Func[func(entries js.Array[filemanagerprivate.DlpMetadata])] {
	return js.RegisterCallback[func(entries js.Array[filemanagerprivate.DlpMetadata])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDlpMetadataCallback[T]) DispatchCallback(
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

		js.Array[filemanagerprivate.DlpMetadata]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDriveQuotaMetadataCallbackFunc func(this js.Ref, driveQuotaMetadata *filemanagerprivate.DriveQuotaMetadata) js.Ref

func (fn GetDriveQuotaMetadataCallbackFunc) Register() js.Func[func(driveQuotaMetadata *filemanagerprivate.DriveQuotaMetadata)] {
	return js.RegisterCallback[func(driveQuotaMetadata *filemanagerprivate.DriveQuotaMetadata)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDriveQuotaMetadataCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 filemanagerprivate.DriveQuotaMetadata
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

type GetDriveQuotaMetadataCallback[T any] struct {
	Fn  func(arg T, this js.Ref, driveQuotaMetadata *filemanagerprivate.DriveQuotaMetadata) js.Ref
	Arg T
}

func (cb *GetDriveQuotaMetadataCallback[T]) Register() js.Func[func(driveQuotaMetadata *filemanagerprivate.DriveQuotaMetadata)] {
	return js.RegisterCallback[func(driveQuotaMetadata *filemanagerprivate.DriveQuotaMetadata)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDriveQuotaMetadataCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 filemanagerprivate.DriveQuotaMetadata
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

type GetEntryPropertiesCallbackFunc func(this js.Ref, entryProperties js.Array[filemanagerprivate.EntryProperties]) js.Ref

func (fn GetEntryPropertiesCallbackFunc) Register() js.Func[func(entryProperties js.Array[filemanagerprivate.EntryProperties])] {
	return js.RegisterCallback[func(entryProperties js.Array[filemanagerprivate.EntryProperties])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetEntryPropertiesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[filemanagerprivate.EntryProperties]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetEntryPropertiesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entryProperties js.Array[filemanagerprivate.EntryProperties]) js.Ref
	Arg T
}

func (cb *GetEntryPropertiesCallback[T]) Register() js.Func[func(entryProperties js.Array[filemanagerprivate.EntryProperties])] {
	return js.RegisterCallback[func(entryProperties js.Array[filemanagerprivate.EntryProperties])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetEntryPropertiesCallback[T]) DispatchCallback(
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

		js.Array[filemanagerprivate.EntryProperties]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetFileTasksCallbackFunc func(this js.Ref, resultingTasks *filemanagerprivate.ResultingTasks) js.Ref

func (fn GetFileTasksCallbackFunc) Register() js.Func[func(resultingTasks *filemanagerprivate.ResultingTasks)] {
	return js.RegisterCallback[func(resultingTasks *filemanagerprivate.ResultingTasks)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetFileTasksCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 filemanagerprivate.ResultingTasks
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

type GetFileTasksCallback[T any] struct {
	Fn  func(arg T, this js.Ref, resultingTasks *filemanagerprivate.ResultingTasks) js.Ref
	Arg T
}

func (cb *GetFileTasksCallback[T]) Register() js.Func[func(resultingTasks *filemanagerprivate.ResultingTasks)] {
	return js.RegisterCallback[func(resultingTasks *filemanagerprivate.ResultingTasks)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetFileTasksCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 filemanagerprivate.ResultingTasks
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

type GetLinuxPackageInfoCallbackFunc func(this js.Ref, linux_package_info *filemanagerprivate.LinuxPackageInfo) js.Ref

func (fn GetLinuxPackageInfoCallbackFunc) Register() js.Func[func(linux_package_info *filemanagerprivate.LinuxPackageInfo)] {
	return js.RegisterCallback[func(linux_package_info *filemanagerprivate.LinuxPackageInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetLinuxPackageInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 filemanagerprivate.LinuxPackageInfo
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

type GetLinuxPackageInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, linux_package_info *filemanagerprivate.LinuxPackageInfo) js.Ref
	Arg T
}

func (cb *GetLinuxPackageInfoCallback[T]) Register() js.Func[func(linux_package_info *filemanagerprivate.LinuxPackageInfo)] {
	return js.RegisterCallback[func(linux_package_info *filemanagerprivate.LinuxPackageInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetLinuxPackageInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 filemanagerprivate.LinuxPackageInfo
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

type GetMimeTypeCallbackFunc func(this js.Ref, result js.String) js.Ref

func (fn GetMimeTypeCallbackFunc) Register() js.Func[func(result js.String)] {
	return js.RegisterCallback[func(result js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetMimeTypeCallbackFunc) DispatchCallback(
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

type GetMimeTypeCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result js.String) js.Ref
	Arg T
}

func (cb *GetMimeTypeCallback[T]) Register() js.Func[func(result js.String)] {
	return js.RegisterCallback[func(result js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetMimeTypeCallback[T]) DispatchCallback(
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

type GetRecentFilesCallbackFunc func(this js.Ref, entries js.Array[EntryDescription]) js.Ref

func (fn GetRecentFilesCallbackFunc) Register() js.Func[func(entries js.Array[EntryDescription])] {
	return js.RegisterCallback[func(entries js.Array[EntryDescription])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetRecentFilesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[EntryDescription]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetRecentFilesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[EntryDescription]) js.Ref
	Arg T
}

func (cb *GetRecentFilesCallback[T]) Register() js.Func[func(entries js.Array[EntryDescription])] {
	return js.RegisterCallback[func(entries js.Array[EntryDescription])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetRecentFilesCallback[T]) DispatchCallback(
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

		js.Array[EntryDescription]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetThumbnailCallbackFunc func(this js.Ref, ThumbnailDataUrl js.String) js.Ref

func (fn GetThumbnailCallbackFunc) Register() js.Func[func(ThumbnailDataUrl js.String)] {
	return js.RegisterCallback[func(ThumbnailDataUrl js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetThumbnailCallbackFunc) DispatchCallback(
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

type GetThumbnailCallback[T any] struct {
	Fn  func(arg T, this js.Ref, ThumbnailDataUrl js.String) js.Ref
	Arg T
}

func (cb *GetThumbnailCallback[T]) Register() js.Func[func(ThumbnailDataUrl js.String)] {
	return js.RegisterCallback[func(ThumbnailDataUrl js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetThumbnailCallback[T]) DispatchCallback(
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

type GetVolumeRootCallbackFunc func(this js.Ref, rootDir *EntryDescription) js.Ref

func (fn GetVolumeRootCallbackFunc) Register() js.Func[func(rootDir *EntryDescription)] {
	return js.RegisterCallback[func(rootDir *EntryDescription)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetVolumeRootCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 EntryDescription
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

type GetVolumeRootCallback[T any] struct {
	Fn  func(arg T, this js.Ref, rootDir *EntryDescription) js.Ref
	Arg T
}

func (cb *GetVolumeRootCallback[T]) Register() js.Func[func(rootDir *EntryDescription)] {
	return js.RegisterCallback[func(rootDir *EntryDescription)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetVolumeRootCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 EntryDescription
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

type IOTaskIdCallbackFunc func(this js.Ref, taskId int32) js.Ref

func (fn IOTaskIdCallbackFunc) Register() js.Func[func(taskId int32)] {
	return js.RegisterCallback[func(taskId int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn IOTaskIdCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type IOTaskIdCallback[T any] struct {
	Fn  func(arg T, this js.Ref, taskId int32) js.Ref
	Arg T
}

func (cb *IOTaskIdCallback[T]) Register() js.Func[func(taskId int32)] {
	return js.RegisterCallback[func(taskId int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *IOTaskIdCallback[T]) DispatchCallback(
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type IOTaskParams struct {
	// DestinationFolderUrl is "IOTaskParams.destinationFolderUrl"
	//
	// Optional
	DestinationFolderUrl js.String
	// Password is "IOTaskParams.password"
	//
	// Optional
	Password js.String
	// ShowNotification is "IOTaskParams.showNotification"
	//
	// Optional
	//
	// NOTE: FFI_USE_ShowNotification MUST be set to true to make this field effective.
	ShowNotification bool

	FFI_USE_ShowNotification bool // for ShowNotification.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IOTaskParams with all fields set.
func (p IOTaskParams) FromRef(ref js.Ref) IOTaskParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new IOTaskParams in the application heap.
func (p IOTaskParams) New() js.Ref {
	return bindings.IOTaskParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *IOTaskParams) UpdateFrom(ref js.Ref) {
	bindings.IOTaskParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IOTaskParams) Update(ref js.Ref) {
	bindings.IOTaskParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IOTaskParams) FreeMembers(recursive bool) {
	js.Free(
		p.DestinationFolderUrl.Ref(),
		p.Password.Ref(),
	)
	p.DestinationFolderUrl = p.DestinationFolderUrl.FromRef(js.Undefined)
	p.Password = p.Password.FromRef(js.Undefined)
}

type InstallLinuxPackageCallbackFunc func(this js.Ref, response filemanagerprivate.InstallLinuxPackageResponse) js.Ref

func (fn InstallLinuxPackageCallbackFunc) Register() js.Func[func(response filemanagerprivate.InstallLinuxPackageResponse)] {
	return js.RegisterCallback[func(response filemanagerprivate.InstallLinuxPackageResponse)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn InstallLinuxPackageCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		filemanagerprivate.InstallLinuxPackageResponse(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type InstallLinuxPackageCallback[T any] struct {
	Fn  func(arg T, this js.Ref, response filemanagerprivate.InstallLinuxPackageResponse) js.Ref
	Arg T
}

func (cb *InstallLinuxPackageCallback[T]) Register() js.Func[func(response filemanagerprivate.InstallLinuxPackageResponse)] {
	return js.RegisterCallback[func(response filemanagerprivate.InstallLinuxPackageResponse)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *InstallLinuxPackageCallback[T]) DispatchCallback(
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

		filemanagerprivate.InstallLinuxPackageResponse(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ParseTrashInfoFilesCallbackFunc func(this js.Ref, files js.Array[ParsedTrashInfoFile]) js.Ref

func (fn ParseTrashInfoFilesCallbackFunc) Register() js.Func[func(files js.Array[ParsedTrashInfoFile])] {
	return js.RegisterCallback[func(files js.Array[ParsedTrashInfoFile])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ParseTrashInfoFilesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[ParsedTrashInfoFile]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ParseTrashInfoFilesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, files js.Array[ParsedTrashInfoFile]) js.Ref
	Arg T
}

func (cb *ParseTrashInfoFilesCallback[T]) Register() js.Func[func(files js.Array[ParsedTrashInfoFile])] {
	return js.RegisterCallback[func(files js.Array[ParsedTrashInfoFile])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ParseTrashInfoFilesCallback[T]) DispatchCallback(
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

		js.Array[ParsedTrashInfoFile]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ParsedTrashInfoFile struct {
	// RestoreEntry is "ParsedTrashInfoFile.restoreEntry"
	//
	// Optional
	//
	// NOTE: RestoreEntry.FFI_USE MUST be set to true to get RestoreEntry used.
	RestoreEntry EntryDescription
	// TrashInfoFileName is "ParsedTrashInfoFile.trashInfoFileName"
	//
	// Optional
	TrashInfoFileName js.String
	// DeletionDate is "ParsedTrashInfoFile.deletionDate"
	//
	// Optional
	//
	// NOTE: FFI_USE_DeletionDate MUST be set to true to make this field effective.
	DeletionDate float64

	FFI_USE_DeletionDate bool // for DeletionDate.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ParsedTrashInfoFile with all fields set.
func (p ParsedTrashInfoFile) FromRef(ref js.Ref) ParsedTrashInfoFile {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ParsedTrashInfoFile in the application heap.
func (p ParsedTrashInfoFile) New() js.Ref {
	return bindings.ParsedTrashInfoFileJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ParsedTrashInfoFile) UpdateFrom(ref js.Ref) {
	bindings.ParsedTrashInfoFileJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ParsedTrashInfoFile) Update(ref js.Ref) {
	bindings.ParsedTrashInfoFileJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ParsedTrashInfoFile) FreeMembers(recursive bool) {
	js.Free(
		p.TrashInfoFileName.Ref(),
	)
	p.TrashInfoFileName = p.TrashInfoFileName.FromRef(js.Undefined)
	if recursive {
		p.RestoreEntry.FreeMembers(true)
	}
}

type RemoveFileWatchCallbackFunc func(this js.Ref, success bool) js.Ref

func (fn RemoveFileWatchCallbackFunc) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RemoveFileWatchCallbackFunc) DispatchCallback(
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

type RemoveFileWatchCallback[T any] struct {
	Fn  func(arg T, this js.Ref, success bool) js.Ref
	Arg T
}

func (cb *RemoveFileWatchCallback[T]) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RemoveFileWatchCallback[T]) DispatchCallback(
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

type ResolveIsolatedEntriesCallbackFunc func(this js.Ref, entries js.Array[EntryDescription]) js.Ref

func (fn ResolveIsolatedEntriesCallbackFunc) Register() js.Func[func(entries js.Array[EntryDescription])] {
	return js.RegisterCallback[func(entries js.Array[EntryDescription])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ResolveIsolatedEntriesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[EntryDescription]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ResolveIsolatedEntriesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[EntryDescription]) js.Ref
	Arg T
}

func (cb *ResolveIsolatedEntriesCallback[T]) Register() js.Func[func(entries js.Array[EntryDescription])] {
	return js.RegisterCallback[func(entries js.Array[EntryDescription])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ResolveIsolatedEntriesCallback[T]) DispatchCallback(
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

		js.Array[EntryDescription]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SearchFilesCallbackFunc func(this js.Ref, entries js.Array[EntryDescription]) js.Ref

func (fn SearchFilesCallbackFunc) Register() js.Func[func(entries js.Array[EntryDescription])] {
	return js.RegisterCallback[func(entries js.Array[EntryDescription])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SearchFilesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[EntryDescription]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SearchFilesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[EntryDescription]) js.Ref
	Arg T
}

func (cb *SearchFilesCallback[T]) Register() js.Func[func(entries js.Array[EntryDescription])] {
	return js.RegisterCallback[func(entries js.Array[EntryDescription])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SearchFilesCallback[T]) DispatchCallback(
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

		js.Array[EntryDescription]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SearchFilesParams struct {
	// RootUrl is "SearchFilesParams.rootUrl"
	//
	// Optional
	RootUrl js.String
	// Query is "SearchFilesParams.query"
	//
	// Optional
	Query js.String
	// Types is "SearchFilesParams.types"
	//
	// Optional
	Types filemanagerprivate.SearchType
	// MaxResults is "SearchFilesParams.maxResults"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxResults MUST be set to true to make this field effective.
	MaxResults int32
	// ModifiedTimestamp is "SearchFilesParams.modifiedTimestamp"
	//
	// Optional
	//
	// NOTE: FFI_USE_ModifiedTimestamp MUST be set to true to make this field effective.
	ModifiedTimestamp float64
	// Category is "SearchFilesParams.category"
	//
	// Optional
	Category filemanagerprivate.FileCategory

	FFI_USE_MaxResults        bool // for MaxResults.
	FFI_USE_ModifiedTimestamp bool // for ModifiedTimestamp.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SearchFilesParams with all fields set.
func (p SearchFilesParams) FromRef(ref js.Ref) SearchFilesParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SearchFilesParams in the application heap.
func (p SearchFilesParams) New() js.Ref {
	return bindings.SearchFilesParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SearchFilesParams) UpdateFrom(ref js.Ref) {
	bindings.SearchFilesParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SearchFilesParams) Update(ref js.Ref) {
	bindings.SearchFilesParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SearchFilesParams) FreeMembers(recursive bool) {
	js.Free(
		p.RootUrl.Ref(),
		p.Query.Ref(),
	)
	p.RootUrl = p.RootUrl.FromRef(js.Undefined)
	p.Query = p.Query.FromRef(js.Undefined)
}

type SimpleCallbackFunc func(this js.Ref) js.Ref

func (fn SimpleCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SimpleCallbackFunc) DispatchCallback(
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

type SimpleCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *SimpleCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SimpleCallback[T]) DispatchCallback(
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

type ValidatePathNameLengthCallbackFunc func(this js.Ref, result bool) js.Ref

func (fn ValidatePathNameLengthCallbackFunc) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ValidatePathNameLengthCallbackFunc) DispatchCallback(
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

type ValidatePathNameLengthCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result bool) js.Ref
	Arg T
}

func (cb *ValidatePathNameLengthCallback[T]) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ValidatePathNameLengthCallback[T]) DispatchCallback(
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

// HasFuncAddFileWatch returns true if the function "WEBEXT.fileManagerPrivateInternal.addFileWatch" exists.
func HasFuncAddFileWatch() bool {
	return js.True == bindings.HasFuncAddFileWatch()
}

// FuncAddFileWatch returns the function "WEBEXT.fileManagerPrivateInternal.addFileWatch".
func FuncAddFileWatch() (fn js.Func[func(url js.String) js.Promise[js.Boolean]]) {
	bindings.FuncAddFileWatch(
		js.Pointer(&fn),
	)
	return
}

// AddFileWatch calls the function "WEBEXT.fileManagerPrivateInternal.addFileWatch" directly.
func AddFileWatch(url js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallAddFileWatch(
		js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryAddFileWatch calls the function "WEBEXT.fileManagerPrivateInternal.addFileWatch"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAddFileWatch(url js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAddFileWatch(
		js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasFuncComputeChecksum returns true if the function "WEBEXT.fileManagerPrivateInternal.computeChecksum" exists.
func HasFuncComputeChecksum() bool {
	return js.True == bindings.HasFuncComputeChecksum()
}

// FuncComputeChecksum returns the function "WEBEXT.fileManagerPrivateInternal.computeChecksum".
func FuncComputeChecksum() (fn js.Func[func(url js.String) js.Promise[js.String]]) {
	bindings.FuncComputeChecksum(
		js.Pointer(&fn),
	)
	return
}

// ComputeChecksum calls the function "WEBEXT.fileManagerPrivateInternal.computeChecksum" directly.
func ComputeChecksum(url js.String) (ret js.Promise[js.String]) {
	bindings.CallComputeChecksum(
		js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryComputeChecksum calls the function "WEBEXT.fileManagerPrivateInternal.computeChecksum"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryComputeChecksum(url js.String) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryComputeChecksum(
		js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasFuncExecuteCustomAction returns true if the function "WEBEXT.fileManagerPrivateInternal.executeCustomAction" exists.
func HasFuncExecuteCustomAction() bool {
	return js.True == bindings.HasFuncExecuteCustomAction()
}

// FuncExecuteCustomAction returns the function "WEBEXT.fileManagerPrivateInternal.executeCustomAction".
func FuncExecuteCustomAction() (fn js.Func[func(urls js.Array[js.String], actionId js.String) js.Promise[js.Void]]) {
	bindings.FuncExecuteCustomAction(
		js.Pointer(&fn),
	)
	return
}

// ExecuteCustomAction calls the function "WEBEXT.fileManagerPrivateInternal.executeCustomAction" directly.
func ExecuteCustomAction(urls js.Array[js.String], actionId js.String) (ret js.Promise[js.Void]) {
	bindings.CallExecuteCustomAction(
		js.Pointer(&ret),
		urls.Ref(),
		actionId.Ref(),
	)

	return
}

// TryExecuteCustomAction calls the function "WEBEXT.fileManagerPrivateInternal.executeCustomAction"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryExecuteCustomAction(urls js.Array[js.String], actionId js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryExecuteCustomAction(
		js.Pointer(&ret), js.Pointer(&exception),
		urls.Ref(),
		actionId.Ref(),
	)

	return
}

// HasFuncExecuteTask returns true if the function "WEBEXT.fileManagerPrivateInternal.executeTask" exists.
func HasFuncExecuteTask() bool {
	return js.True == bindings.HasFuncExecuteTask()
}

// FuncExecuteTask returns the function "WEBEXT.fileManagerPrivateInternal.executeTask".
func FuncExecuteTask() (fn js.Func[func(descriptor filemanagerprivate.FileTaskDescriptor, urls js.Array[js.String]) js.Promise[filemanagerprivate.TaskResult]]) {
	bindings.FuncExecuteTask(
		js.Pointer(&fn),
	)
	return
}

// ExecuteTask calls the function "WEBEXT.fileManagerPrivateInternal.executeTask" directly.
func ExecuteTask(descriptor filemanagerprivate.FileTaskDescriptor, urls js.Array[js.String]) (ret js.Promise[filemanagerprivate.TaskResult]) {
	bindings.CallExecuteTask(
		js.Pointer(&ret),
		js.Pointer(&descriptor),
		urls.Ref(),
	)

	return
}

// TryExecuteTask calls the function "WEBEXT.fileManagerPrivateInternal.executeTask"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryExecuteTask(descriptor filemanagerprivate.FileTaskDescriptor, urls js.Array[js.String]) (ret js.Promise[filemanagerprivate.TaskResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryExecuteTask(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
		urls.Ref(),
	)

	return
}

// HasFuncGetContentMetadata returns true if the function "WEBEXT.fileManagerPrivateInternal.getContentMetadata" exists.
func HasFuncGetContentMetadata() bool {
	return js.True == bindings.HasFuncGetContentMetadata()
}

// FuncGetContentMetadata returns the function "WEBEXT.fileManagerPrivateInternal.getContentMetadata".
func FuncGetContentMetadata() (fn js.Func[func(blobUUID js.String, mimeType js.String, includeImages bool) js.Promise[filemanagerprivate.MediaMetadata]]) {
	bindings.FuncGetContentMetadata(
		js.Pointer(&fn),
	)
	return
}

// GetContentMetadata calls the function "WEBEXT.fileManagerPrivateInternal.getContentMetadata" directly.
func GetContentMetadata(blobUUID js.String, mimeType js.String, includeImages bool) (ret js.Promise[filemanagerprivate.MediaMetadata]) {
	bindings.CallGetContentMetadata(
		js.Pointer(&ret),
		blobUUID.Ref(),
		mimeType.Ref(),
		js.Bool(bool(includeImages)),
	)

	return
}

// TryGetContentMetadata calls the function "WEBEXT.fileManagerPrivateInternal.getContentMetadata"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetContentMetadata(blobUUID js.String, mimeType js.String, includeImages bool) (ret js.Promise[filemanagerprivate.MediaMetadata], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetContentMetadata(
		js.Pointer(&ret), js.Pointer(&exception),
		blobUUID.Ref(),
		mimeType.Ref(),
		js.Bool(bool(includeImages)),
	)

	return
}

// HasFuncGetContentMimeType returns true if the function "WEBEXT.fileManagerPrivateInternal.getContentMimeType" exists.
func HasFuncGetContentMimeType() bool {
	return js.True == bindings.HasFuncGetContentMimeType()
}

// FuncGetContentMimeType returns the function "WEBEXT.fileManagerPrivateInternal.getContentMimeType".
func FuncGetContentMimeType() (fn js.Func[func(blobUUID js.String) js.Promise[js.String]]) {
	bindings.FuncGetContentMimeType(
		js.Pointer(&fn),
	)
	return
}

// GetContentMimeType calls the function "WEBEXT.fileManagerPrivateInternal.getContentMimeType" directly.
func GetContentMimeType(blobUUID js.String) (ret js.Promise[js.String]) {
	bindings.CallGetContentMimeType(
		js.Pointer(&ret),
		blobUUID.Ref(),
	)

	return
}

// TryGetContentMimeType calls the function "WEBEXT.fileManagerPrivateInternal.getContentMimeType"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetContentMimeType(blobUUID js.String) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetContentMimeType(
		js.Pointer(&ret), js.Pointer(&exception),
		blobUUID.Ref(),
	)

	return
}

// HasFuncGetCrostiniSharedPaths returns true if the function "WEBEXT.fileManagerPrivateInternal.getCrostiniSharedPaths" exists.
func HasFuncGetCrostiniSharedPaths() bool {
	return js.True == bindings.HasFuncGetCrostiniSharedPaths()
}

// FuncGetCrostiniSharedPaths returns the function "WEBEXT.fileManagerPrivateInternal.getCrostiniSharedPaths".
func FuncGetCrostiniSharedPaths() (fn js.Func[func(observeFirstForSession bool, vmName js.String, callback js.Func[func(entries js.Array[EntryDescription], firstForSession bool)])]) {
	bindings.FuncGetCrostiniSharedPaths(
		js.Pointer(&fn),
	)
	return
}

// GetCrostiniSharedPaths calls the function "WEBEXT.fileManagerPrivateInternal.getCrostiniSharedPaths" directly.
func GetCrostiniSharedPaths(observeFirstForSession bool, vmName js.String, callback js.Func[func(entries js.Array[EntryDescription], firstForSession bool)]) (ret js.Void) {
	bindings.CallGetCrostiniSharedPaths(
		js.Pointer(&ret),
		js.Bool(bool(observeFirstForSession)),
		vmName.Ref(),
		callback.Ref(),
	)

	return
}

// TryGetCrostiniSharedPaths calls the function "WEBEXT.fileManagerPrivateInternal.getCrostiniSharedPaths"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetCrostiniSharedPaths(observeFirstForSession bool, vmName js.String, callback js.Func[func(entries js.Array[EntryDescription], firstForSession bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetCrostiniSharedPaths(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(observeFirstForSession)),
		vmName.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncGetCustomActions returns true if the function "WEBEXT.fileManagerPrivateInternal.getCustomActions" exists.
func HasFuncGetCustomActions() bool {
	return js.True == bindings.HasFuncGetCustomActions()
}

// FuncGetCustomActions returns the function "WEBEXT.fileManagerPrivateInternal.getCustomActions".
func FuncGetCustomActions() (fn js.Func[func(urls js.Array[js.String]) js.Promise[js.Array[filesystemprovider.Action]]]) {
	bindings.FuncGetCustomActions(
		js.Pointer(&fn),
	)
	return
}

// GetCustomActions calls the function "WEBEXT.fileManagerPrivateInternal.getCustomActions" directly.
func GetCustomActions(urls js.Array[js.String]) (ret js.Promise[js.Array[filesystemprovider.Action]]) {
	bindings.CallGetCustomActions(
		js.Pointer(&ret),
		urls.Ref(),
	)

	return
}

// TryGetCustomActions calls the function "WEBEXT.fileManagerPrivateInternal.getCustomActions"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetCustomActions(urls js.Array[js.String]) (ret js.Promise[js.Array[filesystemprovider.Action]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetCustomActions(
		js.Pointer(&ret), js.Pointer(&exception),
		urls.Ref(),
	)

	return
}

// HasFuncGetDirectorySize returns true if the function "WEBEXT.fileManagerPrivateInternal.getDirectorySize" exists.
func HasFuncGetDirectorySize() bool {
	return js.True == bindings.HasFuncGetDirectorySize()
}

// FuncGetDirectorySize returns the function "WEBEXT.fileManagerPrivateInternal.getDirectorySize".
func FuncGetDirectorySize() (fn js.Func[func(url js.String) js.Promise[js.Number[float64]]]) {
	bindings.FuncGetDirectorySize(
		js.Pointer(&fn),
	)
	return
}

// GetDirectorySize calls the function "WEBEXT.fileManagerPrivateInternal.getDirectorySize" directly.
func GetDirectorySize(url js.String) (ret js.Promise[js.Number[float64]]) {
	bindings.CallGetDirectorySize(
		js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryGetDirectorySize calls the function "WEBEXT.fileManagerPrivateInternal.getDirectorySize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDirectorySize(url js.String) (ret js.Promise[js.Number[float64]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDirectorySize(
		js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasFuncGetDisallowedTransfers returns true if the function "WEBEXT.fileManagerPrivateInternal.getDisallowedTransfers" exists.
func HasFuncGetDisallowedTransfers() bool {
	return js.True == bindings.HasFuncGetDisallowedTransfers()
}

// FuncGetDisallowedTransfers returns the function "WEBEXT.fileManagerPrivateInternal.getDisallowedTransfers".
func FuncGetDisallowedTransfers() (fn js.Func[func(entries js.Array[js.String], destinationEntry js.String, isMove bool) js.Promise[js.Array[EntryDescription]]]) {
	bindings.FuncGetDisallowedTransfers(
		js.Pointer(&fn),
	)
	return
}

// GetDisallowedTransfers calls the function "WEBEXT.fileManagerPrivateInternal.getDisallowedTransfers" directly.
func GetDisallowedTransfers(entries js.Array[js.String], destinationEntry js.String, isMove bool) (ret js.Promise[js.Array[EntryDescription]]) {
	bindings.CallGetDisallowedTransfers(
		js.Pointer(&ret),
		entries.Ref(),
		destinationEntry.Ref(),
		js.Bool(bool(isMove)),
	)

	return
}

// TryGetDisallowedTransfers calls the function "WEBEXT.fileManagerPrivateInternal.getDisallowedTransfers"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDisallowedTransfers(entries js.Array[js.String], destinationEntry js.String, isMove bool) (ret js.Promise[js.Array[EntryDescription]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDisallowedTransfers(
		js.Pointer(&ret), js.Pointer(&exception),
		entries.Ref(),
		destinationEntry.Ref(),
		js.Bool(bool(isMove)),
	)

	return
}

// HasFuncGetDlpMetadata returns true if the function "WEBEXT.fileManagerPrivateInternal.getDlpMetadata" exists.
func HasFuncGetDlpMetadata() bool {
	return js.True == bindings.HasFuncGetDlpMetadata()
}

// FuncGetDlpMetadata returns the function "WEBEXT.fileManagerPrivateInternal.getDlpMetadata".
func FuncGetDlpMetadata() (fn js.Func[func(entries js.Array[js.String]) js.Promise[js.Array[filemanagerprivate.DlpMetadata]]]) {
	bindings.FuncGetDlpMetadata(
		js.Pointer(&fn),
	)
	return
}

// GetDlpMetadata calls the function "WEBEXT.fileManagerPrivateInternal.getDlpMetadata" directly.
func GetDlpMetadata(entries js.Array[js.String]) (ret js.Promise[js.Array[filemanagerprivate.DlpMetadata]]) {
	bindings.CallGetDlpMetadata(
		js.Pointer(&ret),
		entries.Ref(),
	)

	return
}

// TryGetDlpMetadata calls the function "WEBEXT.fileManagerPrivateInternal.getDlpMetadata"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDlpMetadata(entries js.Array[js.String]) (ret js.Promise[js.Array[filemanagerprivate.DlpMetadata]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDlpMetadata(
		js.Pointer(&ret), js.Pointer(&exception),
		entries.Ref(),
	)

	return
}

// HasFuncGetDriveQuotaMetadata returns true if the function "WEBEXT.fileManagerPrivateInternal.getDriveQuotaMetadata" exists.
func HasFuncGetDriveQuotaMetadata() bool {
	return js.True == bindings.HasFuncGetDriveQuotaMetadata()
}

// FuncGetDriveQuotaMetadata returns the function "WEBEXT.fileManagerPrivateInternal.getDriveQuotaMetadata".
func FuncGetDriveQuotaMetadata() (fn js.Func[func(url js.String) js.Promise[filemanagerprivate.DriveQuotaMetadata]]) {
	bindings.FuncGetDriveQuotaMetadata(
		js.Pointer(&fn),
	)
	return
}

// GetDriveQuotaMetadata calls the function "WEBEXT.fileManagerPrivateInternal.getDriveQuotaMetadata" directly.
func GetDriveQuotaMetadata(url js.String) (ret js.Promise[filemanagerprivate.DriveQuotaMetadata]) {
	bindings.CallGetDriveQuotaMetadata(
		js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryGetDriveQuotaMetadata calls the function "WEBEXT.fileManagerPrivateInternal.getDriveQuotaMetadata"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDriveQuotaMetadata(url js.String) (ret js.Promise[filemanagerprivate.DriveQuotaMetadata], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDriveQuotaMetadata(
		js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasFuncGetEntryProperties returns true if the function "WEBEXT.fileManagerPrivateInternal.getEntryProperties" exists.
func HasFuncGetEntryProperties() bool {
	return js.True == bindings.HasFuncGetEntryProperties()
}

// FuncGetEntryProperties returns the function "WEBEXT.fileManagerPrivateInternal.getEntryProperties".
func FuncGetEntryProperties() (fn js.Func[func(urls js.Array[js.String], names js.Array[filemanagerprivate.EntryPropertyName]) js.Promise[js.Array[filemanagerprivate.EntryProperties]]]) {
	bindings.FuncGetEntryProperties(
		js.Pointer(&fn),
	)
	return
}

// GetEntryProperties calls the function "WEBEXT.fileManagerPrivateInternal.getEntryProperties" directly.
func GetEntryProperties(urls js.Array[js.String], names js.Array[filemanagerprivate.EntryPropertyName]) (ret js.Promise[js.Array[filemanagerprivate.EntryProperties]]) {
	bindings.CallGetEntryProperties(
		js.Pointer(&ret),
		urls.Ref(),
		names.Ref(),
	)

	return
}

// TryGetEntryProperties calls the function "WEBEXT.fileManagerPrivateInternal.getEntryProperties"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetEntryProperties(urls js.Array[js.String], names js.Array[filemanagerprivate.EntryPropertyName]) (ret js.Promise[js.Array[filemanagerprivate.EntryProperties]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetEntryProperties(
		js.Pointer(&ret), js.Pointer(&exception),
		urls.Ref(),
		names.Ref(),
	)

	return
}

// HasFuncGetFileTasks returns true if the function "WEBEXT.fileManagerPrivateInternal.getFileTasks" exists.
func HasFuncGetFileTasks() bool {
	return js.True == bindings.HasFuncGetFileTasks()
}

// FuncGetFileTasks returns the function "WEBEXT.fileManagerPrivateInternal.getFileTasks".
func FuncGetFileTasks() (fn js.Func[func(urls js.Array[js.String], dlpSourceUrls js.Array[js.String]) js.Promise[filemanagerprivate.ResultingTasks]]) {
	bindings.FuncGetFileTasks(
		js.Pointer(&fn),
	)
	return
}

// GetFileTasks calls the function "WEBEXT.fileManagerPrivateInternal.getFileTasks" directly.
func GetFileTasks(urls js.Array[js.String], dlpSourceUrls js.Array[js.String]) (ret js.Promise[filemanagerprivate.ResultingTasks]) {
	bindings.CallGetFileTasks(
		js.Pointer(&ret),
		urls.Ref(),
		dlpSourceUrls.Ref(),
	)

	return
}

// TryGetFileTasks calls the function "WEBEXT.fileManagerPrivateInternal.getFileTasks"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetFileTasks(urls js.Array[js.String], dlpSourceUrls js.Array[js.String]) (ret js.Promise[filemanagerprivate.ResultingTasks], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetFileTasks(
		js.Pointer(&ret), js.Pointer(&exception),
		urls.Ref(),
		dlpSourceUrls.Ref(),
	)

	return
}

// HasFuncGetLinuxPackageInfo returns true if the function "WEBEXT.fileManagerPrivateInternal.getLinuxPackageInfo" exists.
func HasFuncGetLinuxPackageInfo() bool {
	return js.True == bindings.HasFuncGetLinuxPackageInfo()
}

// FuncGetLinuxPackageInfo returns the function "WEBEXT.fileManagerPrivateInternal.getLinuxPackageInfo".
func FuncGetLinuxPackageInfo() (fn js.Func[func(url js.String) js.Promise[filemanagerprivate.LinuxPackageInfo]]) {
	bindings.FuncGetLinuxPackageInfo(
		js.Pointer(&fn),
	)
	return
}

// GetLinuxPackageInfo calls the function "WEBEXT.fileManagerPrivateInternal.getLinuxPackageInfo" directly.
func GetLinuxPackageInfo(url js.String) (ret js.Promise[filemanagerprivate.LinuxPackageInfo]) {
	bindings.CallGetLinuxPackageInfo(
		js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryGetLinuxPackageInfo calls the function "WEBEXT.fileManagerPrivateInternal.getLinuxPackageInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetLinuxPackageInfo(url js.String) (ret js.Promise[filemanagerprivate.LinuxPackageInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetLinuxPackageInfo(
		js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasFuncGetMimeType returns true if the function "WEBEXT.fileManagerPrivateInternal.getMimeType" exists.
func HasFuncGetMimeType() bool {
	return js.True == bindings.HasFuncGetMimeType()
}

// FuncGetMimeType returns the function "WEBEXT.fileManagerPrivateInternal.getMimeType".
func FuncGetMimeType() (fn js.Func[func(url js.String) js.Promise[js.String]]) {
	bindings.FuncGetMimeType(
		js.Pointer(&fn),
	)
	return
}

// GetMimeType calls the function "WEBEXT.fileManagerPrivateInternal.getMimeType" directly.
func GetMimeType(url js.String) (ret js.Promise[js.String]) {
	bindings.CallGetMimeType(
		js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryGetMimeType calls the function "WEBEXT.fileManagerPrivateInternal.getMimeType"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetMimeType(url js.String) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetMimeType(
		js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasFuncGetRecentFiles returns true if the function "WEBEXT.fileManagerPrivateInternal.getRecentFiles" exists.
func HasFuncGetRecentFiles() bool {
	return js.True == bindings.HasFuncGetRecentFiles()
}

// FuncGetRecentFiles returns the function "WEBEXT.fileManagerPrivateInternal.getRecentFiles".
func FuncGetRecentFiles() (fn js.Func[func(restriction filemanagerprivate.SourceRestriction, file_category filemanagerprivate.FileCategory, invalidate_cache bool) js.Promise[js.Array[EntryDescription]]]) {
	bindings.FuncGetRecentFiles(
		js.Pointer(&fn),
	)
	return
}

// GetRecentFiles calls the function "WEBEXT.fileManagerPrivateInternal.getRecentFiles" directly.
func GetRecentFiles(restriction filemanagerprivate.SourceRestriction, file_category filemanagerprivate.FileCategory, invalidate_cache bool) (ret js.Promise[js.Array[EntryDescription]]) {
	bindings.CallGetRecentFiles(
		js.Pointer(&ret),
		uint32(restriction),
		uint32(file_category),
		js.Bool(bool(invalidate_cache)),
	)

	return
}

// TryGetRecentFiles calls the function "WEBEXT.fileManagerPrivateInternal.getRecentFiles"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetRecentFiles(restriction filemanagerprivate.SourceRestriction, file_category filemanagerprivate.FileCategory, invalidate_cache bool) (ret js.Promise[js.Array[EntryDescription]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetRecentFiles(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(restriction),
		uint32(file_category),
		js.Bool(bool(invalidate_cache)),
	)

	return
}

// HasFuncGetVolumeRoot returns true if the function "WEBEXT.fileManagerPrivateInternal.getVolumeRoot" exists.
func HasFuncGetVolumeRoot() bool {
	return js.True == bindings.HasFuncGetVolumeRoot()
}

// FuncGetVolumeRoot returns the function "WEBEXT.fileManagerPrivateInternal.getVolumeRoot".
func FuncGetVolumeRoot() (fn js.Func[func(options filemanagerprivate.GetVolumeRootOptions) js.Promise[EntryDescription]]) {
	bindings.FuncGetVolumeRoot(
		js.Pointer(&fn),
	)
	return
}

// GetVolumeRoot calls the function "WEBEXT.fileManagerPrivateInternal.getVolumeRoot" directly.
func GetVolumeRoot(options filemanagerprivate.GetVolumeRootOptions) (ret js.Promise[EntryDescription]) {
	bindings.CallGetVolumeRoot(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetVolumeRoot calls the function "WEBEXT.fileManagerPrivateInternal.getVolumeRoot"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetVolumeRoot(options filemanagerprivate.GetVolumeRootOptions) (ret js.Promise[EntryDescription], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetVolumeRoot(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncImportCrostiniImage returns true if the function "WEBEXT.fileManagerPrivateInternal.importCrostiniImage" exists.
func HasFuncImportCrostiniImage() bool {
	return js.True == bindings.HasFuncImportCrostiniImage()
}

// FuncImportCrostiniImage returns the function "WEBEXT.fileManagerPrivateInternal.importCrostiniImage".
func FuncImportCrostiniImage() (fn js.Func[func(url js.String)]) {
	bindings.FuncImportCrostiniImage(
		js.Pointer(&fn),
	)
	return
}

// ImportCrostiniImage calls the function "WEBEXT.fileManagerPrivateInternal.importCrostiniImage" directly.
func ImportCrostiniImage(url js.String) (ret js.Void) {
	bindings.CallImportCrostiniImage(
		js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryImportCrostiniImage calls the function "WEBEXT.fileManagerPrivateInternal.importCrostiniImage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryImportCrostiniImage(url js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryImportCrostiniImage(
		js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasFuncInstallLinuxPackage returns true if the function "WEBEXT.fileManagerPrivateInternal.installLinuxPackage" exists.
func HasFuncInstallLinuxPackage() bool {
	return js.True == bindings.HasFuncInstallLinuxPackage()
}

// FuncInstallLinuxPackage returns the function "WEBEXT.fileManagerPrivateInternal.installLinuxPackage".
func FuncInstallLinuxPackage() (fn js.Func[func(url js.String, callback js.Func[func(response filemanagerprivate.InstallLinuxPackageResponse)])]) {
	bindings.FuncInstallLinuxPackage(
		js.Pointer(&fn),
	)
	return
}

// InstallLinuxPackage calls the function "WEBEXT.fileManagerPrivateInternal.installLinuxPackage" directly.
func InstallLinuxPackage(url js.String, callback js.Func[func(response filemanagerprivate.InstallLinuxPackageResponse)]) (ret js.Void) {
	bindings.CallInstallLinuxPackage(
		js.Pointer(&ret),
		url.Ref(),
		callback.Ref(),
	)

	return
}

// TryInstallLinuxPackage calls the function "WEBEXT.fileManagerPrivateInternal.installLinuxPackage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryInstallLinuxPackage(url js.String, callback js.Func[func(response filemanagerprivate.InstallLinuxPackageResponse)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryInstallLinuxPackage(
		js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncInvokeSharesheet returns true if the function "WEBEXT.fileManagerPrivateInternal.invokeSharesheet" exists.
func HasFuncInvokeSharesheet() bool {
	return js.True == bindings.HasFuncInvokeSharesheet()
}

// FuncInvokeSharesheet returns the function "WEBEXT.fileManagerPrivateInternal.invokeSharesheet".
func FuncInvokeSharesheet() (fn js.Func[func(urls js.Array[js.String], launchSource filemanagerprivate.SharesheetLaunchSource, dlpSourceUrls js.Array[js.String]) js.Promise[js.Void]]) {
	bindings.FuncInvokeSharesheet(
		js.Pointer(&fn),
	)
	return
}

// InvokeSharesheet calls the function "WEBEXT.fileManagerPrivateInternal.invokeSharesheet" directly.
func InvokeSharesheet(urls js.Array[js.String], launchSource filemanagerprivate.SharesheetLaunchSource, dlpSourceUrls js.Array[js.String]) (ret js.Promise[js.Void]) {
	bindings.CallInvokeSharesheet(
		js.Pointer(&ret),
		urls.Ref(),
		uint32(launchSource),
		dlpSourceUrls.Ref(),
	)

	return
}

// TryInvokeSharesheet calls the function "WEBEXT.fileManagerPrivateInternal.invokeSharesheet"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryInvokeSharesheet(urls js.Array[js.String], launchSource filemanagerprivate.SharesheetLaunchSource, dlpSourceUrls js.Array[js.String]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryInvokeSharesheet(
		js.Pointer(&ret), js.Pointer(&exception),
		urls.Ref(),
		uint32(launchSource),
		dlpSourceUrls.Ref(),
	)

	return
}

// HasFuncParseTrashInfoFiles returns true if the function "WEBEXT.fileManagerPrivateInternal.parseTrashInfoFiles" exists.
func HasFuncParseTrashInfoFiles() bool {
	return js.True == bindings.HasFuncParseTrashInfoFiles()
}

// FuncParseTrashInfoFiles returns the function "WEBEXT.fileManagerPrivateInternal.parseTrashInfoFiles".
func FuncParseTrashInfoFiles() (fn js.Func[func(urls js.Array[js.String]) js.Promise[js.Array[ParsedTrashInfoFile]]]) {
	bindings.FuncParseTrashInfoFiles(
		js.Pointer(&fn),
	)
	return
}

// ParseTrashInfoFiles calls the function "WEBEXT.fileManagerPrivateInternal.parseTrashInfoFiles" directly.
func ParseTrashInfoFiles(urls js.Array[js.String]) (ret js.Promise[js.Array[ParsedTrashInfoFile]]) {
	bindings.CallParseTrashInfoFiles(
		js.Pointer(&ret),
		urls.Ref(),
	)

	return
}

// TryParseTrashInfoFiles calls the function "WEBEXT.fileManagerPrivateInternal.parseTrashInfoFiles"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryParseTrashInfoFiles(urls js.Array[js.String]) (ret js.Promise[js.Array[ParsedTrashInfoFile]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryParseTrashInfoFiles(
		js.Pointer(&ret), js.Pointer(&exception),
		urls.Ref(),
	)

	return
}

// HasFuncPinDriveFile returns true if the function "WEBEXT.fileManagerPrivateInternal.pinDriveFile" exists.
func HasFuncPinDriveFile() bool {
	return js.True == bindings.HasFuncPinDriveFile()
}

// FuncPinDriveFile returns the function "WEBEXT.fileManagerPrivateInternal.pinDriveFile".
func FuncPinDriveFile() (fn js.Func[func(url js.String, pin bool) js.Promise[js.Void]]) {
	bindings.FuncPinDriveFile(
		js.Pointer(&fn),
	)
	return
}

// PinDriveFile calls the function "WEBEXT.fileManagerPrivateInternal.pinDriveFile" directly.
func PinDriveFile(url js.String, pin bool) (ret js.Promise[js.Void]) {
	bindings.CallPinDriveFile(
		js.Pointer(&ret),
		url.Ref(),
		js.Bool(bool(pin)),
	)

	return
}

// TryPinDriveFile calls the function "WEBEXT.fileManagerPrivateInternal.pinDriveFile"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryPinDriveFile(url js.String, pin bool) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPinDriveFile(
		js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
		js.Bool(bool(pin)),
	)

	return
}

// HasFuncRemoveFileWatch returns true if the function "WEBEXT.fileManagerPrivateInternal.removeFileWatch" exists.
func HasFuncRemoveFileWatch() bool {
	return js.True == bindings.HasFuncRemoveFileWatch()
}

// FuncRemoveFileWatch returns the function "WEBEXT.fileManagerPrivateInternal.removeFileWatch".
func FuncRemoveFileWatch() (fn js.Func[func(url js.String) js.Promise[js.Boolean]]) {
	bindings.FuncRemoveFileWatch(
		js.Pointer(&fn),
	)
	return
}

// RemoveFileWatch calls the function "WEBEXT.fileManagerPrivateInternal.removeFileWatch" directly.
func RemoveFileWatch(url js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallRemoveFileWatch(
		js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryRemoveFileWatch calls the function "WEBEXT.fileManagerPrivateInternal.removeFileWatch"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveFileWatch(url js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveFileWatch(
		js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasFuncResolveIsolatedEntries returns true if the function "WEBEXT.fileManagerPrivateInternal.resolveIsolatedEntries" exists.
func HasFuncResolveIsolatedEntries() bool {
	return js.True == bindings.HasFuncResolveIsolatedEntries()
}

// FuncResolveIsolatedEntries returns the function "WEBEXT.fileManagerPrivateInternal.resolveIsolatedEntries".
func FuncResolveIsolatedEntries() (fn js.Func[func(urls js.Array[js.String]) js.Promise[js.Array[EntryDescription]]]) {
	bindings.FuncResolveIsolatedEntries(
		js.Pointer(&fn),
	)
	return
}

// ResolveIsolatedEntries calls the function "WEBEXT.fileManagerPrivateInternal.resolveIsolatedEntries" directly.
func ResolveIsolatedEntries(urls js.Array[js.String]) (ret js.Promise[js.Array[EntryDescription]]) {
	bindings.CallResolveIsolatedEntries(
		js.Pointer(&ret),
		urls.Ref(),
	)

	return
}

// TryResolveIsolatedEntries calls the function "WEBEXT.fileManagerPrivateInternal.resolveIsolatedEntries"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryResolveIsolatedEntries(urls js.Array[js.String]) (ret js.Promise[js.Array[EntryDescription]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryResolveIsolatedEntries(
		js.Pointer(&ret), js.Pointer(&exception),
		urls.Ref(),
	)

	return
}

// HasFuncSearchFiles returns true if the function "WEBEXT.fileManagerPrivateInternal.searchFiles" exists.
func HasFuncSearchFiles() bool {
	return js.True == bindings.HasFuncSearchFiles()
}

// FuncSearchFiles returns the function "WEBEXT.fileManagerPrivateInternal.searchFiles".
func FuncSearchFiles() (fn js.Func[func(searchParams SearchFilesParams) js.Promise[js.Array[EntryDescription]]]) {
	bindings.FuncSearchFiles(
		js.Pointer(&fn),
	)
	return
}

// SearchFiles calls the function "WEBEXT.fileManagerPrivateInternal.searchFiles" directly.
func SearchFiles(searchParams SearchFilesParams) (ret js.Promise[js.Array[EntryDescription]]) {
	bindings.CallSearchFiles(
		js.Pointer(&ret),
		js.Pointer(&searchParams),
	)

	return
}

// TrySearchFiles calls the function "WEBEXT.fileManagerPrivateInternal.searchFiles"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySearchFiles(searchParams SearchFilesParams) (ret js.Promise[js.Array[EntryDescription]], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySearchFiles(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&searchParams),
	)

	return
}

// HasFuncSetDefaultTask returns true if the function "WEBEXT.fileManagerPrivateInternal.setDefaultTask" exists.
func HasFuncSetDefaultTask() bool {
	return js.True == bindings.HasFuncSetDefaultTask()
}

// FuncSetDefaultTask returns the function "WEBEXT.fileManagerPrivateInternal.setDefaultTask".
func FuncSetDefaultTask() (fn js.Func[func(descriptor filemanagerprivate.FileTaskDescriptor, urls js.Array[js.String], mimeTypes js.Array[js.String]) js.Promise[js.Void]]) {
	bindings.FuncSetDefaultTask(
		js.Pointer(&fn),
	)
	return
}

// SetDefaultTask calls the function "WEBEXT.fileManagerPrivateInternal.setDefaultTask" directly.
func SetDefaultTask(descriptor filemanagerprivate.FileTaskDescriptor, urls js.Array[js.String], mimeTypes js.Array[js.String]) (ret js.Promise[js.Void]) {
	bindings.CallSetDefaultTask(
		js.Pointer(&ret),
		js.Pointer(&descriptor),
		urls.Ref(),
		mimeTypes.Ref(),
	)

	return
}

// TrySetDefaultTask calls the function "WEBEXT.fileManagerPrivateInternal.setDefaultTask"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetDefaultTask(descriptor filemanagerprivate.FileTaskDescriptor, urls js.Array[js.String], mimeTypes js.Array[js.String]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetDefaultTask(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
		urls.Ref(),
		mimeTypes.Ref(),
	)

	return
}

// HasFuncSharePathsWithCrostini returns true if the function "WEBEXT.fileManagerPrivateInternal.sharePathsWithCrostini" exists.
func HasFuncSharePathsWithCrostini() bool {
	return js.True == bindings.HasFuncSharePathsWithCrostini()
}

// FuncSharePathsWithCrostini returns the function "WEBEXT.fileManagerPrivateInternal.sharePathsWithCrostini".
func FuncSharePathsWithCrostini() (fn js.Func[func(vmName js.String, urls js.Array[js.String], persist bool) js.Promise[js.Void]]) {
	bindings.FuncSharePathsWithCrostini(
		js.Pointer(&fn),
	)
	return
}

// SharePathsWithCrostini calls the function "WEBEXT.fileManagerPrivateInternal.sharePathsWithCrostini" directly.
func SharePathsWithCrostini(vmName js.String, urls js.Array[js.String], persist bool) (ret js.Promise[js.Void]) {
	bindings.CallSharePathsWithCrostini(
		js.Pointer(&ret),
		vmName.Ref(),
		urls.Ref(),
		js.Bool(bool(persist)),
	)

	return
}

// TrySharePathsWithCrostini calls the function "WEBEXT.fileManagerPrivateInternal.sharePathsWithCrostini"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySharePathsWithCrostini(vmName js.String, urls js.Array[js.String], persist bool) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySharePathsWithCrostini(
		js.Pointer(&ret), js.Pointer(&exception),
		vmName.Ref(),
		urls.Ref(),
		js.Bool(bool(persist)),
	)

	return
}

// HasFuncSharesheetHasTargets returns true if the function "WEBEXT.fileManagerPrivateInternal.sharesheetHasTargets" exists.
func HasFuncSharesheetHasTargets() bool {
	return js.True == bindings.HasFuncSharesheetHasTargets()
}

// FuncSharesheetHasTargets returns the function "WEBEXT.fileManagerPrivateInternal.sharesheetHasTargets".
func FuncSharesheetHasTargets() (fn js.Func[func(urls js.Array[js.String]) js.Promise[js.Boolean]]) {
	bindings.FuncSharesheetHasTargets(
		js.Pointer(&fn),
	)
	return
}

// SharesheetHasTargets calls the function "WEBEXT.fileManagerPrivateInternal.sharesheetHasTargets" directly.
func SharesheetHasTargets(urls js.Array[js.String]) (ret js.Promise[js.Boolean]) {
	bindings.CallSharesheetHasTargets(
		js.Pointer(&ret),
		urls.Ref(),
	)

	return
}

// TrySharesheetHasTargets calls the function "WEBEXT.fileManagerPrivateInternal.sharesheetHasTargets"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySharesheetHasTargets(urls js.Array[js.String]) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySharesheetHasTargets(
		js.Pointer(&ret), js.Pointer(&exception),
		urls.Ref(),
	)

	return
}

// HasFuncStartIOTask returns true if the function "WEBEXT.fileManagerPrivateInternal.startIOTask" exists.
func HasFuncStartIOTask() bool {
	return js.True == bindings.HasFuncStartIOTask()
}

// FuncStartIOTask returns the function "WEBEXT.fileManagerPrivateInternal.startIOTask".
func FuncStartIOTask() (fn js.Func[func(typ filemanagerprivate.IOTaskType, urls js.Array[js.String], params IOTaskParams) js.Promise[js.Number[int32]]]) {
	bindings.FuncStartIOTask(
		js.Pointer(&fn),
	)
	return
}

// StartIOTask calls the function "WEBEXT.fileManagerPrivateInternal.startIOTask" directly.
func StartIOTask(typ filemanagerprivate.IOTaskType, urls js.Array[js.String], params IOTaskParams) (ret js.Promise[js.Number[int32]]) {
	bindings.CallStartIOTask(
		js.Pointer(&ret),
		uint32(typ),
		urls.Ref(),
		js.Pointer(&params),
	)

	return
}

// TryStartIOTask calls the function "WEBEXT.fileManagerPrivateInternal.startIOTask"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStartIOTask(typ filemanagerprivate.IOTaskType, urls js.Array[js.String], params IOTaskParams) (ret js.Promise[js.Number[int32]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStartIOTask(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(typ),
		urls.Ref(),
		js.Pointer(&params),
	)

	return
}

// HasFuncToggleAddedToHoldingSpace returns true if the function "WEBEXT.fileManagerPrivateInternal.toggleAddedToHoldingSpace" exists.
func HasFuncToggleAddedToHoldingSpace() bool {
	return js.True == bindings.HasFuncToggleAddedToHoldingSpace()
}

// FuncToggleAddedToHoldingSpace returns the function "WEBEXT.fileManagerPrivateInternal.toggleAddedToHoldingSpace".
func FuncToggleAddedToHoldingSpace() (fn js.Func[func(urls js.Array[js.String], add bool) js.Promise[js.Void]]) {
	bindings.FuncToggleAddedToHoldingSpace(
		js.Pointer(&fn),
	)
	return
}

// ToggleAddedToHoldingSpace calls the function "WEBEXT.fileManagerPrivateInternal.toggleAddedToHoldingSpace" directly.
func ToggleAddedToHoldingSpace(urls js.Array[js.String], add bool) (ret js.Promise[js.Void]) {
	bindings.CallToggleAddedToHoldingSpace(
		js.Pointer(&ret),
		urls.Ref(),
		js.Bool(bool(add)),
	)

	return
}

// TryToggleAddedToHoldingSpace calls the function "WEBEXT.fileManagerPrivateInternal.toggleAddedToHoldingSpace"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryToggleAddedToHoldingSpace(urls js.Array[js.String], add bool) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryToggleAddedToHoldingSpace(
		js.Pointer(&ret), js.Pointer(&exception),
		urls.Ref(),
		js.Bool(bool(add)),
	)

	return
}

// HasFuncUnsharePathWithCrostini returns true if the function "WEBEXT.fileManagerPrivateInternal.unsharePathWithCrostini" exists.
func HasFuncUnsharePathWithCrostini() bool {
	return js.True == bindings.HasFuncUnsharePathWithCrostini()
}

// FuncUnsharePathWithCrostini returns the function "WEBEXT.fileManagerPrivateInternal.unsharePathWithCrostini".
func FuncUnsharePathWithCrostini() (fn js.Func[func(vmName js.String, url js.String) js.Promise[js.Void]]) {
	bindings.FuncUnsharePathWithCrostini(
		js.Pointer(&fn),
	)
	return
}

// UnsharePathWithCrostini calls the function "WEBEXT.fileManagerPrivateInternal.unsharePathWithCrostini" directly.
func UnsharePathWithCrostini(vmName js.String, url js.String) (ret js.Promise[js.Void]) {
	bindings.CallUnsharePathWithCrostini(
		js.Pointer(&ret),
		vmName.Ref(),
		url.Ref(),
	)

	return
}

// TryUnsharePathWithCrostini calls the function "WEBEXT.fileManagerPrivateInternal.unsharePathWithCrostini"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUnsharePathWithCrostini(vmName js.String, url js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUnsharePathWithCrostini(
		js.Pointer(&ret), js.Pointer(&exception),
		vmName.Ref(),
		url.Ref(),
	)

	return
}

// HasFuncValidatePathNameLength returns true if the function "WEBEXT.fileManagerPrivateInternal.validatePathNameLength" exists.
func HasFuncValidatePathNameLength() bool {
	return js.True == bindings.HasFuncValidatePathNameLength()
}

// FuncValidatePathNameLength returns the function "WEBEXT.fileManagerPrivateInternal.validatePathNameLength".
func FuncValidatePathNameLength() (fn js.Func[func(parentUrl js.String, name js.String) js.Promise[js.Boolean]]) {
	bindings.FuncValidatePathNameLength(
		js.Pointer(&fn),
	)
	return
}

// ValidatePathNameLength calls the function "WEBEXT.fileManagerPrivateInternal.validatePathNameLength" directly.
func ValidatePathNameLength(parentUrl js.String, name js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallValidatePathNameLength(
		js.Pointer(&ret),
		parentUrl.Ref(),
		name.Ref(),
	)

	return
}

// TryValidatePathNameLength calls the function "WEBEXT.fileManagerPrivateInternal.validatePathNameLength"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryValidatePathNameLength(parentUrl js.String, name js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryValidatePathNameLength(
		js.Pointer(&ret), js.Pointer(&exception),
		parentUrl.Ref(),
		name.Ref(),
	)

	return
}
