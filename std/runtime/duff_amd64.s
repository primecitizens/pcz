// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Code generated by mkduff.go; DO NOT EDIT.
// Run go generate from src/runtime to update.
// See mkduff.go for comments.

//go:build pcz && amd64

#include "textflag.h"

TEXT runtime·duffzero<ABIInternal>(SB), NOSPLIT|NOFRAME, $0-0
	MOVUPS X15,(DI)
	MOVUPS X15,16(DI)
	MOVUPS X15,32(DI)
	MOVUPS X15,48(DI)
	LEAQ 64(DI),DI

	MOVUPS X15,(DI)
	MOVUPS X15,16(DI)
	MOVUPS X15,32(DI)
	MOVUPS X15,48(DI)
	LEAQ 64(DI),DI

	MOVUPS X15,(DI)
	MOVUPS X15,16(DI)
	MOVUPS X15,32(DI)
	MOVUPS X15,48(DI)
	LEAQ 64(DI),DI

	MOVUPS X15,(DI)
	MOVUPS X15,16(DI)
	MOVUPS X15,32(DI)
	MOVUPS X15,48(DI)
	LEAQ 64(DI),DI

	MOVUPS X15,(DI)
	MOVUPS X15,16(DI)
	MOVUPS X15,32(DI)
	MOVUPS X15,48(DI)
	LEAQ 64(DI),DI

	MOVUPS X15,(DI)
	MOVUPS X15,16(DI)
	MOVUPS X15,32(DI)
	MOVUPS X15,48(DI)
	LEAQ 64(DI),DI

	MOVUPS X15,(DI)
	MOVUPS X15,16(DI)
	MOVUPS X15,32(DI)
	MOVUPS X15,48(DI)
	LEAQ 64(DI),DI

	MOVUPS X15,(DI)
	MOVUPS X15,16(DI)
	MOVUPS X15,32(DI)
	MOVUPS X15,48(DI)
	LEAQ 64(DI),DI

	MOVUPS X15,(DI)
	MOVUPS X15,16(DI)
	MOVUPS X15,32(DI)
	MOVUPS X15,48(DI)
	LEAQ 64(DI),DI

	MOVUPS X15,(DI)
	MOVUPS X15,16(DI)
	MOVUPS X15,32(DI)
	MOVUPS X15,48(DI)
	LEAQ 64(DI),DI

	MOVUPS X15,(DI)
	MOVUPS X15,16(DI)
	MOVUPS X15,32(DI)
	MOVUPS X15,48(DI)
	LEAQ 64(DI),DI

	MOVUPS X15,(DI)
	MOVUPS X15,16(DI)
	MOVUPS X15,32(DI)
	MOVUPS X15,48(DI)
	LEAQ 64(DI),DI

	MOVUPS X15,(DI)
	MOVUPS X15,16(DI)
	MOVUPS X15,32(DI)
	MOVUPS X15,48(DI)
	LEAQ 64(DI),DI

	MOVUPS X15,(DI)
	MOVUPS X15,16(DI)
	MOVUPS X15,32(DI)
	MOVUPS X15,48(DI)
	LEAQ 64(DI),DI

	MOVUPS X15,(DI)
	MOVUPS X15,16(DI)
	MOVUPS X15,32(DI)
	MOVUPS X15,48(DI)
	LEAQ 64(DI),DI

	MOVUPS X15,(DI)
	MOVUPS X15,16(DI)
	MOVUPS X15,32(DI)
	MOVUPS X15,48(DI)
	LEAQ 64(DI),DI

	RET

TEXT runtime·duffcopy<ABIInternal>(SB), NOSPLIT|NOFRAME, $0-0
	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	MOVUPS (SI), X0
	ADDQ $16, SI
	MOVUPS X0, (DI)
	ADDQ $16, DI

	RET