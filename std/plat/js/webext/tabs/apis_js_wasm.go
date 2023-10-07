// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package tabs

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/extensiontypes"
	"github.com/primecitizens/pcz/std/plat/js/webext/tabs/bindings"
)

type ConnectArgConnectInfo struct {
	// DocumentId is "ConnectArgConnectInfo.documentId"
	//
	// Optional
	DocumentId js.String
	// FrameId is "ConnectArgConnectInfo.frameId"
	//
	// Optional
	//
	// NOTE: FFI_USE_FrameId MUST be set to true to make this field effective.
	FrameId int64
	// Name is "ConnectArgConnectInfo.name"
	//
	// Optional
	Name js.String

	FFI_USE_FrameId bool // for FrameId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ConnectArgConnectInfo with all fields set.
func (p ConnectArgConnectInfo) FromRef(ref js.Ref) ConnectArgConnectInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ConnectArgConnectInfo in the application heap.
func (p ConnectArgConnectInfo) New() js.Ref {
	return bindings.ConnectArgConnectInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ConnectArgConnectInfo) UpdateFrom(ref js.Ref) {
	bindings.ConnectArgConnectInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ConnectArgConnectInfo) Update(ref js.Ref) {
	bindings.ConnectArgConnectInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ConnectArgConnectInfo) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentId.Ref(),
		p.Name.Ref(),
	)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
}

type MutedInfoReason uint32

const (
	_ MutedInfoReason = iota

	MutedInfoReason_USER
	MutedInfoReason_CAPTURE
	MutedInfoReason_EXTENSION
)

func (MutedInfoReason) FromRef(str js.Ref) MutedInfoReason {
	return MutedInfoReason(bindings.ConstOfMutedInfoReason(str))
}

func (x MutedInfoReason) String() (string, bool) {
	switch x {
	case MutedInfoReason_USER:
		return "user", true
	case MutedInfoReason_CAPTURE:
		return "capture", true
	case MutedInfoReason_EXTENSION:
		return "extension", true
	default:
		return "", false
	}
}

