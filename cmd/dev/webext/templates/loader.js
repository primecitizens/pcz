(function () {
  // {{- if eq .Comp "service-worker" }}
  importScripts("./bindings.js");
  importScripts("./embed.js");
  // + {{ end }}

  // + const app = new globalThis.{{ .ModuleName }}.Application();
  const app = new globalThis.bindings.Application();
  app.args = ["{{ .Comp }}"];

  // + const res = app.initSync({{ .WasmEmbed }});
  const res = app.initSync(wasmBlob.buffer);

  app.run(res.instance).then((code) => {
    console.log(`{{ .Comp }} exited, code=${code}`);
  });
})();
