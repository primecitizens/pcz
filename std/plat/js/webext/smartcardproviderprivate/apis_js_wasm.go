// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package smartcardproviderprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/smartcardproviderprivate/bindings"
)

type ConnectionState uint32

const (
	_ ConnectionState = iota

	ConnectionState_ABSENT
	ConnectionState_PRESENT
	ConnectionState_SWALLOWED
	ConnectionState_POWERED
	ConnectionState_NEGOTIABLE
	ConnectionState_SPECIFIC
)

func (ConnectionState) FromRef(str js.Ref) ConnectionState {
	return ConnectionState(bindings.ConstOfConnectionState(str))
}

func (x ConnectionState) String() (string, bool) {
	switch x {
	case ConnectionState_ABSENT:
		return "ABSENT", true
	case ConnectionState_PRESENT:
		return "PRESENT", true
	case ConnectionState_SWALLOWED:
		return "SWALLOWED", true
	case ConnectionState_POWERED:
		return "POWERED", true
	case ConnectionState_NEGOTIABLE:
		return "NEGOTIABLE", true
	case ConnectionState_SPECIFIC:
		return "SPECIFIC", true
	default:
		return "", false
	}
}

type Disposition uint32

const (
	_ Disposition = iota

	Disposition_LEAVE_CARD
	Disposition_RESET_CARD
	Disposition_UNPOWER_CARD
	Disposition_EJECT_CARD
)

func (Disposition) FromRef(str js.Ref) Disposition {
	return Disposition(bindings.ConstOfDisposition(str))
}

func (x Disposition) String() (string, bool) {
	switch x {
	case Disposition_LEAVE_CARD:
		return "LEAVE_CARD", true
	case Disposition_RESET_CARD:
		return "RESET_CARD", true
	case Disposition_UNPOWER_CARD:
		return "UNPOWER_CARD", true
	case Disposition_EJECT_CARD:
		return "EJECT_CARD", true
	default:
		return "", false
	}
}

type Protocol uint32

const (
	_ Protocol = iota

	Protocol_UNDEFINED
	Protocol_T0
	Protocol_T1
	Protocol_RAW
)

func (Protocol) FromRef(str js.Ref) Protocol {
	return Protocol(bindings.ConstOfProtocol(str))
}

func (x Protocol) String() (string, bool) {
	switch x {
	case Protocol_UNDEFINED:
		return "UNDEFINED", true
	case Protocol_T0:
		return "T0", true
	case Protocol_T1:
		return "T1", true
	case Protocol_RAW:
		return "RAW", true
	default:
		return "", false
	}
}

