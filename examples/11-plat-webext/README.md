# WebExtension Platform Example

- [Platform Development Guide](../../docs/user/11-plat-webext.md)

> **NOTE**
> This example only applies to `js/wasm` target and only runs in web browsers with manifest version 3 support.

A simple WebExtension created with `pcz dev webext new --name "Example Go WebExtension" --popup --service-worker ./examples/11-plat-webext`, that:

- Runs a ServiceWorker
  - Logs tab open/close, opening a new tab to `chrome://extensions` on tab closed.
  - Registers a context menu entry `Example Go ContextMenu` for selected text.
- Provides a popup action in toolbar that lists all bookmarks on click.

## Quick Start

```bash
export DIR="./examples/11-plat-webext"

pcz dev webext "${DIR}" \
  --wasm "${DIR}/main.wasm" \
  --bindings "${DIR}/bindings.js" \
  --embedding "${DIR}/embed.js" \
  --embed "var wasmBlob=${DIR}/main.wasm" \
  --zip "${DIR}/webext.zip" \
  --pack-base "${DIR}"
```

Open your browser, and load the built extension (`${DIR}/webext.zip`), if you're using a `Chromium`-like browser, you can load the unpacked extension from the `${DIR}`.

Then:

- Right click the icon (should be a `E`) to inspect popup.
- Select some text and right click to find the example menu entry.
- Create a new tab and close it.
