// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package alarms

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/alarms/bindings"
)

type Alarm struct {
	// Name is "Alarm.name"
	//
	// Optional
	Name js.String
	// ScheduledTime is "Alarm.scheduledTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_ScheduledTime MUST be set to true to make this field effective.
	ScheduledTime float64
	// PeriodInMinutes is "Alarm.periodInMinutes"
	//
	// Optional
	//
	// NOTE: FFI_USE_PeriodInMinutes MUST be set to true to make this field effective.
	PeriodInMinutes float64

	FFI_USE_ScheduledTime   bool // for ScheduledTime.
	FFI_USE_PeriodInMinutes bool // for PeriodInMinutes.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Alarm with all fields set.
func (p Alarm) FromRef(ref js.Ref) Alarm {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Alarm in the application heap.
func (p Alarm) New() js.Ref {
	return bindings.AlarmJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Alarm) UpdateFrom(ref js.Ref) {
	bindings.AlarmJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Alarm) Update(ref js.Ref) {
	bindings.AlarmJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Alarm) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
}

type AlarmCallbackFunc func(this js.Ref, alarm *Alarm) js.Ref

func (fn AlarmCallbackFunc) Register() js.Func[func(alarm *Alarm)] {
	return js.RegisterCallback[func(alarm *Alarm)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn AlarmCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Alarm
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

type AlarmCallback[T any] struct {
	Fn  func(arg T, this js.Ref, alarm *Alarm) js.Ref
	Arg T
}

func (cb *AlarmCallback[T]) Register() js.Func[func(alarm *Alarm)] {
	return js.RegisterCallback[func(alarm *Alarm)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *AlarmCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Alarm
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

type AlarmCreateInfo struct {
	// When is "AlarmCreateInfo.when"
	//
	// Optional
	//
	// NOTE: FFI_USE_When MUST be set to true to make this field effective.
	When float64
	// DelayInMinutes is "AlarmCreateInfo.delayInMinutes"
	//
	// Optional
	//
	// NOTE: FFI_USE_DelayInMinutes MUST be set to true to make this field effective.
	DelayInMinutes float64
	// PeriodInMinutes is "AlarmCreateInfo.periodInMinutes"
	//
	// Optional
	//
	// NOTE: FFI_USE_PeriodInMinutes MUST be set to true to make this field effective.
	PeriodInMinutes float64

	FFI_USE_When            bool // for When.
	FFI_USE_DelayInMinutes  bool // for DelayInMinutes.
	FFI_USE_PeriodInMinutes bool // for PeriodInMinutes.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AlarmCreateInfo with all fields set.
func (p AlarmCreateInfo) FromRef(ref js.Ref) AlarmCreateInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AlarmCreateInfo in the application heap.
func (p AlarmCreateInfo) New() js.Ref {
	return bindings.AlarmCreateInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AlarmCreateInfo) UpdateFrom(ref js.Ref) {
	bindings.AlarmCreateInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AlarmCreateInfo) Update(ref js.Ref) {
	bindings.AlarmCreateInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AlarmCreateInfo) FreeMembers(recursive bool) {
}

type AlarmListCallbackFunc func(this js.Ref, alarms js.Array[Alarm]) js.Ref

func (fn AlarmListCallbackFunc) Register() js.Func[func(alarms js.Array[Alarm])] {
	return js.RegisterCallback[func(alarms js.Array[Alarm])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn AlarmListCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[Alarm]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type AlarmListCallback[T any] struct {
	Fn  func(arg T, this js.Ref, alarms js.Array[Alarm]) js.Ref
	Arg T
}

func (cb *AlarmListCallback[T]) Register() js.Func[func(alarms js.Array[Alarm])] {
	return js.RegisterCallback[func(alarms js.Array[Alarm])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *AlarmListCallback[T]) DispatchCallback(
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

		js.Array[Alarm]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ClearCallbackFunc func(this js.Ref, wasCleared bool) js.Ref

func (fn ClearCallbackFunc) Register() js.Func[func(wasCleared bool)] {
	return js.RegisterCallback[func(wasCleared bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ClearCallbackFunc) DispatchCallback(
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

type ClearCallback[T any] struct {
	Fn  func(arg T, this js.Ref, wasCleared bool) js.Ref
	Arg T
}

func (cb *ClearCallback[T]) Register() js.Func[func(wasCleared bool)] {
	return js.RegisterCallback[func(wasCleared bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ClearCallback[T]) DispatchCallback(
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

type VoidCallbackFunc func(this js.Ref) js.Ref

func (fn VoidCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn VoidCallbackFunc) DispatchCallback(
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

type VoidCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *VoidCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *VoidCallback[T]) DispatchCallback(
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

// HasFuncClear returns true if the function "WEBEXT.alarms.clear" exists.
func HasFuncClear() bool {
	return js.True == bindings.HasFuncClear()
}

// FuncClear returns the function "WEBEXT.alarms.clear".
func FuncClear() (fn js.Func[func(name js.String) js.Promise[js.Boolean]]) {
	bindings.FuncClear(
		js.Pointer(&fn),
	)
	return
}

// Clear calls the function "WEBEXT.alarms.clear" directly.
func Clear(name js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallClear(
		js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryClear calls the function "WEBEXT.alarms.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryClear(name js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClear(
		js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncClearAll returns true if the function "WEBEXT.alarms.clearAll" exists.
func HasFuncClearAll() bool {
	return js.True == bindings.HasFuncClearAll()
}

// FuncClearAll returns the function "WEBEXT.alarms.clearAll".
func FuncClearAll() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncClearAll(
		js.Pointer(&fn),
	)
	return
}

// ClearAll calls the function "WEBEXT.alarms.clearAll" directly.
func ClearAll() (ret js.Promise[js.Boolean]) {
	bindings.CallClearAll(
		js.Pointer(&ret),
	)

	return
}

// TryClearAll calls the function "WEBEXT.alarms.clearAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryClearAll() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClearAll(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreate returns true if the function "WEBEXT.alarms.create" exists.
func HasFuncCreate() bool {
	return js.True == bindings.HasFuncCreate()
}

// FuncCreate returns the function "WEBEXT.alarms.create".
func FuncCreate() (fn js.Func[func(name js.String, alarmInfo AlarmCreateInfo) js.Promise[js.Void]]) {
	bindings.FuncCreate(
		js.Pointer(&fn),
	)
	return
}

// Create calls the function "WEBEXT.alarms.create" directly.
func Create(name js.String, alarmInfo AlarmCreateInfo) (ret js.Promise[js.Void]) {
	bindings.CallCreate(
		js.Pointer(&ret),
		name.Ref(),
		js.Pointer(&alarmInfo),
	)

	return
}

// TryCreate calls the function "WEBEXT.alarms.create"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCreate(name js.String, alarmInfo AlarmCreateInfo) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCreate(
		js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		js.Pointer(&alarmInfo),
	)

	return
}

// HasFuncGet returns true if the function "WEBEXT.alarms.get" exists.
func HasFuncGet() bool {
	return js.True == bindings.HasFuncGet()
}

// FuncGet returns the function "WEBEXT.alarms.get".
func FuncGet() (fn js.Func[func(name js.String) js.Promise[Alarm]]) {
	bindings.FuncGet(
		js.Pointer(&fn),
	)
	return
}

// Get calls the function "WEBEXT.alarms.get" directly.
func Get(name js.String) (ret js.Promise[Alarm]) {
	bindings.CallGet(
		js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGet calls the function "WEBEXT.alarms.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGet(name js.String) (ret js.Promise[Alarm], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGet(
		js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncGetAll returns true if the function "WEBEXT.alarms.getAll" exists.
func HasFuncGetAll() bool {
	return js.True == bindings.HasFuncGetAll()
}

// FuncGetAll returns the function "WEBEXT.alarms.getAll".
func FuncGetAll() (fn js.Func[func() js.Promise[js.Array[Alarm]]]) {
	bindings.FuncGetAll(
		js.Pointer(&fn),
	)
	return
}

// GetAll calls the function "WEBEXT.alarms.getAll" directly.
func GetAll() (ret js.Promise[js.Array[Alarm]]) {
	bindings.CallGetAll(
		js.Pointer(&ret),
	)

	return
}

// TryGetAll calls the function "WEBEXT.alarms.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAll() (ret js.Promise[js.Array[Alarm]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAll(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnAlarmEventCallbackFunc func(this js.Ref, alarm *Alarm) js.Ref

func (fn OnAlarmEventCallbackFunc) Register() js.Func[func(alarm *Alarm)] {
	return js.RegisterCallback[func(alarm *Alarm)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnAlarmEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Alarm
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

type OnAlarmEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, alarm *Alarm) js.Ref
	Arg T
}

func (cb *OnAlarmEventCallback[T]) Register() js.Func[func(alarm *Alarm)] {
	return js.RegisterCallback[func(alarm *Alarm)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnAlarmEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Alarm
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

// HasFuncOnAlarm returns true if the function "WEBEXT.alarms.onAlarm.addListener" exists.
func HasFuncOnAlarm() bool {
	return js.True == bindings.HasFuncOnAlarm()
}

// FuncOnAlarm returns the function "WEBEXT.alarms.onAlarm.addListener".
func FuncOnAlarm() (fn js.Func[func(callback js.Func[func(alarm *Alarm)])]) {
	bindings.FuncOnAlarm(
		js.Pointer(&fn),
	)
	return
}

// OnAlarm calls the function "WEBEXT.alarms.onAlarm.addListener" directly.
func OnAlarm(callback js.Func[func(alarm *Alarm)]) (ret js.Void) {
	bindings.CallOnAlarm(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAlarm calls the function "WEBEXT.alarms.onAlarm.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAlarm(callback js.Func[func(alarm *Alarm)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAlarm(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAlarm returns true if the function "WEBEXT.alarms.onAlarm.removeListener" exists.
func HasFuncOffAlarm() bool {
	return js.True == bindings.HasFuncOffAlarm()
}

// FuncOffAlarm returns the function "WEBEXT.alarms.onAlarm.removeListener".
func FuncOffAlarm() (fn js.Func[func(callback js.Func[func(alarm *Alarm)])]) {
	bindings.FuncOffAlarm(
		js.Pointer(&fn),
	)
	return
}

// OffAlarm calls the function "WEBEXT.alarms.onAlarm.removeListener" directly.
func OffAlarm(callback js.Func[func(alarm *Alarm)]) (ret js.Void) {
	bindings.CallOffAlarm(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAlarm calls the function "WEBEXT.alarms.onAlarm.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAlarm(callback js.Func[func(alarm *Alarm)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAlarm(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAlarm returns true if the function "WEBEXT.alarms.onAlarm.hasListener" exists.
func HasFuncHasOnAlarm() bool {
	return js.True == bindings.HasFuncHasOnAlarm()
}

// FuncHasOnAlarm returns the function "WEBEXT.alarms.onAlarm.hasListener".
func FuncHasOnAlarm() (fn js.Func[func(callback js.Func[func(alarm *Alarm)]) bool]) {
	bindings.FuncHasOnAlarm(
		js.Pointer(&fn),
	)
	return
}

// HasOnAlarm calls the function "WEBEXT.alarms.onAlarm.hasListener" directly.
func HasOnAlarm(callback js.Func[func(alarm *Alarm)]) (ret bool) {
	bindings.CallHasOnAlarm(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAlarm calls the function "WEBEXT.alarms.onAlarm.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnAlarm(callback js.Func[func(alarm *Alarm)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnAlarm(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}
