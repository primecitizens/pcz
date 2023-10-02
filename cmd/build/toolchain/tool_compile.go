// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package toolchain

import (
	"encoding/json"
	"os"
	"path"
	"path/filepath"

	"github.com/primecitizens/pcz/cmd/utils/fsutils"
)

func (t *Toolchain) appendTrimpathFlag(p *Package, flags []string) []string {
	version := p.Module.Version
	if len(version) == 0 {
		version = "0.0.0"
	}

	modulePath := p.Module.Path
	if modulePath == "github.com/primecitizens/std" {
		modulePath = p.ImportComment
	}

	return append(flags,
		"-trimpath",
		p.Dir+"=>"+modulePath+"@"+version+";"+t.buildDir+"=>",
	)
}

// asmhdrFile is the path to "go_asm.h"
func (t *Toolchain) compile(
	lang string, // flag value to -lang
	p *Package, //
	importcfg string, // flag value to -importcfg
	symabi string, // flag value to -symabis
) (
	asmhdrFile, archiveFile string, err error,
) {
	archiveFile = t.getPathInBuildDir(p, "pkg.a")
	flags := []string{
		t.binComplie,
		"-nolocalimports",     //
		"-L",                  // show full file names in error messages
		"-std",                // building std
		"-e",                  // no limit on number of errors reported
		"-+",                  // always compile as runtime (to use ABIInternal/ABI0 outside runtime pakcage)
		"-wb=false",           // disable write barrier
		"-pack",               // create archive of objects
		"-p", p.ImportComment, // ImportComment is the real ImportPath
		"-importcfg", importcfg,
		"-lang", "go" + p.Module.GoVersion,
		"-goversion", t.ver,
		"-o", archiveFile,
	}

	if t.opts.Trimpath {
		flags = t.appendTrimpathFlag(p, flags)
	}

	if len(symabi) != 0 {
		asmhdrFile = t.getPathInBuildDir(p, "go_asm.h")
		flags = append(flags,
			"-symabis", symabi,
			"-asmhdr", asmhdrFile,
		)
	}

	if len(p.EmbedFiles) != 0 {
		type Schema struct {
			// key: pattern
			// value: matched file paths
			Patterns map[string][]string `json:"Patterns"`
			// key: path as matched by pattern
			// value: absolute local filesystem path
			Files map[string]string `json:"Files"`
		}

		embedcfg := t.getPathInBuildDir(p, "embedcfg")
		cfg := Schema{
			Patterns: map[string][]string{},
			Files:    map[string]string{},
		}
		for _, file := range p.EmbedFiles {
			for _, ptn := range p.EmbedPatterns {
				if ok, _ := path.Match(ptn, file); !ok {
					continue
				}

				cfg.Patterns[ptn] = append(cfg.Patterns[ptn], file)
				break
			}
			cfg.Files[file] = filepath.Join(p.Dir, file)
		}

		var data []byte
		if data, err = handleErr2(json.Marshal(&cfg)); err != nil {
			return
		} else if err = handleErr(fsutils.WriteReadOnly(embedcfg, data)); err != nil {
			return
		}

		flags = append(flags, "-embedcfg", embedcfg)
	}

	flags = t.opts.AppendGCFlags(flags)

	// although we can use pkg.GoFiles directly (by setting cwd to pkg.Dir)
	// but in that case, generated archives are not friendly for debugging
	//
	// replace the for loop with following line if we changed our mind
	//flags = append(flags, pkg.GoFiles...)
	for _, file := range p.GoFiles {
		if !filepath.IsAbs(file) {
			file = filepath.Join(p.Dir, file)
		}

		flags = append(flags, file)
	}

	err = fsutils.TemporaryWritable(
		func() error {
			return t.run(
				p.Dir, nil, os.Stdout, os.Stderr,
				flags...,
			)
		},
		asmhdrFile, archiveFile,
	)

	return
}
