// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

#ifdef GOARCH_386
#define PTRSIZE 4
#endif
#ifdef GOARCH_arm
#define PTRSIZE 4
#endif
#ifdef GOARCH_mips
#define PTRSIZE 4
#endif
#ifdef GOARCH_mipsle
#define PTRSIZE 4
#endif
#ifndef PTRSIZE
#define PTRSIZE 8
#endif

TEXT core/abi·FuncPCTestFn(SB),NOSPLIT,$0-0
	RET

GLOBL core/abi·FuncPCTestFnAddr(SB), NOPTR, $PTRSIZE
DATA core/abi·FuncPCTestFnAddr(SB)/PTRSIZE, $core/abi·FuncPCTestFn(SB)
