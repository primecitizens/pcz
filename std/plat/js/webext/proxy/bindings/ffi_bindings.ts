import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/proxy", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_Mode": (ref: heap.Ref<string>): number => {
      const idx = ["direct", "auto_detect", "pac_script", "fixed_servers", "system"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_OnProxyErrorArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 8, false);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Ref(ptr + 0, x["details"]);
        A.store.Ref(ptr + 4, x["error"]);
        A.store.Bool(ptr + 8, x["fatal"] ? true : false);
      }
    },
    "load_OnProxyErrorArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["details"] = A.load.Ref(ptr + 0, undefined);
      x["error"] = A.load.Ref(ptr + 4, undefined);
      x["fatal"] = A.load.Bool(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_PacScript": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Ref(ptr + 0, x["data"]);
        A.store.Bool(ptr + 12, "mandatory" in x ? true : false);
        A.store.Bool(ptr + 4, x["mandatory"] ? true : false);
        A.store.Ref(ptr + 8, x["url"]);
      }
    },
    "load_PacScript": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["data"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 12)) {
        x["mandatory"] = A.load.Bool(ptr + 4);
      } else {
        delete x["mandatory"];
      }
      x["url"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_Scheme": (ref: heap.Ref<string>): number => {
      const idx = ["http", "https", "quic", "socks4", "socks5"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ProxyServer": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 21, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 20, false);
        A.store.Int64(ptr + 8, 0);
        A.store.Enum(ptr + 16, -1);
      } else {
        A.store.Bool(ptr + 21, true);
        A.store.Ref(ptr + 0, x["host"]);
        A.store.Bool(ptr + 20, "port" in x ? true : false);
        A.store.Int64(ptr + 8, x["port"] === undefined ? 0 : (x["port"] as number));
        A.store.Enum(ptr + 16, ["http", "https", "quic", "socks4", "socks5"].indexOf(x["scheme"] as string));
      }
    },
    "load_ProxyServer": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["host"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 20)) {
        x["port"] = A.load.Int64(ptr + 8);
      } else {
        delete x["port"];
      }
      x["scheme"] = A.load.Enum(ptr + 16, ["http", "https", "quic", "socks4", "socks5"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ProxyRules": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 126, false);
        A.store.Ref(ptr + 0, undefined);

        A.store.Bool(ptr + 8 + 21, false);
        A.store.Ref(ptr + 8 + 0, undefined);
        A.store.Bool(ptr + 8 + 20, false);
        A.store.Int64(ptr + 8 + 8, 0);
        A.store.Enum(ptr + 8 + 16, -1);

        A.store.Bool(ptr + 32 + 21, false);
        A.store.Ref(ptr + 32 + 0, undefined);
        A.store.Bool(ptr + 32 + 20, false);
        A.store.Int64(ptr + 32 + 8, 0);
        A.store.Enum(ptr + 32 + 16, -1);

        A.store.Bool(ptr + 56 + 21, false);
        A.store.Ref(ptr + 56 + 0, undefined);
        A.store.Bool(ptr + 56 + 20, false);
        A.store.Int64(ptr + 56 + 8, 0);
        A.store.Enum(ptr + 56 + 16, -1);

        A.store.Bool(ptr + 80 + 21, false);
        A.store.Ref(ptr + 80 + 0, undefined);
        A.store.Bool(ptr + 80 + 20, false);
        A.store.Int64(ptr + 80 + 8, 0);
        A.store.Enum(ptr + 80 + 16, -1);

        A.store.Bool(ptr + 104 + 21, false);
        A.store.Ref(ptr + 104 + 0, undefined);
        A.store.Bool(ptr + 104 + 20, false);
        A.store.Int64(ptr + 104 + 8, 0);
        A.store.Enum(ptr + 104 + 16, -1);
      } else {
        A.store.Bool(ptr + 126, true);
        A.store.Ref(ptr + 0, x["bypassList"]);

        if (typeof x["fallbackProxy"] === "undefined") {
          A.store.Bool(ptr + 8 + 21, false);
          A.store.Ref(ptr + 8 + 0, undefined);
          A.store.Bool(ptr + 8 + 20, false);
          A.store.Int64(ptr + 8 + 8, 0);
          A.store.Enum(ptr + 8 + 16, -1);
        } else {
          A.store.Bool(ptr + 8 + 21, true);
          A.store.Ref(ptr + 8 + 0, x["fallbackProxy"]["host"]);
          A.store.Bool(ptr + 8 + 20, "port" in x["fallbackProxy"] ? true : false);
          A.store.Int64(
            ptr + 8 + 8,
            x["fallbackProxy"]["port"] === undefined ? 0 : (x["fallbackProxy"]["port"] as number)
          );
          A.store.Enum(
            ptr + 8 + 16,
            ["http", "https", "quic", "socks4", "socks5"].indexOf(x["fallbackProxy"]["scheme"] as string)
          );
        }

        if (typeof x["proxyForFtp"] === "undefined") {
          A.store.Bool(ptr + 32 + 21, false);
          A.store.Ref(ptr + 32 + 0, undefined);
          A.store.Bool(ptr + 32 + 20, false);
          A.store.Int64(ptr + 32 + 8, 0);
          A.store.Enum(ptr + 32 + 16, -1);
        } else {
          A.store.Bool(ptr + 32 + 21, true);
          A.store.Ref(ptr + 32 + 0, x["proxyForFtp"]["host"]);
          A.store.Bool(ptr + 32 + 20, "port" in x["proxyForFtp"] ? true : false);
          A.store.Int64(
            ptr + 32 + 8,
            x["proxyForFtp"]["port"] === undefined ? 0 : (x["proxyForFtp"]["port"] as number)
          );
          A.store.Enum(
            ptr + 32 + 16,
            ["http", "https", "quic", "socks4", "socks5"].indexOf(x["proxyForFtp"]["scheme"] as string)
          );
        }

        if (typeof x["proxyForHttp"] === "undefined") {
          A.store.Bool(ptr + 56 + 21, false);
          A.store.Ref(ptr + 56 + 0, undefined);
          A.store.Bool(ptr + 56 + 20, false);
          A.store.Int64(ptr + 56 + 8, 0);
          A.store.Enum(ptr + 56 + 16, -1);
        } else {
          A.store.Bool(ptr + 56 + 21, true);
          A.store.Ref(ptr + 56 + 0, x["proxyForHttp"]["host"]);
          A.store.Bool(ptr + 56 + 20, "port" in x["proxyForHttp"] ? true : false);
          A.store.Int64(
            ptr + 56 + 8,
            x["proxyForHttp"]["port"] === undefined ? 0 : (x["proxyForHttp"]["port"] as number)
          );
          A.store.Enum(
            ptr + 56 + 16,
            ["http", "https", "quic", "socks4", "socks5"].indexOf(x["proxyForHttp"]["scheme"] as string)
          );
        }

        if (typeof x["proxyForHttps"] === "undefined") {
          A.store.Bool(ptr + 80 + 21, false);
          A.store.Ref(ptr + 80 + 0, undefined);
          A.store.Bool(ptr + 80 + 20, false);
          A.store.Int64(ptr + 80 + 8, 0);
          A.store.Enum(ptr + 80 + 16, -1);
        } else {
          A.store.Bool(ptr + 80 + 21, true);
          A.store.Ref(ptr + 80 + 0, x["proxyForHttps"]["host"]);
          A.store.Bool(ptr + 80 + 20, "port" in x["proxyForHttps"] ? true : false);
          A.store.Int64(
            ptr + 80 + 8,
            x["proxyForHttps"]["port"] === undefined ? 0 : (x["proxyForHttps"]["port"] as number)
          );
          A.store.Enum(
            ptr + 80 + 16,
            ["http", "https", "quic", "socks4", "socks5"].indexOf(x["proxyForHttps"]["scheme"] as string)
          );
        }

        if (typeof x["singleProxy"] === "undefined") {
          A.store.Bool(ptr + 104 + 21, false);
          A.store.Ref(ptr + 104 + 0, undefined);
          A.store.Bool(ptr + 104 + 20, false);
          A.store.Int64(ptr + 104 + 8, 0);
          A.store.Enum(ptr + 104 + 16, -1);
        } else {
          A.store.Bool(ptr + 104 + 21, true);
          A.store.Ref(ptr + 104 + 0, x["singleProxy"]["host"]);
          A.store.Bool(ptr + 104 + 20, "port" in x["singleProxy"] ? true : false);
          A.store.Int64(
            ptr + 104 + 8,
            x["singleProxy"]["port"] === undefined ? 0 : (x["singleProxy"]["port"] as number)
          );
          A.store.Enum(
            ptr + 104 + 16,
            ["http", "https", "quic", "socks4", "socks5"].indexOf(x["singleProxy"]["scheme"] as string)
          );
        }
      }
    },
    "load_ProxyRules": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["bypassList"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 8 + 21)) {
        x["fallbackProxy"] = {};
        x["fallbackProxy"]["host"] = A.load.Ref(ptr + 8 + 0, undefined);
        if (A.load.Bool(ptr + 8 + 20)) {
          x["fallbackProxy"]["port"] = A.load.Int64(ptr + 8 + 8);
        } else {
          delete x["fallbackProxy"]["port"];
        }
        x["fallbackProxy"]["scheme"] = A.load.Enum(ptr + 8 + 16, ["http", "https", "quic", "socks4", "socks5"]);
      } else {
        delete x["fallbackProxy"];
      }
      if (A.load.Bool(ptr + 32 + 21)) {
        x["proxyForFtp"] = {};
        x["proxyForFtp"]["host"] = A.load.Ref(ptr + 32 + 0, undefined);
        if (A.load.Bool(ptr + 32 + 20)) {
          x["proxyForFtp"]["port"] = A.load.Int64(ptr + 32 + 8);
        } else {
          delete x["proxyForFtp"]["port"];
        }
        x["proxyForFtp"]["scheme"] = A.load.Enum(ptr + 32 + 16, ["http", "https", "quic", "socks4", "socks5"]);
      } else {
        delete x["proxyForFtp"];
      }
      if (A.load.Bool(ptr + 56 + 21)) {
        x["proxyForHttp"] = {};
        x["proxyForHttp"]["host"] = A.load.Ref(ptr + 56 + 0, undefined);
        if (A.load.Bool(ptr + 56 + 20)) {
          x["proxyForHttp"]["port"] = A.load.Int64(ptr + 56 + 8);
        } else {
          delete x["proxyForHttp"]["port"];
        }
        x["proxyForHttp"]["scheme"] = A.load.Enum(ptr + 56 + 16, ["http", "https", "quic", "socks4", "socks5"]);
      } else {
        delete x["proxyForHttp"];
      }
      if (A.load.Bool(ptr + 80 + 21)) {
        x["proxyForHttps"] = {};
        x["proxyForHttps"]["host"] = A.load.Ref(ptr + 80 + 0, undefined);
        if (A.load.Bool(ptr + 80 + 20)) {
          x["proxyForHttps"]["port"] = A.load.Int64(ptr + 80 + 8);
        } else {
          delete x["proxyForHttps"]["port"];
        }
        x["proxyForHttps"]["scheme"] = A.load.Enum(ptr + 80 + 16, ["http", "https", "quic", "socks4", "socks5"]);
      } else {
        delete x["proxyForHttps"];
      }
      if (A.load.Bool(ptr + 104 + 21)) {
        x["singleProxy"] = {};
        x["singleProxy"]["host"] = A.load.Ref(ptr + 104 + 0, undefined);
        if (A.load.Bool(ptr + 104 + 20)) {
          x["singleProxy"]["port"] = A.load.Int64(ptr + 104 + 8);
        } else {
          delete x["singleProxy"]["port"];
        }
        x["singleProxy"]["scheme"] = A.load.Enum(ptr + 104 + 16, ["http", "https", "quic", "socks4", "socks5"]);
      } else {
        delete x["singleProxy"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ProxyConfig": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 151, false);
        A.store.Enum(ptr + 0, -1);

        A.store.Bool(ptr + 4 + 13, false);
        A.store.Ref(ptr + 4 + 0, undefined);
        A.store.Bool(ptr + 4 + 12, false);
        A.store.Bool(ptr + 4 + 4, false);
        A.store.Ref(ptr + 4 + 8, undefined);

        A.store.Bool(ptr + 24 + 126, false);
        A.store.Ref(ptr + 24 + 0, undefined);

        A.store.Bool(ptr + 24 + 8 + 21, false);
        A.store.Ref(ptr + 24 + 8 + 0, undefined);
        A.store.Bool(ptr + 24 + 8 + 20, false);
        A.store.Int64(ptr + 24 + 8 + 8, 0);
        A.store.Enum(ptr + 24 + 8 + 16, -1);

        A.store.Bool(ptr + 24 + 32 + 21, false);
        A.store.Ref(ptr + 24 + 32 + 0, undefined);
        A.store.Bool(ptr + 24 + 32 + 20, false);
        A.store.Int64(ptr + 24 + 32 + 8, 0);
        A.store.Enum(ptr + 24 + 32 + 16, -1);

        A.store.Bool(ptr + 24 + 56 + 21, false);
        A.store.Ref(ptr + 24 + 56 + 0, undefined);
        A.store.Bool(ptr + 24 + 56 + 20, false);
        A.store.Int64(ptr + 24 + 56 + 8, 0);
        A.store.Enum(ptr + 24 + 56 + 16, -1);

        A.store.Bool(ptr + 24 + 80 + 21, false);
        A.store.Ref(ptr + 24 + 80 + 0, undefined);
        A.store.Bool(ptr + 24 + 80 + 20, false);
        A.store.Int64(ptr + 24 + 80 + 8, 0);
        A.store.Enum(ptr + 24 + 80 + 16, -1);

        A.store.Bool(ptr + 24 + 104 + 21, false);
        A.store.Ref(ptr + 24 + 104 + 0, undefined);
        A.store.Bool(ptr + 24 + 104 + 20, false);
        A.store.Int64(ptr + 24 + 104 + 8, 0);
        A.store.Enum(ptr + 24 + 104 + 16, -1);
      } else {
        A.store.Bool(ptr + 151, true);
        A.store.Enum(
          ptr + 0,
          ["direct", "auto_detect", "pac_script", "fixed_servers", "system"].indexOf(x["mode"] as string)
        );

        if (typeof x["pacScript"] === "undefined") {
          A.store.Bool(ptr + 4 + 13, false);
          A.store.Ref(ptr + 4 + 0, undefined);
          A.store.Bool(ptr + 4 + 12, false);
          A.store.Bool(ptr + 4 + 4, false);
          A.store.Ref(ptr + 4 + 8, undefined);
        } else {
          A.store.Bool(ptr + 4 + 13, true);
          A.store.Ref(ptr + 4 + 0, x["pacScript"]["data"]);
          A.store.Bool(ptr + 4 + 12, "mandatory" in x["pacScript"] ? true : false);
          A.store.Bool(ptr + 4 + 4, x["pacScript"]["mandatory"] ? true : false);
          A.store.Ref(ptr + 4 + 8, x["pacScript"]["url"]);
        }

        if (typeof x["rules"] === "undefined") {
          A.store.Bool(ptr + 24 + 126, false);
          A.store.Ref(ptr + 24 + 0, undefined);

          A.store.Bool(ptr + 24 + 8 + 21, false);
          A.store.Ref(ptr + 24 + 8 + 0, undefined);
          A.store.Bool(ptr + 24 + 8 + 20, false);
          A.store.Int64(ptr + 24 + 8 + 8, 0);
          A.store.Enum(ptr + 24 + 8 + 16, -1);

          A.store.Bool(ptr + 24 + 32 + 21, false);
          A.store.Ref(ptr + 24 + 32 + 0, undefined);
          A.store.Bool(ptr + 24 + 32 + 20, false);
          A.store.Int64(ptr + 24 + 32 + 8, 0);
          A.store.Enum(ptr + 24 + 32 + 16, -1);

          A.store.Bool(ptr + 24 + 56 + 21, false);
          A.store.Ref(ptr + 24 + 56 + 0, undefined);
          A.store.Bool(ptr + 24 + 56 + 20, false);
          A.store.Int64(ptr + 24 + 56 + 8, 0);
          A.store.Enum(ptr + 24 + 56 + 16, -1);

          A.store.Bool(ptr + 24 + 80 + 21, false);
          A.store.Ref(ptr + 24 + 80 + 0, undefined);
          A.store.Bool(ptr + 24 + 80 + 20, false);
          A.store.Int64(ptr + 24 + 80 + 8, 0);
          A.store.Enum(ptr + 24 + 80 + 16, -1);

          A.store.Bool(ptr + 24 + 104 + 21, false);
          A.store.Ref(ptr + 24 + 104 + 0, undefined);
          A.store.Bool(ptr + 24 + 104 + 20, false);
          A.store.Int64(ptr + 24 + 104 + 8, 0);
          A.store.Enum(ptr + 24 + 104 + 16, -1);
        } else {
          A.store.Bool(ptr + 24 + 126, true);
          A.store.Ref(ptr + 24 + 0, x["rules"]["bypassList"]);

          if (typeof x["rules"]["fallbackProxy"] === "undefined") {
            A.store.Bool(ptr + 24 + 8 + 21, false);
            A.store.Ref(ptr + 24 + 8 + 0, undefined);
            A.store.Bool(ptr + 24 + 8 + 20, false);
            A.store.Int64(ptr + 24 + 8 + 8, 0);
            A.store.Enum(ptr + 24 + 8 + 16, -1);
          } else {
            A.store.Bool(ptr + 24 + 8 + 21, true);
            A.store.Ref(ptr + 24 + 8 + 0, x["rules"]["fallbackProxy"]["host"]);
            A.store.Bool(ptr + 24 + 8 + 20, "port" in x["rules"]["fallbackProxy"] ? true : false);
            A.store.Int64(
              ptr + 24 + 8 + 8,
              x["rules"]["fallbackProxy"]["port"] === undefined ? 0 : (x["rules"]["fallbackProxy"]["port"] as number)
            );
            A.store.Enum(
              ptr + 24 + 8 + 16,
              ["http", "https", "quic", "socks4", "socks5"].indexOf(x["rules"]["fallbackProxy"]["scheme"] as string)
            );
          }

          if (typeof x["rules"]["proxyForFtp"] === "undefined") {
            A.store.Bool(ptr + 24 + 32 + 21, false);
            A.store.Ref(ptr + 24 + 32 + 0, undefined);
            A.store.Bool(ptr + 24 + 32 + 20, false);
            A.store.Int64(ptr + 24 + 32 + 8, 0);
            A.store.Enum(ptr + 24 + 32 + 16, -1);
          } else {
            A.store.Bool(ptr + 24 + 32 + 21, true);
            A.store.Ref(ptr + 24 + 32 + 0, x["rules"]["proxyForFtp"]["host"]);
            A.store.Bool(ptr + 24 + 32 + 20, "port" in x["rules"]["proxyForFtp"] ? true : false);
            A.store.Int64(
              ptr + 24 + 32 + 8,
              x["rules"]["proxyForFtp"]["port"] === undefined ? 0 : (x["rules"]["proxyForFtp"]["port"] as number)
            );
            A.store.Enum(
              ptr + 24 + 32 + 16,
              ["http", "https", "quic", "socks4", "socks5"].indexOf(x["rules"]["proxyForFtp"]["scheme"] as string)
            );
          }

          if (typeof x["rules"]["proxyForHttp"] === "undefined") {
            A.store.Bool(ptr + 24 + 56 + 21, false);
            A.store.Ref(ptr + 24 + 56 + 0, undefined);
            A.store.Bool(ptr + 24 + 56 + 20, false);
            A.store.Int64(ptr + 24 + 56 + 8, 0);
            A.store.Enum(ptr + 24 + 56 + 16, -1);
          } else {
            A.store.Bool(ptr + 24 + 56 + 21, true);
            A.store.Ref(ptr + 24 + 56 + 0, x["rules"]["proxyForHttp"]["host"]);
            A.store.Bool(ptr + 24 + 56 + 20, "port" in x["rules"]["proxyForHttp"] ? true : false);
            A.store.Int64(
              ptr + 24 + 56 + 8,
              x["rules"]["proxyForHttp"]["port"] === undefined ? 0 : (x["rules"]["proxyForHttp"]["port"] as number)
            );
            A.store.Enum(
              ptr + 24 + 56 + 16,
              ["http", "https", "quic", "socks4", "socks5"].indexOf(x["rules"]["proxyForHttp"]["scheme"] as string)
            );
          }

          if (typeof x["rules"]["proxyForHttps"] === "undefined") {
            A.store.Bool(ptr + 24 + 80 + 21, false);
            A.store.Ref(ptr + 24 + 80 + 0, undefined);
            A.store.Bool(ptr + 24 + 80 + 20, false);
            A.store.Int64(ptr + 24 + 80 + 8, 0);
            A.store.Enum(ptr + 24 + 80 + 16, -1);
          } else {
            A.store.Bool(ptr + 24 + 80 + 21, true);
            A.store.Ref(ptr + 24 + 80 + 0, x["rules"]["proxyForHttps"]["host"]);
            A.store.Bool(ptr + 24 + 80 + 20, "port" in x["rules"]["proxyForHttps"] ? true : false);
            A.store.Int64(
              ptr + 24 + 80 + 8,
              x["rules"]["proxyForHttps"]["port"] === undefined ? 0 : (x["rules"]["proxyForHttps"]["port"] as number)
            );
            A.store.Enum(
              ptr + 24 + 80 + 16,
              ["http", "https", "quic", "socks4", "socks5"].indexOf(x["rules"]["proxyForHttps"]["scheme"] as string)
            );
          }

          if (typeof x["rules"]["singleProxy"] === "undefined") {
            A.store.Bool(ptr + 24 + 104 + 21, false);
            A.store.Ref(ptr + 24 + 104 + 0, undefined);
            A.store.Bool(ptr + 24 + 104 + 20, false);
            A.store.Int64(ptr + 24 + 104 + 8, 0);
            A.store.Enum(ptr + 24 + 104 + 16, -1);
          } else {
            A.store.Bool(ptr + 24 + 104 + 21, true);
            A.store.Ref(ptr + 24 + 104 + 0, x["rules"]["singleProxy"]["host"]);
            A.store.Bool(ptr + 24 + 104 + 20, "port" in x["rules"]["singleProxy"] ? true : false);
            A.store.Int64(
              ptr + 24 + 104 + 8,
              x["rules"]["singleProxy"]["port"] === undefined ? 0 : (x["rules"]["singleProxy"]["port"] as number)
            );
            A.store.Enum(
              ptr + 24 + 104 + 16,
              ["http", "https", "quic", "socks4", "socks5"].indexOf(x["rules"]["singleProxy"]["scheme"] as string)
            );
          }
        }
      }
    },
    "load_ProxyConfig": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["mode"] = A.load.Enum(ptr + 0, ["direct", "auto_detect", "pac_script", "fixed_servers", "system"]);
      if (A.load.Bool(ptr + 4 + 13)) {
        x["pacScript"] = {};
        x["pacScript"]["data"] = A.load.Ref(ptr + 4 + 0, undefined);
        if (A.load.Bool(ptr + 4 + 12)) {
          x["pacScript"]["mandatory"] = A.load.Bool(ptr + 4 + 4);
        } else {
          delete x["pacScript"]["mandatory"];
        }
        x["pacScript"]["url"] = A.load.Ref(ptr + 4 + 8, undefined);
      } else {
        delete x["pacScript"];
      }
      if (A.load.Bool(ptr + 24 + 126)) {
        x["rules"] = {};
        x["rules"]["bypassList"] = A.load.Ref(ptr + 24 + 0, undefined);
        if (A.load.Bool(ptr + 24 + 8 + 21)) {
          x["rules"]["fallbackProxy"] = {};
          x["rules"]["fallbackProxy"]["host"] = A.load.Ref(ptr + 24 + 8 + 0, undefined);
          if (A.load.Bool(ptr + 24 + 8 + 20)) {
            x["rules"]["fallbackProxy"]["port"] = A.load.Int64(ptr + 24 + 8 + 8);
          } else {
            delete x["rules"]["fallbackProxy"]["port"];
          }
          x["rules"]["fallbackProxy"]["scheme"] = A.load.Enum(ptr + 24 + 8 + 16, [
            "http",
            "https",
            "quic",
            "socks4",
            "socks5",
          ]);
        } else {
          delete x["rules"]["fallbackProxy"];
        }
        if (A.load.Bool(ptr + 24 + 32 + 21)) {
          x["rules"]["proxyForFtp"] = {};
          x["rules"]["proxyForFtp"]["host"] = A.load.Ref(ptr + 24 + 32 + 0, undefined);
          if (A.load.Bool(ptr + 24 + 32 + 20)) {
            x["rules"]["proxyForFtp"]["port"] = A.load.Int64(ptr + 24 + 32 + 8);
          } else {
            delete x["rules"]["proxyForFtp"]["port"];
          }
          x["rules"]["proxyForFtp"]["scheme"] = A.load.Enum(ptr + 24 + 32 + 16, [
            "http",
            "https",
            "quic",
            "socks4",
            "socks5",
          ]);
        } else {
          delete x["rules"]["proxyForFtp"];
        }
        if (A.load.Bool(ptr + 24 + 56 + 21)) {
          x["rules"]["proxyForHttp"] = {};
          x["rules"]["proxyForHttp"]["host"] = A.load.Ref(ptr + 24 + 56 + 0, undefined);
          if (A.load.Bool(ptr + 24 + 56 + 20)) {
            x["rules"]["proxyForHttp"]["port"] = A.load.Int64(ptr + 24 + 56 + 8);
          } else {
            delete x["rules"]["proxyForHttp"]["port"];
          }
          x["rules"]["proxyForHttp"]["scheme"] = A.load.Enum(ptr + 24 + 56 + 16, [
            "http",
            "https",
            "quic",
            "socks4",
            "socks5",
          ]);
        } else {
          delete x["rules"]["proxyForHttp"];
        }
        if (A.load.Bool(ptr + 24 + 80 + 21)) {
          x["rules"]["proxyForHttps"] = {};
          x["rules"]["proxyForHttps"]["host"] = A.load.Ref(ptr + 24 + 80 + 0, undefined);
          if (A.load.Bool(ptr + 24 + 80 + 20)) {
            x["rules"]["proxyForHttps"]["port"] = A.load.Int64(ptr + 24 + 80 + 8);
          } else {
            delete x["rules"]["proxyForHttps"]["port"];
          }
          x["rules"]["proxyForHttps"]["scheme"] = A.load.Enum(ptr + 24 + 80 + 16, [
            "http",
            "https",
            "quic",
            "socks4",
            "socks5",
          ]);
        } else {
          delete x["rules"]["proxyForHttps"];
        }
        if (A.load.Bool(ptr + 24 + 104 + 21)) {
          x["rules"]["singleProxy"] = {};
          x["rules"]["singleProxy"]["host"] = A.load.Ref(ptr + 24 + 104 + 0, undefined);
          if (A.load.Bool(ptr + 24 + 104 + 20)) {
            x["rules"]["singleProxy"]["port"] = A.load.Int64(ptr + 24 + 104 + 8);
          } else {
            delete x["rules"]["singleProxy"]["port"];
          }
          x["rules"]["singleProxy"]["scheme"] = A.load.Enum(ptr + 24 + 104 + 16, [
            "http",
            "https",
            "quic",
            "socks4",
            "socks5",
          ]);
        } else {
          delete x["rules"]["singleProxy"];
        }
      } else {
        delete x["rules"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_OnProxyError": (): heap.Ref<boolean> => {
      if (WEBEXT?.proxy?.onProxyError && "addListener" in WEBEXT?.proxy?.onProxyError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnProxyError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.proxy.onProxyError.addListener);
    },
    "call_OnProxyError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.proxy.onProxyError.addListener(A.H.get<object>(callback));
    },
    "try_OnProxyError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.proxy.onProxyError.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffProxyError": (): heap.Ref<boolean> => {
      if (WEBEXT?.proxy?.onProxyError && "removeListener" in WEBEXT?.proxy?.onProxyError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffProxyError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.proxy.onProxyError.removeListener);
    },
    "call_OffProxyError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.proxy.onProxyError.removeListener(A.H.get<object>(callback));
    },
    "try_OffProxyError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.proxy.onProxyError.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnProxyError": (): heap.Ref<boolean> => {
      if (WEBEXT?.proxy?.onProxyError && "hasListener" in WEBEXT?.proxy?.onProxyError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnProxyError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.proxy.onProxyError.hasListener);
    },
    "call_HasOnProxyError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.proxy.onProxyError.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnProxyError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.proxy.onProxyError.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "get_Settings": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.proxy && "settings" in WEBEXT?.proxy) {
        const val = WEBEXT.proxy.settings;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Settings": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.proxy, "settings", A.H.get<object>(val), WEBEXT.proxy) ? A.H.TRUE : A.H.FALSE;
    },
  };
});
