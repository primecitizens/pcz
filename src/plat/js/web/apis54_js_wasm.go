// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
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

func NewSyncEvent(typ js.String, init SyncEventInit) (ret SyncEvent) {
	ret.ref = bindings.NewSyncEventBySyncEvent(
		typ.Ref(),
		js.Pointer(&init))
	return
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
// It returns ok=false if there is no such property.
func (this SyncEvent) Tag() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSyncEventTag(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LastChance returns the value of property "SyncEvent.lastChance".
//
// It returns ok=false if there is no such property.
func (this SyncEvent) LastChance() (ret bool, ok bool) {
	ok = js.True == bindings.GetSyncEventLastChance(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

func NewTable(descriptor TableDescriptor, value js.Any) (ret Table) {
	ret.ref = bindings.NewTableByTable(
		js.Pointer(&descriptor),
		value.Ref())
	return
}

func NewTableByTable1(descriptor TableDescriptor) (ret Table) {
	ret.ref = bindings.NewTableByTable1(
		js.Pointer(&descriptor))
	return
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
// It returns ok=false if there is no such property.
func (this Table) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetTableLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGrow returns true if the method "Table.grow" exists.
func (this Table) HasGrow() bool {
	return js.True == bindings.HasTableGrow(
		this.Ref(),
	)
}

// GrowFunc returns the method "Table.grow".
func (this Table) GrowFunc() (fn js.Func[func(delta uint32, value js.Any) uint32]) {
	return fn.FromRef(
		bindings.TableGrowFunc(
			this.Ref(),
		),
	)
}

// Grow calls the method "Table.grow".
func (this Table) Grow(delta uint32, value js.Any) (ret uint32) {
	bindings.CallTableGrow(
		this.Ref(), js.Pointer(&ret),
		uint32(delta),
		value.Ref(),
	)

	return
}

// TryGrow calls the method "Table.grow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Table) TryGrow(delta uint32, value js.Any) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTableGrow(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(delta),
		value.Ref(),
	)

	return
}

// HasGrow1 returns true if the method "Table.grow" exists.
func (this Table) HasGrow1() bool {
	return js.True == bindings.HasTableGrow1(
		this.Ref(),
	)
}

// Grow1Func returns the method "Table.grow".
func (this Table) Grow1Func() (fn js.Func[func(delta uint32) uint32]) {
	return fn.FromRef(
		bindings.TableGrow1Func(
			this.Ref(),
		),
	)
}

// Grow1 calls the method "Table.grow".
func (this Table) Grow1(delta uint32) (ret uint32) {
	bindings.CallTableGrow1(
		this.Ref(), js.Pointer(&ret),
		uint32(delta),
	)

	return
}

// TryGrow1 calls the method "Table.grow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Table) TryGrow1(delta uint32) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTableGrow1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(delta),
	)

	return
}

// HasGet returns true if the method "Table.get" exists.
func (this Table) HasGet() bool {
	return js.True == bindings.HasTableGet(
		this.Ref(),
	)
}

// GetFunc returns the method "Table.get".
func (this Table) GetFunc() (fn js.Func[func(index uint32) js.Any]) {
	return fn.FromRef(
		bindings.TableGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "Table.get".
func (this Table) Get(index uint32) (ret js.Any) {
	bindings.CallTableGet(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "Table.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Table) TryGet(index uint32) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTableGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasSet returns true if the method "Table.set" exists.
func (this Table) HasSet() bool {
	return js.True == bindings.HasTableSet(
		this.Ref(),
	)
}

// SetFunc returns the method "Table.set".
func (this Table) SetFunc() (fn js.Func[func(index uint32, value js.Any)]) {
	return fn.FromRef(
		bindings.TableSetFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "Table.set".
func (this Table) Set(index uint32, value js.Any) (ret js.Void) {
	bindings.CallTableSet(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		value.Ref(),
	)

	return
}

// TrySet calls the method "Table.set"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Table) TrySet(index uint32, value js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTableSet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		value.Ref(),
	)

	return
}

// HasSet1 returns true if the method "Table.set" exists.
func (this Table) HasSet1() bool {
	return js.True == bindings.HasTableSet1(
		this.Ref(),
	)
}

// Set1Func returns the method "Table.set".
func (this Table) Set1Func() (fn js.Func[func(index uint32)]) {
	return fn.FromRef(
		bindings.TableSet1Func(
			this.Ref(),
		),
	)
}

// Set1 calls the method "Table.set".
func (this Table) Set1(index uint32) (ret js.Void) {
	bindings.CallTableSet1(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TrySet1 calls the method "Table.set"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Table) TrySet1(index uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTableSet1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
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

func NewTaskController(init TaskControllerInit) (ret TaskController) {
	ret.ref = bindings.NewTaskControllerByTaskController(
		js.Pointer(&init))
	return
}

func NewTaskControllerByTaskController1() (ret TaskController) {
	ret.ref = bindings.NewTaskControllerByTaskController1()
	return
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

// HasSetPriority returns true if the method "TaskController.setPriority" exists.
func (this TaskController) HasSetPriority() bool {
	return js.True == bindings.HasTaskControllerSetPriority(
		this.Ref(),
	)
}

// SetPriorityFunc returns the method "TaskController.setPriority".
func (this TaskController) SetPriorityFunc() (fn js.Func[func(priority TaskPriority)]) {
	return fn.FromRef(
		bindings.TaskControllerSetPriorityFunc(
			this.Ref(),
		),
	)
}

// SetPriority calls the method "TaskController.setPriority".
func (this TaskController) SetPriority(priority TaskPriority) (ret js.Void) {
	bindings.CallTaskControllerSetPriority(
		this.Ref(), js.Pointer(&ret),
		uint32(priority),
	)

	return
}

// TrySetPriority calls the method "TaskController.setPriority"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TaskController) TrySetPriority(priority TaskPriority) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTaskControllerSetPriority(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(priority),
	)

	return
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

func NewTaskPriorityChangeEvent(typ js.String, priorityChangeEventInitDict TaskPriorityChangeEventInit) (ret TaskPriorityChangeEvent) {
	ret.ref = bindings.NewTaskPriorityChangeEventByTaskPriorityChangeEvent(
		typ.Ref(),
		js.Pointer(&priorityChangeEventInitDict))
	return
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
// It returns ok=false if there is no such property.
func (this TaskPriorityChangeEvent) PreviousPriority() (ret TaskPriority, ok bool) {
	ok = js.True == bindings.GetTaskPriorityChangeEventPreviousPriority(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this TaskSignal) Priority() (ret TaskPriority, ok bool) {
	ok = js.True == bindings.GetTaskSignalPriority(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasAny returns true if the staticmethod "TaskSignal.any" exists.
func (this TaskSignal) HasAny() bool {
	return js.True == bindings.HasTaskSignalAny(
		this.Ref(),
	)
}

// AnyFunc returns the staticmethod "TaskSignal.any".
func (this TaskSignal) AnyFunc() (fn js.Func[func(signals js.Array[AbortSignal], init TaskSignalAnyInit) TaskSignal]) {
	return fn.FromRef(
		bindings.TaskSignalAnyFunc(
			this.Ref(),
		),
	)
}

// Any calls the staticmethod "TaskSignal.any".
func (this TaskSignal) Any(signals js.Array[AbortSignal], init TaskSignalAnyInit) (ret TaskSignal) {
	bindings.CallTaskSignalAny(
		this.Ref(), js.Pointer(&ret),
		signals.Ref(),
		js.Pointer(&init),
	)

	return
}

// TryAny calls the staticmethod "TaskSignal.any"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TaskSignal) TryAny(signals js.Array[AbortSignal], init TaskSignalAnyInit) (ret TaskSignal, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTaskSignalAny(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		signals.Ref(),
		js.Pointer(&init),
	)

	return
}

// HasAny1 returns true if the staticmethod "TaskSignal.any" exists.
func (this TaskSignal) HasAny1() bool {
	return js.True == bindings.HasTaskSignalAny1(
		this.Ref(),
	)
}

// Any1Func returns the staticmethod "TaskSignal.any".
func (this TaskSignal) Any1Func() (fn js.Func[func(signals js.Array[AbortSignal]) TaskSignal]) {
	return fn.FromRef(
		bindings.TaskSignalAny1Func(
			this.Ref(),
		),
	)
}

// Any1 calls the staticmethod "TaskSignal.any".
func (this TaskSignal) Any1(signals js.Array[AbortSignal]) (ret TaskSignal) {
	bindings.CallTaskSignalAny1(
		this.Ref(), js.Pointer(&ret),
		signals.Ref(),
	)

	return
}

// TryAny1 calls the staticmethod "TaskSignal.any"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TaskSignal) TryAny1(signals js.Array[AbortSignal]) (ret TaskSignal, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTaskSignalAny1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		signals.Ref(),
	)

	return
}

type TestUtils struct{}

// HasGc returns ture if the function "TestUtils.gc" exists.
func (TestUtils) HasGc() bool {
	return js.True == bindings.HasTestUtilsGc()
}

// GcFunc returns the function "TestUtils.gc".
func (TestUtils) GcFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.TestUtilsGcFunc(),
	)
}

// Gc calls the function "TestUtils.gc".
func (TestUtils) Gc() (ret js.Promise[js.Void]) {
	bindings.CallTestUtilsGc(
		js.Pointer(&ret),
	)
	return
}

// TryGc calls the function "TestUtils.gc"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (TestUtils) TryGc() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryTestUtilsGc(
		js.Pointer(&ret), js.Pointer(&exception),
	)
	return
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

func NewTextDecoder(label js.String, options TextDecoderOptions) (ret TextDecoder) {
	ret.ref = bindings.NewTextDecoderByTextDecoder(
		label.Ref(),
		js.Pointer(&options))
	return
}

func NewTextDecoderByTextDecoder1(label js.String) (ret TextDecoder) {
	ret.ref = bindings.NewTextDecoderByTextDecoder1(
		label.Ref())
	return
}

func NewTextDecoderByTextDecoder2() (ret TextDecoder) {
	ret.ref = bindings.NewTextDecoderByTextDecoder2()
	return
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
// It returns ok=false if there is no such property.
func (this TextDecoder) Encoding() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTextDecoderEncoding(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Fatal returns the value of property "TextDecoder.fatal".
//
// It returns ok=false if there is no such property.
func (this TextDecoder) Fatal() (ret bool, ok bool) {
	ok = js.True == bindings.GetTextDecoderFatal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IgnoreBOM returns the value of property "TextDecoder.ignoreBOM".
//
// It returns ok=false if there is no such property.
func (this TextDecoder) IgnoreBOM() (ret bool, ok bool) {
	ok = js.True == bindings.GetTextDecoderIgnoreBOM(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasDecode returns true if the method "TextDecoder.decode" exists.
func (this TextDecoder) HasDecode() bool {
	return js.True == bindings.HasTextDecoderDecode(
		this.Ref(),
	)
}

// DecodeFunc returns the method "TextDecoder.decode".
func (this TextDecoder) DecodeFunc() (fn js.Func[func(input AllowSharedBufferSource, options TextDecodeOptions) js.String]) {
	return fn.FromRef(
		bindings.TextDecoderDecodeFunc(
			this.Ref(),
		),
	)
}

// Decode calls the method "TextDecoder.decode".
func (this TextDecoder) Decode(input AllowSharedBufferSource, options TextDecodeOptions) (ret js.String) {
	bindings.CallTextDecoderDecode(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryDecode calls the method "TextDecoder.decode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TextDecoder) TryDecode(input AllowSharedBufferSource, options TextDecodeOptions) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextDecoderDecode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasDecode1 returns true if the method "TextDecoder.decode" exists.
func (this TextDecoder) HasDecode1() bool {
	return js.True == bindings.HasTextDecoderDecode1(
		this.Ref(),
	)
}

// Decode1Func returns the method "TextDecoder.decode".
func (this TextDecoder) Decode1Func() (fn js.Func[func(input AllowSharedBufferSource) js.String]) {
	return fn.FromRef(
		bindings.TextDecoderDecode1Func(
			this.Ref(),
		),
	)
}

// Decode1 calls the method "TextDecoder.decode".
func (this TextDecoder) Decode1(input AllowSharedBufferSource) (ret js.String) {
	bindings.CallTextDecoderDecode1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryDecode1 calls the method "TextDecoder.decode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TextDecoder) TryDecode1(input AllowSharedBufferSource) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextDecoderDecode1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasDecode2 returns true if the method "TextDecoder.decode" exists.
func (this TextDecoder) HasDecode2() bool {
	return js.True == bindings.HasTextDecoderDecode2(
		this.Ref(),
	)
}

// Decode2Func returns the method "TextDecoder.decode".
func (this TextDecoder) Decode2Func() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.TextDecoderDecode2Func(
			this.Ref(),
		),
	)
}

// Decode2 calls the method "TextDecoder.decode".
func (this TextDecoder) Decode2() (ret js.String) {
	bindings.CallTextDecoderDecode2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDecode2 calls the method "TextDecoder.decode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TextDecoder) TryDecode2() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextDecoderDecode2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewTextDecoderStream(label js.String, options TextDecoderOptions) (ret TextDecoderStream) {
	ret.ref = bindings.NewTextDecoderStreamByTextDecoderStream(
		label.Ref(),
		js.Pointer(&options))
	return
}

func NewTextDecoderStreamByTextDecoderStream1(label js.String) (ret TextDecoderStream) {
	ret.ref = bindings.NewTextDecoderStreamByTextDecoderStream1(
		label.Ref())
	return
}

func NewTextDecoderStreamByTextDecoderStream2() (ret TextDecoderStream) {
	ret.ref = bindings.NewTextDecoderStreamByTextDecoderStream2()
	return
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
// It returns ok=false if there is no such property.
func (this TextDecoderStream) Encoding() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTextDecoderStreamEncoding(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Fatal returns the value of property "TextDecoderStream.fatal".
//
// It returns ok=false if there is no such property.
func (this TextDecoderStream) Fatal() (ret bool, ok bool) {
	ok = js.True == bindings.GetTextDecoderStreamFatal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IgnoreBOM returns the value of property "TextDecoderStream.ignoreBOM".
//
// It returns ok=false if there is no such property.
func (this TextDecoderStream) IgnoreBOM() (ret bool, ok bool) {
	ok = js.True == bindings.GetTextDecoderStreamIgnoreBOM(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Readable returns the value of property "TextDecoderStream.readable".
//
// It returns ok=false if there is no such property.
func (this TextDecoderStream) Readable() (ret ReadableStream, ok bool) {
	ok = js.True == bindings.GetTextDecoderStreamReadable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Writable returns the value of property "TextDecoderStream.writable".
//
// It returns ok=false if there is no such property.
func (this TextDecoderStream) Writable() (ret WritableStream, ok bool) {
	ok = js.True == bindings.GetTextDecoderStreamWritable(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

// HasDetect returns true if the method "TextDetector.detect" exists.
func (this TextDetector) HasDetect() bool {
	return js.True == bindings.HasTextDetectorDetect(
		this.Ref(),
	)
}

// DetectFunc returns the method "TextDetector.detect".
func (this TextDetector) DetectFunc() (fn js.Func[func(image ImageBitmapSource) js.Promise[js.Array[DetectedText]]]) {
	return fn.FromRef(
		bindings.TextDetectorDetectFunc(
			this.Ref(),
		),
	)
}

// Detect calls the method "TextDetector.detect".
func (this TextDetector) Detect(image ImageBitmapSource) (ret js.Promise[js.Array[DetectedText]]) {
	bindings.CallTextDetectorDetect(
		this.Ref(), js.Pointer(&ret),
		image.Ref(),
	)

	return
}

// TryDetect calls the method "TextDetector.detect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TextDetector) TryDetect(image ImageBitmapSource) (ret js.Promise[js.Array[DetectedText]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextDetectorDetect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this TextEncoder) Encoding() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTextEncoderEncoding(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasEncode returns true if the method "TextEncoder.encode" exists.
func (this TextEncoder) HasEncode() bool {
	return js.True == bindings.HasTextEncoderEncode(
		this.Ref(),
	)
}

// EncodeFunc returns the method "TextEncoder.encode".
func (this TextEncoder) EncodeFunc() (fn js.Func[func(input js.String) js.TypedArray[uint8]]) {
	return fn.FromRef(
		bindings.TextEncoderEncodeFunc(
			this.Ref(),
		),
	)
}

// Encode calls the method "TextEncoder.encode".
func (this TextEncoder) Encode(input js.String) (ret js.TypedArray[uint8]) {
	bindings.CallTextEncoderEncode(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryEncode calls the method "TextEncoder.encode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TextEncoder) TryEncode(input js.String) (ret js.TypedArray[uint8], exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextEncoderEncode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasEncode1 returns true if the method "TextEncoder.encode" exists.
func (this TextEncoder) HasEncode1() bool {
	return js.True == bindings.HasTextEncoderEncode1(
		this.Ref(),
	)
}

// Encode1Func returns the method "TextEncoder.encode".
func (this TextEncoder) Encode1Func() (fn js.Func[func() js.TypedArray[uint8]]) {
	return fn.FromRef(
		bindings.TextEncoderEncode1Func(
			this.Ref(),
		),
	)
}

// Encode1 calls the method "TextEncoder.encode".
func (this TextEncoder) Encode1() (ret js.TypedArray[uint8]) {
	bindings.CallTextEncoderEncode1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryEncode1 calls the method "TextEncoder.encode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TextEncoder) TryEncode1() (ret js.TypedArray[uint8], exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextEncoderEncode1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasEncodeInto returns true if the method "TextEncoder.encodeInto" exists.
func (this TextEncoder) HasEncodeInto() bool {
	return js.True == bindings.HasTextEncoderEncodeInto(
		this.Ref(),
	)
}

// EncodeIntoFunc returns the method "TextEncoder.encodeInto".
func (this TextEncoder) EncodeIntoFunc() (fn js.Func[func(source js.String, destination js.TypedArray[uint8]) TextEncoderEncodeIntoResult]) {
	return fn.FromRef(
		bindings.TextEncoderEncodeIntoFunc(
			this.Ref(),
		),
	)
}

// EncodeInto calls the method "TextEncoder.encodeInto".
func (this TextEncoder) EncodeInto(source js.String, destination js.TypedArray[uint8]) (ret TextEncoderEncodeIntoResult) {
	bindings.CallTextEncoderEncodeInto(
		this.Ref(), js.Pointer(&ret),
		source.Ref(),
		destination.Ref(),
	)

	return
}

// TryEncodeInto calls the method "TextEncoder.encodeInto"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TextEncoder) TryEncodeInto(source js.String, destination js.TypedArray[uint8]) (ret TextEncoderEncodeIntoResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextEncoderEncodeInto(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		source.Ref(),
		destination.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this TextEncoderStream) Encoding() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTextEncoderStreamEncoding(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Readable returns the value of property "TextEncoderStream.readable".
//
// It returns ok=false if there is no such property.
func (this TextEncoderStream) Readable() (ret ReadableStream, ok bool) {
	ok = js.True == bindings.GetTextEncoderStreamReadable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Writable returns the value of property "TextEncoderStream.writable".
//
// It returns ok=false if there is no such property.
func (this TextEncoderStream) Writable() (ret WritableStream, ok bool) {
	ok = js.True == bindings.GetTextEncoderStreamWritable(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

func NewTextFormat(options TextFormatInit) (ret TextFormat) {
	ret.ref = bindings.NewTextFormatByTextFormat(
		js.Pointer(&options))
	return
}

func NewTextFormatByTextFormat1() (ret TextFormat) {
	ret.ref = bindings.NewTextFormatByTextFormat1()
	return
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
// It returns ok=false if there is no such property.
func (this TextFormat) RangeStart() (ret uint32, ok bool) {
	ok = js.True == bindings.GetTextFormatRangeStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RangeEnd returns the value of property "TextFormat.rangeEnd".
//
// It returns ok=false if there is no such property.
func (this TextFormat) RangeEnd() (ret uint32, ok bool) {
	ok = js.True == bindings.GetTextFormatRangeEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// UnderlineStyle returns the value of property "TextFormat.underlineStyle".
//
// It returns ok=false if there is no such property.
func (this TextFormat) UnderlineStyle() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTextFormatUnderlineStyle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// UnderlineThickness returns the value of property "TextFormat.underlineThickness".
//
// It returns ok=false if there is no such property.
func (this TextFormat) UnderlineThickness() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTextFormatUnderlineThickness(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

func NewTextFormatUpdateEvent(typ js.String, options TextFormatUpdateEventInit) (ret TextFormatUpdateEvent) {
	ret.ref = bindings.NewTextFormatUpdateEventByTextFormatUpdateEvent(
		typ.Ref(),
		js.Pointer(&options))
	return
}

func NewTextFormatUpdateEventByTextFormatUpdateEvent1(typ js.String) (ret TextFormatUpdateEvent) {
	ret.ref = bindings.NewTextFormatUpdateEventByTextFormatUpdateEvent1(
		typ.Ref())
	return
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

// HasGetTextFormats returns true if the method "TextFormatUpdateEvent.getTextFormats" exists.
func (this TextFormatUpdateEvent) HasGetTextFormats() bool {
	return js.True == bindings.HasTextFormatUpdateEventGetTextFormats(
		this.Ref(),
	)
}

// GetTextFormatsFunc returns the method "TextFormatUpdateEvent.getTextFormats".
func (this TextFormatUpdateEvent) GetTextFormatsFunc() (fn js.Func[func() js.Array[TextFormat]]) {
	return fn.FromRef(
		bindings.TextFormatUpdateEventGetTextFormatsFunc(
			this.Ref(),
		),
	)
}

// GetTextFormats calls the method "TextFormatUpdateEvent.getTextFormats".
func (this TextFormatUpdateEvent) GetTextFormats() (ret js.Array[TextFormat]) {
	bindings.CallTextFormatUpdateEventGetTextFormats(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetTextFormats calls the method "TextFormatUpdateEvent.getTextFormats"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TextFormatUpdateEvent) TryGetTextFormats() (ret js.Array[TextFormat], exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextFormatUpdateEventGetTextFormats(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

func NewTextUpdateEvent(typ js.String, options TextUpdateEventInit) (ret TextUpdateEvent) {
	ret.ref = bindings.NewTextUpdateEventByTextUpdateEvent(
		typ.Ref(),
		js.Pointer(&options))
	return
}

func NewTextUpdateEventByTextUpdateEvent1(typ js.String) (ret TextUpdateEvent) {
	ret.ref = bindings.NewTextUpdateEventByTextUpdateEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this TextUpdateEvent) UpdateRangeStart() (ret uint32, ok bool) {
	ok = js.True == bindings.GetTextUpdateEventUpdateRangeStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// UpdateRangeEnd returns the value of property "TextUpdateEvent.updateRangeEnd".
//
// It returns ok=false if there is no such property.
func (this TextUpdateEvent) UpdateRangeEnd() (ret uint32, ok bool) {
	ok = js.True == bindings.GetTextUpdateEventUpdateRangeEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Text returns the value of property "TextUpdateEvent.text".
//
// It returns ok=false if there is no such property.
func (this TextUpdateEvent) Text() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTextUpdateEventText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SelectionStart returns the value of property "TextUpdateEvent.selectionStart".
//
// It returns ok=false if there is no such property.
func (this TextUpdateEvent) SelectionStart() (ret uint32, ok bool) {
	ok = js.True == bindings.GetTextUpdateEventSelectionStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SelectionEnd returns the value of property "TextUpdateEvent.selectionEnd".
//
// It returns ok=false if there is no such property.
func (this TextUpdateEvent) SelectionEnd() (ret uint32, ok bool) {
	ok = js.True == bindings.GetTextUpdateEventSelectionEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CompositionStart returns the value of property "TextUpdateEvent.compositionStart".
//
// It returns ok=false if there is no such property.
func (this TextUpdateEvent) CompositionStart() (ret uint32, ok bool) {
	ok = js.True == bindings.GetTextUpdateEventCompositionStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CompositionEnd returns the value of property "TextUpdateEvent.compositionEnd".
//
// It returns ok=false if there is no such property.
func (this TextUpdateEvent) CompositionEnd() (ret uint32, ok bool) {
	ok = js.True == bindings.GetTextUpdateEventCompositionEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewTimeEvent(typ js.String, eventInitDict EventInit) (ret TimeEvent) {
	ret.ref = bindings.NewTimeEventByTimeEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewTimeEventByTimeEvent1(typ js.String) (ret TimeEvent) {
	ret.ref = bindings.NewTimeEventByTimeEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this TimeEvent) View() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetTimeEventView(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Detail returns the value of property "TimeEvent.detail".
//
// It returns ok=false if there is no such property.
func (this TimeEvent) Detail() (ret int32, ok bool) {
	ok = js.True == bindings.GetTimeEventDetail(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasInitTimeEvent returns true if the method "TimeEvent.initTimeEvent" exists.
func (this TimeEvent) HasInitTimeEvent() bool {
	return js.True == bindings.HasTimeEventInitTimeEvent(
		this.Ref(),
	)
}

// InitTimeEventFunc returns the method "TimeEvent.initTimeEvent".
func (this TimeEvent) InitTimeEventFunc() (fn js.Func[func(typeArg js.String, viewArg Window, detailArg int32)]) {
	return fn.FromRef(
		bindings.TimeEventInitTimeEventFunc(
			this.Ref(),
		),
	)
}

// InitTimeEvent calls the method "TimeEvent.initTimeEvent".
func (this TimeEvent) InitTimeEvent(typeArg js.String, viewArg Window, detailArg int32) (ret js.Void) {
	bindings.CallTimeEventInitTimeEvent(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		viewArg.Ref(),
		int32(detailArg),
	)

	return
}

// TryInitTimeEvent calls the method "TimeEvent.initTimeEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TimeEvent) TryInitTimeEvent(typeArg js.String, viewArg Window, detailArg int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTimeEventInitTimeEvent(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		viewArg.Ref(),
		int32(detailArg),
	)

	return
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

func NewToggleEvent(typ js.String, eventInitDict ToggleEventInit) (ret ToggleEvent) {
	ret.ref = bindings.NewToggleEventByToggleEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewToggleEventByToggleEvent1(typ js.String) (ret ToggleEvent) {
	ret.ref = bindings.NewToggleEventByToggleEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this ToggleEvent) OldState() (ret js.String, ok bool) {
	ok = js.True == bindings.GetToggleEventOldState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NewState returns the value of property "ToggleEvent.newState".
//
// It returns ok=false if there is no such property.
func (this ToggleEvent) NewState() (ret js.String, ok bool) {
	ok = js.True == bindings.GetToggleEventNewState(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

func NewTouch(touchInitDict TouchInit) (ret Touch) {
	ret.ref = bindings.NewTouchByTouch(
		js.Pointer(&touchInitDict))
	return
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
// It returns ok=false if there is no such property.
func (this Touch) Identifier() (ret int32, ok bool) {
	ok = js.True == bindings.GetTouchIdentifier(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Target returns the value of property "Touch.target".
//
// It returns ok=false if there is no such property.
func (this Touch) Target() (ret EventTarget, ok bool) {
	ok = js.True == bindings.GetTouchTarget(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ScreenX returns the value of property "Touch.screenX".
//
// It returns ok=false if there is no such property.
func (this Touch) ScreenX() (ret float64, ok bool) {
	ok = js.True == bindings.GetTouchScreenX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ScreenY returns the value of property "Touch.screenY".
//
// It returns ok=false if there is no such property.
func (this Touch) ScreenY() (ret float64, ok bool) {
	ok = js.True == bindings.GetTouchScreenY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ClientX returns the value of property "Touch.clientX".
//
// It returns ok=false if there is no such property.
func (this Touch) ClientX() (ret float64, ok bool) {
	ok = js.True == bindings.GetTouchClientX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ClientY returns the value of property "Touch.clientY".
//
// It returns ok=false if there is no such property.
func (this Touch) ClientY() (ret float64, ok bool) {
	ok = js.True == bindings.GetTouchClientY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PageX returns the value of property "Touch.pageX".
//
// It returns ok=false if there is no such property.
func (this Touch) PageX() (ret float64, ok bool) {
	ok = js.True == bindings.GetTouchPageX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PageY returns the value of property "Touch.pageY".
//
// It returns ok=false if there is no such property.
func (this Touch) PageY() (ret float64, ok bool) {
	ok = js.True == bindings.GetTouchPageY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RadiusX returns the value of property "Touch.radiusX".
//
// It returns ok=false if there is no such property.
func (this Touch) RadiusX() (ret float32, ok bool) {
	ok = js.True == bindings.GetTouchRadiusX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RadiusY returns the value of property "Touch.radiusY".
//
// It returns ok=false if there is no such property.
func (this Touch) RadiusY() (ret float32, ok bool) {
	ok = js.True == bindings.GetTouchRadiusY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RotationAngle returns the value of property "Touch.rotationAngle".
//
// It returns ok=false if there is no such property.
func (this Touch) RotationAngle() (ret float32, ok bool) {
	ok = js.True == bindings.GetTouchRotationAngle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Force returns the value of property "Touch.force".
//
// It returns ok=false if there is no such property.
func (this Touch) Force() (ret float32, ok bool) {
	ok = js.True == bindings.GetTouchForce(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AltitudeAngle returns the value of property "Touch.altitudeAngle".
//
// It returns ok=false if there is no such property.
func (this Touch) AltitudeAngle() (ret float32, ok bool) {
	ok = js.True == bindings.GetTouchAltitudeAngle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AzimuthAngle returns the value of property "Touch.azimuthAngle".
//
// It returns ok=false if there is no such property.
func (this Touch) AzimuthAngle() (ret float32, ok bool) {
	ok = js.True == bindings.GetTouchAzimuthAngle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// TouchType returns the value of property "Touch.touchType".
//
// It returns ok=false if there is no such property.
func (this Touch) TouchType() (ret TouchType, ok bool) {
	ok = js.True == bindings.GetTouchTouchType(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this TouchList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetTouchListLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasItem returns true if the method "TouchList.item" exists.
func (this TouchList) HasItem() bool {
	return js.True == bindings.HasTouchListItem(
		this.Ref(),
	)
}

// ItemFunc returns the method "TouchList.item".
func (this TouchList) ItemFunc() (fn js.Func[func(index uint32) Touch]) {
	return fn.FromRef(
		bindings.TouchListItemFunc(
			this.Ref(),
		),
	)
}

// Item calls the method "TouchList.item".
func (this TouchList) Item(index uint32) (ret Touch) {
	bindings.CallTouchListItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "TouchList.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TouchList) TryItem(index uint32) (ret Touch, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTouchListItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

func NewTouchEvent(typ js.String, eventInitDict TouchEventInit) (ret TouchEvent) {
	ret.ref = bindings.NewTouchEventByTouchEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewTouchEventByTouchEvent1(typ js.String) (ret TouchEvent) {
	ret.ref = bindings.NewTouchEventByTouchEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this TouchEvent) Touches() (ret TouchList, ok bool) {
	ok = js.True == bindings.GetTouchEventTouches(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// TargetTouches returns the value of property "TouchEvent.targetTouches".
//
// It returns ok=false if there is no such property.
func (this TouchEvent) TargetTouches() (ret TouchList, ok bool) {
	ok = js.True == bindings.GetTouchEventTargetTouches(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ChangedTouches returns the value of property "TouchEvent.changedTouches".
//
// It returns ok=false if there is no such property.
func (this TouchEvent) ChangedTouches() (ret TouchList, ok bool) {
	ok = js.True == bindings.GetTouchEventChangedTouches(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AltKey returns the value of property "TouchEvent.altKey".
//
// It returns ok=false if there is no such property.
func (this TouchEvent) AltKey() (ret bool, ok bool) {
	ok = js.True == bindings.GetTouchEventAltKey(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MetaKey returns the value of property "TouchEvent.metaKey".
//
// It returns ok=false if there is no such property.
func (this TouchEvent) MetaKey() (ret bool, ok bool) {
	ok = js.True == bindings.GetTouchEventMetaKey(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CtrlKey returns the value of property "TouchEvent.ctrlKey".
//
// It returns ok=false if there is no such property.
func (this TouchEvent) CtrlKey() (ret bool, ok bool) {
	ok = js.True == bindings.GetTouchEventCtrlKey(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ShiftKey returns the value of property "TouchEvent.shiftKey".
//
// It returns ok=false if there is no such property.
func (this TouchEvent) ShiftKey() (ret bool, ok bool) {
	ok = js.True == bindings.GetTouchEventShiftKey(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetModifierState returns true if the method "TouchEvent.getModifierState" exists.
func (this TouchEvent) HasGetModifierState() bool {
	return js.True == bindings.HasTouchEventGetModifierState(
		this.Ref(),
	)
}

// GetModifierStateFunc returns the method "TouchEvent.getModifierState".
func (this TouchEvent) GetModifierStateFunc() (fn js.Func[func(keyArg js.String) bool]) {
	return fn.FromRef(
		bindings.TouchEventGetModifierStateFunc(
			this.Ref(),
		),
	)
}

// GetModifierState calls the method "TouchEvent.getModifierState".
func (this TouchEvent) GetModifierState(keyArg js.String) (ret bool) {
	bindings.CallTouchEventGetModifierState(
		this.Ref(), js.Pointer(&ret),
		keyArg.Ref(),
	)

	return
}

// TryGetModifierState calls the method "TouchEvent.getModifierState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TouchEvent) TryGetModifierState(keyArg js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTouchEventGetModifierState(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		keyArg.Ref(),
	)

	return
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

func NewTrackEvent(typ js.String, eventInitDict TrackEventInit) (ret TrackEvent) {
	ret.ref = bindings.NewTrackEventByTrackEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewTrackEventByTrackEvent1(typ js.String) (ret TrackEvent) {
	ret.ref = bindings.NewTrackEventByTrackEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this TrackEvent) Track() (ret OneOf_VideoTrack_AudioTrack_TextTrack, ok bool) {
	ok = js.True == bindings.GetTrackEventTrack(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewTransformStream(transformer js.Object, writableStrategy QueuingStrategy, readableStrategy QueuingStrategy) (ret TransformStream) {
	ret.ref = bindings.NewTransformStreamByTransformStream(
		transformer.Ref(),
		js.Pointer(&writableStrategy),
		js.Pointer(&readableStrategy))
	return
}

func NewTransformStreamByTransformStream1(transformer js.Object, writableStrategy QueuingStrategy) (ret TransformStream) {
	ret.ref = bindings.NewTransformStreamByTransformStream1(
		transformer.Ref(),
		js.Pointer(&writableStrategy))
	return
}

func NewTransformStreamByTransformStream2(transformer js.Object) (ret TransformStream) {
	ret.ref = bindings.NewTransformStreamByTransformStream2(
		transformer.Ref())
	return
}

func NewTransformStreamByTransformStream3() (ret TransformStream) {
	ret.ref = bindings.NewTransformStreamByTransformStream3()
	return
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
// It returns ok=false if there is no such property.
func (this TransformStream) Readable() (ret ReadableStream, ok bool) {
	ok = js.True == bindings.GetTransformStreamReadable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Writable returns the value of property "TransformStream.writable".
//
// It returns ok=false if there is no such property.
func (this TransformStream) Writable() (ret WritableStream, ok bool) {
	ok = js.True == bindings.GetTransformStreamWritable(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this TransformStreamDefaultController) DesiredSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetTransformStreamDefaultControllerDesiredSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasEnqueue returns true if the method "TransformStreamDefaultController.enqueue" exists.
func (this TransformStreamDefaultController) HasEnqueue() bool {
	return js.True == bindings.HasTransformStreamDefaultControllerEnqueue(
		this.Ref(),
	)
}

// EnqueueFunc returns the method "TransformStreamDefaultController.enqueue".
func (this TransformStreamDefaultController) EnqueueFunc() (fn js.Func[func(chunk js.Any)]) {
	return fn.FromRef(
		bindings.TransformStreamDefaultControllerEnqueueFunc(
			this.Ref(),
		),
	)
}

// Enqueue calls the method "TransformStreamDefaultController.enqueue".
func (this TransformStreamDefaultController) Enqueue(chunk js.Any) (ret js.Void) {
	bindings.CallTransformStreamDefaultControllerEnqueue(
		this.Ref(), js.Pointer(&ret),
		chunk.Ref(),
	)

	return
}

// TryEnqueue calls the method "TransformStreamDefaultController.enqueue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TransformStreamDefaultController) TryEnqueue(chunk js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTransformStreamDefaultControllerEnqueue(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		chunk.Ref(),
	)

	return
}

// HasEnqueue1 returns true if the method "TransformStreamDefaultController.enqueue" exists.
func (this TransformStreamDefaultController) HasEnqueue1() bool {
	return js.True == bindings.HasTransformStreamDefaultControllerEnqueue1(
		this.Ref(),
	)
}

// Enqueue1Func returns the method "TransformStreamDefaultController.enqueue".
func (this TransformStreamDefaultController) Enqueue1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.TransformStreamDefaultControllerEnqueue1Func(
			this.Ref(),
		),
	)
}

// Enqueue1 calls the method "TransformStreamDefaultController.enqueue".
func (this TransformStreamDefaultController) Enqueue1() (ret js.Void) {
	bindings.CallTransformStreamDefaultControllerEnqueue1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryEnqueue1 calls the method "TransformStreamDefaultController.enqueue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TransformStreamDefaultController) TryEnqueue1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTransformStreamDefaultControllerEnqueue1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasError returns true if the method "TransformStreamDefaultController.error" exists.
func (this TransformStreamDefaultController) HasError() bool {
	return js.True == bindings.HasTransformStreamDefaultControllerError(
		this.Ref(),
	)
}

// ErrorFunc returns the method "TransformStreamDefaultController.error".
func (this TransformStreamDefaultController) ErrorFunc() (fn js.Func[func(reason js.Any)]) {
	return fn.FromRef(
		bindings.TransformStreamDefaultControllerErrorFunc(
			this.Ref(),
		),
	)
}

// Error calls the method "TransformStreamDefaultController.error".
func (this TransformStreamDefaultController) Error(reason js.Any) (ret js.Void) {
	bindings.CallTransformStreamDefaultControllerError(
		this.Ref(), js.Pointer(&ret),
		reason.Ref(),
	)

	return
}

// TryError calls the method "TransformStreamDefaultController.error"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TransformStreamDefaultController) TryError(reason js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTransformStreamDefaultControllerError(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		reason.Ref(),
	)

	return
}

// HasError1 returns true if the method "TransformStreamDefaultController.error" exists.
func (this TransformStreamDefaultController) HasError1() bool {
	return js.True == bindings.HasTransformStreamDefaultControllerError1(
		this.Ref(),
	)
}

// Error1Func returns the method "TransformStreamDefaultController.error".
func (this TransformStreamDefaultController) Error1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.TransformStreamDefaultControllerError1Func(
			this.Ref(),
		),
	)
}

// Error1 calls the method "TransformStreamDefaultController.error".
func (this TransformStreamDefaultController) Error1() (ret js.Void) {
	bindings.CallTransformStreamDefaultControllerError1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryError1 calls the method "TransformStreamDefaultController.error"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TransformStreamDefaultController) TryError1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTransformStreamDefaultControllerError1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasTerminate returns true if the method "TransformStreamDefaultController.terminate" exists.
func (this TransformStreamDefaultController) HasTerminate() bool {
	return js.True == bindings.HasTransformStreamDefaultControllerTerminate(
		this.Ref(),
	)
}

// TerminateFunc returns the method "TransformStreamDefaultController.terminate".
func (this TransformStreamDefaultController) TerminateFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.TransformStreamDefaultControllerTerminateFunc(
			this.Ref(),
		),
	)
}

// Terminate calls the method "TransformStreamDefaultController.terminate".
func (this TransformStreamDefaultController) Terminate() (ret js.Void) {
	bindings.CallTransformStreamDefaultControllerTerminate(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryTerminate calls the method "TransformStreamDefaultController.terminate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TransformStreamDefaultController) TryTerminate() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTransformStreamDefaultControllerTerminate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