type Protocols struct {
	// T0 is "Protocols.t0"
	//
	// Optional
	//
	// NOTE: FFI_USE_T0 MUST be set to true to make this field effective.
	T0 bool
	// T1 is "Protocols.t1"
	//
	// Optional
	//
	// NOTE: FFI_USE_T1 MUST be set to true to make this field effective.
	T1 bool
	// Raw is "Protocols.raw"
	//
	// Optional
	//
	// NOTE: FFI_USE_Raw MUST be set to true to make this field effective.
	Raw bool

	FFI_USE_T0  bool // for T0.
	FFI_USE_T1  bool // for T1.
	FFI_USE_Raw bool // for Raw.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Protocols with all fields set.
func (p Protocols) FromRef(ref js.Ref) Protocols {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Protocols in the application heap.
func (p Protocols) New() js.Ref {
	return bindings.ProtocolsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Protocols) UpdateFrom(ref js.Ref) {
	bindings.ProtocolsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Protocols) Update(ref js.Ref) {
	bindings.ProtocolsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Protocols) FreeMembers(recursive bool) {
}

type ReaderStateFlags struct {
	// Unaware is "ReaderStateFlags.unaware"
	//
	// Optional
	//
	// NOTE: FFI_USE_Unaware MUST be set to true to make this field effective.
	Unaware bool
	// Ignore is "ReaderStateFlags.ignore"
	//
	// Optional
	//
	// NOTE: FFI_USE_Ignore MUST be set to true to make this field effective.
	Ignore bool
	// Changed is "ReaderStateFlags.changed"
	//
	// Optional
	//
	// NOTE: FFI_USE_Changed MUST be set to true to make this field effective.
	Changed bool
	// Unknown is "ReaderStateFlags.unknown"
	//
	// Optional
	//
	// NOTE: FFI_USE_Unknown MUST be set to true to make this field effective.
	Unknown bool
	// Unavailable is "ReaderStateFlags.unavailable"
	//
	// Optional
	//
	// NOTE: FFI_USE_Unavailable MUST be set to true to make this field effective.
	Unavailable bool
	// Empty is "ReaderStateFlags.empty"
	//
	// Optional
	//
	// NOTE: FFI_USE_Empty MUST be set to true to make this field effective.
	Empty bool
	// Present is "ReaderStateFlags.present"
	//
	// Optional
	//
	// NOTE: FFI_USE_Present MUST be set to true to make this field effective.
	Present bool
	// Exclusive is "ReaderStateFlags.exclusive"
	//
	// Optional
	//
	// NOTE: FFI_USE_Exclusive MUST be set to true to make this field effective.
	Exclusive bool
	// Inuse is "ReaderStateFlags.inuse"
	//
	// Optional
	//
	// NOTE: FFI_USE_Inuse MUST be set to true to make this field effective.
	Inuse bool
	// Mute is "ReaderStateFlags.mute"
	//
	// Optional
	//
	// NOTE: FFI_USE_Mute MUST be set to true to make this field effective.
	Mute bool
	// Unpowered is "ReaderStateFlags.unpowered"
	//
	// Optional
	//
	// NOTE: FFI_USE_Unpowered MUST be set to true to make this field effective.
	Unpowered bool

	FFI_USE_Unaware     bool // for Unaware.
	FFI_USE_Ignore      bool // for Ignore.
	FFI_USE_Changed     bool // for Changed.
	FFI_USE_Unknown     bool // for Unknown.
	FFI_USE_Unavailable bool // for Unavailable.
	FFI_USE_Empty       bool // for Empty.
	FFI_USE_Present     bool // for Present.
	FFI_USE_Exclusive   bool // for Exclusive.
	FFI_USE_Inuse       bool // for Inuse.
	FFI_USE_Mute        bool // for Mute.
	FFI_USE_Unpowered   bool // for Unpowered.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReaderStateFlags with all fields set.
func (p ReaderStateFlags) FromRef(ref js.Ref) ReaderStateFlags {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReaderStateFlags in the application heap.
func (p ReaderStateFlags) New() js.Ref {
	return bindings.ReaderStateFlagsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ReaderStateFlags) UpdateFrom(ref js.Ref) {
	bindings.ReaderStateFlagsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ReaderStateFlags) Update(ref js.Ref) {
	bindings.ReaderStateFlagsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ReaderStateFlags) FreeMembers(recursive bool) {
}

type ReaderStateIn struct {
	// Reader is "ReaderStateIn.reader"
	//
	// Optional
	Reader js.String
	// CurrentState is "ReaderStateIn.currentState"
	//
	// Optional
	//
	// NOTE: CurrentState.FFI_USE MUST be set to true to get CurrentState used.
	CurrentState ReaderStateFlags
	// CurrentCount is "ReaderStateIn.currentCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_CurrentCount MUST be set to true to make this field effective.
	CurrentCount int32

	FFI_USE_CurrentCount bool // for CurrentCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReaderStateIn with all fields set.
func (p ReaderStateIn) FromRef(ref js.Ref) ReaderStateIn {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReaderStateIn in the application heap.
func (p ReaderStateIn) New() js.Ref {
	return bindings.ReaderStateInJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ReaderStateIn) UpdateFrom(ref js.Ref) {
	bindings.ReaderStateInJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ReaderStateIn) Update(ref js.Ref) {
	bindings.ReaderStateInJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ReaderStateIn) FreeMembers(recursive bool) {
	js.Free(
		p.Reader.Ref(),
	)
	p.Reader = p.Reader.FromRef(js.Undefined)
	if recursive {
		p.CurrentState.FreeMembers(true)
	}
}

type ReaderStateOut struct {
	// Reader is "ReaderStateOut.reader"
	//
	// Optional
	Reader js.String
	// EventState is "ReaderStateOut.eventState"
	//
	// Optional
	//
	// NOTE: EventState.FFI_USE MUST be set to true to get EventState used.
	EventState ReaderStateFlags
	// EventCount is "ReaderStateOut.eventCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_EventCount MUST be set to true to make this field effective.
	EventCount int32
	// Atr is "ReaderStateOut.atr"
	//
	// Optional
	Atr js.ArrayBuffer

	FFI_USE_EventCount bool // for EventCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReaderStateOut with all fields set.
func (p ReaderStateOut) FromRef(ref js.Ref) ReaderStateOut {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReaderStateOut in the application heap.
func (p ReaderStateOut) New() js.Ref {
	return bindings.ReaderStateOutJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ReaderStateOut) UpdateFrom(ref js.Ref) {
	bindings.ReaderStateOutJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ReaderStateOut) Update(ref js.Ref) {
	bindings.ReaderStateOutJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ReaderStateOut) FreeMembers(recursive bool) {
	js.Free(
		p.Reader.Ref(),
		p.Atr.Ref(),
	)
	p.Reader = p.Reader.FromRef(js.Undefined)
	p.Atr = p.Atr.FromRef(js.Undefined)
	if recursive {
		p.EventState.FreeMembers(true)
	}
}

type ResultCode uint32

const (
	_ ResultCode = iota

	ResultCode_SUCCESS
	ResultCode_REMOVED_CARD
	ResultCode_RESET_CARD
	ResultCode_UNPOWERED_CARD
	ResultCode_UNRESPONSIVE_CARD
	ResultCode_UNSUPPORTED_CARD
	ResultCode_READER_UNAVAILABLE
	ResultCode_SHARING_VIOLATION
	ResultCode_NOT_TRANSACTED
	ResultCode_NO_SMARTCARD
	ResultCode_PROTO_MISMATCH
	ResultCode_SYSTEM_CANCELLED
	ResultCode_NOT_READY
	ResultCode_CANCELLED
	ResultCode_INSUFFICIENT_BUFFER
	ResultCode_INVALID_HANDLE
	ResultCode_INVALID_PARAMETER
	ResultCode_INVALID_VALUE
	ResultCode_NO_MEMORY
	ResultCode_TIMEOUT
	ResultCode_UNKNOWN_READER
	ResultCode_UNSUPPORTED_FEATURE
	ResultCode_NO_READERS_AVAILABLE
	ResultCode_SERVICE_STOPPED
	ResultCode_NO_SERVICE
	ResultCode_COMM_ERROR
	ResultCode_INTERNAL_ERROR
	ResultCode_UNKNOWN_ERROR
	ResultCode_SERVER_TOO_BUSY
	ResultCode_UNEXPECTED
	ResultCode_SHUTDOWN
	ResultCode_UNKNOWN
)

func (ResultCode) FromRef(str js.Ref) ResultCode {
	return ResultCode(bindings.ConstOfResultCode(str))
}

func (x ResultCode) String() (string, bool) {
	switch x {
	case ResultCode_SUCCESS:
		return "SUCCESS", true
	case ResultCode_REMOVED_CARD:
		return "REMOVED_CARD", true
	case ResultCode_RESET_CARD:
		return "RESET_CARD", true
	case ResultCode_UNPOWERED_CARD:
		return "UNPOWERED_CARD", true
	case ResultCode_UNRESPONSIVE_CARD:
		return "UNRESPONSIVE_CARD", true
	case ResultCode_UNSUPPORTED_CARD:
		return "UNSUPPORTED_CARD", true
	case ResultCode_READER_UNAVAILABLE:
		return "READER_UNAVAILABLE", true
	case ResultCode_SHARING_VIOLATION:
		return "SHARING_VIOLATION", true
	case ResultCode_NOT_TRANSACTED:
		return "NOT_TRANSACTED", true
	case ResultCode_NO_SMARTCARD:
		return "NO_SMARTCARD", true
	case ResultCode_PROTO_MISMATCH:
		return "PROTO_MISMATCH", true
	case ResultCode_SYSTEM_CANCELLED:
		return "SYSTEM_CANCELLED", true
	case ResultCode_NOT_READY:
		return "NOT_READY", true
	case ResultCode_CANCELLED:
		return "CANCELLED", true
	case ResultCode_INSUFFICIENT_BUFFER:
		return "INSUFFICIENT_BUFFER", true
	case ResultCode_INVALID_HANDLE:
		return "INVALID_HANDLE", true
	case ResultCode_INVALID_PARAMETER:
		return "INVALID_PARAMETER", true
	case ResultCode_INVALID_VALUE:
		return "INVALID_VALUE", true
	case ResultCode_NO_MEMORY:
		return "NO_MEMORY", true
	case ResultCode_TIMEOUT:
		return "TIMEOUT", true
	case ResultCode_UNKNOWN_READER:
		return "UNKNOWN_READER", true
	case ResultCode_UNSUPPORTED_FEATURE:
		return "UNSUPPORTED_FEATURE", true
	case ResultCode_NO_READERS_AVAILABLE:
		return "NO_READERS_AVAILABLE", true
	case ResultCode_SERVICE_STOPPED:
		return "SERVICE_STOPPED", true
	case ResultCode_NO_SERVICE:
		return "NO_SERVICE", true
	case ResultCode_COMM_ERROR:
		return "COMM_ERROR", true
	case ResultCode_INTERNAL_ERROR:
		return "INTERNAL_ERROR", true
	case ResultCode_UNKNOWN_ERROR:
		return "UNKNOWN_ERROR", true
	case ResultCode_SERVER_TOO_BUSY:
		return "SERVER_TOO_BUSY", true
	case ResultCode_UNEXPECTED:
		return "UNEXPECTED", true
	case ResultCode_SHUTDOWN:
		return "SHUTDOWN", true
	case ResultCode_UNKNOWN:
		return "UNKNOWN", true
	default:
		return "", false
	}
}

type ShareMode uint32

const (
	_ ShareMode = iota

	ShareMode_SHARED
	ShareMode_EXCLUSIVE
	ShareMode_DIRECT
)

func (ShareMode) FromRef(str js.Ref) ShareMode {
	return ShareMode(bindings.ConstOfShareMode(str))
}

func (x ShareMode) String() (string, bool) {
	switch x {
	case ShareMode_SHARED:
		return "SHARED", true
	case ShareMode_EXCLUSIVE:
		return "EXCLUSIVE", true
	case ShareMode_DIRECT:
		return "DIRECT", true
	default:
		return "", false
	}
}

type Timeout struct {
	// Milliseconds is "Timeout.milliseconds"
	//
	// Optional
	//
	// NOTE: FFI_USE_Milliseconds MUST be set to true to make this field effective.
	Milliseconds int32

	FFI_USE_Milliseconds bool // for Milliseconds.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Timeout with all fields set.
func (p Timeout) FromRef(ref js.Ref) Timeout {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Timeout in the application heap.
func (p Timeout) New() js.Ref {
	return bindings.TimeoutJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Timeout) UpdateFrom(ref js.Ref) {
	bindings.TimeoutJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Timeout) Update(ref js.Ref) {
	bindings.TimeoutJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Timeout) FreeMembers(recursive bool) {
}

type OnBeginTransactionRequestedEventCallbackFunc func(this js.Ref, requestId int32, scardHandle int32) js.Ref

func (fn OnBeginTransactionRequestedEventCallbackFunc) Register() js.Func[func(requestId int32, scardHandle int32)] {
	return js.RegisterCallback[func(requestId int32, scardHandle int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnBeginTransactionRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnBeginTransactionRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, requestId int32, scardHandle int32) js.Ref
	Arg T
}

func (cb *OnBeginTransactionRequestedEventCallback[T]) Register() js.Func[func(requestId int32, scardHandle int32)] {
	return js.RegisterCallback[func(requestId int32, scardHandle int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnBeginTransactionRequestedEventCallback[T]) DispatchCallback(
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnBeginTransactionRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onBeginTransactionRequested.addListener" exists.
func HasFuncOnBeginTransactionRequested() bool {
	return js.True == bindings.HasFuncOnBeginTransactionRequested()
}

// FuncOnBeginTransactionRequested returns the function "WEBEXT.smartCardProviderPrivate.onBeginTransactionRequested.addListener".
func FuncOnBeginTransactionRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32)])]) {
	bindings.FuncOnBeginTransactionRequested(
		js.Pointer(&fn),
	)
	return
}

// OnBeginTransactionRequested calls the function "WEBEXT.smartCardProviderPrivate.onBeginTransactionRequested.addListener" directly.
func OnBeginTransactionRequested(callback js.Func[func(requestId int32, scardHandle int32)]) (ret js.Void) {
	bindings.CallOnBeginTransactionRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnBeginTransactionRequested calls the function "WEBEXT.smartCardProviderPrivate.onBeginTransactionRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnBeginTransactionRequested(callback js.Func[func(requestId int32, scardHandle int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnBeginTransactionRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffBeginTransactionRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onBeginTransactionRequested.removeListener" exists.
func HasFuncOffBeginTransactionRequested() bool {
	return js.True == bindings.HasFuncOffBeginTransactionRequested()
}

// FuncOffBeginTransactionRequested returns the function "WEBEXT.smartCardProviderPrivate.onBeginTransactionRequested.removeListener".
func FuncOffBeginTransactionRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32)])]) {
	bindings.FuncOffBeginTransactionRequested(
		js.Pointer(&fn),
	)
	return
}

// OffBeginTransactionRequested calls the function "WEBEXT.smartCardProviderPrivate.onBeginTransactionRequested.removeListener" directly.
func OffBeginTransactionRequested(callback js.Func[func(requestId int32, scardHandle int32)]) (ret js.Void) {
	bindings.CallOffBeginTransactionRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffBeginTransactionRequested calls the function "WEBEXT.smartCardProviderPrivate.onBeginTransactionRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffBeginTransactionRequested(callback js.Func[func(requestId int32, scardHandle int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffBeginTransactionRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnBeginTransactionRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onBeginTransactionRequested.hasListener" exists.
func HasFuncHasOnBeginTransactionRequested() bool {
	return js.True == bindings.HasFuncHasOnBeginTransactionRequested()
}

// FuncHasOnBeginTransactionRequested returns the function "WEBEXT.smartCardProviderPrivate.onBeginTransactionRequested.hasListener".
func FuncHasOnBeginTransactionRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32)]) bool]) {
	bindings.FuncHasOnBeginTransactionRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnBeginTransactionRequested calls the function "WEBEXT.smartCardProviderPrivate.onBeginTransactionRequested.hasListener" directly.
func HasOnBeginTransactionRequested(callback js.Func[func(requestId int32, scardHandle int32)]) (ret bool) {
	bindings.CallHasOnBeginTransactionRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnBeginTransactionRequested calls the function "WEBEXT.smartCardProviderPrivate.onBeginTransactionRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnBeginTransactionRequested(callback js.Func[func(requestId int32, scardHandle int32)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnBeginTransactionRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnCancelRequestedEventCallbackFunc func(this js.Ref, requestId int32, scardContext int32) js.Ref

func (fn OnCancelRequestedEventCallbackFunc) Register() js.Func[func(requestId int32, scardContext int32)] {
	return js.RegisterCallback[func(requestId int32, scardContext int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCancelRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnCancelRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, requestId int32, scardContext int32) js.Ref
	Arg T
}

func (cb *OnCancelRequestedEventCallback[T]) Register() js.Func[func(requestId int32, scardContext int32)] {
	return js.RegisterCallback[func(requestId int32, scardContext int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCancelRequestedEventCallback[T]) DispatchCallback(
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnCancelRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onCancelRequested.addListener" exists.
func HasFuncOnCancelRequested() bool {
	return js.True == bindings.HasFuncOnCancelRequested()
}

// FuncOnCancelRequested returns the function "WEBEXT.smartCardProviderPrivate.onCancelRequested.addListener".
func FuncOnCancelRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardContext int32)])]) {
	bindings.FuncOnCancelRequested(
		js.Pointer(&fn),
	)
	return
}

// OnCancelRequested calls the function "WEBEXT.smartCardProviderPrivate.onCancelRequested.addListener" directly.
func OnCancelRequested(callback js.Func[func(requestId int32, scardContext int32)]) (ret js.Void) {
	bindings.CallOnCancelRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCancelRequested calls the function "WEBEXT.smartCardProviderPrivate.onCancelRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCancelRequested(callback js.Func[func(requestId int32, scardContext int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCancelRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCancelRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onCancelRequested.removeListener" exists.
func HasFuncOffCancelRequested() bool {
	return js.True == bindings.HasFuncOffCancelRequested()
}

// FuncOffCancelRequested returns the function "WEBEXT.smartCardProviderPrivate.onCancelRequested.removeListener".
func FuncOffCancelRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardContext int32)])]) {
	bindings.FuncOffCancelRequested(
		js.Pointer(&fn),
	)
	return
}

