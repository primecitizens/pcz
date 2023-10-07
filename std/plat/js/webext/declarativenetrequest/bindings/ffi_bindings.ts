import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/declarativenetrequest", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_Ruleset": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Ref(ptr + 4, x["path"]);
        A.store.Bool(ptr + 9, "enabled" in x ? true : false);
        A.store.Bool(ptr + 8, x["enabled"] ? true : false);
      }
    },
    "load_Ruleset": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["path"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 9)) {
        x["enabled"] = A.load.Bool(ptr + 8);
      } else {
        delete x["enabled"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DNRInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["rule_resources"]);
      }
    },
    "load_DNRInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["rule_resources"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_DomainType": (ref: heap.Ref<string>): number => {
      const idx = ["firstParty", "thirdParty"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_TabActionCountUpdate": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 9, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Bool(ptr + 8, "tabId" in x ? true : false);
        A.store.Int32(ptr + 0, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Bool(ptr + 9, "increment" in x ? true : false);
        A.store.Int32(ptr + 4, x["increment"] === undefined ? 0 : (x["increment"] as number));
      }
    },
    "load_TabActionCountUpdate": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["tabId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["tabId"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["increment"] = A.load.Int32(ptr + 4);
      } else {
        delete x["increment"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ExtensionActionOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 15, false);
        A.store.Bool(ptr + 0, false);

        A.store.Bool(ptr + 4 + 10, false);
        A.store.Bool(ptr + 4 + 8, false);
        A.store.Int32(ptr + 4 + 0, 0);
        A.store.Bool(ptr + 4 + 9, false);
        A.store.Int32(ptr + 4 + 4, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Bool(ptr + 15, "displayActionCountAsBadgeText" in x ? true : false);
        A.store.Bool(ptr + 0, x["displayActionCountAsBadgeText"] ? true : false);

        if (typeof x["tabUpdate"] === "undefined") {
          A.store.Bool(ptr + 4 + 10, false);
          A.store.Bool(ptr + 4 + 8, false);
          A.store.Int32(ptr + 4 + 0, 0);
          A.store.Bool(ptr + 4 + 9, false);
          A.store.Int32(ptr + 4 + 4, 0);
        } else {
          A.store.Bool(ptr + 4 + 10, true);
          A.store.Bool(ptr + 4 + 8, "tabId" in x["tabUpdate"] ? true : false);
          A.store.Int32(ptr + 4 + 0, x["tabUpdate"]["tabId"] === undefined ? 0 : (x["tabUpdate"]["tabId"] as number));
          A.store.Bool(ptr + 4 + 9, "increment" in x["tabUpdate"] ? true : false);
          A.store.Int32(
            ptr + 4 + 4,
            x["tabUpdate"]["increment"] === undefined ? 0 : (x["tabUpdate"]["increment"] as number)
          );
        }
      }
    },
    "load_ExtensionActionOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 15)) {
        x["displayActionCountAsBadgeText"] = A.load.Bool(ptr + 0);
      } else {
        delete x["displayActionCountAsBadgeText"];
      }
      if (A.load.Bool(ptr + 4 + 10)) {
        x["tabUpdate"] = {};
        if (A.load.Bool(ptr + 4 + 8)) {
          x["tabUpdate"]["tabId"] = A.load.Int32(ptr + 4 + 0);
        } else {
          delete x["tabUpdate"]["tabId"];
        }
        if (A.load.Bool(ptr + 4 + 9)) {
          x["tabUpdate"]["increment"] = A.load.Int32(ptr + 4 + 4);
        } else {
          delete x["tabUpdate"]["increment"];
        }
      } else {
        delete x["tabUpdate"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetDisabledRuleIdsOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["rulesetId"]);
      }
    },
    "load_GetDisabledRuleIdsOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["rulesetId"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_MatchedRule": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "ruleId" in x ? true : false);
        A.store.Int32(ptr + 0, x["ruleId"] === undefined ? 0 : (x["ruleId"] as number));
        A.store.Ref(ptr + 4, x["rulesetId"]);
      }
    },
    "load_MatchedRule": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["ruleId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["ruleId"];
      }
      x["rulesetId"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_MatchedRuleInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 30, false);

        A.store.Bool(ptr + 0 + 9, false);
        A.store.Bool(ptr + 0 + 8, false);
        A.store.Int32(ptr + 0 + 0, 0);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Bool(ptr + 28, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Bool(ptr + 29, false);
        A.store.Int32(ptr + 24, 0);
      } else {
        A.store.Bool(ptr + 30, true);

        if (typeof x["rule"] === "undefined") {
          A.store.Bool(ptr + 0 + 9, false);
          A.store.Bool(ptr + 0 + 8, false);
          A.store.Int32(ptr + 0 + 0, 0);
          A.store.Ref(ptr + 0 + 4, undefined);
        } else {
          A.store.Bool(ptr + 0 + 9, true);
          A.store.Bool(ptr + 0 + 8, "ruleId" in x["rule"] ? true : false);
          A.store.Int32(ptr + 0 + 0, x["rule"]["ruleId"] === undefined ? 0 : (x["rule"]["ruleId"] as number));
          A.store.Ref(ptr + 0 + 4, x["rule"]["rulesetId"]);
        }
        A.store.Bool(ptr + 28, "timeStamp" in x ? true : false);
        A.store.Float64(ptr + 16, x["timeStamp"] === undefined ? 0 : (x["timeStamp"] as number));
        A.store.Bool(ptr + 29, "tabId" in x ? true : false);
        A.store.Int32(ptr + 24, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
      }
    },
    "load_MatchedRuleInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 9)) {
        x["rule"] = {};
        if (A.load.Bool(ptr + 0 + 8)) {
          x["rule"]["ruleId"] = A.load.Int32(ptr + 0 + 0);
        } else {
          delete x["rule"]["ruleId"];
        }
        x["rule"]["rulesetId"] = A.load.Ref(ptr + 0 + 4, undefined);
      } else {
        delete x["rule"];
      }
      if (A.load.Bool(ptr + 28)) {
        x["timeStamp"] = A.load.Float64(ptr + 16);
      } else {
        delete x["timeStamp"];
      }
      if (A.load.Bool(ptr + 29)) {
        x["tabId"] = A.load.Int32(ptr + 24);
      } else {
        delete x["tabId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RulesMatchedDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["rulesMatchedInfo"]);
      }
    },
    "load_RulesMatchedDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["rulesMatchedInfo"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ResourceType": (ref: heap.Ref<string>): number => {
      const idx = [
        "main_frame",
        "sub_frame",
        "stylesheet",
        "script",
        "image",
        "font",
        "object",
        "xmlhttprequest",
        "ping",
        "csp_report",
        "media",
        "websocket",
        "webtransport",
        "webbundle",
        "other",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_RequestMethod": (ref: heap.Ref<string>): number => {
      const idx = ["connect", "delete", "get", "head", "options", "patch", "post", "put", "other"].indexOf(
        A.H.get(ref)
      );
      return idx < 0 ? 0 : idx + 1;
    },

    "store_RuleCondition": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 65, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 64, false);
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Ref(ptr + 32, undefined);
        A.store.Ref(ptr + 36, undefined);
        A.store.Ref(ptr + 40, undefined);
        A.store.Ref(ptr + 44, undefined);
        A.store.Ref(ptr + 48, undefined);
        A.store.Enum(ptr + 52, -1);
        A.store.Ref(ptr + 56, undefined);
        A.store.Ref(ptr + 60, undefined);
      } else {
        A.store.Bool(ptr + 65, true);
        A.store.Ref(ptr + 0, x["urlFilter"]);
        A.store.Ref(ptr + 4, x["regexFilter"]);
        A.store.Bool(ptr + 64, "isUrlFilterCaseSensitive" in x ? true : false);
        A.store.Bool(ptr + 8, x["isUrlFilterCaseSensitive"] ? true : false);
        A.store.Ref(ptr + 12, x["initiatorDomains"]);
        A.store.Ref(ptr + 16, x["excludedInitiatorDomains"]);
        A.store.Ref(ptr + 20, x["requestDomains"]);
        A.store.Ref(ptr + 24, x["excludedRequestDomains"]);
        A.store.Ref(ptr + 28, x["domains"]);
        A.store.Ref(ptr + 32, x["excludedDomains"]);
        A.store.Ref(ptr + 36, x["resourceTypes"]);
        A.store.Ref(ptr + 40, x["excludedResourceTypes"]);
        A.store.Ref(ptr + 44, x["requestMethods"]);
        A.store.Ref(ptr + 48, x["excludedRequestMethods"]);
        A.store.Enum(ptr + 52, ["firstParty", "thirdParty"].indexOf(x["domainType"] as string));
        A.store.Ref(ptr + 56, x["tabIds"]);
        A.store.Ref(ptr + 60, x["excludedTabIds"]);
      }
    },
    "load_RuleCondition": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["urlFilter"] = A.load.Ref(ptr + 0, undefined);
      x["regexFilter"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 64)) {
        x["isUrlFilterCaseSensitive"] = A.load.Bool(ptr + 8);
      } else {
        delete x["isUrlFilterCaseSensitive"];
      }
      x["initiatorDomains"] = A.load.Ref(ptr + 12, undefined);
      x["excludedInitiatorDomains"] = A.load.Ref(ptr + 16, undefined);
      x["requestDomains"] = A.load.Ref(ptr + 20, undefined);
      x["excludedRequestDomains"] = A.load.Ref(ptr + 24, undefined);
      x["domains"] = A.load.Ref(ptr + 28, undefined);
      x["excludedDomains"] = A.load.Ref(ptr + 32, undefined);
      x["resourceTypes"] = A.load.Ref(ptr + 36, undefined);
      x["excludedResourceTypes"] = A.load.Ref(ptr + 40, undefined);
      x["requestMethods"] = A.load.Ref(ptr + 44, undefined);
      x["excludedRequestMethods"] = A.load.Ref(ptr + 48, undefined);
      x["domainType"] = A.load.Enum(ptr + 52, ["firstParty", "thirdParty"]);
      x["tabIds"] = A.load.Ref(ptr + 56, undefined);
      x["excludedTabIds"] = A.load.Ref(ptr + 60, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_RuleActionType": (ref: heap.Ref<string>): number => {
      const idx = ["block", "redirect", "allow", "upgradeScheme", "modifyHeaders", "allowAllRequests"].indexOf(
        A.H.get(ref)
      );
      return idx < 0 ? 0 : idx + 1;
    },

    "store_QueryKeyValue": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Ref(ptr + 0, x["key"]);
        A.store.Ref(ptr + 4, x["value"]);
        A.store.Bool(ptr + 9, "replaceOnly" in x ? true : false);
        A.store.Bool(ptr + 8, x["replaceOnly"] ? true : false);
      }
    },
    "load_QueryKeyValue": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["key"] = A.load.Ref(ptr + 0, undefined);
      x["value"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 9)) {
        x["replaceOnly"] = A.load.Bool(ptr + 8);
      } else {
        delete x["replaceOnly"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_QueryTransform": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["removeParams"]);
        A.store.Ref(ptr + 4, x["addOrReplaceParams"]);
      }
    },
    "load_QueryTransform": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["removeParams"] = A.load.Ref(ptr + 0, undefined);
      x["addOrReplaceParams"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_URLTransform": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 44, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);

        A.store.Bool(ptr + 20 + 8, false);
        A.store.Ref(ptr + 20 + 0, undefined);
        A.store.Ref(ptr + 20 + 4, undefined);
        A.store.Ref(ptr + 32, undefined);
        A.store.Ref(ptr + 36, undefined);
        A.store.Ref(ptr + 40, undefined);
      } else {
        A.store.Bool(ptr + 44, true);
        A.store.Ref(ptr + 0, x["scheme"]);
        A.store.Ref(ptr + 4, x["host"]);
        A.store.Ref(ptr + 8, x["port"]);
        A.store.Ref(ptr + 12, x["path"]);
        A.store.Ref(ptr + 16, x["query"]);

        if (typeof x["queryTransform"] === "undefined") {
          A.store.Bool(ptr + 20 + 8, false);
          A.store.Ref(ptr + 20 + 0, undefined);
          A.store.Ref(ptr + 20 + 4, undefined);
        } else {
          A.store.Bool(ptr + 20 + 8, true);
          A.store.Ref(ptr + 20 + 0, x["queryTransform"]["removeParams"]);
          A.store.Ref(ptr + 20 + 4, x["queryTransform"]["addOrReplaceParams"]);
        }
        A.store.Ref(ptr + 32, x["fragment"]);
        A.store.Ref(ptr + 36, x["username"]);
        A.store.Ref(ptr + 40, x["password"]);
      }
    },
    "load_URLTransform": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["scheme"] = A.load.Ref(ptr + 0, undefined);
      x["host"] = A.load.Ref(ptr + 4, undefined);
      x["port"] = A.load.Ref(ptr + 8, undefined);
      x["path"] = A.load.Ref(ptr + 12, undefined);
      x["query"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 20 + 8)) {
        x["queryTransform"] = {};
        x["queryTransform"]["removeParams"] = A.load.Ref(ptr + 20 + 0, undefined);
        x["queryTransform"]["addOrReplaceParams"] = A.load.Ref(ptr + 20 + 4, undefined);
      } else {
        delete x["queryTransform"];
      }
      x["fragment"] = A.load.Ref(ptr + 32, undefined);
      x["username"] = A.load.Ref(ptr + 36, undefined);
      x["password"] = A.load.Ref(ptr + 40, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Redirect": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 60, false);
        A.store.Ref(ptr + 0, undefined);

        A.store.Bool(ptr + 4 + 44, false);
        A.store.Ref(ptr + 4 + 0, undefined);
        A.store.Ref(ptr + 4 + 4, undefined);
        A.store.Ref(ptr + 4 + 8, undefined);
        A.store.Ref(ptr + 4 + 12, undefined);
        A.store.Ref(ptr + 4 + 16, undefined);

        A.store.Bool(ptr + 4 + 20 + 8, false);
        A.store.Ref(ptr + 4 + 20 + 0, undefined);
        A.store.Ref(ptr + 4 + 20 + 4, undefined);
        A.store.Ref(ptr + 4 + 32, undefined);
        A.store.Ref(ptr + 4 + 36, undefined);
        A.store.Ref(ptr + 4 + 40, undefined);
        A.store.Ref(ptr + 52, undefined);
        A.store.Ref(ptr + 56, undefined);
      } else {
        A.store.Bool(ptr + 60, true);
        A.store.Ref(ptr + 0, x["extensionPath"]);

        if (typeof x["transform"] === "undefined") {
          A.store.Bool(ptr + 4 + 44, false);
          A.store.Ref(ptr + 4 + 0, undefined);
          A.store.Ref(ptr + 4 + 4, undefined);
          A.store.Ref(ptr + 4 + 8, undefined);
          A.store.Ref(ptr + 4 + 12, undefined);
          A.store.Ref(ptr + 4 + 16, undefined);

          A.store.Bool(ptr + 4 + 20 + 8, false);
          A.store.Ref(ptr + 4 + 20 + 0, undefined);
          A.store.Ref(ptr + 4 + 20 + 4, undefined);
          A.store.Ref(ptr + 4 + 32, undefined);
          A.store.Ref(ptr + 4 + 36, undefined);
          A.store.Ref(ptr + 4 + 40, undefined);
        } else {
          A.store.Bool(ptr + 4 + 44, true);
          A.store.Ref(ptr + 4 + 0, x["transform"]["scheme"]);
          A.store.Ref(ptr + 4 + 4, x["transform"]["host"]);
          A.store.Ref(ptr + 4 + 8, x["transform"]["port"]);
          A.store.Ref(ptr + 4 + 12, x["transform"]["path"]);
          A.store.Ref(ptr + 4 + 16, x["transform"]["query"]);

          if (typeof x["transform"]["queryTransform"] === "undefined") {
            A.store.Bool(ptr + 4 + 20 + 8, false);
            A.store.Ref(ptr + 4 + 20 + 0, undefined);
            A.store.Ref(ptr + 4 + 20 + 4, undefined);
          } else {
            A.store.Bool(ptr + 4 + 20 + 8, true);
            A.store.Ref(ptr + 4 + 20 + 0, x["transform"]["queryTransform"]["removeParams"]);
            A.store.Ref(ptr + 4 + 20 + 4, x["transform"]["queryTransform"]["addOrReplaceParams"]);
          }
          A.store.Ref(ptr + 4 + 32, x["transform"]["fragment"]);
          A.store.Ref(ptr + 4 + 36, x["transform"]["username"]);
          A.store.Ref(ptr + 4 + 40, x["transform"]["password"]);
        }
        A.store.Ref(ptr + 52, x["url"]);
        A.store.Ref(ptr + 56, x["regexSubstitution"]);
      }
    },
    "load_Redirect": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["extensionPath"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 4 + 44)) {
        x["transform"] = {};
        x["transform"]["scheme"] = A.load.Ref(ptr + 4 + 0, undefined);
        x["transform"]["host"] = A.load.Ref(ptr + 4 + 4, undefined);
        x["transform"]["port"] = A.load.Ref(ptr + 4 + 8, undefined);
        x["transform"]["path"] = A.load.Ref(ptr + 4 + 12, undefined);
        x["transform"]["query"] = A.load.Ref(ptr + 4 + 16, undefined);
        if (A.load.Bool(ptr + 4 + 20 + 8)) {
          x["transform"]["queryTransform"] = {};
          x["transform"]["queryTransform"]["removeParams"] = A.load.Ref(ptr + 4 + 20 + 0, undefined);
          x["transform"]["queryTransform"]["addOrReplaceParams"] = A.load.Ref(ptr + 4 + 20 + 4, undefined);
        } else {
          delete x["transform"]["queryTransform"];
        }
        x["transform"]["fragment"] = A.load.Ref(ptr + 4 + 32, undefined);
        x["transform"]["username"] = A.load.Ref(ptr + 4 + 36, undefined);
        x["transform"]["password"] = A.load.Ref(ptr + 4 + 40, undefined);
      } else {
        delete x["transform"];
      }
      x["url"] = A.load.Ref(ptr + 52, undefined);
      x["regexSubstitution"] = A.load.Ref(ptr + 56, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_HeaderOperation": (ref: heap.Ref<string>): number => {
      const idx = ["append", "set", "remove"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ModifyHeaderInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["header"]);
        A.store.Enum(ptr + 4, ["append", "set", "remove"].indexOf(x["operation"] as string));
        A.store.Ref(ptr + 8, x["value"]);
      }
    },
    "load_ModifyHeaderInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["header"] = A.load.Ref(ptr + 0, undefined);
      x["operation"] = A.load.Enum(ptr + 4, ["append", "set", "remove"]);
      x["value"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RuleAction": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 76, false);
        A.store.Enum(ptr + 0, -1);

        A.store.Bool(ptr + 4 + 60, false);
        A.store.Ref(ptr + 4 + 0, undefined);

        A.store.Bool(ptr + 4 + 4 + 44, false);
        A.store.Ref(ptr + 4 + 4 + 0, undefined);
        A.store.Ref(ptr + 4 + 4 + 4, undefined);
        A.store.Ref(ptr + 4 + 4 + 8, undefined);
        A.store.Ref(ptr + 4 + 4 + 12, undefined);
        A.store.Ref(ptr + 4 + 4 + 16, undefined);

        A.store.Bool(ptr + 4 + 4 + 20 + 8, false);
        A.store.Ref(ptr + 4 + 4 + 20 + 0, undefined);
        A.store.Ref(ptr + 4 + 4 + 20 + 4, undefined);
        A.store.Ref(ptr + 4 + 4 + 32, undefined);
        A.store.Ref(ptr + 4 + 4 + 36, undefined);
        A.store.Ref(ptr + 4 + 4 + 40, undefined);
        A.store.Ref(ptr + 4 + 52, undefined);
        A.store.Ref(ptr + 4 + 56, undefined);
        A.store.Ref(ptr + 68, undefined);
        A.store.Ref(ptr + 72, undefined);
      } else {
        A.store.Bool(ptr + 76, true);
        A.store.Enum(
          ptr + 0,
          ["block", "redirect", "allow", "upgradeScheme", "modifyHeaders", "allowAllRequests"].indexOf(
            x["type"] as string
          )
        );

        if (typeof x["redirect"] === "undefined") {
          A.store.Bool(ptr + 4 + 60, false);
          A.store.Ref(ptr + 4 + 0, undefined);

          A.store.Bool(ptr + 4 + 4 + 44, false);
          A.store.Ref(ptr + 4 + 4 + 0, undefined);
          A.store.Ref(ptr + 4 + 4 + 4, undefined);
          A.store.Ref(ptr + 4 + 4 + 8, undefined);
          A.store.Ref(ptr + 4 + 4 + 12, undefined);
          A.store.Ref(ptr + 4 + 4 + 16, undefined);

          A.store.Bool(ptr + 4 + 4 + 20 + 8, false);
          A.store.Ref(ptr + 4 + 4 + 20 + 0, undefined);
          A.store.Ref(ptr + 4 + 4 + 20 + 4, undefined);
          A.store.Ref(ptr + 4 + 4 + 32, undefined);
          A.store.Ref(ptr + 4 + 4 + 36, undefined);
          A.store.Ref(ptr + 4 + 4 + 40, undefined);
          A.store.Ref(ptr + 4 + 52, undefined);
          A.store.Ref(ptr + 4 + 56, undefined);
        } else {
          A.store.Bool(ptr + 4 + 60, true);
          A.store.Ref(ptr + 4 + 0, x["redirect"]["extensionPath"]);

          if (typeof x["redirect"]["transform"] === "undefined") {
            A.store.Bool(ptr + 4 + 4 + 44, false);
            A.store.Ref(ptr + 4 + 4 + 0, undefined);
            A.store.Ref(ptr + 4 + 4 + 4, undefined);
            A.store.Ref(ptr + 4 + 4 + 8, undefined);
            A.store.Ref(ptr + 4 + 4 + 12, undefined);
            A.store.Ref(ptr + 4 + 4 + 16, undefined);

            A.store.Bool(ptr + 4 + 4 + 20 + 8, false);
            A.store.Ref(ptr + 4 + 4 + 20 + 0, undefined);
            A.store.Ref(ptr + 4 + 4 + 20 + 4, undefined);
            A.store.Ref(ptr + 4 + 4 + 32, undefined);
            A.store.Ref(ptr + 4 + 4 + 36, undefined);
            A.store.Ref(ptr + 4 + 4 + 40, undefined);
          } else {
            A.store.Bool(ptr + 4 + 4 + 44, true);
            A.store.Ref(ptr + 4 + 4 + 0, x["redirect"]["transform"]["scheme"]);
            A.store.Ref(ptr + 4 + 4 + 4, x["redirect"]["transform"]["host"]);
            A.store.Ref(ptr + 4 + 4 + 8, x["redirect"]["transform"]["port"]);
            A.store.Ref(ptr + 4 + 4 + 12, x["redirect"]["transform"]["path"]);
            A.store.Ref(ptr + 4 + 4 + 16, x["redirect"]["transform"]["query"]);

            if (typeof x["redirect"]["transform"]["queryTransform"] === "undefined") {
              A.store.Bool(ptr + 4 + 4 + 20 + 8, false);
              A.store.Ref(ptr + 4 + 4 + 20 + 0, undefined);
              A.store.Ref(ptr + 4 + 4 + 20 + 4, undefined);
            } else {
              A.store.Bool(ptr + 4 + 4 + 20 + 8, true);
              A.store.Ref(ptr + 4 + 4 + 20 + 0, x["redirect"]["transform"]["queryTransform"]["removeParams"]);
              A.store.Ref(ptr + 4 + 4 + 20 + 4, x["redirect"]["transform"]["queryTransform"]["addOrReplaceParams"]);
            }
            A.store.Ref(ptr + 4 + 4 + 32, x["redirect"]["transform"]["fragment"]);
            A.store.Ref(ptr + 4 + 4 + 36, x["redirect"]["transform"]["username"]);
            A.store.Ref(ptr + 4 + 4 + 40, x["redirect"]["transform"]["password"]);
          }
          A.store.Ref(ptr + 4 + 52, x["redirect"]["url"]);
          A.store.Ref(ptr + 4 + 56, x["redirect"]["regexSubstitution"]);
        }
        A.store.Ref(ptr + 68, x["requestHeaders"]);
        A.store.Ref(ptr + 72, x["responseHeaders"]);
      }
    },
    "load_RuleAction": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["type"] = A.load.Enum(ptr + 0, [
        "block",
        "redirect",
        "allow",
        "upgradeScheme",
        "modifyHeaders",
        "allowAllRequests",
      ]);
      if (A.load.Bool(ptr + 4 + 60)) {
        x["redirect"] = {};
        x["redirect"]["extensionPath"] = A.load.Ref(ptr + 4 + 0, undefined);
        if (A.load.Bool(ptr + 4 + 4 + 44)) {
          x["redirect"]["transform"] = {};
          x["redirect"]["transform"]["scheme"] = A.load.Ref(ptr + 4 + 4 + 0, undefined);
          x["redirect"]["transform"]["host"] = A.load.Ref(ptr + 4 + 4 + 4, undefined);
          x["redirect"]["transform"]["port"] = A.load.Ref(ptr + 4 + 4 + 8, undefined);
          x["redirect"]["transform"]["path"] = A.load.Ref(ptr + 4 + 4 + 12, undefined);
          x["redirect"]["transform"]["query"] = A.load.Ref(ptr + 4 + 4 + 16, undefined);
          if (A.load.Bool(ptr + 4 + 4 + 20 + 8)) {
            x["redirect"]["transform"]["queryTransform"] = {};
            x["redirect"]["transform"]["queryTransform"]["removeParams"] = A.load.Ref(ptr + 4 + 4 + 20 + 0, undefined);
            x["redirect"]["transform"]["queryTransform"]["addOrReplaceParams"] = A.load.Ref(
              ptr + 4 + 4 + 20 + 4,
              undefined
            );
          } else {
            delete x["redirect"]["transform"]["queryTransform"];
          }
          x["redirect"]["transform"]["fragment"] = A.load.Ref(ptr + 4 + 4 + 32, undefined);
          x["redirect"]["transform"]["username"] = A.load.Ref(ptr + 4 + 4 + 36, undefined);
          x["redirect"]["transform"]["password"] = A.load.Ref(ptr + 4 + 4 + 40, undefined);
        } else {
          delete x["redirect"]["transform"];
        }
        x["redirect"]["url"] = A.load.Ref(ptr + 4 + 52, undefined);
        x["redirect"]["regexSubstitution"] = A.load.Ref(ptr + 4 + 56, undefined);
      } else {
        delete x["redirect"];
      }
      x["requestHeaders"] = A.load.Ref(ptr + 68, undefined);
      x["responseHeaders"] = A.load.Ref(ptr + 72, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Rule": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 155, false);
        A.store.Bool(ptr + 153, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 154, false);
        A.store.Int32(ptr + 4, 0);

        A.store.Bool(ptr + 8 + 65, false);
        A.store.Ref(ptr + 8 + 0, undefined);
        A.store.Ref(ptr + 8 + 4, undefined);
        A.store.Bool(ptr + 8 + 64, false);
        A.store.Bool(ptr + 8 + 8, false);
        A.store.Ref(ptr + 8 + 12, undefined);
        A.store.Ref(ptr + 8 + 16, undefined);
        A.store.Ref(ptr + 8 + 20, undefined);
        A.store.Ref(ptr + 8 + 24, undefined);
        A.store.Ref(ptr + 8 + 28, undefined);
        A.store.Ref(ptr + 8 + 32, undefined);
        A.store.Ref(ptr + 8 + 36, undefined);
        A.store.Ref(ptr + 8 + 40, undefined);
        A.store.Ref(ptr + 8 + 44, undefined);
        A.store.Ref(ptr + 8 + 48, undefined);
        A.store.Enum(ptr + 8 + 52, -1);
        A.store.Ref(ptr + 8 + 56, undefined);
        A.store.Ref(ptr + 8 + 60, undefined);

        A.store.Bool(ptr + 76 + 76, false);
        A.store.Enum(ptr + 76 + 0, -1);

        A.store.Bool(ptr + 76 + 4 + 60, false);
        A.store.Ref(ptr + 76 + 4 + 0, undefined);

        A.store.Bool(ptr + 76 + 4 + 4 + 44, false);
        A.store.Ref(ptr + 76 + 4 + 4 + 0, undefined);
        A.store.Ref(ptr + 76 + 4 + 4 + 4, undefined);
        A.store.Ref(ptr + 76 + 4 + 4 + 8, undefined);
        A.store.Ref(ptr + 76 + 4 + 4 + 12, undefined);
        A.store.Ref(ptr + 76 + 4 + 4 + 16, undefined);

        A.store.Bool(ptr + 76 + 4 + 4 + 20 + 8, false);
        A.store.Ref(ptr + 76 + 4 + 4 + 20 + 0, undefined);
        A.store.Ref(ptr + 76 + 4 + 4 + 20 + 4, undefined);
        A.store.Ref(ptr + 76 + 4 + 4 + 32, undefined);
        A.store.Ref(ptr + 76 + 4 + 4 + 36, undefined);
        A.store.Ref(ptr + 76 + 4 + 4 + 40, undefined);
        A.store.Ref(ptr + 76 + 4 + 52, undefined);
        A.store.Ref(ptr + 76 + 4 + 56, undefined);
        A.store.Ref(ptr + 76 + 68, undefined);
        A.store.Ref(ptr + 76 + 72, undefined);
      } else {
        A.store.Bool(ptr + 155, true);
        A.store.Bool(ptr + 153, "id" in x ? true : false);
        A.store.Int32(ptr + 0, x["id"] === undefined ? 0 : (x["id"] as number));
        A.store.Bool(ptr + 154, "priority" in x ? true : false);
        A.store.Int32(ptr + 4, x["priority"] === undefined ? 0 : (x["priority"] as number));

        if (typeof x["condition"] === "undefined") {
          A.store.Bool(ptr + 8 + 65, false);
          A.store.Ref(ptr + 8 + 0, undefined);
          A.store.Ref(ptr + 8 + 4, undefined);
          A.store.Bool(ptr + 8 + 64, false);
          A.store.Bool(ptr + 8 + 8, false);
          A.store.Ref(ptr + 8 + 12, undefined);
          A.store.Ref(ptr + 8 + 16, undefined);
          A.store.Ref(ptr + 8 + 20, undefined);
          A.store.Ref(ptr + 8 + 24, undefined);
          A.store.Ref(ptr + 8 + 28, undefined);
          A.store.Ref(ptr + 8 + 32, undefined);
          A.store.Ref(ptr + 8 + 36, undefined);
          A.store.Ref(ptr + 8 + 40, undefined);
          A.store.Ref(ptr + 8 + 44, undefined);
          A.store.Ref(ptr + 8 + 48, undefined);
          A.store.Enum(ptr + 8 + 52, -1);
          A.store.Ref(ptr + 8 + 56, undefined);
          A.store.Ref(ptr + 8 + 60, undefined);
        } else {
          A.store.Bool(ptr + 8 + 65, true);
          A.store.Ref(ptr + 8 + 0, x["condition"]["urlFilter"]);
          A.store.Ref(ptr + 8 + 4, x["condition"]["regexFilter"]);
          A.store.Bool(ptr + 8 + 64, "isUrlFilterCaseSensitive" in x["condition"] ? true : false);
          A.store.Bool(ptr + 8 + 8, x["condition"]["isUrlFilterCaseSensitive"] ? true : false);
          A.store.Ref(ptr + 8 + 12, x["condition"]["initiatorDomains"]);
          A.store.Ref(ptr + 8 + 16, x["condition"]["excludedInitiatorDomains"]);
          A.store.Ref(ptr + 8 + 20, x["condition"]["requestDomains"]);
          A.store.Ref(ptr + 8 + 24, x["condition"]["excludedRequestDomains"]);
          A.store.Ref(ptr + 8 + 28, x["condition"]["domains"]);
          A.store.Ref(ptr + 8 + 32, x["condition"]["excludedDomains"]);
          A.store.Ref(ptr + 8 + 36, x["condition"]["resourceTypes"]);
          A.store.Ref(ptr + 8 + 40, x["condition"]["excludedResourceTypes"]);
          A.store.Ref(ptr + 8 + 44, x["condition"]["requestMethods"]);
          A.store.Ref(ptr + 8 + 48, x["condition"]["excludedRequestMethods"]);
          A.store.Enum(ptr + 8 + 52, ["firstParty", "thirdParty"].indexOf(x["condition"]["domainType"] as string));
          A.store.Ref(ptr + 8 + 56, x["condition"]["tabIds"]);
          A.store.Ref(ptr + 8 + 60, x["condition"]["excludedTabIds"]);
        }

        if (typeof x["action"] === "undefined") {
          A.store.Bool(ptr + 76 + 76, false);
          A.store.Enum(ptr + 76 + 0, -1);

          A.store.Bool(ptr + 76 + 4 + 60, false);
          A.store.Ref(ptr + 76 + 4 + 0, undefined);

          A.store.Bool(ptr + 76 + 4 + 4 + 44, false);
          A.store.Ref(ptr + 76 + 4 + 4 + 0, undefined);
          A.store.Ref(ptr + 76 + 4 + 4 + 4, undefined);
          A.store.Ref(ptr + 76 + 4 + 4 + 8, undefined);
          A.store.Ref(ptr + 76 + 4 + 4 + 12, undefined);
          A.store.Ref(ptr + 76 + 4 + 4 + 16, undefined);

          A.store.Bool(ptr + 76 + 4 + 4 + 20 + 8, false);
          A.store.Ref(ptr + 76 + 4 + 4 + 20 + 0, undefined);
          A.store.Ref(ptr + 76 + 4 + 4 + 20 + 4, undefined);
          A.store.Ref(ptr + 76 + 4 + 4 + 32, undefined);
          A.store.Ref(ptr + 76 + 4 + 4 + 36, undefined);
          A.store.Ref(ptr + 76 + 4 + 4 + 40, undefined);
          A.store.Ref(ptr + 76 + 4 + 52, undefined);
          A.store.Ref(ptr + 76 + 4 + 56, undefined);
          A.store.Ref(ptr + 76 + 68, undefined);
          A.store.Ref(ptr + 76 + 72, undefined);
        } else {
          A.store.Bool(ptr + 76 + 76, true);
          A.store.Enum(
            ptr + 76 + 0,
            ["block", "redirect", "allow", "upgradeScheme", "modifyHeaders", "allowAllRequests"].indexOf(
              x["action"]["type"] as string
            )
          );

          if (typeof x["action"]["redirect"] === "undefined") {
            A.store.Bool(ptr + 76 + 4 + 60, false);
            A.store.Ref(ptr + 76 + 4 + 0, undefined);

            A.store.Bool(ptr + 76 + 4 + 4 + 44, false);
            A.store.Ref(ptr + 76 + 4 + 4 + 0, undefined);
            A.store.Ref(ptr + 76 + 4 + 4 + 4, undefined);
            A.store.Ref(ptr + 76 + 4 + 4 + 8, undefined);
            A.store.Ref(ptr + 76 + 4 + 4 + 12, undefined);
            A.store.Ref(ptr + 76 + 4 + 4 + 16, undefined);

            A.store.Bool(ptr + 76 + 4 + 4 + 20 + 8, false);
            A.store.Ref(ptr + 76 + 4 + 4 + 20 + 0, undefined);
            A.store.Ref(ptr + 76 + 4 + 4 + 20 + 4, undefined);
            A.store.Ref(ptr + 76 + 4 + 4 + 32, undefined);
            A.store.Ref(ptr + 76 + 4 + 4 + 36, undefined);
            A.store.Ref(ptr + 76 + 4 + 4 + 40, undefined);
            A.store.Ref(ptr + 76 + 4 + 52, undefined);
            A.store.Ref(ptr + 76 + 4 + 56, undefined);
          } else {
            A.store.Bool(ptr + 76 + 4 + 60, true);
            A.store.Ref(ptr + 76 + 4 + 0, x["action"]["redirect"]["extensionPath"]);

            if (typeof x["action"]["redirect"]["transform"] === "undefined") {
              A.store.Bool(ptr + 76 + 4 + 4 + 44, false);
              A.store.Ref(ptr + 76 + 4 + 4 + 0, undefined);
              A.store.Ref(ptr + 76 + 4 + 4 + 4, undefined);
              A.store.Ref(ptr + 76 + 4 + 4 + 8, undefined);
              A.store.Ref(ptr + 76 + 4 + 4 + 12, undefined);
              A.store.Ref(ptr + 76 + 4 + 4 + 16, undefined);

              A.store.Bool(ptr + 76 + 4 + 4 + 20 + 8, false);
              A.store.Ref(ptr + 76 + 4 + 4 + 20 + 0, undefined);
              A.store.Ref(ptr + 76 + 4 + 4 + 20 + 4, undefined);
              A.store.Ref(ptr + 76 + 4 + 4 + 32, undefined);
              A.store.Ref(ptr + 76 + 4 + 4 + 36, undefined);
              A.store.Ref(ptr + 76 + 4 + 4 + 40, undefined);
            } else {
              A.store.Bool(ptr + 76 + 4 + 4 + 44, true);
              A.store.Ref(ptr + 76 + 4 + 4 + 0, x["action"]["redirect"]["transform"]["scheme"]);
              A.store.Ref(ptr + 76 + 4 + 4 + 4, x["action"]["redirect"]["transform"]["host"]);
              A.store.Ref(ptr + 76 + 4 + 4 + 8, x["action"]["redirect"]["transform"]["port"]);
              A.store.Ref(ptr + 76 + 4 + 4 + 12, x["action"]["redirect"]["transform"]["path"]);
              A.store.Ref(ptr + 76 + 4 + 4 + 16, x["action"]["redirect"]["transform"]["query"]);

              if (typeof x["action"]["redirect"]["transform"]["queryTransform"] === "undefined") {
                A.store.Bool(ptr + 76 + 4 + 4 + 20 + 8, false);
                A.store.Ref(ptr + 76 + 4 + 4 + 20 + 0, undefined);
                A.store.Ref(ptr + 76 + 4 + 4 + 20 + 4, undefined);
              } else {
                A.store.Bool(ptr + 76 + 4 + 4 + 20 + 8, true);
                A.store.Ref(
                  ptr + 76 + 4 + 4 + 20 + 0,
                  x["action"]["redirect"]["transform"]["queryTransform"]["removeParams"]
                );
                A.store.Ref(
                  ptr + 76 + 4 + 4 + 20 + 4,
                  x["action"]["redirect"]["transform"]["queryTransform"]["addOrReplaceParams"]
                );
              }
              A.store.Ref(ptr + 76 + 4 + 4 + 32, x["action"]["redirect"]["transform"]["fragment"]);
              A.store.Ref(ptr + 76 + 4 + 4 + 36, x["action"]["redirect"]["transform"]["username"]);
              A.store.Ref(ptr + 76 + 4 + 4 + 40, x["action"]["redirect"]["transform"]["password"]);
            }
            A.store.Ref(ptr + 76 + 4 + 52, x["action"]["redirect"]["url"]);
            A.store.Ref(ptr + 76 + 4 + 56, x["action"]["redirect"]["regexSubstitution"]);
          }
          A.store.Ref(ptr + 76 + 68, x["action"]["requestHeaders"]);
          A.store.Ref(ptr + 76 + 72, x["action"]["responseHeaders"]);
        }
      }
    },
    "load_Rule": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 153)) {
        x["id"] = A.load.Int32(ptr + 0);
      } else {
        delete x["id"];
      }
      if (A.load.Bool(ptr + 154)) {
        x["priority"] = A.load.Int32(ptr + 4);
      } else {
        delete x["priority"];
      }
      if (A.load.Bool(ptr + 8 + 65)) {
        x["condition"] = {};
        x["condition"]["urlFilter"] = A.load.Ref(ptr + 8 + 0, undefined);
        x["condition"]["regexFilter"] = A.load.Ref(ptr + 8 + 4, undefined);
        if (A.load.Bool(ptr + 8 + 64)) {
          x["condition"]["isUrlFilterCaseSensitive"] = A.load.Bool(ptr + 8 + 8);
        } else {
          delete x["condition"]["isUrlFilterCaseSensitive"];
        }
        x["condition"]["initiatorDomains"] = A.load.Ref(ptr + 8 + 12, undefined);
        x["condition"]["excludedInitiatorDomains"] = A.load.Ref(ptr + 8 + 16, undefined);
        x["condition"]["requestDomains"] = A.load.Ref(ptr + 8 + 20, undefined);
        x["condition"]["excludedRequestDomains"] = A.load.Ref(ptr + 8 + 24, undefined);
        x["condition"]["domains"] = A.load.Ref(ptr + 8 + 28, undefined);
        x["condition"]["excludedDomains"] = A.load.Ref(ptr + 8 + 32, undefined);
        x["condition"]["resourceTypes"] = A.load.Ref(ptr + 8 + 36, undefined);
        x["condition"]["excludedResourceTypes"] = A.load.Ref(ptr + 8 + 40, undefined);
        x["condition"]["requestMethods"] = A.load.Ref(ptr + 8 + 44, undefined);
        x["condition"]["excludedRequestMethods"] = A.load.Ref(ptr + 8 + 48, undefined);
        x["condition"]["domainType"] = A.load.Enum(ptr + 8 + 52, ["firstParty", "thirdParty"]);
        x["condition"]["tabIds"] = A.load.Ref(ptr + 8 + 56, undefined);
        x["condition"]["excludedTabIds"] = A.load.Ref(ptr + 8 + 60, undefined);
      } else {
        delete x["condition"];
      }
      if (A.load.Bool(ptr + 76 + 76)) {
        x["action"] = {};
        x["action"]["type"] = A.load.Enum(ptr + 76 + 0, [
          "block",
          "redirect",
          "allow",
          "upgradeScheme",
          "modifyHeaders",
          "allowAllRequests",
        ]);
        if (A.load.Bool(ptr + 76 + 4 + 60)) {
          x["action"]["redirect"] = {};
          x["action"]["redirect"]["extensionPath"] = A.load.Ref(ptr + 76 + 4 + 0, undefined);
          if (A.load.Bool(ptr + 76 + 4 + 4 + 44)) {
            x["action"]["redirect"]["transform"] = {};
            x["action"]["redirect"]["transform"]["scheme"] = A.load.Ref(ptr + 76 + 4 + 4 + 0, undefined);
            x["action"]["redirect"]["transform"]["host"] = A.load.Ref(ptr + 76 + 4 + 4 + 4, undefined);
            x["action"]["redirect"]["transform"]["port"] = A.load.Ref(ptr + 76 + 4 + 4 + 8, undefined);
            x["action"]["redirect"]["transform"]["path"] = A.load.Ref(ptr + 76 + 4 + 4 + 12, undefined);
            x["action"]["redirect"]["transform"]["query"] = A.load.Ref(ptr + 76 + 4 + 4 + 16, undefined);
            if (A.load.Bool(ptr + 76 + 4 + 4 + 20 + 8)) {
              x["action"]["redirect"]["transform"]["queryTransform"] = {};
              x["action"]["redirect"]["transform"]["queryTransform"]["removeParams"] = A.load.Ref(
                ptr + 76 + 4 + 4 + 20 + 0,
                undefined
              );
              x["action"]["redirect"]["transform"]["queryTransform"]["addOrReplaceParams"] = A.load.Ref(
                ptr + 76 + 4 + 4 + 20 + 4,
                undefined
              );
            } else {
              delete x["action"]["redirect"]["transform"]["queryTransform"];
            }
            x["action"]["redirect"]["transform"]["fragment"] = A.load.Ref(ptr + 76 + 4 + 4 + 32, undefined);
            x["action"]["redirect"]["transform"]["username"] = A.load.Ref(ptr + 76 + 4 + 4 + 36, undefined);
            x["action"]["redirect"]["transform"]["password"] = A.load.Ref(ptr + 76 + 4 + 4 + 40, undefined);
          } else {
            delete x["action"]["redirect"]["transform"];
          }
          x["action"]["redirect"]["url"] = A.load.Ref(ptr + 76 + 4 + 52, undefined);
          x["action"]["redirect"]["regexSubstitution"] = A.load.Ref(ptr + 76 + 4 + 56, undefined);
        } else {
          delete x["action"]["redirect"];
        }
        x["action"]["requestHeaders"] = A.load.Ref(ptr + 76 + 68, undefined);
        x["action"]["responseHeaders"] = A.load.Ref(ptr + 76 + 72, undefined);
      } else {
        delete x["action"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetRulesFilter": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["ruleIds"]);
      }
    },
    "load_GetRulesFilter": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["ruleIds"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_UnsupportedRegexReason": (ref: heap.Ref<string>): number => {
      const idx = ["syntaxError", "memoryLimitExceeded"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_IsRegexSupportedResult": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 0, false);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "isSupported" in x ? true : false);
        A.store.Bool(ptr + 0, x["isSupported"] ? true : false);
        A.store.Enum(ptr + 4, ["syntaxError", "memoryLimitExceeded"].indexOf(x["reason"] as string));
      }
    },
    "load_IsRegexSupportedResult": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["isSupported"] = A.load.Bool(ptr + 0);
      } else {
        delete x["isSupported"];
      }
      x["reason"] = A.load.Enum(ptr + 4, ["syntaxError", "memoryLimitExceeded"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ManifestKeys": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 5, false);

        A.store.Bool(ptr + 0 + 4, false);
        A.store.Ref(ptr + 0 + 0, undefined);
      } else {
        A.store.Bool(ptr + 5, true);

        if (typeof x["declarative_net_request"] === "undefined") {
          A.store.Bool(ptr + 0 + 4, false);
          A.store.Ref(ptr + 0 + 0, undefined);
        } else {
          A.store.Bool(ptr + 0 + 4, true);
          A.store.Ref(ptr + 0 + 0, x["declarative_net_request"]["rule_resources"]);
        }
      }
    },
    "load_ManifestKeys": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 4)) {
        x["declarative_net_request"] = {};
        x["declarative_net_request"]["rule_resources"] = A.load.Ref(ptr + 0 + 0, undefined);
      } else {
        delete x["declarative_net_request"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RequestDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 51, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 48, false);
        A.store.Int32(ptr + 16, 0);
        A.store.Ref(ptr + 20, undefined);
        A.store.Enum(ptr + 24, -1);
        A.store.Enum(ptr + 28, -1);
        A.store.Bool(ptr + 49, false);
        A.store.Int32(ptr + 32, 0);
        A.store.Ref(ptr + 36, undefined);
        A.store.Bool(ptr + 50, false);
        A.store.Int32(ptr + 40, 0);
        A.store.Enum(ptr + 44, -1);
      } else {
        A.store.Bool(ptr + 51, true);
        A.store.Ref(ptr + 0, x["requestId"]);
        A.store.Ref(ptr + 4, x["url"]);
        A.store.Ref(ptr + 8, x["initiator"]);
        A.store.Ref(ptr + 12, x["method"]);
        A.store.Bool(ptr + 48, "frameId" in x ? true : false);
        A.store.Int32(ptr + 16, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Ref(ptr + 20, x["documentId"]);
        A.store.Enum(ptr + 24, ["outermost_frame", "fenced_frame", "sub_frame"].indexOf(x["frameType"] as string));
        A.store.Enum(
          ptr + 28,
          ["prerender", "active", "cached", "pending_deletion"].indexOf(x["documentLifecycle"] as string)
        );
        A.store.Bool(ptr + 49, "parentFrameId" in x ? true : false);
        A.store.Int32(ptr + 32, x["parentFrameId"] === undefined ? 0 : (x["parentFrameId"] as number));
        A.store.Ref(ptr + 36, x["parentDocumentId"]);
        A.store.Bool(ptr + 50, "tabId" in x ? true : false);
        A.store.Int32(ptr + 40, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Enum(
          ptr + 44,
          [
            "main_frame",
            "sub_frame",
            "stylesheet",
            "script",
            "image",
            "font",
            "object",
            "xmlhttprequest",
            "ping",
            "csp_report",
            "media",
            "websocket",
            "webtransport",
            "webbundle",
            "other",
          ].indexOf(x["type"] as string)
        );
      }
    },
    "load_RequestDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["requestId"] = A.load.Ref(ptr + 0, undefined);
      x["url"] = A.load.Ref(ptr + 4, undefined);
      x["initiator"] = A.load.Ref(ptr + 8, undefined);
      x["method"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 48)) {
        x["frameId"] = A.load.Int32(ptr + 16);
      } else {
        delete x["frameId"];
      }
      x["documentId"] = A.load.Ref(ptr + 20, undefined);
      x["frameType"] = A.load.Enum(ptr + 24, ["outermost_frame", "fenced_frame", "sub_frame"]);
      x["documentLifecycle"] = A.load.Enum(ptr + 28, ["prerender", "active", "cached", "pending_deletion"]);
      if (A.load.Bool(ptr + 49)) {
        x["parentFrameId"] = A.load.Int32(ptr + 32);
      } else {
        delete x["parentFrameId"];
      }
      x["parentDocumentId"] = A.load.Ref(ptr + 36, undefined);
      if (A.load.Bool(ptr + 50)) {
        x["tabId"] = A.load.Int32(ptr + 40);
      } else {
        delete x["tabId"];
      }
      x["type"] = A.load.Enum(ptr + 44, [
        "main_frame",
        "sub_frame",
        "stylesheet",
        "script",
        "image",
        "font",
        "object",
        "xmlhttprequest",
        "ping",
        "csp_report",
        "media",
        "websocket",
        "webtransport",
        "webbundle",
        "other",
      ]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_MatchedRuleInfoDebug": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 64, false);

        A.store.Bool(ptr + 0 + 9, false);
        A.store.Bool(ptr + 0 + 8, false);
        A.store.Int32(ptr + 0 + 0, 0);
        A.store.Ref(ptr + 0 + 4, undefined);

        A.store.Bool(ptr + 12 + 51, false);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 12 + 4, undefined);
        A.store.Ref(ptr + 12 + 8, undefined);
        A.store.Ref(ptr + 12 + 12, undefined);
        A.store.Bool(ptr + 12 + 48, false);
        A.store.Int32(ptr + 12 + 16, 0);
        A.store.Ref(ptr + 12 + 20, undefined);
        A.store.Enum(ptr + 12 + 24, -1);
        A.store.Enum(ptr + 12 + 28, -1);
        A.store.Bool(ptr + 12 + 49, false);
        A.store.Int32(ptr + 12 + 32, 0);
        A.store.Ref(ptr + 12 + 36, undefined);
        A.store.Bool(ptr + 12 + 50, false);
        A.store.Int32(ptr + 12 + 40, 0);
        A.store.Enum(ptr + 12 + 44, -1);
      } else {
        A.store.Bool(ptr + 64, true);

        if (typeof x["rule"] === "undefined") {
          A.store.Bool(ptr + 0 + 9, false);
          A.store.Bool(ptr + 0 + 8, false);
          A.store.Int32(ptr + 0 + 0, 0);
          A.store.Ref(ptr + 0 + 4, undefined);
        } else {
          A.store.Bool(ptr + 0 + 9, true);
          A.store.Bool(ptr + 0 + 8, "ruleId" in x["rule"] ? true : false);
          A.store.Int32(ptr + 0 + 0, x["rule"]["ruleId"] === undefined ? 0 : (x["rule"]["ruleId"] as number));
          A.store.Ref(ptr + 0 + 4, x["rule"]["rulesetId"]);
        }

        if (typeof x["request"] === "undefined") {
          A.store.Bool(ptr + 12 + 51, false);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Ref(ptr + 12 + 4, undefined);
          A.store.Ref(ptr + 12 + 8, undefined);
          A.store.Ref(ptr + 12 + 12, undefined);
          A.store.Bool(ptr + 12 + 48, false);
          A.store.Int32(ptr + 12 + 16, 0);
          A.store.Ref(ptr + 12 + 20, undefined);
          A.store.Enum(ptr + 12 + 24, -1);
          A.store.Enum(ptr + 12 + 28, -1);
          A.store.Bool(ptr + 12 + 49, false);
          A.store.Int32(ptr + 12 + 32, 0);
          A.store.Ref(ptr + 12 + 36, undefined);
          A.store.Bool(ptr + 12 + 50, false);
          A.store.Int32(ptr + 12 + 40, 0);
          A.store.Enum(ptr + 12 + 44, -1);
        } else {
          A.store.Bool(ptr + 12 + 51, true);
          A.store.Ref(ptr + 12 + 0, x["request"]["requestId"]);
          A.store.Ref(ptr + 12 + 4, x["request"]["url"]);
          A.store.Ref(ptr + 12 + 8, x["request"]["initiator"]);
          A.store.Ref(ptr + 12 + 12, x["request"]["method"]);
          A.store.Bool(ptr + 12 + 48, "frameId" in x["request"] ? true : false);
          A.store.Int32(ptr + 12 + 16, x["request"]["frameId"] === undefined ? 0 : (x["request"]["frameId"] as number));
          A.store.Ref(ptr + 12 + 20, x["request"]["documentId"]);
          A.store.Enum(
            ptr + 12 + 24,
            ["outermost_frame", "fenced_frame", "sub_frame"].indexOf(x["request"]["frameType"] as string)
          );
          A.store.Enum(
            ptr + 12 + 28,
            ["prerender", "active", "cached", "pending_deletion"].indexOf(x["request"]["documentLifecycle"] as string)
          );
          A.store.Bool(ptr + 12 + 49, "parentFrameId" in x["request"] ? true : false);
          A.store.Int32(
            ptr + 12 + 32,
            x["request"]["parentFrameId"] === undefined ? 0 : (x["request"]["parentFrameId"] as number)
          );
          A.store.Ref(ptr + 12 + 36, x["request"]["parentDocumentId"]);
          A.store.Bool(ptr + 12 + 50, "tabId" in x["request"] ? true : false);
          A.store.Int32(ptr + 12 + 40, x["request"]["tabId"] === undefined ? 0 : (x["request"]["tabId"] as number));
          A.store.Enum(
            ptr + 12 + 44,
            [
              "main_frame",
              "sub_frame",
              "stylesheet",
              "script",
              "image",
              "font",
              "object",
              "xmlhttprequest",
              "ping",
              "csp_report",
              "media",
              "websocket",
              "webtransport",
              "webbundle",
              "other",
            ].indexOf(x["request"]["type"] as string)
          );
        }
      }
    },
    "load_MatchedRuleInfoDebug": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 9)) {
        x["rule"] = {};
        if (A.load.Bool(ptr + 0 + 8)) {
          x["rule"]["ruleId"] = A.load.Int32(ptr + 0 + 0);
        } else {
          delete x["rule"]["ruleId"];
        }
        x["rule"]["rulesetId"] = A.load.Ref(ptr + 0 + 4, undefined);
      } else {
        delete x["rule"];
      }
      if (A.load.Bool(ptr + 12 + 51)) {
        x["request"] = {};
        x["request"]["requestId"] = A.load.Ref(ptr + 12 + 0, undefined);
        x["request"]["url"] = A.load.Ref(ptr + 12 + 4, undefined);
        x["request"]["initiator"] = A.load.Ref(ptr + 12 + 8, undefined);
        x["request"]["method"] = A.load.Ref(ptr + 12 + 12, undefined);
        if (A.load.Bool(ptr + 12 + 48)) {
          x["request"]["frameId"] = A.load.Int32(ptr + 12 + 16);
        } else {
          delete x["request"]["frameId"];
        }
        x["request"]["documentId"] = A.load.Ref(ptr + 12 + 20, undefined);
        x["request"]["frameType"] = A.load.Enum(ptr + 12 + 24, ["outermost_frame", "fenced_frame", "sub_frame"]);
        x["request"]["documentLifecycle"] = A.load.Enum(ptr + 12 + 28, [
          "prerender",
          "active",
          "cached",
          "pending_deletion",
        ]);
        if (A.load.Bool(ptr + 12 + 49)) {
          x["request"]["parentFrameId"] = A.load.Int32(ptr + 12 + 32);
        } else {
          delete x["request"]["parentFrameId"];
        }
        x["request"]["parentDocumentId"] = A.load.Ref(ptr + 12 + 36, undefined);
        if (A.load.Bool(ptr + 12 + 50)) {
          x["request"]["tabId"] = A.load.Int32(ptr + 12 + 40);
        } else {
          delete x["request"]["tabId"];
        }
        x["request"]["type"] = A.load.Enum(ptr + 12 + 44, [
          "main_frame",
          "sub_frame",
          "stylesheet",
          "script",
          "image",
          "font",
          "object",
          "xmlhttprequest",
          "ping",
          "csp_report",
          "media",
          "websocket",
          "webtransport",
          "webbundle",
          "other",
        ]);
      } else {
        delete x["request"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_MatchedRulesFilter": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 17, false);
        A.store.Float64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 18, true);
        A.store.Bool(ptr + 16, "tabId" in x ? true : false);
        A.store.Int32(ptr + 0, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Bool(ptr + 17, "minTimeStamp" in x ? true : false);
        A.store.Float64(ptr + 8, x["minTimeStamp"] === undefined ? 0 : (x["minTimeStamp"] as number));
      }
    },
    "load_MatchedRulesFilter": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["tabId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["tabId"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["minTimeStamp"] = A.load.Float64(ptr + 8);
      } else {
        delete x["minTimeStamp"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "has_Properties_GUARANTEED_MINIMUM_STATIC_RULES": (self: heap.Ref<any>): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeNetRequest && "GUARANTEED_MINIMUM_STATIC_RULES" in WEBEXT?.declarativeNetRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Properties_GUARANTEED_MINIMUM_STATIC_RULES": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.GUARANTEED_MINIMUM_STATIC_RULES);
    },

    "call_Properties_GUARANTEED_MINIMUM_STATIC_RULES": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = Reflect.apply(WEBEXT.declarativeNetRequest.GUARANTEED_MINIMUM_STATIC_RULES, thiz, []);
      A.store.Int32(retPtr, _ret);
    },
    "try_Properties_GUARANTEED_MINIMUM_STATIC_RULES": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = Reflect.apply(WEBEXT.declarativeNetRequest.GUARANTEED_MINIMUM_STATIC_RULES, thiz, []);
        A.store.Int32(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Properties_MAX_NUMBER_OF_DYNAMIC_RULES": (self: heap.Ref<any>): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeNetRequest && "MAX_NUMBER_OF_DYNAMIC_RULES" in WEBEXT?.declarativeNetRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Properties_MAX_NUMBER_OF_DYNAMIC_RULES": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.MAX_NUMBER_OF_DYNAMIC_RULES);
    },

    "call_Properties_MAX_NUMBER_OF_DYNAMIC_RULES": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = Reflect.apply(WEBEXT.declarativeNetRequest.MAX_NUMBER_OF_DYNAMIC_RULES, thiz, []);
      A.store.Int32(retPtr, _ret);
    },
    "try_Properties_MAX_NUMBER_OF_DYNAMIC_RULES": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = Reflect.apply(WEBEXT.declarativeNetRequest.MAX_NUMBER_OF_DYNAMIC_RULES, thiz, []);
        A.store.Int32(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Properties_MAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES": (self: heap.Ref<any>): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeNetRequest && "MAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES" in WEBEXT?.declarativeNetRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Properties_MAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.MAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES);
    },

    "call_Properties_MAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = Reflect.apply(WEBEXT.declarativeNetRequest.MAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES, thiz, []);
      A.store.Int32(retPtr, _ret);
    },
    "try_Properties_MAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = Reflect.apply(WEBEXT.declarativeNetRequest.MAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES, thiz, []);
        A.store.Int32(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Properties_GETMATCHEDRULES_QUOTA_INTERVAL": (self: heap.Ref<any>): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeNetRequest && "GETMATCHEDRULES_QUOTA_INTERVAL" in WEBEXT?.declarativeNetRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Properties_GETMATCHEDRULES_QUOTA_INTERVAL": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.GETMATCHEDRULES_QUOTA_INTERVAL);
    },

    "call_Properties_GETMATCHEDRULES_QUOTA_INTERVAL": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = Reflect.apply(WEBEXT.declarativeNetRequest.GETMATCHEDRULES_QUOTA_INTERVAL, thiz, []);
      A.store.Int32(retPtr, _ret);
    },
    "try_Properties_GETMATCHEDRULES_QUOTA_INTERVAL": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = Reflect.apply(WEBEXT.declarativeNetRequest.GETMATCHEDRULES_QUOTA_INTERVAL, thiz, []);
        A.store.Int32(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Properties_MAX_GETMATCHEDRULES_CALLS_PER_INTERVAL": (self: heap.Ref<any>): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeNetRequest && "MAX_GETMATCHEDRULES_CALLS_PER_INTERVAL" in WEBEXT?.declarativeNetRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Properties_MAX_GETMATCHEDRULES_CALLS_PER_INTERVAL": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.MAX_GETMATCHEDRULES_CALLS_PER_INTERVAL);
    },

    "call_Properties_MAX_GETMATCHEDRULES_CALLS_PER_INTERVAL": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = Reflect.apply(WEBEXT.declarativeNetRequest.MAX_GETMATCHEDRULES_CALLS_PER_INTERVAL, thiz, []);
      A.store.Int32(retPtr, _ret);
    },
    "try_Properties_MAX_GETMATCHEDRULES_CALLS_PER_INTERVAL": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = Reflect.apply(WEBEXT.declarativeNetRequest.MAX_GETMATCHEDRULES_CALLS_PER_INTERVAL, thiz, []);
        A.store.Int32(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Properties_MAX_NUMBER_OF_REGEX_RULES": (self: heap.Ref<any>): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeNetRequest && "MAX_NUMBER_OF_REGEX_RULES" in WEBEXT?.declarativeNetRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Properties_MAX_NUMBER_OF_REGEX_RULES": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.MAX_NUMBER_OF_REGEX_RULES);
    },

    "call_Properties_MAX_NUMBER_OF_REGEX_RULES": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = Reflect.apply(WEBEXT.declarativeNetRequest.MAX_NUMBER_OF_REGEX_RULES, thiz, []);
      A.store.Int32(retPtr, _ret);
    },
    "try_Properties_MAX_NUMBER_OF_REGEX_RULES": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = Reflect.apply(WEBEXT.declarativeNetRequest.MAX_NUMBER_OF_REGEX_RULES, thiz, []);
        A.store.Int32(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Properties_MAX_NUMBER_OF_STATIC_RULESETS": (self: heap.Ref<any>): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeNetRequest && "MAX_NUMBER_OF_STATIC_RULESETS" in WEBEXT?.declarativeNetRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Properties_MAX_NUMBER_OF_STATIC_RULESETS": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.MAX_NUMBER_OF_STATIC_RULESETS);
    },

    "call_Properties_MAX_NUMBER_OF_STATIC_RULESETS": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = Reflect.apply(WEBEXT.declarativeNetRequest.MAX_NUMBER_OF_STATIC_RULESETS, thiz, []);
      A.store.Int32(retPtr, _ret);
    },
    "try_Properties_MAX_NUMBER_OF_STATIC_RULESETS": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = Reflect.apply(WEBEXT.declarativeNetRequest.MAX_NUMBER_OF_STATIC_RULESETS, thiz, []);
        A.store.Int32(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Properties_MAX_NUMBER_OF_ENABLED_STATIC_RULESETS": (self: heap.Ref<any>): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeNetRequest && "MAX_NUMBER_OF_ENABLED_STATIC_RULESETS" in WEBEXT?.declarativeNetRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Properties_MAX_NUMBER_OF_ENABLED_STATIC_RULESETS": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.MAX_NUMBER_OF_ENABLED_STATIC_RULESETS);
    },

    "call_Properties_MAX_NUMBER_OF_ENABLED_STATIC_RULESETS": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = Reflect.apply(WEBEXT.declarativeNetRequest.MAX_NUMBER_OF_ENABLED_STATIC_RULESETS, thiz, []);
      A.store.Int32(retPtr, _ret);
    },
    "try_Properties_MAX_NUMBER_OF_ENABLED_STATIC_RULESETS": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = Reflect.apply(WEBEXT.declarativeNetRequest.MAX_NUMBER_OF_ENABLED_STATIC_RULESETS, thiz, []);
        A.store.Int32(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Properties_DYNAMIC_RULESET_ID": (self: heap.Ref<any>): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeNetRequest && "DYNAMIC_RULESET_ID" in WEBEXT?.declarativeNetRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Properties_DYNAMIC_RULESET_ID": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.DYNAMIC_RULESET_ID);
    },

    "call_Properties_DYNAMIC_RULESET_ID": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = Reflect.apply(WEBEXT.declarativeNetRequest.DYNAMIC_RULESET_ID, thiz, []);
      A.store.Ref(retPtr, _ret);
    },
    "try_Properties_DYNAMIC_RULESET_ID": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = Reflect.apply(WEBEXT.declarativeNetRequest.DYNAMIC_RULESET_ID, thiz, []);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Properties_SESSION_RULESET_ID": (self: heap.Ref<any>): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeNetRequest && "SESSION_RULESET_ID" in WEBEXT?.declarativeNetRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Properties_SESSION_RULESET_ID": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.SESSION_RULESET_ID);
    },

    "call_Properties_SESSION_RULESET_ID": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = Reflect.apply(WEBEXT.declarativeNetRequest.SESSION_RULESET_ID, thiz, []);
      A.store.Ref(retPtr, _ret);
    },
    "try_Properties_SESSION_RULESET_ID": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = Reflect.apply(WEBEXT.declarativeNetRequest.SESSION_RULESET_ID, thiz, []);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },

    "store_RegexOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 6, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 7, false);
        A.store.Bool(ptr + 5, false);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["regex"]);
        A.store.Bool(ptr + 6, "isCaseSensitive" in x ? true : false);
        A.store.Bool(ptr + 4, x["isCaseSensitive"] ? true : false);
        A.store.Bool(ptr + 7, "requireCapturing" in x ? true : false);
        A.store.Bool(ptr + 5, x["requireCapturing"] ? true : false);
      }
    },
    "load_RegexOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["regex"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 6)) {
        x["isCaseSensitive"] = A.load.Bool(ptr + 4);
      } else {
        delete x["isCaseSensitive"];
      }
      if (A.load.Bool(ptr + 7)) {
        x["requireCapturing"] = A.load.Bool(ptr + 5);
      } else {
        delete x["requireCapturing"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_TestMatchOutcomeResult": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["matchedRules"]);
      }
    },
    "load_TestMatchOutcomeResult": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["matchedRules"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_TestMatchRequestDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 21, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
        A.store.Enum(ptr + 12, -1);
        A.store.Bool(ptr + 20, false);
        A.store.Int32(ptr + 16, 0);
      } else {
        A.store.Bool(ptr + 21, true);
        A.store.Ref(ptr + 0, x["url"]);
        A.store.Ref(ptr + 4, x["initiator"]);
        A.store.Enum(
          ptr + 8,
          ["connect", "delete", "get", "head", "options", "patch", "post", "put", "other"].indexOf(
            x["method"] as string
          )
        );
        A.store.Enum(
          ptr + 12,
          [
            "main_frame",
            "sub_frame",
            "stylesheet",
            "script",
            "image",
            "font",
            "object",
            "xmlhttprequest",
            "ping",
            "csp_report",
            "media",
            "websocket",
            "webtransport",
            "webbundle",
            "other",
          ].indexOf(x["type"] as string)
        );
        A.store.Bool(ptr + 20, "tabId" in x ? true : false);
        A.store.Int32(ptr + 16, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
      }
    },
    "load_TestMatchRequestDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["url"] = A.load.Ref(ptr + 0, undefined);
      x["initiator"] = A.load.Ref(ptr + 4, undefined);
      x["method"] = A.load.Enum(ptr + 8, [
        "connect",
        "delete",
        "get",
        "head",
        "options",
        "patch",
        "post",
        "put",
        "other",
      ]);
      x["type"] = A.load.Enum(ptr + 12, [
        "main_frame",
        "sub_frame",
        "stylesheet",
        "script",
        "image",
        "font",
        "object",
        "xmlhttprequest",
        "ping",
        "csp_report",
        "media",
        "websocket",
        "webtransport",
        "webbundle",
        "other",
      ]);
      if (A.load.Bool(ptr + 20)) {
        x["tabId"] = A.load.Int32(ptr + 16);
      } else {
        delete x["tabId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UpdateRuleOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["removeRuleIds"]);
        A.store.Ref(ptr + 4, x["addRules"]);
      }
    },
    "load_UpdateRuleOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["removeRuleIds"] = A.load.Ref(ptr + 0, undefined);
      x["addRules"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UpdateRulesetOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["disableRulesetIds"]);
        A.store.Ref(ptr + 4, x["enableRulesetIds"]);
      }
    },
    "load_UpdateRulesetOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["disableRulesetIds"] = A.load.Ref(ptr + 0, undefined);
      x["enableRulesetIds"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UpdateStaticRulesOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["rulesetId"]);
        A.store.Ref(ptr + 4, x["disableRuleIds"]);
        A.store.Ref(ptr + 8, x["enableRuleIds"]);
      }
    },
    "load_UpdateStaticRulesOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["rulesetId"] = A.load.Ref(ptr + 0, undefined);
      x["disableRuleIds"] = A.load.Ref(ptr + 4, undefined);
      x["enableRuleIds"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetAvailableStaticRuleCount": (): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeNetRequest && "getAvailableStaticRuleCount" in WEBEXT?.declarativeNetRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAvailableStaticRuleCount": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.getAvailableStaticRuleCount);
    },
    "call_GetAvailableStaticRuleCount": (retPtr: Pointer): void => {
      const _ret = WEBEXT.declarativeNetRequest.getAvailableStaticRuleCount();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAvailableStaticRuleCount": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.declarativeNetRequest.getAvailableStaticRuleCount();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDisabledRuleIds": (): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeNetRequest && "getDisabledRuleIds" in WEBEXT?.declarativeNetRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDisabledRuleIds": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.getDisabledRuleIds);
    },
    "call_GetDisabledRuleIds": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["rulesetId"] = A.load.Ref(options + 0, undefined);

      const _ret = WEBEXT.declarativeNetRequest.getDisabledRuleIds(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDisabledRuleIds": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["rulesetId"] = A.load.Ref(options + 0, undefined);

        const _ret = WEBEXT.declarativeNetRequest.getDisabledRuleIds(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDynamicRules": (): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeNetRequest && "getDynamicRules" in WEBEXT?.declarativeNetRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDynamicRules": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.getDynamicRules);
    },
    "call_GetDynamicRules": (retPtr: Pointer, filter: Pointer): void => {
      const filter_ffi = {};

      filter_ffi["ruleIds"] = A.load.Ref(filter + 0, undefined);

      const _ret = WEBEXT.declarativeNetRequest.getDynamicRules(filter_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDynamicRules": (retPtr: Pointer, errPtr: Pointer, filter: Pointer): heap.Ref<boolean> => {
      try {
        const filter_ffi = {};

        filter_ffi["ruleIds"] = A.load.Ref(filter + 0, undefined);

        const _ret = WEBEXT.declarativeNetRequest.getDynamicRules(filter_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetEnabledRulesets": (): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeNetRequest && "getEnabledRulesets" in WEBEXT?.declarativeNetRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetEnabledRulesets": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.getEnabledRulesets);
    },
    "call_GetEnabledRulesets": (retPtr: Pointer): void => {
      const _ret = WEBEXT.declarativeNetRequest.getEnabledRulesets();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetEnabledRulesets": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.declarativeNetRequest.getEnabledRulesets();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetMatchedRules": (): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeNetRequest && "getMatchedRules" in WEBEXT?.declarativeNetRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetMatchedRules": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.getMatchedRules);
    },
    "call_GetMatchedRules": (retPtr: Pointer, filter: Pointer): void => {
      const filter_ffi = {};

      if (A.load.Bool(filter + 16)) {
        filter_ffi["tabId"] = A.load.Int32(filter + 0);
      }
      if (A.load.Bool(filter + 17)) {
        filter_ffi["minTimeStamp"] = A.load.Float64(filter + 8);
      }

      const _ret = WEBEXT.declarativeNetRequest.getMatchedRules(filter_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetMatchedRules": (retPtr: Pointer, errPtr: Pointer, filter: Pointer): heap.Ref<boolean> => {
      try {
        const filter_ffi = {};

        if (A.load.Bool(filter + 16)) {
          filter_ffi["tabId"] = A.load.Int32(filter + 0);
        }
        if (A.load.Bool(filter + 17)) {
          filter_ffi["minTimeStamp"] = A.load.Float64(filter + 8);
        }

        const _ret = WEBEXT.declarativeNetRequest.getMatchedRules(filter_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetSessionRules": (): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeNetRequest && "getSessionRules" in WEBEXT?.declarativeNetRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetSessionRules": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.getSessionRules);
    },
    "call_GetSessionRules": (retPtr: Pointer, filter: Pointer): void => {
      const filter_ffi = {};

      filter_ffi["ruleIds"] = A.load.Ref(filter + 0, undefined);

      const _ret = WEBEXT.declarativeNetRequest.getSessionRules(filter_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetSessionRules": (retPtr: Pointer, errPtr: Pointer, filter: Pointer): heap.Ref<boolean> => {
      try {
        const filter_ffi = {};

        filter_ffi["ruleIds"] = A.load.Ref(filter + 0, undefined);

        const _ret = WEBEXT.declarativeNetRequest.getSessionRules(filter_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsRegexSupported": (): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeNetRequest && "isRegexSupported" in WEBEXT?.declarativeNetRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsRegexSupported": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.isRegexSupported);
    },
    "call_IsRegexSupported": (retPtr: Pointer, regexOptions: Pointer): void => {
      const regexOptions_ffi = {};

      regexOptions_ffi["regex"] = A.load.Ref(regexOptions + 0, undefined);
      if (A.load.Bool(regexOptions + 6)) {
        regexOptions_ffi["isCaseSensitive"] = A.load.Bool(regexOptions + 4);
      }
      if (A.load.Bool(regexOptions + 7)) {
        regexOptions_ffi["requireCapturing"] = A.load.Bool(regexOptions + 5);
      }

      const _ret = WEBEXT.declarativeNetRequest.isRegexSupported(regexOptions_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_IsRegexSupported": (retPtr: Pointer, errPtr: Pointer, regexOptions: Pointer): heap.Ref<boolean> => {
      try {
        const regexOptions_ffi = {};

        regexOptions_ffi["regex"] = A.load.Ref(regexOptions + 0, undefined);
        if (A.load.Bool(regexOptions + 6)) {
          regexOptions_ffi["isCaseSensitive"] = A.load.Bool(regexOptions + 4);
        }
        if (A.load.Bool(regexOptions + 7)) {
          regexOptions_ffi["requireCapturing"] = A.load.Bool(regexOptions + 5);
        }

        const _ret = WEBEXT.declarativeNetRequest.isRegexSupported(regexOptions_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnRuleMatchedDebug": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.declarativeNetRequest?.onRuleMatchedDebug &&
        "addListener" in WEBEXT?.declarativeNetRequest?.onRuleMatchedDebug
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnRuleMatchedDebug": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.onRuleMatchedDebug.addListener);
    },
    "call_OnRuleMatchedDebug": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.declarativeNetRequest.onRuleMatchedDebug.addListener(A.H.get<object>(callback));
    },
    "try_OnRuleMatchedDebug": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.declarativeNetRequest.onRuleMatchedDebug.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffRuleMatchedDebug": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.declarativeNetRequest?.onRuleMatchedDebug &&
        "removeListener" in WEBEXT?.declarativeNetRequest?.onRuleMatchedDebug
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffRuleMatchedDebug": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.onRuleMatchedDebug.removeListener);
    },
    "call_OffRuleMatchedDebug": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.declarativeNetRequest.onRuleMatchedDebug.removeListener(A.H.get<object>(callback));
    },
    "try_OffRuleMatchedDebug": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.declarativeNetRequest.onRuleMatchedDebug.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnRuleMatchedDebug": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.declarativeNetRequest?.onRuleMatchedDebug &&
        "hasListener" in WEBEXT?.declarativeNetRequest?.onRuleMatchedDebug
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnRuleMatchedDebug": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.onRuleMatchedDebug.hasListener);
    },
    "call_HasOnRuleMatchedDebug": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.declarativeNetRequest.onRuleMatchedDebug.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnRuleMatchedDebug": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.declarativeNetRequest.onRuleMatchedDebug.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetExtensionActionOptions": (): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeNetRequest && "setExtensionActionOptions" in WEBEXT?.declarativeNetRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetExtensionActionOptions": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.setExtensionActionOptions);
    },
    "call_SetExtensionActionOptions": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 15)) {
        options_ffi["displayActionCountAsBadgeText"] = A.load.Bool(options + 0);
      }
      if (A.load.Bool(options + 4 + 10)) {
        options_ffi["tabUpdate"] = {};
        if (A.load.Bool(options + 4 + 8)) {
          options_ffi["tabUpdate"]["tabId"] = A.load.Int32(options + 4 + 0);
        }
        if (A.load.Bool(options + 4 + 9)) {
          options_ffi["tabUpdate"]["increment"] = A.load.Int32(options + 4 + 4);
        }
      }

      const _ret = WEBEXT.declarativeNetRequest.setExtensionActionOptions(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetExtensionActionOptions": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 15)) {
          options_ffi["displayActionCountAsBadgeText"] = A.load.Bool(options + 0);
        }
        if (A.load.Bool(options + 4 + 10)) {
          options_ffi["tabUpdate"] = {};
          if (A.load.Bool(options + 4 + 8)) {
            options_ffi["tabUpdate"]["tabId"] = A.load.Int32(options + 4 + 0);
          }
          if (A.load.Bool(options + 4 + 9)) {
            options_ffi["tabUpdate"]["increment"] = A.load.Int32(options + 4 + 4);
          }
        }

        const _ret = WEBEXT.declarativeNetRequest.setExtensionActionOptions(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_TestMatchOutcome": (): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeNetRequest && "testMatchOutcome" in WEBEXT?.declarativeNetRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_TestMatchOutcome": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.testMatchOutcome);
    },
    "call_TestMatchOutcome": (retPtr: Pointer, request: Pointer): void => {
      const request_ffi = {};

      request_ffi["url"] = A.load.Ref(request + 0, undefined);
      request_ffi["initiator"] = A.load.Ref(request + 4, undefined);
      request_ffi["method"] = A.load.Enum(request + 8, [
        "connect",
        "delete",
        "get",
        "head",
        "options",
        "patch",
        "post",
        "put",
        "other",
      ]);
      request_ffi["type"] = A.load.Enum(request + 12, [
        "main_frame",
        "sub_frame",
        "stylesheet",
        "script",
        "image",
        "font",
        "object",
        "xmlhttprequest",
        "ping",
        "csp_report",
        "media",
        "websocket",
        "webtransport",
        "webbundle",
        "other",
      ]);
      if (A.load.Bool(request + 20)) {
        request_ffi["tabId"] = A.load.Int32(request + 16);
      }

      const _ret = WEBEXT.declarativeNetRequest.testMatchOutcome(request_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_TestMatchOutcome": (retPtr: Pointer, errPtr: Pointer, request: Pointer): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        request_ffi["url"] = A.load.Ref(request + 0, undefined);
        request_ffi["initiator"] = A.load.Ref(request + 4, undefined);
        request_ffi["method"] = A.load.Enum(request + 8, [
          "connect",
          "delete",
          "get",
          "head",
          "options",
          "patch",
          "post",
          "put",
          "other",
        ]);
        request_ffi["type"] = A.load.Enum(request + 12, [
          "main_frame",
          "sub_frame",
          "stylesheet",
          "script",
          "image",
          "font",
          "object",
          "xmlhttprequest",
          "ping",
          "csp_report",
          "media",
          "websocket",
          "webtransport",
          "webbundle",
          "other",
        ]);
        if (A.load.Bool(request + 20)) {
          request_ffi["tabId"] = A.load.Int32(request + 16);
        }

        const _ret = WEBEXT.declarativeNetRequest.testMatchOutcome(request_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UpdateDynamicRules": (): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeNetRequest && "updateDynamicRules" in WEBEXT?.declarativeNetRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UpdateDynamicRules": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.updateDynamicRules);
    },
    "call_UpdateDynamicRules": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["removeRuleIds"] = A.load.Ref(options + 0, undefined);
      options_ffi["addRules"] = A.load.Ref(options + 4, undefined);

      const _ret = WEBEXT.declarativeNetRequest.updateDynamicRules(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_UpdateDynamicRules": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["removeRuleIds"] = A.load.Ref(options + 0, undefined);
        options_ffi["addRules"] = A.load.Ref(options + 4, undefined);

        const _ret = WEBEXT.declarativeNetRequest.updateDynamicRules(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UpdateEnabledRulesets": (): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeNetRequest && "updateEnabledRulesets" in WEBEXT?.declarativeNetRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UpdateEnabledRulesets": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.updateEnabledRulesets);
    },
    "call_UpdateEnabledRulesets": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["disableRulesetIds"] = A.load.Ref(options + 0, undefined);
      options_ffi["enableRulesetIds"] = A.load.Ref(options + 4, undefined);

      const _ret = WEBEXT.declarativeNetRequest.updateEnabledRulesets(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_UpdateEnabledRulesets": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["disableRulesetIds"] = A.load.Ref(options + 0, undefined);
        options_ffi["enableRulesetIds"] = A.load.Ref(options + 4, undefined);

        const _ret = WEBEXT.declarativeNetRequest.updateEnabledRulesets(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UpdateSessionRules": (): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeNetRequest && "updateSessionRules" in WEBEXT?.declarativeNetRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UpdateSessionRules": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.updateSessionRules);
    },
    "call_UpdateSessionRules": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["removeRuleIds"] = A.load.Ref(options + 0, undefined);
      options_ffi["addRules"] = A.load.Ref(options + 4, undefined);

      const _ret = WEBEXT.declarativeNetRequest.updateSessionRules(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_UpdateSessionRules": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["removeRuleIds"] = A.load.Ref(options + 0, undefined);
        options_ffi["addRules"] = A.load.Ref(options + 4, undefined);

        const _ret = WEBEXT.declarativeNetRequest.updateSessionRules(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UpdateStaticRules": (): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeNetRequest && "updateStaticRules" in WEBEXT?.declarativeNetRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UpdateStaticRules": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeNetRequest.updateStaticRules);
    },
    "call_UpdateStaticRules": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["rulesetId"] = A.load.Ref(options + 0, undefined);
      options_ffi["disableRuleIds"] = A.load.Ref(options + 4, undefined);
      options_ffi["enableRuleIds"] = A.load.Ref(options + 8, undefined);

      const _ret = WEBEXT.declarativeNetRequest.updateStaticRules(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_UpdateStaticRules": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["rulesetId"] = A.load.Ref(options + 0, undefined);
        options_ffi["disableRuleIds"] = A.load.Ref(options + 4, undefined);
        options_ffi["enableRuleIds"] = A.load.Ref(options + 8, undefined);

        const _ret = WEBEXT.declarativeNetRequest.updateStaticRules(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
