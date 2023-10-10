// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

export namespace heap {
  export type Ref<T> = number;
}

export type CallbackFunc = (this: any, ...args: any[]) => void;

export type Pointer = number;

export type ModuleFactory = (App: Application) => Record<string, Function>;

export type ModuleImporter = (module: string, factory: ModuleFactory) => void;

// following constants MUST agree with constant values in ../js.go#const:smallIntCacheSize
const UNDEFINED = 0;
const NULL = 1;
const ZERO = 2;
const FALSE = 3;
const TRUE = 4;
const Const_NaN = 5;
const INF = 6;
const NEG_INF = 7;
const GLOBAL = 8;
const _LAST_CONSTANT_REF_ = GLOBAL;
const SMALL_INT_CACHE_SIZE = 128;

const _static_heap_size_ = SMALL_INT_CACHE_SIZE + (_LAST_CONSTANT_REF_ + 1);

export class Heap {
  // index MUST agree with constant values in ../js.go#const
  public readonly UNDEFINED: heap.Ref<undefined> = UNDEFINED;
  public readonly NULL: heap.Ref<null> = NULL;
  public readonly ZERO: heap.Ref<number> = ZERO;
  public readonly FALSE: heap.Ref<boolean> = FALSE;
  public readonly TRUE: heap.Ref<boolean> = TRUE;
  public readonly NaN: heap.Ref<number> = Const_NaN;
  public readonly INF: heap.Ref<number> = INF;
  public readonly NEG_INF: heap.Ref<number> = NEG_INF;
  public readonly GLOBAL: heap.Ref<object> = GLOBAL;

  private readonly H: Array<any>;
  private readonly ttl: Map<heap.Ref<any>, number>;
  private next: heap.Ref<any>;

  public constructor() {
    this.H = new Array(_static_heap_size_);
    this.H[UNDEFINED] = undefined;
    this.H[NULL] = null;
    this.H[ZERO] = 0;
    this.H[FALSE] = false;
    this.H[TRUE] = true;
    this.H[Const_NaN] = NaN;
    this.H[INF] = Infinity;
    this.H[NEG_INF] = -Infinity;
    this.H[GLOBAL] = globalThis;
    for (let i = 0; i < SMALL_INT_CACHE_SIZE; i++) {
      this.H[_LAST_CONSTANT_REF_ + 1 + i] = i + 1;
    }
    this.next = this.H.length;
    this.ttl = new Map<heap.Ref<any>, number>();
  }

  public get<T>(ref: heap.Ref<T>): T {
    if (!this.ttl.has(ref)) {
      return this.H[ref];
    }

    const ttl = this.ttl.get(ref) - 1;
    if (ttl <= 0) {
      return this.take(ref);
    }

    this.ttl.set(ref, ttl);
    return this.H[ref];
  }

  public freeAfterNGets(ref: heap.Ref<any>, n: number): void {
    if (ref < _static_heap_size_) {
      return;
    }

    if (this.ttl.has(ref)) {
      throw new Error("invalid heap operation");
    }

    if (n <= 0) {
      this.free(ref);
      return;
    }

    this.ttl.set(ref, n);
  }

  public free<T>(ref: heap.Ref<T>): void {
    if (ref < _static_heap_size_) {
      return;
    }

    this.ttl.delete(ref);
    this.H[ref] = this.next;
    this.next = ref;
  }

  // take is the combination of get() and free().
  public take<T>(ref: heap.Ref<T>): T {
    const x = this.H[ref];
    this.free(ref);
    return x;
  }

  replace<T>(idx: heap.Ref<T>, val: T): heap.Ref<boolean> {
    if (idx < _static_heap_size_) {
      return FALSE;
    }

    this.H[idx] = val;
    return TRUE;
  }

  public push(obj: any): heap.Ref<any> {
    if (obj === undefined) return UNDEFINED;
    if (obj === null) return NULL;

    switch (typeof obj) {
      case "boolean":
        return obj ? TRUE : FALSE;
      case "number":
        if (Number.isInteger(obj) && obj >= 0 && obj < SMALL_INT_CACHE_SIZE) {
          return obj + _LAST_CONSTANT_REF_ + 1;
        }

        if (Number.isNaN(obj)) return Const_NaN;
        if (!Number.isFinite(obj)) return obj > 0 ? INF : NEG_INF;
        break;
    }

    if (this.next === this.H.length) this.H.push(this.H.length + 1);
    const idx = this.next;
    this.next = this.H[idx];

    if (typeof this.next !== "number") throw new Error("corrupt heap");

    this.H[idx] = obj;
    return idx;
  }

