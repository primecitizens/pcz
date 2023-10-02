// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package toolchain

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/primecitizens/pcz/cmd/utils/fsutils"
)

func (t *Toolchain) asm(
	p *Package,
	symabiFile string,
	asmhdrFile string,
) (objFiles []string, err error) {
	if len(p.SFiles) == 0 {
		return
	}

	for _, file := range p.SFiles {
		output := t.getPathInBuildDir(p, strings.TrimSuffix(file, ".s")+".o")
		err = fsutils.TemporaryWritable(
			func() error {
				return t.run(p.Dir, nil, os.Stdout, os.Stderr,
					t.createASMArgs(p, nil,
						"-I", filepath.Dir(asmhdrFile),
						"-o", output,
						file,
					)...,
				)
			},
			output,
		)
		if err != nil {
			return
		}

		objFiles = append(objFiles, output)
	}

	return
}

func (t *Toolchain) createASMArgs(p *Package, sfiles []string, args ...string) []string {
	flags := t.opts.AppendASMFlags(
		t.binAsm,
		"-I", p.Dir,
		"-I", GoDefaultIncludeDir,
		"-e", // no limit on number of errors reported
		"-shared",
		"-compiling-runtime", // always compile as runtime (so that we can use ABIInternal)
		"-p", p.ImportComment,
	)

	if t.opts.Trimpath {
		flags = t.appendTrimpathFlag(p, flags)
	}

	flags = append(flags, args...)
	for _, sfile := range sfiles {
		if !filepath.IsAbs(sfile) {
			sfile = filepath.Join(p.Dir, sfile)
		}

		flags = append(flags, sfile)
	}
	return flags
}
