async function setup() {
  const go = new globalThis.Go();

  const result = await WebAssembly.instantiateStreaming(
    fetch("go.wasm"),
    go.importObject
  );

  go.run(result.instance);

  globalThis.setup()
}

setup().catch(console.error);
