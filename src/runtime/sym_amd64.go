// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build pcz && amd64

package runtime

// Called from compiled code; declared for vet; do NOT call from Go.
func gcWriteBarrierCX() {}
func gcWriteBarrierDX() {}
func gcWriteBarrierBX() {}
func gcWriteBarrierBP() {}
func gcWriteBarrierSI() {}
func gcWriteBarrierR8() {}
func gcWriteBarrierR9() {}

// stackcheck checks that SP is in range [g->stack.lo, g->stack.hi).
func stackcheck()

// Called from assembly only; declared for go vet.
func settls() // argument in DI

// Retpolines, used by -spectre=ret flag in cmd/asm, cmd/compile.
func retpolineAX()
func retpolineCX()
func retpolineDX()
func retpolineBX()
func retpolineBP()
func retpolineSI()
func retpolineDI()
func retpolineR8()
func retpolineR9()
func retpolineR10()
func retpolineR11()
func retpolineR12()
func retpolineR13()
func retpolineR14()
func retpolineR15()