// OffCancelRequested calls the function "WEBEXT.smartCardProviderPrivate.onCancelRequested.removeListener" directly.
func OffCancelRequested(callback js.Func[func(requestId int32, scardContext int32)]) (ret js.Void) {
	bindings.CallOffCancelRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCancelRequested calls the function "WEBEXT.smartCardProviderPrivate.onCancelRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCancelRequested(callback js.Func[func(requestId int32, scardContext int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCancelRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCancelRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onCancelRequested.hasListener" exists.
func HasFuncHasOnCancelRequested() bool {
	return js.True == bindings.HasFuncHasOnCancelRequested()
}

// FuncHasOnCancelRequested returns the function "WEBEXT.smartCardProviderPrivate.onCancelRequested.hasListener".
func FuncHasOnCancelRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardContext int32)]) bool]) {
	bindings.FuncHasOnCancelRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnCancelRequested calls the function "WEBEXT.smartCardProviderPrivate.onCancelRequested.hasListener" directly.
func HasOnCancelRequested(callback js.Func[func(requestId int32, scardContext int32)]) (ret bool) {
	bindings.CallHasOnCancelRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCancelRequested calls the function "WEBEXT.smartCardProviderPrivate.onCancelRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCancelRequested(callback js.Func[func(requestId int32, scardContext int32)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCancelRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnConnectRequestedEventCallbackFunc func(this js.Ref, requestId int32, scardContext int32, reader js.String, shareMode ShareMode, preferredProtocols *Protocols) js.Ref

func (fn OnConnectRequestedEventCallbackFunc) Register() js.Func[func(requestId int32, scardContext int32, reader js.String, shareMode ShareMode, preferredProtocols *Protocols)] {
	return js.RegisterCallback[func(requestId int32, scardContext int32, reader js.String, shareMode ShareMode, preferredProtocols *Protocols)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnConnectRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 5+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg4 Protocols
	arg4.UpdateFrom(args[4+1])
	defer arg4.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
		js.String{}.FromRef(args[2+1]),
		ShareMode(0).FromRef(args[3+1]),
		mark.NoEscape(&arg4),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnConnectRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, requestId int32, scardContext int32, reader js.String, shareMode ShareMode, preferredProtocols *Protocols) js.Ref
	Arg T
}

func (cb *OnConnectRequestedEventCallback[T]) Register() js.Func[func(requestId int32, scardContext int32, reader js.String, shareMode ShareMode, preferredProtocols *Protocols)] {
	return js.RegisterCallback[func(requestId int32, scardContext int32, reader js.String, shareMode ShareMode, preferredProtocols *Protocols)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnConnectRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 5+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg4 Protocols
	arg4.UpdateFrom(args[4+1])
	defer arg4.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
		js.String{}.FromRef(args[2+1]),
		ShareMode(0).FromRef(args[3+1]),
		mark.NoEscape(&arg4),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnConnectRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onConnectRequested.addListener" exists.
func HasFuncOnConnectRequested() bool {
	return js.True == bindings.HasFuncOnConnectRequested()
}

// FuncOnConnectRequested returns the function "WEBEXT.smartCardProviderPrivate.onConnectRequested.addListener".
func FuncOnConnectRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardContext int32, reader js.String, shareMode ShareMode, preferredProtocols *Protocols)])]) {
	bindings.FuncOnConnectRequested(
		js.Pointer(&fn),
	)
	return
}

// OnConnectRequested calls the function "WEBEXT.smartCardProviderPrivate.onConnectRequested.addListener" directly.
func OnConnectRequested(callback js.Func[func(requestId int32, scardContext int32, reader js.String, shareMode ShareMode, preferredProtocols *Protocols)]) (ret js.Void) {
	bindings.CallOnConnectRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnConnectRequested calls the function "WEBEXT.smartCardProviderPrivate.onConnectRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnConnectRequested(callback js.Func[func(requestId int32, scardContext int32, reader js.String, shareMode ShareMode, preferredProtocols *Protocols)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnConnectRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffConnectRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onConnectRequested.removeListener" exists.
func HasFuncOffConnectRequested() bool {
	return js.True == bindings.HasFuncOffConnectRequested()
}

// FuncOffConnectRequested returns the function "WEBEXT.smartCardProviderPrivate.onConnectRequested.removeListener".
func FuncOffConnectRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardContext int32, reader js.String, shareMode ShareMode, preferredProtocols *Protocols)])]) {
	bindings.FuncOffConnectRequested(
		js.Pointer(&fn),
	)
	return
}

// OffConnectRequested calls the function "WEBEXT.smartCardProviderPrivate.onConnectRequested.removeListener" directly.
func OffConnectRequested(callback js.Func[func(requestId int32, scardContext int32, reader js.String, shareMode ShareMode, preferredProtocols *Protocols)]) (ret js.Void) {
	bindings.CallOffConnectRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffConnectRequested calls the function "WEBEXT.smartCardProviderPrivate.onConnectRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffConnectRequested(callback js.Func[func(requestId int32, scardContext int32, reader js.String, shareMode ShareMode, preferredProtocols *Protocols)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffConnectRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnConnectRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onConnectRequested.hasListener" exists.
func HasFuncHasOnConnectRequested() bool {
	return js.True == bindings.HasFuncHasOnConnectRequested()
}

// FuncHasOnConnectRequested returns the function "WEBEXT.smartCardProviderPrivate.onConnectRequested.hasListener".
func FuncHasOnConnectRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardContext int32, reader js.String, shareMode ShareMode, preferredProtocols *Protocols)]) bool]) {
	bindings.FuncHasOnConnectRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnConnectRequested calls the function "WEBEXT.smartCardProviderPrivate.onConnectRequested.hasListener" directly.
func HasOnConnectRequested(callback js.Func[func(requestId int32, scardContext int32, reader js.String, shareMode ShareMode, preferredProtocols *Protocols)]) (ret bool) {
	bindings.CallHasOnConnectRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnConnectRequested calls the function "WEBEXT.smartCardProviderPrivate.onConnectRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnConnectRequested(callback js.Func[func(requestId int32, scardContext int32, reader js.String, shareMode ShareMode, preferredProtocols *Protocols)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnConnectRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnControlRequestedEventCallbackFunc func(this js.Ref, requestId int32, scardHandle int32, controlCode int32, data js.ArrayBuffer) js.Ref

func (fn OnControlRequestedEventCallbackFunc) Register() js.Func[func(requestId int32, scardHandle int32, controlCode int32, data js.ArrayBuffer)] {
	return js.RegisterCallback[func(requestId int32, scardHandle int32, controlCode int32, data js.ArrayBuffer)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnControlRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 4+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
		js.Number[int32]{}.FromRef(args[2+1]).Get(),
		js.ArrayBuffer{}.FromRef(args[3+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnControlRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, requestId int32, scardHandle int32, controlCode int32, data js.ArrayBuffer) js.Ref
	Arg T
}

func (cb *OnControlRequestedEventCallback[T]) Register() js.Func[func(requestId int32, scardHandle int32, controlCode int32, data js.ArrayBuffer)] {
	return js.RegisterCallback[func(requestId int32, scardHandle int32, controlCode int32, data js.ArrayBuffer)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnControlRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 4+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
		js.Number[int32]{}.FromRef(args[2+1]).Get(),
		js.ArrayBuffer{}.FromRef(args[3+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnControlRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onControlRequested.addListener" exists.
func HasFuncOnControlRequested() bool {
	return js.True == bindings.HasFuncOnControlRequested()
}

// FuncOnControlRequested returns the function "WEBEXT.smartCardProviderPrivate.onControlRequested.addListener".
func FuncOnControlRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32, controlCode int32, data js.ArrayBuffer)])]) {
	bindings.FuncOnControlRequested(
		js.Pointer(&fn),
	)
	return
}

// OnControlRequested calls the function "WEBEXT.smartCardProviderPrivate.onControlRequested.addListener" directly.
func OnControlRequested(callback js.Func[func(requestId int32, scardHandle int32, controlCode int32, data js.ArrayBuffer)]) (ret js.Void) {
	bindings.CallOnControlRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnControlRequested calls the function "WEBEXT.smartCardProviderPrivate.onControlRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnControlRequested(callback js.Func[func(requestId int32, scardHandle int32, controlCode int32, data js.ArrayBuffer)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnControlRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffControlRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onControlRequested.removeListener" exists.
func HasFuncOffControlRequested() bool {
	return js.True == bindings.HasFuncOffControlRequested()
}

// FuncOffControlRequested returns the function "WEBEXT.smartCardProviderPrivate.onControlRequested.removeListener".
func FuncOffControlRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32, controlCode int32, data js.ArrayBuffer)])]) {
	bindings.FuncOffControlRequested(
		js.Pointer(&fn),
	)
	return
}

