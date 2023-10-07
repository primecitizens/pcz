// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package downloads

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/downloads/bindings"
)

type BooleanDelta struct {
	// Previous is "BooleanDelta.previous"
	//
	// Optional
	//
	// NOTE: FFI_USE_Previous MUST be set to true to make this field effective.
	Previous bool
	// Current is "BooleanDelta.current"
	//
	// Optional
	//
	// NOTE: FFI_USE_Current MUST be set to true to make this field effective.
	Current bool

	FFI_USE_Previous bool // for Previous.
	FFI_USE_Current  bool // for Current.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BooleanDelta with all fields set.
func (p BooleanDelta) FromRef(ref js.Ref) BooleanDelta {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BooleanDelta in the application heap.
func (p BooleanDelta) New() js.Ref {
	return bindings.BooleanDeltaJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *BooleanDelta) UpdateFrom(ref js.Ref) {
	bindings.BooleanDeltaJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BooleanDelta) Update(ref js.Ref) {
	bindings.BooleanDeltaJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BooleanDelta) FreeMembers(recursive bool) {
}

type DangerType uint32

const (
	_ DangerType = iota

	DangerType_FILE
	DangerType_URL
	DangerType_CONTENT
	DangerType_UNCOMMON
	DangerType_HOST
	DangerType_UNWANTED
	DangerType_SAFE
	DangerType_ACCEPTED
	DangerType_ALLOWLISTED_BY_POLICY
	DangerType_ASYNC_SCANNING
	DangerType_PASSWORD_PROTECTED
	DangerType_BLOCKED_TOO_LARGE
	DangerType_SENSITIVE_CONTENT_WARNING
	DangerType_SENSITIVE_CONTENT_BLOCK
	DangerType_UNSUPPORTED_FILE_TYPE
	DangerType_DEEP_SCANNED_FAILED
	DangerType_DEEP_SCANNED_SAFE
	DangerType_DEEP_SCANNED_OPENED_DANGEROUS
	DangerType_PROMPT_FOR_SCANING
	DangerType_ACCOUNT_COMPROMISE
)

func (DangerType) FromRef(str js.Ref) DangerType {
	return DangerType(bindings.ConstOfDangerType(str))
}

func (x DangerType) String() (string, bool) {
	switch x {
	case DangerType_FILE:
		return "file", true
	case DangerType_URL:
		return "url", true
	case DangerType_CONTENT:
		return "content", true
	case DangerType_UNCOMMON:
		return "uncommon", true
	case DangerType_HOST:
		return "host", true
	case DangerType_UNWANTED:
		return "unwanted", true
	case DangerType_SAFE:
		return "safe", true
	case DangerType_ACCEPTED:
		return "accepted", true
	case DangerType_ALLOWLISTED_BY_POLICY:
		return "allowlistedByPolicy", true
	case DangerType_ASYNC_SCANNING:
		return "asyncScanning", true
	case DangerType_PASSWORD_PROTECTED:
		return "passwordProtected", true
	case DangerType_BLOCKED_TOO_LARGE:
		return "blockedTooLarge", true
	case DangerType_SENSITIVE_CONTENT_WARNING:
		return "sensitiveContentWarning", true
	case DangerType_SENSITIVE_CONTENT_BLOCK:
		return "sensitiveContentBlock", true
	case DangerType_UNSUPPORTED_FILE_TYPE:
		return "unsupportedFileType", true
	case DangerType_DEEP_SCANNED_FAILED:
		return "deepScannedFailed", true
	case DangerType_DEEP_SCANNED_SAFE:
		return "deepScannedSafe", true
	case DangerType_DEEP_SCANNED_OPENED_DANGEROUS:
		return "deepScannedOpenedDangerous", true
	case DangerType_PROMPT_FOR_SCANING:
		return "promptForScaning", true
	case DangerType_ACCOUNT_COMPROMISE:
		return "accountCompromise", true
	default:
		return "", false
	}
}

type DoubleDelta struct {
	// Previous is "DoubleDelta.previous"
	//
	// Optional
	//
	// NOTE: FFI_USE_Previous MUST be set to true to make this field effective.
	Previous float64
	// Current is "DoubleDelta.current"
	//
	// Optional
	//
	// NOTE: FFI_USE_Current MUST be set to true to make this field effective.
	Current float64

	FFI_USE_Previous bool // for Previous.
	FFI_USE_Current  bool // for Current.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DoubleDelta with all fields set.
func (p DoubleDelta) FromRef(ref js.Ref) DoubleDelta {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DoubleDelta in the application heap.
func (p DoubleDelta) New() js.Ref {
	return bindings.DoubleDeltaJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DoubleDelta) UpdateFrom(ref js.Ref) {
	bindings.DoubleDeltaJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DoubleDelta) Update(ref js.Ref) {
	bindings.DoubleDeltaJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DoubleDelta) FreeMembers(recursive bool) {
}

type DownloadCallbackFunc func(this js.Ref, downloadId int32) js.Ref

func (fn DownloadCallbackFunc) Register() js.Func[func(downloadId int32)] {
	return js.RegisterCallback[func(downloadId int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DownloadCallbackFunc) DispatchCallback(
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

type DownloadCallback[T any] struct {
	Fn  func(arg T, this js.Ref, downloadId int32) js.Ref
	Arg T
}

func (cb *DownloadCallback[T]) Register() js.Func[func(downloadId int32)] {
	return js.RegisterCallback[func(downloadId int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DownloadCallback[T]) DispatchCallback(
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

type StringDelta struct {
	// Previous is "StringDelta.previous"
	//
	// Optional
	Previous js.String
	// Current is "StringDelta.current"
	//
	// Optional
	Current js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StringDelta with all fields set.
func (p StringDelta) FromRef(ref js.Ref) StringDelta {
	p.UpdateFrom(ref)
	return p
}

// New creates a new StringDelta in the application heap.
func (p StringDelta) New() js.Ref {
	return bindings.StringDeltaJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *StringDelta) UpdateFrom(ref js.Ref) {
	bindings.StringDeltaJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StringDelta) Update(ref js.Ref) {
	bindings.StringDeltaJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StringDelta) FreeMembers(recursive bool) {
	js.Free(
		p.Previous.Ref(),
		p.Current.Ref(),
	)
	p.Previous = p.Previous.FromRef(js.Undefined)
	p.Current = p.Current.FromRef(js.Undefined)
}

type DownloadDelta struct {
	// Id is "DownloadDelta.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id int32
	// Url is "DownloadDelta.url"
	//
	// Optional
	//
	// NOTE: Url.FFI_USE MUST be set to true to get Url used.
	Url StringDelta
	// FinalUrl is "DownloadDelta.finalUrl"
	//
	// Optional
	//
	// NOTE: FinalUrl.FFI_USE MUST be set to true to get FinalUrl used.
	FinalUrl StringDelta
	// Filename is "DownloadDelta.filename"
	//
	// Optional
	//
	// NOTE: Filename.FFI_USE MUST be set to true to get Filename used.
	Filename StringDelta
	// Danger is "DownloadDelta.danger"
	//
	// Optional
	//
	// NOTE: Danger.FFI_USE MUST be set to true to get Danger used.
	Danger StringDelta
	// Mime is "DownloadDelta.mime"
	//
	// Optional
	//
	// NOTE: Mime.FFI_USE MUST be set to true to get Mime used.
	Mime StringDelta
	// StartTime is "DownloadDelta.startTime"
	//
	// Optional
	//
	// NOTE: StartTime.FFI_USE MUST be set to true to get StartTime used.
	StartTime StringDelta
	// EndTime is "DownloadDelta.endTime"
	//
	// Optional
	//
	// NOTE: EndTime.FFI_USE MUST be set to true to get EndTime used.
	EndTime StringDelta
	// State is "DownloadDelta.state"
	//
	// Optional
	//
	// NOTE: State.FFI_USE MUST be set to true to get State used.
	State StringDelta
	// CanResume is "DownloadDelta.canResume"
	//
	// Optional
	//
	// NOTE: CanResume.FFI_USE MUST be set to true to get CanResume used.
	CanResume BooleanDelta
	// Paused is "DownloadDelta.paused"
	//
	// Optional
	//
	// NOTE: Paused.FFI_USE MUST be set to true to get Paused used.
	Paused BooleanDelta
	// Error is "DownloadDelta.error"
	//
	// Optional
	//
	// NOTE: Error.FFI_USE MUST be set to true to get Error used.
	Error StringDelta
	// TotalBytes is "DownloadDelta.totalBytes"
	//
	// Optional
	//
	// NOTE: TotalBytes.FFI_USE MUST be set to true to get TotalBytes used.
	TotalBytes DoubleDelta
	// FileSize is "DownloadDelta.fileSize"
	//
	// Optional
	//
	// NOTE: FileSize.FFI_USE MUST be set to true to get FileSize used.
	FileSize DoubleDelta
	// Exists is "DownloadDelta.exists"
	//
	// Optional
	//
	// NOTE: Exists.FFI_USE MUST be set to true to get Exists used.
	Exists BooleanDelta

	FFI_USE_Id bool // for Id.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DownloadDelta with all fields set.
func (p DownloadDelta) FromRef(ref js.Ref) DownloadDelta {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DownloadDelta in the application heap.
func (p DownloadDelta) New() js.Ref {
	return bindings.DownloadDeltaJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DownloadDelta) UpdateFrom(ref js.Ref) {
	bindings.DownloadDeltaJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DownloadDelta) Update(ref js.Ref) {
	bindings.DownloadDeltaJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DownloadDelta) FreeMembers(recursive bool) {
	if recursive {
		p.Url.FreeMembers(true)
		p.FinalUrl.FreeMembers(true)
		p.Filename.FreeMembers(true)
		p.Danger.FreeMembers(true)
		p.Mime.FreeMembers(true)
		p.StartTime.FreeMembers(true)
		p.EndTime.FreeMembers(true)
		p.State.FreeMembers(true)
		p.CanResume.FreeMembers(true)
		p.Paused.FreeMembers(true)
		p.Error.FreeMembers(true)
		p.TotalBytes.FreeMembers(true)
		p.FileSize.FreeMembers(true)
		p.Exists.FreeMembers(true)
	}
}

type State uint32

const (
	_ State = iota

	State_IN_PROGRESS
	State_INTERRUPTED
	State_COMPLETE
)

func (State) FromRef(str js.Ref) State {
	return State(bindings.ConstOfState(str))
}

func (x State) String() (string, bool) {
	switch x {
	case State_IN_PROGRESS:
		return "in_progress", true
	case State_INTERRUPTED:
		return "interrupted", true
	case State_COMPLETE:
		return "complete", true
	default:
		return "", false
	}
}

type InterruptReason uint32

const (
	_ InterruptReason = iota

	InterruptReason_FILE_FAILED
	InterruptReason_FILE_ACCESS_DENIED
	InterruptReason_FILE_NO_SPACE
	InterruptReason_FILE_NAME_TOO_LONG
	InterruptReason_FILE_TOO_LARGE
	InterruptReason_FILE_VIRUS_INFECTED
	InterruptReason_FILE_TRANSIENT_ERROR
	InterruptReason_FILE_BLOCKED
	InterruptReason_FILE_SECURITY_CHECK_FAILED
	InterruptReason_FILE_TOO_SHORT
	InterruptReason_FILE_HASH_MISMATCH
	InterruptReason_FILE_SAME_AS_SOURCE
	InterruptReason_NETWORK_FAILED
	InterruptReason_NETWORK_TIMEOUT
	InterruptReason_NETWORK_DISCONNECTED
	InterruptReason_NETWORK_SERVER_DOWN
	InterruptReason_NETWORK_INVALID_REQUEST
	InterruptReason_SERVER_FAILED
	InterruptReason_SERVER_NO_RANGE
	InterruptReason_SERVER_BAD_CONTENT
	InterruptReason_SERVER_UNAUTHORIZED
	InterruptReason_SERVER_CERT_PROBLEM
	InterruptReason_SERVER_FORBIDDEN
	InterruptReason_SERVER_UNREACHABLE
	InterruptReason_SERVER_CONTENT_LENGTH_MISMATCH
	InterruptReason_SERVER_CROSS_ORIGIN_REDIRECT
	InterruptReason_USER_CANCELED
	InterruptReason_USER_SHUTDOWN
	InterruptReason_CRASH
)

func (InterruptReason) FromRef(str js.Ref) InterruptReason {
	return InterruptReason(bindings.ConstOfInterruptReason(str))
}

func (x InterruptReason) String() (string, bool) {
	switch x {
	case InterruptReason_FILE_FAILED:
		return "FILE_FAILED", true
	case InterruptReason_FILE_ACCESS_DENIED:
		return "FILE_ACCESS_DENIED", true
	case InterruptReason_FILE_NO_SPACE:
		return "FILE_NO_SPACE", true
	case InterruptReason_FILE_NAME_TOO_LONG:
		return "FILE_NAME_TOO_LONG", true
	case InterruptReason_FILE_TOO_LARGE:
		return "FILE_TOO_LARGE", true
	case InterruptReason_FILE_VIRUS_INFECTED:
		return "FILE_VIRUS_INFECTED", true
	case InterruptReason_FILE_TRANSIENT_ERROR:
		return "FILE_TRANSIENT_ERROR", true
	case InterruptReason_FILE_BLOCKED:
		return "FILE_BLOCKED", true
	case InterruptReason_FILE_SECURITY_CHECK_FAILED:
		return "FILE_SECURITY_CHECK_FAILED", true
	case InterruptReason_FILE_TOO_SHORT:
		return "FILE_TOO_SHORT", true
	case InterruptReason_FILE_HASH_MISMATCH:
		return "FILE_HASH_MISMATCH", true
	case InterruptReason_FILE_SAME_AS_SOURCE:
		return "FILE_SAME_AS_SOURCE", true
	case InterruptReason_NETWORK_FAILED:
		return "NETWORK_FAILED", true
	case InterruptReason_NETWORK_TIMEOUT:
		return "NETWORK_TIMEOUT", true
	case InterruptReason_NETWORK_DISCONNECTED:
		return "NETWORK_DISCONNECTED", true
	case InterruptReason_NETWORK_SERVER_DOWN:
		return "NETWORK_SERVER_DOWN", true
	case InterruptReason_NETWORK_INVALID_REQUEST:
		return "NETWORK_INVALID_REQUEST", true
	case InterruptReason_SERVER_FAILED:
		return "SERVER_FAILED", true
	case InterruptReason_SERVER_NO_RANGE:
		return "SERVER_NO_RANGE", true
	case InterruptReason_SERVER_BAD_CONTENT:
		return "SERVER_BAD_CONTENT", true
	case InterruptReason_SERVER_UNAUTHORIZED:
		return "SERVER_UNAUTHORIZED", true
	case InterruptReason_SERVER_CERT_PROBLEM:
		return "SERVER_CERT_PROBLEM", true
	case InterruptReason_SERVER_FORBIDDEN:
		return "SERVER_FORBIDDEN", true
	case InterruptReason_SERVER_UNREACHABLE:
		return "SERVER_UNREACHABLE", true
	case InterruptReason_SERVER_CONTENT_LENGTH_MISMATCH:
		return "SERVER_CONTENT_LENGTH_MISMATCH", true
	case InterruptReason_SERVER_CROSS_ORIGIN_REDIRECT:
		return "SERVER_CROSS_ORIGIN_REDIRECT", true
	case InterruptReason_USER_CANCELED:
		return "USER_CANCELED", true
	case InterruptReason_USER_SHUTDOWN:
		return "USER_SHUTDOWN", true
	case InterruptReason_CRASH:
		return "CRASH", true
	default:
		return "", false
	}
}

type DownloadItem struct {
	// Id is "DownloadItem.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id int32
	// Url is "DownloadItem.url"
	//
	// Optional
	Url js.String
	// FinalUrl is "DownloadItem.finalUrl"
	//
	// Optional
	FinalUrl js.String
	// Referrer is "DownloadItem.referrer"
	//
	// Optional
	Referrer js.String
	// Filename is "DownloadItem.filename"
	//
	// Optional
	Filename js.String
	// Incognito is "DownloadItem.incognito"
	//
	// Optional
	//
	// NOTE: FFI_USE_Incognito MUST be set to true to make this field effective.
	Incognito bool
	// Danger is "DownloadItem.danger"
	//
	// Optional
	Danger DangerType
	// Mime is "DownloadItem.mime"
	//
	// Optional
	Mime js.String
	// StartTime is "DownloadItem.startTime"
	//
	// Optional
	StartTime js.String
	// EndTime is "DownloadItem.endTime"
	//
	// Optional
	EndTime js.String
	// EstimatedEndTime is "DownloadItem.estimatedEndTime"
	//
	// Optional
	EstimatedEndTime js.String
	// State is "DownloadItem.state"
	//
	// Optional
	State State
	// Paused is "DownloadItem.paused"
	//
	// Optional
	//
	// NOTE: FFI_USE_Paused MUST be set to true to make this field effective.
	Paused bool
	// CanResume is "DownloadItem.canResume"
	//
	// Optional
	//
	// NOTE: FFI_USE_CanResume MUST be set to true to make this field effective.
	CanResume bool
	// Error is "DownloadItem.error"
	//
	// Optional
	Error InterruptReason
	// BytesReceived is "DownloadItem.bytesReceived"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesReceived MUST be set to true to make this field effective.
	BytesReceived float64
	// TotalBytes is "DownloadItem.totalBytes"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalBytes MUST be set to true to make this field effective.
	TotalBytes float64
	// FileSize is "DownloadItem.fileSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_FileSize MUST be set to true to make this field effective.
	FileSize float64
	// Exists is "DownloadItem.exists"
	//
	// Optional
	//
	// NOTE: FFI_USE_Exists MUST be set to true to make this field effective.
	Exists bool
	// ByExtensionId is "DownloadItem.byExtensionId"
	//
	// Optional
	ByExtensionId js.String
	// ByExtensionName is "DownloadItem.byExtensionName"
	//
	// Optional
	ByExtensionName js.String

	FFI_USE_Id            bool // for Id.
	FFI_USE_Incognito     bool // for Incognito.
	FFI_USE_Paused        bool // for Paused.
	FFI_USE_CanResume     bool // for CanResume.
	FFI_USE_BytesReceived bool // for BytesReceived.
	FFI_USE_TotalBytes    bool // for TotalBytes.
	FFI_USE_FileSize      bool // for FileSize.
	FFI_USE_Exists        bool // for Exists.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DownloadItem with all fields set.
func (p DownloadItem) FromRef(ref js.Ref) DownloadItem {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DownloadItem in the application heap.
func (p DownloadItem) New() js.Ref {
	return bindings.DownloadItemJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DownloadItem) UpdateFrom(ref js.Ref) {
	bindings.DownloadItemJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DownloadItem) Update(ref js.Ref) {
	bindings.DownloadItemJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DownloadItem) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
		p.FinalUrl.Ref(),
		p.Referrer.Ref(),
		p.Filename.Ref(),
		p.Mime.Ref(),
		p.StartTime.Ref(),
		p.EndTime.Ref(),
		p.EstimatedEndTime.Ref(),
		p.ByExtensionId.Ref(),
		p.ByExtensionName.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
	p.FinalUrl = p.FinalUrl.FromRef(js.Undefined)
	p.Referrer = p.Referrer.FromRef(js.Undefined)
	p.Filename = p.Filename.FromRef(js.Undefined)
	p.Mime = p.Mime.FromRef(js.Undefined)
	p.StartTime = p.StartTime.FromRef(js.Undefined)
	p.EndTime = p.EndTime.FromRef(js.Undefined)
	p.EstimatedEndTime = p.EstimatedEndTime.FromRef(js.Undefined)
	p.ByExtensionId = p.ByExtensionId.FromRef(js.Undefined)
	p.ByExtensionName = p.ByExtensionName.FromRef(js.Undefined)
}

type FilenameConflictAction uint32

const (
	_ FilenameConflictAction = iota

	FilenameConflictAction_UNIQUIFY
	FilenameConflictAction_OVERWRITE
	FilenameConflictAction_PROMPT
)

func (FilenameConflictAction) FromRef(str js.Ref) FilenameConflictAction {
	return FilenameConflictAction(bindings.ConstOfFilenameConflictAction(str))
}

func (x FilenameConflictAction) String() (string, bool) {
	switch x {
	case FilenameConflictAction_UNIQUIFY:
		return "uniquify", true
	case FilenameConflictAction_OVERWRITE:
		return "overwrite", true
	case FilenameConflictAction_PROMPT:
		return "prompt", true
	default:
		return "", false
	}
}

type HttpMethod uint32

const (
	_ HttpMethod = iota

	HttpMethod_GET
	HttpMethod_POST
)

func (HttpMethod) FromRef(str js.Ref) HttpMethod {
	return HttpMethod(bindings.ConstOfHttpMethod(str))
}

func (x HttpMethod) String() (string, bool) {
	switch x {
	case HttpMethod_GET:
		return "GET", true
	case HttpMethod_POST:
		return "POST", true
	default:
		return "", false
	}
}

type HeaderNameValuePair struct {
	// Name is "HeaderNameValuePair.name"
	//
	// Optional
	Name js.String
	// Value is "HeaderNameValuePair.value"
	//
	// Optional
	Value js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HeaderNameValuePair with all fields set.
func (p HeaderNameValuePair) FromRef(ref js.Ref) HeaderNameValuePair {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HeaderNameValuePair in the application heap.
func (p HeaderNameValuePair) New() js.Ref {
	return bindings.HeaderNameValuePairJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HeaderNameValuePair) UpdateFrom(ref js.Ref) {
	bindings.HeaderNameValuePairJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HeaderNameValuePair) Update(ref js.Ref) {
	bindings.HeaderNameValuePairJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HeaderNameValuePair) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.Value.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Value = p.Value.FromRef(js.Undefined)
}

type DownloadOptions struct {
	// Url is "DownloadOptions.url"
	//
	// Optional
	Url js.String
	// Filename is "DownloadOptions.filename"
	//
	// Optional
	Filename js.String
	// ConflictAction is "DownloadOptions.conflictAction"
	//
	// Optional
	ConflictAction FilenameConflictAction
	// SaveAs is "DownloadOptions.saveAs"
	//
	// Optional
	//
	// NOTE: FFI_USE_SaveAs MUST be set to true to make this field effective.
	SaveAs bool
	// Method is "DownloadOptions.method"
	//
	// Optional
	Method HttpMethod
	// Headers is "DownloadOptions.headers"
	//
	// Optional
	Headers js.Array[HeaderNameValuePair]
	// Body is "DownloadOptions.body"
	//
	// Optional
	Body js.String

	FFI_USE_SaveAs bool // for SaveAs.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DownloadOptions with all fields set.
func (p DownloadOptions) FromRef(ref js.Ref) DownloadOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DownloadOptions in the application heap.
func (p DownloadOptions) New() js.Ref {
	return bindings.DownloadOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DownloadOptions) UpdateFrom(ref js.Ref) {
	bindings.DownloadOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DownloadOptions) Update(ref js.Ref) {
	bindings.DownloadOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DownloadOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
		p.Filename.Ref(),
		p.Headers.Ref(),
		p.Body.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
	p.Filename = p.Filename.FromRef(js.Undefined)
	p.Headers = p.Headers.FromRef(js.Undefined)
	p.Body = p.Body.FromRef(js.Undefined)
}

type DownloadQuery struct {
	// Query is "DownloadQuery.query"
	//
	// Optional
	Query js.Array[js.String]
	// StartedBefore is "DownloadQuery.startedBefore"
	//
	// Optional
	StartedBefore js.String
	// StartedAfter is "DownloadQuery.startedAfter"
	//
	// Optional
	StartedAfter js.String
	// EndedBefore is "DownloadQuery.endedBefore"
	//
	// Optional
	EndedBefore js.String
	// EndedAfter is "DownloadQuery.endedAfter"
	//
	// Optional
	EndedAfter js.String
	// TotalBytesGreater is "DownloadQuery.totalBytesGreater"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalBytesGreater MUST be set to true to make this field effective.
	TotalBytesGreater float64
	// TotalBytesLess is "DownloadQuery.totalBytesLess"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalBytesLess MUST be set to true to make this field effective.
	TotalBytesLess float64
	// FilenameRegex is "DownloadQuery.filenameRegex"
	//
	// Optional
	FilenameRegex js.String
	// UrlRegex is "DownloadQuery.urlRegex"
	//
	// Optional
	UrlRegex js.String
	// FinalUrlRegex is "DownloadQuery.finalUrlRegex"
	//
	// Optional
	FinalUrlRegex js.String
	// Limit is "DownloadQuery.limit"
	//
	// Optional
	//
	// NOTE: FFI_USE_Limit MUST be set to true to make this field effective.
	Limit int32
	// OrderBy is "DownloadQuery.orderBy"
	//
	// Optional
	OrderBy js.Array[js.String]
	// Id is "DownloadQuery.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id int32
	// Url is "DownloadQuery.url"
	//
	// Optional
	Url js.String
	// FinalUrl is "DownloadQuery.finalUrl"
	//
	// Optional
	FinalUrl js.String
	// Filename is "DownloadQuery.filename"
	//
	// Optional
	Filename js.String
	// Danger is "DownloadQuery.danger"
	//
	// Optional
	Danger DangerType
	// Mime is "DownloadQuery.mime"
	//
	// Optional
	Mime js.String
	// StartTime is "DownloadQuery.startTime"
	//
	// Optional
	StartTime js.String
	// EndTime is "DownloadQuery.endTime"
	//
	// Optional
	EndTime js.String
	// State is "DownloadQuery.state"
	//
	// Optional
	State State
	// Paused is "DownloadQuery.paused"
	//
	// Optional
	//
	// NOTE: FFI_USE_Paused MUST be set to true to make this field effective.
	Paused bool
	// Error is "DownloadQuery.error"
	//
	// Optional
	Error InterruptReason
	// BytesReceived is "DownloadQuery.bytesReceived"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesReceived MUST be set to true to make this field effective.
	BytesReceived float64
	// TotalBytes is "DownloadQuery.totalBytes"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalBytes MUST be set to true to make this field effective.
	TotalBytes float64
	// FileSize is "DownloadQuery.fileSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_FileSize MUST be set to true to make this field effective.
	FileSize float64
	// Exists is "DownloadQuery.exists"
	//
	// Optional
	//
	// NOTE: FFI_USE_Exists MUST be set to true to make this field effective.
	Exists bool

	FFI_USE_TotalBytesGreater bool // for TotalBytesGreater.
	FFI_USE_TotalBytesLess    bool // for TotalBytesLess.
	FFI_USE_Limit             bool // for Limit.
	FFI_USE_Id                bool // for Id.
	FFI_USE_Paused            bool // for Paused.
	FFI_USE_BytesReceived     bool // for BytesReceived.
	FFI_USE_TotalBytes        bool // for TotalBytes.
	FFI_USE_FileSize          bool // for FileSize.
	FFI_USE_Exists            bool // for Exists.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DownloadQuery with all fields set.
func (p DownloadQuery) FromRef(ref js.Ref) DownloadQuery {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DownloadQuery in the application heap.
func (p DownloadQuery) New() js.Ref {
	return bindings.DownloadQueryJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DownloadQuery) UpdateFrom(ref js.Ref) {
	bindings.DownloadQueryJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DownloadQuery) Update(ref js.Ref) {
	bindings.DownloadQueryJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DownloadQuery) FreeMembers(recursive bool) {
	js.Free(
		p.Query.Ref(),
		p.StartedBefore.Ref(),
		p.StartedAfter.Ref(),
		p.EndedBefore.Ref(),
		p.EndedAfter.Ref(),
		p.FilenameRegex.Ref(),
		p.UrlRegex.Ref(),
		p.FinalUrlRegex.Ref(),
		p.OrderBy.Ref(),
		p.Url.Ref(),
		p.FinalUrl.Ref(),
		p.Filename.Ref(),
		p.Mime.Ref(),
		p.StartTime.Ref(),
		p.EndTime.Ref(),
	)
	p.Query = p.Query.FromRef(js.Undefined)
	p.StartedBefore = p.StartedBefore.FromRef(js.Undefined)
	p.StartedAfter = p.StartedAfter.FromRef(js.Undefined)
	p.EndedBefore = p.EndedBefore.FromRef(js.Undefined)
	p.EndedAfter = p.EndedAfter.FromRef(js.Undefined)
	p.FilenameRegex = p.FilenameRegex.FromRef(js.Undefined)
	p.UrlRegex = p.UrlRegex.FromRef(js.Undefined)
	p.FinalUrlRegex = p.FinalUrlRegex.FromRef(js.Undefined)
	p.OrderBy = p.OrderBy.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
	p.FinalUrl = p.FinalUrl.FromRef(js.Undefined)
	p.Filename = p.Filename.FromRef(js.Undefined)
	p.Mime = p.Mime.FromRef(js.Undefined)
	p.StartTime = p.StartTime.FromRef(js.Undefined)
	p.EndTime = p.EndTime.FromRef(js.Undefined)
}

type EraseCallbackFunc func(this js.Ref, erasedIds js.Array[int32]) js.Ref

func (fn EraseCallbackFunc) Register() js.Func[func(erasedIds js.Array[int32])] {
	return js.RegisterCallback[func(erasedIds js.Array[int32])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn EraseCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[int32]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type EraseCallback[T any] struct {
	Fn  func(arg T, this js.Ref, erasedIds js.Array[int32]) js.Ref
	Arg T
}

func (cb *EraseCallback[T]) Register() js.Func[func(erasedIds js.Array[int32])] {
	return js.RegisterCallback[func(erasedIds js.Array[int32])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *EraseCallback[T]) DispatchCallback(
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

		js.Array[int32]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type FilenameSuggestion struct {
	// Filename is "FilenameSuggestion.filename"
	//
	// Optional
	Filename js.String
	// ConflictAction is "FilenameSuggestion.conflictAction"
	//
	// Optional
	ConflictAction FilenameConflictAction

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FilenameSuggestion with all fields set.
func (p FilenameSuggestion) FromRef(ref js.Ref) FilenameSuggestion {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FilenameSuggestion in the application heap.
func (p FilenameSuggestion) New() js.Ref {
	return bindings.FilenameSuggestionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FilenameSuggestion) UpdateFrom(ref js.Ref) {
	bindings.FilenameSuggestionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FilenameSuggestion) Update(ref js.Ref) {
	bindings.FilenameSuggestionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FilenameSuggestion) FreeMembers(recursive bool) {
	js.Free(
		p.Filename.Ref(),
	)
	p.Filename = p.Filename.FromRef(js.Undefined)
}

type GetFileIconCallbackFunc func(this js.Ref, iconURL js.String) js.Ref

func (fn GetFileIconCallbackFunc) Register() js.Func[func(iconURL js.String)] {
	return js.RegisterCallback[func(iconURL js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetFileIconCallbackFunc) DispatchCallback(
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

type GetFileIconCallback[T any] struct {
	Fn  func(arg T, this js.Ref, iconURL js.String) js.Ref
	Arg T
}

func (cb *GetFileIconCallback[T]) Register() js.Func[func(iconURL js.String)] {
	return js.RegisterCallback[func(iconURL js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetFileIconCallback[T]) DispatchCallback(
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

type GetFileIconOptions struct {
	// Size is "GetFileIconOptions.size"
	//
	// Optional
	//
	// NOTE: FFI_USE_Size MUST be set to true to make this field effective.
	Size int32

	FFI_USE_Size bool // for Size.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetFileIconOptions with all fields set.
func (p GetFileIconOptions) FromRef(ref js.Ref) GetFileIconOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetFileIconOptions in the application heap.
func (p GetFileIconOptions) New() js.Ref {
	return bindings.GetFileIconOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetFileIconOptions) UpdateFrom(ref js.Ref) {
	bindings.GetFileIconOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetFileIconOptions) Update(ref js.Ref) {
	bindings.GetFileIconOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetFileIconOptions) FreeMembers(recursive bool) {
}

type NullCallbackFunc func(this js.Ref) js.Ref

func (fn NullCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn NullCallbackFunc) DispatchCallback(
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

type NullCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *NullCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *NullCallback[T]) DispatchCallback(
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

type SearchCallbackFunc func(this js.Ref, results js.Array[DownloadItem]) js.Ref

func (fn SearchCallbackFunc) Register() js.Func[func(results js.Array[DownloadItem])] {
	return js.RegisterCallback[func(results js.Array[DownloadItem])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SearchCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[DownloadItem]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SearchCallback[T any] struct {
	Fn  func(arg T, this js.Ref, results js.Array[DownloadItem]) js.Ref
	Arg T
}

func (cb *SearchCallback[T]) Register() js.Func[func(results js.Array[DownloadItem])] {
	return js.RegisterCallback[func(results js.Array[DownloadItem])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SearchCallback[T]) DispatchCallback(
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

		js.Array[DownloadItem]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SuggestFilenameCallbackFunc func(this js.Ref, suggestion *FilenameSuggestion) js.Ref

func (fn SuggestFilenameCallbackFunc) Register() js.Func[func(suggestion *FilenameSuggestion)] {
	return js.RegisterCallback[func(suggestion *FilenameSuggestion)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SuggestFilenameCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 FilenameSuggestion
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

type SuggestFilenameCallback[T any] struct {
	Fn  func(arg T, this js.Ref, suggestion *FilenameSuggestion) js.Ref
	Arg T
}

func (cb *SuggestFilenameCallback[T]) Register() js.Func[func(suggestion *FilenameSuggestion)] {
	return js.RegisterCallback[func(suggestion *FilenameSuggestion)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SuggestFilenameCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 FilenameSuggestion
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

type UiOptions struct {
	// Enabled is "UiOptions.enabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_Enabled MUST be set to true to make this field effective.
	Enabled bool

	FFI_USE_Enabled bool // for Enabled.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UiOptions with all fields set.
func (p UiOptions) FromRef(ref js.Ref) UiOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UiOptions in the application heap.
func (p UiOptions) New() js.Ref {
	return bindings.UiOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UiOptions) UpdateFrom(ref js.Ref) {
	bindings.UiOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UiOptions) Update(ref js.Ref) {
	bindings.UiOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UiOptions) FreeMembers(recursive bool) {
}

// HasFuncAcceptDanger returns true if the function "WEBEXT.downloads.acceptDanger" exists.
func HasFuncAcceptDanger() bool {
	return js.True == bindings.HasFuncAcceptDanger()
}

// FuncAcceptDanger returns the function "WEBEXT.downloads.acceptDanger".
func FuncAcceptDanger() (fn js.Func[func(downloadId int32) js.Promise[js.Void]]) {
	bindings.FuncAcceptDanger(
		js.Pointer(&fn),
	)
	return
}

// AcceptDanger calls the function "WEBEXT.downloads.acceptDanger" directly.
func AcceptDanger(downloadId int32) (ret js.Promise[js.Void]) {
	bindings.CallAcceptDanger(
		js.Pointer(&ret),
		int32(downloadId),
	)

	return
}

// TryAcceptDanger calls the function "WEBEXT.downloads.acceptDanger"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAcceptDanger(downloadId int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAcceptDanger(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(downloadId),
	)

	return
}

// HasFuncCancel returns true if the function "WEBEXT.downloads.cancel" exists.
func HasFuncCancel() bool {
	return js.True == bindings.HasFuncCancel()
}

// FuncCancel returns the function "WEBEXT.downloads.cancel".
func FuncCancel() (fn js.Func[func(downloadId int32) js.Promise[js.Void]]) {
	bindings.FuncCancel(
		js.Pointer(&fn),
	)
	return
}

// Cancel calls the function "WEBEXT.downloads.cancel" directly.
func Cancel(downloadId int32) (ret js.Promise[js.Void]) {
	bindings.CallCancel(
		js.Pointer(&ret),
		int32(downloadId),
	)

	return
}

// TryCancel calls the function "WEBEXT.downloads.cancel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCancel(downloadId int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCancel(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(downloadId),
	)

	return
}

// HasFuncDownload returns true if the function "WEBEXT.downloads.download" exists.
func HasFuncDownload() bool {
	return js.True == bindings.HasFuncDownload()
}

// FuncDownload returns the function "WEBEXT.downloads.download".
func FuncDownload() (fn js.Func[func(options DownloadOptions) js.Promise[js.Number[int32]]]) {
	bindings.FuncDownload(
		js.Pointer(&fn),
	)
	return
}

// Download calls the function "WEBEXT.downloads.download" directly.
func Download(options DownloadOptions) (ret js.Promise[js.Number[int32]]) {
	bindings.CallDownload(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryDownload calls the function "WEBEXT.downloads.download"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDownload(options DownloadOptions) (ret js.Promise[js.Number[int32]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDownload(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncErase returns true if the function "WEBEXT.downloads.erase" exists.
func HasFuncErase() bool {
	return js.True == bindings.HasFuncErase()
}

// FuncErase returns the function "WEBEXT.downloads.erase".
func FuncErase() (fn js.Func[func(query DownloadQuery) js.Promise[js.Array[int32]]]) {
	bindings.FuncErase(
		js.Pointer(&fn),
	)
	return
}

// Erase calls the function "WEBEXT.downloads.erase" directly.
func Erase(query DownloadQuery) (ret js.Promise[js.Array[int32]]) {
	bindings.CallErase(
		js.Pointer(&ret),
		js.Pointer(&query),
	)

	return
}

// TryErase calls the function "WEBEXT.downloads.erase"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryErase(query DownloadQuery) (ret js.Promise[js.Array[int32]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryErase(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&query),
	)

	return
}

// HasFuncGetFileIcon returns true if the function "WEBEXT.downloads.getFileIcon" exists.
func HasFuncGetFileIcon() bool {
	return js.True == bindings.HasFuncGetFileIcon()
}

// FuncGetFileIcon returns the function "WEBEXT.downloads.getFileIcon".
func FuncGetFileIcon() (fn js.Func[func(downloadId int32, options GetFileIconOptions) js.Promise[js.String]]) {
	bindings.FuncGetFileIcon(
		js.Pointer(&fn),
	)
	return
}

// GetFileIcon calls the function "WEBEXT.downloads.getFileIcon" directly.
func GetFileIcon(downloadId int32, options GetFileIconOptions) (ret js.Promise[js.String]) {
	bindings.CallGetFileIcon(
		js.Pointer(&ret),
		int32(downloadId),
		js.Pointer(&options),
	)

	return
}

// TryGetFileIcon calls the function "WEBEXT.downloads.getFileIcon"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetFileIcon(downloadId int32, options GetFileIconOptions) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetFileIcon(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(downloadId),
		js.Pointer(&options),
	)

	return
}

type OnChangedEventCallbackFunc func(this js.Ref, downloadDelta *DownloadDelta) js.Ref

func (fn OnChangedEventCallbackFunc) Register() js.Func[func(downloadDelta *DownloadDelta)] {
	return js.RegisterCallback[func(downloadDelta *DownloadDelta)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DownloadDelta
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

type OnChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, downloadDelta *DownloadDelta) js.Ref
	Arg T
}

func (cb *OnChangedEventCallback[T]) Register() js.Func[func(downloadDelta *DownloadDelta)] {
	return js.RegisterCallback[func(downloadDelta *DownloadDelta)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DownloadDelta
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

// HasFuncOnChanged returns true if the function "WEBEXT.downloads.onChanged.addListener" exists.
func HasFuncOnChanged() bool {
	return js.True == bindings.HasFuncOnChanged()
}

// FuncOnChanged returns the function "WEBEXT.downloads.onChanged.addListener".
func FuncOnChanged() (fn js.Func[func(callback js.Func[func(downloadDelta *DownloadDelta)])]) {
	bindings.FuncOnChanged(
		js.Pointer(&fn),
	)
	return
}

// OnChanged calls the function "WEBEXT.downloads.onChanged.addListener" directly.
func OnChanged(callback js.Func[func(downloadDelta *DownloadDelta)]) (ret js.Void) {
	bindings.CallOnChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnChanged calls the function "WEBEXT.downloads.onChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnChanged(callback js.Func[func(downloadDelta *DownloadDelta)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffChanged returns true if the function "WEBEXT.downloads.onChanged.removeListener" exists.
func HasFuncOffChanged() bool {
	return js.True == bindings.HasFuncOffChanged()
}

// FuncOffChanged returns the function "WEBEXT.downloads.onChanged.removeListener".
func FuncOffChanged() (fn js.Func[func(callback js.Func[func(downloadDelta *DownloadDelta)])]) {
	bindings.FuncOffChanged(
		js.Pointer(&fn),
	)
	return
}

// OffChanged calls the function "WEBEXT.downloads.onChanged.removeListener" directly.
func OffChanged(callback js.Func[func(downloadDelta *DownloadDelta)]) (ret js.Void) {
	bindings.CallOffChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffChanged calls the function "WEBEXT.downloads.onChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffChanged(callback js.Func[func(downloadDelta *DownloadDelta)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnChanged returns true if the function "WEBEXT.downloads.onChanged.hasListener" exists.
func HasFuncHasOnChanged() bool {
	return js.True == bindings.HasFuncHasOnChanged()
}

// FuncHasOnChanged returns the function "WEBEXT.downloads.onChanged.hasListener".
func FuncHasOnChanged() (fn js.Func[func(callback js.Func[func(downloadDelta *DownloadDelta)]) bool]) {
	bindings.FuncHasOnChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnChanged calls the function "WEBEXT.downloads.onChanged.hasListener" directly.
func HasOnChanged(callback js.Func[func(downloadDelta *DownloadDelta)]) (ret bool) {
	bindings.CallHasOnChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnChanged calls the function "WEBEXT.downloads.onChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnChanged(callback js.Func[func(downloadDelta *DownloadDelta)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnCreatedEventCallbackFunc func(this js.Ref, downloadItem *DownloadItem) js.Ref

func (fn OnCreatedEventCallbackFunc) Register() js.Func[func(downloadItem *DownloadItem)] {
	return js.RegisterCallback[func(downloadItem *DownloadItem)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCreatedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DownloadItem
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

type OnCreatedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, downloadItem *DownloadItem) js.Ref
	Arg T
}

func (cb *OnCreatedEventCallback[T]) Register() js.Func[func(downloadItem *DownloadItem)] {
	return js.RegisterCallback[func(downloadItem *DownloadItem)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCreatedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DownloadItem
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

// HasFuncOnCreated returns true if the function "WEBEXT.downloads.onCreated.addListener" exists.
func HasFuncOnCreated() bool {
	return js.True == bindings.HasFuncOnCreated()
}

// FuncOnCreated returns the function "WEBEXT.downloads.onCreated.addListener".
func FuncOnCreated() (fn js.Func[func(callback js.Func[func(downloadItem *DownloadItem)])]) {
	bindings.FuncOnCreated(
		js.Pointer(&fn),
	)
	return
}

// OnCreated calls the function "WEBEXT.downloads.onCreated.addListener" directly.
func OnCreated(callback js.Func[func(downloadItem *DownloadItem)]) (ret js.Void) {
	bindings.CallOnCreated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCreated calls the function "WEBEXT.downloads.onCreated.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCreated(callback js.Func[func(downloadItem *DownloadItem)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCreated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCreated returns true if the function "WEBEXT.downloads.onCreated.removeListener" exists.
func HasFuncOffCreated() bool {
	return js.True == bindings.HasFuncOffCreated()
}

// FuncOffCreated returns the function "WEBEXT.downloads.onCreated.removeListener".
func FuncOffCreated() (fn js.Func[func(callback js.Func[func(downloadItem *DownloadItem)])]) {
	bindings.FuncOffCreated(
		js.Pointer(&fn),
	)
	return
}

// OffCreated calls the function "WEBEXT.downloads.onCreated.removeListener" directly.
func OffCreated(callback js.Func[func(downloadItem *DownloadItem)]) (ret js.Void) {
	bindings.CallOffCreated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCreated calls the function "WEBEXT.downloads.onCreated.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCreated(callback js.Func[func(downloadItem *DownloadItem)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCreated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCreated returns true if the function "WEBEXT.downloads.onCreated.hasListener" exists.
func HasFuncHasOnCreated() bool {
	return js.True == bindings.HasFuncHasOnCreated()
}

// FuncHasOnCreated returns the function "WEBEXT.downloads.onCreated.hasListener".
func FuncHasOnCreated() (fn js.Func[func(callback js.Func[func(downloadItem *DownloadItem)]) bool]) {
	bindings.FuncHasOnCreated(
		js.Pointer(&fn),
	)
	return
}

// HasOnCreated calls the function "WEBEXT.downloads.onCreated.hasListener" directly.
func HasOnCreated(callback js.Func[func(downloadItem *DownloadItem)]) (ret bool) {
	bindings.CallHasOnCreated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCreated calls the function "WEBEXT.downloads.onCreated.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCreated(callback js.Func[func(downloadItem *DownloadItem)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCreated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDeterminingFilenameEventCallbackFunc func(this js.Ref, downloadItem *DownloadItem, suggest js.Func[func(suggestion *FilenameSuggestion)]) js.Ref

func (fn OnDeterminingFilenameEventCallbackFunc) Register() js.Func[func(downloadItem *DownloadItem, suggest js.Func[func(suggestion *FilenameSuggestion)])] {
	return js.RegisterCallback[func(downloadItem *DownloadItem, suggest js.Func[func(suggestion *FilenameSuggestion)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDeterminingFilenameEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DownloadItem
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func(suggestion *FilenameSuggestion)]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnDeterminingFilenameEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, downloadItem *DownloadItem, suggest js.Func[func(suggestion *FilenameSuggestion)]) js.Ref
	Arg T
}

func (cb *OnDeterminingFilenameEventCallback[T]) Register() js.Func[func(downloadItem *DownloadItem, suggest js.Func[func(suggestion *FilenameSuggestion)])] {
	return js.RegisterCallback[func(downloadItem *DownloadItem, suggest js.Func[func(suggestion *FilenameSuggestion)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDeterminingFilenameEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DownloadItem
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func(suggestion *FilenameSuggestion)]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnDeterminingFilename returns true if the function "WEBEXT.downloads.onDeterminingFilename.addListener" exists.
func HasFuncOnDeterminingFilename() bool {
	return js.True == bindings.HasFuncOnDeterminingFilename()
}

// FuncOnDeterminingFilename returns the function "WEBEXT.downloads.onDeterminingFilename.addListener".
func FuncOnDeterminingFilename() (fn js.Func[func(callback js.Func[func(downloadItem *DownloadItem, suggest js.Func[func(suggestion *FilenameSuggestion)])])]) {
	bindings.FuncOnDeterminingFilename(
		js.Pointer(&fn),
	)
	return
}

// OnDeterminingFilename calls the function "WEBEXT.downloads.onDeterminingFilename.addListener" directly.
func OnDeterminingFilename(callback js.Func[func(downloadItem *DownloadItem, suggest js.Func[func(suggestion *FilenameSuggestion)])]) (ret js.Void) {
	bindings.CallOnDeterminingFilename(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDeterminingFilename calls the function "WEBEXT.downloads.onDeterminingFilename.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDeterminingFilename(callback js.Func[func(downloadItem *DownloadItem, suggest js.Func[func(suggestion *FilenameSuggestion)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDeterminingFilename(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDeterminingFilename returns true if the function "WEBEXT.downloads.onDeterminingFilename.removeListener" exists.
func HasFuncOffDeterminingFilename() bool {
	return js.True == bindings.HasFuncOffDeterminingFilename()
}

// FuncOffDeterminingFilename returns the function "WEBEXT.downloads.onDeterminingFilename.removeListener".
func FuncOffDeterminingFilename() (fn js.Func[func(callback js.Func[func(downloadItem *DownloadItem, suggest js.Func[func(suggestion *FilenameSuggestion)])])]) {
	bindings.FuncOffDeterminingFilename(
		js.Pointer(&fn),
	)
	return
}

// OffDeterminingFilename calls the function "WEBEXT.downloads.onDeterminingFilename.removeListener" directly.
func OffDeterminingFilename(callback js.Func[func(downloadItem *DownloadItem, suggest js.Func[func(suggestion *FilenameSuggestion)])]) (ret js.Void) {
	bindings.CallOffDeterminingFilename(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDeterminingFilename calls the function "WEBEXT.downloads.onDeterminingFilename.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDeterminingFilename(callback js.Func[func(downloadItem *DownloadItem, suggest js.Func[func(suggestion *FilenameSuggestion)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDeterminingFilename(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDeterminingFilename returns true if the function "WEBEXT.downloads.onDeterminingFilename.hasListener" exists.
func HasFuncHasOnDeterminingFilename() bool {
	return js.True == bindings.HasFuncHasOnDeterminingFilename()
}

// FuncHasOnDeterminingFilename returns the function "WEBEXT.downloads.onDeterminingFilename.hasListener".
func FuncHasOnDeterminingFilename() (fn js.Func[func(callback js.Func[func(downloadItem *DownloadItem, suggest js.Func[func(suggestion *FilenameSuggestion)])]) bool]) {
	bindings.FuncHasOnDeterminingFilename(
		js.Pointer(&fn),
	)
	return
}

// HasOnDeterminingFilename calls the function "WEBEXT.downloads.onDeterminingFilename.hasListener" directly.
func HasOnDeterminingFilename(callback js.Func[func(downloadItem *DownloadItem, suggest js.Func[func(suggestion *FilenameSuggestion)])]) (ret bool) {
	bindings.CallHasOnDeterminingFilename(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDeterminingFilename calls the function "WEBEXT.downloads.onDeterminingFilename.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDeterminingFilename(callback js.Func[func(downloadItem *DownloadItem, suggest js.Func[func(suggestion *FilenameSuggestion)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDeterminingFilename(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnErasedEventCallbackFunc func(this js.Ref, downloadId int32) js.Ref

func (fn OnErasedEventCallbackFunc) Register() js.Func[func(downloadId int32)] {
	return js.RegisterCallback[func(downloadId int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnErasedEventCallbackFunc) DispatchCallback(
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

type OnErasedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, downloadId int32) js.Ref
	Arg T
}

func (cb *OnErasedEventCallback[T]) Register() js.Func[func(downloadId int32)] {
	return js.RegisterCallback[func(downloadId int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnErasedEventCallback[T]) DispatchCallback(
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

// HasFuncOnErased returns true if the function "WEBEXT.downloads.onErased.addListener" exists.
func HasFuncOnErased() bool {
	return js.True == bindings.HasFuncOnErased()
}

// FuncOnErased returns the function "WEBEXT.downloads.onErased.addListener".
func FuncOnErased() (fn js.Func[func(callback js.Func[func(downloadId int32)])]) {
	bindings.FuncOnErased(
		js.Pointer(&fn),
	)
	return
}

// OnErased calls the function "WEBEXT.downloads.onErased.addListener" directly.
func OnErased(callback js.Func[func(downloadId int32)]) (ret js.Void) {
	bindings.CallOnErased(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnErased calls the function "WEBEXT.downloads.onErased.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnErased(callback js.Func[func(downloadId int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnErased(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffErased returns true if the function "WEBEXT.downloads.onErased.removeListener" exists.
func HasFuncOffErased() bool {
	return js.True == bindings.HasFuncOffErased()
}

// FuncOffErased returns the function "WEBEXT.downloads.onErased.removeListener".
func FuncOffErased() (fn js.Func[func(callback js.Func[func(downloadId int32)])]) {
	bindings.FuncOffErased(
		js.Pointer(&fn),
	)
	return
}

// OffErased calls the function "WEBEXT.downloads.onErased.removeListener" directly.
func OffErased(callback js.Func[func(downloadId int32)]) (ret js.Void) {
	bindings.CallOffErased(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffErased calls the function "WEBEXT.downloads.onErased.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffErased(callback js.Func[func(downloadId int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffErased(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnErased returns true if the function "WEBEXT.downloads.onErased.hasListener" exists.
func HasFuncHasOnErased() bool {
	return js.True == bindings.HasFuncHasOnErased()
}

// FuncHasOnErased returns the function "WEBEXT.downloads.onErased.hasListener".
func FuncHasOnErased() (fn js.Func[func(callback js.Func[func(downloadId int32)]) bool]) {
	bindings.FuncHasOnErased(
		js.Pointer(&fn),
	)
	return
}

// HasOnErased calls the function "WEBEXT.downloads.onErased.hasListener" directly.
func HasOnErased(callback js.Func[func(downloadId int32)]) (ret bool) {
	bindings.CallHasOnErased(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnErased calls the function "WEBEXT.downloads.onErased.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnErased(callback js.Func[func(downloadId int32)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnErased(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOpen returns true if the function "WEBEXT.downloads.open" exists.
func HasFuncOpen() bool {
	return js.True == bindings.HasFuncOpen()
}

// FuncOpen returns the function "WEBEXT.downloads.open".
func FuncOpen() (fn js.Func[func(downloadId int32)]) {
	bindings.FuncOpen(
		js.Pointer(&fn),
	)
	return
}

// Open calls the function "WEBEXT.downloads.open" directly.
func Open(downloadId int32) (ret js.Void) {
	bindings.CallOpen(
		js.Pointer(&ret),
		int32(downloadId),
	)

	return
}

// TryOpen calls the function "WEBEXT.downloads.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpen(downloadId int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpen(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(downloadId),
	)

	return
}

// HasFuncPause returns true if the function "WEBEXT.downloads.pause" exists.
func HasFuncPause() bool {
	return js.True == bindings.HasFuncPause()
}

// FuncPause returns the function "WEBEXT.downloads.pause".
func FuncPause() (fn js.Func[func(downloadId int32) js.Promise[js.Void]]) {
	bindings.FuncPause(
		js.Pointer(&fn),
	)
	return
}

// Pause calls the function "WEBEXT.downloads.pause" directly.
func Pause(downloadId int32) (ret js.Promise[js.Void]) {
	bindings.CallPause(
		js.Pointer(&ret),
		int32(downloadId),
	)

	return
}

// TryPause calls the function "WEBEXT.downloads.pause"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryPause(downloadId int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPause(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(downloadId),
	)

	return
}

// HasFuncRemoveFile returns true if the function "WEBEXT.downloads.removeFile" exists.
func HasFuncRemoveFile() bool {
	return js.True == bindings.HasFuncRemoveFile()
}

// FuncRemoveFile returns the function "WEBEXT.downloads.removeFile".
func FuncRemoveFile() (fn js.Func[func(downloadId int32) js.Promise[js.Void]]) {
	bindings.FuncRemoveFile(
		js.Pointer(&fn),
	)
	return
}

// RemoveFile calls the function "WEBEXT.downloads.removeFile" directly.
func RemoveFile(downloadId int32) (ret js.Promise[js.Void]) {
	bindings.CallRemoveFile(
		js.Pointer(&ret),
		int32(downloadId),
	)

	return
}

// TryRemoveFile calls the function "WEBEXT.downloads.removeFile"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveFile(downloadId int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveFile(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(downloadId),
	)

	return
}

// HasFuncResume returns true if the function "WEBEXT.downloads.resume" exists.
func HasFuncResume() bool {
	return js.True == bindings.HasFuncResume()
}

// FuncResume returns the function "WEBEXT.downloads.resume".
func FuncResume() (fn js.Func[func(downloadId int32) js.Promise[js.Void]]) {
	bindings.FuncResume(
		js.Pointer(&fn),
	)
	return
}

// Resume calls the function "WEBEXT.downloads.resume" directly.
func Resume(downloadId int32) (ret js.Promise[js.Void]) {
	bindings.CallResume(
		js.Pointer(&ret),
		int32(downloadId),
	)

	return
}

// TryResume calls the function "WEBEXT.downloads.resume"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryResume(downloadId int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryResume(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(downloadId),
	)

	return
}

// HasFuncSearch returns true if the function "WEBEXT.downloads.search" exists.
func HasFuncSearch() bool {
	return js.True == bindings.HasFuncSearch()
}

// FuncSearch returns the function "WEBEXT.downloads.search".
func FuncSearch() (fn js.Func[func(query DownloadQuery) js.Promise[js.Array[DownloadItem]]]) {
	bindings.FuncSearch(
		js.Pointer(&fn),
	)
	return
}

// Search calls the function "WEBEXT.downloads.search" directly.
func Search(query DownloadQuery) (ret js.Promise[js.Array[DownloadItem]]) {
	bindings.CallSearch(
		js.Pointer(&ret),
		js.Pointer(&query),
	)

	return
}

// TrySearch calls the function "WEBEXT.downloads.search"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySearch(query DownloadQuery) (ret js.Promise[js.Array[DownloadItem]], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySearch(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&query),
	)

	return
}

// HasFuncSetShelfEnabled returns true if the function "WEBEXT.downloads.setShelfEnabled" exists.
func HasFuncSetShelfEnabled() bool {
	return js.True == bindings.HasFuncSetShelfEnabled()
}

// FuncSetShelfEnabled returns the function "WEBEXT.downloads.setShelfEnabled".
func FuncSetShelfEnabled() (fn js.Func[func(enabled bool)]) {
	bindings.FuncSetShelfEnabled(
		js.Pointer(&fn),
	)
	return
}

// SetShelfEnabled calls the function "WEBEXT.downloads.setShelfEnabled" directly.
func SetShelfEnabled(enabled bool) (ret js.Void) {
	bindings.CallSetShelfEnabled(
		js.Pointer(&ret),
		js.Bool(bool(enabled)),
	)

	return
}

// TrySetShelfEnabled calls the function "WEBEXT.downloads.setShelfEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetShelfEnabled(enabled bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetShelfEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(enabled)),
	)

	return
}

// HasFuncSetUiOptions returns true if the function "WEBEXT.downloads.setUiOptions" exists.
func HasFuncSetUiOptions() bool {
	return js.True == bindings.HasFuncSetUiOptions()
}

// FuncSetUiOptions returns the function "WEBEXT.downloads.setUiOptions".
func FuncSetUiOptions() (fn js.Func[func(options UiOptions) js.Promise[js.Void]]) {
	bindings.FuncSetUiOptions(
		js.Pointer(&fn),
	)
	return
}

// SetUiOptions calls the function "WEBEXT.downloads.setUiOptions" directly.
func SetUiOptions(options UiOptions) (ret js.Promise[js.Void]) {
	bindings.CallSetUiOptions(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TrySetUiOptions calls the function "WEBEXT.downloads.setUiOptions"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetUiOptions(options UiOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetUiOptions(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncShow returns true if the function "WEBEXT.downloads.show" exists.
func HasFuncShow() bool {
	return js.True == bindings.HasFuncShow()
}

// FuncShow returns the function "WEBEXT.downloads.show".
func FuncShow() (fn js.Func[func(downloadId int32)]) {
	bindings.FuncShow(
		js.Pointer(&fn),
	)
	return
}

// Show calls the function "WEBEXT.downloads.show" directly.
func Show(downloadId int32) (ret js.Void) {
	bindings.CallShow(
		js.Pointer(&ret),
		int32(downloadId),
	)

	return
}

// TryShow calls the function "WEBEXT.downloads.show"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryShow(downloadId int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryShow(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(downloadId),
	)

	return
}

// HasFuncShowDefaultFolder returns true if the function "WEBEXT.downloads.showDefaultFolder" exists.
func HasFuncShowDefaultFolder() bool {
	return js.True == bindings.HasFuncShowDefaultFolder()
}

// FuncShowDefaultFolder returns the function "WEBEXT.downloads.showDefaultFolder".
func FuncShowDefaultFolder() (fn js.Func[func()]) {
	bindings.FuncShowDefaultFolder(
		js.Pointer(&fn),
	)
	return
}

// ShowDefaultFolder calls the function "WEBEXT.downloads.showDefaultFolder" directly.
func ShowDefaultFolder() (ret js.Void) {
	bindings.CallShowDefaultFolder(
		js.Pointer(&ret),
	)

	return
}

// TryShowDefaultFolder calls the function "WEBEXT.downloads.showDefaultFolder"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryShowDefaultFolder() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryShowDefaultFolder(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}
