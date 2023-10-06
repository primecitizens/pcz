// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

import { Application, Pointer, heap, importModule } from "@ffi";

importModule("ui/html/builder", (A: Application) => {
  const registry = new Map<heap.Ref<HTMLElement>, string[]>();

  return {
    "reset": (elem: heap.Ref<HTMLElement>) => {
      if (registry.has(elem)) {
        registry.get(elem).splice(0);
        return;
      }

      registry.set(elem, []);
    },
    "free": (elem: heap.Ref<HTMLElement>) => {
      registry.delete(elem);
    },
    "integer": (elem: heap.Ref<HTMLElement>, signed: heap.Ref<boolean>, p64: Pointer, radix: number) => {
      registry.get(elem).push((signed === A.H.TRUE ? A.load.BigInt64(p64) : A.load.BigUint64(p64)).toString(radix));
    },
    "float": (elem: heap.Ref<HTMLElement>, num: number) => {
      registry.get(elem).push(num.toString(10));
    },
    "buf": (elem: heap.Ref<HTMLElement>, pBytes: Pointer, len: number): void => {
      registry.get(elem).push(A.load.String(pBytes, len));
    },
    "flush": (elem: heap.Ref<HTMLElement>, append: heap.Ref<boolean>): void => {
      const data = registry.get(elem).splice(0).join("");
      if (append === A.H.TRUE) {
        A.H.get<HTMLElement>(elem).innerHTML += data;
        return;
      }

      A.H.get<HTMLElement>(elem).innerHTML = data;
    },
  };
});
