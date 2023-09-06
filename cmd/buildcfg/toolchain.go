// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package buildcfg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
)

func NewToolchain(opts *BuildOptions) *Toolchain {
	toolsDir := GOTOOLDIR
	if len(toolsDir) == 0 {
		toolsDir = filepath.Join(opts.Context.GOROOT, "pkg", "tool", runtime.GOOS+"_"+runtime.GOARCH)
	}

	buildDir := opts.BuildDir
	if len(buildDir) == 0 {
		buildDir = ".pcz/build"
	}

	buildDir, err := filepath.Abs(buildDir)
	if err != nil {
		panic(err)
	}

	tc := &Toolchain{
		opts: opts,
		tr:   &PczStdTransformer{},

		binGo:      filepath.Join(opts.Context.GOROOT, "bin", "go"),
		binComplie: filepath.Join(toolsDir, "compile"),
		binAsm:     filepath.Join(toolsDir, "asm"),
		binLink:    filepath.Join(toolsDir, "link"),
		binPack:    filepath.Join(toolsDir, "pack"),

		ver: "", // see below

		buildDir: buildDir,

		env: opts.AppendEnviron(os.Environ()),
	}

	tc.ver = tc.GoVersion()
	return tc
}

type Toolchain struct {
	opts *BuildOptions
	tr   Transformer

	binGo      string
	binComplie string
	binAsm     string
	binLink    string
	binPack    string

	ver string

	buildDir string

	env []string
}

func (t *Toolchain) Build(pkg, output string) (err error) {
	pkgs, err := t.List(pkg)
	if err != nil {
		panic(err)
	}

	if len(pkgs) == 0 {
		panic(fmt.Errorf("nothing to build for %q", pkg))
	}

	// last package if the pkg
	last := &pkgs[len(pkgs)-1]
	if last.Module == nil || len(last.Module.GoVersion) == 0 {
		panic(fmt.Errorf("missing go version in go.mod: %v", last.Module))
	}
	lang := last.Module.GoVersion

	var deps *CompiledPackages

	built, err := t.build(lang, pkgs, deps)
	if err != nil {
		panic(err)
	}
	defer built.Remove()

	if last.Name == "main" { // need to link
		err = t.link(last, built, output)
		if err != nil {
			panic(err)
		}
	} else {
		ar, ok := built.Lookup(last.ImportPath)
		if !ok {
			panic(fmt.Errorf("importpath not found: %q", last.ImportPath))
		}
		err = os.Rename(ar, output)
		if err != nil {
			panic(err)
		}
	}

	return
}

func (t *Toolchain) link(p *Package, deps *CompiledPackages, output string) (err error) {
	importcfgFile := t.getPathInBuildDir(p, "importcfg.link")
	file, err := os.OpenFile(importcfgFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0640)
	if err != nil {
		panic(err)
	}

	func() {
		defer file.Close()

		var line string
		cfg := string(deps.Importcfg())
		for len(cfg) != 0 {
			line, cfg, _ = strings.Cut(cfg, "\n")
			if strings.HasPrefix(line, "importmap") {
				continue
			}

			_, err = file.WriteString(line)
			if err != nil {
				panic(err)
			}

			_, err = file.WriteString("\n")
			if err != nil {
				panic(err)
			}
		}

		// TODO: set modinfo
		bi := debug.BuildInfo{
			GoVersion: t.ver,
			Path:      "",
			Main:      debug.Module{},
			Deps:      []*debug.Module{},
			Settings:  []debug.BuildSetting{},
		}
		_ = bi
		// _, err = fmt.Fprintf(file, "modinfo %q\n", ModInfoData(bi.String()))
		// if err != nil {
		// 	panic(err)
		// }
	}()

	archive, ok := deps.Lookup(p.ImportPath)
	if !ok {
		panic(fmt.Errorf("importpath not found: %q", p.ImportPath))
	}

	ldflags := t.opts.AppendLDFlags(
		t.binLink,
		"-o", output,
		"-importcfg", importcfgFile,
		// TODO: generate buildid
		// "-buildid", "",
	)
	err = t.run(p.Dir, nil, os.Stdout, os.Stderr,
		append(ldflags, archive)...,
	)
	if err != nil {
		panic(err)
	}

	return err
}

