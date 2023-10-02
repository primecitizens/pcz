// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package toolchain

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"time"

	"github.com/primecitizens/pcz/cmd/utils/fsutils"
	"gopkg.in/yaml.v3"
)

func (t *Toolchain) Build(pkgs []*Package, output string) (err error) {
	if len(pkgs) == 0 {
		return fmt.Errorf("nothing to build")
	}

	// the last package is the last package to build
	last := pkgs[len(pkgs)-1]
	if last.Module == nil || len(last.Module.GoVersion) == 0 {
		return fmt.Errorf("missing go version in go.mod: %v", last.Module)
	}
	lang := last.Module.GoVersion

	var deps *CompiledPackages

	built, err := handleErr2(t.build(lang, pkgs, deps))
	if err != nil {
		return
	}

	if last.Name == "main" { // need to link
		if err = handleErr(t.Link(last, built, output)); err != nil {
			return
		}
	} else {
		ar, ok := built.Lookup(last.ImportPath)
		if !ok {
			panic(fmt.Errorf("importpath not found: %q", last.ImportPath))
		}
		if err = handleErr(os.Rename(ar, output)); err != nil {
			return
		}
	}

	return
}

func (t *Toolchain) build(
	lang string, // flag value for for compile -lang,
	pkgs []*Package,
	deps *CompiledPackages,
) (built *CompiledPackages, err error) {
	if deps != nil {
		built = deps
	} else {
		built = NewCompiledPackages()
	}

	var (
		startTime = time.Now()
		buf       bytes.Buffer
		enc       = yaml.NewEncoder(&buf)
	)
	enc.SetIndent(2)

	cacheChanged := make(map[string]bool)
	for _, p := range pkgs {
		if _, ok := built.Lookup(p.ImportPath); ok {
			continue
		}

		// TODO: merge packages using the same ImportComment

		if len(p.EmbedFiles) != 0 {
			// expecting the "embed" package
			if _, ok := built.Lookup("embed"); !ok {
				var ar string
				if ar, err = handleErr2(t.buildPesudoEmbedPkg(lang, p)); err != nil {
					return
				}

				built.Add("embed", "embed", ar)
			}
		}

		var (
			cacheFile = t.getPathInBuildDir(p, "cache.yaml")
			cacheData []byte
			dir       = filepath.Dir(cacheFile)
			old       CacheRecord
			newcache  *CacheRecord
		)
		cacheData, err = os.ReadFile(cacheFile)
		if err == nil {
			if err = handleErr(yaml.Unmarshal(cacheData, &old)); err != nil {
				return
			}
		}

		if newcache, err = handleErr2(t.CreateCacheRecord(startTime, dir, p)); err != nil {
			return
		}

		if t.useCache && old.SourceUnchanged(newcache) {
			for _, dep := range p.Deps {
				if cacheChanged[dep] {
					goto BUILD
				}
			}

			if old.Build.Find("pkg.a").VerifyUnchanged(old.Build.Dir) {
				built.Add(p.ImportPath, p.ImportComment, filepath.Join(old.Build.Dir, "pkg.a"))
				continue
			}
		}

	BUILD:
		cacheChanged[p.ImportPath] = true

		var (
			importcfg       = t.getPathInBuildDir(p, "importcfg")
			symabi          string
			asmhdr, archive string
			objFiles        []string
		)

		if err = handleErr(fsutils.WriteReadOnly(importcfg, built.Importcfg())); err != nil {
			return
		} else if symabi, err = handleErr2(t.Symabis(p)); err != nil {
			return
		} else if asmhdr, archive, err = t.compile(lang, p, importcfg, symabi); err != nil {
			err = handleErr(err)
			return
		} else if objFiles, err = handleErr2(t.asm(p, symabi, asmhdr)); err != nil {
			return
		} else if archive, err = handleErr2(t.Pack(archive, objFiles...)); err != nil {
			return
		}

		// accumulate importcfg
		// ImportComment is the real ImportPath
		built.Add(p.ImportPath, p.ImportComment, archive)

		files := append(
			slices.DeleteFunc(
				[]string{importcfg, symabi, asmhdr, archive},
				func(s string) bool { return len(s) == 0 },
			),
			objFiles...,
		)
		for i, f := range files {
			files[i] = filepath.Base(f)
		}

		if err = handleErr(newcache.Build.Add(time.Now(), dir, files...)); err != nil {
			return
		}

		// write cache record
		if err = handleErr(enc.Encode(newcache)); err != nil {
			return
		} else if err = handleErr(fsutils.WriteReadOnly(cacheFile, buf.Bytes())); err != nil {
			return
		}

		buf.Reset()
	}

	return
}

func (t *Toolchain) buildPesudoEmbedPkg(lang string, p *Package) (ar string, err error) {
	// TODO: check whether the package uses embed.FS, if so, rewrite the code
	// to import the io/fs/memfs and generate memfs implementation.

	pkg := &Package{
		ImportPath:    "embed",
		ImportComment: "embed",
		Dir:           "", // see below
		GoFiles:       []string{"embed.go"},
		Module:        p.Module,
	}

	src := t.getPathInBuildDir(pkg, "embed.go")
	pkg.Dir = filepath.Dir(src)

	importcfg := t.getPathInBuildDir(pkg, "importcfg")
	if err = handleErr(fsutils.WriteReadOnly(src, []byte("package embed\n"))); err != nil {
		return
	} else if err = handleErr(fsutils.WriteReadOnly(importcfg, nil)); err != nil {
		return
	} else if _, ar, err = t.compile(lang, pkg, importcfg, ""); err != nil {
		err = handleErr(err)
		return
	}
	return
}
