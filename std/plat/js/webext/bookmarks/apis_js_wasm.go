// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package bookmarks

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/bookmarks/bindings"
)

type BookmarkTreeNodeUnmodifiable uint32

const (
	_ BookmarkTreeNodeUnmodifiable = iota

	BookmarkTreeNodeUnmodifiable_MANAGED
)

func (BookmarkTreeNodeUnmodifiable) FromRef(str js.Ref) BookmarkTreeNodeUnmodifiable {
	return BookmarkTreeNodeUnmodifiable(bindings.ConstOfBookmarkTreeNodeUnmodifiable(str))
}

func (x BookmarkTreeNodeUnmodifiable) String() (string, bool) {
	switch x {
	case BookmarkTreeNodeUnmodifiable_MANAGED:
		return "managed", true
	default:
		return "", false
	}
}

type BookmarkTreeNode struct {
	// Children is "BookmarkTreeNode.children"
	//
	// Optional
	Children js.Array[BookmarkTreeNode]
	// DateAdded is "BookmarkTreeNode.dateAdded"
	//
	// Optional
	//
	// NOTE: FFI_USE_DateAdded MUST be set to true to make this field effective.
	DateAdded float64
	// DateGroupModified is "BookmarkTreeNode.dateGroupModified"
	//
	// Optional
	//
	// NOTE: FFI_USE_DateGroupModified MUST be set to true to make this field effective.
	DateGroupModified float64
	// DateLastUsed is "BookmarkTreeNode.dateLastUsed"
	//
	// Optional
	//
	// NOTE: FFI_USE_DateLastUsed MUST be set to true to make this field effective.
	DateLastUsed float64
	// Id is "BookmarkTreeNode.id"
	//
	// Required
	Id js.String
	// Index is "BookmarkTreeNode.index"
	//
	// Optional
	//
	// NOTE: FFI_USE_Index MUST be set to true to make this field effective.
	Index int64
	// ParentId is "BookmarkTreeNode.parentId"
	//
	// Optional
	ParentId js.String
	// Title is "BookmarkTreeNode.title"
	//
	// Required
	Title js.String
	// Unmodifiable is "BookmarkTreeNode.unmodifiable"
	//
	// Optional
	Unmodifiable BookmarkTreeNodeUnmodifiable
	// Url is "BookmarkTreeNode.url"
	//
	// Optional
	Url js.String

	FFI_USE_DateAdded         bool // for DateAdded.
	FFI_USE_DateGroupModified bool // for DateGroupModified.
	FFI_USE_DateLastUsed      bool // for DateLastUsed.
	FFI_USE_Index             bool // for Index.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BookmarkTreeNode with all fields set.
func (p BookmarkTreeNode) FromRef(ref js.Ref) BookmarkTreeNode {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BookmarkTreeNode in the application heap.
func (p BookmarkTreeNode) New() js.Ref {
	return bindings.BookmarkTreeNodeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *BookmarkTreeNode) UpdateFrom(ref js.Ref) {
	bindings.BookmarkTreeNodeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BookmarkTreeNode) Update(ref js.Ref) {
	bindings.BookmarkTreeNodeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BookmarkTreeNode) FreeMembers(recursive bool) {
	js.Free(
		p.Children.Ref(),
		p.Id.Ref(),
		p.ParentId.Ref(),
		p.Title.Ref(),
		p.Url.Ref(),
	)
	p.Children = p.Children.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
	p.ParentId = p.ParentId.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type CreateDetails struct {
	// Index is "CreateDetails.index"
	//
	// Optional
	//
	// NOTE: FFI_USE_Index MUST be set to true to make this field effective.
	Index int64
	// ParentId is "CreateDetails.parentId"
	//
	// Optional
	ParentId js.String
	// Title is "CreateDetails.title"
	//
	// Optional
	Title js.String
	// Url is "CreateDetails.url"
	//
	// Optional
	Url js.String

	FFI_USE_Index bool // for Index.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CreateDetails with all fields set.
func (p CreateDetails) FromRef(ref js.Ref) CreateDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CreateDetails in the application heap.
func (p CreateDetails) New() js.Ref {
	return bindings.CreateDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CreateDetails) UpdateFrom(ref js.Ref) {
	bindings.CreateDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CreateDetails) Update(ref js.Ref) {
	bindings.CreateDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CreateDetails) FreeMembers(recursive bool) {
	js.Free(
		p.ParentId.Ref(),
		p.Title.Ref(),
		p.Url.Ref(),
	)
	p.ParentId = p.ParentId.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

// MAX_SUSTAINED_WRITE_OPERATIONS_PER_MINUTE returns the value of property "WEBEXT.bookmarks.MAX_SUSTAINED_WRITE_OPERATIONS_PER_MINUTE".
//
// The returned bool will be false if there is no such property.
func MAX_SUSTAINED_WRITE_OPERATIONS_PER_MINUTE() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMAX_SUSTAINED_WRITE_OPERATIONS_PER_MINUTE(
		js.Pointer(&ret),
	)

	return
}

// SetMAX_SUSTAINED_WRITE_OPERATIONS_PER_MINUTE sets the value of property "WEBEXT.bookmarks.MAX_SUSTAINED_WRITE_OPERATIONS_PER_MINUTE" to val.
//
// It returns false if the property cannot be set.
func SetMAX_SUSTAINED_WRITE_OPERATIONS_PER_MINUTE(val js.String) bool {
	return js.True == bindings.SetMAX_SUSTAINED_WRITE_OPERATIONS_PER_MINUTE(
		val.Ref())
}

// MAX_WRITE_OPERATIONS_PER_HOUR returns the value of property "WEBEXT.bookmarks.MAX_WRITE_OPERATIONS_PER_HOUR".
//
// The returned bool will be false if there is no such property.
func MAX_WRITE_OPERATIONS_PER_HOUR() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMAX_WRITE_OPERATIONS_PER_HOUR(
		js.Pointer(&ret),
	)

	return
}

// SetMAX_WRITE_OPERATIONS_PER_HOUR sets the value of property "WEBEXT.bookmarks.MAX_WRITE_OPERATIONS_PER_HOUR" to val.
//
// It returns false if the property cannot be set.
func SetMAX_WRITE_OPERATIONS_PER_HOUR(val js.String) bool {
	return js.True == bindings.SetMAX_WRITE_OPERATIONS_PER_HOUR(
		val.Ref())
}

type MoveArgDestination struct {
	// Index is "MoveArgDestination.index"
	//
	// Optional
	//
	// NOTE: FFI_USE_Index MUST be set to true to make this field effective.
	Index int64
	// ParentId is "MoveArgDestination.parentId"
	//
	// Optional
	ParentId js.String

	FFI_USE_Index bool // for Index.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MoveArgDestination with all fields set.
func (p MoveArgDestination) FromRef(ref js.Ref) MoveArgDestination {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MoveArgDestination in the application heap.
func (p MoveArgDestination) New() js.Ref {
	return bindings.MoveArgDestinationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MoveArgDestination) UpdateFrom(ref js.Ref) {
	bindings.MoveArgDestinationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MoveArgDestination) Update(ref js.Ref) {
	bindings.MoveArgDestinationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MoveArgDestination) FreeMembers(recursive bool) {
	js.Free(
		p.ParentId.Ref(),
	)
	p.ParentId = p.ParentId.FromRef(js.Undefined)
}

type OnChangedArgChangeInfo struct {
	// Title is "OnChangedArgChangeInfo.title"
	//
	// Required
	Title js.String
	// Url is "OnChangedArgChangeInfo.url"
	//
	// Optional
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnChangedArgChangeInfo with all fields set.
func (p OnChangedArgChangeInfo) FromRef(ref js.Ref) OnChangedArgChangeInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnChangedArgChangeInfo in the application heap.
func (p OnChangedArgChangeInfo) New() js.Ref {
	return bindings.OnChangedArgChangeInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnChangedArgChangeInfo) UpdateFrom(ref js.Ref) {
	bindings.OnChangedArgChangeInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnChangedArgChangeInfo) Update(ref js.Ref) {
	bindings.OnChangedArgChangeInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnChangedArgChangeInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Title.Ref(),
		p.Url.Ref(),
	)
	p.Title = p.Title.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type OnChildrenReorderedArgReorderInfo struct {
	// ChildIds is "OnChildrenReorderedArgReorderInfo.childIds"
	//
	// Required
	ChildIds js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnChildrenReorderedArgReorderInfo with all fields set.
func (p OnChildrenReorderedArgReorderInfo) FromRef(ref js.Ref) OnChildrenReorderedArgReorderInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnChildrenReorderedArgReorderInfo in the application heap.
func (p OnChildrenReorderedArgReorderInfo) New() js.Ref {
	return bindings.OnChildrenReorderedArgReorderInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnChildrenReorderedArgReorderInfo) UpdateFrom(ref js.Ref) {
	bindings.OnChildrenReorderedArgReorderInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnChildrenReorderedArgReorderInfo) Update(ref js.Ref) {
	bindings.OnChildrenReorderedArgReorderInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnChildrenReorderedArgReorderInfo) FreeMembers(recursive bool) {
	js.Free(
		p.ChildIds.Ref(),
	)
	p.ChildIds = p.ChildIds.FromRef(js.Undefined)
}

type OnMovedArgMoveInfo struct {
	// Index is "OnMovedArgMoveInfo.index"
	//
	// Required
	Index int64
	// OldIndex is "OnMovedArgMoveInfo.oldIndex"
	//
	// Required
	OldIndex int64
	// OldParentId is "OnMovedArgMoveInfo.oldParentId"
	//
	// Required
	OldParentId js.String
	// ParentId is "OnMovedArgMoveInfo.parentId"
	//
	// Required
	ParentId js.String

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
	js.Free(
		p.OldParentId.Ref(),
		p.ParentId.Ref(),
	)
	p.OldParentId = p.OldParentId.FromRef(js.Undefined)
	p.ParentId = p.ParentId.FromRef(js.Undefined)
}

type OnRemovedArgRemoveInfo struct {
	// Index is "OnRemovedArgRemoveInfo.index"
	//
	// Required
	Index int64
	// Node is "OnRemovedArgRemoveInfo.node"
	//
	// Required
	//
	// NOTE: Node.FFI_USE MUST be set to true to get Node used.
	Node BookmarkTreeNode
	// ParentId is "OnRemovedArgRemoveInfo.parentId"
	//
	// Required
	ParentId js.String

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
	js.Free(
		p.ParentId.Ref(),
	)
	p.ParentId = p.ParentId.FromRef(js.Undefined)
	if recursive {
		p.Node.FreeMembers(true)
	}
}

type SearchArgQueryChoice1 struct {
	// Query is "SearchArgQueryChoice1.query"
	//
	// Optional
	Query js.String
	// Title is "SearchArgQueryChoice1.title"
	//
	// Optional
	Title js.String
	// Url is "SearchArgQueryChoice1.url"
	//
	// Optional
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SearchArgQueryChoice1 with all fields set.
func (p SearchArgQueryChoice1) FromRef(ref js.Ref) SearchArgQueryChoice1 {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SearchArgQueryChoice1 in the application heap.
func (p SearchArgQueryChoice1) New() js.Ref {
	return bindings.SearchArgQueryChoice1JSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SearchArgQueryChoice1) UpdateFrom(ref js.Ref) {
	bindings.SearchArgQueryChoice1JSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SearchArgQueryChoice1) Update(ref js.Ref) {
	bindings.SearchArgQueryChoice1JSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SearchArgQueryChoice1) FreeMembers(recursive bool) {
	js.Free(
		p.Query.Ref(),
		p.Title.Ref(),
		p.Url.Ref(),
	)
	p.Query = p.Query.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type UpdateArgChanges struct {
	// Title is "UpdateArgChanges.title"
	//
	// Optional
	Title js.String
	// Url is "UpdateArgChanges.url"
	//
	// Optional
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UpdateArgChanges with all fields set.
func (p UpdateArgChanges) FromRef(ref js.Ref) UpdateArgChanges {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UpdateArgChanges in the application heap.
func (p UpdateArgChanges) New() js.Ref {
	return bindings.UpdateArgChangesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UpdateArgChanges) UpdateFrom(ref js.Ref) {
	bindings.UpdateArgChangesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UpdateArgChanges) Update(ref js.Ref) {
	bindings.UpdateArgChangesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UpdateArgChanges) FreeMembers(recursive bool) {
	js.Free(
		p.Title.Ref(),
		p.Url.Ref(),
	)
	p.Title = p.Title.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

// HasFuncCreate returns true if the function "WEBEXT.bookmarks.create" exists.
func HasFuncCreate() bool {
	return js.True == bindings.HasFuncCreate()
}

// FuncCreate returns the function "WEBEXT.bookmarks.create".
func FuncCreate() (fn js.Func[func(bookmark CreateDetails) js.Promise[BookmarkTreeNode]]) {
	bindings.FuncCreate(
		js.Pointer(&fn),
	)
	return
}

// Create calls the function "WEBEXT.bookmarks.create" directly.
func Create(bookmark CreateDetails) (ret js.Promise[BookmarkTreeNode]) {
	bindings.CallCreate(
		js.Pointer(&ret),
		js.Pointer(&bookmark),
	)

	return
}

// TryCreate calls the function "WEBEXT.bookmarks.create"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCreate(bookmark CreateDetails) (ret js.Promise[BookmarkTreeNode], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCreate(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&bookmark),
	)

	return
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

// HasFuncGet returns true if the function "WEBEXT.bookmarks.get" exists.
func HasFuncGet() bool {
	return js.True == bindings.HasFuncGet()
}

// FuncGet returns the function "WEBEXT.bookmarks.get".
func FuncGet() (fn js.Func[func(idOrIdList OneOf_String_ArrayString) js.Promise[js.Array[BookmarkTreeNode]]]) {
	bindings.FuncGet(
		js.Pointer(&fn),
	)
	return
}

// Get calls the function "WEBEXT.bookmarks.get" directly.
func Get(idOrIdList OneOf_String_ArrayString) (ret js.Promise[js.Array[BookmarkTreeNode]]) {
	bindings.CallGet(
		js.Pointer(&ret),
		idOrIdList.Ref(),
	)

	return
}

// TryGet calls the function "WEBEXT.bookmarks.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGet(idOrIdList OneOf_String_ArrayString) (ret js.Promise[js.Array[BookmarkTreeNode]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGet(
		js.Pointer(&ret), js.Pointer(&exception),
		idOrIdList.Ref(),
	)

	return
}

// HasFuncGetChildren returns true if the function "WEBEXT.bookmarks.getChildren" exists.
func HasFuncGetChildren() bool {
	return js.True == bindings.HasFuncGetChildren()
}

// FuncGetChildren returns the function "WEBEXT.bookmarks.getChildren".
func FuncGetChildren() (fn js.Func[func(id js.String) js.Promise[js.Array[BookmarkTreeNode]]]) {
	bindings.FuncGetChildren(
		js.Pointer(&fn),
	)
	return
}

// GetChildren calls the function "WEBEXT.bookmarks.getChildren" directly.
func GetChildren(id js.String) (ret js.Promise[js.Array[BookmarkTreeNode]]) {
	bindings.CallGetChildren(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryGetChildren calls the function "WEBEXT.bookmarks.getChildren"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetChildren(id js.String) (ret js.Promise[js.Array[BookmarkTreeNode]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetChildren(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncGetRecent returns true if the function "WEBEXT.bookmarks.getRecent" exists.
func HasFuncGetRecent() bool {
	return js.True == bindings.HasFuncGetRecent()
}

// FuncGetRecent returns the function "WEBEXT.bookmarks.getRecent".
func FuncGetRecent() (fn js.Func[func(numberOfItems int64) js.Promise[js.Array[BookmarkTreeNode]]]) {
	bindings.FuncGetRecent(
		js.Pointer(&fn),
	)
	return
}

// GetRecent calls the function "WEBEXT.bookmarks.getRecent" directly.
func GetRecent(numberOfItems int64) (ret js.Promise[js.Array[BookmarkTreeNode]]) {
	bindings.CallGetRecent(
		js.Pointer(&ret),
		float64(numberOfItems),
	)

	return
}

// TryGetRecent calls the function "WEBEXT.bookmarks.getRecent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetRecent(numberOfItems int64) (ret js.Promise[js.Array[BookmarkTreeNode]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetRecent(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(numberOfItems),
	)

	return
}

// HasFuncGetSubTree returns true if the function "WEBEXT.bookmarks.getSubTree" exists.
func HasFuncGetSubTree() bool {
	return js.True == bindings.HasFuncGetSubTree()
}

// FuncGetSubTree returns the function "WEBEXT.bookmarks.getSubTree".
func FuncGetSubTree() (fn js.Func[func(id js.String) js.Promise[js.Array[BookmarkTreeNode]]]) {
	bindings.FuncGetSubTree(
		js.Pointer(&fn),
	)
	return
}

// GetSubTree calls the function "WEBEXT.bookmarks.getSubTree" directly.
func GetSubTree(id js.String) (ret js.Promise[js.Array[BookmarkTreeNode]]) {
	bindings.CallGetSubTree(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryGetSubTree calls the function "WEBEXT.bookmarks.getSubTree"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetSubTree(id js.String) (ret js.Promise[js.Array[BookmarkTreeNode]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetSubTree(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncGetTree returns true if the function "WEBEXT.bookmarks.getTree" exists.
func HasFuncGetTree() bool {
	return js.True == bindings.HasFuncGetTree()
}

// FuncGetTree returns the function "WEBEXT.bookmarks.getTree".
func FuncGetTree() (fn js.Func[func() js.Promise[js.Array[BookmarkTreeNode]]]) {
	bindings.FuncGetTree(
		js.Pointer(&fn),
	)
	return
}

// GetTree calls the function "WEBEXT.bookmarks.getTree" directly.
func GetTree() (ret js.Promise[js.Array[BookmarkTreeNode]]) {
	bindings.CallGetTree(
		js.Pointer(&ret),
	)

	return
}

// TryGetTree calls the function "WEBEXT.bookmarks.getTree"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetTree() (ret js.Promise[js.Array[BookmarkTreeNode]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetTree(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMove returns true if the function "WEBEXT.bookmarks.move" exists.
func HasFuncMove() bool {
	return js.True == bindings.HasFuncMove()
}

// FuncMove returns the function "WEBEXT.bookmarks.move".
func FuncMove() (fn js.Func[func(id js.String, destination MoveArgDestination) js.Promise[BookmarkTreeNode]]) {
	bindings.FuncMove(
		js.Pointer(&fn),
	)
	return
}

// Move calls the function "WEBEXT.bookmarks.move" directly.
func Move(id js.String, destination MoveArgDestination) (ret js.Promise[BookmarkTreeNode]) {
	bindings.CallMove(
		js.Pointer(&ret),
		id.Ref(),
		js.Pointer(&destination),
	)

	return
}

// TryMove calls the function "WEBEXT.bookmarks.move"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryMove(id js.String, destination MoveArgDestination) (ret js.Promise[BookmarkTreeNode], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMove(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		js.Pointer(&destination),
	)

	return
}

type OnChangedEventCallbackFunc func(this js.Ref, id js.String, changeInfo *OnChangedArgChangeInfo) js.Ref

func (fn OnChangedEventCallbackFunc) Register() js.Func[func(id js.String, changeInfo *OnChangedArgChangeInfo)] {
	return js.RegisterCallback[func(id js.String, changeInfo *OnChangedArgChangeInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 OnChangedArgChangeInfo
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, id js.String, changeInfo *OnChangedArgChangeInfo) js.Ref
	Arg T
}

func (cb *OnChangedEventCallback[T]) Register() js.Func[func(id js.String, changeInfo *OnChangedArgChangeInfo)] {
	return js.RegisterCallback[func(id js.String, changeInfo *OnChangedArgChangeInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 OnChangedArgChangeInfo
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnChanged returns true if the function "WEBEXT.bookmarks.onChanged.addListener" exists.
func HasFuncOnChanged() bool {
	return js.True == bindings.HasFuncOnChanged()
}

// FuncOnChanged returns the function "WEBEXT.bookmarks.onChanged.addListener".
func FuncOnChanged() (fn js.Func[func(callback js.Func[func(id js.String, changeInfo *OnChangedArgChangeInfo)])]) {
	bindings.FuncOnChanged(
		js.Pointer(&fn),
	)
	return
}

// OnChanged calls the function "WEBEXT.bookmarks.onChanged.addListener" directly.
func OnChanged(callback js.Func[func(id js.String, changeInfo *OnChangedArgChangeInfo)]) (ret js.Void) {
	bindings.CallOnChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnChanged calls the function "WEBEXT.bookmarks.onChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnChanged(callback js.Func[func(id js.String, changeInfo *OnChangedArgChangeInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffChanged returns true if the function "WEBEXT.bookmarks.onChanged.removeListener" exists.
func HasFuncOffChanged() bool {
	return js.True == bindings.HasFuncOffChanged()
}

// FuncOffChanged returns the function "WEBEXT.bookmarks.onChanged.removeListener".
func FuncOffChanged() (fn js.Func[func(callback js.Func[func(id js.String, changeInfo *OnChangedArgChangeInfo)])]) {
	bindings.FuncOffChanged(
		js.Pointer(&fn),
	)
	return
}

// OffChanged calls the function "WEBEXT.bookmarks.onChanged.removeListener" directly.
func OffChanged(callback js.Func[func(id js.String, changeInfo *OnChangedArgChangeInfo)]) (ret js.Void) {
	bindings.CallOffChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffChanged calls the function "WEBEXT.bookmarks.onChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffChanged(callback js.Func[func(id js.String, changeInfo *OnChangedArgChangeInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnChanged returns true if the function "WEBEXT.bookmarks.onChanged.hasListener" exists.
func HasFuncHasOnChanged() bool {
	return js.True == bindings.HasFuncHasOnChanged()
}

// FuncHasOnChanged returns the function "WEBEXT.bookmarks.onChanged.hasListener".
func FuncHasOnChanged() (fn js.Func[func(callback js.Func[func(id js.String, changeInfo *OnChangedArgChangeInfo)]) bool]) {
	bindings.FuncHasOnChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnChanged calls the function "WEBEXT.bookmarks.onChanged.hasListener" directly.
func HasOnChanged(callback js.Func[func(id js.String, changeInfo *OnChangedArgChangeInfo)]) (ret bool) {
	bindings.CallHasOnChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnChanged calls the function "WEBEXT.bookmarks.onChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnChanged(callback js.Func[func(id js.String, changeInfo *OnChangedArgChangeInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnChildrenReorderedEventCallbackFunc func(this js.Ref, id js.String, reorderInfo *OnChildrenReorderedArgReorderInfo) js.Ref

func (fn OnChildrenReorderedEventCallbackFunc) Register() js.Func[func(id js.String, reorderInfo *OnChildrenReorderedArgReorderInfo)] {
	return js.RegisterCallback[func(id js.String, reorderInfo *OnChildrenReorderedArgReorderInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnChildrenReorderedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 OnChildrenReorderedArgReorderInfo
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnChildrenReorderedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, id js.String, reorderInfo *OnChildrenReorderedArgReorderInfo) js.Ref
	Arg T
}

func (cb *OnChildrenReorderedEventCallback[T]) Register() js.Func[func(id js.String, reorderInfo *OnChildrenReorderedArgReorderInfo)] {
	return js.RegisterCallback[func(id js.String, reorderInfo *OnChildrenReorderedArgReorderInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnChildrenReorderedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 OnChildrenReorderedArgReorderInfo
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnChildrenReordered returns true if the function "WEBEXT.bookmarks.onChildrenReordered.addListener" exists.
func HasFuncOnChildrenReordered() bool {
	return js.True == bindings.HasFuncOnChildrenReordered()
}

// FuncOnChildrenReordered returns the function "WEBEXT.bookmarks.onChildrenReordered.addListener".
func FuncOnChildrenReordered() (fn js.Func[func(callback js.Func[func(id js.String, reorderInfo *OnChildrenReorderedArgReorderInfo)])]) {
	bindings.FuncOnChildrenReordered(
		js.Pointer(&fn),
	)
	return
}

// OnChildrenReordered calls the function "WEBEXT.bookmarks.onChildrenReordered.addListener" directly.
func OnChildrenReordered(callback js.Func[func(id js.String, reorderInfo *OnChildrenReorderedArgReorderInfo)]) (ret js.Void) {
	bindings.CallOnChildrenReordered(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnChildrenReordered calls the function "WEBEXT.bookmarks.onChildrenReordered.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnChildrenReordered(callback js.Func[func(id js.String, reorderInfo *OnChildrenReorderedArgReorderInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnChildrenReordered(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffChildrenReordered returns true if the function "WEBEXT.bookmarks.onChildrenReordered.removeListener" exists.
func HasFuncOffChildrenReordered() bool {
	return js.True == bindings.HasFuncOffChildrenReordered()
}

// FuncOffChildrenReordered returns the function "WEBEXT.bookmarks.onChildrenReordered.removeListener".
func FuncOffChildrenReordered() (fn js.Func[func(callback js.Func[func(id js.String, reorderInfo *OnChildrenReorderedArgReorderInfo)])]) {
	bindings.FuncOffChildrenReordered(
		js.Pointer(&fn),
	)
	return
}

// OffChildrenReordered calls the function "WEBEXT.bookmarks.onChildrenReordered.removeListener" directly.
func OffChildrenReordered(callback js.Func[func(id js.String, reorderInfo *OnChildrenReorderedArgReorderInfo)]) (ret js.Void) {
	bindings.CallOffChildrenReordered(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffChildrenReordered calls the function "WEBEXT.bookmarks.onChildrenReordered.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffChildrenReordered(callback js.Func[func(id js.String, reorderInfo *OnChildrenReorderedArgReorderInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffChildrenReordered(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnChildrenReordered returns true if the function "WEBEXT.bookmarks.onChildrenReordered.hasListener" exists.
func HasFuncHasOnChildrenReordered() bool {
	return js.True == bindings.HasFuncHasOnChildrenReordered()
}

// FuncHasOnChildrenReordered returns the function "WEBEXT.bookmarks.onChildrenReordered.hasListener".
func FuncHasOnChildrenReordered() (fn js.Func[func(callback js.Func[func(id js.String, reorderInfo *OnChildrenReorderedArgReorderInfo)]) bool]) {
	bindings.FuncHasOnChildrenReordered(
		js.Pointer(&fn),
	)
	return
}

// HasOnChildrenReordered calls the function "WEBEXT.bookmarks.onChildrenReordered.hasListener" directly.
func HasOnChildrenReordered(callback js.Func[func(id js.String, reorderInfo *OnChildrenReorderedArgReorderInfo)]) (ret bool) {
	bindings.CallHasOnChildrenReordered(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnChildrenReordered calls the function "WEBEXT.bookmarks.onChildrenReordered.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnChildrenReordered(callback js.Func[func(id js.String, reorderInfo *OnChildrenReorderedArgReorderInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnChildrenReordered(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnCreatedEventCallbackFunc func(this js.Ref, id js.String, bookmark *BookmarkTreeNode) js.Ref

func (fn OnCreatedEventCallbackFunc) Register() js.Func[func(id js.String, bookmark *BookmarkTreeNode)] {
	return js.RegisterCallback[func(id js.String, bookmark *BookmarkTreeNode)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCreatedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 BookmarkTreeNode
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnCreatedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, id js.String, bookmark *BookmarkTreeNode) js.Ref
	Arg T
}

func (cb *OnCreatedEventCallback[T]) Register() js.Func[func(id js.String, bookmark *BookmarkTreeNode)] {
	return js.RegisterCallback[func(id js.String, bookmark *BookmarkTreeNode)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCreatedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 BookmarkTreeNode
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnCreated returns true if the function "WEBEXT.bookmarks.onCreated.addListener" exists.
func HasFuncOnCreated() bool {
	return js.True == bindings.HasFuncOnCreated()
}

// FuncOnCreated returns the function "WEBEXT.bookmarks.onCreated.addListener".
func FuncOnCreated() (fn js.Func[func(callback js.Func[func(id js.String, bookmark *BookmarkTreeNode)])]) {
	bindings.FuncOnCreated(
		js.Pointer(&fn),
	)
	return
}

// OnCreated calls the function "WEBEXT.bookmarks.onCreated.addListener" directly.
func OnCreated(callback js.Func[func(id js.String, bookmark *BookmarkTreeNode)]) (ret js.Void) {
	bindings.CallOnCreated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCreated calls the function "WEBEXT.bookmarks.onCreated.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCreated(callback js.Func[func(id js.String, bookmark *BookmarkTreeNode)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCreated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCreated returns true if the function "WEBEXT.bookmarks.onCreated.removeListener" exists.
func HasFuncOffCreated() bool {
	return js.True == bindings.HasFuncOffCreated()
}

// FuncOffCreated returns the function "WEBEXT.bookmarks.onCreated.removeListener".
func FuncOffCreated() (fn js.Func[func(callback js.Func[func(id js.String, bookmark *BookmarkTreeNode)])]) {
	bindings.FuncOffCreated(
		js.Pointer(&fn),
	)
	return
}

// OffCreated calls the function "WEBEXT.bookmarks.onCreated.removeListener" directly.
func OffCreated(callback js.Func[func(id js.String, bookmark *BookmarkTreeNode)]) (ret js.Void) {
	bindings.CallOffCreated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCreated calls the function "WEBEXT.bookmarks.onCreated.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCreated(callback js.Func[func(id js.String, bookmark *BookmarkTreeNode)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCreated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCreated returns true if the function "WEBEXT.bookmarks.onCreated.hasListener" exists.
func HasFuncHasOnCreated() bool {
	return js.True == bindings.HasFuncHasOnCreated()
}

// FuncHasOnCreated returns the function "WEBEXT.bookmarks.onCreated.hasListener".
func FuncHasOnCreated() (fn js.Func[func(callback js.Func[func(id js.String, bookmark *BookmarkTreeNode)]) bool]) {
	bindings.FuncHasOnCreated(
		js.Pointer(&fn),
	)
	return
}

// HasOnCreated calls the function "WEBEXT.bookmarks.onCreated.hasListener" directly.
func HasOnCreated(callback js.Func[func(id js.String, bookmark *BookmarkTreeNode)]) (ret bool) {
	bindings.CallHasOnCreated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCreated calls the function "WEBEXT.bookmarks.onCreated.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCreated(callback js.Func[func(id js.String, bookmark *BookmarkTreeNode)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCreated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnImportBeganEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnImportBeganEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnImportBeganEventCallbackFunc) DispatchCallback(
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

type OnImportBeganEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnImportBeganEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnImportBeganEventCallback[T]) DispatchCallback(
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

// HasFuncOnImportBegan returns true if the function "WEBEXT.bookmarks.onImportBegan.addListener" exists.
func HasFuncOnImportBegan() bool {
	return js.True == bindings.HasFuncOnImportBegan()
}

// FuncOnImportBegan returns the function "WEBEXT.bookmarks.onImportBegan.addListener".
func FuncOnImportBegan() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnImportBegan(
		js.Pointer(&fn),
	)
	return
}

// OnImportBegan calls the function "WEBEXT.bookmarks.onImportBegan.addListener" directly.
func OnImportBegan(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnImportBegan(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnImportBegan calls the function "WEBEXT.bookmarks.onImportBegan.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnImportBegan(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnImportBegan(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffImportBegan returns true if the function "WEBEXT.bookmarks.onImportBegan.removeListener" exists.
func HasFuncOffImportBegan() bool {
	return js.True == bindings.HasFuncOffImportBegan()
}

// FuncOffImportBegan returns the function "WEBEXT.bookmarks.onImportBegan.removeListener".
func FuncOffImportBegan() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffImportBegan(
		js.Pointer(&fn),
	)
	return
}

// OffImportBegan calls the function "WEBEXT.bookmarks.onImportBegan.removeListener" directly.
func OffImportBegan(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffImportBegan(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffImportBegan calls the function "WEBEXT.bookmarks.onImportBegan.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffImportBegan(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffImportBegan(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnImportBegan returns true if the function "WEBEXT.bookmarks.onImportBegan.hasListener" exists.
func HasFuncHasOnImportBegan() bool {
	return js.True == bindings.HasFuncHasOnImportBegan()
}

// FuncHasOnImportBegan returns the function "WEBEXT.bookmarks.onImportBegan.hasListener".
func FuncHasOnImportBegan() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnImportBegan(
		js.Pointer(&fn),
	)
	return
}

// HasOnImportBegan calls the function "WEBEXT.bookmarks.onImportBegan.hasListener" directly.
func HasOnImportBegan(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnImportBegan(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnImportBegan calls the function "WEBEXT.bookmarks.onImportBegan.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnImportBegan(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnImportBegan(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnImportEndedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnImportEndedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnImportEndedEventCallbackFunc) DispatchCallback(
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

type OnImportEndedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnImportEndedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnImportEndedEventCallback[T]) DispatchCallback(
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

// HasFuncOnImportEnded returns true if the function "WEBEXT.bookmarks.onImportEnded.addListener" exists.
func HasFuncOnImportEnded() bool {
	return js.True == bindings.HasFuncOnImportEnded()
}

// FuncOnImportEnded returns the function "WEBEXT.bookmarks.onImportEnded.addListener".
func FuncOnImportEnded() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnImportEnded(
		js.Pointer(&fn),
	)
	return
}

// OnImportEnded calls the function "WEBEXT.bookmarks.onImportEnded.addListener" directly.
func OnImportEnded(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnImportEnded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnImportEnded calls the function "WEBEXT.bookmarks.onImportEnded.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnImportEnded(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnImportEnded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffImportEnded returns true if the function "WEBEXT.bookmarks.onImportEnded.removeListener" exists.
func HasFuncOffImportEnded() bool {
	return js.True == bindings.HasFuncOffImportEnded()
}

// FuncOffImportEnded returns the function "WEBEXT.bookmarks.onImportEnded.removeListener".
func FuncOffImportEnded() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffImportEnded(
		js.Pointer(&fn),
	)
	return
}

// OffImportEnded calls the function "WEBEXT.bookmarks.onImportEnded.removeListener" directly.
func OffImportEnded(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffImportEnded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffImportEnded calls the function "WEBEXT.bookmarks.onImportEnded.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffImportEnded(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffImportEnded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnImportEnded returns true if the function "WEBEXT.bookmarks.onImportEnded.hasListener" exists.
func HasFuncHasOnImportEnded() bool {
	return js.True == bindings.HasFuncHasOnImportEnded()
}

// FuncHasOnImportEnded returns the function "WEBEXT.bookmarks.onImportEnded.hasListener".
func FuncHasOnImportEnded() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnImportEnded(
		js.Pointer(&fn),
	)
	return
}

// HasOnImportEnded calls the function "WEBEXT.bookmarks.onImportEnded.hasListener" directly.
func HasOnImportEnded(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnImportEnded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnImportEnded calls the function "WEBEXT.bookmarks.onImportEnded.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnImportEnded(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnImportEnded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnMovedEventCallbackFunc func(this js.Ref, id js.String, moveInfo *OnMovedArgMoveInfo) js.Ref

func (fn OnMovedEventCallbackFunc) Register() js.Func[func(id js.String, moveInfo *OnMovedArgMoveInfo)] {
	return js.RegisterCallback[func(id js.String, moveInfo *OnMovedArgMoveInfo)](
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

		js.String{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnMovedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, id js.String, moveInfo *OnMovedArgMoveInfo) js.Ref
	Arg T
}

func (cb *OnMovedEventCallback[T]) Register() js.Func[func(id js.String, moveInfo *OnMovedArgMoveInfo)] {
	return js.RegisterCallback[func(id js.String, moveInfo *OnMovedArgMoveInfo)](
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

		js.String{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnMoved returns true if the function "WEBEXT.bookmarks.onMoved.addListener" exists.
func HasFuncOnMoved() bool {
	return js.True == bindings.HasFuncOnMoved()
}

// FuncOnMoved returns the function "WEBEXT.bookmarks.onMoved.addListener".
func FuncOnMoved() (fn js.Func[func(callback js.Func[func(id js.String, moveInfo *OnMovedArgMoveInfo)])]) {
	bindings.FuncOnMoved(
		js.Pointer(&fn),
	)
	return
}

// OnMoved calls the function "WEBEXT.bookmarks.onMoved.addListener" directly.
func OnMoved(callback js.Func[func(id js.String, moveInfo *OnMovedArgMoveInfo)]) (ret js.Void) {
	bindings.CallOnMoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMoved calls the function "WEBEXT.bookmarks.onMoved.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMoved(callback js.Func[func(id js.String, moveInfo *OnMovedArgMoveInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMoved returns true if the function "WEBEXT.bookmarks.onMoved.removeListener" exists.
func HasFuncOffMoved() bool {
	return js.True == bindings.HasFuncOffMoved()
}

// FuncOffMoved returns the function "WEBEXT.bookmarks.onMoved.removeListener".
func FuncOffMoved() (fn js.Func[func(callback js.Func[func(id js.String, moveInfo *OnMovedArgMoveInfo)])]) {
	bindings.FuncOffMoved(
		js.Pointer(&fn),
	)
	return
}

// OffMoved calls the function "WEBEXT.bookmarks.onMoved.removeListener" directly.
func OffMoved(callback js.Func[func(id js.String, moveInfo *OnMovedArgMoveInfo)]) (ret js.Void) {
	bindings.CallOffMoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMoved calls the function "WEBEXT.bookmarks.onMoved.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMoved(callback js.Func[func(id js.String, moveInfo *OnMovedArgMoveInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMoved returns true if the function "WEBEXT.bookmarks.onMoved.hasListener" exists.
func HasFuncHasOnMoved() bool {
	return js.True == bindings.HasFuncHasOnMoved()
}

// FuncHasOnMoved returns the function "WEBEXT.bookmarks.onMoved.hasListener".
func FuncHasOnMoved() (fn js.Func[func(callback js.Func[func(id js.String, moveInfo *OnMovedArgMoveInfo)]) bool]) {
	bindings.FuncHasOnMoved(
		js.Pointer(&fn),
	)
	return
}

// HasOnMoved calls the function "WEBEXT.bookmarks.onMoved.hasListener" directly.
func HasOnMoved(callback js.Func[func(id js.String, moveInfo *OnMovedArgMoveInfo)]) (ret bool) {
	bindings.CallHasOnMoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMoved calls the function "WEBEXT.bookmarks.onMoved.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnMoved(callback js.Func[func(id js.String, moveInfo *OnMovedArgMoveInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnMoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnRemovedEventCallbackFunc func(this js.Ref, id js.String, removeInfo *OnRemovedArgRemoveInfo) js.Ref

func (fn OnRemovedEventCallbackFunc) Register() js.Func[func(id js.String, removeInfo *OnRemovedArgRemoveInfo)] {
	return js.RegisterCallback[func(id js.String, removeInfo *OnRemovedArgRemoveInfo)](
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

		js.String{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnRemovedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, id js.String, removeInfo *OnRemovedArgRemoveInfo) js.Ref
	Arg T
}

func (cb *OnRemovedEventCallback[T]) Register() js.Func[func(id js.String, removeInfo *OnRemovedArgRemoveInfo)] {
	return js.RegisterCallback[func(id js.String, removeInfo *OnRemovedArgRemoveInfo)](
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

		js.String{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnRemoved returns true if the function "WEBEXT.bookmarks.onRemoved.addListener" exists.
func HasFuncOnRemoved() bool {
	return js.True == bindings.HasFuncOnRemoved()
}

// FuncOnRemoved returns the function "WEBEXT.bookmarks.onRemoved.addListener".
func FuncOnRemoved() (fn js.Func[func(callback js.Func[func(id js.String, removeInfo *OnRemovedArgRemoveInfo)])]) {
	bindings.FuncOnRemoved(
		js.Pointer(&fn),
	)
	return
}

// OnRemoved calls the function "WEBEXT.bookmarks.onRemoved.addListener" directly.
func OnRemoved(callback js.Func[func(id js.String, removeInfo *OnRemovedArgRemoveInfo)]) (ret js.Void) {
	bindings.CallOnRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnRemoved calls the function "WEBEXT.bookmarks.onRemoved.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnRemoved(callback js.Func[func(id js.String, removeInfo *OnRemovedArgRemoveInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffRemoved returns true if the function "WEBEXT.bookmarks.onRemoved.removeListener" exists.
func HasFuncOffRemoved() bool {
	return js.True == bindings.HasFuncOffRemoved()
}

// FuncOffRemoved returns the function "WEBEXT.bookmarks.onRemoved.removeListener".
func FuncOffRemoved() (fn js.Func[func(callback js.Func[func(id js.String, removeInfo *OnRemovedArgRemoveInfo)])]) {
	bindings.FuncOffRemoved(
		js.Pointer(&fn),
	)
	return
}

// OffRemoved calls the function "WEBEXT.bookmarks.onRemoved.removeListener" directly.
func OffRemoved(callback js.Func[func(id js.String, removeInfo *OnRemovedArgRemoveInfo)]) (ret js.Void) {
	bindings.CallOffRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffRemoved calls the function "WEBEXT.bookmarks.onRemoved.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffRemoved(callback js.Func[func(id js.String, removeInfo *OnRemovedArgRemoveInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnRemoved returns true if the function "WEBEXT.bookmarks.onRemoved.hasListener" exists.
func HasFuncHasOnRemoved() bool {
	return js.True == bindings.HasFuncHasOnRemoved()
}

// FuncHasOnRemoved returns the function "WEBEXT.bookmarks.onRemoved.hasListener".
func FuncHasOnRemoved() (fn js.Func[func(callback js.Func[func(id js.String, removeInfo *OnRemovedArgRemoveInfo)]) bool]) {
	bindings.FuncHasOnRemoved(
		js.Pointer(&fn),
	)
	return
}

// HasOnRemoved calls the function "WEBEXT.bookmarks.onRemoved.hasListener" directly.
func HasOnRemoved(callback js.Func[func(id js.String, removeInfo *OnRemovedArgRemoveInfo)]) (ret bool) {
	bindings.CallHasOnRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnRemoved calls the function "WEBEXT.bookmarks.onRemoved.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnRemoved(callback js.Func[func(id js.String, removeInfo *OnRemovedArgRemoveInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncRemove returns true if the function "WEBEXT.bookmarks.remove" exists.
func HasFuncRemove() bool {
	return js.True == bindings.HasFuncRemove()
}

// FuncRemove returns the function "WEBEXT.bookmarks.remove".
func FuncRemove() (fn js.Func[func(id js.String) js.Promise[js.Void]]) {
	bindings.FuncRemove(
		js.Pointer(&fn),
	)
	return
}

// Remove calls the function "WEBEXT.bookmarks.remove" directly.
func Remove(id js.String) (ret js.Promise[js.Void]) {
	bindings.CallRemove(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryRemove calls the function "WEBEXT.bookmarks.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemove(id js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemove(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncRemoveTree returns true if the function "WEBEXT.bookmarks.removeTree" exists.
func HasFuncRemoveTree() bool {
	return js.True == bindings.HasFuncRemoveTree()
}

// FuncRemoveTree returns the function "WEBEXT.bookmarks.removeTree".
func FuncRemoveTree() (fn js.Func[func(id js.String) js.Promise[js.Void]]) {
	bindings.FuncRemoveTree(
		js.Pointer(&fn),
	)
	return
}

// RemoveTree calls the function "WEBEXT.bookmarks.removeTree" directly.
func RemoveTree(id js.String) (ret js.Promise[js.Void]) {
	bindings.CallRemoveTree(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryRemoveTree calls the function "WEBEXT.bookmarks.removeTree"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveTree(id js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveTree(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

type OneOf_String_SearchArgQueryChoice1 struct {
	ref js.Ref
}

func (x OneOf_String_SearchArgQueryChoice1) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_SearchArgQueryChoice1) Free() {
	x.ref.Free()
}

func (x OneOf_String_SearchArgQueryChoice1) FromRef(ref js.Ref) OneOf_String_SearchArgQueryChoice1 {
	return OneOf_String_SearchArgQueryChoice1{
		ref: ref,
	}
}

func (x OneOf_String_SearchArgQueryChoice1) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_SearchArgQueryChoice1) SearchArgQueryChoice1() SearchArgQueryChoice1 {
	var ret SearchArgQueryChoice1
	ret.UpdateFrom(x.ref)
	return ret
}

// HasFuncSearch returns true if the function "WEBEXT.bookmarks.search" exists.
func HasFuncSearch() bool {
	return js.True == bindings.HasFuncSearch()
}

// FuncSearch returns the function "WEBEXT.bookmarks.search".
func FuncSearch() (fn js.Func[func(query OneOf_String_SearchArgQueryChoice1) js.Promise[js.Array[BookmarkTreeNode]]]) {
	bindings.FuncSearch(
		js.Pointer(&fn),
	)
	return
}

// Search calls the function "WEBEXT.bookmarks.search" directly.
func Search(query OneOf_String_SearchArgQueryChoice1) (ret js.Promise[js.Array[BookmarkTreeNode]]) {
	bindings.CallSearch(
		js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TrySearch calls the function "WEBEXT.bookmarks.search"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySearch(query OneOf_String_SearchArgQueryChoice1) (ret js.Promise[js.Array[BookmarkTreeNode]], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySearch(
		js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasFuncUpdate returns true if the function "WEBEXT.bookmarks.update" exists.
func HasFuncUpdate() bool {
	return js.True == bindings.HasFuncUpdate()
}

// FuncUpdate returns the function "WEBEXT.bookmarks.update".
func FuncUpdate() (fn js.Func[func(id js.String, changes UpdateArgChanges) js.Promise[BookmarkTreeNode]]) {
	bindings.FuncUpdate(
		js.Pointer(&fn),
	)
	return
}

// Update calls the function "WEBEXT.bookmarks.update" directly.
func Update(id js.String, changes UpdateArgChanges) (ret js.Promise[BookmarkTreeNode]) {
	bindings.CallUpdate(
		js.Pointer(&ret),
		id.Ref(),
		js.Pointer(&changes),
	)

	return
}

// TryUpdate calls the function "WEBEXT.bookmarks.update"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdate(id js.String, changes UpdateArgChanges) (ret js.Promise[BookmarkTreeNode], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdate(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		js.Pointer(&changes),
	)

	return
}