func (t *Toolchain) build(
	lang string, // flag value for for compile -lang,
	pkgs []Package,
	deps *CompiledPackages,
) (built *CompiledPackages, err error) {
	if deps != nil {
		built = deps
	} else {
		built = NewCompiledPackages()
	}

	for i := range pkgs {
		p := &pkgs[i]

		importcfgFile := t.getPathInBuildDir(p, "importcfg")
		err = os.WriteFile(importcfgFile, built.Importcfg(), 0644)
		if err != nil {
			panic(err)
		}

		symabiFile, err := t.symabis(p)
		if err != nil {
			panic(err)
		}

		asmhdrFile, archiveFile, err := t.compile(lang, p, importcfgFile, symabiFile)
		if err != nil {
			panic(err)
		}

		objFiles, err := t.asm(p, symabiFile, asmhdrFile)
		if err != nil {
			panic(err)
		}

		archiveFile, err = t.Pack(archiveFile, objFiles...)
		if err != nil {
			panic(err)
		}

		// accumulate importcfg
		// ImportComment is the real ImportPath
		built.Add(p.ImportPath, p.ImportComment, archiveFile)
	}

	return
}

func (t *Toolchain) getPathInBuildDir(p *Package, s ...string) string {
	ret := filepath.Join(append([]string{t.buildDir, p.ImportComment}, s...)...)
	if err := os.MkdirAll(filepath.Dir(ret), 0750); err != nil {
		panic(err)
	}

	return ret
}

func (t *Toolchain) symabis(p *Package) (outputfile string, err error) {
	if len(p.SFiles) == 0 {
		return
	}

	// -gensymabis will look for "go_asm.h"
	asmhdrFile := t.getPathInBuildDir(p, "go_asm.h")
	err = os.WriteFile(asmhdrFile, nil, 0640)
	if err != nil {
		panic(err)
	}

	outputfile = t.getPathInBuildDir(p, "symabis")
	err = t.run(p.Dir, nil, os.Stdout, os.Stderr,
		t.createASMArgs(p, p.SFiles,
			"-I", filepath.Dir(asmhdrFile), "-gensymabis", "-o", outputfile,
		)...,
	)
	if err != nil {
		panic(err)
	}

	return
}

// asmhdrFile is the path to "go_asm.h"
func (t *Toolchain) compile(
	lang string, // flag value to -lang
	p *Package, //
	importcfgFile string, // flag value to -importcfg
	symabiFile string, // flag value to -symabis
) (asmhdrFile, archiveFile string, err error) {
	archiveFile = t.getPathInBuildDir(p, p.Name+".a")
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
		"-importcfg", importcfgFile,
		"-lang", "go" + p.Module.GoVersion,
		"-goversion", t.ver,
		// "-trimpath", pkg.Module.Dir + "=>",
		"-o", archiveFile,
	}
	if len(symabiFile) != 0 {
		asmhdrFile = t.getPathInBuildDir(p, "go_asm.h")
		flags = append(flags,
			"-symabis", symabiFile,
			"-asmhdr", asmhdrFile,
		)
	}

	flags = t.opts.AppendGCFlags(flags)

	// although we can use pkg.GoFiles directly, but in that case, generated archives
	// are not friendly for debugging
	//
	// replace the for loop with following line if we changed our mind
	//flags = append(flags, pkg.GoFiles...)
	for _, file := range p.GoFiles {
		flags = append(flags, filepath.Join(p.Dir, file))
	}

	err = t.run(p.Dir, nil, os.Stdout, os.Stderr, flags...)
	if err != nil {
		panic(err)
	}

	return
}

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

		err = t.run(p.Dir, nil, os.Stdout, os.Stderr,
			t.createASMArgs(p, nil,
				"-I", filepath.Dir(asmhdrFile),
				"-o", output,
				file,
			)...,
		)
		if err != nil {
			return
		}

		objFiles = append(objFiles, output)
	}

	return
}

func (t *Toolchain) createASMArgs(p *Package, sfiles []string, args ...string) []string {
	asmFlags := t.opts.AppendASMFlags([]string{
		t.binAsm,
		"-I", p.Dir,
		"-I", GoDefaultIncludeDir,
		"-e", // no limit on number of errors reported
		"-shared",
		"-compiling-runtime",  // always compile as runtime (so that we can use ABIInternal)
		"-p", p.ImportComment, // ImportComment is the real ImportPath
		// "-trimpath", pkg.Module.Dir + "=>",
	})

	asmFlags = append(asmFlags, args...)
	for _, sfile := range sfiles {
		asmFlags = append(asmFlags, filepath.Join(p.Dir, sfile))
	}
	return asmFlags
}

