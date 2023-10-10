(function () {
  const now = Date.now();

  const app = new globalThis.bindings.Application();
  app.args = ["popup", now.toString(10)];

  // wasmBlob is a Uint8Array generated from main.wasm (inside embed.js, loaded by popup.html).
  const res = app.initSync(wasmBlob.buffer);

  app.run(res.instance).then((code) => {
    console.log(`popup exit, code=${code}`);
  });
})();
