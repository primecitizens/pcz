// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

import { Application, importModule } from "@ffi";

importModule("core/alloc/sbrkalloc", (A: Application) => {
  return {
    "resetMemoryDataView": (): void => {
      A.resetMemoryDataView();
    },
  };
});
