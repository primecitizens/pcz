// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package webrtcloggingprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/webrtcloggingprivate/bindings"
)

type GenericDoneCallbackFunc func(this js.Ref) js.Ref

func (fn GenericDoneCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GenericDoneCallbackFunc) DispatchCallback(
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

type GenericDoneCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *GenericDoneCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GenericDoneCallback[T]) DispatchCallback(
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

type GetLogsDirectoryCallbackFunc func(this js.Ref, entry js.Object) js.Ref

func (fn GetLogsDirectoryCallbackFunc) Register() js.Func[func(entry js.Object)] {
	return js.RegisterCallback[func(entry js.Object)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetLogsDirectoryCallbackFunc) DispatchCallback(
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

type GetLogsDirectoryCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entry js.Object) js.Ref
	Arg T
}

func (cb *GetLogsDirectoryCallback[T]) Register() js.Func[func(entry js.Object)] {
	return js.RegisterCallback[func(entry js.Object)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetLogsDirectoryCallback[T]) DispatchCallback(
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

type MetaDataEntry struct {
	// Key is "MetaDataEntry.key"
	//
	// Optional
	Key js.String
	// Value is "MetaDataEntry.value"
	//
	// Optional
	Value js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MetaDataEntry with all fields set.
func (p MetaDataEntry) FromRef(ref js.Ref) MetaDataEntry {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MetaDataEntry in the application heap.
func (p MetaDataEntry) New() js.Ref {
	return bindings.MetaDataEntryJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MetaDataEntry) UpdateFrom(ref js.Ref) {
	bindings.MetaDataEntryJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MetaDataEntry) Update(ref js.Ref) {
	bindings.MetaDataEntryJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MetaDataEntry) FreeMembers(recursive bool) {
	js.Free(
		p.Key.Ref(),
		p.Value.Ref(),
	)
	p.Key = p.Key.FromRef(js.Undefined)
	p.Value = p.Value.FromRef(js.Undefined)
}

type RecordingDoneCallbackFunc func(this js.Ref, info *RecordingInfo) js.Ref

func (fn RecordingDoneCallbackFunc) Register() js.Func[func(info *RecordingInfo)] {
	return js.RegisterCallback[func(info *RecordingInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RecordingDoneCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RecordingInfo
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

type RecordingDoneCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *RecordingInfo) js.Ref
	Arg T
}

func (cb *RecordingDoneCallback[T]) Register() js.Func[func(info *RecordingInfo)] {
	return js.RegisterCallback[func(info *RecordingInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RecordingDoneCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RecordingInfo
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

type RecordingInfo struct {
	// PrefixPath is "RecordingInfo.prefixPath"
	//
	// Optional
	PrefixPath js.String
	// DidStop is "RecordingInfo.didStop"
	//
	// Optional
	//
	// NOTE: FFI_USE_DidStop MUST be set to true to make this field effective.
	DidStop bool
	// DidManualStop is "RecordingInfo.didManualStop"
	//
	// Optional
	//
	// NOTE: FFI_USE_DidManualStop MUST be set to true to make this field effective.
	DidManualStop bool

	FFI_USE_DidStop       bool // for DidStop.
	FFI_USE_DidManualStop bool // for DidManualStop.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RecordingInfo with all fields set.
func (p RecordingInfo) FromRef(ref js.Ref) RecordingInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RecordingInfo in the application heap.
func (p RecordingInfo) New() js.Ref {
	return bindings.RecordingInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RecordingInfo) UpdateFrom(ref js.Ref) {
	bindings.RecordingInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RecordingInfo) Update(ref js.Ref) {
	bindings.RecordingInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RecordingInfo) FreeMembers(recursive bool) {
	js.Free(
		p.PrefixPath.Ref(),
	)
	p.PrefixPath = p.PrefixPath.FromRef(js.Undefined)
}

type RequestInfo struct {
	// TabId is "RequestInfo.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int32
	// GuestProcessId is "RequestInfo.guestProcessId"
	//
	// Optional
	//
	// NOTE: FFI_USE_GuestProcessId MUST be set to true to make this field effective.
	GuestProcessId int32
	// TargetWebview is "RequestInfo.targetWebview"
	//
	// Optional
	//
	// NOTE: FFI_USE_TargetWebview MUST be set to true to make this field effective.
	TargetWebview bool

	FFI_USE_TabId          bool // for TabId.
	FFI_USE_GuestProcessId bool // for GuestProcessId.
	FFI_USE_TargetWebview  bool // for TargetWebview.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RequestInfo with all fields set.
func (p RequestInfo) FromRef(ref js.Ref) RequestInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RequestInfo in the application heap.
func (p RequestInfo) New() js.Ref {
	return bindings.RequestInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RequestInfo) UpdateFrom(ref js.Ref) {
	bindings.RequestInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RequestInfo) Update(ref js.Ref) {
	bindings.RequestInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RequestInfo) FreeMembers(recursive bool) {
}

type StartEventLoggingCallbackFunc func(this js.Ref, result *StartEventLoggingResult) js.Ref

func (fn StartEventLoggingCallbackFunc) Register() js.Func[func(result *StartEventLoggingResult)] {
	return js.RegisterCallback[func(result *StartEventLoggingResult)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn StartEventLoggingCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 StartEventLoggingResult
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

type StartEventLoggingCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *StartEventLoggingResult) js.Ref
	Arg T
}

func (cb *StartEventLoggingCallback[T]) Register() js.Func[func(result *StartEventLoggingResult)] {
	return js.RegisterCallback[func(result *StartEventLoggingResult)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *StartEventLoggingCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 StartEventLoggingResult
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

type StartEventLoggingResult struct {
	// LogId is "StartEventLoggingResult.logId"
	//
	// Optional
	LogId js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StartEventLoggingResult with all fields set.
func (p StartEventLoggingResult) FromRef(ref js.Ref) StartEventLoggingResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new StartEventLoggingResult in the application heap.
func (p StartEventLoggingResult) New() js.Ref {
	return bindings.StartEventLoggingResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *StartEventLoggingResult) UpdateFrom(ref js.Ref) {
	bindings.StartEventLoggingResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StartEventLoggingResult) Update(ref js.Ref) {
	bindings.StartEventLoggingResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StartEventLoggingResult) FreeMembers(recursive bool) {
	js.Free(
		p.LogId.Ref(),
	)
	p.LogId = p.LogId.FromRef(js.Undefined)
}

type UploadDoneCallbackFunc func(this js.Ref, result *UploadResult) js.Ref

func (fn UploadDoneCallbackFunc) Register() js.Func[func(result *UploadResult)] {
	return js.RegisterCallback[func(result *UploadResult)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UploadDoneCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 UploadResult
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

type UploadDoneCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *UploadResult) js.Ref
	Arg T
}

func (cb *UploadDoneCallback[T]) Register() js.Func[func(result *UploadResult)] {
	return js.RegisterCallback[func(result *UploadResult)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UploadDoneCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 UploadResult
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

type UploadResult struct {
	// ReportId is "UploadResult.reportId"
	//
	// Optional
	ReportId js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UploadResult with all fields set.
func (p UploadResult) FromRef(ref js.Ref) UploadResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UploadResult in the application heap.
func (p UploadResult) New() js.Ref {
	return bindings.UploadResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UploadResult) UpdateFrom(ref js.Ref) {
	bindings.UploadResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UploadResult) Update(ref js.Ref) {
	bindings.UploadResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UploadResult) FreeMembers(recursive bool) {
	js.Free(
		p.ReportId.Ref(),
	)
	p.ReportId = p.ReportId.FromRef(js.Undefined)
}

// HasFuncDiscard returns true if the function "WEBEXT.webrtcLoggingPrivate.discard" exists.
func HasFuncDiscard() bool {
	return js.True == bindings.HasFuncDiscard()
}

// FuncDiscard returns the function "WEBEXT.webrtcLoggingPrivate.discard".
func FuncDiscard() (fn js.Func[func(request RequestInfo, securityOrigin js.String, callback js.Func[func()])]) {
	bindings.FuncDiscard(
		js.Pointer(&fn),
	)
	return
}

// Discard calls the function "WEBEXT.webrtcLoggingPrivate.discard" directly.
func Discard(request RequestInfo, securityOrigin js.String, callback js.Func[func()]) (ret js.Void) {
	bindings.CallDiscard(
		js.Pointer(&ret),
		js.Pointer(&request),
		securityOrigin.Ref(),
		callback.Ref(),
	)

	return
}

// TryDiscard calls the function "WEBEXT.webrtcLoggingPrivate.discard"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDiscard(request RequestInfo, securityOrigin js.String, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDiscard(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
		securityOrigin.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncGetLogsDirectory returns true if the function "WEBEXT.webrtcLoggingPrivate.getLogsDirectory" exists.
func HasFuncGetLogsDirectory() bool {
	return js.True == bindings.HasFuncGetLogsDirectory()
}

// FuncGetLogsDirectory returns the function "WEBEXT.webrtcLoggingPrivate.getLogsDirectory".
func FuncGetLogsDirectory() (fn js.Func[func(callback js.Func[func(entry js.Object)])]) {
	bindings.FuncGetLogsDirectory(
		js.Pointer(&fn),
	)
	return
}

// GetLogsDirectory calls the function "WEBEXT.webrtcLoggingPrivate.getLogsDirectory" directly.
func GetLogsDirectory(callback js.Func[func(entry js.Object)]) (ret js.Void) {
	bindings.CallGetLogsDirectory(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryGetLogsDirectory calls the function "WEBEXT.webrtcLoggingPrivate.getLogsDirectory"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetLogsDirectory(callback js.Func[func(entry js.Object)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetLogsDirectory(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSetMetaData returns true if the function "WEBEXT.webrtcLoggingPrivate.setMetaData" exists.
func HasFuncSetMetaData() bool {
	return js.True == bindings.HasFuncSetMetaData()
}

// FuncSetMetaData returns the function "WEBEXT.webrtcLoggingPrivate.setMetaData".
func FuncSetMetaData() (fn js.Func[func(request RequestInfo, securityOrigin js.String, metaData js.Array[MetaDataEntry], callback js.Func[func()])]) {
	bindings.FuncSetMetaData(
		js.Pointer(&fn),
	)
	return
}

// SetMetaData calls the function "WEBEXT.webrtcLoggingPrivate.setMetaData" directly.
func SetMetaData(request RequestInfo, securityOrigin js.String, metaData js.Array[MetaDataEntry], callback js.Func[func()]) (ret js.Void) {
	bindings.CallSetMetaData(
		js.Pointer(&ret),
		js.Pointer(&request),
		securityOrigin.Ref(),
		metaData.Ref(),
		callback.Ref(),
	)

	return
}

// TrySetMetaData calls the function "WEBEXT.webrtcLoggingPrivate.setMetaData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetMetaData(request RequestInfo, securityOrigin js.String, metaData js.Array[MetaDataEntry], callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetMetaData(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
		securityOrigin.Ref(),
		metaData.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncSetUploadOnRenderClose returns true if the function "WEBEXT.webrtcLoggingPrivate.setUploadOnRenderClose" exists.
func HasFuncSetUploadOnRenderClose() bool {
	return js.True == bindings.HasFuncSetUploadOnRenderClose()
}

// FuncSetUploadOnRenderClose returns the function "WEBEXT.webrtcLoggingPrivate.setUploadOnRenderClose".
func FuncSetUploadOnRenderClose() (fn js.Func[func(request RequestInfo, securityOrigin js.String, shouldUpload bool)]) {
	bindings.FuncSetUploadOnRenderClose(
		js.Pointer(&fn),
	)
	return
}

// SetUploadOnRenderClose calls the function "WEBEXT.webrtcLoggingPrivate.setUploadOnRenderClose" directly.
func SetUploadOnRenderClose(request RequestInfo, securityOrigin js.String, shouldUpload bool) (ret js.Void) {
	bindings.CallSetUploadOnRenderClose(
		js.Pointer(&ret),
		js.Pointer(&request),
		securityOrigin.Ref(),
		js.Bool(bool(shouldUpload)),
	)

	return
}

// TrySetUploadOnRenderClose calls the function "WEBEXT.webrtcLoggingPrivate.setUploadOnRenderClose"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetUploadOnRenderClose(request RequestInfo, securityOrigin js.String, shouldUpload bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetUploadOnRenderClose(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
		securityOrigin.Ref(),
		js.Bool(bool(shouldUpload)),
	)

	return
}

// HasFuncStart returns true if the function "WEBEXT.webrtcLoggingPrivate.start" exists.
func HasFuncStart() bool {
	return js.True == bindings.HasFuncStart()
}

// FuncStart returns the function "WEBEXT.webrtcLoggingPrivate.start".
func FuncStart() (fn js.Func[func(request RequestInfo, securityOrigin js.String, callback js.Func[func()])]) {
	bindings.FuncStart(
		js.Pointer(&fn),
	)
	return
}

// Start calls the function "WEBEXT.webrtcLoggingPrivate.start" directly.
func Start(request RequestInfo, securityOrigin js.String, callback js.Func[func()]) (ret js.Void) {
	bindings.CallStart(
		js.Pointer(&ret),
		js.Pointer(&request),
		securityOrigin.Ref(),
		callback.Ref(),
	)

	return
}

// TryStart calls the function "WEBEXT.webrtcLoggingPrivate.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStart(request RequestInfo, securityOrigin js.String, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStart(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
		securityOrigin.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncStartAudioDebugRecordings returns true if the function "WEBEXT.webrtcLoggingPrivate.startAudioDebugRecordings" exists.
func HasFuncStartAudioDebugRecordings() bool {
	return js.True == bindings.HasFuncStartAudioDebugRecordings()
}

// FuncStartAudioDebugRecordings returns the function "WEBEXT.webrtcLoggingPrivate.startAudioDebugRecordings".
func FuncStartAudioDebugRecordings() (fn js.Func[func(request RequestInfo, securityOrigin js.String, seconds int32, callback js.Func[func(info *RecordingInfo)])]) {
	bindings.FuncStartAudioDebugRecordings(
		js.Pointer(&fn),
	)
	return
}

// StartAudioDebugRecordings calls the function "WEBEXT.webrtcLoggingPrivate.startAudioDebugRecordings" directly.
func StartAudioDebugRecordings(request RequestInfo, securityOrigin js.String, seconds int32, callback js.Func[func(info *RecordingInfo)]) (ret js.Void) {
	bindings.CallStartAudioDebugRecordings(
		js.Pointer(&ret),
		js.Pointer(&request),
		securityOrigin.Ref(),
		int32(seconds),
		callback.Ref(),
	)

	return
}

// TryStartAudioDebugRecordings calls the function "WEBEXT.webrtcLoggingPrivate.startAudioDebugRecordings"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStartAudioDebugRecordings(request RequestInfo, securityOrigin js.String, seconds int32, callback js.Func[func(info *RecordingInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStartAudioDebugRecordings(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
		securityOrigin.Ref(),
		int32(seconds),
		callback.Ref(),
	)

	return
}

// HasFuncStartEventLogging returns true if the function "WEBEXT.webrtcLoggingPrivate.startEventLogging" exists.
func HasFuncStartEventLogging() bool {
	return js.True == bindings.HasFuncStartEventLogging()
}

// FuncStartEventLogging returns the function "WEBEXT.webrtcLoggingPrivate.startEventLogging".
func FuncStartEventLogging() (fn js.Func[func(request RequestInfo, securityOrigin js.String, sessionId js.String, maxLogSizeBytes int32, outputPeriodMs int32, webAppId int32, callback js.Func[func(result *StartEventLoggingResult)])]) {
	bindings.FuncStartEventLogging(
		js.Pointer(&fn),
	)
	return
}

// StartEventLogging calls the function "WEBEXT.webrtcLoggingPrivate.startEventLogging" directly.
func StartEventLogging(request RequestInfo, securityOrigin js.String, sessionId js.String, maxLogSizeBytes int32, outputPeriodMs int32, webAppId int32, callback js.Func[func(result *StartEventLoggingResult)]) (ret js.Void) {
	bindings.CallStartEventLogging(
		js.Pointer(&ret),
		js.Pointer(&request),
		securityOrigin.Ref(),
		sessionId.Ref(),
		int32(maxLogSizeBytes),
		int32(outputPeriodMs),
		int32(webAppId),
		callback.Ref(),
	)

	return
}

// TryStartEventLogging calls the function "WEBEXT.webrtcLoggingPrivate.startEventLogging"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStartEventLogging(request RequestInfo, securityOrigin js.String, sessionId js.String, maxLogSizeBytes int32, outputPeriodMs int32, webAppId int32, callback js.Func[func(result *StartEventLoggingResult)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStartEventLogging(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
		securityOrigin.Ref(),
		sessionId.Ref(),
		int32(maxLogSizeBytes),
		int32(outputPeriodMs),
		int32(webAppId),
		callback.Ref(),
	)

	return
}

// HasFuncStartRtpDump returns true if the function "WEBEXT.webrtcLoggingPrivate.startRtpDump" exists.
func HasFuncStartRtpDump() bool {
	return js.True == bindings.HasFuncStartRtpDump()
}

// FuncStartRtpDump returns the function "WEBEXT.webrtcLoggingPrivate.startRtpDump".
func FuncStartRtpDump() (fn js.Func[func(request RequestInfo, securityOrigin js.String, incoming bool, outgoing bool, callback js.Func[func()])]) {
	bindings.FuncStartRtpDump(
		js.Pointer(&fn),
	)
	return
}

// StartRtpDump calls the function "WEBEXT.webrtcLoggingPrivate.startRtpDump" directly.
func StartRtpDump(request RequestInfo, securityOrigin js.String, incoming bool, outgoing bool, callback js.Func[func()]) (ret js.Void) {
	bindings.CallStartRtpDump(
		js.Pointer(&ret),
		js.Pointer(&request),
		securityOrigin.Ref(),
		js.Bool(bool(incoming)),
		js.Bool(bool(outgoing)),
		callback.Ref(),
	)

	return
}

// TryStartRtpDump calls the function "WEBEXT.webrtcLoggingPrivate.startRtpDump"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStartRtpDump(request RequestInfo, securityOrigin js.String, incoming bool, outgoing bool, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStartRtpDump(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
		securityOrigin.Ref(),
		js.Bool(bool(incoming)),
		js.Bool(bool(outgoing)),
		callback.Ref(),
	)

	return
}

// HasFuncStop returns true if the function "WEBEXT.webrtcLoggingPrivate.stop" exists.
func HasFuncStop() bool {
	return js.True == bindings.HasFuncStop()
}

// FuncStop returns the function "WEBEXT.webrtcLoggingPrivate.stop".
func FuncStop() (fn js.Func[func(request RequestInfo, securityOrigin js.String, callback js.Func[func()])]) {
	bindings.FuncStop(
		js.Pointer(&fn),
	)
	return
}

// Stop calls the function "WEBEXT.webrtcLoggingPrivate.stop" directly.
func Stop(request RequestInfo, securityOrigin js.String, callback js.Func[func()]) (ret js.Void) {
	bindings.CallStop(
		js.Pointer(&ret),
		js.Pointer(&request),
		securityOrigin.Ref(),
		callback.Ref(),
	)

	return
}

// TryStop calls the function "WEBEXT.webrtcLoggingPrivate.stop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStop(request RequestInfo, securityOrigin js.String, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStop(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
		securityOrigin.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncStopAudioDebugRecordings returns true if the function "WEBEXT.webrtcLoggingPrivate.stopAudioDebugRecordings" exists.
func HasFuncStopAudioDebugRecordings() bool {
	return js.True == bindings.HasFuncStopAudioDebugRecordings()
}

// FuncStopAudioDebugRecordings returns the function "WEBEXT.webrtcLoggingPrivate.stopAudioDebugRecordings".
func FuncStopAudioDebugRecordings() (fn js.Func[func(request RequestInfo, securityOrigin js.String, callback js.Func[func(info *RecordingInfo)])]) {
	bindings.FuncStopAudioDebugRecordings(
		js.Pointer(&fn),
	)
	return
}

// StopAudioDebugRecordings calls the function "WEBEXT.webrtcLoggingPrivate.stopAudioDebugRecordings" directly.
func StopAudioDebugRecordings(request RequestInfo, securityOrigin js.String, callback js.Func[func(info *RecordingInfo)]) (ret js.Void) {
	bindings.CallStopAudioDebugRecordings(
		js.Pointer(&ret),
		js.Pointer(&request),
		securityOrigin.Ref(),
		callback.Ref(),
	)

	return
}

// TryStopAudioDebugRecordings calls the function "WEBEXT.webrtcLoggingPrivate.stopAudioDebugRecordings"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStopAudioDebugRecordings(request RequestInfo, securityOrigin js.String, callback js.Func[func(info *RecordingInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStopAudioDebugRecordings(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
		securityOrigin.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncStopRtpDump returns true if the function "WEBEXT.webrtcLoggingPrivate.stopRtpDump" exists.
func HasFuncStopRtpDump() bool {
	return js.True == bindings.HasFuncStopRtpDump()
}

// FuncStopRtpDump returns the function "WEBEXT.webrtcLoggingPrivate.stopRtpDump".
func FuncStopRtpDump() (fn js.Func[func(request RequestInfo, securityOrigin js.String, incoming bool, outgoing bool, callback js.Func[func()])]) {
	bindings.FuncStopRtpDump(
		js.Pointer(&fn),
	)
	return
}

// StopRtpDump calls the function "WEBEXT.webrtcLoggingPrivate.stopRtpDump" directly.
func StopRtpDump(request RequestInfo, securityOrigin js.String, incoming bool, outgoing bool, callback js.Func[func()]) (ret js.Void) {
	bindings.CallStopRtpDump(
		js.Pointer(&ret),
		js.Pointer(&request),
		securityOrigin.Ref(),
		js.Bool(bool(incoming)),
		js.Bool(bool(outgoing)),
		callback.Ref(),
	)

	return
}

// TryStopRtpDump calls the function "WEBEXT.webrtcLoggingPrivate.stopRtpDump"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStopRtpDump(request RequestInfo, securityOrigin js.String, incoming bool, outgoing bool, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStopRtpDump(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
		securityOrigin.Ref(),
		js.Bool(bool(incoming)),
		js.Bool(bool(outgoing)),
		callback.Ref(),
	)

	return
}

// HasFuncStore returns true if the function "WEBEXT.webrtcLoggingPrivate.store" exists.
func HasFuncStore() bool {
	return js.True == bindings.HasFuncStore()
}

// FuncStore returns the function "WEBEXT.webrtcLoggingPrivate.store".
func FuncStore() (fn js.Func[func(request RequestInfo, securityOrigin js.String, logId js.String, callback js.Func[func()])]) {
	bindings.FuncStore(
		js.Pointer(&fn),
	)
	return
}

// Store calls the function "WEBEXT.webrtcLoggingPrivate.store" directly.
func Store(request RequestInfo, securityOrigin js.String, logId js.String, callback js.Func[func()]) (ret js.Void) {
	bindings.CallStore(
		js.Pointer(&ret),
		js.Pointer(&request),
		securityOrigin.Ref(),
		logId.Ref(),
		callback.Ref(),
	)

	return
}

// TryStore calls the function "WEBEXT.webrtcLoggingPrivate.store"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStore(request RequestInfo, securityOrigin js.String, logId js.String, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStore(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
		securityOrigin.Ref(),
		logId.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncUpload returns true if the function "WEBEXT.webrtcLoggingPrivate.upload" exists.
func HasFuncUpload() bool {
	return js.True == bindings.HasFuncUpload()
}

// FuncUpload returns the function "WEBEXT.webrtcLoggingPrivate.upload".
func FuncUpload() (fn js.Func[func(request RequestInfo, securityOrigin js.String, callback js.Func[func(result *UploadResult)])]) {
	bindings.FuncUpload(
		js.Pointer(&fn),
	)
	return
}

// Upload calls the function "WEBEXT.webrtcLoggingPrivate.upload" directly.
func Upload(request RequestInfo, securityOrigin js.String, callback js.Func[func(result *UploadResult)]) (ret js.Void) {
	bindings.CallUpload(
		js.Pointer(&ret),
		js.Pointer(&request),
		securityOrigin.Ref(),
		callback.Ref(),
	)

	return
}

// TryUpload calls the function "WEBEXT.webrtcLoggingPrivate.upload"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpload(request RequestInfo, securityOrigin js.String, callback js.Func[func(result *UploadResult)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpload(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
		securityOrigin.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncUploadStored returns true if the function "WEBEXT.webrtcLoggingPrivate.uploadStored" exists.
func HasFuncUploadStored() bool {
	return js.True == bindings.HasFuncUploadStored()
}

// FuncUploadStored returns the function "WEBEXT.webrtcLoggingPrivate.uploadStored".
func FuncUploadStored() (fn js.Func[func(request RequestInfo, securityOrigin js.String, logId js.String, callback js.Func[func(result *UploadResult)])]) {
	bindings.FuncUploadStored(
		js.Pointer(&fn),
	)
	return
}

// UploadStored calls the function "WEBEXT.webrtcLoggingPrivate.uploadStored" directly.
func UploadStored(request RequestInfo, securityOrigin js.String, logId js.String, callback js.Func[func(result *UploadResult)]) (ret js.Void) {
	bindings.CallUploadStored(
		js.Pointer(&ret),
		js.Pointer(&request),
		securityOrigin.Ref(),
		logId.Ref(),
		callback.Ref(),
	)

	return
}

// TryUploadStored calls the function "WEBEXT.webrtcLoggingPrivate.uploadStored"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUploadStored(request RequestInfo, securityOrigin js.String, logId js.String, callback js.Func[func(result *UploadResult)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryUploadStored(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
		securityOrigin.Ref(),
		logId.Ref(),
		callback.Ref(),
	)

	return
}
