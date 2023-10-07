// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package imagewriterprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/imagewriterprivate/bindings"
)

type DestroyPartitionsCallbackFunc func(this js.Ref) js.Ref

func (fn DestroyPartitionsCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DestroyPartitionsCallbackFunc) DispatchCallback(
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

type DestroyPartitionsCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *DestroyPartitionsCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DestroyPartitionsCallback[T]) DispatchCallback(
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

type ListRemovableStorageDevicesCallbackFunc func(this js.Ref, devices js.Array[RemovableStorageDevice]) js.Ref

func (fn ListRemovableStorageDevicesCallbackFunc) Register() js.Func[func(devices js.Array[RemovableStorageDevice])] {
	return js.RegisterCallback[func(devices js.Array[RemovableStorageDevice])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ListRemovableStorageDevicesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[RemovableStorageDevice]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ListRemovableStorageDevicesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, devices js.Array[RemovableStorageDevice]) js.Ref
	Arg T
}

func (cb *ListRemovableStorageDevicesCallback[T]) Register() js.Func[func(devices js.Array[RemovableStorageDevice])] {
	return js.RegisterCallback[func(devices js.Array[RemovableStorageDevice])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ListRemovableStorageDevicesCallback[T]) DispatchCallback(
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

		js.Array[RemovableStorageDevice]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type RemovableStorageDevice struct {
	// StorageUnitId is "RemovableStorageDevice.storageUnitId"
	//
	// Optional
	StorageUnitId js.String
	// Capacity is "RemovableStorageDevice.capacity"
	//
	// Optional
	//
	// NOTE: FFI_USE_Capacity MUST be set to true to make this field effective.
	Capacity float64
	// Vendor is "RemovableStorageDevice.vendor"
	//
	// Optional
	Vendor js.String
	// Model is "RemovableStorageDevice.model"
	//
	// Optional
	Model js.String
	// Removable is "RemovableStorageDevice.removable"
	//
	// Optional
	//
	// NOTE: FFI_USE_Removable MUST be set to true to make this field effective.
	Removable bool

	FFI_USE_Capacity  bool // for Capacity.
	FFI_USE_Removable bool // for Removable.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RemovableStorageDevice with all fields set.
func (p RemovableStorageDevice) FromRef(ref js.Ref) RemovableStorageDevice {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RemovableStorageDevice in the application heap.
func (p RemovableStorageDevice) New() js.Ref {
	return bindings.RemovableStorageDeviceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RemovableStorageDevice) UpdateFrom(ref js.Ref) {
	bindings.RemovableStorageDeviceJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RemovableStorageDevice) Update(ref js.Ref) {
	bindings.RemovableStorageDeviceJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RemovableStorageDevice) FreeMembers(recursive bool) {
	js.Free(
		p.StorageUnitId.Ref(),
		p.Vendor.Ref(),
		p.Model.Ref(),
	)
	p.StorageUnitId = p.StorageUnitId.FromRef(js.Undefined)
	p.Vendor = p.Vendor.FromRef(js.Undefined)
	p.Model = p.Model.FromRef(js.Undefined)
}

type Stage uint32

const (
	_ Stage = iota

	Stage_CONFIRMATION
	Stage_DOWNLOAD
	Stage_VERIFY_DOWNLOAD
	Stage_UNZIP
	Stage_WRITE
	Stage_VERIFY_WRITE
	Stage_UNKNOWN
)

func (Stage) FromRef(str js.Ref) Stage {
	return Stage(bindings.ConstOfStage(str))
}

func (x Stage) String() (string, bool) {
	switch x {
	case Stage_CONFIRMATION:
		return "confirmation", true
	case Stage_DOWNLOAD:
		return "download", true
	case Stage_VERIFY_DOWNLOAD:
		return "verifyDownload", true
	case Stage_UNZIP:
		return "unzip", true
	case Stage_WRITE:
		return "write", true
	case Stage_VERIFY_WRITE:
		return "verifyWrite", true
	case Stage_UNKNOWN:
		return "unknown", true
	default:
		return "", false
	}
}

type ProgressInfo struct {
	// Stage is "ProgressInfo.stage"
	//
	// Optional
	Stage Stage
	// PercentComplete is "ProgressInfo.percentComplete"
	//
	// Optional
	//
	// NOTE: FFI_USE_PercentComplete MUST be set to true to make this field effective.
	PercentComplete int32

	FFI_USE_PercentComplete bool // for PercentComplete.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProgressInfo with all fields set.
func (p ProgressInfo) FromRef(ref js.Ref) ProgressInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ProgressInfo in the application heap.
func (p ProgressInfo) New() js.Ref {
	return bindings.ProgressInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ProgressInfo) UpdateFrom(ref js.Ref) {
	bindings.ProgressInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ProgressInfo) Update(ref js.Ref) {
	bindings.ProgressInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ProgressInfo) FreeMembers(recursive bool) {
}

type UrlWriteOptions struct {
	// ImageHash is "UrlWriteOptions.imageHash"
	//
	// Optional
	ImageHash js.String
	// SaveAsDownload is "UrlWriteOptions.saveAsDownload"
	//
	// Optional
	//
	// NOTE: FFI_USE_SaveAsDownload MUST be set to true to make this field effective.
	SaveAsDownload bool

	FFI_USE_SaveAsDownload bool // for SaveAsDownload.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UrlWriteOptions with all fields set.
func (p UrlWriteOptions) FromRef(ref js.Ref) UrlWriteOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UrlWriteOptions in the application heap.
func (p UrlWriteOptions) New() js.Ref {
	return bindings.UrlWriteOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UrlWriteOptions) UpdateFrom(ref js.Ref) {
	bindings.UrlWriteOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UrlWriteOptions) Update(ref js.Ref) {
	bindings.UrlWriteOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UrlWriteOptions) FreeMembers(recursive bool) {
	js.Free(
		p.ImageHash.Ref(),
	)
	p.ImageHash = p.ImageHash.FromRef(js.Undefined)
}

type WriteCancelCallbackFunc func(this js.Ref) js.Ref

func (fn WriteCancelCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn WriteCancelCallbackFunc) DispatchCallback(
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

type WriteCancelCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *WriteCancelCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *WriteCancelCallback[T]) DispatchCallback(
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

type WriteImageCallbackFunc func(this js.Ref) js.Ref

func (fn WriteImageCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn WriteImageCallbackFunc) DispatchCallback(
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

type WriteImageCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *WriteImageCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *WriteImageCallback[T]) DispatchCallback(
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

// HasFuncCancelWrite returns true if the function "WEBEXT.imageWriterPrivate.cancelWrite" exists.
func HasFuncCancelWrite() bool {
	return js.True == bindings.HasFuncCancelWrite()
}

// FuncCancelWrite returns the function "WEBEXT.imageWriterPrivate.cancelWrite".
func FuncCancelWrite() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncCancelWrite(
		js.Pointer(&fn),
	)
	return
}

// CancelWrite calls the function "WEBEXT.imageWriterPrivate.cancelWrite" directly.
func CancelWrite() (ret js.Promise[js.Void]) {
	bindings.CallCancelWrite(
		js.Pointer(&ret),
	)

	return
}

// TryCancelWrite calls the function "WEBEXT.imageWriterPrivate.cancelWrite"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCancelWrite() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCancelWrite(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDestroyPartitions returns true if the function "WEBEXT.imageWriterPrivate.destroyPartitions" exists.
func HasFuncDestroyPartitions() bool {
	return js.True == bindings.HasFuncDestroyPartitions()
}

// FuncDestroyPartitions returns the function "WEBEXT.imageWriterPrivate.destroyPartitions".
func FuncDestroyPartitions() (fn js.Func[func(storageUnitId js.String) js.Promise[js.Void]]) {
	bindings.FuncDestroyPartitions(
		js.Pointer(&fn),
	)
	return
}

// DestroyPartitions calls the function "WEBEXT.imageWriterPrivate.destroyPartitions" directly.
func DestroyPartitions(storageUnitId js.String) (ret js.Promise[js.Void]) {
	bindings.CallDestroyPartitions(
		js.Pointer(&ret),
		storageUnitId.Ref(),
	)

	return
}

// TryDestroyPartitions calls the function "WEBEXT.imageWriterPrivate.destroyPartitions"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDestroyPartitions(storageUnitId js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDestroyPartitions(
		js.Pointer(&ret), js.Pointer(&exception),
		storageUnitId.Ref(),
	)

	return
}

// HasFuncListRemovableStorageDevices returns true if the function "WEBEXT.imageWriterPrivate.listRemovableStorageDevices" exists.
func HasFuncListRemovableStorageDevices() bool {
	return js.True == bindings.HasFuncListRemovableStorageDevices()
}

// FuncListRemovableStorageDevices returns the function "WEBEXT.imageWriterPrivate.listRemovableStorageDevices".
func FuncListRemovableStorageDevices() (fn js.Func[func() js.Promise[js.Array[RemovableStorageDevice]]]) {
	bindings.FuncListRemovableStorageDevices(
		js.Pointer(&fn),
	)
	return
}

// ListRemovableStorageDevices calls the function "WEBEXT.imageWriterPrivate.listRemovableStorageDevices" directly.
func ListRemovableStorageDevices() (ret js.Promise[js.Array[RemovableStorageDevice]]) {
	bindings.CallListRemovableStorageDevices(
		js.Pointer(&ret),
	)

	return
}

// TryListRemovableStorageDevices calls the function "WEBEXT.imageWriterPrivate.listRemovableStorageDevices"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryListRemovableStorageDevices() (ret js.Promise[js.Array[RemovableStorageDevice]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryListRemovableStorageDevices(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnDeviceInsertedEventCallbackFunc func(this js.Ref, device *RemovableStorageDevice) js.Ref

func (fn OnDeviceInsertedEventCallbackFunc) Register() js.Func[func(device *RemovableStorageDevice)] {
	return js.RegisterCallback[func(device *RemovableStorageDevice)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDeviceInsertedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RemovableStorageDevice
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

type OnDeviceInsertedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, device *RemovableStorageDevice) js.Ref
	Arg T
}

func (cb *OnDeviceInsertedEventCallback[T]) Register() js.Func[func(device *RemovableStorageDevice)] {
	return js.RegisterCallback[func(device *RemovableStorageDevice)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDeviceInsertedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RemovableStorageDevice
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

// HasFuncOnDeviceInserted returns true if the function "WEBEXT.imageWriterPrivate.onDeviceInserted.addListener" exists.
func HasFuncOnDeviceInserted() bool {
	return js.True == bindings.HasFuncOnDeviceInserted()
}

// FuncOnDeviceInserted returns the function "WEBEXT.imageWriterPrivate.onDeviceInserted.addListener".
func FuncOnDeviceInserted() (fn js.Func[func(callback js.Func[func(device *RemovableStorageDevice)])]) {
	bindings.FuncOnDeviceInserted(
		js.Pointer(&fn),
	)
	return
}

// OnDeviceInserted calls the function "WEBEXT.imageWriterPrivate.onDeviceInserted.addListener" directly.
func OnDeviceInserted(callback js.Func[func(device *RemovableStorageDevice)]) (ret js.Void) {
	bindings.CallOnDeviceInserted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDeviceInserted calls the function "WEBEXT.imageWriterPrivate.onDeviceInserted.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDeviceInserted(callback js.Func[func(device *RemovableStorageDevice)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDeviceInserted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDeviceInserted returns true if the function "WEBEXT.imageWriterPrivate.onDeviceInserted.removeListener" exists.
func HasFuncOffDeviceInserted() bool {
	return js.True == bindings.HasFuncOffDeviceInserted()
}

// FuncOffDeviceInserted returns the function "WEBEXT.imageWriterPrivate.onDeviceInserted.removeListener".
func FuncOffDeviceInserted() (fn js.Func[func(callback js.Func[func(device *RemovableStorageDevice)])]) {
	bindings.FuncOffDeviceInserted(
		js.Pointer(&fn),
	)
	return
}

// OffDeviceInserted calls the function "WEBEXT.imageWriterPrivate.onDeviceInserted.removeListener" directly.
func OffDeviceInserted(callback js.Func[func(device *RemovableStorageDevice)]) (ret js.Void) {
	bindings.CallOffDeviceInserted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDeviceInserted calls the function "WEBEXT.imageWriterPrivate.onDeviceInserted.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDeviceInserted(callback js.Func[func(device *RemovableStorageDevice)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDeviceInserted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDeviceInserted returns true if the function "WEBEXT.imageWriterPrivate.onDeviceInserted.hasListener" exists.
func HasFuncHasOnDeviceInserted() bool {
	return js.True == bindings.HasFuncHasOnDeviceInserted()
}

// FuncHasOnDeviceInserted returns the function "WEBEXT.imageWriterPrivate.onDeviceInserted.hasListener".
func FuncHasOnDeviceInserted() (fn js.Func[func(callback js.Func[func(device *RemovableStorageDevice)]) bool]) {
	bindings.FuncHasOnDeviceInserted(
		js.Pointer(&fn),
	)
	return
}

// HasOnDeviceInserted calls the function "WEBEXT.imageWriterPrivate.onDeviceInserted.hasListener" directly.
func HasOnDeviceInserted(callback js.Func[func(device *RemovableStorageDevice)]) (ret bool) {
	bindings.CallHasOnDeviceInserted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDeviceInserted calls the function "WEBEXT.imageWriterPrivate.onDeviceInserted.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDeviceInserted(callback js.Func[func(device *RemovableStorageDevice)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDeviceInserted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDeviceRemovedEventCallbackFunc func(this js.Ref, device *RemovableStorageDevice) js.Ref

func (fn OnDeviceRemovedEventCallbackFunc) Register() js.Func[func(device *RemovableStorageDevice)] {
	return js.RegisterCallback[func(device *RemovableStorageDevice)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDeviceRemovedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RemovableStorageDevice
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

type OnDeviceRemovedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, device *RemovableStorageDevice) js.Ref
	Arg T
}

func (cb *OnDeviceRemovedEventCallback[T]) Register() js.Func[func(device *RemovableStorageDevice)] {
	return js.RegisterCallback[func(device *RemovableStorageDevice)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDeviceRemovedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RemovableStorageDevice
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

// HasFuncOnDeviceRemoved returns true if the function "WEBEXT.imageWriterPrivate.onDeviceRemoved.addListener" exists.
func HasFuncOnDeviceRemoved() bool {
	return js.True == bindings.HasFuncOnDeviceRemoved()
}

// FuncOnDeviceRemoved returns the function "WEBEXT.imageWriterPrivate.onDeviceRemoved.addListener".
func FuncOnDeviceRemoved() (fn js.Func[func(callback js.Func[func(device *RemovableStorageDevice)])]) {
	bindings.FuncOnDeviceRemoved(
		js.Pointer(&fn),
	)
	return
}

// OnDeviceRemoved calls the function "WEBEXT.imageWriterPrivate.onDeviceRemoved.addListener" directly.
func OnDeviceRemoved(callback js.Func[func(device *RemovableStorageDevice)]) (ret js.Void) {
	bindings.CallOnDeviceRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDeviceRemoved calls the function "WEBEXT.imageWriterPrivate.onDeviceRemoved.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDeviceRemoved(callback js.Func[func(device *RemovableStorageDevice)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDeviceRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDeviceRemoved returns true if the function "WEBEXT.imageWriterPrivate.onDeviceRemoved.removeListener" exists.
func HasFuncOffDeviceRemoved() bool {
	return js.True == bindings.HasFuncOffDeviceRemoved()
}

// FuncOffDeviceRemoved returns the function "WEBEXT.imageWriterPrivate.onDeviceRemoved.removeListener".
func FuncOffDeviceRemoved() (fn js.Func[func(callback js.Func[func(device *RemovableStorageDevice)])]) {
	bindings.FuncOffDeviceRemoved(
		js.Pointer(&fn),
	)
	return
}

// OffDeviceRemoved calls the function "WEBEXT.imageWriterPrivate.onDeviceRemoved.removeListener" directly.
func OffDeviceRemoved(callback js.Func[func(device *RemovableStorageDevice)]) (ret js.Void) {
	bindings.CallOffDeviceRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDeviceRemoved calls the function "WEBEXT.imageWriterPrivate.onDeviceRemoved.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDeviceRemoved(callback js.Func[func(device *RemovableStorageDevice)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDeviceRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDeviceRemoved returns true if the function "WEBEXT.imageWriterPrivate.onDeviceRemoved.hasListener" exists.
func HasFuncHasOnDeviceRemoved() bool {
	return js.True == bindings.HasFuncHasOnDeviceRemoved()
}

// FuncHasOnDeviceRemoved returns the function "WEBEXT.imageWriterPrivate.onDeviceRemoved.hasListener".
func FuncHasOnDeviceRemoved() (fn js.Func[func(callback js.Func[func(device *RemovableStorageDevice)]) bool]) {
	bindings.FuncHasOnDeviceRemoved(
		js.Pointer(&fn),
	)
	return
}

// HasOnDeviceRemoved calls the function "WEBEXT.imageWriterPrivate.onDeviceRemoved.hasListener" directly.
func HasOnDeviceRemoved(callback js.Func[func(device *RemovableStorageDevice)]) (ret bool) {
	bindings.CallHasOnDeviceRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDeviceRemoved calls the function "WEBEXT.imageWriterPrivate.onDeviceRemoved.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDeviceRemoved(callback js.Func[func(device *RemovableStorageDevice)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDeviceRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnWriteCompleteEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnWriteCompleteEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnWriteCompleteEventCallbackFunc) DispatchCallback(
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

type OnWriteCompleteEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnWriteCompleteEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnWriteCompleteEventCallback[T]) DispatchCallback(
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

// HasFuncOnWriteComplete returns true if the function "WEBEXT.imageWriterPrivate.onWriteComplete.addListener" exists.
func HasFuncOnWriteComplete() bool {
	return js.True == bindings.HasFuncOnWriteComplete()
}

// FuncOnWriteComplete returns the function "WEBEXT.imageWriterPrivate.onWriteComplete.addListener".
func FuncOnWriteComplete() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnWriteComplete(
		js.Pointer(&fn),
	)
	return
}

// OnWriteComplete calls the function "WEBEXT.imageWriterPrivate.onWriteComplete.addListener" directly.
func OnWriteComplete(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnWriteComplete(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnWriteComplete calls the function "WEBEXT.imageWriterPrivate.onWriteComplete.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnWriteComplete(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnWriteComplete(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffWriteComplete returns true if the function "WEBEXT.imageWriterPrivate.onWriteComplete.removeListener" exists.
func HasFuncOffWriteComplete() bool {
	return js.True == bindings.HasFuncOffWriteComplete()
}

// FuncOffWriteComplete returns the function "WEBEXT.imageWriterPrivate.onWriteComplete.removeListener".
func FuncOffWriteComplete() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffWriteComplete(
		js.Pointer(&fn),
	)
	return
}

// OffWriteComplete calls the function "WEBEXT.imageWriterPrivate.onWriteComplete.removeListener" directly.
func OffWriteComplete(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffWriteComplete(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffWriteComplete calls the function "WEBEXT.imageWriterPrivate.onWriteComplete.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffWriteComplete(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffWriteComplete(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnWriteComplete returns true if the function "WEBEXT.imageWriterPrivate.onWriteComplete.hasListener" exists.
func HasFuncHasOnWriteComplete() bool {
	return js.True == bindings.HasFuncHasOnWriteComplete()
}

// FuncHasOnWriteComplete returns the function "WEBEXT.imageWriterPrivate.onWriteComplete.hasListener".
func FuncHasOnWriteComplete() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnWriteComplete(
		js.Pointer(&fn),
	)
	return
}

// HasOnWriteComplete calls the function "WEBEXT.imageWriterPrivate.onWriteComplete.hasListener" directly.
func HasOnWriteComplete(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnWriteComplete(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnWriteComplete calls the function "WEBEXT.imageWriterPrivate.onWriteComplete.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnWriteComplete(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnWriteComplete(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnWriteErrorEventCallbackFunc func(this js.Ref, info *ProgressInfo, err js.String) js.Ref

func (fn OnWriteErrorEventCallbackFunc) Register() js.Func[func(info *ProgressInfo, err js.String)] {
	return js.RegisterCallback[func(info *ProgressInfo, err js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnWriteErrorEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ProgressInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnWriteErrorEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *ProgressInfo, err js.String) js.Ref
	Arg T
}

func (cb *OnWriteErrorEventCallback[T]) Register() js.Func[func(info *ProgressInfo, err js.String)] {
	return js.RegisterCallback[func(info *ProgressInfo, err js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnWriteErrorEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ProgressInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnWriteError returns true if the function "WEBEXT.imageWriterPrivate.onWriteError.addListener" exists.
func HasFuncOnWriteError() bool {
	return js.True == bindings.HasFuncOnWriteError()
}

// FuncOnWriteError returns the function "WEBEXT.imageWriterPrivate.onWriteError.addListener".
func FuncOnWriteError() (fn js.Func[func(callback js.Func[func(info *ProgressInfo, err js.String)])]) {
	bindings.FuncOnWriteError(
		js.Pointer(&fn),
	)
	return
}

// OnWriteError calls the function "WEBEXT.imageWriterPrivate.onWriteError.addListener" directly.
func OnWriteError(callback js.Func[func(info *ProgressInfo, err js.String)]) (ret js.Void) {
	bindings.CallOnWriteError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnWriteError calls the function "WEBEXT.imageWriterPrivate.onWriteError.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnWriteError(callback js.Func[func(info *ProgressInfo, err js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnWriteError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffWriteError returns true if the function "WEBEXT.imageWriterPrivate.onWriteError.removeListener" exists.
func HasFuncOffWriteError() bool {
	return js.True == bindings.HasFuncOffWriteError()
}

// FuncOffWriteError returns the function "WEBEXT.imageWriterPrivate.onWriteError.removeListener".
func FuncOffWriteError() (fn js.Func[func(callback js.Func[func(info *ProgressInfo, err js.String)])]) {
	bindings.FuncOffWriteError(
		js.Pointer(&fn),
	)
	return
}

// OffWriteError calls the function "WEBEXT.imageWriterPrivate.onWriteError.removeListener" directly.
func OffWriteError(callback js.Func[func(info *ProgressInfo, err js.String)]) (ret js.Void) {
	bindings.CallOffWriteError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffWriteError calls the function "WEBEXT.imageWriterPrivate.onWriteError.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffWriteError(callback js.Func[func(info *ProgressInfo, err js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffWriteError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnWriteError returns true if the function "WEBEXT.imageWriterPrivate.onWriteError.hasListener" exists.
func HasFuncHasOnWriteError() bool {
	return js.True == bindings.HasFuncHasOnWriteError()
}

// FuncHasOnWriteError returns the function "WEBEXT.imageWriterPrivate.onWriteError.hasListener".
func FuncHasOnWriteError() (fn js.Func[func(callback js.Func[func(info *ProgressInfo, err js.String)]) bool]) {
	bindings.FuncHasOnWriteError(
		js.Pointer(&fn),
	)
	return
}

// HasOnWriteError calls the function "WEBEXT.imageWriterPrivate.onWriteError.hasListener" directly.
func HasOnWriteError(callback js.Func[func(info *ProgressInfo, err js.String)]) (ret bool) {
	bindings.CallHasOnWriteError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnWriteError calls the function "WEBEXT.imageWriterPrivate.onWriteError.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnWriteError(callback js.Func[func(info *ProgressInfo, err js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnWriteError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnWriteProgressEventCallbackFunc func(this js.Ref, info *ProgressInfo) js.Ref

func (fn OnWriteProgressEventCallbackFunc) Register() js.Func[func(info *ProgressInfo)] {
	return js.RegisterCallback[func(info *ProgressInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnWriteProgressEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ProgressInfo
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

type OnWriteProgressEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *ProgressInfo) js.Ref
	Arg T
}

func (cb *OnWriteProgressEventCallback[T]) Register() js.Func[func(info *ProgressInfo)] {
	return js.RegisterCallback[func(info *ProgressInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnWriteProgressEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ProgressInfo
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

// HasFuncOnWriteProgress returns true if the function "WEBEXT.imageWriterPrivate.onWriteProgress.addListener" exists.
func HasFuncOnWriteProgress() bool {
	return js.True == bindings.HasFuncOnWriteProgress()
}

// FuncOnWriteProgress returns the function "WEBEXT.imageWriterPrivate.onWriteProgress.addListener".
func FuncOnWriteProgress() (fn js.Func[func(callback js.Func[func(info *ProgressInfo)])]) {
	bindings.FuncOnWriteProgress(
		js.Pointer(&fn),
	)
	return
}

// OnWriteProgress calls the function "WEBEXT.imageWriterPrivate.onWriteProgress.addListener" directly.
func OnWriteProgress(callback js.Func[func(info *ProgressInfo)]) (ret js.Void) {
	bindings.CallOnWriteProgress(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnWriteProgress calls the function "WEBEXT.imageWriterPrivate.onWriteProgress.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnWriteProgress(callback js.Func[func(info *ProgressInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnWriteProgress(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffWriteProgress returns true if the function "WEBEXT.imageWriterPrivate.onWriteProgress.removeListener" exists.
func HasFuncOffWriteProgress() bool {
	return js.True == bindings.HasFuncOffWriteProgress()
}

// FuncOffWriteProgress returns the function "WEBEXT.imageWriterPrivate.onWriteProgress.removeListener".
func FuncOffWriteProgress() (fn js.Func[func(callback js.Func[func(info *ProgressInfo)])]) {
	bindings.FuncOffWriteProgress(
		js.Pointer(&fn),
	)
	return
}

// OffWriteProgress calls the function "WEBEXT.imageWriterPrivate.onWriteProgress.removeListener" directly.
func OffWriteProgress(callback js.Func[func(info *ProgressInfo)]) (ret js.Void) {
	bindings.CallOffWriteProgress(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffWriteProgress calls the function "WEBEXT.imageWriterPrivate.onWriteProgress.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffWriteProgress(callback js.Func[func(info *ProgressInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffWriteProgress(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnWriteProgress returns true if the function "WEBEXT.imageWriterPrivate.onWriteProgress.hasListener" exists.
func HasFuncHasOnWriteProgress() bool {
	return js.True == bindings.HasFuncHasOnWriteProgress()
}

// FuncHasOnWriteProgress returns the function "WEBEXT.imageWriterPrivate.onWriteProgress.hasListener".
func FuncHasOnWriteProgress() (fn js.Func[func(callback js.Func[func(info *ProgressInfo)]) bool]) {
	bindings.FuncHasOnWriteProgress(
		js.Pointer(&fn),
	)
	return
}

// HasOnWriteProgress calls the function "WEBEXT.imageWriterPrivate.onWriteProgress.hasListener" directly.
func HasOnWriteProgress(callback js.Func[func(info *ProgressInfo)]) (ret bool) {
	bindings.CallHasOnWriteProgress(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnWriteProgress calls the function "WEBEXT.imageWriterPrivate.onWriteProgress.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnWriteProgress(callback js.Func[func(info *ProgressInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnWriteProgress(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncWriteFromFile returns true if the function "WEBEXT.imageWriterPrivate.writeFromFile" exists.
func HasFuncWriteFromFile() bool {
	return js.True == bindings.HasFuncWriteFromFile()
}

// FuncWriteFromFile returns the function "WEBEXT.imageWriterPrivate.writeFromFile".
func FuncWriteFromFile() (fn js.Func[func(storageUnitId js.String, fileEntry js.Object) js.Promise[js.Void]]) {
	bindings.FuncWriteFromFile(
		js.Pointer(&fn),
	)
	return
}

// WriteFromFile calls the function "WEBEXT.imageWriterPrivate.writeFromFile" directly.
func WriteFromFile(storageUnitId js.String, fileEntry js.Object) (ret js.Promise[js.Void]) {
	bindings.CallWriteFromFile(
		js.Pointer(&ret),
		storageUnitId.Ref(),
		fileEntry.Ref(),
	)

	return
}

// TryWriteFromFile calls the function "WEBEXT.imageWriterPrivate.writeFromFile"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryWriteFromFile(storageUnitId js.String, fileEntry js.Object) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWriteFromFile(
		js.Pointer(&ret), js.Pointer(&exception),
		storageUnitId.Ref(),
		fileEntry.Ref(),
	)

	return
}

// HasFuncWriteFromUrl returns true if the function "WEBEXT.imageWriterPrivate.writeFromUrl" exists.
func HasFuncWriteFromUrl() bool {
	return js.True == bindings.HasFuncWriteFromUrl()
}

// FuncWriteFromUrl returns the function "WEBEXT.imageWriterPrivate.writeFromUrl".
func FuncWriteFromUrl() (fn js.Func[func(storageUnitId js.String, imageUrl js.String, options UrlWriteOptions) js.Promise[js.Void]]) {
	bindings.FuncWriteFromUrl(
		js.Pointer(&fn),
	)
	return
}

// WriteFromUrl calls the function "WEBEXT.imageWriterPrivate.writeFromUrl" directly.
func WriteFromUrl(storageUnitId js.String, imageUrl js.String, options UrlWriteOptions) (ret js.Promise[js.Void]) {
	bindings.CallWriteFromUrl(
		js.Pointer(&ret),
		storageUnitId.Ref(),
		imageUrl.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryWriteFromUrl calls the function "WEBEXT.imageWriterPrivate.writeFromUrl"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryWriteFromUrl(storageUnitId js.String, imageUrl js.String, options UrlWriteOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWriteFromUrl(
		js.Pointer(&ret), js.Pointer(&exception),
		storageUnitId.Ref(),
		imageUrl.Ref(),
		js.Pointer(&options),
	)

	return
}
