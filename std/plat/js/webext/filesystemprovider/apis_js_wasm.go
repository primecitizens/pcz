// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package filesystemprovider

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/filesystemprovider/bindings"
)

type AbortRequestedOptions struct {
	// FileSystemId is "AbortRequestedOptions.fileSystemId"
	//
	// Optional
	FileSystemId js.String
	// RequestId is "AbortRequestedOptions.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// OperationRequestId is "AbortRequestedOptions.operationRequestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_OperationRequestId MUST be set to true to make this field effective.
	OperationRequestId int32

	FFI_USE_RequestId          bool // for RequestId.
	FFI_USE_OperationRequestId bool // for OperationRequestId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AbortRequestedOptions with all fields set.
func (p AbortRequestedOptions) FromRef(ref js.Ref) AbortRequestedOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AbortRequestedOptions in the application heap.
func (p AbortRequestedOptions) New() js.Ref {
	return bindings.AbortRequestedOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AbortRequestedOptions) UpdateFrom(ref js.Ref) {
	bindings.AbortRequestedOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AbortRequestedOptions) Update(ref js.Ref) {
	bindings.AbortRequestedOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AbortRequestedOptions) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemId.Ref(),
	)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
}

type Action struct {
	// Id is "Action.id"
	//
	// Optional
	Id js.String
	// Title is "Action.title"
	//
	// Optional
	Title js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Action with all fields set.
func (p Action) FromRef(ref js.Ref) Action {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Action in the application heap.
func (p Action) New() js.Ref {
	return bindings.ActionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Action) UpdateFrom(ref js.Ref) {
	bindings.ActionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Action) Update(ref js.Ref) {
	bindings.ActionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Action) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.Title.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
}

type ActionsCallbackFunc func(this js.Ref, actions js.Array[Action]) js.Ref

func (fn ActionsCallbackFunc) Register() js.Func[func(actions js.Array[Action])] {
	return js.RegisterCallback[func(actions js.Array[Action])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ActionsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[Action]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ActionsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, actions js.Array[Action]) js.Ref
	Arg T
}

func (cb *ActionsCallback[T]) Register() js.Func[func(actions js.Array[Action])] {
	return js.RegisterCallback[func(actions js.Array[Action])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ActionsCallback[T]) DispatchCallback(
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

		js.Array[Action]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type AddWatcherRequestedOptions struct {
	// FileSystemId is "AddWatcherRequestedOptions.fileSystemId"
	//
	// Optional
	FileSystemId js.String
	// RequestId is "AddWatcherRequestedOptions.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// EntryPath is "AddWatcherRequestedOptions.entryPath"
	//
	// Optional
	EntryPath js.String
	// Recursive is "AddWatcherRequestedOptions.recursive"
	//
	// Optional
	//
	// NOTE: FFI_USE_Recursive MUST be set to true to make this field effective.
	Recursive bool

	FFI_USE_RequestId bool // for RequestId.
	FFI_USE_Recursive bool // for Recursive.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AddWatcherRequestedOptions with all fields set.
func (p AddWatcherRequestedOptions) FromRef(ref js.Ref) AddWatcherRequestedOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AddWatcherRequestedOptions in the application heap.
func (p AddWatcherRequestedOptions) New() js.Ref {
	return bindings.AddWatcherRequestedOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AddWatcherRequestedOptions) UpdateFrom(ref js.Ref) {
	bindings.AddWatcherRequestedOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AddWatcherRequestedOptions) Update(ref js.Ref) {
	bindings.AddWatcherRequestedOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AddWatcherRequestedOptions) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemId.Ref(),
		p.EntryPath.Ref(),
	)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
	p.EntryPath = p.EntryPath.FromRef(js.Undefined)
}

type ChangeType uint32

const (
	_ ChangeType = iota

	ChangeType_CHANGED
	ChangeType_DELETED
)

func (ChangeType) FromRef(str js.Ref) ChangeType {
	return ChangeType(bindings.ConstOfChangeType(str))
}

func (x ChangeType) String() (string, bool) {
	switch x {
	case ChangeType_CHANGED:
		return "CHANGED", true
	case ChangeType_DELETED:
		return "DELETED", true
	default:
		return "", false
	}
}

type Change struct {
	// EntryPath is "Change.entryPath"
	//
	// Optional
	EntryPath js.String
	// ChangeType is "Change.changeType"
	//
	// Optional
	ChangeType ChangeType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Change with all fields set.
func (p Change) FromRef(ref js.Ref) Change {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Change in the application heap.
func (p Change) New() js.Ref {
	return bindings.ChangeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Change) UpdateFrom(ref js.Ref) {
	bindings.ChangeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Change) Update(ref js.Ref) {
	bindings.ChangeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Change) FreeMembers(recursive bool) {
	js.Free(
		p.EntryPath.Ref(),
	)
	p.EntryPath = p.EntryPath.FromRef(js.Undefined)
}

type CloseFileRequestedOptions struct {
	// FileSystemId is "CloseFileRequestedOptions.fileSystemId"
	//
	// Optional
	FileSystemId js.String
	// RequestId is "CloseFileRequestedOptions.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// OpenRequestId is "CloseFileRequestedOptions.openRequestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_OpenRequestId MUST be set to true to make this field effective.
	OpenRequestId int32

	FFI_USE_RequestId     bool // for RequestId.
	FFI_USE_OpenRequestId bool // for OpenRequestId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CloseFileRequestedOptions with all fields set.
func (p CloseFileRequestedOptions) FromRef(ref js.Ref) CloseFileRequestedOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CloseFileRequestedOptions in the application heap.
func (p CloseFileRequestedOptions) New() js.Ref {
	return bindings.CloseFileRequestedOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CloseFileRequestedOptions) UpdateFrom(ref js.Ref) {
	bindings.CloseFileRequestedOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CloseFileRequestedOptions) Update(ref js.Ref) {
	bindings.CloseFileRequestedOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CloseFileRequestedOptions) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemId.Ref(),
	)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
}

type CloudIdentifier struct {
	// ProviderName is "CloudIdentifier.providerName"
	//
	// Optional
	ProviderName js.String
	// Id is "CloudIdentifier.id"
	//
	// Optional
	Id js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CloudIdentifier with all fields set.
func (p CloudIdentifier) FromRef(ref js.Ref) CloudIdentifier {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CloudIdentifier in the application heap.
func (p CloudIdentifier) New() js.Ref {
	return bindings.CloudIdentifierJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CloudIdentifier) UpdateFrom(ref js.Ref) {
	bindings.CloudIdentifierJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CloudIdentifier) Update(ref js.Ref) {
	bindings.CloudIdentifierJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CloudIdentifier) FreeMembers(recursive bool) {
	js.Free(
		p.ProviderName.Ref(),
		p.Id.Ref(),
	)
	p.ProviderName = p.ProviderName.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
}

type CommonActionId uint32

const (
	_ CommonActionId = iota

	CommonActionId_SAVE_FOR_OFFLINE
	CommonActionId_OFFLINE_NOT_NECESSARY
	CommonActionId_SHARE
)

func (CommonActionId) FromRef(str js.Ref) CommonActionId {
	return CommonActionId(bindings.ConstOfCommonActionId(str))
}

func (x CommonActionId) String() (string, bool) {
	switch x {
	case CommonActionId_SAVE_FOR_OFFLINE:
		return "SAVE_FOR_OFFLINE", true
	case CommonActionId_OFFLINE_NOT_NECESSARY:
		return "OFFLINE_NOT_NECESSARY", true
	case CommonActionId_SHARE:
		return "SHARE", true
	default:
		return "", false
	}
}

type ConfigureRequestedOptions struct {
	// FileSystemId is "ConfigureRequestedOptions.fileSystemId"
	//
	// Optional
	FileSystemId js.String
	// RequestId is "ConfigureRequestedOptions.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32

	FFI_USE_RequestId bool // for RequestId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ConfigureRequestedOptions with all fields set.
func (p ConfigureRequestedOptions) FromRef(ref js.Ref) ConfigureRequestedOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ConfigureRequestedOptions in the application heap.
func (p ConfigureRequestedOptions) New() js.Ref {
	return bindings.ConfigureRequestedOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ConfigureRequestedOptions) UpdateFrom(ref js.Ref) {
	bindings.ConfigureRequestedOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ConfigureRequestedOptions) Update(ref js.Ref) {
	bindings.ConfigureRequestedOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ConfigureRequestedOptions) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemId.Ref(),
	)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
}

type CopyEntryRequestedOptions struct {
	// FileSystemId is "CopyEntryRequestedOptions.fileSystemId"
	//
	// Optional
	FileSystemId js.String
	// RequestId is "CopyEntryRequestedOptions.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// SourcePath is "CopyEntryRequestedOptions.sourcePath"
	//
	// Optional
	SourcePath js.String
	// TargetPath is "CopyEntryRequestedOptions.targetPath"
	//
	// Optional
	TargetPath js.String

	FFI_USE_RequestId bool // for RequestId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CopyEntryRequestedOptions with all fields set.
func (p CopyEntryRequestedOptions) FromRef(ref js.Ref) CopyEntryRequestedOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CopyEntryRequestedOptions in the application heap.
func (p CopyEntryRequestedOptions) New() js.Ref {
	return bindings.CopyEntryRequestedOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CopyEntryRequestedOptions) UpdateFrom(ref js.Ref) {
	bindings.CopyEntryRequestedOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CopyEntryRequestedOptions) Update(ref js.Ref) {
	bindings.CopyEntryRequestedOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CopyEntryRequestedOptions) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemId.Ref(),
		p.SourcePath.Ref(),
		p.TargetPath.Ref(),
	)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
	p.SourcePath = p.SourcePath.FromRef(js.Undefined)
	p.TargetPath = p.TargetPath.FromRef(js.Undefined)
}

type CreateDirectoryRequestedOptions struct {
	// FileSystemId is "CreateDirectoryRequestedOptions.fileSystemId"
	//
	// Optional
	FileSystemId js.String
	// RequestId is "CreateDirectoryRequestedOptions.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// DirectoryPath is "CreateDirectoryRequestedOptions.directoryPath"
	//
	// Optional
	DirectoryPath js.String
	// Recursive is "CreateDirectoryRequestedOptions.recursive"
	//
	// Optional
	//
	// NOTE: FFI_USE_Recursive MUST be set to true to make this field effective.
	Recursive bool

	FFI_USE_RequestId bool // for RequestId.
	FFI_USE_Recursive bool // for Recursive.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CreateDirectoryRequestedOptions with all fields set.
func (p CreateDirectoryRequestedOptions) FromRef(ref js.Ref) CreateDirectoryRequestedOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CreateDirectoryRequestedOptions in the application heap.
func (p CreateDirectoryRequestedOptions) New() js.Ref {
	return bindings.CreateDirectoryRequestedOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CreateDirectoryRequestedOptions) UpdateFrom(ref js.Ref) {
	bindings.CreateDirectoryRequestedOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CreateDirectoryRequestedOptions) Update(ref js.Ref) {
	bindings.CreateDirectoryRequestedOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CreateDirectoryRequestedOptions) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemId.Ref(),
		p.DirectoryPath.Ref(),
	)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
	p.DirectoryPath = p.DirectoryPath.FromRef(js.Undefined)
}

type CreateFileRequestedOptions struct {
	// FileSystemId is "CreateFileRequestedOptions.fileSystemId"
	//
	// Optional
	FileSystemId js.String
	// RequestId is "CreateFileRequestedOptions.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// FilePath is "CreateFileRequestedOptions.filePath"
	//
	// Optional
	FilePath js.String

	FFI_USE_RequestId bool // for RequestId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CreateFileRequestedOptions with all fields set.
func (p CreateFileRequestedOptions) FromRef(ref js.Ref) CreateFileRequestedOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CreateFileRequestedOptions in the application heap.
func (p CreateFileRequestedOptions) New() js.Ref {
	return bindings.CreateFileRequestedOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CreateFileRequestedOptions) UpdateFrom(ref js.Ref) {
	bindings.CreateFileRequestedOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CreateFileRequestedOptions) Update(ref js.Ref) {
	bindings.CreateFileRequestedOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CreateFileRequestedOptions) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemId.Ref(),
		p.FilePath.Ref(),
	)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
	p.FilePath = p.FilePath.FromRef(js.Undefined)
}

type DeleteEntryRequestedOptions struct {
	// FileSystemId is "DeleteEntryRequestedOptions.fileSystemId"
	//
	// Optional
	FileSystemId js.String
	// RequestId is "DeleteEntryRequestedOptions.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// EntryPath is "DeleteEntryRequestedOptions.entryPath"
	//
	// Optional
	EntryPath js.String
	// Recursive is "DeleteEntryRequestedOptions.recursive"
	//
	// Optional
	//
	// NOTE: FFI_USE_Recursive MUST be set to true to make this field effective.
	Recursive bool

	FFI_USE_RequestId bool // for RequestId.
	FFI_USE_Recursive bool // for Recursive.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DeleteEntryRequestedOptions with all fields set.
func (p DeleteEntryRequestedOptions) FromRef(ref js.Ref) DeleteEntryRequestedOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DeleteEntryRequestedOptions in the application heap.
func (p DeleteEntryRequestedOptions) New() js.Ref {
	return bindings.DeleteEntryRequestedOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DeleteEntryRequestedOptions) UpdateFrom(ref js.Ref) {
	bindings.DeleteEntryRequestedOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DeleteEntryRequestedOptions) Update(ref js.Ref) {
	bindings.DeleteEntryRequestedOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DeleteEntryRequestedOptions) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemId.Ref(),
		p.EntryPath.Ref(),
	)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
	p.EntryPath = p.EntryPath.FromRef(js.Undefined)
}

type EntriesCallbackFunc func(this js.Ref, entries js.Array[EntryMetadata], hasMore bool) js.Ref

