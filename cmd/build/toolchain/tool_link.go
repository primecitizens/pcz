// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package toolchain

import (
	"bytes"
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	"github.com/primecitizens/pcz/cmd/utils/fsutils"
)

func (t *Toolchain) Link(p *Package, deps *CompiledPackages, output string) (err error) {
	var (
		buf bytes.Buffer
	)

	for cfg := string(deps.Importcfg()); len(cfg) != 0; {
		var line string
		line, cfg, _ = strings.Cut(cfg, "\n")
		if strings.HasPrefix(line, "importmap") {
			continue
		}

		buf.WriteString(line)
		buf.WriteString("\n")
	}

	// TODO: set modinfo in the importcfg file
	bi := debug.BuildInfo{
		GoVersion: t.ver,
		Path:      "",
		Main:      debug.Module{},
		Deps:      []*debug.Module{},
		Settings:  []debug.BuildSetting{},
	}
	_ = bi
	// _, err = handleErr(fmt.Fprintf(file, "modinfo %q\n", ModInfoData(bi.String())))
	// if err != nil {
	// 	return
	// }

	importcfgFile := t.getPathInBuildDir(p, "importcfg.link")
	if err = handleErr(fsutils.WriteReadOnly(importcfgFile, buf.Bytes())); err != nil {
		return
	}

	archive, ok := deps.Lookup(p.ImportPath)
	if !ok {
		panic(fmt.Errorf("importpath not found: %q", p.ImportPath))
	}

	ldflags := t.opts.AppendLDFlags(
		t.binLink,
		"-o", output,
		"-importcfg", importcfgFile,
		// TODO: generate buildid
		//       why? to work with go cache? seems not so useful when the runtime is different
		// "-buildid", "",
	)
	err = handleErr(
		t.run(p.Dir, nil, os.Stdout, os.Stderr, append(ldflags, archive)...),
	)
	if err != nil {
		return
	}

	return err
}
