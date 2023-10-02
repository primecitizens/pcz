// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package toolchain

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"unsafe"
)

// List runs go list for the target package.
func (t *Toolchain) List(dir string) (pkgs []*Package, err error) {
	var buf bytes.Buffer
	if err = handleErr(
		t.run(dir, nil, &buf, nil,
			t.binGo, "list", "-tags", strings.Join(t.opts.GetBuildTags(), ","), "-json", "-deps", "./",
		),
	); err != nil {
		return
	}

	// key: Package.ImportPath
	ignore := make(map[string]bool)

	// key: Package.ImportComment
	index := make(map[string]*Package)

	omitIgnoredPackages := func(importpaths []string) []string {
		return filterOut(importpaths, func(dep string) bool {
			return ignore[dep]
		})
	}

	dec := json.NewDecoder(&buf)
	for dec.More() {
		pkg := new(Package)
		if err = handleErr(dec.Decode(pkg)); err != nil {
			return
		}

		if pkg.Standard {
			ignore[pkg.ImportPath] = true
			continue
		}

		// normalize package details

		if pkg.Name == "main" {
			pkg.ImportPath = "main"
			pkg.ImportComment = "main"
		}

		pkg.Deps = omitIgnoredPackages(pkg.Deps)
		pkg.Imports = omitIgnoredPackages(pkg.Imports)
		pkg.SFiles = filterOut(pkg.SFiles, func(sfile string) bool {
			return strings.HasSuffix(sfile, "_test.s")
		})

		if err = handleErr(t.tr.Transform(pkg, pkgs)); err != nil {
			return
		}

		if existing := index[pkg.ImportComment]; existing != nil {
			// merge duplicate package
			for _, file := range pkg.GoFiles {
				existing.GoFiles = append(
					existing.GoFiles, filepath.Join(pkg.Dir, file),
				)
			}

			for _, file := range pkg.SFiles {
				existing.SFiles = append(
					existing.SFiles, filepath.Join(pkg.Dir, file),
				)
			}
		} else {
			index[pkg.ImportComment] = pkg
			pkgs = append(pkgs, pkg)
		}
	}

	return
}

type Transformer interface {
	Transform(p *Package, deps []*Package) error
}

type PczStdTransformer struct{}

func (t *PczStdTransformer) Transform(p *Package, deps []*Package) (err error) {
	if p.ImportComment, err = handleErr2(t.renameImportpath(p)); err != nil {
		return
	}

	return nil
}

func (r *PczStdTransformer) renameImportpath(p *Package) (importpath string, err error) {
	importpath = p.ImportPath
	importpath, isSTD := strings.CutPrefix(importpath, "github.com/primecitizens/std/")
	if !isSTD {
		if len(p.ImportComment) != 0 {
			importpath = p.ImportComment
		}

		return
	}

	if slices.Contains(p.GoFiles, "doc.go") {
		var file []byte
		if file, err = handleErr2(os.ReadFile(filepath.Join(p.Dir, "doc.go"))); err != nil {
			return
		}

		lines := strings.Split(unsafe.String(unsafe.SliceData(file), len(file)), "\n")
		for _, line := range lines {
			pkg, ok := strings.CutPrefix(line, "//pcz:importpath ")
			if ok {
				return strings.TrimSpace(pkg), nil
			}
		}
	}

	switch {
	case strings.HasPrefix(importpath, "builtin/"):
		return p.Name, nil
	default:
		return importpath, nil
	}
}

func filterOut[T any](slice []T, omit func(T) bool) []T {
	if len(slice) == 0 {
		return slice
	}

	var i int
	for j, s := range slice {
		if omit(s) {
			continue
		}

		slice[i] = slice[j]
		i++
	}

	return slice[:i]
}
