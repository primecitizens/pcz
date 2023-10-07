# Web Platform Development

## Supported Targets

- `js/wasm`

## Development Tools

- `pcz dev web`
  - package: `github.com/primecitizens/pcz/cmd/dev/web`
- `pcz dev web bindgen`: generate js bindings for wasm.
  - package: `github.com/primecitizens/pcz/cmd/dev/internal/js`

## Platform Native APIs

- `github.com/primecitizens/pcz/std/plat/js/web`: A single package exposing all apis generated from [w3c/webref@curated ed/idl](https://github.com/w3c/webref)
  - It is generated as a single mega package because:
    - (major) We haven't found a reasonable way to categorize apis into their own packages (suggestions are welcome).
    - (minor) The web development is done with all things global, and most apis are accessed from the `Window` and `Navigator` type.
  - We are in the process planning to make it a standalon module in a separate repository.
