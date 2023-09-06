// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

// NOTE: this file is only meant to expose APIs for module registration
// and serves as requirements to ./bindgen.ts#Application.

import { Heap } from "./bindgen";

export function addModuleImport(
  namespace: string,
  entries: Record<string, Function>
) {}

export function loadArgs(ptr: number, n: number, free: boolean): any[] {
  return [];
}

export namespace heap {
  export type Ref<T> = number;
}

export const HEAP = new Heap();
