# Go WebAssembly

## References

- Wasm implementation notes in `$GOROOT/src/cmd/compile/internal/wasm/ssa.go`
- Comments inside `$GOROOT/src/cmd/link/internal/wasm/asm.go#func:assignAddress`
- [https://webassembly.github.io/spec/core/syntax/instructions.html](https://webassembly.github.io/spec/core/syntax/instructions.html)

## Tools

- [WebAssembly/wabt](https://github.com/WebAssembly/wabt)
  - `wasm-objdump --disassemble --reloc --detailx FILE` (or `wasm-objdump -drx FILE`)
  - wasm-decompile

## Global Variables

- `$global0` = `SP` (stack pointer register in Go)
- `$global1` = `CTXT` (closure pointer register in Go)
- `$global2` = `g` (goroutine pointer register in Go)
- `$global3` (reflectcall RET0)
- `$global4` (reflectcall RET1)
- `$global5` (reflectcall RET2)
- `$global6` (reflectcall RET3)
- `$global7` = `PAUSE`

## Local Variables

- Go functions
  - `$var0` = `PC_B`
  - `$var1` = local copy of `SP`

## Terminology

- `PC_F` (uint16): index of the function in the compiled wasm module
- `PC_B` (uint16): the resume point inside a go function (a branch label index to `br_table`)
- `PC` (uint32): the Go implementation of traditional program counter (`PC` = `PC_F`<<16 + `PC_B`)
