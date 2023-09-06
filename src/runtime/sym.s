// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

TEXT runtime·mapinitnoop<ABIInternal>(SB),NOSPLIT,$0-0
	RET

#ifdef GOARCH_riscv64
TEXT runtime·getcallerpc<ABIInternal>(SB),NOSPLIT|NOFRAME,$0-8
	MOV 0(X2), X10 // LR saved by caller
	RET
#endif
