import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/fontsettings", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_ClearDefaultFixedFontSizeArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 0, true);
      }
    },
    "load_ClearDefaultFixedFontSizeArgDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ClearDefaultFontSizeArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 0, true);
      }
    },
    "load_ClearDefaultFontSizeArgDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_GenericFamily": (ref: heap.Ref<string>): number => {
      const idx = ["standard", "sansserif", "serif", "fixed", "cursive", "fantasy", "math"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_ScriptCode": (ref: heap.Ref<string>): number => {
      const idx = [
        "Afak",
        "Arab",
        "Armi",
        "Armn",
        "Avst",
        "Bali",
        "Bamu",
        "Bass",
        "Batk",
        "Beng",
        "Blis",
        "Bopo",
        "Brah",
        "Brai",
        "Bugi",
        "Buhd",
        "Cakm",
        "Cans",
        "Cari",
        "Cham",
        "Cher",
        "Cirt",
        "Copt",
        "Cprt",
        "Cyrl",
        "Cyrs",
        "Deva",
        "Dsrt",
        "Dupl",
        "Egyd",
        "Egyh",
        "Egyp",
        "Elba",
        "Ethi",
        "Geor",
        "Geok",
        "Glag",
        "Goth",
        "Gran",
        "Grek",
        "Gujr",
        "Guru",
        "Hang",
        "Hani",
        "Hano",
        "Hans",
        "Hant",
        "Hebr",
        "Hluw",
        "Hmng",
        "Hung",
        "Inds",
        "Ital",
        "Java",
        "Jpan",
        "Jurc",
        "Kali",
        "Khar",
        "Khmr",
        "Khoj",
        "Knda",
        "Kpel",
        "Kthi",
        "Lana",
        "Laoo",
        "Latf",
        "Latg",
        "Latn",
        "Lepc",
        "Limb",
        "Lina",
        "Linb",
        "Lisu",
        "Loma",
        "Lyci",
        "Lydi",
        "Mand",
        "Mani",
        "Maya",
        "Mend",
        "Merc",
        "Mero",
        "Mlym",
        "Moon",
        "Mong",
        "Mroo",
        "Mtei",
        "Mymr",
        "Narb",
        "Nbat",
        "Nkgb",
        "Nkoo",
        "Nshu",
        "Ogam",
        "Olck",
        "Orkh",
        "Orya",
        "Osma",
        "Palm",
        "Perm",
        "Phag",
        "Phli",
        "Phlp",
        "Phlv",
        "Phnx",
        "Plrd",
        "Prti",
        "Rjng",
        "Roro",
        "Runr",
        "Samr",
        "Sara",
        "Sarb",
        "Saur",
        "Sgnw",
        "Shaw",
        "Shrd",
        "Sind",
        "Sinh",
        "Sora",
        "Sund",
        "Sylo",
        "Syrc",
        "Syre",
        "Syrj",
        "Syrn",
        "Tagb",
        "Takr",
        "Tale",
        "Talu",
        "Taml",
        "Tang",
        "Tavt",
        "Telu",
        "Teng",
        "Tfng",
        "Tglg",
        "Thaa",
        "Thai",
        "Tibt",
        "Tirh",
        "Ugar",
        "Vaii",
        "Visp",
        "Wara",
        "Wole",
        "Xpeo",
        "Xsux",
        "Yiii",
        "Zmth",
        "Zsym",
        "Zyyy",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ClearFontArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(
          ptr + 0,
          ["standard", "sansserif", "serif", "fixed", "cursive", "fantasy", "math"].indexOf(
            x["genericFamily"] as string
          )
        );
        A.store.Enum(
          ptr + 4,
          [
            "Afak",
            "Arab",
            "Armi",
            "Armn",
            "Avst",
            "Bali",
            "Bamu",
            "Bass",
            "Batk",
            "Beng",
            "Blis",
            "Bopo",
            "Brah",
            "Brai",
            "Bugi",
            "Buhd",
            "Cakm",
            "Cans",
            "Cari",
            "Cham",
            "Cher",
            "Cirt",
            "Copt",
            "Cprt",
            "Cyrl",
            "Cyrs",
            "Deva",
            "Dsrt",
            "Dupl",
            "Egyd",
            "Egyh",
            "Egyp",
            "Elba",
            "Ethi",
            "Geor",
            "Geok",
            "Glag",
            "Goth",
            "Gran",
            "Grek",
            "Gujr",
            "Guru",
            "Hang",
            "Hani",
            "Hano",
            "Hans",
            "Hant",
            "Hebr",
            "Hluw",
            "Hmng",
            "Hung",
            "Inds",
            "Ital",
            "Java",
            "Jpan",
            "Jurc",
            "Kali",
            "Khar",
            "Khmr",
            "Khoj",
            "Knda",
            "Kpel",
            "Kthi",
            "Lana",
            "Laoo",
            "Latf",
            "Latg",
            "Latn",
            "Lepc",
            "Limb",
            "Lina",
            "Linb",
            "Lisu",
            "Loma",
            "Lyci",
            "Lydi",
            "Mand",
            "Mani",
            "Maya",
            "Mend",
            "Merc",
            "Mero",
            "Mlym",
            "Moon",
            "Mong",
            "Mroo",
            "Mtei",
            "Mymr",
            "Narb",
            "Nbat",
            "Nkgb",
            "Nkoo",
            "Nshu",
            "Ogam",
            "Olck",
            "Orkh",
            "Orya",
            "Osma",
            "Palm",
            "Perm",
            "Phag",
            "Phli",
            "Phlp",
            "Phlv",
            "Phnx",
            "Plrd",
            "Prti",
            "Rjng",
            "Roro",
            "Runr",
            "Samr",
            "Sara",
            "Sarb",
            "Saur",
            "Sgnw",
            "Shaw",
            "Shrd",
            "Sind",
            "Sinh",
            "Sora",
            "Sund",
            "Sylo",
            "Syrc",
            "Syre",
            "Syrj",
            "Syrn",
            "Tagb",
            "Takr",
            "Tale",
            "Talu",
            "Taml",
            "Tang",
            "Tavt",
            "Telu",
            "Teng",
            "Tfng",
            "Tglg",
            "Thaa",
            "Thai",
            "Tibt",
            "Tirh",
            "Ugar",
            "Vaii",
            "Visp",
            "Wara",
            "Wole",
            "Xpeo",
            "Xsux",
            "Yiii",
            "Zmth",
            "Zsym",
            "Zyyy",
          ].indexOf(x["script"] as string)
        );
      }
    },
    "load_ClearFontArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["genericFamily"] = A.load.Enum(ptr + 0, [
        "standard",
        "sansserif",
        "serif",
        "fixed",
        "cursive",
        "fantasy",
        "math",
      ]);
      x["script"] = A.load.Enum(ptr + 4, [
        "Afak",
        "Arab",
        "Armi",
        "Armn",
        "Avst",
        "Bali",
        "Bamu",
        "Bass",
        "Batk",
        "Beng",
        "Blis",
        "Bopo",
        "Brah",
        "Brai",
        "Bugi",
        "Buhd",
        "Cakm",
        "Cans",
        "Cari",
        "Cham",
        "Cher",
        "Cirt",
        "Copt",
        "Cprt",
        "Cyrl",
        "Cyrs",
        "Deva",
        "Dsrt",
        "Dupl",
        "Egyd",
        "Egyh",
        "Egyp",
        "Elba",
        "Ethi",
        "Geor",
        "Geok",
        "Glag",
        "Goth",
        "Gran",
        "Grek",
        "Gujr",
        "Guru",
        "Hang",
        "Hani",
        "Hano",
        "Hans",
        "Hant",
        "Hebr",
        "Hluw",
        "Hmng",
        "Hung",
        "Inds",
        "Ital",
        "Java",
        "Jpan",
        "Jurc",
        "Kali",
        "Khar",
        "Khmr",
        "Khoj",
        "Knda",
        "Kpel",
        "Kthi",
        "Lana",
        "Laoo",
        "Latf",
        "Latg",
        "Latn",
        "Lepc",
        "Limb",
        "Lina",
        "Linb",
        "Lisu",
        "Loma",
        "Lyci",
        "Lydi",
        "Mand",
        "Mani",
        "Maya",
        "Mend",
        "Merc",
        "Mero",
        "Mlym",
        "Moon",
        "Mong",
        "Mroo",
        "Mtei",
        "Mymr",
        "Narb",
        "Nbat",
        "Nkgb",
        "Nkoo",
        "Nshu",
        "Ogam",
        "Olck",
        "Orkh",
        "Orya",
        "Osma",
        "Palm",
        "Perm",
        "Phag",
        "Phli",
        "Phlp",
        "Phlv",
        "Phnx",
        "Plrd",
        "Prti",
        "Rjng",
        "Roro",
        "Runr",
        "Samr",
        "Sara",
        "Sarb",
        "Saur",
        "Sgnw",
        "Shaw",
        "Shrd",
        "Sind",
        "Sinh",
        "Sora",
        "Sund",
        "Sylo",
        "Syrc",
        "Syre",
        "Syrj",
        "Syrn",
        "Tagb",
        "Takr",
        "Tale",
        "Talu",
        "Taml",
        "Tang",
        "Tavt",
        "Telu",
        "Teng",
        "Tfng",
        "Tglg",
        "Thaa",
        "Thai",
        "Tibt",
        "Tirh",
        "Ugar",
        "Vaii",
        "Visp",
        "Wara",
        "Wole",
        "Xpeo",
        "Xsux",
        "Yiii",
        "Zmth",
        "Zsym",
        "Zyyy",
      ]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ClearMinimumFontSizeArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 0, true);
      }
    },
    "load_ClearMinimumFontSizeArgDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_FontName": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["displayName"]);
        A.store.Ref(ptr + 4, x["fontId"]);
      }
    },
    "load_FontName": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["displayName"] = A.load.Ref(ptr + 0, undefined);
      x["fontId"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetDefaultFixedFontSizeArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 0, true);
      }
    },
    "load_GetDefaultFixedFontSizeArgDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_LevelOfControl": (ref: heap.Ref<string>): number => {
      const idx = [
        "not_controllable",
        "controlled_by_other_extensions",
        "controllable_by_this_extension",
        "controlled_by_this_extension",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_GetDefaultFixedFontSizeReturnType": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Enum(
          ptr + 0,
          [
            "not_controllable",
            "controlled_by_other_extensions",
            "controllable_by_this_extension",
            "controlled_by_this_extension",
          ].indexOf(x["levelOfControl"] as string)
        );
        A.store.Int64(ptr + 8, x["pixelSize"] === undefined ? 0 : (x["pixelSize"] as number));
      }
    },
    "load_GetDefaultFixedFontSizeReturnType": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["levelOfControl"] = A.load.Enum(ptr + 0, [
        "not_controllable",
        "controlled_by_other_extensions",
        "controllable_by_this_extension",
        "controlled_by_this_extension",
      ]);
      x["pixelSize"] = A.load.Int64(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetDefaultFontSizeArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 0, true);
      }
    },
    "load_GetDefaultFontSizeArgDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetDefaultFontSizeReturnType": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Enum(
          ptr + 0,
          [
            "not_controllable",
            "controlled_by_other_extensions",
            "controllable_by_this_extension",
            "controlled_by_this_extension",
          ].indexOf(x["levelOfControl"] as string)
        );
        A.store.Int64(ptr + 8, x["pixelSize"] === undefined ? 0 : (x["pixelSize"] as number));
      }
    },
    "load_GetDefaultFontSizeReturnType": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["levelOfControl"] = A.load.Enum(ptr + 0, [
        "not_controllable",
        "controlled_by_other_extensions",
        "controllable_by_this_extension",
        "controlled_by_this_extension",
      ]);
      x["pixelSize"] = A.load.Int64(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetFontArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(
          ptr + 0,
          ["standard", "sansserif", "serif", "fixed", "cursive", "fantasy", "math"].indexOf(
            x["genericFamily"] as string
          )
        );
        A.store.Enum(
          ptr + 4,
          [
            "Afak",
            "Arab",
            "Armi",
            "Armn",
            "Avst",
            "Bali",
            "Bamu",
            "Bass",
            "Batk",
            "Beng",
            "Blis",
            "Bopo",
            "Brah",
            "Brai",
            "Bugi",
            "Buhd",
            "Cakm",
            "Cans",
            "Cari",
            "Cham",
            "Cher",
            "Cirt",
            "Copt",
            "Cprt",
            "Cyrl",
            "Cyrs",
            "Deva",
            "Dsrt",
            "Dupl",
            "Egyd",
            "Egyh",
            "Egyp",
            "Elba",
            "Ethi",
            "Geor",
            "Geok",
            "Glag",
            "Goth",
            "Gran",
            "Grek",
            "Gujr",
            "Guru",
            "Hang",
            "Hani",
            "Hano",
            "Hans",
            "Hant",
            "Hebr",
            "Hluw",
            "Hmng",
            "Hung",
            "Inds",
            "Ital",
            "Java",
            "Jpan",
            "Jurc",
            "Kali",
            "Khar",
            "Khmr",
            "Khoj",
            "Knda",
            "Kpel",
            "Kthi",
            "Lana",
            "Laoo",
            "Latf",
            "Latg",
            "Latn",
            "Lepc",
            "Limb",
            "Lina",
            "Linb",
            "Lisu",
            "Loma",
            "Lyci",
            "Lydi",
            "Mand",
            "Mani",
            "Maya",
            "Mend",
            "Merc",
            "Mero",
            "Mlym",
            "Moon",
            "Mong",
            "Mroo",
            "Mtei",
            "Mymr",
            "Narb",
            "Nbat",
            "Nkgb",
            "Nkoo",
            "Nshu",
            "Ogam",
            "Olck",
            "Orkh",
            "Orya",
            "Osma",
            "Palm",
            "Perm",
            "Phag",
            "Phli",
            "Phlp",
            "Phlv",
            "Phnx",
            "Plrd",
            "Prti",
            "Rjng",
            "Roro",
            "Runr",
            "Samr",
            "Sara",
            "Sarb",
            "Saur",
            "Sgnw",
            "Shaw",
            "Shrd",
            "Sind",
            "Sinh",
            "Sora",
            "Sund",
            "Sylo",
            "Syrc",
            "Syre",
            "Syrj",
            "Syrn",
            "Tagb",
            "Takr",
            "Tale",
            "Talu",
            "Taml",
            "Tang",
            "Tavt",
            "Telu",
            "Teng",
            "Tfng",
            "Tglg",
            "Thaa",
            "Thai",
            "Tibt",
            "Tirh",
            "Ugar",
            "Vaii",
            "Visp",
            "Wara",
            "Wole",
            "Xpeo",
            "Xsux",
            "Yiii",
            "Zmth",
            "Zsym",
            "Zyyy",
          ].indexOf(x["script"] as string)
        );
      }
    },
    "load_GetFontArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["genericFamily"] = A.load.Enum(ptr + 0, [
        "standard",
        "sansserif",
        "serif",
        "fixed",
        "cursive",
        "fantasy",
        "math",
      ]);
      x["script"] = A.load.Enum(ptr + 4, [
        "Afak",
        "Arab",
        "Armi",
        "Armn",
        "Avst",
        "Bali",
        "Bamu",
        "Bass",
        "Batk",
        "Beng",
        "Blis",
        "Bopo",
        "Brah",
        "Brai",
        "Bugi",
        "Buhd",
        "Cakm",
        "Cans",
        "Cari",
        "Cham",
        "Cher",
        "Cirt",
        "Copt",
        "Cprt",
        "Cyrl",
        "Cyrs",
        "Deva",
        "Dsrt",
        "Dupl",
        "Egyd",
        "Egyh",
        "Egyp",
        "Elba",
        "Ethi",
        "Geor",
        "Geok",
        "Glag",
        "Goth",
        "Gran",
        "Grek",
        "Gujr",
        "Guru",
        "Hang",
        "Hani",
        "Hano",
        "Hans",
        "Hant",
        "Hebr",
        "Hluw",
        "Hmng",
        "Hung",
        "Inds",
        "Ital",
        "Java",
        "Jpan",
        "Jurc",
        "Kali",
        "Khar",
        "Khmr",
        "Khoj",
        "Knda",
        "Kpel",
        "Kthi",
        "Lana",
        "Laoo",
        "Latf",
        "Latg",
        "Latn",
        "Lepc",
        "Limb",
        "Lina",
        "Linb",
        "Lisu",
        "Loma",
        "Lyci",
        "Lydi",
        "Mand",
        "Mani",
        "Maya",
        "Mend",
        "Merc",
        "Mero",
        "Mlym",
        "Moon",
        "Mong",
        "Mroo",
        "Mtei",
        "Mymr",
        "Narb",
        "Nbat",
        "Nkgb",
        "Nkoo",
        "Nshu",
        "Ogam",
        "Olck",
        "Orkh",
        "Orya",
        "Osma",
        "Palm",
        "Perm",
        "Phag",
        "Phli",
        "Phlp",
        "Phlv",
        "Phnx",
        "Plrd",
        "Prti",
        "Rjng",
        "Roro",
        "Runr",
        "Samr",
        "Sara",
        "Sarb",
        "Saur",
        "Sgnw",
        "Shaw",
        "Shrd",
        "Sind",
        "Sinh",
        "Sora",
        "Sund",
        "Sylo",
        "Syrc",
        "Syre",
        "Syrj",
        "Syrn",
        "Tagb",
        "Takr",
        "Tale",
        "Talu",
        "Taml",
        "Tang",
        "Tavt",
        "Telu",
        "Teng",
        "Tfng",
        "Tglg",
        "Thaa",
        "Thai",
        "Tibt",
        "Tirh",
        "Ugar",
        "Vaii",
        "Visp",
        "Wara",
        "Wole",
        "Xpeo",
        "Xsux",
        "Yiii",
        "Zmth",
        "Zsym",
        "Zyyy",
      ]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetFontReturnType": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["fontId"]);
        A.store.Enum(
          ptr + 4,
          [
            "not_controllable",
            "controlled_by_other_extensions",
            "controllable_by_this_extension",
            "controlled_by_this_extension",
          ].indexOf(x["levelOfControl"] as string)
        );
      }
    },
    "load_GetFontReturnType": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fontId"] = A.load.Ref(ptr + 0, undefined);
      x["levelOfControl"] = A.load.Enum(ptr + 4, [
        "not_controllable",
        "controlled_by_other_extensions",
        "controllable_by_this_extension",
        "controlled_by_this_extension",
      ]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetMinimumFontSizeArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 0, true);
      }
    },
    "load_GetMinimumFontSizeArgDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetMinimumFontSizeReturnType": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Enum(
          ptr + 0,
          [
            "not_controllable",
            "controlled_by_other_extensions",
            "controllable_by_this_extension",
            "controlled_by_this_extension",
          ].indexOf(x["levelOfControl"] as string)
        );
        A.store.Int64(ptr + 8, x["pixelSize"] === undefined ? 0 : (x["pixelSize"] as number));
      }
    },
    "load_GetMinimumFontSizeReturnType": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["levelOfControl"] = A.load.Enum(ptr + 0, [
        "not_controllable",
        "controlled_by_other_extensions",
        "controllable_by_this_extension",
        "controlled_by_this_extension",
      ]);
      x["pixelSize"] = A.load.Int64(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnDefaultFixedFontSizeChangedArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Enum(
          ptr + 0,
          [
            "not_controllable",
            "controlled_by_other_extensions",
            "controllable_by_this_extension",
            "controlled_by_this_extension",
          ].indexOf(x["levelOfControl"] as string)
        );
        A.store.Int64(ptr + 8, x["pixelSize"] === undefined ? 0 : (x["pixelSize"] as number));
      }
    },
    "load_OnDefaultFixedFontSizeChangedArgDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["levelOfControl"] = A.load.Enum(ptr + 0, [
        "not_controllable",
        "controlled_by_other_extensions",
        "controllable_by_this_extension",
        "controlled_by_this_extension",
      ]);
      x["pixelSize"] = A.load.Int64(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnDefaultFontSizeChangedArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Enum(
          ptr + 0,
          [
            "not_controllable",
            "controlled_by_other_extensions",
            "controllable_by_this_extension",
            "controlled_by_this_extension",
          ].indexOf(x["levelOfControl"] as string)
        );
        A.store.Int64(ptr + 8, x["pixelSize"] === undefined ? 0 : (x["pixelSize"] as number));
      }
    },
    "load_OnDefaultFontSizeChangedArgDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["levelOfControl"] = A.load.Enum(ptr + 0, [
        "not_controllable",
        "controlled_by_other_extensions",
        "controllable_by_this_extension",
        "controlled_by_this_extension",
      ]);
      x["pixelSize"] = A.load.Int64(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnFontChangedArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Enum(ptr + 8, -1);
        A.store.Enum(ptr + 12, -1);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Ref(ptr + 0, x["fontId"]);
        A.store.Enum(
          ptr + 4,
          ["standard", "sansserif", "serif", "fixed", "cursive", "fantasy", "math"].indexOf(
            x["genericFamily"] as string
          )
        );
        A.store.Enum(
          ptr + 8,
          [
            "not_controllable",
            "controlled_by_other_extensions",
            "controllable_by_this_extension",
            "controlled_by_this_extension",
          ].indexOf(x["levelOfControl"] as string)
        );
        A.store.Enum(
          ptr + 12,
          [
            "Afak",
            "Arab",
            "Armi",
            "Armn",
            "Avst",
            "Bali",
            "Bamu",
            "Bass",
            "Batk",
            "Beng",
            "Blis",
            "Bopo",
            "Brah",
            "Brai",
            "Bugi",
            "Buhd",
            "Cakm",
            "Cans",
            "Cari",
            "Cham",
            "Cher",
            "Cirt",
            "Copt",
            "Cprt",
            "Cyrl",
            "Cyrs",
            "Deva",
            "Dsrt",
            "Dupl",
            "Egyd",
            "Egyh",
            "Egyp",
            "Elba",
            "Ethi",
            "Geor",
            "Geok",
            "Glag",
            "Goth",
            "Gran",
            "Grek",
            "Gujr",
            "Guru",
            "Hang",
            "Hani",
            "Hano",
            "Hans",
            "Hant",
            "Hebr",
            "Hluw",
            "Hmng",
            "Hung",
            "Inds",
            "Ital",
            "Java",
            "Jpan",
            "Jurc",
            "Kali",
            "Khar",
            "Khmr",
            "Khoj",
            "Knda",
            "Kpel",
            "Kthi",
            "Lana",
            "Laoo",
            "Latf",
            "Latg",
            "Latn",
            "Lepc",
            "Limb",
            "Lina",
            "Linb",
            "Lisu",
            "Loma",
            "Lyci",
            "Lydi",
            "Mand",
            "Mani",
            "Maya",
            "Mend",
            "Merc",
            "Mero",
            "Mlym",
            "Moon",
            "Mong",
            "Mroo",
            "Mtei",
            "Mymr",
            "Narb",
            "Nbat",
            "Nkgb",
            "Nkoo",
            "Nshu",
            "Ogam",
            "Olck",
            "Orkh",
            "Orya",
            "Osma",
            "Palm",
            "Perm",
            "Phag",
            "Phli",
            "Phlp",
            "Phlv",
            "Phnx",
            "Plrd",
            "Prti",
            "Rjng",
            "Roro",
            "Runr",
            "Samr",
            "Sara",
            "Sarb",
            "Saur",
            "Sgnw",
            "Shaw",
            "Shrd",
            "Sind",
            "Sinh",
            "Sora",
            "Sund",
            "Sylo",
            "Syrc",
            "Syre",
            "Syrj",
            "Syrn",
            "Tagb",
            "Takr",
            "Tale",
            "Talu",
            "Taml",
            "Tang",
            "Tavt",
            "Telu",
            "Teng",
            "Tfng",
            "Tglg",
            "Thaa",
            "Thai",
            "Tibt",
            "Tirh",
            "Ugar",
            "Vaii",
            "Visp",
            "Wara",
            "Wole",
            "Xpeo",
            "Xsux",
            "Yiii",
            "Zmth",
            "Zsym",
            "Zyyy",
          ].indexOf(x["script"] as string)
        );
      }
    },
    "load_OnFontChangedArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fontId"] = A.load.Ref(ptr + 0, undefined);
      x["genericFamily"] = A.load.Enum(ptr + 4, [
        "standard",
        "sansserif",
        "serif",
        "fixed",
        "cursive",
        "fantasy",
        "math",
      ]);
      x["levelOfControl"] = A.load.Enum(ptr + 8, [
        "not_controllable",
        "controlled_by_other_extensions",
        "controllable_by_this_extension",
        "controlled_by_this_extension",
      ]);
      x["script"] = A.load.Enum(ptr + 12, [
        "Afak",
        "Arab",
        "Armi",
        "Armn",
        "Avst",
        "Bali",
        "Bamu",
        "Bass",
        "Batk",
        "Beng",
        "Blis",
        "Bopo",
        "Brah",
        "Brai",
        "Bugi",
        "Buhd",
        "Cakm",
        "Cans",
        "Cari",
        "Cham",
        "Cher",
        "Cirt",
        "Copt",
        "Cprt",
        "Cyrl",
        "Cyrs",
        "Deva",
        "Dsrt",
        "Dupl",
        "Egyd",
        "Egyh",
        "Egyp",
        "Elba",
        "Ethi",
        "Geor",
        "Geok",
        "Glag",
        "Goth",
        "Gran",
        "Grek",
        "Gujr",
        "Guru",
        "Hang",
        "Hani",
        "Hano",
        "Hans",
        "Hant",
        "Hebr",
        "Hluw",
        "Hmng",
        "Hung",
        "Inds",
        "Ital",
        "Java",
        "Jpan",
        "Jurc",
        "Kali",
        "Khar",
        "Khmr",
        "Khoj",
        "Knda",
        "Kpel",
        "Kthi",
        "Lana",
        "Laoo",
        "Latf",
        "Latg",
        "Latn",
        "Lepc",
        "Limb",
        "Lina",
        "Linb",
        "Lisu",
        "Loma",
        "Lyci",
        "Lydi",
        "Mand",
        "Mani",
        "Maya",
        "Mend",
        "Merc",
        "Mero",
        "Mlym",
        "Moon",
        "Mong",
        "Mroo",
        "Mtei",
        "Mymr",
        "Narb",
        "Nbat",
        "Nkgb",
        "Nkoo",
        "Nshu",
        "Ogam",
        "Olck",
        "Orkh",
        "Orya",
        "Osma",
        "Palm",
        "Perm",
        "Phag",
        "Phli",
        "Phlp",
        "Phlv",
        "Phnx",
        "Plrd",
        "Prti",
        "Rjng",
        "Roro",
        "Runr",
        "Samr",
        "Sara",
        "Sarb",
        "Saur",
        "Sgnw",
        "Shaw",
        "Shrd",
        "Sind",
        "Sinh",
        "Sora",
        "Sund",
        "Sylo",
        "Syrc",
        "Syre",
        "Syrj",
        "Syrn",
        "Tagb",
        "Takr",
        "Tale",
        "Talu",
        "Taml",
        "Tang",
        "Tavt",
        "Telu",
        "Teng",
        "Tfng",
        "Tglg",
        "Thaa",
        "Thai",
        "Tibt",
        "Tirh",
        "Ugar",
        "Vaii",
        "Visp",
        "Wara",
        "Wole",
        "Xpeo",
        "Xsux",
        "Yiii",
        "Zmth",
        "Zsym",
        "Zyyy",
      ]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnMinimumFontSizeChangedArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Enum(
          ptr + 0,
          [
            "not_controllable",
            "controlled_by_other_extensions",
            "controllable_by_this_extension",
            "controlled_by_this_extension",
          ].indexOf(x["levelOfControl"] as string)
        );
        A.store.Int64(ptr + 8, x["pixelSize"] === undefined ? 0 : (x["pixelSize"] as number));
      }
    },
    "load_OnMinimumFontSizeChangedArgDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["levelOfControl"] = A.load.Enum(ptr + 0, [
        "not_controllable",
        "controlled_by_other_extensions",
        "controllable_by_this_extension",
        "controlled_by_this_extension",
      ]);
      x["pixelSize"] = A.load.Int64(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetDefaultFixedFontSizeArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Int64(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Int64(ptr + 0, x["pixelSize"] === undefined ? 0 : (x["pixelSize"] as number));
      }
    },
    "load_SetDefaultFixedFontSizeArgDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["pixelSize"] = A.load.Int64(ptr + 0);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetDefaultFontSizeArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Int64(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Int64(ptr + 0, x["pixelSize"] === undefined ? 0 : (x["pixelSize"] as number));
      }
    },
    "load_SetDefaultFontSizeArgDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["pixelSize"] = A.load.Int64(ptr + 0);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetFontArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Enum(ptr + 8, -1);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["fontId"]);
        A.store.Enum(
          ptr + 4,
          ["standard", "sansserif", "serif", "fixed", "cursive", "fantasy", "math"].indexOf(
            x["genericFamily"] as string
          )
        );
        A.store.Enum(
          ptr + 8,
          [
            "Afak",
            "Arab",
            "Armi",
            "Armn",
            "Avst",
            "Bali",
            "Bamu",
            "Bass",
            "Batk",
            "Beng",
            "Blis",
            "Bopo",
            "Brah",
            "Brai",
            "Bugi",
            "Buhd",
            "Cakm",
            "Cans",
            "Cari",
            "Cham",
            "Cher",
            "Cirt",
            "Copt",
            "Cprt",
            "Cyrl",
            "Cyrs",
            "Deva",
            "Dsrt",
            "Dupl",
            "Egyd",
            "Egyh",
            "Egyp",
            "Elba",
            "Ethi",
            "Geor",
            "Geok",
            "Glag",
            "Goth",
            "Gran",
            "Grek",
            "Gujr",
            "Guru",
            "Hang",
            "Hani",
            "Hano",
            "Hans",
            "Hant",
            "Hebr",
            "Hluw",
            "Hmng",
            "Hung",
            "Inds",
            "Ital",
            "Java",
            "Jpan",
            "Jurc",
            "Kali",
            "Khar",
            "Khmr",
            "Khoj",
            "Knda",
            "Kpel",
            "Kthi",
            "Lana",
            "Laoo",
            "Latf",
            "Latg",
            "Latn",
            "Lepc",
            "Limb",
            "Lina",
            "Linb",
            "Lisu",
            "Loma",
            "Lyci",
            "Lydi",
            "Mand",
            "Mani",
            "Maya",
            "Mend",
            "Merc",
            "Mero",
            "Mlym",
            "Moon",
            "Mong",
            "Mroo",
            "Mtei",
            "Mymr",
            "Narb",
            "Nbat",
            "Nkgb",
            "Nkoo",
            "Nshu",
            "Ogam",
            "Olck",
            "Orkh",
            "Orya",
            "Osma",
            "Palm",
            "Perm",
            "Phag",
            "Phli",
            "Phlp",
            "Phlv",
            "Phnx",
            "Plrd",
            "Prti",
            "Rjng",
            "Roro",
            "Runr",
            "Samr",
            "Sara",
            "Sarb",
            "Saur",
            "Sgnw",
            "Shaw",
            "Shrd",
            "Sind",
            "Sinh",
            "Sora",
            "Sund",
            "Sylo",
            "Syrc",
            "Syre",
            "Syrj",
            "Syrn",
            "Tagb",
            "Takr",
            "Tale",
            "Talu",
            "Taml",
            "Tang",
            "Tavt",
            "Telu",
            "Teng",
            "Tfng",
            "Tglg",
            "Thaa",
            "Thai",
            "Tibt",
            "Tirh",
            "Ugar",
            "Vaii",
            "Visp",
            "Wara",
            "Wole",
            "Xpeo",
            "Xsux",
            "Yiii",
            "Zmth",
            "Zsym",
            "Zyyy",
          ].indexOf(x["script"] as string)
        );
      }
    },
    "load_SetFontArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fontId"] = A.load.Ref(ptr + 0, undefined);
      x["genericFamily"] = A.load.Enum(ptr + 4, [
        "standard",
        "sansserif",
        "serif",
        "fixed",
        "cursive",
        "fantasy",
        "math",
      ]);
      x["script"] = A.load.Enum(ptr + 8, [
        "Afak",
        "Arab",
        "Armi",
        "Armn",
        "Avst",
        "Bali",
        "Bamu",
        "Bass",
        "Batk",
        "Beng",
        "Blis",
        "Bopo",
        "Brah",
        "Brai",
        "Bugi",
        "Buhd",
        "Cakm",
        "Cans",
        "Cari",
        "Cham",
        "Cher",
        "Cirt",
        "Copt",
        "Cprt",
        "Cyrl",
        "Cyrs",
        "Deva",
        "Dsrt",
        "Dupl",
        "Egyd",
        "Egyh",
        "Egyp",
        "Elba",
        "Ethi",
        "Geor",
        "Geok",
        "Glag",
        "Goth",
        "Gran",
        "Grek",
        "Gujr",
        "Guru",
        "Hang",
        "Hani",
        "Hano",
        "Hans",
        "Hant",
        "Hebr",
        "Hluw",
        "Hmng",
        "Hung",
        "Inds",
        "Ital",
        "Java",
        "Jpan",
        "Jurc",
        "Kali",
        "Khar",
        "Khmr",
        "Khoj",
        "Knda",
        "Kpel",
        "Kthi",
        "Lana",
        "Laoo",
        "Latf",
        "Latg",
        "Latn",
        "Lepc",
        "Limb",
        "Lina",
        "Linb",
        "Lisu",
        "Loma",
        "Lyci",
        "Lydi",
        "Mand",
        "Mani",
        "Maya",
        "Mend",
        "Merc",
        "Mero",
        "Mlym",
        "Moon",
        "Mong",
        "Mroo",
        "Mtei",
        "Mymr",
        "Narb",
        "Nbat",
        "Nkgb",
        "Nkoo",
        "Nshu",
        "Ogam",
        "Olck",
        "Orkh",
        "Orya",
        "Osma",
        "Palm",
        "Perm",
        "Phag",
        "Phli",
        "Phlp",
        "Phlv",
        "Phnx",
        "Plrd",
        "Prti",
        "Rjng",
        "Roro",
        "Runr",
        "Samr",
        "Sara",
        "Sarb",
        "Saur",
        "Sgnw",
        "Shaw",
        "Shrd",
        "Sind",
        "Sinh",
        "Sora",
        "Sund",
        "Sylo",
        "Syrc",
        "Syre",
        "Syrj",
        "Syrn",
        "Tagb",
        "Takr",
        "Tale",
        "Talu",
        "Taml",
        "Tang",
        "Tavt",
        "Telu",
        "Teng",
        "Tfng",
        "Tglg",
        "Thaa",
        "Thai",
        "Tibt",
        "Tirh",
        "Ugar",
        "Vaii",
        "Visp",
        "Wara",
        "Wole",
        "Xpeo",
        "Xsux",
        "Yiii",
        "Zmth",
        "Zsym",
        "Zyyy",
      ]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetMinimumFontSizeArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Int64(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Int64(ptr + 0, x["pixelSize"] === undefined ? 0 : (x["pixelSize"] as number));
      }
    },
    "load_SetMinimumFontSizeArgDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["pixelSize"] = A.load.Int64(ptr + 0);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_ClearDefaultFixedFontSize": (): heap.Ref<boolean> => {
      if (WEBEXT?.fontSettings && "clearDefaultFixedFontSize" in WEBEXT?.fontSettings) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ClearDefaultFixedFontSize": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.clearDefaultFixedFontSize);
    },
    "call_ClearDefaultFixedFontSize": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      const _ret = WEBEXT.fontSettings.clearDefaultFixedFontSize(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ClearDefaultFixedFontSize": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        const _ret = WEBEXT.fontSettings.clearDefaultFixedFontSize(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ClearDefaultFontSize": (): heap.Ref<boolean> => {
      if (WEBEXT?.fontSettings && "clearDefaultFontSize" in WEBEXT?.fontSettings) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ClearDefaultFontSize": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.clearDefaultFontSize);
    },
    "call_ClearDefaultFontSize": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      const _ret = WEBEXT.fontSettings.clearDefaultFontSize(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ClearDefaultFontSize": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        const _ret = WEBEXT.fontSettings.clearDefaultFontSize(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ClearFont": (): heap.Ref<boolean> => {
      if (WEBEXT?.fontSettings && "clearFont" in WEBEXT?.fontSettings) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ClearFont": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.clearFont);
    },
    "call_ClearFont": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["genericFamily"] = A.load.Enum(details + 0, [
        "standard",
        "sansserif",
        "serif",
        "fixed",
        "cursive",
        "fantasy",
        "math",
      ]);
      details_ffi["script"] = A.load.Enum(details + 4, [
        "Afak",
        "Arab",
        "Armi",
        "Armn",
        "Avst",
        "Bali",
        "Bamu",
        "Bass",
        "Batk",
        "Beng",
        "Blis",
        "Bopo",
        "Brah",
        "Brai",
        "Bugi",
        "Buhd",
        "Cakm",
        "Cans",
        "Cari",
        "Cham",
        "Cher",
        "Cirt",
        "Copt",
        "Cprt",
        "Cyrl",
        "Cyrs",
        "Deva",
        "Dsrt",
        "Dupl",
        "Egyd",
        "Egyh",
        "Egyp",
        "Elba",
        "Ethi",
        "Geor",
        "Geok",
        "Glag",
        "Goth",
        "Gran",
        "Grek",
        "Gujr",
        "Guru",
        "Hang",
        "Hani",
        "Hano",
        "Hans",
        "Hant",
        "Hebr",
        "Hluw",
        "Hmng",
        "Hung",
        "Inds",
        "Ital",
        "Java",
        "Jpan",
        "Jurc",
        "Kali",
        "Khar",
        "Khmr",
        "Khoj",
        "Knda",
        "Kpel",
        "Kthi",
        "Lana",
        "Laoo",
        "Latf",
        "Latg",
        "Latn",
        "Lepc",
        "Limb",
        "Lina",
        "Linb",
        "Lisu",
        "Loma",
        "Lyci",
        "Lydi",
        "Mand",
        "Mani",
        "Maya",
        "Mend",
        "Merc",
        "Mero",
        "Mlym",
        "Moon",
        "Mong",
        "Mroo",
        "Mtei",
        "Mymr",
        "Narb",
        "Nbat",
        "Nkgb",
        "Nkoo",
        "Nshu",
        "Ogam",
        "Olck",
        "Orkh",
        "Orya",
        "Osma",
        "Palm",
        "Perm",
        "Phag",
        "Phli",
        "Phlp",
        "Phlv",
        "Phnx",
        "Plrd",
        "Prti",
        "Rjng",
        "Roro",
        "Runr",
        "Samr",
        "Sara",
        "Sarb",
        "Saur",
        "Sgnw",
        "Shaw",
        "Shrd",
        "Sind",
        "Sinh",
        "Sora",
        "Sund",
        "Sylo",
        "Syrc",
        "Syre",
        "Syrj",
        "Syrn",
        "Tagb",
        "Takr",
        "Tale",
        "Talu",
        "Taml",
        "Tang",
        "Tavt",
        "Telu",
        "Teng",
        "Tfng",
        "Tglg",
        "Thaa",
        "Thai",
        "Tibt",
        "Tirh",
        "Ugar",
        "Vaii",
        "Visp",
        "Wara",
        "Wole",
        "Xpeo",
        "Xsux",
        "Yiii",
        "Zmth",
        "Zsym",
        "Zyyy",
      ]);

      const _ret = WEBEXT.fontSettings.clearFont(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ClearFont": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["genericFamily"] = A.load.Enum(details + 0, [
          "standard",
          "sansserif",
          "serif",
          "fixed",
          "cursive",
          "fantasy",
          "math",
        ]);
        details_ffi["script"] = A.load.Enum(details + 4, [
          "Afak",
          "Arab",
          "Armi",
          "Armn",
          "Avst",
          "Bali",
          "Bamu",
          "Bass",
          "Batk",
          "Beng",
          "Blis",
          "Bopo",
          "Brah",
          "Brai",
          "Bugi",
          "Buhd",
          "Cakm",
          "Cans",
          "Cari",
          "Cham",
          "Cher",
          "Cirt",
          "Copt",
          "Cprt",
          "Cyrl",
          "Cyrs",
          "Deva",
          "Dsrt",
          "Dupl",
          "Egyd",
          "Egyh",
          "Egyp",
          "Elba",
          "Ethi",
          "Geor",
          "Geok",
          "Glag",
          "Goth",
          "Gran",
          "Grek",
          "Gujr",
          "Guru",
          "Hang",
          "Hani",
          "Hano",
          "Hans",
          "Hant",
          "Hebr",
          "Hluw",
          "Hmng",
          "Hung",
          "Inds",
          "Ital",
          "Java",
          "Jpan",
          "Jurc",
          "Kali",
          "Khar",
          "Khmr",
          "Khoj",
          "Knda",
          "Kpel",
          "Kthi",
          "Lana",
          "Laoo",
          "Latf",
          "Latg",
          "Latn",
          "Lepc",
          "Limb",
          "Lina",
          "Linb",
          "Lisu",
          "Loma",
          "Lyci",
          "Lydi",
          "Mand",
          "Mani",
          "Maya",
          "Mend",
          "Merc",
          "Mero",
          "Mlym",
          "Moon",
          "Mong",
          "Mroo",
          "Mtei",
          "Mymr",
          "Narb",
          "Nbat",
          "Nkgb",
          "Nkoo",
          "Nshu",
          "Ogam",
          "Olck",
          "Orkh",
          "Orya",
          "Osma",
          "Palm",
          "Perm",
          "Phag",
          "Phli",
          "Phlp",
          "Phlv",
          "Phnx",
          "Plrd",
          "Prti",
          "Rjng",
          "Roro",
          "Runr",
          "Samr",
          "Sara",
          "Sarb",
          "Saur",
          "Sgnw",
          "Shaw",
          "Shrd",
          "Sind",
          "Sinh",
          "Sora",
          "Sund",
          "Sylo",
          "Syrc",
          "Syre",
          "Syrj",
          "Syrn",
          "Tagb",
          "Takr",
          "Tale",
          "Talu",
          "Taml",
          "Tang",
          "Tavt",
          "Telu",
          "Teng",
          "Tfng",
          "Tglg",
          "Thaa",
          "Thai",
          "Tibt",
          "Tirh",
          "Ugar",
          "Vaii",
          "Visp",
          "Wara",
          "Wole",
          "Xpeo",
          "Xsux",
          "Yiii",
          "Zmth",
          "Zsym",
          "Zyyy",
        ]);

        const _ret = WEBEXT.fontSettings.clearFont(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ClearMinimumFontSize": (): heap.Ref<boolean> => {
      if (WEBEXT?.fontSettings && "clearMinimumFontSize" in WEBEXT?.fontSettings) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ClearMinimumFontSize": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.clearMinimumFontSize);
    },
    "call_ClearMinimumFontSize": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      const _ret = WEBEXT.fontSettings.clearMinimumFontSize(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ClearMinimumFontSize": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        const _ret = WEBEXT.fontSettings.clearMinimumFontSize(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDefaultFixedFontSize": (): heap.Ref<boolean> => {
      if (WEBEXT?.fontSettings && "getDefaultFixedFontSize" in WEBEXT?.fontSettings) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDefaultFixedFontSize": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.getDefaultFixedFontSize);
    },
    "call_GetDefaultFixedFontSize": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      const _ret = WEBEXT.fontSettings.getDefaultFixedFontSize(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDefaultFixedFontSize": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        const _ret = WEBEXT.fontSettings.getDefaultFixedFontSize(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDefaultFontSize": (): heap.Ref<boolean> => {
      if (WEBEXT?.fontSettings && "getDefaultFontSize" in WEBEXT?.fontSettings) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDefaultFontSize": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.getDefaultFontSize);
    },
    "call_GetDefaultFontSize": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      const _ret = WEBEXT.fontSettings.getDefaultFontSize(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDefaultFontSize": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        const _ret = WEBEXT.fontSettings.getDefaultFontSize(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetFont": (): heap.Ref<boolean> => {
      if (WEBEXT?.fontSettings && "getFont" in WEBEXT?.fontSettings) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetFont": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.getFont);
    },
    "call_GetFont": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["genericFamily"] = A.load.Enum(details + 0, [
        "standard",
        "sansserif",
        "serif",
        "fixed",
        "cursive",
        "fantasy",
        "math",
      ]);
      details_ffi["script"] = A.load.Enum(details + 4, [
        "Afak",
        "Arab",
        "Armi",
        "Armn",
        "Avst",
        "Bali",
        "Bamu",
        "Bass",
        "Batk",
        "Beng",
        "Blis",
        "Bopo",
        "Brah",
        "Brai",
        "Bugi",
        "Buhd",
        "Cakm",
        "Cans",
        "Cari",
        "Cham",
        "Cher",
        "Cirt",
        "Copt",
        "Cprt",
        "Cyrl",
        "Cyrs",
        "Deva",
        "Dsrt",
        "Dupl",
        "Egyd",
        "Egyh",
        "Egyp",
        "Elba",
        "Ethi",
        "Geor",
        "Geok",
        "Glag",
        "Goth",
        "Gran",
        "Grek",
        "Gujr",
        "Guru",
        "Hang",
        "Hani",
        "Hano",
        "Hans",
        "Hant",
        "Hebr",
        "Hluw",
        "Hmng",
        "Hung",
        "Inds",
        "Ital",
        "Java",
        "Jpan",
        "Jurc",
        "Kali",
        "Khar",
        "Khmr",
        "Khoj",
        "Knda",
        "Kpel",
        "Kthi",
        "Lana",
        "Laoo",
        "Latf",
        "Latg",
        "Latn",
        "Lepc",
        "Limb",
        "Lina",
        "Linb",
        "Lisu",
        "Loma",
        "Lyci",
        "Lydi",
        "Mand",
        "Mani",
        "Maya",
        "Mend",
        "Merc",
        "Mero",
        "Mlym",
        "Moon",
        "Mong",
        "Mroo",
        "Mtei",
        "Mymr",
        "Narb",
        "Nbat",
        "Nkgb",
        "Nkoo",
        "Nshu",
        "Ogam",
        "Olck",
        "Orkh",
        "Orya",
        "Osma",
        "Palm",
        "Perm",
        "Phag",
        "Phli",
        "Phlp",
        "Phlv",
        "Phnx",
        "Plrd",
        "Prti",
        "Rjng",
        "Roro",
        "Runr",
        "Samr",
        "Sara",
        "Sarb",
        "Saur",
        "Sgnw",
        "Shaw",
        "Shrd",
        "Sind",
        "Sinh",
        "Sora",
        "Sund",
        "Sylo",
        "Syrc",
        "Syre",
        "Syrj",
        "Syrn",
        "Tagb",
        "Takr",
        "Tale",
        "Talu",
        "Taml",
        "Tang",
        "Tavt",
        "Telu",
        "Teng",
        "Tfng",
        "Tglg",
        "Thaa",
        "Thai",
        "Tibt",
        "Tirh",
        "Ugar",
        "Vaii",
        "Visp",
        "Wara",
        "Wole",
        "Xpeo",
        "Xsux",
        "Yiii",
        "Zmth",
        "Zsym",
        "Zyyy",
      ]);

      const _ret = WEBEXT.fontSettings.getFont(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetFont": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["genericFamily"] = A.load.Enum(details + 0, [
          "standard",
          "sansserif",
          "serif",
          "fixed",
          "cursive",
          "fantasy",
          "math",
        ]);
        details_ffi["script"] = A.load.Enum(details + 4, [
          "Afak",
          "Arab",
          "Armi",
          "Armn",
          "Avst",
          "Bali",
          "Bamu",
          "Bass",
          "Batk",
          "Beng",
          "Blis",
          "Bopo",
          "Brah",
          "Brai",
          "Bugi",
          "Buhd",
          "Cakm",
          "Cans",
          "Cari",
          "Cham",
          "Cher",
          "Cirt",
          "Copt",
          "Cprt",
          "Cyrl",
          "Cyrs",
          "Deva",
          "Dsrt",
          "Dupl",
          "Egyd",
          "Egyh",
          "Egyp",
          "Elba",
          "Ethi",
          "Geor",
          "Geok",
          "Glag",
          "Goth",
          "Gran",
          "Grek",
          "Gujr",
          "Guru",
          "Hang",
          "Hani",
          "Hano",
          "Hans",
          "Hant",
          "Hebr",
          "Hluw",
          "Hmng",
          "Hung",
          "Inds",
          "Ital",
          "Java",
          "Jpan",
          "Jurc",
          "Kali",
          "Khar",
          "Khmr",
          "Khoj",
          "Knda",
          "Kpel",
          "Kthi",
          "Lana",
          "Laoo",
          "Latf",
          "Latg",
          "Latn",
          "Lepc",
          "Limb",
          "Lina",
          "Linb",
          "Lisu",
          "Loma",
          "Lyci",
          "Lydi",
          "Mand",
          "Mani",
          "Maya",
          "Mend",
          "Merc",
          "Mero",
          "Mlym",
          "Moon",
          "Mong",
          "Mroo",
          "Mtei",
          "Mymr",
          "Narb",
          "Nbat",
          "Nkgb",
          "Nkoo",
          "Nshu",
          "Ogam",
          "Olck",
          "Orkh",
          "Orya",
          "Osma",
          "Palm",
          "Perm",
          "Phag",
          "Phli",
          "Phlp",
          "Phlv",
          "Phnx",
          "Plrd",
          "Prti",
          "Rjng",
          "Roro",
          "Runr",
          "Samr",
          "Sara",
          "Sarb",
          "Saur",
          "Sgnw",
          "Shaw",
          "Shrd",
          "Sind",
          "Sinh",
          "Sora",
          "Sund",
          "Sylo",
          "Syrc",
          "Syre",
          "Syrj",
          "Syrn",
          "Tagb",
          "Takr",
          "Tale",
          "Talu",
          "Taml",
          "Tang",
          "Tavt",
          "Telu",
          "Teng",
          "Tfng",
          "Tglg",
          "Thaa",
          "Thai",
          "Tibt",
          "Tirh",
          "Ugar",
          "Vaii",
          "Visp",
          "Wara",
          "Wole",
          "Xpeo",
          "Xsux",
          "Yiii",
          "Zmth",
          "Zsym",
          "Zyyy",
        ]);

        const _ret = WEBEXT.fontSettings.getFont(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetFontList": (): heap.Ref<boolean> => {
      if (WEBEXT?.fontSettings && "getFontList" in WEBEXT?.fontSettings) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetFontList": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.getFontList);
    },
    "call_GetFontList": (retPtr: Pointer): void => {
      const _ret = WEBEXT.fontSettings.getFontList();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetFontList": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fontSettings.getFontList();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetMinimumFontSize": (): heap.Ref<boolean> => {
      if (WEBEXT?.fontSettings && "getMinimumFontSize" in WEBEXT?.fontSettings) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetMinimumFontSize": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.getMinimumFontSize);
    },
    "call_GetMinimumFontSize": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      const _ret = WEBEXT.fontSettings.getMinimumFontSize(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetMinimumFontSize": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        const _ret = WEBEXT.fontSettings.getMinimumFontSize(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDefaultFixedFontSizeChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fontSettings?.onDefaultFixedFontSizeChanged &&
        "addListener" in WEBEXT?.fontSettings?.onDefaultFixedFontSizeChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDefaultFixedFontSizeChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.onDefaultFixedFontSizeChanged.addListener);
    },
    "call_OnDefaultFixedFontSizeChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fontSettings.onDefaultFixedFontSizeChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnDefaultFixedFontSizeChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fontSettings.onDefaultFixedFontSizeChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDefaultFixedFontSizeChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fontSettings?.onDefaultFixedFontSizeChanged &&
        "removeListener" in WEBEXT?.fontSettings?.onDefaultFixedFontSizeChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDefaultFixedFontSizeChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.onDefaultFixedFontSizeChanged.removeListener);
    },
    "call_OffDefaultFixedFontSizeChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fontSettings.onDefaultFixedFontSizeChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffDefaultFixedFontSizeChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fontSettings.onDefaultFixedFontSizeChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDefaultFixedFontSizeChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fontSettings?.onDefaultFixedFontSizeChanged &&
        "hasListener" in WEBEXT?.fontSettings?.onDefaultFixedFontSizeChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDefaultFixedFontSizeChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.onDefaultFixedFontSizeChanged.hasListener);
    },
    "call_HasOnDefaultFixedFontSizeChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fontSettings.onDefaultFixedFontSizeChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDefaultFixedFontSizeChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fontSettings.onDefaultFixedFontSizeChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDefaultFontSizeChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fontSettings?.onDefaultFontSizeChanged &&
        "addListener" in WEBEXT?.fontSettings?.onDefaultFontSizeChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDefaultFontSizeChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.onDefaultFontSizeChanged.addListener);
    },
    "call_OnDefaultFontSizeChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fontSettings.onDefaultFontSizeChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnDefaultFontSizeChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fontSettings.onDefaultFontSizeChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDefaultFontSizeChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fontSettings?.onDefaultFontSizeChanged &&
        "removeListener" in WEBEXT?.fontSettings?.onDefaultFontSizeChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDefaultFontSizeChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.onDefaultFontSizeChanged.removeListener);
    },
    "call_OffDefaultFontSizeChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fontSettings.onDefaultFontSizeChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffDefaultFontSizeChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fontSettings.onDefaultFontSizeChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDefaultFontSizeChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fontSettings?.onDefaultFontSizeChanged &&
        "hasListener" in WEBEXT?.fontSettings?.onDefaultFontSizeChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDefaultFontSizeChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.onDefaultFontSizeChanged.hasListener);
    },
    "call_HasOnDefaultFontSizeChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fontSettings.onDefaultFontSizeChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDefaultFontSizeChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fontSettings.onDefaultFontSizeChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnFontChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.fontSettings?.onFontChanged && "addListener" in WEBEXT?.fontSettings?.onFontChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnFontChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.onFontChanged.addListener);
    },
    "call_OnFontChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fontSettings.onFontChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnFontChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fontSettings.onFontChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffFontChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.fontSettings?.onFontChanged && "removeListener" in WEBEXT?.fontSettings?.onFontChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffFontChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.onFontChanged.removeListener);
    },
    "call_OffFontChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fontSettings.onFontChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffFontChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fontSettings.onFontChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnFontChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.fontSettings?.onFontChanged && "hasListener" in WEBEXT?.fontSettings?.onFontChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnFontChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.onFontChanged.hasListener);
    },
    "call_HasOnFontChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fontSettings.onFontChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnFontChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fontSettings.onFontChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnMinimumFontSizeChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fontSettings?.onMinimumFontSizeChanged &&
        "addListener" in WEBEXT?.fontSettings?.onMinimumFontSizeChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMinimumFontSizeChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.onMinimumFontSizeChanged.addListener);
    },
    "call_OnMinimumFontSizeChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fontSettings.onMinimumFontSizeChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnMinimumFontSizeChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fontSettings.onMinimumFontSizeChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMinimumFontSizeChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fontSettings?.onMinimumFontSizeChanged &&
        "removeListener" in WEBEXT?.fontSettings?.onMinimumFontSizeChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMinimumFontSizeChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.onMinimumFontSizeChanged.removeListener);
    },
    "call_OffMinimumFontSizeChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fontSettings.onMinimumFontSizeChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffMinimumFontSizeChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fontSettings.onMinimumFontSizeChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMinimumFontSizeChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fontSettings?.onMinimumFontSizeChanged &&
        "hasListener" in WEBEXT?.fontSettings?.onMinimumFontSizeChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMinimumFontSizeChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.onMinimumFontSizeChanged.hasListener);
    },
    "call_HasOnMinimumFontSizeChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fontSettings.onMinimumFontSizeChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMinimumFontSizeChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fontSettings.onMinimumFontSizeChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetDefaultFixedFontSize": (): heap.Ref<boolean> => {
      if (WEBEXT?.fontSettings && "setDefaultFixedFontSize" in WEBEXT?.fontSettings) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetDefaultFixedFontSize": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.setDefaultFixedFontSize);
    },
    "call_SetDefaultFixedFontSize": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["pixelSize"] = A.load.Int64(details + 0);

      const _ret = WEBEXT.fontSettings.setDefaultFixedFontSize(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetDefaultFixedFontSize": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["pixelSize"] = A.load.Int64(details + 0);

        const _ret = WEBEXT.fontSettings.setDefaultFixedFontSize(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetDefaultFontSize": (): heap.Ref<boolean> => {
      if (WEBEXT?.fontSettings && "setDefaultFontSize" in WEBEXT?.fontSettings) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetDefaultFontSize": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.setDefaultFontSize);
    },
    "call_SetDefaultFontSize": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["pixelSize"] = A.load.Int64(details + 0);

      const _ret = WEBEXT.fontSettings.setDefaultFontSize(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetDefaultFontSize": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["pixelSize"] = A.load.Int64(details + 0);

        const _ret = WEBEXT.fontSettings.setDefaultFontSize(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetFont": (): heap.Ref<boolean> => {
      if (WEBEXT?.fontSettings && "setFont" in WEBEXT?.fontSettings) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetFont": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.setFont);
    },
    "call_SetFont": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["fontId"] = A.load.Ref(details + 0, undefined);
      details_ffi["genericFamily"] = A.load.Enum(details + 4, [
        "standard",
        "sansserif",
        "serif",
        "fixed",
        "cursive",
        "fantasy",
        "math",
      ]);
      details_ffi["script"] = A.load.Enum(details + 8, [
        "Afak",
        "Arab",
        "Armi",
        "Armn",
        "Avst",
        "Bali",
        "Bamu",
        "Bass",
        "Batk",
        "Beng",
        "Blis",
        "Bopo",
        "Brah",
        "Brai",
        "Bugi",
        "Buhd",
        "Cakm",
        "Cans",
        "Cari",
        "Cham",
        "Cher",
        "Cirt",
        "Copt",
        "Cprt",
        "Cyrl",
        "Cyrs",
        "Deva",
        "Dsrt",
        "Dupl",
        "Egyd",
        "Egyh",
        "Egyp",
        "Elba",
        "Ethi",
        "Geor",
        "Geok",
        "Glag",
        "Goth",
        "Gran",
        "Grek",
        "Gujr",
        "Guru",
        "Hang",
        "Hani",
        "Hano",
        "Hans",
        "Hant",
        "Hebr",
        "Hluw",
        "Hmng",
        "Hung",
        "Inds",
        "Ital",
        "Java",
        "Jpan",
        "Jurc",
        "Kali",
        "Khar",
        "Khmr",
        "Khoj",
        "Knda",
        "Kpel",
        "Kthi",
        "Lana",
        "Laoo",
        "Latf",
        "Latg",
        "Latn",
        "Lepc",
        "Limb",
        "Lina",
        "Linb",
        "Lisu",
        "Loma",
        "Lyci",
        "Lydi",
        "Mand",
        "Mani",
        "Maya",
        "Mend",
        "Merc",
        "Mero",
        "Mlym",
        "Moon",
        "Mong",
        "Mroo",
        "Mtei",
        "Mymr",
        "Narb",
        "Nbat",
        "Nkgb",
        "Nkoo",
        "Nshu",
        "Ogam",
        "Olck",
        "Orkh",
        "Orya",
        "Osma",
        "Palm",
        "Perm",
        "Phag",
        "Phli",
        "Phlp",
        "Phlv",
        "Phnx",
        "Plrd",
        "Prti",
        "Rjng",
        "Roro",
        "Runr",
        "Samr",
        "Sara",
        "Sarb",
        "Saur",
        "Sgnw",
        "Shaw",
        "Shrd",
        "Sind",
        "Sinh",
        "Sora",
        "Sund",
        "Sylo",
        "Syrc",
        "Syre",
        "Syrj",
        "Syrn",
        "Tagb",
        "Takr",
        "Tale",
        "Talu",
        "Taml",
        "Tang",
        "Tavt",
        "Telu",
        "Teng",
        "Tfng",
        "Tglg",
        "Thaa",
        "Thai",
        "Tibt",
        "Tirh",
        "Ugar",
        "Vaii",
        "Visp",
        "Wara",
        "Wole",
        "Xpeo",
        "Xsux",
        "Yiii",
        "Zmth",
        "Zsym",
        "Zyyy",
      ]);

      const _ret = WEBEXT.fontSettings.setFont(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetFont": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["fontId"] = A.load.Ref(details + 0, undefined);
        details_ffi["genericFamily"] = A.load.Enum(details + 4, [
          "standard",
          "sansserif",
          "serif",
          "fixed",
          "cursive",
          "fantasy",
          "math",
        ]);
        details_ffi["script"] = A.load.Enum(details + 8, [
          "Afak",
          "Arab",
          "Armi",
          "Armn",
          "Avst",
          "Bali",
          "Bamu",
          "Bass",
          "Batk",
          "Beng",
          "Blis",
          "Bopo",
          "Brah",
          "Brai",
          "Bugi",
          "Buhd",
          "Cakm",
          "Cans",
          "Cari",
          "Cham",
          "Cher",
          "Cirt",
          "Copt",
          "Cprt",
          "Cyrl",
          "Cyrs",
          "Deva",
          "Dsrt",
          "Dupl",
          "Egyd",
          "Egyh",
          "Egyp",
          "Elba",
          "Ethi",
          "Geor",
          "Geok",
          "Glag",
          "Goth",
          "Gran",
          "Grek",
          "Gujr",
          "Guru",
          "Hang",
          "Hani",
          "Hano",
          "Hans",
          "Hant",
          "Hebr",
          "Hluw",
          "Hmng",
          "Hung",
          "Inds",
          "Ital",
          "Java",
          "Jpan",
          "Jurc",
          "Kali",
          "Khar",
          "Khmr",
          "Khoj",
          "Knda",
          "Kpel",
          "Kthi",
          "Lana",
          "Laoo",
          "Latf",
          "Latg",
          "Latn",
          "Lepc",
          "Limb",
          "Lina",
          "Linb",
          "Lisu",
          "Loma",
          "Lyci",
          "Lydi",
          "Mand",
          "Mani",
          "Maya",
          "Mend",
          "Merc",
          "Mero",
          "Mlym",
          "Moon",
          "Mong",
          "Mroo",
          "Mtei",
          "Mymr",
          "Narb",
          "Nbat",
          "Nkgb",
          "Nkoo",
          "Nshu",
          "Ogam",
          "Olck",
          "Orkh",
          "Orya",
          "Osma",
          "Palm",
          "Perm",
          "Phag",
          "Phli",
          "Phlp",
          "Phlv",
          "Phnx",
          "Plrd",
          "Prti",
          "Rjng",
          "Roro",
          "Runr",
          "Samr",
          "Sara",
          "Sarb",
          "Saur",
          "Sgnw",
          "Shaw",
          "Shrd",
          "Sind",
          "Sinh",
          "Sora",
          "Sund",
          "Sylo",
          "Syrc",
          "Syre",
          "Syrj",
          "Syrn",
          "Tagb",
          "Takr",
          "Tale",
          "Talu",
          "Taml",
          "Tang",
          "Tavt",
          "Telu",
          "Teng",
          "Tfng",
          "Tglg",
          "Thaa",
          "Thai",
          "Tibt",
          "Tirh",
          "Ugar",
          "Vaii",
          "Visp",
          "Wara",
          "Wole",
          "Xpeo",
          "Xsux",
          "Yiii",
          "Zmth",
          "Zsym",
          "Zyyy",
        ]);

        const _ret = WEBEXT.fontSettings.setFont(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetMinimumFontSize": (): heap.Ref<boolean> => {
      if (WEBEXT?.fontSettings && "setMinimumFontSize" in WEBEXT?.fontSettings) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetMinimumFontSize": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fontSettings.setMinimumFontSize);
    },
    "call_SetMinimumFontSize": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["pixelSize"] = A.load.Int64(details + 0);

      const _ret = WEBEXT.fontSettings.setMinimumFontSize(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetMinimumFontSize": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["pixelSize"] = A.load.Int64(details + 0);

        const _ret = WEBEXT.fontSettings.setMinimumFontSize(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