// OffControlRequested calls the function "WEBEXT.smartCardProviderPrivate.onControlRequested.removeListener" directly.
func OffControlRequested(callback js.Func[func(requestId int32, scardHandle int32, controlCode int32, data js.ArrayBuffer)]) (ret js.Void) {
	bindings.CallOffControlRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffControlRequested calls the function "WEBEXT.smartCardProviderPrivate.onControlRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffControlRequested(callback js.Func[func(requestId int32, scardHandle int32, controlCode int32, data js.ArrayBuffer)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffControlRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnControlRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onControlRequested.hasListener" exists.
func HasFuncHasOnControlRequested() bool {
	return js.True == bindings.HasFuncHasOnControlRequested()
}

// FuncHasOnControlRequested returns the function "WEBEXT.smartCardProviderPrivate.onControlRequested.hasListener".
func FuncHasOnControlRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32, controlCode int32, data js.ArrayBuffer)]) bool]) {
	bindings.FuncHasOnControlRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnControlRequested calls the function "WEBEXT.smartCardProviderPrivate.onControlRequested.hasListener" directly.
func HasOnControlRequested(callback js.Func[func(requestId int32, scardHandle int32, controlCode int32, data js.ArrayBuffer)]) (ret bool) {
	bindings.CallHasOnControlRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnControlRequested calls the function "WEBEXT.smartCardProviderPrivate.onControlRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnControlRequested(callback js.Func[func(requestId int32, scardHandle int32, controlCode int32, data js.ArrayBuffer)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnControlRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDisconnectRequestedEventCallbackFunc func(this js.Ref, requestId int32, scardHandle int32, disposition Disposition) js.Ref

func (fn OnDisconnectRequestedEventCallbackFunc) Register() js.Func[func(requestId int32, scardHandle int32, disposition Disposition)] {
	return js.RegisterCallback[func(requestId int32, scardHandle int32, disposition Disposition)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDisconnectRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
		Disposition(0).FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnDisconnectRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, requestId int32, scardHandle int32, disposition Disposition) js.Ref
	Arg T
}

func (cb *OnDisconnectRequestedEventCallback[T]) Register() js.Func[func(requestId int32, scardHandle int32, disposition Disposition)] {
	return js.RegisterCallback[func(requestId int32, scardHandle int32, disposition Disposition)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDisconnectRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
		Disposition(0).FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnDisconnectRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onDisconnectRequested.addListener" exists.
func HasFuncOnDisconnectRequested() bool {
	return js.True == bindings.HasFuncOnDisconnectRequested()
}

// FuncOnDisconnectRequested returns the function "WEBEXT.smartCardProviderPrivate.onDisconnectRequested.addListener".
func FuncOnDisconnectRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32, disposition Disposition)])]) {
	bindings.FuncOnDisconnectRequested(
		js.Pointer(&fn),
	)
	return
}

// OnDisconnectRequested calls the function "WEBEXT.smartCardProviderPrivate.onDisconnectRequested.addListener" directly.
func OnDisconnectRequested(callback js.Func[func(requestId int32, scardHandle int32, disposition Disposition)]) (ret js.Void) {
	bindings.CallOnDisconnectRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDisconnectRequested calls the function "WEBEXT.smartCardProviderPrivate.onDisconnectRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDisconnectRequested(callback js.Func[func(requestId int32, scardHandle int32, disposition Disposition)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDisconnectRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDisconnectRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onDisconnectRequested.removeListener" exists.
func HasFuncOffDisconnectRequested() bool {
	return js.True == bindings.HasFuncOffDisconnectRequested()
}

// FuncOffDisconnectRequested returns the function "WEBEXT.smartCardProviderPrivate.onDisconnectRequested.removeListener".
func FuncOffDisconnectRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32, disposition Disposition)])]) {
	bindings.FuncOffDisconnectRequested(
		js.Pointer(&fn),
	)
	return
}

// OffDisconnectRequested calls the function "WEBEXT.smartCardProviderPrivate.onDisconnectRequested.removeListener" directly.
func OffDisconnectRequested(callback js.Func[func(requestId int32, scardHandle int32, disposition Disposition)]) (ret js.Void) {
	bindings.CallOffDisconnectRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDisconnectRequested calls the function "WEBEXT.smartCardProviderPrivate.onDisconnectRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDisconnectRequested(callback js.Func[func(requestId int32, scardHandle int32, disposition Disposition)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDisconnectRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDisconnectRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onDisconnectRequested.hasListener" exists.
func HasFuncHasOnDisconnectRequested() bool {
	return js.True == bindings.HasFuncHasOnDisconnectRequested()
}

// FuncHasOnDisconnectRequested returns the function "WEBEXT.smartCardProviderPrivate.onDisconnectRequested.hasListener".
func FuncHasOnDisconnectRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32, disposition Disposition)]) bool]) {
	bindings.FuncHasOnDisconnectRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnDisconnectRequested calls the function "WEBEXT.smartCardProviderPrivate.onDisconnectRequested.hasListener" directly.
func HasOnDisconnectRequested(callback js.Func[func(requestId int32, scardHandle int32, disposition Disposition)]) (ret bool) {
	bindings.CallHasOnDisconnectRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDisconnectRequested calls the function "WEBEXT.smartCardProviderPrivate.onDisconnectRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDisconnectRequested(callback js.Func[func(requestId int32, scardHandle int32, disposition Disposition)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDisconnectRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnEndTransactionRequestedEventCallbackFunc func(this js.Ref, requestId int32, scardHandle int32, disposition Disposition) js.Ref

func (fn OnEndTransactionRequestedEventCallbackFunc) Register() js.Func[func(requestId int32, scardHandle int32, disposition Disposition)] {
	return js.RegisterCallback[func(requestId int32, scardHandle int32, disposition Disposition)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnEndTransactionRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
		Disposition(0).FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnEndTransactionRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, requestId int32, scardHandle int32, disposition Disposition) js.Ref
	Arg T
}

func (cb *OnEndTransactionRequestedEventCallback[T]) Register() js.Func[func(requestId int32, scardHandle int32, disposition Disposition)] {
	return js.RegisterCallback[func(requestId int32, scardHandle int32, disposition Disposition)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnEndTransactionRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
		Disposition(0).FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnEndTransactionRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onEndTransactionRequested.addListener" exists.
func HasFuncOnEndTransactionRequested() bool {
	return js.True == bindings.HasFuncOnEndTransactionRequested()
}

// FuncOnEndTransactionRequested returns the function "WEBEXT.smartCardProviderPrivate.onEndTransactionRequested.addListener".
func FuncOnEndTransactionRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32, disposition Disposition)])]) {
	bindings.FuncOnEndTransactionRequested(
		js.Pointer(&fn),
	)
	return
}

// OnEndTransactionRequested calls the function "WEBEXT.smartCardProviderPrivate.onEndTransactionRequested.addListener" directly.
func OnEndTransactionRequested(callback js.Func[func(requestId int32, scardHandle int32, disposition Disposition)]) (ret js.Void) {
	bindings.CallOnEndTransactionRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnEndTransactionRequested calls the function "WEBEXT.smartCardProviderPrivate.onEndTransactionRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnEndTransactionRequested(callback js.Func[func(requestId int32, scardHandle int32, disposition Disposition)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnEndTransactionRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffEndTransactionRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onEndTransactionRequested.removeListener" exists.
func HasFuncOffEndTransactionRequested() bool {
	return js.True == bindings.HasFuncOffEndTransactionRequested()
}

// FuncOffEndTransactionRequested returns the function "WEBEXT.smartCardProviderPrivate.onEndTransactionRequested.removeListener".
func FuncOffEndTransactionRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32, disposition Disposition)])]) {
	bindings.FuncOffEndTransactionRequested(
		js.Pointer(&fn),
	)
	return
}

