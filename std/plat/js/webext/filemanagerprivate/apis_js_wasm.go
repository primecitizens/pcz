// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package filemanagerprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/filemanagerprivate/bindings"
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

type AddMountCallbackFunc func(this js.Ref, sourcePath js.String) js.Ref

func (fn AddMountCallbackFunc) Register() js.Func[func(sourcePath js.String)] {
	return js.RegisterCallback[func(sourcePath js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn AddMountCallbackFunc) DispatchCallback(
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

type AddMountCallback[T any] struct {
	Fn  func(arg T, this js.Ref, sourcePath js.String) js.Ref
	Arg T
}

func (cb *AddMountCallback[T]) Register() js.Func[func(sourcePath js.String)] {
	return js.RegisterCallback[func(sourcePath js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *AddMountCallback[T]) DispatchCallback(
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

type IconSet struct {
	// Icon16x16Url is "IconSet.icon16x16Url"
	//
	// Optional
	Icon16x16Url js.String
	// Icon32x32Url is "IconSet.icon32x32Url"
	//
	// Optional
	Icon32x32Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IconSet with all fields set.
func (p IconSet) FromRef(ref js.Ref) IconSet {
	p.UpdateFrom(ref)
	return p
}

// New creates a new IconSet in the application heap.
func (p IconSet) New() js.Ref {
	return bindings.IconSetJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *IconSet) UpdateFrom(ref js.Ref) {
	bindings.IconSetJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IconSet) Update(ref js.Ref) {
	bindings.IconSetJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IconSet) FreeMembers(recursive bool) {
	js.Free(
		p.Icon16x16Url.Ref(),
		p.Icon32x32Url.Ref(),
	)
	p.Icon16x16Url = p.Icon16x16Url.FromRef(js.Undefined)
	p.Icon32x32Url = p.Icon32x32Url.FromRef(js.Undefined)
}

type AndroidApp struct {
	// Name is "AndroidApp.name"
	//
	// Optional
	Name js.String
	// PackageName is "AndroidApp.packageName"
	//
	// Optional
	PackageName js.String
	// ActivityName is "AndroidApp.activityName"
	//
	// Optional
	ActivityName js.String
	// IconSet is "AndroidApp.iconSet"
	//
	// Optional
	//
	// NOTE: IconSet.FFI_USE MUST be set to true to get IconSet used.
	IconSet IconSet

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AndroidApp with all fields set.
func (p AndroidApp) FromRef(ref js.Ref) AndroidApp {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AndroidApp in the application heap.
func (p AndroidApp) New() js.Ref {
	return bindings.AndroidAppJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AndroidApp) UpdateFrom(ref js.Ref) {
	bindings.AndroidAppJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AndroidApp) Update(ref js.Ref) {
	bindings.AndroidAppJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AndroidApp) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.PackageName.Ref(),
		p.ActivityName.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.PackageName = p.PackageName.FromRef(js.Undefined)
	p.ActivityName = p.ActivityName.FromRef(js.Undefined)
	if recursive {
		p.IconSet.FreeMembers(true)
	}
}

type AttachedImages struct {
	// Data is "AttachedImages.data"
	//
	// Optional
	Data js.String
	// Type is "AttachedImages.type"
	//
	// Optional
	Type js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AttachedImages with all fields set.
func (p AttachedImages) FromRef(ref js.Ref) AttachedImages {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AttachedImages in the application heap.
func (p AttachedImages) New() js.Ref {
	return bindings.AttachedImagesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AttachedImages) UpdateFrom(ref js.Ref) {
	bindings.AttachedImagesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AttachedImages) Update(ref js.Ref) {
	bindings.AttachedImagesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AttachedImages) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
		p.Type.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
	p.Type = p.Type.FromRef(js.Undefined)
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

type BulkPinStage uint32

const (
	_ BulkPinStage = iota

	BulkPinStage_STOPPED
	BulkPinStage_PAUSED_OFFLINE
	BulkPinStage_PAUSED_BATTERY_SAVER
	BulkPinStage_GETTING_FREE_SPACE
	BulkPinStage_LISTING_FILES
	BulkPinStage_SYNCING
	BulkPinStage_SUCCESS
	BulkPinStage_NOT_ENOUGH_SPACE
	BulkPinStage_CANNOT_GET_FREE_SPACE
	BulkPinStage_CANNOT_LIST_FILES
	BulkPinStage_CANNOT_ENABLE_DOCS_OFFLINE
)

func (BulkPinStage) FromRef(str js.Ref) BulkPinStage {
	return BulkPinStage(bindings.ConstOfBulkPinStage(str))
}

func (x BulkPinStage) String() (string, bool) {
	switch x {
	case BulkPinStage_STOPPED:
		return "stopped", true
	case BulkPinStage_PAUSED_OFFLINE:
		return "paused_offline", true
	case BulkPinStage_PAUSED_BATTERY_SAVER:
		return "paused_battery_saver", true
	case BulkPinStage_GETTING_FREE_SPACE:
		return "getting_free_space", true
	case BulkPinStage_LISTING_FILES:
		return "listing_files", true
	case BulkPinStage_SYNCING:
		return "syncing", true
	case BulkPinStage_SUCCESS:
		return "success", true
	case BulkPinStage_NOT_ENOUGH_SPACE:
		return "not_enough_space", true
	case BulkPinStage_CANNOT_GET_FREE_SPACE:
		return "cannot_get_free_space", true
	case BulkPinStage_CANNOT_LIST_FILES:
		return "cannot_list_files", true
	case BulkPinStage_CANNOT_ENABLE_DOCS_OFFLINE:
		return "cannot_enable_docs_offline", true
	default:
		return "", false
	}
}

type BulkPinProgress struct {
	// Stage is "BulkPinProgress.stage"
	//
	// Optional
	Stage BulkPinStage
	// FreeSpaceBytes is "BulkPinProgress.freeSpaceBytes"
	//
	// Optional
	//
	// NOTE: FFI_USE_FreeSpaceBytes MUST be set to true to make this field effective.
	FreeSpaceBytes float64
	// RequiredSpaceBytes is "BulkPinProgress.requiredSpaceBytes"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequiredSpaceBytes MUST be set to true to make this field effective.
	RequiredSpaceBytes float64
	// BytesToPin is "BulkPinProgress.bytesToPin"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesToPin MUST be set to true to make this field effective.
	BytesToPin float64
	// PinnedBytes is "BulkPinProgress.pinnedBytes"
	//
	// Optional
	//
	// NOTE: FFI_USE_PinnedBytes MUST be set to true to make this field effective.
	PinnedBytes float64
	// FilesToPin is "BulkPinProgress.filesToPin"
	//
	// Optional
	//
	// NOTE: FFI_USE_FilesToPin MUST be set to true to make this field effective.
	FilesToPin int32
	// ListedFiles is "BulkPinProgress.listedFiles"
	//
	// Optional
	//
	// NOTE: FFI_USE_ListedFiles MUST be set to true to make this field effective.
	ListedFiles int32
	// RemainingSeconds is "BulkPinProgress.remainingSeconds"
	//
	// Optional
	//
	// NOTE: FFI_USE_RemainingSeconds MUST be set to true to make this field effective.
	RemainingSeconds float64
	// EmptiedQueue is "BulkPinProgress.emptiedQueue"
	//
	// Optional
	//
	// NOTE: FFI_USE_EmptiedQueue MUST be set to true to make this field effective.
	EmptiedQueue bool

	FFI_USE_FreeSpaceBytes     bool // for FreeSpaceBytes.
	FFI_USE_RequiredSpaceBytes bool // for RequiredSpaceBytes.
	FFI_USE_BytesToPin         bool // for BytesToPin.
	FFI_USE_PinnedBytes        bool // for PinnedBytes.
	FFI_USE_FilesToPin         bool // for FilesToPin.
	FFI_USE_ListedFiles        bool // for ListedFiles.
	FFI_USE_RemainingSeconds   bool // for RemainingSeconds.
	FFI_USE_EmptiedQueue       bool // for EmptiedQueue.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BulkPinProgress with all fields set.
func (p BulkPinProgress) FromRef(ref js.Ref) BulkPinProgress {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BulkPinProgress in the application heap.
func (p BulkPinProgress) New() js.Ref {
	return bindings.BulkPinProgressJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *BulkPinProgress) UpdateFrom(ref js.Ref) {
	bindings.BulkPinProgressJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BulkPinProgress) Update(ref js.Ref) {
	bindings.BulkPinProgressJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BulkPinProgress) FreeMembers(recursive bool) {
}

type ChangeType uint32

const (
	_ ChangeType = iota

	ChangeType_ADD_OR_UPDATE
	ChangeType_DELETE
)

func (ChangeType) FromRef(str js.Ref) ChangeType {
	return ChangeType(bindings.ConstOfChangeType(str))
}

func (x ChangeType) String() (string, bool) {
	switch x {
	case ChangeType_ADD_OR_UPDATE:
		return "add_or_update", true
	case ChangeType_DELETE:
		return "delete", true
	default:
		return "", false
	}
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

type ConflictPauseParams struct {
	// ConflictName is "ConflictPauseParams.conflictName"
	//
	// Optional
	ConflictName js.String
	// ConflictIsDirectory is "ConflictPauseParams.conflictIsDirectory"
	//
	// Optional
	//
	// NOTE: FFI_USE_ConflictIsDirectory MUST be set to true to make this field effective.
	ConflictIsDirectory bool
	// ConflictMultiple is "ConflictPauseParams.conflictMultiple"
	//
	// Optional
	//
	// NOTE: FFI_USE_ConflictMultiple MUST be set to true to make this field effective.
	ConflictMultiple bool
	// ConflictTargetUrl is "ConflictPauseParams.conflictTargetUrl"
	//
	// Optional
	ConflictTargetUrl js.String

	FFI_USE_ConflictIsDirectory bool // for ConflictIsDirectory.
	FFI_USE_ConflictMultiple    bool // for ConflictMultiple.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ConflictPauseParams with all fields set.
func (p ConflictPauseParams) FromRef(ref js.Ref) ConflictPauseParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ConflictPauseParams in the application heap.
func (p ConflictPauseParams) New() js.Ref {
	return bindings.ConflictPauseParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ConflictPauseParams) UpdateFrom(ref js.Ref) {
	bindings.ConflictPauseParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ConflictPauseParams) Update(ref js.Ref) {
	bindings.ConflictPauseParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ConflictPauseParams) FreeMembers(recursive bool) {
	js.Free(
		p.ConflictName.Ref(),
		p.ConflictTargetUrl.Ref(),
	)
	p.ConflictName = p.ConflictName.FromRef(js.Undefined)
	p.ConflictTargetUrl = p.ConflictTargetUrl.FromRef(js.Undefined)
}

type ConflictResumeParams struct {
	// ConflictResolve is "ConflictResumeParams.conflictResolve"
	//
	// Optional
	ConflictResolve js.String
	// ConflictApplyToAll is "ConflictResumeParams.conflictApplyToAll"
	//
	// Optional
	//
	// NOTE: FFI_USE_ConflictApplyToAll MUST be set to true to make this field effective.
	ConflictApplyToAll bool

	FFI_USE_ConflictApplyToAll bool // for ConflictApplyToAll.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ConflictResumeParams with all fields set.
func (p ConflictResumeParams) FromRef(ref js.Ref) ConflictResumeParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ConflictResumeParams in the application heap.
func (p ConflictResumeParams) New() js.Ref {
	return bindings.ConflictResumeParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ConflictResumeParams) UpdateFrom(ref js.Ref) {
	bindings.ConflictResumeParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ConflictResumeParams) Update(ref js.Ref) {
	bindings.ConflictResumeParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ConflictResumeParams) FreeMembers(recursive bool) {
	js.Free(
		p.ConflictResolve.Ref(),
	)
	p.ConflictResolve = p.ConflictResolve.FromRef(js.Undefined)
}

type CrostiniEventType uint32

const (
	_ CrostiniEventType = iota

	CrostiniEventType_ENABLE
	CrostiniEventType_DISABLE
	CrostiniEventType_SHARE
	CrostiniEventType_UNSHARE
	CrostiniEventType_DROP_FAILED_PLUGIN_VM_DIRECTORY_NOT_SHARED
)

func (CrostiniEventType) FromRef(str js.Ref) CrostiniEventType {
	return CrostiniEventType(bindings.ConstOfCrostiniEventType(str))
}

func (x CrostiniEventType) String() (string, bool) {
	switch x {
	case CrostiniEventType_ENABLE:
		return "enable", true
	case CrostiniEventType_DISABLE:
		return "disable", true
	case CrostiniEventType_SHARE:
		return "share", true
	case CrostiniEventType_UNSHARE:
		return "unshare", true
	case CrostiniEventType_DROP_FAILED_PLUGIN_VM_DIRECTORY_NOT_SHARED:
		return "drop_failed_plugin_vm_directory_not_shared", true
	default:
		return "", false
	}
}

type CrostiniEvent struct {
	// EventType is "CrostiniEvent.eventType"
	//
	// Optional
	EventType CrostiniEventType
	// VmName is "CrostiniEvent.vmName"
	//
	// Optional
	VmName js.String
	// ContainerName is "CrostiniEvent.containerName"
	//
	// Optional
	ContainerName js.String
	// Entries is "CrostiniEvent.entries"
	//
	// Optional
	Entries js.Array[js.Object]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CrostiniEvent with all fields set.
func (p CrostiniEvent) FromRef(ref js.Ref) CrostiniEvent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CrostiniEvent in the application heap.
func (p CrostiniEvent) New() js.Ref {
	return bindings.CrostiniEventJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CrostiniEvent) UpdateFrom(ref js.Ref) {
	bindings.CrostiniEventJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CrostiniEvent) Update(ref js.Ref) {
	bindings.CrostiniEventJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CrostiniEvent) FreeMembers(recursive bool) {
	js.Free(
		p.VmName.Ref(),
		p.ContainerName.Ref(),
		p.Entries.Ref(),
	)
	p.VmName = p.VmName.FromRef(js.Undefined)
	p.ContainerName = p.ContainerName.FromRef(js.Undefined)
	p.Entries = p.Entries.FromRef(js.Undefined)
}

type DeviceConnectionState uint32

const (
	_ DeviceConnectionState = iota

	DeviceConnectionState_OFFLINE
	DeviceConnectionState_ONLINE
)

func (DeviceConnectionState) FromRef(str js.Ref) DeviceConnectionState {
	return DeviceConnectionState(bindings.ConstOfDeviceConnectionState(str))
}

func (x DeviceConnectionState) String() (string, bool) {
	switch x {
	case DeviceConnectionState_OFFLINE:
		return "OFFLINE", true
	case DeviceConnectionState_ONLINE:
		return "ONLINE", true
	default:
		return "", false
	}
}

type DeviceEventType uint32

const (
	_ DeviceEventType = iota

	DeviceEventType_DISABLED
	DeviceEventType_REMOVED
	DeviceEventType_HARD_UNPLUGGED
	DeviceEventType_FORMAT_START
	DeviceEventType_FORMAT_SUCCESS
	DeviceEventType_FORMAT_FAIL
	DeviceEventType_RENAME_START
	DeviceEventType_RENAME_SUCCESS
	DeviceEventType_RENAME_FAIL
	DeviceEventType_PARTITION_START
	DeviceEventType_PARTITION_SUCCESS
	DeviceEventType_PARTITION_FAIL
)

func (DeviceEventType) FromRef(str js.Ref) DeviceEventType {
	return DeviceEventType(bindings.ConstOfDeviceEventType(str))
}

func (x DeviceEventType) String() (string, bool) {
	switch x {
	case DeviceEventType_DISABLED:
		return "disabled", true
	case DeviceEventType_REMOVED:
		return "removed", true
	case DeviceEventType_HARD_UNPLUGGED:
		return "hard_unplugged", true
	case DeviceEventType_FORMAT_START:
		return "format_start", true
	case DeviceEventType_FORMAT_SUCCESS:
		return "format_success", true
	case DeviceEventType_FORMAT_FAIL:
		return "format_fail", true
	case DeviceEventType_RENAME_START:
		return "rename_start", true
	case DeviceEventType_RENAME_SUCCESS:
		return "rename_success", true
	case DeviceEventType_RENAME_FAIL:
		return "rename_fail", true
	case DeviceEventType_PARTITION_START:
		return "partition_start", true
	case DeviceEventType_PARTITION_SUCCESS:
		return "partition_success", true
	case DeviceEventType_PARTITION_FAIL:
		return "partition_fail", true
	default:
		return "", false
	}
}

type DeviceEvent struct {
	// Type is "DeviceEvent.type"
	//
	// Optional
	Type DeviceEventType
	// DevicePath is "DeviceEvent.devicePath"
	//
	// Optional
	DevicePath js.String
	// DeviceLabel is "DeviceEvent.deviceLabel"
	//
	// Optional
	DeviceLabel js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DeviceEvent with all fields set.
func (p DeviceEvent) FromRef(ref js.Ref) DeviceEvent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DeviceEvent in the application heap.
func (p DeviceEvent) New() js.Ref {
	return bindings.DeviceEventJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DeviceEvent) UpdateFrom(ref js.Ref) {
	bindings.DeviceEventJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DeviceEvent) Update(ref js.Ref) {
	bindings.DeviceEventJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DeviceEvent) FreeMembers(recursive bool) {
	js.Free(
		p.DevicePath.Ref(),
		p.DeviceLabel.Ref(),
	)
	p.DevicePath = p.DevicePath.FromRef(js.Undefined)
	p.DeviceLabel = p.DeviceLabel.FromRef(js.Undefined)
}

type DeviceType uint32

const (
	_ DeviceType = iota

	DeviceType_USB
	DeviceType_SD
	DeviceType_OPTICAL
	DeviceType_MOBILE
	DeviceType_UNKNOWN
)

func (DeviceType) FromRef(str js.Ref) DeviceType {
	return DeviceType(bindings.ConstOfDeviceType(str))
}

func (x DeviceType) String() (string, bool) {
	switch x {
	case DeviceType_USB:
		return "usb", true
	case DeviceType_SD:
		return "sd", true
	case DeviceType_OPTICAL:
		return "optical", true
	case DeviceType_MOBILE:
		return "mobile", true
	case DeviceType_UNKNOWN:
		return "unknown", true
	default:
		return "", false
	}
}

type VolumeType uint32

const (
	_ VolumeType = iota

	VolumeType_DRIVE
	VolumeType_DOWNLOADS
	VolumeType_REMOVABLE
	VolumeType_ARCHIVE
	VolumeType_PROVIDED
	VolumeType_MTP
	VolumeType_MEDIA_VIEW
	VolumeType_CROSTINI
	VolumeType_ANDROID_FILES
	VolumeType_DOCUMENTS_PROVIDER
	VolumeType_TESTING
	VolumeType_SMB
	VolumeType_SYSTEM_INTERNAL
	VolumeType_GUEST_OS
)

func (VolumeType) FromRef(str js.Ref) VolumeType {
	return VolumeType(bindings.ConstOfVolumeType(str))
}

func (x VolumeType) String() (string, bool) {
	switch x {
	case VolumeType_DRIVE:
		return "drive", true
	case VolumeType_DOWNLOADS:
		return "downloads", true
	case VolumeType_REMOVABLE:
		return "removable", true
	case VolumeType_ARCHIVE:
		return "archive", true
	case VolumeType_PROVIDED:
		return "provided", true
	case VolumeType_MTP:
		return "mtp", true
	case VolumeType_MEDIA_VIEW:
		return "media_view", true
	case VolumeType_CROSTINI:
		return "crostini", true
	case VolumeType_ANDROID_FILES:
		return "android_files", true
	case VolumeType_DOCUMENTS_PROVIDER:
		return "documents_provider", true
	case VolumeType_TESTING:
		return "testing", true
	case VolumeType_SMB:
		return "smb", true
	case VolumeType_SYSTEM_INTERNAL:
		return "system_internal", true
	case VolumeType_GUEST_OS:
		return "guest_os", true
	default:
		return "", false
	}
}

type DialogCallerInformation struct {
	// Url is "DialogCallerInformation.url"
	//
	// Optional
	Url js.String
	// Component is "DialogCallerInformation.component"
	//
	// Optional
	Component VolumeType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DialogCallerInformation with all fields set.
func (p DialogCallerInformation) FromRef(ref js.Ref) DialogCallerInformation {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DialogCallerInformation in the application heap.
func (p DialogCallerInformation) New() js.Ref {
	return bindings.DialogCallerInformationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DialogCallerInformation) UpdateFrom(ref js.Ref) {
	bindings.DialogCallerInformationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DialogCallerInformation) Update(ref js.Ref) {
	bindings.DialogCallerInformationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DialogCallerInformation) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
}

type DlpLevel uint32

const (
	_ DlpLevel = iota

	DlpLevel_REPORT
	DlpLevel_WARN
	DlpLevel_BLOCK
	DlpLevel_ALLOW
)

func (DlpLevel) FromRef(str js.Ref) DlpLevel {
	return DlpLevel(bindings.ConstOfDlpLevel(str))
}

func (x DlpLevel) String() (string, bool) {
	switch x {
	case DlpLevel_REPORT:
		return "report", true
	case DlpLevel_WARN:
		return "warn", true
	case DlpLevel_BLOCK:
		return "block", true
	case DlpLevel_ALLOW:
		return "allow", true
	default:
		return "", false
	}
}

type DlpMetadata struct {
	// SourceUrl is "DlpMetadata.sourceUrl"
	//
	// Optional
	SourceUrl js.String
	// IsDlpRestricted is "DlpMetadata.isDlpRestricted"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsDlpRestricted MUST be set to true to make this field effective.
	IsDlpRestricted bool
	// IsRestrictedForDestination is "DlpMetadata.isRestrictedForDestination"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsRestrictedForDestination MUST be set to true to make this field effective.
	IsRestrictedForDestination bool

	FFI_USE_IsDlpRestricted            bool // for IsDlpRestricted.
	FFI_USE_IsRestrictedForDestination bool // for IsRestrictedForDestination.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DlpMetadata with all fields set.
func (p DlpMetadata) FromRef(ref js.Ref) DlpMetadata {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DlpMetadata in the application heap.
func (p DlpMetadata) New() js.Ref {
	return bindings.DlpMetadataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DlpMetadata) UpdateFrom(ref js.Ref) {
	bindings.DlpMetadataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DlpMetadata) Update(ref js.Ref) {
	bindings.DlpMetadataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DlpMetadata) FreeMembers(recursive bool) {
	js.Free(
		p.SourceUrl.Ref(),
	)
	p.SourceUrl = p.SourceUrl.FromRef(js.Undefined)
}

type DlpRestrictionDetails struct {
	// Level is "DlpRestrictionDetails.level"
	//
	// Optional
	Level DlpLevel
	// Urls is "DlpRestrictionDetails.urls"
	//
	// Optional
	Urls js.Array[js.String]
	// Components is "DlpRestrictionDetails.components"
	//
	// Optional
	Components js.Array[VolumeType]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DlpRestrictionDetails with all fields set.
func (p DlpRestrictionDetails) FromRef(ref js.Ref) DlpRestrictionDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DlpRestrictionDetails in the application heap.
func (p DlpRestrictionDetails) New() js.Ref {
	return bindings.DlpRestrictionDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DlpRestrictionDetails) UpdateFrom(ref js.Ref) {
	bindings.DlpRestrictionDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DlpRestrictionDetails) Update(ref js.Ref) {
	bindings.DlpRestrictionDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DlpRestrictionDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Urls.Ref(),
		p.Components.Ref(),
	)
	p.Urls = p.Urls.FromRef(js.Undefined)
	p.Components = p.Components.FromRef(js.Undefined)
}

type DriveConfirmDialogType uint32

const (
	_ DriveConfirmDialogType = iota

	DriveConfirmDialogType_ENABLE_DOCS_OFFLINE
)

func (DriveConfirmDialogType) FromRef(str js.Ref) DriveConfirmDialogType {
	return DriveConfirmDialogType(bindings.ConstOfDriveConfirmDialogType(str))
}

func (x DriveConfirmDialogType) String() (string, bool) {
	switch x {
	case DriveConfirmDialogType_ENABLE_DOCS_OFFLINE:
		return "enable_docs_offline", true
	default:
		return "", false
	}
}

type DriveConfirmDialogEvent struct {
	// Type is "DriveConfirmDialogEvent.type"
	//
	// Optional
	Type DriveConfirmDialogType
	// FileUrl is "DriveConfirmDialogEvent.fileUrl"
	//
	// Optional
	FileUrl js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DriveConfirmDialogEvent with all fields set.
func (p DriveConfirmDialogEvent) FromRef(ref js.Ref) DriveConfirmDialogEvent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DriveConfirmDialogEvent in the application heap.
func (p DriveConfirmDialogEvent) New() js.Ref {
	return bindings.DriveConfirmDialogEventJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DriveConfirmDialogEvent) UpdateFrom(ref js.Ref) {
	bindings.DriveConfirmDialogEventJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DriveConfirmDialogEvent) Update(ref js.Ref) {
	bindings.DriveConfirmDialogEventJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DriveConfirmDialogEvent) FreeMembers(recursive bool) {
	js.Free(
		p.FileUrl.Ref(),
	)
	p.FileUrl = p.FileUrl.FromRef(js.Undefined)
}

type DriveConnectionStateType uint32

const (
	_ DriveConnectionStateType = iota

	DriveConnectionStateType_OFFLINE
	DriveConnectionStateType_METERED
	DriveConnectionStateType_ONLINE
)

func (DriveConnectionStateType) FromRef(str js.Ref) DriveConnectionStateType {
	return DriveConnectionStateType(bindings.ConstOfDriveConnectionStateType(str))
}

func (x DriveConnectionStateType) String() (string, bool) {
	switch x {
	case DriveConnectionStateType_OFFLINE:
		return "OFFLINE", true
	case DriveConnectionStateType_METERED:
		return "METERED", true
	case DriveConnectionStateType_ONLINE:
		return "ONLINE", true
	default:
		return "", false
	}
}

type DriveOfflineReason uint32

const (
	_ DriveOfflineReason = iota

	DriveOfflineReason_NOT_READY
	DriveOfflineReason_NO_NETWORK
	DriveOfflineReason_NO_SERVICE
)

func (DriveOfflineReason) FromRef(str js.Ref) DriveOfflineReason {
	return DriveOfflineReason(bindings.ConstOfDriveOfflineReason(str))
}

func (x DriveOfflineReason) String() (string, bool) {
	switch x {
	case DriveOfflineReason_NOT_READY:
		return "NOT_READY", true
	case DriveOfflineReason_NO_NETWORK:
		return "NO_NETWORK", true
	case DriveOfflineReason_NO_SERVICE:
		return "NO_SERVICE", true
	default:
		return "", false
	}
}

type DriveConnectionState struct {
	// Type is "DriveConnectionState.type"
	//
	// Optional
	Type DriveConnectionStateType
	// Reason is "DriveConnectionState.reason"
	//
	// Optional
	Reason DriveOfflineReason

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DriveConnectionState with all fields set.
func (p DriveConnectionState) FromRef(ref js.Ref) DriveConnectionState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DriveConnectionState in the application heap.
func (p DriveConnectionState) New() js.Ref {
	return bindings.DriveConnectionStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DriveConnectionState) UpdateFrom(ref js.Ref) {
	bindings.DriveConnectionStateJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DriveConnectionState) Update(ref js.Ref) {
	bindings.DriveConnectionStateJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DriveConnectionState) FreeMembers(recursive bool) {
}

type DriveDialogResult uint32

const (
	_ DriveDialogResult = iota

	DriveDialogResult_NOT_DISPLAYED
	DriveDialogResult_ACCEPT
	DriveDialogResult_REJECT
	DriveDialogResult_DISMISS
)

func (DriveDialogResult) FromRef(str js.Ref) DriveDialogResult {
	return DriveDialogResult(bindings.ConstOfDriveDialogResult(str))
}

func (x DriveDialogResult) String() (string, bool) {
	switch x {
	case DriveDialogResult_NOT_DISPLAYED:
		return "not_displayed", true
	case DriveDialogResult_ACCEPT:
		return "accept", true
	case DriveDialogResult_REJECT:
		return "reject", true
	case DriveDialogResult_DISMISS:
		return "dismiss", true
	default:
		return "", false
	}
}

type DriveMetadataSearchResult struct {
	// Entry is "DriveMetadataSearchResult.entry"
	//
	// Optional
	Entry js.Object
	// HighlightedBaseName is "DriveMetadataSearchResult.highlightedBaseName"
	//
	// Optional
	HighlightedBaseName js.String
	// AvailableOffline is "DriveMetadataSearchResult.availableOffline"
	//
	// Optional
	//
	// NOTE: FFI_USE_AvailableOffline MUST be set to true to make this field effective.
	AvailableOffline bool

	FFI_USE_AvailableOffline bool // for AvailableOffline.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DriveMetadataSearchResult with all fields set.
func (p DriveMetadataSearchResult) FromRef(ref js.Ref) DriveMetadataSearchResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DriveMetadataSearchResult in the application heap.
func (p DriveMetadataSearchResult) New() js.Ref {
	return bindings.DriveMetadataSearchResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DriveMetadataSearchResult) UpdateFrom(ref js.Ref) {
	bindings.DriveMetadataSearchResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DriveMetadataSearchResult) Update(ref js.Ref) {
	bindings.DriveMetadataSearchResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DriveMetadataSearchResult) FreeMembers(recursive bool) {
	js.Free(
		p.Entry.Ref(),
		p.HighlightedBaseName.Ref(),
	)
	p.Entry = p.Entry.FromRef(js.Undefined)
	p.HighlightedBaseName = p.HighlightedBaseName.FromRef(js.Undefined)
}

type UserType uint32

const (
	_ UserType = iota

	UserType_UNMANAGED
	UserType_ORGANIZATION
)

func (UserType) FromRef(str js.Ref) UserType {
	return UserType(bindings.ConstOfUserType(str))
}

func (x UserType) String() (string, bool) {
	switch x {
	case UserType_UNMANAGED:
		return "unmanaged", true
	case UserType_ORGANIZATION:
		return "organization", true
	default:
		return "", false
	}
}

type DriveQuotaMetadata struct {
	// UserType is "DriveQuotaMetadata.userType"
	//
	// Optional
	UserType UserType
	// UsedBytes is "DriveQuotaMetadata.usedBytes"
	//
	// Optional
	//
	// NOTE: FFI_USE_UsedBytes MUST be set to true to make this field effective.
	UsedBytes float64
	// TotalBytes is "DriveQuotaMetadata.totalBytes"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalBytes MUST be set to true to make this field effective.
	TotalBytes float64
	// OrganizationLimitExceeded is "DriveQuotaMetadata.organizationLimitExceeded"
	//
	// Optional
	//
	// NOTE: FFI_USE_OrganizationLimitExceeded MUST be set to true to make this field effective.
	OrganizationLimitExceeded bool
	// OrganizationName is "DriveQuotaMetadata.organizationName"
	//
	// Optional
	OrganizationName js.String

	FFI_USE_UsedBytes                 bool // for UsedBytes.
	FFI_USE_TotalBytes                bool // for TotalBytes.
	FFI_USE_OrganizationLimitExceeded bool // for OrganizationLimitExceeded.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DriveQuotaMetadata with all fields set.
func (p DriveQuotaMetadata) FromRef(ref js.Ref) DriveQuotaMetadata {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DriveQuotaMetadata in the application heap.
func (p DriveQuotaMetadata) New() js.Ref {
	return bindings.DriveQuotaMetadataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DriveQuotaMetadata) UpdateFrom(ref js.Ref) {
	bindings.DriveQuotaMetadataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DriveQuotaMetadata) Update(ref js.Ref) {
	bindings.DriveQuotaMetadataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DriveQuotaMetadata) FreeMembers(recursive bool) {
	js.Free(
		p.OrganizationName.Ref(),
	)
	p.OrganizationName = p.OrganizationName.FromRef(js.Undefined)
}

type DriveShareType uint32

const (
	_ DriveShareType = iota

	DriveShareType_CAN_EDIT
	DriveShareType_CAN_COMMENT
	DriveShareType_CAN_VIEW
)

func (DriveShareType) FromRef(str js.Ref) DriveShareType {
	return DriveShareType(bindings.ConstOfDriveShareType(str))
}

func (x DriveShareType) String() (string, bool) {
	switch x {
	case DriveShareType_CAN_EDIT:
		return "can_edit", true
	case DriveShareType_CAN_COMMENT:
		return "can_comment", true
	case DriveShareType_CAN_VIEW:
		return "can_view", true
	default:
		return "", false
	}
}

type DriveSyncErrorType uint32

const (
	_ DriveSyncErrorType = iota

	DriveSyncErrorType_DELETE_WITHOUT_PERMISSION
	DriveSyncErrorType_SERVICE_UNAVAILABLE
	DriveSyncErrorType_NO_SERVER_SPACE
	DriveSyncErrorType_NO_SERVER_SPACE_ORGANIZATION
	DriveSyncErrorType_NO_LOCAL_SPACE
	DriveSyncErrorType_NO_SHARED_DRIVE_SPACE
	DriveSyncErrorType_MISC
)

func (DriveSyncErrorType) FromRef(str js.Ref) DriveSyncErrorType {
	return DriveSyncErrorType(bindings.ConstOfDriveSyncErrorType(str))
}

func (x DriveSyncErrorType) String() (string, bool) {
	switch x {
	case DriveSyncErrorType_DELETE_WITHOUT_PERMISSION:
		return "delete_without_permission", true
	case DriveSyncErrorType_SERVICE_UNAVAILABLE:
		return "service_unavailable", true
	case DriveSyncErrorType_NO_SERVER_SPACE:
		return "no_server_space", true
	case DriveSyncErrorType_NO_SERVER_SPACE_ORGANIZATION:
		return "no_server_space_organization", true
	case DriveSyncErrorType_NO_LOCAL_SPACE:
		return "no_local_space", true
	case DriveSyncErrorType_NO_SHARED_DRIVE_SPACE:
		return "no_shared_drive_space", true
	case DriveSyncErrorType_MISC:
		return "misc", true
	default:
		return "", false
	}
}

type DriveSyncErrorEvent struct {
	// Type is "DriveSyncErrorEvent.type"
	//
	// Optional
	Type DriveSyncErrorType
	// FileUrl is "DriveSyncErrorEvent.fileUrl"
	//
	// Optional
	FileUrl js.String
	// SharedDrive is "DriveSyncErrorEvent.sharedDrive"
	//
	// Optional
	SharedDrive js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DriveSyncErrorEvent with all fields set.
func (p DriveSyncErrorEvent) FromRef(ref js.Ref) DriveSyncErrorEvent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DriveSyncErrorEvent in the application heap.
func (p DriveSyncErrorEvent) New() js.Ref {
	return bindings.DriveSyncErrorEventJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DriveSyncErrorEvent) UpdateFrom(ref js.Ref) {
	bindings.DriveSyncErrorEventJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DriveSyncErrorEvent) Update(ref js.Ref) {
	bindings.DriveSyncErrorEventJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DriveSyncErrorEvent) FreeMembers(recursive bool) {
	js.Free(
		p.FileUrl.Ref(),
		p.SharedDrive.Ref(),
	)
	p.FileUrl = p.FileUrl.FromRef(js.Undefined)
	p.SharedDrive = p.SharedDrive.FromRef(js.Undefined)
}

type RecentDateBucket uint32

const (
	_ RecentDateBucket = iota

	RecentDateBucket_TODAY
	RecentDateBucket_YESTERDAY
	RecentDateBucket_EARLIER_THIS_WEEK
	RecentDateBucket_EARLIER_THIS_MONTH
	RecentDateBucket_EARLIER_THIS_YEAR
	RecentDateBucket_OLDER
)

func (RecentDateBucket) FromRef(str js.Ref) RecentDateBucket {
	return RecentDateBucket(bindings.ConstOfRecentDateBucket(str))
}

func (x RecentDateBucket) String() (string, bool) {
	switch x {
	case RecentDateBucket_TODAY:
		return "today", true
	case RecentDateBucket_YESTERDAY:
		return "yesterday", true
	case RecentDateBucket_EARLIER_THIS_WEEK:
		return "earlier_this_week", true
	case RecentDateBucket_EARLIER_THIS_MONTH:
		return "earlier_this_month", true
	case RecentDateBucket_EARLIER_THIS_YEAR:
		return "earlier_this_year", true
	case RecentDateBucket_OLDER:
		return "older", true
	default:
		return "", false
	}
}

type SyncStatus uint32

const (
	_ SyncStatus = iota

	SyncStatus_NOT_FOUND
	SyncStatus_QUEUED
	SyncStatus_IN_PROGRESS
	SyncStatus_COMPLETED
	SyncStatus_ERROR
)

func (SyncStatus) FromRef(str js.Ref) SyncStatus {
	return SyncStatus(bindings.ConstOfSyncStatus(str))
}

func (x SyncStatus) String() (string, bool) {
	switch x {
	case SyncStatus_NOT_FOUND:
		return "not_found", true
	case SyncStatus_QUEUED:
		return "queued", true
	case SyncStatus_IN_PROGRESS:
		return "in_progress", true
	case SyncStatus_COMPLETED:
		return "completed", true
	case SyncStatus_ERROR:
		return "error", true
	default:
		return "", false
	}
}

type EntryProperties struct {
	// Size is "EntryProperties.size"
	//
	// Optional
	//
	// NOTE: FFI_USE_Size MUST be set to true to make this field effective.
	Size float64
	// ModificationTime is "EntryProperties.modificationTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_ModificationTime MUST be set to true to make this field effective.
	ModificationTime float64
	// ModificationByMeTime is "EntryProperties.modificationByMeTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_ModificationByMeTime MUST be set to true to make this field effective.
	ModificationByMeTime float64
	// RecentDateBucket is "EntryProperties.recentDateBucket"
	//
	// Optional
	RecentDateBucket RecentDateBucket
	// ThumbnailUrl is "EntryProperties.thumbnailUrl"
	//
	// Optional
	ThumbnailUrl js.String
	// CroppedThumbnailUrl is "EntryProperties.croppedThumbnailUrl"
	//
	// Optional
	CroppedThumbnailUrl js.String
	// ImageWidth is "EntryProperties.imageWidth"
	//
	// Optional
	//
	// NOTE: FFI_USE_ImageWidth MUST be set to true to make this field effective.
	ImageWidth int32
	// ImageHeight is "EntryProperties.imageHeight"
	//
	// Optional
	//
	// NOTE: FFI_USE_ImageHeight MUST be set to true to make this field effective.
	ImageHeight int32
	// ImageRotation is "EntryProperties.imageRotation"
	//
	// Optional
	//
	// NOTE: FFI_USE_ImageRotation MUST be set to true to make this field effective.
	ImageRotation int32
	// Pinned is "EntryProperties.pinned"
	//
	// Optional
	//
	// NOTE: FFI_USE_Pinned MUST be set to true to make this field effective.
	Pinned bool
	// Present is "EntryProperties.present"
	//
	// Optional
	//
	// NOTE: FFI_USE_Present MUST be set to true to make this field effective.
	Present bool
	// Hosted is "EntryProperties.hosted"
	//
	// Optional
	//
	// NOTE: FFI_USE_Hosted MUST be set to true to make this field effective.
	Hosted bool
	// AvailableOffline is "EntryProperties.availableOffline"
	//
	// Optional
	//
	// NOTE: FFI_USE_AvailableOffline MUST be set to true to make this field effective.
	AvailableOffline bool
	// AvailableWhenMetered is "EntryProperties.availableWhenMetered"
	//
	// Optional
	//
	// NOTE: FFI_USE_AvailableWhenMetered MUST be set to true to make this field effective.
	AvailableWhenMetered bool
	// Dirty is "EntryProperties.dirty"
	//
	// Optional
	//
	// NOTE: FFI_USE_Dirty MUST be set to true to make this field effective.
	Dirty bool
	// CustomIconUrl is "EntryProperties.customIconUrl"
	//
	// Optional
	CustomIconUrl js.String
	// ContentMimeType is "EntryProperties.contentMimeType"
	//
	// Optional
	ContentMimeType js.String
	// SharedWithMe is "EntryProperties.sharedWithMe"
	//
	// Optional
	//
	// NOTE: FFI_USE_SharedWithMe MUST be set to true to make this field effective.
	SharedWithMe bool
	// Shared is "EntryProperties.shared"
	//
	// Optional
	//
	// NOTE: FFI_USE_Shared MUST be set to true to make this field effective.
	Shared bool
	// Starred is "EntryProperties.starred"
	//
	// Optional
	//
	// NOTE: FFI_USE_Starred MUST be set to true to make this field effective.
	Starred bool
	// ExternalFileUrl is "EntryProperties.externalFileUrl"
	//
	// Optional
	ExternalFileUrl js.String
	// AlternateUrl is "EntryProperties.alternateUrl"
	//
	// Optional
	AlternateUrl js.String
	// ShareUrl is "EntryProperties.shareUrl"
	//
	// Optional
	ShareUrl js.String
	// CanCopy is "EntryProperties.canCopy"
	//
	// Optional
	//
	// NOTE: FFI_USE_CanCopy MUST be set to true to make this field effective.
	CanCopy bool
	// CanDelete is "EntryProperties.canDelete"
	//
	// Optional
	//
	// NOTE: FFI_USE_CanDelete MUST be set to true to make this field effective.
	CanDelete bool
	// CanRename is "EntryProperties.canRename"
	//
	// Optional
	//
	// NOTE: FFI_USE_CanRename MUST be set to true to make this field effective.
	CanRename bool
	// CanAddChildren is "EntryProperties.canAddChildren"
	//
	// Optional
	//
	// NOTE: FFI_USE_CanAddChildren MUST be set to true to make this field effective.
	CanAddChildren bool
	// CanShare is "EntryProperties.canShare"
	//
	// Optional
	//
	// NOTE: FFI_USE_CanShare MUST be set to true to make this field effective.
	CanShare bool
	// CanPin is "EntryProperties.canPin"
	//
	// Optional
	//
	// NOTE: FFI_USE_CanPin MUST be set to true to make this field effective.
	CanPin bool
	// IsMachineRoot is "EntryProperties.isMachineRoot"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsMachineRoot MUST be set to true to make this field effective.
	IsMachineRoot bool
	// IsExternalMedia is "EntryProperties.isExternalMedia"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsExternalMedia MUST be set to true to make this field effective.
	IsExternalMedia bool
	// IsArbitrarySyncFolder is "EntryProperties.isArbitrarySyncFolder"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsArbitrarySyncFolder MUST be set to true to make this field effective.
	IsArbitrarySyncFolder bool
	// SyncStatus is "EntryProperties.syncStatus"
	//
	// Optional
	SyncStatus SyncStatus
	// Progress is "EntryProperties.progress"
	//
	// Optional
	//
	// NOTE: FFI_USE_Progress MUST be set to true to make this field effective.
	Progress float64
	// SyncCompletedTime is "EntryProperties.syncCompletedTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_SyncCompletedTime MUST be set to true to make this field effective.
	SyncCompletedTime float64
	// Shortcut is "EntryProperties.shortcut"
	//
	// Optional
	//
	// NOTE: FFI_USE_Shortcut MUST be set to true to make this field effective.
	Shortcut bool

	FFI_USE_Size                  bool // for Size.
	FFI_USE_ModificationTime      bool // for ModificationTime.
	FFI_USE_ModificationByMeTime  bool // for ModificationByMeTime.
	FFI_USE_ImageWidth            bool // for ImageWidth.
	FFI_USE_ImageHeight           bool // for ImageHeight.
	FFI_USE_ImageRotation         bool // for ImageRotation.
	FFI_USE_Pinned                bool // for Pinned.
	FFI_USE_Present               bool // for Present.
	FFI_USE_Hosted                bool // for Hosted.
	FFI_USE_AvailableOffline      bool // for AvailableOffline.
	FFI_USE_AvailableWhenMetered  bool // for AvailableWhenMetered.
	FFI_USE_Dirty                 bool // for Dirty.
	FFI_USE_SharedWithMe          bool // for SharedWithMe.
	FFI_USE_Shared                bool // for Shared.
	FFI_USE_Starred               bool // for Starred.
	FFI_USE_CanCopy               bool // for CanCopy.
	FFI_USE_CanDelete             bool // for CanDelete.
	FFI_USE_CanRename             bool // for CanRename.
	FFI_USE_CanAddChildren        bool // for CanAddChildren.
	FFI_USE_CanShare              bool // for CanShare.
	FFI_USE_CanPin                bool // for CanPin.
	FFI_USE_IsMachineRoot         bool // for IsMachineRoot.
	FFI_USE_IsExternalMedia       bool // for IsExternalMedia.
	FFI_USE_IsArbitrarySyncFolder bool // for IsArbitrarySyncFolder.
	FFI_USE_Progress              bool // for Progress.
	FFI_USE_SyncCompletedTime     bool // for SyncCompletedTime.
	FFI_USE_Shortcut              bool // for Shortcut.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EntryProperties with all fields set.
func (p EntryProperties) FromRef(ref js.Ref) EntryProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EntryProperties in the application heap.
func (p EntryProperties) New() js.Ref {
	return bindings.EntryPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *EntryProperties) UpdateFrom(ref js.Ref) {
	bindings.EntryPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EntryProperties) Update(ref js.Ref) {
	bindings.EntryPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EntryProperties) FreeMembers(recursive bool) {
	js.Free(
		p.ThumbnailUrl.Ref(),
		p.CroppedThumbnailUrl.Ref(),
		p.CustomIconUrl.Ref(),
		p.ContentMimeType.Ref(),
		p.ExternalFileUrl.Ref(),
		p.AlternateUrl.Ref(),
		p.ShareUrl.Ref(),
	)
	p.ThumbnailUrl = p.ThumbnailUrl.FromRef(js.Undefined)
	p.CroppedThumbnailUrl = p.CroppedThumbnailUrl.FromRef(js.Undefined)
	p.CustomIconUrl = p.CustomIconUrl.FromRef(js.Undefined)
	p.ContentMimeType = p.ContentMimeType.FromRef(js.Undefined)
	p.ExternalFileUrl = p.ExternalFileUrl.FromRef(js.Undefined)
	p.AlternateUrl = p.AlternateUrl.FromRef(js.Undefined)
	p.ShareUrl = p.ShareUrl.FromRef(js.Undefined)
}

type EntryPropertyName uint32

const (
	_ EntryPropertyName = iota

	EntryPropertyName_SIZE
	EntryPropertyName_MODIFICATION_TIME
	EntryPropertyName_MODIFICATION_BY_ME_TIME
	EntryPropertyName_THUMBNAIL_URL
	EntryPropertyName_CROPPED_THUMBNAIL_URL
	EntryPropertyName_IMAGE_WIDTH
	EntryPropertyName_IMAGE_HEIGHT
	EntryPropertyName_IMAGE_ROTATION
	EntryPropertyName_PINNED
	EntryPropertyName_PRESENT
	EntryPropertyName_HOSTED
	EntryPropertyName_AVAILABLE_OFFLINE
	EntryPropertyName_AVAILABLE_WHEN_METERED
	EntryPropertyName_DIRTY
	EntryPropertyName_CUSTOM_ICON_URL
	EntryPropertyName_CONTENT_MIME_TYPE
	EntryPropertyName_SHARED_WITH_ME
	EntryPropertyName_SHARED
	EntryPropertyName_STARRED
	EntryPropertyName_EXTERNAL_FILE_URL
	EntryPropertyName_ALTERNATE_URL
	EntryPropertyName_SHARE_URL
	EntryPropertyName_CAN_COPY
	EntryPropertyName_CAN_DELETE
	EntryPropertyName_CAN_RENAME
	EntryPropertyName_CAN_ADD_CHILDREN
	EntryPropertyName_CAN_SHARE
	EntryPropertyName_CAN_PIN
	EntryPropertyName_IS_MACHINE_ROOT
	EntryPropertyName_IS_EXTERNAL_MEDIA
	EntryPropertyName_IS_ARBITRARY_SYNC_FOLDER
	EntryPropertyName_SYNC_STATUS
	EntryPropertyName_PROGRESS
	EntryPropertyName_SHORTCUT
	EntryPropertyName_SYNC_COMPLETED_TIME
)

func (EntryPropertyName) FromRef(str js.Ref) EntryPropertyName {
	return EntryPropertyName(bindings.ConstOfEntryPropertyName(str))
}

func (x EntryPropertyName) String() (string, bool) {
	switch x {
	case EntryPropertyName_SIZE:
		return "size", true
	case EntryPropertyName_MODIFICATION_TIME:
		return "modificationTime", true
	case EntryPropertyName_MODIFICATION_BY_ME_TIME:
		return "modificationByMeTime", true
	case EntryPropertyName_THUMBNAIL_URL:
		return "thumbnailUrl", true
	case EntryPropertyName_CROPPED_THUMBNAIL_URL:
		return "croppedThumbnailUrl", true
	case EntryPropertyName_IMAGE_WIDTH:
		return "imageWidth", true
	case EntryPropertyName_IMAGE_HEIGHT:
		return "imageHeight", true
	case EntryPropertyName_IMAGE_ROTATION:
		return "imageRotation", true
	case EntryPropertyName_PINNED:
		return "pinned", true
	case EntryPropertyName_PRESENT:
		return "present", true
	case EntryPropertyName_HOSTED:
		return "hosted", true
	case EntryPropertyName_AVAILABLE_OFFLINE:
		return "availableOffline", true
	case EntryPropertyName_AVAILABLE_WHEN_METERED:
		return "availableWhenMetered", true
	case EntryPropertyName_DIRTY:
		return "dirty", true
	case EntryPropertyName_CUSTOM_ICON_URL:
		return "customIconUrl", true
	case EntryPropertyName_CONTENT_MIME_TYPE:
		return "contentMimeType", true
	case EntryPropertyName_SHARED_WITH_ME:
		return "sharedWithMe", true
	case EntryPropertyName_SHARED:
		return "shared", true
	case EntryPropertyName_STARRED:
		return "starred", true
	case EntryPropertyName_EXTERNAL_FILE_URL:
		return "externalFileUrl", true
	case EntryPropertyName_ALTERNATE_URL:
		return "alternateUrl", true
	case EntryPropertyName_SHARE_URL:
		return "shareUrl", true
	case EntryPropertyName_CAN_COPY:
		return "canCopy", true
	case EntryPropertyName_CAN_DELETE:
		return "canDelete", true
	case EntryPropertyName_CAN_RENAME:
		return "canRename", true
	case EntryPropertyName_CAN_ADD_CHILDREN:
		return "canAddChildren", true
	case EntryPropertyName_CAN_SHARE:
		return "canShare", true
	case EntryPropertyName_CAN_PIN:
		return "canPin", true
	case EntryPropertyName_IS_MACHINE_ROOT:
		return "isMachineRoot", true
	case EntryPropertyName_IS_EXTERNAL_MEDIA:
		return "isExternalMedia", true
	case EntryPropertyName_IS_ARBITRARY_SYNC_FOLDER:
		return "isArbitrarySyncFolder", true
	case EntryPropertyName_SYNC_STATUS:
		return "syncStatus", true
	case EntryPropertyName_PROGRESS:
		return "progress", true
	case EntryPropertyName_SHORTCUT:
		return "shortcut", true
	case EntryPropertyName_SYNC_COMPLETED_TIME:
		return "syncCompletedTime", true
	default:
		return "", false
	}
}

type TaskResult uint32

const (
	_ TaskResult = iota

	TaskResult_OPENED
	TaskResult_MESSAGE_SENT
	TaskResult_FAILED
	TaskResult_EMPTY
	TaskResult_FAILED_PLUGIN_VM_DIRECTORY_NOT_SHARED
)

func (TaskResult) FromRef(str js.Ref) TaskResult {
	return TaskResult(bindings.ConstOfTaskResult(str))
}

func (x TaskResult) String() (string, bool) {
	switch x {
	case TaskResult_OPENED:
		return "opened", true
	case TaskResult_MESSAGE_SENT:
		return "message_sent", true
	case TaskResult_FAILED:
		return "failed", true
	case TaskResult_EMPTY:
		return "empty", true
	case TaskResult_FAILED_PLUGIN_VM_DIRECTORY_NOT_SHARED:
		return "failed_plugin_vm_directory_not_shared", true
	default:
		return "", false
	}
}

type ExecuteTaskCallbackFunc func(this js.Ref, result TaskResult) js.Ref

func (fn ExecuteTaskCallbackFunc) Register() js.Func[func(result TaskResult)] {
	return js.RegisterCallback[func(result TaskResult)](
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

		TaskResult(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ExecuteTaskCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result TaskResult) js.Ref
	Arg T
}

func (cb *ExecuteTaskCallback[T]) Register() js.Func[func(result TaskResult)] {
	return js.RegisterCallback[func(result TaskResult)](
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

		TaskResult(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type FileCategory uint32

const (
	_ FileCategory = iota

	FileCategory_ALL
	FileCategory_AUDIO
	FileCategory_IMAGE
	FileCategory_VIDEO
	FileCategory_DOCUMENT
)

func (FileCategory) FromRef(str js.Ref) FileCategory {
	return FileCategory(bindings.ConstOfFileCategory(str))
}

func (x FileCategory) String() (string, bool) {
	switch x {
	case FileCategory_ALL:
		return "all", true
	case FileCategory_AUDIO:
		return "audio", true
	case FileCategory_IMAGE:
		return "image", true
	case FileCategory_VIDEO:
		return "video", true
	case FileCategory_DOCUMENT:
		return "document", true
	default:
		return "", false
	}
}

type FileChange struct {
	// Url is "FileChange.url"
	//
	// Optional
	Url js.String
	// Changes is "FileChange.changes"
	//
	// Optional
	Changes js.Array[ChangeType]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FileChange with all fields set.
func (p FileChange) FromRef(ref js.Ref) FileChange {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FileChange in the application heap.
func (p FileChange) New() js.Ref {
	return bindings.FileChangeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FileChange) UpdateFrom(ref js.Ref) {
	bindings.FileChangeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FileChange) Update(ref js.Ref) {
	bindings.FileChangeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FileChange) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
		p.Changes.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
	p.Changes = p.Changes.FromRef(js.Undefined)
}

type FileSystemProviderAction struct {
	// Id is "FileSystemProviderAction.id"
	//
	// Optional
	Id js.String
	// Title is "FileSystemProviderAction.title"
	//
	// Optional
	Title js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FileSystemProviderAction with all fields set.
func (p FileSystemProviderAction) FromRef(ref js.Ref) FileSystemProviderAction {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FileSystemProviderAction in the application heap.
func (p FileSystemProviderAction) New() js.Ref {
	return bindings.FileSystemProviderActionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FileSystemProviderAction) UpdateFrom(ref js.Ref) {
	bindings.FileSystemProviderActionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FileSystemProviderAction) Update(ref js.Ref) {
	bindings.FileSystemProviderActionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FileSystemProviderAction) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.Title.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
}

type FileTaskDescriptor struct {
	// AppId is "FileTaskDescriptor.appId"
	//
	// Optional
	AppId js.String
	// TaskType is "FileTaskDescriptor.taskType"
	//
	// Optional
	TaskType js.String
	// ActionId is "FileTaskDescriptor.actionId"
	//
	// Optional
	ActionId js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FileTaskDescriptor with all fields set.
func (p FileTaskDescriptor) FromRef(ref js.Ref) FileTaskDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FileTaskDescriptor in the application heap.
func (p FileTaskDescriptor) New() js.Ref {
	return bindings.FileTaskDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FileTaskDescriptor) UpdateFrom(ref js.Ref) {
	bindings.FileTaskDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FileTaskDescriptor) Update(ref js.Ref) {
	bindings.FileTaskDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FileTaskDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.AppId.Ref(),
		p.TaskType.Ref(),
		p.ActionId.Ref(),
	)
	p.AppId = p.AppId.FromRef(js.Undefined)
	p.TaskType = p.TaskType.FromRef(js.Undefined)
	p.ActionId = p.ActionId.FromRef(js.Undefined)
}

type FileTask struct {
	// Descriptor is "FileTask.descriptor"
	//
	// Optional
	//
	// NOTE: Descriptor.FFI_USE MUST be set to true to get Descriptor used.
	Descriptor FileTaskDescriptor
	// Title is "FileTask.title"
	//
	// Optional
	Title js.String
	// IconUrl is "FileTask.iconUrl"
	//
	// Optional
	IconUrl js.String
	// IsDefault is "FileTask.isDefault"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsDefault MUST be set to true to make this field effective.
	IsDefault bool
	// IsGenericFileHandler is "FileTask.isGenericFileHandler"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsGenericFileHandler MUST be set to true to make this field effective.
	IsGenericFileHandler bool
	// IsDlpBlocked is "FileTask.isDlpBlocked"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsDlpBlocked MUST be set to true to make this field effective.
	IsDlpBlocked bool

	FFI_USE_IsDefault            bool // for IsDefault.
	FFI_USE_IsGenericFileHandler bool // for IsGenericFileHandler.
	FFI_USE_IsDlpBlocked         bool // for IsDlpBlocked.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FileTask with all fields set.
func (p FileTask) FromRef(ref js.Ref) FileTask {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FileTask in the application heap.
func (p FileTask) New() js.Ref {
	return bindings.FileTaskJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FileTask) UpdateFrom(ref js.Ref) {
	bindings.FileTaskJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FileTask) Update(ref js.Ref) {
	bindings.FileTaskJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FileTask) FreeMembers(recursive bool) {
	js.Free(
		p.Title.Ref(),
		p.IconUrl.Ref(),
	)
	p.Title = p.Title.FromRef(js.Undefined)
	p.IconUrl = p.IconUrl.FromRef(js.Undefined)
	if recursive {
		p.Descriptor.FreeMembers(true)
	}
}

type TransferState uint32

const (
	_ TransferState = iota

	TransferState_IN_PROGRESS
	TransferState_QUEUED
	TransferState_COMPLETED
	TransferState_FAILED
)

func (TransferState) FromRef(str js.Ref) TransferState {
	return TransferState(bindings.ConstOfTransferState(str))
}

func (x TransferState) String() (string, bool) {
	switch x {
	case TransferState_IN_PROGRESS:
		return "in_progress", true
	case TransferState_QUEUED:
		return "queued", true
	case TransferState_COMPLETED:
		return "completed", true
	case TransferState_FAILED:
		return "failed", true
	default:
		return "", false
	}
}

type FileTransferStatus struct {
	// FileUrl is "FileTransferStatus.fileUrl"
	//
	// Optional
	FileUrl js.String
	// TransferState is "FileTransferStatus.transferState"
	//
	// Optional
	TransferState TransferState
	// Processed is "FileTransferStatus.processed"
	//
	// Optional
	//
	// NOTE: FFI_USE_Processed MUST be set to true to make this field effective.
	Processed float64
	// Total is "FileTransferStatus.total"
	//
	// Optional
	//
	// NOTE: FFI_USE_Total MUST be set to true to make this field effective.
	Total float64
	// NumTotalJobs is "FileTransferStatus.numTotalJobs"
	//
	// Optional
	//
	// NOTE: FFI_USE_NumTotalJobs MUST be set to true to make this field effective.
	NumTotalJobs int32
	// ShowNotification is "FileTransferStatus.showNotification"
	//
	// Optional
	//
	// NOTE: FFI_USE_ShowNotification MUST be set to true to make this field effective.
	ShowNotification bool
	// HideWhenZeroJobs is "FileTransferStatus.hideWhenZeroJobs"
	//
	// Optional
	//
	// NOTE: FFI_USE_HideWhenZeroJobs MUST be set to true to make this field effective.
	HideWhenZeroJobs bool

	FFI_USE_Processed        bool // for Processed.
	FFI_USE_Total            bool // for Total.
	FFI_USE_NumTotalJobs     bool // for NumTotalJobs.
	FFI_USE_ShowNotification bool // for ShowNotification.
	FFI_USE_HideWhenZeroJobs bool // for HideWhenZeroJobs.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FileTransferStatus with all fields set.
func (p FileTransferStatus) FromRef(ref js.Ref) FileTransferStatus {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FileTransferStatus in the application heap.
func (p FileTransferStatus) New() js.Ref {
	return bindings.FileTransferStatusJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FileTransferStatus) UpdateFrom(ref js.Ref) {
	bindings.FileTransferStatusJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FileTransferStatus) Update(ref js.Ref) {
	bindings.FileTransferStatusJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FileTransferStatus) FreeMembers(recursive bool) {
	js.Free(
		p.FileUrl.Ref(),
	)
	p.FileUrl = p.FileUrl.FromRef(js.Undefined)
}

type FileWatchEventType uint32

const (
	_ FileWatchEventType = iota

	FileWatchEventType_CHANGED
	FileWatchEventType_ERROR
)

func (FileWatchEventType) FromRef(str js.Ref) FileWatchEventType {
	return FileWatchEventType(bindings.ConstOfFileWatchEventType(str))
}

func (x FileWatchEventType) String() (string, bool) {
	switch x {
	case FileWatchEventType_CHANGED:
		return "changed", true
	case FileWatchEventType_ERROR:
		return "error", true
	default:
		return "", false
	}
}

type FileWatchEvent struct {
	// EventType is "FileWatchEvent.eventType"
	//
	// Optional
	EventType FileWatchEventType
	// Entry is "FileWatchEvent.entry"
	//
	// Optional
	Entry js.Object
	// ChangedFiles is "FileWatchEvent.changedFiles"
	//
	// Optional
	ChangedFiles js.Array[FileChange]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FileWatchEvent with all fields set.
func (p FileWatchEvent) FromRef(ref js.Ref) FileWatchEvent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FileWatchEvent in the application heap.
func (p FileWatchEvent) New() js.Ref {
	return bindings.FileWatchEventJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FileWatchEvent) UpdateFrom(ref js.Ref) {
	bindings.FileWatchEventJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FileWatchEvent) Update(ref js.Ref) {
	bindings.FileWatchEventJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FileWatchEvent) FreeMembers(recursive bool) {
	js.Free(
		p.Entry.Ref(),
		p.ChangedFiles.Ref(),
	)
	p.Entry = p.Entry.FromRef(js.Undefined)
	p.ChangedFiles = p.ChangedFiles.FromRef(js.Undefined)
}

type FormatFileSystemType uint32

const (
	_ FormatFileSystemType = iota

	FormatFileSystemType_VFAT
	FormatFileSystemType_EXFAT
	FormatFileSystemType_NTFS
)

func (FormatFileSystemType) FromRef(str js.Ref) FormatFileSystemType {
	return FormatFileSystemType(bindings.ConstOfFormatFileSystemType(str))
}

func (x FormatFileSystemType) String() (string, bool) {
	switch x {
	case FormatFileSystemType_VFAT:
		return "vfat", true
	case FormatFileSystemType_EXFAT:
		return "exfat", true
	case FormatFileSystemType_NTFS:
		return "ntfs", true
	default:
		return "", false
	}
}

type GetAndroidPickerAppsCallbackFunc func(this js.Ref, apps js.Array[AndroidApp]) js.Ref

func (fn GetAndroidPickerAppsCallbackFunc) Register() js.Func[func(apps js.Array[AndroidApp])] {
	return js.RegisterCallback[func(apps js.Array[AndroidApp])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetAndroidPickerAppsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[AndroidApp]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetAndroidPickerAppsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, apps js.Array[AndroidApp]) js.Ref
	Arg T
}

func (cb *GetAndroidPickerAppsCallback[T]) Register() js.Func[func(apps js.Array[AndroidApp])] {
	return js.RegisterCallback[func(apps js.Array[AndroidApp])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetAndroidPickerAppsCallback[T]) DispatchCallback(
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

		js.Array[AndroidApp]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetBulkPinProgressCallbackFunc func(this js.Ref, progress *BulkPinProgress) js.Ref

func (fn GetBulkPinProgressCallbackFunc) Register() js.Func[func(progress *BulkPinProgress)] {
	return js.RegisterCallback[func(progress *BulkPinProgress)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetBulkPinProgressCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 BulkPinProgress
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

type GetBulkPinProgressCallback[T any] struct {
	Fn  func(arg T, this js.Ref, progress *BulkPinProgress) js.Ref
	Arg T
}

func (cb *GetBulkPinProgressCallback[T]) Register() js.Func[func(progress *BulkPinProgress)] {
	return js.RegisterCallback[func(progress *BulkPinProgress)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetBulkPinProgressCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 BulkPinProgress
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

type StreamInfo struct {
	// Type is "StreamInfo.type"
	//
	// Optional
	Type js.String
	// Tags is "StreamInfo.tags"
	//
	// Optional
	Tags js.Object

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StreamInfo with all fields set.
func (p StreamInfo) FromRef(ref js.Ref) StreamInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new StreamInfo in the application heap.
func (p StreamInfo) New() js.Ref {
	return bindings.StreamInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *StreamInfo) UpdateFrom(ref js.Ref) {
	bindings.StreamInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StreamInfo) Update(ref js.Ref) {
	bindings.StreamInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StreamInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Type.Ref(),
		p.Tags.Ref(),
	)
	p.Type = p.Type.FromRef(js.Undefined)
	p.Tags = p.Tags.FromRef(js.Undefined)
}

type MediaMetadata struct {
	// MimeType is "MediaMetadata.mimeType"
	//
	// Optional
	MimeType js.String
	// Height is "MediaMetadata.height"
	//
	// Optional
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height int32
	// Width is "MediaMetadata.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width int32
	// Duration is "MediaMetadata.duration"
	//
	// Optional
	//
	// NOTE: FFI_USE_Duration MUST be set to true to make this field effective.
	Duration float64
	// Rotation is "MediaMetadata.rotation"
	//
	// Optional
	//
	// NOTE: FFI_USE_Rotation MUST be set to true to make this field effective.
	Rotation int32
	// Album is "MediaMetadata.album"
	//
	// Optional
	Album js.String
	// Artist is "MediaMetadata.artist"
	//
	// Optional
	Artist js.String
	// Comment is "MediaMetadata.comment"
	//
	// Optional
	Comment js.String
	// Copyright is "MediaMetadata.copyright"
	//
	// Optional
	Copyright js.String
	// Disc is "MediaMetadata.disc"
	//
	// Optional
	//
	// NOTE: FFI_USE_Disc MUST be set to true to make this field effective.
	Disc int32
	// Genre is "MediaMetadata.genre"
	//
	// Optional
	Genre js.String
	// Language is "MediaMetadata.language"
	//
	// Optional
	Language js.String
	// Title is "MediaMetadata.title"
	//
	// Optional
	Title js.String
	// Track is "MediaMetadata.track"
	//
	// Optional
	//
	// NOTE: FFI_USE_Track MUST be set to true to make this field effective.
	Track int32
	// RawTags is "MediaMetadata.rawTags"
	//
	// Optional
	RawTags js.Array[StreamInfo]
	// AttachedImages is "MediaMetadata.attachedImages"
	//
	// Optional
	AttachedImages js.Array[AttachedImages]

	FFI_USE_Height   bool // for Height.
	FFI_USE_Width    bool // for Width.
	FFI_USE_Duration bool // for Duration.
	FFI_USE_Rotation bool // for Rotation.
	FFI_USE_Disc     bool // for Disc.
	FFI_USE_Track    bool // for Track.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaMetadata with all fields set.
func (p MediaMetadata) FromRef(ref js.Ref) MediaMetadata {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaMetadata in the application heap.
func (p MediaMetadata) New() js.Ref {
	return bindings.MediaMetadataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MediaMetadata) UpdateFrom(ref js.Ref) {
	bindings.MediaMetadataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MediaMetadata) Update(ref js.Ref) {
	bindings.MediaMetadataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MediaMetadata) FreeMembers(recursive bool) {
	js.Free(
		p.MimeType.Ref(),
		p.Album.Ref(),
		p.Artist.Ref(),
		p.Comment.Ref(),
		p.Copyright.Ref(),
		p.Genre.Ref(),
		p.Language.Ref(),
		p.Title.Ref(),
		p.RawTags.Ref(),
		p.AttachedImages.Ref(),
	)
	p.MimeType = p.MimeType.FromRef(js.Undefined)
	p.Album = p.Album.FromRef(js.Undefined)
	p.Artist = p.Artist.FromRef(js.Undefined)
	p.Comment = p.Comment.FromRef(js.Undefined)
	p.Copyright = p.Copyright.FromRef(js.Undefined)
	p.Genre = p.Genre.FromRef(js.Undefined)
	p.Language = p.Language.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
	p.RawTags = p.RawTags.FromRef(js.Undefined)
	p.AttachedImages = p.AttachedImages.FromRef(js.Undefined)
}

type GetContentMetadataCallbackFunc func(this js.Ref, result *MediaMetadata) js.Ref

func (fn GetContentMetadataCallbackFunc) Register() js.Func[func(result *MediaMetadata)] {
	return js.RegisterCallback[func(result *MediaMetadata)](
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
	var arg0 MediaMetadata
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
	Fn  func(arg T, this js.Ref, result *MediaMetadata) js.Ref
	Arg T
}

func (cb *GetContentMetadataCallback[T]) Register() js.Func[func(result *MediaMetadata)] {
	return js.RegisterCallback[func(result *MediaMetadata)](
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
	var arg0 MediaMetadata
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

type GetCrostiniSharedPathsCallbackFunc func(this js.Ref, entries js.Array[js.Object], firstForSession bool) js.Ref

func (fn GetCrostiniSharedPathsCallbackFunc) Register() js.Func[func(entries js.Array[js.Object], firstForSession bool)] {
	return js.RegisterCallback[func(entries js.Array[js.Object], firstForSession bool)](
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

		js.Array[js.Object]{}.FromRef(args[0+1]),
		args[1+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetCrostiniSharedPathsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[js.Object], firstForSession bool) js.Ref
	Arg T
}

func (cb *GetCrostiniSharedPathsCallback[T]) Register() js.Func[func(entries js.Array[js.Object], firstForSession bool)] {
	return js.RegisterCallback[func(entries js.Array[js.Object], firstForSession bool)](
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

		js.Array[js.Object]{}.FromRef(args[0+1]),
		args[1+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetCustomActionsCallbackFunc func(this js.Ref, actions js.Array[FileSystemProviderAction]) js.Ref

func (fn GetCustomActionsCallbackFunc) Register() js.Func[func(actions js.Array[FileSystemProviderAction])] {
	return js.RegisterCallback[func(actions js.Array[FileSystemProviderAction])](
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

		js.Array[FileSystemProviderAction]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetCustomActionsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, actions js.Array[FileSystemProviderAction]) js.Ref
	Arg T
}

func (cb *GetCustomActionsCallback[T]) Register() js.Func[func(actions js.Array[FileSystemProviderAction])] {
	return js.RegisterCallback[func(actions js.Array[FileSystemProviderAction])](
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

		js.Array[FileSystemProviderAction]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDeviceConnectionStateCallbackFunc func(this js.Ref, result DeviceConnectionState) js.Ref

func (fn GetDeviceConnectionStateCallbackFunc) Register() js.Func[func(result DeviceConnectionState)] {
	return js.RegisterCallback[func(result DeviceConnectionState)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDeviceConnectionStateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		DeviceConnectionState(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDeviceConnectionStateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result DeviceConnectionState) js.Ref
	Arg T
}

func (cb *GetDeviceConnectionStateCallback[T]) Register() js.Func[func(result DeviceConnectionState)] {
	return js.RegisterCallback[func(result DeviceConnectionState)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDeviceConnectionStateCallback[T]) DispatchCallback(
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

		DeviceConnectionState(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDialogCallerCallbackFunc func(this js.Ref, caller *DialogCallerInformation) js.Ref

func (fn GetDialogCallerCallbackFunc) Register() js.Func[func(caller *DialogCallerInformation)] {
	return js.RegisterCallback[func(caller *DialogCallerInformation)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDialogCallerCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DialogCallerInformation
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

type GetDialogCallerCallback[T any] struct {
	Fn  func(arg T, this js.Ref, caller *DialogCallerInformation) js.Ref
	Arg T
}

func (cb *GetDialogCallerCallback[T]) Register() js.Func[func(caller *DialogCallerInformation)] {
	return js.RegisterCallback[func(caller *DialogCallerInformation)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDialogCallerCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DialogCallerInformation
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

type GetDisallowedTransfersCallbackFunc func(this js.Ref, disallowedEntries js.Array[js.Object]) js.Ref

func (fn GetDisallowedTransfersCallbackFunc) Register() js.Func[func(disallowedEntries js.Array[js.Object])] {
	return js.RegisterCallback[func(disallowedEntries js.Array[js.Object])](
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

		js.Array[js.Object]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDisallowedTransfersCallback[T any] struct {
	Fn  func(arg T, this js.Ref, disallowedEntries js.Array[js.Object]) js.Ref
	Arg T
}

func (cb *GetDisallowedTransfersCallback[T]) Register() js.Func[func(disallowedEntries js.Array[js.Object])] {
	return js.RegisterCallback[func(disallowedEntries js.Array[js.Object])](
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

		js.Array[js.Object]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDlpBlockedComponentsCallbackFunc func(this js.Ref, blockedComponents js.Array[VolumeType]) js.Ref

func (fn GetDlpBlockedComponentsCallbackFunc) Register() js.Func[func(blockedComponents js.Array[VolumeType])] {
	return js.RegisterCallback[func(blockedComponents js.Array[VolumeType])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDlpBlockedComponentsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[VolumeType]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDlpBlockedComponentsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, blockedComponents js.Array[VolumeType]) js.Ref
	Arg T
}

func (cb *GetDlpBlockedComponentsCallback[T]) Register() js.Func[func(blockedComponents js.Array[VolumeType])] {
	return js.RegisterCallback[func(blockedComponents js.Array[VolumeType])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDlpBlockedComponentsCallback[T]) DispatchCallback(
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

		js.Array[VolumeType]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDlpMetadataCallbackFunc func(this js.Ref, dlpMetadata js.Array[DlpMetadata]) js.Ref

func (fn GetDlpMetadataCallbackFunc) Register() js.Func[func(dlpMetadata js.Array[DlpMetadata])] {
	return js.RegisterCallback[func(dlpMetadata js.Array[DlpMetadata])](
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

		js.Array[DlpMetadata]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDlpMetadataCallback[T any] struct {
	Fn  func(arg T, this js.Ref, dlpMetadata js.Array[DlpMetadata]) js.Ref
	Arg T
}

func (cb *GetDlpMetadataCallback[T]) Register() js.Func[func(dlpMetadata js.Array[DlpMetadata])] {
	return js.RegisterCallback[func(dlpMetadata js.Array[DlpMetadata])](
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

		js.Array[DlpMetadata]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDlpRestrictionDetailsCallbackFunc func(this js.Ref, restrictionDetails js.Array[DlpRestrictionDetails]) js.Ref

func (fn GetDlpRestrictionDetailsCallbackFunc) Register() js.Func[func(restrictionDetails js.Array[DlpRestrictionDetails])] {
	return js.RegisterCallback[func(restrictionDetails js.Array[DlpRestrictionDetails])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDlpRestrictionDetailsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[DlpRestrictionDetails]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDlpRestrictionDetailsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, restrictionDetails js.Array[DlpRestrictionDetails]) js.Ref
	Arg T
}

func (cb *GetDlpRestrictionDetailsCallback[T]) Register() js.Func[func(restrictionDetails js.Array[DlpRestrictionDetails])] {
	return js.RegisterCallback[func(restrictionDetails js.Array[DlpRestrictionDetails])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDlpRestrictionDetailsCallback[T]) DispatchCallback(
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

		js.Array[DlpRestrictionDetails]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDriveConnectionStateCallbackFunc func(this js.Ref, result *DriveConnectionState) js.Ref

func (fn GetDriveConnectionStateCallbackFunc) Register() js.Func[func(result *DriveConnectionState)] {
	return js.RegisterCallback[func(result *DriveConnectionState)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDriveConnectionStateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DriveConnectionState
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

type GetDriveConnectionStateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *DriveConnectionState) js.Ref
	Arg T
}

func (cb *GetDriveConnectionStateCallback[T]) Register() js.Func[func(result *DriveConnectionState)] {
	return js.RegisterCallback[func(result *DriveConnectionState)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDriveConnectionStateCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DriveConnectionState
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

type GetDriveQuotaMetadataCallbackFunc func(this js.Ref, driveQuotaMetadata *DriveQuotaMetadata) js.Ref

func (fn GetDriveQuotaMetadataCallbackFunc) Register() js.Func[func(driveQuotaMetadata *DriveQuotaMetadata)] {
	return js.RegisterCallback[func(driveQuotaMetadata *DriveQuotaMetadata)](
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
	var arg0 DriveQuotaMetadata
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
	Fn  func(arg T, this js.Ref, driveQuotaMetadata *DriveQuotaMetadata) js.Ref
	Arg T
}

func (cb *GetDriveQuotaMetadataCallback[T]) Register() js.Func[func(driveQuotaMetadata *DriveQuotaMetadata)] {
	return js.RegisterCallback[func(driveQuotaMetadata *DriveQuotaMetadata)](
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
	var arg0 DriveQuotaMetadata
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

type GetEntryPropertiesCallbackFunc func(this js.Ref, entryProperties js.Array[EntryProperties]) js.Ref

func (fn GetEntryPropertiesCallbackFunc) Register() js.Func[func(entryProperties js.Array[EntryProperties])] {
	return js.RegisterCallback[func(entryProperties js.Array[EntryProperties])](
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

		js.Array[EntryProperties]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetEntryPropertiesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entryProperties js.Array[EntryProperties]) js.Ref
	Arg T
}

func (cb *GetEntryPropertiesCallback[T]) Register() js.Func[func(entryProperties js.Array[EntryProperties])] {
	return js.RegisterCallback[func(entryProperties js.Array[EntryProperties])](
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

		js.Array[EntryProperties]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetFileTasksCallbackFunc func(this js.Ref, resultingTasks *ResultingTasks) js.Ref

func (fn GetFileTasksCallbackFunc) Register() js.Func[func(resultingTasks *ResultingTasks)] {
	return js.RegisterCallback[func(resultingTasks *ResultingTasks)](
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
	var arg0 ResultingTasks
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
	Fn  func(arg T, this js.Ref, resultingTasks *ResultingTasks) js.Ref
	Arg T
}

func (cb *GetFileTasksCallback[T]) Register() js.Func[func(resultingTasks *ResultingTasks)] {
	return js.RegisterCallback[func(resultingTasks *ResultingTasks)](
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
	var arg0 ResultingTasks
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

type PolicyDefaultHandlerStatus uint32

const (
	_ PolicyDefaultHandlerStatus = iota

	PolicyDefaultHandlerStatus_DEFAULT_HANDLER_ASSIGNED_BY_POLICY
	PolicyDefaultHandlerStatus_INCORRECT_ASSIGNMENT
)

func (PolicyDefaultHandlerStatus) FromRef(str js.Ref) PolicyDefaultHandlerStatus {
	return PolicyDefaultHandlerStatus(bindings.ConstOfPolicyDefaultHandlerStatus(str))
}

func (x PolicyDefaultHandlerStatus) String() (string, bool) {
	switch x {
	case PolicyDefaultHandlerStatus_DEFAULT_HANDLER_ASSIGNED_BY_POLICY:
		return "default_handler_assigned_by_policy", true
	case PolicyDefaultHandlerStatus_INCORRECT_ASSIGNMENT:
		return "incorrect_assignment", true
	default:
		return "", false
	}
}

type ResultingTasks struct {
	// Tasks is "ResultingTasks.tasks"
	//
	// Optional
	Tasks js.Array[FileTask]
	// PolicyDefaultHandlerStatus is "ResultingTasks.policyDefaultHandlerStatus"
	//
	// Optional
	PolicyDefaultHandlerStatus PolicyDefaultHandlerStatus

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ResultingTasks with all fields set.
func (p ResultingTasks) FromRef(ref js.Ref) ResultingTasks {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ResultingTasks in the application heap.
func (p ResultingTasks) New() js.Ref {
	return bindings.ResultingTasksJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ResultingTasks) UpdateFrom(ref js.Ref) {
	bindings.ResultingTasksJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ResultingTasks) Update(ref js.Ref) {
	bindings.ResultingTasksJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ResultingTasks) FreeMembers(recursive bool) {
	js.Free(
		p.Tasks.Ref(),
	)
	p.Tasks = p.Tasks.FromRef(js.Undefined)
}

type GetLinuxPackageInfoCallbackFunc func(this js.Ref, linux_package_info *LinuxPackageInfo) js.Ref

func (fn GetLinuxPackageInfoCallbackFunc) Register() js.Func[func(linux_package_info *LinuxPackageInfo)] {
	return js.RegisterCallback[func(linux_package_info *LinuxPackageInfo)](
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
	var arg0 LinuxPackageInfo
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
	Fn  func(arg T, this js.Ref, linux_package_info *LinuxPackageInfo) js.Ref
	Arg T
}

func (cb *GetLinuxPackageInfoCallback[T]) Register() js.Func[func(linux_package_info *LinuxPackageInfo)] {
	return js.RegisterCallback[func(linux_package_info *LinuxPackageInfo)](
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
	var arg0 LinuxPackageInfo
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

type LinuxPackageInfo struct {
	// Name is "LinuxPackageInfo.name"
	//
	// Optional
	Name js.String
	// Version is "LinuxPackageInfo.version"
	//
	// Optional
	Version js.String
	// Summary is "LinuxPackageInfo.summary"
	//
	// Optional
	Summary js.String
	// Description is "LinuxPackageInfo.description"
	//
	// Optional
	Description js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LinuxPackageInfo with all fields set.
func (p LinuxPackageInfo) FromRef(ref js.Ref) LinuxPackageInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new LinuxPackageInfo in the application heap.
func (p LinuxPackageInfo) New() js.Ref {
	return bindings.LinuxPackageInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *LinuxPackageInfo) UpdateFrom(ref js.Ref) {
	bindings.LinuxPackageInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *LinuxPackageInfo) Update(ref js.Ref) {
	bindings.LinuxPackageInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *LinuxPackageInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.Version.Ref(),
		p.Summary.Ref(),
		p.Description.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Version = p.Version.FromRef(js.Undefined)
	p.Summary = p.Summary.FromRef(js.Undefined)
	p.Description = p.Description.FromRef(js.Undefined)
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

type GetPreferencesCallbackFunc func(this js.Ref, result *Preferences) js.Ref

func (fn GetPreferencesCallbackFunc) Register() js.Func[func(result *Preferences)] {
	return js.RegisterCallback[func(result *Preferences)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetPreferencesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Preferences
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

type GetPreferencesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *Preferences) js.Ref
	Arg T
}

func (cb *GetPreferencesCallback[T]) Register() js.Func[func(result *Preferences)] {
	return js.RegisterCallback[func(result *Preferences)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetPreferencesCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Preferences
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

type Preferences struct {
	// DriveEnabled is "Preferences.driveEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_DriveEnabled MUST be set to true to make this field effective.
	DriveEnabled bool
	// DriveSyncEnabledOnMeteredNetwork is "Preferences.driveSyncEnabledOnMeteredNetwork"
	//
	// Optional
	//
	// NOTE: FFI_USE_DriveSyncEnabledOnMeteredNetwork MUST be set to true to make this field effective.
	DriveSyncEnabledOnMeteredNetwork bool
	// SearchSuggestEnabled is "Preferences.searchSuggestEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_SearchSuggestEnabled MUST be set to true to make this field effective.
	SearchSuggestEnabled bool
	// Use24hourClock is "Preferences.use24hourClock"
	//
	// Optional
	//
	// NOTE: FFI_USE_Use24hourClock MUST be set to true to make this field effective.
	Use24hourClock bool
	// Timezone is "Preferences.timezone"
	//
	// Optional
	Timezone js.String
	// ArcEnabled is "Preferences.arcEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_ArcEnabled MUST be set to true to make this field effective.
	ArcEnabled bool
	// ArcRemovableMediaAccessEnabled is "Preferences.arcRemovableMediaAccessEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_ArcRemovableMediaAccessEnabled MUST be set to true to make this field effective.
	ArcRemovableMediaAccessEnabled bool
	// FolderShortcuts is "Preferences.folderShortcuts"
	//
	// Optional
	FolderShortcuts js.Array[js.String]
	// TrashEnabled is "Preferences.trashEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_TrashEnabled MUST be set to true to make this field effective.
	TrashEnabled bool
	// OfficeFileMovedOneDrive is "Preferences.officeFileMovedOneDrive"
	//
	// Optional
	//
	// NOTE: FFI_USE_OfficeFileMovedOneDrive MUST be set to true to make this field effective.
	OfficeFileMovedOneDrive float64
	// OfficeFileMovedGoogleDrive is "Preferences.officeFileMovedGoogleDrive"
	//
	// Optional
	//
	// NOTE: FFI_USE_OfficeFileMovedGoogleDrive MUST be set to true to make this field effective.
	OfficeFileMovedGoogleDrive float64
	// DriveFsBulkPinningEnabled is "Preferences.driveFsBulkPinningEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_DriveFsBulkPinningEnabled MUST be set to true to make this field effective.
	DriveFsBulkPinningEnabled bool

	FFI_USE_DriveEnabled                     bool // for DriveEnabled.
	FFI_USE_DriveSyncEnabledOnMeteredNetwork bool // for DriveSyncEnabledOnMeteredNetwork.
	FFI_USE_SearchSuggestEnabled             bool // for SearchSuggestEnabled.
	FFI_USE_Use24hourClock                   bool // for Use24hourClock.
	FFI_USE_ArcEnabled                       bool // for ArcEnabled.
	FFI_USE_ArcRemovableMediaAccessEnabled   bool // for ArcRemovableMediaAccessEnabled.
	FFI_USE_TrashEnabled                     bool // for TrashEnabled.
	FFI_USE_OfficeFileMovedOneDrive          bool // for OfficeFileMovedOneDrive.
	FFI_USE_OfficeFileMovedGoogleDrive       bool // for OfficeFileMovedGoogleDrive.
	FFI_USE_DriveFsBulkPinningEnabled        bool // for DriveFsBulkPinningEnabled.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Preferences with all fields set.
func (p Preferences) FromRef(ref js.Ref) Preferences {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Preferences in the application heap.
func (p Preferences) New() js.Ref {
	return bindings.PreferencesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Preferences) UpdateFrom(ref js.Ref) {
	bindings.PreferencesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Preferences) Update(ref js.Ref) {
	bindings.PreferencesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Preferences) FreeMembers(recursive bool) {
	js.Free(
		p.Timezone.Ref(),
		p.FolderShortcuts.Ref(),
	)
	p.Timezone = p.Timezone.FromRef(js.Undefined)
	p.FolderShortcuts = p.FolderShortcuts.FromRef(js.Undefined)
}

type GetProfilesCallbackFunc func(this js.Ref, profiles js.Array[ProfileInfo], runningProfile js.String, displayProfile js.String) js.Ref

func (fn GetProfilesCallbackFunc) Register() js.Func[func(profiles js.Array[ProfileInfo], runningProfile js.String, displayProfile js.String)] {
	return js.RegisterCallback[func(profiles js.Array[ProfileInfo], runningProfile js.String, displayProfile js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetProfilesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[ProfileInfo]{}.FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
		js.String{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetProfilesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, profiles js.Array[ProfileInfo], runningProfile js.String, displayProfile js.String) js.Ref
	Arg T
}

func (cb *GetProfilesCallback[T]) Register() js.Func[func(profiles js.Array[ProfileInfo], runningProfile js.String, displayProfile js.String)] {
	return js.RegisterCallback[func(profiles js.Array[ProfileInfo], runningProfile js.String, displayProfile js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetProfilesCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Array[ProfileInfo]{}.FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
		js.String{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ProfileInfo struct {
	// ProfileId is "ProfileInfo.profileId"
	//
	// Optional
	ProfileId js.String
	// DisplayName is "ProfileInfo.displayName"
	//
	// Optional
	DisplayName js.String
	// IsCurrentProfile is "ProfileInfo.isCurrentProfile"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsCurrentProfile MUST be set to true to make this field effective.
	IsCurrentProfile bool

	FFI_USE_IsCurrentProfile bool // for IsCurrentProfile.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProfileInfo with all fields set.
func (p ProfileInfo) FromRef(ref js.Ref) ProfileInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ProfileInfo in the application heap.
func (p ProfileInfo) New() js.Ref {
	return bindings.ProfileInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ProfileInfo) UpdateFrom(ref js.Ref) {
	bindings.ProfileInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ProfileInfo) Update(ref js.Ref) {
	bindings.ProfileInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ProfileInfo) FreeMembers(recursive bool) {
	js.Free(
		p.ProfileId.Ref(),
		p.DisplayName.Ref(),
	)
	p.ProfileId = p.ProfileId.FromRef(js.Undefined)
	p.DisplayName = p.DisplayName.FromRef(js.Undefined)
}

type GetProvidersCallbackFunc func(this js.Ref, extensions js.Array[Provider]) js.Ref

func (fn GetProvidersCallbackFunc) Register() js.Func[func(extensions js.Array[Provider])] {
	return js.RegisterCallback[func(extensions js.Array[Provider])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetProvidersCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[Provider]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetProvidersCallback[T any] struct {
	Fn  func(arg T, this js.Ref, extensions js.Array[Provider]) js.Ref
	Arg T
}

func (cb *GetProvidersCallback[T]) Register() js.Func[func(extensions js.Array[Provider])] {
	return js.RegisterCallback[func(extensions js.Array[Provider])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetProvidersCallback[T]) DispatchCallback(
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

		js.Array[Provider]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ProviderSource uint32

const (
	_ ProviderSource = iota

	ProviderSource_FILE
	ProviderSource_DEVICE
	ProviderSource_NETWORK
)

func (ProviderSource) FromRef(str js.Ref) ProviderSource {
	return ProviderSource(bindings.ConstOfProviderSource(str))
}

func (x ProviderSource) String() (string, bool) {
	switch x {
	case ProviderSource_FILE:
		return "file", true
	case ProviderSource_DEVICE:
		return "device", true
	case ProviderSource_NETWORK:
		return "network", true
	default:
		return "", false
	}
}

type Provider struct {
	// ProviderId is "Provider.providerId"
	//
	// Optional
	ProviderId js.String
	// IconSet is "Provider.iconSet"
	//
	// Optional
	//
	// NOTE: IconSet.FFI_USE MUST be set to true to get IconSet used.
	IconSet IconSet
	// Name is "Provider.name"
	//
	// Optional
	Name js.String
	// Configurable is "Provider.configurable"
	//
	// Optional
	//
	// NOTE: FFI_USE_Configurable MUST be set to true to make this field effective.
	Configurable bool
	// Watchable is "Provider.watchable"
	//
	// Optional
	//
	// NOTE: FFI_USE_Watchable MUST be set to true to make this field effective.
	Watchable bool
	// MultipleMounts is "Provider.multipleMounts"
	//
	// Optional
	//
	// NOTE: FFI_USE_MultipleMounts MUST be set to true to make this field effective.
	MultipleMounts bool
	// Source is "Provider.source"
	//
	// Optional
	Source ProviderSource

	FFI_USE_Configurable   bool // for Configurable.
	FFI_USE_Watchable      bool // for Watchable.
	FFI_USE_MultipleMounts bool // for MultipleMounts.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Provider with all fields set.
func (p Provider) FromRef(ref js.Ref) Provider {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Provider in the application heap.
func (p Provider) New() js.Ref {
	return bindings.ProviderJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Provider) UpdateFrom(ref js.Ref) {
	bindings.ProviderJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Provider) Update(ref js.Ref) {
	bindings.ProviderJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Provider) FreeMembers(recursive bool) {
	js.Free(
		p.ProviderId.Ref(),
		p.Name.Ref(),
	)
	p.ProviderId = p.ProviderId.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	if recursive {
		p.IconSet.FreeMembers(true)
	}
}

type GetRecentFilesCallbackFunc func(this js.Ref, entries js.Array[js.Object]) js.Ref

func (fn GetRecentFilesCallbackFunc) Register() js.Func[func(entries js.Array[js.Object])] {
	return js.RegisterCallback[func(entries js.Array[js.Object])](
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

		js.Array[js.Object]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetRecentFilesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[js.Object]) js.Ref
	Arg T
}

func (cb *GetRecentFilesCallback[T]) Register() js.Func[func(entries js.Array[js.Object])] {
	return js.RegisterCallback[func(entries js.Array[js.Object])](
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

		js.Array[js.Object]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetSizeStatsCallbackFunc func(this js.Ref, sizeStats *MountPointSizeStats) js.Ref

func (fn GetSizeStatsCallbackFunc) Register() js.Func[func(sizeStats *MountPointSizeStats)] {
	return js.RegisterCallback[func(sizeStats *MountPointSizeStats)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetSizeStatsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 MountPointSizeStats
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

type GetSizeStatsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, sizeStats *MountPointSizeStats) js.Ref
	Arg T
}

func (cb *GetSizeStatsCallback[T]) Register() js.Func[func(sizeStats *MountPointSizeStats)] {
	return js.RegisterCallback[func(sizeStats *MountPointSizeStats)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetSizeStatsCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 MountPointSizeStats
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

type MountPointSizeStats struct {
	// TotalSize is "MountPointSizeStats.totalSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalSize MUST be set to true to make this field effective.
	TotalSize float64
	// RemainingSize is "MountPointSizeStats.remainingSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_RemainingSize MUST be set to true to make this field effective.
	RemainingSize float64

	FFI_USE_TotalSize     bool // for TotalSize.
	FFI_USE_RemainingSize bool // for RemainingSize.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MountPointSizeStats with all fields set.
func (p MountPointSizeStats) FromRef(ref js.Ref) MountPointSizeStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MountPointSizeStats in the application heap.
func (p MountPointSizeStats) New() js.Ref {
	return bindings.MountPointSizeStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MountPointSizeStats) UpdateFrom(ref js.Ref) {
	bindings.MountPointSizeStatsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MountPointSizeStats) Update(ref js.Ref) {
	bindings.MountPointSizeStatsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MountPointSizeStats) FreeMembers(recursive bool) {
}

type GetStringsCallbackFunc func(this js.Ref, result js.Object) js.Ref

func (fn GetStringsCallbackFunc) Register() js.Func[func(result js.Object)] {
	return js.RegisterCallback[func(result js.Object)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetStringsCallbackFunc) DispatchCallback(
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

type GetStringsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result js.Object) js.Ref
	Arg T
}

func (cb *GetStringsCallback[T]) Register() js.Func[func(result js.Object)] {
	return js.RegisterCallback[func(result js.Object)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetStringsCallback[T]) DispatchCallback(
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

type GetVolumeMetadataListCallbackFunc func(this js.Ref, volumeMetadataList js.Array[VolumeMetadata]) js.Ref

func (fn GetVolumeMetadataListCallbackFunc) Register() js.Func[func(volumeMetadataList js.Array[VolumeMetadata])] {
	return js.RegisterCallback[func(volumeMetadataList js.Array[VolumeMetadata])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetVolumeMetadataListCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[VolumeMetadata]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetVolumeMetadataListCallback[T any] struct {
	Fn  func(arg T, this js.Ref, volumeMetadataList js.Array[VolumeMetadata]) js.Ref
	Arg T
}

func (cb *GetVolumeMetadataListCallback[T]) Register() js.Func[func(volumeMetadataList js.Array[VolumeMetadata])] {
	return js.RegisterCallback[func(volumeMetadataList js.Array[VolumeMetadata])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetVolumeMetadataListCallback[T]) DispatchCallback(
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

		js.Array[VolumeMetadata]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type Source uint32

const (
	_ Source = iota

	Source_FILE
	Source_DEVICE
	Source_NETWORK
	Source_SYSTEM
)

func (Source) FromRef(str js.Ref) Source {
	return Source(bindings.ConstOfSource(str))
}

func (x Source) String() (string, bool) {
	switch x {
	case Source_FILE:
		return "file", true
	case Source_DEVICE:
		return "device", true
	case Source_NETWORK:
		return "network", true
	case Source_SYSTEM:
		return "system", true
	default:
		return "", false
	}
}

type MountError uint32

const (
	_ MountError = iota

	MountError_SUCCESS
	MountError_IN_PROGRESS
	MountError_UNKNOWN_ERROR
	MountError_INTERNAL_ERROR
	MountError_INVALID_ARGUMENT
	MountError_INVALID_PATH
	MountError_PATH_ALREADY_MOUNTED
	MountError_PATH_NOT_MOUNTED
	MountError_DIRECTORY_CREATION_FAILED
	MountError_INVALID_MOUNT_OPTIONS
	MountError_INSUFFICIENT_PERMISSIONS
	MountError_MOUNT_PROGRAM_NOT_FOUND
	MountError_MOUNT_PROGRAM_FAILED
	MountError_INVALID_DEVICE_PATH
	MountError_UNKNOWN_FILESYSTEM
	MountError_UNSUPPORTED_FILESYSTEM
	MountError_NEED_PASSWORD
	MountError_CANCELLED
	MountError_BUSY
)

func (MountError) FromRef(str js.Ref) MountError {
	return MountError(bindings.ConstOfMountError(str))
}

func (x MountError) String() (string, bool) {
	switch x {
	case MountError_SUCCESS:
		return "success", true
	case MountError_IN_PROGRESS:
		return "in_progress", true
	case MountError_UNKNOWN_ERROR:
		return "unknown_error", true
	case MountError_INTERNAL_ERROR:
		return "internal_error", true
	case MountError_INVALID_ARGUMENT:
		return "invalid_argument", true
	case MountError_INVALID_PATH:
		return "invalid_path", true
	case MountError_PATH_ALREADY_MOUNTED:
		return "path_already_mounted", true
	case MountError_PATH_NOT_MOUNTED:
		return "path_not_mounted", true
	case MountError_DIRECTORY_CREATION_FAILED:
		return "directory_creation_failed", true
	case MountError_INVALID_MOUNT_OPTIONS:
		return "invalid_mount_options", true
	case MountError_INSUFFICIENT_PERMISSIONS:
		return "insufficient_permissions", true
	case MountError_MOUNT_PROGRAM_NOT_FOUND:
		return "mount_program_not_found", true
	case MountError_MOUNT_PROGRAM_FAILED:
		return "mount_program_failed", true
	case MountError_INVALID_DEVICE_PATH:
		return "invalid_device_path", true
	case MountError_UNKNOWN_FILESYSTEM:
		return "unknown_filesystem", true
	case MountError_UNSUPPORTED_FILESYSTEM:
		return "unsupported_filesystem", true
	case MountError_NEED_PASSWORD:
		return "need_password", true
	case MountError_CANCELLED:
		return "cancelled", true
	case MountError_BUSY:
		return "busy", true
	default:
		return "", false
	}
}

type MountContext uint32

const (
	_ MountContext = iota

	MountContext_USER
	MountContext_AUTO
)

func (MountContext) FromRef(str js.Ref) MountContext {
	return MountContext(bindings.ConstOfMountContext(str))
}

func (x MountContext) String() (string, bool) {
	switch x {
	case MountContext_USER:
		return "user", true
	case MountContext_AUTO:
		return "auto", true
	default:
		return "", false
	}
}

type VmType uint32

const (
	_ VmType = iota

	VmType_TERMINA
	VmType_PLUGIN_VM
	VmType_BOREALIS
	VmType_BRUSCHETTA
	VmType_ARCVM
)

func (VmType) FromRef(str js.Ref) VmType {
	return VmType(bindings.ConstOfVmType(str))
}

func (x VmType) String() (string, bool) {
	switch x {
	case VmType_TERMINA:
		return "termina", true
	case VmType_PLUGIN_VM:
		return "plugin_vm", true
	case VmType_BOREALIS:
		return "borealis", true
	case VmType_BRUSCHETTA:
		return "bruschetta", true
	case VmType_ARCVM:
		return "arcvm", true
	default:
		return "", false
	}
}

type VolumeMetadata struct {
	// VolumeId is "VolumeMetadata.volumeId"
	//
	// Optional
	VolumeId js.String
	// FileSystemId is "VolumeMetadata.fileSystemId"
	//
	// Optional
	FileSystemId js.String
	// ProviderId is "VolumeMetadata.providerId"
	//
	// Optional
	ProviderId js.String
	// Source is "VolumeMetadata.source"
	//
	// Optional
	Source Source
	// VolumeLabel is "VolumeMetadata.volumeLabel"
	//
	// Optional
	VolumeLabel js.String
	// Profile is "VolumeMetadata.profile"
	//
	// Optional
	//
	// NOTE: Profile.FFI_USE MUST be set to true to get Profile used.
	Profile ProfileInfo
	// SourcePath is "VolumeMetadata.sourcePath"
	//
	// Optional
	SourcePath js.String
	// VolumeType is "VolumeMetadata.volumeType"
	//
	// Optional
	VolumeType VolumeType
	// DeviceType is "VolumeMetadata.deviceType"
	//
	// Optional
	DeviceType DeviceType
	// DevicePath is "VolumeMetadata.devicePath"
	//
	// Optional
	DevicePath js.String
	// IsParentDevice is "VolumeMetadata.isParentDevice"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsParentDevice MUST be set to true to make this field effective.
	IsParentDevice bool
	// IsReadOnly is "VolumeMetadata.isReadOnly"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsReadOnly MUST be set to true to make this field effective.
	IsReadOnly bool
	// IsReadOnlyRemovableDevice is "VolumeMetadata.isReadOnlyRemovableDevice"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsReadOnlyRemovableDevice MUST be set to true to make this field effective.
	IsReadOnlyRemovableDevice bool
	// HasMedia is "VolumeMetadata.hasMedia"
	//
	// Optional
	//
	// NOTE: FFI_USE_HasMedia MUST be set to true to make this field effective.
	HasMedia bool
	// Configurable is "VolumeMetadata.configurable"
	//
	// Optional
	//
	// NOTE: FFI_USE_Configurable MUST be set to true to make this field effective.
	Configurable bool
	// Watchable is "VolumeMetadata.watchable"
	//
	// Optional
	//
	// NOTE: FFI_USE_Watchable MUST be set to true to make this field effective.
	Watchable bool
	// MountCondition is "VolumeMetadata.mountCondition"
	//
	// Optional
	MountCondition MountError
	// MountContext is "VolumeMetadata.mountContext"
	//
	// Optional
	MountContext MountContext
	// DiskFileSystemType is "VolumeMetadata.diskFileSystemType"
	//
	// Optional
	DiskFileSystemType js.String
	// IconSet is "VolumeMetadata.iconSet"
	//
	// Optional
	//
	// NOTE: IconSet.FFI_USE MUST be set to true to get IconSet used.
	IconSet IconSet
	// DriveLabel is "VolumeMetadata.driveLabel"
	//
	// Optional
	DriveLabel js.String
	// RemoteMountPath is "VolumeMetadata.remoteMountPath"
	//
	// Optional
	RemoteMountPath js.String
	// Hidden is "VolumeMetadata.hidden"
	//
	// Optional
	//
	// NOTE: FFI_USE_Hidden MUST be set to true to make this field effective.
	Hidden bool
	// VmType is "VolumeMetadata.vmType"
	//
	// Optional
	VmType VmType

	FFI_USE_IsParentDevice            bool // for IsParentDevice.
	FFI_USE_IsReadOnly                bool // for IsReadOnly.
	FFI_USE_IsReadOnlyRemovableDevice bool // for IsReadOnlyRemovableDevice.
	FFI_USE_HasMedia                  bool // for HasMedia.
	FFI_USE_Configurable              bool // for Configurable.
	FFI_USE_Watchable                 bool // for Watchable.
	FFI_USE_Hidden                    bool // for Hidden.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VolumeMetadata with all fields set.
func (p VolumeMetadata) FromRef(ref js.Ref) VolumeMetadata {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VolumeMetadata in the application heap.
func (p VolumeMetadata) New() js.Ref {
	return bindings.VolumeMetadataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *VolumeMetadata) UpdateFrom(ref js.Ref) {
	bindings.VolumeMetadataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *VolumeMetadata) Update(ref js.Ref) {
	bindings.VolumeMetadataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *VolumeMetadata) FreeMembers(recursive bool) {
	js.Free(
		p.VolumeId.Ref(),
		p.FileSystemId.Ref(),
		p.ProviderId.Ref(),
		p.VolumeLabel.Ref(),
		p.SourcePath.Ref(),
		p.DevicePath.Ref(),
		p.DiskFileSystemType.Ref(),
		p.DriveLabel.Ref(),
		p.RemoteMountPath.Ref(),
	)
	p.VolumeId = p.VolumeId.FromRef(js.Undefined)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
	p.ProviderId = p.ProviderId.FromRef(js.Undefined)
	p.VolumeLabel = p.VolumeLabel.FromRef(js.Undefined)
	p.SourcePath = p.SourcePath.FromRef(js.Undefined)
	p.DevicePath = p.DevicePath.FromRef(js.Undefined)
	p.DiskFileSystemType = p.DiskFileSystemType.FromRef(js.Undefined)
	p.DriveLabel = p.DriveLabel.FromRef(js.Undefined)
	p.RemoteMountPath = p.RemoteMountPath.FromRef(js.Undefined)
	if recursive {
		p.Profile.FreeMembers(true)
		p.IconSet.FreeMembers(true)
	}
}

type GetVolumeRootCallbackFunc func(this js.Ref, rootDir js.Object) js.Ref

func (fn GetVolumeRootCallbackFunc) Register() js.Func[func(rootDir js.Object)] {
	return js.RegisterCallback[func(rootDir js.Object)](
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

	if ctx.Return(fn(
		args[0],

		js.Object{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetVolumeRootCallback[T any] struct {
	Fn  func(arg T, this js.Ref, rootDir js.Object) js.Ref
	Arg T
}

func (cb *GetVolumeRootCallback[T]) Register() js.Func[func(rootDir js.Object)] {
	return js.RegisterCallback[func(rootDir js.Object)](
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

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Object{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetVolumeRootOptions struct {
	// VolumeId is "GetVolumeRootOptions.volumeId"
	//
	// Optional
	VolumeId js.String
	// Writable is "GetVolumeRootOptions.writable"
	//
	// Optional
	//
	// NOTE: FFI_USE_Writable MUST be set to true to make this field effective.
	Writable bool

	FFI_USE_Writable bool // for Writable.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetVolumeRootOptions with all fields set.
func (p GetVolumeRootOptions) FromRef(ref js.Ref) GetVolumeRootOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetVolumeRootOptions in the application heap.
func (p GetVolumeRootOptions) New() js.Ref {
	return bindings.GetVolumeRootOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetVolumeRootOptions) UpdateFrom(ref js.Ref) {
	bindings.GetVolumeRootOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetVolumeRootOptions) Update(ref js.Ref) {
	bindings.GetVolumeRootOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetVolumeRootOptions) FreeMembers(recursive bool) {
	js.Free(
		p.VolumeId.Ref(),
	)
	p.VolumeId = p.VolumeId.FromRef(js.Undefined)
}

type HoldingSpaceState struct {
	// ItemUrls is "HoldingSpaceState.itemUrls"
	//
	// Optional
	ItemUrls js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HoldingSpaceState with all fields set.
func (p HoldingSpaceState) FromRef(ref js.Ref) HoldingSpaceState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HoldingSpaceState in the application heap.
func (p HoldingSpaceState) New() js.Ref {
	return bindings.HoldingSpaceStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HoldingSpaceState) UpdateFrom(ref js.Ref) {
	bindings.HoldingSpaceStateJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HoldingSpaceState) Update(ref js.Ref) {
	bindings.HoldingSpaceStateJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HoldingSpaceState) FreeMembers(recursive bool) {
	js.Free(
		p.ItemUrls.Ref(),
	)
	p.ItemUrls = p.ItemUrls.FromRef(js.Undefined)
}

type HoldingSpaceStateCallbackFunc func(this js.Ref, state *HoldingSpaceState) js.Ref

func (fn HoldingSpaceStateCallbackFunc) Register() js.Func[func(state *HoldingSpaceState)] {
	return js.RegisterCallback[func(state *HoldingSpaceState)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn HoldingSpaceStateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 HoldingSpaceState
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

type HoldingSpaceStateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, state *HoldingSpaceState) js.Ref
	Arg T
}

func (cb *HoldingSpaceStateCallback[T]) Register() js.Func[func(state *HoldingSpaceState)] {
	return js.RegisterCallback[func(state *HoldingSpaceState)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *HoldingSpaceStateCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 HoldingSpaceState
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
	// DestinationFolder is "IOTaskParams.destinationFolder"
	//
	// Optional
	DestinationFolder js.Object
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
		p.DestinationFolder.Ref(),
		p.Password.Ref(),
	)
	p.DestinationFolder = p.DestinationFolder.FromRef(js.Undefined)
	p.Password = p.Password.FromRef(js.Undefined)
}

type IOTaskState uint32

const (
	_ IOTaskState = iota

	IOTaskState_QUEUED
	IOTaskState_SCANNING
	IOTaskState_IN_PROGRESS
	IOTaskState_PAUSED
	IOTaskState_SUCCESS
	IOTaskState_ERROR
	IOTaskState_NEED_PASSWORD
	IOTaskState_CANCELLED
)

func (IOTaskState) FromRef(str js.Ref) IOTaskState {
	return IOTaskState(bindings.ConstOfIOTaskState(str))
}

func (x IOTaskState) String() (string, bool) {
	switch x {
	case IOTaskState_QUEUED:
		return "queued", true
	case IOTaskState_SCANNING:
		return "scanning", true
	case IOTaskState_IN_PROGRESS:
		return "in_progress", true
	case IOTaskState_PAUSED:
		return "paused", true
	case IOTaskState_SUCCESS:
		return "success", true
	case IOTaskState_ERROR:
		return "error", true
	case IOTaskState_NEED_PASSWORD:
		return "need_password", true
	case IOTaskState_CANCELLED:
		return "cancelled", true
	default:
		return "", false
	}
}

type IOTaskType uint32

const (
	_ IOTaskType = iota

	IOTaskType_COPY
	IOTaskType_DELETE
	IOTaskType_EMPTY_TRASH
	IOTaskType_EXTRACT
	IOTaskType_MOVE
	IOTaskType_RESTORE
	IOTaskType_RESTORE_TO_DESTINATION
	IOTaskType_TRASH
	IOTaskType_ZIP
)

func (IOTaskType) FromRef(str js.Ref) IOTaskType {
	return IOTaskType(bindings.ConstOfIOTaskType(str))
}

func (x IOTaskType) String() (string, bool) {
	switch x {
	case IOTaskType_COPY:
		return "copy", true
	case IOTaskType_DELETE:
		return "delete", true
	case IOTaskType_EMPTY_TRASH:
		return "empty_trash", true
	case IOTaskType_EXTRACT:
		return "extract", true
	case IOTaskType_MOVE:
		return "move", true
	case IOTaskType_RESTORE:
		return "restore", true
	case IOTaskType_RESTORE_TO_DESTINATION:
		return "restore_to_destination", true
	case IOTaskType_TRASH:
		return "trash", true
	case IOTaskType_ZIP:
		return "zip", true
	default:
		return "", false
	}
}

type InspectionType uint32

const (
	_ InspectionType = iota

	InspectionType_NORMAL
	InspectionType_CONSOLE
	InspectionType_ELEMENT
	InspectionType_BACKGROUND
)

func (InspectionType) FromRef(str js.Ref) InspectionType {
	return InspectionType(bindings.ConstOfInspectionType(str))
}

func (x InspectionType) String() (string, bool) {
	switch x {
	case InspectionType_NORMAL:
		return "normal", true
	case InspectionType_CONSOLE:
		return "console", true
	case InspectionType_ELEMENT:
		return "element", true
	case InspectionType_BACKGROUND:
		return "background", true
	default:
		return "", false
	}
}

type InstallLinuxPackageResponse uint32

const (
	_ InstallLinuxPackageResponse = iota

	InstallLinuxPackageResponse_STARTED
	InstallLinuxPackageResponse_FAILED
	InstallLinuxPackageResponse_INSTALL_ALREADY_ACTIVE
)

func (InstallLinuxPackageResponse) FromRef(str js.Ref) InstallLinuxPackageResponse {
	return InstallLinuxPackageResponse(bindings.ConstOfInstallLinuxPackageResponse(str))
}

func (x InstallLinuxPackageResponse) String() (string, bool) {
	switch x {
	case InstallLinuxPackageResponse_STARTED:
		return "started", true
	case InstallLinuxPackageResponse_FAILED:
		return "failed", true
	case InstallLinuxPackageResponse_INSTALL_ALREADY_ACTIVE:
		return "install_already_active", true
	default:
		return "", false
	}
}

type InstallLinuxPackageCallbackFunc func(this js.Ref, response InstallLinuxPackageResponse, failure_reason js.String) js.Ref

func (fn InstallLinuxPackageCallbackFunc) Register() js.Func[func(response InstallLinuxPackageResponse, failure_reason js.String)] {
	return js.RegisterCallback[func(response InstallLinuxPackageResponse, failure_reason js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn InstallLinuxPackageCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		InstallLinuxPackageResponse(0).FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type InstallLinuxPackageCallback[T any] struct {
	Fn  func(arg T, this js.Ref, response InstallLinuxPackageResponse, failure_reason js.String) js.Ref
	Arg T
}

func (cb *InstallLinuxPackageCallback[T]) Register() js.Func[func(response InstallLinuxPackageResponse, failure_reason js.String)] {
	return js.RegisterCallback[func(response InstallLinuxPackageResponse, failure_reason js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *InstallLinuxPackageCallback[T]) DispatchCallback(
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

		InstallLinuxPackageResponse(0).FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ListMountableGuestsCallbackFunc func(this js.Ref, guest js.Array[MountableGuest]) js.Ref

func (fn ListMountableGuestsCallbackFunc) Register() js.Func[func(guest js.Array[MountableGuest])] {
	return js.RegisterCallback[func(guest js.Array[MountableGuest])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ListMountableGuestsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[MountableGuest]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ListMountableGuestsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, guest js.Array[MountableGuest]) js.Ref
	Arg T
}

func (cb *ListMountableGuestsCallback[T]) Register() js.Func[func(guest js.Array[MountableGuest])] {
	return js.RegisterCallback[func(guest js.Array[MountableGuest])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ListMountableGuestsCallback[T]) DispatchCallback(
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

		js.Array[MountableGuest]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type MountableGuest struct {
	// Id is "MountableGuest.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id int32
	// DisplayName is "MountableGuest.displayName"
	//
	// Optional
	DisplayName js.String
	// VmType is "MountableGuest.vmType"
	//
	// Optional
	VmType VmType

	FFI_USE_Id bool // for Id.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MountableGuest with all fields set.
func (p MountableGuest) FromRef(ref js.Ref) MountableGuest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MountableGuest in the application heap.
func (p MountableGuest) New() js.Ref {
	return bindings.MountableGuestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MountableGuest) UpdateFrom(ref js.Ref) {
	bindings.MountableGuestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MountableGuest) Update(ref js.Ref) {
	bindings.MountableGuestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MountableGuest) FreeMembers(recursive bool) {
	js.Free(
		p.DisplayName.Ref(),
	)
	p.DisplayName = p.DisplayName.FromRef(js.Undefined)
}

type MountCompletedEventType uint32

const (
	_ MountCompletedEventType = iota

	MountCompletedEventType_MOUNT
	MountCompletedEventType_UNMOUNT
)

func (MountCompletedEventType) FromRef(str js.Ref) MountCompletedEventType {
	return MountCompletedEventType(bindings.ConstOfMountCompletedEventType(str))
}

func (x MountCompletedEventType) String() (string, bool) {
	switch x {
	case MountCompletedEventType_MOUNT:
		return "mount", true
	case MountCompletedEventType_UNMOUNT:
		return "unmount", true
	default:
		return "", false
	}
}

type MountCompletedEvent struct {
	// EventType is "MountCompletedEvent.eventType"
	//
	// Optional
	EventType MountCompletedEventType
	// Status is "MountCompletedEvent.status"
	//
	// Optional
	Status MountError
	// VolumeMetadata is "MountCompletedEvent.volumeMetadata"
	//
	// Optional
	//
	// NOTE: VolumeMetadata.FFI_USE MUST be set to true to get VolumeMetadata used.
	VolumeMetadata VolumeMetadata
	// ShouldNotify is "MountCompletedEvent.shouldNotify"
	//
	// Optional
	//
	// NOTE: FFI_USE_ShouldNotify MUST be set to true to make this field effective.
	ShouldNotify bool

	FFI_USE_ShouldNotify bool // for ShouldNotify.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MountCompletedEvent with all fields set.
func (p MountCompletedEvent) FromRef(ref js.Ref) MountCompletedEvent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MountCompletedEvent in the application heap.
func (p MountCompletedEvent) New() js.Ref {
	return bindings.MountCompletedEventJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MountCompletedEvent) UpdateFrom(ref js.Ref) {
	bindings.MountCompletedEventJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MountCompletedEvent) Update(ref js.Ref) {
	bindings.MountCompletedEventJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MountCompletedEvent) FreeMembers(recursive bool) {
	if recursive {
		p.VolumeMetadata.FreeMembers(true)
	}
}

type OpenWindowParams struct {
	// CurrentDirectoryURL is "OpenWindowParams.currentDirectoryURL"
	//
	// Optional
	CurrentDirectoryURL js.String
	// SelectionURL is "OpenWindowParams.selectionURL"
	//
	// Optional
	SelectionURL js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OpenWindowParams with all fields set.
func (p OpenWindowParams) FromRef(ref js.Ref) OpenWindowParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OpenWindowParams in the application heap.
func (p OpenWindowParams) New() js.Ref {
	return bindings.OpenWindowParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OpenWindowParams) UpdateFrom(ref js.Ref) {
	bindings.OpenWindowParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OpenWindowParams) Update(ref js.Ref) {
	bindings.OpenWindowParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OpenWindowParams) FreeMembers(recursive bool) {
	js.Free(
		p.CurrentDirectoryURL.Ref(),
		p.SelectionURL.Ref(),
	)
	p.CurrentDirectoryURL = p.CurrentDirectoryURL.FromRef(js.Undefined)
	p.SelectionURL = p.SelectionURL.FromRef(js.Undefined)
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
	RestoreEntry js.Object
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
		p.RestoreEntry.Ref(),
		p.TrashInfoFileName.Ref(),
	)
	p.RestoreEntry = p.RestoreEntry.FromRef(js.Undefined)
	p.TrashInfoFileName = p.TrashInfoFileName.FromRef(js.Undefined)
}

type PolicyErrorType uint32

const (
	_ PolicyErrorType = iota

	PolicyErrorType_DLP
	PolicyErrorType_ENTERPRISE_CONNECTORS
	PolicyErrorType_DLP_WARNING_TIMEOUT
)

func (PolicyErrorType) FromRef(str js.Ref) PolicyErrorType {
	return PolicyErrorType(bindings.ConstOfPolicyErrorType(str))
}

func (x PolicyErrorType) String() (string, bool) {
	switch x {
	case PolicyErrorType_DLP:
		return "dlp", true
	case PolicyErrorType_ENTERPRISE_CONNECTORS:
		return "enterprise_connectors", true
	case PolicyErrorType_DLP_WARNING_TIMEOUT:
		return "dlp_warning_timeout", true
	default:
		return "", false
	}
}

type PolicyPauseParams struct {
	// Type is "PolicyPauseParams.type"
	//
	// Optional
	Type PolicyErrorType
	// PolicyFileCount is "PolicyPauseParams.policyFileCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_PolicyFileCount MUST be set to true to make this field effective.
	PolicyFileCount int32
	// FileName is "PolicyPauseParams.fileName"
	//
	// Optional
	FileName js.String

	FFI_USE_PolicyFileCount bool // for PolicyFileCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PolicyPauseParams with all fields set.
func (p PolicyPauseParams) FromRef(ref js.Ref) PolicyPauseParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PolicyPauseParams in the application heap.
func (p PolicyPauseParams) New() js.Ref {
	return bindings.PolicyPauseParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PolicyPauseParams) UpdateFrom(ref js.Ref) {
	bindings.PolicyPauseParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PolicyPauseParams) Update(ref js.Ref) {
	bindings.PolicyPauseParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PolicyPauseParams) FreeMembers(recursive bool) {
	js.Free(
		p.FileName.Ref(),
	)
	p.FileName = p.FileName.FromRef(js.Undefined)
}

type PauseParams struct {
	// ConflictParams is "PauseParams.conflictParams"
	//
	// Optional
	//
	// NOTE: ConflictParams.FFI_USE MUST be set to true to get ConflictParams used.
	ConflictParams ConflictPauseParams
	// PolicyParams is "PauseParams.policyParams"
	//
	// Optional
	//
	// NOTE: PolicyParams.FFI_USE MUST be set to true to get PolicyParams used.
	PolicyParams PolicyPauseParams

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PauseParams with all fields set.
func (p PauseParams) FromRef(ref js.Ref) PauseParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PauseParams in the application heap.
func (p PauseParams) New() js.Ref {
	return bindings.PauseParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PauseParams) UpdateFrom(ref js.Ref) {
	bindings.PauseParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PauseParams) Update(ref js.Ref) {
	bindings.PauseParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PauseParams) FreeMembers(recursive bool) {
	if recursive {
		p.ConflictParams.FreeMembers(true)
		p.PolicyParams.FreeMembers(true)
	}
}

type PolicyDialogType uint32

const (
	_ PolicyDialogType = iota

	PolicyDialogType_WARNING
	PolicyDialogType_ERROR
)

func (PolicyDialogType) FromRef(str js.Ref) PolicyDialogType {
	return PolicyDialogType(bindings.ConstOfPolicyDialogType(str))
}

func (x PolicyDialogType) String() (string, bool) {
	switch x {
	case PolicyDialogType_WARNING:
		return "warning", true
	case PolicyDialogType_ERROR:
		return "error", true
	default:
		return "", false
	}
}

type PolicyError struct {
	// Type is "PolicyError.type"
	//
	// Optional
	Type PolicyErrorType
	// PolicyFileCount is "PolicyError.policyFileCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_PolicyFileCount MUST be set to true to make this field effective.
	PolicyFileCount int32
	// FileName is "PolicyError.fileName"
	//
	// Optional
	FileName js.String

	FFI_USE_PolicyFileCount bool // for PolicyFileCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PolicyError with all fields set.
func (p PolicyError) FromRef(ref js.Ref) PolicyError {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PolicyError in the application heap.
func (p PolicyError) New() js.Ref {
	return bindings.PolicyErrorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PolicyError) UpdateFrom(ref js.Ref) {
	bindings.PolicyErrorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PolicyError) Update(ref js.Ref) {
	bindings.PolicyErrorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PolicyError) FreeMembers(recursive bool) {
	js.Free(
		p.FileName.Ref(),
	)
	p.FileName = p.FileName.FromRef(js.Undefined)
}

type PolicyResumeParams struct {
	// Type is "PolicyResumeParams.type"
	//
	// Optional
	Type PolicyErrorType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PolicyResumeParams with all fields set.
func (p PolicyResumeParams) FromRef(ref js.Ref) PolicyResumeParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PolicyResumeParams in the application heap.
func (p PolicyResumeParams) New() js.Ref {
	return bindings.PolicyResumeParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PolicyResumeParams) UpdateFrom(ref js.Ref) {
	bindings.PolicyResumeParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PolicyResumeParams) Update(ref js.Ref) {
	bindings.PolicyResumeParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PolicyResumeParams) FreeMembers(recursive bool) {
}

type PreferencesChange struct {
	// DriveSyncEnabledOnMeteredNetwork is "PreferencesChange.driveSyncEnabledOnMeteredNetwork"
	//
	// Optional
	//
	// NOTE: FFI_USE_DriveSyncEnabledOnMeteredNetwork MUST be set to true to make this field effective.
	DriveSyncEnabledOnMeteredNetwork bool
	// ArcEnabled is "PreferencesChange.arcEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_ArcEnabled MUST be set to true to make this field effective.
	ArcEnabled bool
	// ArcRemovableMediaAccessEnabled is "PreferencesChange.arcRemovableMediaAccessEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_ArcRemovableMediaAccessEnabled MUST be set to true to make this field effective.
	ArcRemovableMediaAccessEnabled bool
	// FolderShortcuts is "PreferencesChange.folderShortcuts"
	//
	// Optional
	FolderShortcuts js.Array[js.String]
	// DriveFsBulkPinningEnabled is "PreferencesChange.driveFsBulkPinningEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_DriveFsBulkPinningEnabled MUST be set to true to make this field effective.
	DriveFsBulkPinningEnabled bool

	FFI_USE_DriveSyncEnabledOnMeteredNetwork bool // for DriveSyncEnabledOnMeteredNetwork.
	FFI_USE_ArcEnabled                       bool // for ArcEnabled.
	FFI_USE_ArcRemovableMediaAccessEnabled   bool // for ArcRemovableMediaAccessEnabled.
	FFI_USE_DriveFsBulkPinningEnabled        bool // for DriveFsBulkPinningEnabled.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PreferencesChange with all fields set.
func (p PreferencesChange) FromRef(ref js.Ref) PreferencesChange {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PreferencesChange in the application heap.
func (p PreferencesChange) New() js.Ref {
	return bindings.PreferencesChangeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PreferencesChange) UpdateFrom(ref js.Ref) {
	bindings.PreferencesChangeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PreferencesChange) Update(ref js.Ref) {
	bindings.PreferencesChangeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PreferencesChange) FreeMembers(recursive bool) {
	js.Free(
		p.FolderShortcuts.Ref(),
	)
	p.FolderShortcuts = p.FolderShortcuts.FromRef(js.Undefined)
}

type ProgressStatus struct {
	// Type is "ProgressStatus.type"
	//
	// Optional
	Type IOTaskType
	// State is "ProgressStatus.state"
	//
	// Optional
	State IOTaskState
	// PolicyError is "ProgressStatus.policyError"
	//
	// Optional
	//
	// NOTE: PolicyError.FFI_USE MUST be set to true to get PolicyError used.
	PolicyError PolicyError
	// SourceName is "ProgressStatus.sourceName"
	//
	// Optional
	SourceName js.String
	// NumRemainingItems is "ProgressStatus.numRemainingItems"
	//
	// Optional
	//
	// NOTE: FFI_USE_NumRemainingItems MUST be set to true to make this field effective.
	NumRemainingItems int32
	// ItemCount is "ProgressStatus.itemCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_ItemCount MUST be set to true to make this field effective.
	ItemCount int32
	// DestinationName is "ProgressStatus.destinationName"
	//
	// Optional
	DestinationName js.String
	// BytesTransferred is "ProgressStatus.bytesTransferred"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesTransferred MUST be set to true to make this field effective.
	BytesTransferred int32
	// TotalBytes is "ProgressStatus.totalBytes"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalBytes MUST be set to true to make this field effective.
	TotalBytes int32
	// TaskId is "ProgressStatus.taskId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TaskId MUST be set to true to make this field effective.
	TaskId int32
	// RemainingSeconds is "ProgressStatus.remainingSeconds"
	//
	// Optional
	//
	// NOTE: FFI_USE_RemainingSeconds MUST be set to true to make this field effective.
	RemainingSeconds float64
	// SourcesScanned is "ProgressStatus.sourcesScanned"
	//
	// Optional
	//
	// NOTE: FFI_USE_SourcesScanned MUST be set to true to make this field effective.
	SourcesScanned int32
	// ShowNotification is "ProgressStatus.showNotification"
	//
	// Optional
	//
	// NOTE: FFI_USE_ShowNotification MUST be set to true to make this field effective.
	ShowNotification bool
	// ErrorName is "ProgressStatus.errorName"
	//
	// Optional
	ErrorName js.String
	// PauseParams is "ProgressStatus.pauseParams"
	//
	// Optional
	//
	// NOTE: PauseParams.FFI_USE MUST be set to true to get PauseParams used.
	PauseParams PauseParams
	// Outputs is "ProgressStatus.outputs"
	//
	// Optional
	Outputs js.Array[js.Object]
	// DestinationVolumeId is "ProgressStatus.destinationVolumeId"
	//
	// Optional
	DestinationVolumeId js.String

	FFI_USE_NumRemainingItems bool // for NumRemainingItems.
	FFI_USE_ItemCount         bool // for ItemCount.
	FFI_USE_BytesTransferred  bool // for BytesTransferred.
	FFI_USE_TotalBytes        bool // for TotalBytes.
	FFI_USE_TaskId            bool // for TaskId.
	FFI_USE_RemainingSeconds  bool // for RemainingSeconds.
	FFI_USE_SourcesScanned    bool // for SourcesScanned.
	FFI_USE_ShowNotification  bool // for ShowNotification.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProgressStatus with all fields set.
func (p ProgressStatus) FromRef(ref js.Ref) ProgressStatus {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ProgressStatus in the application heap.
func (p ProgressStatus) New() js.Ref {
	return bindings.ProgressStatusJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ProgressStatus) UpdateFrom(ref js.Ref) {
	bindings.ProgressStatusJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ProgressStatus) Update(ref js.Ref) {
	bindings.ProgressStatusJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ProgressStatus) FreeMembers(recursive bool) {
	js.Free(
		p.SourceName.Ref(),
		p.DestinationName.Ref(),
		p.ErrorName.Ref(),
		p.Outputs.Ref(),
		p.DestinationVolumeId.Ref(),
	)
	p.SourceName = p.SourceName.FromRef(js.Undefined)
	p.DestinationName = p.DestinationName.FromRef(js.Undefined)
	p.ErrorName = p.ErrorName.FromRef(js.Undefined)
	p.Outputs = p.Outputs.FromRef(js.Undefined)
	p.DestinationVolumeId = p.DestinationVolumeId.FromRef(js.Undefined)
	if recursive {
		p.PolicyError.FreeMembers(true)
		p.PauseParams.FreeMembers(true)
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

type ResolveEntriesCallbackFunc func(this js.Ref, entries js.Array[js.Object]) js.Ref

func (fn ResolveEntriesCallbackFunc) Register() js.Func[func(entries js.Array[js.Object])] {
	return js.RegisterCallback[func(entries js.Array[js.Object])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ResolveEntriesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[js.Object]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ResolveEntriesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[js.Object]) js.Ref
	Arg T
}

func (cb *ResolveEntriesCallback[T]) Register() js.Func[func(entries js.Array[js.Object])] {
	return js.RegisterCallback[func(entries js.Array[js.Object])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ResolveEntriesCallback[T]) DispatchCallback(
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

		js.Array[js.Object]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ResumeParams struct {
	// ConflictParams is "ResumeParams.conflictParams"
	//
	// Optional
	//
	// NOTE: ConflictParams.FFI_USE MUST be set to true to get ConflictParams used.
	ConflictParams ConflictResumeParams
	// PolicyParams is "ResumeParams.policyParams"
	//
	// Optional
	//
	// NOTE: PolicyParams.FFI_USE MUST be set to true to get PolicyParams used.
	PolicyParams PolicyResumeParams

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ResumeParams with all fields set.
func (p ResumeParams) FromRef(ref js.Ref) ResumeParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ResumeParams in the application heap.
func (p ResumeParams) New() js.Ref {
	return bindings.ResumeParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ResumeParams) UpdateFrom(ref js.Ref) {
	bindings.ResumeParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ResumeParams) Update(ref js.Ref) {
	bindings.ResumeParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ResumeParams) FreeMembers(recursive bool) {
	if recursive {
		p.ConflictParams.FreeMembers(true)
		p.PolicyParams.FreeMembers(true)
	}
}

type SearchDriveCallbackFunc func(this js.Ref, entries js.Array[js.Object], nextFeed js.String) js.Ref

func (fn SearchDriveCallbackFunc) Register() js.Func[func(entries js.Array[js.Object], nextFeed js.String)] {
	return js.RegisterCallback[func(entries js.Array[js.Object], nextFeed js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SearchDriveCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[js.Object]{}.FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SearchDriveCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[js.Object], nextFeed js.String) js.Ref
	Arg T
}

func (cb *SearchDriveCallback[T]) Register() js.Func[func(entries js.Array[js.Object], nextFeed js.String)] {
	return js.RegisterCallback[func(entries js.Array[js.Object], nextFeed js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SearchDriveCallback[T]) DispatchCallback(
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

		js.Array[js.Object]{}.FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SearchDriveMetadataCallbackFunc func(this js.Ref, results js.Array[DriveMetadataSearchResult]) js.Ref

func (fn SearchDriveMetadataCallbackFunc) Register() js.Func[func(results js.Array[DriveMetadataSearchResult])] {
	return js.RegisterCallback[func(results js.Array[DriveMetadataSearchResult])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SearchDriveMetadataCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[DriveMetadataSearchResult]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SearchDriveMetadataCallback[T any] struct {
	Fn  func(arg T, this js.Ref, results js.Array[DriveMetadataSearchResult]) js.Ref
	Arg T
}

func (cb *SearchDriveMetadataCallback[T]) Register() js.Func[func(results js.Array[DriveMetadataSearchResult])] {
	return js.RegisterCallback[func(results js.Array[DriveMetadataSearchResult])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SearchDriveMetadataCallback[T]) DispatchCallback(
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

		js.Array[DriveMetadataSearchResult]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SearchFilesByHashesCallbackFunc func(this js.Ref, paths js.Object) js.Ref

func (fn SearchFilesByHashesCallbackFunc) Register() js.Func[func(paths js.Object)] {
	return js.RegisterCallback[func(paths js.Object)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SearchFilesByHashesCallbackFunc) DispatchCallback(
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

type SearchFilesByHashesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, paths js.Object) js.Ref
	Arg T
}

func (cb *SearchFilesByHashesCallback[T]) Register() js.Func[func(paths js.Object)] {
	return js.RegisterCallback[func(paths js.Object)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SearchFilesByHashesCallback[T]) DispatchCallback(
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

type SearchFilesCallbackFunc func(this js.Ref, entries js.Array[js.Object]) js.Ref

func (fn SearchFilesCallbackFunc) Register() js.Func[func(entries js.Array[js.Object])] {
	return js.RegisterCallback[func(entries js.Array[js.Object])](
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

		js.Array[js.Object]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SearchFilesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[js.Object]) js.Ref
	Arg T
}

func (cb *SearchFilesCallback[T]) Register() js.Func[func(entries js.Array[js.Object])] {
	return js.RegisterCallback[func(entries js.Array[js.Object])](
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

		js.Array[js.Object]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SearchType uint32

const (
	_ SearchType = iota

	SearchType_EXCLUDE_DIRECTORIES
	SearchType_SHARED_WITH_ME
	SearchType_OFFLINE
	SearchType_ALL
)

func (SearchType) FromRef(str js.Ref) SearchType {
	return SearchType(bindings.ConstOfSearchType(str))
}

func (x SearchType) String() (string, bool) {
	switch x {
	case SearchType_EXCLUDE_DIRECTORIES:
		return "EXCLUDE_DIRECTORIES", true
	case SearchType_SHARED_WITH_ME:
		return "SHARED_WITH_ME", true
	case SearchType_OFFLINE:
		return "OFFLINE", true
	case SearchType_ALL:
		return "ALL", true
	default:
		return "", false
	}
}

type SearchMetadataParams struct {
	// RootDir is "SearchMetadataParams.rootDir"
	//
	// Optional
	RootDir js.Object
	// Query is "SearchMetadataParams.query"
	//
	// Optional
	Query js.String
	// Types is "SearchMetadataParams.types"
	//
	// Optional
	Types SearchType
	// MaxResults is "SearchMetadataParams.maxResults"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxResults MUST be set to true to make this field effective.
	MaxResults int32
	// ModifiedTimestamp is "SearchMetadataParams.modifiedTimestamp"
	//
	// Optional
	//
	// NOTE: FFI_USE_ModifiedTimestamp MUST be set to true to make this field effective.
	ModifiedTimestamp float64
	// Category is "SearchMetadataParams.category"
	//
	// Optional
	Category FileCategory

	FFI_USE_MaxResults        bool // for MaxResults.
	FFI_USE_ModifiedTimestamp bool // for ModifiedTimestamp.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SearchMetadataParams with all fields set.
func (p SearchMetadataParams) FromRef(ref js.Ref) SearchMetadataParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SearchMetadataParams in the application heap.
func (p SearchMetadataParams) New() js.Ref {
	return bindings.SearchMetadataParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SearchMetadataParams) UpdateFrom(ref js.Ref) {
	bindings.SearchMetadataParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SearchMetadataParams) Update(ref js.Ref) {
	bindings.SearchMetadataParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SearchMetadataParams) FreeMembers(recursive bool) {
	js.Free(
		p.RootDir.Ref(),
		p.Query.Ref(),
	)
	p.RootDir = p.RootDir.FromRef(js.Undefined)
	p.Query = p.Query.FromRef(js.Undefined)
}

type SearchParams struct {
	// Query is "SearchParams.query"
	//
	// Optional
	Query js.String
	// Category is "SearchParams.category"
	//
	// Optional
	Category FileCategory
	// ModifiedTimestamp is "SearchParams.modifiedTimestamp"
	//
	// Optional
	//
	// NOTE: FFI_USE_ModifiedTimestamp MUST be set to true to make this field effective.
	ModifiedTimestamp float64
	// NextFeed is "SearchParams.nextFeed"
	//
	// Optional
	NextFeed js.String

	FFI_USE_ModifiedTimestamp bool // for ModifiedTimestamp.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SearchParams with all fields set.
func (p SearchParams) FromRef(ref js.Ref) SearchParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SearchParams in the application heap.
func (p SearchParams) New() js.Ref {
	return bindings.SearchParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SearchParams) UpdateFrom(ref js.Ref) {
	bindings.SearchParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SearchParams) Update(ref js.Ref) {
	bindings.SearchParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SearchParams) FreeMembers(recursive bool) {
	js.Free(
		p.Query.Ref(),
		p.NextFeed.Ref(),
	)
	p.Query = p.Query.FromRef(js.Undefined)
	p.NextFeed = p.NextFeed.FromRef(js.Undefined)
}

type SharesheetLaunchSource uint32

const (
	_ SharesheetLaunchSource = iota

	SharesheetLaunchSource_CONTEXT_MENU
	SharesheetLaunchSource_SHARESHEET_BUTTON
	SharesheetLaunchSource_UNKNOWN
)

func (SharesheetLaunchSource) FromRef(str js.Ref) SharesheetLaunchSource {
	return SharesheetLaunchSource(bindings.ConstOfSharesheetLaunchSource(str))
}

func (x SharesheetLaunchSource) String() (string, bool) {
	switch x {
	case SharesheetLaunchSource_CONTEXT_MENU:
		return "context_menu", true
	case SharesheetLaunchSource_SHARESHEET_BUTTON:
		return "sharesheet_button", true
	case SharesheetLaunchSource_UNKNOWN:
		return "unknown", true
	default:
		return "", false
	}
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

type SourceRestriction uint32

const (
	_ SourceRestriction = iota

	SourceRestriction_ANY_SOURCE
	SourceRestriction_NATIVE_SOURCE
)

func (SourceRestriction) FromRef(str js.Ref) SourceRestriction {
	return SourceRestriction(bindings.ConstOfSourceRestriction(str))
}

func (x SourceRestriction) String() (string, bool) {
	switch x {
	case SourceRestriction_ANY_SOURCE:
		return "any_source", true
	case SourceRestriction_NATIVE_SOURCE:
		return "native_source", true
	default:
		return "", false
	}
}

type SyncState struct {
	// FileUrl is "SyncState.fileUrl"
	//
	// Optional
	FileUrl js.String
	// SyncStatus is "SyncState.syncStatus"
	//
	// Optional
	SyncStatus SyncStatus
	// Progress is "SyncState.progress"
	//
	// Optional
	//
	// NOTE: FFI_USE_Progress MUST be set to true to make this field effective.
	Progress float64

	FFI_USE_Progress bool // for Progress.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SyncState with all fields set.
func (p SyncState) FromRef(ref js.Ref) SyncState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SyncState in the application heap.
func (p SyncState) New() js.Ref {
	return bindings.SyncStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SyncState) UpdateFrom(ref js.Ref) {
	bindings.SyncStateJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SyncState) Update(ref js.Ref) {
	bindings.SyncStateJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SyncState) FreeMembers(recursive bool) {
	js.Free(
		p.FileUrl.Ref(),
	)
	p.FileUrl = p.FileUrl.FromRef(js.Undefined)
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

type ZoomOperationType uint32

const (
	_ ZoomOperationType = iota

	ZoomOperationType_IN
	ZoomOperationType_OUT
	ZoomOperationType_RESET
)

func (ZoomOperationType) FromRef(str js.Ref) ZoomOperationType {
	return ZoomOperationType(bindings.ConstOfZoomOperationType(str))
}

func (x ZoomOperationType) String() (string, bool) {
	switch x {
	case ZoomOperationType_IN:
		return "in", true
	case ZoomOperationType_OUT:
		return "out", true
	case ZoomOperationType_RESET:
		return "reset", true
	default:
		return "", false
	}
}

// HasFuncAddFileWatch returns true if the function "WEBEXT.fileManagerPrivate.addFileWatch" exists.
func HasFuncAddFileWatch() bool {
	return js.True == bindings.HasFuncAddFileWatch()
}

// FuncAddFileWatch returns the function "WEBEXT.fileManagerPrivate.addFileWatch".
func FuncAddFileWatch() (fn js.Func[func(entry js.Object) js.Promise[js.Boolean]]) {
	bindings.FuncAddFileWatch(
		js.Pointer(&fn),
	)
	return
}

// AddFileWatch calls the function "WEBEXT.fileManagerPrivate.addFileWatch" directly.
func AddFileWatch(entry js.Object) (ret js.Promise[js.Boolean]) {
	bindings.CallAddFileWatch(
		js.Pointer(&ret),
		entry.Ref(),
	)

	return
}

// TryAddFileWatch calls the function "WEBEXT.fileManagerPrivate.addFileWatch"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAddFileWatch(entry js.Object) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAddFileWatch(
		js.Pointer(&ret), js.Pointer(&exception),
		entry.Ref(),
	)

	return
}

// HasFuncAddMount returns true if the function "WEBEXT.fileManagerPrivate.addMount" exists.
func HasFuncAddMount() bool {
	return js.True == bindings.HasFuncAddMount()
}

// FuncAddMount returns the function "WEBEXT.fileManagerPrivate.addMount".
func FuncAddMount() (fn js.Func[func(fileUrl js.String, password js.String) js.Promise[js.String]]) {
	bindings.FuncAddMount(
		js.Pointer(&fn),
	)
	return
}

// AddMount calls the function "WEBEXT.fileManagerPrivate.addMount" directly.
func AddMount(fileUrl js.String, password js.String) (ret js.Promise[js.String]) {
	bindings.CallAddMount(
		js.Pointer(&ret),
		fileUrl.Ref(),
		password.Ref(),
	)

	return
}

// TryAddMount calls the function "WEBEXT.fileManagerPrivate.addMount"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAddMount(fileUrl js.String, password js.String) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAddMount(
		js.Pointer(&ret), js.Pointer(&exception),
		fileUrl.Ref(),
		password.Ref(),
	)

	return
}

// HasFuncAddProvidedFileSystem returns true if the function "WEBEXT.fileManagerPrivate.addProvidedFileSystem" exists.
func HasFuncAddProvidedFileSystem() bool {
	return js.True == bindings.HasFuncAddProvidedFileSystem()
}

// FuncAddProvidedFileSystem returns the function "WEBEXT.fileManagerPrivate.addProvidedFileSystem".
func FuncAddProvidedFileSystem() (fn js.Func[func(providerId js.String) js.Promise[js.Void]]) {
	bindings.FuncAddProvidedFileSystem(
		js.Pointer(&fn),
	)
	return
}

// AddProvidedFileSystem calls the function "WEBEXT.fileManagerPrivate.addProvidedFileSystem" directly.
func AddProvidedFileSystem(providerId js.String) (ret js.Promise[js.Void]) {
	bindings.CallAddProvidedFileSystem(
		js.Pointer(&ret),
		providerId.Ref(),
	)

	return
}

// TryAddProvidedFileSystem calls the function "WEBEXT.fileManagerPrivate.addProvidedFileSystem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAddProvidedFileSystem(providerId js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAddProvidedFileSystem(
		js.Pointer(&ret), js.Pointer(&exception),
		providerId.Ref(),
	)

	return
}

// HasFuncCalculateBulkPinRequiredSpace returns true if the function "WEBEXT.fileManagerPrivate.calculateBulkPinRequiredSpace" exists.
func HasFuncCalculateBulkPinRequiredSpace() bool {
	return js.True == bindings.HasFuncCalculateBulkPinRequiredSpace()
}

// FuncCalculateBulkPinRequiredSpace returns the function "WEBEXT.fileManagerPrivate.calculateBulkPinRequiredSpace".
func FuncCalculateBulkPinRequiredSpace() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncCalculateBulkPinRequiredSpace(
		js.Pointer(&fn),
	)
	return
}

// CalculateBulkPinRequiredSpace calls the function "WEBEXT.fileManagerPrivate.calculateBulkPinRequiredSpace" directly.
func CalculateBulkPinRequiredSpace() (ret js.Promise[js.Void]) {
	bindings.CallCalculateBulkPinRequiredSpace(
		js.Pointer(&ret),
	)

	return
}

// TryCalculateBulkPinRequiredSpace calls the function "WEBEXT.fileManagerPrivate.calculateBulkPinRequiredSpace"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCalculateBulkPinRequiredSpace() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCalculateBulkPinRequiredSpace(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCancelDialog returns true if the function "WEBEXT.fileManagerPrivate.cancelDialog" exists.
func HasFuncCancelDialog() bool {
	return js.True == bindings.HasFuncCancelDialog()
}

// FuncCancelDialog returns the function "WEBEXT.fileManagerPrivate.cancelDialog".
func FuncCancelDialog() (fn js.Func[func()]) {
	bindings.FuncCancelDialog(
		js.Pointer(&fn),
	)
	return
}

// CancelDialog calls the function "WEBEXT.fileManagerPrivate.cancelDialog" directly.
func CancelDialog() (ret js.Void) {
	bindings.CallCancelDialog(
		js.Pointer(&ret),
	)

	return
}

// TryCancelDialog calls the function "WEBEXT.fileManagerPrivate.cancelDialog"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCancelDialog() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCancelDialog(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCancelIOTask returns true if the function "WEBEXT.fileManagerPrivate.cancelIOTask" exists.
func HasFuncCancelIOTask() bool {
	return js.True == bindings.HasFuncCancelIOTask()
}

// FuncCancelIOTask returns the function "WEBEXT.fileManagerPrivate.cancelIOTask".
func FuncCancelIOTask() (fn js.Func[func(taskId int32)]) {
	bindings.FuncCancelIOTask(
		js.Pointer(&fn),
	)
	return
}

// CancelIOTask calls the function "WEBEXT.fileManagerPrivate.cancelIOTask" directly.
func CancelIOTask(taskId int32) (ret js.Void) {
	bindings.CallCancelIOTask(
		js.Pointer(&ret),
		int32(taskId),
	)

	return
}

// TryCancelIOTask calls the function "WEBEXT.fileManagerPrivate.cancelIOTask"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCancelIOTask(taskId int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCancelIOTask(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(taskId),
	)

	return
}

// HasFuncCancelMounting returns true if the function "WEBEXT.fileManagerPrivate.cancelMounting" exists.
func HasFuncCancelMounting() bool {
	return js.True == bindings.HasFuncCancelMounting()
}

// FuncCancelMounting returns the function "WEBEXT.fileManagerPrivate.cancelMounting".
func FuncCancelMounting() (fn js.Func[func(fileUrl js.String) js.Promise[js.Void]]) {
	bindings.FuncCancelMounting(
		js.Pointer(&fn),
	)
	return
}

// CancelMounting calls the function "WEBEXT.fileManagerPrivate.cancelMounting" directly.
func CancelMounting(fileUrl js.String) (ret js.Promise[js.Void]) {
	bindings.CallCancelMounting(
		js.Pointer(&ret),
		fileUrl.Ref(),
	)

	return
}

// TryCancelMounting calls the function "WEBEXT.fileManagerPrivate.cancelMounting"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCancelMounting(fileUrl js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCancelMounting(
		js.Pointer(&ret), js.Pointer(&exception),
		fileUrl.Ref(),
	)

	return
}

// HasFuncComputeChecksum returns true if the function "WEBEXT.fileManagerPrivate.computeChecksum" exists.
func HasFuncComputeChecksum() bool {
	return js.True == bindings.HasFuncComputeChecksum()
}

// FuncComputeChecksum returns the function "WEBEXT.fileManagerPrivate.computeChecksum".
func FuncComputeChecksum() (fn js.Func[func(entry js.Object) js.Promise[js.String]]) {
	bindings.FuncComputeChecksum(
		js.Pointer(&fn),
	)
	return
}

// ComputeChecksum calls the function "WEBEXT.fileManagerPrivate.computeChecksum" directly.
func ComputeChecksum(entry js.Object) (ret js.Promise[js.String]) {
	bindings.CallComputeChecksum(
		js.Pointer(&ret),
		entry.Ref(),
	)

	return
}

// TryComputeChecksum calls the function "WEBEXT.fileManagerPrivate.computeChecksum"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryComputeChecksum(entry js.Object) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryComputeChecksum(
		js.Pointer(&ret), js.Pointer(&exception),
		entry.Ref(),
	)

	return
}

// HasFuncConfigureVolume returns true if the function "WEBEXT.fileManagerPrivate.configureVolume" exists.
func HasFuncConfigureVolume() bool {
	return js.True == bindings.HasFuncConfigureVolume()
}

// FuncConfigureVolume returns the function "WEBEXT.fileManagerPrivate.configureVolume".
func FuncConfigureVolume() (fn js.Func[func(volumeId js.String) js.Promise[js.Void]]) {
	bindings.FuncConfigureVolume(
		js.Pointer(&fn),
	)
	return
}

// ConfigureVolume calls the function "WEBEXT.fileManagerPrivate.configureVolume" directly.
func ConfigureVolume(volumeId js.String) (ret js.Promise[js.Void]) {
	bindings.CallConfigureVolume(
		js.Pointer(&ret),
		volumeId.Ref(),
	)

	return
}

// TryConfigureVolume calls the function "WEBEXT.fileManagerPrivate.configureVolume"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryConfigureVolume(volumeId js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryConfigureVolume(
		js.Pointer(&ret), js.Pointer(&exception),
		volumeId.Ref(),
	)

	return
}

// HasFuncDismissIOTask returns true if the function "WEBEXT.fileManagerPrivate.dismissIOTask" exists.
func HasFuncDismissIOTask() bool {
	return js.True == bindings.HasFuncDismissIOTask()
}

// FuncDismissIOTask returns the function "WEBEXT.fileManagerPrivate.dismissIOTask".
func FuncDismissIOTask() (fn js.Func[func(taskId int32) js.Promise[js.Void]]) {
	bindings.FuncDismissIOTask(
		js.Pointer(&fn),
	)
	return
}

// DismissIOTask calls the function "WEBEXT.fileManagerPrivate.dismissIOTask" directly.
func DismissIOTask(taskId int32) (ret js.Promise[js.Void]) {
	bindings.CallDismissIOTask(
		js.Pointer(&ret),
		int32(taskId),
	)

	return
}

// TryDismissIOTask calls the function "WEBEXT.fileManagerPrivate.dismissIOTask"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDismissIOTask(taskId int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDismissIOTask(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(taskId),
	)

	return
}

// HasFuncEnableExternalFileScheme returns true if the function "WEBEXT.fileManagerPrivate.enableExternalFileScheme" exists.
func HasFuncEnableExternalFileScheme() bool {
	return js.True == bindings.HasFuncEnableExternalFileScheme()
}

// FuncEnableExternalFileScheme returns the function "WEBEXT.fileManagerPrivate.enableExternalFileScheme".
func FuncEnableExternalFileScheme() (fn js.Func[func()]) {
	bindings.FuncEnableExternalFileScheme(
		js.Pointer(&fn),
	)
	return
}

// EnableExternalFileScheme calls the function "WEBEXT.fileManagerPrivate.enableExternalFileScheme" directly.
func EnableExternalFileScheme() (ret js.Void) {
	bindings.CallEnableExternalFileScheme(
		js.Pointer(&ret),
	)

	return
}

// TryEnableExternalFileScheme calls the function "WEBEXT.fileManagerPrivate.enableExternalFileScheme"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryEnableExternalFileScheme() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEnableExternalFileScheme(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncExecuteCustomAction returns true if the function "WEBEXT.fileManagerPrivate.executeCustomAction" exists.
func HasFuncExecuteCustomAction() bool {
	return js.True == bindings.HasFuncExecuteCustomAction()
}

// FuncExecuteCustomAction returns the function "WEBEXT.fileManagerPrivate.executeCustomAction".
func FuncExecuteCustomAction() (fn js.Func[func(entries js.Array[js.Object], actionId js.String) js.Promise[js.Void]]) {
	bindings.FuncExecuteCustomAction(
		js.Pointer(&fn),
	)
	return
}

// ExecuteCustomAction calls the function "WEBEXT.fileManagerPrivate.executeCustomAction" directly.
func ExecuteCustomAction(entries js.Array[js.Object], actionId js.String) (ret js.Promise[js.Void]) {
	bindings.CallExecuteCustomAction(
		js.Pointer(&ret),
		entries.Ref(),
		actionId.Ref(),
	)

	return
}

// TryExecuteCustomAction calls the function "WEBEXT.fileManagerPrivate.executeCustomAction"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryExecuteCustomAction(entries js.Array[js.Object], actionId js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryExecuteCustomAction(
		js.Pointer(&ret), js.Pointer(&exception),
		entries.Ref(),
		actionId.Ref(),
	)

	return
}

// HasFuncExecuteTask returns true if the function "WEBEXT.fileManagerPrivate.executeTask" exists.
func HasFuncExecuteTask() bool {
	return js.True == bindings.HasFuncExecuteTask()
}

// FuncExecuteTask returns the function "WEBEXT.fileManagerPrivate.executeTask".
func FuncExecuteTask() (fn js.Func[func(descriptor FileTaskDescriptor, entries js.Array[js.Object]) js.Promise[TaskResult]]) {
	bindings.FuncExecuteTask(
		js.Pointer(&fn),
	)
	return
}

// ExecuteTask calls the function "WEBEXT.fileManagerPrivate.executeTask" directly.
func ExecuteTask(descriptor FileTaskDescriptor, entries js.Array[js.Object]) (ret js.Promise[TaskResult]) {
	bindings.CallExecuteTask(
		js.Pointer(&ret),
		js.Pointer(&descriptor),
		entries.Ref(),
	)

	return
}

// TryExecuteTask calls the function "WEBEXT.fileManagerPrivate.executeTask"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryExecuteTask(descriptor FileTaskDescriptor, entries js.Array[js.Object]) (ret js.Promise[TaskResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryExecuteTask(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
		entries.Ref(),
	)

	return
}

// HasFuncFormatVolume returns true if the function "WEBEXT.fileManagerPrivate.formatVolume" exists.
func HasFuncFormatVolume() bool {
	return js.True == bindings.HasFuncFormatVolume()
}

// FuncFormatVolume returns the function "WEBEXT.fileManagerPrivate.formatVolume".
func FuncFormatVolume() (fn js.Func[func(volumeId js.String, filesystem FormatFileSystemType, volumeLabel js.String)]) {
	bindings.FuncFormatVolume(
		js.Pointer(&fn),
	)
	return
}

// FormatVolume calls the function "WEBEXT.fileManagerPrivate.formatVolume" directly.
func FormatVolume(volumeId js.String, filesystem FormatFileSystemType, volumeLabel js.String) (ret js.Void) {
	bindings.CallFormatVolume(
		js.Pointer(&ret),
		volumeId.Ref(),
		uint32(filesystem),
		volumeLabel.Ref(),
	)

	return
}

// TryFormatVolume calls the function "WEBEXT.fileManagerPrivate.formatVolume"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryFormatVolume(volumeId js.String, filesystem FormatFileSystemType, volumeLabel js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFormatVolume(
		js.Pointer(&ret), js.Pointer(&exception),
		volumeId.Ref(),
		uint32(filesystem),
		volumeLabel.Ref(),
	)

	return
}

// HasFuncGetAndroidPickerApps returns true if the function "WEBEXT.fileManagerPrivate.getAndroidPickerApps" exists.
func HasFuncGetAndroidPickerApps() bool {
	return js.True == bindings.HasFuncGetAndroidPickerApps()
}

// FuncGetAndroidPickerApps returns the function "WEBEXT.fileManagerPrivate.getAndroidPickerApps".
func FuncGetAndroidPickerApps() (fn js.Func[func(extensions js.Array[js.String]) js.Promise[js.Array[AndroidApp]]]) {
	bindings.FuncGetAndroidPickerApps(
		js.Pointer(&fn),
	)
	return
}

// GetAndroidPickerApps calls the function "WEBEXT.fileManagerPrivate.getAndroidPickerApps" directly.
func GetAndroidPickerApps(extensions js.Array[js.String]) (ret js.Promise[js.Array[AndroidApp]]) {
	bindings.CallGetAndroidPickerApps(
		js.Pointer(&ret),
		extensions.Ref(),
	)

	return
}

// TryGetAndroidPickerApps calls the function "WEBEXT.fileManagerPrivate.getAndroidPickerApps"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAndroidPickerApps(extensions js.Array[js.String]) (ret js.Promise[js.Array[AndroidApp]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAndroidPickerApps(
		js.Pointer(&ret), js.Pointer(&exception),
		extensions.Ref(),
	)

	return
}

// HasFuncGetBulkPinProgress returns true if the function "WEBEXT.fileManagerPrivate.getBulkPinProgress" exists.
func HasFuncGetBulkPinProgress() bool {
	return js.True == bindings.HasFuncGetBulkPinProgress()
}

// FuncGetBulkPinProgress returns the function "WEBEXT.fileManagerPrivate.getBulkPinProgress".
func FuncGetBulkPinProgress() (fn js.Func[func() js.Promise[BulkPinProgress]]) {
	bindings.FuncGetBulkPinProgress(
		js.Pointer(&fn),
	)
	return
}

// GetBulkPinProgress calls the function "WEBEXT.fileManagerPrivate.getBulkPinProgress" directly.
func GetBulkPinProgress() (ret js.Promise[BulkPinProgress]) {
	bindings.CallGetBulkPinProgress(
		js.Pointer(&ret),
	)

	return
}

// TryGetBulkPinProgress calls the function "WEBEXT.fileManagerPrivate.getBulkPinProgress"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetBulkPinProgress() (ret js.Promise[BulkPinProgress], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetBulkPinProgress(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetContentMetadata returns true if the function "WEBEXT.fileManagerPrivate.getContentMetadata" exists.
func HasFuncGetContentMetadata() bool {
	return js.True == bindings.HasFuncGetContentMetadata()
}

// FuncGetContentMetadata returns the function "WEBEXT.fileManagerPrivate.getContentMetadata".
func FuncGetContentMetadata() (fn js.Func[func(fileEntry js.Object, mimeType js.String, includeImages bool) js.Promise[MediaMetadata]]) {
	bindings.FuncGetContentMetadata(
		js.Pointer(&fn),
	)
	return
}

// GetContentMetadata calls the function "WEBEXT.fileManagerPrivate.getContentMetadata" directly.
func GetContentMetadata(fileEntry js.Object, mimeType js.String, includeImages bool) (ret js.Promise[MediaMetadata]) {
	bindings.CallGetContentMetadata(
		js.Pointer(&ret),
		fileEntry.Ref(),
		mimeType.Ref(),
		js.Bool(bool(includeImages)),
	)

	return
}

// TryGetContentMetadata calls the function "WEBEXT.fileManagerPrivate.getContentMetadata"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetContentMetadata(fileEntry js.Object, mimeType js.String, includeImages bool) (ret js.Promise[MediaMetadata], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetContentMetadata(
		js.Pointer(&ret), js.Pointer(&exception),
		fileEntry.Ref(),
		mimeType.Ref(),
		js.Bool(bool(includeImages)),
	)

	return
}

// HasFuncGetContentMimeType returns true if the function "WEBEXT.fileManagerPrivate.getContentMimeType" exists.
func HasFuncGetContentMimeType() bool {
	return js.True == bindings.HasFuncGetContentMimeType()
}

// FuncGetContentMimeType returns the function "WEBEXT.fileManagerPrivate.getContentMimeType".
func FuncGetContentMimeType() (fn js.Func[func(fileEntry js.Object) js.Promise[js.String]]) {
	bindings.FuncGetContentMimeType(
		js.Pointer(&fn),
	)
	return
}

// GetContentMimeType calls the function "WEBEXT.fileManagerPrivate.getContentMimeType" directly.
func GetContentMimeType(fileEntry js.Object) (ret js.Promise[js.String]) {
	bindings.CallGetContentMimeType(
		js.Pointer(&ret),
		fileEntry.Ref(),
	)

	return
}

// TryGetContentMimeType calls the function "WEBEXT.fileManagerPrivate.getContentMimeType"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetContentMimeType(fileEntry js.Object) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetContentMimeType(
		js.Pointer(&ret), js.Pointer(&exception),
		fileEntry.Ref(),
	)

	return
}

// HasFuncGetCrostiniSharedPaths returns true if the function "WEBEXT.fileManagerPrivate.getCrostiniSharedPaths" exists.
func HasFuncGetCrostiniSharedPaths() bool {
	return js.True == bindings.HasFuncGetCrostiniSharedPaths()
}

// FuncGetCrostiniSharedPaths returns the function "WEBEXT.fileManagerPrivate.getCrostiniSharedPaths".
func FuncGetCrostiniSharedPaths() (fn js.Func[func(observeFirstForSession bool, vmName js.String, callback js.Func[func(entries js.Array[js.Object], firstForSession bool)])]) {
	bindings.FuncGetCrostiniSharedPaths(
		js.Pointer(&fn),
	)
	return
}

// GetCrostiniSharedPaths calls the function "WEBEXT.fileManagerPrivate.getCrostiniSharedPaths" directly.
func GetCrostiniSharedPaths(observeFirstForSession bool, vmName js.String, callback js.Func[func(entries js.Array[js.Object], firstForSession bool)]) (ret js.Void) {
	bindings.CallGetCrostiniSharedPaths(
		js.Pointer(&ret),
		js.Bool(bool(observeFirstForSession)),
		vmName.Ref(),
		callback.Ref(),
	)

	return
}

// TryGetCrostiniSharedPaths calls the function "WEBEXT.fileManagerPrivate.getCrostiniSharedPaths"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetCrostiniSharedPaths(observeFirstForSession bool, vmName js.String, callback js.Func[func(entries js.Array[js.Object], firstForSession bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetCrostiniSharedPaths(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(observeFirstForSession)),
		vmName.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncGetCustomActions returns true if the function "WEBEXT.fileManagerPrivate.getCustomActions" exists.
func HasFuncGetCustomActions() bool {
	return js.True == bindings.HasFuncGetCustomActions()
}

// FuncGetCustomActions returns the function "WEBEXT.fileManagerPrivate.getCustomActions".
func FuncGetCustomActions() (fn js.Func[func(entries js.Array[js.Object]) js.Promise[js.Array[FileSystemProviderAction]]]) {
	bindings.FuncGetCustomActions(
		js.Pointer(&fn),
	)
	return
}

// GetCustomActions calls the function "WEBEXT.fileManagerPrivate.getCustomActions" directly.
func GetCustomActions(entries js.Array[js.Object]) (ret js.Promise[js.Array[FileSystemProviderAction]]) {
	bindings.CallGetCustomActions(
		js.Pointer(&ret),
		entries.Ref(),
	)

	return
}

// TryGetCustomActions calls the function "WEBEXT.fileManagerPrivate.getCustomActions"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetCustomActions(entries js.Array[js.Object]) (ret js.Promise[js.Array[FileSystemProviderAction]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetCustomActions(
		js.Pointer(&ret), js.Pointer(&exception),
		entries.Ref(),
	)

	return
}

// HasFuncGetDeviceConnectionState returns true if the function "WEBEXT.fileManagerPrivate.getDeviceConnectionState" exists.
func HasFuncGetDeviceConnectionState() bool {
	return js.True == bindings.HasFuncGetDeviceConnectionState()
}

// FuncGetDeviceConnectionState returns the function "WEBEXT.fileManagerPrivate.getDeviceConnectionState".
func FuncGetDeviceConnectionState() (fn js.Func[func(callback js.Func[func(result DeviceConnectionState)])]) {
	bindings.FuncGetDeviceConnectionState(
		js.Pointer(&fn),
	)
	return
}

// GetDeviceConnectionState calls the function "WEBEXT.fileManagerPrivate.getDeviceConnectionState" directly.
func GetDeviceConnectionState(callback js.Func[func(result DeviceConnectionState)]) (ret js.Void) {
	bindings.CallGetDeviceConnectionState(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryGetDeviceConnectionState calls the function "WEBEXT.fileManagerPrivate.getDeviceConnectionState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDeviceConnectionState(callback js.Func[func(result DeviceConnectionState)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDeviceConnectionState(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncGetDialogCaller returns true if the function "WEBEXT.fileManagerPrivate.getDialogCaller" exists.
func HasFuncGetDialogCaller() bool {
	return js.True == bindings.HasFuncGetDialogCaller()
}

// FuncGetDialogCaller returns the function "WEBEXT.fileManagerPrivate.getDialogCaller".
func FuncGetDialogCaller() (fn js.Func[func() js.Promise[DialogCallerInformation]]) {
	bindings.FuncGetDialogCaller(
		js.Pointer(&fn),
	)
	return
}

// GetDialogCaller calls the function "WEBEXT.fileManagerPrivate.getDialogCaller" directly.
func GetDialogCaller() (ret js.Promise[DialogCallerInformation]) {
	bindings.CallGetDialogCaller(
		js.Pointer(&ret),
	)

	return
}

// TryGetDialogCaller calls the function "WEBEXT.fileManagerPrivate.getDialogCaller"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDialogCaller() (ret js.Promise[DialogCallerInformation], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDialogCaller(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDirectorySize returns true if the function "WEBEXT.fileManagerPrivate.getDirectorySize" exists.
func HasFuncGetDirectorySize() bool {
	return js.True == bindings.HasFuncGetDirectorySize()
}

// FuncGetDirectorySize returns the function "WEBEXT.fileManagerPrivate.getDirectorySize".
func FuncGetDirectorySize() (fn js.Func[func(entry js.Object) js.Promise[js.Number[float64]]]) {
	bindings.FuncGetDirectorySize(
		js.Pointer(&fn),
	)
	return
}

// GetDirectorySize calls the function "WEBEXT.fileManagerPrivate.getDirectorySize" directly.
func GetDirectorySize(entry js.Object) (ret js.Promise[js.Number[float64]]) {
	bindings.CallGetDirectorySize(
		js.Pointer(&ret),
		entry.Ref(),
	)

	return
}

// TryGetDirectorySize calls the function "WEBEXT.fileManagerPrivate.getDirectorySize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDirectorySize(entry js.Object) (ret js.Promise[js.Number[float64]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDirectorySize(
		js.Pointer(&ret), js.Pointer(&exception),
		entry.Ref(),
	)

	return
}

// HasFuncGetDisallowedTransfers returns true if the function "WEBEXT.fileManagerPrivate.getDisallowedTransfers" exists.
func HasFuncGetDisallowedTransfers() bool {
	return js.True == bindings.HasFuncGetDisallowedTransfers()
}

// FuncGetDisallowedTransfers returns the function "WEBEXT.fileManagerPrivate.getDisallowedTransfers".
func FuncGetDisallowedTransfers() (fn js.Func[func(entries js.Array[js.Object], destinationEntry js.Object, isMove bool) js.Promise[js.Array[js.Object]]]) {
	bindings.FuncGetDisallowedTransfers(
		js.Pointer(&fn),
	)
	return
}

// GetDisallowedTransfers calls the function "WEBEXT.fileManagerPrivate.getDisallowedTransfers" directly.
func GetDisallowedTransfers(entries js.Array[js.Object], destinationEntry js.Object, isMove bool) (ret js.Promise[js.Array[js.Object]]) {
	bindings.CallGetDisallowedTransfers(
		js.Pointer(&ret),
		entries.Ref(),
		destinationEntry.Ref(),
		js.Bool(bool(isMove)),
	)

	return
}

// TryGetDisallowedTransfers calls the function "WEBEXT.fileManagerPrivate.getDisallowedTransfers"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDisallowedTransfers(entries js.Array[js.Object], destinationEntry js.Object, isMove bool) (ret js.Promise[js.Array[js.Object]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDisallowedTransfers(
		js.Pointer(&ret), js.Pointer(&exception),
		entries.Ref(),
		destinationEntry.Ref(),
		js.Bool(bool(isMove)),
	)

	return
}

// HasFuncGetDlpBlockedComponents returns true if the function "WEBEXT.fileManagerPrivate.getDlpBlockedComponents" exists.
func HasFuncGetDlpBlockedComponents() bool {
	return js.True == bindings.HasFuncGetDlpBlockedComponents()
}

// FuncGetDlpBlockedComponents returns the function "WEBEXT.fileManagerPrivate.getDlpBlockedComponents".
func FuncGetDlpBlockedComponents() (fn js.Func[func(sourceUrl js.String) js.Promise[js.Array[VolumeType]]]) {
	bindings.FuncGetDlpBlockedComponents(
		js.Pointer(&fn),
	)
	return
}

// GetDlpBlockedComponents calls the function "WEBEXT.fileManagerPrivate.getDlpBlockedComponents" directly.
func GetDlpBlockedComponents(sourceUrl js.String) (ret js.Promise[js.Array[VolumeType]]) {
	bindings.CallGetDlpBlockedComponents(
		js.Pointer(&ret),
		sourceUrl.Ref(),
	)

	return
}

// TryGetDlpBlockedComponents calls the function "WEBEXT.fileManagerPrivate.getDlpBlockedComponents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDlpBlockedComponents(sourceUrl js.String) (ret js.Promise[js.Array[VolumeType]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDlpBlockedComponents(
		js.Pointer(&ret), js.Pointer(&exception),
		sourceUrl.Ref(),
	)

	return
}

// HasFuncGetDlpMetadata returns true if the function "WEBEXT.fileManagerPrivate.getDlpMetadata" exists.
func HasFuncGetDlpMetadata() bool {
	return js.True == bindings.HasFuncGetDlpMetadata()
}

// FuncGetDlpMetadata returns the function "WEBEXT.fileManagerPrivate.getDlpMetadata".
func FuncGetDlpMetadata() (fn js.Func[func(entries js.Array[js.Object]) js.Promise[js.Array[DlpMetadata]]]) {
	bindings.FuncGetDlpMetadata(
		js.Pointer(&fn),
	)
	return
}

// GetDlpMetadata calls the function "WEBEXT.fileManagerPrivate.getDlpMetadata" directly.
func GetDlpMetadata(entries js.Array[js.Object]) (ret js.Promise[js.Array[DlpMetadata]]) {
	bindings.CallGetDlpMetadata(
		js.Pointer(&ret),
		entries.Ref(),
	)

	return
}

// TryGetDlpMetadata calls the function "WEBEXT.fileManagerPrivate.getDlpMetadata"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDlpMetadata(entries js.Array[js.Object]) (ret js.Promise[js.Array[DlpMetadata]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDlpMetadata(
		js.Pointer(&ret), js.Pointer(&exception),
		entries.Ref(),
	)

	return
}

// HasFuncGetDlpRestrictionDetails returns true if the function "WEBEXT.fileManagerPrivate.getDlpRestrictionDetails" exists.
func HasFuncGetDlpRestrictionDetails() bool {
	return js.True == bindings.HasFuncGetDlpRestrictionDetails()
}

// FuncGetDlpRestrictionDetails returns the function "WEBEXT.fileManagerPrivate.getDlpRestrictionDetails".
func FuncGetDlpRestrictionDetails() (fn js.Func[func(sourceUrl js.String) js.Promise[js.Array[DlpRestrictionDetails]]]) {
	bindings.FuncGetDlpRestrictionDetails(
		js.Pointer(&fn),
	)
	return
}

// GetDlpRestrictionDetails calls the function "WEBEXT.fileManagerPrivate.getDlpRestrictionDetails" directly.
func GetDlpRestrictionDetails(sourceUrl js.String) (ret js.Promise[js.Array[DlpRestrictionDetails]]) {
	bindings.CallGetDlpRestrictionDetails(
		js.Pointer(&ret),
		sourceUrl.Ref(),
	)

	return
}

// TryGetDlpRestrictionDetails calls the function "WEBEXT.fileManagerPrivate.getDlpRestrictionDetails"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDlpRestrictionDetails(sourceUrl js.String) (ret js.Promise[js.Array[DlpRestrictionDetails]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDlpRestrictionDetails(
		js.Pointer(&ret), js.Pointer(&exception),
		sourceUrl.Ref(),
	)

	return
}

// HasFuncGetDriveConnectionState returns true if the function "WEBEXT.fileManagerPrivate.getDriveConnectionState" exists.
func HasFuncGetDriveConnectionState() bool {
	return js.True == bindings.HasFuncGetDriveConnectionState()
}

// FuncGetDriveConnectionState returns the function "WEBEXT.fileManagerPrivate.getDriveConnectionState".
func FuncGetDriveConnectionState() (fn js.Func[func() js.Promise[DriveConnectionState]]) {
	bindings.FuncGetDriveConnectionState(
		js.Pointer(&fn),
	)
	return
}

// GetDriveConnectionState calls the function "WEBEXT.fileManagerPrivate.getDriveConnectionState" directly.
func GetDriveConnectionState() (ret js.Promise[DriveConnectionState]) {
	bindings.CallGetDriveConnectionState(
		js.Pointer(&ret),
	)

	return
}

// TryGetDriveConnectionState calls the function "WEBEXT.fileManagerPrivate.getDriveConnectionState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDriveConnectionState() (ret js.Promise[DriveConnectionState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDriveConnectionState(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDriveQuotaMetadata returns true if the function "WEBEXT.fileManagerPrivate.getDriveQuotaMetadata" exists.
func HasFuncGetDriveQuotaMetadata() bool {
	return js.True == bindings.HasFuncGetDriveQuotaMetadata()
}

// FuncGetDriveQuotaMetadata returns the function "WEBEXT.fileManagerPrivate.getDriveQuotaMetadata".
func FuncGetDriveQuotaMetadata() (fn js.Func[func(entry js.Object) js.Promise[DriveQuotaMetadata]]) {
	bindings.FuncGetDriveQuotaMetadata(
		js.Pointer(&fn),
	)
	return
}

// GetDriveQuotaMetadata calls the function "WEBEXT.fileManagerPrivate.getDriveQuotaMetadata" directly.
func GetDriveQuotaMetadata(entry js.Object) (ret js.Promise[DriveQuotaMetadata]) {
	bindings.CallGetDriveQuotaMetadata(
		js.Pointer(&ret),
		entry.Ref(),
	)

	return
}

// TryGetDriveQuotaMetadata calls the function "WEBEXT.fileManagerPrivate.getDriveQuotaMetadata"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDriveQuotaMetadata(entry js.Object) (ret js.Promise[DriveQuotaMetadata], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDriveQuotaMetadata(
		js.Pointer(&ret), js.Pointer(&exception),
		entry.Ref(),
	)

	return
}

// HasFuncGetEntryProperties returns true if the function "WEBEXT.fileManagerPrivate.getEntryProperties" exists.
func HasFuncGetEntryProperties() bool {
	return js.True == bindings.HasFuncGetEntryProperties()
}

// FuncGetEntryProperties returns the function "WEBEXT.fileManagerPrivate.getEntryProperties".
func FuncGetEntryProperties() (fn js.Func[func(entries js.Array[js.Object], names js.Array[EntryPropertyName]) js.Promise[js.Array[EntryProperties]]]) {
	bindings.FuncGetEntryProperties(
		js.Pointer(&fn),
	)
	return
}

// GetEntryProperties calls the function "WEBEXT.fileManagerPrivate.getEntryProperties" directly.
func GetEntryProperties(entries js.Array[js.Object], names js.Array[EntryPropertyName]) (ret js.Promise[js.Array[EntryProperties]]) {
	bindings.CallGetEntryProperties(
		js.Pointer(&ret),
		entries.Ref(),
		names.Ref(),
	)

	return
}

// TryGetEntryProperties calls the function "WEBEXT.fileManagerPrivate.getEntryProperties"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetEntryProperties(entries js.Array[js.Object], names js.Array[EntryPropertyName]) (ret js.Promise[js.Array[EntryProperties]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetEntryProperties(
		js.Pointer(&ret), js.Pointer(&exception),
		entries.Ref(),
		names.Ref(),
	)

	return
}

// HasFuncGetFileTasks returns true if the function "WEBEXT.fileManagerPrivate.getFileTasks" exists.
func HasFuncGetFileTasks() bool {
	return js.True == bindings.HasFuncGetFileTasks()
}

// FuncGetFileTasks returns the function "WEBEXT.fileManagerPrivate.getFileTasks".
func FuncGetFileTasks() (fn js.Func[func(entries js.Array[js.Object], dlpSourceUrls js.Array[js.String]) js.Promise[ResultingTasks]]) {
	bindings.FuncGetFileTasks(
		js.Pointer(&fn),
	)
	return
}

// GetFileTasks calls the function "WEBEXT.fileManagerPrivate.getFileTasks" directly.
func GetFileTasks(entries js.Array[js.Object], dlpSourceUrls js.Array[js.String]) (ret js.Promise[ResultingTasks]) {
	bindings.CallGetFileTasks(
		js.Pointer(&ret),
		entries.Ref(),
		dlpSourceUrls.Ref(),
	)

	return
}

// TryGetFileTasks calls the function "WEBEXT.fileManagerPrivate.getFileTasks"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetFileTasks(entries js.Array[js.Object], dlpSourceUrls js.Array[js.String]) (ret js.Promise[ResultingTasks], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetFileTasks(
		js.Pointer(&ret), js.Pointer(&exception),
		entries.Ref(),
		dlpSourceUrls.Ref(),
	)

	return
}

// HasFuncGetHoldingSpaceState returns true if the function "WEBEXT.fileManagerPrivate.getHoldingSpaceState" exists.
func HasFuncGetHoldingSpaceState() bool {
	return js.True == bindings.HasFuncGetHoldingSpaceState()
}

// FuncGetHoldingSpaceState returns the function "WEBEXT.fileManagerPrivate.getHoldingSpaceState".
func FuncGetHoldingSpaceState() (fn js.Func[func() js.Promise[HoldingSpaceState]]) {
	bindings.FuncGetHoldingSpaceState(
		js.Pointer(&fn),
	)
	return
}

// GetHoldingSpaceState calls the function "WEBEXT.fileManagerPrivate.getHoldingSpaceState" directly.
func GetHoldingSpaceState() (ret js.Promise[HoldingSpaceState]) {
	bindings.CallGetHoldingSpaceState(
		js.Pointer(&ret),
	)

	return
}

// TryGetHoldingSpaceState calls the function "WEBEXT.fileManagerPrivate.getHoldingSpaceState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetHoldingSpaceState() (ret js.Promise[HoldingSpaceState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetHoldingSpaceState(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetLinuxPackageInfo returns true if the function "WEBEXT.fileManagerPrivate.getLinuxPackageInfo" exists.
func HasFuncGetLinuxPackageInfo() bool {
	return js.True == bindings.HasFuncGetLinuxPackageInfo()
}

// FuncGetLinuxPackageInfo returns the function "WEBEXT.fileManagerPrivate.getLinuxPackageInfo".
func FuncGetLinuxPackageInfo() (fn js.Func[func(entry js.Object) js.Promise[LinuxPackageInfo]]) {
	bindings.FuncGetLinuxPackageInfo(
		js.Pointer(&fn),
	)
	return
}

// GetLinuxPackageInfo calls the function "WEBEXT.fileManagerPrivate.getLinuxPackageInfo" directly.
func GetLinuxPackageInfo(entry js.Object) (ret js.Promise[LinuxPackageInfo]) {
	bindings.CallGetLinuxPackageInfo(
		js.Pointer(&ret),
		entry.Ref(),
	)

	return
}

// TryGetLinuxPackageInfo calls the function "WEBEXT.fileManagerPrivate.getLinuxPackageInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetLinuxPackageInfo(entry js.Object) (ret js.Promise[LinuxPackageInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetLinuxPackageInfo(
		js.Pointer(&ret), js.Pointer(&exception),
		entry.Ref(),
	)

	return
}

// HasFuncGetMimeType returns true if the function "WEBEXT.fileManagerPrivate.getMimeType" exists.
func HasFuncGetMimeType() bool {
	return js.True == bindings.HasFuncGetMimeType()
}

// FuncGetMimeType returns the function "WEBEXT.fileManagerPrivate.getMimeType".
func FuncGetMimeType() (fn js.Func[func(entry js.Object) js.Promise[js.String]]) {
	bindings.FuncGetMimeType(
		js.Pointer(&fn),
	)
	return
}

// GetMimeType calls the function "WEBEXT.fileManagerPrivate.getMimeType" directly.
func GetMimeType(entry js.Object) (ret js.Promise[js.String]) {
	bindings.CallGetMimeType(
		js.Pointer(&ret),
		entry.Ref(),
	)

	return
}

// TryGetMimeType calls the function "WEBEXT.fileManagerPrivate.getMimeType"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetMimeType(entry js.Object) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetMimeType(
		js.Pointer(&ret), js.Pointer(&exception),
		entry.Ref(),
	)

	return
}

// HasFuncGetPreferences returns true if the function "WEBEXT.fileManagerPrivate.getPreferences" exists.
func HasFuncGetPreferences() bool {
	return js.True == bindings.HasFuncGetPreferences()
}

// FuncGetPreferences returns the function "WEBEXT.fileManagerPrivate.getPreferences".
func FuncGetPreferences() (fn js.Func[func() js.Promise[Preferences]]) {
	bindings.FuncGetPreferences(
		js.Pointer(&fn),
	)
	return
}

// GetPreferences calls the function "WEBEXT.fileManagerPrivate.getPreferences" directly.
func GetPreferences() (ret js.Promise[Preferences]) {
	bindings.CallGetPreferences(
		js.Pointer(&ret),
	)

	return
}

// TryGetPreferences calls the function "WEBEXT.fileManagerPrivate.getPreferences"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetPreferences() (ret js.Promise[Preferences], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetPreferences(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetProfiles returns true if the function "WEBEXT.fileManagerPrivate.getProfiles" exists.
func HasFuncGetProfiles() bool {
	return js.True == bindings.HasFuncGetProfiles()
}

// FuncGetProfiles returns the function "WEBEXT.fileManagerPrivate.getProfiles".
func FuncGetProfiles() (fn js.Func[func(callback js.Func[func(profiles js.Array[ProfileInfo], runningProfile js.String, displayProfile js.String)])]) {
	bindings.FuncGetProfiles(
		js.Pointer(&fn),
	)
	return
}

// GetProfiles calls the function "WEBEXT.fileManagerPrivate.getProfiles" directly.
func GetProfiles(callback js.Func[func(profiles js.Array[ProfileInfo], runningProfile js.String, displayProfile js.String)]) (ret js.Void) {
	bindings.CallGetProfiles(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryGetProfiles calls the function "WEBEXT.fileManagerPrivate.getProfiles"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetProfiles(callback js.Func[func(profiles js.Array[ProfileInfo], runningProfile js.String, displayProfile js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetProfiles(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncGetProviders returns true if the function "WEBEXT.fileManagerPrivate.getProviders" exists.
func HasFuncGetProviders() bool {
	return js.True == bindings.HasFuncGetProviders()
}

// FuncGetProviders returns the function "WEBEXT.fileManagerPrivate.getProviders".
func FuncGetProviders() (fn js.Func[func() js.Promise[js.Array[Provider]]]) {
	bindings.FuncGetProviders(
		js.Pointer(&fn),
	)
	return
}

// GetProviders calls the function "WEBEXT.fileManagerPrivate.getProviders" directly.
func GetProviders() (ret js.Promise[js.Array[Provider]]) {
	bindings.CallGetProviders(
		js.Pointer(&ret),
	)

	return
}

// TryGetProviders calls the function "WEBEXT.fileManagerPrivate.getProviders"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetProviders() (ret js.Promise[js.Array[Provider]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetProviders(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetRecentFiles returns true if the function "WEBEXT.fileManagerPrivate.getRecentFiles" exists.
func HasFuncGetRecentFiles() bool {
	return js.True == bindings.HasFuncGetRecentFiles()
}

// FuncGetRecentFiles returns the function "WEBEXT.fileManagerPrivate.getRecentFiles".
func FuncGetRecentFiles() (fn js.Func[func(restriction SourceRestriction, fileCategory FileCategory, invalidateCache bool) js.Promise[js.Array[js.Object]]]) {
	bindings.FuncGetRecentFiles(
		js.Pointer(&fn),
	)
	return
}

// GetRecentFiles calls the function "WEBEXT.fileManagerPrivate.getRecentFiles" directly.
func GetRecentFiles(restriction SourceRestriction, fileCategory FileCategory, invalidateCache bool) (ret js.Promise[js.Array[js.Object]]) {
	bindings.CallGetRecentFiles(
		js.Pointer(&ret),
		uint32(restriction),
		uint32(fileCategory),
		js.Bool(bool(invalidateCache)),
	)

	return
}

// TryGetRecentFiles calls the function "WEBEXT.fileManagerPrivate.getRecentFiles"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetRecentFiles(restriction SourceRestriction, fileCategory FileCategory, invalidateCache bool) (ret js.Promise[js.Array[js.Object]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetRecentFiles(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(restriction),
		uint32(fileCategory),
		js.Bool(bool(invalidateCache)),
	)

	return
}

// HasFuncGetSizeStats returns true if the function "WEBEXT.fileManagerPrivate.getSizeStats" exists.
func HasFuncGetSizeStats() bool {
	return js.True == bindings.HasFuncGetSizeStats()
}

// FuncGetSizeStats returns the function "WEBEXT.fileManagerPrivate.getSizeStats".
func FuncGetSizeStats() (fn js.Func[func(volumeId js.String) js.Promise[MountPointSizeStats]]) {
	bindings.FuncGetSizeStats(
		js.Pointer(&fn),
	)
	return
}

// GetSizeStats calls the function "WEBEXT.fileManagerPrivate.getSizeStats" directly.
func GetSizeStats(volumeId js.String) (ret js.Promise[MountPointSizeStats]) {
	bindings.CallGetSizeStats(
		js.Pointer(&ret),
		volumeId.Ref(),
	)

	return
}

// TryGetSizeStats calls the function "WEBEXT.fileManagerPrivate.getSizeStats"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetSizeStats(volumeId js.String) (ret js.Promise[MountPointSizeStats], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetSizeStats(
		js.Pointer(&ret), js.Pointer(&exception),
		volumeId.Ref(),
	)

	return
}

// HasFuncGetStrings returns true if the function "WEBEXT.fileManagerPrivate.getStrings" exists.
func HasFuncGetStrings() bool {
	return js.True == bindings.HasFuncGetStrings()
}

// FuncGetStrings returns the function "WEBEXT.fileManagerPrivate.getStrings".
func FuncGetStrings() (fn js.Func[func() js.Promise[js.Object]]) {
	bindings.FuncGetStrings(
		js.Pointer(&fn),
	)
	return
}

// GetStrings calls the function "WEBEXT.fileManagerPrivate.getStrings" directly.
func GetStrings() (ret js.Promise[js.Object]) {
	bindings.CallGetStrings(
		js.Pointer(&ret),
	)

	return
}

// TryGetStrings calls the function "WEBEXT.fileManagerPrivate.getStrings"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetStrings() (ret js.Promise[js.Object], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetStrings(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetVolumeMetadataList returns true if the function "WEBEXT.fileManagerPrivate.getVolumeMetadataList" exists.
func HasFuncGetVolumeMetadataList() bool {
	return js.True == bindings.HasFuncGetVolumeMetadataList()
}

// FuncGetVolumeMetadataList returns the function "WEBEXT.fileManagerPrivate.getVolumeMetadataList".
func FuncGetVolumeMetadataList() (fn js.Func[func() js.Promise[js.Array[VolumeMetadata]]]) {
	bindings.FuncGetVolumeMetadataList(
		js.Pointer(&fn),
	)
	return
}

// GetVolumeMetadataList calls the function "WEBEXT.fileManagerPrivate.getVolumeMetadataList" directly.
func GetVolumeMetadataList() (ret js.Promise[js.Array[VolumeMetadata]]) {
	bindings.CallGetVolumeMetadataList(
		js.Pointer(&ret),
	)

	return
}

// TryGetVolumeMetadataList calls the function "WEBEXT.fileManagerPrivate.getVolumeMetadataList"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetVolumeMetadataList() (ret js.Promise[js.Array[VolumeMetadata]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetVolumeMetadataList(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetVolumeRoot returns true if the function "WEBEXT.fileManagerPrivate.getVolumeRoot" exists.
func HasFuncGetVolumeRoot() bool {
	return js.True == bindings.HasFuncGetVolumeRoot()
}

// FuncGetVolumeRoot returns the function "WEBEXT.fileManagerPrivate.getVolumeRoot".
func FuncGetVolumeRoot() (fn js.Func[func(options GetVolumeRootOptions) js.Promise[js.Object]]) {
	bindings.FuncGetVolumeRoot(
		js.Pointer(&fn),
	)
	return
}

// GetVolumeRoot calls the function "WEBEXT.fileManagerPrivate.getVolumeRoot" directly.
func GetVolumeRoot(options GetVolumeRootOptions) (ret js.Promise[js.Object]) {
	bindings.CallGetVolumeRoot(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetVolumeRoot calls the function "WEBEXT.fileManagerPrivate.getVolumeRoot"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetVolumeRoot(options GetVolumeRootOptions) (ret js.Promise[js.Object], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetVolumeRoot(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncGrantAccess returns true if the function "WEBEXT.fileManagerPrivate.grantAccess" exists.
func HasFuncGrantAccess() bool {
	return js.True == bindings.HasFuncGrantAccess()
}

// FuncGrantAccess returns the function "WEBEXT.fileManagerPrivate.grantAccess".
func FuncGrantAccess() (fn js.Func[func(entryUrls js.Array[js.String]) js.Promise[js.Void]]) {
	bindings.FuncGrantAccess(
		js.Pointer(&fn),
	)
	return
}

// GrantAccess calls the function "WEBEXT.fileManagerPrivate.grantAccess" directly.
func GrantAccess(entryUrls js.Array[js.String]) (ret js.Promise[js.Void]) {
	bindings.CallGrantAccess(
		js.Pointer(&ret),
		entryUrls.Ref(),
	)

	return
}

// TryGrantAccess calls the function "WEBEXT.fileManagerPrivate.grantAccess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGrantAccess(entryUrls js.Array[js.String]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGrantAccess(
		js.Pointer(&ret), js.Pointer(&exception),
		entryUrls.Ref(),
	)

	return
}

// HasFuncImportCrostiniImage returns true if the function "WEBEXT.fileManagerPrivate.importCrostiniImage" exists.
func HasFuncImportCrostiniImage() bool {
	return js.True == bindings.HasFuncImportCrostiniImage()
}

// FuncImportCrostiniImage returns the function "WEBEXT.fileManagerPrivate.importCrostiniImage".
func FuncImportCrostiniImage() (fn js.Func[func(entry js.Object)]) {
	bindings.FuncImportCrostiniImage(
		js.Pointer(&fn),
	)
	return
}

// ImportCrostiniImage calls the function "WEBEXT.fileManagerPrivate.importCrostiniImage" directly.
func ImportCrostiniImage(entry js.Object) (ret js.Void) {
	bindings.CallImportCrostiniImage(
		js.Pointer(&ret),
		entry.Ref(),
	)

	return
}

// TryImportCrostiniImage calls the function "WEBEXT.fileManagerPrivate.importCrostiniImage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryImportCrostiniImage(entry js.Object) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryImportCrostiniImage(
		js.Pointer(&ret), js.Pointer(&exception),
		entry.Ref(),
	)

	return
}

// HasFuncInstallLinuxPackage returns true if the function "WEBEXT.fileManagerPrivate.installLinuxPackage" exists.
func HasFuncInstallLinuxPackage() bool {
	return js.True == bindings.HasFuncInstallLinuxPackage()
}

// FuncInstallLinuxPackage returns the function "WEBEXT.fileManagerPrivate.installLinuxPackage".
func FuncInstallLinuxPackage() (fn js.Func[func(entry js.Object, callback js.Func[func(response InstallLinuxPackageResponse, failure_reason js.String)])]) {
	bindings.FuncInstallLinuxPackage(
		js.Pointer(&fn),
	)
	return
}

// InstallLinuxPackage calls the function "WEBEXT.fileManagerPrivate.installLinuxPackage" directly.
func InstallLinuxPackage(entry js.Object, callback js.Func[func(response InstallLinuxPackageResponse, failure_reason js.String)]) (ret js.Void) {
	bindings.CallInstallLinuxPackage(
		js.Pointer(&ret),
		entry.Ref(),
		callback.Ref(),
	)

	return
}

// TryInstallLinuxPackage calls the function "WEBEXT.fileManagerPrivate.installLinuxPackage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryInstallLinuxPackage(entry js.Object, callback js.Func[func(response InstallLinuxPackageResponse, failure_reason js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryInstallLinuxPackage(
		js.Pointer(&ret), js.Pointer(&exception),
		entry.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncInvokeSharesheet returns true if the function "WEBEXT.fileManagerPrivate.invokeSharesheet" exists.
func HasFuncInvokeSharesheet() bool {
	return js.True == bindings.HasFuncInvokeSharesheet()
}

// FuncInvokeSharesheet returns the function "WEBEXT.fileManagerPrivate.invokeSharesheet".
func FuncInvokeSharesheet() (fn js.Func[func(entries js.Array[js.Object], launchSource SharesheetLaunchSource, dlpSourceUrls js.Array[js.String]) js.Promise[js.Void]]) {
	bindings.FuncInvokeSharesheet(
		js.Pointer(&fn),
	)
	return
}

// InvokeSharesheet calls the function "WEBEXT.fileManagerPrivate.invokeSharesheet" directly.
func InvokeSharesheet(entries js.Array[js.Object], launchSource SharesheetLaunchSource, dlpSourceUrls js.Array[js.String]) (ret js.Promise[js.Void]) {
	bindings.CallInvokeSharesheet(
		js.Pointer(&ret),
		entries.Ref(),
		uint32(launchSource),
		dlpSourceUrls.Ref(),
	)

	return
}

// TryInvokeSharesheet calls the function "WEBEXT.fileManagerPrivate.invokeSharesheet"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryInvokeSharesheet(entries js.Array[js.Object], launchSource SharesheetLaunchSource, dlpSourceUrls js.Array[js.String]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryInvokeSharesheet(
		js.Pointer(&ret), js.Pointer(&exception),
		entries.Ref(),
		uint32(launchSource),
		dlpSourceUrls.Ref(),
	)

	return
}

// HasFuncIsTabletModeEnabled returns true if the function "WEBEXT.fileManagerPrivate.isTabletModeEnabled" exists.
func HasFuncIsTabletModeEnabled() bool {
	return js.True == bindings.HasFuncIsTabletModeEnabled()
}

// FuncIsTabletModeEnabled returns the function "WEBEXT.fileManagerPrivate.isTabletModeEnabled".
func FuncIsTabletModeEnabled() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncIsTabletModeEnabled(
		js.Pointer(&fn),
	)
	return
}

// IsTabletModeEnabled calls the function "WEBEXT.fileManagerPrivate.isTabletModeEnabled" directly.
func IsTabletModeEnabled() (ret js.Promise[js.Boolean]) {
	bindings.CallIsTabletModeEnabled(
		js.Pointer(&ret),
	)

	return
}

// TryIsTabletModeEnabled calls the function "WEBEXT.fileManagerPrivate.isTabletModeEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsTabletModeEnabled() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsTabletModeEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncListMountableGuests returns true if the function "WEBEXT.fileManagerPrivate.listMountableGuests" exists.
func HasFuncListMountableGuests() bool {
	return js.True == bindings.HasFuncListMountableGuests()
}

// FuncListMountableGuests returns the function "WEBEXT.fileManagerPrivate.listMountableGuests".
func FuncListMountableGuests() (fn js.Func[func() js.Promise[js.Array[MountableGuest]]]) {
	bindings.FuncListMountableGuests(
		js.Pointer(&fn),
	)
	return
}

// ListMountableGuests calls the function "WEBEXT.fileManagerPrivate.listMountableGuests" directly.
func ListMountableGuests() (ret js.Promise[js.Array[MountableGuest]]) {
	bindings.CallListMountableGuests(
		js.Pointer(&ret),
	)

	return
}

// TryListMountableGuests calls the function "WEBEXT.fileManagerPrivate.listMountableGuests"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryListMountableGuests() (ret js.Promise[js.Array[MountableGuest]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryListMountableGuests(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMountCrostini returns true if the function "WEBEXT.fileManagerPrivate.mountCrostini" exists.
func HasFuncMountCrostini() bool {
	return js.True == bindings.HasFuncMountCrostini()
}

// FuncMountCrostini returns the function "WEBEXT.fileManagerPrivate.mountCrostini".
func FuncMountCrostini() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncMountCrostini(
		js.Pointer(&fn),
	)
	return
}

// MountCrostini calls the function "WEBEXT.fileManagerPrivate.mountCrostini" directly.
func MountCrostini() (ret js.Promise[js.Void]) {
	bindings.CallMountCrostini(
		js.Pointer(&ret),
	)

	return
}

// TryMountCrostini calls the function "WEBEXT.fileManagerPrivate.mountCrostini"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryMountCrostini() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMountCrostini(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMountGuest returns true if the function "WEBEXT.fileManagerPrivate.mountGuest" exists.
func HasFuncMountGuest() bool {
	return js.True == bindings.HasFuncMountGuest()
}

// FuncMountGuest returns the function "WEBEXT.fileManagerPrivate.mountGuest".
func FuncMountGuest() (fn js.Func[func(id int32) js.Promise[js.Void]]) {
	bindings.FuncMountGuest(
		js.Pointer(&fn),
	)
	return
}

// MountGuest calls the function "WEBEXT.fileManagerPrivate.mountGuest" directly.
func MountGuest(id int32) (ret js.Promise[js.Void]) {
	bindings.CallMountGuest(
		js.Pointer(&ret),
		int32(id),
	)

	return
}

// TryMountGuest calls the function "WEBEXT.fileManagerPrivate.mountGuest"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryMountGuest(id int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMountGuest(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(id),
	)

	return
}

// HasFuncNotifyDriveDialogResult returns true if the function "WEBEXT.fileManagerPrivate.notifyDriveDialogResult" exists.
func HasFuncNotifyDriveDialogResult() bool {
	return js.True == bindings.HasFuncNotifyDriveDialogResult()
}

// FuncNotifyDriveDialogResult returns the function "WEBEXT.fileManagerPrivate.notifyDriveDialogResult".
func FuncNotifyDriveDialogResult() (fn js.Func[func(result DriveDialogResult)]) {
	bindings.FuncNotifyDriveDialogResult(
		js.Pointer(&fn),
	)
	return
}

// NotifyDriveDialogResult calls the function "WEBEXT.fileManagerPrivate.notifyDriveDialogResult" directly.
func NotifyDriveDialogResult(result DriveDialogResult) (ret js.Void) {
	bindings.CallNotifyDriveDialogResult(
		js.Pointer(&ret),
		uint32(result),
	)

	return
}

// TryNotifyDriveDialogResult calls the function "WEBEXT.fileManagerPrivate.notifyDriveDialogResult"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryNotifyDriveDialogResult(result DriveDialogResult) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNotifyDriveDialogResult(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(result),
	)

	return
}

type OnAppsUpdatedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnAppsUpdatedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnAppsUpdatedEventCallbackFunc) DispatchCallback(
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

type OnAppsUpdatedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnAppsUpdatedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnAppsUpdatedEventCallback[T]) DispatchCallback(
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

// HasFuncOnAppsUpdated returns true if the function "WEBEXT.fileManagerPrivate.onAppsUpdated.addListener" exists.
func HasFuncOnAppsUpdated() bool {
	return js.True == bindings.HasFuncOnAppsUpdated()
}

// FuncOnAppsUpdated returns the function "WEBEXT.fileManagerPrivate.onAppsUpdated.addListener".
func FuncOnAppsUpdated() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnAppsUpdated(
		js.Pointer(&fn),
	)
	return
}

// OnAppsUpdated calls the function "WEBEXT.fileManagerPrivate.onAppsUpdated.addListener" directly.
func OnAppsUpdated(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnAppsUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAppsUpdated calls the function "WEBEXT.fileManagerPrivate.onAppsUpdated.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAppsUpdated(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAppsUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAppsUpdated returns true if the function "WEBEXT.fileManagerPrivate.onAppsUpdated.removeListener" exists.
func HasFuncOffAppsUpdated() bool {
	return js.True == bindings.HasFuncOffAppsUpdated()
}

// FuncOffAppsUpdated returns the function "WEBEXT.fileManagerPrivate.onAppsUpdated.removeListener".
func FuncOffAppsUpdated() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffAppsUpdated(
		js.Pointer(&fn),
	)
	return
}

// OffAppsUpdated calls the function "WEBEXT.fileManagerPrivate.onAppsUpdated.removeListener" directly.
func OffAppsUpdated(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffAppsUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAppsUpdated calls the function "WEBEXT.fileManagerPrivate.onAppsUpdated.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAppsUpdated(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAppsUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAppsUpdated returns true if the function "WEBEXT.fileManagerPrivate.onAppsUpdated.hasListener" exists.
func HasFuncHasOnAppsUpdated() bool {
	return js.True == bindings.HasFuncHasOnAppsUpdated()
}

// FuncHasOnAppsUpdated returns the function "WEBEXT.fileManagerPrivate.onAppsUpdated.hasListener".
func FuncHasOnAppsUpdated() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnAppsUpdated(
		js.Pointer(&fn),
	)
	return
}

// HasOnAppsUpdated calls the function "WEBEXT.fileManagerPrivate.onAppsUpdated.hasListener" directly.
func HasOnAppsUpdated(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnAppsUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAppsUpdated calls the function "WEBEXT.fileManagerPrivate.onAppsUpdated.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnAppsUpdated(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnAppsUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnBulkPinProgressEventCallbackFunc func(this js.Ref, progress *BulkPinProgress) js.Ref

func (fn OnBulkPinProgressEventCallbackFunc) Register() js.Func[func(progress *BulkPinProgress)] {
	return js.RegisterCallback[func(progress *BulkPinProgress)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnBulkPinProgressEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 BulkPinProgress
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

type OnBulkPinProgressEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, progress *BulkPinProgress) js.Ref
	Arg T
}

func (cb *OnBulkPinProgressEventCallback[T]) Register() js.Func[func(progress *BulkPinProgress)] {
	return js.RegisterCallback[func(progress *BulkPinProgress)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnBulkPinProgressEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 BulkPinProgress
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

// HasFuncOnBulkPinProgress returns true if the function "WEBEXT.fileManagerPrivate.onBulkPinProgress.addListener" exists.
func HasFuncOnBulkPinProgress() bool {
	return js.True == bindings.HasFuncOnBulkPinProgress()
}

// FuncOnBulkPinProgress returns the function "WEBEXT.fileManagerPrivate.onBulkPinProgress.addListener".
func FuncOnBulkPinProgress() (fn js.Func[func(callback js.Func[func(progress *BulkPinProgress)])]) {
	bindings.FuncOnBulkPinProgress(
		js.Pointer(&fn),
	)
	return
}

// OnBulkPinProgress calls the function "WEBEXT.fileManagerPrivate.onBulkPinProgress.addListener" directly.
func OnBulkPinProgress(callback js.Func[func(progress *BulkPinProgress)]) (ret js.Void) {
	bindings.CallOnBulkPinProgress(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnBulkPinProgress calls the function "WEBEXT.fileManagerPrivate.onBulkPinProgress.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnBulkPinProgress(callback js.Func[func(progress *BulkPinProgress)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnBulkPinProgress(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffBulkPinProgress returns true if the function "WEBEXT.fileManagerPrivate.onBulkPinProgress.removeListener" exists.
func HasFuncOffBulkPinProgress() bool {
	return js.True == bindings.HasFuncOffBulkPinProgress()
}

// FuncOffBulkPinProgress returns the function "WEBEXT.fileManagerPrivate.onBulkPinProgress.removeListener".
func FuncOffBulkPinProgress() (fn js.Func[func(callback js.Func[func(progress *BulkPinProgress)])]) {
	bindings.FuncOffBulkPinProgress(
		js.Pointer(&fn),
	)
	return
}

// OffBulkPinProgress calls the function "WEBEXT.fileManagerPrivate.onBulkPinProgress.removeListener" directly.
func OffBulkPinProgress(callback js.Func[func(progress *BulkPinProgress)]) (ret js.Void) {
	bindings.CallOffBulkPinProgress(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffBulkPinProgress calls the function "WEBEXT.fileManagerPrivate.onBulkPinProgress.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffBulkPinProgress(callback js.Func[func(progress *BulkPinProgress)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffBulkPinProgress(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnBulkPinProgress returns true if the function "WEBEXT.fileManagerPrivate.onBulkPinProgress.hasListener" exists.
func HasFuncHasOnBulkPinProgress() bool {
	return js.True == bindings.HasFuncHasOnBulkPinProgress()
}

// FuncHasOnBulkPinProgress returns the function "WEBEXT.fileManagerPrivate.onBulkPinProgress.hasListener".
func FuncHasOnBulkPinProgress() (fn js.Func[func(callback js.Func[func(progress *BulkPinProgress)]) bool]) {
	bindings.FuncHasOnBulkPinProgress(
		js.Pointer(&fn),
	)
	return
}

// HasOnBulkPinProgress calls the function "WEBEXT.fileManagerPrivate.onBulkPinProgress.hasListener" directly.
func HasOnBulkPinProgress(callback js.Func[func(progress *BulkPinProgress)]) (ret bool) {
	bindings.CallHasOnBulkPinProgress(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnBulkPinProgress calls the function "WEBEXT.fileManagerPrivate.onBulkPinProgress.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnBulkPinProgress(callback js.Func[func(progress *BulkPinProgress)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnBulkPinProgress(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnCrostiniChangedEventCallbackFunc func(this js.Ref, event *CrostiniEvent) js.Ref

func (fn OnCrostiniChangedEventCallbackFunc) Register() js.Func[func(event *CrostiniEvent)] {
	return js.RegisterCallback[func(event *CrostiniEvent)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCrostiniChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CrostiniEvent
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

type OnCrostiniChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event *CrostiniEvent) js.Ref
	Arg T
}

func (cb *OnCrostiniChangedEventCallback[T]) Register() js.Func[func(event *CrostiniEvent)] {
	return js.RegisterCallback[func(event *CrostiniEvent)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCrostiniChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CrostiniEvent
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

// HasFuncOnCrostiniChanged returns true if the function "WEBEXT.fileManagerPrivate.onCrostiniChanged.addListener" exists.
func HasFuncOnCrostiniChanged() bool {
	return js.True == bindings.HasFuncOnCrostiniChanged()
}

// FuncOnCrostiniChanged returns the function "WEBEXT.fileManagerPrivate.onCrostiniChanged.addListener".
func FuncOnCrostiniChanged() (fn js.Func[func(callback js.Func[func(event *CrostiniEvent)])]) {
	bindings.FuncOnCrostiniChanged(
		js.Pointer(&fn),
	)
	return
}

// OnCrostiniChanged calls the function "WEBEXT.fileManagerPrivate.onCrostiniChanged.addListener" directly.
func OnCrostiniChanged(callback js.Func[func(event *CrostiniEvent)]) (ret js.Void) {
	bindings.CallOnCrostiniChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCrostiniChanged calls the function "WEBEXT.fileManagerPrivate.onCrostiniChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCrostiniChanged(callback js.Func[func(event *CrostiniEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCrostiniChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCrostiniChanged returns true if the function "WEBEXT.fileManagerPrivate.onCrostiniChanged.removeListener" exists.
func HasFuncOffCrostiniChanged() bool {
	return js.True == bindings.HasFuncOffCrostiniChanged()
}

// FuncOffCrostiniChanged returns the function "WEBEXT.fileManagerPrivate.onCrostiniChanged.removeListener".
func FuncOffCrostiniChanged() (fn js.Func[func(callback js.Func[func(event *CrostiniEvent)])]) {
	bindings.FuncOffCrostiniChanged(
		js.Pointer(&fn),
	)
	return
}

// OffCrostiniChanged calls the function "WEBEXT.fileManagerPrivate.onCrostiniChanged.removeListener" directly.
func OffCrostiniChanged(callback js.Func[func(event *CrostiniEvent)]) (ret js.Void) {
	bindings.CallOffCrostiniChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCrostiniChanged calls the function "WEBEXT.fileManagerPrivate.onCrostiniChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCrostiniChanged(callback js.Func[func(event *CrostiniEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCrostiniChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCrostiniChanged returns true if the function "WEBEXT.fileManagerPrivate.onCrostiniChanged.hasListener" exists.
func HasFuncHasOnCrostiniChanged() bool {
	return js.True == bindings.HasFuncHasOnCrostiniChanged()
}

// FuncHasOnCrostiniChanged returns the function "WEBEXT.fileManagerPrivate.onCrostiniChanged.hasListener".
func FuncHasOnCrostiniChanged() (fn js.Func[func(callback js.Func[func(event *CrostiniEvent)]) bool]) {
	bindings.FuncHasOnCrostiniChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnCrostiniChanged calls the function "WEBEXT.fileManagerPrivate.onCrostiniChanged.hasListener" directly.
func HasOnCrostiniChanged(callback js.Func[func(event *CrostiniEvent)]) (ret bool) {
	bindings.CallHasOnCrostiniChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCrostiniChanged calls the function "WEBEXT.fileManagerPrivate.onCrostiniChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCrostiniChanged(callback js.Func[func(event *CrostiniEvent)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCrostiniChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDeviceChangedEventCallbackFunc func(this js.Ref, event *DeviceEvent) js.Ref

func (fn OnDeviceChangedEventCallbackFunc) Register() js.Func[func(event *DeviceEvent)] {
	return js.RegisterCallback[func(event *DeviceEvent)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDeviceChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DeviceEvent
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

type OnDeviceChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event *DeviceEvent) js.Ref
	Arg T
}

func (cb *OnDeviceChangedEventCallback[T]) Register() js.Func[func(event *DeviceEvent)] {
	return js.RegisterCallback[func(event *DeviceEvent)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDeviceChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DeviceEvent
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

// HasFuncOnDeviceChanged returns true if the function "WEBEXT.fileManagerPrivate.onDeviceChanged.addListener" exists.
func HasFuncOnDeviceChanged() bool {
	return js.True == bindings.HasFuncOnDeviceChanged()
}

// FuncOnDeviceChanged returns the function "WEBEXT.fileManagerPrivate.onDeviceChanged.addListener".
func FuncOnDeviceChanged() (fn js.Func[func(callback js.Func[func(event *DeviceEvent)])]) {
	bindings.FuncOnDeviceChanged(
		js.Pointer(&fn),
	)
	return
}

// OnDeviceChanged calls the function "WEBEXT.fileManagerPrivate.onDeviceChanged.addListener" directly.
func OnDeviceChanged(callback js.Func[func(event *DeviceEvent)]) (ret js.Void) {
	bindings.CallOnDeviceChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDeviceChanged calls the function "WEBEXT.fileManagerPrivate.onDeviceChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDeviceChanged(callback js.Func[func(event *DeviceEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDeviceChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDeviceChanged returns true if the function "WEBEXT.fileManagerPrivate.onDeviceChanged.removeListener" exists.
func HasFuncOffDeviceChanged() bool {
	return js.True == bindings.HasFuncOffDeviceChanged()
}

// FuncOffDeviceChanged returns the function "WEBEXT.fileManagerPrivate.onDeviceChanged.removeListener".
func FuncOffDeviceChanged() (fn js.Func[func(callback js.Func[func(event *DeviceEvent)])]) {
	bindings.FuncOffDeviceChanged(
		js.Pointer(&fn),
	)
	return
}

// OffDeviceChanged calls the function "WEBEXT.fileManagerPrivate.onDeviceChanged.removeListener" directly.
func OffDeviceChanged(callback js.Func[func(event *DeviceEvent)]) (ret js.Void) {
	bindings.CallOffDeviceChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDeviceChanged calls the function "WEBEXT.fileManagerPrivate.onDeviceChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDeviceChanged(callback js.Func[func(event *DeviceEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDeviceChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDeviceChanged returns true if the function "WEBEXT.fileManagerPrivate.onDeviceChanged.hasListener" exists.
func HasFuncHasOnDeviceChanged() bool {
	return js.True == bindings.HasFuncHasOnDeviceChanged()
}

// FuncHasOnDeviceChanged returns the function "WEBEXT.fileManagerPrivate.onDeviceChanged.hasListener".
func FuncHasOnDeviceChanged() (fn js.Func[func(callback js.Func[func(event *DeviceEvent)]) bool]) {
	bindings.FuncHasOnDeviceChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnDeviceChanged calls the function "WEBEXT.fileManagerPrivate.onDeviceChanged.hasListener" directly.
func HasOnDeviceChanged(callback js.Func[func(event *DeviceEvent)]) (ret bool) {
	bindings.CallHasOnDeviceChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDeviceChanged calls the function "WEBEXT.fileManagerPrivate.onDeviceChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDeviceChanged(callback js.Func[func(event *DeviceEvent)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDeviceChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDeviceConnectionStatusChangedEventCallbackFunc func(this js.Ref, state DeviceConnectionState) js.Ref

func (fn OnDeviceConnectionStatusChangedEventCallbackFunc) Register() js.Func[func(state DeviceConnectionState)] {
	return js.RegisterCallback[func(state DeviceConnectionState)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDeviceConnectionStatusChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		DeviceConnectionState(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnDeviceConnectionStatusChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, state DeviceConnectionState) js.Ref
	Arg T
}

func (cb *OnDeviceConnectionStatusChangedEventCallback[T]) Register() js.Func[func(state DeviceConnectionState)] {
	return js.RegisterCallback[func(state DeviceConnectionState)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDeviceConnectionStatusChangedEventCallback[T]) DispatchCallback(
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

		DeviceConnectionState(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnDeviceConnectionStatusChanged returns true if the function "WEBEXT.fileManagerPrivate.onDeviceConnectionStatusChanged.addListener" exists.
func HasFuncOnDeviceConnectionStatusChanged() bool {
	return js.True == bindings.HasFuncOnDeviceConnectionStatusChanged()
}

// FuncOnDeviceConnectionStatusChanged returns the function "WEBEXT.fileManagerPrivate.onDeviceConnectionStatusChanged.addListener".
func FuncOnDeviceConnectionStatusChanged() (fn js.Func[func(callback js.Func[func(state DeviceConnectionState)])]) {
	bindings.FuncOnDeviceConnectionStatusChanged(
		js.Pointer(&fn),
	)
	return
}

// OnDeviceConnectionStatusChanged calls the function "WEBEXT.fileManagerPrivate.onDeviceConnectionStatusChanged.addListener" directly.
func OnDeviceConnectionStatusChanged(callback js.Func[func(state DeviceConnectionState)]) (ret js.Void) {
	bindings.CallOnDeviceConnectionStatusChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDeviceConnectionStatusChanged calls the function "WEBEXT.fileManagerPrivate.onDeviceConnectionStatusChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDeviceConnectionStatusChanged(callback js.Func[func(state DeviceConnectionState)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDeviceConnectionStatusChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDeviceConnectionStatusChanged returns true if the function "WEBEXT.fileManagerPrivate.onDeviceConnectionStatusChanged.removeListener" exists.
func HasFuncOffDeviceConnectionStatusChanged() bool {
	return js.True == bindings.HasFuncOffDeviceConnectionStatusChanged()
}

// FuncOffDeviceConnectionStatusChanged returns the function "WEBEXT.fileManagerPrivate.onDeviceConnectionStatusChanged.removeListener".
func FuncOffDeviceConnectionStatusChanged() (fn js.Func[func(callback js.Func[func(state DeviceConnectionState)])]) {
	bindings.FuncOffDeviceConnectionStatusChanged(
		js.Pointer(&fn),
	)
	return
}

// OffDeviceConnectionStatusChanged calls the function "WEBEXT.fileManagerPrivate.onDeviceConnectionStatusChanged.removeListener" directly.
func OffDeviceConnectionStatusChanged(callback js.Func[func(state DeviceConnectionState)]) (ret js.Void) {
	bindings.CallOffDeviceConnectionStatusChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDeviceConnectionStatusChanged calls the function "WEBEXT.fileManagerPrivate.onDeviceConnectionStatusChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDeviceConnectionStatusChanged(callback js.Func[func(state DeviceConnectionState)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDeviceConnectionStatusChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDeviceConnectionStatusChanged returns true if the function "WEBEXT.fileManagerPrivate.onDeviceConnectionStatusChanged.hasListener" exists.
func HasFuncHasOnDeviceConnectionStatusChanged() bool {
	return js.True == bindings.HasFuncHasOnDeviceConnectionStatusChanged()
}

// FuncHasOnDeviceConnectionStatusChanged returns the function "WEBEXT.fileManagerPrivate.onDeviceConnectionStatusChanged.hasListener".
func FuncHasOnDeviceConnectionStatusChanged() (fn js.Func[func(callback js.Func[func(state DeviceConnectionState)]) bool]) {
	bindings.FuncHasOnDeviceConnectionStatusChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnDeviceConnectionStatusChanged calls the function "WEBEXT.fileManagerPrivate.onDeviceConnectionStatusChanged.hasListener" directly.
func HasOnDeviceConnectionStatusChanged(callback js.Func[func(state DeviceConnectionState)]) (ret bool) {
	bindings.CallHasOnDeviceConnectionStatusChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDeviceConnectionStatusChanged calls the function "WEBEXT.fileManagerPrivate.onDeviceConnectionStatusChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDeviceConnectionStatusChanged(callback js.Func[func(state DeviceConnectionState)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDeviceConnectionStatusChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDirectoryChangedEventCallbackFunc func(this js.Ref, event *FileWatchEvent) js.Ref

func (fn OnDirectoryChangedEventCallbackFunc) Register() js.Func[func(event *FileWatchEvent)] {
	return js.RegisterCallback[func(event *FileWatchEvent)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDirectoryChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 FileWatchEvent
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

type OnDirectoryChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event *FileWatchEvent) js.Ref
	Arg T
}

func (cb *OnDirectoryChangedEventCallback[T]) Register() js.Func[func(event *FileWatchEvent)] {
	return js.RegisterCallback[func(event *FileWatchEvent)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDirectoryChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 FileWatchEvent
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

// HasFuncOnDirectoryChanged returns true if the function "WEBEXT.fileManagerPrivate.onDirectoryChanged.addListener" exists.
func HasFuncOnDirectoryChanged() bool {
	return js.True == bindings.HasFuncOnDirectoryChanged()
}

// FuncOnDirectoryChanged returns the function "WEBEXT.fileManagerPrivate.onDirectoryChanged.addListener".
func FuncOnDirectoryChanged() (fn js.Func[func(callback js.Func[func(event *FileWatchEvent)])]) {
	bindings.FuncOnDirectoryChanged(
		js.Pointer(&fn),
	)
	return
}

// OnDirectoryChanged calls the function "WEBEXT.fileManagerPrivate.onDirectoryChanged.addListener" directly.
func OnDirectoryChanged(callback js.Func[func(event *FileWatchEvent)]) (ret js.Void) {
	bindings.CallOnDirectoryChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDirectoryChanged calls the function "WEBEXT.fileManagerPrivate.onDirectoryChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDirectoryChanged(callback js.Func[func(event *FileWatchEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDirectoryChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDirectoryChanged returns true if the function "WEBEXT.fileManagerPrivate.onDirectoryChanged.removeListener" exists.
func HasFuncOffDirectoryChanged() bool {
	return js.True == bindings.HasFuncOffDirectoryChanged()
}

// FuncOffDirectoryChanged returns the function "WEBEXT.fileManagerPrivate.onDirectoryChanged.removeListener".
func FuncOffDirectoryChanged() (fn js.Func[func(callback js.Func[func(event *FileWatchEvent)])]) {
	bindings.FuncOffDirectoryChanged(
		js.Pointer(&fn),
	)
	return
}

// OffDirectoryChanged calls the function "WEBEXT.fileManagerPrivate.onDirectoryChanged.removeListener" directly.
func OffDirectoryChanged(callback js.Func[func(event *FileWatchEvent)]) (ret js.Void) {
	bindings.CallOffDirectoryChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDirectoryChanged calls the function "WEBEXT.fileManagerPrivate.onDirectoryChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDirectoryChanged(callback js.Func[func(event *FileWatchEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDirectoryChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDirectoryChanged returns true if the function "WEBEXT.fileManagerPrivate.onDirectoryChanged.hasListener" exists.
func HasFuncHasOnDirectoryChanged() bool {
	return js.True == bindings.HasFuncHasOnDirectoryChanged()
}

// FuncHasOnDirectoryChanged returns the function "WEBEXT.fileManagerPrivate.onDirectoryChanged.hasListener".
func FuncHasOnDirectoryChanged() (fn js.Func[func(callback js.Func[func(event *FileWatchEvent)]) bool]) {
	bindings.FuncHasOnDirectoryChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnDirectoryChanged calls the function "WEBEXT.fileManagerPrivate.onDirectoryChanged.hasListener" directly.
func HasOnDirectoryChanged(callback js.Func[func(event *FileWatchEvent)]) (ret bool) {
	bindings.CallHasOnDirectoryChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDirectoryChanged calls the function "WEBEXT.fileManagerPrivate.onDirectoryChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDirectoryChanged(callback js.Func[func(event *FileWatchEvent)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDirectoryChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDriveConfirmDialogEventCallbackFunc func(this js.Ref, event *DriveConfirmDialogEvent) js.Ref

func (fn OnDriveConfirmDialogEventCallbackFunc) Register() js.Func[func(event *DriveConfirmDialogEvent)] {
	return js.RegisterCallback[func(event *DriveConfirmDialogEvent)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDriveConfirmDialogEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DriveConfirmDialogEvent
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

type OnDriveConfirmDialogEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event *DriveConfirmDialogEvent) js.Ref
	Arg T
}

func (cb *OnDriveConfirmDialogEventCallback[T]) Register() js.Func[func(event *DriveConfirmDialogEvent)] {
	return js.RegisterCallback[func(event *DriveConfirmDialogEvent)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDriveConfirmDialogEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DriveConfirmDialogEvent
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

// HasFuncOnDriveConfirmDialog returns true if the function "WEBEXT.fileManagerPrivate.onDriveConfirmDialog.addListener" exists.
func HasFuncOnDriveConfirmDialog() bool {
	return js.True == bindings.HasFuncOnDriveConfirmDialog()
}

// FuncOnDriveConfirmDialog returns the function "WEBEXT.fileManagerPrivate.onDriveConfirmDialog.addListener".
func FuncOnDriveConfirmDialog() (fn js.Func[func(callback js.Func[func(event *DriveConfirmDialogEvent)])]) {
	bindings.FuncOnDriveConfirmDialog(
		js.Pointer(&fn),
	)
	return
}

// OnDriveConfirmDialog calls the function "WEBEXT.fileManagerPrivate.onDriveConfirmDialog.addListener" directly.
func OnDriveConfirmDialog(callback js.Func[func(event *DriveConfirmDialogEvent)]) (ret js.Void) {
	bindings.CallOnDriveConfirmDialog(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDriveConfirmDialog calls the function "WEBEXT.fileManagerPrivate.onDriveConfirmDialog.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDriveConfirmDialog(callback js.Func[func(event *DriveConfirmDialogEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDriveConfirmDialog(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDriveConfirmDialog returns true if the function "WEBEXT.fileManagerPrivate.onDriveConfirmDialog.removeListener" exists.
func HasFuncOffDriveConfirmDialog() bool {
	return js.True == bindings.HasFuncOffDriveConfirmDialog()
}

// FuncOffDriveConfirmDialog returns the function "WEBEXT.fileManagerPrivate.onDriveConfirmDialog.removeListener".
func FuncOffDriveConfirmDialog() (fn js.Func[func(callback js.Func[func(event *DriveConfirmDialogEvent)])]) {
	bindings.FuncOffDriveConfirmDialog(
		js.Pointer(&fn),
	)
	return
}

// OffDriveConfirmDialog calls the function "WEBEXT.fileManagerPrivate.onDriveConfirmDialog.removeListener" directly.
func OffDriveConfirmDialog(callback js.Func[func(event *DriveConfirmDialogEvent)]) (ret js.Void) {
	bindings.CallOffDriveConfirmDialog(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDriveConfirmDialog calls the function "WEBEXT.fileManagerPrivate.onDriveConfirmDialog.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDriveConfirmDialog(callback js.Func[func(event *DriveConfirmDialogEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDriveConfirmDialog(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDriveConfirmDialog returns true if the function "WEBEXT.fileManagerPrivate.onDriveConfirmDialog.hasListener" exists.
func HasFuncHasOnDriveConfirmDialog() bool {
	return js.True == bindings.HasFuncHasOnDriveConfirmDialog()
}

// FuncHasOnDriveConfirmDialog returns the function "WEBEXT.fileManagerPrivate.onDriveConfirmDialog.hasListener".
func FuncHasOnDriveConfirmDialog() (fn js.Func[func(callback js.Func[func(event *DriveConfirmDialogEvent)]) bool]) {
	bindings.FuncHasOnDriveConfirmDialog(
		js.Pointer(&fn),
	)
	return
}

// HasOnDriveConfirmDialog calls the function "WEBEXT.fileManagerPrivate.onDriveConfirmDialog.hasListener" directly.
func HasOnDriveConfirmDialog(callback js.Func[func(event *DriveConfirmDialogEvent)]) (ret bool) {
	bindings.CallHasOnDriveConfirmDialog(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDriveConfirmDialog calls the function "WEBEXT.fileManagerPrivate.onDriveConfirmDialog.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDriveConfirmDialog(callback js.Func[func(event *DriveConfirmDialogEvent)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDriveConfirmDialog(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDriveConnectionStatusChangedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnDriveConnectionStatusChangedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDriveConnectionStatusChangedEventCallbackFunc) DispatchCallback(
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

type OnDriveConnectionStatusChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnDriveConnectionStatusChangedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDriveConnectionStatusChangedEventCallback[T]) DispatchCallback(
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

// HasFuncOnDriveConnectionStatusChanged returns true if the function "WEBEXT.fileManagerPrivate.onDriveConnectionStatusChanged.addListener" exists.
func HasFuncOnDriveConnectionStatusChanged() bool {
	return js.True == bindings.HasFuncOnDriveConnectionStatusChanged()
}

// FuncOnDriveConnectionStatusChanged returns the function "WEBEXT.fileManagerPrivate.onDriveConnectionStatusChanged.addListener".
func FuncOnDriveConnectionStatusChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnDriveConnectionStatusChanged(
		js.Pointer(&fn),
	)
	return
}

// OnDriveConnectionStatusChanged calls the function "WEBEXT.fileManagerPrivate.onDriveConnectionStatusChanged.addListener" directly.
func OnDriveConnectionStatusChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnDriveConnectionStatusChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDriveConnectionStatusChanged calls the function "WEBEXT.fileManagerPrivate.onDriveConnectionStatusChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDriveConnectionStatusChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDriveConnectionStatusChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDriveConnectionStatusChanged returns true if the function "WEBEXT.fileManagerPrivate.onDriveConnectionStatusChanged.removeListener" exists.
func HasFuncOffDriveConnectionStatusChanged() bool {
	return js.True == bindings.HasFuncOffDriveConnectionStatusChanged()
}

// FuncOffDriveConnectionStatusChanged returns the function "WEBEXT.fileManagerPrivate.onDriveConnectionStatusChanged.removeListener".
func FuncOffDriveConnectionStatusChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffDriveConnectionStatusChanged(
		js.Pointer(&fn),
	)
	return
}

// OffDriveConnectionStatusChanged calls the function "WEBEXT.fileManagerPrivate.onDriveConnectionStatusChanged.removeListener" directly.
func OffDriveConnectionStatusChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffDriveConnectionStatusChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDriveConnectionStatusChanged calls the function "WEBEXT.fileManagerPrivate.onDriveConnectionStatusChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDriveConnectionStatusChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDriveConnectionStatusChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDriveConnectionStatusChanged returns true if the function "WEBEXT.fileManagerPrivate.onDriveConnectionStatusChanged.hasListener" exists.
func HasFuncHasOnDriveConnectionStatusChanged() bool {
	return js.True == bindings.HasFuncHasOnDriveConnectionStatusChanged()
}

// FuncHasOnDriveConnectionStatusChanged returns the function "WEBEXT.fileManagerPrivate.onDriveConnectionStatusChanged.hasListener".
func FuncHasOnDriveConnectionStatusChanged() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnDriveConnectionStatusChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnDriveConnectionStatusChanged calls the function "WEBEXT.fileManagerPrivate.onDriveConnectionStatusChanged.hasListener" directly.
func HasOnDriveConnectionStatusChanged(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnDriveConnectionStatusChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDriveConnectionStatusChanged calls the function "WEBEXT.fileManagerPrivate.onDriveConnectionStatusChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDriveConnectionStatusChanged(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDriveConnectionStatusChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDriveSyncErrorEventCallbackFunc func(this js.Ref, event *DriveSyncErrorEvent) js.Ref

func (fn OnDriveSyncErrorEventCallbackFunc) Register() js.Func[func(event *DriveSyncErrorEvent)] {
	return js.RegisterCallback[func(event *DriveSyncErrorEvent)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDriveSyncErrorEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DriveSyncErrorEvent
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

type OnDriveSyncErrorEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event *DriveSyncErrorEvent) js.Ref
	Arg T
}

func (cb *OnDriveSyncErrorEventCallback[T]) Register() js.Func[func(event *DriveSyncErrorEvent)] {
	return js.RegisterCallback[func(event *DriveSyncErrorEvent)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDriveSyncErrorEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DriveSyncErrorEvent
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

// HasFuncOnDriveSyncError returns true if the function "WEBEXT.fileManagerPrivate.onDriveSyncError.addListener" exists.
func HasFuncOnDriveSyncError() bool {
	return js.True == bindings.HasFuncOnDriveSyncError()
}

// FuncOnDriveSyncError returns the function "WEBEXT.fileManagerPrivate.onDriveSyncError.addListener".
func FuncOnDriveSyncError() (fn js.Func[func(callback js.Func[func(event *DriveSyncErrorEvent)])]) {
	bindings.FuncOnDriveSyncError(
		js.Pointer(&fn),
	)
	return
}

// OnDriveSyncError calls the function "WEBEXT.fileManagerPrivate.onDriveSyncError.addListener" directly.
func OnDriveSyncError(callback js.Func[func(event *DriveSyncErrorEvent)]) (ret js.Void) {
	bindings.CallOnDriveSyncError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDriveSyncError calls the function "WEBEXT.fileManagerPrivate.onDriveSyncError.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDriveSyncError(callback js.Func[func(event *DriveSyncErrorEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDriveSyncError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDriveSyncError returns true if the function "WEBEXT.fileManagerPrivate.onDriveSyncError.removeListener" exists.
func HasFuncOffDriveSyncError() bool {
	return js.True == bindings.HasFuncOffDriveSyncError()
}

// FuncOffDriveSyncError returns the function "WEBEXT.fileManagerPrivate.onDriveSyncError.removeListener".
func FuncOffDriveSyncError() (fn js.Func[func(callback js.Func[func(event *DriveSyncErrorEvent)])]) {
	bindings.FuncOffDriveSyncError(
		js.Pointer(&fn),
	)
	return
}

// OffDriveSyncError calls the function "WEBEXT.fileManagerPrivate.onDriveSyncError.removeListener" directly.
func OffDriveSyncError(callback js.Func[func(event *DriveSyncErrorEvent)]) (ret js.Void) {
	bindings.CallOffDriveSyncError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDriveSyncError calls the function "WEBEXT.fileManagerPrivate.onDriveSyncError.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDriveSyncError(callback js.Func[func(event *DriveSyncErrorEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDriveSyncError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDriveSyncError returns true if the function "WEBEXT.fileManagerPrivate.onDriveSyncError.hasListener" exists.
func HasFuncHasOnDriveSyncError() bool {
	return js.True == bindings.HasFuncHasOnDriveSyncError()
}

// FuncHasOnDriveSyncError returns the function "WEBEXT.fileManagerPrivate.onDriveSyncError.hasListener".
func FuncHasOnDriveSyncError() (fn js.Func[func(callback js.Func[func(event *DriveSyncErrorEvent)]) bool]) {
	bindings.FuncHasOnDriveSyncError(
		js.Pointer(&fn),
	)
	return
}

// HasOnDriveSyncError calls the function "WEBEXT.fileManagerPrivate.onDriveSyncError.hasListener" directly.
func HasOnDriveSyncError(callback js.Func[func(event *DriveSyncErrorEvent)]) (ret bool) {
	bindings.CallHasOnDriveSyncError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDriveSyncError calls the function "WEBEXT.fileManagerPrivate.onDriveSyncError.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDriveSyncError(callback js.Func[func(event *DriveSyncErrorEvent)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDriveSyncError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnFileTransfersUpdatedEventCallbackFunc func(this js.Ref, event *FileTransferStatus) js.Ref

func (fn OnFileTransfersUpdatedEventCallbackFunc) Register() js.Func[func(event *FileTransferStatus)] {
	return js.RegisterCallback[func(event *FileTransferStatus)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnFileTransfersUpdatedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 FileTransferStatus
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

type OnFileTransfersUpdatedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event *FileTransferStatus) js.Ref
	Arg T
}

func (cb *OnFileTransfersUpdatedEventCallback[T]) Register() js.Func[func(event *FileTransferStatus)] {
	return js.RegisterCallback[func(event *FileTransferStatus)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnFileTransfersUpdatedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 FileTransferStatus
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

// HasFuncOnFileTransfersUpdated returns true if the function "WEBEXT.fileManagerPrivate.onFileTransfersUpdated.addListener" exists.
func HasFuncOnFileTransfersUpdated() bool {
	return js.True == bindings.HasFuncOnFileTransfersUpdated()
}

// FuncOnFileTransfersUpdated returns the function "WEBEXT.fileManagerPrivate.onFileTransfersUpdated.addListener".
func FuncOnFileTransfersUpdated() (fn js.Func[func(callback js.Func[func(event *FileTransferStatus)])]) {
	bindings.FuncOnFileTransfersUpdated(
		js.Pointer(&fn),
	)
	return
}

// OnFileTransfersUpdated calls the function "WEBEXT.fileManagerPrivate.onFileTransfersUpdated.addListener" directly.
func OnFileTransfersUpdated(callback js.Func[func(event *FileTransferStatus)]) (ret js.Void) {
	bindings.CallOnFileTransfersUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnFileTransfersUpdated calls the function "WEBEXT.fileManagerPrivate.onFileTransfersUpdated.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnFileTransfersUpdated(callback js.Func[func(event *FileTransferStatus)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnFileTransfersUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffFileTransfersUpdated returns true if the function "WEBEXT.fileManagerPrivate.onFileTransfersUpdated.removeListener" exists.
func HasFuncOffFileTransfersUpdated() bool {
	return js.True == bindings.HasFuncOffFileTransfersUpdated()
}

// FuncOffFileTransfersUpdated returns the function "WEBEXT.fileManagerPrivate.onFileTransfersUpdated.removeListener".
func FuncOffFileTransfersUpdated() (fn js.Func[func(callback js.Func[func(event *FileTransferStatus)])]) {
	bindings.FuncOffFileTransfersUpdated(
		js.Pointer(&fn),
	)
	return
}

// OffFileTransfersUpdated calls the function "WEBEXT.fileManagerPrivate.onFileTransfersUpdated.removeListener" directly.
func OffFileTransfersUpdated(callback js.Func[func(event *FileTransferStatus)]) (ret js.Void) {
	bindings.CallOffFileTransfersUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffFileTransfersUpdated calls the function "WEBEXT.fileManagerPrivate.onFileTransfersUpdated.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffFileTransfersUpdated(callback js.Func[func(event *FileTransferStatus)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffFileTransfersUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnFileTransfersUpdated returns true if the function "WEBEXT.fileManagerPrivate.onFileTransfersUpdated.hasListener" exists.
func HasFuncHasOnFileTransfersUpdated() bool {
	return js.True == bindings.HasFuncHasOnFileTransfersUpdated()
}

// FuncHasOnFileTransfersUpdated returns the function "WEBEXT.fileManagerPrivate.onFileTransfersUpdated.hasListener".
func FuncHasOnFileTransfersUpdated() (fn js.Func[func(callback js.Func[func(event *FileTransferStatus)]) bool]) {
	bindings.FuncHasOnFileTransfersUpdated(
		js.Pointer(&fn),
	)
	return
}

// HasOnFileTransfersUpdated calls the function "WEBEXT.fileManagerPrivate.onFileTransfersUpdated.hasListener" directly.
func HasOnFileTransfersUpdated(callback js.Func[func(event *FileTransferStatus)]) (ret bool) {
	bindings.CallHasOnFileTransfersUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnFileTransfersUpdated calls the function "WEBEXT.fileManagerPrivate.onFileTransfersUpdated.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnFileTransfersUpdated(callback js.Func[func(event *FileTransferStatus)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnFileTransfersUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnIOTaskProgressStatusEventCallbackFunc func(this js.Ref, status *ProgressStatus) js.Ref

func (fn OnIOTaskProgressStatusEventCallbackFunc) Register() js.Func[func(status *ProgressStatus)] {
	return js.RegisterCallback[func(status *ProgressStatus)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnIOTaskProgressStatusEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ProgressStatus
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

type OnIOTaskProgressStatusEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, status *ProgressStatus) js.Ref
	Arg T
}

func (cb *OnIOTaskProgressStatusEventCallback[T]) Register() js.Func[func(status *ProgressStatus)] {
	return js.RegisterCallback[func(status *ProgressStatus)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnIOTaskProgressStatusEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ProgressStatus
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

// HasFuncOnIOTaskProgressStatus returns true if the function "WEBEXT.fileManagerPrivate.onIOTaskProgressStatus.addListener" exists.
func HasFuncOnIOTaskProgressStatus() bool {
	return js.True == bindings.HasFuncOnIOTaskProgressStatus()
}

// FuncOnIOTaskProgressStatus returns the function "WEBEXT.fileManagerPrivate.onIOTaskProgressStatus.addListener".
func FuncOnIOTaskProgressStatus() (fn js.Func[func(callback js.Func[func(status *ProgressStatus)])]) {
	bindings.FuncOnIOTaskProgressStatus(
		js.Pointer(&fn),
	)
	return
}

// OnIOTaskProgressStatus calls the function "WEBEXT.fileManagerPrivate.onIOTaskProgressStatus.addListener" directly.
func OnIOTaskProgressStatus(callback js.Func[func(status *ProgressStatus)]) (ret js.Void) {
	bindings.CallOnIOTaskProgressStatus(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnIOTaskProgressStatus calls the function "WEBEXT.fileManagerPrivate.onIOTaskProgressStatus.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnIOTaskProgressStatus(callback js.Func[func(status *ProgressStatus)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnIOTaskProgressStatus(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffIOTaskProgressStatus returns true if the function "WEBEXT.fileManagerPrivate.onIOTaskProgressStatus.removeListener" exists.
func HasFuncOffIOTaskProgressStatus() bool {
	return js.True == bindings.HasFuncOffIOTaskProgressStatus()
}

// FuncOffIOTaskProgressStatus returns the function "WEBEXT.fileManagerPrivate.onIOTaskProgressStatus.removeListener".
func FuncOffIOTaskProgressStatus() (fn js.Func[func(callback js.Func[func(status *ProgressStatus)])]) {
	bindings.FuncOffIOTaskProgressStatus(
		js.Pointer(&fn),
	)
	return
}

// OffIOTaskProgressStatus calls the function "WEBEXT.fileManagerPrivate.onIOTaskProgressStatus.removeListener" directly.
func OffIOTaskProgressStatus(callback js.Func[func(status *ProgressStatus)]) (ret js.Void) {
	bindings.CallOffIOTaskProgressStatus(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffIOTaskProgressStatus calls the function "WEBEXT.fileManagerPrivate.onIOTaskProgressStatus.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffIOTaskProgressStatus(callback js.Func[func(status *ProgressStatus)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffIOTaskProgressStatus(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnIOTaskProgressStatus returns true if the function "WEBEXT.fileManagerPrivate.onIOTaskProgressStatus.hasListener" exists.
func HasFuncHasOnIOTaskProgressStatus() bool {
	return js.True == bindings.HasFuncHasOnIOTaskProgressStatus()
}

// FuncHasOnIOTaskProgressStatus returns the function "WEBEXT.fileManagerPrivate.onIOTaskProgressStatus.hasListener".
func FuncHasOnIOTaskProgressStatus() (fn js.Func[func(callback js.Func[func(status *ProgressStatus)]) bool]) {
	bindings.FuncHasOnIOTaskProgressStatus(
		js.Pointer(&fn),
	)
	return
}

// HasOnIOTaskProgressStatus calls the function "WEBEXT.fileManagerPrivate.onIOTaskProgressStatus.hasListener" directly.
func HasOnIOTaskProgressStatus(callback js.Func[func(status *ProgressStatus)]) (ret bool) {
	bindings.CallHasOnIOTaskProgressStatus(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnIOTaskProgressStatus calls the function "WEBEXT.fileManagerPrivate.onIOTaskProgressStatus.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnIOTaskProgressStatus(callback js.Func[func(status *ProgressStatus)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnIOTaskProgressStatus(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnIndividualFileTransfersUpdatedEventCallbackFunc func(this js.Ref, event js.Array[SyncState]) js.Ref

func (fn OnIndividualFileTransfersUpdatedEventCallbackFunc) Register() js.Func[func(event js.Array[SyncState])] {
	return js.RegisterCallback[func(event js.Array[SyncState])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnIndividualFileTransfersUpdatedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[SyncState]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnIndividualFileTransfersUpdatedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event js.Array[SyncState]) js.Ref
	Arg T
}

func (cb *OnIndividualFileTransfersUpdatedEventCallback[T]) Register() js.Func[func(event js.Array[SyncState])] {
	return js.RegisterCallback[func(event js.Array[SyncState])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnIndividualFileTransfersUpdatedEventCallback[T]) DispatchCallback(
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

		js.Array[SyncState]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnIndividualFileTransfersUpdated returns true if the function "WEBEXT.fileManagerPrivate.onIndividualFileTransfersUpdated.addListener" exists.
func HasFuncOnIndividualFileTransfersUpdated() bool {
	return js.True == bindings.HasFuncOnIndividualFileTransfersUpdated()
}

// FuncOnIndividualFileTransfersUpdated returns the function "WEBEXT.fileManagerPrivate.onIndividualFileTransfersUpdated.addListener".
func FuncOnIndividualFileTransfersUpdated() (fn js.Func[func(callback js.Func[func(event js.Array[SyncState])])]) {
	bindings.FuncOnIndividualFileTransfersUpdated(
		js.Pointer(&fn),
	)
	return
}

// OnIndividualFileTransfersUpdated calls the function "WEBEXT.fileManagerPrivate.onIndividualFileTransfersUpdated.addListener" directly.
func OnIndividualFileTransfersUpdated(callback js.Func[func(event js.Array[SyncState])]) (ret js.Void) {
	bindings.CallOnIndividualFileTransfersUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnIndividualFileTransfersUpdated calls the function "WEBEXT.fileManagerPrivate.onIndividualFileTransfersUpdated.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnIndividualFileTransfersUpdated(callback js.Func[func(event js.Array[SyncState])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnIndividualFileTransfersUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffIndividualFileTransfersUpdated returns true if the function "WEBEXT.fileManagerPrivate.onIndividualFileTransfersUpdated.removeListener" exists.
func HasFuncOffIndividualFileTransfersUpdated() bool {
	return js.True == bindings.HasFuncOffIndividualFileTransfersUpdated()
}

// FuncOffIndividualFileTransfersUpdated returns the function "WEBEXT.fileManagerPrivate.onIndividualFileTransfersUpdated.removeListener".
func FuncOffIndividualFileTransfersUpdated() (fn js.Func[func(callback js.Func[func(event js.Array[SyncState])])]) {
	bindings.FuncOffIndividualFileTransfersUpdated(
		js.Pointer(&fn),
	)
	return
}

// OffIndividualFileTransfersUpdated calls the function "WEBEXT.fileManagerPrivate.onIndividualFileTransfersUpdated.removeListener" directly.
func OffIndividualFileTransfersUpdated(callback js.Func[func(event js.Array[SyncState])]) (ret js.Void) {
	bindings.CallOffIndividualFileTransfersUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffIndividualFileTransfersUpdated calls the function "WEBEXT.fileManagerPrivate.onIndividualFileTransfersUpdated.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffIndividualFileTransfersUpdated(callback js.Func[func(event js.Array[SyncState])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffIndividualFileTransfersUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnIndividualFileTransfersUpdated returns true if the function "WEBEXT.fileManagerPrivate.onIndividualFileTransfersUpdated.hasListener" exists.
func HasFuncHasOnIndividualFileTransfersUpdated() bool {
	return js.True == bindings.HasFuncHasOnIndividualFileTransfersUpdated()
}

// FuncHasOnIndividualFileTransfersUpdated returns the function "WEBEXT.fileManagerPrivate.onIndividualFileTransfersUpdated.hasListener".
func FuncHasOnIndividualFileTransfersUpdated() (fn js.Func[func(callback js.Func[func(event js.Array[SyncState])]) bool]) {
	bindings.FuncHasOnIndividualFileTransfersUpdated(
		js.Pointer(&fn),
	)
	return
}

// HasOnIndividualFileTransfersUpdated calls the function "WEBEXT.fileManagerPrivate.onIndividualFileTransfersUpdated.hasListener" directly.
func HasOnIndividualFileTransfersUpdated(callback js.Func[func(event js.Array[SyncState])]) (ret bool) {
	bindings.CallHasOnIndividualFileTransfersUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnIndividualFileTransfersUpdated calls the function "WEBEXT.fileManagerPrivate.onIndividualFileTransfersUpdated.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnIndividualFileTransfersUpdated(callback js.Func[func(event js.Array[SyncState])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnIndividualFileTransfersUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnMountCompletedEventCallbackFunc func(this js.Ref, event *MountCompletedEvent) js.Ref

func (fn OnMountCompletedEventCallbackFunc) Register() js.Func[func(event *MountCompletedEvent)] {
	return js.RegisterCallback[func(event *MountCompletedEvent)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnMountCompletedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 MountCompletedEvent
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

type OnMountCompletedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event *MountCompletedEvent) js.Ref
	Arg T
}

func (cb *OnMountCompletedEventCallback[T]) Register() js.Func[func(event *MountCompletedEvent)] {
	return js.RegisterCallback[func(event *MountCompletedEvent)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnMountCompletedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 MountCompletedEvent
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

// HasFuncOnMountCompleted returns true if the function "WEBEXT.fileManagerPrivate.onMountCompleted.addListener" exists.
func HasFuncOnMountCompleted() bool {
	return js.True == bindings.HasFuncOnMountCompleted()
}

// FuncOnMountCompleted returns the function "WEBEXT.fileManagerPrivate.onMountCompleted.addListener".
func FuncOnMountCompleted() (fn js.Func[func(callback js.Func[func(event *MountCompletedEvent)])]) {
	bindings.FuncOnMountCompleted(
		js.Pointer(&fn),
	)
	return
}

// OnMountCompleted calls the function "WEBEXT.fileManagerPrivate.onMountCompleted.addListener" directly.
func OnMountCompleted(callback js.Func[func(event *MountCompletedEvent)]) (ret js.Void) {
	bindings.CallOnMountCompleted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMountCompleted calls the function "WEBEXT.fileManagerPrivate.onMountCompleted.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMountCompleted(callback js.Func[func(event *MountCompletedEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMountCompleted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMountCompleted returns true if the function "WEBEXT.fileManagerPrivate.onMountCompleted.removeListener" exists.
func HasFuncOffMountCompleted() bool {
	return js.True == bindings.HasFuncOffMountCompleted()
}

// FuncOffMountCompleted returns the function "WEBEXT.fileManagerPrivate.onMountCompleted.removeListener".
func FuncOffMountCompleted() (fn js.Func[func(callback js.Func[func(event *MountCompletedEvent)])]) {
	bindings.FuncOffMountCompleted(
		js.Pointer(&fn),
	)
	return
}

// OffMountCompleted calls the function "WEBEXT.fileManagerPrivate.onMountCompleted.removeListener" directly.
func OffMountCompleted(callback js.Func[func(event *MountCompletedEvent)]) (ret js.Void) {
	bindings.CallOffMountCompleted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMountCompleted calls the function "WEBEXT.fileManagerPrivate.onMountCompleted.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMountCompleted(callback js.Func[func(event *MountCompletedEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMountCompleted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMountCompleted returns true if the function "WEBEXT.fileManagerPrivate.onMountCompleted.hasListener" exists.
func HasFuncHasOnMountCompleted() bool {
	return js.True == bindings.HasFuncHasOnMountCompleted()
}

// FuncHasOnMountCompleted returns the function "WEBEXT.fileManagerPrivate.onMountCompleted.hasListener".
func FuncHasOnMountCompleted() (fn js.Func[func(callback js.Func[func(event *MountCompletedEvent)]) bool]) {
	bindings.FuncHasOnMountCompleted(
		js.Pointer(&fn),
	)
	return
}

// HasOnMountCompleted calls the function "WEBEXT.fileManagerPrivate.onMountCompleted.hasListener" directly.
func HasOnMountCompleted(callback js.Func[func(event *MountCompletedEvent)]) (ret bool) {
	bindings.CallHasOnMountCompleted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMountCompleted calls the function "WEBEXT.fileManagerPrivate.onMountCompleted.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnMountCompleted(callback js.Func[func(event *MountCompletedEvent)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnMountCompleted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnMountableGuestsChangedEventCallbackFunc func(this js.Ref, guests js.Array[MountableGuest]) js.Ref

func (fn OnMountableGuestsChangedEventCallbackFunc) Register() js.Func[func(guests js.Array[MountableGuest])] {
	return js.RegisterCallback[func(guests js.Array[MountableGuest])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnMountableGuestsChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[MountableGuest]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnMountableGuestsChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, guests js.Array[MountableGuest]) js.Ref
	Arg T
}

func (cb *OnMountableGuestsChangedEventCallback[T]) Register() js.Func[func(guests js.Array[MountableGuest])] {
	return js.RegisterCallback[func(guests js.Array[MountableGuest])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnMountableGuestsChangedEventCallback[T]) DispatchCallback(
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

		js.Array[MountableGuest]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnMountableGuestsChanged returns true if the function "WEBEXT.fileManagerPrivate.onMountableGuestsChanged.addListener" exists.
func HasFuncOnMountableGuestsChanged() bool {
	return js.True == bindings.HasFuncOnMountableGuestsChanged()
}

// FuncOnMountableGuestsChanged returns the function "WEBEXT.fileManagerPrivate.onMountableGuestsChanged.addListener".
func FuncOnMountableGuestsChanged() (fn js.Func[func(callback js.Func[func(guests js.Array[MountableGuest])])]) {
	bindings.FuncOnMountableGuestsChanged(
		js.Pointer(&fn),
	)
	return
}

// OnMountableGuestsChanged calls the function "WEBEXT.fileManagerPrivate.onMountableGuestsChanged.addListener" directly.
func OnMountableGuestsChanged(callback js.Func[func(guests js.Array[MountableGuest])]) (ret js.Void) {
	bindings.CallOnMountableGuestsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMountableGuestsChanged calls the function "WEBEXT.fileManagerPrivate.onMountableGuestsChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMountableGuestsChanged(callback js.Func[func(guests js.Array[MountableGuest])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMountableGuestsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMountableGuestsChanged returns true if the function "WEBEXT.fileManagerPrivate.onMountableGuestsChanged.removeListener" exists.
func HasFuncOffMountableGuestsChanged() bool {
	return js.True == bindings.HasFuncOffMountableGuestsChanged()
}

// FuncOffMountableGuestsChanged returns the function "WEBEXT.fileManagerPrivate.onMountableGuestsChanged.removeListener".
func FuncOffMountableGuestsChanged() (fn js.Func[func(callback js.Func[func(guests js.Array[MountableGuest])])]) {
	bindings.FuncOffMountableGuestsChanged(
		js.Pointer(&fn),
	)
	return
}

// OffMountableGuestsChanged calls the function "WEBEXT.fileManagerPrivate.onMountableGuestsChanged.removeListener" directly.
func OffMountableGuestsChanged(callback js.Func[func(guests js.Array[MountableGuest])]) (ret js.Void) {
	bindings.CallOffMountableGuestsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMountableGuestsChanged calls the function "WEBEXT.fileManagerPrivate.onMountableGuestsChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMountableGuestsChanged(callback js.Func[func(guests js.Array[MountableGuest])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMountableGuestsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMountableGuestsChanged returns true if the function "WEBEXT.fileManagerPrivate.onMountableGuestsChanged.hasListener" exists.
func HasFuncHasOnMountableGuestsChanged() bool {
	return js.True == bindings.HasFuncHasOnMountableGuestsChanged()
}

// FuncHasOnMountableGuestsChanged returns the function "WEBEXT.fileManagerPrivate.onMountableGuestsChanged.hasListener".
func FuncHasOnMountableGuestsChanged() (fn js.Func[func(callback js.Func[func(guests js.Array[MountableGuest])]) bool]) {
	bindings.FuncHasOnMountableGuestsChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnMountableGuestsChanged calls the function "WEBEXT.fileManagerPrivate.onMountableGuestsChanged.hasListener" directly.
func HasOnMountableGuestsChanged(callback js.Func[func(guests js.Array[MountableGuest])]) (ret bool) {
	bindings.CallHasOnMountableGuestsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMountableGuestsChanged calls the function "WEBEXT.fileManagerPrivate.onMountableGuestsChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnMountableGuestsChanged(callback js.Func[func(guests js.Array[MountableGuest])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnMountableGuestsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnPinTransfersUpdatedEventCallbackFunc func(this js.Ref, event *FileTransferStatus) js.Ref

func (fn OnPinTransfersUpdatedEventCallbackFunc) Register() js.Func[func(event *FileTransferStatus)] {
	return js.RegisterCallback[func(event *FileTransferStatus)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPinTransfersUpdatedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 FileTransferStatus
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

type OnPinTransfersUpdatedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event *FileTransferStatus) js.Ref
	Arg T
}

func (cb *OnPinTransfersUpdatedEventCallback[T]) Register() js.Func[func(event *FileTransferStatus)] {
	return js.RegisterCallback[func(event *FileTransferStatus)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPinTransfersUpdatedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 FileTransferStatus
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

// HasFuncOnPinTransfersUpdated returns true if the function "WEBEXT.fileManagerPrivate.onPinTransfersUpdated.addListener" exists.
func HasFuncOnPinTransfersUpdated() bool {
	return js.True == bindings.HasFuncOnPinTransfersUpdated()
}

// FuncOnPinTransfersUpdated returns the function "WEBEXT.fileManagerPrivate.onPinTransfersUpdated.addListener".
func FuncOnPinTransfersUpdated() (fn js.Func[func(callback js.Func[func(event *FileTransferStatus)])]) {
	bindings.FuncOnPinTransfersUpdated(
		js.Pointer(&fn),
	)
	return
}

// OnPinTransfersUpdated calls the function "WEBEXT.fileManagerPrivate.onPinTransfersUpdated.addListener" directly.
func OnPinTransfersUpdated(callback js.Func[func(event *FileTransferStatus)]) (ret js.Void) {
	bindings.CallOnPinTransfersUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPinTransfersUpdated calls the function "WEBEXT.fileManagerPrivate.onPinTransfersUpdated.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPinTransfersUpdated(callback js.Func[func(event *FileTransferStatus)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPinTransfersUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPinTransfersUpdated returns true if the function "WEBEXT.fileManagerPrivate.onPinTransfersUpdated.removeListener" exists.
func HasFuncOffPinTransfersUpdated() bool {
	return js.True == bindings.HasFuncOffPinTransfersUpdated()
}

// FuncOffPinTransfersUpdated returns the function "WEBEXT.fileManagerPrivate.onPinTransfersUpdated.removeListener".
func FuncOffPinTransfersUpdated() (fn js.Func[func(callback js.Func[func(event *FileTransferStatus)])]) {
	bindings.FuncOffPinTransfersUpdated(
		js.Pointer(&fn),
	)
	return
}

// OffPinTransfersUpdated calls the function "WEBEXT.fileManagerPrivate.onPinTransfersUpdated.removeListener" directly.
func OffPinTransfersUpdated(callback js.Func[func(event *FileTransferStatus)]) (ret js.Void) {
	bindings.CallOffPinTransfersUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPinTransfersUpdated calls the function "WEBEXT.fileManagerPrivate.onPinTransfersUpdated.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPinTransfersUpdated(callback js.Func[func(event *FileTransferStatus)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPinTransfersUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPinTransfersUpdated returns true if the function "WEBEXT.fileManagerPrivate.onPinTransfersUpdated.hasListener" exists.
func HasFuncHasOnPinTransfersUpdated() bool {
	return js.True == bindings.HasFuncHasOnPinTransfersUpdated()
}

// FuncHasOnPinTransfersUpdated returns the function "WEBEXT.fileManagerPrivate.onPinTransfersUpdated.hasListener".
func FuncHasOnPinTransfersUpdated() (fn js.Func[func(callback js.Func[func(event *FileTransferStatus)]) bool]) {
	bindings.FuncHasOnPinTransfersUpdated(
		js.Pointer(&fn),
	)
	return
}

// HasOnPinTransfersUpdated calls the function "WEBEXT.fileManagerPrivate.onPinTransfersUpdated.hasListener" directly.
func HasOnPinTransfersUpdated(callback js.Func[func(event *FileTransferStatus)]) (ret bool) {
	bindings.CallHasOnPinTransfersUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPinTransfersUpdated calls the function "WEBEXT.fileManagerPrivate.onPinTransfersUpdated.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPinTransfersUpdated(callback js.Func[func(event *FileTransferStatus)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPinTransfersUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnPreferencesChangedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnPreferencesChangedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPreferencesChangedEventCallbackFunc) DispatchCallback(
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

type OnPreferencesChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnPreferencesChangedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPreferencesChangedEventCallback[T]) DispatchCallback(
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

// HasFuncOnPreferencesChanged returns true if the function "WEBEXT.fileManagerPrivate.onPreferencesChanged.addListener" exists.
func HasFuncOnPreferencesChanged() bool {
	return js.True == bindings.HasFuncOnPreferencesChanged()
}

// FuncOnPreferencesChanged returns the function "WEBEXT.fileManagerPrivate.onPreferencesChanged.addListener".
func FuncOnPreferencesChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnPreferencesChanged(
		js.Pointer(&fn),
	)
	return
}

// OnPreferencesChanged calls the function "WEBEXT.fileManagerPrivate.onPreferencesChanged.addListener" directly.
func OnPreferencesChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnPreferencesChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPreferencesChanged calls the function "WEBEXT.fileManagerPrivate.onPreferencesChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPreferencesChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPreferencesChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPreferencesChanged returns true if the function "WEBEXT.fileManagerPrivate.onPreferencesChanged.removeListener" exists.
func HasFuncOffPreferencesChanged() bool {
	return js.True == bindings.HasFuncOffPreferencesChanged()
}

// FuncOffPreferencesChanged returns the function "WEBEXT.fileManagerPrivate.onPreferencesChanged.removeListener".
func FuncOffPreferencesChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffPreferencesChanged(
		js.Pointer(&fn),
	)
	return
}

// OffPreferencesChanged calls the function "WEBEXT.fileManagerPrivate.onPreferencesChanged.removeListener" directly.
func OffPreferencesChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffPreferencesChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPreferencesChanged calls the function "WEBEXT.fileManagerPrivate.onPreferencesChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPreferencesChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPreferencesChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPreferencesChanged returns true if the function "WEBEXT.fileManagerPrivate.onPreferencesChanged.hasListener" exists.
func HasFuncHasOnPreferencesChanged() bool {
	return js.True == bindings.HasFuncHasOnPreferencesChanged()
}

// FuncHasOnPreferencesChanged returns the function "WEBEXT.fileManagerPrivate.onPreferencesChanged.hasListener".
func FuncHasOnPreferencesChanged() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnPreferencesChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnPreferencesChanged calls the function "WEBEXT.fileManagerPrivate.onPreferencesChanged.hasListener" directly.
func HasOnPreferencesChanged(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnPreferencesChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPreferencesChanged calls the function "WEBEXT.fileManagerPrivate.onPreferencesChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPreferencesChanged(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPreferencesChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnTabletModeChangedEventCallbackFunc func(this js.Ref, enabled bool) js.Ref

func (fn OnTabletModeChangedEventCallbackFunc) Register() js.Func[func(enabled bool)] {
	return js.RegisterCallback[func(enabled bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnTabletModeChangedEventCallbackFunc) DispatchCallback(
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

type OnTabletModeChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, enabled bool) js.Ref
	Arg T
}

func (cb *OnTabletModeChangedEventCallback[T]) Register() js.Func[func(enabled bool)] {
	return js.RegisterCallback[func(enabled bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnTabletModeChangedEventCallback[T]) DispatchCallback(
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

// HasFuncOnTabletModeChanged returns true if the function "WEBEXT.fileManagerPrivate.onTabletModeChanged.addListener" exists.
func HasFuncOnTabletModeChanged() bool {
	return js.True == bindings.HasFuncOnTabletModeChanged()
}

// FuncOnTabletModeChanged returns the function "WEBEXT.fileManagerPrivate.onTabletModeChanged.addListener".
func FuncOnTabletModeChanged() (fn js.Func[func(callback js.Func[func(enabled bool)])]) {
	bindings.FuncOnTabletModeChanged(
		js.Pointer(&fn),
	)
	return
}

// OnTabletModeChanged calls the function "WEBEXT.fileManagerPrivate.onTabletModeChanged.addListener" directly.
func OnTabletModeChanged(callback js.Func[func(enabled bool)]) (ret js.Void) {
	bindings.CallOnTabletModeChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnTabletModeChanged calls the function "WEBEXT.fileManagerPrivate.onTabletModeChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnTabletModeChanged(callback js.Func[func(enabled bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnTabletModeChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffTabletModeChanged returns true if the function "WEBEXT.fileManagerPrivate.onTabletModeChanged.removeListener" exists.
func HasFuncOffTabletModeChanged() bool {
	return js.True == bindings.HasFuncOffTabletModeChanged()
}

// FuncOffTabletModeChanged returns the function "WEBEXT.fileManagerPrivate.onTabletModeChanged.removeListener".
func FuncOffTabletModeChanged() (fn js.Func[func(callback js.Func[func(enabled bool)])]) {
	bindings.FuncOffTabletModeChanged(
		js.Pointer(&fn),
	)
	return
}

// OffTabletModeChanged calls the function "WEBEXT.fileManagerPrivate.onTabletModeChanged.removeListener" directly.
func OffTabletModeChanged(callback js.Func[func(enabled bool)]) (ret js.Void) {
	bindings.CallOffTabletModeChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffTabletModeChanged calls the function "WEBEXT.fileManagerPrivate.onTabletModeChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffTabletModeChanged(callback js.Func[func(enabled bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffTabletModeChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnTabletModeChanged returns true if the function "WEBEXT.fileManagerPrivate.onTabletModeChanged.hasListener" exists.
func HasFuncHasOnTabletModeChanged() bool {
	return js.True == bindings.HasFuncHasOnTabletModeChanged()
}

// FuncHasOnTabletModeChanged returns the function "WEBEXT.fileManagerPrivate.onTabletModeChanged.hasListener".
func FuncHasOnTabletModeChanged() (fn js.Func[func(callback js.Func[func(enabled bool)]) bool]) {
	bindings.FuncHasOnTabletModeChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnTabletModeChanged calls the function "WEBEXT.fileManagerPrivate.onTabletModeChanged.hasListener" directly.
func HasOnTabletModeChanged(callback js.Func[func(enabled bool)]) (ret bool) {
	bindings.CallHasOnTabletModeChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnTabletModeChanged calls the function "WEBEXT.fileManagerPrivate.onTabletModeChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnTabletModeChanged(callback js.Func[func(enabled bool)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnTabletModeChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOpenInspector returns true if the function "WEBEXT.fileManagerPrivate.openInspector" exists.
func HasFuncOpenInspector() bool {
	return js.True == bindings.HasFuncOpenInspector()
}

// FuncOpenInspector returns the function "WEBEXT.fileManagerPrivate.openInspector".
func FuncOpenInspector() (fn js.Func[func(typ InspectionType)]) {
	bindings.FuncOpenInspector(
		js.Pointer(&fn),
	)
	return
}

// OpenInspector calls the function "WEBEXT.fileManagerPrivate.openInspector" directly.
func OpenInspector(typ InspectionType) (ret js.Void) {
	bindings.CallOpenInspector(
		js.Pointer(&ret),
		uint32(typ),
	)

	return
}

// TryOpenInspector calls the function "WEBEXT.fileManagerPrivate.openInspector"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenInspector(typ InspectionType) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenInspector(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(typ),
	)

	return
}

// HasFuncOpenManageSyncSettings returns true if the function "WEBEXT.fileManagerPrivate.openManageSyncSettings" exists.
func HasFuncOpenManageSyncSettings() bool {
	return js.True == bindings.HasFuncOpenManageSyncSettings()
}

// FuncOpenManageSyncSettings returns the function "WEBEXT.fileManagerPrivate.openManageSyncSettings".
func FuncOpenManageSyncSettings() (fn js.Func[func()]) {
	bindings.FuncOpenManageSyncSettings(
		js.Pointer(&fn),
	)
	return
}

// OpenManageSyncSettings calls the function "WEBEXT.fileManagerPrivate.openManageSyncSettings" directly.
func OpenManageSyncSettings() (ret js.Void) {
	bindings.CallOpenManageSyncSettings(
		js.Pointer(&ret),
	)

	return
}

// TryOpenManageSyncSettings calls the function "WEBEXT.fileManagerPrivate.openManageSyncSettings"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenManageSyncSettings() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenManageSyncSettings(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncOpenSettingsSubpage returns true if the function "WEBEXT.fileManagerPrivate.openSettingsSubpage" exists.
func HasFuncOpenSettingsSubpage() bool {
	return js.True == bindings.HasFuncOpenSettingsSubpage()
}

// FuncOpenSettingsSubpage returns the function "WEBEXT.fileManagerPrivate.openSettingsSubpage".
func FuncOpenSettingsSubpage() (fn js.Func[func(sub_page js.String)]) {
	bindings.FuncOpenSettingsSubpage(
		js.Pointer(&fn),
	)
	return
}

// OpenSettingsSubpage calls the function "WEBEXT.fileManagerPrivate.openSettingsSubpage" directly.
func OpenSettingsSubpage(sub_page js.String) (ret js.Void) {
	bindings.CallOpenSettingsSubpage(
		js.Pointer(&ret),
		sub_page.Ref(),
	)

	return
}

// TryOpenSettingsSubpage calls the function "WEBEXT.fileManagerPrivate.openSettingsSubpage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenSettingsSubpage(sub_page js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenSettingsSubpage(
		js.Pointer(&ret), js.Pointer(&exception),
		sub_page.Ref(),
	)

	return
}

// HasFuncOpenURL returns true if the function "WEBEXT.fileManagerPrivate.openURL" exists.
func HasFuncOpenURL() bool {
	return js.True == bindings.HasFuncOpenURL()
}

// FuncOpenURL returns the function "WEBEXT.fileManagerPrivate.openURL".
func FuncOpenURL() (fn js.Func[func(url js.String)]) {
	bindings.FuncOpenURL(
		js.Pointer(&fn),
	)
	return
}

// OpenURL calls the function "WEBEXT.fileManagerPrivate.openURL" directly.
func OpenURL(url js.String) (ret js.Void) {
	bindings.CallOpenURL(
		js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryOpenURL calls the function "WEBEXT.fileManagerPrivate.openURL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenURL(url js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenURL(
		js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasFuncOpenWindow returns true if the function "WEBEXT.fileManagerPrivate.openWindow" exists.
func HasFuncOpenWindow() bool {
	return js.True == bindings.HasFuncOpenWindow()
}

// FuncOpenWindow returns the function "WEBEXT.fileManagerPrivate.openWindow".
func FuncOpenWindow() (fn js.Func[func(params OpenWindowParams) js.Promise[js.Boolean]]) {
	bindings.FuncOpenWindow(
		js.Pointer(&fn),
	)
	return
}

// OpenWindow calls the function "WEBEXT.fileManagerPrivate.openWindow" directly.
func OpenWindow(params OpenWindowParams) (ret js.Promise[js.Boolean]) {
	bindings.CallOpenWindow(
		js.Pointer(&ret),
		js.Pointer(&params),
	)

	return
}

// TryOpenWindow calls the function "WEBEXT.fileManagerPrivate.openWindow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenWindow(params OpenWindowParams) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenWindow(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&params),
	)

	return
}

// HasFuncParseTrashInfoFiles returns true if the function "WEBEXT.fileManagerPrivate.parseTrashInfoFiles" exists.
func HasFuncParseTrashInfoFiles() bool {
	return js.True == bindings.HasFuncParseTrashInfoFiles()
}

// FuncParseTrashInfoFiles returns the function "WEBEXT.fileManagerPrivate.parseTrashInfoFiles".
func FuncParseTrashInfoFiles() (fn js.Func[func(entries js.Array[js.Object]) js.Promise[js.Array[ParsedTrashInfoFile]]]) {
	bindings.FuncParseTrashInfoFiles(
		js.Pointer(&fn),
	)
	return
}

// ParseTrashInfoFiles calls the function "WEBEXT.fileManagerPrivate.parseTrashInfoFiles" directly.
func ParseTrashInfoFiles(entries js.Array[js.Object]) (ret js.Promise[js.Array[ParsedTrashInfoFile]]) {
	bindings.CallParseTrashInfoFiles(
		js.Pointer(&ret),
		entries.Ref(),
	)

	return
}

// TryParseTrashInfoFiles calls the function "WEBEXT.fileManagerPrivate.parseTrashInfoFiles"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryParseTrashInfoFiles(entries js.Array[js.Object]) (ret js.Promise[js.Array[ParsedTrashInfoFile]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryParseTrashInfoFiles(
		js.Pointer(&ret), js.Pointer(&exception),
		entries.Ref(),
	)

	return
}

// HasFuncPinDriveFile returns true if the function "WEBEXT.fileManagerPrivate.pinDriveFile" exists.
func HasFuncPinDriveFile() bool {
	return js.True == bindings.HasFuncPinDriveFile()
}

// FuncPinDriveFile returns the function "WEBEXT.fileManagerPrivate.pinDriveFile".
func FuncPinDriveFile() (fn js.Func[func(entry js.Object, pin bool) js.Promise[js.Void]]) {
	bindings.FuncPinDriveFile(
		js.Pointer(&fn),
	)
	return
}

// PinDriveFile calls the function "WEBEXT.fileManagerPrivate.pinDriveFile" directly.
func PinDriveFile(entry js.Object, pin bool) (ret js.Promise[js.Void]) {
	bindings.CallPinDriveFile(
		js.Pointer(&ret),
		entry.Ref(),
		js.Bool(bool(pin)),
	)

	return
}

// TryPinDriveFile calls the function "WEBEXT.fileManagerPrivate.pinDriveFile"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryPinDriveFile(entry js.Object, pin bool) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPinDriveFile(
		js.Pointer(&ret), js.Pointer(&exception),
		entry.Ref(),
		js.Bool(bool(pin)),
	)

	return
}

// HasFuncPollDriveHostedFilePinStates returns true if the function "WEBEXT.fileManagerPrivate.pollDriveHostedFilePinStates" exists.
func HasFuncPollDriveHostedFilePinStates() bool {
	return js.True == bindings.HasFuncPollDriveHostedFilePinStates()
}

// FuncPollDriveHostedFilePinStates returns the function "WEBEXT.fileManagerPrivate.pollDriveHostedFilePinStates".
func FuncPollDriveHostedFilePinStates() (fn js.Func[func()]) {
	bindings.FuncPollDriveHostedFilePinStates(
		js.Pointer(&fn),
	)
	return
}

// PollDriveHostedFilePinStates calls the function "WEBEXT.fileManagerPrivate.pollDriveHostedFilePinStates" directly.
func PollDriveHostedFilePinStates() (ret js.Void) {
	bindings.CallPollDriveHostedFilePinStates(
		js.Pointer(&ret),
	)

	return
}

// TryPollDriveHostedFilePinStates calls the function "WEBEXT.fileManagerPrivate.pollDriveHostedFilePinStates"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryPollDriveHostedFilePinStates() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPollDriveHostedFilePinStates(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncProgressPausedTasks returns true if the function "WEBEXT.fileManagerPrivate.progressPausedTasks" exists.
func HasFuncProgressPausedTasks() bool {
	return js.True == bindings.HasFuncProgressPausedTasks()
}

// FuncProgressPausedTasks returns the function "WEBEXT.fileManagerPrivate.progressPausedTasks".
func FuncProgressPausedTasks() (fn js.Func[func()]) {
	bindings.FuncProgressPausedTasks(
		js.Pointer(&fn),
	)
	return
}

// ProgressPausedTasks calls the function "WEBEXT.fileManagerPrivate.progressPausedTasks" directly.
func ProgressPausedTasks() (ret js.Void) {
	bindings.CallProgressPausedTasks(
		js.Pointer(&ret),
	)

	return
}

// TryProgressPausedTasks calls the function "WEBEXT.fileManagerPrivate.progressPausedTasks"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryProgressPausedTasks() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryProgressPausedTasks(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRemoveFileWatch returns true if the function "WEBEXT.fileManagerPrivate.removeFileWatch" exists.
func HasFuncRemoveFileWatch() bool {
	return js.True == bindings.HasFuncRemoveFileWatch()
}

// FuncRemoveFileWatch returns the function "WEBEXT.fileManagerPrivate.removeFileWatch".
func FuncRemoveFileWatch() (fn js.Func[func(entry js.Object) js.Promise[js.Boolean]]) {
	bindings.FuncRemoveFileWatch(
		js.Pointer(&fn),
	)
	return
}

// RemoveFileWatch calls the function "WEBEXT.fileManagerPrivate.removeFileWatch" directly.
func RemoveFileWatch(entry js.Object) (ret js.Promise[js.Boolean]) {
	bindings.CallRemoveFileWatch(
		js.Pointer(&ret),
		entry.Ref(),
	)

	return
}

// TryRemoveFileWatch calls the function "WEBEXT.fileManagerPrivate.removeFileWatch"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveFileWatch(entry js.Object) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveFileWatch(
		js.Pointer(&ret), js.Pointer(&exception),
		entry.Ref(),
	)

	return
}

// HasFuncRemoveMount returns true if the function "WEBEXT.fileManagerPrivate.removeMount" exists.
func HasFuncRemoveMount() bool {
	return js.True == bindings.HasFuncRemoveMount()
}

// FuncRemoveMount returns the function "WEBEXT.fileManagerPrivate.removeMount".
func FuncRemoveMount() (fn js.Func[func(volumeId js.String) js.Promise[js.Void]]) {
	bindings.FuncRemoveMount(
		js.Pointer(&fn),
	)
	return
}

// RemoveMount calls the function "WEBEXT.fileManagerPrivate.removeMount" directly.
func RemoveMount(volumeId js.String) (ret js.Promise[js.Void]) {
	bindings.CallRemoveMount(
		js.Pointer(&ret),
		volumeId.Ref(),
	)

	return
}

// TryRemoveMount calls the function "WEBEXT.fileManagerPrivate.removeMount"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveMount(volumeId js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveMount(
		js.Pointer(&ret), js.Pointer(&exception),
		volumeId.Ref(),
	)

	return
}

// HasFuncRenameVolume returns true if the function "WEBEXT.fileManagerPrivate.renameVolume" exists.
func HasFuncRenameVolume() bool {
	return js.True == bindings.HasFuncRenameVolume()
}

// FuncRenameVolume returns the function "WEBEXT.fileManagerPrivate.renameVolume".
func FuncRenameVolume() (fn js.Func[func(volumeId js.String, newName js.String)]) {
	bindings.FuncRenameVolume(
		js.Pointer(&fn),
	)
	return
}

// RenameVolume calls the function "WEBEXT.fileManagerPrivate.renameVolume" directly.
func RenameVolume(volumeId js.String, newName js.String) (ret js.Void) {
	bindings.CallRenameVolume(
		js.Pointer(&ret),
		volumeId.Ref(),
		newName.Ref(),
	)

	return
}

// TryRenameVolume calls the function "WEBEXT.fileManagerPrivate.renameVolume"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRenameVolume(volumeId js.String, newName js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRenameVolume(
		js.Pointer(&ret), js.Pointer(&exception),
		volumeId.Ref(),
		newName.Ref(),
	)

	return
}

// HasFuncResolveIsolatedEntries returns true if the function "WEBEXT.fileManagerPrivate.resolveIsolatedEntries" exists.
func HasFuncResolveIsolatedEntries() bool {
	return js.True == bindings.HasFuncResolveIsolatedEntries()
}

// FuncResolveIsolatedEntries returns the function "WEBEXT.fileManagerPrivate.resolveIsolatedEntries".
func FuncResolveIsolatedEntries() (fn js.Func[func(entries js.Array[js.Object]) js.Promise[js.Array[js.Object]]]) {
	bindings.FuncResolveIsolatedEntries(
		js.Pointer(&fn),
	)
	return
}

// ResolveIsolatedEntries calls the function "WEBEXT.fileManagerPrivate.resolveIsolatedEntries" directly.
func ResolveIsolatedEntries(entries js.Array[js.Object]) (ret js.Promise[js.Array[js.Object]]) {
	bindings.CallResolveIsolatedEntries(
		js.Pointer(&ret),
		entries.Ref(),
	)

	return
}

// TryResolveIsolatedEntries calls the function "WEBEXT.fileManagerPrivate.resolveIsolatedEntries"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryResolveIsolatedEntries(entries js.Array[js.Object]) (ret js.Promise[js.Array[js.Object]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryResolveIsolatedEntries(
		js.Pointer(&ret), js.Pointer(&exception),
		entries.Ref(),
	)

	return
}

// HasFuncResumeIOTask returns true if the function "WEBEXT.fileManagerPrivate.resumeIOTask" exists.
func HasFuncResumeIOTask() bool {
	return js.True == bindings.HasFuncResumeIOTask()
}

// FuncResumeIOTask returns the function "WEBEXT.fileManagerPrivate.resumeIOTask".
func FuncResumeIOTask() (fn js.Func[func(taskId int32, params ResumeParams)]) {
	bindings.FuncResumeIOTask(
		js.Pointer(&fn),
	)
	return
}

// ResumeIOTask calls the function "WEBEXT.fileManagerPrivate.resumeIOTask" directly.
func ResumeIOTask(taskId int32, params ResumeParams) (ret js.Void) {
	bindings.CallResumeIOTask(
		js.Pointer(&ret),
		int32(taskId),
		js.Pointer(&params),
	)

	return
}

// TryResumeIOTask calls the function "WEBEXT.fileManagerPrivate.resumeIOTask"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryResumeIOTask(taskId int32, params ResumeParams) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryResumeIOTask(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(taskId),
		js.Pointer(&params),
	)

	return
}

// HasFuncSearchDrive returns true if the function "WEBEXT.fileManagerPrivate.searchDrive" exists.
func HasFuncSearchDrive() bool {
	return js.True == bindings.HasFuncSearchDrive()
}

// FuncSearchDrive returns the function "WEBEXT.fileManagerPrivate.searchDrive".
func FuncSearchDrive() (fn js.Func[func(searchParams SearchParams, callback js.Func[func(entries js.Array[js.Object], nextFeed js.String)])]) {
	bindings.FuncSearchDrive(
		js.Pointer(&fn),
	)
	return
}

// SearchDrive calls the function "WEBEXT.fileManagerPrivate.searchDrive" directly.
func SearchDrive(searchParams SearchParams, callback js.Func[func(entries js.Array[js.Object], nextFeed js.String)]) (ret js.Void) {
	bindings.CallSearchDrive(
		js.Pointer(&ret),
		js.Pointer(&searchParams),
		callback.Ref(),
	)

	return
}

// TrySearchDrive calls the function "WEBEXT.fileManagerPrivate.searchDrive"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySearchDrive(searchParams SearchParams, callback js.Func[func(entries js.Array[js.Object], nextFeed js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySearchDrive(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&searchParams),
		callback.Ref(),
	)

	return
}

// HasFuncSearchDriveMetadata returns true if the function "WEBEXT.fileManagerPrivate.searchDriveMetadata" exists.
func HasFuncSearchDriveMetadata() bool {
	return js.True == bindings.HasFuncSearchDriveMetadata()
}

// FuncSearchDriveMetadata returns the function "WEBEXT.fileManagerPrivate.searchDriveMetadata".
func FuncSearchDriveMetadata() (fn js.Func[func(searchParams SearchMetadataParams) js.Promise[js.Array[DriveMetadataSearchResult]]]) {
	bindings.FuncSearchDriveMetadata(
		js.Pointer(&fn),
	)
	return
}

// SearchDriveMetadata calls the function "WEBEXT.fileManagerPrivate.searchDriveMetadata" directly.
func SearchDriveMetadata(searchParams SearchMetadataParams) (ret js.Promise[js.Array[DriveMetadataSearchResult]]) {
	bindings.CallSearchDriveMetadata(
		js.Pointer(&ret),
		js.Pointer(&searchParams),
	)

	return
}

// TrySearchDriveMetadata calls the function "WEBEXT.fileManagerPrivate.searchDriveMetadata"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySearchDriveMetadata(searchParams SearchMetadataParams) (ret js.Promise[js.Array[DriveMetadataSearchResult]], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySearchDriveMetadata(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&searchParams),
	)

	return
}

// HasFuncSearchFiles returns true if the function "WEBEXT.fileManagerPrivate.searchFiles" exists.
func HasFuncSearchFiles() bool {
	return js.True == bindings.HasFuncSearchFiles()
}

// FuncSearchFiles returns the function "WEBEXT.fileManagerPrivate.searchFiles".
func FuncSearchFiles() (fn js.Func[func(searchParams SearchMetadataParams) js.Promise[js.Array[js.Object]]]) {
	bindings.FuncSearchFiles(
		js.Pointer(&fn),
	)
	return
}

// SearchFiles calls the function "WEBEXT.fileManagerPrivate.searchFiles" directly.
func SearchFiles(searchParams SearchMetadataParams) (ret js.Promise[js.Array[js.Object]]) {
	bindings.CallSearchFiles(
		js.Pointer(&ret),
		js.Pointer(&searchParams),
	)

	return
}

// TrySearchFiles calls the function "WEBEXT.fileManagerPrivate.searchFiles"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySearchFiles(searchParams SearchMetadataParams) (ret js.Promise[js.Array[js.Object]], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySearchFiles(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&searchParams),
	)

	return
}

// HasFuncSearchFilesByHashes returns true if the function "WEBEXT.fileManagerPrivate.searchFilesByHashes" exists.
func HasFuncSearchFilesByHashes() bool {
	return js.True == bindings.HasFuncSearchFilesByHashes()
}

// FuncSearchFilesByHashes returns the function "WEBEXT.fileManagerPrivate.searchFilesByHashes".
func FuncSearchFilesByHashes() (fn js.Func[func(volumeId js.String, hashList js.Array[js.String]) js.Promise[js.Object]]) {
	bindings.FuncSearchFilesByHashes(
		js.Pointer(&fn),
	)
	return
}

// SearchFilesByHashes calls the function "WEBEXT.fileManagerPrivate.searchFilesByHashes" directly.
func SearchFilesByHashes(volumeId js.String, hashList js.Array[js.String]) (ret js.Promise[js.Object]) {
	bindings.CallSearchFilesByHashes(
		js.Pointer(&ret),
		volumeId.Ref(),
		hashList.Ref(),
	)

	return
}

// TrySearchFilesByHashes calls the function "WEBEXT.fileManagerPrivate.searchFilesByHashes"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySearchFilesByHashes(volumeId js.String, hashList js.Array[js.String]) (ret js.Promise[js.Object], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySearchFilesByHashes(
		js.Pointer(&ret), js.Pointer(&exception),
		volumeId.Ref(),
		hashList.Ref(),
	)

	return
}

// HasFuncSelectAndroidPickerApp returns true if the function "WEBEXT.fileManagerPrivate.selectAndroidPickerApp" exists.
func HasFuncSelectAndroidPickerApp() bool {
	return js.True == bindings.HasFuncSelectAndroidPickerApp()
}

// FuncSelectAndroidPickerApp returns the function "WEBEXT.fileManagerPrivate.selectAndroidPickerApp".
func FuncSelectAndroidPickerApp() (fn js.Func[func(androidApp AndroidApp) js.Promise[js.Void]]) {
	bindings.FuncSelectAndroidPickerApp(
		js.Pointer(&fn),
	)
	return
}

// SelectAndroidPickerApp calls the function "WEBEXT.fileManagerPrivate.selectAndroidPickerApp" directly.
func SelectAndroidPickerApp(androidApp AndroidApp) (ret js.Promise[js.Void]) {
	bindings.CallSelectAndroidPickerApp(
		js.Pointer(&ret),
		js.Pointer(&androidApp),
	)

	return
}

// TrySelectAndroidPickerApp calls the function "WEBEXT.fileManagerPrivate.selectAndroidPickerApp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySelectAndroidPickerApp(androidApp AndroidApp) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectAndroidPickerApp(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&androidApp),
	)

	return
}

// HasFuncSelectFile returns true if the function "WEBEXT.fileManagerPrivate.selectFile" exists.
func HasFuncSelectFile() bool {
	return js.True == bindings.HasFuncSelectFile()
}

// FuncSelectFile returns the function "WEBEXT.fileManagerPrivate.selectFile".
func FuncSelectFile() (fn js.Func[func(selectedPath js.String, index int32, forOpening bool, shouldReturnLocalPath bool) js.Promise[js.Void]]) {
	bindings.FuncSelectFile(
		js.Pointer(&fn),
	)
	return
}

// SelectFile calls the function "WEBEXT.fileManagerPrivate.selectFile" directly.
func SelectFile(selectedPath js.String, index int32, forOpening bool, shouldReturnLocalPath bool) (ret js.Promise[js.Void]) {
	bindings.CallSelectFile(
		js.Pointer(&ret),
		selectedPath.Ref(),
		int32(index),
		js.Bool(bool(forOpening)),
		js.Bool(bool(shouldReturnLocalPath)),
	)

	return
}

// TrySelectFile calls the function "WEBEXT.fileManagerPrivate.selectFile"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySelectFile(selectedPath js.String, index int32, forOpening bool, shouldReturnLocalPath bool) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectFile(
		js.Pointer(&ret), js.Pointer(&exception),
		selectedPath.Ref(),
		int32(index),
		js.Bool(bool(forOpening)),
		js.Bool(bool(shouldReturnLocalPath)),
	)

	return
}

// HasFuncSelectFiles returns true if the function "WEBEXT.fileManagerPrivate.selectFiles" exists.
func HasFuncSelectFiles() bool {
	return js.True == bindings.HasFuncSelectFiles()
}

// FuncSelectFiles returns the function "WEBEXT.fileManagerPrivate.selectFiles".
func FuncSelectFiles() (fn js.Func[func(selectedPaths js.Array[js.String], shouldReturnLocalPath bool) js.Promise[js.Void]]) {
	bindings.FuncSelectFiles(
		js.Pointer(&fn),
	)
	return
}

// SelectFiles calls the function "WEBEXT.fileManagerPrivate.selectFiles" directly.
func SelectFiles(selectedPaths js.Array[js.String], shouldReturnLocalPath bool) (ret js.Promise[js.Void]) {
	bindings.CallSelectFiles(
		js.Pointer(&ret),
		selectedPaths.Ref(),
		js.Bool(bool(shouldReturnLocalPath)),
	)

	return
}

// TrySelectFiles calls the function "WEBEXT.fileManagerPrivate.selectFiles"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySelectFiles(selectedPaths js.Array[js.String], shouldReturnLocalPath bool) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectFiles(
		js.Pointer(&ret), js.Pointer(&exception),
		selectedPaths.Ref(),
		js.Bool(bool(shouldReturnLocalPath)),
	)

	return
}

// HasFuncSendFeedback returns true if the function "WEBEXT.fileManagerPrivate.sendFeedback" exists.
func HasFuncSendFeedback() bool {
	return js.True == bindings.HasFuncSendFeedback()
}

// FuncSendFeedback returns the function "WEBEXT.fileManagerPrivate.sendFeedback".
func FuncSendFeedback() (fn js.Func[func()]) {
	bindings.FuncSendFeedback(
		js.Pointer(&fn),
	)
	return
}

// SendFeedback calls the function "WEBEXT.fileManagerPrivate.sendFeedback" directly.
func SendFeedback() (ret js.Void) {
	bindings.CallSendFeedback(
		js.Pointer(&ret),
	)

	return
}

// TrySendFeedback calls the function "WEBEXT.fileManagerPrivate.sendFeedback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendFeedback() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendFeedback(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetDefaultTask returns true if the function "WEBEXT.fileManagerPrivate.setDefaultTask" exists.
func HasFuncSetDefaultTask() bool {
	return js.True == bindings.HasFuncSetDefaultTask()
}

// FuncSetDefaultTask returns the function "WEBEXT.fileManagerPrivate.setDefaultTask".
func FuncSetDefaultTask() (fn js.Func[func(descriptor FileTaskDescriptor, entries js.Array[js.Object], mimeTypes js.Array[js.String]) js.Promise[js.Void]]) {
	bindings.FuncSetDefaultTask(
		js.Pointer(&fn),
	)
	return
}

// SetDefaultTask calls the function "WEBEXT.fileManagerPrivate.setDefaultTask" directly.
func SetDefaultTask(descriptor FileTaskDescriptor, entries js.Array[js.Object], mimeTypes js.Array[js.String]) (ret js.Promise[js.Void]) {
	bindings.CallSetDefaultTask(
		js.Pointer(&ret),
		js.Pointer(&descriptor),
		entries.Ref(),
		mimeTypes.Ref(),
	)

	return
}

// TrySetDefaultTask calls the function "WEBEXT.fileManagerPrivate.setDefaultTask"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetDefaultTask(descriptor FileTaskDescriptor, entries js.Array[js.Object], mimeTypes js.Array[js.String]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetDefaultTask(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
		entries.Ref(),
		mimeTypes.Ref(),
	)

	return
}

// HasFuncSetPreferences returns true if the function "WEBEXT.fileManagerPrivate.setPreferences" exists.
func HasFuncSetPreferences() bool {
	return js.True == bindings.HasFuncSetPreferences()
}

// FuncSetPreferences returns the function "WEBEXT.fileManagerPrivate.setPreferences".
func FuncSetPreferences() (fn js.Func[func(changeInfo PreferencesChange)]) {
	bindings.FuncSetPreferences(
		js.Pointer(&fn),
	)
	return
}

// SetPreferences calls the function "WEBEXT.fileManagerPrivate.setPreferences" directly.
func SetPreferences(changeInfo PreferencesChange) (ret js.Void) {
	bindings.CallSetPreferences(
		js.Pointer(&ret),
		js.Pointer(&changeInfo),
	)

	return
}

// TrySetPreferences calls the function "WEBEXT.fileManagerPrivate.setPreferences"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetPreferences(changeInfo PreferencesChange) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetPreferences(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&changeInfo),
	)

	return
}

// HasFuncSharePathsWithCrostini returns true if the function "WEBEXT.fileManagerPrivate.sharePathsWithCrostini" exists.
func HasFuncSharePathsWithCrostini() bool {
	return js.True == bindings.HasFuncSharePathsWithCrostini()
}

// FuncSharePathsWithCrostini returns the function "WEBEXT.fileManagerPrivate.sharePathsWithCrostini".
func FuncSharePathsWithCrostini() (fn js.Func[func(vmName js.String, entries js.Array[js.Object], persist bool) js.Promise[js.Void]]) {
	bindings.FuncSharePathsWithCrostini(
		js.Pointer(&fn),
	)
	return
}

// SharePathsWithCrostini calls the function "WEBEXT.fileManagerPrivate.sharePathsWithCrostini" directly.
func SharePathsWithCrostini(vmName js.String, entries js.Array[js.Object], persist bool) (ret js.Promise[js.Void]) {
	bindings.CallSharePathsWithCrostini(
		js.Pointer(&ret),
		vmName.Ref(),
		entries.Ref(),
		js.Bool(bool(persist)),
	)

	return
}

// TrySharePathsWithCrostini calls the function "WEBEXT.fileManagerPrivate.sharePathsWithCrostini"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySharePathsWithCrostini(vmName js.String, entries js.Array[js.Object], persist bool) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySharePathsWithCrostini(
		js.Pointer(&ret), js.Pointer(&exception),
		vmName.Ref(),
		entries.Ref(),
		js.Bool(bool(persist)),
	)

	return
}

// HasFuncSharesheetHasTargets returns true if the function "WEBEXT.fileManagerPrivate.sharesheetHasTargets" exists.
func HasFuncSharesheetHasTargets() bool {
	return js.True == bindings.HasFuncSharesheetHasTargets()
}

// FuncSharesheetHasTargets returns the function "WEBEXT.fileManagerPrivate.sharesheetHasTargets".
func FuncSharesheetHasTargets() (fn js.Func[func(entries js.Array[js.Object]) js.Promise[js.Boolean]]) {
	bindings.FuncSharesheetHasTargets(
		js.Pointer(&fn),
	)
	return
}

// SharesheetHasTargets calls the function "WEBEXT.fileManagerPrivate.sharesheetHasTargets" directly.
func SharesheetHasTargets(entries js.Array[js.Object]) (ret js.Promise[js.Boolean]) {
	bindings.CallSharesheetHasTargets(
		js.Pointer(&ret),
		entries.Ref(),
	)

	return
}

// TrySharesheetHasTargets calls the function "WEBEXT.fileManagerPrivate.sharesheetHasTargets"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySharesheetHasTargets(entries js.Array[js.Object]) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySharesheetHasTargets(
		js.Pointer(&ret), js.Pointer(&exception),
		entries.Ref(),
	)

	return
}

// HasFuncShowPolicyDialog returns true if the function "WEBEXT.fileManagerPrivate.showPolicyDialog" exists.
func HasFuncShowPolicyDialog() bool {
	return js.True == bindings.HasFuncShowPolicyDialog()
}

// FuncShowPolicyDialog returns the function "WEBEXT.fileManagerPrivate.showPolicyDialog".
func FuncShowPolicyDialog() (fn js.Func[func(taskId int32, typ PolicyDialogType)]) {
	bindings.FuncShowPolicyDialog(
		js.Pointer(&fn),
	)
	return
}

// ShowPolicyDialog calls the function "WEBEXT.fileManagerPrivate.showPolicyDialog" directly.
func ShowPolicyDialog(taskId int32, typ PolicyDialogType) (ret js.Void) {
	bindings.CallShowPolicyDialog(
		js.Pointer(&ret),
		int32(taskId),
		uint32(typ),
	)

	return
}

// TryShowPolicyDialog calls the function "WEBEXT.fileManagerPrivate.showPolicyDialog"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryShowPolicyDialog(taskId int32, typ PolicyDialogType) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryShowPolicyDialog(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(taskId),
		uint32(typ),
	)

	return
}

// HasFuncSinglePartitionFormat returns true if the function "WEBEXT.fileManagerPrivate.singlePartitionFormat" exists.
func HasFuncSinglePartitionFormat() bool {
	return js.True == bindings.HasFuncSinglePartitionFormat()
}

// FuncSinglePartitionFormat returns the function "WEBEXT.fileManagerPrivate.singlePartitionFormat".
func FuncSinglePartitionFormat() (fn js.Func[func(deviceStoragePath js.String, filesystem FormatFileSystemType, volumeLabel js.String)]) {
	bindings.FuncSinglePartitionFormat(
		js.Pointer(&fn),
	)
	return
}

// SinglePartitionFormat calls the function "WEBEXT.fileManagerPrivate.singlePartitionFormat" directly.
func SinglePartitionFormat(deviceStoragePath js.String, filesystem FormatFileSystemType, volumeLabel js.String) (ret js.Void) {
	bindings.CallSinglePartitionFormat(
		js.Pointer(&ret),
		deviceStoragePath.Ref(),
		uint32(filesystem),
		volumeLabel.Ref(),
	)

	return
}

// TrySinglePartitionFormat calls the function "WEBEXT.fileManagerPrivate.singlePartitionFormat"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySinglePartitionFormat(deviceStoragePath js.String, filesystem FormatFileSystemType, volumeLabel js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySinglePartitionFormat(
		js.Pointer(&ret), js.Pointer(&exception),
		deviceStoragePath.Ref(),
		uint32(filesystem),
		volumeLabel.Ref(),
	)

	return
}

// HasFuncStartIOTask returns true if the function "WEBEXT.fileManagerPrivate.startIOTask" exists.
func HasFuncStartIOTask() bool {
	return js.True == bindings.HasFuncStartIOTask()
}

// FuncStartIOTask returns the function "WEBEXT.fileManagerPrivate.startIOTask".
func FuncStartIOTask() (fn js.Func[func(typ IOTaskType, entries js.Array[js.Object], params IOTaskParams) js.Promise[js.Number[int32]]]) {
	bindings.FuncStartIOTask(
		js.Pointer(&fn),
	)
	return
}

// StartIOTask calls the function "WEBEXT.fileManagerPrivate.startIOTask" directly.
func StartIOTask(typ IOTaskType, entries js.Array[js.Object], params IOTaskParams) (ret js.Promise[js.Number[int32]]) {
	bindings.CallStartIOTask(
		js.Pointer(&ret),
		uint32(typ),
		entries.Ref(),
		js.Pointer(&params),
	)

	return
}

// TryStartIOTask calls the function "WEBEXT.fileManagerPrivate.startIOTask"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStartIOTask(typ IOTaskType, entries js.Array[js.Object], params IOTaskParams) (ret js.Promise[js.Number[int32]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStartIOTask(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(typ),
		entries.Ref(),
		js.Pointer(&params),
	)

	return
}

// HasFuncToggleAddedToHoldingSpace returns true if the function "WEBEXT.fileManagerPrivate.toggleAddedToHoldingSpace" exists.
func HasFuncToggleAddedToHoldingSpace() bool {
	return js.True == bindings.HasFuncToggleAddedToHoldingSpace()
}

// FuncToggleAddedToHoldingSpace returns the function "WEBEXT.fileManagerPrivate.toggleAddedToHoldingSpace".
func FuncToggleAddedToHoldingSpace() (fn js.Func[func(entries js.Array[js.Object], added bool) js.Promise[js.Void]]) {
	bindings.FuncToggleAddedToHoldingSpace(
		js.Pointer(&fn),
	)
	return
}

// ToggleAddedToHoldingSpace calls the function "WEBEXT.fileManagerPrivate.toggleAddedToHoldingSpace" directly.
func ToggleAddedToHoldingSpace(entries js.Array[js.Object], added bool) (ret js.Promise[js.Void]) {
	bindings.CallToggleAddedToHoldingSpace(
		js.Pointer(&ret),
		entries.Ref(),
		js.Bool(bool(added)),
	)

	return
}

// TryToggleAddedToHoldingSpace calls the function "WEBEXT.fileManagerPrivate.toggleAddedToHoldingSpace"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryToggleAddedToHoldingSpace(entries js.Array[js.Object], added bool) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryToggleAddedToHoldingSpace(
		js.Pointer(&ret), js.Pointer(&exception),
		entries.Ref(),
		js.Bool(bool(added)),
	)

	return
}

// HasFuncUnsharePathWithCrostini returns true if the function "WEBEXT.fileManagerPrivate.unsharePathWithCrostini" exists.
func HasFuncUnsharePathWithCrostini() bool {
	return js.True == bindings.HasFuncUnsharePathWithCrostini()
}

// FuncUnsharePathWithCrostini returns the function "WEBEXT.fileManagerPrivate.unsharePathWithCrostini".
func FuncUnsharePathWithCrostini() (fn js.Func[func(vmName js.String, entry js.Object) js.Promise[js.Void]]) {
	bindings.FuncUnsharePathWithCrostini(
		js.Pointer(&fn),
	)
	return
}

// UnsharePathWithCrostini calls the function "WEBEXT.fileManagerPrivate.unsharePathWithCrostini" directly.
func UnsharePathWithCrostini(vmName js.String, entry js.Object) (ret js.Promise[js.Void]) {
	bindings.CallUnsharePathWithCrostini(
		js.Pointer(&ret),
		vmName.Ref(),
		entry.Ref(),
	)

	return
}

// TryUnsharePathWithCrostini calls the function "WEBEXT.fileManagerPrivate.unsharePathWithCrostini"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUnsharePathWithCrostini(vmName js.String, entry js.Object) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUnsharePathWithCrostini(
		js.Pointer(&ret), js.Pointer(&exception),
		vmName.Ref(),
		entry.Ref(),
	)

	return
}

// HasFuncValidatePathNameLength returns true if the function "WEBEXT.fileManagerPrivate.validatePathNameLength" exists.
func HasFuncValidatePathNameLength() bool {
	return js.True == bindings.HasFuncValidatePathNameLength()
}

// FuncValidatePathNameLength returns the function "WEBEXT.fileManagerPrivate.validatePathNameLength".
func FuncValidatePathNameLength() (fn js.Func[func(parentEntry js.Object, name js.String) js.Promise[js.Boolean]]) {
	bindings.FuncValidatePathNameLength(
		js.Pointer(&fn),
	)
	return
}

// ValidatePathNameLength calls the function "WEBEXT.fileManagerPrivate.validatePathNameLength" directly.
func ValidatePathNameLength(parentEntry js.Object, name js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallValidatePathNameLength(
		js.Pointer(&ret),
		parentEntry.Ref(),
		name.Ref(),
	)

	return
}

// TryValidatePathNameLength calls the function "WEBEXT.fileManagerPrivate.validatePathNameLength"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryValidatePathNameLength(parentEntry js.Object, name js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryValidatePathNameLength(
		js.Pointer(&ret), js.Pointer(&exception),
		parentEntry.Ref(),
		name.Ref(),
	)

	return
}

// HasFuncZoom returns true if the function "WEBEXT.fileManagerPrivate.zoom" exists.
func HasFuncZoom() bool {
	return js.True == bindings.HasFuncZoom()
}

// FuncZoom returns the function "WEBEXT.fileManagerPrivate.zoom".
func FuncZoom() (fn js.Func[func(operation ZoomOperationType)]) {
	bindings.FuncZoom(
		js.Pointer(&fn),
	)
	return
}

// Zoom calls the function "WEBEXT.fileManagerPrivate.zoom" directly.
func Zoom(operation ZoomOperationType) (ret js.Void) {
	bindings.CallZoom(
		js.Pointer(&ret),
		uint32(operation),
	)

	return
}

// TryZoom calls the function "WEBEXT.fileManagerPrivate.zoom"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryZoom(operation ZoomOperationType) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryZoom(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(operation),
	)

	return
}
