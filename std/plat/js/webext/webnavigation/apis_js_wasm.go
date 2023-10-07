// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package webnavigation

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/extensiontypes"
	"github.com/primecitizens/pcz/std/plat/js/webext/webnavigation/bindings"
)

type GetAllFramesArgDetails struct {
	// TabId is "GetAllFramesArgDetails.tabId"
	//
	// Required
	TabId int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetAllFramesArgDetails with all fields set.
func (p GetAllFramesArgDetails) FromRef(ref js.Ref) GetAllFramesArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetAllFramesArgDetails in the application heap.
func (p GetAllFramesArgDetails) New() js.Ref {
	return bindings.GetAllFramesArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetAllFramesArgDetails) UpdateFrom(ref js.Ref) {
	bindings.GetAllFramesArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetAllFramesArgDetails) Update(ref js.Ref) {
	bindings.GetAllFramesArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetAllFramesArgDetails) FreeMembers(recursive bool) {
}

type GetAllFramesReturnTypeElem struct {
	// DocumentId is "GetAllFramesReturnTypeElem.documentId"
	//
	// Required
	DocumentId js.String
	// DocumentLifecycle is "GetAllFramesReturnTypeElem.documentLifecycle"
	//
	// Required
	DocumentLifecycle extensiontypes.DocumentLifecycle
	// ErrorOccurred is "GetAllFramesReturnTypeElem.errorOccurred"
	//
	// Required
	ErrorOccurred bool
	// FrameId is "GetAllFramesReturnTypeElem.frameId"
	//
	// Required
	FrameId int64
	// FrameType is "GetAllFramesReturnTypeElem.frameType"
	//
	// Required
	FrameType extensiontypes.FrameType
	// ParentDocumentId is "GetAllFramesReturnTypeElem.parentDocumentId"
	//
	// Optional
	ParentDocumentId js.String
	// ParentFrameId is "GetAllFramesReturnTypeElem.parentFrameId"
	//
	// Required
	ParentFrameId int64
	// ProcessId is "GetAllFramesReturnTypeElem.processId"
	//
	// Required
	ProcessId int64
	// Url is "GetAllFramesReturnTypeElem.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetAllFramesReturnTypeElem with all fields set.
func (p GetAllFramesReturnTypeElem) FromRef(ref js.Ref) GetAllFramesReturnTypeElem {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetAllFramesReturnTypeElem in the application heap.
func (p GetAllFramesReturnTypeElem) New() js.Ref {
	return bindings.GetAllFramesReturnTypeElemJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetAllFramesReturnTypeElem) UpdateFrom(ref js.Ref) {
	bindings.GetAllFramesReturnTypeElemJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetAllFramesReturnTypeElem) Update(ref js.Ref) {
	bindings.GetAllFramesReturnTypeElemJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetAllFramesReturnTypeElem) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentId.Ref(),
		p.ParentDocumentId.Ref(),
		p.Url.Ref(),
	)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
	p.ParentDocumentId = p.ParentDocumentId.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type GetFrameArgDetails struct {
	// DocumentId is "GetFrameArgDetails.documentId"
	//
	// Optional
	DocumentId js.String
	// FrameId is "GetFrameArgDetails.frameId"
	//
	// Optional
	//
	// NOTE: FFI_USE_FrameId MUST be set to true to make this field effective.
	FrameId int64
	// ProcessId is "GetFrameArgDetails.processId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ProcessId MUST be set to true to make this field effective.
	ProcessId int64
	// TabId is "GetFrameArgDetails.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int64

	FFI_USE_FrameId   bool // for FrameId.
	FFI_USE_ProcessId bool // for ProcessId.
	FFI_USE_TabId     bool // for TabId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetFrameArgDetails with all fields set.
func (p GetFrameArgDetails) FromRef(ref js.Ref) GetFrameArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetFrameArgDetails in the application heap.
func (p GetFrameArgDetails) New() js.Ref {
	return bindings.GetFrameArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetFrameArgDetails) UpdateFrom(ref js.Ref) {
	bindings.GetFrameArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetFrameArgDetails) Update(ref js.Ref) {
	bindings.GetFrameArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetFrameArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentId.Ref(),
	)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
}

type GetFrameReturnType struct {
	// DocumentId is "GetFrameReturnType.documentId"
	//
	// Required
	DocumentId js.String
	// DocumentLifecycle is "GetFrameReturnType.documentLifecycle"
	//
	// Required
	DocumentLifecycle extensiontypes.DocumentLifecycle
	// ErrorOccurred is "GetFrameReturnType.errorOccurred"
	//
	// Required
	ErrorOccurred bool
	// FrameType is "GetFrameReturnType.frameType"
	//
	// Required
	FrameType extensiontypes.FrameType
	// ParentDocumentId is "GetFrameReturnType.parentDocumentId"
	//
	// Optional
	ParentDocumentId js.String
	// ParentFrameId is "GetFrameReturnType.parentFrameId"
	//
	// Required
	ParentFrameId int64
	// Url is "GetFrameReturnType.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetFrameReturnType with all fields set.
func (p GetFrameReturnType) FromRef(ref js.Ref) GetFrameReturnType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetFrameReturnType in the application heap.
func (p GetFrameReturnType) New() js.Ref {
	return bindings.GetFrameReturnTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetFrameReturnType) UpdateFrom(ref js.Ref) {
	bindings.GetFrameReturnTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetFrameReturnType) Update(ref js.Ref) {
	bindings.GetFrameReturnTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetFrameReturnType) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentId.Ref(),
		p.ParentDocumentId.Ref(),
		p.Url.Ref(),
	)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
	p.ParentDocumentId = p.ParentDocumentId.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type OnBeforeNavigateArgDetails struct {
	// DocumentLifecycle is "OnBeforeNavigateArgDetails.documentLifecycle"
	//
	// Required
	DocumentLifecycle extensiontypes.DocumentLifecycle
	// FrameId is "OnBeforeNavigateArgDetails.frameId"
	//
	// Required
	FrameId int64
	// FrameType is "OnBeforeNavigateArgDetails.frameType"
	//
	// Required
	FrameType extensiontypes.FrameType
	// ParentDocumentId is "OnBeforeNavigateArgDetails.parentDocumentId"
	//
	// Optional
	ParentDocumentId js.String
	// ParentFrameId is "OnBeforeNavigateArgDetails.parentFrameId"
	//
	// Required
	ParentFrameId int64
	// ProcessId is "OnBeforeNavigateArgDetails.processId"
	//
	// Required
	ProcessId int64
	// TabId is "OnBeforeNavigateArgDetails.tabId"
	//
	// Required
	TabId int64
	// TimeStamp is "OnBeforeNavigateArgDetails.timeStamp"
	//
	// Required
	TimeStamp float64
	// Url is "OnBeforeNavigateArgDetails.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnBeforeNavigateArgDetails with all fields set.
func (p OnBeforeNavigateArgDetails) FromRef(ref js.Ref) OnBeforeNavigateArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnBeforeNavigateArgDetails in the application heap.
func (p OnBeforeNavigateArgDetails) New() js.Ref {
	return bindings.OnBeforeNavigateArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnBeforeNavigateArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnBeforeNavigateArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnBeforeNavigateArgDetails) Update(ref js.Ref) {
	bindings.OnBeforeNavigateArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnBeforeNavigateArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.ParentDocumentId.Ref(),
		p.Url.Ref(),
	)
	p.ParentDocumentId = p.ParentDocumentId.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type TransitionQualifier uint32

const (
	_ TransitionQualifier = iota

	TransitionQualifier_CLIENT_REDIRECT
	TransitionQualifier_SERVER_REDIRECT
	TransitionQualifier_FORWARD_BACK
	TransitionQualifier_FROM_ADDRESS_BAR
)

func (TransitionQualifier) FromRef(str js.Ref) TransitionQualifier {
	return TransitionQualifier(bindings.ConstOfTransitionQualifier(str))
}

func (x TransitionQualifier) String() (string, bool) {
	switch x {
	case TransitionQualifier_CLIENT_REDIRECT:
		return "client_redirect", true
	case TransitionQualifier_SERVER_REDIRECT:
		return "server_redirect", true
	case TransitionQualifier_FORWARD_BACK:
		return "forward_back", true
	case TransitionQualifier_FROM_ADDRESS_BAR:
		return "from_address_bar", true
	default:
		return "", false
	}
}

type TransitionType uint32

const (
	_ TransitionType = iota

	TransitionType_LINK
	TransitionType_TYPED
	TransitionType_AUTO_BOOKMARK
	TransitionType_AUTO_SUBFRAME
	TransitionType_MANUAL_SUBFRAME
	TransitionType_GENERATED
	TransitionType_START_PAGE
	TransitionType_FORM_SUBMIT
	TransitionType_RELOAD
	TransitionType_KEYWORD
	TransitionType_KEYWORD_GENERATED
)

func (TransitionType) FromRef(str js.Ref) TransitionType {
	return TransitionType(bindings.ConstOfTransitionType(str))
}

func (x TransitionType) String() (string, bool) {
	switch x {
	case TransitionType_LINK:
		return "link", true
	case TransitionType_TYPED:
		return "typed", true
	case TransitionType_AUTO_BOOKMARK:
		return "auto_bookmark", true
	case TransitionType_AUTO_SUBFRAME:
		return "auto_subframe", true
	case TransitionType_MANUAL_SUBFRAME:
		return "manual_subframe", true
	case TransitionType_GENERATED:
		return "generated", true
	case TransitionType_START_PAGE:
		return "start_page", true
	case TransitionType_FORM_SUBMIT:
		return "form_submit", true
	case TransitionType_RELOAD:
		return "reload", true
	case TransitionType_KEYWORD:
		return "keyword", true
	case TransitionType_KEYWORD_GENERATED:
		return "keyword_generated", true
	default:
		return "", false
	}
}

type OnCommittedArgDetails struct {
	// DocumentId is "OnCommittedArgDetails.documentId"
	//
	// Required
	DocumentId js.String
	// DocumentLifecycle is "OnCommittedArgDetails.documentLifecycle"
	//
	// Required
	DocumentLifecycle extensiontypes.DocumentLifecycle
	// FrameId is "OnCommittedArgDetails.frameId"
	//
	// Required
	FrameId int64
	// FrameType is "OnCommittedArgDetails.frameType"
	//
	// Required
	FrameType extensiontypes.FrameType
	// ParentDocumentId is "OnCommittedArgDetails.parentDocumentId"
	//
	// Optional
	ParentDocumentId js.String
	// ParentFrameId is "OnCommittedArgDetails.parentFrameId"
	//
	// Required
	ParentFrameId int64
	// ProcessId is "OnCommittedArgDetails.processId"
	//
	// Required
	ProcessId int64
	// TabId is "OnCommittedArgDetails.tabId"
	//
	// Required
	TabId int64
	// TimeStamp is "OnCommittedArgDetails.timeStamp"
	//
	// Required
	TimeStamp float64
	// TransitionQualifiers is "OnCommittedArgDetails.transitionQualifiers"
	//
	// Required
	TransitionQualifiers js.Array[TransitionQualifier]
	// TransitionType is "OnCommittedArgDetails.transitionType"
	//
	// Required
	TransitionType TransitionType
	// Url is "OnCommittedArgDetails.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnCommittedArgDetails with all fields set.
func (p OnCommittedArgDetails) FromRef(ref js.Ref) OnCommittedArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnCommittedArgDetails in the application heap.
func (p OnCommittedArgDetails) New() js.Ref {
	return bindings.OnCommittedArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnCommittedArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnCommittedArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnCommittedArgDetails) Update(ref js.Ref) {
	bindings.OnCommittedArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnCommittedArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentId.Ref(),
		p.ParentDocumentId.Ref(),
		p.TransitionQualifiers.Ref(),
		p.Url.Ref(),
	)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
	p.ParentDocumentId = p.ParentDocumentId.FromRef(js.Undefined)
	p.TransitionQualifiers = p.TransitionQualifiers.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type OnCompletedArgDetails struct {
	// DocumentId is "OnCompletedArgDetails.documentId"
	//
	// Required
	DocumentId js.String
	// DocumentLifecycle is "OnCompletedArgDetails.documentLifecycle"
	//
	// Required
	DocumentLifecycle extensiontypes.DocumentLifecycle
	// FrameId is "OnCompletedArgDetails.frameId"
	//
	// Required
	FrameId int64
	// FrameType is "OnCompletedArgDetails.frameType"
	//
	// Required
	FrameType extensiontypes.FrameType
	// ParentDocumentId is "OnCompletedArgDetails.parentDocumentId"
	//
	// Optional
	ParentDocumentId js.String
	// ParentFrameId is "OnCompletedArgDetails.parentFrameId"
	//
	// Required
	ParentFrameId int64
	// ProcessId is "OnCompletedArgDetails.processId"
	//
	// Required
	ProcessId int64
	// TabId is "OnCompletedArgDetails.tabId"
	//
	// Required
	TabId int64
	// TimeStamp is "OnCompletedArgDetails.timeStamp"
	//
	// Required
	TimeStamp float64
	// Url is "OnCompletedArgDetails.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnCompletedArgDetails with all fields set.
func (p OnCompletedArgDetails) FromRef(ref js.Ref) OnCompletedArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnCompletedArgDetails in the application heap.
func (p OnCompletedArgDetails) New() js.Ref {
	return bindings.OnCompletedArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnCompletedArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnCompletedArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnCompletedArgDetails) Update(ref js.Ref) {
	bindings.OnCompletedArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnCompletedArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentId.Ref(),
		p.ParentDocumentId.Ref(),
		p.Url.Ref(),
	)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
	p.ParentDocumentId = p.ParentDocumentId.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type OnCreatedNavigationTargetArgDetails struct {
	// SourceFrameId is "OnCreatedNavigationTargetArgDetails.sourceFrameId"
	//
	// Required
	SourceFrameId int64
	// SourceProcessId is "OnCreatedNavigationTargetArgDetails.sourceProcessId"
	//
	// Required
	SourceProcessId int64
	// SourceTabId is "OnCreatedNavigationTargetArgDetails.sourceTabId"
	//
	// Required
	SourceTabId int64
	// TabId is "OnCreatedNavigationTargetArgDetails.tabId"
	//
	// Required
	TabId int64
	// TimeStamp is "OnCreatedNavigationTargetArgDetails.timeStamp"
	//
	// Required
	TimeStamp float64
	// Url is "OnCreatedNavigationTargetArgDetails.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnCreatedNavigationTargetArgDetails with all fields set.
func (p OnCreatedNavigationTargetArgDetails) FromRef(ref js.Ref) OnCreatedNavigationTargetArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnCreatedNavigationTargetArgDetails in the application heap.
func (p OnCreatedNavigationTargetArgDetails) New() js.Ref {
	return bindings.OnCreatedNavigationTargetArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnCreatedNavigationTargetArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnCreatedNavigationTargetArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnCreatedNavigationTargetArgDetails) Update(ref js.Ref) {
	bindings.OnCreatedNavigationTargetArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnCreatedNavigationTargetArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
}

type OnDOMContentLoadedArgDetails struct {
	// DocumentId is "OnDOMContentLoadedArgDetails.documentId"
	//
	// Required
	DocumentId js.String
	// DocumentLifecycle is "OnDOMContentLoadedArgDetails.documentLifecycle"
	//
	// Required
	DocumentLifecycle extensiontypes.DocumentLifecycle
	// FrameId is "OnDOMContentLoadedArgDetails.frameId"
	//
	// Required
	FrameId int64
	// FrameType is "OnDOMContentLoadedArgDetails.frameType"
	//
	// Required
	FrameType extensiontypes.FrameType
	// ParentDocumentId is "OnDOMContentLoadedArgDetails.parentDocumentId"
	//
	// Optional
	ParentDocumentId js.String
	// ParentFrameId is "OnDOMContentLoadedArgDetails.parentFrameId"
	//
	// Required
	ParentFrameId int64
	// ProcessId is "OnDOMContentLoadedArgDetails.processId"
	//
	// Required
	ProcessId int64
	// TabId is "OnDOMContentLoadedArgDetails.tabId"
	//
	// Required
	TabId int64
	// TimeStamp is "OnDOMContentLoadedArgDetails.timeStamp"
	//
	// Required
	TimeStamp float64
	// Url is "OnDOMContentLoadedArgDetails.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnDOMContentLoadedArgDetails with all fields set.
func (p OnDOMContentLoadedArgDetails) FromRef(ref js.Ref) OnDOMContentLoadedArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnDOMContentLoadedArgDetails in the application heap.
func (p OnDOMContentLoadedArgDetails) New() js.Ref {
	return bindings.OnDOMContentLoadedArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnDOMContentLoadedArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnDOMContentLoadedArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnDOMContentLoadedArgDetails) Update(ref js.Ref) {
	bindings.OnDOMContentLoadedArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnDOMContentLoadedArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentId.Ref(),
		p.ParentDocumentId.Ref(),
		p.Url.Ref(),
	)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
	p.ParentDocumentId = p.ParentDocumentId.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type OnErrorOccurredArgDetails struct {
	// DocumentId is "OnErrorOccurredArgDetails.documentId"
	//
	// Required
	DocumentId js.String
	// DocumentLifecycle is "OnErrorOccurredArgDetails.documentLifecycle"
	//
	// Required
	DocumentLifecycle extensiontypes.DocumentLifecycle
	// Error is "OnErrorOccurredArgDetails.error"
	//
	// Required
	Error js.String
	// FrameId is "OnErrorOccurredArgDetails.frameId"
	//
	// Required
	FrameId int64
	// FrameType is "OnErrorOccurredArgDetails.frameType"
	//
	// Required
	FrameType extensiontypes.FrameType
	// ParentDocumentId is "OnErrorOccurredArgDetails.parentDocumentId"
	//
	// Optional
	ParentDocumentId js.String
	// ParentFrameId is "OnErrorOccurredArgDetails.parentFrameId"
	//
	// Required
	ParentFrameId int64
	// ProcessId is "OnErrorOccurredArgDetails.processId"
	//
	// Required
	ProcessId int64
	// TabId is "OnErrorOccurredArgDetails.tabId"
	//
	// Required
	TabId int64
	// TimeStamp is "OnErrorOccurredArgDetails.timeStamp"
	//
	// Required
	TimeStamp float64
	// Url is "OnErrorOccurredArgDetails.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnErrorOccurredArgDetails with all fields set.
func (p OnErrorOccurredArgDetails) FromRef(ref js.Ref) OnErrorOccurredArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnErrorOccurredArgDetails in the application heap.
func (p OnErrorOccurredArgDetails) New() js.Ref {
	return bindings.OnErrorOccurredArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnErrorOccurredArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnErrorOccurredArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnErrorOccurredArgDetails) Update(ref js.Ref) {
	bindings.OnErrorOccurredArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnErrorOccurredArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentId.Ref(),
		p.Error.Ref(),
		p.ParentDocumentId.Ref(),
		p.Url.Ref(),
	)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
	p.Error = p.Error.FromRef(js.Undefined)
	p.ParentDocumentId = p.ParentDocumentId.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type OnHistoryStateUpdatedArgDetails struct {
	// DocumentId is "OnHistoryStateUpdatedArgDetails.documentId"
	//
	// Required
	DocumentId js.String
	// DocumentLifecycle is "OnHistoryStateUpdatedArgDetails.documentLifecycle"
	//
	// Required
	DocumentLifecycle extensiontypes.DocumentLifecycle
	// FrameId is "OnHistoryStateUpdatedArgDetails.frameId"
	//
	// Required
	FrameId int64
	// FrameType is "OnHistoryStateUpdatedArgDetails.frameType"
	//
	// Required
	FrameType extensiontypes.FrameType
	// ParentDocumentId is "OnHistoryStateUpdatedArgDetails.parentDocumentId"
	//
	// Optional
	ParentDocumentId js.String
	// ParentFrameId is "OnHistoryStateUpdatedArgDetails.parentFrameId"
	//
	// Required
	ParentFrameId int64
	// ProcessId is "OnHistoryStateUpdatedArgDetails.processId"
	//
	// Required
	ProcessId int64
	// TabId is "OnHistoryStateUpdatedArgDetails.tabId"
	//
	// Required
	TabId int64
	// TimeStamp is "OnHistoryStateUpdatedArgDetails.timeStamp"
	//
	// Required
	TimeStamp float64
	// TransitionQualifiers is "OnHistoryStateUpdatedArgDetails.transitionQualifiers"
	//
	// Required
	TransitionQualifiers js.Array[TransitionQualifier]
	// TransitionType is "OnHistoryStateUpdatedArgDetails.transitionType"
	//
	// Required
	TransitionType TransitionType
	// Url is "OnHistoryStateUpdatedArgDetails.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnHistoryStateUpdatedArgDetails with all fields set.
func (p OnHistoryStateUpdatedArgDetails) FromRef(ref js.Ref) OnHistoryStateUpdatedArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnHistoryStateUpdatedArgDetails in the application heap.
func (p OnHistoryStateUpdatedArgDetails) New() js.Ref {
	return bindings.OnHistoryStateUpdatedArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnHistoryStateUpdatedArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnHistoryStateUpdatedArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnHistoryStateUpdatedArgDetails) Update(ref js.Ref) {
	bindings.OnHistoryStateUpdatedArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnHistoryStateUpdatedArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentId.Ref(),
		p.ParentDocumentId.Ref(),
		p.TransitionQualifiers.Ref(),
		p.Url.Ref(),
	)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
	p.ParentDocumentId = p.ParentDocumentId.FromRef(js.Undefined)
	p.TransitionQualifiers = p.TransitionQualifiers.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type OnReferenceFragmentUpdatedArgDetails struct {
	// DocumentId is "OnReferenceFragmentUpdatedArgDetails.documentId"
	//
	// Required
	DocumentId js.String
	// DocumentLifecycle is "OnReferenceFragmentUpdatedArgDetails.documentLifecycle"
	//
	// Required
	DocumentLifecycle extensiontypes.DocumentLifecycle
	// FrameId is "OnReferenceFragmentUpdatedArgDetails.frameId"
	//
	// Required
	FrameId int64
	// FrameType is "OnReferenceFragmentUpdatedArgDetails.frameType"
	//
	// Required
	FrameType extensiontypes.FrameType
	// ParentDocumentId is "OnReferenceFragmentUpdatedArgDetails.parentDocumentId"
	//
	// Optional
	ParentDocumentId js.String
	// ParentFrameId is "OnReferenceFragmentUpdatedArgDetails.parentFrameId"
	//
	// Required
	ParentFrameId int64
	// ProcessId is "OnReferenceFragmentUpdatedArgDetails.processId"
	//
	// Required
	ProcessId int64
	// TabId is "OnReferenceFragmentUpdatedArgDetails.tabId"
	//
	// Required
	TabId int64
	// TimeStamp is "OnReferenceFragmentUpdatedArgDetails.timeStamp"
	//
	// Required
	TimeStamp float64
	// TransitionQualifiers is "OnReferenceFragmentUpdatedArgDetails.transitionQualifiers"
	//
	// Required
	TransitionQualifiers js.Array[TransitionQualifier]
	// TransitionType is "OnReferenceFragmentUpdatedArgDetails.transitionType"
	//
	// Required
	TransitionType TransitionType
	// Url is "OnReferenceFragmentUpdatedArgDetails.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnReferenceFragmentUpdatedArgDetails with all fields set.
func (p OnReferenceFragmentUpdatedArgDetails) FromRef(ref js.Ref) OnReferenceFragmentUpdatedArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnReferenceFragmentUpdatedArgDetails in the application heap.
func (p OnReferenceFragmentUpdatedArgDetails) New() js.Ref {
	return bindings.OnReferenceFragmentUpdatedArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnReferenceFragmentUpdatedArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnReferenceFragmentUpdatedArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnReferenceFragmentUpdatedArgDetails) Update(ref js.Ref) {
	bindings.OnReferenceFragmentUpdatedArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnReferenceFragmentUpdatedArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentId.Ref(),
		p.ParentDocumentId.Ref(),
		p.TransitionQualifiers.Ref(),
		p.Url.Ref(),
	)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
	p.ParentDocumentId = p.ParentDocumentId.FromRef(js.Undefined)
	p.TransitionQualifiers = p.TransitionQualifiers.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type OnTabReplacedArgDetails struct {
	// ReplacedTabId is "OnTabReplacedArgDetails.replacedTabId"
	//
	// Required
	ReplacedTabId int64
	// TabId is "OnTabReplacedArgDetails.tabId"
	//
	// Required
	TabId int64
	// TimeStamp is "OnTabReplacedArgDetails.timeStamp"
	//
	// Required
	TimeStamp float64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnTabReplacedArgDetails with all fields set.
func (p OnTabReplacedArgDetails) FromRef(ref js.Ref) OnTabReplacedArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnTabReplacedArgDetails in the application heap.
func (p OnTabReplacedArgDetails) New() js.Ref {
	return bindings.OnTabReplacedArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnTabReplacedArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnTabReplacedArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnTabReplacedArgDetails) Update(ref js.Ref) {
	bindings.OnTabReplacedArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnTabReplacedArgDetails) FreeMembers(recursive bool) {
}

// HasFuncGetAllFrames returns true if the function "WEBEXT.webNavigation.getAllFrames" exists.
func HasFuncGetAllFrames() bool {
	return js.True == bindings.HasFuncGetAllFrames()
}

// FuncGetAllFrames returns the function "WEBEXT.webNavigation.getAllFrames".
func FuncGetAllFrames() (fn js.Func[func(details GetAllFramesArgDetails) js.Promise[js.Array[GetAllFramesReturnTypeElem]]]) {
	bindings.FuncGetAllFrames(
		js.Pointer(&fn),
	)
	return
}

// GetAllFrames calls the function "WEBEXT.webNavigation.getAllFrames" directly.
func GetAllFrames(details GetAllFramesArgDetails) (ret js.Promise[js.Array[GetAllFramesReturnTypeElem]]) {
	bindings.CallGetAllFrames(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryGetAllFrames calls the function "WEBEXT.webNavigation.getAllFrames"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAllFrames(details GetAllFramesArgDetails) (ret js.Promise[js.Array[GetAllFramesReturnTypeElem]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAllFrames(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncGetFrame returns true if the function "WEBEXT.webNavigation.getFrame" exists.
func HasFuncGetFrame() bool {
	return js.True == bindings.HasFuncGetFrame()
}

// FuncGetFrame returns the function "WEBEXT.webNavigation.getFrame".
func FuncGetFrame() (fn js.Func[func(details GetFrameArgDetails) js.Promise[GetFrameReturnType]]) {
	bindings.FuncGetFrame(
		js.Pointer(&fn),
	)
	return
}

// GetFrame calls the function "WEBEXT.webNavigation.getFrame" directly.
func GetFrame(details GetFrameArgDetails) (ret js.Promise[GetFrameReturnType]) {
	bindings.CallGetFrame(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryGetFrame calls the function "WEBEXT.webNavigation.getFrame"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetFrame(details GetFrameArgDetails) (ret js.Promise[GetFrameReturnType], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetFrame(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

type OnBeforeNavigateEventCallbackFunc func(this js.Ref, details *OnBeforeNavigateArgDetails) js.Ref

func (fn OnBeforeNavigateEventCallbackFunc) Register() js.Func[func(details *OnBeforeNavigateArgDetails)] {
	return js.RegisterCallback[func(details *OnBeforeNavigateArgDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnBeforeNavigateEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnBeforeNavigateArgDetails
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

type OnBeforeNavigateEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnBeforeNavigateArgDetails) js.Ref
	Arg T
}

func (cb *OnBeforeNavigateEventCallback[T]) Register() js.Func[func(details *OnBeforeNavigateArgDetails)] {
	return js.RegisterCallback[func(details *OnBeforeNavigateArgDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnBeforeNavigateEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnBeforeNavigateArgDetails
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

// HasFuncOnBeforeNavigate returns true if the function "WEBEXT.webNavigation.onBeforeNavigate.addListener" exists.
func HasFuncOnBeforeNavigate() bool {
	return js.True == bindings.HasFuncOnBeforeNavigate()
}

// FuncOnBeforeNavigate returns the function "WEBEXT.webNavigation.onBeforeNavigate.addListener".
func FuncOnBeforeNavigate() (fn js.Func[func(callback js.Func[func(details *OnBeforeNavigateArgDetails)])]) {
	bindings.FuncOnBeforeNavigate(
		js.Pointer(&fn),
	)
	return
}

// OnBeforeNavigate calls the function "WEBEXT.webNavigation.onBeforeNavigate.addListener" directly.
func OnBeforeNavigate(callback js.Func[func(details *OnBeforeNavigateArgDetails)]) (ret js.Void) {
	bindings.CallOnBeforeNavigate(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnBeforeNavigate calls the function "WEBEXT.webNavigation.onBeforeNavigate.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnBeforeNavigate(callback js.Func[func(details *OnBeforeNavigateArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnBeforeNavigate(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffBeforeNavigate returns true if the function "WEBEXT.webNavigation.onBeforeNavigate.removeListener" exists.
func HasFuncOffBeforeNavigate() bool {
	return js.True == bindings.HasFuncOffBeforeNavigate()
}

// FuncOffBeforeNavigate returns the function "WEBEXT.webNavigation.onBeforeNavigate.removeListener".
func FuncOffBeforeNavigate() (fn js.Func[func(callback js.Func[func(details *OnBeforeNavigateArgDetails)])]) {
	bindings.FuncOffBeforeNavigate(
		js.Pointer(&fn),
	)
	return
}

// OffBeforeNavigate calls the function "WEBEXT.webNavigation.onBeforeNavigate.removeListener" directly.
func OffBeforeNavigate(callback js.Func[func(details *OnBeforeNavigateArgDetails)]) (ret js.Void) {
	bindings.CallOffBeforeNavigate(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffBeforeNavigate calls the function "WEBEXT.webNavigation.onBeforeNavigate.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffBeforeNavigate(callback js.Func[func(details *OnBeforeNavigateArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffBeforeNavigate(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnBeforeNavigate returns true if the function "WEBEXT.webNavigation.onBeforeNavigate.hasListener" exists.
func HasFuncHasOnBeforeNavigate() bool {
	return js.True == bindings.HasFuncHasOnBeforeNavigate()
}

// FuncHasOnBeforeNavigate returns the function "WEBEXT.webNavigation.onBeforeNavigate.hasListener".
func FuncHasOnBeforeNavigate() (fn js.Func[func(callback js.Func[func(details *OnBeforeNavigateArgDetails)]) bool]) {
	bindings.FuncHasOnBeforeNavigate(
		js.Pointer(&fn),
	)
	return
}

// HasOnBeforeNavigate calls the function "WEBEXT.webNavigation.onBeforeNavigate.hasListener" directly.
func HasOnBeforeNavigate(callback js.Func[func(details *OnBeforeNavigateArgDetails)]) (ret bool) {
	bindings.CallHasOnBeforeNavigate(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnBeforeNavigate calls the function "WEBEXT.webNavigation.onBeforeNavigate.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnBeforeNavigate(callback js.Func[func(details *OnBeforeNavigateArgDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnBeforeNavigate(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnCommittedEventCallbackFunc func(this js.Ref, details *OnCommittedArgDetails) js.Ref

func (fn OnCommittedEventCallbackFunc) Register() js.Func[func(details *OnCommittedArgDetails)] {
	return js.RegisterCallback[func(details *OnCommittedArgDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCommittedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnCommittedArgDetails
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

type OnCommittedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnCommittedArgDetails) js.Ref
	Arg T
}

func (cb *OnCommittedEventCallback[T]) Register() js.Func[func(details *OnCommittedArgDetails)] {
	return js.RegisterCallback[func(details *OnCommittedArgDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCommittedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnCommittedArgDetails
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

// HasFuncOnCommitted returns true if the function "WEBEXT.webNavigation.onCommitted.addListener" exists.
func HasFuncOnCommitted() bool {
	return js.True == bindings.HasFuncOnCommitted()
}

// FuncOnCommitted returns the function "WEBEXT.webNavigation.onCommitted.addListener".
func FuncOnCommitted() (fn js.Func[func(callback js.Func[func(details *OnCommittedArgDetails)])]) {
	bindings.FuncOnCommitted(
		js.Pointer(&fn),
	)
	return
}

// OnCommitted calls the function "WEBEXT.webNavigation.onCommitted.addListener" directly.
func OnCommitted(callback js.Func[func(details *OnCommittedArgDetails)]) (ret js.Void) {
	bindings.CallOnCommitted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCommitted calls the function "WEBEXT.webNavigation.onCommitted.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCommitted(callback js.Func[func(details *OnCommittedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCommitted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCommitted returns true if the function "WEBEXT.webNavigation.onCommitted.removeListener" exists.
func HasFuncOffCommitted() bool {
	return js.True == bindings.HasFuncOffCommitted()
}

// FuncOffCommitted returns the function "WEBEXT.webNavigation.onCommitted.removeListener".
func FuncOffCommitted() (fn js.Func[func(callback js.Func[func(details *OnCommittedArgDetails)])]) {
	bindings.FuncOffCommitted(
		js.Pointer(&fn),
	)
	return
}

// OffCommitted calls the function "WEBEXT.webNavigation.onCommitted.removeListener" directly.
func OffCommitted(callback js.Func[func(details *OnCommittedArgDetails)]) (ret js.Void) {
	bindings.CallOffCommitted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCommitted calls the function "WEBEXT.webNavigation.onCommitted.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCommitted(callback js.Func[func(details *OnCommittedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCommitted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCommitted returns true if the function "WEBEXT.webNavigation.onCommitted.hasListener" exists.
func HasFuncHasOnCommitted() bool {
	return js.True == bindings.HasFuncHasOnCommitted()
}

// FuncHasOnCommitted returns the function "WEBEXT.webNavigation.onCommitted.hasListener".
func FuncHasOnCommitted() (fn js.Func[func(callback js.Func[func(details *OnCommittedArgDetails)]) bool]) {
	bindings.FuncHasOnCommitted(
		js.Pointer(&fn),
	)
	return
}

// HasOnCommitted calls the function "WEBEXT.webNavigation.onCommitted.hasListener" directly.
func HasOnCommitted(callback js.Func[func(details *OnCommittedArgDetails)]) (ret bool) {
	bindings.CallHasOnCommitted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCommitted calls the function "WEBEXT.webNavigation.onCommitted.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCommitted(callback js.Func[func(details *OnCommittedArgDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCommitted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnCompletedEventCallbackFunc func(this js.Ref, details *OnCompletedArgDetails) js.Ref

func (fn OnCompletedEventCallbackFunc) Register() js.Func[func(details *OnCompletedArgDetails)] {
	return js.RegisterCallback[func(details *OnCompletedArgDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCompletedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnCompletedArgDetails
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

type OnCompletedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnCompletedArgDetails) js.Ref
	Arg T
}

func (cb *OnCompletedEventCallback[T]) Register() js.Func[func(details *OnCompletedArgDetails)] {
	return js.RegisterCallback[func(details *OnCompletedArgDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCompletedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnCompletedArgDetails
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

// HasFuncOnCompleted returns true if the function "WEBEXT.webNavigation.onCompleted.addListener" exists.
func HasFuncOnCompleted() bool {
	return js.True == bindings.HasFuncOnCompleted()
}

// FuncOnCompleted returns the function "WEBEXT.webNavigation.onCompleted.addListener".
func FuncOnCompleted() (fn js.Func[func(callback js.Func[func(details *OnCompletedArgDetails)])]) {
	bindings.FuncOnCompleted(
		js.Pointer(&fn),
	)
	return
}

// OnCompleted calls the function "WEBEXT.webNavigation.onCompleted.addListener" directly.
func OnCompleted(callback js.Func[func(details *OnCompletedArgDetails)]) (ret js.Void) {
	bindings.CallOnCompleted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCompleted calls the function "WEBEXT.webNavigation.onCompleted.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCompleted(callback js.Func[func(details *OnCompletedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCompleted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCompleted returns true if the function "WEBEXT.webNavigation.onCompleted.removeListener" exists.
func HasFuncOffCompleted() bool {
	return js.True == bindings.HasFuncOffCompleted()
}

// FuncOffCompleted returns the function "WEBEXT.webNavigation.onCompleted.removeListener".
func FuncOffCompleted() (fn js.Func[func(callback js.Func[func(details *OnCompletedArgDetails)])]) {
	bindings.FuncOffCompleted(
		js.Pointer(&fn),
	)
	return
}

// OffCompleted calls the function "WEBEXT.webNavigation.onCompleted.removeListener" directly.
func OffCompleted(callback js.Func[func(details *OnCompletedArgDetails)]) (ret js.Void) {
	bindings.CallOffCompleted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCompleted calls the function "WEBEXT.webNavigation.onCompleted.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCompleted(callback js.Func[func(details *OnCompletedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCompleted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCompleted returns true if the function "WEBEXT.webNavigation.onCompleted.hasListener" exists.
func HasFuncHasOnCompleted() bool {
	return js.True == bindings.HasFuncHasOnCompleted()
}

// FuncHasOnCompleted returns the function "WEBEXT.webNavigation.onCompleted.hasListener".
func FuncHasOnCompleted() (fn js.Func[func(callback js.Func[func(details *OnCompletedArgDetails)]) bool]) {
	bindings.FuncHasOnCompleted(
		js.Pointer(&fn),
	)
	return
}

// HasOnCompleted calls the function "WEBEXT.webNavigation.onCompleted.hasListener" directly.
func HasOnCompleted(callback js.Func[func(details *OnCompletedArgDetails)]) (ret bool) {
	bindings.CallHasOnCompleted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCompleted calls the function "WEBEXT.webNavigation.onCompleted.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCompleted(callback js.Func[func(details *OnCompletedArgDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCompleted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnCreatedNavigationTargetEventCallbackFunc func(this js.Ref, details *OnCreatedNavigationTargetArgDetails) js.Ref

func (fn OnCreatedNavigationTargetEventCallbackFunc) Register() js.Func[func(details *OnCreatedNavigationTargetArgDetails)] {
	return js.RegisterCallback[func(details *OnCreatedNavigationTargetArgDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCreatedNavigationTargetEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnCreatedNavigationTargetArgDetails
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

type OnCreatedNavigationTargetEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnCreatedNavigationTargetArgDetails) js.Ref
	Arg T
}

func (cb *OnCreatedNavigationTargetEventCallback[T]) Register() js.Func[func(details *OnCreatedNavigationTargetArgDetails)] {
	return js.RegisterCallback[func(details *OnCreatedNavigationTargetArgDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCreatedNavigationTargetEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnCreatedNavigationTargetArgDetails
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

// HasFuncOnCreatedNavigationTarget returns true if the function "WEBEXT.webNavigation.onCreatedNavigationTarget.addListener" exists.
func HasFuncOnCreatedNavigationTarget() bool {
	return js.True == bindings.HasFuncOnCreatedNavigationTarget()
}

// FuncOnCreatedNavigationTarget returns the function "WEBEXT.webNavigation.onCreatedNavigationTarget.addListener".
func FuncOnCreatedNavigationTarget() (fn js.Func[func(callback js.Func[func(details *OnCreatedNavigationTargetArgDetails)])]) {
	bindings.FuncOnCreatedNavigationTarget(
		js.Pointer(&fn),
	)
	return
}

// OnCreatedNavigationTarget calls the function "WEBEXT.webNavigation.onCreatedNavigationTarget.addListener" directly.
func OnCreatedNavigationTarget(callback js.Func[func(details *OnCreatedNavigationTargetArgDetails)]) (ret js.Void) {
	bindings.CallOnCreatedNavigationTarget(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCreatedNavigationTarget calls the function "WEBEXT.webNavigation.onCreatedNavigationTarget.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCreatedNavigationTarget(callback js.Func[func(details *OnCreatedNavigationTargetArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCreatedNavigationTarget(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCreatedNavigationTarget returns true if the function "WEBEXT.webNavigation.onCreatedNavigationTarget.removeListener" exists.
func HasFuncOffCreatedNavigationTarget() bool {
	return js.True == bindings.HasFuncOffCreatedNavigationTarget()
}

// FuncOffCreatedNavigationTarget returns the function "WEBEXT.webNavigation.onCreatedNavigationTarget.removeListener".
func FuncOffCreatedNavigationTarget() (fn js.Func[func(callback js.Func[func(details *OnCreatedNavigationTargetArgDetails)])]) {
	bindings.FuncOffCreatedNavigationTarget(
		js.Pointer(&fn),
	)
	return
}

// OffCreatedNavigationTarget calls the function "WEBEXT.webNavigation.onCreatedNavigationTarget.removeListener" directly.
func OffCreatedNavigationTarget(callback js.Func[func(details *OnCreatedNavigationTargetArgDetails)]) (ret js.Void) {
	bindings.CallOffCreatedNavigationTarget(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCreatedNavigationTarget calls the function "WEBEXT.webNavigation.onCreatedNavigationTarget.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCreatedNavigationTarget(callback js.Func[func(details *OnCreatedNavigationTargetArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCreatedNavigationTarget(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCreatedNavigationTarget returns true if the function "WEBEXT.webNavigation.onCreatedNavigationTarget.hasListener" exists.
func HasFuncHasOnCreatedNavigationTarget() bool {
	return js.True == bindings.HasFuncHasOnCreatedNavigationTarget()
}

// FuncHasOnCreatedNavigationTarget returns the function "WEBEXT.webNavigation.onCreatedNavigationTarget.hasListener".
func FuncHasOnCreatedNavigationTarget() (fn js.Func[func(callback js.Func[func(details *OnCreatedNavigationTargetArgDetails)]) bool]) {
	bindings.FuncHasOnCreatedNavigationTarget(
		js.Pointer(&fn),
	)
	return
}

// HasOnCreatedNavigationTarget calls the function "WEBEXT.webNavigation.onCreatedNavigationTarget.hasListener" directly.
func HasOnCreatedNavigationTarget(callback js.Func[func(details *OnCreatedNavigationTargetArgDetails)]) (ret bool) {
	bindings.CallHasOnCreatedNavigationTarget(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCreatedNavigationTarget calls the function "WEBEXT.webNavigation.onCreatedNavigationTarget.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCreatedNavigationTarget(callback js.Func[func(details *OnCreatedNavigationTargetArgDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCreatedNavigationTarget(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDOMContentLoadedEventCallbackFunc func(this js.Ref, details *OnDOMContentLoadedArgDetails) js.Ref

func (fn OnDOMContentLoadedEventCallbackFunc) Register() js.Func[func(details *OnDOMContentLoadedArgDetails)] {
	return js.RegisterCallback[func(details *OnDOMContentLoadedArgDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDOMContentLoadedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnDOMContentLoadedArgDetails
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

type OnDOMContentLoadedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnDOMContentLoadedArgDetails) js.Ref
	Arg T
}

func (cb *OnDOMContentLoadedEventCallback[T]) Register() js.Func[func(details *OnDOMContentLoadedArgDetails)] {
	return js.RegisterCallback[func(details *OnDOMContentLoadedArgDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDOMContentLoadedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnDOMContentLoadedArgDetails
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

// HasFuncOnDOMContentLoaded returns true if the function "WEBEXT.webNavigation.onDOMContentLoaded.addListener" exists.
func HasFuncOnDOMContentLoaded() bool {
	return js.True == bindings.HasFuncOnDOMContentLoaded()
}

// FuncOnDOMContentLoaded returns the function "WEBEXT.webNavigation.onDOMContentLoaded.addListener".
func FuncOnDOMContentLoaded() (fn js.Func[func(callback js.Func[func(details *OnDOMContentLoadedArgDetails)])]) {
	bindings.FuncOnDOMContentLoaded(
		js.Pointer(&fn),
	)
	return
}

// OnDOMContentLoaded calls the function "WEBEXT.webNavigation.onDOMContentLoaded.addListener" directly.
func OnDOMContentLoaded(callback js.Func[func(details *OnDOMContentLoadedArgDetails)]) (ret js.Void) {
	bindings.CallOnDOMContentLoaded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDOMContentLoaded calls the function "WEBEXT.webNavigation.onDOMContentLoaded.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDOMContentLoaded(callback js.Func[func(details *OnDOMContentLoadedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDOMContentLoaded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDOMContentLoaded returns true if the function "WEBEXT.webNavigation.onDOMContentLoaded.removeListener" exists.
func HasFuncOffDOMContentLoaded() bool {
	return js.True == bindings.HasFuncOffDOMContentLoaded()
}

// FuncOffDOMContentLoaded returns the function "WEBEXT.webNavigation.onDOMContentLoaded.removeListener".
func FuncOffDOMContentLoaded() (fn js.Func[func(callback js.Func[func(details *OnDOMContentLoadedArgDetails)])]) {
	bindings.FuncOffDOMContentLoaded(
		js.Pointer(&fn),
	)
	return
}

// OffDOMContentLoaded calls the function "WEBEXT.webNavigation.onDOMContentLoaded.removeListener" directly.
func OffDOMContentLoaded(callback js.Func[func(details *OnDOMContentLoadedArgDetails)]) (ret js.Void) {
	bindings.CallOffDOMContentLoaded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDOMContentLoaded calls the function "WEBEXT.webNavigation.onDOMContentLoaded.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDOMContentLoaded(callback js.Func[func(details *OnDOMContentLoadedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDOMContentLoaded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDOMContentLoaded returns true if the function "WEBEXT.webNavigation.onDOMContentLoaded.hasListener" exists.
func HasFuncHasOnDOMContentLoaded() bool {
	return js.True == bindings.HasFuncHasOnDOMContentLoaded()
}

// FuncHasOnDOMContentLoaded returns the function "WEBEXT.webNavigation.onDOMContentLoaded.hasListener".
func FuncHasOnDOMContentLoaded() (fn js.Func[func(callback js.Func[func(details *OnDOMContentLoadedArgDetails)]) bool]) {
	bindings.FuncHasOnDOMContentLoaded(
		js.Pointer(&fn),
	)
	return
}

// HasOnDOMContentLoaded calls the function "WEBEXT.webNavigation.onDOMContentLoaded.hasListener" directly.
func HasOnDOMContentLoaded(callback js.Func[func(details *OnDOMContentLoadedArgDetails)]) (ret bool) {
	bindings.CallHasOnDOMContentLoaded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDOMContentLoaded calls the function "WEBEXT.webNavigation.onDOMContentLoaded.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDOMContentLoaded(callback js.Func[func(details *OnDOMContentLoadedArgDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDOMContentLoaded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnErrorOccurredEventCallbackFunc func(this js.Ref, details *OnErrorOccurredArgDetails) js.Ref

func (fn OnErrorOccurredEventCallbackFunc) Register() js.Func[func(details *OnErrorOccurredArgDetails)] {
	return js.RegisterCallback[func(details *OnErrorOccurredArgDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnErrorOccurredEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnErrorOccurredArgDetails
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

type OnErrorOccurredEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnErrorOccurredArgDetails) js.Ref
	Arg T
}

func (cb *OnErrorOccurredEventCallback[T]) Register() js.Func[func(details *OnErrorOccurredArgDetails)] {
	return js.RegisterCallback[func(details *OnErrorOccurredArgDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnErrorOccurredEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnErrorOccurredArgDetails
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

// HasFuncOnErrorOccurred returns true if the function "WEBEXT.webNavigation.onErrorOccurred.addListener" exists.
func HasFuncOnErrorOccurred() bool {
	return js.True == bindings.HasFuncOnErrorOccurred()
}

// FuncOnErrorOccurred returns the function "WEBEXT.webNavigation.onErrorOccurred.addListener".
func FuncOnErrorOccurred() (fn js.Func[func(callback js.Func[func(details *OnErrorOccurredArgDetails)])]) {
	bindings.FuncOnErrorOccurred(
		js.Pointer(&fn),
	)
	return
}

// OnErrorOccurred calls the function "WEBEXT.webNavigation.onErrorOccurred.addListener" directly.
func OnErrorOccurred(callback js.Func[func(details *OnErrorOccurredArgDetails)]) (ret js.Void) {
	bindings.CallOnErrorOccurred(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnErrorOccurred calls the function "WEBEXT.webNavigation.onErrorOccurred.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnErrorOccurred(callback js.Func[func(details *OnErrorOccurredArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnErrorOccurred(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffErrorOccurred returns true if the function "WEBEXT.webNavigation.onErrorOccurred.removeListener" exists.
func HasFuncOffErrorOccurred() bool {
	return js.True == bindings.HasFuncOffErrorOccurred()
}

// FuncOffErrorOccurred returns the function "WEBEXT.webNavigation.onErrorOccurred.removeListener".
func FuncOffErrorOccurred() (fn js.Func[func(callback js.Func[func(details *OnErrorOccurredArgDetails)])]) {
	bindings.FuncOffErrorOccurred(
		js.Pointer(&fn),
	)
	return
}

// OffErrorOccurred calls the function "WEBEXT.webNavigation.onErrorOccurred.removeListener" directly.
func OffErrorOccurred(callback js.Func[func(details *OnErrorOccurredArgDetails)]) (ret js.Void) {
	bindings.CallOffErrorOccurred(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffErrorOccurred calls the function "WEBEXT.webNavigation.onErrorOccurred.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffErrorOccurred(callback js.Func[func(details *OnErrorOccurredArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffErrorOccurred(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnErrorOccurred returns true if the function "WEBEXT.webNavigation.onErrorOccurred.hasListener" exists.
func HasFuncHasOnErrorOccurred() bool {
	return js.True == bindings.HasFuncHasOnErrorOccurred()
}

// FuncHasOnErrorOccurred returns the function "WEBEXT.webNavigation.onErrorOccurred.hasListener".
func FuncHasOnErrorOccurred() (fn js.Func[func(callback js.Func[func(details *OnErrorOccurredArgDetails)]) bool]) {
	bindings.FuncHasOnErrorOccurred(
		js.Pointer(&fn),
	)
	return
}

// HasOnErrorOccurred calls the function "WEBEXT.webNavigation.onErrorOccurred.hasListener" directly.
func HasOnErrorOccurred(callback js.Func[func(details *OnErrorOccurredArgDetails)]) (ret bool) {
	bindings.CallHasOnErrorOccurred(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnErrorOccurred calls the function "WEBEXT.webNavigation.onErrorOccurred.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnErrorOccurred(callback js.Func[func(details *OnErrorOccurredArgDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnErrorOccurred(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnHistoryStateUpdatedEventCallbackFunc func(this js.Ref, details *OnHistoryStateUpdatedArgDetails) js.Ref

func (fn OnHistoryStateUpdatedEventCallbackFunc) Register() js.Func[func(details *OnHistoryStateUpdatedArgDetails)] {
	return js.RegisterCallback[func(details *OnHistoryStateUpdatedArgDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnHistoryStateUpdatedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnHistoryStateUpdatedArgDetails
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

type OnHistoryStateUpdatedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnHistoryStateUpdatedArgDetails) js.Ref
	Arg T
}

func (cb *OnHistoryStateUpdatedEventCallback[T]) Register() js.Func[func(details *OnHistoryStateUpdatedArgDetails)] {
	return js.RegisterCallback[func(details *OnHistoryStateUpdatedArgDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnHistoryStateUpdatedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnHistoryStateUpdatedArgDetails
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

// HasFuncOnHistoryStateUpdated returns true if the function "WEBEXT.webNavigation.onHistoryStateUpdated.addListener" exists.
func HasFuncOnHistoryStateUpdated() bool {
	return js.True == bindings.HasFuncOnHistoryStateUpdated()
}

// FuncOnHistoryStateUpdated returns the function "WEBEXT.webNavigation.onHistoryStateUpdated.addListener".
func FuncOnHistoryStateUpdated() (fn js.Func[func(callback js.Func[func(details *OnHistoryStateUpdatedArgDetails)])]) {
	bindings.FuncOnHistoryStateUpdated(
		js.Pointer(&fn),
	)
	return
}

// OnHistoryStateUpdated calls the function "WEBEXT.webNavigation.onHistoryStateUpdated.addListener" directly.
func OnHistoryStateUpdated(callback js.Func[func(details *OnHistoryStateUpdatedArgDetails)]) (ret js.Void) {
	bindings.CallOnHistoryStateUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnHistoryStateUpdated calls the function "WEBEXT.webNavigation.onHistoryStateUpdated.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnHistoryStateUpdated(callback js.Func[func(details *OnHistoryStateUpdatedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnHistoryStateUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffHistoryStateUpdated returns true if the function "WEBEXT.webNavigation.onHistoryStateUpdated.removeListener" exists.
func HasFuncOffHistoryStateUpdated() bool {
	return js.True == bindings.HasFuncOffHistoryStateUpdated()
}

// FuncOffHistoryStateUpdated returns the function "WEBEXT.webNavigation.onHistoryStateUpdated.removeListener".
func FuncOffHistoryStateUpdated() (fn js.Func[func(callback js.Func[func(details *OnHistoryStateUpdatedArgDetails)])]) {
	bindings.FuncOffHistoryStateUpdated(
		js.Pointer(&fn),
	)
	return
}

// OffHistoryStateUpdated calls the function "WEBEXT.webNavigation.onHistoryStateUpdated.removeListener" directly.
func OffHistoryStateUpdated(callback js.Func[func(details *OnHistoryStateUpdatedArgDetails)]) (ret js.Void) {
	bindings.CallOffHistoryStateUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffHistoryStateUpdated calls the function "WEBEXT.webNavigation.onHistoryStateUpdated.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffHistoryStateUpdated(callback js.Func[func(details *OnHistoryStateUpdatedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffHistoryStateUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnHistoryStateUpdated returns true if the function "WEBEXT.webNavigation.onHistoryStateUpdated.hasListener" exists.
func HasFuncHasOnHistoryStateUpdated() bool {
	return js.True == bindings.HasFuncHasOnHistoryStateUpdated()
}

// FuncHasOnHistoryStateUpdated returns the function "WEBEXT.webNavigation.onHistoryStateUpdated.hasListener".
func FuncHasOnHistoryStateUpdated() (fn js.Func[func(callback js.Func[func(details *OnHistoryStateUpdatedArgDetails)]) bool]) {
	bindings.FuncHasOnHistoryStateUpdated(
		js.Pointer(&fn),
	)
	return
}

// HasOnHistoryStateUpdated calls the function "WEBEXT.webNavigation.onHistoryStateUpdated.hasListener" directly.
func HasOnHistoryStateUpdated(callback js.Func[func(details *OnHistoryStateUpdatedArgDetails)]) (ret bool) {
	bindings.CallHasOnHistoryStateUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnHistoryStateUpdated calls the function "WEBEXT.webNavigation.onHistoryStateUpdated.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnHistoryStateUpdated(callback js.Func[func(details *OnHistoryStateUpdatedArgDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnHistoryStateUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnReferenceFragmentUpdatedEventCallbackFunc func(this js.Ref, details *OnReferenceFragmentUpdatedArgDetails) js.Ref

func (fn OnReferenceFragmentUpdatedEventCallbackFunc) Register() js.Func[func(details *OnReferenceFragmentUpdatedArgDetails)] {
	return js.RegisterCallback[func(details *OnReferenceFragmentUpdatedArgDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnReferenceFragmentUpdatedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnReferenceFragmentUpdatedArgDetails
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

type OnReferenceFragmentUpdatedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnReferenceFragmentUpdatedArgDetails) js.Ref
	Arg T
}

func (cb *OnReferenceFragmentUpdatedEventCallback[T]) Register() js.Func[func(details *OnReferenceFragmentUpdatedArgDetails)] {
	return js.RegisterCallback[func(details *OnReferenceFragmentUpdatedArgDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnReferenceFragmentUpdatedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnReferenceFragmentUpdatedArgDetails
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

// HasFuncOnReferenceFragmentUpdated returns true if the function "WEBEXT.webNavigation.onReferenceFragmentUpdated.addListener" exists.
func HasFuncOnReferenceFragmentUpdated() bool {
	return js.True == bindings.HasFuncOnReferenceFragmentUpdated()
}

// FuncOnReferenceFragmentUpdated returns the function "WEBEXT.webNavigation.onReferenceFragmentUpdated.addListener".
func FuncOnReferenceFragmentUpdated() (fn js.Func[func(callback js.Func[func(details *OnReferenceFragmentUpdatedArgDetails)])]) {
	bindings.FuncOnReferenceFragmentUpdated(
		js.Pointer(&fn),
	)
	return
}

// OnReferenceFragmentUpdated calls the function "WEBEXT.webNavigation.onReferenceFragmentUpdated.addListener" directly.
func OnReferenceFragmentUpdated(callback js.Func[func(details *OnReferenceFragmentUpdatedArgDetails)]) (ret js.Void) {
	bindings.CallOnReferenceFragmentUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnReferenceFragmentUpdated calls the function "WEBEXT.webNavigation.onReferenceFragmentUpdated.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnReferenceFragmentUpdated(callback js.Func[func(details *OnReferenceFragmentUpdatedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnReferenceFragmentUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffReferenceFragmentUpdated returns true if the function "WEBEXT.webNavigation.onReferenceFragmentUpdated.removeListener" exists.
func HasFuncOffReferenceFragmentUpdated() bool {
	return js.True == bindings.HasFuncOffReferenceFragmentUpdated()
}

// FuncOffReferenceFragmentUpdated returns the function "WEBEXT.webNavigation.onReferenceFragmentUpdated.removeListener".
func FuncOffReferenceFragmentUpdated() (fn js.Func[func(callback js.Func[func(details *OnReferenceFragmentUpdatedArgDetails)])]) {
	bindings.FuncOffReferenceFragmentUpdated(
		js.Pointer(&fn),
	)
	return
}

// OffReferenceFragmentUpdated calls the function "WEBEXT.webNavigation.onReferenceFragmentUpdated.removeListener" directly.
func OffReferenceFragmentUpdated(callback js.Func[func(details *OnReferenceFragmentUpdatedArgDetails)]) (ret js.Void) {
	bindings.CallOffReferenceFragmentUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffReferenceFragmentUpdated calls the function "WEBEXT.webNavigation.onReferenceFragmentUpdated.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffReferenceFragmentUpdated(callback js.Func[func(details *OnReferenceFragmentUpdatedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffReferenceFragmentUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnReferenceFragmentUpdated returns true if the function "WEBEXT.webNavigation.onReferenceFragmentUpdated.hasListener" exists.
func HasFuncHasOnReferenceFragmentUpdated() bool {
	return js.True == bindings.HasFuncHasOnReferenceFragmentUpdated()
}

// FuncHasOnReferenceFragmentUpdated returns the function "WEBEXT.webNavigation.onReferenceFragmentUpdated.hasListener".
func FuncHasOnReferenceFragmentUpdated() (fn js.Func[func(callback js.Func[func(details *OnReferenceFragmentUpdatedArgDetails)]) bool]) {
	bindings.FuncHasOnReferenceFragmentUpdated(
		js.Pointer(&fn),
	)
	return
}

// HasOnReferenceFragmentUpdated calls the function "WEBEXT.webNavigation.onReferenceFragmentUpdated.hasListener" directly.
func HasOnReferenceFragmentUpdated(callback js.Func[func(details *OnReferenceFragmentUpdatedArgDetails)]) (ret bool) {
	bindings.CallHasOnReferenceFragmentUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnReferenceFragmentUpdated calls the function "WEBEXT.webNavigation.onReferenceFragmentUpdated.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnReferenceFragmentUpdated(callback js.Func[func(details *OnReferenceFragmentUpdatedArgDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnReferenceFragmentUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnTabReplacedEventCallbackFunc func(this js.Ref, details *OnTabReplacedArgDetails) js.Ref

func (fn OnTabReplacedEventCallbackFunc) Register() js.Func[func(details *OnTabReplacedArgDetails)] {
	return js.RegisterCallback[func(details *OnTabReplacedArgDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnTabReplacedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnTabReplacedArgDetails
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

type OnTabReplacedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnTabReplacedArgDetails) js.Ref
	Arg T
}

func (cb *OnTabReplacedEventCallback[T]) Register() js.Func[func(details *OnTabReplacedArgDetails)] {
	return js.RegisterCallback[func(details *OnTabReplacedArgDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnTabReplacedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnTabReplacedArgDetails
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

// HasFuncOnTabReplaced returns true if the function "WEBEXT.webNavigation.onTabReplaced.addListener" exists.
func HasFuncOnTabReplaced() bool {
	return js.True == bindings.HasFuncOnTabReplaced()
}

// FuncOnTabReplaced returns the function "WEBEXT.webNavigation.onTabReplaced.addListener".
func FuncOnTabReplaced() (fn js.Func[func(callback js.Func[func(details *OnTabReplacedArgDetails)])]) {
	bindings.FuncOnTabReplaced(
		js.Pointer(&fn),
	)
	return
}

// OnTabReplaced calls the function "WEBEXT.webNavigation.onTabReplaced.addListener" directly.
func OnTabReplaced(callback js.Func[func(details *OnTabReplacedArgDetails)]) (ret js.Void) {
	bindings.CallOnTabReplaced(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnTabReplaced calls the function "WEBEXT.webNavigation.onTabReplaced.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnTabReplaced(callback js.Func[func(details *OnTabReplacedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnTabReplaced(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffTabReplaced returns true if the function "WEBEXT.webNavigation.onTabReplaced.removeListener" exists.
func HasFuncOffTabReplaced() bool {
	return js.True == bindings.HasFuncOffTabReplaced()
}

// FuncOffTabReplaced returns the function "WEBEXT.webNavigation.onTabReplaced.removeListener".
func FuncOffTabReplaced() (fn js.Func[func(callback js.Func[func(details *OnTabReplacedArgDetails)])]) {
	bindings.FuncOffTabReplaced(
		js.Pointer(&fn),
	)
	return
}

// OffTabReplaced calls the function "WEBEXT.webNavigation.onTabReplaced.removeListener" directly.
func OffTabReplaced(callback js.Func[func(details *OnTabReplacedArgDetails)]) (ret js.Void) {
	bindings.CallOffTabReplaced(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffTabReplaced calls the function "WEBEXT.webNavigation.onTabReplaced.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffTabReplaced(callback js.Func[func(details *OnTabReplacedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffTabReplaced(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnTabReplaced returns true if the function "WEBEXT.webNavigation.onTabReplaced.hasListener" exists.
func HasFuncHasOnTabReplaced() bool {
	return js.True == bindings.HasFuncHasOnTabReplaced()
}

// FuncHasOnTabReplaced returns the function "WEBEXT.webNavigation.onTabReplaced.hasListener".
func FuncHasOnTabReplaced() (fn js.Func[func(callback js.Func[func(details *OnTabReplacedArgDetails)]) bool]) {
	bindings.FuncHasOnTabReplaced(
		js.Pointer(&fn),
	)
	return
}

// HasOnTabReplaced calls the function "WEBEXT.webNavigation.onTabReplaced.hasListener" directly.
func HasOnTabReplaced(callback js.Func[func(details *OnTabReplacedArgDetails)]) (ret bool) {
	bindings.CallHasOnTabReplaced(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnTabReplaced calls the function "WEBEXT.webNavigation.onTabReplaced.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnTabReplaced(callback js.Func[func(details *OnTabReplacedArgDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnTabReplaced(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}
