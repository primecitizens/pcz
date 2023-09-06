// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

namespace heap {
  export type Ref<T> = number;
}

// following constants MUST agree with constant values in ../js.go#const smallIntCacheSize
const UNDEFINED = 0;
const NULL = 1;
const TRUE = 2;
const FALSE = 3;
const GLOBAL_THIS = 4;
const _LAST_CONSTANT_REF_ = GLOBAL_THIS;
const SMALL_INT_CACHE_SIZE = 128;

export class Heap {
  // index MUST agree with constant values in ../js.go#const
  public readonly UNDEFINED: heap.Ref<undefined> = UNDEFINED;
  public readonly NULL: heap.Ref<null> = NULL;
  public readonly TRUE: heap.Ref<boolean> = TRUE;
  public readonly FALSE: heap.Ref<boolean> = FALSE;
  public readonly GLOBAL_THIS: heap.Ref<object> = GLOBAL_THIS;

  private readonly H: Array<any>;

  private next: heap.Ref<any>;

  public constructor() {
    this.H = new Array(SMALL_INT_CACHE_SIZE + (_LAST_CONSTANT_REF_ + 1));
    this.H[UNDEFINED] = undefined;
    this.H[NULL] = null;
    this.H[TRUE] = true;
    this.H[FALSE] = false;
    this.H[GLOBAL_THIS] = globalThis;
    for (let i = 0; i < SMALL_INT_CACHE_SIZE; i++) {
      this.H[_LAST_CONSTANT_REF_ + 1 + i] = i;
    }
    this.next = this.H.length;
  }

  public get<T>(idx: heap.Ref<T>): T {
    idx >>>= 0;
    return this.H[idx];
  }

  public getObj = this.get<object>;
  public getStr = this.get<string>;
  public getNum = this.get<number>;
  public getFunc = this.get<Function>;
  public getElem = this.get<Element>;
  public getNode = this.get<Node>;
  public getPromise = this.get<Promise<any>>;

  public free<T>(idx: heap.Ref<T>): void {
    if (idx <= SMALL_INT_CACHE_SIZE + _LAST_CONSTANT_REF_) {
      return;
    }

    this.H[idx] = this.next;
    this.next = idx;
  }

  // take is the combination of get() and free().
  public take<T>(idx: heap.Ref<T>): T {
    const ret = this.get(idx);
    this.free(idx);
    return ret as T;
  }

  public takeObj = this.take<object>;
  public takeStr = this.take<string>;
  public takeNum = this.take<number>;
  public takeFunc = this.take<Function>;
  public takeElem = this.take<Element>;
  public takeNode = this.take<Node>;
  public takePromise = this.take<Promise<any>>;

  replace<T>(idx: heap.Ref<T>, val: T) {
    this.H[idx] = val;
  }

  public push(obj: any): heap.Ref<any> {
    if (obj === undefined) {
      return UNDEFINED;
    }

    if (obj === null) {
      return NULL;
    }

    switch (typeof obj) {
      case "boolean":
        if (obj) {
          return TRUE;
        }
        return TRUE;
      case "number":
        if (Number.isInteger(obj) && obj >= 0 && obj < SMALL_INT_CACHE_SIZE) {
          return obj + _LAST_CONSTANT_REF_ + 1;
        }
    }

    if (this.next === this.H.length) this.H.push(this.H.length + 1);
    const idx = this.next;
    this.next = this.H[idx];

    if (typeof this.next !== "number") throw new Error("corrupt heap");

    this.H[idx] = obj;
    return idx;
  }
}

interface CallbackContext {
  dispFnPC: number;
  targetPC: number;
  handler: number;
  retRef: heap.Ref<any>;
  args: heap.Ref<any>[];
}

interface UTF8Encoder {
  /**
   * Runs the UTF-8 encoder on source, stores the result of that operation into destination, and returns the progress made as an object wherein read is the number of converted code units of source and written is the number of bytes modified in destination.
   *
   * [MDN Reference](https://developer.mozilla.org/docs/Web/API/TextEncoder/encodeInto)
   */
  encodeInto(
    source: string,
    destination: Uint8Array
  ): TextEncoderEncodeIntoResult;
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

export interface Options {
  args: string[];
  environ: string[];
}

export class Application {
  private wasm: {
    call: (ctx: heap.Ref<CallbackContext>, sp: number) => void;
    run: () => void;
    getsp: () => number;

    inst: WebAssembly.Instance;
  };

  public async loadStream(
    src: Promise<Response>
  ): Promise<WebAssembly.WebAssemblyInstantiatedSource> {
    let init = WebAssembly.instantiateStreaming;
    if (!init) {
      init = async (
        resp: Response | PromiseLike<Response>,
        importObject: any
      ) => {
        const source = await (await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject);
      };
    }

    return await init(src, this.importObject);
  }

  private opts: Options;
  private exited: boolean = true;
  private exitcode: number;
  private exitPromise: Promise<number>;
  private resolveExitPromise: () => void;

