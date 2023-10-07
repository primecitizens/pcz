// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package sidepanel

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/sidepanel/bindings"
)

type GetPanelOptions struct {
	// TabId is "GetPanelOptions.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int32

	FFI_USE_TabId bool // for TabId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetPanelOptions with all fields set.
func (p GetPanelOptions) FromRef(ref js.Ref) GetPanelOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetPanelOptions in the application heap.
func (p GetPanelOptions) New() js.Ref {
	return bindings.GetPanelOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetPanelOptions) UpdateFrom(ref js.Ref) {
	bindings.GetPanelOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetPanelOptions) Update(ref js.Ref) {
	bindings.GetPanelOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetPanelOptions) FreeMembers(recursive bool) {
}

type SidePanel struct {
	// DefaultPath is "SidePanel.default_path"
	//
	// Optional
	DefaultPath js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SidePanel with all fields set.
func (p SidePanel) FromRef(ref js.Ref) SidePanel {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SidePanel in the application heap.
func (p SidePanel) New() js.Ref {
	return bindings.SidePanelJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SidePanel) UpdateFrom(ref js.Ref) {
	bindings.SidePanelJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SidePanel) Update(ref js.Ref) {
	bindings.SidePanelJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SidePanel) FreeMembers(recursive bool) {
	js.Free(
		p.DefaultPath.Ref(),
	)
	p.DefaultPath = p.DefaultPath.FromRef(js.Undefined)
}

type ManifestKeys struct {
	// SidePanel is "ManifestKeys.side_panel"
	//
	// Optional
	//
	// NOTE: SidePanel.FFI_USE MUST be set to true to get SidePanel used.
	SidePanel SidePanel

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ManifestKeys with all fields set.
func (p ManifestKeys) FromRef(ref js.Ref) ManifestKeys {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ManifestKeys in the application heap.
func (p ManifestKeys) New() js.Ref {
	return bindings.ManifestKeysJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ManifestKeys) UpdateFrom(ref js.Ref) {
	bindings.ManifestKeysJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ManifestKeys) Update(ref js.Ref) {
	bindings.ManifestKeysJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ManifestKeys) FreeMembers(recursive bool) {
	if recursive {
		p.SidePanel.FreeMembers(true)
	}
}

type OpenOptions struct {
	// WindowId is "OpenOptions.windowId"
	//
	// Optional
	//
	// NOTE: FFI_USE_WindowId MUST be set to true to make this field effective.
	WindowId int32
	// TabId is "OpenOptions.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int32

	FFI_USE_WindowId bool // for WindowId.
	FFI_USE_TabId    bool // for TabId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OpenOptions with all fields set.
func (p OpenOptions) FromRef(ref js.Ref) OpenOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OpenOptions in the application heap.
func (p OpenOptions) New() js.Ref {
	return bindings.OpenOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OpenOptions) UpdateFrom(ref js.Ref) {
	bindings.OpenOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OpenOptions) Update(ref js.Ref) {
	bindings.OpenOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OpenOptions) FreeMembers(recursive bool) {
}

type PanelBehavior struct {
	// OpenPanelOnActionClick is "PanelBehavior.openPanelOnActionClick"
	//
	// Optional
	//
	// NOTE: FFI_USE_OpenPanelOnActionClick MUST be set to true to make this field effective.
	OpenPanelOnActionClick bool

	FFI_USE_OpenPanelOnActionClick bool // for OpenPanelOnActionClick.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PanelBehavior with all fields set.
func (p PanelBehavior) FromRef(ref js.Ref) PanelBehavior {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PanelBehavior in the application heap.
func (p PanelBehavior) New() js.Ref {
	return bindings.PanelBehaviorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PanelBehavior) UpdateFrom(ref js.Ref) {
	bindings.PanelBehaviorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PanelBehavior) Update(ref js.Ref) {
	bindings.PanelBehaviorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PanelBehavior) FreeMembers(recursive bool) {
}

type PanelBehaviorCallbackFunc func(this js.Ref, behavior *PanelBehavior) js.Ref

func (fn PanelBehaviorCallbackFunc) Register() js.Func[func(behavior *PanelBehavior)] {
	return js.RegisterCallback[func(behavior *PanelBehavior)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn PanelBehaviorCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PanelBehavior
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

type PanelBehaviorCallback[T any] struct {
	Fn  func(arg T, this js.Ref, behavior *PanelBehavior) js.Ref
	Arg T
}

func (cb *PanelBehaviorCallback[T]) Register() js.Func[func(behavior *PanelBehavior)] {
	return js.RegisterCallback[func(behavior *PanelBehavior)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *PanelBehaviorCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PanelBehavior
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

type PanelOptions struct {
	// TabId is "PanelOptions.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int32
	// Path is "PanelOptions.path"
	//
	// Optional
	Path js.String
	// Enabled is "PanelOptions.enabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_Enabled MUST be set to true to make this field effective.
	Enabled bool

	FFI_USE_TabId   bool // for TabId.
	FFI_USE_Enabled bool // for Enabled.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PanelOptions with all fields set.
func (p PanelOptions) FromRef(ref js.Ref) PanelOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PanelOptions in the application heap.
func (p PanelOptions) New() js.Ref {
	return bindings.PanelOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PanelOptions) UpdateFrom(ref js.Ref) {
	bindings.PanelOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PanelOptions) Update(ref js.Ref) {
	bindings.PanelOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PanelOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Path.Ref(),
	)
	p.Path = p.Path.FromRef(js.Undefined)
}

type PanelOptionsCallbackFunc func(this js.Ref, options *PanelOptions) js.Ref

func (fn PanelOptionsCallbackFunc) Register() js.Func[func(options *PanelOptions)] {
	return js.RegisterCallback[func(options *PanelOptions)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn PanelOptionsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PanelOptions
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

type PanelOptionsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, options *PanelOptions) js.Ref
	Arg T
}

func (cb *PanelOptionsCallback[T]) Register() js.Func[func(options *PanelOptions)] {
	return js.RegisterCallback[func(options *PanelOptions)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *PanelOptionsCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PanelOptions
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

// HasFuncGetOptions returns true if the function "WEBEXT.sidePanel.getOptions" exists.
func HasFuncGetOptions() bool {
	return js.True == bindings.HasFuncGetOptions()
}

// FuncGetOptions returns the function "WEBEXT.sidePanel.getOptions".
func FuncGetOptions() (fn js.Func[func(options GetPanelOptions) js.Promise[PanelOptions]]) {
	bindings.FuncGetOptions(
		js.Pointer(&fn),
	)
	return
}

// GetOptions calls the function "WEBEXT.sidePanel.getOptions" directly.
func GetOptions(options GetPanelOptions) (ret js.Promise[PanelOptions]) {
	bindings.CallGetOptions(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetOptions calls the function "WEBEXT.sidePanel.getOptions"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetOptions(options GetPanelOptions) (ret js.Promise[PanelOptions], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetOptions(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncGetPanelBehavior returns true if the function "WEBEXT.sidePanel.getPanelBehavior" exists.
func HasFuncGetPanelBehavior() bool {
	return js.True == bindings.HasFuncGetPanelBehavior()
}

// FuncGetPanelBehavior returns the function "WEBEXT.sidePanel.getPanelBehavior".
func FuncGetPanelBehavior() (fn js.Func[func() js.Promise[PanelBehavior]]) {
	bindings.FuncGetPanelBehavior(
		js.Pointer(&fn),
	)
	return
}

// GetPanelBehavior calls the function "WEBEXT.sidePanel.getPanelBehavior" directly.
func GetPanelBehavior() (ret js.Promise[PanelBehavior]) {
	bindings.CallGetPanelBehavior(
		js.Pointer(&ret),
	)

	return
}

// TryGetPanelBehavior calls the function "WEBEXT.sidePanel.getPanelBehavior"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetPanelBehavior() (ret js.Promise[PanelBehavior], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetPanelBehavior(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncOpen returns true if the function "WEBEXT.sidePanel.open" exists.
func HasFuncOpen() bool {
	return js.True == bindings.HasFuncOpen()
}

// FuncOpen returns the function "WEBEXT.sidePanel.open".
func FuncOpen() (fn js.Func[func(options OpenOptions) js.Promise[js.Void]]) {
	bindings.FuncOpen(
		js.Pointer(&fn),
	)
	return
}

// Open calls the function "WEBEXT.sidePanel.open" directly.
func Open(options OpenOptions) (ret js.Promise[js.Void]) {
	bindings.CallOpen(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryOpen calls the function "WEBEXT.sidePanel.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpen(options OpenOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpen(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncSetOptions returns true if the function "WEBEXT.sidePanel.setOptions" exists.
func HasFuncSetOptions() bool {
	return js.True == bindings.HasFuncSetOptions()
}

// FuncSetOptions returns the function "WEBEXT.sidePanel.setOptions".
func FuncSetOptions() (fn js.Func[func(options PanelOptions) js.Promise[js.Void]]) {
	bindings.FuncSetOptions(
		js.Pointer(&fn),
	)
	return
}

// SetOptions calls the function "WEBEXT.sidePanel.setOptions" directly.
func SetOptions(options PanelOptions) (ret js.Promise[js.Void]) {
	bindings.CallSetOptions(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TrySetOptions calls the function "WEBEXT.sidePanel.setOptions"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetOptions(options PanelOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetOptions(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncSetPanelBehavior returns true if the function "WEBEXT.sidePanel.setPanelBehavior" exists.
func HasFuncSetPanelBehavior() bool {
	return js.True == bindings.HasFuncSetPanelBehavior()
}

// FuncSetPanelBehavior returns the function "WEBEXT.sidePanel.setPanelBehavior".
func FuncSetPanelBehavior() (fn js.Func[func(behavior PanelBehavior) js.Promise[js.Void]]) {
	bindings.FuncSetPanelBehavior(
		js.Pointer(&fn),
	)
	return
}

// SetPanelBehavior calls the function "WEBEXT.sidePanel.setPanelBehavior" directly.
func SetPanelBehavior(behavior PanelBehavior) (ret js.Promise[js.Void]) {
	bindings.CallSetPanelBehavior(
		js.Pointer(&ret),
		js.Pointer(&behavior),
	)

	return
}

// TrySetPanelBehavior calls the function "WEBEXT.sidePanel.setPanelBehavior"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetPanelBehavior(behavior PanelBehavior) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetPanelBehavior(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&behavior),
	)

	return
}