type MutedInfo struct {
	// ExtensionId is "MutedInfo.extensionId"
	//
	// Optional
	ExtensionId js.String
	// Muted is "MutedInfo.muted"
	//
	// Required
	Muted bool
	// Reason is "MutedInfo.reason"
	//
	// Optional
	Reason MutedInfoReason

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MutedInfo with all fields set.
func (p MutedInfo) FromRef(ref js.Ref) MutedInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MutedInfo in the application heap.
func (p MutedInfo) New() js.Ref {
	return bindings.MutedInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MutedInfo) UpdateFrom(ref js.Ref) {
	bindings.MutedInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MutedInfo) Update(ref js.Ref) {
	bindings.MutedInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MutedInfo) FreeMembers(recursive bool) {
	js.Free(
		p.ExtensionId.Ref(),
	)
	p.ExtensionId = p.ExtensionId.FromRef(js.Undefined)
}

type TabStatus uint32

const (
	_ TabStatus = iota

	TabStatus_UNLOADED
	TabStatus_LOADING
	TabStatus_COMPLETE
)

func (TabStatus) FromRef(str js.Ref) TabStatus {
	return TabStatus(bindings.ConstOfTabStatus(str))
}

func (x TabStatus) String() (string, bool) {
	switch x {
	case TabStatus_UNLOADED:
		return "unloaded", true
	case TabStatus_LOADING:
		return "loading", true
	case TabStatus_COMPLETE:
		return "complete", true
	default:
		return "", false
	}
}

type Tab struct {
	// Active is "Tab.active"
	//
	// Required
	Active bool
	// Audible is "Tab.audible"
	//
	// Optional
	//
	// NOTE: FFI_USE_Audible MUST be set to true to make this field effective.
	Audible bool
	// AutoDiscardable is "Tab.autoDiscardable"
	//
	// Required
	AutoDiscardable bool
	// Discarded is "Tab.discarded"
	//
	// Required
	Discarded bool
	// FavIconUrl is "Tab.favIconUrl"
	//
	// Optional
	FavIconUrl js.String
	// GroupId is "Tab.groupId"
	//
	// Required
	GroupId int64
	// Height is "Tab.height"
	//
	// Optional
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height int64
	// Highlighted is "Tab.highlighted"
	//
	// Required
	Highlighted bool
	// Id is "Tab.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id int64
	// Incognito is "Tab.incognito"
	//
	// Required
	Incognito bool
	// Index is "Tab.index"
	//
	// Required
	Index int64
	// MutedInfo is "Tab.mutedInfo"
	//
	// Optional
	//
	// NOTE: MutedInfo.FFI_USE MUST be set to true to get MutedInfo used.
	MutedInfo MutedInfo
	// OpenerTabId is "Tab.openerTabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_OpenerTabId MUST be set to true to make this field effective.
	OpenerTabId int64
	// PendingUrl is "Tab.pendingUrl"
	//
	// Optional
	PendingUrl js.String
	// Pinned is "Tab.pinned"
	//
	// Required
	Pinned bool
	// Selected is "Tab.selected"
	//
	// Required
	Selected bool
	// SessionId is "Tab.sessionId"
	//
	// Optional
	SessionId js.String
	// Status is "Tab.status"
	//
	// Optional
	Status TabStatus
	// Title is "Tab.title"
	//
	// Optional
	Title js.String
	// Url is "Tab.url"
	//
	// Optional
	Url js.String
	// Width is "Tab.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width int64
	// WindowId is "Tab.windowId"
	//
	// Required
	WindowId int64

	FFI_USE_Audible     bool // for Audible.
	FFI_USE_Height      bool // for Height.
	FFI_USE_Id          bool // for Id.
	FFI_USE_OpenerTabId bool // for OpenerTabId.
	FFI_USE_Width       bool // for Width.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Tab with all fields set.
func (p Tab) FromRef(ref js.Ref) Tab {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Tab in the application heap.
func (p Tab) New() js.Ref {
	return bindings.TabJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Tab) UpdateFrom(ref js.Ref) {
	bindings.TabJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Tab) Update(ref js.Ref) {
	bindings.TabJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Tab) FreeMembers(recursive bool) {
	js.Free(
		p.FavIconUrl.Ref(),
		p.PendingUrl.Ref(),
		p.SessionId.Ref(),
		p.Title.Ref(),
		p.Url.Ref(),
	)
	p.FavIconUrl = p.FavIconUrl.FromRef(js.Undefined)
	p.PendingUrl = p.PendingUrl.FromRef(js.Undefined)
	p.SessionId = p.SessionId.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
	if recursive {
		p.MutedInfo.FreeMembers(true)
	}
}

type CreateArgCreateProperties struct {
	// Active is "CreateArgCreateProperties.active"
	//
	// Optional
	//
	// NOTE: FFI_USE_Active MUST be set to true to make this field effective.
	Active bool
	// Index is "CreateArgCreateProperties.index"
	//
	// Optional
	//
	// NOTE: FFI_USE_Index MUST be set to true to make this field effective.
	Index int64
	// OpenerTabId is "CreateArgCreateProperties.openerTabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_OpenerTabId MUST be set to true to make this field effective.
	OpenerTabId int64
	// Pinned is "CreateArgCreateProperties.pinned"
	//
	// Optional
	//
	// NOTE: FFI_USE_Pinned MUST be set to true to make this field effective.
	Pinned bool
	// Selected is "CreateArgCreateProperties.selected"
	//
	// Optional
	//
	// NOTE: FFI_USE_Selected MUST be set to true to make this field effective.
	Selected bool
	// Url is "CreateArgCreateProperties.url"
	//
	// Optional
	Url js.String
	// WindowId is "CreateArgCreateProperties.windowId"
	//
	// Optional
	//
	// NOTE: FFI_USE_WindowId MUST be set to true to make this field effective.
	WindowId int64

	FFI_USE_Active      bool // for Active.
	FFI_USE_Index       bool // for Index.
	FFI_USE_OpenerTabId bool // for OpenerTabId.
	FFI_USE_Pinned      bool // for Pinned.
	FFI_USE_Selected    bool // for Selected.
	FFI_USE_WindowId    bool // for WindowId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CreateArgCreateProperties with all fields set.
func (p CreateArgCreateProperties) FromRef(ref js.Ref) CreateArgCreateProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CreateArgCreateProperties in the application heap.
func (p CreateArgCreateProperties) New() js.Ref {
	return bindings.CreateArgCreatePropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CreateArgCreateProperties) UpdateFrom(ref js.Ref) {
	bindings.CreateArgCreatePropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CreateArgCreateProperties) Update(ref js.Ref) {
	bindings.CreateArgCreatePropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CreateArgCreateProperties) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
}

type GroupArgOptionsFieldCreateProperties struct {
	// WindowId is "GroupArgOptionsFieldCreateProperties.windowId"
	//
	// Optional
	//
	// NOTE: FFI_USE_WindowId MUST be set to true to make this field effective.
	WindowId int64

	FFI_USE_WindowId bool // for WindowId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GroupArgOptionsFieldCreateProperties with all fields set.
func (p GroupArgOptionsFieldCreateProperties) FromRef(ref js.Ref) GroupArgOptionsFieldCreateProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GroupArgOptionsFieldCreateProperties in the application heap.
func (p GroupArgOptionsFieldCreateProperties) New() js.Ref {
	return bindings.GroupArgOptionsFieldCreatePropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GroupArgOptionsFieldCreateProperties) UpdateFrom(ref js.Ref) {
	bindings.GroupArgOptionsFieldCreatePropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GroupArgOptionsFieldCreateProperties) Update(ref js.Ref) {
	bindings.GroupArgOptionsFieldCreatePropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GroupArgOptionsFieldCreateProperties) FreeMembers(recursive bool) {
}

type OneOf_Int64_ArrayInt64 struct {
	ref js.Ref
}

func (x OneOf_Int64_ArrayInt64) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Int64_ArrayInt64) Free() {
	x.ref.Free()
}

func (x OneOf_Int64_ArrayInt64) FromRef(ref js.Ref) OneOf_Int64_ArrayInt64 {
	return OneOf_Int64_ArrayInt64{
		ref: ref,
	}
}

func (x OneOf_Int64_ArrayInt64) Int64() int64 {
	return js.BigInt[int64]{}.FromRef(x.ref).Get()
}

func (x OneOf_Int64_ArrayInt64) ArrayInt64() js.Array[int64] {
	return js.Array[int64]{}.FromRef(x.ref)
}

type GroupArgOptions struct {
	// CreateProperties is "GroupArgOptions.createProperties"
	//
	// Optional
	//
	// NOTE: CreateProperties.FFI_USE MUST be set to true to get CreateProperties used.
	CreateProperties GroupArgOptionsFieldCreateProperties
	// GroupId is "GroupArgOptions.groupId"
	//
	// Optional
	//
	// NOTE: FFI_USE_GroupId MUST be set to true to make this field effective.
	GroupId int64
	// TabIds is "GroupArgOptions.tabIds"
	//
	// Required
	TabIds OneOf_Int64_ArrayInt64

	FFI_USE_GroupId bool // for GroupId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GroupArgOptions with all fields set.
func (p GroupArgOptions) FromRef(ref js.Ref) GroupArgOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GroupArgOptions in the application heap.
func (p GroupArgOptions) New() js.Ref {
	return bindings.GroupArgOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GroupArgOptions) UpdateFrom(ref js.Ref) {
	bindings.GroupArgOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GroupArgOptions) Update(ref js.Ref) {
	bindings.GroupArgOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GroupArgOptions) FreeMembers(recursive bool) {
	js.Free(
		p.TabIds.Ref(),
	)
	p.TabIds = p.TabIds.FromRef(js.Undefined)
	if recursive {
		p.CreateProperties.FreeMembers(true)
	}
}

type OneOf_ArrayInt64_Int64 struct {
	ref js.Ref
}

func (x OneOf_ArrayInt64_Int64) Ref() js.Ref {
	return x.ref
}

func (x OneOf_ArrayInt64_Int64) Free() {
	x.ref.Free()
}

func (x OneOf_ArrayInt64_Int64) FromRef(ref js.Ref) OneOf_ArrayInt64_Int64 {
	return OneOf_ArrayInt64_Int64{
		ref: ref,
	}
}

func (x OneOf_ArrayInt64_Int64) ArrayInt64() js.Array[int64] {
	return js.Array[int64]{}.FromRef(x.ref)
}

func (x OneOf_ArrayInt64_Int64) Int64() int64 {
	return js.BigInt[int64]{}.FromRef(x.ref).Get()
}

type HighlightArgHighlightInfo struct {
	// Tabs is "HighlightArgHighlightInfo.tabs"
	//
	// Required
	Tabs OneOf_ArrayInt64_Int64
	// WindowId is "HighlightArgHighlightInfo.windowId"
	//
	// Optional
	//
	// NOTE: FFI_USE_WindowId MUST be set to true to make this field effective.
	WindowId int64

	FFI_USE_WindowId bool // for WindowId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HighlightArgHighlightInfo with all fields set.
func (p HighlightArgHighlightInfo) FromRef(ref js.Ref) HighlightArgHighlightInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HighlightArgHighlightInfo in the application heap.
func (p HighlightArgHighlightInfo) New() js.Ref {
	return bindings.HighlightArgHighlightInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HighlightArgHighlightInfo) UpdateFrom(ref js.Ref) {
	bindings.HighlightArgHighlightInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HighlightArgHighlightInfo) Update(ref js.Ref) {
	bindings.HighlightArgHighlightInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HighlightArgHighlightInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Tabs.Ref(),
	)
	p.Tabs = p.Tabs.FromRef(js.Undefined)
}

// MAX_CAPTURE_VISIBLE_TAB_CALLS_PER_SECOND returns the value of property "WEBEXT.tabs.MAX_CAPTURE_VISIBLE_TAB_CALLS_PER_SECOND".
//
// The returned bool will be false if there is no such property.
func MAX_CAPTURE_VISIBLE_TAB_CALLS_PER_SECOND() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMAX_CAPTURE_VISIBLE_TAB_CALLS_PER_SECOND(
		js.Pointer(&ret),
	)

	return
}

// SetMAX_CAPTURE_VISIBLE_TAB_CALLS_PER_SECOND sets the value of property "WEBEXT.tabs.MAX_CAPTURE_VISIBLE_TAB_CALLS_PER_SECOND" to val.
//
// It returns false if the property cannot be set.
func SetMAX_CAPTURE_VISIBLE_TAB_CALLS_PER_SECOND(val js.String) bool {
	return js.True == bindings.SetMAX_CAPTURE_VISIBLE_TAB_CALLS_PER_SECOND(
		val.Ref())
}

type MoveArgMoveProperties struct {
	// Index is "MoveArgMoveProperties.index"
	//
	// Required
	Index int64
	// WindowId is "MoveArgMoveProperties.windowId"
	//
	// Optional
	//
	// NOTE: FFI_USE_WindowId MUST be set to true to make this field effective.
	WindowId int64

	FFI_USE_WindowId bool // for WindowId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MoveArgMoveProperties with all fields set.
func (p MoveArgMoveProperties) FromRef(ref js.Ref) MoveArgMoveProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MoveArgMoveProperties in the application heap.
func (p MoveArgMoveProperties) New() js.Ref {
	return bindings.MoveArgMovePropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MoveArgMoveProperties) UpdateFrom(ref js.Ref) {
	bindings.MoveArgMovePropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MoveArgMoveProperties) Update(ref js.Ref) {
	bindings.MoveArgMovePropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MoveArgMoveProperties) FreeMembers(recursive bool) {
}

type OnActivatedArgActiveInfo struct {
	// TabId is "OnActivatedArgActiveInfo.tabId"
	//
	// Required
	TabId int64
	// WindowId is "OnActivatedArgActiveInfo.windowId"
	//
	// Required
	WindowId int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnActivatedArgActiveInfo with all fields set.
func (p OnActivatedArgActiveInfo) FromRef(ref js.Ref) OnActivatedArgActiveInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnActivatedArgActiveInfo in the application heap.
func (p OnActivatedArgActiveInfo) New() js.Ref {
	return bindings.OnActivatedArgActiveInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnActivatedArgActiveInfo) UpdateFrom(ref js.Ref) {
	bindings.OnActivatedArgActiveInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnActivatedArgActiveInfo) Update(ref js.Ref) {
	bindings.OnActivatedArgActiveInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnActivatedArgActiveInfo) FreeMembers(recursive bool) {
}

type OnActiveChangedArgSelectInfo struct {
	// WindowId is "OnActiveChangedArgSelectInfo.windowId"
	//
	// Required
	WindowId int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnActiveChangedArgSelectInfo with all fields set.
func (p OnActiveChangedArgSelectInfo) FromRef(ref js.Ref) OnActiveChangedArgSelectInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnActiveChangedArgSelectInfo in the application heap.
func (p OnActiveChangedArgSelectInfo) New() js.Ref {
	return bindings.OnActiveChangedArgSelectInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnActiveChangedArgSelectInfo) UpdateFrom(ref js.Ref) {
	bindings.OnActiveChangedArgSelectInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnActiveChangedArgSelectInfo) Update(ref js.Ref) {
	bindings.OnActiveChangedArgSelectInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnActiveChangedArgSelectInfo) FreeMembers(recursive bool) {
}

type OnAttachedArgAttachInfo struct {
	// NewPosition is "OnAttachedArgAttachInfo.newPosition"
	//
	// Required
	NewPosition int64
	// NewWindowId is "OnAttachedArgAttachInfo.newWindowId"
	//
	// Required
	NewWindowId int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnAttachedArgAttachInfo with all fields set.
func (p OnAttachedArgAttachInfo) FromRef(ref js.Ref) OnAttachedArgAttachInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnAttachedArgAttachInfo in the application heap.
func (p OnAttachedArgAttachInfo) New() js.Ref {
	return bindings.OnAttachedArgAttachInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnAttachedArgAttachInfo) UpdateFrom(ref js.Ref) {
	bindings.OnAttachedArgAttachInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnAttachedArgAttachInfo) Update(ref js.Ref) {
	bindings.OnAttachedArgAttachInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnAttachedArgAttachInfo) FreeMembers(recursive bool) {
}

type OnDetachedArgDetachInfo struct {
	// OldPosition is "OnDetachedArgDetachInfo.oldPosition"
	//
	// Required
	OldPosition int64
	// OldWindowId is "OnDetachedArgDetachInfo.oldWindowId"
	//
	// Required
	OldWindowId int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnDetachedArgDetachInfo with all fields set.
func (p OnDetachedArgDetachInfo) FromRef(ref js.Ref) OnDetachedArgDetachInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnDetachedArgDetachInfo in the application heap.
func (p OnDetachedArgDetachInfo) New() js.Ref {
	return bindings.OnDetachedArgDetachInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnDetachedArgDetachInfo) UpdateFrom(ref js.Ref) {
	bindings.OnDetachedArgDetachInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnDetachedArgDetachInfo) Update(ref js.Ref) {
	bindings.OnDetachedArgDetachInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnDetachedArgDetachInfo) FreeMembers(recursive bool) {
}

type OnHighlightChangedArgSelectInfo struct {
	// TabIds is "OnHighlightChangedArgSelectInfo.tabIds"
	//
	// Required
	TabIds js.Array[int64]
	// WindowId is "OnHighlightChangedArgSelectInfo.windowId"
	//
	// Required
	WindowId int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnHighlightChangedArgSelectInfo with all fields set.
func (p OnHighlightChangedArgSelectInfo) FromRef(ref js.Ref) OnHighlightChangedArgSelectInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnHighlightChangedArgSelectInfo in the application heap.
func (p OnHighlightChangedArgSelectInfo) New() js.Ref {
	return bindings.OnHighlightChangedArgSelectInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnHighlightChangedArgSelectInfo) UpdateFrom(ref js.Ref) {
	bindings.OnHighlightChangedArgSelectInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnHighlightChangedArgSelectInfo) Update(ref js.Ref) {
	bindings.OnHighlightChangedArgSelectInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnHighlightChangedArgSelectInfo) FreeMembers(recursive bool) {
	js.Free(
		p.TabIds.Ref(),
	)
	p.TabIds = p.TabIds.FromRef(js.Undefined)
}

type OnHighlightedArgHighlightInfo struct {
	// TabIds is "OnHighlightedArgHighlightInfo.tabIds"
	//
	// Required
	TabIds js.Array[int64]
	// WindowId is "OnHighlightedArgHighlightInfo.windowId"
	//
	// Required
	WindowId int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnHighlightedArgHighlightInfo with all fields set.
func (p OnHighlightedArgHighlightInfo) FromRef(ref js.Ref) OnHighlightedArgHighlightInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnHighlightedArgHighlightInfo in the application heap.
func (p OnHighlightedArgHighlightInfo) New() js.Ref {
	return bindings.OnHighlightedArgHighlightInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnHighlightedArgHighlightInfo) UpdateFrom(ref js.Ref) {
	bindings.OnHighlightedArgHighlightInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnHighlightedArgHighlightInfo) Update(ref js.Ref) {
	bindings.OnHighlightedArgHighlightInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnHighlightedArgHighlightInfo) FreeMembers(recursive bool) {
	js.Free(
		p.TabIds.Ref(),
	)
	p.TabIds = p.TabIds.FromRef(js.Undefined)
}

type OnMovedArgMoveInfo struct {
	// FromIndex is "OnMovedArgMoveInfo.fromIndex"
	//
	// Required
	FromIndex int64
	// ToIndex is "OnMovedArgMoveInfo.toIndex"
	//
	// Required
	ToIndex int64
	// WindowId is "OnMovedArgMoveInfo.windowId"
	//
	// Required
	WindowId int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnMovedArgMoveInfo with all fields set.
func (p OnMovedArgMoveInfo) FromRef(ref js.Ref) OnMovedArgMoveInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnMovedArgMoveInfo in the application heap.
func (p OnMovedArgMoveInfo) New() js.Ref {
	return bindings.OnMovedArgMoveInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnMovedArgMoveInfo) UpdateFrom(ref js.Ref) {
	bindings.OnMovedArgMoveInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnMovedArgMoveInfo) Update(ref js.Ref) {
	bindings.OnMovedArgMoveInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnMovedArgMoveInfo) FreeMembers(recursive bool) {
}

type OnRemovedArgRemoveInfo struct {
	// IsWindowClosing is "OnRemovedArgRemoveInfo.isWindowClosing"
	//
	// Required
	IsWindowClosing bool
	// WindowId is "OnRemovedArgRemoveInfo.windowId"
	//
	// Required
	WindowId int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnRemovedArgRemoveInfo with all fields set.
func (p OnRemovedArgRemoveInfo) FromRef(ref js.Ref) OnRemovedArgRemoveInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnRemovedArgRemoveInfo in the application heap.
func (p OnRemovedArgRemoveInfo) New() js.Ref {
	return bindings.OnRemovedArgRemoveInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnRemovedArgRemoveInfo) UpdateFrom(ref js.Ref) {
	bindings.OnRemovedArgRemoveInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnRemovedArgRemoveInfo) Update(ref js.Ref) {
	bindings.OnRemovedArgRemoveInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnRemovedArgRemoveInfo) FreeMembers(recursive bool) {
}

type OnSelectionChangedArgSelectInfo struct {
	// WindowId is "OnSelectionChangedArgSelectInfo.windowId"
	//
	// Required
	WindowId int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnSelectionChangedArgSelectInfo with all fields set.
func (p OnSelectionChangedArgSelectInfo) FromRef(ref js.Ref) OnSelectionChangedArgSelectInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnSelectionChangedArgSelectInfo in the application heap.
func (p OnSelectionChangedArgSelectInfo) New() js.Ref {
	return bindings.OnSelectionChangedArgSelectInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnSelectionChangedArgSelectInfo) UpdateFrom(ref js.Ref) {
	bindings.OnSelectionChangedArgSelectInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnSelectionChangedArgSelectInfo) Update(ref js.Ref) {
	bindings.OnSelectionChangedArgSelectInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnSelectionChangedArgSelectInfo) FreeMembers(recursive bool) {
}

type OnUpdatedArgChangeInfo struct {
	// Audible is "OnUpdatedArgChangeInfo.audible"
	//
	// Optional
	//
	// NOTE: FFI_USE_Audible MUST be set to true to make this field effective.
	Audible bool
	// AutoDiscardable is "OnUpdatedArgChangeInfo.autoDiscardable"
	//
	// Optional
	//
	// NOTE: FFI_USE_AutoDiscardable MUST be set to true to make this field effective.
	AutoDiscardable bool
	// Discarded is "OnUpdatedArgChangeInfo.discarded"
	//
	// Optional
	//
	// NOTE: FFI_USE_Discarded MUST be set to true to make this field effective.
	Discarded bool
	// FavIconUrl is "OnUpdatedArgChangeInfo.favIconUrl"
	//
	// Optional
	FavIconUrl js.String
	// GroupId is "OnUpdatedArgChangeInfo.groupId"
	//
	// Optional
	//
	// NOTE: FFI_USE_GroupId MUST be set to true to make this field effective.
	GroupId int64
	// MutedInfo is "OnUpdatedArgChangeInfo.mutedInfo"
	//
	// Optional
	//
	// NOTE: MutedInfo.FFI_USE MUST be set to true to get MutedInfo used.
	MutedInfo MutedInfo
	// Pinned is "OnUpdatedArgChangeInfo.pinned"
	//
	// Optional
	//
	// NOTE: FFI_USE_Pinned MUST be set to true to make this field effective.
	Pinned bool
	// Status is "OnUpdatedArgChangeInfo.status"
	//
	// Optional
	Status TabStatus
	// Title is "OnUpdatedArgChangeInfo.title"
	//
	// Optional
	Title js.String
	// Url is "OnUpdatedArgChangeInfo.url"
	//
	// Optional
	Url js.String

	FFI_USE_Audible         bool // for Audible.
	FFI_USE_AutoDiscardable bool // for AutoDiscardable.
	FFI_USE_Discarded       bool // for Discarded.
	FFI_USE_GroupId         bool // for GroupId.
	FFI_USE_Pinned          bool // for Pinned.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnUpdatedArgChangeInfo with all fields set.
func (p OnUpdatedArgChangeInfo) FromRef(ref js.Ref) OnUpdatedArgChangeInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnUpdatedArgChangeInfo in the application heap.
func (p OnUpdatedArgChangeInfo) New() js.Ref {
	return bindings.OnUpdatedArgChangeInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnUpdatedArgChangeInfo) UpdateFrom(ref js.Ref) {
	bindings.OnUpdatedArgChangeInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnUpdatedArgChangeInfo) Update(ref js.Ref) {
	bindings.OnUpdatedArgChangeInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnUpdatedArgChangeInfo) FreeMembers(recursive bool) {
	js.Free(
		p.FavIconUrl.Ref(),
		p.Title.Ref(),
		p.Url.Ref(),
	)
	p.FavIconUrl = p.FavIconUrl.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
	if recursive {
		p.MutedInfo.FreeMembers(true)
	}
}

type ZoomSettingsMode uint32

const (
	_ ZoomSettingsMode = iota

	ZoomSettingsMode_AUTOMATIC
	ZoomSettingsMode_MANUAL
	ZoomSettingsMode_DISABLED
)

func (ZoomSettingsMode) FromRef(str js.Ref) ZoomSettingsMode {
	return ZoomSettingsMode(bindings.ConstOfZoomSettingsMode(str))
}

func (x ZoomSettingsMode) String() (string, bool) {
	switch x {
	case ZoomSettingsMode_AUTOMATIC:
		return "automatic", true
	case ZoomSettingsMode_MANUAL:
		return "manual", true
	case ZoomSettingsMode_DISABLED:
		return "disabled", true
	default:
		return "", false
	}
}

type ZoomSettingsScope uint32

const (
	_ ZoomSettingsScope = iota

	ZoomSettingsScope_PER_ORIGIN
	ZoomSettingsScope_PER_TAB
)

func (ZoomSettingsScope) FromRef(str js.Ref) ZoomSettingsScope {
	return ZoomSettingsScope(bindings.ConstOfZoomSettingsScope(str))
}

func (x ZoomSettingsScope) String() (string, bool) {
	switch x {
	case ZoomSettingsScope_PER_ORIGIN:
		return "per-origin", true
	case ZoomSettingsScope_PER_TAB:
		return "per-tab", true
	default:
		return "", false
	}
}

type ZoomSettings struct {
	// DefaultZoomFactor is "ZoomSettings.defaultZoomFactor"
	//
	// Optional
	//
	// NOTE: FFI_USE_DefaultZoomFactor MUST be set to true to make this field effective.
	DefaultZoomFactor float64
	// Mode is "ZoomSettings.mode"
	//
	// Optional
	Mode ZoomSettingsMode
	// Scope is "ZoomSettings.scope"
	//
	// Optional
	Scope ZoomSettingsScope

	FFI_USE_DefaultZoomFactor bool // for DefaultZoomFactor.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ZoomSettings with all fields set.
func (p ZoomSettings) FromRef(ref js.Ref) ZoomSettings {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ZoomSettings in the application heap.
func (p ZoomSettings) New() js.Ref {
	return bindings.ZoomSettingsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ZoomSettings) UpdateFrom(ref js.Ref) {
	bindings.ZoomSettingsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ZoomSettings) Update(ref js.Ref) {
	bindings.ZoomSettingsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ZoomSettings) FreeMembers(recursive bool) {
}

type OnZoomChangeArgZoomChangeInfo struct {
	// NewZoomFactor is "OnZoomChangeArgZoomChangeInfo.newZoomFactor"
	//
	// Required
	NewZoomFactor float64
	// OldZoomFactor is "OnZoomChangeArgZoomChangeInfo.oldZoomFactor"
	//
	// Required
	OldZoomFactor float64
	// TabId is "OnZoomChangeArgZoomChangeInfo.tabId"
	//
	// Required
	TabId int64
	// ZoomSettings is "OnZoomChangeArgZoomChangeInfo.zoomSettings"
	//
	// Required
	//
	// NOTE: ZoomSettings.FFI_USE MUST be set to true to get ZoomSettings used.
	ZoomSettings ZoomSettings

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnZoomChangeArgZoomChangeInfo with all fields set.
func (p OnZoomChangeArgZoomChangeInfo) FromRef(ref js.Ref) OnZoomChangeArgZoomChangeInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnZoomChangeArgZoomChangeInfo in the application heap.
func (p OnZoomChangeArgZoomChangeInfo) New() js.Ref {
	return bindings.OnZoomChangeArgZoomChangeInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnZoomChangeArgZoomChangeInfo) UpdateFrom(ref js.Ref) {
	bindings.OnZoomChangeArgZoomChangeInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnZoomChangeArgZoomChangeInfo) Update(ref js.Ref) {
	bindings.OnZoomChangeArgZoomChangeInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnZoomChangeArgZoomChangeInfo) FreeMembers(recursive bool) {
	if recursive {
		p.ZoomSettings.FreeMembers(true)
	}
}

type OneOf_String_ArrayString struct {
	ref js.Ref
}

func (x OneOf_String_ArrayString) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_ArrayString) Free() {
	x.ref.Free()
}

func (x OneOf_String_ArrayString) FromRef(ref js.Ref) OneOf_String_ArrayString {
	return OneOf_String_ArrayString{
		ref: ref,
	}
}

func (x OneOf_String_ArrayString) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayString) ArrayString() js.Array[js.String] {
	return js.Array[js.String]{}.FromRef(x.ref)
}

type WindowType uint32

const (
	_ WindowType = iota

	WindowType_NORMAL
	WindowType_POPUP
	WindowType_PANEL
	WindowType_APP
	WindowType_DEVTOOLS
)

func (WindowType) FromRef(str js.Ref) WindowType {
	return WindowType(bindings.ConstOfWindowType(str))
}

func (x WindowType) String() (string, bool) {
	switch x {
	case WindowType_NORMAL:
		return "normal", true
	case WindowType_POPUP:
		return "popup", true
	case WindowType_PANEL:
		return "panel", true
	case WindowType_APP:
		return "app", true
	case WindowType_DEVTOOLS:
		return "devtools", true
	default:
		return "", false
	}
}

type QueryArgQueryInfo struct {
	// Active is "QueryArgQueryInfo.active"
	//
	// Optional
	//
	// NOTE: FFI_USE_Active MUST be set to true to make this field effective.
	Active bool
	// Audible is "QueryArgQueryInfo.audible"
	//
	// Optional
	//
	// NOTE: FFI_USE_Audible MUST be set to true to make this field effective.
	Audible bool
	// AutoDiscardable is "QueryArgQueryInfo.autoDiscardable"
	//
	// Optional
	//
	// NOTE: FFI_USE_AutoDiscardable MUST be set to true to make this field effective.
	AutoDiscardable bool
	// CurrentWindow is "QueryArgQueryInfo.currentWindow"
	//
	// Optional
	//
	// NOTE: FFI_USE_CurrentWindow MUST be set to true to make this field effective.
	CurrentWindow bool
	// Discarded is "QueryArgQueryInfo.discarded"
	//
	// Optional
	//
	// NOTE: FFI_USE_Discarded MUST be set to true to make this field effective.
	Discarded bool
	// GroupId is "QueryArgQueryInfo.groupId"
	//
	// Optional
	//
	// NOTE: FFI_USE_GroupId MUST be set to true to make this field effective.
	GroupId int64
	// Highlighted is "QueryArgQueryInfo.highlighted"
	//
	// Optional
	//
	// NOTE: FFI_USE_Highlighted MUST be set to true to make this field effective.
	Highlighted bool
	// Index is "QueryArgQueryInfo.index"
	//
	// Optional
	//
	// NOTE: FFI_USE_Index MUST be set to true to make this field effective.
	Index int64
	// LastFocusedWindow is "QueryArgQueryInfo.lastFocusedWindow"
	//
	// Optional
	//
	// NOTE: FFI_USE_LastFocusedWindow MUST be set to true to make this field effective.
	LastFocusedWindow bool
	// Muted is "QueryArgQueryInfo.muted"
	//
	// Optional
	//
	// NOTE: FFI_USE_Muted MUST be set to true to make this field effective.
	Muted bool
	// Pinned is "QueryArgQueryInfo.pinned"
	//
	// Optional
	//
	// NOTE: FFI_USE_Pinned MUST be set to true to make this field effective.
	Pinned bool
	// Status is "QueryArgQueryInfo.status"
	//
	// Optional
	Status TabStatus
	// Title is "QueryArgQueryInfo.title"
	//
	// Optional
	Title js.String
	// Url is "QueryArgQueryInfo.url"
	//
	// Optional
	Url OneOf_String_ArrayString
	// WindowId is "QueryArgQueryInfo.windowId"
	//
	// Optional
	//
	// NOTE: FFI_USE_WindowId MUST be set to true to make this field effective.
	WindowId int64
	// WindowType is "QueryArgQueryInfo.windowType"
	//
	// Optional
	WindowType WindowType

	FFI_USE_Active            bool // for Active.
	FFI_USE_Audible           bool // for Audible.
	FFI_USE_AutoDiscardable   bool // for AutoDiscardable.
	FFI_USE_CurrentWindow     bool // for CurrentWindow.
	FFI_USE_Discarded         bool // for Discarded.
	FFI_USE_GroupId           bool // for GroupId.
	FFI_USE_Highlighted       bool // for Highlighted.
	FFI_USE_Index             bool // for Index.
	FFI_USE_LastFocusedWindow bool // for LastFocusedWindow.
	FFI_USE_Muted             bool // for Muted.
	FFI_USE_Pinned            bool // for Pinned.
	FFI_USE_WindowId          bool // for WindowId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a QueryArgQueryInfo with all fields set.
func (p QueryArgQueryInfo) FromRef(ref js.Ref) QueryArgQueryInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new QueryArgQueryInfo in the application heap.
func (p QueryArgQueryInfo) New() js.Ref {
	return bindings.QueryArgQueryInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *QueryArgQueryInfo) UpdateFrom(ref js.Ref) {
	bindings.QueryArgQueryInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *QueryArgQueryInfo) Update(ref js.Ref) {
	bindings.QueryArgQueryInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *QueryArgQueryInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Title.Ref(),
		p.Url.Ref(),
	)
	p.Title = p.Title.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type ReloadArgReloadProperties struct {
	// BypassCache is "ReloadArgReloadProperties.bypassCache"
	//
	// Optional
	//
	// NOTE: FFI_USE_BypassCache MUST be set to true to make this field effective.
	BypassCache bool

	FFI_USE_BypassCache bool // for BypassCache.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReloadArgReloadProperties with all fields set.
func (p ReloadArgReloadProperties) FromRef(ref js.Ref) ReloadArgReloadProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReloadArgReloadProperties in the application heap.
func (p ReloadArgReloadProperties) New() js.Ref {
	return bindings.ReloadArgReloadPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ReloadArgReloadProperties) UpdateFrom(ref js.Ref) {
	bindings.ReloadArgReloadPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ReloadArgReloadProperties) Update(ref js.Ref) {
	bindings.ReloadArgReloadPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ReloadArgReloadProperties) FreeMembers(recursive bool) {
}

type SendMessageArgOptions struct {
	// DocumentId is "SendMessageArgOptions.documentId"
	//
	// Optional
	DocumentId js.String
	// FrameId is "SendMessageArgOptions.frameId"
	//
	// Optional
	//
	// NOTE: FFI_USE_FrameId MUST be set to true to make this field effective.
	FrameId int64

	FFI_USE_FrameId bool // for FrameId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SendMessageArgOptions with all fields set.
func (p SendMessageArgOptions) FromRef(ref js.Ref) SendMessageArgOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SendMessageArgOptions in the application heap.
func (p SendMessageArgOptions) New() js.Ref {
	return bindings.SendMessageArgOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SendMessageArgOptions) UpdateFrom(ref js.Ref) {
	bindings.SendMessageArgOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SendMessageArgOptions) Update(ref js.Ref) {
	bindings.SendMessageArgOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SendMessageArgOptions) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentId.Ref(),
	)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
}

// TAB_ID_NONE returns the value of property "WEBEXT.tabs.TAB_ID_NONE".
//
// The returned bool will be false if there is no such property.
func TAB_ID_NONE() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTAB_ID_NONE(
		js.Pointer(&ret),
	)

	return
}

// SetTAB_ID_NONE sets the value of property "WEBEXT.tabs.TAB_ID_NONE" to val.
//
// It returns false if the property cannot be set.
func SetTAB_ID_NONE(val js.String) bool {
	return js.True == bindings.SetTAB_ID_NONE(
		val.Ref())
}

type UpdateArgUpdateProperties struct {
	// Active is "UpdateArgUpdateProperties.active"
	//
	// Optional
	//
	// NOTE: FFI_USE_Active MUST be set to true to make this field effective.
	Active bool
	// AutoDiscardable is "UpdateArgUpdateProperties.autoDiscardable"
	//
	// Optional
	//
	// NOTE: FFI_USE_AutoDiscardable MUST be set to true to make this field effective.
	AutoDiscardable bool
	// Highlighted is "UpdateArgUpdateProperties.highlighted"
	//
	// Optional
	//
	// NOTE: FFI_USE_Highlighted MUST be set to true to make this field effective.
	Highlighted bool
	// Muted is "UpdateArgUpdateProperties.muted"
	//
	// Optional
	//
	// NOTE: FFI_USE_Muted MUST be set to true to make this field effective.
	Muted bool
	// OpenerTabId is "UpdateArgUpdateProperties.openerTabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_OpenerTabId MUST be set to true to make this field effective.
	OpenerTabId int64
	// Pinned is "UpdateArgUpdateProperties.pinned"
	//
	// Optional
	//
	// NOTE: FFI_USE_Pinned MUST be set to true to make this field effective.
	Pinned bool
	// Selected is "UpdateArgUpdateProperties.selected"
	//
	// Optional
	//
	// NOTE: FFI_USE_Selected MUST be set to true to make this field effective.
	Selected bool
	// Url is "UpdateArgUpdateProperties.url"
	//
	// Optional
	Url js.String

	FFI_USE_Active          bool // for Active.
	FFI_USE_AutoDiscardable bool // for AutoDiscardable.
	FFI_USE_Highlighted     bool // for Highlighted.
	FFI_USE_Muted           bool // for Muted.
	FFI_USE_OpenerTabId     bool // for OpenerTabId.
	FFI_USE_Pinned          bool // for Pinned.
	FFI_USE_Selected        bool // for Selected.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UpdateArgUpdateProperties with all fields set.
func (p UpdateArgUpdateProperties) FromRef(ref js.Ref) UpdateArgUpdateProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UpdateArgUpdateProperties in the application heap.
func (p UpdateArgUpdateProperties) New() js.Ref {
	return bindings.UpdateArgUpdatePropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UpdateArgUpdateProperties) UpdateFrom(ref js.Ref) {
	bindings.UpdateArgUpdatePropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UpdateArgUpdateProperties) Update(ref js.Ref) {
	bindings.UpdateArgUpdatePropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UpdateArgUpdateProperties) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
}

// HasFuncCaptureVisibleTab returns true if the function "WEBEXT.tabs.captureVisibleTab" exists.
func HasFuncCaptureVisibleTab() bool {
	return js.True == bindings.HasFuncCaptureVisibleTab()
}

// FuncCaptureVisibleTab returns the function "WEBEXT.tabs.captureVisibleTab".
func FuncCaptureVisibleTab() (fn js.Func[func(windowId int64, options extensiontypes.ImageDetails) js.Promise[js.String]]) {
	bindings.FuncCaptureVisibleTab(
		js.Pointer(&fn),
	)
	return
}

// CaptureVisibleTab calls the function "WEBEXT.tabs.captureVisibleTab" directly.
func CaptureVisibleTab(windowId int64, options extensiontypes.ImageDetails) (ret js.Promise[js.String]) {
	bindings.CallCaptureVisibleTab(
		js.Pointer(&ret),
		float64(windowId),
		js.Pointer(&options),
	)

	return
}

// TryCaptureVisibleTab calls the function "WEBEXT.tabs.captureVisibleTab"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCaptureVisibleTab(windowId int64, options extensiontypes.ImageDetails) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCaptureVisibleTab(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(windowId),
		js.Pointer(&options),
	)

	return
}

// HasFuncCreate returns true if the function "WEBEXT.tabs.create" exists.
func HasFuncCreate() bool {
	return js.True == bindings.HasFuncCreate()
}

// FuncCreate returns the function "WEBEXT.tabs.create".
func FuncCreate() (fn js.Func[func(createProperties CreateArgCreateProperties) js.Promise[Tab]]) {
	bindings.FuncCreate(
		js.Pointer(&fn),
	)
	return
}

// Create calls the function "WEBEXT.tabs.create" directly.
func Create(createProperties CreateArgCreateProperties) (ret js.Promise[Tab]) {
	bindings.CallCreate(
		js.Pointer(&ret),
		js.Pointer(&createProperties),
	)

	return
}

// TryCreate calls the function "WEBEXT.tabs.create"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCreate(createProperties CreateArgCreateProperties) (ret js.Promise[Tab], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCreate(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&createProperties),
	)

	return
}

// HasFuncDetectLanguage returns true if the function "WEBEXT.tabs.detectLanguage" exists.
func HasFuncDetectLanguage() bool {
	return js.True == bindings.HasFuncDetectLanguage()
}

// FuncDetectLanguage returns the function "WEBEXT.tabs.detectLanguage".
func FuncDetectLanguage() (fn js.Func[func(tabId int64) js.Promise[js.String]]) {
	bindings.FuncDetectLanguage(
		js.Pointer(&fn),
	)
	return
}

// DetectLanguage calls the function "WEBEXT.tabs.detectLanguage" directly.
func DetectLanguage(tabId int64) (ret js.Promise[js.String]) {
	bindings.CallDetectLanguage(
		js.Pointer(&ret),
		float64(tabId),
	)

	return
}

// TryDetectLanguage calls the function "WEBEXT.tabs.detectLanguage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDetectLanguage(tabId int64) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDetectLanguage(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(tabId),
	)

	return
}

// HasFuncDiscard returns true if the function "WEBEXT.tabs.discard" exists.
func HasFuncDiscard() bool {
	return js.True == bindings.HasFuncDiscard()
}

// FuncDiscard returns the function "WEBEXT.tabs.discard".
func FuncDiscard() (fn js.Func[func(tabId int64) js.Promise[Tab]]) {
	bindings.FuncDiscard(
		js.Pointer(&fn),
	)
	return
}

// Discard calls the function "WEBEXT.tabs.discard" directly.
func Discard(tabId int64) (ret js.Promise[Tab]) {
	bindings.CallDiscard(
		js.Pointer(&ret),
		float64(tabId),
	)

	return
}

// TryDiscard calls the function "WEBEXT.tabs.discard"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDiscard(tabId int64) (ret js.Promise[Tab], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDiscard(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(tabId),
	)

	return
}

// HasFuncDuplicate returns true if the function "WEBEXT.tabs.duplicate" exists.
func HasFuncDuplicate() bool {
	return js.True == bindings.HasFuncDuplicate()
}

// FuncDuplicate returns the function "WEBEXT.tabs.duplicate".
func FuncDuplicate() (fn js.Func[func(tabId int64) js.Promise[Tab]]) {
	bindings.FuncDuplicate(
		js.Pointer(&fn),
	)
	return
}

// Duplicate calls the function "WEBEXT.tabs.duplicate" directly.
func Duplicate(tabId int64) (ret js.Promise[Tab]) {
	bindings.CallDuplicate(
		js.Pointer(&ret),
		float64(tabId),
	)

	return
}

// TryDuplicate calls the function "WEBEXT.tabs.duplicate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDuplicate(tabId int64) (ret js.Promise[Tab], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDuplicate(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(tabId),
	)

	return
}

// HasFuncExecuteScript returns true if the function "WEBEXT.tabs.executeScript" exists.
func HasFuncExecuteScript() bool {
	return js.True == bindings.HasFuncExecuteScript()
}

// FuncExecuteScript returns the function "WEBEXT.tabs.executeScript".
func FuncExecuteScript() (fn js.Func[func(tabId int64, details extensiontypes.InjectDetails) js.Promise[js.Array[js.Any]]]) {
	bindings.FuncExecuteScript(
		js.Pointer(&fn),
	)
	return
}

// ExecuteScript calls the function "WEBEXT.tabs.executeScript" directly.
func ExecuteScript(tabId int64, details extensiontypes.InjectDetails) (ret js.Promise[js.Array[js.Any]]) {
	bindings.CallExecuteScript(
		js.Pointer(&ret),
		float64(tabId),
		js.Pointer(&details),
	)

	return
}

// TryExecuteScript calls the function "WEBEXT.tabs.executeScript"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryExecuteScript(tabId int64, details extensiontypes.InjectDetails) (ret js.Promise[js.Array[js.Any]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryExecuteScript(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(tabId),
		js.Pointer(&details),
	)

	return
}

// HasFuncGet returns true if the function "WEBEXT.tabs.get" exists.
func HasFuncGet() bool {
	return js.True == bindings.HasFuncGet()
}

// FuncGet returns the function "WEBEXT.tabs.get".
func FuncGet() (fn js.Func[func(tabId int64) js.Promise[Tab]]) {
	bindings.FuncGet(
		js.Pointer(&fn),
	)
	return
}

// Get calls the function "WEBEXT.tabs.get" directly.
func Get(tabId int64) (ret js.Promise[Tab]) {
	bindings.CallGet(
		js.Pointer(&ret),
		float64(tabId),
	)

	return
}

// TryGet calls the function "WEBEXT.tabs.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGet(tabId int64) (ret js.Promise[Tab], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGet(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(tabId),
	)

	return
}

// HasFuncGetAllInWindow returns true if the function "WEBEXT.tabs.getAllInWindow" exists.
func HasFuncGetAllInWindow() bool {
	return js.True == bindings.HasFuncGetAllInWindow()
}

// FuncGetAllInWindow returns the function "WEBEXT.tabs.getAllInWindow".
func FuncGetAllInWindow() (fn js.Func[func(windowId int64) js.Promise[js.Array[Tab]]]) {
	bindings.FuncGetAllInWindow(
		js.Pointer(&fn),
	)
	return
}

// GetAllInWindow calls the function "WEBEXT.tabs.getAllInWindow" directly.
func GetAllInWindow(windowId int64) (ret js.Promise[js.Array[Tab]]) {
	bindings.CallGetAllInWindow(
		js.Pointer(&ret),
		float64(windowId),
	)

	return
}

// TryGetAllInWindow calls the function "WEBEXT.tabs.getAllInWindow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAllInWindow(windowId int64) (ret js.Promise[js.Array[Tab]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAllInWindow(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(windowId),
	)

	return
}

