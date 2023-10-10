(function () {
  importScripts("./bindings.js");
  importScripts("./embed.js");

  const now = Date.now();

  const app = new globalThis.bindings.Application();
  app.args = ["service-worker", now.toString(10)];

  // wasmBlob is a Uint8Array generated from main.wasm (inside embed.js).
  const res = app.initSync(wasmBlob.buffer);

  app.run(res.instance).then((code) => {
    console.log(`service-worker exit, code=${code}`);
  });
})();
