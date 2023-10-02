// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

// NOTE: this file is only meant to expose APIs for module registration.

import { ModuleFactory, ModuleImporter } from "./bindgen";

export * from "./bindgen";

export namespace heap {
  export type Ref<T> = number;
}

export const importModule: ModuleImporter = (module: string, factory: ModuleFactory) => {};