// HasFuncGetCurrent returns true if the function "WEBEXT.tabs.getCurrent" exists.
func HasFuncGetCurrent() bool {
	return js.True == bindings.HasFuncGetCurrent()
}

// FuncGetCurrent returns the function "WEBEXT.tabs.getCurrent".
func FuncGetCurrent() (fn js.Func[func() js.Promise[Tab]]) {
	bindings.FuncGetCurrent(
		js.Pointer(&fn),
	)
	return
}

// GetCurrent calls the function "WEBEXT.tabs.getCurrent" directly.
func GetCurrent() (ret js.Promise[Tab]) {
	bindings.CallGetCurrent(
		js.Pointer(&ret),
	)

	return
}

// TryGetCurrent calls the function "WEBEXT.tabs.getCurrent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetCurrent() (ret js.Promise[Tab], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetCurrent(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetSelected returns true if the function "WEBEXT.tabs.getSelected" exists.
func HasFuncGetSelected() bool {
	return js.True == bindings.HasFuncGetSelected()
}

// FuncGetSelected returns the function "WEBEXT.tabs.getSelected".
func FuncGetSelected() (fn js.Func[func(windowId int64) js.Promise[Tab]]) {
	bindings.FuncGetSelected(
		js.Pointer(&fn),
	)
	return
}

