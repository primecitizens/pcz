# {{ .Name }}

File paths:

```bash
export WASM_FILE="./main.wasm"
export BINDINGS_FILE="./bindings.js"
export EMBEDDING_FILE="./embed.js"
export ZIP_FILE="./webext.zip"
```

## Development

```bash
pcz dev webext ./ \
  --wasm "${WASM_FILE}" \
  --module-name "{{ .ModuleName }}" \
  --bindings "${BINDINGS_FILE}" \
  --embedding "${EMBEDDING_FILE}" \
  --embed "const wasmBlob=${WASM_FILE}" \
  --zip "${ZIP_FILE}"
```

## Release

1. Build and strip the wasm blob.

    ```bash
    pcz build ./ \
      -p js/wasm \
      -o "${WASM_FILE}" \
      --trimpath \
      -L-s \
      -L-w
    ```

2. Generate bindings for the wasm.

    ```bash
    pcz dev webext bindgen ./ \
      -o "${BINDINGS_FILE}" \
      --wasm "${WASM_FILE}" \
      --module-system "umd" \
      --module-name "{{ .ModuleName }}" \
      --minify \
      --es 6

    mv "${BINDINGS_FILE}.min.js" "${BINDINGS_FILE}"
    ```

3. Generate embedding for loading wasm synchronously

    ```bash
    pcz embed js \
      -o "${EMBEDDING_FILE}" \
      -f "const wasmBlob=${WASM_FILE}"
    ```

4. Package the web extension

    ```bash
    # this example uses the default glob to match files
    pcz dev webext pack -o "./webext.zip"
    ```

5. Distribute the web extension
   - [Chrome](https://developer.chrome.com/docs/extensions/mv3/linux_hosting/)
   - [Firefox](https://extensionworkshop.com/documentation/publish/package-your-extension/)
   - [Edge](https://learn.microsoft.com/en-us/microsoft-edge/extensions-chromium/publish/publish-extension)
   - [Safari](https://developer.apple.com/documentation/xcode/distributing-your-app-for-beta-testing-and-releases)
   - [Opera](https://dev.opera.com/extensions/publishing-guidelines/)
