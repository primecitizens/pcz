import { Application, Pointer, heap, importModule } from "@ffi";

importModule("ffi/js/array", (A: Application) => {
  type TypedArray =
    | Uint16Array
    | Int16Array
    | Uint32Array
    | Int32Array
    | Float32Array
    | Float64Array
    | BigUint64Array
    | BigInt64Array;

  // append go numeric values to js array.
  const append = (
    load: (ptr: number) => number | bigint,
    elemSz: number,
    refJsArray: heap.Ref<TypedArray>,
    offsetInJsArray: number,
    goArray: Pointer,
    goArrayLen: number
  ): number => {
    const jsArray = A.H.get<TypedArray>(refJsArray);

    if (jsArray.length <= offsetInJsArray) return 0;

    const cap = jsArray.length - offsetInJsArray;
    if (cap < goArrayLen) {
      goArrayLen = cap;
    }

    for (let i = offsetInJsArray; i < goArrayLen; i++) {
      jsArray[i] = load(goArray + i * elemSz);
    }
    return goArrayLen;
  };

  // copy js numeric values in array to go array.
  const copy = (
    store: (ptr: number, val: number | bigint) => void,
    elemSz: number,
    refJsArray: heap.Ref<TypedArray>,
    offsetInJsArray: number,
    goArray: Pointer,
    goArrayLen: number
  ): number => {
    const jsArray = A.H.get<TypedArray>(refJsArray);

    if (jsArray.length <= offsetInJsArray) return 0;

    const cap = jsArray.length - offsetInJsArray;
    if (cap < goArrayLen) {
      goArrayLen = cap;
    }

    for (let i = offsetInJsArray; i < goArrayLen; i++) {
      store(goArray + i * elemSz, jsArray[i]);
    }
    return goArrayLen;
  };

  return {
    //
    // Array
    //
    "new": (
      sz: number,
      elemSz: number,
      signed: heap.Ref<boolean>,
      float: heap.Ref<boolean>
    ): heap.Ref<Array<any> | TypedArray> => {
      switch (elemSz) {
        case 0:
          return A.H.push(new Array(sz));
        case 1:
          return signed === A.H.TRUE ? A.H.push(new Int8Array(sz)) : A.H.push(new Uint8Array(sz));
        case 2:
          return signed === A.H.TRUE ? A.H.push(new Int16Array(sz)) : A.H.push(new Uint16Array(sz));
        case 4:
          if (float === A.H.TRUE) return A.H.push(new Float32Array(sz));

          return signed === A.H.TRUE ? A.H.push(new Int32Array(sz)) : A.H.push(new Uint32Array(sz));
        case 8:
          if (float === A.H.TRUE) return A.H.push(new Float64Array(sz));

          // available since ES2020
          return signed === A.H.TRUE ? A.H.push(new BigInt64Array(sz)) : A.H.push(new BigUint64Array(sz));
        default:
          throw new Error(`unsupported element size=${elemSz}`);
      }
    },
    "length": (arr: heap.Ref<Array<any> | TypedArray>): number => {
      return A.H.get<Array<any> | TypedArray>(arr).length;
    },
    "slice": (
      arr: heap.Ref<Array<any> | TypedArray>,
      start: number,
      end: number
    ): heap.Ref<Array<any> | TypedArray> => {
      if (end) {
        return A.H.push(A.H.get<Array<any> | TypedArray>(arr).slice(start, end));
      }

      return A.H.push(A.H.get<Array<any> | TypedArray>(arr).slice(start));
    },
    "append": (
      refJsArray: heap.Ref<Array<any> | TypedArray>,
      take: heap.Ref<boolean>,
      elemSz: number,
      signed: heap.Ref<boolean>,
      float: heap.Ref<boolean>,
      offsetInJsArray: number,
      goArray: Pointer,
      goArrayLen: number
    ): number => {
      switch (elemSz) {
        case 0: {
          let data = A.load.Refs(goArray, goArrayLen, take === A.H.TRUE);
          const a = A.H.get<Array<any>>(refJsArray);
          let j = 0;
          for (let i = offsetInJsArray; i < a.length && j < data.length; i++) {
            a[i] = data[j];
            j++;
          }
          A.H.replace(refJsArray, a.concat(data.slice(j)));
          return goArrayLen;
        }
        case 1: {
          const a = A.H.get<Int8Array | Uint8Array | Uint8ClampedArray>(refJsArray);
          if (a.length <= offsetInJsArray) return 0;
          const cap = a.length - offsetInJsArray;
          if (cap < goArrayLen) goArrayLen = cap;
          a.set(A.load.Raw(goArray, goArrayLen), offsetInJsArray);
          return goArrayLen;
        }
        case 2:
          return signed === A.H.TRUE
            ? append(A.load.Int16, 2, refJsArray, offsetInJsArray, goArray, goArrayLen)
            : append(A.load.Uint16, 2, refJsArray, offsetInJsArray, goArray, goArrayLen);
        case 4:
          if (float === A.H.TRUE) return append(A.load.Float32, 4, refJsArray, offsetInJsArray, goArray, goArrayLen);

          return signed === A.H.TRUE
            ? append(A.load.Int32, 4, refJsArray, offsetInJsArray, goArray, goArrayLen)
            : append(A.load.Uint32, 4, refJsArray, offsetInJsArray, goArray, goArrayLen);
        case 8:
          if (float === A.H.TRUE) return append(A.load.Float64, 8, refJsArray, offsetInJsArray, goArray, goArrayLen);

          return signed === A.H.TRUE
            ? append(A.load.Int64, 8, refJsArray, offsetInJsArray, goArray, goArrayLen)
            : append(A.load.Uint64, 8, refJsArray, offsetInJsArray, goArray, goArrayLen);
        default:
          throw new Error(`unsupported element size=${elemSz}`);
      }
    },
    "copy": (
      refJsArray: heap.Ref<Array<any> | TypedArray>,
      elemSz: number,
      signed: heap.Ref<boolean>,
      float: heap.Ref<boolean>,
      offsetInJsArray: number,
      goArray: Pointer,
      goArrayLen: number
    ): number => {
      switch (elemSz) {
        case 1:
          const a = A.H.get<Int8Array | Uint8Array | Uint8ClampedArray>(refJsArray);
          if (a.length <= offsetInJsArray) return 0;
          const cap = a.length - offsetInJsArray;
          if (cap < goArrayLen) goArrayLen = cap;
          A.load.Raw(goArray, goArrayLen).set(a, offsetInJsArray);
          return goArrayLen;
        case 2:
          return signed === A.H.TRUE
            ? copy(A.store.Int16, 2, refJsArray, offsetInJsArray, goArray, goArrayLen)
            : copy(A.store.Uint16, 2, refJsArray, offsetInJsArray, goArray, goArrayLen);
        case 4:
          if (float === A.H.TRUE) return copy(A.store.Float32, 4, refJsArray, offsetInJsArray, goArray, goArrayLen);

          return signed === A.H.TRUE
            ? copy(A.store.Int32, 4, refJsArray, offsetInJsArray, goArray, goArrayLen)
            : copy(A.store.Uint32, 4, refJsArray, offsetInJsArray, goArray, goArrayLen);
        case 8:
          if (float === A.H.TRUE) return copy(A.store.Float64, 8, refJsArray, offsetInJsArray, goArray, goArrayLen);

          return signed === A.H.TRUE
            ? copy(A.store.Int64, 8, refJsArray, offsetInJsArray, goArray, goArrayLen)
            : copy(A.store.Uint64, 8, refJsArray, offsetInJsArray, goArray, goArrayLen);
        default:
          throw new Error(`unsupported element size=${elemSz}`);
      }
    },
    "set": (a: heap.Ref<TypedArray>, b: heap.Ref<TypedArray>): void => {
      A.H.get<Uint32Array>(a).set(A.H.get<Uint32Array>(b));
    },
    "buffer": (take: heap.Ref<boolean>, ref: heap.Ref<TypedArray>): heap.Ref<ArrayBufferLike> => {
      return A.H.push((take === A.H.TRUE ? A.H.take<TypedArray>(ref) : A.H.get<TypedArray>(ref)).buffer);
    },
    "fromBuffer": (
      take: heap.Ref<boolean>,
      refBuf: heap.Ref<ArrayBuffer | SharedArrayBuffer>,
      elemSz: number,
      signed: heap.Ref<boolean>,
      float: heap.Ref<boolean>
    ): heap.Ref<TypedArray> => {
      const buf = (take === A.H.TRUE ? A.H.take(refBuf) : A.H.get(refBuf)) as ArrayBuffer;
      switch (elemSz) {
        case 1:
          return signed === A.H.TRUE ? A.H.push(new Int8Array(buf)) : A.H.push(new Uint8Array(buf));
        case 2:
          return signed === A.H.TRUE ? A.H.push(new Int16Array(buf)) : A.H.push(new Uint16Array(buf));
        case 4:
          if (float === A.H.TRUE) return A.H.push(new Float32Array(buf));

          return signed === A.H.TRUE ? A.H.push(new Int32Array(buf)) : A.H.push(new Uint32Array(buf));
        case 8:
          if (float === A.H.TRUE) return A.H.push(new Float64Array(buf));

          return signed === A.H.TRUE ? A.H.push(new BigInt64Array(buf)) : A.H.push(new BigUint64Array(buf));
        default:
          throw new Error(`unsupported element size=${elemSz}`);
      }
    },
    "byteOffset": (ref: heap.Ref<TypedArray>): number => {
      return A.H.get<TypedArray>(ref).byteOffset;
    },
    "byteLength": (ref: heap.Ref<TypedArray>): number => {
      return A.H.get<TypedArray>(ref).byteLength;
    },
    "nth": (
      self: heap.Ref<Array<any>>,
      elemSz: number,
      signed: heap.Ref<boolean>,
      float: heap.Ref<boolean>,
      i: number,
      ptr: number
    ): heap.Ref<any> => {
      const thiz = A.H.get<Array<any> | TypedArray>(self);
      if (i >= thiz.length) return A.H.FALSE;

      switch (elemSz) {
        case 0:
          A.store.Ref(ptr, thiz[i]);
          break;
        case 1:
          if (signed === A.H.TRUE) {
            A.store.Int8(ptr, thiz[i]);
          } else {
            A.store.Uint8(ptr, thiz[i]);
          }
          break;
        case 2:
          if (signed === A.H.TRUE) {
            A.store.Int16(ptr, thiz[i]);
          } else {
            A.store.Uint16(ptr, thiz[i]);
          }
          break;
        case 4:
          if (float === A.H.TRUE) {
            A.store.Float32(ptr, thiz[i]);
          } else if (signed === A.H.TRUE) {
            A.store.Int32(ptr, thiz[i]);
          } else {
            A.store.Uint32(ptr, thiz[i]);
          }
          break;
        case 8:
          if (float === A.H.TRUE) {
            A.store.Float64(ptr, thiz[i]);
          } else if (signed === A.H.TRUE) {
            A.store.Int64(ptr, thiz[i]);
          } else {
            A.store.Uint64(ptr, thiz[i]);
          }
          break;
        default:
          throw new Error(`unsupported element size=${elemSz}`);
      }

      return A.H.TRUE;
    },
    "nthNum": (self: heap.Ref<Array<any>>, i: number, ptr: number): heap.Ref<boolean> => {
      const thiz = A.H.get<Array<any>>(self);
      if (i >= thiz.length) return A.H.FALSE;

      A.store.Float64(ptr, thiz[i]);
      return A.H.TRUE;
    },
    "nthBool": (self: heap.Ref<Array<any>>, i: number, ptr: number): heap.Ref<boolean> => {
      const thiz = A.H.get<Array<any>>(self);
      if (i >= thiz.length) return A.H.FALSE;

      A.store.Bool(ptr, thiz[i] ? true : false);
      return A.H.TRUE;
    },

    "setNth": (
      self: heap.Ref<Array<any>>,
      elemSz: number,
      signed: heap.Ref<boolean>,
      float: heap.Ref<boolean>,
      i: number,
      take: heap.Ref<boolean>,
      ptr: Pointer
    ): heap.Ref<boolean> => {
      const thiz = A.H.get<Array<any>>(self);
      if (i >= thiz.length) return A.H.FALSE;

      switch (elemSz) {
        case 0:
          thiz[i] = A.load.Ref(ptr, undefined, take === A.H.TRUE);
          break;
        case 1:
          if (signed === A.H.TRUE) {
            thiz[i] = A.load.Int8(ptr);
          } else {
            thiz[i] = A.load.Uint8(ptr);
          }
          break;
        case 2:
          if (signed === A.H.TRUE) {
            thiz[i] = A.load.Int16(ptr);
          } else {
            thiz[i] = A.load.Uint16(ptr);
          }
          break;
        case 4:
          if (float === A.H.TRUE) {
            thiz[i] = A.load.Float32(ptr);
          } else if (signed === A.H.TRUE) {
            thiz[i] = A.load.Int32(ptr);
          } else {
            thiz[i] = A.load.Uint32(ptr);
          }
          break;
        case 8:
          if (float === A.H.TRUE) {
            thiz[i] = A.load.Float64(ptr);
          } else if (signed === A.H.TRUE) {
            thiz[i] = A.load.Int64(ptr);
          } else {
            thiz[i] = A.load.Uint64(ptr);
          }
          break;
        default:
          throw new Error(`unsupported element size=${elemSz}`);
      }
      return A.H.TRUE;
    },
    "setNthNum": (self: heap.Ref<Array<any>>, i: number, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<Array<any>>(self);
      if (i >= thiz.length) return A.H.FALSE;
      thiz[i] = val;
      return A.H.TRUE;
    },
    "setNthBool": (self: heap.Ref<Array<any>>, i: number, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<Array<any>>(self);
      if (i >= thiz.length) return A.H.FALSE;
      thiz[i] = val === A.H.TRUE ? true : false;
      return A.H.TRUE;
    },
    "setNthString": (self: heap.Ref<Array<any>>, i: number, ptr: number, sz: number): heap.Ref<boolean> => {
      const thiz = A.H.get<Array<any>>(self);
      if (i >= thiz.length) return A.H.FALSE;
      thiz[i] = A.load.String(ptr, sz);
      return A.H.TRUE;
    },
  };
});