  reset() {
    this.H.splice(_static_heap_size_);
    this.ttl.clear();
    this.next = this.H.length;
  }
}

interface UTF8Encoder {
  /**
   * Runs the UTF-8 encoder on source, stores the result of that operation into destination, and returns the progress made as an object wherein read is the number of converted code units of source and written is the number of bytes modified in destination.
   *
   * [MDN Reference](https://developer.mozilla.org/docs/Web/API/TextEncoder/encodeInto)
   */
  encodeInto(source: string, destination: Uint8Array): TextEncoderEncodeIntoResult;
}

interface UTF8Decoder {
  /**
   * Returns the result of running encoding's decoder. The method can be invoked zero or more times with options's stream set to true, and then once without options's stream (or set to false), to process a fragmented input. If the invocation without options's stream (or set to false) has no input, it's clearest to omit both arguments.
   *
   * ```
   * var string = "", decoder = new TextDecoder(encoding), buffer;
   * while(buffer = next_chunk()) {
   *   string += decoder.decode(buffer, {stream:true});
   * }
   * string += decoder.decode(); // end-of-queue
   * ```
   *
   * If the error mode is "fatal" and encoding's decoder returns error, throws a TypeError.
   *
   * [MDN Reference](https://developer.mozilla.org/docs/Web/API/TextDecoder/decode)
   */
  decode(input?: BufferSource, options?: TextDecodeOptions): string;
}

export interface AppOptions {
  utf8enc?: UTF8Encoder;
  utf8dec?: UTF8Decoder;
}

export class Application {
  // init instantiates the wasm asynchronously.
  //
  // @param src is a Promise<Response>, normally returned by a fetch() call.
  public async init(src: Promise<Response>): Promise<WebAssembly.WebAssemblyInstantiatedSource> {
    let fn = WebAssembly.instantiateStreaming;
    if (!fn) {
      fn = async (resp: Response | PromiseLike<Response>, importObject: any) => {
        const source = await (await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject);
      };
    }

    return await fn(src, this.importObject);
  }

  // initSync instantiates the wasm blob synchronously.
  //
  // @param src is the wasm blob, whose type should be one of:
  //   - BufferSource
  //   - base64 encoded string
  //   - data url string (with content base64 encoded)
  public initSync(src: BufferSource | string): WebAssembly.WebAssemblyInstantiatedSource {
    if (typeof src === "string") {
      if (src.startsWith("data:")) {
        // data url format is like `data:application/wasm,<base64 encoded string>`
        src = src.substring(src.indexOf(","));
      }

      src = atob(src);
      const buf = new Uint8Array(src.length);
      for (let i = 0; i < buf.length; i++) {
        buf[i] = src.charCodeAt(i);
      }

      src = buf.buffer;
    }

    const mod = new WebAssembly.Module(src);
    return {
      instance: new WebAssembly.Instance(mod, this.importObject),
      module: mod,
    };
  }

  private _runid: number = 0;
  private _exited: boolean = true;
  private _exitcode: number;

  public get runid(): number {
    return this._runid;
  }

  public get exited(): boolean {
    return this._exited;
  }

  private _exitPromise: Promise<number>;
  private _exitResolveFn: (value: number) => void;
  public resolveExitPromise() {
    this._exitResolveFn(this._exitcode);
  }

  private _wasm: {
    // call is wasm_export_run in go assembly (special link symbol)
    //
    // In the standard go runtime, it is the function to start the wasm program:
    //  wasm_export_run(argc i32, argv i32)
    //
    // but in pcz we have repurposed its usage to run callback functions.
    readonly call: (ctx: heap.Ref<any>, sp: number) => void;

    // run is wasm_export_resume in go assembly (special link symbol)
    //
    // In pcz, it is repurposed to be the program entrypoint.
    readonly run: () => void;

    // getsp is wasm_export_getsp in go assembly (special link symbol)
    //
    // (currently not used)
    readonly getsp: () => number;

    readonly inst: WebAssembly.Instance;
  };

  public call(ctx: heap.Ref<any>, sp: number): void {
    this._wasm.call(ctx, sp);
  }

  // public get mem(): WebAssembly.Memory {
  //   return this._wasm.inst.exports.mem as WebAssembly.Memory;
  // }

