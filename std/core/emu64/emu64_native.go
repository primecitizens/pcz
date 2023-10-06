// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !(arm || 386 || mips || mipsle)

package emu64

func Float64ToInt64(n float64) uint64  { return uint64(n) }
func Float64ToUint64(n float64) uint64 { return uint64(n) }
func Int64ToFloat64(n int64) float64   { return float64(n) }
func Uint64ToFloat64(n uint64) float64 { return float64(n) }
func Int64ToFloat32(n int64) float32   { return float32(n) }
func Uint64ToFloat32(n uint64) float32 { return float32(n) }
func Uint64Div(n, d uint64) uint64     { return n / d }
func Uint64Mod(n, d uint64) uint64     { return n % d }
func Int64Div(n, d int64) int64        { return n / d }
func Int64Mod(n, d int64) int64        { return n % d }
func Float64ToUint32(n float64) uint32 { return uint32(n) }
func Uint32ToFloat64(n uint32) float64 { return float64(n) }
