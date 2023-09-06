// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package buildcfg

import (
	"encoding/hex"
	"time"
)

// The first 14 bytes are buildInfoMagic.
// The next two bytes indicate pointer size in bytes (4 or 8) and endianness
// (0 for little, 1 for big).
// Two virtual addresses to Go strings follow that: runtime.buildVersion,
// and runtime.modinfo.
// On 32-bit platforms, the last 8 bytes are unused.
// If the endianness has the 2 bit set, then the pointers are zero
// and the 32-byte header is followed by varint-prefixed string data
// for the two string values we care about.

var (
	infoStart, _ = hex.DecodeString("3077af0c9274080241e1c107e6d618e6")
	infoEnd, _   = hex.DecodeString("f932433186182072008242104116d8f2")
)

func ModInfoData(info string) []byte {
	return []byte(string(infoStart) + info + string(infoEnd))
}

// Package
//
// copied from src/cmd/go/internal/load/pkg.go#PackagePublic
type Package struct {
	// Note: These fields are part of the go command's public API.
	// See list.go. It is okay to add fields, but not to change or
	// remove existing ones. Keep in sync with ../list/list.go
	Dir           string   `json:",omitempty"` // directory containing package sources
	ImportPath    string   `json:",omitempty"` // import path of package in dir
	ImportComment string   `json:",omitempty"` // path in import comment on package statement
	Name          string   `json:",omitempty"` // package name
	Doc           string   `json:",omitempty"` // package documentation string
	Target        string   `json:",omitempty"` // installed target for this package (may be executable)
	Shlib         string   `json:",omitempty"` // the shared library that contains this package (only set when -linkshared)
	Root          string   `json:",omitempty"` // Go root, Go path dir, or module root dir containing this package
	ConflictDir   string   `json:",omitempty"` // Dir is hidden by this other directory
	ForTest       string   `json:",omitempty"` // package is only for use in named test
	Export        string   `json:",omitempty"` // file containing export data (set by go list -export)
	BuildID       string   `json:",omitempty"` // build ID of the compiled package (set by go list -export)
	Module        *Module  `json:",omitempty"` // info about package's module, if any
	Match         []string `json:",omitempty"` // command-line patterns matching this package
	Goroot        bool     `json:",omitempty"` // is this package found in the Go root?
	Standard      bool     `json:",omitempty"` // is this package part of the standard Go library?
	DepOnly       bool     `json:",omitempty"` // package is only as a dependency, not explicitly listed
	BinaryOnly    bool     `json:",omitempty"` // package cannot be recompiled
	Incomplete    bool     `json:",omitempty"` // was there an error loading this package or dependencies?

	DefaultGODEBUG string `json:",omitempty"` // default GODEBUG setting (only for Name=="main")

	// Stale and StaleReason remain here *only* for the list command.
	// They are only initialized in preparation for list execution.
	// The regular build determines staleness on the fly during action execution.
	Stale       bool   `json:",omitempty"` // would 'go install' do anything for this package?
	StaleReason string `json:",omitempty"` // why is Stale true?

	// Source files
	// If you add to this list you MUST add to p.AllFiles (below) too.
	// Otherwise file name security lists will not apply to any new additions.
	GoFiles           []string `json:",omitempty"` // .go source files (excluding CgoFiles, TestGoFiles, XTestGoFiles)
	CgoFiles          []string `json:",omitempty"` // .go source files that import "C"
	CompiledGoFiles   []string `json:",omitempty"` // .go output from running cgo on CgoFiles
	IgnoredGoFiles    []string `json:",omitempty"` // .go source files ignored due to build constraints
	InvalidGoFiles    []string `json:",omitempty"` // .go source files with detected problems (parse error, wrong package name, and so on)
	IgnoredOtherFiles []string `json:",omitempty"` // non-.go source files ignored due to build constraints
	CFiles            []string `json:",omitempty"` // .c source files
	CXXFiles          []string `json:",omitempty"` // .cc, .cpp and .cxx source files
	MFiles            []string `json:",omitempty"` // .m source files
	HFiles            []string `json:",omitempty"` // .h, .hh, .hpp and .hxx source files
	FFiles            []string `json:",omitempty"` // .f, .F, .for and .f90 Fortran source files
	SFiles            []string `json:",omitempty"` // .s source files
	SwigFiles         []string `json:",omitempty"` // .swig files
	SwigCXXFiles      []string `json:",omitempty"` // .swigcxx files
	SysoFiles         []string `json:",omitempty"` // .syso system object files added to package

	// Embedded files
	EmbedPatterns []string `json:",omitempty"` // //go:embed patterns
	EmbedFiles    []string `json:",omitempty"` // files matched by EmbedPatterns

	// Cgo directives
	CgoCFLAGS    []string `json:",omitempty"` // cgo: flags for C compiler
	CgoCPPFLAGS  []string `json:",omitempty"` // cgo: flags for C preprocessor
	CgoCXXFLAGS  []string `json:",omitempty"` // cgo: flags for C++ compiler
	CgoFFLAGS    []string `json:",omitempty"` // cgo: flags for Fortran compiler
	CgoLDFLAGS   []string `json:",omitempty"` // cgo: flags for linker
	CgoPkgConfig []string `json:",omitempty"` // cgo: pkg-config names

	// Dependency information
	Imports   []string          `json:",omitempty"` // import paths used by this package
	ImportMap map[string]string `json:",omitempty"` // map from source import to ImportPath (identity entries omitted)
	Deps      []string          `json:",omitempty"` // all (recursively) imported dependencies

	// Error information
	// Incomplete is above, packed into the other bools
	Error      *PackageError   `json:",omitempty"` // error loading this package (not dependencies)
	DepsErrors []*PackageError `json:",omitempty"` // errors loading dependencies, collected by go list before output

	// Test information
	// If you add to this list you MUST add to p.AllFiles (below) too.
	// Otherwise file name security lists will not apply to any new additions.
	TestGoFiles        []string `json:",omitempty"` // _test.go files in package
	TestImports        []string `json:",omitempty"` // imports from TestGoFiles
	TestEmbedPatterns  []string `json:",omitempty"` // //go:embed patterns
	TestEmbedFiles     []string `json:",omitempty"` // files matched by TestEmbedPatterns
	XTestGoFiles       []string `json:",omitempty"` // _test.go files outside package
	XTestImports       []string `json:",omitempty"` // imports from XTestGoFiles
	XTestEmbedPatterns []string `json:",omitempty"` // //go:embed patterns
	XTestEmbedFiles    []string `json:",omitempty"` // files matched by XTestEmbedPatterns
}

