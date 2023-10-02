// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/ffi/js"
	"github.com/primecitizens/std/plat/js/web/bindings"
)

func _() {
	var (
		_ abi.FuncID
	)
	assert.TODO()
}

type SyncEventInit struct {
	// Tag is "SyncEventInit.tag"
	//
	// Required
	Tag js.String
	// LastChance is "SyncEventInit.lastChance"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_LastChance MUST be set to true to make this field effective.
	LastChance bool
	// Bubbles is "SyncEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "SyncEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "SyncEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_LastChance bool // for LastChance.
	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SyncEventInit with all fields set.
func (p SyncEventInit) FromRef(ref js.Ref) SyncEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SyncEventInit in the application heap.
func (p SyncEventInit) New() js.Ref {
	return bindings.SyncEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SyncEventInit) UpdateFrom(ref js.Ref) {
	bindings.SyncEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SyncEventInit) Update(ref js.Ref) {
	bindings.SyncEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewSyncEvent(typ js.String, init SyncEventInit) SyncEvent {
	return SyncEvent{}.FromRef(
		bindings.NewSyncEventBySyncEvent(
			typ.Ref(),
			js.Pointer(&init)),
	)
}

type SyncEvent struct {
	ExtendableEvent
}

func (this SyncEvent) Once() SyncEvent {
	this.Ref().Once()
	return this
}

func (this SyncEvent) Ref() js.Ref {
	return this.ExtendableEvent.Ref()
}

func (this SyncEvent) FromRef(ref js.Ref) SyncEvent {
	this.ExtendableEvent = this.ExtendableEvent.FromRef(ref)
	return this
}

func (this SyncEvent) Free() {
	this.Ref().Free()
}

// Tag returns the value of property "SyncEvent.tag".
//
// The returned bool will be false if there is no such property.
func (this SyncEvent) Tag() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSyncEventTag(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// LastChance returns the value of property "SyncEvent.lastChance".
//
// The returned bool will be false if there is no such property.
func (this SyncEvent) LastChance() (bool, bool) {
	var _ok bool
	_ret := bindings.GetSyncEventLastChance(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

type TableKind uint32

const (
	_ TableKind = iota

	TableKind_EXTERNREF
	TableKind_ANYFUNC
)

func (TableKind) FromRef(str js.Ref) TableKind {
	return TableKind(bindings.ConstOfTableKind(str))
}

func (x TableKind) String() (string, bool) {
	switch x {
	case TableKind_EXTERNREF:
		return "externref", true
	case TableKind_ANYFUNC:
		return "anyfunc", true
	default:
		return "", false
	}
}

type TableDescriptor struct {
	// Element is "TableDescriptor.element"
	//
	// Required
	Element TableKind
	// Initial is "TableDescriptor.initial"
	//
	// Required
	Initial uint32
	// Maximum is "TableDescriptor.maximum"
	//
	// Optional
	//
	// NOTE: FFI_USE_Maximum MUST be set to true to make this field effective.
	Maximum uint32

	FFI_USE_Maximum bool // for Maximum.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TableDescriptor with all fields set.
func (p TableDescriptor) FromRef(ref js.Ref) TableDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TableDescriptor in the application heap.
func (p TableDescriptor) New() js.Ref {
	return bindings.TableDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p TableDescriptor) UpdateFrom(ref js.Ref) {
	bindings.TableDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p TableDescriptor) Update(ref js.Ref) {
	bindings.TableDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewTable(descriptor TableDescriptor, value js.Any) Table {
	return Table{}.FromRef(
		bindings.NewTableByTable(
			js.Pointer(&descriptor),
			value.Ref()),
	)
}

func NewTableByTable1(descriptor TableDescriptor) Table {
	return Table{}.FromRef(
		bindings.NewTableByTable1(
			js.Pointer(&descriptor)),
	)
}

type Table struct {
	ref js.Ref
}

func (this Table) Once() Table {
	this.Ref().Once()
	return this
}

func (this Table) Ref() js.Ref {
	return this.ref
}

func (this Table) FromRef(ref js.Ref) Table {
	this.ref = ref
	return this
}

func (this Table) Free() {
	this.Ref().Free()
}

// Length returns the value of property "Table.length".
//
// The returned bool will be false if there is no such property.
func (this Table) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetTableLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Grow calls the method "Table.grow".
//
// The returned bool will be false if there is no such method.
func (this Table) Grow(delta uint32, value js.Any) (uint32, bool) {
	var _ok bool
	_ret := bindings.CallTableGrow(
		this.Ref(), js.Pointer(&_ok),
		uint32(delta),
		value.Ref(),
	)

	return uint32(_ret), _ok
}

// GrowFunc returns the method "Table.grow".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Table) GrowFunc() (fn js.Func[func(delta uint32, value js.Any) uint32]) {
	return fn.FromRef(
		bindings.TableGrowFunc(
			this.Ref(),
		),
	)
}

// Grow1 calls the method "Table.grow".
//
// The returned bool will be false if there is no such method.
func (this Table) Grow1(delta uint32) (uint32, bool) {
	var _ok bool
	_ret := bindings.CallTableGrow1(
		this.Ref(), js.Pointer(&_ok),
		uint32(delta),
	)

	return uint32(_ret), _ok
}

// Grow1Func returns the method "Table.grow".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Table) Grow1Func() (fn js.Func[func(delta uint32) uint32]) {
	return fn.FromRef(
		bindings.TableGrow1Func(
			this.Ref(),
		),
	)
}

// Get calls the method "Table.get".
//
// The returned bool will be false if there is no such method.
func (this Table) Get(index uint32) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallTableGet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetFunc returns the method "Table.get".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Table) GetFunc() (fn js.Func[func(index uint32) js.Any]) {
	return fn.FromRef(
		bindings.TableGetFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "Table.set".
//
// The returned bool will be false if there is no such method.
func (this Table) Set(index uint32, value js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallTableSet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		value.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetFunc returns the method "Table.set".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Table) SetFunc() (fn js.Func[func(index uint32, value js.Any)]) {
	return fn.FromRef(
		bindings.TableSetFunc(
			this.Ref(),
		),
	)
}

// Set1 calls the method "Table.set".
//
// The returned bool will be false if there is no such method.
func (this Table) Set1(index uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallTableSet1(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Set1Func returns the method "Table.set".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Table) Set1Func() (fn js.Func[func(index uint32)]) {
	return fn.FromRef(
		bindings.TableSet1Func(
			this.Ref(),
		),
	)
}

type TaskControllerInit struct {
	// Priority is "TaskControllerInit.priority"
	//
	// Optional, defaults to "user-visible".
	Priority TaskPriority

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TaskControllerInit with all fields set.
func (p TaskControllerInit) FromRef(ref js.Ref) TaskControllerInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TaskControllerInit in the application heap.
func (p TaskControllerInit) New() js.Ref {
	return bindings.TaskControllerInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p TaskControllerInit) UpdateFrom(ref js.Ref) {
	bindings.TaskControllerInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p TaskControllerInit) Update(ref js.Ref) {
	bindings.TaskControllerInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewTaskController(init TaskControllerInit) TaskController {
	return TaskController{}.FromRef(
		bindings.NewTaskControllerByTaskController(
			js.Pointer(&init)),
	)
}

func NewTaskControllerByTaskController1() TaskController {
	return TaskController{}.FromRef(
		bindings.NewTaskControllerByTaskController1(),
	)
}

type TaskController struct {
	AbortController
}

func (this TaskController) Once() TaskController {
	this.Ref().Once()
	return this
}

func (this TaskController) Ref() js.Ref {
	return this.AbortController.Ref()
}

func (this TaskController) FromRef(ref js.Ref) TaskController {
	this.AbortController = this.AbortController.FromRef(ref)
	return this
}

func (this TaskController) Free() {
	this.Ref().Free()
}

// SetPriority calls the method "TaskController.setPriority".
//
// The returned bool will be false if there is no such method.
func (this TaskController) SetPriority(priority TaskPriority) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallTaskControllerSetPriority(
		this.Ref(), js.Pointer(&_ok),
		uint32(priority),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetPriorityFunc returns the method "TaskController.setPriority".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TaskController) SetPriorityFunc() (fn js.Func[func(priority TaskPriority)]) {
	return fn.FromRef(
		bindings.TaskControllerSetPriorityFunc(
			this.Ref(),
		),
	)
}

type TaskPriorityChangeEventInit struct {
	// PreviousPriority is "TaskPriorityChangeEventInit.previousPriority"
	//
	// Required
	PreviousPriority TaskPriority
	// Bubbles is "TaskPriorityChangeEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "TaskPriorityChangeEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "TaskPriorityChangeEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TaskPriorityChangeEventInit with all fields set.
func (p TaskPriorityChangeEventInit) FromRef(ref js.Ref) TaskPriorityChangeEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TaskPriorityChangeEventInit in the application heap.
func (p TaskPriorityChangeEventInit) New() js.Ref {
	return bindings.TaskPriorityChangeEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p TaskPriorityChangeEventInit) UpdateFrom(ref js.Ref) {
	bindings.TaskPriorityChangeEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p TaskPriorityChangeEventInit) Update(ref js.Ref) {
	bindings.TaskPriorityChangeEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewTaskPriorityChangeEvent(typ js.String, priorityChangeEventInitDict TaskPriorityChangeEventInit) TaskPriorityChangeEvent {
	return TaskPriorityChangeEvent{}.FromRef(
		bindings.NewTaskPriorityChangeEventByTaskPriorityChangeEvent(
			typ.Ref(),
			js.Pointer(&priorityChangeEventInitDict)),
	)
}

type TaskPriorityChangeEvent struct {
	Event
}

func (this TaskPriorityChangeEvent) Once() TaskPriorityChangeEvent {
	this.Ref().Once()
	return this
}

func (this TaskPriorityChangeEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this TaskPriorityChangeEvent) FromRef(ref js.Ref) TaskPriorityChangeEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this TaskPriorityChangeEvent) Free() {
	this.Ref().Free()
}

// PreviousPriority returns the value of property "TaskPriorityChangeEvent.previousPriority".
//
// The returned bool will be false if there is no such property.
func (this TaskPriorityChangeEvent) PreviousPriority() (TaskPriority, bool) {
	var _ok bool
	_ret := bindings.GetTaskPriorityChangeEventPreviousPriority(
		this.Ref(), js.Pointer(&_ok),
	)
	return TaskPriority(_ret), _ok
}

type OneOf_TaskPriority_TaskSignal struct {
	ref js.Ref
}

func (x OneOf_TaskPriority_TaskSignal) Ref() js.Ref {
	return x.ref
}

func (x OneOf_TaskPriority_TaskSignal) Free() {
	x.ref.Free()
}

func (x OneOf_TaskPriority_TaskSignal) FromRef(ref js.Ref) OneOf_TaskPriority_TaskSignal {
	return OneOf_TaskPriority_TaskSignal{
		ref: ref,
	}
}

func (x OneOf_TaskPriority_TaskSignal) TaskPriority() TaskPriority {
	return TaskPriority(0).FromRef(x.ref)
}

func (x OneOf_TaskPriority_TaskSignal) TaskSignal() TaskSignal {
	return TaskSignal{}.FromRef(x.ref)
}

type TaskSignalAnyInit struct {
	// Priority is "TaskSignalAnyInit.priority"
	//
	// Optional, defaults to "user-visible".
	Priority OneOf_TaskPriority_TaskSignal

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TaskSignalAnyInit with all fields set.
func (p TaskSignalAnyInit) FromRef(ref js.Ref) TaskSignalAnyInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TaskSignalAnyInit in the application heap.
func (p TaskSignalAnyInit) New() js.Ref {
	return bindings.TaskSignalAnyInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p TaskSignalAnyInit) UpdateFrom(ref js.Ref) {
	bindings.TaskSignalAnyInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p TaskSignalAnyInit) Update(ref js.Ref) {
	bindings.TaskSignalAnyInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type TaskSignal struct {
	AbortSignal
}

func (this TaskSignal) Once() TaskSignal {
	this.Ref().Once()
	return this
}

func (this TaskSignal) Ref() js.Ref {
	return this.AbortSignal.Ref()
}

func (this TaskSignal) FromRef(ref js.Ref) TaskSignal {
	this.AbortSignal = this.AbortSignal.FromRef(ref)
	return this
}

func (this TaskSignal) Free() {
	this.Ref().Free()
}

// Priority returns the value of property "TaskSignal.priority".
//
// The returned bool will be false if there is no such property.
func (this TaskSignal) Priority() (TaskPriority, bool) {
	var _ok bool
	_ret := bindings.GetTaskSignalPriority(
		this.Ref(), js.Pointer(&_ok),
	)
	return TaskPriority(_ret), _ok
}

// Any calls the staticmethod "TaskSignal.any".
//
// The returned bool will be false if there is no such method.
func (this TaskSignal) Any(signals js.Array[AbortSignal], init TaskSignalAnyInit) (TaskSignal, bool) {
	var _ok bool
	_ret := bindings.CallTaskSignalAny(
		this.Ref(), js.Pointer(&_ok),
		signals.Ref(),
		js.Pointer(&init),
	)

	return TaskSignal{}.FromRef(_ret), _ok
}

// AnyFunc returns the staticmethod "TaskSignal.any".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TaskSignal) AnyFunc() (fn js.Func[func(signals js.Array[AbortSignal], init TaskSignalAnyInit) TaskSignal]) {
	return fn.FromRef(
		bindings.TaskSignalAnyFunc(
			this.Ref(),
		),
	)
}

// Any1 calls the staticmethod "TaskSignal.any".
//
// The returned bool will be false if there is no such method.
func (this TaskSignal) Any1(signals js.Array[AbortSignal]) (TaskSignal, bool) {
	var _ok bool
	_ret := bindings.CallTaskSignalAny1(
		this.Ref(), js.Pointer(&_ok),
		signals.Ref(),
	)

	return TaskSignal{}.FromRef(_ret), _ok
}

// Any1Func returns the staticmethod "TaskSignal.any".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TaskSignal) Any1Func() (fn js.Func[func(signals js.Array[AbortSignal]) TaskSignal]) {
	return fn.FromRef(
		bindings.TaskSignalAny1Func(
			this.Ref(),
		),
	)
}

type TestUtils struct{}

// Gc calls the function "TestUtils.gc".
//
// The returned bool will be false if there is no such method.
func (TestUtils) Gc() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallTestUtilsGc(
		js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// GcFunc returns the function "TestUtils.gc".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TestUtils) GcFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.TestUtilsGcFunc(),
	)
}

type TextDecodeOptions struct {
	// Stream is "TextDecodeOptions.stream"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Stream MUST be set to true to make this field effective.
	Stream bool

	FFI_USE_Stream bool // for Stream.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TextDecodeOptions with all fields set.
func (p TextDecodeOptions) FromRef(ref js.Ref) TextDecodeOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TextDecodeOptions in the application heap.
func (p TextDecodeOptions) New() js.Ref {
	return bindings.TextDecodeOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p TextDecodeOptions) UpdateFrom(ref js.Ref) {
	bindings.TextDecodeOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p TextDecodeOptions) Update(ref js.Ref) {
	bindings.TextDecodeOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type TextDecoderOptions struct {
	// Fatal is "TextDecoderOptions.fatal"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Fatal MUST be set to true to make this field effective.
	Fatal bool
	// IgnoreBOM is "TextDecoderOptions.ignoreBOM"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_IgnoreBOM MUST be set to true to make this field effective.
	IgnoreBOM bool

	FFI_USE_Fatal     bool // for Fatal.
	FFI_USE_IgnoreBOM bool // for IgnoreBOM.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TextDecoderOptions with all fields set.
func (p TextDecoderOptions) FromRef(ref js.Ref) TextDecoderOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TextDecoderOptions in the application heap.
func (p TextDecoderOptions) New() js.Ref {
	return bindings.TextDecoderOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p TextDecoderOptions) UpdateFrom(ref js.Ref) {
	bindings.TextDecoderOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p TextDecoderOptions) Update(ref js.Ref) {
	bindings.TextDecoderOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewTextDecoder(label js.String, options TextDecoderOptions) TextDecoder {
	return TextDecoder{}.FromRef(
		bindings.NewTextDecoderByTextDecoder(
			label.Ref(),
			js.Pointer(&options)),
	)
}

func NewTextDecoderByTextDecoder1(label js.String) TextDecoder {
	return TextDecoder{}.FromRef(
		bindings.NewTextDecoderByTextDecoder1(
			label.Ref()),
	)
}

func NewTextDecoderByTextDecoder2() TextDecoder {
	return TextDecoder{}.FromRef(
		bindings.NewTextDecoderByTextDecoder2(),
	)
}

type TextDecoder struct {
	ref js.Ref
}

func (this TextDecoder) Once() TextDecoder {
	this.Ref().Once()
	return this
}

func (this TextDecoder) Ref() js.Ref {
	return this.ref
}

func (this TextDecoder) FromRef(ref js.Ref) TextDecoder {
	this.ref = ref
	return this
}

func (this TextDecoder) Free() {
	this.Ref().Free()
}

// Encoding returns the value of property "TextDecoder.encoding".
//
// The returned bool will be false if there is no such property.
func (this TextDecoder) Encoding() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetTextDecoderEncoding(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Fatal returns the value of property "TextDecoder.fatal".
//
// The returned bool will be false if there is no such property.
func (this TextDecoder) Fatal() (bool, bool) {
	var _ok bool
	_ret := bindings.GetTextDecoderFatal(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// IgnoreBOM returns the value of property "TextDecoder.ignoreBOM".
//
// The returned bool will be false if there is no such property.
func (this TextDecoder) IgnoreBOM() (bool, bool) {
	var _ok bool
	_ret := bindings.GetTextDecoderIgnoreBOM(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Decode calls the method "TextDecoder.decode".
//
// The returned bool will be false if there is no such method.
func (this TextDecoder) Decode(input AllowSharedBufferSource, options TextDecodeOptions) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallTextDecoderDecode(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return js.String{}.FromRef(_ret), _ok
}

// DecodeFunc returns the method "TextDecoder.decode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TextDecoder) DecodeFunc() (fn js.Func[func(input AllowSharedBufferSource, options TextDecodeOptions) js.String]) {
	return fn.FromRef(
		bindings.TextDecoderDecodeFunc(
			this.Ref(),
		),
	)
}

// Decode1 calls the method "TextDecoder.decode".
//
// The returned bool will be false if there is no such method.
func (this TextDecoder) Decode1(input AllowSharedBufferSource) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallTextDecoderDecode1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// Decode1Func returns the method "TextDecoder.decode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TextDecoder) Decode1Func() (fn js.Func[func(input AllowSharedBufferSource) js.String]) {
	return fn.FromRef(
		bindings.TextDecoderDecode1Func(
			this.Ref(),
		),
	)
}

// Decode2 calls the method "TextDecoder.decode".
//
// The returned bool will be false if there is no such method.
func (this TextDecoder) Decode2() (js.String, bool) {
	var _ok bool
	_ret := bindings.CallTextDecoderDecode2(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.String{}.FromRef(_ret), _ok
}

// Decode2Func returns the method "TextDecoder.decode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TextDecoder) Decode2Func() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.TextDecoderDecode2Func(
			this.Ref(),
		),
	)
}

func NewTextDecoderStream(label js.String, options TextDecoderOptions) TextDecoderStream {
	return TextDecoderStream{}.FromRef(
		bindings.NewTextDecoderStreamByTextDecoderStream(
			label.Ref(),
			js.Pointer(&options)),
	)
}

func NewTextDecoderStreamByTextDecoderStream1(label js.String) TextDecoderStream {
	return TextDecoderStream{}.FromRef(
		bindings.NewTextDecoderStreamByTextDecoderStream1(
			label.Ref()),
	)
}

func NewTextDecoderStreamByTextDecoderStream2() TextDecoderStream {
	return TextDecoderStream{}.FromRef(
		bindings.NewTextDecoderStreamByTextDecoderStream2(),
	)
}

type TextDecoderStream struct {
	ref js.Ref
}

func (this TextDecoderStream) Once() TextDecoderStream {
	this.Ref().Once()
	return this
}

func (this TextDecoderStream) Ref() js.Ref {
	return this.ref
}

func (this TextDecoderStream) FromRef(ref js.Ref) TextDecoderStream {
	this.ref = ref
	return this
}

func (this TextDecoderStream) Free() {
	this.Ref().Free()
}

// Encoding returns the value of property "TextDecoderStream.encoding".
//
// The returned bool will be false if there is no such property.
func (this TextDecoderStream) Encoding() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetTextDecoderStreamEncoding(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Fatal returns the value of property "TextDecoderStream.fatal".
//
// The returned bool will be false if there is no such property.
func (this TextDecoderStream) Fatal() (bool, bool) {
	var _ok bool
	_ret := bindings.GetTextDecoderStreamFatal(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// IgnoreBOM returns the value of property "TextDecoderStream.ignoreBOM".
//
// The returned bool will be false if there is no such property.
func (this TextDecoderStream) IgnoreBOM() (bool, bool) {
	var _ok bool
	_ret := bindings.GetTextDecoderStreamIgnoreBOM(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Readable returns the value of property "TextDecoderStream.readable".
//
// The returned bool will be false if there is no such property.
func (this TextDecoderStream) Readable() (ReadableStream, bool) {
	var _ok bool
	_ret := bindings.GetTextDecoderStreamReadable(
		this.Ref(), js.Pointer(&_ok),
	)
	return ReadableStream{}.FromRef(_ret), _ok
}

// Writable returns the value of property "TextDecoderStream.writable".
//
// The returned bool will be false if there is no such property.
func (this TextDecoderStream) Writable() (WritableStream, bool) {
	var _ok bool
	_ret := bindings.GetTextDecoderStreamWritable(
		this.Ref(), js.Pointer(&_ok),
	)
	return WritableStream{}.FromRef(_ret), _ok
}

type TextDetector struct {
	ref js.Ref
}

func (this TextDetector) Once() TextDetector {
	this.Ref().Once()
	return this
}

func (this TextDetector) Ref() js.Ref {
	return this.ref
}

func (this TextDetector) FromRef(ref js.Ref) TextDetector {
	this.ref = ref
	return this
}

func (this TextDetector) Free() {
	this.Ref().Free()
}

// Detect calls the method "TextDetector.detect".
//
// The returned bool will be false if there is no such method.
func (this TextDetector) Detect(image ImageBitmapSource) (js.Promise[js.Array[DetectedText]], bool) {
	var _ok bool
	_ret := bindings.CallTextDetectorDetect(
		this.Ref(), js.Pointer(&_ok),
		image.Ref(),
	)

	return js.Promise[js.Array[DetectedText]]{}.FromRef(_ret), _ok
}

// DetectFunc returns the method "TextDetector.detect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TextDetector) DetectFunc() (fn js.Func[func(image ImageBitmapSource) js.Promise[js.Array[DetectedText]]]) {
	return fn.FromRef(
		bindings.TextDetectorDetectFunc(
			this.Ref(),
		),
	)
}

type TextEncoderEncodeIntoResult struct {
	// Read is "TextEncoderEncodeIntoResult.read"
	//
	// Optional
	//
	// NOTE: FFI_USE_Read MUST be set to true to make this field effective.
	Read uint64
	// Written is "TextEncoderEncodeIntoResult.written"
	//
	// Optional
	//
	// NOTE: FFI_USE_Written MUST be set to true to make this field effective.
	Written uint64

	FFI_USE_Read    bool // for Read.
	FFI_USE_Written bool // for Written.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TextEncoderEncodeIntoResult with all fields set.
func (p TextEncoderEncodeIntoResult) FromRef(ref js.Ref) TextEncoderEncodeIntoResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TextEncoderEncodeIntoResult in the application heap.
func (p TextEncoderEncodeIntoResult) New() js.Ref {
	return bindings.TextEncoderEncodeIntoResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p TextEncoderEncodeIntoResult) UpdateFrom(ref js.Ref) {
	bindings.TextEncoderEncodeIntoResultJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p TextEncoderEncodeIntoResult) Update(ref js.Ref) {
	bindings.TextEncoderEncodeIntoResultJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type TextEncoder struct {
	ref js.Ref
}

func (this TextEncoder) Once() TextEncoder {
	this.Ref().Once()
	return this
}

func (this TextEncoder) Ref() js.Ref {
	return this.ref
}

func (this TextEncoder) FromRef(ref js.Ref) TextEncoder {
	this.ref = ref
	return this
}

func (this TextEncoder) Free() {
	this.Ref().Free()
}

// Encoding returns the value of property "TextEncoder.encoding".
//
// The returned bool will be false if there is no such property.
func (this TextEncoder) Encoding() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetTextEncoderEncoding(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Encode calls the method "TextEncoder.encode".
//
// The returned bool will be false if there is no such method.
func (this TextEncoder) Encode(input js.String) (js.TypedArray[uint8], bool) {
	var _ok bool
	_ret := bindings.CallTextEncoderEncode(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return js.TypedArray[uint8]{}.FromRef(_ret), _ok
}

// EncodeFunc returns the method "TextEncoder.encode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TextEncoder) EncodeFunc() (fn js.Func[func(input js.String) js.TypedArray[uint8]]) {
	return fn.FromRef(
		bindings.TextEncoderEncodeFunc(
			this.Ref(),
		),
	)
}

// Encode1 calls the method "TextEncoder.encode".
//
// The returned bool will be false if there is no such method.
func (this TextEncoder) Encode1() (js.TypedArray[uint8], bool) {
	var _ok bool
	_ret := bindings.CallTextEncoderEncode1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.TypedArray[uint8]{}.FromRef(_ret), _ok
}

// Encode1Func returns the method "TextEncoder.encode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TextEncoder) Encode1Func() (fn js.Func[func() js.TypedArray[uint8]]) {
	return fn.FromRef(
		bindings.TextEncoderEncode1Func(
			this.Ref(),
		),
	)
}

// EncodeInto calls the method "TextEncoder.encodeInto".
//
// The returned bool will be false if there is no such method.
func (this TextEncoder) EncodeInto(source js.String, destination js.TypedArray[uint8]) (TextEncoderEncodeIntoResult, bool) {
	var _ret TextEncoderEncodeIntoResult
	_ok := js.True == bindings.CallTextEncoderEncodeInto(
		this.Ref(), js.Pointer(&_ret),
		source.Ref(),
		destination.Ref(),
	)

	return _ret, _ok
}

// EncodeIntoFunc returns the method "TextEncoder.encodeInto".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TextEncoder) EncodeIntoFunc() (fn js.Func[func(source js.String, destination js.TypedArray[uint8]) TextEncoderEncodeIntoResult]) {
	return fn.FromRef(
		bindings.TextEncoderEncodeIntoFunc(
			this.Ref(),
		),
	)
}

type TextEncoderStream struct {
	ref js.Ref
}

func (this TextEncoderStream) Once() TextEncoderStream {
	this.Ref().Once()
	return this
}

func (this TextEncoderStream) Ref() js.Ref {
	return this.ref
}

func (this TextEncoderStream) FromRef(ref js.Ref) TextEncoderStream {
	this.ref = ref
	return this
}

func (this TextEncoderStream) Free() {
	this.Ref().Free()
}

// Encoding returns the value of property "TextEncoderStream.encoding".
//
// The returned bool will be false if there is no such property.
func (this TextEncoderStream) Encoding() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetTextEncoderStreamEncoding(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Readable returns the value of property "TextEncoderStream.readable".
//
// The returned bool will be false if there is no such property.
func (this TextEncoderStream) Readable() (ReadableStream, bool) {
	var _ok bool
	_ret := bindings.GetTextEncoderStreamReadable(
		this.Ref(), js.Pointer(&_ok),
	)
	return ReadableStream{}.FromRef(_ret), _ok
}

// Writable returns the value of property "TextEncoderStream.writable".
//
// The returned bool will be false if there is no such property.
func (this TextEncoderStream) Writable() (WritableStream, bool) {
	var _ok bool
	_ret := bindings.GetTextEncoderStreamWritable(
		this.Ref(), js.Pointer(&_ok),
	)
	return WritableStream{}.FromRef(_ret), _ok
}

type TextFormatInit struct {
	// RangeStart is "TextFormatInit.rangeStart"
	//
	// Optional
	//
	// NOTE: FFI_USE_RangeStart MUST be set to true to make this field effective.
	RangeStart uint32
	// RangeEnd is "TextFormatInit.rangeEnd"
	//
	// Optional
	//
	// NOTE: FFI_USE_RangeEnd MUST be set to true to make this field effective.
	RangeEnd uint32
	// UnderlineStyle is "TextFormatInit.underlineStyle"
	//
	// Optional
	UnderlineStyle js.String
	// UnderlineThickness is "TextFormatInit.underlineThickness"
	//
	// Optional
	UnderlineThickness js.String

	FFI_USE_RangeStart bool // for RangeStart.
	FFI_USE_RangeEnd   bool // for RangeEnd.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TextFormatInit with all fields set.
func (p TextFormatInit) FromRef(ref js.Ref) TextFormatInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TextFormatInit in the application heap.
func (p TextFormatInit) New() js.Ref {
	return bindings.TextFormatInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p TextFormatInit) UpdateFrom(ref js.Ref) {
	bindings.TextFormatInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p TextFormatInit) Update(ref js.Ref) {
	bindings.TextFormatInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewTextFormat(options TextFormatInit) TextFormat {
	return TextFormat{}.FromRef(
		bindings.NewTextFormatByTextFormat(
			js.Pointer(&options)),
	)
}

func NewTextFormatByTextFormat1() TextFormat {
	return TextFormat{}.FromRef(
		bindings.NewTextFormatByTextFormat1(),
	)
}

type TextFormat struct {
	ref js.Ref
}

func (this TextFormat) Once() TextFormat {
	this.Ref().Once()
	return this
}

func (this TextFormat) Ref() js.Ref {
	return this.ref
}

func (this TextFormat) FromRef(ref js.Ref) TextFormat {
	this.ref = ref
	return this
}

func (this TextFormat) Free() {
	this.Ref().Free()
}

// RangeStart returns the value of property "TextFormat.rangeStart".
//
// The returned bool will be false if there is no such property.
func (this TextFormat) RangeStart() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetTextFormatRangeStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// RangeEnd returns the value of property "TextFormat.rangeEnd".
//
// The returned bool will be false if there is no such property.
func (this TextFormat) RangeEnd() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetTextFormatRangeEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// UnderlineStyle returns the value of property "TextFormat.underlineStyle".
//
// The returned bool will be false if there is no such property.
func (this TextFormat) UnderlineStyle() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetTextFormatUnderlineStyle(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// UnderlineThickness returns the value of property "TextFormat.underlineThickness".
//
// The returned bool will be false if there is no such property.
func (this TextFormat) UnderlineThickness() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetTextFormatUnderlineThickness(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type TextFormatUpdateEventInit struct {
	// TextFormats is "TextFormatUpdateEventInit.textFormats"
	//
	// Optional
	TextFormats js.Array[TextFormat]
	// Bubbles is "TextFormatUpdateEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "TextFormatUpdateEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "TextFormatUpdateEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TextFormatUpdateEventInit with all fields set.
func (p TextFormatUpdateEventInit) FromRef(ref js.Ref) TextFormatUpdateEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TextFormatUpdateEventInit in the application heap.
func (p TextFormatUpdateEventInit) New() js.Ref {
	return bindings.TextFormatUpdateEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p TextFormatUpdateEventInit) UpdateFrom(ref js.Ref) {
	bindings.TextFormatUpdateEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p TextFormatUpdateEventInit) Update(ref js.Ref) {
	bindings.TextFormatUpdateEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewTextFormatUpdateEvent(typ js.String, options TextFormatUpdateEventInit) TextFormatUpdateEvent {
	return TextFormatUpdateEvent{}.FromRef(
		bindings.NewTextFormatUpdateEventByTextFormatUpdateEvent(
			typ.Ref(),
			js.Pointer(&options)),
	)
}

func NewTextFormatUpdateEventByTextFormatUpdateEvent1(typ js.String) TextFormatUpdateEvent {
	return TextFormatUpdateEvent{}.FromRef(
		bindings.NewTextFormatUpdateEventByTextFormatUpdateEvent1(
			typ.Ref()),
	)
}

type TextFormatUpdateEvent struct {
	Event
}

func (this TextFormatUpdateEvent) Once() TextFormatUpdateEvent {
	this.Ref().Once()
	return this
}

func (this TextFormatUpdateEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this TextFormatUpdateEvent) FromRef(ref js.Ref) TextFormatUpdateEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this TextFormatUpdateEvent) Free() {
	this.Ref().Free()
}

// GetTextFormats calls the method "TextFormatUpdateEvent.getTextFormats".
//
// The returned bool will be false if there is no such method.
func (this TextFormatUpdateEvent) GetTextFormats() (js.Array[TextFormat], bool) {
	var _ok bool
	_ret := bindings.CallTextFormatUpdateEventGetTextFormats(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[TextFormat]{}.FromRef(_ret), _ok
}

// GetTextFormatsFunc returns the method "TextFormatUpdateEvent.getTextFormats".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TextFormatUpdateEvent) GetTextFormatsFunc() (fn js.Func[func() js.Array[TextFormat]]) {
	return fn.FromRef(
		bindings.TextFormatUpdateEventGetTextFormatsFunc(
			this.Ref(),
		),
	)
}

type TextUpdateEventInit struct {
	// UpdateRangeStart is "TextUpdateEventInit.updateRangeStart"
	//
	// Optional
	//
	// NOTE: FFI_USE_UpdateRangeStart MUST be set to true to make this field effective.
	UpdateRangeStart uint32
	// UpdateRangeEnd is "TextUpdateEventInit.updateRangeEnd"
	//
	// Optional
	//
	// NOTE: FFI_USE_UpdateRangeEnd MUST be set to true to make this field effective.
	UpdateRangeEnd uint32
	// Text is "TextUpdateEventInit.text"
	//
	// Optional
	Text js.String
	// SelectionStart is "TextUpdateEventInit.selectionStart"
	//
	// Optional
	//
	// NOTE: FFI_USE_SelectionStart MUST be set to true to make this field effective.
	SelectionStart uint32
	// SelectionEnd is "TextUpdateEventInit.selectionEnd"
	//
	// Optional
	//
	// NOTE: FFI_USE_SelectionEnd MUST be set to true to make this field effective.
	SelectionEnd uint32
	// CompositionStart is "TextUpdateEventInit.compositionStart"
	//
	// Optional
	//
	// NOTE: FFI_USE_CompositionStart MUST be set to true to make this field effective.
	CompositionStart uint32
	// CompositionEnd is "TextUpdateEventInit.compositionEnd"
	//
	// Optional
	//
	// NOTE: FFI_USE_CompositionEnd MUST be set to true to make this field effective.
	CompositionEnd uint32
	// Bubbles is "TextUpdateEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "TextUpdateEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "TextUpdateEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_UpdateRangeStart bool // for UpdateRangeStart.
	FFI_USE_UpdateRangeEnd   bool // for UpdateRangeEnd.
	FFI_USE_SelectionStart   bool // for SelectionStart.
	FFI_USE_SelectionEnd     bool // for SelectionEnd.
	FFI_USE_CompositionStart bool // for CompositionStart.
	FFI_USE_CompositionEnd   bool // for CompositionEnd.
	FFI_USE_Bubbles          bool // for Bubbles.
	FFI_USE_Cancelable       bool // for Cancelable.
	FFI_USE_Composed         bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TextUpdateEventInit with all fields set.
func (p TextUpdateEventInit) FromRef(ref js.Ref) TextUpdateEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TextUpdateEventInit in the application heap.
func (p TextUpdateEventInit) New() js.Ref {
	return bindings.TextUpdateEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p TextUpdateEventInit) UpdateFrom(ref js.Ref) {
	bindings.TextUpdateEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p TextUpdateEventInit) Update(ref js.Ref) {
	bindings.TextUpdateEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewTextUpdateEvent(typ js.String, options TextUpdateEventInit) TextUpdateEvent {
	return TextUpdateEvent{}.FromRef(
		bindings.NewTextUpdateEventByTextUpdateEvent(
			typ.Ref(),
			js.Pointer(&options)),
	)
}

func NewTextUpdateEventByTextUpdateEvent1(typ js.String) TextUpdateEvent {
	return TextUpdateEvent{}.FromRef(
		bindings.NewTextUpdateEventByTextUpdateEvent1(
			typ.Ref()),
	)
}

type TextUpdateEvent struct {
	Event
}

func (this TextUpdateEvent) Once() TextUpdateEvent {
	this.Ref().Once()
	return this
}

func (this TextUpdateEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this TextUpdateEvent) FromRef(ref js.Ref) TextUpdateEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this TextUpdateEvent) Free() {
	this.Ref().Free()
}

// UpdateRangeStart returns the value of property "TextUpdateEvent.updateRangeStart".
//
// The returned bool will be false if there is no such property.
func (this TextUpdateEvent) UpdateRangeStart() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetTextUpdateEventUpdateRangeStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// UpdateRangeEnd returns the value of property "TextUpdateEvent.updateRangeEnd".
//
// The returned bool will be false if there is no such property.
func (this TextUpdateEvent) UpdateRangeEnd() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetTextUpdateEventUpdateRangeEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Text returns the value of property "TextUpdateEvent.text".
//
// The returned bool will be false if there is no such property.
func (this TextUpdateEvent) Text() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetTextUpdateEventText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// SelectionStart returns the value of property "TextUpdateEvent.selectionStart".
//
// The returned bool will be false if there is no such property.
func (this TextUpdateEvent) SelectionStart() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetTextUpdateEventSelectionStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// SelectionEnd returns the value of property "TextUpdateEvent.selectionEnd".
//
// The returned bool will be false if there is no such property.
func (this TextUpdateEvent) SelectionEnd() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetTextUpdateEventSelectionEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// CompositionStart returns the value of property "TextUpdateEvent.compositionStart".
//
// The returned bool will be false if there is no such property.
func (this TextUpdateEvent) CompositionStart() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetTextUpdateEventCompositionStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// CompositionEnd returns the value of property "TextUpdateEvent.compositionEnd".
//
// The returned bool will be false if there is no such property.
func (this TextUpdateEvent) CompositionEnd() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetTextUpdateEventCompositionEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

func NewTimeEvent(typ js.String, eventInitDict EventInit) TimeEvent {
	return TimeEvent{}.FromRef(
		bindings.NewTimeEventByTimeEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewTimeEventByTimeEvent1(typ js.String) TimeEvent {
	return TimeEvent{}.FromRef(
		bindings.NewTimeEventByTimeEvent1(
			typ.Ref()),
	)
}

type TimeEvent struct {
	Event
}

func (this TimeEvent) Once() TimeEvent {
	this.Ref().Once()
	return this
}

func (this TimeEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this TimeEvent) FromRef(ref js.Ref) TimeEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this TimeEvent) Free() {
	this.Ref().Free()
}

// View returns the value of property "TimeEvent.view".
//
// The returned bool will be false if there is no such property.
func (this TimeEvent) View() (js.Object, bool) {
	var _ok bool
	_ret := bindings.GetTimeEventView(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Object{}.FromRef(_ret), _ok
}

// Detail returns the value of property "TimeEvent.detail".
//
// The returned bool will be false if there is no such property.
func (this TimeEvent) Detail() (int32, bool) {
	var _ok bool
	_ret := bindings.GetTimeEventDetail(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// InitTimeEvent calls the method "TimeEvent.initTimeEvent".
//
// The returned bool will be false if there is no such method.
func (this TimeEvent) InitTimeEvent(typeArg js.String, viewArg Window, detailArg int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallTimeEventInitTimeEvent(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		viewArg.Ref(),
		int32(detailArg),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitTimeEventFunc returns the method "TimeEvent.initTimeEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TimeEvent) InitTimeEventFunc() (fn js.Func[func(typeArg js.String, viewArg Window, detailArg int32)]) {
	return fn.FromRef(
		bindings.TimeEventInitTimeEventFunc(
			this.Ref(),
		),
	)
}

type ToggleEventInit struct {
	// OldState is "ToggleEventInit.oldState"
	//
	// Optional, defaults to "".
	OldState js.String
	// NewState is "ToggleEventInit.newState"
	//
	// Optional, defaults to "".
	NewState js.String
	// Bubbles is "ToggleEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "ToggleEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "ToggleEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ToggleEventInit with all fields set.
func (p ToggleEventInit) FromRef(ref js.Ref) ToggleEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ToggleEventInit in the application heap.
func (p ToggleEventInit) New() js.Ref {
	return bindings.ToggleEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ToggleEventInit) UpdateFrom(ref js.Ref) {
	bindings.ToggleEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ToggleEventInit) Update(ref js.Ref) {
	bindings.ToggleEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewToggleEvent(typ js.String, eventInitDict ToggleEventInit) ToggleEvent {
	return ToggleEvent{}.FromRef(
		bindings.NewToggleEventByToggleEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewToggleEventByToggleEvent1(typ js.String) ToggleEvent {
	return ToggleEvent{}.FromRef(
		bindings.NewToggleEventByToggleEvent1(
			typ.Ref()),
	)
}

type ToggleEvent struct {
	Event
}

func (this ToggleEvent) Once() ToggleEvent {
	this.Ref().Once()
	return this
}

func (this ToggleEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this ToggleEvent) FromRef(ref js.Ref) ToggleEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this ToggleEvent) Free() {
	this.Ref().Free()
}

// OldState returns the value of property "ToggleEvent.oldState".
//
// The returned bool will be false if there is no such property.
func (this ToggleEvent) OldState() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetToggleEventOldState(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// NewState returns the value of property "ToggleEvent.newState".
//
// The returned bool will be false if there is no such property.
func (this ToggleEvent) NewState() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetToggleEventNewState(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type TokenBinding struct {
	// Status is "TokenBinding.status"
	//
	// Required
	Status js.String
	// Id is "TokenBinding.id"
	//
	// Optional
	Id js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TokenBinding with all fields set.
func (p TokenBinding) FromRef(ref js.Ref) TokenBinding {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TokenBinding in the application heap.
func (p TokenBinding) New() js.Ref {
	return bindings.TokenBindingJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p TokenBinding) UpdateFrom(ref js.Ref) {
	bindings.TokenBindingJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p TokenBinding) Update(ref js.Ref) {
	bindings.TokenBindingJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type TokenBindingStatus uint32

const (
	_ TokenBindingStatus = iota

	TokenBindingStatus_PRESENT
	TokenBindingStatus_SUPPORTED
)

func (TokenBindingStatus) FromRef(str js.Ref) TokenBindingStatus {
	return TokenBindingStatus(bindings.ConstOfTokenBindingStatus(str))
}

func (x TokenBindingStatus) String() (string, bool) {
	switch x {
	case TokenBindingStatus_PRESENT:
		return "present", true
	case TokenBindingStatus_SUPPORTED:
		return "supported", true
	default:
		return "", false
	}
}

type TopLevelStorageAccessPermissionDescriptor struct {
	// RequestedOrigin is "TopLevelStorageAccessPermissionDescriptor.requestedOrigin"
	//
	// Optional, defaults to "".
	RequestedOrigin js.String
	// Name is "TopLevelStorageAccessPermissionDescriptor.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TopLevelStorageAccessPermissionDescriptor with all fields set.
func (p TopLevelStorageAccessPermissionDescriptor) FromRef(ref js.Ref) TopLevelStorageAccessPermissionDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TopLevelStorageAccessPermissionDescriptor in the application heap.
func (p TopLevelStorageAccessPermissionDescriptor) New() js.Ref {
	return bindings.TopLevelStorageAccessPermissionDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p TopLevelStorageAccessPermissionDescriptor) UpdateFrom(ref js.Ref) {
	bindings.TopLevelStorageAccessPermissionDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p TopLevelStorageAccessPermissionDescriptor) Update(ref js.Ref) {
	bindings.TopLevelStorageAccessPermissionDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type TouchType uint32

const (
	_ TouchType = iota

	TouchType_DIRECT
	TouchType_STYLUS
)

func (TouchType) FromRef(str js.Ref) TouchType {
	return TouchType(bindings.ConstOfTouchType(str))
}

func (x TouchType) String() (string, bool) {
	switch x {
	case TouchType_DIRECT:
		return "direct", true
	case TouchType_STYLUS:
		return "stylus", true
	default:
		return "", false
	}
}

type TouchInit struct {
	// Identifier is "TouchInit.identifier"
	//
	// Required
	Identifier int32
	// Target is "TouchInit.target"
	//
	// Required
	Target EventTarget
	// ClientX is "TouchInit.clientX"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_ClientX MUST be set to true to make this field effective.
	ClientX float64
	// ClientY is "TouchInit.clientY"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_ClientY MUST be set to true to make this field effective.
	ClientY float64
	// ScreenX is "TouchInit.screenX"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_ScreenX MUST be set to true to make this field effective.
	ScreenX float64
	// ScreenY is "TouchInit.screenY"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_ScreenY MUST be set to true to make this field effective.
	ScreenY float64
	// PageX is "TouchInit.pageX"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_PageX MUST be set to true to make this field effective.
	PageX float64
	// PageY is "TouchInit.pageY"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_PageY MUST be set to true to make this field effective.
	PageY float64
	// RadiusX is "TouchInit.radiusX"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_RadiusX MUST be set to true to make this field effective.
	RadiusX float32
	// RadiusY is "TouchInit.radiusY"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_RadiusY MUST be set to true to make this field effective.
	RadiusY float32
	// RotationAngle is "TouchInit.rotationAngle"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_RotationAngle MUST be set to true to make this field effective.
	RotationAngle float32
	// Force is "TouchInit.force"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Force MUST be set to true to make this field effective.
	Force float32
	// AltitudeAngle is "TouchInit.altitudeAngle"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_AltitudeAngle MUST be set to true to make this field effective.
	AltitudeAngle float64
	// AzimuthAngle is "TouchInit.azimuthAngle"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_AzimuthAngle MUST be set to true to make this field effective.
	AzimuthAngle float64
	// TouchType is "TouchInit.touchType"
	//
	// Optional, defaults to "direct".
	TouchType TouchType

	FFI_USE_ClientX       bool // for ClientX.
	FFI_USE_ClientY       bool // for ClientY.
	FFI_USE_ScreenX       bool // for ScreenX.
	FFI_USE_ScreenY       bool // for ScreenY.
	FFI_USE_PageX         bool // for PageX.
	FFI_USE_PageY         bool // for PageY.
	FFI_USE_RadiusX       bool // for RadiusX.
	FFI_USE_RadiusY       bool // for RadiusY.
	FFI_USE_RotationAngle bool // for RotationAngle.
	FFI_USE_Force         bool // for Force.
	FFI_USE_AltitudeAngle bool // for AltitudeAngle.
	FFI_USE_AzimuthAngle  bool // for AzimuthAngle.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TouchInit with all fields set.
func (p TouchInit) FromRef(ref js.Ref) TouchInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TouchInit in the application heap.
func (p TouchInit) New() js.Ref {
	return bindings.TouchInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p TouchInit) UpdateFrom(ref js.Ref) {
	bindings.TouchInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p TouchInit) Update(ref js.Ref) {
	bindings.TouchInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewTouch(touchInitDict TouchInit) Touch {
	return Touch{}.FromRef(
		bindings.NewTouchByTouch(
			js.Pointer(&touchInitDict)),
	)
}

type Touch struct {
	ref js.Ref
}

func (this Touch) Once() Touch {
	this.Ref().Once()
	return this
}

func (this Touch) Ref() js.Ref {
	return this.ref
}

func (this Touch) FromRef(ref js.Ref) Touch {
	this.ref = ref
	return this
}

func (this Touch) Free() {
	this.Ref().Free()
}

// Identifier returns the value of property "Touch.identifier".
//
// The returned bool will be false if there is no such property.
func (this Touch) Identifier() (int32, bool) {
	var _ok bool
	_ret := bindings.GetTouchIdentifier(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Target returns the value of property "Touch.target".
//
// The returned bool will be false if there is no such property.
func (this Touch) Target() (EventTarget, bool) {
	var _ok bool
	_ret := bindings.GetTouchTarget(
		this.Ref(), js.Pointer(&_ok),
	)
	return EventTarget{}.FromRef(_ret), _ok
}

// ScreenX returns the value of property "Touch.screenX".
//
// The returned bool will be false if there is no such property.
func (this Touch) ScreenX() (float64, bool) {
	var _ok bool
	_ret := bindings.GetTouchScreenX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ScreenY returns the value of property "Touch.screenY".
//
// The returned bool will be false if there is no such property.
func (this Touch) ScreenY() (float64, bool) {
	var _ok bool
	_ret := bindings.GetTouchScreenY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ClientX returns the value of property "Touch.clientX".
//
// The returned bool will be false if there is no such property.
func (this Touch) ClientX() (float64, bool) {
	var _ok bool
	_ret := bindings.GetTouchClientX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ClientY returns the value of property "Touch.clientY".
//
// The returned bool will be false if there is no such property.
func (this Touch) ClientY() (float64, bool) {
	var _ok bool
	_ret := bindings.GetTouchClientY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// PageX returns the value of property "Touch.pageX".
//
// The returned bool will be false if there is no such property.
func (this Touch) PageX() (float64, bool) {
	var _ok bool
	_ret := bindings.GetTouchPageX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// PageY returns the value of property "Touch.pageY".
//
// The returned bool will be false if there is no such property.
func (this Touch) PageY() (float64, bool) {
	var _ok bool
	_ret := bindings.GetTouchPageY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// RadiusX returns the value of property "Touch.radiusX".
//
// The returned bool will be false if there is no such property.
func (this Touch) RadiusX() (float32, bool) {
	var _ok bool
	_ret := bindings.GetTouchRadiusX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// RadiusY returns the value of property "Touch.radiusY".
//
// The returned bool will be false if there is no such property.
func (this Touch) RadiusY() (float32, bool) {
	var _ok bool
	_ret := bindings.GetTouchRadiusY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// RotationAngle returns the value of property "Touch.rotationAngle".
//
// The returned bool will be false if there is no such property.
func (this Touch) RotationAngle() (float32, bool) {
	var _ok bool
	_ret := bindings.GetTouchRotationAngle(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// Force returns the value of property "Touch.force".
//
// The returned bool will be false if there is no such property.
func (this Touch) Force() (float32, bool) {
	var _ok bool
	_ret := bindings.GetTouchForce(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// AltitudeAngle returns the value of property "Touch.altitudeAngle".
//
// The returned bool will be false if there is no such property.
func (this Touch) AltitudeAngle() (float32, bool) {
	var _ok bool
	_ret := bindings.GetTouchAltitudeAngle(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// AzimuthAngle returns the value of property "Touch.azimuthAngle".
//
// The returned bool will be false if there is no such property.
func (this Touch) AzimuthAngle() (float32, bool) {
	var _ok bool
	_ret := bindings.GetTouchAzimuthAngle(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// TouchType returns the value of property "Touch.touchType".
//
// The returned bool will be false if there is no such property.
func (this Touch) TouchType() (TouchType, bool) {
	var _ok bool
	_ret := bindings.GetTouchTouchType(
		this.Ref(), js.Pointer(&_ok),
	)
	return TouchType(_ret), _ok
}

type TouchEventInit struct {
	// Touches is "TouchEventInit.touches"
	//
	// Optional, defaults to [].
	Touches js.Array[Touch]
	// TargetTouches is "TouchEventInit.targetTouches"
	//
	// Optional, defaults to [].
	TargetTouches js.Array[Touch]
	// ChangedTouches is "TouchEventInit.changedTouches"
	//
	// Optional, defaults to [].
	ChangedTouches js.Array[Touch]
	// CtrlKey is "TouchEventInit.ctrlKey"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_CtrlKey MUST be set to true to make this field effective.
	CtrlKey bool
	// ShiftKey is "TouchEventInit.shiftKey"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ShiftKey MUST be set to true to make this field effective.
	ShiftKey bool
	// AltKey is "TouchEventInit.altKey"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_AltKey MUST be set to true to make this field effective.
	AltKey bool
	// MetaKey is "TouchEventInit.metaKey"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_MetaKey MUST be set to true to make this field effective.
	MetaKey bool
	// ModifierAltGraph is "TouchEventInit.modifierAltGraph"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierAltGraph MUST be set to true to make this field effective.
	ModifierAltGraph bool
	// ModifierCapsLock is "TouchEventInit.modifierCapsLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierCapsLock MUST be set to true to make this field effective.
	ModifierCapsLock bool
	// ModifierFn is "TouchEventInit.modifierFn"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierFn MUST be set to true to make this field effective.
	ModifierFn bool
	// ModifierFnLock is "TouchEventInit.modifierFnLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierFnLock MUST be set to true to make this field effective.
	ModifierFnLock bool
	// ModifierHyper is "TouchEventInit.modifierHyper"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierHyper MUST be set to true to make this field effective.
	ModifierHyper bool
	// ModifierNumLock is "TouchEventInit.modifierNumLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierNumLock MUST be set to true to make this field effective.
	ModifierNumLock bool
	// ModifierScrollLock is "TouchEventInit.modifierScrollLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierScrollLock MUST be set to true to make this field effective.
	ModifierScrollLock bool
	// ModifierSuper is "TouchEventInit.modifierSuper"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierSuper MUST be set to true to make this field effective.
	ModifierSuper bool
	// ModifierSymbol is "TouchEventInit.modifierSymbol"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierSymbol MUST be set to true to make this field effective.
	ModifierSymbol bool
	// ModifierSymbolLock is "TouchEventInit.modifierSymbolLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierSymbolLock MUST be set to true to make this field effective.
	ModifierSymbolLock bool
	// View is "TouchEventInit.view"
	//
	// Optional, defaults to null.
	View Window
	// Detail is "TouchEventInit.detail"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Detail MUST be set to true to make this field effective.
	Detail int32
	// Bubbles is "TouchEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "TouchEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "TouchEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_CtrlKey            bool // for CtrlKey.
	FFI_USE_ShiftKey           bool // for ShiftKey.
	FFI_USE_AltKey             bool // for AltKey.
	FFI_USE_MetaKey            bool // for MetaKey.
	FFI_USE_ModifierAltGraph   bool // for ModifierAltGraph.
	FFI_USE_ModifierCapsLock   bool // for ModifierCapsLock.
	FFI_USE_ModifierFn         bool // for ModifierFn.
	FFI_USE_ModifierFnLock     bool // for ModifierFnLock.
	FFI_USE_ModifierHyper      bool // for ModifierHyper.
	FFI_USE_ModifierNumLock    bool // for ModifierNumLock.
	FFI_USE_ModifierScrollLock bool // for ModifierScrollLock.
	FFI_USE_ModifierSuper      bool // for ModifierSuper.
	FFI_USE_ModifierSymbol     bool // for ModifierSymbol.
	FFI_USE_ModifierSymbolLock bool // for ModifierSymbolLock.
	FFI_USE_Detail             bool // for Detail.
	FFI_USE_Bubbles            bool // for Bubbles.
	FFI_USE_Cancelable         bool // for Cancelable.
	FFI_USE_Composed           bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TouchEventInit with all fields set.
func (p TouchEventInit) FromRef(ref js.Ref) TouchEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TouchEventInit in the application heap.
func (p TouchEventInit) New() js.Ref {
	return bindings.TouchEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p TouchEventInit) UpdateFrom(ref js.Ref) {
	bindings.TouchEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p TouchEventInit) Update(ref js.Ref) {
	bindings.TouchEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type TouchList struct {
	ref js.Ref
}

func (this TouchList) Once() TouchList {
	this.Ref().Once()
	return this
}

func (this TouchList) Ref() js.Ref {
	return this.ref
}

func (this TouchList) FromRef(ref js.Ref) TouchList {
	this.ref = ref
	return this
}

func (this TouchList) Free() {
	this.Ref().Free()
}

// Length returns the value of property "TouchList.length".
//
// The returned bool will be false if there is no such property.
func (this TouchList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetTouchListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Item calls the method "TouchList.item".
//
// The returned bool will be false if there is no such method.
func (this TouchList) Item(index uint32) (Touch, bool) {
	var _ok bool
	_ret := bindings.CallTouchListItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return Touch{}.FromRef(_ret), _ok
}

// ItemFunc returns the method "TouchList.item".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TouchList) ItemFunc() (fn js.Func[func(index uint32) Touch]) {
	return fn.FromRef(
		bindings.TouchListItemFunc(
			this.Ref(),
		),
	)
}

func NewTouchEvent(typ js.String, eventInitDict TouchEventInit) TouchEvent {
	return TouchEvent{}.FromRef(
		bindings.NewTouchEventByTouchEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewTouchEventByTouchEvent1(typ js.String) TouchEvent {
	return TouchEvent{}.FromRef(
		bindings.NewTouchEventByTouchEvent1(
			typ.Ref()),
	)
}

type TouchEvent struct {
	UIEvent
}

func (this TouchEvent) Once() TouchEvent {
	this.Ref().Once()
	return this
}

func (this TouchEvent) Ref() js.Ref {
	return this.UIEvent.Ref()
}

func (this TouchEvent) FromRef(ref js.Ref) TouchEvent {
	this.UIEvent = this.UIEvent.FromRef(ref)
	return this
}

func (this TouchEvent) Free() {
	this.Ref().Free()
}

// Touches returns the value of property "TouchEvent.touches".
//
// The returned bool will be false if there is no such property.
func (this TouchEvent) Touches() (TouchList, bool) {
	var _ok bool
	_ret := bindings.GetTouchEventTouches(
		this.Ref(), js.Pointer(&_ok),
	)
	return TouchList{}.FromRef(_ret), _ok
}

// TargetTouches returns the value of property "TouchEvent.targetTouches".
//
// The returned bool will be false if there is no such property.
func (this TouchEvent) TargetTouches() (TouchList, bool) {
	var _ok bool
	_ret := bindings.GetTouchEventTargetTouches(
		this.Ref(), js.Pointer(&_ok),
	)
	return TouchList{}.FromRef(_ret), _ok
}

// ChangedTouches returns the value of property "TouchEvent.changedTouches".
//
// The returned bool will be false if there is no such property.
func (this TouchEvent) ChangedTouches() (TouchList, bool) {
	var _ok bool
	_ret := bindings.GetTouchEventChangedTouches(
		this.Ref(), js.Pointer(&_ok),
	)
	return TouchList{}.FromRef(_ret), _ok
}

// AltKey returns the value of property "TouchEvent.altKey".
//
// The returned bool will be false if there is no such property.
func (this TouchEvent) AltKey() (bool, bool) {
	var _ok bool
	_ret := bindings.GetTouchEventAltKey(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// MetaKey returns the value of property "TouchEvent.metaKey".
//
// The returned bool will be false if there is no such property.
func (this TouchEvent) MetaKey() (bool, bool) {
	var _ok bool
	_ret := bindings.GetTouchEventMetaKey(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// CtrlKey returns the value of property "TouchEvent.ctrlKey".
//
// The returned bool will be false if there is no such property.
func (this TouchEvent) CtrlKey() (bool, bool) {
	var _ok bool
	_ret := bindings.GetTouchEventCtrlKey(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// ShiftKey returns the value of property "TouchEvent.shiftKey".
//
// The returned bool will be false if there is no such property.
func (this TouchEvent) ShiftKey() (bool, bool) {
	var _ok bool
	_ret := bindings.GetTouchEventShiftKey(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// GetModifierState calls the method "TouchEvent.getModifierState".
//
// The returned bool will be false if there is no such method.
func (this TouchEvent) GetModifierState(keyArg js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallTouchEventGetModifierState(
		this.Ref(), js.Pointer(&_ok),
		keyArg.Ref(),
	)

	return _ret == js.True, _ok
}

// GetModifierStateFunc returns the method "TouchEvent.getModifierState".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TouchEvent) GetModifierStateFunc() (fn js.Func[func(keyArg js.String) bool]) {
	return fn.FromRef(
		bindings.TouchEventGetModifierStateFunc(
			this.Ref(),
		),
	)
}

type OneOf_VideoTrack_AudioTrack_TextTrack struct {
	ref js.Ref
}

func (x OneOf_VideoTrack_AudioTrack_TextTrack) Ref() js.Ref {
	return x.ref
}

func (x OneOf_VideoTrack_AudioTrack_TextTrack) Free() {
	x.ref.Free()
}

func (x OneOf_VideoTrack_AudioTrack_TextTrack) FromRef(ref js.Ref) OneOf_VideoTrack_AudioTrack_TextTrack {
	return OneOf_VideoTrack_AudioTrack_TextTrack{
		ref: ref,
	}
}

func (x OneOf_VideoTrack_AudioTrack_TextTrack) VideoTrack() VideoTrack {
	return VideoTrack{}.FromRef(x.ref)
}

func (x OneOf_VideoTrack_AudioTrack_TextTrack) AudioTrack() AudioTrack {
	return AudioTrack{}.FromRef(x.ref)
}

func (x OneOf_VideoTrack_AudioTrack_TextTrack) TextTrack() TextTrack {
	return TextTrack{}.FromRef(x.ref)
}

type TrackEventInit struct {
	// Track is "TrackEventInit.track"
	//
	// Optional, defaults to null.
	Track OneOf_VideoTrack_AudioTrack_TextTrack
	// Bubbles is "TrackEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "TrackEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "TrackEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TrackEventInit with all fields set.
func (p TrackEventInit) FromRef(ref js.Ref) TrackEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TrackEventInit in the application heap.
func (p TrackEventInit) New() js.Ref {
	return bindings.TrackEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p TrackEventInit) UpdateFrom(ref js.Ref) {
	bindings.TrackEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p TrackEventInit) Update(ref js.Ref) {
	bindings.TrackEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewTrackEvent(typ js.String, eventInitDict TrackEventInit) TrackEvent {
	return TrackEvent{}.FromRef(
		bindings.NewTrackEventByTrackEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewTrackEventByTrackEvent1(typ js.String) TrackEvent {
	return TrackEvent{}.FromRef(
		bindings.NewTrackEventByTrackEvent1(
			typ.Ref()),
	)
}

type TrackEvent struct {
	Event
}

func (this TrackEvent) Once() TrackEvent {
	this.Ref().Once()
	return this
}

func (this TrackEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this TrackEvent) FromRef(ref js.Ref) TrackEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this TrackEvent) Free() {
	this.Ref().Free()
}

// Track returns the value of property "TrackEvent.track".
//
// The returned bool will be false if there is no such property.
func (this TrackEvent) Track() (OneOf_VideoTrack_AudioTrack_TextTrack, bool) {
	var _ok bool
	_ret := bindings.GetTrackEventTrack(
		this.Ref(), js.Pointer(&_ok),
	)
	return OneOf_VideoTrack_AudioTrack_TextTrack{}.FromRef(_ret), _ok
}

func NewTransformStream(transformer js.Object, writableStrategy QueuingStrategy, readableStrategy QueuingStrategy) TransformStream {
	return TransformStream{}.FromRef(
		bindings.NewTransformStreamByTransformStream(
			transformer.Ref(),
			js.Pointer(&writableStrategy),
			js.Pointer(&readableStrategy)),
	)
}

func NewTransformStreamByTransformStream1(transformer js.Object, writableStrategy QueuingStrategy) TransformStream {
	return TransformStream{}.FromRef(
		bindings.NewTransformStreamByTransformStream1(
			transformer.Ref(),
			js.Pointer(&writableStrategy)),
	)
}

func NewTransformStreamByTransformStream2(transformer js.Object) TransformStream {
	return TransformStream{}.FromRef(
		bindings.NewTransformStreamByTransformStream2(
			transformer.Ref()),
	)
}

func NewTransformStreamByTransformStream3() TransformStream {
	return TransformStream{}.FromRef(
		bindings.NewTransformStreamByTransformStream3(),
	)
}

type TransformStream struct {
	ref js.Ref
}

func (this TransformStream) Once() TransformStream {
	this.Ref().Once()
	return this
}

func (this TransformStream) Ref() js.Ref {
	return this.ref
}

func (this TransformStream) FromRef(ref js.Ref) TransformStream {
	this.ref = ref
	return this
}

func (this TransformStream) Free() {
	this.Ref().Free()
}

// Readable returns the value of property "TransformStream.readable".
//
// The returned bool will be false if there is no such property.
func (this TransformStream) Readable() (ReadableStream, bool) {
	var _ok bool
	_ret := bindings.GetTransformStreamReadable(
		this.Ref(), js.Pointer(&_ok),
	)
	return ReadableStream{}.FromRef(_ret), _ok
}

// Writable returns the value of property "TransformStream.writable".
//
// The returned bool will be false if there is no such property.
func (this TransformStream) Writable() (WritableStream, bool) {
	var _ok bool
	_ret := bindings.GetTransformStreamWritable(
		this.Ref(), js.Pointer(&_ok),
	)
	return WritableStream{}.FromRef(_ret), _ok
}

type TransformStreamDefaultController struct {
	ref js.Ref
}

func (this TransformStreamDefaultController) Once() TransformStreamDefaultController {
	this.Ref().Once()
	return this
}

func (this TransformStreamDefaultController) Ref() js.Ref {
	return this.ref
}

func (this TransformStreamDefaultController) FromRef(ref js.Ref) TransformStreamDefaultController {
	this.ref = ref
	return this
}

func (this TransformStreamDefaultController) Free() {
	this.Ref().Free()
}

// DesiredSize returns the value of property "TransformStreamDefaultController.desiredSize".
//
// The returned bool will be false if there is no such property.
func (this TransformStreamDefaultController) DesiredSize() (float64, bool) {
	var _ok bool
	_ret := bindings.GetTransformStreamDefaultControllerDesiredSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Enqueue calls the method "TransformStreamDefaultController.enqueue".
//
// The returned bool will be false if there is no such method.
func (this TransformStreamDefaultController) Enqueue(chunk js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallTransformStreamDefaultControllerEnqueue(
		this.Ref(), js.Pointer(&_ok),
		chunk.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// EnqueueFunc returns the method "TransformStreamDefaultController.enqueue".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TransformStreamDefaultController) EnqueueFunc() (fn js.Func[func(chunk js.Any)]) {
	return fn.FromRef(
		bindings.TransformStreamDefaultControllerEnqueueFunc(
			this.Ref(),
		),
	)
}

// Enqueue1 calls the method "TransformStreamDefaultController.enqueue".
//
// The returned bool will be false if there is no such method.
func (this TransformStreamDefaultController) Enqueue1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallTransformStreamDefaultControllerEnqueue1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Enqueue1Func returns the method "TransformStreamDefaultController.enqueue".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TransformStreamDefaultController) Enqueue1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.TransformStreamDefaultControllerEnqueue1Func(
			this.Ref(),
		),
	)
}

// Error calls the method "TransformStreamDefaultController.error".
//
// The returned bool will be false if there is no such method.
func (this TransformStreamDefaultController) Error(reason js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallTransformStreamDefaultControllerError(
		this.Ref(), js.Pointer(&_ok),
		reason.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ErrorFunc returns the method "TransformStreamDefaultController.error".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TransformStreamDefaultController) ErrorFunc() (fn js.Func[func(reason js.Any)]) {
	return fn.FromRef(
		bindings.TransformStreamDefaultControllerErrorFunc(
			this.Ref(),
		),
	)
}

// Error1 calls the method "TransformStreamDefaultController.error".
//
// The returned bool will be false if there is no such method.
func (this TransformStreamDefaultController) Error1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallTransformStreamDefaultControllerError1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Error1Func returns the method "TransformStreamDefaultController.error".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TransformStreamDefaultController) Error1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.TransformStreamDefaultControllerError1Func(
			this.Ref(),
		),
	)
}

// Terminate calls the method "TransformStreamDefaultController.terminate".
//
// The returned bool will be false if there is no such method.
func (this TransformStreamDefaultController) Terminate() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallTransformStreamDefaultControllerTerminate(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TerminateFunc returns the method "TransformStreamDefaultController.terminate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TransformStreamDefaultController) TerminateFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.TransformStreamDefaultControllerTerminateFunc(
			this.Ref(),
		),
	)
}

type TransformerStartCallbackFunc func(this js.Ref, controller TransformStreamDefaultController) js.Ref

func (fn TransformerStartCallbackFunc) Register() js.Func[func(controller TransformStreamDefaultController) js.Any] {
	return js.RegisterCallback[func(controller TransformStreamDefaultController) js.Any](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn TransformerStartCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		TransformStreamDefaultController{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type TransformerStartCallback[T any] struct {
	Fn  func(arg T, this js.Ref, controller TransformStreamDefaultController) js.Ref
	Arg T
}

func (cb *TransformerStartCallback[T]) Register() js.Func[func(controller TransformStreamDefaultController) js.Any] {
	return js.RegisterCallback[func(controller TransformStreamDefaultController) js.Any](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *TransformerStartCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		TransformStreamDefaultController{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type TransformerTransformCallbackFunc func(this js.Ref, chunk js.Any, controller TransformStreamDefaultController) js.Ref

func (fn TransformerTransformCallbackFunc) Register() js.Func[func(chunk js.Any, controller TransformStreamDefaultController) js.Promise[js.Void]] {
	return js.RegisterCallback[func(chunk js.Any, controller TransformStreamDefaultController) js.Promise[js.Void]](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn TransformerTransformCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Any{}.FromRef(args[0+1]),
		TransformStreamDefaultController{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type TransformerTransformCallback[T any] struct {
	Fn  func(arg T, this js.Ref, chunk js.Any, controller TransformStreamDefaultController) js.Ref
	Arg T
}

func (cb *TransformerTransformCallback[T]) Register() js.Func[func(chunk js.Any, controller TransformStreamDefaultController) js.Promise[js.Void]] {
	return js.RegisterCallback[func(chunk js.Any, controller TransformStreamDefaultController) js.Promise[js.Void]](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *TransformerTransformCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Any{}.FromRef(args[0+1]),
		TransformStreamDefaultController{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type TransformerFlushCallbackFunc func(this js.Ref, controller TransformStreamDefaultController) js.Ref

func (fn TransformerFlushCallbackFunc) Register() js.Func[func(controller TransformStreamDefaultController) js.Promise[js.Void]] {
	return js.RegisterCallback[func(controller TransformStreamDefaultController) js.Promise[js.Void]](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn TransformerFlushCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		TransformStreamDefaultController{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type TransformerFlushCallback[T any] struct {
	Fn  func(arg T, this js.Ref, controller TransformStreamDefaultController) js.Ref
	Arg T
}

func (cb *TransformerFlushCallback[T]) Register() js.Func[func(controller TransformStreamDefaultController) js.Promise[js.Void]] {
	return js.RegisterCallback[func(controller TransformStreamDefaultController) js.Promise[js.Void]](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *TransformerFlushCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		TransformStreamDefaultController{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}
