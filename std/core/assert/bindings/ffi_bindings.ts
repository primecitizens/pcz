// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

import { Application, Pointer, importModule } from "@ffi";

importModule("core/assert", (A: Application) => {
  let buf: string[] = [];
  return {
    "append": (pBytes: Pointer, len: number) => {
      buf.push(A.load.String(pBytes, len));
    },
    "throw": () => {
      const msg = buf.join("");
      buf.splice(0);
      throw new Error(msg);
    },
  };
});
