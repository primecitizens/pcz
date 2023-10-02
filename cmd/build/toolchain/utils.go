// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package toolchain

import (
	"bytes"
	"fmt"
	"os"
)

func NewCompiledPackages() *CompiledPackages {
	return &CompiledPackages{
		compiled: map[string]string{},
	}
}

type CompiledPackages struct {
	compiled map[string]string

	importcfg bytes.Buffer
}

// Add a built archive (.a file) for the importpath.
func (cp *CompiledPackages) Add(importpath, actualImportpath, ar string) {
	if oldAR, ok := cp.compiled[importpath]; ok {
		fmt.Fprintf(os.Stderr, "duplicate importpath %q: old = %s ; new = %s\n", importpath, oldAR, ar)
		// panic(fmt.Errorf("duplicate importpath %q: old = %s ; new = %s", importpath, oldAR, ar))
	}

	cp.compiled[importpath] = ar
	_, _ = fmt.Fprintf(&cp.importcfg, "importmap %s=%s\n", importpath, actualImportpath)
	_, _ = fmt.Fprintf(&cp.importcfg, "packagefile %s=%s\n", actualImportpath, ar)
}

func (cp *CompiledPackages) Importcfg() []byte {
	return cp.importcfg.Bytes()
}

func (cp *CompiledPackages) Lookup(importpath string) (string, bool) {
	ret, ok := cp.compiled[importpath]
	return ret, ok
}
