// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

// when building with the "noos" tag, the GOOS is actually "linux".
//
//go:build noos && riscv64


#include "textflag.h"


#ifdef GOOS_efi

TEXT rt0(SB),NOSPLIT|NOFRAME,$0
	// TODO: pass args (EFI_HANDLE imgHandle, EFI_SYSTEM_TABLE *sysTab)
	JMP main·efi_main<ABIInternal>(SB)

#endif