func (t *Toolchain) Pack(archiveFile string, objFiles ...string) (newArchiveFile string, err error) {
	newArchiveFile = archiveFile
	if len(objFiles) == 0 {
		return
	}

	err = t.run("", nil, os.Stdout, os.Stderr,
		append([]string{t.binPack, "r", archiveFile}, objFiles...)...,
	)
	if err != nil {
		panic(err)
	}
	return
}

// List runs go list for the target package.
func (t *Toolchain) List(dir string) (pkgs []Package, err error) {
	var buf bytes.Buffer
	err = t.run(dir, nil, &buf, nil,
		t.binGo, "list", "-tags", strings.Join(t.opts.GetBuildTags(), ","), "-json", "-deps", "./",
	)
	if err != nil {
		panic(err)
	}

	ignore := make(map[string]bool)

	omitIgnoredPackages := func(importpaths []string) []string {
		return filterOut(importpaths, func(dep string) bool {
			return ignore[dep]
		})
	}

	dec := json.NewDecoder(&buf)
	for dec.More() {
		var pkg Package
		err = dec.Decode(&pkg)
		if err != nil {
			panic(err)
		}

		if pkg.Standard {
			switch pkg.ImportPath {
			case "runtime", // we don't use official runtime
				"unsafe", "builtin": // ignore identity packages
				ignore[pkg.ImportPath] = true
				continue
			}

			switch {
			case strings.HasPrefix(pkg.ImportPath, "internal/"),
				strings.HasPrefix(pkg.ImportPath, "runtime/"): // ignore runtime dependencies
				ignore[pkg.ImportPath] = true
				continue
			default:
				// deny official std usage
				panic(fmt.Sprintf("go std package %q is not allowed", pkg.ImportPath))
			}

			panic("unreachable")
		}

		// normalize package details

		if pkg.Name == "main" {
			pkg.ImportPath = "main"
		}

		pkg.Deps = omitIgnoredPackages(pkg.Deps)
		pkg.Imports = omitIgnoredPackages(pkg.Imports)
		pkg.SFiles = filterOut(pkg.SFiles, func(sfile string) bool {
			return strings.HasSuffix(sfile, "_test.s")
		})

		err = t.tr.Transform(&pkg, pkgs)
		if err != nil {
			panic(err)
		}

		pkgs = append(pkgs, pkg)
	}

	return
}

// GoVersion returns the version string as printed by `go version`
func (t *Toolchain) GoVersion() string {
	var sb strings.Builder
	err := t.run("", nil, &sb, nil,
		t.binGo, "version",
	)
	if err != nil {
		panic(err)
	}

	parts := strings.Split(sb.String(), " ")
	return strings.TrimSpace(parts[len(parts)-2])
}

// VerifyVersion returns error if gover is different from versions reported by
// compile/asm/link.
func (tc *Toolchain) VerifyVersion(gover string) error {
	if v := tc.toolVer(tc.binComplie); v != gover {
		return fmt.Errorf("version of gc and go missmatch: go = %q, gc = %q", gover, v)
	}
	if v := tc.toolVer(tc.binAsm); v != gover {
		return fmt.Errorf("version of asm and go missmatch: go = %q, gc = %q", gover, v)
	}
	if v := tc.toolVer(tc.binLink); v != gover {
		return fmt.Errorf("version of link and go missmatch: go = %q, gc = %q", gover, v)
	}

	return nil
}

func (t *Toolchain) toolVer(tool string) string {
	var sb strings.Builder
	err := t.run("", nil, &sb, nil,
		tool, "-V=full",
	)
	if err != nil {
		panic(err)
	}

	parts := strings.Split(sb.String(), " ")
	return strings.TrimSpace(parts[len(parts)-1])
}

func (t *Toolchain) run(
	cwd string,
	stdin io.Reader,
	stdout, stderr io.Writer,
	args ...string,
) (err error) {
	var sb strings.Builder
	if stderr == nil {
		stderr = &sb
	}

	if t.opts.Debug {
		fmt.Fprintf(os.Stderr, "exec: %v\n", args)
	}

	cmd := &exec.Cmd{
		Path:   args[0],
		Args:   args,
		Dir:    cwd,
		Env:    t.env,
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
	}

	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to run %q: %w", args[0], err)
	}

	err = cmd.Wait()
	if err != nil {
		if sb.Len() > 0 {
			return fmt.Errorf("%q error: %w\n\n%s", args[0], err, sb.String())
		}

		return fmt.Errorf("%s error: %w", args[0], err)
	}

	return nil
}