  public run(inst: WebAssembly.Instance): Promise<number> {
    if (!this._exited) {
      throw new Error("Concurrent run() is not supported, create multiple Applications instead.");
    }

    this._runid++;
    this._exited = false;
    this._exitPromise = new Promise((resolve) => {
      this._exitResolveFn = resolve;
    }).then((code: number): number => {
      this.H.reset();
      for (let i = 0; i < this.importedModules.length; i++) {
        this.importedModules[i][1]?.["_onexit"]?.(code);
      }

      return code;
    });

    this._wasm = {
      inst,
      run: inst.exports.resume as () => void,
      call: inst.exports.run as (ctx: heap.Ref<any>, sp: number) => void,
      getsp: inst.exports.getsp as () => number,
    };

    // NOTE: always reset the memory view before running any wasm func
    this.resetMemoryDataView();

    for (let i = 0; i < this.importedModules.length; i++) {
      this.importedModules[i][1]?.["_onstart"]?.();
    }

    this._wasm.run();
    if (this._exited) {
      this.resolveExitPromise();
    }

    return this._exitPromise;
  }

  private view: DataView;
  public resetMemoryDataView(): void {
    this.view = new DataView((this._wasm.inst.exports.mem as WebAssembly.Memory).buffer);
  }

  private readonly buf_utf8size = new Uint8Array(32);
  public UTF8Sizeof(s: string): number {
    let sum: number = 0;
    while (s.length > 0) {
      const progress = this.UTF8enc.encodeInto(s, this.buf_utf8size);
      if ((progress.read ?? 0) === 0 || (progress.written ?? 0) === 0) {
        break;
      }

      sum += progress.written ?? 0;
      s = s.substring(progress.read ?? 0);
    }
    return sum;
  }

  public readonly UTF8enc: UTF8Encoder;
  public readonly UTF8dec: UTF8Decoder;
  public readonly H: Heap;

  public readonly load = {
    Raw: (pByte: Pointer, len: number): Uint8Array => {
      return new Uint8Array(this.view.buffer).subarray(pByte, pByte + len);
    },
    Bool: (pBool: Pointer): boolean => {
      return this.view.getUint8(pBool >>> 0) !== 0;
    },
    Uint8: (pU8: Pointer): number => {
      return this.view.getUint8(pU8 >>> 0);
    },
    Uint16: (pU16: Pointer): number => {
      return this.view.getUint16(pU16 >>> 0, true);
    },
    Uint32: (pU32: Pointer): number => {
      return this.view.getUint32(pU32 >>> 0, true);
    },
    Uint64: (pU64: Pointer): number => {
      return Number(this.view.getBigUint64(pU64 >>> 0, true));
    },
    BigUint64: (pU64: Pointer): bigint => {
      return this.view.getBigUint64(pU64 >>> 0, true);
    },
    Int8: (pI8: Pointer): number => {
      return this.view.getInt8(pI8 >>> 0);
    },
    Int16: (pI16: Pointer): number => {
      return this.view.getInt16(pI16 >>> 0, true);
    },
    Int32: (pI32: Pointer): number => {
      return this.view.getInt32(pI32 >>> 0, true);
    },
    Int64: (pI64: Pointer): number => {
      return Number(this.view.getBigInt64(pI64 >>> 0, true));
    },
    BigInt64: (pI64: Pointer): bigint => {
      return this.view.getBigInt64(pI64 >>> 0, true);
    },
    Float32: (pF32: Pointer): number => {
      return this.view.getFloat32(pF32 >>> 0, true);
    },
    Float64: (pF64: Pointer): number => {
      return this.view.getFloat64(pF64 >>> 0, true);
    },
    String: (pByte: Pointer, len: number): string => {
      pByte >>>= 0;
      return this.UTF8dec.decode(this.view.buffer.slice(pByte, pByte + len));
    },

    Ref: <T = any>(pU32: Pointer, def?: T, take?: boolean): T => {
      const ref = this.view.getUint32(pU32 >>> 0, true);
      if (ref === UNDEFINED) return def;

      return take ? this.H.take<T>(ref) : this.H.get<T>(ref);
    },
    Enum: (pU32: Pointer, enums: string[]): string | undefined => {
      const x = this.view.getUint32(pU32 >>> 0, true);
      if (x <= 0 || x > enums.length) {
        return undefined;
      }

      return enums[x - 1];
    },

    Refs: (pU32: Pointer, n: number, free?: boolean): any[] => {
      if (n === 0) {
        return [];
      }

      pU32 >>>= 0;
      const args: any[] = new Array(n);
      if (free) {
        for (let i = 0; i < n; i++) {
          args[i] = this.H.take(this.view.getUint32(pU32 + i * 4, true));
        }
      } else {
        for (let i = 0; i < n; i++) {
          args[i] = this.H.get(this.view.getUint32(pU32 + i * 4, true));
        }
      }
      return args;
    },
  };

