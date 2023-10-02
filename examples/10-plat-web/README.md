# Web Platform Example

> **NOTE**
> This example only applies to `js/wasm` target and only runs in web browsers.

## Quick Start

```bash
pcz dev web --port 8080 ./examples/10-plat-web
```

Open up your local browser and navigate to `http://localhost:8080`, click button to activate/switch example.

## Build (and prepare for deployment)

1. Build the wasm blob, the [example-fetch.go](./example-fetch.go) contains a link-time variable (`wasmPath`), you may want to set it to the public path of your wasm.

    ```bash
    export WASM_FILE="./dist/app.wasm"
    export WASM_WEB_PATH="/static/app.wasm"

    pcz build -p js/wasm \
      -o "${WASM_FILE}" \
      -X "main.wasmPath=${WASM_WEB_PATH}" \
      ./examples/10-plat-web
    ```

2. Generate required bindings for your wasm blob to run in web browser, the command below will generate `./dist/bindings.js` in es6 syntax (as specified by `--es 6`), the module name is set to `bindings`, with module system support for `commonjs`, `amd.js` and web browser global (as specified by `--mode umd`).

    ```bash
    # NOTE: the WASM_FILE here is the one we defined in the first step.
    export BINDINGS_FILE="./dist/bindings.js"
    pcz dev web bindgen \
      --wasm "${WASM_FILE}" \
      -o "${BINDINGS_FILE}" \
      --name bindings \
      --es 6 \
      --mode umd \
      ./examples/10-plat-web
    ```

    You may want to add cli flag `--minify`, which will generate 2 additional files:
      - `./dist/bindings.min.js` (minified version of `./dist/bindings.js`)
      - `./dist/bindings.min.js.map` (sourcemap of `./dist/bindings.min.js`)

3. Write a `html` file to load `${WASM_FILE}` and `${BINDINGS_FILE}`, here we will use a basic one ([./index.html](./index.html)) that runs the wasm application on page load, remember to update `URL`s in the `<script src="URL">` and `fetch(URL)` according to your own deployment environment.
