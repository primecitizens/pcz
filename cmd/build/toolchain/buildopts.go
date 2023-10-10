// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package toolchain

import (
	"fmt"
	"go/build"
	"strings"
)

type Options struct {
	Context *build.Context

	Debug bool `cli:"debug,#debug build process"`

	BuildDir string `cli:"build-dir,def=.pcz/build,#set directory to store intermediate build outputs"`

	Mode       string // `cli:"mode,#the same as go build -buildmode"`
	EnableMSan bool   // `cli:"msan,#enable memory san interface"`
	EnableASan bool   // `cli:"asan,#enable address san interface"`
	EnableRace bool   // `cli:"race,#enable race detector"`
	NoCache    bool   `cli:"no-cache,#build without cache"`
	Trimpath   bool   `cli:"trimpath,#remove all filesystem paths from the resulting executable"`

	Defs        []string `cli:"X|define,#link-time string value definition of the form importpath.name=value"`
	EntrySymbol string   `cli:"entry,def=rt0,#set entry symbol name"`
	Platform    Platform `cli:"p|platform,#\"os/arch\" pair"`
	Tags        []string `cli:"t|tag,#add build tags"`

	LDFlags  []string `cli:"L|ldflag,#add flags to go tool link"`
	GCFlags  []string `cli:"G|gcflag,#add flags to go tool compile (not including gccgo)"`
	ASMFlags []string `cli:"A|asmflag,#add flags to go tool asm"`
}

// GOOS returns a valid GOOS value suggested by the .Platform or GOOS
// environment variable.
func (opts *Options) GOOS() string {
	return opts.Platform.GOOS(opts.Context.GOOS)
}

// GOARCH returns a valid GOARCH value suggested by the .Platform or GOARCH
// environment variable.
func (opts *Options) GOARCH() string {
	return opts.Platform.GOARCH(opts.Context.GOARCH)
}

func (opts *Options) GetBuildTags() []string {
	tags := append(
		[]string{
			// TODO: currently we only supports go1.21 toolchain
			"go1.21",
			"pcz", // build tag used to make pcz std work along with official go command
		},
		opts.Tags...,
	)

	switch os := opts.Platform.OS(); os {
	case "efi":
		tags = append(tags, "noos", "efi")
	case "noos":
		tags = append(tags, "noos")
	case "nodejs", "deno",
		"oculus",
		"webview-android",
		"opera", "opera-android",
		"chrome", "chrome-android",
		"firefox", "firefox-android",
		"safari", "safari-ios",
		"edge", "ie":
		tags = append(tags, "js_"+os)
	}

	return tags
}

func (opts *Options) AppendEnviron(env []string) []string {
	env = append(env, "GOOS="+opts.GOOS())

	switch goarch := opts.GOARCH(); goarch {
	case "386":
		env = append(env, "GOARCH=386", "GO386="+opts.Platform.GO386())
	case "amd64":
		env = append(env, "GOARCH=amd64", "GOAMD64="+opts.Platform.GOAMD64())
	case "arm":
		env = append(env, "GOARCH=arm", "GOARM="+opts.Platform.GOARM())
	case "arm64":
		env = append(env, "GOARCH=arm64")
	case "mips64le":
		env = append(env, "GOARCH=mips64le", "GOMIPS64="+opts.Platform.GOMIPS64())
	case "mips64":
		env = append(env, "GOARCH=mips64", "GOMIPS64="+opts.Platform.GOMIPS64())
	case "mipsle":
		env = append(env, "GOARCH=mipsle", "GOMIPS="+opts.Platform.GOMIPS())
	case "mips":
		env = append(env, "GOARCH=mips", "GOMIPS="+opts.Platform.GOMIPS())
	case "ppc64":
		env = append(env, "GOARCH=ppc64", "GOPPC64="+opts.Platform.GOPPC64())
	case "wasm":
		env = append(env, "GOARCH=wasm", "GOWASM="+opts.Platform.GOWASM())
	default:
		panic(fmt.Errorf("invalid GOARCH value: %q", goarch))
	}

	return env
}