func (fn EntriesCallbackFunc) Register() js.Func[func(entries js.Array[EntryMetadata], hasMore bool)] {
	return js.RegisterCallback[func(entries js.Array[EntryMetadata], hasMore bool)](
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

		js.Array[EntryMetadata]{}.FromRef(args[0+1]),
		args[1+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type EntriesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[EntryMetadata], hasMore bool) js.Ref
	Arg T
}

func (cb *EntriesCallback[T]) Register() js.Func[func(entries js.Array[EntryMetadata], hasMore bool)] {
	return js.RegisterCallback[func(entries js.Array[EntryMetadata], hasMore bool)](
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

		js.Array[EntryMetadata]{}.FromRef(args[0+1]),
		args[1+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type EntryMetadata struct {
	// IsDirectory is "EntryMetadata.isDirectory"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsDirectory MUST be set to true to make this field effective.
	IsDirectory bool
	// Name is "EntryMetadata.name"
	//
	// Optional
	Name js.String
	// Size is "EntryMetadata.size"
	//
	// Optional
	//
	// NOTE: FFI_USE_Size MUST be set to true to make this field effective.
	Size float64
	// ModificationTime is "EntryMetadata.modificationTime"
	//
	// Optional
	ModificationTime js.Object
	// MimeType is "EntryMetadata.mimeType"
	//
	// Optional
	MimeType js.String
	// Thumbnail is "EntryMetadata.thumbnail"
	//
	// Optional
	Thumbnail js.String
	// CloudIdentifier is "EntryMetadata.cloudIdentifier"
	//
	// Optional
	//
	// NOTE: CloudIdentifier.FFI_USE MUST be set to true to get CloudIdentifier used.
	CloudIdentifier CloudIdentifier

	FFI_USE_IsDirectory bool // for IsDirectory.
	FFI_USE_Size        bool // for Size.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EntryMetadata with all fields set.
func (p EntryMetadata) FromRef(ref js.Ref) EntryMetadata {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EntryMetadata in the application heap.
func (p EntryMetadata) New() js.Ref {
	return bindings.EntryMetadataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *EntryMetadata) UpdateFrom(ref js.Ref) {
	bindings.EntryMetadataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EntryMetadata) Update(ref js.Ref) {
	bindings.EntryMetadataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EntryMetadata) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.ModificationTime.Ref(),
		p.MimeType.Ref(),
		p.Thumbnail.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.ModificationTime = p.ModificationTime.FromRef(js.Undefined)
	p.MimeType = p.MimeType.FromRef(js.Undefined)
	p.Thumbnail = p.Thumbnail.FromRef(js.Undefined)
	if recursive {
		p.CloudIdentifier.FreeMembers(true)
	}
}

type ExecuteActionRequestedOptions struct {
	// FileSystemId is "ExecuteActionRequestedOptions.fileSystemId"
	//
	// Optional
	FileSystemId js.String
	// RequestId is "ExecuteActionRequestedOptions.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// EntryPaths is "ExecuteActionRequestedOptions.entryPaths"
	//
	// Optional
	EntryPaths js.Array[js.String]
	// ActionId is "ExecuteActionRequestedOptions.actionId"
	//
	// Optional
	ActionId js.String

	FFI_USE_RequestId bool // for RequestId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ExecuteActionRequestedOptions with all fields set.
func (p ExecuteActionRequestedOptions) FromRef(ref js.Ref) ExecuteActionRequestedOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ExecuteActionRequestedOptions in the application heap.
func (p ExecuteActionRequestedOptions) New() js.Ref {
	return bindings.ExecuteActionRequestedOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ExecuteActionRequestedOptions) UpdateFrom(ref js.Ref) {
	bindings.ExecuteActionRequestedOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ExecuteActionRequestedOptions) Update(ref js.Ref) {
	bindings.ExecuteActionRequestedOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ExecuteActionRequestedOptions) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemId.Ref(),
		p.EntryPaths.Ref(),
		p.ActionId.Ref(),
	)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
	p.EntryPaths = p.EntryPaths.FromRef(js.Undefined)
	p.ActionId = p.ActionId.FromRef(js.Undefined)
}

type FileDataCallbackFunc func(this js.Ref, data js.ArrayBuffer, hasMore bool) js.Ref

func (fn FileDataCallbackFunc) Register() js.Func[func(data js.ArrayBuffer, hasMore bool)] {
	return js.RegisterCallback[func(data js.ArrayBuffer, hasMore bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn FileDataCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.ArrayBuffer{}.FromRef(args[0+1]),
		args[1+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type FileDataCallback[T any] struct {
	Fn  func(arg T, this js.Ref, data js.ArrayBuffer, hasMore bool) js.Ref
	Arg T
}

func (cb *FileDataCallback[T]) Register() js.Func[func(data js.ArrayBuffer, hasMore bool)] {
	return js.RegisterCallback[func(data js.ArrayBuffer, hasMore bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *FileDataCallback[T]) DispatchCallback(
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

		js.ArrayBuffer{}.FromRef(args[0+1]),
		args[1+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OpenFileMode uint32

const (
	_ OpenFileMode = iota

	OpenFileMode_READ
	OpenFileMode_WRITE
)

func (OpenFileMode) FromRef(str js.Ref) OpenFileMode {
	return OpenFileMode(bindings.ConstOfOpenFileMode(str))
}

func (x OpenFileMode) String() (string, bool) {
	switch x {
	case OpenFileMode_READ:
		return "READ", true
	case OpenFileMode_WRITE:
		return "WRITE", true
	default:
		return "", false
	}
}

type OpenedFile struct {
	// OpenRequestId is "OpenedFile.openRequestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_OpenRequestId MUST be set to true to make this field effective.
	OpenRequestId int32
	// FilePath is "OpenedFile.filePath"
	//
	// Optional
	FilePath js.String
	// Mode is "OpenedFile.mode"
	//
	// Optional
	Mode OpenFileMode

	FFI_USE_OpenRequestId bool // for OpenRequestId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OpenedFile with all fields set.
func (p OpenedFile) FromRef(ref js.Ref) OpenedFile {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OpenedFile in the application heap.
func (p OpenedFile) New() js.Ref {
	return bindings.OpenedFileJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OpenedFile) UpdateFrom(ref js.Ref) {
	bindings.OpenedFileJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OpenedFile) Update(ref js.Ref) {
	bindings.OpenedFileJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OpenedFile) FreeMembers(recursive bool) {
	js.Free(
		p.FilePath.Ref(),
	)
	p.FilePath = p.FilePath.FromRef(js.Undefined)
}

type Watcher struct {
	// EntryPath is "Watcher.entryPath"
	//
	// Optional
	EntryPath js.String
	// Recursive is "Watcher.recursive"
	//
	// Optional
	//
	// NOTE: FFI_USE_Recursive MUST be set to true to make this field effective.
	Recursive bool
	// LastTag is "Watcher.lastTag"
	//
	// Optional
	LastTag js.String

	FFI_USE_Recursive bool // for Recursive.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Watcher with all fields set.
func (p Watcher) FromRef(ref js.Ref) Watcher {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Watcher in the application heap.
func (p Watcher) New() js.Ref {
	return bindings.WatcherJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Watcher) UpdateFrom(ref js.Ref) {
	bindings.WatcherJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Watcher) Update(ref js.Ref) {
	bindings.WatcherJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Watcher) FreeMembers(recursive bool) {
	js.Free(
		p.EntryPath.Ref(),
		p.LastTag.Ref(),
	)
	p.EntryPath = p.EntryPath.FromRef(js.Undefined)
	p.LastTag = p.LastTag.FromRef(js.Undefined)
}

type FileSystemInfo struct {
	// FileSystemId is "FileSystemInfo.fileSystemId"
	//
	// Optional
	FileSystemId js.String
	// DisplayName is "FileSystemInfo.displayName"
	//
	// Optional
	DisplayName js.String
	// Writable is "FileSystemInfo.writable"
	//
	// Optional
	//
	// NOTE: FFI_USE_Writable MUST be set to true to make this field effective.
	Writable bool
	// OpenedFilesLimit is "FileSystemInfo.openedFilesLimit"
	//
	// Optional
	//
	// NOTE: FFI_USE_OpenedFilesLimit MUST be set to true to make this field effective.
	OpenedFilesLimit int32
	// OpenedFiles is "FileSystemInfo.openedFiles"
	//
	// Optional
	OpenedFiles js.Array[OpenedFile]
	// SupportsNotifyTag is "FileSystemInfo.supportsNotifyTag"
	//
	// Optional
	//
	// NOTE: FFI_USE_SupportsNotifyTag MUST be set to true to make this field effective.
	SupportsNotifyTag bool
	// Watchers is "FileSystemInfo.watchers"
	//
	// Optional
	Watchers js.Array[Watcher]

	FFI_USE_Writable          bool // for Writable.
	FFI_USE_OpenedFilesLimit  bool // for OpenedFilesLimit.
	FFI_USE_SupportsNotifyTag bool // for SupportsNotifyTag.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FileSystemInfo with all fields set.
func (p FileSystemInfo) FromRef(ref js.Ref) FileSystemInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FileSystemInfo in the application heap.
func (p FileSystemInfo) New() js.Ref {
	return bindings.FileSystemInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FileSystemInfo) UpdateFrom(ref js.Ref) {
	bindings.FileSystemInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FileSystemInfo) Update(ref js.Ref) {
	bindings.FileSystemInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FileSystemInfo) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemId.Ref(),
		p.DisplayName.Ref(),
		p.OpenedFiles.Ref(),
		p.Watchers.Ref(),
	)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
	p.DisplayName = p.DisplayName.FromRef(js.Undefined)
	p.OpenedFiles = p.OpenedFiles.FromRef(js.Undefined)
	p.Watchers = p.Watchers.FromRef(js.Undefined)
}

type GetActionsRequestedOptions struct {
	// FileSystemId is "GetActionsRequestedOptions.fileSystemId"
	//
	// Optional
	FileSystemId js.String
	// RequestId is "GetActionsRequestedOptions.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// EntryPaths is "GetActionsRequestedOptions.entryPaths"
	//
	// Optional
	EntryPaths js.Array[js.String]

	FFI_USE_RequestId bool // for RequestId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetActionsRequestedOptions with all fields set.
func (p GetActionsRequestedOptions) FromRef(ref js.Ref) GetActionsRequestedOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetActionsRequestedOptions in the application heap.
func (p GetActionsRequestedOptions) New() js.Ref {
	return bindings.GetActionsRequestedOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetActionsRequestedOptions) UpdateFrom(ref js.Ref) {
	bindings.GetActionsRequestedOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetActionsRequestedOptions) Update(ref js.Ref) {
	bindings.GetActionsRequestedOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetActionsRequestedOptions) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemId.Ref(),
		p.EntryPaths.Ref(),
	)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
	p.EntryPaths = p.EntryPaths.FromRef(js.Undefined)
}

type GetAllCallbackFunc func(this js.Ref, fileSystems js.Array[FileSystemInfo]) js.Ref

func (fn GetAllCallbackFunc) Register() js.Func[func(fileSystems js.Array[FileSystemInfo])] {
	return js.RegisterCallback[func(fileSystems js.Array[FileSystemInfo])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetAllCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[FileSystemInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetAllCallback[T any] struct {
	Fn  func(arg T, this js.Ref, fileSystems js.Array[FileSystemInfo]) js.Ref
	Arg T
}

func (cb *GetAllCallback[T]) Register() js.Func[func(fileSystems js.Array[FileSystemInfo])] {
	return js.RegisterCallback[func(fileSystems js.Array[FileSystemInfo])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetAllCallback[T]) DispatchCallback(
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

		js.Array[FileSystemInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetCallbackFunc func(this js.Ref, fileSystem *FileSystemInfo) js.Ref

func (fn GetCallbackFunc) Register() js.Func[func(fileSystem *FileSystemInfo)] {
	return js.RegisterCallback[func(fileSystem *FileSystemInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 FileSystemInfo
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

type GetCallback[T any] struct {
	Fn  func(arg T, this js.Ref, fileSystem *FileSystemInfo) js.Ref
	Arg T
}

func (cb *GetCallback[T]) Register() js.Func[func(fileSystem *FileSystemInfo)] {
	return js.RegisterCallback[func(fileSystem *FileSystemInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 FileSystemInfo
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

type GetMetadataRequestedOptions struct {
	// FileSystemId is "GetMetadataRequestedOptions.fileSystemId"
	//
	// Optional
	FileSystemId js.String
	// RequestId is "GetMetadataRequestedOptions.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// EntryPath is "GetMetadataRequestedOptions.entryPath"
	//
	// Optional
	EntryPath js.String
	// IsDirectory is "GetMetadataRequestedOptions.isDirectory"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsDirectory MUST be set to true to make this field effective.
	IsDirectory bool
	// Name is "GetMetadataRequestedOptions.name"
	//
	// Optional
	//
	// NOTE: FFI_USE_Name MUST be set to true to make this field effective.
	Name bool
	// Size is "GetMetadataRequestedOptions.size"
	//
	// Optional
	//
	// NOTE: FFI_USE_Size MUST be set to true to make this field effective.
	Size bool
	// ModificationTime is "GetMetadataRequestedOptions.modificationTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_ModificationTime MUST be set to true to make this field effective.
	ModificationTime bool
	// MimeType is "GetMetadataRequestedOptions.mimeType"
	//
	// Optional
	//
	// NOTE: FFI_USE_MimeType MUST be set to true to make this field effective.
	MimeType bool
	// Thumbnail is "GetMetadataRequestedOptions.thumbnail"
	//
	// Optional
	//
	// NOTE: FFI_USE_Thumbnail MUST be set to true to make this field effective.
	Thumbnail bool
	// CloudIdentifier is "GetMetadataRequestedOptions.cloudIdentifier"
	//
	// Optional
	//
	// NOTE: FFI_USE_CloudIdentifier MUST be set to true to make this field effective.
	CloudIdentifier bool

	FFI_USE_RequestId        bool // for RequestId.
	FFI_USE_IsDirectory      bool // for IsDirectory.
	FFI_USE_Name             bool // for Name.
	FFI_USE_Size             bool // for Size.
	FFI_USE_ModificationTime bool // for ModificationTime.
	FFI_USE_MimeType         bool // for MimeType.
	FFI_USE_Thumbnail        bool // for Thumbnail.
	FFI_USE_CloudIdentifier  bool // for CloudIdentifier.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetMetadataRequestedOptions with all fields set.
func (p GetMetadataRequestedOptions) FromRef(ref js.Ref) GetMetadataRequestedOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetMetadataRequestedOptions in the application heap.
func (p GetMetadataRequestedOptions) New() js.Ref {
	return bindings.GetMetadataRequestedOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetMetadataRequestedOptions) UpdateFrom(ref js.Ref) {
	bindings.GetMetadataRequestedOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetMetadataRequestedOptions) Update(ref js.Ref) {
	bindings.GetMetadataRequestedOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetMetadataRequestedOptions) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemId.Ref(),
		p.EntryPath.Ref(),
	)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
	p.EntryPath = p.EntryPath.FromRef(js.Undefined)
}

type MetadataCallbackFunc func(this js.Ref, metadata *EntryMetadata) js.Ref

func (fn MetadataCallbackFunc) Register() js.Func[func(metadata *EntryMetadata)] {
	return js.RegisterCallback[func(metadata *EntryMetadata)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn MetadataCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 EntryMetadata
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

type MetadataCallback[T any] struct {
	Fn  func(arg T, this js.Ref, metadata *EntryMetadata) js.Ref
	Arg T
}

func (cb *MetadataCallback[T]) Register() js.Func[func(metadata *EntryMetadata)] {
	return js.RegisterCallback[func(metadata *EntryMetadata)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *MetadataCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 EntryMetadata
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

type MountOptions struct {
	// FileSystemId is "MountOptions.fileSystemId"
	//
	// Optional
	FileSystemId js.String
	// DisplayName is "MountOptions.displayName"
	//
	// Optional
	DisplayName js.String
	// Writable is "MountOptions.writable"
	//
	// Optional
	//
	// NOTE: FFI_USE_Writable MUST be set to true to make this field effective.
	Writable bool
	// OpenedFilesLimit is "MountOptions.openedFilesLimit"
	//
	// Optional
	//
	// NOTE: FFI_USE_OpenedFilesLimit MUST be set to true to make this field effective.
	OpenedFilesLimit int32
	// SupportsNotifyTag is "MountOptions.supportsNotifyTag"
	//
	// Optional
	//
	// NOTE: FFI_USE_SupportsNotifyTag MUST be set to true to make this field effective.
	SupportsNotifyTag bool
	// Persistent is "MountOptions.persistent"
	//
	// Optional
	//
	// NOTE: FFI_USE_Persistent MUST be set to true to make this field effective.
	Persistent bool

	FFI_USE_Writable          bool // for Writable.
	FFI_USE_OpenedFilesLimit  bool // for OpenedFilesLimit.
	FFI_USE_SupportsNotifyTag bool // for SupportsNotifyTag.
	FFI_USE_Persistent        bool // for Persistent.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MountOptions with all fields set.
func (p MountOptions) FromRef(ref js.Ref) MountOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MountOptions in the application heap.
func (p MountOptions) New() js.Ref {
	return bindings.MountOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MountOptions) UpdateFrom(ref js.Ref) {
	bindings.MountOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MountOptions) Update(ref js.Ref) {
	bindings.MountOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MountOptions) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemId.Ref(),
		p.DisplayName.Ref(),
	)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
	p.DisplayName = p.DisplayName.FromRef(js.Undefined)
}

type MoveEntryRequestedOptions struct {
	// FileSystemId is "MoveEntryRequestedOptions.fileSystemId"
	//
	// Optional
	FileSystemId js.String
	// RequestId is "MoveEntryRequestedOptions.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// SourcePath is "MoveEntryRequestedOptions.sourcePath"
	//
	// Optional
	SourcePath js.String
	// TargetPath is "MoveEntryRequestedOptions.targetPath"
	//
	// Optional
	TargetPath js.String

	FFI_USE_RequestId bool // for RequestId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MoveEntryRequestedOptions with all fields set.
func (p MoveEntryRequestedOptions) FromRef(ref js.Ref) MoveEntryRequestedOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MoveEntryRequestedOptions in the application heap.
func (p MoveEntryRequestedOptions) New() js.Ref {
	return bindings.MoveEntryRequestedOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MoveEntryRequestedOptions) UpdateFrom(ref js.Ref) {
	bindings.MoveEntryRequestedOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MoveEntryRequestedOptions) Update(ref js.Ref) {
	bindings.MoveEntryRequestedOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MoveEntryRequestedOptions) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemId.Ref(),
		p.SourcePath.Ref(),
		p.TargetPath.Ref(),
	)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
	p.SourcePath = p.SourcePath.FromRef(js.Undefined)
	p.TargetPath = p.TargetPath.FromRef(js.Undefined)
}

type NotifyOptions struct {
	// FileSystemId is "NotifyOptions.fileSystemId"
	//
	// Optional
	FileSystemId js.String
	// ObservedPath is "NotifyOptions.observedPath"
	//
	// Optional
	ObservedPath js.String
	// Recursive is "NotifyOptions.recursive"
	//
	// Optional
	//
	// NOTE: FFI_USE_Recursive MUST be set to true to make this field effective.
	Recursive bool
	// ChangeType is "NotifyOptions.changeType"
	//
	// Optional
	ChangeType ChangeType
	// Changes is "NotifyOptions.changes"
	//
	// Optional
	Changes js.Array[Change]
	// Tag is "NotifyOptions.tag"
	//
	// Optional
	Tag js.String

	FFI_USE_Recursive bool // for Recursive.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NotifyOptions with all fields set.
func (p NotifyOptions) FromRef(ref js.Ref) NotifyOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NotifyOptions in the application heap.
func (p NotifyOptions) New() js.Ref {
	return bindings.NotifyOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *NotifyOptions) UpdateFrom(ref js.Ref) {
	bindings.NotifyOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NotifyOptions) Update(ref js.Ref) {
	bindings.NotifyOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NotifyOptions) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemId.Ref(),
		p.ObservedPath.Ref(),
		p.Changes.Ref(),
		p.Tag.Ref(),
	)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
	p.ObservedPath = p.ObservedPath.FromRef(js.Undefined)
	p.Changes = p.Changes.FromRef(js.Undefined)
	p.Tag = p.Tag.FromRef(js.Undefined)
}

type OpenFileRequestedOptions struct {
	// FileSystemId is "OpenFileRequestedOptions.fileSystemId"
	//
	// Optional
	FileSystemId js.String
	// RequestId is "OpenFileRequestedOptions.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// FilePath is "OpenFileRequestedOptions.filePath"
	//
	// Optional
	FilePath js.String
	// Mode is "OpenFileRequestedOptions.mode"
	//
	// Optional
	Mode OpenFileMode

	FFI_USE_RequestId bool // for RequestId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OpenFileRequestedOptions with all fields set.
func (p OpenFileRequestedOptions) FromRef(ref js.Ref) OpenFileRequestedOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OpenFileRequestedOptions in the application heap.
func (p OpenFileRequestedOptions) New() js.Ref {
	return bindings.OpenFileRequestedOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OpenFileRequestedOptions) UpdateFrom(ref js.Ref) {
	bindings.OpenFileRequestedOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OpenFileRequestedOptions) Update(ref js.Ref) {
	bindings.OpenFileRequestedOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OpenFileRequestedOptions) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemId.Ref(),
		p.FilePath.Ref(),
	)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
	p.FilePath = p.FilePath.FromRef(js.Undefined)
}

type ProviderError uint32

const (
	_ ProviderError = iota

	ProviderError_OK
	ProviderError_FAILED
	ProviderError_IN_USE
	ProviderError_EXISTS
	ProviderError_NOT_FOUND
	ProviderError_ACCESS_DENIED
	ProviderError_TOO_MANY_OPENED
	ProviderError_NO_MEMORY
	ProviderError_NO_SPACE
	ProviderError_NOT_A_DIRECTORY
	ProviderError_INVALID_OPERATION
	ProviderError_SECURITY
	ProviderError_ABORT
	ProviderError_NOT_A_FILE
	ProviderError_NOT_EMPTY
	ProviderError_INVALID_URL
	ProviderError_IO
)

func (ProviderError) FromRef(str js.Ref) ProviderError {
	return ProviderError(bindings.ConstOfProviderError(str))
}

func (x ProviderError) String() (string, bool) {
	switch x {
	case ProviderError_OK:
		return "OK", true
	case ProviderError_FAILED:
		return "FAILED", true
	case ProviderError_IN_USE:
		return "IN_USE", true
	case ProviderError_EXISTS:
		return "EXISTS", true
	case ProviderError_NOT_FOUND:
		return "NOT_FOUND", true
	case ProviderError_ACCESS_DENIED:
		return "ACCESS_DENIED", true
	case ProviderError_TOO_MANY_OPENED:
		return "TOO_MANY_OPENED", true
	case ProviderError_NO_MEMORY:
		return "NO_MEMORY", true
	case ProviderError_NO_SPACE:
		return "NO_SPACE", true
	case ProviderError_NOT_A_DIRECTORY:
		return "NOT_A_DIRECTORY", true
	case ProviderError_INVALID_OPERATION:
		return "INVALID_OPERATION", true
	case ProviderError_SECURITY:
		return "SECURITY", true
	case ProviderError_ABORT:
		return "ABORT", true
	case ProviderError_NOT_A_FILE:
		return "NOT_A_FILE", true
	case ProviderError_NOT_EMPTY:
		return "NOT_EMPTY", true
	case ProviderError_INVALID_URL:
		return "INVALID_URL", true
	case ProviderError_IO:
		return "IO", true
	default:
		return "", false
	}
}

type ProviderErrorCallbackFunc func(this js.Ref, err ProviderError) js.Ref

func (fn ProviderErrorCallbackFunc) Register() js.Func[func(err ProviderError)] {
	return js.RegisterCallback[func(err ProviderError)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ProviderErrorCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		ProviderError(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ProviderErrorCallback[T any] struct {
	Fn  func(arg T, this js.Ref, err ProviderError) js.Ref
	Arg T
}

func (cb *ProviderErrorCallback[T]) Register() js.Func[func(err ProviderError)] {
	return js.RegisterCallback[func(err ProviderError)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ProviderErrorCallback[T]) DispatchCallback(
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

		ProviderError(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ProviderSuccessCallbackFunc func(this js.Ref) js.Ref

func (fn ProviderSuccessCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ProviderSuccessCallbackFunc) DispatchCallback(
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

type ProviderSuccessCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *ProviderSuccessCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ProviderSuccessCallback[T]) DispatchCallback(
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

type ReadDirectoryRequestedOptions struct {
	// FileSystemId is "ReadDirectoryRequestedOptions.fileSystemId"
	//
	// Optional
	FileSystemId js.String
	// RequestId is "ReadDirectoryRequestedOptions.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// DirectoryPath is "ReadDirectoryRequestedOptions.directoryPath"
	//
	// Optional
	DirectoryPath js.String
	// IsDirectory is "ReadDirectoryRequestedOptions.isDirectory"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsDirectory MUST be set to true to make this field effective.
	IsDirectory bool
	// Name is "ReadDirectoryRequestedOptions.name"
	//
	// Optional
	//
	// NOTE: FFI_USE_Name MUST be set to true to make this field effective.
	Name bool
	// Size is "ReadDirectoryRequestedOptions.size"
	//
	// Optional
	//
	// NOTE: FFI_USE_Size MUST be set to true to make this field effective.
	Size bool
	// ModificationTime is "ReadDirectoryRequestedOptions.modificationTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_ModificationTime MUST be set to true to make this field effective.
	ModificationTime bool
	// MimeType is "ReadDirectoryRequestedOptions.mimeType"
	//
	// Optional
	//
	// NOTE: FFI_USE_MimeType MUST be set to true to make this field effective.
	MimeType bool
	// Thumbnail is "ReadDirectoryRequestedOptions.thumbnail"
	//
	// Optional
	//
	// NOTE: FFI_USE_Thumbnail MUST be set to true to make this field effective.
	Thumbnail bool

	FFI_USE_RequestId        bool // for RequestId.
	FFI_USE_IsDirectory      bool // for IsDirectory.
	FFI_USE_Name             bool // for Name.
	FFI_USE_Size             bool // for Size.
	FFI_USE_ModificationTime bool // for ModificationTime.
	FFI_USE_MimeType         bool // for MimeType.
	FFI_USE_Thumbnail        bool // for Thumbnail.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReadDirectoryRequestedOptions with all fields set.
func (p ReadDirectoryRequestedOptions) FromRef(ref js.Ref) ReadDirectoryRequestedOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReadDirectoryRequestedOptions in the application heap.
func (p ReadDirectoryRequestedOptions) New() js.Ref {
	return bindings.ReadDirectoryRequestedOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ReadDirectoryRequestedOptions) UpdateFrom(ref js.Ref) {
	bindings.ReadDirectoryRequestedOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ReadDirectoryRequestedOptions) Update(ref js.Ref) {
	bindings.ReadDirectoryRequestedOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ReadDirectoryRequestedOptions) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemId.Ref(),
		p.DirectoryPath.Ref(),
	)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
	p.DirectoryPath = p.DirectoryPath.FromRef(js.Undefined)
}

type ReadFileRequestedOptions struct {
	// FileSystemId is "ReadFileRequestedOptions.fileSystemId"
	//
	// Optional
	FileSystemId js.String
	// RequestId is "ReadFileRequestedOptions.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// OpenRequestId is "ReadFileRequestedOptions.openRequestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_OpenRequestId MUST be set to true to make this field effective.
	OpenRequestId int32
	// Offset is "ReadFileRequestedOptions.offset"
	//
	// Optional
	//
	// NOTE: FFI_USE_Offset MUST be set to true to make this field effective.
	Offset float64
	// Length is "ReadFileRequestedOptions.length"
	//
	// Optional
	//
	// NOTE: FFI_USE_Length MUST be set to true to make this field effective.
	Length float64

	FFI_USE_RequestId     bool // for RequestId.
	FFI_USE_OpenRequestId bool // for OpenRequestId.
	FFI_USE_Offset        bool // for Offset.
	FFI_USE_Length        bool // for Length.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReadFileRequestedOptions with all fields set.
func (p ReadFileRequestedOptions) FromRef(ref js.Ref) ReadFileRequestedOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReadFileRequestedOptions in the application heap.
func (p ReadFileRequestedOptions) New() js.Ref {
	return bindings.ReadFileRequestedOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ReadFileRequestedOptions) UpdateFrom(ref js.Ref) {
	bindings.ReadFileRequestedOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ReadFileRequestedOptions) Update(ref js.Ref) {
	bindings.ReadFileRequestedOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ReadFileRequestedOptions) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemId.Ref(),
	)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
}

type RemoveWatcherRequestedOptions struct {
	// FileSystemId is "RemoveWatcherRequestedOptions.fileSystemId"
	//
	// Optional
	FileSystemId js.String
	// RequestId is "RemoveWatcherRequestedOptions.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// EntryPath is "RemoveWatcherRequestedOptions.entryPath"
	//
	// Optional
	EntryPath js.String
	// Recursive is "RemoveWatcherRequestedOptions.recursive"
	//
	// Optional
	//
	// NOTE: FFI_USE_Recursive MUST be set to true to make this field effective.
	Recursive bool

	FFI_USE_RequestId bool // for RequestId.
	FFI_USE_Recursive bool // for Recursive.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RemoveWatcherRequestedOptions with all fields set.
func (p RemoveWatcherRequestedOptions) FromRef(ref js.Ref) RemoveWatcherRequestedOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RemoveWatcherRequestedOptions in the application heap.
func (p RemoveWatcherRequestedOptions) New() js.Ref {
	return bindings.RemoveWatcherRequestedOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RemoveWatcherRequestedOptions) UpdateFrom(ref js.Ref) {
	bindings.RemoveWatcherRequestedOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RemoveWatcherRequestedOptions) Update(ref js.Ref) {
	bindings.RemoveWatcherRequestedOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RemoveWatcherRequestedOptions) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemId.Ref(),
		p.EntryPath.Ref(),
	)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
	p.EntryPath = p.EntryPath.FromRef(js.Undefined)
}

type ResultCallbackFunc func(this js.Ref) js.Ref

func (fn ResultCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ResultCallbackFunc) DispatchCallback(
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

type ResultCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *ResultCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ResultCallback[T]) DispatchCallback(
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

type TruncateRequestedOptions struct {
	// FileSystemId is "TruncateRequestedOptions.fileSystemId"
	//
	// Optional
	FileSystemId js.String
	// RequestId is "TruncateRequestedOptions.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// FilePath is "TruncateRequestedOptions.filePath"
	//
	// Optional
	FilePath js.String
	// Length is "TruncateRequestedOptions.length"
	//
	// Optional
	//
	// NOTE: FFI_USE_Length MUST be set to true to make this field effective.
	Length float64

	FFI_USE_RequestId bool // for RequestId.
	FFI_USE_Length    bool // for Length.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TruncateRequestedOptions with all fields set.
func (p TruncateRequestedOptions) FromRef(ref js.Ref) TruncateRequestedOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TruncateRequestedOptions in the application heap.
func (p TruncateRequestedOptions) New() js.Ref {
	return bindings.TruncateRequestedOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TruncateRequestedOptions) UpdateFrom(ref js.Ref) {
	bindings.TruncateRequestedOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TruncateRequestedOptions) Update(ref js.Ref) {
	bindings.TruncateRequestedOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TruncateRequestedOptions) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemId.Ref(),
		p.FilePath.Ref(),
	)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
	p.FilePath = p.FilePath.FromRef(js.Undefined)
}

type UnmountOptions struct {
	// FileSystemId is "UnmountOptions.fileSystemId"
	//
	// Optional
	FileSystemId js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UnmountOptions with all fields set.
func (p UnmountOptions) FromRef(ref js.Ref) UnmountOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UnmountOptions in the application heap.
func (p UnmountOptions) New() js.Ref {
	return bindings.UnmountOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UnmountOptions) UpdateFrom(ref js.Ref) {
	bindings.UnmountOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UnmountOptions) Update(ref js.Ref) {
	bindings.UnmountOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UnmountOptions) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemId.Ref(),
	)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
}

type UnmountRequestedOptions struct {
	// FileSystemId is "UnmountRequestedOptions.fileSystemId"
	//
	// Optional
	FileSystemId js.String
	// RequestId is "UnmountRequestedOptions.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32

	FFI_USE_RequestId bool // for RequestId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UnmountRequestedOptions with all fields set.
func (p UnmountRequestedOptions) FromRef(ref js.Ref) UnmountRequestedOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UnmountRequestedOptions in the application heap.
func (p UnmountRequestedOptions) New() js.Ref {
	return bindings.UnmountRequestedOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UnmountRequestedOptions) UpdateFrom(ref js.Ref) {
	bindings.UnmountRequestedOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UnmountRequestedOptions) Update(ref js.Ref) {
	bindings.UnmountRequestedOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UnmountRequestedOptions) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemId.Ref(),
	)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
}

type WriteFileRequestedOptions struct {
	// FileSystemId is "WriteFileRequestedOptions.fileSystemId"
	//
	// Optional
	FileSystemId js.String
	// RequestId is "WriteFileRequestedOptions.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// OpenRequestId is "WriteFileRequestedOptions.openRequestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_OpenRequestId MUST be set to true to make this field effective.
	OpenRequestId int32
	// Offset is "WriteFileRequestedOptions.offset"
	//
	// Optional
	//
	// NOTE: FFI_USE_Offset MUST be set to true to make this field effective.
	Offset float64
	// Data is "WriteFileRequestedOptions.data"
	//
	// Optional
	Data js.ArrayBuffer

	FFI_USE_RequestId     bool // for RequestId.
	FFI_USE_OpenRequestId bool // for OpenRequestId.
	FFI_USE_Offset        bool // for Offset.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WriteFileRequestedOptions with all fields set.
func (p WriteFileRequestedOptions) FromRef(ref js.Ref) WriteFileRequestedOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WriteFileRequestedOptions in the application heap.
func (p WriteFileRequestedOptions) New() js.Ref {
	return bindings.WriteFileRequestedOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *WriteFileRequestedOptions) UpdateFrom(ref js.Ref) {
	bindings.WriteFileRequestedOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WriteFileRequestedOptions) Update(ref js.Ref) {
	bindings.WriteFileRequestedOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WriteFileRequestedOptions) FreeMembers(recursive bool) {
	js.Free(
		p.FileSystemId.Ref(),
		p.Data.Ref(),
	)
	p.FileSystemId = p.FileSystemId.FromRef(js.Undefined)
	p.Data = p.Data.FromRef(js.Undefined)
}

// HasFuncGet returns true if the function "WEBEXT.fileSystemProvider.get" exists.
func HasFuncGet() bool {
	return js.True == bindings.HasFuncGet()
}

// FuncGet returns the function "WEBEXT.fileSystemProvider.get".
func FuncGet() (fn js.Func[func(fileSystemId js.String) js.Promise[FileSystemInfo]]) {
	bindings.FuncGet(
		js.Pointer(&fn),
	)
	return
}

// Get calls the function "WEBEXT.fileSystemProvider.get" directly.
func Get(fileSystemId js.String) (ret js.Promise[FileSystemInfo]) {
	bindings.CallGet(
		js.Pointer(&ret),
		fileSystemId.Ref(),
	)

	return
}

// TryGet calls the function "WEBEXT.fileSystemProvider.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGet(fileSystemId js.String) (ret js.Promise[FileSystemInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGet(
		js.Pointer(&ret), js.Pointer(&exception),
		fileSystemId.Ref(),
	)

	return
}

// HasFuncGetAll returns true if the function "WEBEXT.fileSystemProvider.getAll" exists.
func HasFuncGetAll() bool {
	return js.True == bindings.HasFuncGetAll()
}

// FuncGetAll returns the function "WEBEXT.fileSystemProvider.getAll".
func FuncGetAll() (fn js.Func[func() js.Promise[js.Array[FileSystemInfo]]]) {
	bindings.FuncGetAll(
		js.Pointer(&fn),
	)
	return
}

// GetAll calls the function "WEBEXT.fileSystemProvider.getAll" directly.
func GetAll() (ret js.Promise[js.Array[FileSystemInfo]]) {
	bindings.CallGetAll(
		js.Pointer(&ret),
	)

	return
}

// TryGetAll calls the function "WEBEXT.fileSystemProvider.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAll() (ret js.Promise[js.Array[FileSystemInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAll(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMount returns true if the function "WEBEXT.fileSystemProvider.mount" exists.
func HasFuncMount() bool {
	return js.True == bindings.HasFuncMount()
}

// FuncMount returns the function "WEBEXT.fileSystemProvider.mount".
func FuncMount() (fn js.Func[func(options MountOptions) js.Promise[js.Void]]) {
	bindings.FuncMount(
		js.Pointer(&fn),
	)
	return
}

// Mount calls the function "WEBEXT.fileSystemProvider.mount" directly.
func Mount(options MountOptions) (ret js.Promise[js.Void]) {
	bindings.CallMount(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryMount calls the function "WEBEXT.fileSystemProvider.mount"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryMount(options MountOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMount(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncNotify returns true if the function "WEBEXT.fileSystemProvider.notify" exists.
func HasFuncNotify() bool {
	return js.True == bindings.HasFuncNotify()
}

// FuncNotify returns the function "WEBEXT.fileSystemProvider.notify".
func FuncNotify() (fn js.Func[func(options NotifyOptions) js.Promise[js.Void]]) {
	bindings.FuncNotify(
		js.Pointer(&fn),
	)
	return
}

// Notify calls the function "WEBEXT.fileSystemProvider.notify" directly.
func Notify(options NotifyOptions) (ret js.Promise[js.Void]) {
	bindings.CallNotify(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryNotify calls the function "WEBEXT.fileSystemProvider.notify"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryNotify(options NotifyOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNotify(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

type OnAbortRequestedEventCallbackFunc func(this js.Ref, options *AbortRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref

func (fn OnAbortRequestedEventCallbackFunc) Register() js.Func[func(options *AbortRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *AbortRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnAbortRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AbortRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnAbortRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, options *AbortRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref
	Arg T
}

func (cb *OnAbortRequestedEventCallback[T]) Register() js.Func[func(options *AbortRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *AbortRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnAbortRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AbortRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnAbortRequested returns true if the function "WEBEXT.fileSystemProvider.onAbortRequested.addListener" exists.
func HasFuncOnAbortRequested() bool {
	return js.True == bindings.HasFuncOnAbortRequested()
}

// FuncOnAbortRequested returns the function "WEBEXT.fileSystemProvider.onAbortRequested.addListener".
func FuncOnAbortRequested() (fn js.Func[func(callback js.Func[func(options *AbortRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOnAbortRequested(
		js.Pointer(&fn),
	)
	return
}

// OnAbortRequested calls the function "WEBEXT.fileSystemProvider.onAbortRequested.addListener" directly.
func OnAbortRequested(callback js.Func[func(options *AbortRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOnAbortRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAbortRequested calls the function "WEBEXT.fileSystemProvider.onAbortRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAbortRequested(callback js.Func[func(options *AbortRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAbortRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAbortRequested returns true if the function "WEBEXT.fileSystemProvider.onAbortRequested.removeListener" exists.
func HasFuncOffAbortRequested() bool {
	return js.True == bindings.HasFuncOffAbortRequested()
}

// FuncOffAbortRequested returns the function "WEBEXT.fileSystemProvider.onAbortRequested.removeListener".
func FuncOffAbortRequested() (fn js.Func[func(callback js.Func[func(options *AbortRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOffAbortRequested(
		js.Pointer(&fn),
	)
	return
}

// OffAbortRequested calls the function "WEBEXT.fileSystemProvider.onAbortRequested.removeListener" directly.
func OffAbortRequested(callback js.Func[func(options *AbortRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOffAbortRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAbortRequested calls the function "WEBEXT.fileSystemProvider.onAbortRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAbortRequested(callback js.Func[func(options *AbortRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAbortRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAbortRequested returns true if the function "WEBEXT.fileSystemProvider.onAbortRequested.hasListener" exists.
func HasFuncHasOnAbortRequested() bool {
	return js.True == bindings.HasFuncHasOnAbortRequested()
}

// FuncHasOnAbortRequested returns the function "WEBEXT.fileSystemProvider.onAbortRequested.hasListener".
func FuncHasOnAbortRequested() (fn js.Func[func(callback js.Func[func(options *AbortRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) bool]) {
	bindings.FuncHasOnAbortRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnAbortRequested calls the function "WEBEXT.fileSystemProvider.onAbortRequested.hasListener" directly.
func HasOnAbortRequested(callback js.Func[func(options *AbortRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool) {
	bindings.CallHasOnAbortRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAbortRequested calls the function "WEBEXT.fileSystemProvider.onAbortRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnAbortRequested(callback js.Func[func(options *AbortRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnAbortRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnAddWatcherRequestedEventCallbackFunc func(this js.Ref, options *AddWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref

func (fn OnAddWatcherRequestedEventCallbackFunc) Register() js.Func[func(options *AddWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *AddWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnAddWatcherRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AddWatcherRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnAddWatcherRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, options *AddWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref
	Arg T
}

func (cb *OnAddWatcherRequestedEventCallback[T]) Register() js.Func[func(options *AddWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *AddWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnAddWatcherRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AddWatcherRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnAddWatcherRequested returns true if the function "WEBEXT.fileSystemProvider.onAddWatcherRequested.addListener" exists.
func HasFuncOnAddWatcherRequested() bool {
	return js.True == bindings.HasFuncOnAddWatcherRequested()
}

// FuncOnAddWatcherRequested returns the function "WEBEXT.fileSystemProvider.onAddWatcherRequested.addListener".
func FuncOnAddWatcherRequested() (fn js.Func[func(callback js.Func[func(options *AddWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOnAddWatcherRequested(
		js.Pointer(&fn),
	)
	return
}

// OnAddWatcherRequested calls the function "WEBEXT.fileSystemProvider.onAddWatcherRequested.addListener" directly.
func OnAddWatcherRequested(callback js.Func[func(options *AddWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOnAddWatcherRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAddWatcherRequested calls the function "WEBEXT.fileSystemProvider.onAddWatcherRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAddWatcherRequested(callback js.Func[func(options *AddWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAddWatcherRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAddWatcherRequested returns true if the function "WEBEXT.fileSystemProvider.onAddWatcherRequested.removeListener" exists.
func HasFuncOffAddWatcherRequested() bool {
	return js.True == bindings.HasFuncOffAddWatcherRequested()
}

// FuncOffAddWatcherRequested returns the function "WEBEXT.fileSystemProvider.onAddWatcherRequested.removeListener".
func FuncOffAddWatcherRequested() (fn js.Func[func(callback js.Func[func(options *AddWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOffAddWatcherRequested(
		js.Pointer(&fn),
	)
	return
}

// OffAddWatcherRequested calls the function "WEBEXT.fileSystemProvider.onAddWatcherRequested.removeListener" directly.
func OffAddWatcherRequested(callback js.Func[func(options *AddWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOffAddWatcherRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAddWatcherRequested calls the function "WEBEXT.fileSystemProvider.onAddWatcherRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAddWatcherRequested(callback js.Func[func(options *AddWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAddWatcherRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAddWatcherRequested returns true if the function "WEBEXT.fileSystemProvider.onAddWatcherRequested.hasListener" exists.
func HasFuncHasOnAddWatcherRequested() bool {
	return js.True == bindings.HasFuncHasOnAddWatcherRequested()
}

// FuncHasOnAddWatcherRequested returns the function "WEBEXT.fileSystemProvider.onAddWatcherRequested.hasListener".
func FuncHasOnAddWatcherRequested() (fn js.Func[func(callback js.Func[func(options *AddWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) bool]) {
	bindings.FuncHasOnAddWatcherRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnAddWatcherRequested calls the function "WEBEXT.fileSystemProvider.onAddWatcherRequested.hasListener" directly.
func HasOnAddWatcherRequested(callback js.Func[func(options *AddWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool) {
	bindings.CallHasOnAddWatcherRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAddWatcherRequested calls the function "WEBEXT.fileSystemProvider.onAddWatcherRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnAddWatcherRequested(callback js.Func[func(options *AddWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnAddWatcherRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnCloseFileRequestedEventCallbackFunc func(this js.Ref, options *CloseFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref

func (fn OnCloseFileRequestedEventCallbackFunc) Register() js.Func[func(options *CloseFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *CloseFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCloseFileRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CloseFileRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnCloseFileRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, options *CloseFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref
	Arg T
}

func (cb *OnCloseFileRequestedEventCallback[T]) Register() js.Func[func(options *CloseFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *CloseFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCloseFileRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CloseFileRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnCloseFileRequested returns true if the function "WEBEXT.fileSystemProvider.onCloseFileRequested.addListener" exists.
func HasFuncOnCloseFileRequested() bool {
	return js.True == bindings.HasFuncOnCloseFileRequested()
}

// FuncOnCloseFileRequested returns the function "WEBEXT.fileSystemProvider.onCloseFileRequested.addListener".
func FuncOnCloseFileRequested() (fn js.Func[func(callback js.Func[func(options *CloseFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOnCloseFileRequested(
		js.Pointer(&fn),
	)
	return
}

// OnCloseFileRequested calls the function "WEBEXT.fileSystemProvider.onCloseFileRequested.addListener" directly.
func OnCloseFileRequested(callback js.Func[func(options *CloseFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOnCloseFileRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCloseFileRequested calls the function "WEBEXT.fileSystemProvider.onCloseFileRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCloseFileRequested(callback js.Func[func(options *CloseFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCloseFileRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCloseFileRequested returns true if the function "WEBEXT.fileSystemProvider.onCloseFileRequested.removeListener" exists.
func HasFuncOffCloseFileRequested() bool {
	return js.True == bindings.HasFuncOffCloseFileRequested()
}

// FuncOffCloseFileRequested returns the function "WEBEXT.fileSystemProvider.onCloseFileRequested.removeListener".
func FuncOffCloseFileRequested() (fn js.Func[func(callback js.Func[func(options *CloseFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOffCloseFileRequested(
		js.Pointer(&fn),
	)
	return
}

// OffCloseFileRequested calls the function "WEBEXT.fileSystemProvider.onCloseFileRequested.removeListener" directly.
func OffCloseFileRequested(callback js.Func[func(options *CloseFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOffCloseFileRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCloseFileRequested calls the function "WEBEXT.fileSystemProvider.onCloseFileRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCloseFileRequested(callback js.Func[func(options *CloseFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCloseFileRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCloseFileRequested returns true if the function "WEBEXT.fileSystemProvider.onCloseFileRequested.hasListener" exists.
func HasFuncHasOnCloseFileRequested() bool {
	return js.True == bindings.HasFuncHasOnCloseFileRequested()
}

// FuncHasOnCloseFileRequested returns the function "WEBEXT.fileSystemProvider.onCloseFileRequested.hasListener".
func FuncHasOnCloseFileRequested() (fn js.Func[func(callback js.Func[func(options *CloseFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) bool]) {
	bindings.FuncHasOnCloseFileRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnCloseFileRequested calls the function "WEBEXT.fileSystemProvider.onCloseFileRequested.hasListener" directly.
func HasOnCloseFileRequested(callback js.Func[func(options *CloseFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool) {
	bindings.CallHasOnCloseFileRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCloseFileRequested calls the function "WEBEXT.fileSystemProvider.onCloseFileRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCloseFileRequested(callback js.Func[func(options *CloseFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCloseFileRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnConfigureRequestedEventCallbackFunc func(this js.Ref, options *ConfigureRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref

func (fn OnConfigureRequestedEventCallbackFunc) Register() js.Func[func(options *ConfigureRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *ConfigureRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnConfigureRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ConfigureRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnConfigureRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, options *ConfigureRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref
	Arg T
}

func (cb *OnConfigureRequestedEventCallback[T]) Register() js.Func[func(options *ConfigureRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *ConfigureRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnConfigureRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ConfigureRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnConfigureRequested returns true if the function "WEBEXT.fileSystemProvider.onConfigureRequested.addListener" exists.
func HasFuncOnConfigureRequested() bool {
	return js.True == bindings.HasFuncOnConfigureRequested()
}

// FuncOnConfigureRequested returns the function "WEBEXT.fileSystemProvider.onConfigureRequested.addListener".
func FuncOnConfigureRequested() (fn js.Func[func(callback js.Func[func(options *ConfigureRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOnConfigureRequested(
		js.Pointer(&fn),
	)
	return
}

// OnConfigureRequested calls the function "WEBEXT.fileSystemProvider.onConfigureRequested.addListener" directly.
func OnConfigureRequested(callback js.Func[func(options *ConfigureRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOnConfigureRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnConfigureRequested calls the function "WEBEXT.fileSystemProvider.onConfigureRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnConfigureRequested(callback js.Func[func(options *ConfigureRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnConfigureRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffConfigureRequested returns true if the function "WEBEXT.fileSystemProvider.onConfigureRequested.removeListener" exists.
func HasFuncOffConfigureRequested() bool {
	return js.True == bindings.HasFuncOffConfigureRequested()
}

// FuncOffConfigureRequested returns the function "WEBEXT.fileSystemProvider.onConfigureRequested.removeListener".
func FuncOffConfigureRequested() (fn js.Func[func(callback js.Func[func(options *ConfigureRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOffConfigureRequested(
		js.Pointer(&fn),
	)
	return
}

// OffConfigureRequested calls the function "WEBEXT.fileSystemProvider.onConfigureRequested.removeListener" directly.
func OffConfigureRequested(callback js.Func[func(options *ConfigureRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOffConfigureRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffConfigureRequested calls the function "WEBEXT.fileSystemProvider.onConfigureRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffConfigureRequested(callback js.Func[func(options *ConfigureRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffConfigureRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnConfigureRequested returns true if the function "WEBEXT.fileSystemProvider.onConfigureRequested.hasListener" exists.
func HasFuncHasOnConfigureRequested() bool {
	return js.True == bindings.HasFuncHasOnConfigureRequested()
}

// FuncHasOnConfigureRequested returns the function "WEBEXT.fileSystemProvider.onConfigureRequested.hasListener".
func FuncHasOnConfigureRequested() (fn js.Func[func(callback js.Func[func(options *ConfigureRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) bool]) {
	bindings.FuncHasOnConfigureRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnConfigureRequested calls the function "WEBEXT.fileSystemProvider.onConfigureRequested.hasListener" directly.
func HasOnConfigureRequested(callback js.Func[func(options *ConfigureRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool) {
	bindings.CallHasOnConfigureRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnConfigureRequested calls the function "WEBEXT.fileSystemProvider.onConfigureRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnConfigureRequested(callback js.Func[func(options *ConfigureRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnConfigureRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnCopyEntryRequestedEventCallbackFunc func(this js.Ref, options *CopyEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref

func (fn OnCopyEntryRequestedEventCallbackFunc) Register() js.Func[func(options *CopyEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *CopyEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCopyEntryRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CopyEntryRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnCopyEntryRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, options *CopyEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref
	Arg T
}

func (cb *OnCopyEntryRequestedEventCallback[T]) Register() js.Func[func(options *CopyEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *CopyEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCopyEntryRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CopyEntryRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnCopyEntryRequested returns true if the function "WEBEXT.fileSystemProvider.onCopyEntryRequested.addListener" exists.
func HasFuncOnCopyEntryRequested() bool {
	return js.True == bindings.HasFuncOnCopyEntryRequested()
}

// FuncOnCopyEntryRequested returns the function "WEBEXT.fileSystemProvider.onCopyEntryRequested.addListener".
func FuncOnCopyEntryRequested() (fn js.Func[func(callback js.Func[func(options *CopyEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOnCopyEntryRequested(
		js.Pointer(&fn),
	)
	return
}

// OnCopyEntryRequested calls the function "WEBEXT.fileSystemProvider.onCopyEntryRequested.addListener" directly.
func OnCopyEntryRequested(callback js.Func[func(options *CopyEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOnCopyEntryRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCopyEntryRequested calls the function "WEBEXT.fileSystemProvider.onCopyEntryRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCopyEntryRequested(callback js.Func[func(options *CopyEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCopyEntryRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCopyEntryRequested returns true if the function "WEBEXT.fileSystemProvider.onCopyEntryRequested.removeListener" exists.
func HasFuncOffCopyEntryRequested() bool {
	return js.True == bindings.HasFuncOffCopyEntryRequested()
}

// FuncOffCopyEntryRequested returns the function "WEBEXT.fileSystemProvider.onCopyEntryRequested.removeListener".
func FuncOffCopyEntryRequested() (fn js.Func[func(callback js.Func[func(options *CopyEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOffCopyEntryRequested(
		js.Pointer(&fn),
	)
	return
}

// OffCopyEntryRequested calls the function "WEBEXT.fileSystemProvider.onCopyEntryRequested.removeListener" directly.
func OffCopyEntryRequested(callback js.Func[func(options *CopyEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOffCopyEntryRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCopyEntryRequested calls the function "WEBEXT.fileSystemProvider.onCopyEntryRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCopyEntryRequested(callback js.Func[func(options *CopyEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCopyEntryRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCopyEntryRequested returns true if the function "WEBEXT.fileSystemProvider.onCopyEntryRequested.hasListener" exists.
func HasFuncHasOnCopyEntryRequested() bool {
	return js.True == bindings.HasFuncHasOnCopyEntryRequested()
}

// FuncHasOnCopyEntryRequested returns the function "WEBEXT.fileSystemProvider.onCopyEntryRequested.hasListener".
func FuncHasOnCopyEntryRequested() (fn js.Func[func(callback js.Func[func(options *CopyEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) bool]) {
	bindings.FuncHasOnCopyEntryRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnCopyEntryRequested calls the function "WEBEXT.fileSystemProvider.onCopyEntryRequested.hasListener" directly.
func HasOnCopyEntryRequested(callback js.Func[func(options *CopyEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool) {
	bindings.CallHasOnCopyEntryRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCopyEntryRequested calls the function "WEBEXT.fileSystemProvider.onCopyEntryRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCopyEntryRequested(callback js.Func[func(options *CopyEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCopyEntryRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnCreateDirectoryRequestedEventCallbackFunc func(this js.Ref, options *CreateDirectoryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref

func (fn OnCreateDirectoryRequestedEventCallbackFunc) Register() js.Func[func(options *CreateDirectoryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *CreateDirectoryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCreateDirectoryRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CreateDirectoryRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnCreateDirectoryRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, options *CreateDirectoryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref
	Arg T
}

func (cb *OnCreateDirectoryRequestedEventCallback[T]) Register() js.Func[func(options *CreateDirectoryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *CreateDirectoryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCreateDirectoryRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CreateDirectoryRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnCreateDirectoryRequested returns true if the function "WEBEXT.fileSystemProvider.onCreateDirectoryRequested.addListener" exists.
func HasFuncOnCreateDirectoryRequested() bool {
	return js.True == bindings.HasFuncOnCreateDirectoryRequested()
}

// FuncOnCreateDirectoryRequested returns the function "WEBEXT.fileSystemProvider.onCreateDirectoryRequested.addListener".
func FuncOnCreateDirectoryRequested() (fn js.Func[func(callback js.Func[func(options *CreateDirectoryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOnCreateDirectoryRequested(
		js.Pointer(&fn),
	)
	return
}

// OnCreateDirectoryRequested calls the function "WEBEXT.fileSystemProvider.onCreateDirectoryRequested.addListener" directly.
func OnCreateDirectoryRequested(callback js.Func[func(options *CreateDirectoryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOnCreateDirectoryRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCreateDirectoryRequested calls the function "WEBEXT.fileSystemProvider.onCreateDirectoryRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCreateDirectoryRequested(callback js.Func[func(options *CreateDirectoryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCreateDirectoryRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCreateDirectoryRequested returns true if the function "WEBEXT.fileSystemProvider.onCreateDirectoryRequested.removeListener" exists.
func HasFuncOffCreateDirectoryRequested() bool {
	return js.True == bindings.HasFuncOffCreateDirectoryRequested()
}

// FuncOffCreateDirectoryRequested returns the function "WEBEXT.fileSystemProvider.onCreateDirectoryRequested.removeListener".
func FuncOffCreateDirectoryRequested() (fn js.Func[func(callback js.Func[func(options *CreateDirectoryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOffCreateDirectoryRequested(
		js.Pointer(&fn),
	)
	return
}

// OffCreateDirectoryRequested calls the function "WEBEXT.fileSystemProvider.onCreateDirectoryRequested.removeListener" directly.
func OffCreateDirectoryRequested(callback js.Func[func(options *CreateDirectoryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOffCreateDirectoryRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCreateDirectoryRequested calls the function "WEBEXT.fileSystemProvider.onCreateDirectoryRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCreateDirectoryRequested(callback js.Func[func(options *CreateDirectoryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCreateDirectoryRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCreateDirectoryRequested returns true if the function "WEBEXT.fileSystemProvider.onCreateDirectoryRequested.hasListener" exists.
func HasFuncHasOnCreateDirectoryRequested() bool {
	return js.True == bindings.HasFuncHasOnCreateDirectoryRequested()
}

// FuncHasOnCreateDirectoryRequested returns the function "WEBEXT.fileSystemProvider.onCreateDirectoryRequested.hasListener".
func FuncHasOnCreateDirectoryRequested() (fn js.Func[func(callback js.Func[func(options *CreateDirectoryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) bool]) {
	bindings.FuncHasOnCreateDirectoryRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnCreateDirectoryRequested calls the function "WEBEXT.fileSystemProvider.onCreateDirectoryRequested.hasListener" directly.
func HasOnCreateDirectoryRequested(callback js.Func[func(options *CreateDirectoryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool) {
	bindings.CallHasOnCreateDirectoryRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCreateDirectoryRequested calls the function "WEBEXT.fileSystemProvider.onCreateDirectoryRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCreateDirectoryRequested(callback js.Func[func(options *CreateDirectoryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCreateDirectoryRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnCreateFileRequestedEventCallbackFunc func(this js.Ref, options *CreateFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref

func (fn OnCreateFileRequestedEventCallbackFunc) Register() js.Func[func(options *CreateFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *CreateFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCreateFileRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CreateFileRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnCreateFileRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, options *CreateFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref
	Arg T
}

func (cb *OnCreateFileRequestedEventCallback[T]) Register() js.Func[func(options *CreateFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *CreateFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCreateFileRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CreateFileRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnCreateFileRequested returns true if the function "WEBEXT.fileSystemProvider.onCreateFileRequested.addListener" exists.
func HasFuncOnCreateFileRequested() bool {
	return js.True == bindings.HasFuncOnCreateFileRequested()
}

// FuncOnCreateFileRequested returns the function "WEBEXT.fileSystemProvider.onCreateFileRequested.addListener".
func FuncOnCreateFileRequested() (fn js.Func[func(callback js.Func[func(options *CreateFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOnCreateFileRequested(
		js.Pointer(&fn),
	)
	return
}

// OnCreateFileRequested calls the function "WEBEXT.fileSystemProvider.onCreateFileRequested.addListener" directly.
func OnCreateFileRequested(callback js.Func[func(options *CreateFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOnCreateFileRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCreateFileRequested calls the function "WEBEXT.fileSystemProvider.onCreateFileRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCreateFileRequested(callback js.Func[func(options *CreateFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCreateFileRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCreateFileRequested returns true if the function "WEBEXT.fileSystemProvider.onCreateFileRequested.removeListener" exists.
func HasFuncOffCreateFileRequested() bool {
	return js.True == bindings.HasFuncOffCreateFileRequested()
}

// FuncOffCreateFileRequested returns the function "WEBEXT.fileSystemProvider.onCreateFileRequested.removeListener".
func FuncOffCreateFileRequested() (fn js.Func[func(callback js.Func[func(options *CreateFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOffCreateFileRequested(
		js.Pointer(&fn),
	)
	return
}

// OffCreateFileRequested calls the function "WEBEXT.fileSystemProvider.onCreateFileRequested.removeListener" directly.
func OffCreateFileRequested(callback js.Func[func(options *CreateFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOffCreateFileRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCreateFileRequested calls the function "WEBEXT.fileSystemProvider.onCreateFileRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCreateFileRequested(callback js.Func[func(options *CreateFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCreateFileRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCreateFileRequested returns true if the function "WEBEXT.fileSystemProvider.onCreateFileRequested.hasListener" exists.
func HasFuncHasOnCreateFileRequested() bool {
	return js.True == bindings.HasFuncHasOnCreateFileRequested()
}

// FuncHasOnCreateFileRequested returns the function "WEBEXT.fileSystemProvider.onCreateFileRequested.hasListener".
func FuncHasOnCreateFileRequested() (fn js.Func[func(callback js.Func[func(options *CreateFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) bool]) {
	bindings.FuncHasOnCreateFileRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnCreateFileRequested calls the function "WEBEXT.fileSystemProvider.onCreateFileRequested.hasListener" directly.
func HasOnCreateFileRequested(callback js.Func[func(options *CreateFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool) {
	bindings.CallHasOnCreateFileRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCreateFileRequested calls the function "WEBEXT.fileSystemProvider.onCreateFileRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCreateFileRequested(callback js.Func[func(options *CreateFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCreateFileRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDeleteEntryRequestedEventCallbackFunc func(this js.Ref, options *DeleteEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref

func (fn OnDeleteEntryRequestedEventCallbackFunc) Register() js.Func[func(options *DeleteEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *DeleteEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDeleteEntryRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DeleteEntryRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnDeleteEntryRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, options *DeleteEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref
	Arg T
}

func (cb *OnDeleteEntryRequestedEventCallback[T]) Register() js.Func[func(options *DeleteEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *DeleteEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDeleteEntryRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DeleteEntryRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnDeleteEntryRequested returns true if the function "WEBEXT.fileSystemProvider.onDeleteEntryRequested.addListener" exists.
func HasFuncOnDeleteEntryRequested() bool {
	return js.True == bindings.HasFuncOnDeleteEntryRequested()
}

// FuncOnDeleteEntryRequested returns the function "WEBEXT.fileSystemProvider.onDeleteEntryRequested.addListener".
func FuncOnDeleteEntryRequested() (fn js.Func[func(callback js.Func[func(options *DeleteEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOnDeleteEntryRequested(
		js.Pointer(&fn),
	)
	return
}

// OnDeleteEntryRequested calls the function "WEBEXT.fileSystemProvider.onDeleteEntryRequested.addListener" directly.
func OnDeleteEntryRequested(callback js.Func[func(options *DeleteEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOnDeleteEntryRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDeleteEntryRequested calls the function "WEBEXT.fileSystemProvider.onDeleteEntryRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDeleteEntryRequested(callback js.Func[func(options *DeleteEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDeleteEntryRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDeleteEntryRequested returns true if the function "WEBEXT.fileSystemProvider.onDeleteEntryRequested.removeListener" exists.
func HasFuncOffDeleteEntryRequested() bool {
	return js.True == bindings.HasFuncOffDeleteEntryRequested()
}

// FuncOffDeleteEntryRequested returns the function "WEBEXT.fileSystemProvider.onDeleteEntryRequested.removeListener".
func FuncOffDeleteEntryRequested() (fn js.Func[func(callback js.Func[func(options *DeleteEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOffDeleteEntryRequested(
		js.Pointer(&fn),
	)
	return
}

// OffDeleteEntryRequested calls the function "WEBEXT.fileSystemProvider.onDeleteEntryRequested.removeListener" directly.
func OffDeleteEntryRequested(callback js.Func[func(options *DeleteEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOffDeleteEntryRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDeleteEntryRequested calls the function "WEBEXT.fileSystemProvider.onDeleteEntryRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDeleteEntryRequested(callback js.Func[func(options *DeleteEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDeleteEntryRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDeleteEntryRequested returns true if the function "WEBEXT.fileSystemProvider.onDeleteEntryRequested.hasListener" exists.
func HasFuncHasOnDeleteEntryRequested() bool {
	return js.True == bindings.HasFuncHasOnDeleteEntryRequested()
}

// FuncHasOnDeleteEntryRequested returns the function "WEBEXT.fileSystemProvider.onDeleteEntryRequested.hasListener".
func FuncHasOnDeleteEntryRequested() (fn js.Func[func(callback js.Func[func(options *DeleteEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) bool]) {
	bindings.FuncHasOnDeleteEntryRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnDeleteEntryRequested calls the function "WEBEXT.fileSystemProvider.onDeleteEntryRequested.hasListener" directly.
func HasOnDeleteEntryRequested(callback js.Func[func(options *DeleteEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool) {
	bindings.CallHasOnDeleteEntryRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDeleteEntryRequested calls the function "WEBEXT.fileSystemProvider.onDeleteEntryRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDeleteEntryRequested(callback js.Func[func(options *DeleteEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDeleteEntryRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnExecuteActionRequestedEventCallbackFunc func(this js.Ref, options *ExecuteActionRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref

func (fn OnExecuteActionRequestedEventCallbackFunc) Register() js.Func[func(options *ExecuteActionRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *ExecuteActionRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnExecuteActionRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ExecuteActionRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnExecuteActionRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, options *ExecuteActionRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref
	Arg T
}

func (cb *OnExecuteActionRequestedEventCallback[T]) Register() js.Func[func(options *ExecuteActionRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *ExecuteActionRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnExecuteActionRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ExecuteActionRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnExecuteActionRequested returns true if the function "WEBEXT.fileSystemProvider.onExecuteActionRequested.addListener" exists.
func HasFuncOnExecuteActionRequested() bool {
	return js.True == bindings.HasFuncOnExecuteActionRequested()
}

// FuncOnExecuteActionRequested returns the function "WEBEXT.fileSystemProvider.onExecuteActionRequested.addListener".
func FuncOnExecuteActionRequested() (fn js.Func[func(callback js.Func[func(options *ExecuteActionRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOnExecuteActionRequested(
		js.Pointer(&fn),
	)
	return
}

// OnExecuteActionRequested calls the function "WEBEXT.fileSystemProvider.onExecuteActionRequested.addListener" directly.
func OnExecuteActionRequested(callback js.Func[func(options *ExecuteActionRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOnExecuteActionRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnExecuteActionRequested calls the function "WEBEXT.fileSystemProvider.onExecuteActionRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnExecuteActionRequested(callback js.Func[func(options *ExecuteActionRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnExecuteActionRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffExecuteActionRequested returns true if the function "WEBEXT.fileSystemProvider.onExecuteActionRequested.removeListener" exists.
func HasFuncOffExecuteActionRequested() bool {
	return js.True == bindings.HasFuncOffExecuteActionRequested()
}

// FuncOffExecuteActionRequested returns the function "WEBEXT.fileSystemProvider.onExecuteActionRequested.removeListener".
func FuncOffExecuteActionRequested() (fn js.Func[func(callback js.Func[func(options *ExecuteActionRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOffExecuteActionRequested(
		js.Pointer(&fn),
	)
	return
}

// OffExecuteActionRequested calls the function "WEBEXT.fileSystemProvider.onExecuteActionRequested.removeListener" directly.
func OffExecuteActionRequested(callback js.Func[func(options *ExecuteActionRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOffExecuteActionRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffExecuteActionRequested calls the function "WEBEXT.fileSystemProvider.onExecuteActionRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffExecuteActionRequested(callback js.Func[func(options *ExecuteActionRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffExecuteActionRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnExecuteActionRequested returns true if the function "WEBEXT.fileSystemProvider.onExecuteActionRequested.hasListener" exists.
func HasFuncHasOnExecuteActionRequested() bool {
	return js.True == bindings.HasFuncHasOnExecuteActionRequested()
}

// FuncHasOnExecuteActionRequested returns the function "WEBEXT.fileSystemProvider.onExecuteActionRequested.hasListener".
func FuncHasOnExecuteActionRequested() (fn js.Func[func(callback js.Func[func(options *ExecuteActionRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) bool]) {
	bindings.FuncHasOnExecuteActionRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnExecuteActionRequested calls the function "WEBEXT.fileSystemProvider.onExecuteActionRequested.hasListener" directly.
func HasOnExecuteActionRequested(callback js.Func[func(options *ExecuteActionRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool) {
	bindings.CallHasOnExecuteActionRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnExecuteActionRequested calls the function "WEBEXT.fileSystemProvider.onExecuteActionRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnExecuteActionRequested(callback js.Func[func(options *ExecuteActionRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnExecuteActionRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnGetActionsRequestedEventCallbackFunc func(this js.Ref, options *GetActionsRequestedOptions, successCallback js.Func[func(actions js.Array[Action])], errorCallback js.Func[func(err ProviderError)]) js.Ref

func (fn OnGetActionsRequestedEventCallbackFunc) Register() js.Func[func(options *GetActionsRequestedOptions, successCallback js.Func[func(actions js.Array[Action])], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *GetActionsRequestedOptions, successCallback js.Func[func(actions js.Array[Action])], errorCallback js.Func[func(err ProviderError)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnGetActionsRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 GetActionsRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func(actions js.Array[Action])]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnGetActionsRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, options *GetActionsRequestedOptions, successCallback js.Func[func(actions js.Array[Action])], errorCallback js.Func[func(err ProviderError)]) js.Ref
	Arg T
}

func (cb *OnGetActionsRequestedEventCallback[T]) Register() js.Func[func(options *GetActionsRequestedOptions, successCallback js.Func[func(actions js.Array[Action])], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *GetActionsRequestedOptions, successCallback js.Func[func(actions js.Array[Action])], errorCallback js.Func[func(err ProviderError)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnGetActionsRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 GetActionsRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func(actions js.Array[Action])]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnGetActionsRequested returns true if the function "WEBEXT.fileSystemProvider.onGetActionsRequested.addListener" exists.
func HasFuncOnGetActionsRequested() bool {
	return js.True == bindings.HasFuncOnGetActionsRequested()
}

// FuncOnGetActionsRequested returns the function "WEBEXT.fileSystemProvider.onGetActionsRequested.addListener".
func FuncOnGetActionsRequested() (fn js.Func[func(callback js.Func[func(options *GetActionsRequestedOptions, successCallback js.Func[func(actions js.Array[Action])], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOnGetActionsRequested(
		js.Pointer(&fn),
	)
	return
}

// OnGetActionsRequested calls the function "WEBEXT.fileSystemProvider.onGetActionsRequested.addListener" directly.
func OnGetActionsRequested(callback js.Func[func(options *GetActionsRequestedOptions, successCallback js.Func[func(actions js.Array[Action])], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOnGetActionsRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnGetActionsRequested calls the function "WEBEXT.fileSystemProvider.onGetActionsRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnGetActionsRequested(callback js.Func[func(options *GetActionsRequestedOptions, successCallback js.Func[func(actions js.Array[Action])], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnGetActionsRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffGetActionsRequested returns true if the function "WEBEXT.fileSystemProvider.onGetActionsRequested.removeListener" exists.
func HasFuncOffGetActionsRequested() bool {
	return js.True == bindings.HasFuncOffGetActionsRequested()
}

// FuncOffGetActionsRequested returns the function "WEBEXT.fileSystemProvider.onGetActionsRequested.removeListener".
func FuncOffGetActionsRequested() (fn js.Func[func(callback js.Func[func(options *GetActionsRequestedOptions, successCallback js.Func[func(actions js.Array[Action])], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOffGetActionsRequested(
		js.Pointer(&fn),
	)
	return
}

// OffGetActionsRequested calls the function "WEBEXT.fileSystemProvider.onGetActionsRequested.removeListener" directly.
func OffGetActionsRequested(callback js.Func[func(options *GetActionsRequestedOptions, successCallback js.Func[func(actions js.Array[Action])], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOffGetActionsRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffGetActionsRequested calls the function "WEBEXT.fileSystemProvider.onGetActionsRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffGetActionsRequested(callback js.Func[func(options *GetActionsRequestedOptions, successCallback js.Func[func(actions js.Array[Action])], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffGetActionsRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnGetActionsRequested returns true if the function "WEBEXT.fileSystemProvider.onGetActionsRequested.hasListener" exists.
func HasFuncHasOnGetActionsRequested() bool {
	return js.True == bindings.HasFuncHasOnGetActionsRequested()
}

// FuncHasOnGetActionsRequested returns the function "WEBEXT.fileSystemProvider.onGetActionsRequested.hasListener".
func FuncHasOnGetActionsRequested() (fn js.Func[func(callback js.Func[func(options *GetActionsRequestedOptions, successCallback js.Func[func(actions js.Array[Action])], errorCallback js.Func[func(err ProviderError)])]) bool]) {
	bindings.FuncHasOnGetActionsRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnGetActionsRequested calls the function "WEBEXT.fileSystemProvider.onGetActionsRequested.hasListener" directly.
func HasOnGetActionsRequested(callback js.Func[func(options *GetActionsRequestedOptions, successCallback js.Func[func(actions js.Array[Action])], errorCallback js.Func[func(err ProviderError)])]) (ret bool) {
	bindings.CallHasOnGetActionsRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnGetActionsRequested calls the function "WEBEXT.fileSystemProvider.onGetActionsRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnGetActionsRequested(callback js.Func[func(options *GetActionsRequestedOptions, successCallback js.Func[func(actions js.Array[Action])], errorCallback js.Func[func(err ProviderError)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnGetActionsRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnGetMetadataRequestedEventCallbackFunc func(this js.Ref, options *GetMetadataRequestedOptions, successCallback js.Func[func(metadata *EntryMetadata)], errorCallback js.Func[func(err ProviderError)]) js.Ref

func (fn OnGetMetadataRequestedEventCallbackFunc) Register() js.Func[func(options *GetMetadataRequestedOptions, successCallback js.Func[func(metadata *EntryMetadata)], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *GetMetadataRequestedOptions, successCallback js.Func[func(metadata *EntryMetadata)], errorCallback js.Func[func(err ProviderError)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnGetMetadataRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 GetMetadataRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func(metadata *EntryMetadata)]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnGetMetadataRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, options *GetMetadataRequestedOptions, successCallback js.Func[func(metadata *EntryMetadata)], errorCallback js.Func[func(err ProviderError)]) js.Ref
	Arg T
}

func (cb *OnGetMetadataRequestedEventCallback[T]) Register() js.Func[func(options *GetMetadataRequestedOptions, successCallback js.Func[func(metadata *EntryMetadata)], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *GetMetadataRequestedOptions, successCallback js.Func[func(metadata *EntryMetadata)], errorCallback js.Func[func(err ProviderError)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnGetMetadataRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 GetMetadataRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func(metadata *EntryMetadata)]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnGetMetadataRequested returns true if the function "WEBEXT.fileSystemProvider.onGetMetadataRequested.addListener" exists.
func HasFuncOnGetMetadataRequested() bool {
	return js.True == bindings.HasFuncOnGetMetadataRequested()
}

// FuncOnGetMetadataRequested returns the function "WEBEXT.fileSystemProvider.onGetMetadataRequested.addListener".
func FuncOnGetMetadataRequested() (fn js.Func[func(callback js.Func[func(options *GetMetadataRequestedOptions, successCallback js.Func[func(metadata *EntryMetadata)], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOnGetMetadataRequested(
		js.Pointer(&fn),
	)
	return
}

// OnGetMetadataRequested calls the function "WEBEXT.fileSystemProvider.onGetMetadataRequested.addListener" directly.
func OnGetMetadataRequested(callback js.Func[func(options *GetMetadataRequestedOptions, successCallback js.Func[func(metadata *EntryMetadata)], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOnGetMetadataRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnGetMetadataRequested calls the function "WEBEXT.fileSystemProvider.onGetMetadataRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnGetMetadataRequested(callback js.Func[func(options *GetMetadataRequestedOptions, successCallback js.Func[func(metadata *EntryMetadata)], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnGetMetadataRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffGetMetadataRequested returns true if the function "WEBEXT.fileSystemProvider.onGetMetadataRequested.removeListener" exists.
func HasFuncOffGetMetadataRequested() bool {
	return js.True == bindings.HasFuncOffGetMetadataRequested()
}

// FuncOffGetMetadataRequested returns the function "WEBEXT.fileSystemProvider.onGetMetadataRequested.removeListener".
func FuncOffGetMetadataRequested() (fn js.Func[func(callback js.Func[func(options *GetMetadataRequestedOptions, successCallback js.Func[func(metadata *EntryMetadata)], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOffGetMetadataRequested(
		js.Pointer(&fn),
	)
	return
}

// OffGetMetadataRequested calls the function "WEBEXT.fileSystemProvider.onGetMetadataRequested.removeListener" directly.
func OffGetMetadataRequested(callback js.Func[func(options *GetMetadataRequestedOptions, successCallback js.Func[func(metadata *EntryMetadata)], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOffGetMetadataRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffGetMetadataRequested calls the function "WEBEXT.fileSystemProvider.onGetMetadataRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffGetMetadataRequested(callback js.Func[func(options *GetMetadataRequestedOptions, successCallback js.Func[func(metadata *EntryMetadata)], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffGetMetadataRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnGetMetadataRequested returns true if the function "WEBEXT.fileSystemProvider.onGetMetadataRequested.hasListener" exists.
func HasFuncHasOnGetMetadataRequested() bool {
	return js.True == bindings.HasFuncHasOnGetMetadataRequested()
}

// FuncHasOnGetMetadataRequested returns the function "WEBEXT.fileSystemProvider.onGetMetadataRequested.hasListener".
func FuncHasOnGetMetadataRequested() (fn js.Func[func(callback js.Func[func(options *GetMetadataRequestedOptions, successCallback js.Func[func(metadata *EntryMetadata)], errorCallback js.Func[func(err ProviderError)])]) bool]) {
	bindings.FuncHasOnGetMetadataRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnGetMetadataRequested calls the function "WEBEXT.fileSystemProvider.onGetMetadataRequested.hasListener" directly.
func HasOnGetMetadataRequested(callback js.Func[func(options *GetMetadataRequestedOptions, successCallback js.Func[func(metadata *EntryMetadata)], errorCallback js.Func[func(err ProviderError)])]) (ret bool) {
	bindings.CallHasOnGetMetadataRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnGetMetadataRequested calls the function "WEBEXT.fileSystemProvider.onGetMetadataRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnGetMetadataRequested(callback js.Func[func(options *GetMetadataRequestedOptions, successCallback js.Func[func(metadata *EntryMetadata)], errorCallback js.Func[func(err ProviderError)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnGetMetadataRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnMountRequestedEventCallbackFunc func(this js.Ref, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref

func (fn OnMountRequestedEventCallbackFunc) Register() js.Func[func(successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnMountRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Func[func()]{}.FromRef(args[0+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnMountRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref
	Arg T
}

func (cb *OnMountRequestedEventCallback[T]) Register() js.Func[func(successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnMountRequestedEventCallback[T]) DispatchCallback(
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

		js.Func[func()]{}.FromRef(args[0+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnMountRequested returns true if the function "WEBEXT.fileSystemProvider.onMountRequested.addListener" exists.
func HasFuncOnMountRequested() bool {
	return js.True == bindings.HasFuncOnMountRequested()
}

// FuncOnMountRequested returns the function "WEBEXT.fileSystemProvider.onMountRequested.addListener".
func FuncOnMountRequested() (fn js.Func[func(callback js.Func[func(successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOnMountRequested(
		js.Pointer(&fn),
	)
	return
}

// OnMountRequested calls the function "WEBEXT.fileSystemProvider.onMountRequested.addListener" directly.
func OnMountRequested(callback js.Func[func(successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOnMountRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMountRequested calls the function "WEBEXT.fileSystemProvider.onMountRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMountRequested(callback js.Func[func(successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMountRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMountRequested returns true if the function "WEBEXT.fileSystemProvider.onMountRequested.removeListener" exists.
func HasFuncOffMountRequested() bool {
	return js.True == bindings.HasFuncOffMountRequested()
}

// FuncOffMountRequested returns the function "WEBEXT.fileSystemProvider.onMountRequested.removeListener".
func FuncOffMountRequested() (fn js.Func[func(callback js.Func[func(successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOffMountRequested(
		js.Pointer(&fn),
	)
	return
}

// OffMountRequested calls the function "WEBEXT.fileSystemProvider.onMountRequested.removeListener" directly.
func OffMountRequested(callback js.Func[func(successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOffMountRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMountRequested calls the function "WEBEXT.fileSystemProvider.onMountRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMountRequested(callback js.Func[func(successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMountRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMountRequested returns true if the function "WEBEXT.fileSystemProvider.onMountRequested.hasListener" exists.
func HasFuncHasOnMountRequested() bool {
	return js.True == bindings.HasFuncHasOnMountRequested()
}

// FuncHasOnMountRequested returns the function "WEBEXT.fileSystemProvider.onMountRequested.hasListener".
func FuncHasOnMountRequested() (fn js.Func[func(callback js.Func[func(successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) bool]) {
	bindings.FuncHasOnMountRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnMountRequested calls the function "WEBEXT.fileSystemProvider.onMountRequested.hasListener" directly.
func HasOnMountRequested(callback js.Func[func(successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool) {
	bindings.CallHasOnMountRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMountRequested calls the function "WEBEXT.fileSystemProvider.onMountRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnMountRequested(callback js.Func[func(successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnMountRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnMoveEntryRequestedEventCallbackFunc func(this js.Ref, options *MoveEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref

func (fn OnMoveEntryRequestedEventCallbackFunc) Register() js.Func[func(options *MoveEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *MoveEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnMoveEntryRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 MoveEntryRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnMoveEntryRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, options *MoveEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref
	Arg T
}

func (cb *OnMoveEntryRequestedEventCallback[T]) Register() js.Func[func(options *MoveEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *MoveEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnMoveEntryRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 MoveEntryRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnMoveEntryRequested returns true if the function "WEBEXT.fileSystemProvider.onMoveEntryRequested.addListener" exists.
func HasFuncOnMoveEntryRequested() bool {
	return js.True == bindings.HasFuncOnMoveEntryRequested()
}

// FuncOnMoveEntryRequested returns the function "WEBEXT.fileSystemProvider.onMoveEntryRequested.addListener".
func FuncOnMoveEntryRequested() (fn js.Func[func(callback js.Func[func(options *MoveEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOnMoveEntryRequested(
		js.Pointer(&fn),
	)
	return
}

// OnMoveEntryRequested calls the function "WEBEXT.fileSystemProvider.onMoveEntryRequested.addListener" directly.
func OnMoveEntryRequested(callback js.Func[func(options *MoveEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOnMoveEntryRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMoveEntryRequested calls the function "WEBEXT.fileSystemProvider.onMoveEntryRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMoveEntryRequested(callback js.Func[func(options *MoveEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMoveEntryRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMoveEntryRequested returns true if the function "WEBEXT.fileSystemProvider.onMoveEntryRequested.removeListener" exists.
func HasFuncOffMoveEntryRequested() bool {
	return js.True == bindings.HasFuncOffMoveEntryRequested()
}

// FuncOffMoveEntryRequested returns the function "WEBEXT.fileSystemProvider.onMoveEntryRequested.removeListener".
func FuncOffMoveEntryRequested() (fn js.Func[func(callback js.Func[func(options *MoveEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOffMoveEntryRequested(
		js.Pointer(&fn),
	)
	return
}

// OffMoveEntryRequested calls the function "WEBEXT.fileSystemProvider.onMoveEntryRequested.removeListener" directly.
func OffMoveEntryRequested(callback js.Func[func(options *MoveEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOffMoveEntryRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMoveEntryRequested calls the function "WEBEXT.fileSystemProvider.onMoveEntryRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMoveEntryRequested(callback js.Func[func(options *MoveEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMoveEntryRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMoveEntryRequested returns true if the function "WEBEXT.fileSystemProvider.onMoveEntryRequested.hasListener" exists.
func HasFuncHasOnMoveEntryRequested() bool {
	return js.True == bindings.HasFuncHasOnMoveEntryRequested()
}

// FuncHasOnMoveEntryRequested returns the function "WEBEXT.fileSystemProvider.onMoveEntryRequested.hasListener".
func FuncHasOnMoveEntryRequested() (fn js.Func[func(callback js.Func[func(options *MoveEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) bool]) {
	bindings.FuncHasOnMoveEntryRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnMoveEntryRequested calls the function "WEBEXT.fileSystemProvider.onMoveEntryRequested.hasListener" directly.
func HasOnMoveEntryRequested(callback js.Func[func(options *MoveEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool) {
	bindings.CallHasOnMoveEntryRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMoveEntryRequested calls the function "WEBEXT.fileSystemProvider.onMoveEntryRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnMoveEntryRequested(callback js.Func[func(options *MoveEntryRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnMoveEntryRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnOpenFileRequestedEventCallbackFunc func(this js.Ref, options *OpenFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref

func (fn OnOpenFileRequestedEventCallbackFunc) Register() js.Func[func(options *OpenFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *OpenFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnOpenFileRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OpenFileRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnOpenFileRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, options *OpenFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref
	Arg T
}

func (cb *OnOpenFileRequestedEventCallback[T]) Register() js.Func[func(options *OpenFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *OpenFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnOpenFileRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OpenFileRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnOpenFileRequested returns true if the function "WEBEXT.fileSystemProvider.onOpenFileRequested.addListener" exists.
func HasFuncOnOpenFileRequested() bool {
	return js.True == bindings.HasFuncOnOpenFileRequested()
}

// FuncOnOpenFileRequested returns the function "WEBEXT.fileSystemProvider.onOpenFileRequested.addListener".
func FuncOnOpenFileRequested() (fn js.Func[func(callback js.Func[func(options *OpenFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOnOpenFileRequested(
		js.Pointer(&fn),
	)
	return
}

// OnOpenFileRequested calls the function "WEBEXT.fileSystemProvider.onOpenFileRequested.addListener" directly.
func OnOpenFileRequested(callback js.Func[func(options *OpenFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOnOpenFileRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnOpenFileRequested calls the function "WEBEXT.fileSystemProvider.onOpenFileRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnOpenFileRequested(callback js.Func[func(options *OpenFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnOpenFileRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffOpenFileRequested returns true if the function "WEBEXT.fileSystemProvider.onOpenFileRequested.removeListener" exists.
func HasFuncOffOpenFileRequested() bool {
	return js.True == bindings.HasFuncOffOpenFileRequested()
}

// FuncOffOpenFileRequested returns the function "WEBEXT.fileSystemProvider.onOpenFileRequested.removeListener".
func FuncOffOpenFileRequested() (fn js.Func[func(callback js.Func[func(options *OpenFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOffOpenFileRequested(
		js.Pointer(&fn),
	)
	return
}

// OffOpenFileRequested calls the function "WEBEXT.fileSystemProvider.onOpenFileRequested.removeListener" directly.
func OffOpenFileRequested(callback js.Func[func(options *OpenFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOffOpenFileRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffOpenFileRequested calls the function "WEBEXT.fileSystemProvider.onOpenFileRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffOpenFileRequested(callback js.Func[func(options *OpenFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffOpenFileRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnOpenFileRequested returns true if the function "WEBEXT.fileSystemProvider.onOpenFileRequested.hasListener" exists.
func HasFuncHasOnOpenFileRequested() bool {
	return js.True == bindings.HasFuncHasOnOpenFileRequested()
}

// FuncHasOnOpenFileRequested returns the function "WEBEXT.fileSystemProvider.onOpenFileRequested.hasListener".
func FuncHasOnOpenFileRequested() (fn js.Func[func(callback js.Func[func(options *OpenFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) bool]) {
	bindings.FuncHasOnOpenFileRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnOpenFileRequested calls the function "WEBEXT.fileSystemProvider.onOpenFileRequested.hasListener" directly.
func HasOnOpenFileRequested(callback js.Func[func(options *OpenFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool) {
	bindings.CallHasOnOpenFileRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnOpenFileRequested calls the function "WEBEXT.fileSystemProvider.onOpenFileRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnOpenFileRequested(callback js.Func[func(options *OpenFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnOpenFileRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnReadDirectoryRequestedEventCallbackFunc func(this js.Ref, options *ReadDirectoryRequestedOptions, successCallback js.Func[func(entries js.Array[EntryMetadata], hasMore bool)], errorCallback js.Func[func(err ProviderError)]) js.Ref

func (fn OnReadDirectoryRequestedEventCallbackFunc) Register() js.Func[func(options *ReadDirectoryRequestedOptions, successCallback js.Func[func(entries js.Array[EntryMetadata], hasMore bool)], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *ReadDirectoryRequestedOptions, successCallback js.Func[func(entries js.Array[EntryMetadata], hasMore bool)], errorCallback js.Func[func(err ProviderError)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnReadDirectoryRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ReadDirectoryRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func(entries js.Array[EntryMetadata], hasMore bool)]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnReadDirectoryRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, options *ReadDirectoryRequestedOptions, successCallback js.Func[func(entries js.Array[EntryMetadata], hasMore bool)], errorCallback js.Func[func(err ProviderError)]) js.Ref
	Arg T
}

func (cb *OnReadDirectoryRequestedEventCallback[T]) Register() js.Func[func(options *ReadDirectoryRequestedOptions, successCallback js.Func[func(entries js.Array[EntryMetadata], hasMore bool)], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *ReadDirectoryRequestedOptions, successCallback js.Func[func(entries js.Array[EntryMetadata], hasMore bool)], errorCallback js.Func[func(err ProviderError)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnReadDirectoryRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ReadDirectoryRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func(entries js.Array[EntryMetadata], hasMore bool)]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnReadDirectoryRequested returns true if the function "WEBEXT.fileSystemProvider.onReadDirectoryRequested.addListener" exists.
func HasFuncOnReadDirectoryRequested() bool {
	return js.True == bindings.HasFuncOnReadDirectoryRequested()
}

// FuncOnReadDirectoryRequested returns the function "WEBEXT.fileSystemProvider.onReadDirectoryRequested.addListener".
func FuncOnReadDirectoryRequested() (fn js.Func[func(callback js.Func[func(options *ReadDirectoryRequestedOptions, successCallback js.Func[func(entries js.Array[EntryMetadata], hasMore bool)], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOnReadDirectoryRequested(
		js.Pointer(&fn),
	)
	return
}

// OnReadDirectoryRequested calls the function "WEBEXT.fileSystemProvider.onReadDirectoryRequested.addListener" directly.
func OnReadDirectoryRequested(callback js.Func[func(options *ReadDirectoryRequestedOptions, successCallback js.Func[func(entries js.Array[EntryMetadata], hasMore bool)], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOnReadDirectoryRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnReadDirectoryRequested calls the function "WEBEXT.fileSystemProvider.onReadDirectoryRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnReadDirectoryRequested(callback js.Func[func(options *ReadDirectoryRequestedOptions, successCallback js.Func[func(entries js.Array[EntryMetadata], hasMore bool)], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnReadDirectoryRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffReadDirectoryRequested returns true if the function "WEBEXT.fileSystemProvider.onReadDirectoryRequested.removeListener" exists.
func HasFuncOffReadDirectoryRequested() bool {
	return js.True == bindings.HasFuncOffReadDirectoryRequested()
}

// FuncOffReadDirectoryRequested returns the function "WEBEXT.fileSystemProvider.onReadDirectoryRequested.removeListener".
func FuncOffReadDirectoryRequested() (fn js.Func[func(callback js.Func[func(options *ReadDirectoryRequestedOptions, successCallback js.Func[func(entries js.Array[EntryMetadata], hasMore bool)], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOffReadDirectoryRequested(
		js.Pointer(&fn),
	)
	return
}

// OffReadDirectoryRequested calls the function "WEBEXT.fileSystemProvider.onReadDirectoryRequested.removeListener" directly.
func OffReadDirectoryRequested(callback js.Func[func(options *ReadDirectoryRequestedOptions, successCallback js.Func[func(entries js.Array[EntryMetadata], hasMore bool)], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOffReadDirectoryRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffReadDirectoryRequested calls the function "WEBEXT.fileSystemProvider.onReadDirectoryRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffReadDirectoryRequested(callback js.Func[func(options *ReadDirectoryRequestedOptions, successCallback js.Func[func(entries js.Array[EntryMetadata], hasMore bool)], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffReadDirectoryRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnReadDirectoryRequested returns true if the function "WEBEXT.fileSystemProvider.onReadDirectoryRequested.hasListener" exists.
func HasFuncHasOnReadDirectoryRequested() bool {
	return js.True == bindings.HasFuncHasOnReadDirectoryRequested()
}

// FuncHasOnReadDirectoryRequested returns the function "WEBEXT.fileSystemProvider.onReadDirectoryRequested.hasListener".
func FuncHasOnReadDirectoryRequested() (fn js.Func[func(callback js.Func[func(options *ReadDirectoryRequestedOptions, successCallback js.Func[func(entries js.Array[EntryMetadata], hasMore bool)], errorCallback js.Func[func(err ProviderError)])]) bool]) {
	bindings.FuncHasOnReadDirectoryRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnReadDirectoryRequested calls the function "WEBEXT.fileSystemProvider.onReadDirectoryRequested.hasListener" directly.
func HasOnReadDirectoryRequested(callback js.Func[func(options *ReadDirectoryRequestedOptions, successCallback js.Func[func(entries js.Array[EntryMetadata], hasMore bool)], errorCallback js.Func[func(err ProviderError)])]) (ret bool) {
	bindings.CallHasOnReadDirectoryRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnReadDirectoryRequested calls the function "WEBEXT.fileSystemProvider.onReadDirectoryRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnReadDirectoryRequested(callback js.Func[func(options *ReadDirectoryRequestedOptions, successCallback js.Func[func(entries js.Array[EntryMetadata], hasMore bool)], errorCallback js.Func[func(err ProviderError)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnReadDirectoryRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnReadFileRequestedEventCallbackFunc func(this js.Ref, options *ReadFileRequestedOptions, successCallback js.Func[func(data js.ArrayBuffer, hasMore bool)], errorCallback js.Func[func(err ProviderError)]) js.Ref

func (fn OnReadFileRequestedEventCallbackFunc) Register() js.Func[func(options *ReadFileRequestedOptions, successCallback js.Func[func(data js.ArrayBuffer, hasMore bool)], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *ReadFileRequestedOptions, successCallback js.Func[func(data js.ArrayBuffer, hasMore bool)], errorCallback js.Func[func(err ProviderError)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnReadFileRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ReadFileRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func(data js.ArrayBuffer, hasMore bool)]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnReadFileRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, options *ReadFileRequestedOptions, successCallback js.Func[func(data js.ArrayBuffer, hasMore bool)], errorCallback js.Func[func(err ProviderError)]) js.Ref
	Arg T
}

func (cb *OnReadFileRequestedEventCallback[T]) Register() js.Func[func(options *ReadFileRequestedOptions, successCallback js.Func[func(data js.ArrayBuffer, hasMore bool)], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *ReadFileRequestedOptions, successCallback js.Func[func(data js.ArrayBuffer, hasMore bool)], errorCallback js.Func[func(err ProviderError)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnReadFileRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ReadFileRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func(data js.ArrayBuffer, hasMore bool)]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnReadFileRequested returns true if the function "WEBEXT.fileSystemProvider.onReadFileRequested.addListener" exists.
func HasFuncOnReadFileRequested() bool {
	return js.True == bindings.HasFuncOnReadFileRequested()
}

// FuncOnReadFileRequested returns the function "WEBEXT.fileSystemProvider.onReadFileRequested.addListener".
func FuncOnReadFileRequested() (fn js.Func[func(callback js.Func[func(options *ReadFileRequestedOptions, successCallback js.Func[func(data js.ArrayBuffer, hasMore bool)], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOnReadFileRequested(
		js.Pointer(&fn),
	)
	return
}

// OnReadFileRequested calls the function "WEBEXT.fileSystemProvider.onReadFileRequested.addListener" directly.
func OnReadFileRequested(callback js.Func[func(options *ReadFileRequestedOptions, successCallback js.Func[func(data js.ArrayBuffer, hasMore bool)], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOnReadFileRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnReadFileRequested calls the function "WEBEXT.fileSystemProvider.onReadFileRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnReadFileRequested(callback js.Func[func(options *ReadFileRequestedOptions, successCallback js.Func[func(data js.ArrayBuffer, hasMore bool)], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnReadFileRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffReadFileRequested returns true if the function "WEBEXT.fileSystemProvider.onReadFileRequested.removeListener" exists.
func HasFuncOffReadFileRequested() bool {
	return js.True == bindings.HasFuncOffReadFileRequested()
}

// FuncOffReadFileRequested returns the function "WEBEXT.fileSystemProvider.onReadFileRequested.removeListener".
func FuncOffReadFileRequested() (fn js.Func[func(callback js.Func[func(options *ReadFileRequestedOptions, successCallback js.Func[func(data js.ArrayBuffer, hasMore bool)], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOffReadFileRequested(
		js.Pointer(&fn),
	)
	return
}

// OffReadFileRequested calls the function "WEBEXT.fileSystemProvider.onReadFileRequested.removeListener" directly.
func OffReadFileRequested(callback js.Func[func(options *ReadFileRequestedOptions, successCallback js.Func[func(data js.ArrayBuffer, hasMore bool)], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOffReadFileRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffReadFileRequested calls the function "WEBEXT.fileSystemProvider.onReadFileRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffReadFileRequested(callback js.Func[func(options *ReadFileRequestedOptions, successCallback js.Func[func(data js.ArrayBuffer, hasMore bool)], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffReadFileRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnReadFileRequested returns true if the function "WEBEXT.fileSystemProvider.onReadFileRequested.hasListener" exists.
func HasFuncHasOnReadFileRequested() bool {
	return js.True == bindings.HasFuncHasOnReadFileRequested()
}

// FuncHasOnReadFileRequested returns the function "WEBEXT.fileSystemProvider.onReadFileRequested.hasListener".
func FuncHasOnReadFileRequested() (fn js.Func[func(callback js.Func[func(options *ReadFileRequestedOptions, successCallback js.Func[func(data js.ArrayBuffer, hasMore bool)], errorCallback js.Func[func(err ProviderError)])]) bool]) {
	bindings.FuncHasOnReadFileRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnReadFileRequested calls the function "WEBEXT.fileSystemProvider.onReadFileRequested.hasListener" directly.
func HasOnReadFileRequested(callback js.Func[func(options *ReadFileRequestedOptions, successCallback js.Func[func(data js.ArrayBuffer, hasMore bool)], errorCallback js.Func[func(err ProviderError)])]) (ret bool) {
	bindings.CallHasOnReadFileRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnReadFileRequested calls the function "WEBEXT.fileSystemProvider.onReadFileRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnReadFileRequested(callback js.Func[func(options *ReadFileRequestedOptions, successCallback js.Func[func(data js.ArrayBuffer, hasMore bool)], errorCallback js.Func[func(err ProviderError)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnReadFileRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnRemoveWatcherRequestedEventCallbackFunc func(this js.Ref, options *RemoveWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref

func (fn OnRemoveWatcherRequestedEventCallbackFunc) Register() js.Func[func(options *RemoveWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *RemoveWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnRemoveWatcherRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RemoveWatcherRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnRemoveWatcherRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, options *RemoveWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref
	Arg T
}

func (cb *OnRemoveWatcherRequestedEventCallback[T]) Register() js.Func[func(options *RemoveWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *RemoveWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnRemoveWatcherRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RemoveWatcherRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnRemoveWatcherRequested returns true if the function "WEBEXT.fileSystemProvider.onRemoveWatcherRequested.addListener" exists.
func HasFuncOnRemoveWatcherRequested() bool {
	return js.True == bindings.HasFuncOnRemoveWatcherRequested()
}

// FuncOnRemoveWatcherRequested returns the function "WEBEXT.fileSystemProvider.onRemoveWatcherRequested.addListener".
func FuncOnRemoveWatcherRequested() (fn js.Func[func(callback js.Func[func(options *RemoveWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOnRemoveWatcherRequested(
		js.Pointer(&fn),
	)
	return
}

// OnRemoveWatcherRequested calls the function "WEBEXT.fileSystemProvider.onRemoveWatcherRequested.addListener" directly.
func OnRemoveWatcherRequested(callback js.Func[func(options *RemoveWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOnRemoveWatcherRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnRemoveWatcherRequested calls the function "WEBEXT.fileSystemProvider.onRemoveWatcherRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnRemoveWatcherRequested(callback js.Func[func(options *RemoveWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnRemoveWatcherRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffRemoveWatcherRequested returns true if the function "WEBEXT.fileSystemProvider.onRemoveWatcherRequested.removeListener" exists.
func HasFuncOffRemoveWatcherRequested() bool {
	return js.True == bindings.HasFuncOffRemoveWatcherRequested()
}

// FuncOffRemoveWatcherRequested returns the function "WEBEXT.fileSystemProvider.onRemoveWatcherRequested.removeListener".
func FuncOffRemoveWatcherRequested() (fn js.Func[func(callback js.Func[func(options *RemoveWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOffRemoveWatcherRequested(
		js.Pointer(&fn),
	)
	return
}

// OffRemoveWatcherRequested calls the function "WEBEXT.fileSystemProvider.onRemoveWatcherRequested.removeListener" directly.
func OffRemoveWatcherRequested(callback js.Func[func(options *RemoveWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOffRemoveWatcherRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffRemoveWatcherRequested calls the function "WEBEXT.fileSystemProvider.onRemoveWatcherRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffRemoveWatcherRequested(callback js.Func[func(options *RemoveWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffRemoveWatcherRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnRemoveWatcherRequested returns true if the function "WEBEXT.fileSystemProvider.onRemoveWatcherRequested.hasListener" exists.
func HasFuncHasOnRemoveWatcherRequested() bool {
	return js.True == bindings.HasFuncHasOnRemoveWatcherRequested()
}

// FuncHasOnRemoveWatcherRequested returns the function "WEBEXT.fileSystemProvider.onRemoveWatcherRequested.hasListener".
func FuncHasOnRemoveWatcherRequested() (fn js.Func[func(callback js.Func[func(options *RemoveWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) bool]) {
	bindings.FuncHasOnRemoveWatcherRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnRemoveWatcherRequested calls the function "WEBEXT.fileSystemProvider.onRemoveWatcherRequested.hasListener" directly.
func HasOnRemoveWatcherRequested(callback js.Func[func(options *RemoveWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool) {
	bindings.CallHasOnRemoveWatcherRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnRemoveWatcherRequested calls the function "WEBEXT.fileSystemProvider.onRemoveWatcherRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnRemoveWatcherRequested(callback js.Func[func(options *RemoveWatcherRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnRemoveWatcherRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnTruncateRequestedEventCallbackFunc func(this js.Ref, options *TruncateRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref

func (fn OnTruncateRequestedEventCallbackFunc) Register() js.Func[func(options *TruncateRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *TruncateRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnTruncateRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TruncateRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnTruncateRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, options *TruncateRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref
	Arg T
}

func (cb *OnTruncateRequestedEventCallback[T]) Register() js.Func[func(options *TruncateRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *TruncateRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnTruncateRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TruncateRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnTruncateRequested returns true if the function "WEBEXT.fileSystemProvider.onTruncateRequested.addListener" exists.
func HasFuncOnTruncateRequested() bool {
	return js.True == bindings.HasFuncOnTruncateRequested()
}

// FuncOnTruncateRequested returns the function "WEBEXT.fileSystemProvider.onTruncateRequested.addListener".
func FuncOnTruncateRequested() (fn js.Func[func(callback js.Func[func(options *TruncateRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOnTruncateRequested(
		js.Pointer(&fn),
	)
	return
}

// OnTruncateRequested calls the function "WEBEXT.fileSystemProvider.onTruncateRequested.addListener" directly.
func OnTruncateRequested(callback js.Func[func(options *TruncateRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOnTruncateRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnTruncateRequested calls the function "WEBEXT.fileSystemProvider.onTruncateRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnTruncateRequested(callback js.Func[func(options *TruncateRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnTruncateRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffTruncateRequested returns true if the function "WEBEXT.fileSystemProvider.onTruncateRequested.removeListener" exists.
func HasFuncOffTruncateRequested() bool {
	return js.True == bindings.HasFuncOffTruncateRequested()
}

// FuncOffTruncateRequested returns the function "WEBEXT.fileSystemProvider.onTruncateRequested.removeListener".
func FuncOffTruncateRequested() (fn js.Func[func(callback js.Func[func(options *TruncateRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOffTruncateRequested(
		js.Pointer(&fn),
	)
	return
}

// OffTruncateRequested calls the function "WEBEXT.fileSystemProvider.onTruncateRequested.removeListener" directly.
func OffTruncateRequested(callback js.Func[func(options *TruncateRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOffTruncateRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffTruncateRequested calls the function "WEBEXT.fileSystemProvider.onTruncateRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffTruncateRequested(callback js.Func[func(options *TruncateRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffTruncateRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnTruncateRequested returns true if the function "WEBEXT.fileSystemProvider.onTruncateRequested.hasListener" exists.
func HasFuncHasOnTruncateRequested() bool {
	return js.True == bindings.HasFuncHasOnTruncateRequested()
}

// FuncHasOnTruncateRequested returns the function "WEBEXT.fileSystemProvider.onTruncateRequested.hasListener".
func FuncHasOnTruncateRequested() (fn js.Func[func(callback js.Func[func(options *TruncateRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) bool]) {
	bindings.FuncHasOnTruncateRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnTruncateRequested calls the function "WEBEXT.fileSystemProvider.onTruncateRequested.hasListener" directly.
func HasOnTruncateRequested(callback js.Func[func(options *TruncateRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool) {
	bindings.CallHasOnTruncateRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnTruncateRequested calls the function "WEBEXT.fileSystemProvider.onTruncateRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnTruncateRequested(callback js.Func[func(options *TruncateRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnTruncateRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnUnmountRequestedEventCallbackFunc func(this js.Ref, options *UnmountRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref

func (fn OnUnmountRequestedEventCallbackFunc) Register() js.Func[func(options *UnmountRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *UnmountRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnUnmountRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 UnmountRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnUnmountRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, options *UnmountRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref
	Arg T
}

func (cb *OnUnmountRequestedEventCallback[T]) Register() js.Func[func(options *UnmountRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *UnmountRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnUnmountRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 UnmountRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnUnmountRequested returns true if the function "WEBEXT.fileSystemProvider.onUnmountRequested.addListener" exists.
func HasFuncOnUnmountRequested() bool {
	return js.True == bindings.HasFuncOnUnmountRequested()
}

// FuncOnUnmountRequested returns the function "WEBEXT.fileSystemProvider.onUnmountRequested.addListener".
func FuncOnUnmountRequested() (fn js.Func[func(callback js.Func[func(options *UnmountRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOnUnmountRequested(
		js.Pointer(&fn),
	)
	return
}

// OnUnmountRequested calls the function "WEBEXT.fileSystemProvider.onUnmountRequested.addListener" directly.
func OnUnmountRequested(callback js.Func[func(options *UnmountRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOnUnmountRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnUnmountRequested calls the function "WEBEXT.fileSystemProvider.onUnmountRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnUnmountRequested(callback js.Func[func(options *UnmountRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnUnmountRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffUnmountRequested returns true if the function "WEBEXT.fileSystemProvider.onUnmountRequested.removeListener" exists.
func HasFuncOffUnmountRequested() bool {
	return js.True == bindings.HasFuncOffUnmountRequested()
}

// FuncOffUnmountRequested returns the function "WEBEXT.fileSystemProvider.onUnmountRequested.removeListener".
func FuncOffUnmountRequested() (fn js.Func[func(callback js.Func[func(options *UnmountRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOffUnmountRequested(
		js.Pointer(&fn),
	)
	return
}

// OffUnmountRequested calls the function "WEBEXT.fileSystemProvider.onUnmountRequested.removeListener" directly.
func OffUnmountRequested(callback js.Func[func(options *UnmountRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOffUnmountRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffUnmountRequested calls the function "WEBEXT.fileSystemProvider.onUnmountRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffUnmountRequested(callback js.Func[func(options *UnmountRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffUnmountRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnUnmountRequested returns true if the function "WEBEXT.fileSystemProvider.onUnmountRequested.hasListener" exists.
func HasFuncHasOnUnmountRequested() bool {
	return js.True == bindings.HasFuncHasOnUnmountRequested()
}

// FuncHasOnUnmountRequested returns the function "WEBEXT.fileSystemProvider.onUnmountRequested.hasListener".
func FuncHasOnUnmountRequested() (fn js.Func[func(callback js.Func[func(options *UnmountRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) bool]) {
	bindings.FuncHasOnUnmountRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnUnmountRequested calls the function "WEBEXT.fileSystemProvider.onUnmountRequested.hasListener" directly.
func HasOnUnmountRequested(callback js.Func[func(options *UnmountRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool) {
	bindings.CallHasOnUnmountRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnUnmountRequested calls the function "WEBEXT.fileSystemProvider.onUnmountRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnUnmountRequested(callback js.Func[func(options *UnmountRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnUnmountRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnWriteFileRequestedEventCallbackFunc func(this js.Ref, options *WriteFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref

func (fn OnWriteFileRequestedEventCallbackFunc) Register() js.Func[func(options *WriteFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *WriteFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnWriteFileRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 WriteFileRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnWriteFileRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, options *WriteFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)]) js.Ref
	Arg T
}

func (cb *OnWriteFileRequestedEventCallback[T]) Register() js.Func[func(options *WriteFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])] {
	return js.RegisterCallback[func(options *WriteFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnWriteFileRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 WriteFileRequestedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func()]{}.FromRef(args[1+1]),
		js.Func[func(err ProviderError)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnWriteFileRequested returns true if the function "WEBEXT.fileSystemProvider.onWriteFileRequested.addListener" exists.
func HasFuncOnWriteFileRequested() bool {
	return js.True == bindings.HasFuncOnWriteFileRequested()
}

// FuncOnWriteFileRequested returns the function "WEBEXT.fileSystemProvider.onWriteFileRequested.addListener".
func FuncOnWriteFileRequested() (fn js.Func[func(callback js.Func[func(options *WriteFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOnWriteFileRequested(
		js.Pointer(&fn),
	)
	return
}

// OnWriteFileRequested calls the function "WEBEXT.fileSystemProvider.onWriteFileRequested.addListener" directly.
func OnWriteFileRequested(callback js.Func[func(options *WriteFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOnWriteFileRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnWriteFileRequested calls the function "WEBEXT.fileSystemProvider.onWriteFileRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnWriteFileRequested(callback js.Func[func(options *WriteFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnWriteFileRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffWriteFileRequested returns true if the function "WEBEXT.fileSystemProvider.onWriteFileRequested.removeListener" exists.
func HasFuncOffWriteFileRequested() bool {
	return js.True == bindings.HasFuncOffWriteFileRequested()
}

// FuncOffWriteFileRequested returns the function "WEBEXT.fileSystemProvider.onWriteFileRequested.removeListener".
func FuncOffWriteFileRequested() (fn js.Func[func(callback js.Func[func(options *WriteFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])])]) {
	bindings.FuncOffWriteFileRequested(
		js.Pointer(&fn),
	)
	return
}

// OffWriteFileRequested calls the function "WEBEXT.fileSystemProvider.onWriteFileRequested.removeListener" directly.
func OffWriteFileRequested(callback js.Func[func(options *WriteFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void) {
	bindings.CallOffWriteFileRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffWriteFileRequested calls the function "WEBEXT.fileSystemProvider.onWriteFileRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffWriteFileRequested(callback js.Func[func(options *WriteFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffWriteFileRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnWriteFileRequested returns true if the function "WEBEXT.fileSystemProvider.onWriteFileRequested.hasListener" exists.
func HasFuncHasOnWriteFileRequested() bool {
	return js.True == bindings.HasFuncHasOnWriteFileRequested()
}

// FuncHasOnWriteFileRequested returns the function "WEBEXT.fileSystemProvider.onWriteFileRequested.hasListener".
func FuncHasOnWriteFileRequested() (fn js.Func[func(callback js.Func[func(options *WriteFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) bool]) {
	bindings.FuncHasOnWriteFileRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnWriteFileRequested calls the function "WEBEXT.fileSystemProvider.onWriteFileRequested.hasListener" directly.
func HasOnWriteFileRequested(callback js.Func[func(options *WriteFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool) {
	bindings.CallHasOnWriteFileRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnWriteFileRequested calls the function "WEBEXT.fileSystemProvider.onWriteFileRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnWriteFileRequested(callback js.Func[func(options *WriteFileRequestedOptions, successCallback js.Func[func()], errorCallback js.Func[func(err ProviderError)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnWriteFileRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncUnmount returns true if the function "WEBEXT.fileSystemProvider.unmount" exists.
func HasFuncUnmount() bool {
	return js.True == bindings.HasFuncUnmount()
}

// FuncUnmount returns the function "WEBEXT.fileSystemProvider.unmount".
func FuncUnmount() (fn js.Func[func(options UnmountOptions) js.Promise[js.Void]]) {
	bindings.FuncUnmount(
		js.Pointer(&fn),
	)
	return
}

// Unmount calls the function "WEBEXT.fileSystemProvider.unmount" directly.
func Unmount(options UnmountOptions) (ret js.Promise[js.Void]) {
	bindings.CallUnmount(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryUnmount calls the function "WEBEXT.fileSystemProvider.unmount"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUnmount(options UnmountOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUnmount(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}
