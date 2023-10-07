// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/ffi/js"
)

type (
	_ unsafe.Pointer
	_ js.Ref
)

//go:wasmimport plat/js/webext/enterprise/platformkeysprivate has_ChallengeMachineKey
//go:noescape
func HasFuncChallengeMachineKey() js.Ref

//go:wasmimport plat/js/webext/enterprise/platformkeysprivate func_ChallengeMachineKey
//go:noescape
func FuncChallengeMachineKey(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/platformkeysprivate call_ChallengeMachineKey
//go:noescape
func CallChallengeMachineKey(
	retPtr unsafe.Pointer,
	challenge js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeysprivate try_ChallengeMachineKey
//go:noescape
func TryChallengeMachineKey(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	challenge js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeysprivate has_ChallengeUserKey
//go:noescape
func HasFuncChallengeUserKey() js.Ref

//go:wasmimport plat/js/webext/enterprise/platformkeysprivate func_ChallengeUserKey
//go:noescape
func FuncChallengeUserKey(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/platformkeysprivate call_ChallengeUserKey
//go:noescape
func CallChallengeUserKey(
	retPtr unsafe.Pointer,
	challenge js.Ref,
	registerKey js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeysprivate try_ChallengeUserKey
//go:noescape
func TryChallengeUserKey(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	challenge js.Ref,
	registerKey js.Ref) (ok js.Ref)
