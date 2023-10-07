// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package toolchain

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// NOTE: ImportComment is the real ImportPath

func NewToolchain(opts *Options) *Toolchain {
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

		useCache: !opts.NoCache,

		goos:   opts.GOOS(),
		goarch: opts.GOARCH(),

		binGo:      filepath.Join(opts.Context.GOROOT, "bin", "go"),
		binComplie: filepath.Join(toolsDir, "compile"),
		binAsm:     filepath.Join(toolsDir, "asm"),
		binLink:    filepath.Join(toolsDir, "link"),
		binPack:    filepath.Join(toolsDir, "pack"),
		binVet:     filepath.Join(toolsDir, "vet"),

		ver: "", // see below

		buildDir: buildDir,

		env: opts.AppendEnviron(os.Environ()),
	}

	tc.ver = tc.GoVersion()
	return tc
}

type Toolchain struct {
	opts *Options
	tr   Transformer

	useCache bool

	goos   string
	goarch string

	binGo      string
	binComplie string
	binAsm     string
	binLink    string
	binPack    string
	binVet     string

	ver string

	buildDir string

	env []string
}

func (t *Toolchain) getPathInBuildDir(p *Package, filename string) string {
	var ret string
	if t.opts.Trimpath {
		ret = filepath.Join(
			t.buildDir, "trimpath", t.ver, t.goos+"_"+t.goarch, p.ImportComment, filename,
		)
	} else {
		// ${build_dir}/go1.21.2/js/wasm/package-real-import-path/filename
		ret = filepath.Join(
			t.buildDir, t.ver, t.goos+"_"+t.goarch, p.ImportComment, filename,
		)
	}
	if err := os.MkdirAll(filepath.Dir(ret), 0750); err != nil {
		panic(err)
	}

	return ret
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
		return fmt.Errorf("version of gc and go mismatch: go = %q, gc = %q", gover, v)
	} else if v := tc.toolVer(tc.binAsm); v != gover {
		return fmt.Errorf("version of asm and go mismatch: go = %q, asm = %q", gover, v)
	} else if v := tc.toolVer(tc.binLink); v != gover {
		return fmt.Errorf("version of link and go mismatch: go = %q, link = %q", gover, v)
	} else if v := tc.toolVer(tc.binVet); v != gover {
		return fmt.Errorf("version of vet and go mismatch: go = %q, vet = %q", gover, v)
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
