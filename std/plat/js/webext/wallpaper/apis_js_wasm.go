// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package wallpaper

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/wallpaper/bindings"
)

type WallpaperLayout uint32

const (
	_ WallpaperLayout = iota

	WallpaperLayout_STRETCH
	WallpaperLayout_CENTER
	WallpaperLayout_CENTER_CROPPED
)

func (WallpaperLayout) FromRef(str js.Ref) WallpaperLayout {
	return WallpaperLayout(bindings.ConstOfWallpaperLayout(str))
}

func (x WallpaperLayout) String() (string, bool) {
	switch x {
	case WallpaperLayout_STRETCH:
		return "STRETCH", true
	case WallpaperLayout_CENTER:
		return "CENTER", true
	case WallpaperLayout_CENTER_CROPPED:
		return "CENTER_CROPPED", true
	default:
		return "", false
	}
}

type SetWallpaperArgDetails struct {
	// Data is "SetWallpaperArgDetails.data"
	//
	// Optional
	Data js.TypedArray[uint8]
	// Filename is "SetWallpaperArgDetails.filename"
	//
	// Required
	Filename js.String
	// Layout is "SetWallpaperArgDetails.layout"
	//
	// Required
	Layout WallpaperLayout
	// Thumbnail is "SetWallpaperArgDetails.thumbnail"
	//
	// Optional
	//
	// NOTE: FFI_USE_Thumbnail MUST be set to true to make this field effective.
	Thumbnail bool
	// Url is "SetWallpaperArgDetails.url"
	//
	// Optional
	Url js.String

	FFI_USE_Thumbnail bool // for Thumbnail.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetWallpaperArgDetails with all fields set.
func (p SetWallpaperArgDetails) FromRef(ref js.Ref) SetWallpaperArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetWallpaperArgDetails in the application heap.
func (p SetWallpaperArgDetails) New() js.Ref {
	return bindings.SetWallpaperArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetWallpaperArgDetails) UpdateFrom(ref js.Ref) {
	bindings.SetWallpaperArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetWallpaperArgDetails) Update(ref js.Ref) {
	bindings.SetWallpaperArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetWallpaperArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
		p.Filename.Ref(),
		p.Url.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
	p.Filename = p.Filename.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

// HasFuncSetWallpaper returns true if the function "WEBEXT.wallpaper.setWallpaper" exists.
func HasFuncSetWallpaper() bool {
	return js.True == bindings.HasFuncSetWallpaper()
}

// FuncSetWallpaper returns the function "WEBEXT.wallpaper.setWallpaper".
func FuncSetWallpaper() (fn js.Func[func(details SetWallpaperArgDetails) js.Promise[js.TypedArray[uint8]]]) {
	bindings.FuncSetWallpaper(
		js.Pointer(&fn),
	)
	return
}

// SetWallpaper calls the function "WEBEXT.wallpaper.setWallpaper" directly.
func SetWallpaper(details SetWallpaperArgDetails) (ret js.Promise[js.TypedArray[uint8]]) {
	bindings.CallSetWallpaper(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TrySetWallpaper calls the function "WEBEXT.wallpaper.setWallpaper"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetWallpaper(details SetWallpaperArgDetails) (ret js.Promise[js.TypedArray[uint8]], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetWallpaper(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}
