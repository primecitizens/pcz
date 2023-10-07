// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package diagnostics

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/chromeos/diagnostics/bindings"
)

type AcPowerStatus uint32

const (
	_ AcPowerStatus = iota

	AcPowerStatus_CONNECTED
	AcPowerStatus_DISCONNECTED
)

func (AcPowerStatus) FromRef(str js.Ref) AcPowerStatus {
	return AcPowerStatus(bindings.ConstOfAcPowerStatus(str))
}

func (x AcPowerStatus) String() (string, bool) {
	switch x {
	case AcPowerStatus_CONNECTED:
		return "connected", true
	case AcPowerStatus_DISCONNECTED:
		return "disconnected", true
	default:
		return "", false
	}
}

type CancelRoutineCallbackFunc func(this js.Ref) js.Ref

func (fn CancelRoutineCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CancelRoutineCallbackFunc) DispatchCallback(
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

type CancelRoutineCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *CancelRoutineCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CancelRoutineCallback[T]) DispatchCallback(
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

type CancelRoutineRequest struct {
	// Uuid is "CancelRoutineRequest.uuid"
	//
	// Optional
	Uuid js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CancelRoutineRequest with all fields set.
func (p CancelRoutineRequest) FromRef(ref js.Ref) CancelRoutineRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CancelRoutineRequest in the application heap.
func (p CancelRoutineRequest) New() js.Ref {
	return bindings.CancelRoutineRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CancelRoutineRequest) UpdateFrom(ref js.Ref) {
	bindings.CancelRoutineRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CancelRoutineRequest) Update(ref js.Ref) {
	bindings.CancelRoutineRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CancelRoutineRequest) FreeMembers(recursive bool) {
	js.Free(
		p.Uuid.Ref(),
	)
	p.Uuid = p.Uuid.FromRef(js.Undefined)
}

type CreateMemoryRoutineCallbackFunc func(this js.Ref, response *CreateMemoryRoutineResponse) js.Ref

func (fn CreateMemoryRoutineCallbackFunc) Register() js.Func[func(response *CreateMemoryRoutineResponse)] {
	return js.RegisterCallback[func(response *CreateMemoryRoutineResponse)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CreateMemoryRoutineCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CreateMemoryRoutineResponse
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

type CreateMemoryRoutineCallback[T any] struct {
	Fn  func(arg T, this js.Ref, response *CreateMemoryRoutineResponse) js.Ref
	Arg T
}

func (cb *CreateMemoryRoutineCallback[T]) Register() js.Func[func(response *CreateMemoryRoutineResponse)] {
	return js.RegisterCallback[func(response *CreateMemoryRoutineResponse)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CreateMemoryRoutineCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CreateMemoryRoutineResponse
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

type CreateMemoryRoutineResponse struct {
	// Uuid is "CreateMemoryRoutineResponse.uuid"
	//
	// Optional
	Uuid js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CreateMemoryRoutineResponse with all fields set.
func (p CreateMemoryRoutineResponse) FromRef(ref js.Ref) CreateMemoryRoutineResponse {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CreateMemoryRoutineResponse in the application heap.
func (p CreateMemoryRoutineResponse) New() js.Ref {
	return bindings.CreateMemoryRoutineResponseJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CreateMemoryRoutineResponse) UpdateFrom(ref js.Ref) {
	bindings.CreateMemoryRoutineResponseJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CreateMemoryRoutineResponse) Update(ref js.Ref) {
	bindings.CreateMemoryRoutineResponseJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CreateMemoryRoutineResponse) FreeMembers(recursive bool) {
	js.Free(
		p.Uuid.Ref(),
	)
	p.Uuid = p.Uuid.FromRef(js.Undefined)
}

type DiskReadRoutineType uint32

const (
	_ DiskReadRoutineType = iota

	DiskReadRoutineType_LINEAR
	DiskReadRoutineType_RANDOM
)

func (DiskReadRoutineType) FromRef(str js.Ref) DiskReadRoutineType {
	return DiskReadRoutineType(bindings.ConstOfDiskReadRoutineType(str))
}

func (x DiskReadRoutineType) String() (string, bool) {
	switch x {
	case DiskReadRoutineType_LINEAR:
		return "linear", true
	case DiskReadRoutineType_RANDOM:
		return "random", true
	default:
		return "", false
	}
}

type EOFFunc func(this js.Ref) js.Ref

func (fn EOFFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn EOFFunc) DispatchCallback(
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

type EOF[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *EOF[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *EOF[T]) DispatchCallback(
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

type ExceptionReason uint32

const (
	_ ExceptionReason = iota

	ExceptionReason_UNKNOWN
	ExceptionReason_UNEXPECTED
	ExceptionReason_UNSUPPORTED
	ExceptionReason_APP_UI_CLOSED
)

func (ExceptionReason) FromRef(str js.Ref) ExceptionReason {
	return ExceptionReason(bindings.ConstOfExceptionReason(str))
}

func (x ExceptionReason) String() (string, bool) {
	switch x {
	case ExceptionReason_UNKNOWN:
		return "unknown", true
	case ExceptionReason_UNEXPECTED:
		return "unexpected", true
	case ExceptionReason_UNSUPPORTED:
		return "unsupported", true
	case ExceptionReason_APP_UI_CLOSED:
		return "app_ui_closed", true
	default:
		return "", false
	}
}

type ExceptionInfo struct {
	// Uuid is "ExceptionInfo.uuid"
	//
	// Optional
	Uuid js.String
	// Reason is "ExceptionInfo.reason"
	//
	// Optional
	Reason ExceptionReason
	// DebugMessage is "ExceptionInfo.debugMessage"
	//
	// Optional
	DebugMessage js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ExceptionInfo with all fields set.
func (p ExceptionInfo) FromRef(ref js.Ref) ExceptionInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ExceptionInfo in the application heap.
func (p ExceptionInfo) New() js.Ref {
	return bindings.ExceptionInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ExceptionInfo) UpdateFrom(ref js.Ref) {
	bindings.ExceptionInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ExceptionInfo) Update(ref js.Ref) {
	bindings.ExceptionInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ExceptionInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Uuid.Ref(),
		p.DebugMessage.Ref(),
	)
	p.Uuid = p.Uuid.FromRef(js.Undefined)
	p.DebugMessage = p.DebugMessage.FromRef(js.Undefined)
}

type GetAvailableRoutinesCallbackFunc func(this js.Ref, response *GetAvailableRoutinesResponse) js.Ref

func (fn GetAvailableRoutinesCallbackFunc) Register() js.Func[func(response *GetAvailableRoutinesResponse)] {
	return js.RegisterCallback[func(response *GetAvailableRoutinesResponse)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetAvailableRoutinesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 GetAvailableRoutinesResponse
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

type GetAvailableRoutinesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, response *GetAvailableRoutinesResponse) js.Ref
	Arg T
}

func (cb *GetAvailableRoutinesCallback[T]) Register() js.Func[func(response *GetAvailableRoutinesResponse)] {
	return js.RegisterCallback[func(response *GetAvailableRoutinesResponse)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetAvailableRoutinesCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 GetAvailableRoutinesResponse
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

type RoutineType uint32

const (
	_ RoutineType = iota

	RoutineType_AC_POWER
	RoutineType_BATTERY_CAPACITY
	RoutineType_BATTERY_CHARGE
	RoutineType_BATTERY_DISCHARGE
	RoutineType_BATTERY_HEALTH
	RoutineType_CPU_CACHE
	RoutineType_CPU_FLOATING_POINT_ACCURACY
	RoutineType_CPU_PRIME_SEARCH
	RoutineType_CPU_STRESS
	RoutineType_DISK_READ
	RoutineType_DNS_RESOLUTION
	RoutineType_MEMORY
	RoutineType_NVME_WEAR_LEVEL
	RoutineType_SMARTCTL_CHECK
	RoutineType_LAN_CONNECTIVITY
	RoutineType_SIGNAL_STRENGTH
	RoutineType_DNS_RESOLVER_PRESENT
	RoutineType_GATEWAY_CAN_BE_PINGED
	RoutineType_SENSITIVE_SENSOR
	RoutineType_NVME_SELF_TEST
	RoutineType_FINGERPRINT_ALIVE
	RoutineType_SMARTCTL_CHECK_WITH_PERCENTAGE_USED
	RoutineType_EMMC_LIFETIME
	RoutineType_BLUETOOTH_POWER
	RoutineType_UFS_LIFETIME
	RoutineType_POWER_BUTTON
	RoutineType_AUDIO_DRIVER
	RoutineType_BLUETOOTH_DISCOVERY
	RoutineType_BLUETOOTH_SCANNING
	RoutineType_BLUETOOTH_PAIRING
)

func (RoutineType) FromRef(str js.Ref) RoutineType {
	return RoutineType(bindings.ConstOfRoutineType(str))
}

func (x RoutineType) String() (string, bool) {
	switch x {
	case RoutineType_AC_POWER:
		return "ac_power", true
	case RoutineType_BATTERY_CAPACITY:
		return "battery_capacity", true
	case RoutineType_BATTERY_CHARGE:
		return "battery_charge", true
	case RoutineType_BATTERY_DISCHARGE:
		return "battery_discharge", true
	case RoutineType_BATTERY_HEALTH:
		return "battery_health", true
	case RoutineType_CPU_CACHE:
		return "cpu_cache", true
	case RoutineType_CPU_FLOATING_POINT_ACCURACY:
		return "cpu_floating_point_accuracy", true
	case RoutineType_CPU_PRIME_SEARCH:
		return "cpu_prime_search", true
	case RoutineType_CPU_STRESS:
		return "cpu_stress", true
	case RoutineType_DISK_READ:
		return "disk_read", true
	case RoutineType_DNS_RESOLUTION:
		return "dns_resolution", true
	case RoutineType_MEMORY:
		return "memory", true
	case RoutineType_NVME_WEAR_LEVEL:
		return "nvme_wear_level", true
	case RoutineType_SMARTCTL_CHECK:
		return "smartctl_check", true
	case RoutineType_LAN_CONNECTIVITY:
		return "lan_connectivity", true
	case RoutineType_SIGNAL_STRENGTH:
		return "signal_strength", true
	case RoutineType_DNS_RESOLVER_PRESENT:
		return "dns_resolver_present", true
	case RoutineType_GATEWAY_CAN_BE_PINGED:
		return "gateway_can_be_pinged", true
	case RoutineType_SENSITIVE_SENSOR:
		return "sensitive_sensor", true
	case RoutineType_NVME_SELF_TEST:
		return "nvme_self_test", true
	case RoutineType_FINGERPRINT_ALIVE:
		return "fingerprint_alive", true
	case RoutineType_SMARTCTL_CHECK_WITH_PERCENTAGE_USED:
		return "smartctl_check_with_percentage_used", true
	case RoutineType_EMMC_LIFETIME:
		return "emmc_lifetime", true
	case RoutineType_BLUETOOTH_POWER:
		return "bluetooth_power", true
	case RoutineType_UFS_LIFETIME:
		return "ufs_lifetime", true
	case RoutineType_POWER_BUTTON:
		return "power_button", true
	case RoutineType_AUDIO_DRIVER:
		return "audio_driver", true
	case RoutineType_BLUETOOTH_DISCOVERY:
		return "bluetooth_discovery", true
	case RoutineType_BLUETOOTH_SCANNING:
		return "bluetooth_scanning", true
	case RoutineType_BLUETOOTH_PAIRING:
		return "bluetooth_pairing", true
	default:
		return "", false
	}
}

type GetAvailableRoutinesResponse struct {
	// Routines is "GetAvailableRoutinesResponse.routines"
	//
	// Optional
	Routines js.Array[RoutineType]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetAvailableRoutinesResponse with all fields set.
func (p GetAvailableRoutinesResponse) FromRef(ref js.Ref) GetAvailableRoutinesResponse {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetAvailableRoutinesResponse in the application heap.
func (p GetAvailableRoutinesResponse) New() js.Ref {
	return bindings.GetAvailableRoutinesResponseJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetAvailableRoutinesResponse) UpdateFrom(ref js.Ref) {
	bindings.GetAvailableRoutinesResponseJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetAvailableRoutinesResponse) Update(ref js.Ref) {
	bindings.GetAvailableRoutinesResponseJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetAvailableRoutinesResponse) FreeMembers(recursive bool) {
	js.Free(
		p.Routines.Ref(),
	)
	p.Routines = p.Routines.FromRef(js.Undefined)
}

type GetRoutineUpdateCallbackFunc func(this js.Ref, response *GetRoutineUpdateResponse) js.Ref

func (fn GetRoutineUpdateCallbackFunc) Register() js.Func[func(response *GetRoutineUpdateResponse)] {
	return js.RegisterCallback[func(response *GetRoutineUpdateResponse)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetRoutineUpdateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 GetRoutineUpdateResponse
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

type GetRoutineUpdateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, response *GetRoutineUpdateResponse) js.Ref
	Arg T
}

func (cb *GetRoutineUpdateCallback[T]) Register() js.Func[func(response *GetRoutineUpdateResponse)] {
	return js.RegisterCallback[func(response *GetRoutineUpdateResponse)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetRoutineUpdateCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 GetRoutineUpdateResponse
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

type RoutineStatus uint32

const (
	_ RoutineStatus = iota

	RoutineStatus_UNKNOWN
	RoutineStatus_READY
	RoutineStatus_RUNNING
	RoutineStatus_WAITING_USER_ACTION
	RoutineStatus_PASSED
	RoutineStatus_FAILED
	RoutineStatus_ERROR
	RoutineStatus_CANCELLED
	RoutineStatus_FAILED_TO_START
	RoutineStatus_REMOVED
	RoutineStatus_CANCELLING
	RoutineStatus_UNSUPPORTED
	RoutineStatus_NOT_RUN
)

func (RoutineStatus) FromRef(str js.Ref) RoutineStatus {
	return RoutineStatus(bindings.ConstOfRoutineStatus(str))
}

func (x RoutineStatus) String() (string, bool) {
	switch x {
	case RoutineStatus_UNKNOWN:
		return "unknown", true
	case RoutineStatus_READY:
		return "ready", true
	case RoutineStatus_RUNNING:
		return "running", true
	case RoutineStatus_WAITING_USER_ACTION:
		return "waiting_user_action", true
	case RoutineStatus_PASSED:
		return "passed", true
	case RoutineStatus_FAILED:
		return "failed", true
	case RoutineStatus_ERROR:
		return "error", true
	case RoutineStatus_CANCELLED:
		return "cancelled", true
	case RoutineStatus_FAILED_TO_START:
		return "failed_to_start", true
	case RoutineStatus_REMOVED:
		return "removed", true
	case RoutineStatus_CANCELLING:
		return "cancelling", true
	case RoutineStatus_UNSUPPORTED:
		return "unsupported", true
	case RoutineStatus_NOT_RUN:
		return "not_run", true
	default:
		return "", false
	}
}

type UserMessageType uint32

const (
	_ UserMessageType = iota

	UserMessageType_UNKNOWN
	UserMessageType_UNPLUG_AC_POWER
	UserMessageType_PLUG_IN_AC_POWER
	UserMessageType_PRESS_POWER_BUTTON
)

func (UserMessageType) FromRef(str js.Ref) UserMessageType {
	return UserMessageType(bindings.ConstOfUserMessageType(str))
}

func (x UserMessageType) String() (string, bool) {
	switch x {
	case UserMessageType_UNKNOWN:
		return "unknown", true
	case UserMessageType_UNPLUG_AC_POWER:
		return "unplug_ac_power", true
	case UserMessageType_PLUG_IN_AC_POWER:
		return "plug_in_ac_power", true
	case UserMessageType_PRESS_POWER_BUTTON:
		return "press_power_button", true
	default:
		return "", false
	}
}

type GetRoutineUpdateResponse struct {
	// ProgressPercent is "GetRoutineUpdateResponse.progress_percent"
	//
	// Optional
	//
	// NOTE: FFI_USE_ProgressPercent MUST be set to true to make this field effective.
	ProgressPercent int32
	// Output is "GetRoutineUpdateResponse.output"
	//
	// Optional
	Output js.String
	// Status is "GetRoutineUpdateResponse.status"
	//
	// Optional
	Status RoutineStatus
	// StatusMessage is "GetRoutineUpdateResponse.status_message"
	//
	// Optional
	StatusMessage js.String
	// UserMessage is "GetRoutineUpdateResponse.user_message"
	//
	// Optional
	UserMessage UserMessageType

	FFI_USE_ProgressPercent bool // for ProgressPercent.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetRoutineUpdateResponse with all fields set.
func (p GetRoutineUpdateResponse) FromRef(ref js.Ref) GetRoutineUpdateResponse {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetRoutineUpdateResponse in the application heap.
func (p GetRoutineUpdateResponse) New() js.Ref {
	return bindings.GetRoutineUpdateResponseJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetRoutineUpdateResponse) UpdateFrom(ref js.Ref) {
	bindings.GetRoutineUpdateResponseJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetRoutineUpdateResponse) Update(ref js.Ref) {
	bindings.GetRoutineUpdateResponseJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetRoutineUpdateResponse) FreeMembers(recursive bool) {
	js.Free(
		p.Output.Ref(),
		p.StatusMessage.Ref(),
	)
	p.Output = p.Output.FromRef(js.Undefined)
	p.StatusMessage = p.StatusMessage.FromRef(js.Undefined)
}

type RoutineCommandType uint32

const (
	_ RoutineCommandType = iota

	RoutineCommandType_CANCEL
	RoutineCommandType_REMOVE
	RoutineCommandType_RESUME
	RoutineCommandType_STATUS
)

func (RoutineCommandType) FromRef(str js.Ref) RoutineCommandType {
	return RoutineCommandType(bindings.ConstOfRoutineCommandType(str))
}

func (x RoutineCommandType) String() (string, bool) {
	switch x {
	case RoutineCommandType_CANCEL:
		return "cancel", true
	case RoutineCommandType_REMOVE:
		return "remove", true
	case RoutineCommandType_RESUME:
		return "resume", true
	case RoutineCommandType_STATUS:
		return "status", true
	default:
		return "", false
	}
}

type GetRoutineUpdateRequest struct {
	// Id is "GetRoutineUpdateRequest.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id int32
	// Command is "GetRoutineUpdateRequest.command"
	//
	// Optional
	Command RoutineCommandType

	FFI_USE_Id bool // for Id.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetRoutineUpdateRequest with all fields set.
func (p GetRoutineUpdateRequest) FromRef(ref js.Ref) GetRoutineUpdateRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetRoutineUpdateRequest in the application heap.
func (p GetRoutineUpdateRequest) New() js.Ref {
	return bindings.GetRoutineUpdateRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetRoutineUpdateRequest) UpdateFrom(ref js.Ref) {
	bindings.GetRoutineUpdateRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetRoutineUpdateRequest) Update(ref js.Ref) {
	bindings.GetRoutineUpdateRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetRoutineUpdateRequest) FreeMembers(recursive bool) {
}

type MemtesterTestItemEnum uint32

const (
	_ MemtesterTestItemEnum = iota

	MemtesterTestItemEnum_UNKNOWN
	MemtesterTestItemEnum_STUCK_ADDRESS
	MemtesterTestItemEnum_COMPARE_AND
	MemtesterTestItemEnum_COMPARE_DIV
	MemtesterTestItemEnum_COMPARE_MUL
	MemtesterTestItemEnum_COMPARE_OR
	MemtesterTestItemEnum_COMPARE_SUB
	MemtesterTestItemEnum_COMPARE_XOR
	MemtesterTestItemEnum_SEQUENTIAL_INCREMENT
	MemtesterTestItemEnum_BIT_FLIP
	MemtesterTestItemEnum_BIT_SPREAD
	MemtesterTestItemEnum_BLOCK_SEQUENTIAL
	MemtesterTestItemEnum_CHECKERBOARD
	MemtesterTestItemEnum_RANDOM_VALUE
	MemtesterTestItemEnum_SOLID_BITS
	MemtesterTestItemEnum_WALKING_ONES
	MemtesterTestItemEnum_WALKING_ZEROES
	MemtesterTestItemEnum_EIGHT_BIT_WRITES
	MemtesterTestItemEnum_SIXTEEN_BIT_WRITES
)

func (MemtesterTestItemEnum) FromRef(str js.Ref) MemtesterTestItemEnum {
	return MemtesterTestItemEnum(bindings.ConstOfMemtesterTestItemEnum(str))
}

func (x MemtesterTestItemEnum) String() (string, bool) {
	switch x {
	case MemtesterTestItemEnum_UNKNOWN:
		return "unknown", true
	case MemtesterTestItemEnum_STUCK_ADDRESS:
		return "stuck_address", true
	case MemtesterTestItemEnum_COMPARE_AND:
		return "compare_and", true
	case MemtesterTestItemEnum_COMPARE_DIV:
		return "compare_div", true
	case MemtesterTestItemEnum_COMPARE_MUL:
		return "compare_mul", true
	case MemtesterTestItemEnum_COMPARE_OR:
		return "compare_or", true
	case MemtesterTestItemEnum_COMPARE_SUB:
		return "compare_sub", true
	case MemtesterTestItemEnum_COMPARE_XOR:
		return "compare_xor", true
	case MemtesterTestItemEnum_SEQUENTIAL_INCREMENT:
		return "sequential_increment", true
	case MemtesterTestItemEnum_BIT_FLIP:
		return "bit_flip", true
	case MemtesterTestItemEnum_BIT_SPREAD:
		return "bit_spread", true
	case MemtesterTestItemEnum_BLOCK_SEQUENTIAL:
		return "block_sequential", true
	case MemtesterTestItemEnum_CHECKERBOARD:
		return "checkerboard", true
	case MemtesterTestItemEnum_RANDOM_VALUE:
		return "random_value", true
	case MemtesterTestItemEnum_SOLID_BITS:
		return "solid_bits", true
	case MemtesterTestItemEnum_WALKING_ONES:
		return "walking_ones", true
	case MemtesterTestItemEnum_WALKING_ZEROES:
		return "walking_zeroes", true
	case MemtesterTestItemEnum_EIGHT_BIT_WRITES:
		return "eight_bit_writes", true
	case MemtesterTestItemEnum_SIXTEEN_BIT_WRITES:
		return "sixteen_bit_writes", true
	default:
		return "", false
	}
}

type MemtesterResult struct {
	// PassedItems is "MemtesterResult.passed_items"
	//
	// Optional
	PassedItems js.Array[MemtesterTestItemEnum]
	// FailedItems is "MemtesterResult.failed_items"
	//
	// Optional
	FailedItems js.Array[MemtesterTestItemEnum]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MemtesterResult with all fields set.
func (p MemtesterResult) FromRef(ref js.Ref) MemtesterResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MemtesterResult in the application heap.
func (p MemtesterResult) New() js.Ref {
	return bindings.MemtesterResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MemtesterResult) UpdateFrom(ref js.Ref) {
	bindings.MemtesterResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MemtesterResult) Update(ref js.Ref) {
	bindings.MemtesterResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MemtesterResult) FreeMembers(recursive bool) {
	js.Free(
		p.PassedItems.Ref(),
		p.FailedItems.Ref(),
	)
	p.PassedItems = p.PassedItems.FromRef(js.Undefined)
	p.FailedItems = p.FailedItems.FromRef(js.Undefined)
}

type MemoryRoutineFinishedInfo struct {
	// Uuid is "MemoryRoutineFinishedInfo.uuid"
	//
	// Optional
	Uuid js.String
	// HasPassed is "MemoryRoutineFinishedInfo.has_passed"
	//
	// Optional
	//
	// NOTE: FFI_USE_HasPassed MUST be set to true to make this field effective.
	HasPassed bool
	// BytesTested is "MemoryRoutineFinishedInfo.bytesTested"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesTested MUST be set to true to make this field effective.
	BytesTested float64
	// Result is "MemoryRoutineFinishedInfo.result"
	//
	// Optional
	//
	// NOTE: Result.FFI_USE MUST be set to true to get Result used.
	Result MemtesterResult

	FFI_USE_HasPassed   bool // for HasPassed.
	FFI_USE_BytesTested bool // for BytesTested.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MemoryRoutineFinishedInfo with all fields set.
func (p MemoryRoutineFinishedInfo) FromRef(ref js.Ref) MemoryRoutineFinishedInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MemoryRoutineFinishedInfo in the application heap.
func (p MemoryRoutineFinishedInfo) New() js.Ref {
	return bindings.MemoryRoutineFinishedInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MemoryRoutineFinishedInfo) UpdateFrom(ref js.Ref) {
	bindings.MemoryRoutineFinishedInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MemoryRoutineFinishedInfo) Update(ref js.Ref) {
	bindings.MemoryRoutineFinishedInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MemoryRoutineFinishedInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Uuid.Ref(),
	)
	p.Uuid = p.Uuid.FromRef(js.Undefined)
	if recursive {
		p.Result.FreeMembers(true)
	}
}

type NvmeSelfTestType uint32

const (
	_ NvmeSelfTestType = iota

	NvmeSelfTestType_SHORT_TEST
	NvmeSelfTestType_LONG_TEST
)

func (NvmeSelfTestType) FromRef(str js.Ref) NvmeSelfTestType {
	return NvmeSelfTestType(bindings.ConstOfNvmeSelfTestType(str))
}

func (x NvmeSelfTestType) String() (string, bool) {
	switch x {
	case NvmeSelfTestType_SHORT_TEST:
		return "short_test", true
	case NvmeSelfTestType_LONG_TEST:
		return "long_test", true
	default:
		return "", false
	}
}

type RoutineInitializedInfo struct {
	// Uuid is "RoutineInitializedInfo.uuid"
	//
	// Optional
	Uuid js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RoutineInitializedInfo with all fields set.
func (p RoutineInitializedInfo) FromRef(ref js.Ref) RoutineInitializedInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RoutineInitializedInfo in the application heap.
func (p RoutineInitializedInfo) New() js.Ref {
	return bindings.RoutineInitializedInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RoutineInitializedInfo) UpdateFrom(ref js.Ref) {
	bindings.RoutineInitializedInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RoutineInitializedInfo) Update(ref js.Ref) {
	bindings.RoutineInitializedInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RoutineInitializedInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Uuid.Ref(),
	)
	p.Uuid = p.Uuid.FromRef(js.Undefined)
}

type RoutineRunningInfo struct {
	// Uuid is "RoutineRunningInfo.uuid"
	//
	// Optional
	Uuid js.String
	// Percentage is "RoutineRunningInfo.percentage"
	//
	// Optional
	//
	// NOTE: FFI_USE_Percentage MUST be set to true to make this field effective.
	Percentage int32

	FFI_USE_Percentage bool // for Percentage.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RoutineRunningInfo with all fields set.
func (p RoutineRunningInfo) FromRef(ref js.Ref) RoutineRunningInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RoutineRunningInfo in the application heap.
func (p RoutineRunningInfo) New() js.Ref {
	return bindings.RoutineRunningInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RoutineRunningInfo) UpdateFrom(ref js.Ref) {
	bindings.RoutineRunningInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RoutineRunningInfo) Update(ref js.Ref) {
	bindings.RoutineRunningInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RoutineRunningInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Uuid.Ref(),
	)
	p.Uuid = p.Uuid.FromRef(js.Undefined)
}

type RoutineSupportStatus uint32

const (
	_ RoutineSupportStatus = iota

	RoutineSupportStatus_SUPPORTED
	RoutineSupportStatus_UNSUPPORTED
)

func (RoutineSupportStatus) FromRef(str js.Ref) RoutineSupportStatus {
	return RoutineSupportStatus(bindings.ConstOfRoutineSupportStatus(str))
}

func (x RoutineSupportStatus) String() (string, bool) {
	switch x {
	case RoutineSupportStatus_SUPPORTED:
		return "supported", true
	case RoutineSupportStatus_UNSUPPORTED:
		return "unsupported", true
	default:
		return "", false
	}
}

type RoutineSupportStatusInfo struct {
	// Status is "RoutineSupportStatusInfo.status"
	//
	// Optional
	Status RoutineSupportStatus

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RoutineSupportStatusInfo with all fields set.
func (p RoutineSupportStatusInfo) FromRef(ref js.Ref) RoutineSupportStatusInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RoutineSupportStatusInfo in the application heap.
func (p RoutineSupportStatusInfo) New() js.Ref {
	return bindings.RoutineSupportStatusInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RoutineSupportStatusInfo) UpdateFrom(ref js.Ref) {
	bindings.RoutineSupportStatusInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RoutineSupportStatusInfo) Update(ref js.Ref) {
	bindings.RoutineSupportStatusInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RoutineSupportStatusInfo) FreeMembers(recursive bool) {
}

type RoutineSupportStatusInfoCallbackFunc func(this js.Ref, info *RoutineSupportStatusInfo) js.Ref

func (fn RoutineSupportStatusInfoCallbackFunc) Register() js.Func[func(info *RoutineSupportStatusInfo)] {
	return js.RegisterCallback[func(info *RoutineSupportStatusInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RoutineSupportStatusInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RoutineSupportStatusInfo
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

type RoutineSupportStatusInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *RoutineSupportStatusInfo) js.Ref
	Arg T
}

func (cb *RoutineSupportStatusInfoCallback[T]) Register() js.Func[func(info *RoutineSupportStatusInfo)] {
	return js.RegisterCallback[func(info *RoutineSupportStatusInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RoutineSupportStatusInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RoutineSupportStatusInfo
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

type RoutineWaitingReason uint32

const (
	_ RoutineWaitingReason = iota

	RoutineWaitingReason_WAITING_TO_BE_SCHEDULED
	RoutineWaitingReason_WAITING_USER_INPUT
)

func (RoutineWaitingReason) FromRef(str js.Ref) RoutineWaitingReason {
	return RoutineWaitingReason(bindings.ConstOfRoutineWaitingReason(str))
}

func (x RoutineWaitingReason) String() (string, bool) {
	switch x {
	case RoutineWaitingReason_WAITING_TO_BE_SCHEDULED:
		return "waiting_to_be_scheduled", true
	case RoutineWaitingReason_WAITING_USER_INPUT:
		return "waiting_user_input", true
	default:
		return "", false
	}
}

type RoutineWaitingInfo struct {
	// Uuid is "RoutineWaitingInfo.uuid"
	//
	// Optional
	Uuid js.String
	// Percentage is "RoutineWaitingInfo.percentage"
	//
	// Optional
	//
	// NOTE: FFI_USE_Percentage MUST be set to true to make this field effective.
	Percentage int32
	// Reason is "RoutineWaitingInfo.reason"
	//
	// Optional
	Reason RoutineWaitingReason
	// Message is "RoutineWaitingInfo.message"
	//
	// Optional
	Message js.String

	FFI_USE_Percentage bool // for Percentage.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RoutineWaitingInfo with all fields set.
func (p RoutineWaitingInfo) FromRef(ref js.Ref) RoutineWaitingInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RoutineWaitingInfo in the application heap.
func (p RoutineWaitingInfo) New() js.Ref {
	return bindings.RoutineWaitingInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RoutineWaitingInfo) UpdateFrom(ref js.Ref) {
	bindings.RoutineWaitingInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RoutineWaitingInfo) Update(ref js.Ref) {
	bindings.RoutineWaitingInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RoutineWaitingInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Uuid.Ref(),
		p.Message.Ref(),
	)
	p.Uuid = p.Uuid.FromRef(js.Undefined)
	p.Message = p.Message.FromRef(js.Undefined)
}

type RunAcPowerRoutineRequest struct {
	// ExpectedStatus is "RunAcPowerRoutineRequest.expected_status"
	//
	// Optional
	ExpectedStatus AcPowerStatus
	// ExpectedPowerType is "RunAcPowerRoutineRequest.expected_power_type"
	//
	// Optional
	ExpectedPowerType js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RunAcPowerRoutineRequest with all fields set.
func (p RunAcPowerRoutineRequest) FromRef(ref js.Ref) RunAcPowerRoutineRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RunAcPowerRoutineRequest in the application heap.
func (p RunAcPowerRoutineRequest) New() js.Ref {
	return bindings.RunAcPowerRoutineRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RunAcPowerRoutineRequest) UpdateFrom(ref js.Ref) {
	bindings.RunAcPowerRoutineRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RunAcPowerRoutineRequest) Update(ref js.Ref) {
	bindings.RunAcPowerRoutineRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RunAcPowerRoutineRequest) FreeMembers(recursive bool) {
	js.Free(
		p.ExpectedPowerType.Ref(),
	)
	p.ExpectedPowerType = p.ExpectedPowerType.FromRef(js.Undefined)
}

type RunBatteryChargeRoutineRequest struct {
	// LengthSeconds is "RunBatteryChargeRoutineRequest.length_seconds"
	//
	// Optional
	//
	// NOTE: FFI_USE_LengthSeconds MUST be set to true to make this field effective.
	LengthSeconds int32
	// MinimumChargePercentRequired is "RunBatteryChargeRoutineRequest.minimum_charge_percent_required"
	//
	// Optional
	//
	// NOTE: FFI_USE_MinimumChargePercentRequired MUST be set to true to make this field effective.
	MinimumChargePercentRequired int32

	FFI_USE_LengthSeconds                bool // for LengthSeconds.
	FFI_USE_MinimumChargePercentRequired bool // for MinimumChargePercentRequired.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RunBatteryChargeRoutineRequest with all fields set.
func (p RunBatteryChargeRoutineRequest) FromRef(ref js.Ref) RunBatteryChargeRoutineRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RunBatteryChargeRoutineRequest in the application heap.
func (p RunBatteryChargeRoutineRequest) New() js.Ref {
	return bindings.RunBatteryChargeRoutineRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RunBatteryChargeRoutineRequest) UpdateFrom(ref js.Ref) {
	bindings.RunBatteryChargeRoutineRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RunBatteryChargeRoutineRequest) Update(ref js.Ref) {
	bindings.RunBatteryChargeRoutineRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RunBatteryChargeRoutineRequest) FreeMembers(recursive bool) {
}

type RunBatteryDischargeRoutineRequest struct {
	// LengthSeconds is "RunBatteryDischargeRoutineRequest.length_seconds"
	//
	// Optional
	//
	// NOTE: FFI_USE_LengthSeconds MUST be set to true to make this field effective.
	LengthSeconds int32
	// MaximumDischargePercentAllowed is "RunBatteryDischargeRoutineRequest.maximum_discharge_percent_allowed"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaximumDischargePercentAllowed MUST be set to true to make this field effective.
	MaximumDischargePercentAllowed int32

	FFI_USE_LengthSeconds                  bool // for LengthSeconds.
	FFI_USE_MaximumDischargePercentAllowed bool // for MaximumDischargePercentAllowed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RunBatteryDischargeRoutineRequest with all fields set.
func (p RunBatteryDischargeRoutineRequest) FromRef(ref js.Ref) RunBatteryDischargeRoutineRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RunBatteryDischargeRoutineRequest in the application heap.
func (p RunBatteryDischargeRoutineRequest) New() js.Ref {
	return bindings.RunBatteryDischargeRoutineRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RunBatteryDischargeRoutineRequest) UpdateFrom(ref js.Ref) {
	bindings.RunBatteryDischargeRoutineRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RunBatteryDischargeRoutineRequest) Update(ref js.Ref) {
	bindings.RunBatteryDischargeRoutineRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RunBatteryDischargeRoutineRequest) FreeMembers(recursive bool) {
}

type RunBluetoothPairingRoutineRequest struct {
	// PeripheralId is "RunBluetoothPairingRoutineRequest.peripheral_id"
	//
	// Optional
	PeripheralId js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RunBluetoothPairingRoutineRequest with all fields set.
func (p RunBluetoothPairingRoutineRequest) FromRef(ref js.Ref) RunBluetoothPairingRoutineRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RunBluetoothPairingRoutineRequest in the application heap.
func (p RunBluetoothPairingRoutineRequest) New() js.Ref {
	return bindings.RunBluetoothPairingRoutineRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RunBluetoothPairingRoutineRequest) UpdateFrom(ref js.Ref) {
	bindings.RunBluetoothPairingRoutineRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RunBluetoothPairingRoutineRequest) Update(ref js.Ref) {
	bindings.RunBluetoothPairingRoutineRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RunBluetoothPairingRoutineRequest) FreeMembers(recursive bool) {
	js.Free(
		p.PeripheralId.Ref(),
	)
	p.PeripheralId = p.PeripheralId.FromRef(js.Undefined)
}

type RunBluetoothScanningRoutineRequest struct {
	// LengthSeconds is "RunBluetoothScanningRoutineRequest.length_seconds"
	//
	// Optional
	//
	// NOTE: FFI_USE_LengthSeconds MUST be set to true to make this field effective.
	LengthSeconds int32

	FFI_USE_LengthSeconds bool // for LengthSeconds.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RunBluetoothScanningRoutineRequest with all fields set.
func (p RunBluetoothScanningRoutineRequest) FromRef(ref js.Ref) RunBluetoothScanningRoutineRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RunBluetoothScanningRoutineRequest in the application heap.
func (p RunBluetoothScanningRoutineRequest) New() js.Ref {
	return bindings.RunBluetoothScanningRoutineRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RunBluetoothScanningRoutineRequest) UpdateFrom(ref js.Ref) {
	bindings.RunBluetoothScanningRoutineRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RunBluetoothScanningRoutineRequest) Update(ref js.Ref) {
	bindings.RunBluetoothScanningRoutineRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RunBluetoothScanningRoutineRequest) FreeMembers(recursive bool) {
}

type RunCpuRoutineRequest struct {
	// LengthSeconds is "RunCpuRoutineRequest.length_seconds"
	//
	// Optional
	//
	// NOTE: FFI_USE_LengthSeconds MUST be set to true to make this field effective.
	LengthSeconds int32

	FFI_USE_LengthSeconds bool // for LengthSeconds.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RunCpuRoutineRequest with all fields set.
func (p RunCpuRoutineRequest) FromRef(ref js.Ref) RunCpuRoutineRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RunCpuRoutineRequest in the application heap.
func (p RunCpuRoutineRequest) New() js.Ref {
	return bindings.RunCpuRoutineRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RunCpuRoutineRequest) UpdateFrom(ref js.Ref) {
	bindings.RunCpuRoutineRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RunCpuRoutineRequest) Update(ref js.Ref) {
	bindings.RunCpuRoutineRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RunCpuRoutineRequest) FreeMembers(recursive bool) {
}

type RunDiskReadRequest struct {
	// Type is "RunDiskReadRequest.type"
	//
	// Optional
	Type DiskReadRoutineType
	// LengthSeconds is "RunDiskReadRequest.length_seconds"
	//
	// Optional
	//
	// NOTE: FFI_USE_LengthSeconds MUST be set to true to make this field effective.
	LengthSeconds int32
	// FileSizeMb is "RunDiskReadRequest.file_size_mb"
	//
	// Optional
	//
	// NOTE: FFI_USE_FileSizeMb MUST be set to true to make this field effective.
	FileSizeMb int32

	FFI_USE_LengthSeconds bool // for LengthSeconds.
	FFI_USE_FileSizeMb    bool // for FileSizeMb.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RunDiskReadRequest with all fields set.
func (p RunDiskReadRequest) FromRef(ref js.Ref) RunDiskReadRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RunDiskReadRequest in the application heap.
func (p RunDiskReadRequest) New() js.Ref {
	return bindings.RunDiskReadRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RunDiskReadRequest) UpdateFrom(ref js.Ref) {
	bindings.RunDiskReadRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RunDiskReadRequest) Update(ref js.Ref) {
	bindings.RunDiskReadRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RunDiskReadRequest) FreeMembers(recursive bool) {
}

type RunMemoryRoutineArguments struct {
	// MaxTestingMemKib is "RunMemoryRoutineArguments.maxTestingMemKib"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxTestingMemKib MUST be set to true to make this field effective.
	MaxTestingMemKib int32

	FFI_USE_MaxTestingMemKib bool // for MaxTestingMemKib.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RunMemoryRoutineArguments with all fields set.
func (p RunMemoryRoutineArguments) FromRef(ref js.Ref) RunMemoryRoutineArguments {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RunMemoryRoutineArguments in the application heap.
func (p RunMemoryRoutineArguments) New() js.Ref {
	return bindings.RunMemoryRoutineArgumentsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RunMemoryRoutineArguments) UpdateFrom(ref js.Ref) {
	bindings.RunMemoryRoutineArgumentsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RunMemoryRoutineArguments) Update(ref js.Ref) {
	bindings.RunMemoryRoutineArgumentsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RunMemoryRoutineArguments) FreeMembers(recursive bool) {
}

type RunNvmeSelfTestRequest struct {
	// TestType is "RunNvmeSelfTestRequest.test_type"
	//
	// Optional
	TestType NvmeSelfTestType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RunNvmeSelfTestRequest with all fields set.
func (p RunNvmeSelfTestRequest) FromRef(ref js.Ref) RunNvmeSelfTestRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RunNvmeSelfTestRequest in the application heap.
func (p RunNvmeSelfTestRequest) New() js.Ref {
	return bindings.RunNvmeSelfTestRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RunNvmeSelfTestRequest) UpdateFrom(ref js.Ref) {
	bindings.RunNvmeSelfTestRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RunNvmeSelfTestRequest) Update(ref js.Ref) {
	bindings.RunNvmeSelfTestRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RunNvmeSelfTestRequest) FreeMembers(recursive bool) {
}

type RunNvmeWearLevelRequest struct {
	// WearLevelThreshold is "RunNvmeWearLevelRequest.wear_level_threshold"
	//
	// Optional
	//
	// NOTE: FFI_USE_WearLevelThreshold MUST be set to true to make this field effective.
	WearLevelThreshold int32

	FFI_USE_WearLevelThreshold bool // for WearLevelThreshold.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RunNvmeWearLevelRequest with all fields set.
func (p RunNvmeWearLevelRequest) FromRef(ref js.Ref) RunNvmeWearLevelRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RunNvmeWearLevelRequest in the application heap.
func (p RunNvmeWearLevelRequest) New() js.Ref {
	return bindings.RunNvmeWearLevelRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RunNvmeWearLevelRequest) UpdateFrom(ref js.Ref) {
	bindings.RunNvmeWearLevelRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RunNvmeWearLevelRequest) Update(ref js.Ref) {
	bindings.RunNvmeWearLevelRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RunNvmeWearLevelRequest) FreeMembers(recursive bool) {
}

type RunPowerButtonRequest struct {
	// TimeoutSeconds is "RunPowerButtonRequest.timeout_seconds"
	//
	// Optional
	//
	// NOTE: FFI_USE_TimeoutSeconds MUST be set to true to make this field effective.
	TimeoutSeconds int32

	FFI_USE_TimeoutSeconds bool // for TimeoutSeconds.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RunPowerButtonRequest with all fields set.
func (p RunPowerButtonRequest) FromRef(ref js.Ref) RunPowerButtonRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RunPowerButtonRequest in the application heap.
func (p RunPowerButtonRequest) New() js.Ref {
	return bindings.RunPowerButtonRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RunPowerButtonRequest) UpdateFrom(ref js.Ref) {
	bindings.RunPowerButtonRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RunPowerButtonRequest) Update(ref js.Ref) {
	bindings.RunPowerButtonRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RunPowerButtonRequest) FreeMembers(recursive bool) {
}

type RunRoutineCallbackFunc func(this js.Ref, response *RunRoutineResponse) js.Ref

func (fn RunRoutineCallbackFunc) Register() js.Func[func(response *RunRoutineResponse)] {
	return js.RegisterCallback[func(response *RunRoutineResponse)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RunRoutineCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RunRoutineResponse
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

type RunRoutineCallback[T any] struct {
	Fn  func(arg T, this js.Ref, response *RunRoutineResponse) js.Ref
	Arg T
}

func (cb *RunRoutineCallback[T]) Register() js.Func[func(response *RunRoutineResponse)] {
	return js.RegisterCallback[func(response *RunRoutineResponse)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RunRoutineCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RunRoutineResponse
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

type RunRoutineResponse struct {
	// Id is "RunRoutineResponse.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id int32
	// Status is "RunRoutineResponse.status"
	//
	// Optional
	Status RoutineStatus

	FFI_USE_Id bool // for Id.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RunRoutineResponse with all fields set.
func (p RunRoutineResponse) FromRef(ref js.Ref) RunRoutineResponse {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RunRoutineResponse in the application heap.
func (p RunRoutineResponse) New() js.Ref {
	return bindings.RunRoutineResponseJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RunRoutineResponse) UpdateFrom(ref js.Ref) {
	bindings.RunRoutineResponseJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RunRoutineResponse) Update(ref js.Ref) {
	bindings.RunRoutineResponseJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RunRoutineResponse) FreeMembers(recursive bool) {
}

type RunSmartctlCheckRequest struct {
	// PercentageUsedThreshold is "RunSmartctlCheckRequest.percentage_used_threshold"
	//
	// Optional
	//
	// NOTE: FFI_USE_PercentageUsedThreshold MUST be set to true to make this field effective.
	PercentageUsedThreshold int32

	FFI_USE_PercentageUsedThreshold bool // for PercentageUsedThreshold.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RunSmartctlCheckRequest with all fields set.
func (p RunSmartctlCheckRequest) FromRef(ref js.Ref) RunSmartctlCheckRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RunSmartctlCheckRequest in the application heap.
func (p RunSmartctlCheckRequest) New() js.Ref {
	return bindings.RunSmartctlCheckRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RunSmartctlCheckRequest) UpdateFrom(ref js.Ref) {
	bindings.RunSmartctlCheckRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RunSmartctlCheckRequest) Update(ref js.Ref) {
	bindings.RunSmartctlCheckRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RunSmartctlCheckRequest) FreeMembers(recursive bool) {
}

type StartRoutineCallbackFunc func(this js.Ref) js.Ref

func (fn StartRoutineCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn StartRoutineCallbackFunc) DispatchCallback(
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

type StartRoutineCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *StartRoutineCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *StartRoutineCallback[T]) DispatchCallback(
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

type StartRoutineRequest struct {
	// Uuid is "StartRoutineRequest.uuid"
	//
	// Optional
	Uuid js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StartRoutineRequest with all fields set.
func (p StartRoutineRequest) FromRef(ref js.Ref) StartRoutineRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new StartRoutineRequest in the application heap.
func (p StartRoutineRequest) New() js.Ref {
	return bindings.StartRoutineRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *StartRoutineRequest) UpdateFrom(ref js.Ref) {
	bindings.StartRoutineRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StartRoutineRequest) Update(ref js.Ref) {
	bindings.StartRoutineRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StartRoutineRequest) FreeMembers(recursive bool) {
	js.Free(
		p.Uuid.Ref(),
	)
	p.Uuid = p.Uuid.FromRef(js.Undefined)
}

// HasFuncCancelRoutine returns true if the function "WEBEXT.os.diagnostics.cancelRoutine" exists.
func HasFuncCancelRoutine() bool {
	return js.True == bindings.HasFuncCancelRoutine()
}

// FuncCancelRoutine returns the function "WEBEXT.os.diagnostics.cancelRoutine".
func FuncCancelRoutine() (fn js.Func[func(request CancelRoutineRequest) js.Promise[js.Void]]) {
	bindings.FuncCancelRoutine(
		js.Pointer(&fn),
	)
	return
}

// CancelRoutine calls the function "WEBEXT.os.diagnostics.cancelRoutine" directly.
func CancelRoutine(request CancelRoutineRequest) (ret js.Promise[js.Void]) {
	bindings.CallCancelRoutine(
		js.Pointer(&ret),
		js.Pointer(&request),
	)

	return
}

// TryCancelRoutine calls the function "WEBEXT.os.diagnostics.cancelRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCancelRoutine(request CancelRoutineRequest) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCancelRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
	)

	return
}

// HasFuncCreateMemoryRoutine returns true if the function "WEBEXT.os.diagnostics.createMemoryRoutine" exists.
func HasFuncCreateMemoryRoutine() bool {
	return js.True == bindings.HasFuncCreateMemoryRoutine()
}

// FuncCreateMemoryRoutine returns the function "WEBEXT.os.diagnostics.createMemoryRoutine".
func FuncCreateMemoryRoutine() (fn js.Func[func(args RunMemoryRoutineArguments) js.Promise[CreateMemoryRoutineResponse]]) {
	bindings.FuncCreateMemoryRoutine(
		js.Pointer(&fn),
	)
	return
}

// CreateMemoryRoutine calls the function "WEBEXT.os.diagnostics.createMemoryRoutine" directly.
func CreateMemoryRoutine(args RunMemoryRoutineArguments) (ret js.Promise[CreateMemoryRoutineResponse]) {
	bindings.CallCreateMemoryRoutine(
		js.Pointer(&ret),
		js.Pointer(&args),
	)

	return
}

// TryCreateMemoryRoutine calls the function "WEBEXT.os.diagnostics.createMemoryRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCreateMemoryRoutine(args RunMemoryRoutineArguments) (ret js.Promise[CreateMemoryRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCreateMemoryRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&args),
	)

	return
}

// HasFuncGetAvailableRoutines returns true if the function "WEBEXT.os.diagnostics.getAvailableRoutines" exists.
func HasFuncGetAvailableRoutines() bool {
	return js.True == bindings.HasFuncGetAvailableRoutines()
}

// FuncGetAvailableRoutines returns the function "WEBEXT.os.diagnostics.getAvailableRoutines".
func FuncGetAvailableRoutines() (fn js.Func[func() js.Promise[GetAvailableRoutinesResponse]]) {
	bindings.FuncGetAvailableRoutines(
		js.Pointer(&fn),
	)
	return
}

// GetAvailableRoutines calls the function "WEBEXT.os.diagnostics.getAvailableRoutines" directly.
func GetAvailableRoutines() (ret js.Promise[GetAvailableRoutinesResponse]) {
	bindings.CallGetAvailableRoutines(
		js.Pointer(&ret),
	)

	return
}

// TryGetAvailableRoutines calls the function "WEBEXT.os.diagnostics.getAvailableRoutines"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAvailableRoutines() (ret js.Promise[GetAvailableRoutinesResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAvailableRoutines(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetRoutineUpdate returns true if the function "WEBEXT.os.diagnostics.getRoutineUpdate" exists.
func HasFuncGetRoutineUpdate() bool {
	return js.True == bindings.HasFuncGetRoutineUpdate()
}

// FuncGetRoutineUpdate returns the function "WEBEXT.os.diagnostics.getRoutineUpdate".
func FuncGetRoutineUpdate() (fn js.Func[func(request GetRoutineUpdateRequest) js.Promise[GetRoutineUpdateResponse]]) {
	bindings.FuncGetRoutineUpdate(
		js.Pointer(&fn),
	)
	return
}

// GetRoutineUpdate calls the function "WEBEXT.os.diagnostics.getRoutineUpdate" directly.
func GetRoutineUpdate(request GetRoutineUpdateRequest) (ret js.Promise[GetRoutineUpdateResponse]) {
	bindings.CallGetRoutineUpdate(
		js.Pointer(&ret),
		js.Pointer(&request),
	)

	return
}

// TryGetRoutineUpdate calls the function "WEBEXT.os.diagnostics.getRoutineUpdate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetRoutineUpdate(request GetRoutineUpdateRequest) (ret js.Promise[GetRoutineUpdateResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetRoutineUpdate(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
	)

	return
}

// HasFuncIsMemoryRoutineArgumentSupported returns true if the function "WEBEXT.os.diagnostics.isMemoryRoutineArgumentSupported" exists.
func HasFuncIsMemoryRoutineArgumentSupported() bool {
	return js.True == bindings.HasFuncIsMemoryRoutineArgumentSupported()
}

// FuncIsMemoryRoutineArgumentSupported returns the function "WEBEXT.os.diagnostics.isMemoryRoutineArgumentSupported".
func FuncIsMemoryRoutineArgumentSupported() (fn js.Func[func(args RunMemoryRoutineArguments) js.Promise[RoutineSupportStatusInfo]]) {
	bindings.FuncIsMemoryRoutineArgumentSupported(
		js.Pointer(&fn),
	)
	return
}

// IsMemoryRoutineArgumentSupported calls the function "WEBEXT.os.diagnostics.isMemoryRoutineArgumentSupported" directly.
func IsMemoryRoutineArgumentSupported(args RunMemoryRoutineArguments) (ret js.Promise[RoutineSupportStatusInfo]) {
	bindings.CallIsMemoryRoutineArgumentSupported(
		js.Pointer(&ret),
		js.Pointer(&args),
	)

	return
}

// TryIsMemoryRoutineArgumentSupported calls the function "WEBEXT.os.diagnostics.isMemoryRoutineArgumentSupported"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsMemoryRoutineArgumentSupported(args RunMemoryRoutineArguments) (ret js.Promise[RoutineSupportStatusInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsMemoryRoutineArgumentSupported(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&args),
	)

	return
}

type OnMemoryRoutineFinishedEventCallbackFunc func(this js.Ref, finishedInfo *MemoryRoutineFinishedInfo) js.Ref

func (fn OnMemoryRoutineFinishedEventCallbackFunc) Register() js.Func[func(finishedInfo *MemoryRoutineFinishedInfo)] {
	return js.RegisterCallback[func(finishedInfo *MemoryRoutineFinishedInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnMemoryRoutineFinishedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 MemoryRoutineFinishedInfo
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

type OnMemoryRoutineFinishedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, finishedInfo *MemoryRoutineFinishedInfo) js.Ref
	Arg T
}

func (cb *OnMemoryRoutineFinishedEventCallback[T]) Register() js.Func[func(finishedInfo *MemoryRoutineFinishedInfo)] {
	return js.RegisterCallback[func(finishedInfo *MemoryRoutineFinishedInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnMemoryRoutineFinishedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 MemoryRoutineFinishedInfo
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

// HasFuncOnMemoryRoutineFinished returns true if the function "WEBEXT.os.diagnostics.onMemoryRoutineFinished.addListener" exists.
func HasFuncOnMemoryRoutineFinished() bool {
	return js.True == bindings.HasFuncOnMemoryRoutineFinished()
}

// FuncOnMemoryRoutineFinished returns the function "WEBEXT.os.diagnostics.onMemoryRoutineFinished.addListener".
func FuncOnMemoryRoutineFinished() (fn js.Func[func(callback js.Func[func(finishedInfo *MemoryRoutineFinishedInfo)])]) {
	bindings.FuncOnMemoryRoutineFinished(
		js.Pointer(&fn),
	)
	return
}

// OnMemoryRoutineFinished calls the function "WEBEXT.os.diagnostics.onMemoryRoutineFinished.addListener" directly.
func OnMemoryRoutineFinished(callback js.Func[func(finishedInfo *MemoryRoutineFinishedInfo)]) (ret js.Void) {
	bindings.CallOnMemoryRoutineFinished(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMemoryRoutineFinished calls the function "WEBEXT.os.diagnostics.onMemoryRoutineFinished.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMemoryRoutineFinished(callback js.Func[func(finishedInfo *MemoryRoutineFinishedInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMemoryRoutineFinished(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMemoryRoutineFinished returns true if the function "WEBEXT.os.diagnostics.onMemoryRoutineFinished.removeListener" exists.
func HasFuncOffMemoryRoutineFinished() bool {
	return js.True == bindings.HasFuncOffMemoryRoutineFinished()
}

// FuncOffMemoryRoutineFinished returns the function "WEBEXT.os.diagnostics.onMemoryRoutineFinished.removeListener".
func FuncOffMemoryRoutineFinished() (fn js.Func[func(callback js.Func[func(finishedInfo *MemoryRoutineFinishedInfo)])]) {
	bindings.FuncOffMemoryRoutineFinished(
		js.Pointer(&fn),
	)
	return
}

// OffMemoryRoutineFinished calls the function "WEBEXT.os.diagnostics.onMemoryRoutineFinished.removeListener" directly.
func OffMemoryRoutineFinished(callback js.Func[func(finishedInfo *MemoryRoutineFinishedInfo)]) (ret js.Void) {
	bindings.CallOffMemoryRoutineFinished(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMemoryRoutineFinished calls the function "WEBEXT.os.diagnostics.onMemoryRoutineFinished.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMemoryRoutineFinished(callback js.Func[func(finishedInfo *MemoryRoutineFinishedInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMemoryRoutineFinished(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMemoryRoutineFinished returns true if the function "WEBEXT.os.diagnostics.onMemoryRoutineFinished.hasListener" exists.
func HasFuncHasOnMemoryRoutineFinished() bool {
	return js.True == bindings.HasFuncHasOnMemoryRoutineFinished()
}

// FuncHasOnMemoryRoutineFinished returns the function "WEBEXT.os.diagnostics.onMemoryRoutineFinished.hasListener".
func FuncHasOnMemoryRoutineFinished() (fn js.Func[func(callback js.Func[func(finishedInfo *MemoryRoutineFinishedInfo)]) bool]) {
	bindings.FuncHasOnMemoryRoutineFinished(
		js.Pointer(&fn),
	)
	return
}

// HasOnMemoryRoutineFinished calls the function "WEBEXT.os.diagnostics.onMemoryRoutineFinished.hasListener" directly.
func HasOnMemoryRoutineFinished(callback js.Func[func(finishedInfo *MemoryRoutineFinishedInfo)]) (ret bool) {
	bindings.CallHasOnMemoryRoutineFinished(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMemoryRoutineFinished calls the function "WEBEXT.os.diagnostics.onMemoryRoutineFinished.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnMemoryRoutineFinished(callback js.Func[func(finishedInfo *MemoryRoutineFinishedInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnMemoryRoutineFinished(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnRoutineExceptionEventCallbackFunc func(this js.Ref, exceptionInfo *ExceptionInfo) js.Ref

func (fn OnRoutineExceptionEventCallbackFunc) Register() js.Func[func(exceptionInfo *ExceptionInfo)] {
	return js.RegisterCallback[func(exceptionInfo *ExceptionInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnRoutineExceptionEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ExceptionInfo
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

type OnRoutineExceptionEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, exceptionInfo *ExceptionInfo) js.Ref
	Arg T
}

func (cb *OnRoutineExceptionEventCallback[T]) Register() js.Func[func(exceptionInfo *ExceptionInfo)] {
	return js.RegisterCallback[func(exceptionInfo *ExceptionInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnRoutineExceptionEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ExceptionInfo
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

// HasFuncOnRoutineException returns true if the function "WEBEXT.os.diagnostics.onRoutineException.addListener" exists.
func HasFuncOnRoutineException() bool {
	return js.True == bindings.HasFuncOnRoutineException()
}

// FuncOnRoutineException returns the function "WEBEXT.os.diagnostics.onRoutineException.addListener".
func FuncOnRoutineException() (fn js.Func[func(callback js.Func[func(exceptionInfo *ExceptionInfo)])]) {
	bindings.FuncOnRoutineException(
		js.Pointer(&fn),
	)
	return
}

// OnRoutineException calls the function "WEBEXT.os.diagnostics.onRoutineException.addListener" directly.
func OnRoutineException(callback js.Func[func(exceptionInfo *ExceptionInfo)]) (ret js.Void) {
	bindings.CallOnRoutineException(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnRoutineException calls the function "WEBEXT.os.diagnostics.onRoutineException.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnRoutineException(callback js.Func[func(exceptionInfo *ExceptionInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnRoutineException(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffRoutineException returns true if the function "WEBEXT.os.diagnostics.onRoutineException.removeListener" exists.
func HasFuncOffRoutineException() bool {
	return js.True == bindings.HasFuncOffRoutineException()
}

// FuncOffRoutineException returns the function "WEBEXT.os.diagnostics.onRoutineException.removeListener".
func FuncOffRoutineException() (fn js.Func[func(callback js.Func[func(exceptionInfo *ExceptionInfo)])]) {
	bindings.FuncOffRoutineException(
		js.Pointer(&fn),
	)
	return
}

// OffRoutineException calls the function "WEBEXT.os.diagnostics.onRoutineException.removeListener" directly.
func OffRoutineException(callback js.Func[func(exceptionInfo *ExceptionInfo)]) (ret js.Void) {
	bindings.CallOffRoutineException(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffRoutineException calls the function "WEBEXT.os.diagnostics.onRoutineException.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffRoutineException(callback js.Func[func(exceptionInfo *ExceptionInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffRoutineException(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnRoutineException returns true if the function "WEBEXT.os.diagnostics.onRoutineException.hasListener" exists.
func HasFuncHasOnRoutineException() bool {
	return js.True == bindings.HasFuncHasOnRoutineException()
}

// FuncHasOnRoutineException returns the function "WEBEXT.os.diagnostics.onRoutineException.hasListener".
func FuncHasOnRoutineException() (fn js.Func[func(callback js.Func[func(exceptionInfo *ExceptionInfo)]) bool]) {
	bindings.FuncHasOnRoutineException(
		js.Pointer(&fn),
	)
	return
}

// HasOnRoutineException calls the function "WEBEXT.os.diagnostics.onRoutineException.hasListener" directly.
func HasOnRoutineException(callback js.Func[func(exceptionInfo *ExceptionInfo)]) (ret bool) {
	bindings.CallHasOnRoutineException(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnRoutineException calls the function "WEBEXT.os.diagnostics.onRoutineException.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnRoutineException(callback js.Func[func(exceptionInfo *ExceptionInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnRoutineException(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnRoutineInitializedEventCallbackFunc func(this js.Ref, initializedInfo *RoutineInitializedInfo) js.Ref

func (fn OnRoutineInitializedEventCallbackFunc) Register() js.Func[func(initializedInfo *RoutineInitializedInfo)] {
	return js.RegisterCallback[func(initializedInfo *RoutineInitializedInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnRoutineInitializedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RoutineInitializedInfo
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

type OnRoutineInitializedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, initializedInfo *RoutineInitializedInfo) js.Ref
	Arg T
}

func (cb *OnRoutineInitializedEventCallback[T]) Register() js.Func[func(initializedInfo *RoutineInitializedInfo)] {
	return js.RegisterCallback[func(initializedInfo *RoutineInitializedInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnRoutineInitializedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RoutineInitializedInfo
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

// HasFuncOnRoutineInitialized returns true if the function "WEBEXT.os.diagnostics.onRoutineInitialized.addListener" exists.
func HasFuncOnRoutineInitialized() bool {
	return js.True == bindings.HasFuncOnRoutineInitialized()
}

// FuncOnRoutineInitialized returns the function "WEBEXT.os.diagnostics.onRoutineInitialized.addListener".
func FuncOnRoutineInitialized() (fn js.Func[func(callback js.Func[func(initializedInfo *RoutineInitializedInfo)])]) {
	bindings.FuncOnRoutineInitialized(
		js.Pointer(&fn),
	)
	return
}

// OnRoutineInitialized calls the function "WEBEXT.os.diagnostics.onRoutineInitialized.addListener" directly.
func OnRoutineInitialized(callback js.Func[func(initializedInfo *RoutineInitializedInfo)]) (ret js.Void) {
	bindings.CallOnRoutineInitialized(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnRoutineInitialized calls the function "WEBEXT.os.diagnostics.onRoutineInitialized.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnRoutineInitialized(callback js.Func[func(initializedInfo *RoutineInitializedInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnRoutineInitialized(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffRoutineInitialized returns true if the function "WEBEXT.os.diagnostics.onRoutineInitialized.removeListener" exists.
func HasFuncOffRoutineInitialized() bool {
	return js.True == bindings.HasFuncOffRoutineInitialized()
}

// FuncOffRoutineInitialized returns the function "WEBEXT.os.diagnostics.onRoutineInitialized.removeListener".
func FuncOffRoutineInitialized() (fn js.Func[func(callback js.Func[func(initializedInfo *RoutineInitializedInfo)])]) {
	bindings.FuncOffRoutineInitialized(
		js.Pointer(&fn),
	)
	return
}

// OffRoutineInitialized calls the function "WEBEXT.os.diagnostics.onRoutineInitialized.removeListener" directly.
func OffRoutineInitialized(callback js.Func[func(initializedInfo *RoutineInitializedInfo)]) (ret js.Void) {
	bindings.CallOffRoutineInitialized(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffRoutineInitialized calls the function "WEBEXT.os.diagnostics.onRoutineInitialized.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffRoutineInitialized(callback js.Func[func(initializedInfo *RoutineInitializedInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffRoutineInitialized(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnRoutineInitialized returns true if the function "WEBEXT.os.diagnostics.onRoutineInitialized.hasListener" exists.
func HasFuncHasOnRoutineInitialized() bool {
	return js.True == bindings.HasFuncHasOnRoutineInitialized()
}

// FuncHasOnRoutineInitialized returns the function "WEBEXT.os.diagnostics.onRoutineInitialized.hasListener".
func FuncHasOnRoutineInitialized() (fn js.Func[func(callback js.Func[func(initializedInfo *RoutineInitializedInfo)]) bool]) {
	bindings.FuncHasOnRoutineInitialized(
		js.Pointer(&fn),
	)
	return
}

// HasOnRoutineInitialized calls the function "WEBEXT.os.diagnostics.onRoutineInitialized.hasListener" directly.
func HasOnRoutineInitialized(callback js.Func[func(initializedInfo *RoutineInitializedInfo)]) (ret bool) {
	bindings.CallHasOnRoutineInitialized(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnRoutineInitialized calls the function "WEBEXT.os.diagnostics.onRoutineInitialized.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnRoutineInitialized(callback js.Func[func(initializedInfo *RoutineInitializedInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnRoutineInitialized(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnRoutineRunningEventCallbackFunc func(this js.Ref, runningInfo *RoutineRunningInfo) js.Ref

func (fn OnRoutineRunningEventCallbackFunc) Register() js.Func[func(runningInfo *RoutineRunningInfo)] {
	return js.RegisterCallback[func(runningInfo *RoutineRunningInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnRoutineRunningEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RoutineRunningInfo
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

type OnRoutineRunningEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, runningInfo *RoutineRunningInfo) js.Ref
	Arg T
}

func (cb *OnRoutineRunningEventCallback[T]) Register() js.Func[func(runningInfo *RoutineRunningInfo)] {
	return js.RegisterCallback[func(runningInfo *RoutineRunningInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnRoutineRunningEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RoutineRunningInfo
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

// HasFuncOnRoutineRunning returns true if the function "WEBEXT.os.diagnostics.onRoutineRunning.addListener" exists.
func HasFuncOnRoutineRunning() bool {
	return js.True == bindings.HasFuncOnRoutineRunning()
}

// FuncOnRoutineRunning returns the function "WEBEXT.os.diagnostics.onRoutineRunning.addListener".
func FuncOnRoutineRunning() (fn js.Func[func(callback js.Func[func(runningInfo *RoutineRunningInfo)])]) {
	bindings.FuncOnRoutineRunning(
		js.Pointer(&fn),
	)
	return
}

// OnRoutineRunning calls the function "WEBEXT.os.diagnostics.onRoutineRunning.addListener" directly.
func OnRoutineRunning(callback js.Func[func(runningInfo *RoutineRunningInfo)]) (ret js.Void) {
	bindings.CallOnRoutineRunning(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnRoutineRunning calls the function "WEBEXT.os.diagnostics.onRoutineRunning.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnRoutineRunning(callback js.Func[func(runningInfo *RoutineRunningInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnRoutineRunning(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffRoutineRunning returns true if the function "WEBEXT.os.diagnostics.onRoutineRunning.removeListener" exists.
func HasFuncOffRoutineRunning() bool {
	return js.True == bindings.HasFuncOffRoutineRunning()
}

// FuncOffRoutineRunning returns the function "WEBEXT.os.diagnostics.onRoutineRunning.removeListener".
func FuncOffRoutineRunning() (fn js.Func[func(callback js.Func[func(runningInfo *RoutineRunningInfo)])]) {
	bindings.FuncOffRoutineRunning(
		js.Pointer(&fn),
	)
	return
}

// OffRoutineRunning calls the function "WEBEXT.os.diagnostics.onRoutineRunning.removeListener" directly.
func OffRoutineRunning(callback js.Func[func(runningInfo *RoutineRunningInfo)]) (ret js.Void) {
	bindings.CallOffRoutineRunning(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffRoutineRunning calls the function "WEBEXT.os.diagnostics.onRoutineRunning.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffRoutineRunning(callback js.Func[func(runningInfo *RoutineRunningInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffRoutineRunning(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnRoutineRunning returns true if the function "WEBEXT.os.diagnostics.onRoutineRunning.hasListener" exists.
func HasFuncHasOnRoutineRunning() bool {
	return js.True == bindings.HasFuncHasOnRoutineRunning()
}

// FuncHasOnRoutineRunning returns the function "WEBEXT.os.diagnostics.onRoutineRunning.hasListener".
func FuncHasOnRoutineRunning() (fn js.Func[func(callback js.Func[func(runningInfo *RoutineRunningInfo)]) bool]) {
	bindings.FuncHasOnRoutineRunning(
		js.Pointer(&fn),
	)
	return
}

// HasOnRoutineRunning calls the function "WEBEXT.os.diagnostics.onRoutineRunning.hasListener" directly.
func HasOnRoutineRunning(callback js.Func[func(runningInfo *RoutineRunningInfo)]) (ret bool) {
	bindings.CallHasOnRoutineRunning(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnRoutineRunning calls the function "WEBEXT.os.diagnostics.onRoutineRunning.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnRoutineRunning(callback js.Func[func(runningInfo *RoutineRunningInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnRoutineRunning(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnRoutineWaitingEventCallbackFunc func(this js.Ref, waitingInfo *RoutineWaitingInfo) js.Ref

func (fn OnRoutineWaitingEventCallbackFunc) Register() js.Func[func(waitingInfo *RoutineWaitingInfo)] {
	return js.RegisterCallback[func(waitingInfo *RoutineWaitingInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnRoutineWaitingEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RoutineWaitingInfo
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

type OnRoutineWaitingEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, waitingInfo *RoutineWaitingInfo) js.Ref
	Arg T
}

func (cb *OnRoutineWaitingEventCallback[T]) Register() js.Func[func(waitingInfo *RoutineWaitingInfo)] {
	return js.RegisterCallback[func(waitingInfo *RoutineWaitingInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnRoutineWaitingEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RoutineWaitingInfo
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

// HasFuncOnRoutineWaiting returns true if the function "WEBEXT.os.diagnostics.onRoutineWaiting.addListener" exists.
func HasFuncOnRoutineWaiting() bool {
	return js.True == bindings.HasFuncOnRoutineWaiting()
}

// FuncOnRoutineWaiting returns the function "WEBEXT.os.diagnostics.onRoutineWaiting.addListener".
func FuncOnRoutineWaiting() (fn js.Func[func(callback js.Func[func(waitingInfo *RoutineWaitingInfo)])]) {
	bindings.FuncOnRoutineWaiting(
		js.Pointer(&fn),
	)
	return
}

// OnRoutineWaiting calls the function "WEBEXT.os.diagnostics.onRoutineWaiting.addListener" directly.
func OnRoutineWaiting(callback js.Func[func(waitingInfo *RoutineWaitingInfo)]) (ret js.Void) {
	bindings.CallOnRoutineWaiting(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnRoutineWaiting calls the function "WEBEXT.os.diagnostics.onRoutineWaiting.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnRoutineWaiting(callback js.Func[func(waitingInfo *RoutineWaitingInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnRoutineWaiting(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffRoutineWaiting returns true if the function "WEBEXT.os.diagnostics.onRoutineWaiting.removeListener" exists.
func HasFuncOffRoutineWaiting() bool {
	return js.True == bindings.HasFuncOffRoutineWaiting()
}

// FuncOffRoutineWaiting returns the function "WEBEXT.os.diagnostics.onRoutineWaiting.removeListener".
func FuncOffRoutineWaiting() (fn js.Func[func(callback js.Func[func(waitingInfo *RoutineWaitingInfo)])]) {
	bindings.FuncOffRoutineWaiting(
		js.Pointer(&fn),
	)
	return
}

// OffRoutineWaiting calls the function "WEBEXT.os.diagnostics.onRoutineWaiting.removeListener" directly.
func OffRoutineWaiting(callback js.Func[func(waitingInfo *RoutineWaitingInfo)]) (ret js.Void) {
	bindings.CallOffRoutineWaiting(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffRoutineWaiting calls the function "WEBEXT.os.diagnostics.onRoutineWaiting.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffRoutineWaiting(callback js.Func[func(waitingInfo *RoutineWaitingInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffRoutineWaiting(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnRoutineWaiting returns true if the function "WEBEXT.os.diagnostics.onRoutineWaiting.hasListener" exists.
func HasFuncHasOnRoutineWaiting() bool {
	return js.True == bindings.HasFuncHasOnRoutineWaiting()
}

// FuncHasOnRoutineWaiting returns the function "WEBEXT.os.diagnostics.onRoutineWaiting.hasListener".
func FuncHasOnRoutineWaiting() (fn js.Func[func(callback js.Func[func(waitingInfo *RoutineWaitingInfo)]) bool]) {
	bindings.FuncHasOnRoutineWaiting(
		js.Pointer(&fn),
	)
	return
}

// HasOnRoutineWaiting calls the function "WEBEXT.os.diagnostics.onRoutineWaiting.hasListener" directly.
func HasOnRoutineWaiting(callback js.Func[func(waitingInfo *RoutineWaitingInfo)]) (ret bool) {
	bindings.CallHasOnRoutineWaiting(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnRoutineWaiting calls the function "WEBEXT.os.diagnostics.onRoutineWaiting.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnRoutineWaiting(callback js.Func[func(waitingInfo *RoutineWaitingInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnRoutineWaiting(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncRunAcPowerRoutine returns true if the function "WEBEXT.os.diagnostics.runAcPowerRoutine" exists.
func HasFuncRunAcPowerRoutine() bool {
	return js.True == bindings.HasFuncRunAcPowerRoutine()
}

// FuncRunAcPowerRoutine returns the function "WEBEXT.os.diagnostics.runAcPowerRoutine".
func FuncRunAcPowerRoutine() (fn js.Func[func(request RunAcPowerRoutineRequest) js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunAcPowerRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunAcPowerRoutine calls the function "WEBEXT.os.diagnostics.runAcPowerRoutine" directly.
func RunAcPowerRoutine(request RunAcPowerRoutineRequest) (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunAcPowerRoutine(
		js.Pointer(&ret),
		js.Pointer(&request),
	)

	return
}

// TryRunAcPowerRoutine calls the function "WEBEXT.os.diagnostics.runAcPowerRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunAcPowerRoutine(request RunAcPowerRoutineRequest) (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunAcPowerRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
	)

	return
}

// HasFuncRunAudioDriverRoutine returns true if the function "WEBEXT.os.diagnostics.runAudioDriverRoutine" exists.
func HasFuncRunAudioDriverRoutine() bool {
	return js.True == bindings.HasFuncRunAudioDriverRoutine()
}

// FuncRunAudioDriverRoutine returns the function "WEBEXT.os.diagnostics.runAudioDriverRoutine".
func FuncRunAudioDriverRoutine() (fn js.Func[func() js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunAudioDriverRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunAudioDriverRoutine calls the function "WEBEXT.os.diagnostics.runAudioDriverRoutine" directly.
func RunAudioDriverRoutine() (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunAudioDriverRoutine(
		js.Pointer(&ret),
	)

	return
}

// TryRunAudioDriverRoutine calls the function "WEBEXT.os.diagnostics.runAudioDriverRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunAudioDriverRoutine() (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunAudioDriverRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRunBatteryCapacityRoutine returns true if the function "WEBEXT.os.diagnostics.runBatteryCapacityRoutine" exists.
func HasFuncRunBatteryCapacityRoutine() bool {
	return js.True == bindings.HasFuncRunBatteryCapacityRoutine()
}

// FuncRunBatteryCapacityRoutine returns the function "WEBEXT.os.diagnostics.runBatteryCapacityRoutine".
func FuncRunBatteryCapacityRoutine() (fn js.Func[func() js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunBatteryCapacityRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunBatteryCapacityRoutine calls the function "WEBEXT.os.diagnostics.runBatteryCapacityRoutine" directly.
func RunBatteryCapacityRoutine() (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunBatteryCapacityRoutine(
		js.Pointer(&ret),
	)

	return
}

// TryRunBatteryCapacityRoutine calls the function "WEBEXT.os.diagnostics.runBatteryCapacityRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunBatteryCapacityRoutine() (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunBatteryCapacityRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRunBatteryChargeRoutine returns true if the function "WEBEXT.os.diagnostics.runBatteryChargeRoutine" exists.
func HasFuncRunBatteryChargeRoutine() bool {
	return js.True == bindings.HasFuncRunBatteryChargeRoutine()
}

// FuncRunBatteryChargeRoutine returns the function "WEBEXT.os.diagnostics.runBatteryChargeRoutine".
func FuncRunBatteryChargeRoutine() (fn js.Func[func(request RunBatteryChargeRoutineRequest) js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunBatteryChargeRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunBatteryChargeRoutine calls the function "WEBEXT.os.diagnostics.runBatteryChargeRoutine" directly.
func RunBatteryChargeRoutine(request RunBatteryChargeRoutineRequest) (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunBatteryChargeRoutine(
		js.Pointer(&ret),
		js.Pointer(&request),
	)

	return
}

// TryRunBatteryChargeRoutine calls the function "WEBEXT.os.diagnostics.runBatteryChargeRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunBatteryChargeRoutine(request RunBatteryChargeRoutineRequest) (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunBatteryChargeRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
	)

	return
}

// HasFuncRunBatteryDischargeRoutine returns true if the function "WEBEXT.os.diagnostics.runBatteryDischargeRoutine" exists.
func HasFuncRunBatteryDischargeRoutine() bool {
	return js.True == bindings.HasFuncRunBatteryDischargeRoutine()
}

// FuncRunBatteryDischargeRoutine returns the function "WEBEXT.os.diagnostics.runBatteryDischargeRoutine".
func FuncRunBatteryDischargeRoutine() (fn js.Func[func(request RunBatteryDischargeRoutineRequest) js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunBatteryDischargeRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunBatteryDischargeRoutine calls the function "WEBEXT.os.diagnostics.runBatteryDischargeRoutine" directly.
func RunBatteryDischargeRoutine(request RunBatteryDischargeRoutineRequest) (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunBatteryDischargeRoutine(
		js.Pointer(&ret),
		js.Pointer(&request),
	)

	return
}

// TryRunBatteryDischargeRoutine calls the function "WEBEXT.os.diagnostics.runBatteryDischargeRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunBatteryDischargeRoutine(request RunBatteryDischargeRoutineRequest) (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunBatteryDischargeRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
	)

	return
}

// HasFuncRunBatteryHealthRoutine returns true if the function "WEBEXT.os.diagnostics.runBatteryHealthRoutine" exists.
func HasFuncRunBatteryHealthRoutine() bool {
	return js.True == bindings.HasFuncRunBatteryHealthRoutine()
}

// FuncRunBatteryHealthRoutine returns the function "WEBEXT.os.diagnostics.runBatteryHealthRoutine".
func FuncRunBatteryHealthRoutine() (fn js.Func[func() js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunBatteryHealthRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunBatteryHealthRoutine calls the function "WEBEXT.os.diagnostics.runBatteryHealthRoutine" directly.
func RunBatteryHealthRoutine() (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunBatteryHealthRoutine(
		js.Pointer(&ret),
	)

	return
}

// TryRunBatteryHealthRoutine calls the function "WEBEXT.os.diagnostics.runBatteryHealthRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunBatteryHealthRoutine() (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunBatteryHealthRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRunBluetoothDiscoveryRoutine returns true if the function "WEBEXT.os.diagnostics.runBluetoothDiscoveryRoutine" exists.
func HasFuncRunBluetoothDiscoveryRoutine() bool {
	return js.True == bindings.HasFuncRunBluetoothDiscoveryRoutine()
}

// FuncRunBluetoothDiscoveryRoutine returns the function "WEBEXT.os.diagnostics.runBluetoothDiscoveryRoutine".
func FuncRunBluetoothDiscoveryRoutine() (fn js.Func[func() js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunBluetoothDiscoveryRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunBluetoothDiscoveryRoutine calls the function "WEBEXT.os.diagnostics.runBluetoothDiscoveryRoutine" directly.
func RunBluetoothDiscoveryRoutine() (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunBluetoothDiscoveryRoutine(
		js.Pointer(&ret),
	)

	return
}

// TryRunBluetoothDiscoveryRoutine calls the function "WEBEXT.os.diagnostics.runBluetoothDiscoveryRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunBluetoothDiscoveryRoutine() (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunBluetoothDiscoveryRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRunBluetoothPairingRoutine returns true if the function "WEBEXT.os.diagnostics.runBluetoothPairingRoutine" exists.
func HasFuncRunBluetoothPairingRoutine() bool {
	return js.True == bindings.HasFuncRunBluetoothPairingRoutine()
}

// FuncRunBluetoothPairingRoutine returns the function "WEBEXT.os.diagnostics.runBluetoothPairingRoutine".
func FuncRunBluetoothPairingRoutine() (fn js.Func[func(request RunBluetoothPairingRoutineRequest) js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunBluetoothPairingRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunBluetoothPairingRoutine calls the function "WEBEXT.os.diagnostics.runBluetoothPairingRoutine" directly.
func RunBluetoothPairingRoutine(request RunBluetoothPairingRoutineRequest) (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunBluetoothPairingRoutine(
		js.Pointer(&ret),
		js.Pointer(&request),
	)

	return
}

// TryRunBluetoothPairingRoutine calls the function "WEBEXT.os.diagnostics.runBluetoothPairingRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunBluetoothPairingRoutine(request RunBluetoothPairingRoutineRequest) (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunBluetoothPairingRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
	)

	return
}

// HasFuncRunBluetoothPowerRoutine returns true if the function "WEBEXT.os.diagnostics.runBluetoothPowerRoutine" exists.
func HasFuncRunBluetoothPowerRoutine() bool {
	return js.True == bindings.HasFuncRunBluetoothPowerRoutine()
}

// FuncRunBluetoothPowerRoutine returns the function "WEBEXT.os.diagnostics.runBluetoothPowerRoutine".
func FuncRunBluetoothPowerRoutine() (fn js.Func[func() js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunBluetoothPowerRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunBluetoothPowerRoutine calls the function "WEBEXT.os.diagnostics.runBluetoothPowerRoutine" directly.
func RunBluetoothPowerRoutine() (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunBluetoothPowerRoutine(
		js.Pointer(&ret),
	)

	return
}

// TryRunBluetoothPowerRoutine calls the function "WEBEXT.os.diagnostics.runBluetoothPowerRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunBluetoothPowerRoutine() (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunBluetoothPowerRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRunBluetoothScanningRoutine returns true if the function "WEBEXT.os.diagnostics.runBluetoothScanningRoutine" exists.
func HasFuncRunBluetoothScanningRoutine() bool {
	return js.True == bindings.HasFuncRunBluetoothScanningRoutine()
}

// FuncRunBluetoothScanningRoutine returns the function "WEBEXT.os.diagnostics.runBluetoothScanningRoutine".
func FuncRunBluetoothScanningRoutine() (fn js.Func[func(request RunBluetoothScanningRoutineRequest) js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunBluetoothScanningRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunBluetoothScanningRoutine calls the function "WEBEXT.os.diagnostics.runBluetoothScanningRoutine" directly.
func RunBluetoothScanningRoutine(request RunBluetoothScanningRoutineRequest) (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunBluetoothScanningRoutine(
		js.Pointer(&ret),
		js.Pointer(&request),
	)

	return
}

// TryRunBluetoothScanningRoutine calls the function "WEBEXT.os.diagnostics.runBluetoothScanningRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunBluetoothScanningRoutine(request RunBluetoothScanningRoutineRequest) (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunBluetoothScanningRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
	)

	return
}

// HasFuncRunCpuCacheRoutine returns true if the function "WEBEXT.os.diagnostics.runCpuCacheRoutine" exists.
func HasFuncRunCpuCacheRoutine() bool {
	return js.True == bindings.HasFuncRunCpuCacheRoutine()
}

// FuncRunCpuCacheRoutine returns the function "WEBEXT.os.diagnostics.runCpuCacheRoutine".
func FuncRunCpuCacheRoutine() (fn js.Func[func(request RunCpuRoutineRequest) js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunCpuCacheRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunCpuCacheRoutine calls the function "WEBEXT.os.diagnostics.runCpuCacheRoutine" directly.
func RunCpuCacheRoutine(request RunCpuRoutineRequest) (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunCpuCacheRoutine(
		js.Pointer(&ret),
		js.Pointer(&request),
	)

	return
}

// TryRunCpuCacheRoutine calls the function "WEBEXT.os.diagnostics.runCpuCacheRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunCpuCacheRoutine(request RunCpuRoutineRequest) (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunCpuCacheRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
	)

	return
}

// HasFuncRunCpuFloatingPointAccuracyRoutine returns true if the function "WEBEXT.os.diagnostics.runCpuFloatingPointAccuracyRoutine" exists.
func HasFuncRunCpuFloatingPointAccuracyRoutine() bool {
	return js.True == bindings.HasFuncRunCpuFloatingPointAccuracyRoutine()
}

// FuncRunCpuFloatingPointAccuracyRoutine returns the function "WEBEXT.os.diagnostics.runCpuFloatingPointAccuracyRoutine".
func FuncRunCpuFloatingPointAccuracyRoutine() (fn js.Func[func(request RunCpuRoutineRequest) js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunCpuFloatingPointAccuracyRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunCpuFloatingPointAccuracyRoutine calls the function "WEBEXT.os.diagnostics.runCpuFloatingPointAccuracyRoutine" directly.
func RunCpuFloatingPointAccuracyRoutine(request RunCpuRoutineRequest) (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunCpuFloatingPointAccuracyRoutine(
		js.Pointer(&ret),
		js.Pointer(&request),
	)

	return
}

// TryRunCpuFloatingPointAccuracyRoutine calls the function "WEBEXT.os.diagnostics.runCpuFloatingPointAccuracyRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunCpuFloatingPointAccuracyRoutine(request RunCpuRoutineRequest) (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunCpuFloatingPointAccuracyRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
	)

	return
}

// HasFuncRunCpuPrimeSearchRoutine returns true if the function "WEBEXT.os.diagnostics.runCpuPrimeSearchRoutine" exists.
func HasFuncRunCpuPrimeSearchRoutine() bool {
	return js.True == bindings.HasFuncRunCpuPrimeSearchRoutine()
}

// FuncRunCpuPrimeSearchRoutine returns the function "WEBEXT.os.diagnostics.runCpuPrimeSearchRoutine".
func FuncRunCpuPrimeSearchRoutine() (fn js.Func[func(request RunCpuRoutineRequest) js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunCpuPrimeSearchRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunCpuPrimeSearchRoutine calls the function "WEBEXT.os.diagnostics.runCpuPrimeSearchRoutine" directly.
func RunCpuPrimeSearchRoutine(request RunCpuRoutineRequest) (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunCpuPrimeSearchRoutine(
		js.Pointer(&ret),
		js.Pointer(&request),
	)

	return
}

// TryRunCpuPrimeSearchRoutine calls the function "WEBEXT.os.diagnostics.runCpuPrimeSearchRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunCpuPrimeSearchRoutine(request RunCpuRoutineRequest) (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunCpuPrimeSearchRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
	)

	return
}

// HasFuncRunCpuStressRoutine returns true if the function "WEBEXT.os.diagnostics.runCpuStressRoutine" exists.
func HasFuncRunCpuStressRoutine() bool {
	return js.True == bindings.HasFuncRunCpuStressRoutine()
}

// FuncRunCpuStressRoutine returns the function "WEBEXT.os.diagnostics.runCpuStressRoutine".
func FuncRunCpuStressRoutine() (fn js.Func[func(request RunCpuRoutineRequest) js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunCpuStressRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunCpuStressRoutine calls the function "WEBEXT.os.diagnostics.runCpuStressRoutine" directly.
func RunCpuStressRoutine(request RunCpuRoutineRequest) (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunCpuStressRoutine(
		js.Pointer(&ret),
		js.Pointer(&request),
	)

	return
}

// TryRunCpuStressRoutine calls the function "WEBEXT.os.diagnostics.runCpuStressRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunCpuStressRoutine(request RunCpuRoutineRequest) (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunCpuStressRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
	)

	return
}

// HasFuncRunDiskReadRoutine returns true if the function "WEBEXT.os.diagnostics.runDiskReadRoutine" exists.
func HasFuncRunDiskReadRoutine() bool {
	return js.True == bindings.HasFuncRunDiskReadRoutine()
}

// FuncRunDiskReadRoutine returns the function "WEBEXT.os.diagnostics.runDiskReadRoutine".
func FuncRunDiskReadRoutine() (fn js.Func[func(request RunDiskReadRequest) js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunDiskReadRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunDiskReadRoutine calls the function "WEBEXT.os.diagnostics.runDiskReadRoutine" directly.
func RunDiskReadRoutine(request RunDiskReadRequest) (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunDiskReadRoutine(
		js.Pointer(&ret),
		js.Pointer(&request),
	)

	return
}

// TryRunDiskReadRoutine calls the function "WEBEXT.os.diagnostics.runDiskReadRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunDiskReadRoutine(request RunDiskReadRequest) (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunDiskReadRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
	)

	return
}

// HasFuncRunDnsResolutionRoutine returns true if the function "WEBEXT.os.diagnostics.runDnsResolutionRoutine" exists.
func HasFuncRunDnsResolutionRoutine() bool {
	return js.True == bindings.HasFuncRunDnsResolutionRoutine()
}

// FuncRunDnsResolutionRoutine returns the function "WEBEXT.os.diagnostics.runDnsResolutionRoutine".
func FuncRunDnsResolutionRoutine() (fn js.Func[func() js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunDnsResolutionRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunDnsResolutionRoutine calls the function "WEBEXT.os.diagnostics.runDnsResolutionRoutine" directly.
func RunDnsResolutionRoutine() (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunDnsResolutionRoutine(
		js.Pointer(&ret),
	)

	return
}

// TryRunDnsResolutionRoutine calls the function "WEBEXT.os.diagnostics.runDnsResolutionRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunDnsResolutionRoutine() (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunDnsResolutionRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRunDnsResolverPresentRoutine returns true if the function "WEBEXT.os.diagnostics.runDnsResolverPresentRoutine" exists.
func HasFuncRunDnsResolverPresentRoutine() bool {
	return js.True == bindings.HasFuncRunDnsResolverPresentRoutine()
}

// FuncRunDnsResolverPresentRoutine returns the function "WEBEXT.os.diagnostics.runDnsResolverPresentRoutine".
func FuncRunDnsResolverPresentRoutine() (fn js.Func[func() js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunDnsResolverPresentRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunDnsResolverPresentRoutine calls the function "WEBEXT.os.diagnostics.runDnsResolverPresentRoutine" directly.
func RunDnsResolverPresentRoutine() (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunDnsResolverPresentRoutine(
		js.Pointer(&ret),
	)

	return
}

// TryRunDnsResolverPresentRoutine calls the function "WEBEXT.os.diagnostics.runDnsResolverPresentRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunDnsResolverPresentRoutine() (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunDnsResolverPresentRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRunEmmcLifetimeRoutine returns true if the function "WEBEXT.os.diagnostics.runEmmcLifetimeRoutine" exists.
func HasFuncRunEmmcLifetimeRoutine() bool {
	return js.True == bindings.HasFuncRunEmmcLifetimeRoutine()
}

// FuncRunEmmcLifetimeRoutine returns the function "WEBEXT.os.diagnostics.runEmmcLifetimeRoutine".
func FuncRunEmmcLifetimeRoutine() (fn js.Func[func() js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunEmmcLifetimeRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunEmmcLifetimeRoutine calls the function "WEBEXT.os.diagnostics.runEmmcLifetimeRoutine" directly.
func RunEmmcLifetimeRoutine() (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunEmmcLifetimeRoutine(
		js.Pointer(&ret),
	)

	return
}

// TryRunEmmcLifetimeRoutine calls the function "WEBEXT.os.diagnostics.runEmmcLifetimeRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunEmmcLifetimeRoutine() (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunEmmcLifetimeRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRunFingerprintAliveRoutine returns true if the function "WEBEXT.os.diagnostics.runFingerprintAliveRoutine" exists.
func HasFuncRunFingerprintAliveRoutine() bool {
	return js.True == bindings.HasFuncRunFingerprintAliveRoutine()
}

// FuncRunFingerprintAliveRoutine returns the function "WEBEXT.os.diagnostics.runFingerprintAliveRoutine".
func FuncRunFingerprintAliveRoutine() (fn js.Func[func() js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunFingerprintAliveRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunFingerprintAliveRoutine calls the function "WEBEXT.os.diagnostics.runFingerprintAliveRoutine" directly.
func RunFingerprintAliveRoutine() (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunFingerprintAliveRoutine(
		js.Pointer(&ret),
	)

	return
}

// TryRunFingerprintAliveRoutine calls the function "WEBEXT.os.diagnostics.runFingerprintAliveRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunFingerprintAliveRoutine() (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunFingerprintAliveRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRunGatewayCanBePingedRoutine returns true if the function "WEBEXT.os.diagnostics.runGatewayCanBePingedRoutine" exists.
func HasFuncRunGatewayCanBePingedRoutine() bool {
	return js.True == bindings.HasFuncRunGatewayCanBePingedRoutine()
}

// FuncRunGatewayCanBePingedRoutine returns the function "WEBEXT.os.diagnostics.runGatewayCanBePingedRoutine".
func FuncRunGatewayCanBePingedRoutine() (fn js.Func[func() js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunGatewayCanBePingedRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunGatewayCanBePingedRoutine calls the function "WEBEXT.os.diagnostics.runGatewayCanBePingedRoutine" directly.
func RunGatewayCanBePingedRoutine() (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunGatewayCanBePingedRoutine(
		js.Pointer(&ret),
	)

	return
}

// TryRunGatewayCanBePingedRoutine calls the function "WEBEXT.os.diagnostics.runGatewayCanBePingedRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunGatewayCanBePingedRoutine() (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunGatewayCanBePingedRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRunLanConnectivityRoutine returns true if the function "WEBEXT.os.diagnostics.runLanConnectivityRoutine" exists.
func HasFuncRunLanConnectivityRoutine() bool {
	return js.True == bindings.HasFuncRunLanConnectivityRoutine()
}

// FuncRunLanConnectivityRoutine returns the function "WEBEXT.os.diagnostics.runLanConnectivityRoutine".
func FuncRunLanConnectivityRoutine() (fn js.Func[func() js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunLanConnectivityRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunLanConnectivityRoutine calls the function "WEBEXT.os.diagnostics.runLanConnectivityRoutine" directly.
func RunLanConnectivityRoutine() (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunLanConnectivityRoutine(
		js.Pointer(&ret),
	)

	return
}

// TryRunLanConnectivityRoutine calls the function "WEBEXT.os.diagnostics.runLanConnectivityRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunLanConnectivityRoutine() (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunLanConnectivityRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRunMemoryRoutine returns true if the function "WEBEXT.os.diagnostics.runMemoryRoutine" exists.
func HasFuncRunMemoryRoutine() bool {
	return js.True == bindings.HasFuncRunMemoryRoutine()
}

// FuncRunMemoryRoutine returns the function "WEBEXT.os.diagnostics.runMemoryRoutine".
func FuncRunMemoryRoutine() (fn js.Func[func() js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunMemoryRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunMemoryRoutine calls the function "WEBEXT.os.diagnostics.runMemoryRoutine" directly.
func RunMemoryRoutine() (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunMemoryRoutine(
		js.Pointer(&ret),
	)

	return
}

// TryRunMemoryRoutine calls the function "WEBEXT.os.diagnostics.runMemoryRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunMemoryRoutine() (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunMemoryRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRunNvmeSelfTestRoutine returns true if the function "WEBEXT.os.diagnostics.runNvmeSelfTestRoutine" exists.
func HasFuncRunNvmeSelfTestRoutine() bool {
	return js.True == bindings.HasFuncRunNvmeSelfTestRoutine()
}

// FuncRunNvmeSelfTestRoutine returns the function "WEBEXT.os.diagnostics.runNvmeSelfTestRoutine".
func FuncRunNvmeSelfTestRoutine() (fn js.Func[func(request RunNvmeSelfTestRequest) js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunNvmeSelfTestRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunNvmeSelfTestRoutine calls the function "WEBEXT.os.diagnostics.runNvmeSelfTestRoutine" directly.
func RunNvmeSelfTestRoutine(request RunNvmeSelfTestRequest) (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunNvmeSelfTestRoutine(
		js.Pointer(&ret),
		js.Pointer(&request),
	)

	return
}

// TryRunNvmeSelfTestRoutine calls the function "WEBEXT.os.diagnostics.runNvmeSelfTestRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunNvmeSelfTestRoutine(request RunNvmeSelfTestRequest) (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunNvmeSelfTestRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
	)

	return
}

// HasFuncRunNvmeWearLevelRoutine returns true if the function "WEBEXT.os.diagnostics.runNvmeWearLevelRoutine" exists.
func HasFuncRunNvmeWearLevelRoutine() bool {
	return js.True == bindings.HasFuncRunNvmeWearLevelRoutine()
}

// FuncRunNvmeWearLevelRoutine returns the function "WEBEXT.os.diagnostics.runNvmeWearLevelRoutine".
func FuncRunNvmeWearLevelRoutine() (fn js.Func[func(request RunNvmeWearLevelRequest) js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunNvmeWearLevelRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunNvmeWearLevelRoutine calls the function "WEBEXT.os.diagnostics.runNvmeWearLevelRoutine" directly.
func RunNvmeWearLevelRoutine(request RunNvmeWearLevelRequest) (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunNvmeWearLevelRoutine(
		js.Pointer(&ret),
		js.Pointer(&request),
	)

	return
}

// TryRunNvmeWearLevelRoutine calls the function "WEBEXT.os.diagnostics.runNvmeWearLevelRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunNvmeWearLevelRoutine(request RunNvmeWearLevelRequest) (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunNvmeWearLevelRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
	)

	return
}

// HasFuncRunPowerButtonRoutine returns true if the function "WEBEXT.os.diagnostics.runPowerButtonRoutine" exists.
func HasFuncRunPowerButtonRoutine() bool {
	return js.True == bindings.HasFuncRunPowerButtonRoutine()
}

// FuncRunPowerButtonRoutine returns the function "WEBEXT.os.diagnostics.runPowerButtonRoutine".
func FuncRunPowerButtonRoutine() (fn js.Func[func(request RunPowerButtonRequest) js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunPowerButtonRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunPowerButtonRoutine calls the function "WEBEXT.os.diagnostics.runPowerButtonRoutine" directly.
func RunPowerButtonRoutine(request RunPowerButtonRequest) (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunPowerButtonRoutine(
		js.Pointer(&ret),
		js.Pointer(&request),
	)

	return
}

// TryRunPowerButtonRoutine calls the function "WEBEXT.os.diagnostics.runPowerButtonRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunPowerButtonRoutine(request RunPowerButtonRequest) (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunPowerButtonRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
	)

	return
}

// HasFuncRunSensitiveSensorRoutine returns true if the function "WEBEXT.os.diagnostics.runSensitiveSensorRoutine" exists.
func HasFuncRunSensitiveSensorRoutine() bool {
	return js.True == bindings.HasFuncRunSensitiveSensorRoutine()
}

// FuncRunSensitiveSensorRoutine returns the function "WEBEXT.os.diagnostics.runSensitiveSensorRoutine".
func FuncRunSensitiveSensorRoutine() (fn js.Func[func() js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunSensitiveSensorRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunSensitiveSensorRoutine calls the function "WEBEXT.os.diagnostics.runSensitiveSensorRoutine" directly.
func RunSensitiveSensorRoutine() (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunSensitiveSensorRoutine(
		js.Pointer(&ret),
	)

	return
}

// TryRunSensitiveSensorRoutine calls the function "WEBEXT.os.diagnostics.runSensitiveSensorRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunSensitiveSensorRoutine() (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunSensitiveSensorRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRunSignalStrengthRoutine returns true if the function "WEBEXT.os.diagnostics.runSignalStrengthRoutine" exists.
func HasFuncRunSignalStrengthRoutine() bool {
	return js.True == bindings.HasFuncRunSignalStrengthRoutine()
}

// FuncRunSignalStrengthRoutine returns the function "WEBEXT.os.diagnostics.runSignalStrengthRoutine".
func FuncRunSignalStrengthRoutine() (fn js.Func[func() js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunSignalStrengthRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunSignalStrengthRoutine calls the function "WEBEXT.os.diagnostics.runSignalStrengthRoutine" directly.
func RunSignalStrengthRoutine() (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunSignalStrengthRoutine(
		js.Pointer(&ret),
	)

	return
}

// TryRunSignalStrengthRoutine calls the function "WEBEXT.os.diagnostics.runSignalStrengthRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunSignalStrengthRoutine() (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunSignalStrengthRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRunSmartctlCheckRoutine returns true if the function "WEBEXT.os.diagnostics.runSmartctlCheckRoutine" exists.
func HasFuncRunSmartctlCheckRoutine() bool {
	return js.True == bindings.HasFuncRunSmartctlCheckRoutine()
}

// FuncRunSmartctlCheckRoutine returns the function "WEBEXT.os.diagnostics.runSmartctlCheckRoutine".
func FuncRunSmartctlCheckRoutine() (fn js.Func[func(request RunSmartctlCheckRequest) js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunSmartctlCheckRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunSmartctlCheckRoutine calls the function "WEBEXT.os.diagnostics.runSmartctlCheckRoutine" directly.
func RunSmartctlCheckRoutine(request RunSmartctlCheckRequest) (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunSmartctlCheckRoutine(
		js.Pointer(&ret),
		js.Pointer(&request),
	)

	return
}

// TryRunSmartctlCheckRoutine calls the function "WEBEXT.os.diagnostics.runSmartctlCheckRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunSmartctlCheckRoutine(request RunSmartctlCheckRequest) (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunSmartctlCheckRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
	)

	return
}

// HasFuncRunUfsLifetimeRoutine returns true if the function "WEBEXT.os.diagnostics.runUfsLifetimeRoutine" exists.
func HasFuncRunUfsLifetimeRoutine() bool {
	return js.True == bindings.HasFuncRunUfsLifetimeRoutine()
}

// FuncRunUfsLifetimeRoutine returns the function "WEBEXT.os.diagnostics.runUfsLifetimeRoutine".
func FuncRunUfsLifetimeRoutine() (fn js.Func[func() js.Promise[RunRoutineResponse]]) {
	bindings.FuncRunUfsLifetimeRoutine(
		js.Pointer(&fn),
	)
	return
}

// RunUfsLifetimeRoutine calls the function "WEBEXT.os.diagnostics.runUfsLifetimeRoutine" directly.
func RunUfsLifetimeRoutine() (ret js.Promise[RunRoutineResponse]) {
	bindings.CallRunUfsLifetimeRoutine(
		js.Pointer(&ret),
	)

	return
}

// TryRunUfsLifetimeRoutine calls the function "WEBEXT.os.diagnostics.runUfsLifetimeRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunUfsLifetimeRoutine() (ret js.Promise[RunRoutineResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunUfsLifetimeRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStartRoutine returns true if the function "WEBEXT.os.diagnostics.startRoutine" exists.
func HasFuncStartRoutine() bool {
	return js.True == bindings.HasFuncStartRoutine()
}

// FuncStartRoutine returns the function "WEBEXT.os.diagnostics.startRoutine".
func FuncStartRoutine() (fn js.Func[func(request StartRoutineRequest) js.Promise[js.Void]]) {
	bindings.FuncStartRoutine(
		js.Pointer(&fn),
	)
	return
}

// StartRoutine calls the function "WEBEXT.os.diagnostics.startRoutine" directly.
func StartRoutine(request StartRoutineRequest) (ret js.Promise[js.Void]) {
	bindings.CallStartRoutine(
		js.Pointer(&ret),
		js.Pointer(&request),
	)

	return
}

// TryStartRoutine calls the function "WEBEXT.os.diagnostics.startRoutine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStartRoutine(request StartRoutineRequest) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStartRoutine(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
	)

	return
}
