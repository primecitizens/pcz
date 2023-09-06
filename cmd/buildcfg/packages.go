// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package buildcfg

import (
	"strings"
)

type Transformer interface {
	Transform(p *Package, deps []Package) error
}

type PczStdTransformer struct{}

func (t *PczStdTransformer) Transform(p *Package, deps []Package) error {
	p.ImportComment = t.renameImportpath(p.ImportPath)
	for i := range p.Deps {
		p.Deps[i] = t.renameImportpath(p.Deps[i])
	}

	return nil
}

func (r *PczStdTransformer) renameImportpath(importpath string) string {
	short, isSTD := strings.CutPrefix(importpath, "github.com/primecitizens/std/")
	if !isSTD {
		return importpath
	}

	switch short {
	case "rt0":
		return "rt0"
	case "core/abi":
		return "internal/abi"
	case "core/mark/internal/sys":
		return "runtime/internal/sys"
	case "core/mark/internal/atomic":
		return "runtime/internal/atomic"
	}

	switch {
	case strings.HasPrefix(short, "builtin/"):
		return "std" + strings.TrimPrefix(short, "builtin/")
	}

	return short
}
