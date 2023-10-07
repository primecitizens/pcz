# WebExtension Platform Development

## Supported Targets

- `js/wasm`

## Development Tools

- `pcz dev webext`: build wasm blob and generate bindings on source change.
  - package: `github.com/primecitizens/pcz/cmd/dev/webext`
- `pcz dev webext bindgen`: generate js bindings for wasm.
  - package: `github.com/primecitizens/pcz/cmd/dev/internal/js`

## Platform Native APIs

- `github.com/primecitizens/pcz/std/plat/js/webext/**`: A collection of packages exposing all apis generated from [chromium source idls](../../cmd/codegen/js/webext/)
  - It is split into multiple packages because:
    - (major) There are a lot of type name collisions.
    - (major) WebExtension apis are defined in their own namepsaces.
  - There is no plan to collect these APIs into a single package.
  - NOTE: We are in the process planning to make it a standalon module in a separate repository.

> **NOTE**
> The `tabs` namepsace is spread in 2 packages (`tabs` and `tabs1`) in order to resolve import cycle with the `runtime` and `windows` namespace.
>
> The `os` namespace is hosted by package `chromeos`.

## Browser Requirements

- Manifest Version 3 Support.
- Expose the WebExtension apis via `browser` (preferred) or `chrome` (fallback)
