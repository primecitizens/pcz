// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package commands

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/commands/bindings"
	"github.com/primecitizens/pcz/std/plat/js/webext/tabs"
)

type Command struct {
	// Description is "Command.description"
	//
	// Optional
	Description js.String
	// Name is "Command.name"
	//
	// Optional
	Name js.String
	// Shortcut is "Command.shortcut"
	//
	// Optional
	Shortcut js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Command with all fields set.
func (p Command) FromRef(ref js.Ref) Command {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Command in the application heap.
func (p Command) New() js.Ref {
	return bindings.CommandJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Command) UpdateFrom(ref js.Ref) {
	bindings.CommandJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Command) Update(ref js.Ref) {
	bindings.CommandJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Command) FreeMembers(recursive bool) {
	js.Free(
		p.Description.Ref(),
		p.Name.Ref(),
		p.Shortcut.Ref(),
	)
	p.Description = p.Description.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Shortcut = p.Shortcut.FromRef(js.Undefined)
}

// HasFuncGetAll returns true if the function "WEBEXT.commands.getAll" exists.
func HasFuncGetAll() bool {
	return js.True == bindings.HasFuncGetAll()
}

// FuncGetAll returns the function "WEBEXT.commands.getAll".
func FuncGetAll() (fn js.Func[func() js.Promise[js.Array[Command]]]) {
	bindings.FuncGetAll(
		js.Pointer(&fn),
	)
	return
}

// GetAll calls the function "WEBEXT.commands.getAll" directly.
func GetAll() (ret js.Promise[js.Array[Command]]) {
	bindings.CallGetAll(
		js.Pointer(&ret),
	)

	return
}

// TryGetAll calls the function "WEBEXT.commands.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAll() (ret js.Promise[js.Array[Command]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAll(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnCommandEventCallbackFunc func(this js.Ref, command js.String, tab *tabs.Tab) js.Ref

func (fn OnCommandEventCallbackFunc) Register() js.Func[func(command js.String, tab *tabs.Tab)] {
	return js.RegisterCallback[func(command js.String, tab *tabs.Tab)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCommandEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 tabs.Tab
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

type OnCommandEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, command js.String, tab *tabs.Tab) js.Ref
	Arg T
}

func (cb *OnCommandEventCallback[T]) Register() js.Func[func(command js.String, tab *tabs.Tab)] {
	return js.RegisterCallback[func(command js.String, tab *tabs.Tab)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCommandEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 tabs.Tab
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

// HasFuncOnCommand returns true if the function "WEBEXT.commands.onCommand.addListener" exists.
func HasFuncOnCommand() bool {
	return js.True == bindings.HasFuncOnCommand()
}

// FuncOnCommand returns the function "WEBEXT.commands.onCommand.addListener".
func FuncOnCommand() (fn js.Func[func(callback js.Func[func(command js.String, tab *tabs.Tab)])]) {
	bindings.FuncOnCommand(
		js.Pointer(&fn),
	)
	return
}

// OnCommand calls the function "WEBEXT.commands.onCommand.addListener" directly.
func OnCommand(callback js.Func[func(command js.String, tab *tabs.Tab)]) (ret js.Void) {
	bindings.CallOnCommand(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCommand calls the function "WEBEXT.commands.onCommand.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCommand(callback js.Func[func(command js.String, tab *tabs.Tab)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCommand(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCommand returns true if the function "WEBEXT.commands.onCommand.removeListener" exists.
func HasFuncOffCommand() bool {
	return js.True == bindings.HasFuncOffCommand()
}

// FuncOffCommand returns the function "WEBEXT.commands.onCommand.removeListener".
func FuncOffCommand() (fn js.Func[func(callback js.Func[func(command js.String, tab *tabs.Tab)])]) {
	bindings.FuncOffCommand(
		js.Pointer(&fn),
	)
	return
}

// OffCommand calls the function "WEBEXT.commands.onCommand.removeListener" directly.
func OffCommand(callback js.Func[func(command js.String, tab *tabs.Tab)]) (ret js.Void) {
	bindings.CallOffCommand(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCommand calls the function "WEBEXT.commands.onCommand.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCommand(callback js.Func[func(command js.String, tab *tabs.Tab)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCommand(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCommand returns true if the function "WEBEXT.commands.onCommand.hasListener" exists.
func HasFuncHasOnCommand() bool {
	return js.True == bindings.HasFuncHasOnCommand()
}

// FuncHasOnCommand returns the function "WEBEXT.commands.onCommand.hasListener".
func FuncHasOnCommand() (fn js.Func[func(callback js.Func[func(command js.String, tab *tabs.Tab)]) bool]) {
	bindings.FuncHasOnCommand(
		js.Pointer(&fn),
	)
	return
}

// HasOnCommand calls the function "WEBEXT.commands.onCommand.hasListener" directly.
func HasOnCommand(callback js.Func[func(command js.String, tab *tabs.Tab)]) (ret bool) {
	bindings.CallHasOnCommand(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCommand calls the function "WEBEXT.commands.onCommand.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCommand(callback js.Func[func(command js.String, tab *tabs.Tab)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCommand(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}