// A PackageError describes an error loading information about a package.
type PackageError struct {
	ImportStack      []string // shortest path from package named on command line to this one
	Pos              string   // position of error
	Err              error    // the error itself
	IsImportCycle    bool     // the error is an import cycle
	Hard             bool     // whether the error is soft or hard; soft errors are ignored in some places
	alwaysPrintStack bool     // whether to always print the ImportStack
}

// Module
//
// copied from src/cmd/go/internal/modinfo/info.go#ModulePublic
type Module struct {
	Path       string       `json:",omitempty"` // module path
	Version    string       `json:",omitempty"` // module version
	Query      string       `json:",omitempty"` // version query corresponding to this version
	Versions   []string     `json:",omitempty"` // available module versions
	Replace    *Module      `json:",omitempty"` // replaced by this module
	Time       *time.Time   `json:",omitempty"` // time version was created
	Update     *Module      `json:",omitempty"` // available update (with -u)
	Main       bool         `json:",omitempty"` // is this the main module?
	Indirect   bool         `json:",omitempty"` // module is only indirectly needed by main module
	Dir        string       `json:",omitempty"` // directory holding local copy of files, if any
	GoMod      string       `json:",omitempty"` // path to go.mod file describing module, if any
	GoVersion  string       `json:",omitempty"` // go version used in module
	Retracted  []string     `json:",omitempty"` // retraction information, if any (with -retracted or -u)
	Deprecated string       `json:",omitempty"` // deprecation message, if any (with -u)
	Error      *ModuleError `json:",omitempty"` // error loading module

	Origin *Origin `json:",omitempty"` // provenance of module
	Reuse  bool    `json:",omitempty"` // reuse of old module info is safe
}

type ModuleError struct {
	Err string // error text
}

// An Origin describes the provenance of a given repo method result.
// It can be passed to CheckReuse (usually in a different go command invocation)
// to see whether the result remains up-to-date.
//
// copied from src/cmd/go/internal/modfetch/codehost/codehost.go#Origin
type Origin struct {
	VCS    string `json:",omitempty"` // "git" etc
	URL    string `json:",omitempty"` // URL of repository
	Subdir string `json:",omitempty"` // subdirectory in repo

	// If TagSum is non-empty, then the resolution of this module version
	// depends on the set of tags present in the repo, specifically the tags
	// of the form TagPrefix + a valid semver version.
	// If the matching repo tags and their commit hashes still hash to TagSum,
	// the Origin is still valid (at least as far as the tags are concerned).
	// The exact checksum is up to the Repo implementation; see (*gitRepo).Tags.
	TagPrefix string `json:",omitempty"`
	TagSum    string `json:",omitempty"`

	// If Ref is non-empty, then the resolution of this module version
	// depends on Ref resolving to the revision identified by Hash.
	// If Ref still resolves to Hash, the Origin is still valid (at least as far as Ref is concerned).
	// For Git, the Ref is a full ref like "refs/heads/main" or "refs/tags/v1.2.3",
	// and the Hash is the Git object hash the ref maps to.
	// Other VCS might choose differently, but the idea is that Ref is the name
	// with a mutable meaning while Hash is a name with an immutable meaning.
	Ref  string `json:",omitempty"`
	Hash string `json:",omitempty"`

	// If RepoSum is non-empty, then the resolution of this module version
	// failed due to the repo being available but the version not being present.
	// This depends on the entire state of the repo, which RepoSum summarizes.
	// For Git, this is a hash of all the refs and their hashes.
	RepoSum string `json:",omitempty"`
}
