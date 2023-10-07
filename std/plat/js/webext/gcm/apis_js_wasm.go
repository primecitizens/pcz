// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package gcm

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/gcm/bindings"
)

// MAX_MESSAGE_SIZE returns the value of property "WEBEXT.gcm.MAX_MESSAGE_SIZE".
//
// The returned bool will be false if there is no such property.
func MAX_MESSAGE_SIZE() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMAX_MESSAGE_SIZE(
		js.Pointer(&ret),
	)

	return
}

// SetMAX_MESSAGE_SIZE sets the value of property "WEBEXT.gcm.MAX_MESSAGE_SIZE" to val.
//
// It returns false if the property cannot be set.
func SetMAX_MESSAGE_SIZE(val js.String) bool {
	return js.True == bindings.SetMAX_MESSAGE_SIZE(
		val.Ref())
}

type OnMessageArgMessage struct {
	// CollapseKey is "OnMessageArgMessage.collapseKey"
	//
	// Optional
	CollapseKey js.String
	// Data is "OnMessageArgMessage.data"
	//
	// Required
	Data js.String
	// From is "OnMessageArgMessage.from"
	//
	// Optional
	From js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnMessageArgMessage with all fields set.
func (p OnMessageArgMessage) FromRef(ref js.Ref) OnMessageArgMessage {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnMessageArgMessage in the application heap.
func (p OnMessageArgMessage) New() js.Ref {
	return bindings.OnMessageArgMessageJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnMessageArgMessage) UpdateFrom(ref js.Ref) {
	bindings.OnMessageArgMessageJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnMessageArgMessage) Update(ref js.Ref) {
	bindings.OnMessageArgMessageJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnMessageArgMessage) FreeMembers(recursive bool) {
	js.Free(
		p.CollapseKey.Ref(),
		p.Data.Ref(),
		p.From.Ref(),
	)
	p.CollapseKey = p.CollapseKey.FromRef(js.Undefined)
	p.Data = p.Data.FromRef(js.Undefined)
	p.From = p.From.FromRef(js.Undefined)
}

type OnSendErrorArgError struct {
	// Details is "OnSendErrorArgError.details"
	//
	// Required
	Details js.String
	// ErrorMessage is "OnSendErrorArgError.errorMessage"
	//
	// Required
	ErrorMessage js.String
	// MessageId is "OnSendErrorArgError.messageId"
	//
	// Optional
	MessageId js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnSendErrorArgError with all fields set.
func (p OnSendErrorArgError) FromRef(ref js.Ref) OnSendErrorArgError {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnSendErrorArgError in the application heap.
func (p OnSendErrorArgError) New() js.Ref {
	return bindings.OnSendErrorArgErrorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnSendErrorArgError) UpdateFrom(ref js.Ref) {
	bindings.OnSendErrorArgErrorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnSendErrorArgError) Update(ref js.Ref) {
	bindings.OnSendErrorArgErrorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnSendErrorArgError) FreeMembers(recursive bool) {
	js.Free(
		p.Details.Ref(),
		p.ErrorMessage.Ref(),
		p.MessageId.Ref(),
	)
	p.Details = p.Details.FromRef(js.Undefined)
	p.ErrorMessage = p.ErrorMessage.FromRef(js.Undefined)
	p.MessageId = p.MessageId.FromRef(js.Undefined)
}

type SendArgMessage struct {
	// Data is "SendArgMessage.data"
	//
	// Required
	Data js.String
	// DestinationId is "SendArgMessage.destinationId"
	//
	// Required
	DestinationId js.String
	// MessageId is "SendArgMessage.messageId"
	//
	// Required
	MessageId js.String
	// TimeToLive is "SendArgMessage.timeToLive"
	//
	// Optional
	//
	// NOTE: FFI_USE_TimeToLive MUST be set to true to make this field effective.
	TimeToLive int64

	FFI_USE_TimeToLive bool // for TimeToLive.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SendArgMessage with all fields set.
func (p SendArgMessage) FromRef(ref js.Ref) SendArgMessage {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SendArgMessage in the application heap.
func (p SendArgMessage) New() js.Ref {
	return bindings.SendArgMessageJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SendArgMessage) UpdateFrom(ref js.Ref) {
	bindings.SendArgMessageJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SendArgMessage) Update(ref js.Ref) {
	bindings.SendArgMessageJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SendArgMessage) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
		p.DestinationId.Ref(),
		p.MessageId.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
	p.DestinationId = p.DestinationId.FromRef(js.Undefined)
	p.MessageId = p.MessageId.FromRef(js.Undefined)
}

type OnMessageEventCallbackFunc func(this js.Ref, message *OnMessageArgMessage) js.Ref

func (fn OnMessageEventCallbackFunc) Register() js.Func[func(message *OnMessageArgMessage)] {
	return js.RegisterCallback[func(message *OnMessageArgMessage)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnMessageEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnMessageArgMessage
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

type OnMessageEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, message *OnMessageArgMessage) js.Ref
	Arg T
}

func (cb *OnMessageEventCallback[T]) Register() js.Func[func(message *OnMessageArgMessage)] {
	return js.RegisterCallback[func(message *OnMessageArgMessage)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnMessageEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnMessageArgMessage
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

// HasFuncOnMessage returns true if the function "WEBEXT.gcm.onMessage.addListener" exists.
func HasFuncOnMessage() bool {
	return js.True == bindings.HasFuncOnMessage()
}

// FuncOnMessage returns the function "WEBEXT.gcm.onMessage.addListener".
func FuncOnMessage() (fn js.Func[func(callback js.Func[func(message *OnMessageArgMessage)])]) {
	bindings.FuncOnMessage(
		js.Pointer(&fn),
	)
	return
}

// OnMessage calls the function "WEBEXT.gcm.onMessage.addListener" directly.
func OnMessage(callback js.Func[func(message *OnMessageArgMessage)]) (ret js.Void) {
	bindings.CallOnMessage(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMessage calls the function "WEBEXT.gcm.onMessage.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMessage(callback js.Func[func(message *OnMessageArgMessage)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMessage(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMessage returns true if the function "WEBEXT.gcm.onMessage.removeListener" exists.
func HasFuncOffMessage() bool {
	return js.True == bindings.HasFuncOffMessage()
}

// FuncOffMessage returns the function "WEBEXT.gcm.onMessage.removeListener".
func FuncOffMessage() (fn js.Func[func(callback js.Func[func(message *OnMessageArgMessage)])]) {
	bindings.FuncOffMessage(
		js.Pointer(&fn),
	)
	return
}

// OffMessage calls the function "WEBEXT.gcm.onMessage.removeListener" directly.
func OffMessage(callback js.Func[func(message *OnMessageArgMessage)]) (ret js.Void) {
	bindings.CallOffMessage(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMessage calls the function "WEBEXT.gcm.onMessage.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMessage(callback js.Func[func(message *OnMessageArgMessage)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMessage(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMessage returns true if the function "WEBEXT.gcm.onMessage.hasListener" exists.
func HasFuncHasOnMessage() bool {
	return js.True == bindings.HasFuncHasOnMessage()
}

// FuncHasOnMessage returns the function "WEBEXT.gcm.onMessage.hasListener".
func FuncHasOnMessage() (fn js.Func[func(callback js.Func[func(message *OnMessageArgMessage)]) bool]) {
	bindings.FuncHasOnMessage(
		js.Pointer(&fn),
	)
	return
}

// HasOnMessage calls the function "WEBEXT.gcm.onMessage.hasListener" directly.
func HasOnMessage(callback js.Func[func(message *OnMessageArgMessage)]) (ret bool) {
	bindings.CallHasOnMessage(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMessage calls the function "WEBEXT.gcm.onMessage.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnMessage(callback js.Func[func(message *OnMessageArgMessage)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnMessage(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnMessagesDeletedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnMessagesDeletedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnMessagesDeletedEventCallbackFunc) DispatchCallback(
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

type OnMessagesDeletedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnMessagesDeletedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnMessagesDeletedEventCallback[T]) DispatchCallback(
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

// HasFuncOnMessagesDeleted returns true if the function "WEBEXT.gcm.onMessagesDeleted.addListener" exists.
func HasFuncOnMessagesDeleted() bool {
	return js.True == bindings.HasFuncOnMessagesDeleted()
}

// FuncOnMessagesDeleted returns the function "WEBEXT.gcm.onMessagesDeleted.addListener".
func FuncOnMessagesDeleted() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnMessagesDeleted(
		js.Pointer(&fn),
	)
	return
}

// OnMessagesDeleted calls the function "WEBEXT.gcm.onMessagesDeleted.addListener" directly.
func OnMessagesDeleted(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnMessagesDeleted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMessagesDeleted calls the function "WEBEXT.gcm.onMessagesDeleted.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMessagesDeleted(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMessagesDeleted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMessagesDeleted returns true if the function "WEBEXT.gcm.onMessagesDeleted.removeListener" exists.
func HasFuncOffMessagesDeleted() bool {
	return js.True == bindings.HasFuncOffMessagesDeleted()
}

// FuncOffMessagesDeleted returns the function "WEBEXT.gcm.onMessagesDeleted.removeListener".
func FuncOffMessagesDeleted() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffMessagesDeleted(
		js.Pointer(&fn),
	)
	return
}

// OffMessagesDeleted calls the function "WEBEXT.gcm.onMessagesDeleted.removeListener" directly.
func OffMessagesDeleted(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffMessagesDeleted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMessagesDeleted calls the function "WEBEXT.gcm.onMessagesDeleted.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMessagesDeleted(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMessagesDeleted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMessagesDeleted returns true if the function "WEBEXT.gcm.onMessagesDeleted.hasListener" exists.
func HasFuncHasOnMessagesDeleted() bool {
	return js.True == bindings.HasFuncHasOnMessagesDeleted()
}

// FuncHasOnMessagesDeleted returns the function "WEBEXT.gcm.onMessagesDeleted.hasListener".
func FuncHasOnMessagesDeleted() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnMessagesDeleted(
		js.Pointer(&fn),
	)
	return
}

// HasOnMessagesDeleted calls the function "WEBEXT.gcm.onMessagesDeleted.hasListener" directly.
func HasOnMessagesDeleted(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnMessagesDeleted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMessagesDeleted calls the function "WEBEXT.gcm.onMessagesDeleted.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnMessagesDeleted(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnMessagesDeleted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSendErrorEventCallbackFunc func(this js.Ref, err *OnSendErrorArgError) js.Ref

func (fn OnSendErrorEventCallbackFunc) Register() js.Func[func(err *OnSendErrorArgError)] {
	return js.RegisterCallback[func(err *OnSendErrorArgError)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSendErrorEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnSendErrorArgError
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

type OnSendErrorEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, err *OnSendErrorArgError) js.Ref
	Arg T
}

func (cb *OnSendErrorEventCallback[T]) Register() js.Func[func(err *OnSendErrorArgError)] {
	return js.RegisterCallback[func(err *OnSendErrorArgError)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSendErrorEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnSendErrorArgError
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

// HasFuncOnSendError returns true if the function "WEBEXT.gcm.onSendError.addListener" exists.
func HasFuncOnSendError() bool {
	return js.True == bindings.HasFuncOnSendError()
}

// FuncOnSendError returns the function "WEBEXT.gcm.onSendError.addListener".
func FuncOnSendError() (fn js.Func[func(callback js.Func[func(err *OnSendErrorArgError)])]) {
	bindings.FuncOnSendError(
		js.Pointer(&fn),
	)
	return
}

// OnSendError calls the function "WEBEXT.gcm.onSendError.addListener" directly.
func OnSendError(callback js.Func[func(err *OnSendErrorArgError)]) (ret js.Void) {
	bindings.CallOnSendError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSendError calls the function "WEBEXT.gcm.onSendError.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSendError(callback js.Func[func(err *OnSendErrorArgError)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSendError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSendError returns true if the function "WEBEXT.gcm.onSendError.removeListener" exists.
func HasFuncOffSendError() bool {
	return js.True == bindings.HasFuncOffSendError()
}

// FuncOffSendError returns the function "WEBEXT.gcm.onSendError.removeListener".
func FuncOffSendError() (fn js.Func[func(callback js.Func[func(err *OnSendErrorArgError)])]) {
	bindings.FuncOffSendError(
		js.Pointer(&fn),
	)
	return
}

// OffSendError calls the function "WEBEXT.gcm.onSendError.removeListener" directly.
func OffSendError(callback js.Func[func(err *OnSendErrorArgError)]) (ret js.Void) {
	bindings.CallOffSendError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSendError calls the function "WEBEXT.gcm.onSendError.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSendError(callback js.Func[func(err *OnSendErrorArgError)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSendError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSendError returns true if the function "WEBEXT.gcm.onSendError.hasListener" exists.
func HasFuncHasOnSendError() bool {
	return js.True == bindings.HasFuncHasOnSendError()
}

// FuncHasOnSendError returns the function "WEBEXT.gcm.onSendError.hasListener".
func FuncHasOnSendError() (fn js.Func[func(callback js.Func[func(err *OnSendErrorArgError)]) bool]) {
	bindings.FuncHasOnSendError(
		js.Pointer(&fn),
	)
	return
}

// HasOnSendError calls the function "WEBEXT.gcm.onSendError.hasListener" directly.
func HasOnSendError(callback js.Func[func(err *OnSendErrorArgError)]) (ret bool) {
	bindings.CallHasOnSendError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSendError calls the function "WEBEXT.gcm.onSendError.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSendError(callback js.Func[func(err *OnSendErrorArgError)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSendError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncRegister returns true if the function "WEBEXT.gcm.register" exists.
func HasFuncRegister() bool {
	return js.True == bindings.HasFuncRegister()
}

// FuncRegister returns the function "WEBEXT.gcm.register".
func FuncRegister() (fn js.Func[func(senderIds js.Array[js.String]) js.Promise[js.String]]) {
	bindings.FuncRegister(
		js.Pointer(&fn),
	)
	return
}

// Register calls the function "WEBEXT.gcm.register" directly.
func Register(senderIds js.Array[js.String]) (ret js.Promise[js.String]) {
	bindings.CallRegister(
		js.Pointer(&ret),
		senderIds.Ref(),
	)

	return
}

// TryRegister calls the function "WEBEXT.gcm.register"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRegister(senderIds js.Array[js.String]) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRegister(
		js.Pointer(&ret), js.Pointer(&exception),
		senderIds.Ref(),
	)

	return
}

// HasFuncSend returns true if the function "WEBEXT.gcm.send" exists.
func HasFuncSend() bool {
	return js.True == bindings.HasFuncSend()
}

// FuncSend returns the function "WEBEXT.gcm.send".
func FuncSend() (fn js.Func[func(message SendArgMessage) js.Promise[js.String]]) {
	bindings.FuncSend(
		js.Pointer(&fn),
	)
	return
}

// Send calls the function "WEBEXT.gcm.send" directly.
func Send(message SendArgMessage) (ret js.Promise[js.String]) {
	bindings.CallSend(
		js.Pointer(&ret),
		js.Pointer(&message),
	)

	return
}

// TrySend calls the function "WEBEXT.gcm.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySend(message SendArgMessage) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySend(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&message),
	)

	return
}

// HasFuncUnregister returns true if the function "WEBEXT.gcm.unregister" exists.
func HasFuncUnregister() bool {
	return js.True == bindings.HasFuncUnregister()
}

// FuncUnregister returns the function "WEBEXT.gcm.unregister".
func FuncUnregister() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncUnregister(
		js.Pointer(&fn),
	)
	return
}

// Unregister calls the function "WEBEXT.gcm.unregister" directly.
func Unregister() (ret js.Promise[js.Void]) {
	bindings.CallUnregister(
		js.Pointer(&ret),
	)

	return
}

// TryUnregister calls the function "WEBEXT.gcm.unregister"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUnregister() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUnregister(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}
