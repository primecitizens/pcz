const app = new globalThis.bindings.Application();
app.args = ["--popup", `${Date.now()}`];

app.loadStream(fetch("/main.wasm")).then((res) => {
  app.run(res.instance).then((code) => {
    console.log(`program exit, code=${code}`);
  });
});
