// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

import { Application, Pointer, importModule } from "@ffi";

importModule("stdprint", (App: Application) => {
  const buf = new Array<string>();
  let timeoutID: any; // using type any to workaround node js Timeout type.

  return {
    "print": (str: Pointer, len: number): void => {
      let s = App.load.String(str, len);
      const lf = s.lastIndexOf("\n");
      if (lf >= 0) {
        if (timeoutID) {
          clearTimeout(timeoutID);
          timeoutID = undefined;
        }

        console.log(buf.splice(0, buf.length).join("") + s.substring(0, lf));
        if (lf === s.length - 1) {
          return;
        }

        s = s.substring(lf + 1);
      }

      buf.push(s);
      if (timeoutID) {
        return; // only keep one timeout job going
      }

      timeoutID = setTimeout(() => {
        timeoutID = undefined;
        console.log(buf.splice(0, buf.length).join(""));
      }, 100);
      return;
    },
  };
});
