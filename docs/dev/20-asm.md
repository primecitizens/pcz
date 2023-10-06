# Go Assembly

- Syntax
  - Introduction: [https://go.dev/doc/asm](https://go.dev/doc/asm)
  - Manual: [https://9p.io/sys/doc/asm.html](https://9p.io/sys/doc/asm.html) ([https://doc.cat-v.org/plan_9/2nd_edition/papers/asm/](https://doc.cat-v.org/plan_9/2nd_edition/papers/asm/))

- Instruction Names
  - [${GOROOT}/cmd/internal/obj/arm64](https://pkg.go.dev/cmd/internal/obj/arm64)
  - [${GOROOT}/cmd/internal/obj/x86](https://pkg.go.dev/cmd/internal/obj/x86)
  - [${GOROOT}/cmd/internal/obj/riscv](https://pkg.go.dev/cmd/internal/obj/riscv)
  - [${GOROOT}/cmd/internal/obj/wasm](https://pkg.go.dev/cmd/internal/obj/wasm)
  - [${GOROOT}/cmd/internal/obj/ppc64](https://pkg.go.dev/cmd/internal/obj/ppc64)
  - [${GOROOT}/cmd/internal/obj/arm](https://pkg.go.dev/cmd/internal/obj/arm)
  - [${GOROOT}/cmd/internal/obj/mips](https://pkg.go.dev/cmd/internal/obj/mips)

## The `g` (goroutine) register

- 386 (TODO)
  - android
  - linux
  - windows
- amd64
  - ABIInternal: REGG = REG_R14 (g register in ABIInternal)
  - ABI0 (TODO)
    - windows
    - darwin
    - linux
      - `-shared`
    - freebsd
      - `-shared`
- arm: REGG = REGEXT - 0 = REG_R10
- arm64: REGG = REG_R28
- mips: REGG = REG_R30
- ppc64: REGG = REG_R30
- riscv: REGG = REG_G = REG_S11
- s390x: REGG = REG_R13
- wasm: REGG = REG_g (global variable `$global2`)

Refs:

- `${GOROOT}/src/cmd/internal/obj/*/a.out.go`

## Instruction Suffix

- `386`

    | Suffix        | Meaning         |
    | ------------- | --------------- |
    | L (long word) | 32 bit          |
    | LU            | 32 bit unsigned |
    | W (word)      | 16 bit          |
    | WU            | 16 bit unsigned |
    | B (byte)      | 8  bit          |
    | BU            | 8  bit unsigned |
    | S (single)    | 32 bit float    |
    | D (double)    | 64 bit float    |

- `amd64`

    | Suffix        | Meaning         |
    | ------------- | --------------- |
    | Q (quad word) | 64 bit          |
    | L (long word) | 32 bit          |
    | W (word)      | 16 bit          |
    | WU            | 16 bit unsigned |
    | B (byte)      | 8  bit          |
    | S (single)    | 32 bit float    |
    | D (double)    | 64 bit float    |

- `arm`

    | Suffix        | Meaning         |
    | ------------- | --------------- |
    | W (word)      | 32 bit          |
    | H (half word) | 16 bit          |
    | HU            | 16 bit unsigned |
    | B (byte)      | 8 bit           |
    | BU            | 8 bit unsigned  |
    | F (float)     | 32 bit float    |
    | D (double)    | 64 bit float    |

- `arm64`

    | Suffix          | Meaning         |
    | --------------- | --------------- |
    | D (double word) | 64 bit          |
    | W (word)        | 32 bit          |
    | H (half word)   | 16 bit          |
    | HU              | 16 bit unsigned |
    | B (byte)        | 8  bit          |
    | BU              | 8  bit unsigned |
    | S (single)      | 32 bit float    |
    | D (double)      | 64 bit float    |

- `riscv64`

    | Suffix          | Meaning                                        |
    | --------------- | ---------------------------------------------- |
    | L               | 64 bit int, used when the opcode starts with F |
    | D (double word) | 64 bit int                                     |
    | W (word)        | 32 bit int                                     |
    | H (half word)   | 16 bit int                                     |
    | B (byte)        | 8  bit int                                     |
    | S (single)      | 32 bit float                                   |
    | D (double)      | 64 bit float                                   |

- `mips`

    | Suffix        | Meaning         |
    | ------------- | --------------- |
    | W (word)      | 32 bit          |
    | H (half word) | 16 bit          |
    | HU            | 16 bit unsigned |
    | B (byte)      | 8 bit           |
    | BU            | 8 bit unsigned  |
    | F (float)     | 32 bit float    |
    | D (double)    | 64 bit float    |

- `mips64`

    | Suffix        | Meaning         |
    | ------------- | --------------- |
    | V (vlong)     | 64 bit          |
    | W (word)      | 32 bit          |
    | WU (word)     | 32 bit unsigned |
    | H (half word) | 16 bit          |
    | HU            | 16 bit unsigned |
    | B (byte)      | 8 bit           |
    | BU            | 8 bit unsigned  |
    | F (float)     | 32 bit float    |
    | D (double)    | 64 bit float    |

- `s390x`

    | Suffix          | Meaning                              |
    | --------------- | ------------------------------------ |
    | D (double word) | 64 bit (frequently omitted)          |
    | W (word)        | 32 bit                               |
    | H (half word)   | 16 bit                               |
    | B (byte)        | 8  bit                               |
    | S (single)      | 32 bit (double precision is omitted) |

- `wasm`

    | Suffix          | Meaning |
    | --------------- | ------- |
    | D (double word) | 64 bit  |
    | W (word)        | 32 bit  |
    | H (half word)   | 16 bit  |
    | B (byte)        | 8  bit  |

Refs:

- `${GOROOT}/src/cmd/compile/internal/ssa/_gen/*.rules`
- `${GOROOT}/src/cmd/compile/internal/ssa/_gen/*Ops.go`

## FPU Mode in Go

- `386`: x87 FPCSR

    | Flag | Value | Meaning                   |
    | ---- | ----- | ------------------------- |
    | RC   | 0     | Round to nearest          |
    | PC   | 11    | Double extended precision |
    | PM   | 1     | Precision masked          |
    | UM   | 1     | Underflow masked          |
    | OM   | 1     | Overflow masked           |
    | ZM   | 1     | Zero divide masked        |
    | DM   | 1     | De-normal operand masked  |
    | IM   | 1     | Invalid operation masked  |

- `amd64`: SSE2 MXCSR (x87 FPCSR is not used by Go on amd64), matches MS and SysV ABI.

    | Flag | Bit   | Value  | Meaning                    |
    | ---- | ----- | ------ | -------------------------- |
    | FZ   | 15    | 0      | Do not flush to zero       |
    | RC   | 14/13 | 0 (RN) | Round to nearest           |
    | PM   | 12    | 1      | Precision masked           |
    | UM   | 11    | 1      | Underflow masked           |
    | OM   | 10    | 1      | Overflow masked            |
    | ZM   | 9     | 1      | Divide-by-zero masked      |
    | DM   | 8     | 1      | Denormal operations masked |
    | IM   | 7     | 1      | Invalid operations masked  |
    | DAZ  | 6     | 0      | Do not zero de-normals     |

- `arm64`: FPCR

    | Flag | Bit   | Value  | Meaning                                                             |
    | ---- | ----- | ------ | ------------------------------------------------------------------- |
    | DN   | 25    | 0      | Propagate NaN operands                                              |
    | FZ   | 24    | 0      | Do not flush to zero                                                |
    | RC   | 23/22 | 0 (RN) | Round to nearest, choose even if tied                               |
    | IDE  | 15    | 0      | Denormal operations trap disabled                                   |
    | IXE  | 12    | 0      | Inexact trap disabled                                               |
    | UFE  | 11    | 0      | Underflow trap disabled                                             |
    | OFE  | 10    | 0      | Overflow trap disabled                                              |
    | DZE  | 9     | 0      | Divide-by-zero trap disabled                                        |
    | IOE  | 8     | 0      | Invalid operations trap disabled                                    |
    | NEP  | 2     | 0      | Scalar operations do not affect higher elements in vector registers |
    | AH   | 1     | 0      | No alternate handling of de-normal inputs                           |
    | FIZ  | 0     | 0      | Do not zero de-normals                                              |

- `ppc64`: FPSCR initialized to 0 by the kernel and untouched in Go
- `riscv64`: CSR initialized by system and untouched in Go

Refs:

- [`${GOROOT}/src/cmd/compile/abi-internal.md`](go.dev/s/regabi)
