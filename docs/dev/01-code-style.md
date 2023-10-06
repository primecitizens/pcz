# Code Style

## Always Do

- import `_ "embed"` when using `//go:embed`
- import `_ "unsafe"` when using `//go:linkname`
- add `//go:build GOARCH` when the file is platform specific, DO NOT rely on filenames.
- add `Ex` version of the function when the function takes a callback.
  - for example: `func Foo(a []T, fn func(T))`, add `func FooEx(a []T, arg Arg, fn func(Arg, T))`
  - NOTE: the additional `Arg` MUST be the first arg of the callback so it works with `Type.Method` and `(*Type).Method`.