// GetSelected calls the function "WEBEXT.tabs.getSelected" directly.
func GetSelected(windowId int64) (ret js.Promise[Tab]) {
	bindings.CallGetSelected(
		js.Pointer(&ret),
		float64(windowId),
	)

	return
}

// TryGetSelected calls the function "WEBEXT.tabs.getSelected"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetSelected(windowId int64) (ret js.Promise[Tab], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetSelected(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(windowId),
	)

	return
}

// HasFuncGetZoom returns true if the function "WEBEXT.tabs.getZoom" exists.
func HasFuncGetZoom() bool {
	return js.True == bindings.HasFuncGetZoom()
}

// FuncGetZoom returns the function "WEBEXT.tabs.getZoom".
func FuncGetZoom() (fn js.Func[func(tabId int64) js.Promise[js.Number[float64]]]) {
	bindings.FuncGetZoom(
		js.Pointer(&fn),
	)
	return
}

// GetZoom calls the function "WEBEXT.tabs.getZoom" directly.
func GetZoom(tabId int64) (ret js.Promise[js.Number[float64]]) {
	bindings.CallGetZoom(
		js.Pointer(&ret),
		float64(tabId),
	)

	return
}

// TryGetZoom calls the function "WEBEXT.tabs.getZoom"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetZoom(tabId int64) (ret js.Promise[js.Number[float64]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetZoom(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(tabId),
	)

	return
}

// HasFuncGetZoomSettings returns true if the function "WEBEXT.tabs.getZoomSettings" exists.
func HasFuncGetZoomSettings() bool {
	return js.True == bindings.HasFuncGetZoomSettings()
}

// FuncGetZoomSettings returns the function "WEBEXT.tabs.getZoomSettings".
func FuncGetZoomSettings() (fn js.Func[func(tabId int64) js.Promise[ZoomSettings]]) {
	bindings.FuncGetZoomSettings(
		js.Pointer(&fn),
	)
	return
}

// GetZoomSettings calls the function "WEBEXT.tabs.getZoomSettings" directly.
func GetZoomSettings(tabId int64) (ret js.Promise[ZoomSettings]) {
	bindings.CallGetZoomSettings(
		js.Pointer(&ret),
		float64(tabId),
	)

	return
}

// TryGetZoomSettings calls the function "WEBEXT.tabs.getZoomSettings"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetZoomSettings(tabId int64) (ret js.Promise[ZoomSettings], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetZoomSettings(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(tabId),
	)

	return
}

// HasFuncGoBack returns true if the function "WEBEXT.tabs.goBack" exists.
func HasFuncGoBack() bool {
	return js.True == bindings.HasFuncGoBack()
}

// FuncGoBack returns the function "WEBEXT.tabs.goBack".
func FuncGoBack() (fn js.Func[func(tabId int64) js.Promise[js.Void]]) {
	bindings.FuncGoBack(
		js.Pointer(&fn),
	)
	return
}

// GoBack calls the function "WEBEXT.tabs.goBack" directly.
func GoBack(tabId int64) (ret js.Promise[js.Void]) {
	bindings.CallGoBack(
		js.Pointer(&ret),
		float64(tabId),
	)

	return
}

// TryGoBack calls the function "WEBEXT.tabs.goBack"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGoBack(tabId int64) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGoBack(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(tabId),
	)

	return
}

// HasFuncGoForward returns true if the function "WEBEXT.tabs.goForward" exists.
func HasFuncGoForward() bool {
	return js.True == bindings.HasFuncGoForward()
}

// FuncGoForward returns the function "WEBEXT.tabs.goForward".
func FuncGoForward() (fn js.Func[func(tabId int64) js.Promise[js.Void]]) {
	bindings.FuncGoForward(
		js.Pointer(&fn),
	)
	return
}

// GoForward calls the function "WEBEXT.tabs.goForward" directly.
func GoForward(tabId int64) (ret js.Promise[js.Void]) {
	bindings.CallGoForward(
		js.Pointer(&ret),
		float64(tabId),
	)

	return
}

// TryGoForward calls the function "WEBEXT.tabs.goForward"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGoForward(tabId int64) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGoForward(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(tabId),
	)

	return
}

// HasFuncGroup returns true if the function "WEBEXT.tabs.group" exists.
func HasFuncGroup() bool {
	return js.True == bindings.HasFuncGroup()
}

// FuncGroup returns the function "WEBEXT.tabs.group".
func FuncGroup() (fn js.Func[func(options GroupArgOptions) js.Promise[js.BigInt[int64]]]) {
	bindings.FuncGroup(
		js.Pointer(&fn),
	)
	return
}

// Group calls the function "WEBEXT.tabs.group" directly.
func Group(options GroupArgOptions) (ret js.Promise[js.BigInt[int64]]) {
	bindings.CallGroup(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGroup calls the function "WEBEXT.tabs.group"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGroup(options GroupArgOptions) (ret js.Promise[js.BigInt[int64]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGroup(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncInsertCSS returns true if the function "WEBEXT.tabs.insertCSS" exists.
func HasFuncInsertCSS() bool {
	return js.True == bindings.HasFuncInsertCSS()
}

// FuncInsertCSS returns the function "WEBEXT.tabs.insertCSS".
func FuncInsertCSS() (fn js.Func[func(tabId int64, details extensiontypes.InjectDetails) js.Promise[js.Void]]) {
	bindings.FuncInsertCSS(
		js.Pointer(&fn),
	)
	return
}

// InsertCSS calls the function "WEBEXT.tabs.insertCSS" directly.
func InsertCSS(tabId int64, details extensiontypes.InjectDetails) (ret js.Promise[js.Void]) {
	bindings.CallInsertCSS(
		js.Pointer(&ret),
		float64(tabId),
		js.Pointer(&details),
	)

	return
}

// TryInsertCSS calls the function "WEBEXT.tabs.insertCSS"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryInsertCSS(tabId int64, details extensiontypes.InjectDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryInsertCSS(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(tabId),
		js.Pointer(&details),
	)

	return
}

type OneOf_Tab_ArrayTab struct {
	ref js.Ref
}

func (x OneOf_Tab_ArrayTab) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Tab_ArrayTab) Free() {
	x.ref.Free()
}

func (x OneOf_Tab_ArrayTab) FromRef(ref js.Ref) OneOf_Tab_ArrayTab {
	return OneOf_Tab_ArrayTab{
		ref: ref,
	}
}

func (x OneOf_Tab_ArrayTab) Tab() Tab {
	var ret Tab
	ret.UpdateFrom(x.ref)
	return ret
}

func (x OneOf_Tab_ArrayTab) ArrayTab() js.Array[Tab] {
	return js.Array[Tab]{}.FromRef(x.ref)
}

// HasFuncMove returns true if the function "WEBEXT.tabs.move" exists.
func HasFuncMove() bool {
	return js.True == bindings.HasFuncMove()
}

// FuncMove returns the function "WEBEXT.tabs.move".
func FuncMove() (fn js.Func[func(tabIds OneOf_Int64_ArrayInt64, moveProperties MoveArgMoveProperties) js.Promise[OneOf_Tab_ArrayTab]]) {
	bindings.FuncMove(
		js.Pointer(&fn),
	)
	return
}

// Move calls the function "WEBEXT.tabs.move" directly.
func Move(tabIds OneOf_Int64_ArrayInt64, moveProperties MoveArgMoveProperties) (ret js.Promise[OneOf_Tab_ArrayTab]) {
	bindings.CallMove(
		js.Pointer(&ret),
		tabIds.Ref(),
		js.Pointer(&moveProperties),
	)

	return
}

// TryMove calls the function "WEBEXT.tabs.move"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryMove(tabIds OneOf_Int64_ArrayInt64, moveProperties MoveArgMoveProperties) (ret js.Promise[OneOf_Tab_ArrayTab], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMove(
		js.Pointer(&ret), js.Pointer(&exception),
		tabIds.Ref(),
		js.Pointer(&moveProperties),
	)

	return
}

type OnActivatedEventCallbackFunc func(this js.Ref, activeInfo *OnActivatedArgActiveInfo) js.Ref

