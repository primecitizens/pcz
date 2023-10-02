// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package toolchain

import (
	"fmt"
	"runtime"
	"strings"
)

// Platform is a os/arch pair
type Platform string

func (p Platform) Pair() (os, arch string) {
	os, arch, ok := strings.Cut(string(p), "/")
	if ok {
		return
	}

	return "noos", string(p)
}

func (p Platform) OS() string {
	ret, _ := p.Pair()
	return ret
}

func (p Platform) Arch() string {
	_, ret := p.Pair()
	return ret
}

func (p Platform) GOOS(envGOOS string) string {
	switch os := p.OS(); os {
	case "":
		if len(envGOOS) != 0 { // from env
			return envGOOS
		}

		return EnvOr("GOHOSTOS", runtime.GOOS) // fallback to host GOOS
	case "noos", "efi":
		switch p.Arch() {
		case "wasm":
			return "js"
		default:
			// use linux for noos target as it supports all available cpu archs
			// with most compiler features available
			return "linux"
		}
	case "js", "nodejs", "web", "webext":
		return "js"

	case "wasip1", "wasi", "wasix":
		return "wasip1"

	case "darwin", "ios",
		"linux", "android",
		"freebsd", "openbsd", "netbsd", "solaris", "aix", "dragonfly", "illumos",
		"plan9",
		"windows":
		return os
	case "watchos", "tvos":
		return "darwin"
	default:
		panic(fmt.Errorf("invalid platform option: %q", os))
	}
}

func (p Platform) GOARCH(envGOARCH string) string {
	switch arch := p.Arch(); arch {
	case "":
		if len(envGOARCH) != 0 {
			return envGOARCH
		}

		return EnvOr("GOHOSTARCH", runtime.GOARCH) // fallback to host GOARCH
	case "x86", "x86sse2", "x86sf":
		return "386"

	case "arm", "armv5", "armv6", "armv7":
		return "arm"

	case "arm64":
		return "arm64"

	case "amd64", "amd64v1", "amd64v2", "amd64v3", "amd64v4":
		return "amd64"

	case "mips", "mipssf", "mipshf":
		return "mips"
	case "mipsle", "mipslesf", "mipslehf":
		return "mipsle"
	case "mips64", "mips64sf", "mips64hf":
		return "mips64"
	case "mips64le", "mips64lesf", "mips64lehf":
		return "mips64le"

	case "ppc64", "ppc64v8", "ppc64v9", "ppc64v10":
		return "ppc64"
	case "ppc64le", "ppc64lev8", "ppc64lev9", "ppc64lev10":
		return "ppc64le"

	case "wasm":
		return "wasm"
	default:
		if strings.HasPrefix(arch, "wasm:") {
			return "wasm"
		}

		panic(fmt.Errorf("unknown arch value: %q", arch))
	}
}

func (p Platform) GOARM() string {
	switch arch := p.Arch(); arch {
	case "":
		return GOARM
	case "arm":
		return "7"
	case "armv5", "armv6", "armv7":
		return strings.TrimPrefix(arch, "armv")
	default:
		panic("bad usage")
	}
}

func (p Platform) GO386() string {
	switch p.Arch() {
	case "":
		return GO386
	case "x86", "x86sse2":
		return "sse2"
	case "x86sf":
		return "softfloat"
	default:
		panic("bad usage")
	}
}

func (p Platform) GOAMD64() string {
	switch arch := p.Arch(); arch {
	case "":
		return GOAMD64
	case "amd64", "amd64v1", "amd64v2", "amd64v3", "amd64v4":
		suffix := strings.TrimPrefix(arch, "amd64")
		if len(suffix) == 0 {
			suffix = "v1"
		}

		return suffix
	default:
		panic("bad usage")
	}
}

func (p Platform) GOMIPS64() string {
	switch p.Arch() {
	case "":
		return GOMIPS64
	case "mips64", "mips64hf", "mips64le", "mips64lehf":
		return "hardfloat"
	case "mips64sf", "mips64lesf":
		return "softfloat"
	default:
		panic("bad usage")
	}
}

func (p Platform) GOMIPS() string {
	switch p.Arch() {
	case "":
		return GOMIPS
	case "mips", "mipshf", "mipsle", "mipslehf":
		return "hardfloat"
	case "mipssf", "mipslesf":
		return "softfloat"
	default:
		panic("bad usage")
	}
}

func (p Platform) GOPPC64() string {
	switch p.Arch() {
	case "":
		return GOPPC64
	case "ppc64", "ppc64v8", "ppc64le", "ppc64lev8":
		return "power8"
	case "ppc64v9", "ppc64lev9":
		return "power9"
	case "ppc64v10", "ppc64lev10":
		return "power10"
	default:
		panic("bad usage")
	}
}

func (p Platform) GOWASM() string {
	switch arch := p.Arch(); arch {
	case "":
		return GOWASM
	case "wasm":
		return ""
	default:
		opts, ok := strings.CutPrefix(arch, "wasm:")
		if ok {
			return opts
		}

		panic("bad usage")
	}
}
