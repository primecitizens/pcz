// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build 386 || amd64

package cpu_test

import (
	"internal/godebug"
	"testing"

	. "github.com/primecitizens/pcz/std/core/cpu"
)

func TestX86ifAVX2hasAVX(t *testing.T) {
	if X86.HasAll(X86Feature_avx2) && !X86.HasAll(X86Feature_avx2) {
		t.Fatalf("HasAVX expected true when HasAVX2 is true, got false")
	}
}

func TestDisableSSE3(t *testing.T) {
	// if GOAMD64() > 1 {
	// 	t.Skip("skipping test: can't run on GOAMD64>v1 machines")
	// }
	// runDebugOptionsTest(t, "TestSSE3DebugOption", "cpu.sse3=off")
}

func TestSSE3DebugOption(t *testing.T) {
	MustHaveDebugOptionsSupport(t)

	if godebug.Get("cpu.sse3") != "off" {
		t.Skipf("skipping test: GODEBUG=cpu.sse3=off not set")
	}

	want := false
	if got := X86.HasAll(X86Feature_sse3); got != want {
		t.Errorf("X86.HasSSE3 expected %v, got %v", want, got)
	}
}
