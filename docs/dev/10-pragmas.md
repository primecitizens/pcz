# Go pragmas

## References

- `${GOROOT}/src/cmd/compile/doc.go`
- `${GOROOT}/src/cmd/cgo/doc.go`
- `${GOROOT}/src/cmd/compile/internal/ir/node.go#type:PragmaFlag`
- `${GOROOT}/src/cmd/compile/internal/noder/lex.go#func:pragmaFlag`

## Precompile-time Pragmas

- `go:build`

## Compile-time Pragmas

- `go:nointerface`
  - used by `fieldtrack` experiment, see [golang/go#47045](https://github.com/golang/go/issues/47045)

- `go:norace`

  > The //go:norace directive must be followed by a function declaration.
  > It specifies that the function's memory accesses must be ignored by the
  > race detector. This is most commonly used in low-level code invoked
  > at times when it is unsafe to call into the race detector runtime.

- `go:nosplit`

  > The //go:nosplit directive must be followed by a function declaration.
  > It specifies that the function must omit its usual stack overflow check.
  > This is most commonly used by low-level runtime code invoked
  > at times when it is unsafe for the calling goroutine to be preempted.

- `go:noinline`
  > The //go:noinline directive must be followed by a function declaration.
  > It specifies that calls to the function should not be inlined, overriding
  > the compiler's usual optimization rules. This is typically only needed
  > for special runtime functions or when debugging the compiler.

- `go:nocheckptr`

  > Func should not be instrumented by checkptr

- `go:systemstack`

  > Func must run on system stack

- `go:nowritebarrier`
  
  > Emit compiler error instead of write barrier

- `go:nowritebarrierrec`

  > Error on write barrier in this or recursive callees

- `go:yeswritebarrierrec`

  > Cancels Nowritebarrierrec in this function and callees

- `go:uintptrkeepalive`

  > Pointers converted to uintptr must be kept alive

- `go:cgo_unsafe_args`

  > Treat a pointer to one arg as a pointer to them all
  >
  > If marked "go:cgo_unsafe_args", don't inline, since the function
  > makes assumptions about its argument frame layout.

- `go:registerparams` (to be removed after register abi is working)

- `go:noescape`

  > The //go:noescape directive must be followed by a function declaration without
  > a body (meaning that the function has an implementation not written in Go).
  > It specifies that the function does not allow any of the pointers passed as
  > arguments to escape into the heap or into the values returned from the function.
  > This information can be used during the compiler's escape analysis of Go code
  > calling the function.

- `go:uintptrescapes`

  > The //go:uintptrescapes directive must be followed by a function declaration.
  > It specifies that the function's uintptr arguments may be pointer values that
  > have been converted to uintptr and must be on the heap and kept alive for the
  > duration of the call, even though from the types alone it would appear that the
  > object is no longer needed during the call. The conversion from pointer to
  > uintptr must appear in the argument list of any call to this function. This
  > directive is necessary for some low-level system call implementations and
  > should be avoided otherwise.

## Link-time Pragmas

- `go:linkname localname [importpath.name]`

  > The //go:linkname directive conventionally precedes the var or func
  > declaration named by `localname`, though its position does not
  > change its effect.
  > This directive determines the object-file symbol used for a Go var or
  > func declaration, allowing two Go symbols to alias the same
  > object-file symbol, thereby enabling one package to access a symbol in
  > another package even when this would violate the usual encapsulation
  > of unexported declarations, or even type safety.
  > For that reason, it is only enabled in files that have imported "unsafe".
  >
  > It may be used in two scenarios. Let's assume that package upper
  > imports package lower, perhaps indirectly. In the first scenario,
  > package lower defines a symbol whose object file name belongs to
  > package upper. Both packages contain a linkname directive: package
  > lower uses the two-argument form and package upper uses the
  > one-argument form. In the example below, lower.f is an alias for the
  > function upper.g:
  >
  >     package upper
  >     import _ "unsafe"
  >     //go:linkname g
  >     func g()
  > 
  >     package lower
  >     import _ "unsafe"
  >     //go:linkname f upper.g
  >     func f() { ... }
  >
  > The linkname directive in package upper suppresses the usual error for
  > a function that lacks a body. (That check may alternatively be
  > suppressed by including a .s file, even an empty one, in the package.)
  >
  > In the second scenario, package upper unilaterally creates an alias
  > for a symbol in package lower. In the example below, upper.g is an alias
  > for the function lower.f.
  >
  >     package upper
  >     import _ "unsafe"
  >     //go:linkname g lower.f
  >     func g()
  > 
  >     package lower
  >     func f() { ... }
  >
  > The declaration of lower.f may also have a linkname directive with a
  > single argument, f. This is optional, but helps alert the reader that
  > the function is accessed from outside the package.

- `go:cgo_import_dynamic <local> [<remote> ["<library>"]]`

  > In internal linking mode, allow an unresolved reference to
  > `<local>`, assuming it will be resolved by a dynamic library
  > symbol. The optional `<remote>` specifies the symbol's name and
  > possibly version in the dynamic library, and the optional `"<library>"`
  > names the specific library where the symbol should be found.
  >
  > On AIX, the library pattern is slightly different. It must be
  > "lib.a/obj.o" with obj.o the member of this library exporting
  > this symbol.
  >
  > In the `<remote>`, # or @ can be used to introduce a symbol version.
  >
  > Examples:
  > //go:cgo_import_dynamic puts
  > //go:cgo_import_dynamic puts puts#GLIBC_2.2.5
  > //go:cgo_import_dynamic puts puts#GLIBC_2.2.5 "libc.so.6"
  >
  > A side effect of the cgo_import_dynamic directive with a
  > library is to make the final binary depend on that dynamic
  > library. To get the dependency without importing any specific
  > symbols, use _ for local and remote.
  >
  > Example:
  > //go:cgo_import_dynamic _ _ "libc.so.6"
  >
  > For compatibility with current versions of SWIG,
  > #pragma dynimport is an alias for //go:cgo_import_dynamic.

- `go:cgo_export_dynamic <local> <remote>`

  > In internal linking mode, put the Go symbol
  > named `<local>` into the program's exported symbol table as
  > `<remote>`, so that C code can refer to it by that name. This
  > mechanism makes it possible for C code to call back into Go or
  > to share Go's data.
  >
  > For compatibility with current versions of SWIG,
  > #pragma dynexport is an alias for //go:cgo_export_dynamic.

- `go:cgo_import_static <local>`

  > In external linking mode, allow unresolved references to
  > `<local>` in the go.o object file prepared for the host linker,
  > under the assumption that `<local>` will be supplied by the
  > other object files that will be linked with go.o.
  >
  > Example:
  > //go:cgo_import_static puts_wrapper

- `go:cgo_export_static <local> <remote>`

  > In external linking mode, put the Go symbol
  > named `<local>` into the program's exported symbol table as
  > `<remote>`, so that C code can refer to it by that name. This
  > mechanism makes it possible for C code to call back into Go or
  > to share Go's data.

- `go:cgo_dynamic_linker "<path>"`

  > In internal linking mode, use `"<path>"` as the dynamic linker
  > in the final binary. This directive is only needed from one
  > package when constructing a binary; by convention it is
  > supplied by runtime/cgo.
  >
  > Example:
  > //go:cgo_dynamic_linker "/lib/ld-linux.so.2"

- `go:cgo_ldflag "<arg>"`

  > In external linking mode, invoke the host linker (usually gcc)
  > with `"<arg>"` as a command-line argument following the .o files.
  > Note that the arguments are for "gcc", not "ld".
  >
  > Example:
  > //go:cgo_ldflag "-lpthread"
  > //go:cgo_ldflag "-L/usr/local/sqlite3/lib"