// OffEndTransactionRequested calls the function "WEBEXT.smartCardProviderPrivate.onEndTransactionRequested.removeListener" directly.
func OffEndTransactionRequested(callback js.Func[func(requestId int32, scardHandle int32, disposition Disposition)]) (ret js.Void) {
	bindings.CallOffEndTransactionRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffEndTransactionRequested calls the function "WEBEXT.smartCardProviderPrivate.onEndTransactionRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffEndTransactionRequested(callback js.Func[func(requestId int32, scardHandle int32, disposition Disposition)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffEndTransactionRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnEndTransactionRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onEndTransactionRequested.hasListener" exists.
func HasFuncHasOnEndTransactionRequested() bool {
	return js.True == bindings.HasFuncHasOnEndTransactionRequested()
}

// FuncHasOnEndTransactionRequested returns the function "WEBEXT.smartCardProviderPrivate.onEndTransactionRequested.hasListener".
func FuncHasOnEndTransactionRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32, disposition Disposition)]) bool]) {
	bindings.FuncHasOnEndTransactionRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnEndTransactionRequested calls the function "WEBEXT.smartCardProviderPrivate.onEndTransactionRequested.hasListener" directly.
func HasOnEndTransactionRequested(callback js.Func[func(requestId int32, scardHandle int32, disposition Disposition)]) (ret bool) {
	bindings.CallHasOnEndTransactionRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnEndTransactionRequested calls the function "WEBEXT.smartCardProviderPrivate.onEndTransactionRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnEndTransactionRequested(callback js.Func[func(requestId int32, scardHandle int32, disposition Disposition)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnEndTransactionRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnEstablishContextRequestedEventCallbackFunc func(this js.Ref, requestId int32) js.Ref

func (fn OnEstablishContextRequestedEventCallbackFunc) Register() js.Func[func(requestId int32)] {
	return js.RegisterCallback[func(requestId int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnEstablishContextRequestedEventCallbackFunc) DispatchCallback(
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

type OnEstablishContextRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, requestId int32) js.Ref
	Arg T
}

func (cb *OnEstablishContextRequestedEventCallback[T]) Register() js.Func[func(requestId int32)] {
	return js.RegisterCallback[func(requestId int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnEstablishContextRequestedEventCallback[T]) DispatchCallback(
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

// HasFuncOnEstablishContextRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onEstablishContextRequested.addListener" exists.
func HasFuncOnEstablishContextRequested() bool {
	return js.True == bindings.HasFuncOnEstablishContextRequested()
}

// FuncOnEstablishContextRequested returns the function "WEBEXT.smartCardProviderPrivate.onEstablishContextRequested.addListener".
func FuncOnEstablishContextRequested() (fn js.Func[func(callback js.Func[func(requestId int32)])]) {
	bindings.FuncOnEstablishContextRequested(
		js.Pointer(&fn),
	)
	return
}

// OnEstablishContextRequested calls the function "WEBEXT.smartCardProviderPrivate.onEstablishContextRequested.addListener" directly.
func OnEstablishContextRequested(callback js.Func[func(requestId int32)]) (ret js.Void) {
	bindings.CallOnEstablishContextRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnEstablishContextRequested calls the function "WEBEXT.smartCardProviderPrivate.onEstablishContextRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnEstablishContextRequested(callback js.Func[func(requestId int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnEstablishContextRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffEstablishContextRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onEstablishContextRequested.removeListener" exists.
func HasFuncOffEstablishContextRequested() bool {
	return js.True == bindings.HasFuncOffEstablishContextRequested()
}

// FuncOffEstablishContextRequested returns the function "WEBEXT.smartCardProviderPrivate.onEstablishContextRequested.removeListener".
func FuncOffEstablishContextRequested() (fn js.Func[func(callback js.Func[func(requestId int32)])]) {
	bindings.FuncOffEstablishContextRequested(
		js.Pointer(&fn),
	)
	return
}

// OffEstablishContextRequested calls the function "WEBEXT.smartCardProviderPrivate.onEstablishContextRequested.removeListener" directly.
func OffEstablishContextRequested(callback js.Func[func(requestId int32)]) (ret js.Void) {
	bindings.CallOffEstablishContextRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffEstablishContextRequested calls the function "WEBEXT.smartCardProviderPrivate.onEstablishContextRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffEstablishContextRequested(callback js.Func[func(requestId int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffEstablishContextRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnEstablishContextRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onEstablishContextRequested.hasListener" exists.
func HasFuncHasOnEstablishContextRequested() bool {
	return js.True == bindings.HasFuncHasOnEstablishContextRequested()
}

// FuncHasOnEstablishContextRequested returns the function "WEBEXT.smartCardProviderPrivate.onEstablishContextRequested.hasListener".
func FuncHasOnEstablishContextRequested() (fn js.Func[func(callback js.Func[func(requestId int32)]) bool]) {
	bindings.FuncHasOnEstablishContextRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnEstablishContextRequested calls the function "WEBEXT.smartCardProviderPrivate.onEstablishContextRequested.hasListener" directly.
func HasOnEstablishContextRequested(callback js.Func[func(requestId int32)]) (ret bool) {
	bindings.CallHasOnEstablishContextRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnEstablishContextRequested calls the function "WEBEXT.smartCardProviderPrivate.onEstablishContextRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnEstablishContextRequested(callback js.Func[func(requestId int32)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnEstablishContextRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnGetAttribRequestedEventCallbackFunc func(this js.Ref, requestId int32, scardHandle int32, attribId int32) js.Ref

func (fn OnGetAttribRequestedEventCallbackFunc) Register() js.Func[func(requestId int32, scardHandle int32, attribId int32)] {
	return js.RegisterCallback[func(requestId int32, scardHandle int32, attribId int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnGetAttribRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
		js.Number[int32]{}.FromRef(args[2+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnGetAttribRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, requestId int32, scardHandle int32, attribId int32) js.Ref
	Arg T
}

func (cb *OnGetAttribRequestedEventCallback[T]) Register() js.Func[func(requestId int32, scardHandle int32, attribId int32)] {
	return js.RegisterCallback[func(requestId int32, scardHandle int32, attribId int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnGetAttribRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
		js.Number[int32]{}.FromRef(args[2+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnGetAttribRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onGetAttribRequested.addListener" exists.
func HasFuncOnGetAttribRequested() bool {
	return js.True == bindings.HasFuncOnGetAttribRequested()
}

// FuncOnGetAttribRequested returns the function "WEBEXT.smartCardProviderPrivate.onGetAttribRequested.addListener".
func FuncOnGetAttribRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32, attribId int32)])]) {
	bindings.FuncOnGetAttribRequested(
		js.Pointer(&fn),
	)
	return
}

// OnGetAttribRequested calls the function "WEBEXT.smartCardProviderPrivate.onGetAttribRequested.addListener" directly.
func OnGetAttribRequested(callback js.Func[func(requestId int32, scardHandle int32, attribId int32)]) (ret js.Void) {
	bindings.CallOnGetAttribRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnGetAttribRequested calls the function "WEBEXT.smartCardProviderPrivate.onGetAttribRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnGetAttribRequested(callback js.Func[func(requestId int32, scardHandle int32, attribId int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnGetAttribRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffGetAttribRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onGetAttribRequested.removeListener" exists.
func HasFuncOffGetAttribRequested() bool {
	return js.True == bindings.HasFuncOffGetAttribRequested()
}

// FuncOffGetAttribRequested returns the function "WEBEXT.smartCardProviderPrivate.onGetAttribRequested.removeListener".
func FuncOffGetAttribRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32, attribId int32)])]) {
	bindings.FuncOffGetAttribRequested(
		js.Pointer(&fn),
	)
	return
}

// OffGetAttribRequested calls the function "WEBEXT.smartCardProviderPrivate.onGetAttribRequested.removeListener" directly.
func OffGetAttribRequested(callback js.Func[func(requestId int32, scardHandle int32, attribId int32)]) (ret js.Void) {
	bindings.CallOffGetAttribRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffGetAttribRequested calls the function "WEBEXT.smartCardProviderPrivate.onGetAttribRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffGetAttribRequested(callback js.Func[func(requestId int32, scardHandle int32, attribId int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffGetAttribRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnGetAttribRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onGetAttribRequested.hasListener" exists.
func HasFuncHasOnGetAttribRequested() bool {
	return js.True == bindings.HasFuncHasOnGetAttribRequested()
}

// FuncHasOnGetAttribRequested returns the function "WEBEXT.smartCardProviderPrivate.onGetAttribRequested.hasListener".
func FuncHasOnGetAttribRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32, attribId int32)]) bool]) {
	bindings.FuncHasOnGetAttribRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnGetAttribRequested calls the function "WEBEXT.smartCardProviderPrivate.onGetAttribRequested.hasListener" directly.
func HasOnGetAttribRequested(callback js.Func[func(requestId int32, scardHandle int32, attribId int32)]) (ret bool) {
	bindings.CallHasOnGetAttribRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnGetAttribRequested calls the function "WEBEXT.smartCardProviderPrivate.onGetAttribRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnGetAttribRequested(callback js.Func[func(requestId int32, scardHandle int32, attribId int32)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnGetAttribRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnGetStatusChangeRequestedEventCallbackFunc func(this js.Ref, requestId int32, scardContext int32, timeout *Timeout, readerStates js.Array[ReaderStateIn]) js.Ref

func (fn OnGetStatusChangeRequestedEventCallbackFunc) Register() js.Func[func(requestId int32, scardContext int32, timeout *Timeout, readerStates js.Array[ReaderStateIn])] {
	return js.RegisterCallback[func(requestId int32, scardContext int32, timeout *Timeout, readerStates js.Array[ReaderStateIn])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnGetStatusChangeRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 4+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg2 Timeout
	arg2.UpdateFrom(args[2+1])
	defer arg2.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
		mark.NoEscape(&arg2),
		js.Array[ReaderStateIn]{}.FromRef(args[3+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnGetStatusChangeRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, requestId int32, scardContext int32, timeout *Timeout, readerStates js.Array[ReaderStateIn]) js.Ref
	Arg T
}

func (cb *OnGetStatusChangeRequestedEventCallback[T]) Register() js.Func[func(requestId int32, scardContext int32, timeout *Timeout, readerStates js.Array[ReaderStateIn])] {
	return js.RegisterCallback[func(requestId int32, scardContext int32, timeout *Timeout, readerStates js.Array[ReaderStateIn])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnGetStatusChangeRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 4+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg2 Timeout
	arg2.UpdateFrom(args[2+1])
	defer arg2.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
		mark.NoEscape(&arg2),
		js.Array[ReaderStateIn]{}.FromRef(args[3+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnGetStatusChangeRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onGetStatusChangeRequested.addListener" exists.
func HasFuncOnGetStatusChangeRequested() bool {
	return js.True == bindings.HasFuncOnGetStatusChangeRequested()
}

// FuncOnGetStatusChangeRequested returns the function "WEBEXT.smartCardProviderPrivate.onGetStatusChangeRequested.addListener".
func FuncOnGetStatusChangeRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardContext int32, timeout *Timeout, readerStates js.Array[ReaderStateIn])])]) {
	bindings.FuncOnGetStatusChangeRequested(
		js.Pointer(&fn),
	)
	return
}

// OnGetStatusChangeRequested calls the function "WEBEXT.smartCardProviderPrivate.onGetStatusChangeRequested.addListener" directly.
func OnGetStatusChangeRequested(callback js.Func[func(requestId int32, scardContext int32, timeout *Timeout, readerStates js.Array[ReaderStateIn])]) (ret js.Void) {
	bindings.CallOnGetStatusChangeRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnGetStatusChangeRequested calls the function "WEBEXT.smartCardProviderPrivate.onGetStatusChangeRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnGetStatusChangeRequested(callback js.Func[func(requestId int32, scardContext int32, timeout *Timeout, readerStates js.Array[ReaderStateIn])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnGetStatusChangeRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffGetStatusChangeRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onGetStatusChangeRequested.removeListener" exists.
func HasFuncOffGetStatusChangeRequested() bool {
	return js.True == bindings.HasFuncOffGetStatusChangeRequested()
}

// FuncOffGetStatusChangeRequested returns the function "WEBEXT.smartCardProviderPrivate.onGetStatusChangeRequested.removeListener".
func FuncOffGetStatusChangeRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardContext int32, timeout *Timeout, readerStates js.Array[ReaderStateIn])])]) {
	bindings.FuncOffGetStatusChangeRequested(
		js.Pointer(&fn),
	)
	return
}

// OffGetStatusChangeRequested calls the function "WEBEXT.smartCardProviderPrivate.onGetStatusChangeRequested.removeListener" directly.
func OffGetStatusChangeRequested(callback js.Func[func(requestId int32, scardContext int32, timeout *Timeout, readerStates js.Array[ReaderStateIn])]) (ret js.Void) {
	bindings.CallOffGetStatusChangeRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffGetStatusChangeRequested calls the function "WEBEXT.smartCardProviderPrivate.onGetStatusChangeRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffGetStatusChangeRequested(callback js.Func[func(requestId int32, scardContext int32, timeout *Timeout, readerStates js.Array[ReaderStateIn])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffGetStatusChangeRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnGetStatusChangeRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onGetStatusChangeRequested.hasListener" exists.
func HasFuncHasOnGetStatusChangeRequested() bool {
	return js.True == bindings.HasFuncHasOnGetStatusChangeRequested()
}

// FuncHasOnGetStatusChangeRequested returns the function "WEBEXT.smartCardProviderPrivate.onGetStatusChangeRequested.hasListener".
func FuncHasOnGetStatusChangeRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardContext int32, timeout *Timeout, readerStates js.Array[ReaderStateIn])]) bool]) {
	bindings.FuncHasOnGetStatusChangeRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnGetStatusChangeRequested calls the function "WEBEXT.smartCardProviderPrivate.onGetStatusChangeRequested.hasListener" directly.
func HasOnGetStatusChangeRequested(callback js.Func[func(requestId int32, scardContext int32, timeout *Timeout, readerStates js.Array[ReaderStateIn])]) (ret bool) {
	bindings.CallHasOnGetStatusChangeRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnGetStatusChangeRequested calls the function "WEBEXT.smartCardProviderPrivate.onGetStatusChangeRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnGetStatusChangeRequested(callback js.Func[func(requestId int32, scardContext int32, timeout *Timeout, readerStates js.Array[ReaderStateIn])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnGetStatusChangeRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnListReadersRequestedEventCallbackFunc func(this js.Ref, requestId int32, scardContext int32) js.Ref

func (fn OnListReadersRequestedEventCallbackFunc) Register() js.Func[func(requestId int32, scardContext int32)] {
	return js.RegisterCallback[func(requestId int32, scardContext int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnListReadersRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnListReadersRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, requestId int32, scardContext int32) js.Ref
	Arg T
}

func (cb *OnListReadersRequestedEventCallback[T]) Register() js.Func[func(requestId int32, scardContext int32)] {
	return js.RegisterCallback[func(requestId int32, scardContext int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnListReadersRequestedEventCallback[T]) DispatchCallback(
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnListReadersRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onListReadersRequested.addListener" exists.
func HasFuncOnListReadersRequested() bool {
	return js.True == bindings.HasFuncOnListReadersRequested()
}

// FuncOnListReadersRequested returns the function "WEBEXT.smartCardProviderPrivate.onListReadersRequested.addListener".
func FuncOnListReadersRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardContext int32)])]) {
	bindings.FuncOnListReadersRequested(
		js.Pointer(&fn),
	)
	return
}

// OnListReadersRequested calls the function "WEBEXT.smartCardProviderPrivate.onListReadersRequested.addListener" directly.
func OnListReadersRequested(callback js.Func[func(requestId int32, scardContext int32)]) (ret js.Void) {
	bindings.CallOnListReadersRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnListReadersRequested calls the function "WEBEXT.smartCardProviderPrivate.onListReadersRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnListReadersRequested(callback js.Func[func(requestId int32, scardContext int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnListReadersRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffListReadersRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onListReadersRequested.removeListener" exists.
func HasFuncOffListReadersRequested() bool {
	return js.True == bindings.HasFuncOffListReadersRequested()
}

// FuncOffListReadersRequested returns the function "WEBEXT.smartCardProviderPrivate.onListReadersRequested.removeListener".
func FuncOffListReadersRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardContext int32)])]) {
	bindings.FuncOffListReadersRequested(
		js.Pointer(&fn),
	)
	return
}

