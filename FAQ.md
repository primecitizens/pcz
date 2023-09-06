# Frequently Asked Questions (maybe)

## What's the next supported toolchain?

When all of following criteria are met, then we would work on support for a new toolchain:

- Some siginificant features landed in compiler (`gc`) or linker (e.g. [golang/go#61502](https://github.com/golang/go/issues/61502))
- Still possible to avoid escape analysis:
  - either `mark.NoEscape` still works (not becoming a compiler intrinsic)
  - or there is option in the new toolchain to turn off escape analysis.
- Still possible to rename packages to `internal/*`
  - requires toolchain support.

## Why this project?

To make software controlled by peasant -- not to create restrictions to software development, but to enable every peasant to build complex software with simple instructions.

See also [goals of the project](./README.md#goals).

## What does `pcz` stand for?

The acronym for Prime Citizens -- a reminder to all of us everyone is their own Prime Minister, think in individuals, not in group of people.

or Peasant Control Zutto (meaning `forever` in Japanese).
