// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package feedbackprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/feedbackprivate/bindings"
)

type AttachedFile struct {
	// Name is "AttachedFile.name"
	//
	// Optional
	Name js.String
	// Data is "AttachedFile.data"
	//
	// Optional
	Data js.Object

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AttachedFile with all fields set.
func (p AttachedFile) FromRef(ref js.Ref) AttachedFile {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AttachedFile in the application heap.
func (p AttachedFile) New() js.Ref {
	return bindings.AttachedFileJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AttachedFile) UpdateFrom(ref js.Ref) {
	bindings.AttachedFileJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AttachedFile) Update(ref js.Ref) {
	bindings.AttachedFileJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AttachedFile) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.Data.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Data = p.Data.FromRef(js.Undefined)
}

type FeedbackFlow uint32

const (
	_ FeedbackFlow = iota

	FeedbackFlow_REGULAR
	FeedbackFlow_LOGIN
	FeedbackFlow_SAD_TAB_CRASH
	FeedbackFlow_GOOGLE_INTERNAL
)

func (FeedbackFlow) FromRef(str js.Ref) FeedbackFlow {
	return FeedbackFlow(bindings.ConstOfFeedbackFlow(str))
}

func (x FeedbackFlow) String() (string, bool) {
	switch x {
	case FeedbackFlow_REGULAR:
		return "regular", true
	case FeedbackFlow_LOGIN:
		return "login", true
	case FeedbackFlow_SAD_TAB_CRASH:
		return "sadTabCrash", true
	case FeedbackFlow_GOOGLE_INTERNAL:
		return "googleInternal", true
	default:
		return "", false
	}
}

type LogsMapEntry struct {
	// Key is "LogsMapEntry.key"
	//
	// Optional
	Key js.String
	// Value is "LogsMapEntry.value"
	//
	// Optional
	Value js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LogsMapEntry with all fields set.
func (p LogsMapEntry) FromRef(ref js.Ref) LogsMapEntry {
	p.UpdateFrom(ref)
	return p
}

// New creates a new LogsMapEntry in the application heap.
func (p LogsMapEntry) New() js.Ref {
	return bindings.LogsMapEntryJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *LogsMapEntry) UpdateFrom(ref js.Ref) {
	bindings.LogsMapEntryJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *LogsMapEntry) Update(ref js.Ref) {
	bindings.LogsMapEntryJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *LogsMapEntry) FreeMembers(recursive bool) {
	js.Free(
		p.Key.Ref(),
		p.Value.Ref(),
	)
	p.Key = p.Key.FromRef(js.Undefined)
	p.Value = p.Value.FromRef(js.Undefined)
}

type FeedbackInfo struct {
	// AttachedFile is "FeedbackInfo.attachedFile"
	//
	// Optional
	//
	// NOTE: AttachedFile.FFI_USE MUST be set to true to get AttachedFile used.
	AttachedFile AttachedFile
	// CategoryTag is "FeedbackInfo.categoryTag"
	//
	// Optional
	CategoryTag js.String
	// Description is "FeedbackInfo.description"
	//
	// Optional
	Description js.String
	// DescriptionPlaceholder is "FeedbackInfo.descriptionPlaceholder"
	//
	// Optional
	DescriptionPlaceholder js.String
	// Email is "FeedbackInfo.email"
	//
	// Optional
	Email js.String
	// PageUrl is "FeedbackInfo.pageUrl"
	//
	// Optional
	PageUrl js.String
	// ProductId is "FeedbackInfo.productId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ProductId MUST be set to true to make this field effective.
	ProductId int32
	// Screenshot is "FeedbackInfo.screenshot"
	//
	// Optional
	Screenshot js.Object
	// TraceId is "FeedbackInfo.traceId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TraceId MUST be set to true to make this field effective.
	TraceId int32
	// SystemInformation is "FeedbackInfo.systemInformation"
	//
	// Optional
	SystemInformation js.Array[LogsMapEntry]
	// SendHistograms is "FeedbackInfo.sendHistograms"
	//
	// Optional
	//
	// NOTE: FFI_USE_SendHistograms MUST be set to true to make this field effective.
	SendHistograms bool
	// Flow is "FeedbackInfo.flow"
	//
	// Optional
	Flow FeedbackFlow
	// AttachedFileBlobUuid is "FeedbackInfo.attachedFileBlobUuid"
	//
	// Optional
	AttachedFileBlobUuid js.String
	// ScreenshotBlobUuid is "FeedbackInfo.screenshotBlobUuid"
	//
	// Optional
	ScreenshotBlobUuid js.String
	// UseSystemWindowFrame is "FeedbackInfo.useSystemWindowFrame"
	//
	// Optional
	//
	// NOTE: FFI_USE_UseSystemWindowFrame MUST be set to true to make this field effective.
	UseSystemWindowFrame bool
	// SendBluetoothLogs is "FeedbackInfo.sendBluetoothLogs"
	//
	// Optional
	//
	// NOTE: FFI_USE_SendBluetoothLogs MUST be set to true to make this field effective.
	SendBluetoothLogs bool
	// SendTabTitles is "FeedbackInfo.sendTabTitles"
	//
	// Optional
	//
	// NOTE: FFI_USE_SendTabTitles MUST be set to true to make this field effective.
	SendTabTitles bool
	// AssistantDebugInfoAllowed is "FeedbackInfo.assistantDebugInfoAllowed"
	//
	// Optional
	//
	// NOTE: FFI_USE_AssistantDebugInfoAllowed MUST be set to true to make this field effective.
	AssistantDebugInfoAllowed bool
	// FromAssistant is "FeedbackInfo.fromAssistant"
	//
	// Optional
	//
	// NOTE: FFI_USE_FromAssistant MUST be set to true to make this field effective.
	FromAssistant bool
	// IncludeBluetoothLogs is "FeedbackInfo.includeBluetoothLogs"
	//
	// Optional
	//
	// NOTE: FFI_USE_IncludeBluetoothLogs MUST be set to true to make this field effective.
	IncludeBluetoothLogs bool
	// ShowQuestionnaire is "FeedbackInfo.showQuestionnaire"
	//
	// Optional
	//
	// NOTE: FFI_USE_ShowQuestionnaire MUST be set to true to make this field effective.
	ShowQuestionnaire bool
	// FromAutofill is "FeedbackInfo.fromAutofill"
	//
	// Optional
	//
	// NOTE: FFI_USE_FromAutofill MUST be set to true to make this field effective.
	FromAutofill bool
	// AutofillMetadata is "FeedbackInfo.autofillMetadata"
	//
	// Optional
	AutofillMetadata js.String
	// SendAutofillMetadata is "FeedbackInfo.sendAutofillMetadata"
	//
	// Optional
	//
	// NOTE: FFI_USE_SendAutofillMetadata MUST be set to true to make this field effective.
	SendAutofillMetadata bool

	FFI_USE_ProductId                 bool // for ProductId.
	FFI_USE_TraceId                   bool // for TraceId.
	FFI_USE_SendHistograms            bool // for SendHistograms.
	FFI_USE_UseSystemWindowFrame      bool // for UseSystemWindowFrame.
	FFI_USE_SendBluetoothLogs         bool // for SendBluetoothLogs.
	FFI_USE_SendTabTitles             bool // for SendTabTitles.
	FFI_USE_AssistantDebugInfoAllowed bool // for AssistantDebugInfoAllowed.
	FFI_USE_FromAssistant             bool // for FromAssistant.
	FFI_USE_IncludeBluetoothLogs      bool // for IncludeBluetoothLogs.
	FFI_USE_ShowQuestionnaire         bool // for ShowQuestionnaire.
	FFI_USE_FromAutofill              bool // for FromAutofill.
	FFI_USE_SendAutofillMetadata      bool // for SendAutofillMetadata.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FeedbackInfo with all fields set.
func (p FeedbackInfo) FromRef(ref js.Ref) FeedbackInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FeedbackInfo in the application heap.
func (p FeedbackInfo) New() js.Ref {
	return bindings.FeedbackInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FeedbackInfo) UpdateFrom(ref js.Ref) {
	bindings.FeedbackInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FeedbackInfo) Update(ref js.Ref) {
	bindings.FeedbackInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FeedbackInfo) FreeMembers(recursive bool) {
	js.Free(
		p.CategoryTag.Ref(),
		p.Description.Ref(),
		p.DescriptionPlaceholder.Ref(),
		p.Email.Ref(),
		p.PageUrl.Ref(),
		p.Screenshot.Ref(),
		p.SystemInformation.Ref(),
		p.AttachedFileBlobUuid.Ref(),
		p.ScreenshotBlobUuid.Ref(),
		p.AutofillMetadata.Ref(),
	)
	p.CategoryTag = p.CategoryTag.FromRef(js.Undefined)
	p.Description = p.Description.FromRef(js.Undefined)
	p.DescriptionPlaceholder = p.DescriptionPlaceholder.FromRef(js.Undefined)
	p.Email = p.Email.FromRef(js.Undefined)
	p.PageUrl = p.PageUrl.FromRef(js.Undefined)
	p.Screenshot = p.Screenshot.FromRef(js.Undefined)
	p.SystemInformation = p.SystemInformation.FromRef(js.Undefined)
	p.AttachedFileBlobUuid = p.AttachedFileBlobUuid.FromRef(js.Undefined)
	p.ScreenshotBlobUuid = p.ScreenshotBlobUuid.FromRef(js.Undefined)
	p.AutofillMetadata = p.AutofillMetadata.FromRef(js.Undefined)
	if recursive {
		p.AttachedFile.FreeMembers(true)
	}
}

type FeedbackSource uint32

const (
	_ FeedbackSource = iota

	FeedbackSource_QUICKOFFICE
)

func (FeedbackSource) FromRef(str js.Ref) FeedbackSource {
	return FeedbackSource(bindings.ConstOfFeedbackSource(str))
}

func (x FeedbackSource) String() (string, bool) {
	switch x {
	case FeedbackSource_QUICKOFFICE:
		return "quickoffice", true
	default:
		return "", false
	}
}

type GetSystemInformationCallbackFunc func(this js.Ref, systemInformation js.Array[LogsMapEntry]) js.Ref

func (fn GetSystemInformationCallbackFunc) Register() js.Func[func(systemInformation js.Array[LogsMapEntry])] {
	return js.RegisterCallback[func(systemInformation js.Array[LogsMapEntry])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetSystemInformationCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[LogsMapEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetSystemInformationCallback[T any] struct {
	Fn  func(arg T, this js.Ref, systemInformation js.Array[LogsMapEntry]) js.Ref
	Arg T
}

func (cb *GetSystemInformationCallback[T]) Register() js.Func[func(systemInformation js.Array[LogsMapEntry])] {
	return js.RegisterCallback[func(systemInformation js.Array[LogsMapEntry])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetSystemInformationCallback[T]) DispatchCallback(
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

		js.Array[LogsMapEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetUserEmailCallbackFunc func(this js.Ref, email js.String) js.Ref

func (fn GetUserEmailCallbackFunc) Register() js.Func[func(email js.String)] {
	return js.RegisterCallback[func(email js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetUserEmailCallbackFunc) DispatchCallback(
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

type GetUserEmailCallback[T any] struct {
	Fn  func(arg T, this js.Ref, email js.String) js.Ref
	Arg T
}

func (cb *GetUserEmailCallback[T]) Register() js.Func[func(email js.String)] {
	return js.RegisterCallback[func(email js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetUserEmailCallback[T]) DispatchCallback(
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

type LandingPageType uint32

const (
	_ LandingPageType = iota

	LandingPageType_NORMAL
	LandingPageType_TECHSTOP
	LandingPageType_NO_LANDING_PAGE
)

func (LandingPageType) FromRef(str js.Ref) LandingPageType {
	return LandingPageType(bindings.ConstOfLandingPageType(str))
}

func (x LandingPageType) String() (string, bool) {
	switch x {
	case LandingPageType_NORMAL:
		return "normal", true
	case LandingPageType_TECHSTOP:
		return "techstop", true
	case LandingPageType_NO_LANDING_PAGE:
		return "noLandingPage", true
	default:
		return "", false
	}
}

type LogSource uint32

const (
	_ LogSource = iota

	LogSource_MESSAGES
	LogSource_UI_LATEST
	LogSource_DRM_MODETEST
	LogSource_LSUSB
	LogSource_ATRUS_LOG
	LogSource_NET_LOG
	LogSource_EVENT_LOG
	LogSource_UPDATE_ENGINE_LOG
	LogSource_POWERD_LATEST
	LogSource_POWERD_PREVIOUS
	LogSource_LSPCI
	LogSource_IFCONFIG
	LogSource_UPTIME
)

func (LogSource) FromRef(str js.Ref) LogSource {
	return LogSource(bindings.ConstOfLogSource(str))
}

func (x LogSource) String() (string, bool) {
	switch x {
	case LogSource_MESSAGES:
		return "messages", true
	case LogSource_UI_LATEST:
		return "uiLatest", true
	case LogSource_DRM_MODETEST:
		return "drmModetest", true
	case LogSource_LSUSB:
		return "lsusb", true
	case LogSource_ATRUS_LOG:
		return "atrusLog", true
	case LogSource_NET_LOG:
		return "netLog", true
	case LogSource_EVENT_LOG:
		return "eventLog", true
	case LogSource_UPDATE_ENGINE_LOG:
		return "updateEngineLog", true
	case LogSource_POWERD_LATEST:
		return "powerdLatest", true
	case LogSource_POWERD_PREVIOUS:
		return "powerdPrevious", true
	case LogSource_LSPCI:
		return "lspci", true
	case LogSource_IFCONFIG:
		return "ifconfig", true
	case LogSource_UPTIME:
		return "uptime", true
	default:
		return "", false
	}
}

type ReadLogSourceCallbackFunc func(this js.Ref, result *ReadLogSourceResult) js.Ref

func (fn ReadLogSourceCallbackFunc) Register() js.Func[func(result *ReadLogSourceResult)] {
	return js.RegisterCallback[func(result *ReadLogSourceResult)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ReadLogSourceCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ReadLogSourceResult
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

type ReadLogSourceCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *ReadLogSourceResult) js.Ref
	Arg T
}

func (cb *ReadLogSourceCallback[T]) Register() js.Func[func(result *ReadLogSourceResult)] {
	return js.RegisterCallback[func(result *ReadLogSourceResult)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ReadLogSourceCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ReadLogSourceResult
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

type ReadLogSourceResult struct {
	// ReaderId is "ReadLogSourceResult.readerId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ReaderId MUST be set to true to make this field effective.
	ReaderId int32
	// LogLines is "ReadLogSourceResult.logLines"
	//
	// Optional
	LogLines js.Array[js.String]

	FFI_USE_ReaderId bool // for ReaderId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReadLogSourceResult with all fields set.
func (p ReadLogSourceResult) FromRef(ref js.Ref) ReadLogSourceResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReadLogSourceResult in the application heap.
func (p ReadLogSourceResult) New() js.Ref {
	return bindings.ReadLogSourceResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ReadLogSourceResult) UpdateFrom(ref js.Ref) {
	bindings.ReadLogSourceResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ReadLogSourceResult) Update(ref js.Ref) {
	bindings.ReadLogSourceResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ReadLogSourceResult) FreeMembers(recursive bool) {
	js.Free(
		p.LogLines.Ref(),
	)
	p.LogLines = p.LogLines.FromRef(js.Undefined)
}

type ReadLogSourceParams struct {
	// Source is "ReadLogSourceParams.source"
	//
	// Optional
	Source LogSource
	// Incremental is "ReadLogSourceParams.incremental"
	//
	// Optional
	//
	// NOTE: FFI_USE_Incremental MUST be set to true to make this field effective.
	Incremental bool
	// ReaderId is "ReadLogSourceParams.readerId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ReaderId MUST be set to true to make this field effective.
	ReaderId int32

	FFI_USE_Incremental bool // for Incremental.
	FFI_USE_ReaderId    bool // for ReaderId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReadLogSourceParams with all fields set.
func (p ReadLogSourceParams) FromRef(ref js.Ref) ReadLogSourceParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReadLogSourceParams in the application heap.
func (p ReadLogSourceParams) New() js.Ref {
	return bindings.ReadLogSourceParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ReadLogSourceParams) UpdateFrom(ref js.Ref) {
	bindings.ReadLogSourceParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ReadLogSourceParams) Update(ref js.Ref) {
	bindings.ReadLogSourceParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ReadLogSourceParams) FreeMembers(recursive bool) {
}

type SendFeedbackCallbackFunc func(this js.Ref, result *SendFeedbackResult) js.Ref

func (fn SendFeedbackCallbackFunc) Register() js.Func[func(result *SendFeedbackResult)] {
	return js.RegisterCallback[func(result *SendFeedbackResult)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SendFeedbackCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SendFeedbackResult
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

type SendFeedbackCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *SendFeedbackResult) js.Ref
	Arg T
}

func (cb *SendFeedbackCallback[T]) Register() js.Func[func(result *SendFeedbackResult)] {
	return js.RegisterCallback[func(result *SendFeedbackResult)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SendFeedbackCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SendFeedbackResult
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

type Status uint32

const (
	_ Status = iota

	Status_SUCCESS
	Status_DELAYED
)

func (Status) FromRef(str js.Ref) Status {
	return Status(bindings.ConstOfStatus(str))
}

func (x Status) String() (string, bool) {
	switch x {
	case Status_SUCCESS:
		return "success", true
	case Status_DELAYED:
		return "delayed", true
	default:
		return "", false
	}
}

type SendFeedbackResult struct {
	// Status is "SendFeedbackResult.status"
	//
	// Optional
	Status Status
	// LandingPageType is "SendFeedbackResult.landingPageType"
	//
	// Optional
	LandingPageType LandingPageType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SendFeedbackResult with all fields set.
func (p SendFeedbackResult) FromRef(ref js.Ref) SendFeedbackResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SendFeedbackResult in the application heap.
func (p SendFeedbackResult) New() js.Ref {
	return bindings.SendFeedbackResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SendFeedbackResult) UpdateFrom(ref js.Ref) {
	bindings.SendFeedbackResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SendFeedbackResult) Update(ref js.Ref) {
	bindings.SendFeedbackResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SendFeedbackResult) FreeMembers(recursive bool) {
}

// HasFuncGetSystemInformation returns true if the function "WEBEXT.feedbackPrivate.getSystemInformation" exists.
func HasFuncGetSystemInformation() bool {
	return js.True == bindings.HasFuncGetSystemInformation()
}

// FuncGetSystemInformation returns the function "WEBEXT.feedbackPrivate.getSystemInformation".
func FuncGetSystemInformation() (fn js.Func[func(callback js.Func[func(systemInformation js.Array[LogsMapEntry])])]) {
	bindings.FuncGetSystemInformation(
		js.Pointer(&fn),
	)
	return
}

// GetSystemInformation calls the function "WEBEXT.feedbackPrivate.getSystemInformation" directly.
func GetSystemInformation(callback js.Func[func(systemInformation js.Array[LogsMapEntry])]) (ret js.Void) {
	bindings.CallGetSystemInformation(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryGetSystemInformation calls the function "WEBEXT.feedbackPrivate.getSystemInformation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetSystemInformation(callback js.Func[func(systemInformation js.Array[LogsMapEntry])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetSystemInformation(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncGetUserEmail returns true if the function "WEBEXT.feedbackPrivate.getUserEmail" exists.
func HasFuncGetUserEmail() bool {
	return js.True == bindings.HasFuncGetUserEmail()
}

// FuncGetUserEmail returns the function "WEBEXT.feedbackPrivate.getUserEmail".
func FuncGetUserEmail() (fn js.Func[func(callback js.Func[func(email js.String)])]) {
	bindings.FuncGetUserEmail(
		js.Pointer(&fn),
	)
	return
}

// GetUserEmail calls the function "WEBEXT.feedbackPrivate.getUserEmail" directly.
func GetUserEmail(callback js.Func[func(email js.String)]) (ret js.Void) {
	bindings.CallGetUserEmail(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryGetUserEmail calls the function "WEBEXT.feedbackPrivate.getUserEmail"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetUserEmail(callback js.Func[func(email js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetUserEmail(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnFeedbackRequestedEventCallbackFunc func(this js.Ref, feedback *FeedbackInfo) js.Ref

func (fn OnFeedbackRequestedEventCallbackFunc) Register() js.Func[func(feedback *FeedbackInfo)] {
	return js.RegisterCallback[func(feedback *FeedbackInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnFeedbackRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 FeedbackInfo
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

type OnFeedbackRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, feedback *FeedbackInfo) js.Ref
	Arg T
}

func (cb *OnFeedbackRequestedEventCallback[T]) Register() js.Func[func(feedback *FeedbackInfo)] {
	return js.RegisterCallback[func(feedback *FeedbackInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnFeedbackRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 FeedbackInfo
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

// HasFuncOnFeedbackRequested returns true if the function "WEBEXT.feedbackPrivate.onFeedbackRequested.addListener" exists.
func HasFuncOnFeedbackRequested() bool {
	return js.True == bindings.HasFuncOnFeedbackRequested()
}

// FuncOnFeedbackRequested returns the function "WEBEXT.feedbackPrivate.onFeedbackRequested.addListener".
func FuncOnFeedbackRequested() (fn js.Func[func(callback js.Func[func(feedback *FeedbackInfo)])]) {
	bindings.FuncOnFeedbackRequested(
		js.Pointer(&fn),
	)
	return
}

// OnFeedbackRequested calls the function "WEBEXT.feedbackPrivate.onFeedbackRequested.addListener" directly.
func OnFeedbackRequested(callback js.Func[func(feedback *FeedbackInfo)]) (ret js.Void) {
	bindings.CallOnFeedbackRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnFeedbackRequested calls the function "WEBEXT.feedbackPrivate.onFeedbackRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnFeedbackRequested(callback js.Func[func(feedback *FeedbackInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnFeedbackRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffFeedbackRequested returns true if the function "WEBEXT.feedbackPrivate.onFeedbackRequested.removeListener" exists.
func HasFuncOffFeedbackRequested() bool {
	return js.True == bindings.HasFuncOffFeedbackRequested()
}

// FuncOffFeedbackRequested returns the function "WEBEXT.feedbackPrivate.onFeedbackRequested.removeListener".
func FuncOffFeedbackRequested() (fn js.Func[func(callback js.Func[func(feedback *FeedbackInfo)])]) {
	bindings.FuncOffFeedbackRequested(
		js.Pointer(&fn),
	)
	return
}

// OffFeedbackRequested calls the function "WEBEXT.feedbackPrivate.onFeedbackRequested.removeListener" directly.
func OffFeedbackRequested(callback js.Func[func(feedback *FeedbackInfo)]) (ret js.Void) {
	bindings.CallOffFeedbackRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffFeedbackRequested calls the function "WEBEXT.feedbackPrivate.onFeedbackRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffFeedbackRequested(callback js.Func[func(feedback *FeedbackInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffFeedbackRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnFeedbackRequested returns true if the function "WEBEXT.feedbackPrivate.onFeedbackRequested.hasListener" exists.
func HasFuncHasOnFeedbackRequested() bool {
	return js.True == bindings.HasFuncHasOnFeedbackRequested()
}

// FuncHasOnFeedbackRequested returns the function "WEBEXT.feedbackPrivate.onFeedbackRequested.hasListener".
func FuncHasOnFeedbackRequested() (fn js.Func[func(callback js.Func[func(feedback *FeedbackInfo)]) bool]) {
	bindings.FuncHasOnFeedbackRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnFeedbackRequested calls the function "WEBEXT.feedbackPrivate.onFeedbackRequested.hasListener" directly.
func HasOnFeedbackRequested(callback js.Func[func(feedback *FeedbackInfo)]) (ret bool) {
	bindings.CallHasOnFeedbackRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnFeedbackRequested calls the function "WEBEXT.feedbackPrivate.onFeedbackRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnFeedbackRequested(callback js.Func[func(feedback *FeedbackInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnFeedbackRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOpenFeedback returns true if the function "WEBEXT.feedbackPrivate.openFeedback" exists.
func HasFuncOpenFeedback() bool {
	return js.True == bindings.HasFuncOpenFeedback()
}

// FuncOpenFeedback returns the function "WEBEXT.feedbackPrivate.openFeedback".
func FuncOpenFeedback() (fn js.Func[func(source FeedbackSource)]) {
	bindings.FuncOpenFeedback(
		js.Pointer(&fn),
	)
	return
}

// OpenFeedback calls the function "WEBEXT.feedbackPrivate.openFeedback" directly.
func OpenFeedback(source FeedbackSource) (ret js.Void) {
	bindings.CallOpenFeedback(
		js.Pointer(&ret),
		uint32(source),
	)

	return
}

// TryOpenFeedback calls the function "WEBEXT.feedbackPrivate.openFeedback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenFeedback(source FeedbackSource) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenFeedback(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(source),
	)

	return
}

// HasFuncReadLogSource returns true if the function "WEBEXT.feedbackPrivate.readLogSource" exists.
func HasFuncReadLogSource() bool {
	return js.True == bindings.HasFuncReadLogSource()
}

// FuncReadLogSource returns the function "WEBEXT.feedbackPrivate.readLogSource".
func FuncReadLogSource() (fn js.Func[func(params ReadLogSourceParams, callback js.Func[func(result *ReadLogSourceResult)])]) {
	bindings.FuncReadLogSource(
		js.Pointer(&fn),
	)
	return
}

// ReadLogSource calls the function "WEBEXT.feedbackPrivate.readLogSource" directly.
func ReadLogSource(params ReadLogSourceParams, callback js.Func[func(result *ReadLogSourceResult)]) (ret js.Void) {
	bindings.CallReadLogSource(
		js.Pointer(&ret),
		js.Pointer(&params),
		callback.Ref(),
	)

	return
}

// TryReadLogSource calls the function "WEBEXT.feedbackPrivate.readLogSource"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReadLogSource(params ReadLogSourceParams, callback js.Func[func(result *ReadLogSourceResult)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadLogSource(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&params),
		callback.Ref(),
	)

	return
}

// HasFuncSendFeedback returns true if the function "WEBEXT.feedbackPrivate.sendFeedback" exists.
func HasFuncSendFeedback() bool {
	return js.True == bindings.HasFuncSendFeedback()
}

// FuncSendFeedback returns the function "WEBEXT.feedbackPrivate.sendFeedback".
func FuncSendFeedback() (fn js.Func[func(feedback FeedbackInfo, loadSystemInfo bool, formOpenTime float64) js.Promise[SendFeedbackResult]]) {
	bindings.FuncSendFeedback(
		js.Pointer(&fn),
	)
	return
}

// SendFeedback calls the function "WEBEXT.feedbackPrivate.sendFeedback" directly.
func SendFeedback(feedback FeedbackInfo, loadSystemInfo bool, formOpenTime float64) (ret js.Promise[SendFeedbackResult]) {
	bindings.CallSendFeedback(
		js.Pointer(&ret),
		js.Pointer(&feedback),
		js.Bool(bool(loadSystemInfo)),
		float64(formOpenTime),
	)

	return
}

// TrySendFeedback calls the function "WEBEXT.feedbackPrivate.sendFeedback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendFeedback(feedback FeedbackInfo, loadSystemInfo bool, formOpenTime float64) (ret js.Promise[SendFeedbackResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendFeedback(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&feedback),
		js.Bool(bool(loadSystemInfo)),
		float64(formOpenTime),
	)

	return
}
