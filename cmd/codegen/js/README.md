# js platform codegen

API sources:

- [w3c/webref@curated ed/idl](https://github.com/w3c/webref)
- [mdn/browser-compat-data](https://github.com/mdn/browser-compat-data)
- [chromium source](https://source.chromium.org/chromium/chromium/src) ([github mirror](https://github.com/chromium/chromium))

## Maintenance Notes

The BCD (Browser Compat Data) can be downloaded from the release area of the official MDN repository (`https://github.com/mdn/browser-compat-data/releases`, find `data.json`, about 12MB in size) or download from CDN sites like `https://unpkg.com/@mdn/browser-compat-data/data.json`

To update idl files in `webext`:

- Checkout chromium repo tree (here we use `main` branch)
  - `git clone --depth=1 -b main https://github.com/chromium/chromium`
  - `export CHROMIUM_ROOT="$(pwd)/chromium"`
- Copy all webextension api declaration files to `./webext`:
  - `rm -rf ./webext && mkdir -p webext/{chrome, chromeos}`
  - `cp "${CHROMIUM_ROOT}/extensions/common/api/"*.{idl,json} ./webext`
  - `cp "${CHROMIUM_ROOT}/chrome/common/extensions/api/"*.{idl,json} ./webext/chrome`
  - `cp "${CHROMIUM_ROOT}/chrome/common/chromeos/extensions/api/"*.{idl,json} ./webext/chromeos`
- Open each idl file (`./webext/**/*.idl`)
  - Remove `^  `
  - Replace `^(dictionary|interface|callback|enum) ` with `[Namespace=(chrome, )]\n$1 `
  - Replace `(\[.*?\]) (dictionary|interface|callback|enum) ` with `[Namespace=(chrome, )]\n$1 $2 `
  - Replace `[Namespace=(chrome, )]` with correct namespace value, append platform values.
  - Remove `namespace xxx {` line and its extended attributes, remove trailing `};` line.
  - Replace `(\w+)\[\]` with `sequence<$1>`.
  - Merge multiple Extended-Attributes into one
    - mostly `[inline_doc]` and `[noinline_doc]`, and we just remove these.
    - `\]\n\[(.*)\] ` to `, $1]\n`
    - `\]\n\[(.*)\]\n` to `, $1]\n`
  - Quote enum values. (a lot, replace `^  (\w+),` to `  "$1",`)
  - Replace `any?` to `any`
  - Rewrite dictionary definitions with methods to interface definitions
  - Run `pcz codegen` to check other errors.

Refs:

- `${CHROMIUM_ROOT}/chrome/common/extensions/api/schemas.md`
- `${CHROMIUM_ROOT}/extensions/renderer/bindings.md`
- `${CHROMIUM_ROOT}/third_party/blink/renderer/bindings/IDLExtendedAttributes.md`
- `https://github.com/mozilla/webextension-polyfill/blob/master/src/browser-polyfill.js`
- `https://github.com/w3c/webextensions`
- `https://developer.chrome.com/docs/extensions/mv2/architecture-overview/#sync`

Commands to generate libraries (from repo root):

```bash
go run ./ codegen js -o ./src/plat/js --webref "${WEBREF_HOME}"
prettier "./src/plat/js/**/bindings/ffi_bindings.ts" --write
```
