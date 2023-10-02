// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && wasip1

package sys

import (
	"unsafe"

	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/ffi/wasm/wasi"
)

var (
	args []string
	envs []string
)

func ntharg(i int) string {
	return args[i]
}

func nthenv(i int) string {
	return envs[i]
}

func init() {
	argc, argSz, errno := wasi.ArgsSizes()
	if errno != 0 {
		assert.Throw("args_sizes_get", "failed")
	}

	envc, envSz, errno := wasi.EnvironSizes()
	if errno != 0 {
		assert.Throw("environ_sizes_get", "failed")
	}

	// defensive: ensure size is zero when count is zero.
	if argc == 0 {
		argSz = 0
	}
	if envc == 0 {
		envSz = 0
	}

	if argSz == 0 && envSz == 0 {
		return
	}

	type uintptr32 = uint32

	totalSize := uintptr(envSz) + uintptr(argSz) +
		unsafe.Sizeof(string(""))*(uintptr(envc)+uintptr(argc))

	buf := make([]byte, totalSize)
	args = unsafe.Slice(
		(*string)(unsafe.Pointer(unsafe.SliceData(buf))),
		argc,
	)

	buf = buf[uintptr(argc)*unsafe.Sizeof(string("")):]
	envs = unsafe.Slice(
		(*string)(unsafe.Pointer(unsafe.SliceData(buf))),
		envc,
	)
	buf = buf[uintptr(envc)*unsafe.Sizeof(string("")):]

	if argSz != 0 {
		argv := unsafe.Slice((*uintptr32)(unsafe.Pointer(unsafe.SliceData(args))), argc)
		if wasi.ArgsGet(
			unsafe.Pointer(unsafe.SliceData(argv)),
			unsafe.Pointer(unsafe.SliceData(buf)),
		) != 0 {
			assert.Throw("args_get", "failed")
		}

		for argc > 0 {
			argc--
			start := argv[argc] - uintptr32(uintptr(unsafe.Pointer(unsafe.SliceData(buf))))
			end := start
			for buf[end] != 0 {
				end++
			}
			args[argc] = unsafe.String(&buf[start], end-start)
		}

		buf = buf[argSz:]
	}

	if envSz != 0 {
		envv := unsafe.Slice((*uintptr32)(unsafe.Pointer(unsafe.SliceData(envs))), envc)
		if wasi.Environ(
			unsafe.Pointer(unsafe.SliceData(envv)),
			unsafe.Pointer(unsafe.SliceData(buf)),
		) != 0 {
			assert.Throw("environ_get", "failed")
		}

		for envc > 0 {
			envc--
			start := envv[envc] - uintptr32(uintptr(unsafe.Pointer(unsafe.SliceData(buf))))
			end := start
			for buf[end] != 0 {
				end++
			}
			envs[envc] = unsafe.String(&buf[start], end-start)
		}
	}
}