// OffListReadersRequested calls the function "WEBEXT.smartCardProviderPrivate.onListReadersRequested.removeListener" directly.
func OffListReadersRequested(callback js.Func[func(requestId int32, scardContext int32)]) (ret js.Void) {
	bindings.CallOffListReadersRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffListReadersRequested calls the function "WEBEXT.smartCardProviderPrivate.onListReadersRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffListReadersRequested(callback js.Func[func(requestId int32, scardContext int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffListReadersRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnListReadersRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onListReadersRequested.hasListener" exists.
func HasFuncHasOnListReadersRequested() bool {
	return js.True == bindings.HasFuncHasOnListReadersRequested()
}

// FuncHasOnListReadersRequested returns the function "WEBEXT.smartCardProviderPrivate.onListReadersRequested.hasListener".
func FuncHasOnListReadersRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardContext int32)]) bool]) {
	bindings.FuncHasOnListReadersRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnListReadersRequested calls the function "WEBEXT.smartCardProviderPrivate.onListReadersRequested.hasListener" directly.
func HasOnListReadersRequested(callback js.Func[func(requestId int32, scardContext int32)]) (ret bool) {
	bindings.CallHasOnListReadersRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnListReadersRequested calls the function "WEBEXT.smartCardProviderPrivate.onListReadersRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnListReadersRequested(callback js.Func[func(requestId int32, scardContext int32)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnListReadersRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnReleaseContextRequestedEventCallbackFunc func(this js.Ref, requestId int32, scardContext int32) js.Ref

func (fn OnReleaseContextRequestedEventCallbackFunc) Register() js.Func[func(requestId int32, scardContext int32)] {
	return js.RegisterCallback[func(requestId int32, scardContext int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnReleaseContextRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnReleaseContextRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, requestId int32, scardContext int32) js.Ref
	Arg T
}

func (cb *OnReleaseContextRequestedEventCallback[T]) Register() js.Func[func(requestId int32, scardContext int32)] {
	return js.RegisterCallback[func(requestId int32, scardContext int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnReleaseContextRequestedEventCallback[T]) DispatchCallback(
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnReleaseContextRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onReleaseContextRequested.addListener" exists.
func HasFuncOnReleaseContextRequested() bool {
	return js.True == bindings.HasFuncOnReleaseContextRequested()
}

// FuncOnReleaseContextRequested returns the function "WEBEXT.smartCardProviderPrivate.onReleaseContextRequested.addListener".
func FuncOnReleaseContextRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardContext int32)])]) {
	bindings.FuncOnReleaseContextRequested(
		js.Pointer(&fn),
	)
	return
}

// OnReleaseContextRequested calls the function "WEBEXT.smartCardProviderPrivate.onReleaseContextRequested.addListener" directly.
func OnReleaseContextRequested(callback js.Func[func(requestId int32, scardContext int32)]) (ret js.Void) {
	bindings.CallOnReleaseContextRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnReleaseContextRequested calls the function "WEBEXT.smartCardProviderPrivate.onReleaseContextRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnReleaseContextRequested(callback js.Func[func(requestId int32, scardContext int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnReleaseContextRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffReleaseContextRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onReleaseContextRequested.removeListener" exists.
func HasFuncOffReleaseContextRequested() bool {
	return js.True == bindings.HasFuncOffReleaseContextRequested()
}

// FuncOffReleaseContextRequested returns the function "WEBEXT.smartCardProviderPrivate.onReleaseContextRequested.removeListener".
func FuncOffReleaseContextRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardContext int32)])]) {
	bindings.FuncOffReleaseContextRequested(
		js.Pointer(&fn),
	)
	return
}

// OffReleaseContextRequested calls the function "WEBEXT.smartCardProviderPrivate.onReleaseContextRequested.removeListener" directly.
func OffReleaseContextRequested(callback js.Func[func(requestId int32, scardContext int32)]) (ret js.Void) {
	bindings.CallOffReleaseContextRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffReleaseContextRequested calls the function "WEBEXT.smartCardProviderPrivate.onReleaseContextRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffReleaseContextRequested(callback js.Func[func(requestId int32, scardContext int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffReleaseContextRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnReleaseContextRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onReleaseContextRequested.hasListener" exists.
func HasFuncHasOnReleaseContextRequested() bool {
	return js.True == bindings.HasFuncHasOnReleaseContextRequested()
}

// FuncHasOnReleaseContextRequested returns the function "WEBEXT.smartCardProviderPrivate.onReleaseContextRequested.hasListener".
func FuncHasOnReleaseContextRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardContext int32)]) bool]) {
	bindings.FuncHasOnReleaseContextRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnReleaseContextRequested calls the function "WEBEXT.smartCardProviderPrivate.onReleaseContextRequested.hasListener" directly.
func HasOnReleaseContextRequested(callback js.Func[func(requestId int32, scardContext int32)]) (ret bool) {
	bindings.CallHasOnReleaseContextRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnReleaseContextRequested calls the function "WEBEXT.smartCardProviderPrivate.onReleaseContextRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnReleaseContextRequested(callback js.Func[func(requestId int32, scardContext int32)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnReleaseContextRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSetAttribRequestedEventCallbackFunc func(this js.Ref, requestId int32, scardHandle int32, attribId int32, data js.ArrayBuffer) js.Ref

func (fn OnSetAttribRequestedEventCallbackFunc) Register() js.Func[func(requestId int32, scardHandle int32, attribId int32, data js.ArrayBuffer)] {
	return js.RegisterCallback[func(requestId int32, scardHandle int32, attribId int32, data js.ArrayBuffer)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSetAttribRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 4+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
		js.Number[int32]{}.FromRef(args[2+1]).Get(),
		js.ArrayBuffer{}.FromRef(args[3+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnSetAttribRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, requestId int32, scardHandle int32, attribId int32, data js.ArrayBuffer) js.Ref
	Arg T
}

func (cb *OnSetAttribRequestedEventCallback[T]) Register() js.Func[func(requestId int32, scardHandle int32, attribId int32, data js.ArrayBuffer)] {
	return js.RegisterCallback[func(requestId int32, scardHandle int32, attribId int32, data js.ArrayBuffer)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSetAttribRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 4+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
		js.Number[int32]{}.FromRef(args[2+1]).Get(),
		js.ArrayBuffer{}.FromRef(args[3+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnSetAttribRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onSetAttribRequested.addListener" exists.
func HasFuncOnSetAttribRequested() bool {
	return js.True == bindings.HasFuncOnSetAttribRequested()
}

// FuncOnSetAttribRequested returns the function "WEBEXT.smartCardProviderPrivate.onSetAttribRequested.addListener".
func FuncOnSetAttribRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32, attribId int32, data js.ArrayBuffer)])]) {
	bindings.FuncOnSetAttribRequested(
		js.Pointer(&fn),
	)
	return
}

// OnSetAttribRequested calls the function "WEBEXT.smartCardProviderPrivate.onSetAttribRequested.addListener" directly.
func OnSetAttribRequested(callback js.Func[func(requestId int32, scardHandle int32, attribId int32, data js.ArrayBuffer)]) (ret js.Void) {
	bindings.CallOnSetAttribRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSetAttribRequested calls the function "WEBEXT.smartCardProviderPrivate.onSetAttribRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSetAttribRequested(callback js.Func[func(requestId int32, scardHandle int32, attribId int32, data js.ArrayBuffer)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSetAttribRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSetAttribRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onSetAttribRequested.removeListener" exists.
func HasFuncOffSetAttribRequested() bool {
	return js.True == bindings.HasFuncOffSetAttribRequested()
}

// FuncOffSetAttribRequested returns the function "WEBEXT.smartCardProviderPrivate.onSetAttribRequested.removeListener".
func FuncOffSetAttribRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32, attribId int32, data js.ArrayBuffer)])]) {
	bindings.FuncOffSetAttribRequested(
		js.Pointer(&fn),
	)
	return
}

