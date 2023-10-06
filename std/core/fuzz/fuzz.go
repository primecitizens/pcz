// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package fuzz

// TODO

func TraceCmp1(arg0, arg1 uint8, fakepc int)       {}
func TraceCmp2(arg0, arg1 uint16, fakepc int)      {}
func TraceCmp4(arg0, arg1 uint32, fakepc int)      {}
func TraceCmp8(arg0, arg1 uint64, fakepc int)      {}
func TraceConstCmp1(arg0, arg1 uint8, fakepc int)  {}
func TraceConstCmp2(arg0, arg1 uint16, fakepc int) {}
func TraceConstCmp4(arg0, arg1 uint32, fakepc int) {}
func TraceConstCmp8(arg0, arg1 uint64, fakepc int) {}
func HookStrCmp(arg0, arg1 string, fakepc int)     {}
func HookEqualFold(arg0, arg1 string, fakepc int)  {}
