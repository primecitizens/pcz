// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

import { Application, Pointer, importModule } from "@ffi";

importModule("time/sysclock", (A: Application) => {
  return {
    "timeOriginMS": (pI64: Pointer) => {
      A.store.Int64(pI64, Date.now() - performance.now());
    },
    "walltime": (pI64: Pointer) => {
      A.store.Int64(pI64, Date.now());
    },
    "millitime": (pI64: Pointer) => {
      A.store.Int64(pI64, performance.now());
    },
  };
});
