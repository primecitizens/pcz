// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package filehandlers

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/filehandlers/bindings"
)

type Icon struct {
	// Src is "Icon.src"
	//
	// Optional
	Src js.String
	// Sizes is "Icon.sizes"
	//
	// Optional
	Sizes js.String
	// Type is "Icon.type"
	//
	// Optional
	Type js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Icon with all fields set.
func (p Icon) FromRef(ref js.Ref) Icon {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Icon in the application heap.
func (p Icon) New() js.Ref {
	return bindings.IconJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Icon) UpdateFrom(ref js.Ref) {
	bindings.IconJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Icon) Update(ref js.Ref) {
	bindings.IconJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Icon) FreeMembers(recursive bool) {
	js.Free(
		p.Src.Ref(),
		p.Sizes.Ref(),
		p.Type.Ref(),
	)
	p.Src = p.Src.FromRef(js.Undefined)
	p.Sizes = p.Sizes.FromRef(js.Undefined)
	p.Type = p.Type.FromRef(js.Undefined)
}

type FileHandler struct {
	// Accept is "FileHandler.accept"
	//
	// Optional
	Accept js.Object
	// Action is "FileHandler.action"
	//
	// Optional
	Action js.String
	// Name is "FileHandler.name"
	//
	// Optional
	Name js.String
	// Icons is "FileHandler.icons"
	//
	// Optional
	Icons js.Array[Icon]
	// LaunchType is "FileHandler.launch_type"
	//
	// Optional
	LaunchType js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FileHandler with all fields set.
func (p FileHandler) FromRef(ref js.Ref) FileHandler {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FileHandler in the application heap.
func (p FileHandler) New() js.Ref {
	return bindings.FileHandlerJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FileHandler) UpdateFrom(ref js.Ref) {
	bindings.FileHandlerJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FileHandler) Update(ref js.Ref) {
	bindings.FileHandlerJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FileHandler) FreeMembers(recursive bool) {
	js.Free(
		p.Accept.Ref(),
		p.Action.Ref(),
		p.Name.Ref(),
		p.Icons.Ref(),
		p.LaunchType.Ref(),
	)
	p.Accept = p.Accept.FromRef(js.Undefined)
	p.Action = p.Action.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Icons = p.Icons.FromRef(js.Undefined)
	p.LaunchType = p.LaunchType.FromRef(js.Undefined)
}

type ManifestKeys struct {
	// FileHandlers is "ManifestKeys.file_handlers"
	//
	// Optional
	FileHandlers js.Array[FileHandler]

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
		p.FileHandlers.Ref(),
	)
	p.FileHandlers = p.FileHandlers.FromRef(js.Undefined)
}
