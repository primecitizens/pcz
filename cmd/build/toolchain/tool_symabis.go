// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package toolchain

import (
	"os"
	"path/filepath"

	"github.com/primecitizens/pcz/cmd/utils/fsutils"
)

func (t *Toolchain) Symabis(p *Package) (outputfile string, err error) {
	if len(p.SFiles) == 0 {
		return
	}

	// -gensymabis will look for "go_asm.h"
	asmhdrFile := t.getPathInBuildDir(p, "go_asm.h")
	if err = handleErr(fsutils.WriteReadOnly(asmhdrFile, nil)); err != nil {
		return
	}

	outputfile = t.getPathInBuildDir(p, "symabis")
	err = fsutils.TemporaryWritable(
		func() error {
			return t.run(p.Dir, nil, os.Stdout, os.Stderr,
				t.createASMArgs(p, p.SFiles,
					"-I", filepath.Dir(asmhdrFile), "-gensymabis", "-o", outputfile,
				)...,
			)
		},
		outputfile,
	)

	return
}
