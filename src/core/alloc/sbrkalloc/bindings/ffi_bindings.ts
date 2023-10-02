import { Application, importModule } from "@ffi";

importModule("core/alloc/sbrkalloc", (A: Application) => {
  return {
    "resetMemoryDataView": (): void => {
      A.resetMemoryDataView();
    },
  };
});
