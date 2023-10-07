# Roadmap

The ultimate goal is to make Go, the branded modern C, become the actual modern C, featuring great usability and platform native support.

We are well aware it is an ambitious project regarding to our small scale, and it takes time and effort in order to make it actually happen, currently we are working on following tasks:

## Long-term

- Encourage Go developers to build more allocation-free modules.

## Phase 1: Make Go Suitable for System Programming (ongoing)

### Stage 1: Up & Running - Single-Thread Application (current)

> **NOTE**
> At this stage, we will prioritize the `js/wasm` support.

- [x] Be compatible with go1.21 toolchain.
- [x] Better structuring of the std.
- [x] Expose all low-level & reusable bits of the runtime.
- [x] Provide apis for custom memory allocator.
- [ ] Stacktrace.
- [ ] Stack growth.
  - [ ] Copy stack.
- [ ] Debugging with `delve`.
- [ ] Program startup
  - [x] `js/wasm`
  - [x] `wasip1/wasm`
  - [ ] `linux/{arm64, amd64}`
  - [ ] `android/{arm64, amd64}`
  - [ ] `darwin/{arm64, amd64}`
  - [ ] `ios/arm64`
  - [ ] `windows/{arm64, amd64}`
- [ ] FFI support (call and callback)
  - [x] `js`
  - [x] `wasm/wasi`
  - [ ] `syscall/linux`
  - [ ] `vdso`
  - [ ] `C`
    - [ ] `aapcs64`
    - [ ] `sysv{32, 64}`
    - [ ] `fastcall`
    - [ ] `stdcall`
    - [ ] `win64`
  - [ ] `objc` (based on `C`)
  - [ ] `jni` (based on `C`)
- [ ] Platform SDK (`pcz codegen`)
  - [ ] `DOM` apis (code generation with WebIDL)
    - [x] Support converting WebIDL to YAML spec in `h2y`.
    - [x] Support Go package generation for WebIDL in `ffigen`.
    - [ ] Use browser-compat-data for comments.
  - [x] `WebExtension` apis (Manifest Version 3).
    - [x] Normalize chromium web extension idls & json.
    - [x] Resolve package conflicts among namespaces.
  - [ ] `nodejs` apis.
  - [ ] `android` NDK
    - [x] Add C header file support in `h2y`
    - [ ] Support Go package generation for C in `ffigen`.
  - [ ] `darwin` XCode (`*.sdk/*.framework`)
    - [x] Add Objective-C header file support in `h2y`
    - [ ] Support Go package generation for Objective-C in `ffigen`.
  - [ ] `win32` Windows SDK (see tasks for `android` NDK)

> **NOTE**
> A lot of design decisions to be made in the stage 1 tasks require taking stage 2 tasks into consideration, we won't create any barrier for any disscussions under any circumstances, stay ontopic and you are always welcome.

### Stage 2: The `go` keyword reMastered - Multi-Thread Application

- [ ] Define apis for building custom `goroutine` support.
- [ ] Port Sched program from the official Go std for `goroutine`.
  - [ ] Generalize and expose `internal/poll`.
- [ ] Port GC program from the official Go std to work with custom allocators.
- [ ] Runtime feature support.
  - [ ] `race`
  - [ ] `asan`
  - [ ] `msan`
  - [ ] `pprof`
  - [ ] `fuzz`

### Stage 3: Dream Come True? - Cross-platform Native GUI

- [ ] GPU support
  - [ ] `vulkan`
  - [ ] `directx`
  - [ ] `metal`
- [ ] Wrap native GUI widgets exposed by platform apis.
- [ ] Constraint layout for native GUI widgets.
- [ ] `opentype` font support.

TBD.

## Phase 2: Compile C in Go

- Support static binary build when working with `C` (`go build -buildmode=static`)
  - Implement a `pcz` subcommand to support go external linking mode.
- Standardize linker scripts in YAML (in our opinion, they resembles a lot)
  - Implement tool (`l2y`) to convert existing linker scripts to YAML.
  - Implement support for linking custom sections in `pcz`.
- Support `unsafe.SelectABI` for foreign ABIs in go compiler (inspired by [golang/go#44065](https://github.com/golang/go/issues/44065)).
- A libc implemented in Go.
- A configurable run-time linker (`ldso`).
  - Define platform agnostic standard for thread-local storage relocation.

## Phase 3: Build kernels in Go

- Add missing privileged instructions to `$GOROOT/src/cmd/asm`.
- UEFI libraries in Go
  - OS Loader
- Drivers
- Unikernel library (Inspired by [eliasnaur/unik](https://git.sr.ht/~eliasnaur/unik)).

## Phase 4: Make Go Toolchain the Modern LLVM?

TBD.
