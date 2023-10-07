// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package contentscripts

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/contentscripts/bindings"
	"github.com/primecitizens/pcz/std/plat/js/webext/extensiontypes"
)

type RunAt uint32

const (
	_ RunAt = iota

	RunAt_DOCUMENT_IDLE
	RunAt_DOCUMENT_START
	RunAt_DOCUMENT_END
)

func (RunAt) FromRef(str js.Ref) RunAt {
	return RunAt(bindings.ConstOfRunAt(str))
}

func (x RunAt) String() (string, bool) {
	switch x {
	case RunAt_DOCUMENT_IDLE:
		return "document_idle", true
	case RunAt_DOCUMENT_START:
		return "document_start", true
	case RunAt_DOCUMENT_END:
		return "document_end", true
	default:
		return "", false
	}
}

type ContentScript struct {
	// Matches is "ContentScript.matches"
	//
	// Optional
	Matches js.Array[js.String]
	// ExcludeMatches is "ContentScript.exclude_matches"
	//
	// Optional
	ExcludeMatches js.Array[js.String]
	// Css is "ContentScript.css"
	//
	// Optional
	Css js.Array[js.String]
	// Js is "ContentScript.js"
	//
	// Optional
	Js js.Array[js.String]
	// AllFrames is "ContentScript.all_frames"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllFrames MUST be set to true to make this field effective.
	AllFrames bool
	// MatchOriginAsFallback is "ContentScript.match_origin_as_fallback"
	//
	// Optional
	//
	// NOTE: FFI_USE_MatchOriginAsFallback MUST be set to true to make this field effective.
	MatchOriginAsFallback bool
	// MatchAboutBlank is "ContentScript.match_about_blank"
	//
	// Optional
	//
	// NOTE: FFI_USE_MatchAboutBlank MUST be set to true to make this field effective.
	MatchAboutBlank bool
	// IncludeGlobs is "ContentScript.include_globs"
	//
	// Optional
	IncludeGlobs js.Array[js.String]
	// ExcludeGlobs is "ContentScript.exclude_globs"
	//
	// Optional
	ExcludeGlobs js.Array[js.String]
	// RunAt is "ContentScript.run_at"
	//
	// Optional
	RunAt RunAt
	// World is "ContentScript.world"
	//
	// Optional
	World extensiontypes.ExecutionWorld

	FFI_USE_AllFrames             bool // for AllFrames.
	FFI_USE_MatchOriginAsFallback bool // for MatchOriginAsFallback.
	FFI_USE_MatchAboutBlank       bool // for MatchAboutBlank.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ContentScript with all fields set.
func (p ContentScript) FromRef(ref js.Ref) ContentScript {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ContentScript in the application heap.
func (p ContentScript) New() js.Ref {
	return bindings.ContentScriptJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ContentScript) UpdateFrom(ref js.Ref) {
	bindings.ContentScriptJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ContentScript) Update(ref js.Ref) {
	bindings.ContentScriptJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ContentScript) FreeMembers(recursive bool) {
	js.Free(
		p.Matches.Ref(),
		p.ExcludeMatches.Ref(),
		p.Css.Ref(),
		p.Js.Ref(),
		p.IncludeGlobs.Ref(),
		p.ExcludeGlobs.Ref(),
	)
	p.Matches = p.Matches.FromRef(js.Undefined)
	p.ExcludeMatches = p.ExcludeMatches.FromRef(js.Undefined)
	p.Css = p.Css.FromRef(js.Undefined)
	p.Js = p.Js.FromRef(js.Undefined)
	p.IncludeGlobs = p.IncludeGlobs.FromRef(js.Undefined)
	p.ExcludeGlobs = p.ExcludeGlobs.FromRef(js.Undefined)
}

type ManifestKeys struct {
	// ContentScripts is "ManifestKeys.content_scripts"
	//
	// Optional
	ContentScripts js.Array[ContentScript]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ManifestKeys with all fields set.
func (p ManifestKeys) FromRef(ref js.Ref) ManifestKeys {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ManifestKeys in the application heap.
func (p ManifestKeys) New() js.Ref {
	return bindings.ManifestKeysJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ManifestKeys) UpdateFrom(ref js.Ref) {
	bindings.ManifestKeysJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ManifestKeys) Update(ref js.Ref) {
	bindings.ManifestKeysJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ManifestKeys) FreeMembers(recursive bool) {
	js.Free(
		p.ContentScripts.Ref(),
	)
	p.ContentScripts = p.ContentScripts.FromRef(js.Undefined)
}