  public readonly store = {
    Raw: (pByte: Pointer, data: Uint8Array): void => {
      pByte >>>= 0;
      new Uint8Array(this.view.buffer).subarray(pByte, pByte + data.length).set(data);
    },
    String: (pByteBuf: number, len: number, val: string): { n: number; remainder: string } => {
      pByteBuf >>>= 0;
      let sum = 0;
      do {
        const r = this.UTF8enc.encodeInto(
          val,
          // NOTE: do not use ArrayBuffer.slice (this.mem.buffer.slice)
          // the returned slice doesn't represent the underlay memory.
          new Uint8Array(this.view.buffer).subarray(pByteBuf, pByteBuf + len)
        );
        if ((r.read ?? 0) === 0 || (r.written ?? 0) === 0) {
          break;
        }
        val = val.substring(r.read ?? 0);
        sum += r.written ?? 0;
        pByteBuf += r.written ?? 0;
        len -= r.written ?? 0;
      } while (val.length > 0 && len > 0);

      return { n: sum, remainder: val };
    },
    Bool: (pBool: Pointer, val: boolean): void => {
      this.view.setUint8(pBool >>> 0, val ? 1 : 0);
    },
    Uint8: (pU8: Pointer, val: number): void => {
      this.view.setUint8(pU8 >>> 0, val);
    },
    Uint16: (pU16: Pointer, val: number): void => {
      this.view.setUint16(pU16 >>> 0, val, true);
    },
    Uint32: (pU32: Pointer, val: number): void => {
      this.view.setUint32(pU32 >>> 0, val, true);
    },
    Uint64: (pU64: Pointer, val: number | bigint): void => {
      if (typeof val !== "bigint") val = BigInt(val);
      this.view.setBigUint64(pU64 >>> 0, val, true);
    },
    Int8: (pI8: Pointer, val: number): void => {
      this.view.setInt8(pI8 >>> 0, val);
    },
    Int16: (pI16: Pointer, val: number): void => {
      this.view.setInt16(pI16 >>> 0, val, true);
    },
    Int32: (pI32: Pointer, val: number): void => {
      this.view.setInt32(pI32 >>> 0, val, true);
    },
    Int64: (pI64: Pointer, val: number | bigint): void => {
      if (typeof val !== "bigint") val = BigInt(val);
      this.view.setBigInt64(pI64 >>> 0, val, true);
    },
    Float32: (pF32: Pointer, val: number): void => {
      this.view.setFloat32(pF32 >>> 0, val, true);
    },
    Float64: (pF64: Pointer, val: number): void => {
      this.view.setFloat64(pF64 >>> 0, val, true);
    },

    Ref: <T = any>(pU32: Pointer, val: T): void => {
      this.view.setUint32(pU32 >>> 0, this.H.push(val), true);
    },
    Enum: (pU32: Pointer, index: number): void => {
      if (index < 0) {
        this.view.setUint32(pU32 >>> 0, 0, true);
      } else {
        this.view.setUint32(pU32 >>> 0, index + 1, true);
      }
    },
  };

  private readonly importedModules: [string, Record<string, Function>][];
  private readonly importObject: Record<string, Record<string, Function>> = {
    "ffi/js.core": {
      exit: (code: number): void => {
        this._exitcode = code;
        this._exited = true;
      },
    },
  };

  public constructor(opts?: AppOptions) {
    if (opts?.utf8enc) {
      this.UTF8enc = opts!.utf8enc;
    } else {
      if (!globalThis.TextEncoder) {
        throw new Error("TextEncoder not available, please construct with custom UTF8Encoder");
      }

      this.UTF8enc = new globalThis.TextEncoder();
    }

    if (opts?.utf8dec) {
      this.UTF8dec = opts!.utf8dec;
    } else {
      if (!globalThis.TextDecoder) {
        throw new Error("TextDecoder not available, please construct with custom UTF8Decoder");
      }

      this.UTF8dec = new globalThis.TextDecoder("utf-8");
    }

    this.H = new Heap();

    // constants below are used by inserted module imports (see ./@ffi.ts).
    const importModule: ModuleImporter = (module: string, factory: ModuleFactory): void => {
      if (this.importObject[module]) {
        // TODO: allow unique function imports in the same namespace.
        throw new Error(`duplicate module name ${module}`);
      }

      this.importObject[module] = factory(this);
    };

    //__TEMPLATE_INSERT_IMPORTS__

    this.importedModules = Object.entries(this.importObject);
  }
}