func (fn OnActivatedEventCallbackFunc) Register() js.Func[func(activeInfo *OnActivatedArgActiveInfo)] {
	return js.RegisterCallback[func(activeInfo *OnActivatedArgActiveInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnActivatedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnActivatedArgActiveInfo
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

type OnActivatedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, activeInfo *OnActivatedArgActiveInfo) js.Ref
	Arg T
}

func (cb *OnActivatedEventCallback[T]) Register() js.Func[func(activeInfo *OnActivatedArgActiveInfo)] {
	return js.RegisterCallback[func(activeInfo *OnActivatedArgActiveInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnActivatedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnActivatedArgActiveInfo
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

// HasFuncOnActivated returns true if the function "WEBEXT.tabs.onActivated.addListener" exists.
func HasFuncOnActivated() bool {
	return js.True == bindings.HasFuncOnActivated()
}

// FuncOnActivated returns the function "WEBEXT.tabs.onActivated.addListener".
func FuncOnActivated() (fn js.Func[func(callback js.Func[func(activeInfo *OnActivatedArgActiveInfo)])]) {
	bindings.FuncOnActivated(
		js.Pointer(&fn),
	)
	return
}

// OnActivated calls the function "WEBEXT.tabs.onActivated.addListener" directly.
func OnActivated(callback js.Func[func(activeInfo *OnActivatedArgActiveInfo)]) (ret js.Void) {
	bindings.CallOnActivated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnActivated calls the function "WEBEXT.tabs.onActivated.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnActivated(callback js.Func[func(activeInfo *OnActivatedArgActiveInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnActivated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffActivated returns true if the function "WEBEXT.tabs.onActivated.removeListener" exists.
func HasFuncOffActivated() bool {
	return js.True == bindings.HasFuncOffActivated()
}

// FuncOffActivated returns the function "WEBEXT.tabs.onActivated.removeListener".
func FuncOffActivated() (fn js.Func[func(callback js.Func[func(activeInfo *OnActivatedArgActiveInfo)])]) {
	bindings.FuncOffActivated(
		js.Pointer(&fn),
	)
	return
}

// OffActivated calls the function "WEBEXT.tabs.onActivated.removeListener" directly.
func OffActivated(callback js.Func[func(activeInfo *OnActivatedArgActiveInfo)]) (ret js.Void) {
	bindings.CallOffActivated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffActivated calls the function "WEBEXT.tabs.onActivated.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffActivated(callback js.Func[func(activeInfo *OnActivatedArgActiveInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffActivated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnActivated returns true if the function "WEBEXT.tabs.onActivated.hasListener" exists.
func HasFuncHasOnActivated() bool {
	return js.True == bindings.HasFuncHasOnActivated()
}

// FuncHasOnActivated returns the function "WEBEXT.tabs.onActivated.hasListener".
func FuncHasOnActivated() (fn js.Func[func(callback js.Func[func(activeInfo *OnActivatedArgActiveInfo)]) bool]) {
	bindings.FuncHasOnActivated(
		js.Pointer(&fn),
	)
	return
}

// HasOnActivated calls the function "WEBEXT.tabs.onActivated.hasListener" directly.
func HasOnActivated(callback js.Func[func(activeInfo *OnActivatedArgActiveInfo)]) (ret bool) {
	bindings.CallHasOnActivated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnActivated calls the function "WEBEXT.tabs.onActivated.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnActivated(callback js.Func[func(activeInfo *OnActivatedArgActiveInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnActivated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnActiveChangedEventCallbackFunc func(this js.Ref, tabId int64, selectInfo *OnActiveChangedArgSelectInfo) js.Ref

func (fn OnActiveChangedEventCallbackFunc) Register() js.Func[func(tabId int64, selectInfo *OnActiveChangedArgSelectInfo)] {
	return js.RegisterCallback[func(tabId int64, selectInfo *OnActiveChangedArgSelectInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnActiveChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 OnActiveChangedArgSelectInfo
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnActiveChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, tabId int64, selectInfo *OnActiveChangedArgSelectInfo) js.Ref
	Arg T
}

func (cb *OnActiveChangedEventCallback[T]) Register() js.Func[func(tabId int64, selectInfo *OnActiveChangedArgSelectInfo)] {
	return js.RegisterCallback[func(tabId int64, selectInfo *OnActiveChangedArgSelectInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnActiveChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 OnActiveChangedArgSelectInfo
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnActiveChanged returns true if the function "WEBEXT.tabs.onActiveChanged.addListener" exists.
func HasFuncOnActiveChanged() bool {
	return js.True == bindings.HasFuncOnActiveChanged()
}

// FuncOnActiveChanged returns the function "WEBEXT.tabs.onActiveChanged.addListener".
func FuncOnActiveChanged() (fn js.Func[func(callback js.Func[func(tabId int64, selectInfo *OnActiveChangedArgSelectInfo)])]) {
	bindings.FuncOnActiveChanged(
		js.Pointer(&fn),
	)
	return
}

// OnActiveChanged calls the function "WEBEXT.tabs.onActiveChanged.addListener" directly.
func OnActiveChanged(callback js.Func[func(tabId int64, selectInfo *OnActiveChangedArgSelectInfo)]) (ret js.Void) {
	bindings.CallOnActiveChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnActiveChanged calls the function "WEBEXT.tabs.onActiveChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnActiveChanged(callback js.Func[func(tabId int64, selectInfo *OnActiveChangedArgSelectInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnActiveChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffActiveChanged returns true if the function "WEBEXT.tabs.onActiveChanged.removeListener" exists.
func HasFuncOffActiveChanged() bool {
	return js.True == bindings.HasFuncOffActiveChanged()
}

// FuncOffActiveChanged returns the function "WEBEXT.tabs.onActiveChanged.removeListener".
func FuncOffActiveChanged() (fn js.Func[func(callback js.Func[func(tabId int64, selectInfo *OnActiveChangedArgSelectInfo)])]) {
	bindings.FuncOffActiveChanged(
		js.Pointer(&fn),
	)
	return
}

// OffActiveChanged calls the function "WEBEXT.tabs.onActiveChanged.removeListener" directly.
func OffActiveChanged(callback js.Func[func(tabId int64, selectInfo *OnActiveChangedArgSelectInfo)]) (ret js.Void) {
	bindings.CallOffActiveChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffActiveChanged calls the function "WEBEXT.tabs.onActiveChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffActiveChanged(callback js.Func[func(tabId int64, selectInfo *OnActiveChangedArgSelectInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffActiveChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnActiveChanged returns true if the function "WEBEXT.tabs.onActiveChanged.hasListener" exists.
func HasFuncHasOnActiveChanged() bool {
	return js.True == bindings.HasFuncHasOnActiveChanged()
}

// FuncHasOnActiveChanged returns the function "WEBEXT.tabs.onActiveChanged.hasListener".
func FuncHasOnActiveChanged() (fn js.Func[func(callback js.Func[func(tabId int64, selectInfo *OnActiveChangedArgSelectInfo)]) bool]) {
	bindings.FuncHasOnActiveChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnActiveChanged calls the function "WEBEXT.tabs.onActiveChanged.hasListener" directly.
func HasOnActiveChanged(callback js.Func[func(tabId int64, selectInfo *OnActiveChangedArgSelectInfo)]) (ret bool) {
	bindings.CallHasOnActiveChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnActiveChanged calls the function "WEBEXT.tabs.onActiveChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnActiveChanged(callback js.Func[func(tabId int64, selectInfo *OnActiveChangedArgSelectInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnActiveChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnAttachedEventCallbackFunc func(this js.Ref, tabId int64, attachInfo *OnAttachedArgAttachInfo) js.Ref

func (fn OnAttachedEventCallbackFunc) Register() js.Func[func(tabId int64, attachInfo *OnAttachedArgAttachInfo)] {
	return js.RegisterCallback[func(tabId int64, attachInfo *OnAttachedArgAttachInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnAttachedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 OnAttachedArgAttachInfo
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnAttachedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, tabId int64, attachInfo *OnAttachedArgAttachInfo) js.Ref
	Arg T
}

func (cb *OnAttachedEventCallback[T]) Register() js.Func[func(tabId int64, attachInfo *OnAttachedArgAttachInfo)] {
	return js.RegisterCallback[func(tabId int64, attachInfo *OnAttachedArgAttachInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnAttachedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 OnAttachedArgAttachInfo
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnAttached returns true if the function "WEBEXT.tabs.onAttached.addListener" exists.
func HasFuncOnAttached() bool {
	return js.True == bindings.HasFuncOnAttached()
}

// FuncOnAttached returns the function "WEBEXT.tabs.onAttached.addListener".
func FuncOnAttached() (fn js.Func[func(callback js.Func[func(tabId int64, attachInfo *OnAttachedArgAttachInfo)])]) {
	bindings.FuncOnAttached(
		js.Pointer(&fn),
	)
	return
}

// OnAttached calls the function "WEBEXT.tabs.onAttached.addListener" directly.
func OnAttached(callback js.Func[func(tabId int64, attachInfo *OnAttachedArgAttachInfo)]) (ret js.Void) {
	bindings.CallOnAttached(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAttached calls the function "WEBEXT.tabs.onAttached.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAttached(callback js.Func[func(tabId int64, attachInfo *OnAttachedArgAttachInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAttached(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAttached returns true if the function "WEBEXT.tabs.onAttached.removeListener" exists.
func HasFuncOffAttached() bool {
	return js.True == bindings.HasFuncOffAttached()
}

// FuncOffAttached returns the function "WEBEXT.tabs.onAttached.removeListener".
func FuncOffAttached() (fn js.Func[func(callback js.Func[func(tabId int64, attachInfo *OnAttachedArgAttachInfo)])]) {
	bindings.FuncOffAttached(
		js.Pointer(&fn),
	)
	return
}

// OffAttached calls the function "WEBEXT.tabs.onAttached.removeListener" directly.
func OffAttached(callback js.Func[func(tabId int64, attachInfo *OnAttachedArgAttachInfo)]) (ret js.Void) {
	bindings.CallOffAttached(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAttached calls the function "WEBEXT.tabs.onAttached.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAttached(callback js.Func[func(tabId int64, attachInfo *OnAttachedArgAttachInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAttached(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAttached returns true if the function "WEBEXT.tabs.onAttached.hasListener" exists.
func HasFuncHasOnAttached() bool {
	return js.True == bindings.HasFuncHasOnAttached()
}

// FuncHasOnAttached returns the function "WEBEXT.tabs.onAttached.hasListener".
func FuncHasOnAttached() (fn js.Func[func(callback js.Func[func(tabId int64, attachInfo *OnAttachedArgAttachInfo)]) bool]) {
	bindings.FuncHasOnAttached(
		js.Pointer(&fn),
	)
	return
}

// HasOnAttached calls the function "WEBEXT.tabs.onAttached.hasListener" directly.
func HasOnAttached(callback js.Func[func(tabId int64, attachInfo *OnAttachedArgAttachInfo)]) (ret bool) {
	bindings.CallHasOnAttached(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAttached calls the function "WEBEXT.tabs.onAttached.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnAttached(callback js.Func[func(tabId int64, attachInfo *OnAttachedArgAttachInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnAttached(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnCreatedEventCallbackFunc func(this js.Ref, tab *Tab) js.Ref

func (fn OnCreatedEventCallbackFunc) Register() js.Func[func(tab *Tab)] {
	return js.RegisterCallback[func(tab *Tab)](
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
	var arg0 Tab
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
	Fn  func(arg T, this js.Ref, tab *Tab) js.Ref
	Arg T
}

func (cb *OnCreatedEventCallback[T]) Register() js.Func[func(tab *Tab)] {
	return js.RegisterCallback[func(tab *Tab)](
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
	var arg0 Tab
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

// HasFuncOnCreated returns true if the function "WEBEXT.tabs.onCreated.addListener" exists.
func HasFuncOnCreated() bool {
	return js.True == bindings.HasFuncOnCreated()
}

// FuncOnCreated returns the function "WEBEXT.tabs.onCreated.addListener".
func FuncOnCreated() (fn js.Func[func(callback js.Func[func(tab *Tab)])]) {
	bindings.FuncOnCreated(
		js.Pointer(&fn),
	)
	return
}

// OnCreated calls the function "WEBEXT.tabs.onCreated.addListener" directly.
func OnCreated(callback js.Func[func(tab *Tab)]) (ret js.Void) {
	bindings.CallOnCreated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCreated calls the function "WEBEXT.tabs.onCreated.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCreated(callback js.Func[func(tab *Tab)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCreated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCreated returns true if the function "WEBEXT.tabs.onCreated.removeListener" exists.
func HasFuncOffCreated() bool {
	return js.True == bindings.HasFuncOffCreated()
}

// FuncOffCreated returns the function "WEBEXT.tabs.onCreated.removeListener".
func FuncOffCreated() (fn js.Func[func(callback js.Func[func(tab *Tab)])]) {
	bindings.FuncOffCreated(
		js.Pointer(&fn),
	)
	return
}

// OffCreated calls the function "WEBEXT.tabs.onCreated.removeListener" directly.
func OffCreated(callback js.Func[func(tab *Tab)]) (ret js.Void) {
	bindings.CallOffCreated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCreated calls the function "WEBEXT.tabs.onCreated.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCreated(callback js.Func[func(tab *Tab)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCreated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCreated returns true if the function "WEBEXT.tabs.onCreated.hasListener" exists.
func HasFuncHasOnCreated() bool {
	return js.True == bindings.HasFuncHasOnCreated()
}

// FuncHasOnCreated returns the function "WEBEXT.tabs.onCreated.hasListener".
func FuncHasOnCreated() (fn js.Func[func(callback js.Func[func(tab *Tab)]) bool]) {
	bindings.FuncHasOnCreated(
		js.Pointer(&fn),
	)
	return
}

// HasOnCreated calls the function "WEBEXT.tabs.onCreated.hasListener" directly.
func HasOnCreated(callback js.Func[func(tab *Tab)]) (ret bool) {
	bindings.CallHasOnCreated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCreated calls the function "WEBEXT.tabs.onCreated.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCreated(callback js.Func[func(tab *Tab)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCreated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDetachedEventCallbackFunc func(this js.Ref, tabId int64, detachInfo *OnDetachedArgDetachInfo) js.Ref

func (fn OnDetachedEventCallbackFunc) Register() js.Func[func(tabId int64, detachInfo *OnDetachedArgDetachInfo)] {
	return js.RegisterCallback[func(tabId int64, detachInfo *OnDetachedArgDetachInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDetachedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 OnDetachedArgDetachInfo
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnDetachedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, tabId int64, detachInfo *OnDetachedArgDetachInfo) js.Ref
	Arg T
}

func (cb *OnDetachedEventCallback[T]) Register() js.Func[func(tabId int64, detachInfo *OnDetachedArgDetachInfo)] {
	return js.RegisterCallback[func(tabId int64, detachInfo *OnDetachedArgDetachInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDetachedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 OnDetachedArgDetachInfo
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnDetached returns true if the function "WEBEXT.tabs.onDetached.addListener" exists.
func HasFuncOnDetached() bool {
	return js.True == bindings.HasFuncOnDetached()
}

// FuncOnDetached returns the function "WEBEXT.tabs.onDetached.addListener".
func FuncOnDetached() (fn js.Func[func(callback js.Func[func(tabId int64, detachInfo *OnDetachedArgDetachInfo)])]) {
	bindings.FuncOnDetached(
		js.Pointer(&fn),
	)
	return
}

// OnDetached calls the function "WEBEXT.tabs.onDetached.addListener" directly.
func OnDetached(callback js.Func[func(tabId int64, detachInfo *OnDetachedArgDetachInfo)]) (ret js.Void) {
	bindings.CallOnDetached(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDetached calls the function "WEBEXT.tabs.onDetached.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDetached(callback js.Func[func(tabId int64, detachInfo *OnDetachedArgDetachInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDetached(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDetached returns true if the function "WEBEXT.tabs.onDetached.removeListener" exists.
func HasFuncOffDetached() bool {
	return js.True == bindings.HasFuncOffDetached()
}

// FuncOffDetached returns the function "WEBEXT.tabs.onDetached.removeListener".
func FuncOffDetached() (fn js.Func[func(callback js.Func[func(tabId int64, detachInfo *OnDetachedArgDetachInfo)])]) {
	bindings.FuncOffDetached(
		js.Pointer(&fn),
	)
	return
}

// OffDetached calls the function "WEBEXT.tabs.onDetached.removeListener" directly.
func OffDetached(callback js.Func[func(tabId int64, detachInfo *OnDetachedArgDetachInfo)]) (ret js.Void) {
	bindings.CallOffDetached(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDetached calls the function "WEBEXT.tabs.onDetached.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDetached(callback js.Func[func(tabId int64, detachInfo *OnDetachedArgDetachInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDetached(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDetached returns true if the function "WEBEXT.tabs.onDetached.hasListener" exists.
func HasFuncHasOnDetached() bool {
	return js.True == bindings.HasFuncHasOnDetached()
}

// FuncHasOnDetached returns the function "WEBEXT.tabs.onDetached.hasListener".
func FuncHasOnDetached() (fn js.Func[func(callback js.Func[func(tabId int64, detachInfo *OnDetachedArgDetachInfo)]) bool]) {
	bindings.FuncHasOnDetached(
		js.Pointer(&fn),
	)
	return
}

// HasOnDetached calls the function "WEBEXT.tabs.onDetached.hasListener" directly.
func HasOnDetached(callback js.Func[func(tabId int64, detachInfo *OnDetachedArgDetachInfo)]) (ret bool) {
	bindings.CallHasOnDetached(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDetached calls the function "WEBEXT.tabs.onDetached.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDetached(callback js.Func[func(tabId int64, detachInfo *OnDetachedArgDetachInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDetached(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnHighlightChangedEventCallbackFunc func(this js.Ref, selectInfo *OnHighlightChangedArgSelectInfo) js.Ref

func (fn OnHighlightChangedEventCallbackFunc) Register() js.Func[func(selectInfo *OnHighlightChangedArgSelectInfo)] {
	return js.RegisterCallback[func(selectInfo *OnHighlightChangedArgSelectInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnHighlightChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnHighlightChangedArgSelectInfo
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

type OnHighlightChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, selectInfo *OnHighlightChangedArgSelectInfo) js.Ref
	Arg T
}

func (cb *OnHighlightChangedEventCallback[T]) Register() js.Func[func(selectInfo *OnHighlightChangedArgSelectInfo)] {
	return js.RegisterCallback[func(selectInfo *OnHighlightChangedArgSelectInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnHighlightChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnHighlightChangedArgSelectInfo
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

// HasFuncOnHighlightChanged returns true if the function "WEBEXT.tabs.onHighlightChanged.addListener" exists.
func HasFuncOnHighlightChanged() bool {
	return js.True == bindings.HasFuncOnHighlightChanged()
}

// FuncOnHighlightChanged returns the function "WEBEXT.tabs.onHighlightChanged.addListener".
func FuncOnHighlightChanged() (fn js.Func[func(callback js.Func[func(selectInfo *OnHighlightChangedArgSelectInfo)])]) {
	bindings.FuncOnHighlightChanged(
		js.Pointer(&fn),
	)
	return
}

// OnHighlightChanged calls the function "WEBEXT.tabs.onHighlightChanged.addListener" directly.
func OnHighlightChanged(callback js.Func[func(selectInfo *OnHighlightChangedArgSelectInfo)]) (ret js.Void) {
	bindings.CallOnHighlightChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnHighlightChanged calls the function "WEBEXT.tabs.onHighlightChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnHighlightChanged(callback js.Func[func(selectInfo *OnHighlightChangedArgSelectInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnHighlightChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffHighlightChanged returns true if the function "WEBEXT.tabs.onHighlightChanged.removeListener" exists.
func HasFuncOffHighlightChanged() bool {
	return js.True == bindings.HasFuncOffHighlightChanged()
}

// FuncOffHighlightChanged returns the function "WEBEXT.tabs.onHighlightChanged.removeListener".
func FuncOffHighlightChanged() (fn js.Func[func(callback js.Func[func(selectInfo *OnHighlightChangedArgSelectInfo)])]) {
	bindings.FuncOffHighlightChanged(
		js.Pointer(&fn),
	)
	return
}

// OffHighlightChanged calls the function "WEBEXT.tabs.onHighlightChanged.removeListener" directly.
func OffHighlightChanged(callback js.Func[func(selectInfo *OnHighlightChangedArgSelectInfo)]) (ret js.Void) {
	bindings.CallOffHighlightChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffHighlightChanged calls the function "WEBEXT.tabs.onHighlightChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffHighlightChanged(callback js.Func[func(selectInfo *OnHighlightChangedArgSelectInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffHighlightChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnHighlightChanged returns true if the function "WEBEXT.tabs.onHighlightChanged.hasListener" exists.
func HasFuncHasOnHighlightChanged() bool {
	return js.True == bindings.HasFuncHasOnHighlightChanged()
}

// FuncHasOnHighlightChanged returns the function "WEBEXT.tabs.onHighlightChanged.hasListener".
func FuncHasOnHighlightChanged() (fn js.Func[func(callback js.Func[func(selectInfo *OnHighlightChangedArgSelectInfo)]) bool]) {
	bindings.FuncHasOnHighlightChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnHighlightChanged calls the function "WEBEXT.tabs.onHighlightChanged.hasListener" directly.
func HasOnHighlightChanged(callback js.Func[func(selectInfo *OnHighlightChangedArgSelectInfo)]) (ret bool) {
	bindings.CallHasOnHighlightChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnHighlightChanged calls the function "WEBEXT.tabs.onHighlightChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnHighlightChanged(callback js.Func[func(selectInfo *OnHighlightChangedArgSelectInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnHighlightChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnHighlightedEventCallbackFunc func(this js.Ref, highlightInfo *OnHighlightedArgHighlightInfo) js.Ref

func (fn OnHighlightedEventCallbackFunc) Register() js.Func[func(highlightInfo *OnHighlightedArgHighlightInfo)] {
	return js.RegisterCallback[func(highlightInfo *OnHighlightedArgHighlightInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnHighlightedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnHighlightedArgHighlightInfo
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

type OnHighlightedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, highlightInfo *OnHighlightedArgHighlightInfo) js.Ref
	Arg T
}

func (cb *OnHighlightedEventCallback[T]) Register() js.Func[func(highlightInfo *OnHighlightedArgHighlightInfo)] {
	return js.RegisterCallback[func(highlightInfo *OnHighlightedArgHighlightInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnHighlightedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnHighlightedArgHighlightInfo
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

// HasFuncOnHighlighted returns true if the function "WEBEXT.tabs.onHighlighted.addListener" exists.
func HasFuncOnHighlighted() bool {
	return js.True == bindings.HasFuncOnHighlighted()
}

// FuncOnHighlighted returns the function "WEBEXT.tabs.onHighlighted.addListener".
func FuncOnHighlighted() (fn js.Func[func(callback js.Func[func(highlightInfo *OnHighlightedArgHighlightInfo)])]) {
	bindings.FuncOnHighlighted(
		js.Pointer(&fn),
	)
	return
}

// OnHighlighted calls the function "WEBEXT.tabs.onHighlighted.addListener" directly.
func OnHighlighted(callback js.Func[func(highlightInfo *OnHighlightedArgHighlightInfo)]) (ret js.Void) {
	bindings.CallOnHighlighted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnHighlighted calls the function "WEBEXT.tabs.onHighlighted.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnHighlighted(callback js.Func[func(highlightInfo *OnHighlightedArgHighlightInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnHighlighted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffHighlighted returns true if the function "WEBEXT.tabs.onHighlighted.removeListener" exists.
func HasFuncOffHighlighted() bool {
	return js.True == bindings.HasFuncOffHighlighted()
}

// FuncOffHighlighted returns the function "WEBEXT.tabs.onHighlighted.removeListener".
func FuncOffHighlighted() (fn js.Func[func(callback js.Func[func(highlightInfo *OnHighlightedArgHighlightInfo)])]) {
	bindings.FuncOffHighlighted(
		js.Pointer(&fn),
	)
	return
}

// OffHighlighted calls the function "WEBEXT.tabs.onHighlighted.removeListener" directly.
func OffHighlighted(callback js.Func[func(highlightInfo *OnHighlightedArgHighlightInfo)]) (ret js.Void) {
	bindings.CallOffHighlighted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffHighlighted calls the function "WEBEXT.tabs.onHighlighted.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffHighlighted(callback js.Func[func(highlightInfo *OnHighlightedArgHighlightInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffHighlighted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnHighlighted returns true if the function "WEBEXT.tabs.onHighlighted.hasListener" exists.
func HasFuncHasOnHighlighted() bool {
	return js.True == bindings.HasFuncHasOnHighlighted()
}

// FuncHasOnHighlighted returns the function "WEBEXT.tabs.onHighlighted.hasListener".
func FuncHasOnHighlighted() (fn js.Func[func(callback js.Func[func(highlightInfo *OnHighlightedArgHighlightInfo)]) bool]) {
	bindings.FuncHasOnHighlighted(
		js.Pointer(&fn),
	)
	return
}

// HasOnHighlighted calls the function "WEBEXT.tabs.onHighlighted.hasListener" directly.
func HasOnHighlighted(callback js.Func[func(highlightInfo *OnHighlightedArgHighlightInfo)]) (ret bool) {
	bindings.CallHasOnHighlighted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnHighlighted calls the function "WEBEXT.tabs.onHighlighted.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnHighlighted(callback js.Func[func(highlightInfo *OnHighlightedArgHighlightInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnHighlighted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnMovedEventCallbackFunc func(this js.Ref, tabId int64, moveInfo *OnMovedArgMoveInfo) js.Ref

func (fn OnMovedEventCallbackFunc) Register() js.Func[func(tabId int64, moveInfo *OnMovedArgMoveInfo)] {
	return js.RegisterCallback[func(tabId int64, moveInfo *OnMovedArgMoveInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnMovedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 OnMovedArgMoveInfo
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnMovedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, tabId int64, moveInfo *OnMovedArgMoveInfo) js.Ref
	Arg T
}

func (cb *OnMovedEventCallback[T]) Register() js.Func[func(tabId int64, moveInfo *OnMovedArgMoveInfo)] {
	return js.RegisterCallback[func(tabId int64, moveInfo *OnMovedArgMoveInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnMovedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 OnMovedArgMoveInfo
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnMoved returns true if the function "WEBEXT.tabs.onMoved.addListener" exists.
func HasFuncOnMoved() bool {
	return js.True == bindings.HasFuncOnMoved()
}

// FuncOnMoved returns the function "WEBEXT.tabs.onMoved.addListener".
func FuncOnMoved() (fn js.Func[func(callback js.Func[func(tabId int64, moveInfo *OnMovedArgMoveInfo)])]) {
	bindings.FuncOnMoved(
		js.Pointer(&fn),
	)
	return
}

// OnMoved calls the function "WEBEXT.tabs.onMoved.addListener" directly.
func OnMoved(callback js.Func[func(tabId int64, moveInfo *OnMovedArgMoveInfo)]) (ret js.Void) {
	bindings.CallOnMoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMoved calls the function "WEBEXT.tabs.onMoved.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMoved(callback js.Func[func(tabId int64, moveInfo *OnMovedArgMoveInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMoved returns true if the function "WEBEXT.tabs.onMoved.removeListener" exists.
func HasFuncOffMoved() bool {
	return js.True == bindings.HasFuncOffMoved()
}

// FuncOffMoved returns the function "WEBEXT.tabs.onMoved.removeListener".
func FuncOffMoved() (fn js.Func[func(callback js.Func[func(tabId int64, moveInfo *OnMovedArgMoveInfo)])]) {
	bindings.FuncOffMoved(
		js.Pointer(&fn),
	)
	return
}

// OffMoved calls the function "WEBEXT.tabs.onMoved.removeListener" directly.
func OffMoved(callback js.Func[func(tabId int64, moveInfo *OnMovedArgMoveInfo)]) (ret js.Void) {
	bindings.CallOffMoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMoved calls the function "WEBEXT.tabs.onMoved.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMoved(callback js.Func[func(tabId int64, moveInfo *OnMovedArgMoveInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMoved returns true if the function "WEBEXT.tabs.onMoved.hasListener" exists.
func HasFuncHasOnMoved() bool {
	return js.True == bindings.HasFuncHasOnMoved()
}

// FuncHasOnMoved returns the function "WEBEXT.tabs.onMoved.hasListener".
func FuncHasOnMoved() (fn js.Func[func(callback js.Func[func(tabId int64, moveInfo *OnMovedArgMoveInfo)]) bool]) {
	bindings.FuncHasOnMoved(
		js.Pointer(&fn),
	)
	return
}

// HasOnMoved calls the function "WEBEXT.tabs.onMoved.hasListener" directly.
func HasOnMoved(callback js.Func[func(tabId int64, moveInfo *OnMovedArgMoveInfo)]) (ret bool) {
	bindings.CallHasOnMoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMoved calls the function "WEBEXT.tabs.onMoved.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnMoved(callback js.Func[func(tabId int64, moveInfo *OnMovedArgMoveInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnMoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnRemovedEventCallbackFunc func(this js.Ref, tabId int64, removeInfo *OnRemovedArgRemoveInfo) js.Ref

func (fn OnRemovedEventCallbackFunc) Register() js.Func[func(tabId int64, removeInfo *OnRemovedArgRemoveInfo)] {
	return js.RegisterCallback[func(tabId int64, removeInfo *OnRemovedArgRemoveInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnRemovedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 OnRemovedArgRemoveInfo
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnRemovedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, tabId int64, removeInfo *OnRemovedArgRemoveInfo) js.Ref
	Arg T
}

func (cb *OnRemovedEventCallback[T]) Register() js.Func[func(tabId int64, removeInfo *OnRemovedArgRemoveInfo)] {
	return js.RegisterCallback[func(tabId int64, removeInfo *OnRemovedArgRemoveInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnRemovedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 OnRemovedArgRemoveInfo
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnRemoved returns true if the function "WEBEXT.tabs.onRemoved.addListener" exists.
func HasFuncOnRemoved() bool {
	return js.True == bindings.HasFuncOnRemoved()
}

// FuncOnRemoved returns the function "WEBEXT.tabs.onRemoved.addListener".
func FuncOnRemoved() (fn js.Func[func(callback js.Func[func(tabId int64, removeInfo *OnRemovedArgRemoveInfo)])]) {
	bindings.FuncOnRemoved(
		js.Pointer(&fn),
	)
	return
}

// OnRemoved calls the function "WEBEXT.tabs.onRemoved.addListener" directly.
func OnRemoved(callback js.Func[func(tabId int64, removeInfo *OnRemovedArgRemoveInfo)]) (ret js.Void) {
	bindings.CallOnRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnRemoved calls the function "WEBEXT.tabs.onRemoved.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnRemoved(callback js.Func[func(tabId int64, removeInfo *OnRemovedArgRemoveInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffRemoved returns true if the function "WEBEXT.tabs.onRemoved.removeListener" exists.
func HasFuncOffRemoved() bool {
	return js.True == bindings.HasFuncOffRemoved()
}

// FuncOffRemoved returns the function "WEBEXT.tabs.onRemoved.removeListener".
func FuncOffRemoved() (fn js.Func[func(callback js.Func[func(tabId int64, removeInfo *OnRemovedArgRemoveInfo)])]) {
	bindings.FuncOffRemoved(
		js.Pointer(&fn),
	)
	return
}

// OffRemoved calls the function "WEBEXT.tabs.onRemoved.removeListener" directly.
func OffRemoved(callback js.Func[func(tabId int64, removeInfo *OnRemovedArgRemoveInfo)]) (ret js.Void) {
	bindings.CallOffRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffRemoved calls the function "WEBEXT.tabs.onRemoved.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffRemoved(callback js.Func[func(tabId int64, removeInfo *OnRemovedArgRemoveInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnRemoved returns true if the function "WEBEXT.tabs.onRemoved.hasListener" exists.
func HasFuncHasOnRemoved() bool {
	return js.True == bindings.HasFuncHasOnRemoved()
}

// FuncHasOnRemoved returns the function "WEBEXT.tabs.onRemoved.hasListener".
func FuncHasOnRemoved() (fn js.Func[func(callback js.Func[func(tabId int64, removeInfo *OnRemovedArgRemoveInfo)]) bool]) {
	bindings.FuncHasOnRemoved(
		js.Pointer(&fn),
	)
	return
}

// HasOnRemoved calls the function "WEBEXT.tabs.onRemoved.hasListener" directly.
func HasOnRemoved(callback js.Func[func(tabId int64, removeInfo *OnRemovedArgRemoveInfo)]) (ret bool) {
	bindings.CallHasOnRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnRemoved calls the function "WEBEXT.tabs.onRemoved.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnRemoved(callback js.Func[func(tabId int64, removeInfo *OnRemovedArgRemoveInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnReplacedEventCallbackFunc func(this js.Ref, addedTabId int64, removedTabId int64) js.Ref

func (fn OnReplacedEventCallbackFunc) Register() js.Func[func(addedTabId int64, removedTabId int64)] {
	return js.RegisterCallback[func(addedTabId int64, removedTabId int64)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnReplacedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
		js.BigInt[int64]{}.FromRef(args[1+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnReplacedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, addedTabId int64, removedTabId int64) js.Ref
	Arg T
}

func (cb *OnReplacedEventCallback[T]) Register() js.Func[func(addedTabId int64, removedTabId int64)] {
	return js.RegisterCallback[func(addedTabId int64, removedTabId int64)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnReplacedEventCallback[T]) DispatchCallback(
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

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
		js.BigInt[int64]{}.FromRef(args[1+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnReplaced returns true if the function "WEBEXT.tabs.onReplaced.addListener" exists.
func HasFuncOnReplaced() bool {
	return js.True == bindings.HasFuncOnReplaced()
}

// FuncOnReplaced returns the function "WEBEXT.tabs.onReplaced.addListener".
func FuncOnReplaced() (fn js.Func[func(callback js.Func[func(addedTabId int64, removedTabId int64)])]) {
	bindings.FuncOnReplaced(
		js.Pointer(&fn),
	)
	return
}

// OnReplaced calls the function "WEBEXT.tabs.onReplaced.addListener" directly.
func OnReplaced(callback js.Func[func(addedTabId int64, removedTabId int64)]) (ret js.Void) {
	bindings.CallOnReplaced(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnReplaced calls the function "WEBEXT.tabs.onReplaced.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnReplaced(callback js.Func[func(addedTabId int64, removedTabId int64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnReplaced(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffReplaced returns true if the function "WEBEXT.tabs.onReplaced.removeListener" exists.
func HasFuncOffReplaced() bool {
	return js.True == bindings.HasFuncOffReplaced()
}

// FuncOffReplaced returns the function "WEBEXT.tabs.onReplaced.removeListener".
func FuncOffReplaced() (fn js.Func[func(callback js.Func[func(addedTabId int64, removedTabId int64)])]) {
	bindings.FuncOffReplaced(
		js.Pointer(&fn),
	)
	return
}

// OffReplaced calls the function "WEBEXT.tabs.onReplaced.removeListener" directly.
func OffReplaced(callback js.Func[func(addedTabId int64, removedTabId int64)]) (ret js.Void) {
	bindings.CallOffReplaced(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffReplaced calls the function "WEBEXT.tabs.onReplaced.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffReplaced(callback js.Func[func(addedTabId int64, removedTabId int64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffReplaced(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnReplaced returns true if the function "WEBEXT.tabs.onReplaced.hasListener" exists.
func HasFuncHasOnReplaced() bool {
	return js.True == bindings.HasFuncHasOnReplaced()
}

// FuncHasOnReplaced returns the function "WEBEXT.tabs.onReplaced.hasListener".
func FuncHasOnReplaced() (fn js.Func[func(callback js.Func[func(addedTabId int64, removedTabId int64)]) bool]) {
	bindings.FuncHasOnReplaced(
		js.Pointer(&fn),
	)
	return
}

// HasOnReplaced calls the function "WEBEXT.tabs.onReplaced.hasListener" directly.
func HasOnReplaced(callback js.Func[func(addedTabId int64, removedTabId int64)]) (ret bool) {
	bindings.CallHasOnReplaced(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnReplaced calls the function "WEBEXT.tabs.onReplaced.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnReplaced(callback js.Func[func(addedTabId int64, removedTabId int64)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnReplaced(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSelectionChangedEventCallbackFunc func(this js.Ref, tabId int64, selectInfo *OnSelectionChangedArgSelectInfo) js.Ref

func (fn OnSelectionChangedEventCallbackFunc) Register() js.Func[func(tabId int64, selectInfo *OnSelectionChangedArgSelectInfo)] {
	return js.RegisterCallback[func(tabId int64, selectInfo *OnSelectionChangedArgSelectInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSelectionChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 OnSelectionChangedArgSelectInfo
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnSelectionChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, tabId int64, selectInfo *OnSelectionChangedArgSelectInfo) js.Ref
	Arg T
}

func (cb *OnSelectionChangedEventCallback[T]) Register() js.Func[func(tabId int64, selectInfo *OnSelectionChangedArgSelectInfo)] {
	return js.RegisterCallback[func(tabId int64, selectInfo *OnSelectionChangedArgSelectInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSelectionChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 OnSelectionChangedArgSelectInfo
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnSelectionChanged returns true if the function "WEBEXT.tabs.onSelectionChanged.addListener" exists.
func HasFuncOnSelectionChanged() bool {
	return js.True == bindings.HasFuncOnSelectionChanged()
}

// FuncOnSelectionChanged returns the function "WEBEXT.tabs.onSelectionChanged.addListener".
func FuncOnSelectionChanged() (fn js.Func[func(callback js.Func[func(tabId int64, selectInfo *OnSelectionChangedArgSelectInfo)])]) {
	bindings.FuncOnSelectionChanged(
		js.Pointer(&fn),
	)
	return
}

// OnSelectionChanged calls the function "WEBEXT.tabs.onSelectionChanged.addListener" directly.
func OnSelectionChanged(callback js.Func[func(tabId int64, selectInfo *OnSelectionChangedArgSelectInfo)]) (ret js.Void) {
	bindings.CallOnSelectionChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSelectionChanged calls the function "WEBEXT.tabs.onSelectionChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSelectionChanged(callback js.Func[func(tabId int64, selectInfo *OnSelectionChangedArgSelectInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSelectionChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSelectionChanged returns true if the function "WEBEXT.tabs.onSelectionChanged.removeListener" exists.
func HasFuncOffSelectionChanged() bool {
	return js.True == bindings.HasFuncOffSelectionChanged()
}

// FuncOffSelectionChanged returns the function "WEBEXT.tabs.onSelectionChanged.removeListener".
func FuncOffSelectionChanged() (fn js.Func[func(callback js.Func[func(tabId int64, selectInfo *OnSelectionChangedArgSelectInfo)])]) {
	bindings.FuncOffSelectionChanged(
		js.Pointer(&fn),
	)
	return
}

// OffSelectionChanged calls the function "WEBEXT.tabs.onSelectionChanged.removeListener" directly.
func OffSelectionChanged(callback js.Func[func(tabId int64, selectInfo *OnSelectionChangedArgSelectInfo)]) (ret js.Void) {
	bindings.CallOffSelectionChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSelectionChanged calls the function "WEBEXT.tabs.onSelectionChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSelectionChanged(callback js.Func[func(tabId int64, selectInfo *OnSelectionChangedArgSelectInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSelectionChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSelectionChanged returns true if the function "WEBEXT.tabs.onSelectionChanged.hasListener" exists.
func HasFuncHasOnSelectionChanged() bool {
	return js.True == bindings.HasFuncHasOnSelectionChanged()
}

// FuncHasOnSelectionChanged returns the function "WEBEXT.tabs.onSelectionChanged.hasListener".
func FuncHasOnSelectionChanged() (fn js.Func[func(callback js.Func[func(tabId int64, selectInfo *OnSelectionChangedArgSelectInfo)]) bool]) {
	bindings.FuncHasOnSelectionChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnSelectionChanged calls the function "WEBEXT.tabs.onSelectionChanged.hasListener" directly.
func HasOnSelectionChanged(callback js.Func[func(tabId int64, selectInfo *OnSelectionChangedArgSelectInfo)]) (ret bool) {
	bindings.CallHasOnSelectionChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSelectionChanged calls the function "WEBEXT.tabs.onSelectionChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSelectionChanged(callback js.Func[func(tabId int64, selectInfo *OnSelectionChangedArgSelectInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSelectionChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnUpdatedEventCallbackFunc func(this js.Ref, tabId int64, changeInfo *OnUpdatedArgChangeInfo, tab *Tab) js.Ref

func (fn OnUpdatedEventCallbackFunc) Register() js.Func[func(tabId int64, changeInfo *OnUpdatedArgChangeInfo, tab *Tab)] {
	return js.RegisterCallback[func(tabId int64, changeInfo *OnUpdatedArgChangeInfo, tab *Tab)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnUpdatedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 OnUpdatedArgChangeInfo
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)
	var arg2 Tab
	arg2.UpdateFrom(args[2+1])
	defer arg2.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
		mark.NoEscape(&arg1),
		mark.NoEscape(&arg2),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnUpdatedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, tabId int64, changeInfo *OnUpdatedArgChangeInfo, tab *Tab) js.Ref
	Arg T
}

func (cb *OnUpdatedEventCallback[T]) Register() js.Func[func(tabId int64, changeInfo *OnUpdatedArgChangeInfo, tab *Tab)] {
	return js.RegisterCallback[func(tabId int64, changeInfo *OnUpdatedArgChangeInfo, tab *Tab)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnUpdatedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 OnUpdatedArgChangeInfo
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)
	var arg2 Tab
	arg2.UpdateFrom(args[2+1])
	defer arg2.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
		mark.NoEscape(&arg1),
		mark.NoEscape(&arg2),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnUpdated returns true if the function "WEBEXT.tabs.onUpdated.addListener" exists.
func HasFuncOnUpdated() bool {
	return js.True == bindings.HasFuncOnUpdated()
}

// FuncOnUpdated returns the function "WEBEXT.tabs.onUpdated.addListener".
func FuncOnUpdated() (fn js.Func[func(callback js.Func[func(tabId int64, changeInfo *OnUpdatedArgChangeInfo, tab *Tab)])]) {
	bindings.FuncOnUpdated(
		js.Pointer(&fn),
	)
	return
}

// OnUpdated calls the function "WEBEXT.tabs.onUpdated.addListener" directly.
func OnUpdated(callback js.Func[func(tabId int64, changeInfo *OnUpdatedArgChangeInfo, tab *Tab)]) (ret js.Void) {
	bindings.CallOnUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnUpdated calls the function "WEBEXT.tabs.onUpdated.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnUpdated(callback js.Func[func(tabId int64, changeInfo *OnUpdatedArgChangeInfo, tab *Tab)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffUpdated returns true if the function "WEBEXT.tabs.onUpdated.removeListener" exists.
func HasFuncOffUpdated() bool {
	return js.True == bindings.HasFuncOffUpdated()
}

// FuncOffUpdated returns the function "WEBEXT.tabs.onUpdated.removeListener".
func FuncOffUpdated() (fn js.Func[func(callback js.Func[func(tabId int64, changeInfo *OnUpdatedArgChangeInfo, tab *Tab)])]) {
	bindings.FuncOffUpdated(
		js.Pointer(&fn),
	)
	return
}

// OffUpdated calls the function "WEBEXT.tabs.onUpdated.removeListener" directly.
func OffUpdated(callback js.Func[func(tabId int64, changeInfo *OnUpdatedArgChangeInfo, tab *Tab)]) (ret js.Void) {
	bindings.CallOffUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffUpdated calls the function "WEBEXT.tabs.onUpdated.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffUpdated(callback js.Func[func(tabId int64, changeInfo *OnUpdatedArgChangeInfo, tab *Tab)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnUpdated returns true if the function "WEBEXT.tabs.onUpdated.hasListener" exists.
func HasFuncHasOnUpdated() bool {
	return js.True == bindings.HasFuncHasOnUpdated()
}

// FuncHasOnUpdated returns the function "WEBEXT.tabs.onUpdated.hasListener".
func FuncHasOnUpdated() (fn js.Func[func(callback js.Func[func(tabId int64, changeInfo *OnUpdatedArgChangeInfo, tab *Tab)]) bool]) {
	bindings.FuncHasOnUpdated(
		js.Pointer(&fn),
	)
	return
}

// HasOnUpdated calls the function "WEBEXT.tabs.onUpdated.hasListener" directly.
func HasOnUpdated(callback js.Func[func(tabId int64, changeInfo *OnUpdatedArgChangeInfo, tab *Tab)]) (ret bool) {
	bindings.CallHasOnUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnUpdated calls the function "WEBEXT.tabs.onUpdated.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnUpdated(callback js.Func[func(tabId int64, changeInfo *OnUpdatedArgChangeInfo, tab *Tab)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnZoomChangeEventCallbackFunc func(this js.Ref, ZoomChangeInfo *OnZoomChangeArgZoomChangeInfo) js.Ref

func (fn OnZoomChangeEventCallbackFunc) Register() js.Func[func(ZoomChangeInfo *OnZoomChangeArgZoomChangeInfo)] {
	return js.RegisterCallback[func(ZoomChangeInfo *OnZoomChangeArgZoomChangeInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnZoomChangeEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnZoomChangeArgZoomChangeInfo
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

type OnZoomChangeEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, ZoomChangeInfo *OnZoomChangeArgZoomChangeInfo) js.Ref
	Arg T
}

func (cb *OnZoomChangeEventCallback[T]) Register() js.Func[func(ZoomChangeInfo *OnZoomChangeArgZoomChangeInfo)] {
	return js.RegisterCallback[func(ZoomChangeInfo *OnZoomChangeArgZoomChangeInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnZoomChangeEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnZoomChangeArgZoomChangeInfo
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

// HasFuncOnZoomChange returns true if the function "WEBEXT.tabs.onZoomChange.addListener" exists.
func HasFuncOnZoomChange() bool {
	return js.True == bindings.HasFuncOnZoomChange()
}

// FuncOnZoomChange returns the function "WEBEXT.tabs.onZoomChange.addListener".
func FuncOnZoomChange() (fn js.Func[func(callback js.Func[func(ZoomChangeInfo *OnZoomChangeArgZoomChangeInfo)])]) {
	bindings.FuncOnZoomChange(
		js.Pointer(&fn),
	)
	return
}

// OnZoomChange calls the function "WEBEXT.tabs.onZoomChange.addListener" directly.
func OnZoomChange(callback js.Func[func(ZoomChangeInfo *OnZoomChangeArgZoomChangeInfo)]) (ret js.Void) {
	bindings.CallOnZoomChange(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnZoomChange calls the function "WEBEXT.tabs.onZoomChange.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnZoomChange(callback js.Func[func(ZoomChangeInfo *OnZoomChangeArgZoomChangeInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnZoomChange(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffZoomChange returns true if the function "WEBEXT.tabs.onZoomChange.removeListener" exists.
func HasFuncOffZoomChange() bool {
	return js.True == bindings.HasFuncOffZoomChange()
}

// FuncOffZoomChange returns the function "WEBEXT.tabs.onZoomChange.removeListener".
func FuncOffZoomChange() (fn js.Func[func(callback js.Func[func(ZoomChangeInfo *OnZoomChangeArgZoomChangeInfo)])]) {
	bindings.FuncOffZoomChange(
		js.Pointer(&fn),
	)
	return
}

// OffZoomChange calls the function "WEBEXT.tabs.onZoomChange.removeListener" directly.
func OffZoomChange(callback js.Func[func(ZoomChangeInfo *OnZoomChangeArgZoomChangeInfo)]) (ret js.Void) {
	bindings.CallOffZoomChange(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffZoomChange calls the function "WEBEXT.tabs.onZoomChange.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffZoomChange(callback js.Func[func(ZoomChangeInfo *OnZoomChangeArgZoomChangeInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffZoomChange(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnZoomChange returns true if the function "WEBEXT.tabs.onZoomChange.hasListener" exists.
func HasFuncHasOnZoomChange() bool {
	return js.True == bindings.HasFuncHasOnZoomChange()
}

// FuncHasOnZoomChange returns the function "WEBEXT.tabs.onZoomChange.hasListener".
func FuncHasOnZoomChange() (fn js.Func[func(callback js.Func[func(ZoomChangeInfo *OnZoomChangeArgZoomChangeInfo)]) bool]) {
	bindings.FuncHasOnZoomChange(
		js.Pointer(&fn),
	)
	return
}

// HasOnZoomChange calls the function "WEBEXT.tabs.onZoomChange.hasListener" directly.
func HasOnZoomChange(callback js.Func[func(ZoomChangeInfo *OnZoomChangeArgZoomChangeInfo)]) (ret bool) {
	bindings.CallHasOnZoomChange(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnZoomChange calls the function "WEBEXT.tabs.onZoomChange.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnZoomChange(callback js.Func[func(ZoomChangeInfo *OnZoomChangeArgZoomChangeInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnZoomChange(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncQuery returns true if the function "WEBEXT.tabs.query" exists.
func HasFuncQuery() bool {
	return js.True == bindings.HasFuncQuery()
}

// FuncQuery returns the function "WEBEXT.tabs.query".
func FuncQuery() (fn js.Func[func(queryInfo QueryArgQueryInfo) js.Promise[js.Array[Tab]]]) {
	bindings.FuncQuery(
		js.Pointer(&fn),
	)
	return
}

// Query calls the function "WEBEXT.tabs.query" directly.
func Query(queryInfo QueryArgQueryInfo) (ret js.Promise[js.Array[Tab]]) {
	bindings.CallQuery(
		js.Pointer(&ret),
		js.Pointer(&queryInfo),
	)

	return
}

// TryQuery calls the function "WEBEXT.tabs.query"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryQuery(queryInfo QueryArgQueryInfo) (ret js.Promise[js.Array[Tab]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryQuery(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&queryInfo),
	)

	return
}

// HasFuncReload returns true if the function "WEBEXT.tabs.reload" exists.
func HasFuncReload() bool {
	return js.True == bindings.HasFuncReload()
}

// FuncReload returns the function "WEBEXT.tabs.reload".
func FuncReload() (fn js.Func[func(tabId int64, reloadProperties ReloadArgReloadProperties) js.Promise[js.Void]]) {
	bindings.FuncReload(
		js.Pointer(&fn),
	)
	return
}

// Reload calls the function "WEBEXT.tabs.reload" directly.
func Reload(tabId int64, reloadProperties ReloadArgReloadProperties) (ret js.Promise[js.Void]) {
	bindings.CallReload(
		js.Pointer(&ret),
		float64(tabId),
		js.Pointer(&reloadProperties),
	)

	return
}

// TryReload calls the function "WEBEXT.tabs.reload"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReload(tabId int64, reloadProperties ReloadArgReloadProperties) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReload(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(tabId),
		js.Pointer(&reloadProperties),
	)

	return
}

// HasFuncRemove returns true if the function "WEBEXT.tabs.remove" exists.
func HasFuncRemove() bool {
	return js.True == bindings.HasFuncRemove()
}

// FuncRemove returns the function "WEBEXT.tabs.remove".
func FuncRemove() (fn js.Func[func(tabIds OneOf_Int64_ArrayInt64) js.Promise[js.Void]]) {
	bindings.FuncRemove(
		js.Pointer(&fn),
	)
	return
}

// Remove calls the function "WEBEXT.tabs.remove" directly.
func Remove(tabIds OneOf_Int64_ArrayInt64) (ret js.Promise[js.Void]) {
	bindings.CallRemove(
		js.Pointer(&ret),
		tabIds.Ref(),
	)

	return
}

// TryRemove calls the function "WEBEXT.tabs.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemove(tabIds OneOf_Int64_ArrayInt64) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemove(
		js.Pointer(&ret), js.Pointer(&exception),
		tabIds.Ref(),
	)

	return
}

// HasFuncRemoveCSS returns true if the function "WEBEXT.tabs.removeCSS" exists.
func HasFuncRemoveCSS() bool {
	return js.True == bindings.HasFuncRemoveCSS()
}

// FuncRemoveCSS returns the function "WEBEXT.tabs.removeCSS".
func FuncRemoveCSS() (fn js.Func[func(tabId int64, details extensiontypes.DeleteInjectionDetails) js.Promise[js.Void]]) {
	bindings.FuncRemoveCSS(
		js.Pointer(&fn),
	)
	return
}

// RemoveCSS calls the function "WEBEXT.tabs.removeCSS" directly.
func RemoveCSS(tabId int64, details extensiontypes.DeleteInjectionDetails) (ret js.Promise[js.Void]) {
	bindings.CallRemoveCSS(
		js.Pointer(&ret),
		float64(tabId),
		js.Pointer(&details),
	)

	return
}

// TryRemoveCSS calls the function "WEBEXT.tabs.removeCSS"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveCSS(tabId int64, details extensiontypes.DeleteInjectionDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveCSS(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(tabId),
		js.Pointer(&details),
	)

	return
}

// HasFuncSendMessage returns true if the function "WEBEXT.tabs.sendMessage" exists.
func HasFuncSendMessage() bool {
	return js.True == bindings.HasFuncSendMessage()
}

// FuncSendMessage returns the function "WEBEXT.tabs.sendMessage".
func FuncSendMessage() (fn js.Func[func(tabId int64, message js.Any, options SendMessageArgOptions) js.Promise[js.Any]]) {
	bindings.FuncSendMessage(
		js.Pointer(&fn),
	)
	return
}

// SendMessage calls the function "WEBEXT.tabs.sendMessage" directly.
func SendMessage(tabId int64, message js.Any, options SendMessageArgOptions) (ret js.Promise[js.Any]) {
	bindings.CallSendMessage(
		js.Pointer(&ret),
		float64(tabId),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// TrySendMessage calls the function "WEBEXT.tabs.sendMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendMessage(tabId int64, message js.Any, options SendMessageArgOptions) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendMessage(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(tabId),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncSendRequest returns true if the function "WEBEXT.tabs.sendRequest" exists.
func HasFuncSendRequest() bool {
	return js.True == bindings.HasFuncSendRequest()
}

// FuncSendRequest returns the function "WEBEXT.tabs.sendRequest".
func FuncSendRequest() (fn js.Func[func(tabId int64, request js.Any) js.Promise[js.Any]]) {
	bindings.FuncSendRequest(
		js.Pointer(&fn),
	)
	return
}

// SendRequest calls the function "WEBEXT.tabs.sendRequest" directly.
func SendRequest(tabId int64, request js.Any) (ret js.Promise[js.Any]) {
	bindings.CallSendRequest(
		js.Pointer(&ret),
		float64(tabId),
		request.Ref(),
	)

	return
}

// TrySendRequest calls the function "WEBEXT.tabs.sendRequest"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendRequest(tabId int64, request js.Any) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(tabId),
		request.Ref(),
	)

	return
}

// HasFuncSetZoom returns true if the function "WEBEXT.tabs.setZoom" exists.
func HasFuncSetZoom() bool {
	return js.True == bindings.HasFuncSetZoom()
}

// FuncSetZoom returns the function "WEBEXT.tabs.setZoom".
func FuncSetZoom() (fn js.Func[func(tabId int64, zoomFactor float64) js.Promise[js.Void]]) {
	bindings.FuncSetZoom(
		js.Pointer(&fn),
	)
	return
}

// SetZoom calls the function "WEBEXT.tabs.setZoom" directly.
func SetZoom(tabId int64, zoomFactor float64) (ret js.Promise[js.Void]) {
	bindings.CallSetZoom(
		js.Pointer(&ret),
		float64(tabId),
		float64(zoomFactor),
	)

	return
}

// TrySetZoom calls the function "WEBEXT.tabs.setZoom"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetZoom(tabId int64, zoomFactor float64) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetZoom(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(tabId),
		float64(zoomFactor),
	)

	return
}

// HasFuncSetZoomSettings returns true if the function "WEBEXT.tabs.setZoomSettings" exists.
func HasFuncSetZoomSettings() bool {
	return js.True == bindings.HasFuncSetZoomSettings()
}

// FuncSetZoomSettings returns the function "WEBEXT.tabs.setZoomSettings".
func FuncSetZoomSettings() (fn js.Func[func(tabId int64, zoomSettings ZoomSettings) js.Promise[js.Void]]) {
	bindings.FuncSetZoomSettings(
		js.Pointer(&fn),
	)
	return
}

// SetZoomSettings calls the function "WEBEXT.tabs.setZoomSettings" directly.
func SetZoomSettings(tabId int64, zoomSettings ZoomSettings) (ret js.Promise[js.Void]) {
	bindings.CallSetZoomSettings(
		js.Pointer(&ret),
		float64(tabId),
		js.Pointer(&zoomSettings),
	)

	return
}

// TrySetZoomSettings calls the function "WEBEXT.tabs.setZoomSettings"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetZoomSettings(tabId int64, zoomSettings ZoomSettings) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetZoomSettings(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(tabId),
		js.Pointer(&zoomSettings),
	)

	return
}

// HasFuncUngroup returns true if the function "WEBEXT.tabs.ungroup" exists.
func HasFuncUngroup() bool {
	return js.True == bindings.HasFuncUngroup()
}

// FuncUngroup returns the function "WEBEXT.tabs.ungroup".
func FuncUngroup() (fn js.Func[func(tabIds OneOf_Int64_ArrayInt64) js.Promise[js.Void]]) {
	bindings.FuncUngroup(
		js.Pointer(&fn),
	)
	return
}

// Ungroup calls the function "WEBEXT.tabs.ungroup" directly.
func Ungroup(tabIds OneOf_Int64_ArrayInt64) (ret js.Promise[js.Void]) {
	bindings.CallUngroup(
		js.Pointer(&ret),
		tabIds.Ref(),
	)

	return
}

// TryUngroup calls the function "WEBEXT.tabs.ungroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUngroup(tabIds OneOf_Int64_ArrayInt64) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUngroup(
		js.Pointer(&ret), js.Pointer(&exception),
		tabIds.Ref(),
	)

	return
}

// HasFuncUpdate returns true if the function "WEBEXT.tabs.update" exists.
func HasFuncUpdate() bool {
	return js.True == bindings.HasFuncUpdate()
}

// FuncUpdate returns the function "WEBEXT.tabs.update".
func FuncUpdate() (fn js.Func[func(tabId int64, updateProperties UpdateArgUpdateProperties) js.Promise[Tab]]) {
	bindings.FuncUpdate(
		js.Pointer(&fn),
	)
	return
}

// Update calls the function "WEBEXT.tabs.update" directly.
func Update(tabId int64, updateProperties UpdateArgUpdateProperties) (ret js.Promise[Tab]) {
	bindings.CallUpdate(
		js.Pointer(&ret),
		float64(tabId),
		js.Pointer(&updateProperties),
	)

	return
}

// TryUpdate calls the function "WEBEXT.tabs.update"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdate(tabId int64, updateProperties UpdateArgUpdateProperties) (ret js.Promise[Tab], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdate(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(tabId),
		js.Pointer(&updateProperties),
	)

	return
}
