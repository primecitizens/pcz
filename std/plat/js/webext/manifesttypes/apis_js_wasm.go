// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package manifesttypes

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/manifesttypes/bindings"
)

type ChromeSettingsOverridesFieldSearchProvider struct {
	// AlternateUrls is "ChromeSettingsOverridesFieldSearchProvider.alternate_urls"
	//
	// Optional
	AlternateUrls js.Array[js.String]
	// Encoding is "ChromeSettingsOverridesFieldSearchProvider.encoding"
	//
	// Optional
	Encoding js.String
	// FaviconUrl is "ChromeSettingsOverridesFieldSearchProvider.favicon_url"
	//
	// Optional
	FaviconUrl js.String
	// ImageUrl is "ChromeSettingsOverridesFieldSearchProvider.image_url"
	//
	// Optional
	ImageUrl js.String
	// ImageUrlPostParams is "ChromeSettingsOverridesFieldSearchProvider.image_url_post_params"
	//
	// Optional
	ImageUrlPostParams js.String
	// IsDefault is "ChromeSettingsOverridesFieldSearchProvider.is_default"
	//
	// Required
	IsDefault bool
	// Keyword is "ChromeSettingsOverridesFieldSearchProvider.keyword"
	//
	// Optional
	Keyword js.String
	// Name is "ChromeSettingsOverridesFieldSearchProvider.name"
	//
	// Optional
	Name js.String
	// PrepopulatedId is "ChromeSettingsOverridesFieldSearchProvider.prepopulated_id"
	//
	// Optional
	//
	// NOTE: FFI_USE_PrepopulatedId MUST be set to true to make this field effective.
	PrepopulatedId int64
	// SearchUrl is "ChromeSettingsOverridesFieldSearchProvider.search_url"
	//
	// Required
	SearchUrl js.String
	// SearchUrlPostParams is "ChromeSettingsOverridesFieldSearchProvider.search_url_post_params"
	//
	// Optional
	SearchUrlPostParams js.String
	// SuggestUrl is "ChromeSettingsOverridesFieldSearchProvider.suggest_url"
	//
	// Optional
	SuggestUrl js.String
	// SuggestUrlPostParams is "ChromeSettingsOverridesFieldSearchProvider.suggest_url_post_params"
	//
	// Optional
	SuggestUrlPostParams js.String

	FFI_USE_PrepopulatedId bool // for PrepopulatedId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ChromeSettingsOverridesFieldSearchProvider with all fields set.
func (p ChromeSettingsOverridesFieldSearchProvider) FromRef(ref js.Ref) ChromeSettingsOverridesFieldSearchProvider {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ChromeSettingsOverridesFieldSearchProvider in the application heap.
func (p ChromeSettingsOverridesFieldSearchProvider) New() js.Ref {
	return bindings.ChromeSettingsOverridesFieldSearchProviderJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ChromeSettingsOverridesFieldSearchProvider) UpdateFrom(ref js.Ref) {
	bindings.ChromeSettingsOverridesFieldSearchProviderJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ChromeSettingsOverridesFieldSearchProvider) Update(ref js.Ref) {
	bindings.ChromeSettingsOverridesFieldSearchProviderJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ChromeSettingsOverridesFieldSearchProvider) FreeMembers(recursive bool) {
	js.Free(
		p.AlternateUrls.Ref(),
		p.Encoding.Ref(),
		p.FaviconUrl.Ref(),
		p.ImageUrl.Ref(),
		p.ImageUrlPostParams.Ref(),
		p.Keyword.Ref(),
		p.Name.Ref(),
		p.SearchUrl.Ref(),
		p.SearchUrlPostParams.Ref(),
		p.SuggestUrl.Ref(),
		p.SuggestUrlPostParams.Ref(),
	)
	p.AlternateUrls = p.AlternateUrls.FromRef(js.Undefined)
	p.Encoding = p.Encoding.FromRef(js.Undefined)
	p.FaviconUrl = p.FaviconUrl.FromRef(js.Undefined)
	p.ImageUrl = p.ImageUrl.FromRef(js.Undefined)
	p.ImageUrlPostParams = p.ImageUrlPostParams.FromRef(js.Undefined)
	p.Keyword = p.Keyword.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.SearchUrl = p.SearchUrl.FromRef(js.Undefined)
	p.SearchUrlPostParams = p.SearchUrlPostParams.FromRef(js.Undefined)
	p.SuggestUrl = p.SuggestUrl.FromRef(js.Undefined)
	p.SuggestUrlPostParams = p.SuggestUrlPostParams.FromRef(js.Undefined)
}

type ChromeSettingsOverrides struct {
	// Homepage is "ChromeSettingsOverrides.homepage"
	//
	// Optional
	Homepage js.String
	// SearchProvider is "ChromeSettingsOverrides.search_provider"
	//
	// Optional
	//
	// NOTE: SearchProvider.FFI_USE MUST be set to true to get SearchProvider used.
	SearchProvider ChromeSettingsOverridesFieldSearchProvider
	// StartupPages is "ChromeSettingsOverrides.startup_pages"
	//
	// Optional
	StartupPages js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ChromeSettingsOverrides with all fields set.
func (p ChromeSettingsOverrides) FromRef(ref js.Ref) ChromeSettingsOverrides {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ChromeSettingsOverrides in the application heap.
func (p ChromeSettingsOverrides) New() js.Ref {
	return bindings.ChromeSettingsOverridesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ChromeSettingsOverrides) UpdateFrom(ref js.Ref) {
	bindings.ChromeSettingsOverridesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ChromeSettingsOverrides) Update(ref js.Ref) {
	bindings.ChromeSettingsOverridesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ChromeSettingsOverrides) FreeMembers(recursive bool) {
	js.Free(
		p.Homepage.Ref(),
		p.StartupPages.Ref(),
	)
	p.Homepage = p.Homepage.FromRef(js.Undefined)
	p.StartupPages = p.StartupPages.FromRef(js.Undefined)
	if recursive {
		p.SearchProvider.FreeMembers(true)
	}
}

type FileSystemProviderSource uint32

const (
	_ FileSystemProviderSource = iota

	FileSystemProviderSource_FILE
	FileSystemProviderSource_DEVICE
	FileSystemProviderSource_NETWORK
)

func (FileSystemProviderSource) FromRef(str js.Ref) FileSystemProviderSource {
	return FileSystemProviderSource(bindings.ConstOfFileSystemProviderSource(str))
}

func (x FileSystemProviderSource) String() (string, bool) {
	switch x {
	case FileSystemProviderSource_FILE:
		return "file", true
	case FileSystemProviderSource_DEVICE:
		return "device", true
	case FileSystemProviderSource_NETWORK:
		return "network", true
	default:
		return "", false
	}
}

type FileSystemProviderCapabilities struct {
	// Configurable is "FileSystemProviderCapabilities.configurable"
	//
	// Optional
	//
	// NOTE: FFI_USE_Configurable MUST be set to true to make this field effective.
	Configurable bool
	// MultipleMounts is "FileSystemProviderCapabilities.multiple_mounts"
	//
	// Optional
	//
	// NOTE: FFI_USE_MultipleMounts MUST be set to true to make this field effective.
	MultipleMounts bool
	// Source is "FileSystemProviderCapabilities.source"
	//
	// Required
	Source FileSystemProviderSource
	// Watchable is "FileSystemProviderCapabilities.watchable"
	//
	// Optional
	//
	// NOTE: FFI_USE_Watchable MUST be set to true to make this field effective.
	Watchable bool

	FFI_USE_Configurable   bool // for Configurable.
	FFI_USE_MultipleMounts bool // for MultipleMounts.
	FFI_USE_Watchable      bool // for Watchable.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FileSystemProviderCapabilities with all fields set.
func (p FileSystemProviderCapabilities) FromRef(ref js.Ref) FileSystemProviderCapabilities {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FileSystemProviderCapabilities in the application heap.
func (p FileSystemProviderCapabilities) New() js.Ref {
	return bindings.FileSystemProviderCapabilitiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FileSystemProviderCapabilities) UpdateFrom(ref js.Ref) {
	bindings.FileSystemProviderCapabilitiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FileSystemProviderCapabilities) Update(ref js.Ref) {
	bindings.FileSystemProviderCapabilitiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FileSystemProviderCapabilities) FreeMembers(recursive bool) {
}
