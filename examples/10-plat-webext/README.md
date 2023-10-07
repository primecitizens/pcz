# WebExtension Platform Example

> **NOTE**
> This example only applies to `js/wasm` target and only runs in web browsers with manifest version 3 support.

A simple web extension that:

- Runs a ServiceWorker in the background
  - logging tab open/close, opening a new tab to `chrome://extensions` on tab closed.
  - providing a context menu entry `Example WASM ContextMenu` for selected text.

  > **ATTENTION**
  > The ServiceWorker example doesn't work after being inactive. (**HELP WANTED**)

- Provides a popup action in toolbar that lists all bookmarks on click.

## Quick Start

> **NOTE**
> In order to load unpacked extension, you need to have a `chromium`-like browser installed, and enable development mode on extension management page (`chrome://extensions`).

```bash
pcz dev webext \
  --bindings ./examples/10-plat-webext/bindings.js \
  --wasm ./examples/10-plat-webext/main.wasm \
  ./examples/10-plat-webext
```

Open up your `chromium`-like browser, load unpacked extension from the example directory (`./examples/10-plat-webext`), then:

- right click the icon (should be a `E`) to inspect popup.
- select some text and right click to find the example context menu entry.
- create a new tab and close it.
