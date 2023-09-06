// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

// Package stack provides low-level access to current stack space
package stack

import (
	stdgo "github.com/primecitizens/std/builtin/go"
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/arch"
	"github.com/primecitizens/std/core/os"
	"github.com/primecitizens/std/core/race"
)

type Stack = stdgo.Stack

const (
	// AIX requires a larger stack for syscalls.
	// The race build also needs more stack. See issue 54291.
	//
	// This arithmetic must match that in
	// $GOROOT/cmd/internal/objabi/stack.go:stackGuardMultiplier.
	StackGuardMultiplier = 1 + os.IsAix + race.IsRace

	// StackSystem is a number of additional bytes to add
	// to each stack below the usual guard area for OS-specific
	// purposes like signal handling. Used on Windows, Plan 9,
	// and iOS because they do not use a separate stack.
	StackSystem = os.IsWindows*512*arch.PtrSize +
		os.IsPlan9*512 +
		os.IsIos*arch.IsArm64*1024

	// The minimum size of stack used by Go code
	stackMin = 2048

	// The minimum stack size to allocate.
	// The hackery here rounds fixedStack0 up to a power of 2.
	fixedStack0 = stackMin + StackSystem
	fixedStack1 = fixedStack0 - 1
	fixedStack2 = fixedStack1 | (fixedStack1 >> 1)
	fixedStack3 = fixedStack2 | (fixedStack2 >> 2)
	fixedStack4 = fixedStack3 | (fixedStack3 >> 4)
	fixedStack5 = fixedStack4 | (fixedStack4 >> 8)
	fixedStack6 = fixedStack5 | (fixedStack5 >> 16)
	fixedStack  = fixedStack6 + 1

	// StackNosplit is the maximum number of bytes that a chain of NOSPLIT
	// functions can use.
	//
	// This arithmetic must match that in cmd/internal/objabi/stack.go:StackNosplit.
	StackNosplit = abi.StackNosplitBase * StackGuardMultiplier

	// The stack guard is a pointer this many bytes above the
	// bottom of the stack.
	//
	// The guard leaves enough room for a stackNosplit chain of NOSPLIT calls
	// plus one stackSmall frame plus stackSystem bytes for the OS.
	//
	// This arithmetic must match that in cmd/internal/objabi/stack.go:StackLimit.
	StackGuard = StackNosplit + StackSystem + abi.StackSmall
)

// GetSP returns the value of the stack pointer register on calling.
func GetSP() uintptr

//go:noescape
func SetSP(sp uintptr)