// OffSetAttribRequested calls the function "WEBEXT.smartCardProviderPrivate.onSetAttribRequested.removeListener" directly.
func OffSetAttribRequested(callback js.Func[func(requestId int32, scardHandle int32, attribId int32, data js.ArrayBuffer)]) (ret js.Void) {
	bindings.CallOffSetAttribRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSetAttribRequested calls the function "WEBEXT.smartCardProviderPrivate.onSetAttribRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSetAttribRequested(callback js.Func[func(requestId int32, scardHandle int32, attribId int32, data js.ArrayBuffer)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSetAttribRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSetAttribRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onSetAttribRequested.hasListener" exists.
func HasFuncHasOnSetAttribRequested() bool {
	return js.True == bindings.HasFuncHasOnSetAttribRequested()
}

// FuncHasOnSetAttribRequested returns the function "WEBEXT.smartCardProviderPrivate.onSetAttribRequested.hasListener".
func FuncHasOnSetAttribRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32, attribId int32, data js.ArrayBuffer)]) bool]) {
	bindings.FuncHasOnSetAttribRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnSetAttribRequested calls the function "WEBEXT.smartCardProviderPrivate.onSetAttribRequested.hasListener" directly.
func HasOnSetAttribRequested(callback js.Func[func(requestId int32, scardHandle int32, attribId int32, data js.ArrayBuffer)]) (ret bool) {
	bindings.CallHasOnSetAttribRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSetAttribRequested calls the function "WEBEXT.smartCardProviderPrivate.onSetAttribRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSetAttribRequested(callback js.Func[func(requestId int32, scardHandle int32, attribId int32, data js.ArrayBuffer)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSetAttribRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnStatusRequestedEventCallbackFunc func(this js.Ref, requestId int32, scardHandle int32) js.Ref

func (fn OnStatusRequestedEventCallbackFunc) Register() js.Func[func(requestId int32, scardHandle int32)] {
	return js.RegisterCallback[func(requestId int32, scardHandle int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnStatusRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnStatusRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, requestId int32, scardHandle int32) js.Ref
	Arg T
}

func (cb *OnStatusRequestedEventCallback[T]) Register() js.Func[func(requestId int32, scardHandle int32)] {
	return js.RegisterCallback[func(requestId int32, scardHandle int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnStatusRequestedEventCallback[T]) DispatchCallback(
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnStatusRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onStatusRequested.addListener" exists.
func HasFuncOnStatusRequested() bool {
	return js.True == bindings.HasFuncOnStatusRequested()
}

// FuncOnStatusRequested returns the function "WEBEXT.smartCardProviderPrivate.onStatusRequested.addListener".
func FuncOnStatusRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32)])]) {
	bindings.FuncOnStatusRequested(
		js.Pointer(&fn),
	)
	return
}

// OnStatusRequested calls the function "WEBEXT.smartCardProviderPrivate.onStatusRequested.addListener" directly.
func OnStatusRequested(callback js.Func[func(requestId int32, scardHandle int32)]) (ret js.Void) {
	bindings.CallOnStatusRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnStatusRequested calls the function "WEBEXT.smartCardProviderPrivate.onStatusRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnStatusRequested(callback js.Func[func(requestId int32, scardHandle int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnStatusRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffStatusRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onStatusRequested.removeListener" exists.
func HasFuncOffStatusRequested() bool {
	return js.True == bindings.HasFuncOffStatusRequested()
}

// FuncOffStatusRequested returns the function "WEBEXT.smartCardProviderPrivate.onStatusRequested.removeListener".
func FuncOffStatusRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32)])]) {
	bindings.FuncOffStatusRequested(
		js.Pointer(&fn),
	)
	return
}

// OffStatusRequested calls the function "WEBEXT.smartCardProviderPrivate.onStatusRequested.removeListener" directly.
func OffStatusRequested(callback js.Func[func(requestId int32, scardHandle int32)]) (ret js.Void) {
	bindings.CallOffStatusRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffStatusRequested calls the function "WEBEXT.smartCardProviderPrivate.onStatusRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffStatusRequested(callback js.Func[func(requestId int32, scardHandle int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffStatusRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnStatusRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onStatusRequested.hasListener" exists.
func HasFuncHasOnStatusRequested() bool {
	return js.True == bindings.HasFuncHasOnStatusRequested()
}

// FuncHasOnStatusRequested returns the function "WEBEXT.smartCardProviderPrivate.onStatusRequested.hasListener".
func FuncHasOnStatusRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32)]) bool]) {
	bindings.FuncHasOnStatusRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnStatusRequested calls the function "WEBEXT.smartCardProviderPrivate.onStatusRequested.hasListener" directly.
func HasOnStatusRequested(callback js.Func[func(requestId int32, scardHandle int32)]) (ret bool) {
	bindings.CallHasOnStatusRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnStatusRequested calls the function "WEBEXT.smartCardProviderPrivate.onStatusRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnStatusRequested(callback js.Func[func(requestId int32, scardHandle int32)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnStatusRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnTransmitRequestedEventCallbackFunc func(this js.Ref, requestId int32, scardHandle int32, protocol Protocol, data js.ArrayBuffer) js.Ref

func (fn OnTransmitRequestedEventCallbackFunc) Register() js.Func[func(requestId int32, scardHandle int32, protocol Protocol, data js.ArrayBuffer)] {
	return js.RegisterCallback[func(requestId int32, scardHandle int32, protocol Protocol, data js.ArrayBuffer)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnTransmitRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 4+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
		Protocol(0).FromRef(args[2+1]),
		js.ArrayBuffer{}.FromRef(args[3+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnTransmitRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, requestId int32, scardHandle int32, protocol Protocol, data js.ArrayBuffer) js.Ref
	Arg T
}

func (cb *OnTransmitRequestedEventCallback[T]) Register() js.Func[func(requestId int32, scardHandle int32, protocol Protocol, data js.ArrayBuffer)] {
	return js.RegisterCallback[func(requestId int32, scardHandle int32, protocol Protocol, data js.ArrayBuffer)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnTransmitRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 4+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
		Protocol(0).FromRef(args[2+1]),
		js.ArrayBuffer{}.FromRef(args[3+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnTransmitRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onTransmitRequested.addListener" exists.
func HasFuncOnTransmitRequested() bool {
	return js.True == bindings.HasFuncOnTransmitRequested()
}

// FuncOnTransmitRequested returns the function "WEBEXT.smartCardProviderPrivate.onTransmitRequested.addListener".
func FuncOnTransmitRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32, protocol Protocol, data js.ArrayBuffer)])]) {
	bindings.FuncOnTransmitRequested(
		js.Pointer(&fn),
	)
	return
}

// OnTransmitRequested calls the function "WEBEXT.smartCardProviderPrivate.onTransmitRequested.addListener" directly.
func OnTransmitRequested(callback js.Func[func(requestId int32, scardHandle int32, protocol Protocol, data js.ArrayBuffer)]) (ret js.Void) {
	bindings.CallOnTransmitRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnTransmitRequested calls the function "WEBEXT.smartCardProviderPrivate.onTransmitRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnTransmitRequested(callback js.Func[func(requestId int32, scardHandle int32, protocol Protocol, data js.ArrayBuffer)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnTransmitRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffTransmitRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onTransmitRequested.removeListener" exists.
func HasFuncOffTransmitRequested() bool {
	return js.True == bindings.HasFuncOffTransmitRequested()
}

// FuncOffTransmitRequested returns the function "WEBEXT.smartCardProviderPrivate.onTransmitRequested.removeListener".
func FuncOffTransmitRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32, protocol Protocol, data js.ArrayBuffer)])]) {
	bindings.FuncOffTransmitRequested(
		js.Pointer(&fn),
	)
	return
}

// OffTransmitRequested calls the function "WEBEXT.smartCardProviderPrivate.onTransmitRequested.removeListener" directly.
func OffTransmitRequested(callback js.Func[func(requestId int32, scardHandle int32, protocol Protocol, data js.ArrayBuffer)]) (ret js.Void) {
	bindings.CallOffTransmitRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffTransmitRequested calls the function "WEBEXT.smartCardProviderPrivate.onTransmitRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffTransmitRequested(callback js.Func[func(requestId int32, scardHandle int32, protocol Protocol, data js.ArrayBuffer)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffTransmitRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnTransmitRequested returns true if the function "WEBEXT.smartCardProviderPrivate.onTransmitRequested.hasListener" exists.
func HasFuncHasOnTransmitRequested() bool {
	return js.True == bindings.HasFuncHasOnTransmitRequested()
}

// FuncHasOnTransmitRequested returns the function "WEBEXT.smartCardProviderPrivate.onTransmitRequested.hasListener".
func FuncHasOnTransmitRequested() (fn js.Func[func(callback js.Func[func(requestId int32, scardHandle int32, protocol Protocol, data js.ArrayBuffer)]) bool]) {
	bindings.FuncHasOnTransmitRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnTransmitRequested calls the function "WEBEXT.smartCardProviderPrivate.onTransmitRequested.hasListener" directly.
func HasOnTransmitRequested(callback js.Func[func(requestId int32, scardHandle int32, protocol Protocol, data js.ArrayBuffer)]) (ret bool) {
	bindings.CallHasOnTransmitRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnTransmitRequested calls the function "WEBEXT.smartCardProviderPrivate.onTransmitRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnTransmitRequested(callback js.Func[func(requestId int32, scardHandle int32, protocol Protocol, data js.ArrayBuffer)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnTransmitRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncReportConnectResult returns true if the function "WEBEXT.smartCardProviderPrivate.reportConnectResult" exists.
func HasFuncReportConnectResult() bool {
	return js.True == bindings.HasFuncReportConnectResult()
}

// FuncReportConnectResult returns the function "WEBEXT.smartCardProviderPrivate.reportConnectResult".
func FuncReportConnectResult() (fn js.Func[func(requestId int32, scardHandle int32, activeProtocol Protocol, resultCode ResultCode)]) {
	bindings.FuncReportConnectResult(
		js.Pointer(&fn),
	)
	return
}

// ReportConnectResult calls the function "WEBEXT.smartCardProviderPrivate.reportConnectResult" directly.
func ReportConnectResult(requestId int32, scardHandle int32, activeProtocol Protocol, resultCode ResultCode) (ret js.Void) {
	bindings.CallReportConnectResult(
		js.Pointer(&ret),
		int32(requestId),
		int32(scardHandle),
		uint32(activeProtocol),
		uint32(resultCode),
	)

	return
}

// TryReportConnectResult calls the function "WEBEXT.smartCardProviderPrivate.reportConnectResult"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReportConnectResult(requestId int32, scardHandle int32, activeProtocol Protocol, resultCode ResultCode) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReportConnectResult(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(requestId),
		int32(scardHandle),
		uint32(activeProtocol),
		uint32(resultCode),
	)

	return
}

// HasFuncReportDataResult returns true if the function "WEBEXT.smartCardProviderPrivate.reportDataResult" exists.
func HasFuncReportDataResult() bool {
	return js.True == bindings.HasFuncReportDataResult()
}

// FuncReportDataResult returns the function "WEBEXT.smartCardProviderPrivate.reportDataResult".
func FuncReportDataResult() (fn js.Func[func(requestId int32, data js.ArrayBuffer, resultCode ResultCode)]) {
	bindings.FuncReportDataResult(
		js.Pointer(&fn),
	)
	return
}

// ReportDataResult calls the function "WEBEXT.smartCardProviderPrivate.reportDataResult" directly.
func ReportDataResult(requestId int32, data js.ArrayBuffer, resultCode ResultCode) (ret js.Void) {
	bindings.CallReportDataResult(
		js.Pointer(&ret),
		int32(requestId),
		data.Ref(),
		uint32(resultCode),
	)

	return
}

// TryReportDataResult calls the function "WEBEXT.smartCardProviderPrivate.reportDataResult"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReportDataResult(requestId int32, data js.ArrayBuffer, resultCode ResultCode) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReportDataResult(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(requestId),
		data.Ref(),
		uint32(resultCode),
	)

	return
}

