// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package extensiontypes

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/extensiontypes/bindings"
)

type CSSOrigin uint32

const (
	_ CSSOrigin = iota

	CSSOrigin_AUTHOR
	CSSOrigin_USER
)

func (CSSOrigin) FromRef(str js.Ref) CSSOrigin {
	return CSSOrigin(bindings.ConstOfCSSOrigin(str))
}

func (x CSSOrigin) String() (string, bool) {
	switch x {
	case CSSOrigin_AUTHOR:
		return "author", true
	case CSSOrigin_USER:
		return "user", true
	default:
		return "", false
	}
}

type ExecutionWorld uint32

const (
	_ ExecutionWorld = iota

	ExecutionWorld_ISOLATED
	ExecutionWorld_MAIN
)

func (ExecutionWorld) FromRef(str js.Ref) ExecutionWorld {
	return ExecutionWorld(bindings.ConstOfExecutionWorld(str))
}

func (x ExecutionWorld) String() (string, bool) {
	switch x {
	case ExecutionWorld_ISOLATED:
		return "ISOLATED", true
	case ExecutionWorld_MAIN:
		return "MAIN", true
	default:
		return "", false
	}
}

type RunAt uint32

const (
	_ RunAt = iota

	RunAt_DOCUMENT_START
	RunAt_DOCUMENT_END
	RunAt_DOCUMENT_IDLE
)

func (RunAt) FromRef(str js.Ref) RunAt {
	return RunAt(bindings.ConstOfRunAt(str))
}

func (x RunAt) String() (string, bool) {
	switch x {
	case RunAt_DOCUMENT_START:
		return "document_start", true
	case RunAt_DOCUMENT_END:
		return "document_end", true
	case RunAt_DOCUMENT_IDLE:
		return "document_idle", true
	default:
		return "", false
	}
}

type DeleteInjectionDetails struct {
	// AllFrames is "DeleteInjectionDetails.allFrames"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllFrames MUST be set to true to make this field effective.
	AllFrames bool
	// Code is "DeleteInjectionDetails.code"
	//
	// Optional
	Code js.String
	// CssOrigin is "DeleteInjectionDetails.cssOrigin"
	//
	// Optional
	CssOrigin CSSOrigin
	// File is "DeleteInjectionDetails.file"
	//
	// Optional
	File js.String
	// FrameId is "DeleteInjectionDetails.frameId"
	//
	// Optional
	//
	// NOTE: FFI_USE_FrameId MUST be set to true to make this field effective.
	FrameId int64
	// MatchAboutBlank is "DeleteInjectionDetails.matchAboutBlank"
	//
	// Optional
	//
	// NOTE: FFI_USE_MatchAboutBlank MUST be set to true to make this field effective.
	MatchAboutBlank bool

	FFI_USE_AllFrames       bool // for AllFrames.
	FFI_USE_FrameId         bool // for FrameId.
	FFI_USE_MatchAboutBlank bool // for MatchAboutBlank.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DeleteInjectionDetails with all fields set.
func (p DeleteInjectionDetails) FromRef(ref js.Ref) DeleteInjectionDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DeleteInjectionDetails in the application heap.
func (p DeleteInjectionDetails) New() js.Ref {
	return bindings.DeleteInjectionDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DeleteInjectionDetails) UpdateFrom(ref js.Ref) {
	bindings.DeleteInjectionDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DeleteInjectionDetails) Update(ref js.Ref) {
	bindings.DeleteInjectionDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DeleteInjectionDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Code.Ref(),
		p.File.Ref(),
	)
	p.Code = p.Code.FromRef(js.Undefined)
	p.File = p.File.FromRef(js.Undefined)
}

type DocumentLifecycle uint32

const (
	_ DocumentLifecycle = iota

	DocumentLifecycle_PRERENDER
	DocumentLifecycle_ACTIVE
	DocumentLifecycle_CACHED
	DocumentLifecycle_PENDING_DELETION
)

func (DocumentLifecycle) FromRef(str js.Ref) DocumentLifecycle {
	return DocumentLifecycle(bindings.ConstOfDocumentLifecycle(str))
}

