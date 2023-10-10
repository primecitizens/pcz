# Example `hello`

## Quick Start

- `js/wasm`

    ```bash
    pcz dev web --port 8080 ./examples/00-basic-hello -- gopher
    ```

    Open up your local browser and navigate to `http://localhost:8080`, open console and you can find the output.

    > **NOTE**
    > For build and deployment examples, please see [../11-plat-web/README.md](../11-plat-web/README.md).

- `wasip1/wasm`

    ```bash
    pcz build -p wasip1/wasm -o ./bin/hello.wasm ./examples/00-basic-hello
    go run ./examples/runwasi.go ./bin/hello.wasm gopher
    # Hello gopher
    ```