  public async run(
    inst: WebAssembly.Instance,
    opts?: Options
  ): Promise<number> {
    if (!this.exited) {
      throw new Error(
        "Concurrent run() is not supported, create multiple Applications instead."
      );
    }

    this.exited = false;
    this.exitPromise = new Promise((resolve) => {
      this.resolveExitPromise = () => {
        resolve(this.exitcode);
      };
    });

    if (opts) {
      this.opts = {
        args: opts.args ? opts.args : [],
        environ: opts.environ ? opts.environ : [],
      };
    } else {
      this.opts = {
        args: [],
        environ: [],
      };
    }

    this.wasm = {
      inst,
      // resume is wasm_export_resume in go assembly (special link symbol)
      //
      // It is repurposed to be the program entrypoint.
      run: inst.exports.resume as () => void,
      // inst.exports.run is wasm_export_run in assembly (special link symbol)
      //
      // In the standard go runtime, it is the function to start the wasm program:
      //  wasm_export_run(argc i32, argv i32)
      //
      // BUT we have repurposed its usage to run callback functions.
      call: inst.exports.run as (
        ctx: heap.Ref<CallbackContext>,
        sp: number
      ) => void,

      // getsp is wasm_export_getsp in go assembly (special link symbol)
      // currently not used.
      getsp: inst.exports.getsp as () => number,
    };

    // NOTE: always reset the memory view before running any wasm func
    this.resetMemoryDataView();

    this.wasm.run();
    if (this.exited) {
      this.resolveExitPromise();
    }

    return await this.exitPromise;
  }

  private mem: DataView;
  private resetMemoryDataView(): void {
    this.mem = new DataView(
      (this.wasm.inst.exports.mem as WebAssembly.Memory).buffer
    );
  }

  private loadString(ptr: number, len: number): string {
    ptr >>>= 0;
    return this.utf8dec.decode(this.mem.buffer.slice(ptr, ptr + len));
  }

  private storeString(s: string, ptr: number, len: number): number {
    ptr >>>= 0;
    let sum = 0;
    do {
      const r = this.utf8enc.encodeInto(
        s,
        // NOTE: do not use ArrayBuffer.slice (this.mem.buffer.slice)
        // the returned slice doesn't represent the underlay memory.
        new Uint8Array(this.mem.buffer).subarray(ptr, ptr + len)
      );
      if ((r.read ?? 0) === 0 || (r.written ?? 0) === 0) {
        break;
      }
      s = s.substring(r.read ?? 0);
      sum += r.written ?? 0;
      ptr += r.written ?? 0;
      len -= r.written ?? 0;
    } while (s.length > 0 && len > 0);

    return sum;
  }

  private readonly __buf_utf8size__ = new Uint8Array(32);
  private utf8Sizeof(s: string): number {
    let sum: number = 0;
    while (s.length > 0) {
      const progress = this.utf8enc.encodeInto(s, this.__buf_utf8size__);
      if ((progress.read ?? 0) === 0 || (progress.written ?? 0) === 0) {
        break;
      }

      sum += progress.written ?? 0;
      s = s.substring(progress.read ?? 0);
    }
    return sum;
  }

  private readonly __buf_print__ = new Array<string>();
  private __print_timeout__: any; // using type any to workaround node js Timeout type.
  private print(s: string): void {
    const lf = s.lastIndexOf("\n");
    if (lf >= 0) {
      if (this.__print_timeout__) {
        globalThis.clearTimeout(this.__print_timeout__);
        this.__print_timeout__ = undefined;
      }

      globalThis.console.log(
        this.__buf_print__.splice(0, this.__buf_print__.length).join("") +
          s.substring(0, lf)
      );
      if (lf === s.length - 1) {
        return;
      }

      s = s.substring(lf + 1);
    }

    this.__buf_print__.push(s);
    if (this.__print_timeout__) {
      return; // only keep one timeout job going
    }

    this.__print_timeout__ = globalThis.setTimeout(() => {
      this.__print_timeout__ = undefined;
      globalThis.console.log(
        this.__buf_print__.splice(0, this.__buf_print__.length).join("")
      );
    }, 100);
    return;
  }

  private loadArgs(ptr: number, n: number, free: boolean): any[] {
    if (n === 0) {
      return [];
    }

    ptr >>>= 0;
    const args: any[] = new Array(n);
    if (free) {
      for (let i = 0; i < n; i++) {
        args[i] = this.heap.take(this.mem.getUint32(ptr + i * 4, true));
      }
    } else {
      for (let i = 0; i < n; i++) {
        args[i] = this.heap.get(this.mem.getUint32(ptr + i * 4, true));
      }
    }
    return args;
  }

