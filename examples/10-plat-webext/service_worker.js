importScripts("./bindings.js");

function run() {
  const app = new globalThis.bindings.Application();
  app.args = ["--service-worker", `${Date.now()}`];

  app.loadStream(fetch("/main.wasm")).then((res) => {
    app.run(res.instance).then((code) => {
      console.log(`program exit, code=${code}`);
    });
  });
}

const WEBEXT = globalThis.browser ? globalThis.browser : globalThis.chrome;

WEBEXT.runtime.onInstalled.addListener(run);
WEBEXT.runtime.onStartup.addListener(run);
