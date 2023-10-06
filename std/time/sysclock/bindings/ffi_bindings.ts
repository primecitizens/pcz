// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

import { Application, importModule } from "@ffi";

importModule("time/sysclock", (A: Application) => {
  return {
    "timeOriginMS": (): number => {
      return Date.now() - performance.now();
    },
    "walltime": (): number => {
      return Date.now();
    },
    "millitime": (): number => {
      return performance.now();
    },
  };
});