// HasFuncReportEstablishContextResult returns true if the function "WEBEXT.smartCardProviderPrivate.reportEstablishContextResult" exists.
func HasFuncReportEstablishContextResult() bool {
	return js.True == bindings.HasFuncReportEstablishContextResult()
}

// FuncReportEstablishContextResult returns the function "WEBEXT.smartCardProviderPrivate.reportEstablishContextResult".
func FuncReportEstablishContextResult() (fn js.Func[func(requestId int32, scardContext int32, resultCode ResultCode)]) {
	bindings.FuncReportEstablishContextResult(
		js.Pointer(&fn),
	)
	return
}

// ReportEstablishContextResult calls the function "WEBEXT.smartCardProviderPrivate.reportEstablishContextResult" directly.
func ReportEstablishContextResult(requestId int32, scardContext int32, resultCode ResultCode) (ret js.Void) {
	bindings.CallReportEstablishContextResult(
		js.Pointer(&ret),
		int32(requestId),
		int32(scardContext),
		uint32(resultCode),
	)

	return
}

// TryReportEstablishContextResult calls the function "WEBEXT.smartCardProviderPrivate.reportEstablishContextResult"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReportEstablishContextResult(requestId int32, scardContext int32, resultCode ResultCode) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReportEstablishContextResult(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(requestId),
		int32(scardContext),
		uint32(resultCode),
	)

	return
}

// HasFuncReportGetStatusChangeResult returns true if the function "WEBEXT.smartCardProviderPrivate.reportGetStatusChangeResult" exists.
func HasFuncReportGetStatusChangeResult() bool {
	return js.True == bindings.HasFuncReportGetStatusChangeResult()
}

// FuncReportGetStatusChangeResult returns the function "WEBEXT.smartCardProviderPrivate.reportGetStatusChangeResult".
func FuncReportGetStatusChangeResult() (fn js.Func[func(requestId int32, readerStates js.Array[ReaderStateOut], resultCode ResultCode)]) {
	bindings.FuncReportGetStatusChangeResult(
		js.Pointer(&fn),
	)
	return
}

// ReportGetStatusChangeResult calls the function "WEBEXT.smartCardProviderPrivate.reportGetStatusChangeResult" directly.
func ReportGetStatusChangeResult(requestId int32, readerStates js.Array[ReaderStateOut], resultCode ResultCode) (ret js.Void) {
	bindings.CallReportGetStatusChangeResult(
		js.Pointer(&ret),
		int32(requestId),
		readerStates.Ref(),
		uint32(resultCode),
	)

	return
}

// TryReportGetStatusChangeResult calls the function "WEBEXT.smartCardProviderPrivate.reportGetStatusChangeResult"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReportGetStatusChangeResult(requestId int32, readerStates js.Array[ReaderStateOut], resultCode ResultCode) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReportGetStatusChangeResult(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(requestId),
		readerStates.Ref(),
		uint32(resultCode),
	)

	return
}

// HasFuncReportListReadersResult returns true if the function "WEBEXT.smartCardProviderPrivate.reportListReadersResult" exists.
func HasFuncReportListReadersResult() bool {
	return js.True == bindings.HasFuncReportListReadersResult()
}

// FuncReportListReadersResult returns the function "WEBEXT.smartCardProviderPrivate.reportListReadersResult".
func FuncReportListReadersResult() (fn js.Func[func(requestId int32, readers js.Array[js.String], resultCode ResultCode)]) {
	bindings.FuncReportListReadersResult(
		js.Pointer(&fn),
	)
	return
}

// ReportListReadersResult calls the function "WEBEXT.smartCardProviderPrivate.reportListReadersResult" directly.
func ReportListReadersResult(requestId int32, readers js.Array[js.String], resultCode ResultCode) (ret js.Void) {
	bindings.CallReportListReadersResult(
		js.Pointer(&ret),
		int32(requestId),
		readers.Ref(),
		uint32(resultCode),
	)

	return
}

// TryReportListReadersResult calls the function "WEBEXT.smartCardProviderPrivate.reportListReadersResult"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReportListReadersResult(requestId int32, readers js.Array[js.String], resultCode ResultCode) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReportListReadersResult(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(requestId),
		readers.Ref(),
		uint32(resultCode),
	)

	return
}

// HasFuncReportPlainResult returns true if the function "WEBEXT.smartCardProviderPrivate.reportPlainResult" exists.
func HasFuncReportPlainResult() bool {
	return js.True == bindings.HasFuncReportPlainResult()
}

// FuncReportPlainResult returns the function "WEBEXT.smartCardProviderPrivate.reportPlainResult".
func FuncReportPlainResult() (fn js.Func[func(requestId int32, resultCode ResultCode)]) {
	bindings.FuncReportPlainResult(
		js.Pointer(&fn),
	)
	return
}

// ReportPlainResult calls the function "WEBEXT.smartCardProviderPrivate.reportPlainResult" directly.
func ReportPlainResult(requestId int32, resultCode ResultCode) (ret js.Void) {
	bindings.CallReportPlainResult(
		js.Pointer(&ret),
		int32(requestId),
		uint32(resultCode),
	)

	return
}

// TryReportPlainResult calls the function "WEBEXT.smartCardProviderPrivate.reportPlainResult"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReportPlainResult(requestId int32, resultCode ResultCode) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReportPlainResult(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(requestId),
		uint32(resultCode),
	)

	return
}

// HasFuncReportReleaseContextResult returns true if the function "WEBEXT.smartCardProviderPrivate.reportReleaseContextResult" exists.
func HasFuncReportReleaseContextResult() bool {
	return js.True == bindings.HasFuncReportReleaseContextResult()
}

// FuncReportReleaseContextResult returns the function "WEBEXT.smartCardProviderPrivate.reportReleaseContextResult".
func FuncReportReleaseContextResult() (fn js.Func[func(requestId int32, resultCode ResultCode)]) {
	bindings.FuncReportReleaseContextResult(
		js.Pointer(&fn),
	)
	return
}

// ReportReleaseContextResult calls the function "WEBEXT.smartCardProviderPrivate.reportReleaseContextResult" directly.
func ReportReleaseContextResult(requestId int32, resultCode ResultCode) (ret js.Void) {
	bindings.CallReportReleaseContextResult(
		js.Pointer(&ret),
		int32(requestId),
		uint32(resultCode),
	)

	return
}

// TryReportReleaseContextResult calls the function "WEBEXT.smartCardProviderPrivate.reportReleaseContextResult"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReportReleaseContextResult(requestId int32, resultCode ResultCode) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReportReleaseContextResult(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(requestId),
		uint32(resultCode),
	)

	return
}

// HasFuncReportStatusResult returns true if the function "WEBEXT.smartCardProviderPrivate.reportStatusResult" exists.
func HasFuncReportStatusResult() bool {
	return js.True == bindings.HasFuncReportStatusResult()
}

// FuncReportStatusResult returns the function "WEBEXT.smartCardProviderPrivate.reportStatusResult".
func FuncReportStatusResult() (fn js.Func[func(requestId int32, readerName js.String, state ConnectionState, protocol Protocol, atr js.ArrayBuffer, resultCode ResultCode)]) {
	bindings.FuncReportStatusResult(
		js.Pointer(&fn),
	)
	return
}

// ReportStatusResult calls the function "WEBEXT.smartCardProviderPrivate.reportStatusResult" directly.
func ReportStatusResult(requestId int32, readerName js.String, state ConnectionState, protocol Protocol, atr js.ArrayBuffer, resultCode ResultCode) (ret js.Void) {
	bindings.CallReportStatusResult(
		js.Pointer(&ret),
		int32(requestId),
		readerName.Ref(),
		uint32(state),
		uint32(protocol),
		atr.Ref(),
		uint32(resultCode),
	)

	return
}

// TryReportStatusResult calls the function "WEBEXT.smartCardProviderPrivate.reportStatusResult"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReportStatusResult(requestId int32, readerName js.String, state ConnectionState, protocol Protocol, atr js.ArrayBuffer, resultCode ResultCode) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReportStatusResult(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(requestId),
		readerName.Ref(),
		uint32(state),
		uint32(protocol),
		atr.Ref(),
		uint32(resultCode),
	)

	return
}