func (opts *Options) AppendASMFlags(flags ...string) []string {
	switch os := opts.Platform.OS(); os {
	case "efi":
		flags = append(flags, "-D", "GOOS_efi")
	case "noos":
	default:
		flags = append(flags, "-D", "GOOS_"+opts.GOOS())
	}

	switch goarch := opts.GOARCH(); goarch {
	case "386":
		flags = append(flags, "-D", "GOARCH_386", "-D", "GO386_"+opts.Platform.GO386())
	case "amd64":
		x := opts.Platform.GOAMD64()
		flags = append(flags, "-D", "GOARCH_amd64", "-D", "GOAMD64_"+x)

		// ref: ${GOROOT}/src/runtime/asm_amd64.h
		switch x {
		case "v4":
			fallthrough
		case "v3":
			flags = append(flags, "-D", "hasAVX", "-D", "hasAVX2")
			fallthrough
		case "v2":
			flags = append(flags, "-D", "hasPOPCNT", "-D", "hasSSE42")
		}

	case "arm":
		flags = append(flags, "-D", "GOARCH_arm", "-D", "GOARM_"+opts.Platform.GOARM())
	case "arm64":
		flags = append(flags, "-D", "GOARCH_arm64")
	case "mips64le":
		flags = append(flags, "-D", "GOARCH_mips64le", "-D", "GOMIPS64_"+opts.Platform.GOMIPS64())
	case "mips64":
		flags = append(flags, "-D", "GOARCH_mips64", "-D", "GOMIPS64_"+opts.Platform.GOMIPS64())
	case "mipsle":
		flags = append(flags, "-D", "GOARCH_mipsle", "-D", "GOMIPS_"+opts.Platform.GOMIPS())
	case "mips":
		flags = append(flags, "-D", "GOARCH_mips", "-D", "GOMIPS_"+opts.Platform.GOMIPS())
	case "ppc64":
		flags = append(flags, "-D", "GOARCH_ppc64", "-D", "GOPPC64_"+opts.Platform.GOPPC64())

		// ref: ${GOROOT}/src/runtime/asm_ppc64x.h
		flags = append(flags, "-D", "FIXED_FRAME=32")
	case "wasm":
		flags = append(flags, "-D", "GOARCH_wasm")
		for _, opt := range strings.Split(opts.Platform.GOWASM(), ",") {
			flags = append(flags, "-D", "GOWASM_"+strings.TrimSpace(opt))
		}
	default:
		panic(fmt.Errorf("invalid GOARCH value: %q", goarch))
	}

	return append(flags, opts.ASMFlags...)
}

func (opts *Options) AppendGCFlags(flags []string) []string {
	if opts.GOARCH() == "wasm" || opts.GOOS() == "plan9" {
		flags = append(flags,
			"-dwarf=false", // wasm has no support for dwarf symbols,
		)
	} else {
		flags = append(flags,
			"-shared", // build shared object
			// "-dwarf",  // generate dwarf symbols (default behavior)
		)
	}

	// TODO: add command-line presets to add these flags
	// if opts.Debug {
	// 	flags = append(flags,
	// 		"-%",    // debug non-static initializers
	// 		"-v",    // increase debug verbosity
	// 		"-E",    // debug symbol export
	// 		"-live", // debug liveness analysis
	// 		"-w",    // debug type checking
	// 		"-r",    // debug generated wrappers
	// 		"-j",    // debug runtime initialized variables
	// 		"-m=10", // show optimization decisions
	// 	)
	// }

	if opts.EnableASan {
		flags = append(flags, "-asan")
	}

	if opts.EnableMSan {
		flags = append(flags, "-msan")
	}

	if opts.EnableRace {
		flags = append(flags, "-race")
	}

	return append(flags, opts.GCFlags...)
}

func (opts *Options) AppendLDFlags(flags ...string) []string {
	if len(opts.EntrySymbol) != 0 {
		if opts.GOOS() == "wasip1" {
			// (go1.21) go linker always expects "_rt0_wasm_wasip1"
			flags = append(flags, "-E", "_rt0_wasm_wasip1")
		} else {
			flags = append(flags, "-E", opts.EntrySymbol)
		}
	}

	if opts.EnableASan {
		flags = append(flags, "-asan")
	}

	if opts.EnableMSan {
		flags = append(flags, "-msan")
	}

	if opts.EnableRace {
		flags = append(flags, "-race")
	}

	for _, def := range opts.Defs {
		flags = append(flags, "-X", def)
	}

	if len(opts.Mode) != 0 {
		flags = append(flags, "-buildmode", opts.Mode)
	}

	return append(flags, opts.LDFlags...)
}
