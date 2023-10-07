// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

import { importModule, heap, Application } from "@ffi";

importModule("ffi/js", (A: Application): Record<string, Function> => {
  return {
    //
    // Heap management
    //
    "once": (ref: heap.Ref<any>): void => {
      A.H.freeAfterNGets(ref, 1);
    },
    "clone": (ref: heap.Ref<any>): heap.Ref<any> => {
      return A.H.push(A.H.get(ref));
    },
    "free": (ref: heap.Ref<any>): void => {
      A.H.take(ref);
    },
    "batchFree": (ptr: number, n: number): void => {
      for (let i = 0; i < n; i++) {
        A.H.free(A.load.Uint32(ptr + i * 4));
      }
    },
    "replace": (old: heap.Ref<any>, ne: heap.Ref<any>): heap.Ref<boolean> => {
      return A.H.replace(old, A.H.take(ne));
    },
    "replaceNum": (old: heap.Ref<any>, num: number): heap.Ref<boolean> => {
      return A.H.replace(old, num);
    },
    "replaceBigInt": (old: heap.Ref<any>, signed: heap.Ref<boolean>, ptr: number): heap.Ref<boolean> => {
      return signed === A.H.TRUE ? A.H.replace(old, A.load.Int64(ptr)) : A.H.replace(old, A.load.Uint64(ptr));
    },
    "replaceBool": (old: heap.Ref<any>, bool: heap.Ref<boolean>): heap.Ref<boolean> => {
      return A.H.replace(old, bool === A.H.TRUE ? true : false);
    },
    "replaceString": (old: heap.Ref<any>, ptr: number, sz: number): heap.Ref<boolean> => {
      return A.H.replace(old, A.load.String(ptr, sz));
    },

    //
    // Number
    //
    "number": (n: number): heap.Ref<number> => {
      return A.H.push(n);
    },
    "getnum": (ref: heap.Ref<number>, ptr: number): heap.Ref<boolean> => {
      const x = A.H.get<number>(ref);
      if (typeof x !== "number") return A.H.FALSE;
      A.store.Float64(ptr, x);
      return A.H.TRUE;
    },
    "bigint": (signed: heap.Ref<boolean>, ptr: number): heap.Ref<number> => {
      return signed === A.H.TRUE ? A.H.push(A.load.Int64(ptr)) : A.H.push(A.load.Uint64(ptr));
    },
    "getBigInt": (ref: heap.Ref<bigint>, signed: heap.Ref<boolean>, ptr: number): heap.Ref<boolean> => {
      const x = A.H.get<bigint | number>(ref);
      // TODO: do we actually want to limit the type to BigInt?
      //       it seems most of the time when the WebIDL expects a int64
      //       js runtime only provides a "number" rather than BigInt.
      // if (typeof x !== "bigint") return A.H.FALSE;
      if (signed === A.H.TRUE) {
        A.store.Int64(ptr, x);
      } else {
        if (x < 0) {
          throw new Error("unexpected negative uint64 value");
        }
        A.store.Uint64(ptr, x);
      }
      return A.H.TRUE;
    },

    //
    // String
    //
    "encodeUTF8": (s: heap.Ref<string>, ptr: number, len: number): number => {
      return A.store.String(ptr, len, A.H.get<string>(s)).n;
    },
    "decodeUTF8": (ptr: number, len: number): heap.Ref<string> => {
      return A.H.push(A.load.String(ptr, len));
    },
    "appendUTF8": (s: heap.Ref<string>, ptr: number, len: number): heap.Ref<string> => {
      return A.H.replace(s, A.H.get<string>(s) + A.load.String(ptr, len));
    },
    "prependUTF8": (s: heap.Ref<string>, ptr: number, len: number): heap.Ref<string> => {
      return A.H.replace(s, A.load.String(ptr, len) + A.H.get<string>(s));
    },
    "equalsUTF8": (s: heap.Ref<string>, ptr: number, len: number): heap.Ref<boolean> => {
      return A.load.String(ptr, len) === A.H.get<string>(s) ? A.H.TRUE : A.H.FALSE;
    },
    "sizeUTF8": (s: heap.Ref<string>): number => {
      return A.UTF8Sizeof(A.H.get<string>(s));
    },

    //
    // Func
    //

    "try": (
      fn: heap.Ref<Function>,
      self: heap.Ref<any>,
      free: heap.Ref<boolean>,
      _catch: heap.Ref<Function>,
      _finally: heap.Ref<Function>,
      ptr: number,
      n: number,
      pCaught: number
    ): heap.Ref<any> => {
      const thiz = A.H.get(self);
      const args = A.load.Refs(ptr, n, free === A.H.TRUE);
      const katch = free === A.H.TRUE ? A.H.take<Function>(_catch) : A.H.get<Function>(_catch);
      const vinally = free === A.H.TRUE ? A.H.take<Function>(_finally) : A.H.get<Function>(_finally);

      try {
        return A.H.push(Reflect.apply(A.H.get<Function>(fn), thiz, args));
      } catch (error) {
        A.store.Bool(pCaught, true);

        if (katch) {
          return A.H.push(Reflect.apply(katch, thiz, [error, ...args]));
        }
        return A.H.UNDEFINED;
      } finally {
        if (vinally) {
          Reflect.apply(vinally, thiz, args);
        }
      }
    },
    "call": (
      fn: heap.Ref<Function>,
      self: heap.Ref<any>,
      free: heap.Ref<boolean>,
      ptr: number,
      n: number
    ): heap.Ref<any> => {
      return A.H.push(Reflect.apply(A.H.get<Function>(fn), A.H.get(self), A.load.Refs(ptr, n, free === A.H.TRUE)));
    },
    "callVoid": (
      fn: heap.Ref<Function>,
      self: heap.Ref<any>,
      free: heap.Ref<boolean>,
      ptr: number,
      n: number
    ): void => {
      Reflect.apply(A.H.get<Function>(fn), A.H.get(self), A.load.Refs(ptr, n, free === A.H.TRUE));
    },
    "callNum": (
      fn: heap.Ref<Function>,
      self: heap.Ref<any>,
      free: heap.Ref<boolean>,
      ptr: number,
      n: number
    ): number => {
      return Reflect.apply(A.H.get<Function>(fn), A.H.get(self), A.load.Refs(ptr, n, free === A.H.TRUE)) as number;
    },
    "callBool": (
      fn: heap.Ref<Function>,
      self: heap.Ref<any>,
      free: heap.Ref<boolean>,
      ptr: number,
      n: number
    ): heap.Ref<boolean> => {
      return Reflect.apply(A.H.get<Function>(fn), A.H.get(self), A.load.Refs(ptr, n, free === A.H.TRUE))
        ? A.H.TRUE
        : A.H.FALSE;
    },

    //
    // Object
    //
    "new": (constructor: heap.Ref<Function>, free: heap.Ref<boolean>, ptr: number, n: number): heap.Ref<any> => {
      return A.H.push(Reflect.construct(A.H.get(constructor), A.load.Refs(ptr, n, free === A.H.TRUE)));
    },
    "instanceof": (a: heap.Ref<any>, cls: heap.Ref<any>): heap.Ref<boolean> => {
      return A.H.get<any>(a) instanceof A.H.get<any>(cls) ? A.H.TRUE : A.H.FALSE;
    },

    "getPrototype": (self: heap.Ref<object>): heap.Ref<object> => {
      const ret = Reflect.getPrototypeOf(A.H.get<object>(self));
      if (ret) {
        return A.H.push(ret);
      }
      return A.H.UNDEFINED;
    },
    "setPrototype": (
      self: heap.Ref<object>,
      free: heap.Ref<object>,
      prototype: heap.Ref<object>
    ): heap.Ref<boolean> => {
      const proto = free === A.H.TRUE ? A.H.take<object>(prototype) : A.H.get<object>(prototype);
      return Reflect.setPrototypeOf(A.H.get<object>(self), proto) ? A.H.TRUE : A.H.FALSE;
    },

    "getPropDesc": (self: heap.Ref<object>, ptr: number, sz: number): number => {
      const desc = Reflect.getOwnPropertyDescriptor(A.H.get<object>(self), A.load.String(ptr, sz));
      if (!desc) {
        return 0;
      }
      // NOTE: MUST agree with ffi/wasm/js/object.go#PropertyDescription_xxx
      let flags: number = 1;
      flags += desc.writable ? 2 : 0;
      flags += desc.configurable ? 4 : 0;
      flags += desc.enumerable ? 8 : 0;
      flags += desc.get ? 16 : 0;
      flags += desc.set ? 32 : 0;
      return flags;
    },
    "defineProp": (
      self: heap.Ref<object>,
      ptr: number,
      sz: number,
      flags: number,
      getter: heap.Ref<Function>,
      setter: heap.Ref<Function>
    ): heap.Ref<boolean> => {
      return Reflect.defineProperty(A.H.get<object>(self), A.load.String(ptr, sz), {
        writable: (flags & 2) !== 0 ? true : false,
        configurable: (flags & 4) !== 0 ? true : false,
        enumerable: (flags & 8) !== 0 ? true : false,
        get: (flags & 16) !== 0 ? (A.H.get<Function>(getter) as () => any) : undefined,
        set: (flags & 32) !== 0 ? (A.H.get<Function>(setter) as (v: any) => void) : undefined,
      })
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "deleteProp": (self: heap.Ref<Object>, ptr: number, sz: number): heap.Ref<boolean> => {
      return Reflect.deleteProperty(A.H.get<object>(self), A.load.String(ptr, sz)) ? A.H.TRUE : A.H.FALSE;
    },
    "prop": (self: heap.Ref<Object>, ptr: number, sz: number): heap.Ref<any> => {
      const thiz = A.H.get<object>(self);
      return A.H.push(Reflect.get(thiz, A.load.String(ptr, sz), thiz));
    },
    "numProp": (self: heap.Ref<Object>, ptr: number, sz: number): number => {
      const thiz = A.H.get<object>(self);
      return Reflect.get(thiz, A.load.String(ptr, sz), thiz) as number;
    },
    "boolProp": (self: heap.Ref<Object>, ptr: number, sz: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      return Reflect.get(thiz, A.load.String(ptr, sz), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "propByString": (self: heap.Ref<Object>, name: heap.Ref<string>): heap.Ref<any> => {
      const thiz = A.H.get<object>(self);
      return A.H.push(Reflect.get(thiz, A.H.get<string>(name), thiz));
    },
    "numPropByString": (self: heap.Ref<Object>, name: heap.Ref<string>): number => {
      const thiz = A.H.get<object>(self);
      return Reflect.get(thiz, A.H.get<string>(name), thiz) as number;
    },
    "boolPropByString": (self: heap.Ref<Object>, name: heap.Ref<string>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      return Reflect.get(thiz, A.H.get<string>(name), thiz) ? A.H.TRUE : A.H.FALSE;
    },

    "setProp": (
      self: heap.Ref<Object>,
      ptr: number,
      sz: number,
      free: heap.Ref<boolean>,
      val: heap.Ref<any>
    ): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      return Reflect.set(thiz, A.load.String(ptr, sz), free === A.H.TRUE ? A.H.take(val) : A.H.get(val), thiz)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "setNumProp": (self: heap.Ref<Object>, ptr: number, sz: number, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      return Reflect.set(thiz, A.load.String(ptr, sz), val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "setBoolProp": (self: heap.Ref<Object>, ptr: number, sz: number, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      return Reflect.set(thiz, A.load.String(ptr, sz), val === A.H.TRUE ? true : false, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "setStringProp": (
      self: heap.Ref<Object>,
      nameptr: number,
      namesz: number,
      valptr: number,
      valsz: number
    ): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      return Reflect.set(thiz, A.load.String(nameptr, namesz), A.load.String(valptr, valsz), thiz)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "setPropByString": (
      self: heap.Ref<Object>,
      name: heap.Ref<string>,
      free: heap.Ref<boolean>,
      val: heap.Ref<any>
    ): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      return Reflect.set(thiz, A.H.get<string>(name), free === A.H.TRUE ? A.H.take(val) : A.H.get(val), thiz)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "setNumPropByString": (self: heap.Ref<Object>, name: heap.Ref<string>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      return Reflect.set(thiz, A.H.get<string>(name), val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "setBoolPropByString": (self: heap.Ref<Object>, name: heap.Ref<string>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      return Reflect.set(thiz, A.H.get<string>(name), val === A.H.TRUE ? true : false, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "setStringPropByString": (
      self: heap.Ref<Object>,
      name: heap.Ref<string>,
      free: heap.Ref<boolean>,
      val: heap.Ref<string>
    ): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      return Reflect.set(thiz, A.H.get<string>(name), free === A.H.TRUE ? A.H.take(val) : A.H.get<string>(val), thiz)
        ? A.H.TRUE
        : A.H.FALSE;
    },
  };
});
