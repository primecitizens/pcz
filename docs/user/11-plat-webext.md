# WebExtension Platform Development

Exemple: [examples/11-plat-webext](../../examples/11-plat-webext)

## Supported Targets

- `js/wasm`

## Development Tools

- `pcz dev webext`: build wasm blob and generate bindings on source change.
  - package: `github.com/primecitizens/pcz/cmd/dev/webext`
- `pcz dev webext new`: create a new project for building WebExtension.
  - package: `github.com/primecitizens/pcz/cmd/dev/webext`
- `pcz dev webext pack`: create zip files with matched files
  - package: `github.com/primecitizens/pcz/cmd/dev/webext`
- `pcz dev webext bindgen`: generate js bindings for wasm.
  - package: `github.com/primecitizens/pcz/cmd/dev/internal/js`
- `pcz embed js`: generate js embedings of wasm blob for loading wasm app synchronously.
  - package: `github.com/primecitizens/pcz/cmd/embed`

## Platform Native APIs

- `github.com/primecitizens/pcz/std/plat/js/webext/**`: A collection of packages exposing all apis generated from [chromium source idls](../../cmd/codegen/js/webext)
  - It is split into multiple packages because:
    - (major) There are a lot of type name collisions.
    - (major) WebExtension apis are defined in their own namepsaces.
  - There is no plan to collect these APIs into a single package.
  - We are in the process planning to make it a standalone module in a separate repository.

  > **NOTE**
  > The `tabs` namepsace is spread in 2 packages (`tabs` and `tabs1`) in order to resolve import cycle with the `runtime` and `windows` namespace.
  >
  > The `os` namespace is hosted by the package `chromeos`.

- `https://json.schemastore.org/chrome-manifest.json`: JSON schema for `manifest.json` (manifest version 3).

## Browser Requirements

- Expose the WebExtension apis via `browser` (preferred) or `chrome` (fallback)
- Asynchronous WebExtension API support (returns `Promise` instead of requiring a callback).
  - For `Chromium`-like browsers: Manifest Version 3+.
  - For `Firefox`, `Safari`: Manifest Version 2+.

## General Guide

- When building `ServiceWorker`s: Callbacks for extension events MUST be registered synchronously in order to get events handled on ServiceWorker startup, see [https://developer.chrome.com/docs/extensions/mv3/service_workers/events/](https://developer.chrome.com/docs/extensions/mv3/service_workers/events/)
  - In order to make sure that is the case for you WebExtension:
    - Load you wasm blob synchronously with `Application.initSync(src)` in the ServiceWorker script referenced by the `manifests.json#.background.service_worker`.
    - In your Go source code, there should be no `yield.Thread` call (which is called in `(js.Promise).Await`) before all desired event callbacks has been registered.