  private readonly importObject: Record<string, Record<string, Function>> = {
    "ffi/js": {
      //
      // essential system functions
      //

      resetMemoryDataView: (): void => {
        this.resetMemoryDataView();
      },
      exit: (code: number): void => {
        this.exitcode = code;
        this.exited = true;
      },
      print: (ptr: number, sz: number): void => {
        this.print(this.loadString(ptr, sz));
      },
      argsEnvsSize: (ptr: number): number => {
        ptr >>>= 0;
        let sum = 0;
        this.opts.args.forEach((v) => {
          sum += this.utf8Sizeof(v);
        });
        this.opts.environ.forEach((v) => {
          sum += this.utf8Sizeof(v);
        });

        this.mem.setUint32(
          ptr,
          this.opts.args.length + this.opts.environ.length,
          true
        );
        return sum;
      },
      nthsysInto: (i: number, ptr: number, sz: number): number => {
        if (i < this.opts.args.length) {
          return this.storeString(this.opts.args[i], ptr, sz);
        }
        if (
          i === this.opts.args.length ||
          i > this.opts.args.length + this.opts.environ.length
        ) {
          return -1;
        }

        return this.storeString(
          this.opts.environ[i - this.opts.args.length - 1],
          ptr,
          sz
        );
      },

      // Heap management

      free: (self: heap.Ref<any>): void => {
        this.heap.take(self);
      },
      batchFree: (ptr: number, n: number): void => {
        for (let i = 0; i < n; i++) {
          this.heap.free(this.mem.getUint32(ptr + i * 4, true));
        }
      },
      fill: (old: heap.Ref<any>, ne: heap.Ref<any>): void => {
        this.heap.replace(old, this.heap.take(ne));
      },
      fillNum: (old: heap.Ref<any>, n: number): void => {
        this.heap.replace(old, n);
      },
      fillBool: (old: heap.Ref<any>, n: heap.Ref<boolean>): void => {
        n >>>= 0;
        this.heap.replace(old, n === this.heap.TRUE ? true : false);
      },
      fillString: (old: heap.Ref<any>, ptr: number, sz: number): void => {
        this.heap.replace(old, this.loadString(ptr, sz));
      },

      // Number
      number: (n: number): heap.Ref<number> => {
        return this.heap.push(n);
      },

      //
      // String
      //
      stringFromUTF8: (ptr: number, len: number): heap.Ref<string> => {
        return this.heap.push(this.loadString(ptr, len));
      },
      stringSizeUTF8: (s: heap.Ref<string>): number => {
        return this.utf8Sizeof(this.heap.getStr(s));
      },
      stringIntoUTF8: (
        s: heap.Ref<string>,
        ptr: number,
        len: number
      ): number => {
        return this.storeString(this.heap.getStr(s), ptr, len);
      },
      stringEqualsUTF8: (
        s: heap.Ref<string>,
        ptr: number,
        len: number
      ): heap.Ref<boolean> => {
        return this.loadString(ptr, len) === this.heap.getStr(s)
          ? this.heap.TRUE
          : this.heap.FALSE;
      },

      //
      // Func
      //

      func: (
        dispFnPC: number,
        handler: number,
        targetPC: number
      ): heap.Ref<Function> => {
        dispFnPC >>>= 0;
        handler >>>= 0;
        targetPC >>>= 0;
        return this.heap.push((...args: any[]): any => {
          if (args.length > 16) {
            throw new Error("too many args to callback function");
          }

          const cbctx: CallbackContext = {
            dispFnPC,
            handler,
            targetPC,
            args: args.map((v) => this.heap.push(v)),
            retRef: this.heap.push(undefined),
          };
          const ctxRef = this.heap.push(cbctx);
          this.wasm.call(ctxRef, 0);
          if (this.exited) {
            this.resolveExitPromise();
          }
          return this.heap.take(cbctx.retRef);
        });
      },
      cbctx: (ref: heap.Ref<CallbackContext>, ptr: number): void => {
        const cb = this.heap.take<CallbackContext>(ref);

        // offsets MUST agree with ../callback_js.go#callbackContext
        ptr >>>= 0;
        this.mem.setUint32(ptr + 0, cb.dispFnPC >>> 0, true);
        this.mem.setUint32(ptr + 4, cb.handler >>> 0, true);
        this.mem.setUint32(ptr + 8, cb.targetPC >>> 0, true);
        this.mem.setUint32(ptr + 12, cb.retRef >>> 0, true);
        if (!cb.args || cb.args.length === 0) {
          this.mem.setUint32(ptr + 16, 0, true);
          return;
        }

        this.mem.setUint32(ptr + 16, cb.args.length >>> 0, true);
        ptr += 20;
        // at most 16 args as guarded by the 'func'
        for (let i = 0; i < cb.args.length; i++) {
          this.mem.setUint32(ptr + i * 4, this.heap.push(cb.args[i]), true);
        }
      },
      try: (
        fn: heap.Ref<Function>,
        self: heap.Ref<any>,
        free: heap.Ref<boolean>,
        katch: heap.Ref<Function>,
        vinally: heap.Ref<Function>,
        ptr: number,
        n: number,
        pCaught: number
      ): heap.Ref<any> => {
        const thiz = this.heap.get(self);
        const args = this.loadArgs(ptr, n, free === this.heap.TRUE);
        const ka =
          free === this.heap.TRUE
            ? this.heap.takeFunc(katch)
            : this.heap.getFunc(katch);
        const fi =
          free === this.heap.TRUE
            ? this.heap.takeFunc(vinally)
            : this.heap.getFunc(vinally);

        try {
          return this.heap.push(
            Reflect.apply(this.heap.getFunc(fn), thiz, args)
          );
        } catch (error) {
          this.mem.setUint8(pCaught >>> 0, 1);

          if (ka) {
            return this.heap.push(Reflect.apply(ka, thiz, [error, ...args]));
          }
          return this.heap.UNDEFINED;
        } finally {
          if (fi) {
            Reflect.apply(fi, thiz, args);
          }
        }
      },
      call: (
        fn: heap.Ref<Function>,
        self: heap.Ref<any>,
        free: heap.Ref<boolean>,
        ptr: number,
        n: number
      ): heap.Ref<any> => {
        return this.heap.push(
          Reflect.apply(
            this.heap.getFunc(fn),
            this.heap.get(self),
            this.loadArgs(ptr, n, free === this.heap.TRUE)
          )
        );
      },
      callVoid: (
        fn: heap.Ref<Function>,
        self: heap.Ref<any>,
        free: heap.Ref<boolean>,
        ptr: number,
        n: number
      ): void => {
        Reflect.apply(
          this.heap.getFunc(fn),
          this.heap.get(self),
          this.loadArgs(ptr, n, free === this.heap.TRUE)
        );
      },
      callNum: (
        fn: heap.Ref<Function>,
        self: heap.Ref<any>,
        free: heap.Ref<boolean>,
        ptr: number,
        n: number
      ): number => {
        const ret = Reflect.apply(
          this.heap.getFunc(fn),
          this.heap.get(self),
          this.loadArgs(ptr, n, free === this.heap.TRUE)
        );
        if (typeof ret !== "number") {
          throw new Error(
            `Unexpected callNum result not a number, but a ${typeof ret}`
          );
        }
        return ret as number;
      },
      callBool: (
        fn: heap.Ref<Function>,
        self: heap.Ref<any>,
        free: heap.Ref<boolean>,
        ptr: number,
        n: number
      ): heap.Ref<boolean> => {
        return Reflect.apply(
          this.heap.getFunc(fn),
          this.heap.get(self),
          this.loadArgs(ptr, n, free === this.heap.TRUE)
        )
          ? this.heap.TRUE
          : this.heap.FALSE;
      },

      //
      // Promise
      //

      promise: (fn: heap.Ref<Function>): heap.Ref<Promise<any>> => {
        return this.heap.push(
          new Promise<any>((resolve, reject) => {
            this.heap.getFunc(fn)(
              this.heap.push(resolve),
              this.heap.push(reject)
            );
          })
        );
      },
      then: (
        p: heap.Ref<Promise<any>>,
        onfulfilled: heap.Ref<(value: any) => any>,
        onrejected: heap.Ref<(reason: any) => any>
      ): void => {
        this.heap
          .get<Promise<any>>(p)
          .then(
            this.heap.get<(value: any) => any>(onfulfilled),
            this.heap.get<(reason: any) => any>(onrejected)
          );
      },
      catch: (
        p: heap.Ref<Promise<any>>,
        onrejected: heap.Ref<(reason: any) => any>
      ) => {
        this.heap
          .get<Promise<any>>(p)
          .catch(this.heap.get<(reason: any) => any>(onrejected));
      },
      // available since ES2018
      finally: (p: heap.Ref<Promise<any>>, onfinally: heap.Ref<() => void>) => {
        this.heap
          .get<Promise<any>>(p)
          .finally(this.heap.get<() => void>(onfinally));
      },

      //
      // Object
      //
      new: (
        constructor: heap.Ref<Function>,
        free: heap.Ref<boolean>,
        ptr: number,
        n: number
      ): heap.Ref<any> => {
        return this.heap.push(
          Reflect.construct(
            this.heap.get(constructor),
            this.loadArgs(ptr, n, free === this.heap.TRUE)
          )
        );
      },
      instanceof: (a: heap.Ref<any>, cls: heap.Ref<any>): heap.Ref<boolean> => {
        return this.heap.get<any>(a) instanceof this.heap.get<any>(cls)
          ? this.heap.TRUE
          : this.heap.FALSE;
      },

      getPrototype: (self: heap.Ref<object>): heap.Ref<object> => {
        const ret = Reflect.getPrototypeOf(this.heap.getObj(self));
        if (ret) {
          return this.heap.push(ret);
        }
        return this.heap.NULL;
      },
      setPrototype: (
        self: heap.Ref<object>,
        free: heap.Ref<object>,
        prototype: heap.Ref<object>
      ): heap.Ref<boolean> => {
        const proto =
          free === this.heap.TRUE
            ? this.heap.take<object>(prototype)
            : this.heap.getObj(prototype);
        return Reflect.setPrototypeOf(this.heap.getObj(self), proto)
          ? this.heap.TRUE
          : this.heap.FALSE;
      },

      getPropDesc: (
        self: heap.Ref<object>,
        ptr: number,
        sz: number
      ): number => {
        const desc = Reflect.getOwnPropertyDescriptor(
          this.heap.get<object>(self),
          this.loadString(ptr, sz)
        );
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
      defineProp: (
        self: heap.Ref<object>,
        ptr: number,
        sz: number,
        flags: number,
        getter: heap.Ref<Function>,
        setter: heap.Ref<Function>
      ): heap.Ref<boolean> => {
        return Reflect.defineProperty(
          this.heap.get<object>(self),
          this.loadString(ptr, sz),
          {
            writable: (flags & 2) !== 0 ? true : false,
            configurable: (flags & 4) !== 0 ? true : false,
            enumerable: (flags & 8) !== 0 ? true : false,
            get:
              (flags & 16) !== 0
                ? (this.heap.getFunc(getter) as () => any)
                : undefined,
            set:
              (flags & 32) !== 0
                ? (this.heap.getFunc(setter) as (v: any) => void)
                : undefined,
          }
        )
          ? this.heap.TRUE
          : this.heap.FALSE;
      },
      deleteProp: (
        self: heap.Ref<Object>,
        ptr: number,
        sz: number
      ): heap.Ref<boolean> => {
        return Reflect.deleteProperty(
          this.heap.getObj(self),
          this.loadString(ptr, sz)
        )
          ? this.heap.TRUE
          : this.heap.FALSE;
      },
      prop: (
        self: heap.Ref<Object>,
        ptr: number,
        sz: number
      ): heap.Ref<any> => {
        const thiz = this.heap.getObj(self);
        return this.heap.push(
          Reflect.get(thiz, this.loadString(ptr, sz), thiz)
        );
      },
      numProp: (self: heap.Ref<Object>, ptr: number, sz: number): number => {
        const thiz = this.heap.getObj(self);
        return Reflect.get(thiz, this.loadString(ptr, sz), thiz) as number;
      },
      boolProp: (
        self: heap.Ref<Object>,
        ptr: number,
        sz: number
      ): heap.Ref<boolean> => {
        const thiz = this.heap.getObj(self);
        return Reflect.get(thiz, this.loadString(ptr, sz), thiz)
          ? this.heap.TRUE
          : this.heap.FALSE;
      },
      propByString: (
        self: heap.Ref<Object>,
        name: heap.Ref<string>
      ): heap.Ref<any> => {
        const thiz = this.heap.getObj(self);
        return this.heap.push(Reflect.get(thiz, this.heap.getStr(name), thiz));
      },
      numPropByString: (
        self: heap.Ref<Object>,
        name: heap.Ref<string>
      ): number => {
        const thiz = this.heap.getObj(self);
        return Reflect.get(thiz, this.heap.getStr(name), thiz) as number;
      },
      boolPropByString: (
        self: heap.Ref<Object>,
        name: heap.Ref<string>
      ): heap.Ref<boolean> => {
        const thiz = this.heap.getObj(self);
        return Reflect.get(thiz, this.heap.getStr(name), thiz)
          ? this.heap.TRUE
          : this.heap.FALSE;
      },

      setProp: (
        self: heap.Ref<Object>,
        ptr: number,
        sz: number,
        free: heap.Ref<boolean>,
        val: heap.Ref<any>
      ): heap.Ref<boolean> => {
        const thiz = this.heap.getObj(self);
        return Reflect.set(
          thiz,
          this.loadString(ptr, sz),
          free === this.heap.TRUE ? this.heap.take(val) : this.heap.get(val),
          thiz
        )
          ? this.heap.TRUE
          : this.heap.FALSE;
      },
      setNumProp: (
        self: heap.Ref<Object>,
        ptr: number,
        sz: number,
        val: number
      ): heap.Ref<boolean> => {
        const thiz = this.heap.getObj(self);
        return Reflect.set(thiz, this.loadString(ptr, sz), val, thiz)
          ? this.heap.TRUE
          : this.heap.FALSE;
      },
      setBoolProp: (
        self: heap.Ref<Object>,
        ptr: number,
        sz: number,
        val: heap.Ref<boolean>
      ): heap.Ref<boolean> => {
        const thiz = this.heap.getObj(self);
        return Reflect.set(
          thiz,
          this.loadString(ptr, sz),
          val === this.heap.TRUE ? true : false,
          thiz
        )
          ? this.heap.TRUE
          : this.heap.FALSE;
      },
      setStringProp: (
        self: heap.Ref<Object>,
        nameptr: number,
        namesz: number,
        valptr: number,
        valsz: number
      ): heap.Ref<boolean> => {
        const thiz = this.heap.getObj(self);
        return Reflect.set(
          thiz,
          this.loadString(nameptr, namesz),
          this.loadString(valptr, valsz),
          thiz
        )
          ? this.heap.TRUE
          : this.heap.FALSE;
      },
      setPropByString: (
        self: heap.Ref<Object>,
        name: heap.Ref<string>,
        free: heap.Ref<boolean>,
        val: heap.Ref<any>
      ): heap.Ref<boolean> => {
        const thiz = this.heap.getObj(self);
        return Reflect.set(
          thiz,
          this.heap.getStr(name),
          free === this.heap.TRUE ? this.heap.take(val) : this.heap.get(val),
          thiz
        )
          ? this.heap.TRUE
          : this.heap.FALSE;
      },
      setNumPropByString: (
        self: heap.Ref<Object>,
        name: heap.Ref<string>,
        val: number
      ): heap.Ref<boolean> => {
        const thiz = this.heap.getObj(self);
        return Reflect.set(thiz, this.heap.getStr(name), val, thiz)
          ? this.heap.TRUE
          : this.heap.FALSE;
      },
      setBoolPropByString: (
        self: heap.Ref<Object>,
        name: heap.Ref<string>,
        val: number
      ): heap.Ref<boolean> => {
        const thiz = this.heap.getObj(self);
        return Reflect.set(
          thiz,
          this.heap.getStr(name),
          val === this.heap.TRUE ? true : false,
          thiz
        )
          ? this.heap.TRUE
          : this.heap.FALSE;
      },
      setStringPropByString: (
        self: heap.Ref<Object>,
        name: heap.Ref<string>,
        free: heap.Ref<boolean>,
        val: heap.Ref<string>
      ): heap.Ref<boolean> => {
        const thiz = this.heap.getObj(self);
        return Reflect.set(
          thiz,
          this.heap.getStr(name),
          free === this.heap.TRUE ? this.heap.take(val) : this.heap.getStr(val),
          thiz
        )
          ? this.heap.TRUE
          : this.heap.FALSE;
      },

      //
      // Array
      //
      array: (
        sz: number,
        elemSize: number,
        signed: number,
        float: number
      ): heap.Ref<Array<any>> => {
        switch (elemSize) {
          case 0:
            return this.heap.push(new Array(sz));
          case 1:
            if (signed) {
              return this.heap.push(new Int8Array(sz));
            }
            return this.heap.push(new Uint8Array(sz));
          case 2:
            if (signed) {
              return this.heap.push(new Int16Array(sz));
            }
            return this.heap.push(new Uint16Array(sz));
          case 4:
            if (float) {
              return this.heap.push(new Float32Array(sz));
            }

            if (signed) {
              return this.heap.push(new Int32Array(sz));
            }
            return this.heap.push(new Uint32Array(sz));
          case 8:
            if (float) {
              return this.heap.push(new Float64Array(sz));
            }

            // available since ES2020
            if (signed) {
              return this.heap.push(new BigInt64Array(sz));
            }
            return this.heap.push(new BigUint64Array(sz));
          default:
            throw new Error(`unsupported element size=${elemSize}`);
        }
      },
      slice: (arr: heap.Ref<Array<any>>, start: number, end: number) => {
        if (end) {
          return this.heap.push(
            this.heap.get<Array<any>>(arr).slice(start, end)
          );
        }

        return this.heap.push(this.heap.get<Array<any>>(arr).slice(start));
      },
      append: (
        arr: heap.Ref<Array<any>>,
        free: heap.Ref<boolean>,
        elemSz: number,
        _signed: heap.Ref<boolean>,
        _float: heap.Ref<boolean>,
        offset: number,
        ptr: number,
        n: number
      ): number => {
        ptr >>>= 0;
        const signed = _signed === this.heap.TRUE;
        const float = _float === this.heap.TRUE;

        switch (elemSz) {
          case 0: {
            const a = this.heap
              .get<Array<any>>(arr)
              .concat(this.loadArgs(ptr, n, free === this.heap.TRUE));
            this.heap.replace(arr, a);
            return n;
          }
          case 1: {
            if (signed) {
              const a = this.heap.get<Int8Array>(arr);
              if (a.length <= offset) {
                return 0;
              }

              const cap = a.length - offset;
              if (cap < n) {
                n = cap;
              }

              a.set(new Int8Array(this.mem.buffer.slice(ptr, ptr + n)), offset);
              return n;
            }

            const a = this.heap.get<Uint8Array>(arr);
            if (a.length <= offset) {
              return 0;
            }

            const cap = a.length - offset;
            if (cap < n) {
              n = cap;
            }

            a.set(new Uint8Array(this.mem.buffer.slice(ptr, ptr + n)), offset);
            return n;
          }
          case 2: {
            if (signed) {
              const a = this.heap.get<Int16Array>(arr);
              if (a.length <= offset) {
                return 0;
              }

              const cap = a.length - offset;
              if (cap < n) {
                n = cap;
              }
              for (let i = offset; i < n; i++) {
                a[i] = this.mem.getInt16(ptr + i * 2, true);
              }
              return n;
            }

            const a = this.heap.get<Uint16Array>(arr);
            if (a.length <= offset) {
              return 0;
            }

            const cap = a.length - offset;
            if (cap < n) {
              n = cap;
            }
            for (let i = offset; i < n; i++) {
              a[i] = this.mem.getUint16(ptr + i * 2, true);
            }
            return n;
          }
          case 4: {
            if (float) {
              const a = this.heap.get<Float32Array>(arr);
              if (a.length <= offset) {
                return 0;
              }

              const cap = a.length - offset;
              if (cap < n) {
                n = cap;
              }
              for (let i = offset; i < n; i++) {
                a[i] = this.mem.getFloat32(ptr + i * 4, true);
              }
              return n;
            }

            if (signed) {
              const a = this.heap.get<Int32Array>(arr);
              if (a.length <= offset) {
                return 0;
              }

              const cap = a.length - offset;
              if (cap < n) {
                n = cap;
              }
              for (let i = offset; i < n; i++) {
                a[i] = this.mem.getInt32(ptr + i * 4, true);
              }
              return n;
            }

            const a = this.heap.get<Uint32Array>(arr);
            if (a.length <= offset) {
              return 0;
            }

            const cap = a.length - offset;
            if (cap < n) {
              n = cap;
            }
            for (let i = offset; i < n; i++) {
              a[i] = this.mem.getUint32(ptr + i * 4, true);
            }
            return n;
          }
          case 8: {
            if (float) {
              const a = this.heap.get<Float64Array>(arr);
              if (a.length <= offset) {
                return 0;
              }

              const cap = a.length - offset;
              if (cap < n) {
                n = cap;
              }
              for (let i = offset; i < n; i++) {
                a[i] = this.mem.getFloat64(ptr + i * 8, true);
              }
              return n;
            }

            if (signed) {
              const a = this.heap.get<BigInt64Array>(arr);
              if (a.length <= offset) {
                return 0;
              }

              const cap = a.length - offset;
              if (cap < n) {
                n = cap;
              }
              for (let i = offset; i < n; i++) {
                a[i] = this.mem.getBigInt64(ptr + i * 8, true);
              }
              return n;
            }

            const a = this.heap.get<BigUint64Array>(arr);
            if (a.length <= offset) {
              return 0;
            }

            const cap = a.length - offset;
            if (cap < n) {
              n = cap;
            }
            for (let i = offset; i < n; i++) {
              a[i] = this.mem.getBigUint64(ptr + i * 8, true);
            }
            return n;
          }
          default:
            throw new Error(`unsupported element size=${elemSz}`);
        }
      },
      copy: (
        arr: heap.Ref<Array<any>>,
        elemSz: number,
        _signed: heap.Ref<boolean>,
        _float: heap.Ref<boolean>,
        offset: number,
        ptr: number,
        n: number
      ): number => {
        ptr >>>= 0;
        const signed = _signed === this.heap.TRUE;
        const float = _float === this.heap.TRUE;

        switch (elemSz) {
          case 0: {
            throw new Error("TODO");
          }
          case 1: {
            if (signed) {
              const a = this.heap.get<Int8Array>(arr);
              if (a.length <= offset) {
                return 0;
              }

              const cap = a.length - offset;
              if (cap < n) {
                n = cap;
              }

              new Int8Array(this.mem.buffer)
                .subarray(ptr, ptr + n)
                .set(a, offset);
              return n;
            }

            const a = this.heap.get<Uint8Array>(arr);
            if (a.length <= offset) {
              return 0;
            }

            const cap = a.length - offset;
            if (cap < n) {
              n = cap;
            }

            new Uint8Array(this.mem.buffer)
              .subarray(ptr, ptr + n)
              .set(a, offset);
            return n;
          }
          case 2: {
            if (signed) {
              const a = this.heap.get<Int16Array>(arr);
              if (a.length <= offset) {
                return 0;
              }

              const cap = a.length - offset;
              if (cap < n) {
                n = cap;
              }
              for (let i = offset; i < n; i++) {
                this.mem.setInt16(ptr + i * 2, a[i], true);
              }
              return n;
            }

            const a = this.heap.get<Uint16Array>(arr);
            if (a.length <= offset) {
              return 0;
            }

            const cap = a.length - offset;
            if (cap < n) {
              n = cap;
            }
            for (let i = offset; i < n; i++) {
              this.mem.setUint16(ptr + i * 2, a[i], true);
            }
            return n;
          }
          case 4: {
            if (float) {
              const a = this.heap.get<Float32Array>(arr);
              if (a.length <= offset) {
                return 0;
              }

              const cap = a.length - offset;
              if (cap < n) {
                n = cap;
              }
              for (let i = offset; i < n; i++) {
                this.mem.setFloat32(ptr + i * 4, a[i], true);
              }
              return n;
            }

            if (signed) {
              const a = this.heap.get<Int32Array>(arr);
              if (a.length <= offset) {
                return 0;
              }

              const cap = a.length - offset;
              if (cap < n) {
                n = cap;
              }
              for (let i = offset; i < n; i++) {
                this.mem.setInt32(ptr + i * 4, a[i], true);
              }
              return n;
            }

            const a = this.heap.get<Uint32Array>(arr);
            if (a.length <= offset) {
              return 0;
            }

            const cap = a.length - offset;
            if (cap < n) {
              n = cap;
            }
            for (let i = offset; i < n; i++) {
              this.mem.setUint32(ptr + i * 4, a[i], true);
            }
            return n;
          }
          case 8: {
            if (float) {
              const a = this.heap.get<Float64Array>(arr);
              if (a.length <= offset) {
                return 0;
              }

              const cap = a.length - offset;
              if (cap < n) {
                n = cap;
              }
              for (let i = offset; i < n; i++) {
                this.mem.setFloat64(ptr + i * 8, a[i], true);
              }
              return n;
            }

            if (signed) {
              const a = this.heap.get<BigInt64Array>(arr);
              if (a.length <= offset) {
                return 0;
              }

              const cap = a.length - offset;
              if (cap < n) {
                n = cap;
              }
              for (let i = offset; i < n; i++) {
                this.mem.setBigInt64(ptr + i * 8, a[i], true);
              }
              return n;
            }

            const a = this.heap.get<BigUint64Array>(arr);
            if (a.length <= offset) {
              return 0;
            }

            const cap = a.length - offset;
            if (cap < n) {
              n = cap;
            }
            for (let i = offset; i < n; i++) {
              this.mem.setBigUint64(ptr + i * 8, a[i], true);
            }
            return n;
          }
          default:
            throw new Error(`unsupported element size=${elemSz}`);
        }
      },
      nth: (
        self: heap.Ref<Array<any>>,
        i: number,
        ptr: number
      ): heap.Ref<any> => {
        const thiz = this.heap.getObj(self);
        if (i >= Reflect.get(thiz, "length")) {
          return this.heap.FALSE;
        }

        this.mem.setUint32(
          ptr,
          this.heap.push(Reflect.get(thiz, i, thiz)),
          true
        );
        return this.heap.TRUE;
      },
      nthNum: (
        self: heap.Ref<Array<any>>,
        i: number,
        ptr: number
      ): heap.Ref<boolean> => {
        const thiz = this.heap.getObj(self);
        if (i >= Reflect.get(thiz, "length")) {
          return this.heap.FALSE;
        }

        this.mem.setFloat64(ptr, Reflect.get(thiz, i, thiz), true);
        return this.heap.TRUE;
      },
      nthBool: (
        self: heap.Ref<Array<any>>,
        i: number,
        ptr: number
      ): heap.Ref<boolean> => {
        const thiz = this.heap.getObj(self);
        if (i >= Reflect.get(thiz, "length")) {
          return this.heap.FALSE;
        }

        this.mem.setUint8(
          ptr,
          Reflect.get(thiz, i, thiz) ? this.heap.TRUE : this.heap.FALSE
        );
        return this.heap.TRUE;
      },

      setNth: (
        self: heap.Ref<Array<any>>,
        i: number,
        free: heap.Ref<boolean>,
        val: heap.Ref<any>
      ): heap.Ref<boolean> => {
        const thiz = this.heap.getObj(self);
        return Reflect.set(
          thiz,
          i,
          free === this.heap.TRUE ? this.heap.take(val) : this.heap.get(val),
          thiz
        )
          ? this.heap.TRUE
          : this.heap.FALSE;
      },
      setNthNum: (
        self: heap.Ref<Array<any>>,
        i: number,
        val: number
      ): heap.Ref<boolean> => {
        const thiz = this.heap.getObj(self);
        return Reflect.set(thiz, i, val, thiz)
          ? this.heap.TRUE
          : this.heap.FALSE;
      },
      setNthBool: (
        self: heap.Ref<Array<any>>,
        i: number,
        val: heap.Ref<boolean>
      ): heap.Ref<boolean> => {
        const thiz = this.heap.getObj(self);
        return Reflect.set(thiz, i, val === this.heap.TRUE ? true : false, thiz)
          ? this.heap.TRUE
          : this.heap.FALSE;
      },
      setNthString: (
        self: heap.Ref<Array<any>>,
        i: number,
        ptr: number,
        sz: number
      ): heap.Ref<boolean> => {
        const thiz = this.heap.getObj(self);
        return Reflect.set(thiz, i, this.loadString(ptr, sz), thiz)
          ? this.heap.TRUE
          : this.heap.FALSE;
      },
    },
  };

  private readonly utf8enc: UTF8Encoder;
  private readonly utf8dec: UTF8Decoder;
  private readonly heap: Heap;

  public constructor(utf8enc?: UTF8Encoder, utf8dec?: UTF8Decoder) {
    if (utf8enc) {
      this.utf8enc = utf8enc;
    } else {
      if (!globalThis.TextEncoder) {
        throw new Error(
          "TextEncoder not found, please create this app with custom UTF8Encoder"
        );
      }

      this.utf8enc = new globalThis.TextEncoder();
    }

    if (utf8dec) {
      this.utf8dec = utf8dec;
    } else {
      if (!globalThis.TextDecoder) {
        throw new Error(
          "TextDecoder not found, please create this app with custom UTF8Decoder"
        );
      }

      this.utf8dec = new globalThis.TextDecoder("utf-8");
    }

    this.heap = new Heap();

    // constants defined below are used by inserted module imports (see ./@ffi.ts).

    const addModuleImport = (
      namespace: string,
      entries: Record<string, any>
    ): void => {
      if (this.importObject[namespace]) {
        // TODO: allow unique function imports in the same namespace.
        throw new Error(`duplicate namespace ${namespace}`);
      }

      this.importObject[namespace] = entries;
    };
    const HEAP = this.heap;
    const loadArgs = (ptr: number, n: number, free: boolean) => {
      // MUST be a wraped function (rather than ... = this.loadArgs), or
      // js `this` context will mess up.
      return this.loadArgs(ptr, n, free);
    };
    //__TEMPLATE_INSERT_IMPORTS__
  }
}