func (x DocumentLifecycle) String() (string, bool) {
	switch x {
	case DocumentLifecycle_PRERENDER:
		return "prerender", true
	case DocumentLifecycle_ACTIVE:
		return "active", true
	case DocumentLifecycle_CACHED:
		return "cached", true
	case DocumentLifecycle_PENDING_DELETION:
		return "pending_deletion", true
	default:
		return "", false
	}
}

type FrameType uint32

const (
	_ FrameType = iota

	FrameType_OUTERMOST_FRAME
	FrameType_FENCED_FRAME
	FrameType_SUB_FRAME
)

func (FrameType) FromRef(str js.Ref) FrameType {
	return FrameType(bindings.ConstOfFrameType(str))
}

func (x FrameType) String() (string, bool) {
	switch x {
	case FrameType_OUTERMOST_FRAME:
		return "outermost_frame", true
	case FrameType_FENCED_FRAME:
		return "fenced_frame", true
	case FrameType_SUB_FRAME:
		return "sub_frame", true
	default:
		return "", false
	}
}

type ImageFormat uint32

const (
	_ ImageFormat = iota

	ImageFormat_JPEG
	ImageFormat_PNG
)

func (ImageFormat) FromRef(str js.Ref) ImageFormat {
	return ImageFormat(bindings.ConstOfImageFormat(str))
}

func (x ImageFormat) String() (string, bool) {
	switch x {
	case ImageFormat_JPEG:
		return "jpeg", true
	case ImageFormat_PNG:
		return "png", true
	default:
		return "", false
	}
}

type ImageDetails struct {
	// Format is "ImageDetails.format"
	//
	// Optional
	Format ImageFormat
	// Quality is "ImageDetails.quality"
	//
	// Optional
	//
	// NOTE: FFI_USE_Quality MUST be set to true to make this field effective.
	Quality int64

	FFI_USE_Quality bool // for Quality.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ImageDetails with all fields set.
func (p ImageDetails) FromRef(ref js.Ref) ImageDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ImageDetails in the application heap.
func (p ImageDetails) New() js.Ref {
	return bindings.ImageDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ImageDetails) UpdateFrom(ref js.Ref) {
	bindings.ImageDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ImageDetails) Update(ref js.Ref) {
	bindings.ImageDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ImageDetails) FreeMembers(recursive bool) {
}

type InjectDetails struct {
	// AllFrames is "InjectDetails.allFrames"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllFrames MUST be set to true to make this field effective.
	AllFrames bool
	// Code is "InjectDetails.code"
	//
	// Optional
	Code js.String
	// CssOrigin is "InjectDetails.cssOrigin"
	//
	// Optional
	CssOrigin CSSOrigin
	// File is "InjectDetails.file"
	//
	// Optional
	File js.String
	// FrameId is "InjectDetails.frameId"
	//
	// Optional
	//
	// NOTE: FFI_USE_FrameId MUST be set to true to make this field effective.
	FrameId int64
	// MatchAboutBlank is "InjectDetails.matchAboutBlank"
	//
	// Optional
	//
	// NOTE: FFI_USE_MatchAboutBlank MUST be set to true to make this field effective.
	MatchAboutBlank bool
	// RunAt is "InjectDetails.runAt"
	//
	// Optional
	RunAt RunAt

	FFI_USE_AllFrames       bool // for AllFrames.
	FFI_USE_FrameId         bool // for FrameId.
	FFI_USE_MatchAboutBlank bool // for MatchAboutBlank.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a InjectDetails with all fields set.
func (p InjectDetails) FromRef(ref js.Ref) InjectDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new InjectDetails in the application heap.
func (p InjectDetails) New() js.Ref {
	return bindings.InjectDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *InjectDetails) UpdateFrom(ref js.Ref) {
	bindings.InjectDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *InjectDetails) Update(ref js.Ref) {
	bindings.InjectDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *InjectDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Code.Ref(),
		p.File.Ref(),
	)
	p.Code = p.Code.FromRef(js.Undefined)
	p.File = p.File.FromRef(js.Undefined)
}
