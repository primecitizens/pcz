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

//go:wasmimport plat/js/webext/wallpaper constof_WallpaperLayout
//go:noescape
func ConstOfWallpaperLayout(str js.Ref) uint32

//go:wasmimport plat/js/webext/wallpaper store_SetWallpaperArgDetails
//go:noescape
func SetWallpaperArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/wallpaper load_SetWallpaperArgDetails
//go:noescape
func SetWallpaperArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/wallpaper has_SetWallpaper
//go:noescape
func HasFuncSetWallpaper() js.Ref

//go:wasmimport plat/js/webext/wallpaper func_SetWallpaper
//go:noescape
func FuncSetWallpaper(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/wallpaper call_SetWallpaper
//go:noescape
func CallSetWallpaper(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/wallpaper try_SetWallpaper
//go:noescape
func TrySetWallpaper(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)
