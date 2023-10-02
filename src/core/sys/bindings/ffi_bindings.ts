import { Application, Pointer, importModule } from "@ffi";

importModule("core/sys", (A: Application) => {
  let _args: string[] = [];
  let _environ: string[] = [];
  Object.defineProperty(A, "args", {
    get: (): string[] => {
      return _args;
    },
    set: (v: string[]) => {
      _args = v;
    },
  });
  Object.defineProperty(A, "environ", {
    get: (): string[] => {
      return _environ;
    },
    set: (v: string[]) => {
      _environ = v;
    },
  });

  return {
    "sizes": (pArgcAndEnvc: Pointer): number => {
      pArgcAndEnvc >>>= 0;
      let sum = 0;
      _args.forEach((v) => {
        sum += A.UTF8Sizeof(v);
      });
      _environ.forEach((v) => {
        sum += A.UTF8Sizeof(v);
      });

      A.store.Uint32(pArgcAndEnvc, _args.length + _environ.length);
      return sum;
    },
    "nth": (i: number, buf: Pointer, sz: number): number => {
      if (i < _args.length) {
        return A.store.String(buf, sz, _args[i]).n;
      }
      if (i === _args.length || i > _args.length + _environ.length) {
        return -1;
      }

      return A.store.String(buf, sz, _environ[i - _args.length - 1]).n;
    },
  };
});
