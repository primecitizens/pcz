import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/autofillprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_AccountInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
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
        A.store.Ref(ptr + 0, x["email"]);
        A.store.Bool(ptr + 6, "isSyncEnabledForAutofillProfiles" in x ? true : false);
        A.store.Bool(ptr + 4, x["isSyncEnabledForAutofillProfiles"] ? true : false);
        A.store.Bool(ptr + 7, "isEligibleForAddressAccountStorage" in x ? true : false);
        A.store.Bool(ptr + 5, x["isEligibleForAddressAccountStorage"] ? true : false);
      }
    },
    "load_AccountInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["email"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 6)) {
        x["isSyncEnabledForAutofillProfiles"] = A.load.Bool(ptr + 4);
      } else {
        delete x["isSyncEnabledForAutofillProfiles"];
      }
      if (A.load.Bool(ptr + 7)) {
        x["isEligibleForAddressAccountStorage"] = A.load.Bool(ptr + 5);
      } else {
        delete x["isEligibleForAddressAccountStorage"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ServerFieldType": (ref: heap.Ref<string>): number => {
      const idx = [
        "NO_SERVER_DATA",
        "UNKNOWN_TYPE",
        "EMPTY_TYPE",
        "NAME_FIRST",
        "NAME_MIDDLE",
        "NAME_LAST",
        "NAME_MIDDLE_INITIAL",
        "NAME_FULL",
        "NAME_SUFFIX",
        "EMAIL_ADDRESS",
        "PHONE_HOME_NUMBER",
        "PHONE_HOME_CITY_CODE",
        "PHONE_HOME_COUNTRY_CODE",
        "PHONE_HOME_CITY_AND_NUMBER",
        "PHONE_HOME_WHOLE_NUMBER",
        "ADDRESS_HOME_LINE1",
        "ADDRESS_HOME_LINE2",
        "ADDRESS_HOME_APT_NUM",
        "ADDRESS_HOME_CITY",
        "ADDRESS_HOME_STATE",
        "ADDRESS_HOME_ZIP",
        "ADDRESS_HOME_COUNTRY",
        "CREDIT_CARD_NAME_FULL",
        "CREDIT_CARD_NUMBER",
        "CREDIT_CARD_EXP_MONTH",
        "CREDIT_CARD_EXP_2_DIGIT_YEAR",
        "CREDIT_CARD_EXP_4_DIGIT_YEAR",
        "CREDIT_CARD_EXP_DATE_2_DIGIT_YEAR",
        "CREDIT_CARD_EXP_DATE_4_DIGIT_YEAR",
        "CREDIT_CARD_TYPE",
        "CREDIT_CARD_VERIFICATION_CODE",
        "COMPANY_NAME",
        "FIELD_WITH_DEFAULT_VALUE",
        "MERCHANT_EMAIL_SIGNUP",
        "MERCHANT_PROMO_CODE",
        "PASSWORD",
        "ACCOUNT_CREATION_PASSWORD",
        "ADDRESS_HOME_STREET_ADDRESS",
        "ADDRESS_HOME_SORTING_CODE",
        "ADDRESS_HOME_DEPENDENT_LOCALITY",
        "ADDRESS_HOME_LINE3",
        "NOT_ACCOUNT_CREATION_PASSWORD",
        "USERNAME",
        "USERNAME_AND_EMAIL_ADDRESS",
        "NEW_PASSWORD",
        "PROBABLY_NEW_PASSWORD",
        "NOT_NEW_PASSWORD",
        "CREDIT_CARD_NAME_FIRST",
        "CREDIT_CARD_NAME_LAST",
        "PHONE_HOME_EXTENSION",
        "CONFIRMATION_PASSWORD",
        "AMBIGUOUS_TYPE",
        "SEARCH_TERM",
        "PRICE",
        "NOT_PASSWORD",
        "SINGLE_USERNAME",
        "NOT_USERNAME",
        "UPI_VPA",
        "ADDRESS_HOME_STREET_NAME",
        "ADDRESS_HOME_HOUSE_NUMBER",
        "ADDRESS_HOME_SUBPREMISE",
        "ADDRESS_HOME_OTHER_SUBUNIT",
        "NAME_LAST_FIRST",
        "NAME_LAST_CONJUNCTION",
        "NAME_LAST_SECOND",
        "NAME_HONORIFIC_PREFIX",
        "ADDRESS_HOME_ADDRESS",
        "ADDRESS_HOME_ADDRESS_WITH_NAME",
        "ADDRESS_HOME_FLOOR",
        "NAME_FULL_WITH_HONORIFIC_PREFIX",
        "BIRTHDATE_DAY",
        "BIRTHDATE_MONTH",
        "BIRTHDATE_4_DIGIT_YEAR",
        "PHONE_HOME_CITY_CODE_WITH_TRUNK_PREFIX",
        "PHONE_HOME_CITY_AND_NUMBER_WITHOUT_TRUNK_PREFIX",
        "PHONE_HOME_NUMBER_PREFIX",
        "PHONE_HOME_NUMBER_SUFFIX",
        "IBAN_VALUE",
        "CREDIT_CARD_STANDALONE_VERIFICATION_CODE",
        "NUMERIC_QUANTITY",
        "ONE_TIME_CODE",
        "DELIVERY_INSTRUCTIONS",
        "ADDRESS_HOME_OVERFLOW",
        "ADDRESS_HOME_LANDMARK",
        "ADDRESS_HOME_OVERFLOW_AND_LANDMARK",
        "ADDRESS_HOME_ADMIN_LEVEL2",
        "ADDRESS_HOME_STREET_LOCATION",
        "ADDRESS_HOME_BETWEEN_STREETS",
        "ADDRESS_HOME_BETWEEN_STREETS_OR_LANDMARK",
        "ADDRESS_HOME_BETWEEN_STREETS_1",
        "ADDRESS_HOME_BETWEEN_STREETS_2",
        "SINGLE_USERNAME_FORGOT_PASSWORD",
        "MAX_VALID_FIELD_TYPE",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_AddressComponent": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 18, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 9, false);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 18, true);
        A.store.Enum(
          ptr + 0,
          [
            "NO_SERVER_DATA",
            "UNKNOWN_TYPE",
            "EMPTY_TYPE",
            "NAME_FIRST",
            "NAME_MIDDLE",
            "NAME_LAST",
            "NAME_MIDDLE_INITIAL",
            "NAME_FULL",
            "NAME_SUFFIX",
            "EMAIL_ADDRESS",
            "PHONE_HOME_NUMBER",
            "PHONE_HOME_CITY_CODE",
            "PHONE_HOME_COUNTRY_CODE",
            "PHONE_HOME_CITY_AND_NUMBER",
            "PHONE_HOME_WHOLE_NUMBER",
            "ADDRESS_HOME_LINE1",
            "ADDRESS_HOME_LINE2",
            "ADDRESS_HOME_APT_NUM",
            "ADDRESS_HOME_CITY",
            "ADDRESS_HOME_STATE",
            "ADDRESS_HOME_ZIP",
            "ADDRESS_HOME_COUNTRY",
            "CREDIT_CARD_NAME_FULL",
            "CREDIT_CARD_NUMBER",
            "CREDIT_CARD_EXP_MONTH",
            "CREDIT_CARD_EXP_2_DIGIT_YEAR",
            "CREDIT_CARD_EXP_4_DIGIT_YEAR",
            "CREDIT_CARD_EXP_DATE_2_DIGIT_YEAR",
            "CREDIT_CARD_EXP_DATE_4_DIGIT_YEAR",
            "CREDIT_CARD_TYPE",
            "CREDIT_CARD_VERIFICATION_CODE",
            "COMPANY_NAME",
            "FIELD_WITH_DEFAULT_VALUE",
            "MERCHANT_EMAIL_SIGNUP",
            "MERCHANT_PROMO_CODE",
            "PASSWORD",
            "ACCOUNT_CREATION_PASSWORD",
            "ADDRESS_HOME_STREET_ADDRESS",
            "ADDRESS_HOME_SORTING_CODE",
            "ADDRESS_HOME_DEPENDENT_LOCALITY",
            "ADDRESS_HOME_LINE3",
            "NOT_ACCOUNT_CREATION_PASSWORD",
            "USERNAME",
            "USERNAME_AND_EMAIL_ADDRESS",
            "NEW_PASSWORD",
            "PROBABLY_NEW_PASSWORD",
            "NOT_NEW_PASSWORD",
            "CREDIT_CARD_NAME_FIRST",
            "CREDIT_CARD_NAME_LAST",
            "PHONE_HOME_EXTENSION",
            "CONFIRMATION_PASSWORD",
            "AMBIGUOUS_TYPE",
            "SEARCH_TERM",
            "PRICE",
            "NOT_PASSWORD",
            "SINGLE_USERNAME",
            "NOT_USERNAME",
            "UPI_VPA",
            "ADDRESS_HOME_STREET_NAME",
            "ADDRESS_HOME_HOUSE_NUMBER",
            "ADDRESS_HOME_SUBPREMISE",
            "ADDRESS_HOME_OTHER_SUBUNIT",
            "NAME_LAST_FIRST",
            "NAME_LAST_CONJUNCTION",
            "NAME_LAST_SECOND",
            "NAME_HONORIFIC_PREFIX",
            "ADDRESS_HOME_ADDRESS",
            "ADDRESS_HOME_ADDRESS_WITH_NAME",
            "ADDRESS_HOME_FLOOR",
            "NAME_FULL_WITH_HONORIFIC_PREFIX",
            "BIRTHDATE_DAY",
            "BIRTHDATE_MONTH",
            "BIRTHDATE_4_DIGIT_YEAR",
            "PHONE_HOME_CITY_CODE_WITH_TRUNK_PREFIX",
            "PHONE_HOME_CITY_AND_NUMBER_WITHOUT_TRUNK_PREFIX",
            "PHONE_HOME_NUMBER_PREFIX",
            "PHONE_HOME_NUMBER_SUFFIX",
            "IBAN_VALUE",
            "CREDIT_CARD_STANDALONE_VERIFICATION_CODE",
            "NUMERIC_QUANTITY",
            "ONE_TIME_CODE",
            "DELIVERY_INSTRUCTIONS",
            "ADDRESS_HOME_OVERFLOW",
            "ADDRESS_HOME_LANDMARK",
            "ADDRESS_HOME_OVERFLOW_AND_LANDMARK",
            "ADDRESS_HOME_ADMIN_LEVEL2",
            "ADDRESS_HOME_STREET_LOCATION",
            "ADDRESS_HOME_BETWEEN_STREETS",
            "ADDRESS_HOME_BETWEEN_STREETS_OR_LANDMARK",
            "ADDRESS_HOME_BETWEEN_STREETS_1",
            "ADDRESS_HOME_BETWEEN_STREETS_2",
            "SINGLE_USERNAME_FORGOT_PASSWORD",
            "MAX_VALID_FIELD_TYPE",
          ].indexOf(x["field"] as string)
        );
        A.store.Ref(ptr + 4, x["fieldName"]);
        A.store.Bool(ptr + 16, "isLongField" in x ? true : false);
        A.store.Bool(ptr + 8, x["isLongField"] ? true : false);
        A.store.Bool(ptr + 17, "isRequired" in x ? true : false);
        A.store.Bool(ptr + 9, x["isRequired"] ? true : false);
        A.store.Ref(ptr + 12, x["placeholder"]);
      }
    },
    "load_AddressComponent": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["field"] = A.load.Enum(ptr + 0, [
        "NO_SERVER_DATA",
        "UNKNOWN_TYPE",
        "EMPTY_TYPE",
        "NAME_FIRST",
        "NAME_MIDDLE",
        "NAME_LAST",
        "NAME_MIDDLE_INITIAL",
        "NAME_FULL",
        "NAME_SUFFIX",
        "EMAIL_ADDRESS",
        "PHONE_HOME_NUMBER",
        "PHONE_HOME_CITY_CODE",
        "PHONE_HOME_COUNTRY_CODE",
        "PHONE_HOME_CITY_AND_NUMBER",
        "PHONE_HOME_WHOLE_NUMBER",
        "ADDRESS_HOME_LINE1",
        "ADDRESS_HOME_LINE2",
        "ADDRESS_HOME_APT_NUM",
        "ADDRESS_HOME_CITY",
        "ADDRESS_HOME_STATE",
        "ADDRESS_HOME_ZIP",
        "ADDRESS_HOME_COUNTRY",
        "CREDIT_CARD_NAME_FULL",
        "CREDIT_CARD_NUMBER",
        "CREDIT_CARD_EXP_MONTH",
        "CREDIT_CARD_EXP_2_DIGIT_YEAR",
        "CREDIT_CARD_EXP_4_DIGIT_YEAR",
        "CREDIT_CARD_EXP_DATE_2_DIGIT_YEAR",
        "CREDIT_CARD_EXP_DATE_4_DIGIT_YEAR",
        "CREDIT_CARD_TYPE",
        "CREDIT_CARD_VERIFICATION_CODE",
        "COMPANY_NAME",
        "FIELD_WITH_DEFAULT_VALUE",
        "MERCHANT_EMAIL_SIGNUP",
        "MERCHANT_PROMO_CODE",
        "PASSWORD",
        "ACCOUNT_CREATION_PASSWORD",
        "ADDRESS_HOME_STREET_ADDRESS",
        "ADDRESS_HOME_SORTING_CODE",
        "ADDRESS_HOME_DEPENDENT_LOCALITY",
        "ADDRESS_HOME_LINE3",
        "NOT_ACCOUNT_CREATION_PASSWORD",
        "USERNAME",
        "USERNAME_AND_EMAIL_ADDRESS",
        "NEW_PASSWORD",
        "PROBABLY_NEW_PASSWORD",
        "NOT_NEW_PASSWORD",
        "CREDIT_CARD_NAME_FIRST",
        "CREDIT_CARD_NAME_LAST",
        "PHONE_HOME_EXTENSION",
        "CONFIRMATION_PASSWORD",
        "AMBIGUOUS_TYPE",
        "SEARCH_TERM",
        "PRICE",
        "NOT_PASSWORD",
        "SINGLE_USERNAME",
        "NOT_USERNAME",
        "UPI_VPA",
        "ADDRESS_HOME_STREET_NAME",
        "ADDRESS_HOME_HOUSE_NUMBER",
        "ADDRESS_HOME_SUBPREMISE",
        "ADDRESS_HOME_OTHER_SUBUNIT",
        "NAME_LAST_FIRST",
        "NAME_LAST_CONJUNCTION",
        "NAME_LAST_SECOND",
        "NAME_HONORIFIC_PREFIX",
        "ADDRESS_HOME_ADDRESS",
        "ADDRESS_HOME_ADDRESS_WITH_NAME",
        "ADDRESS_HOME_FLOOR",
        "NAME_FULL_WITH_HONORIFIC_PREFIX",
        "BIRTHDATE_DAY",
        "BIRTHDATE_MONTH",
        "BIRTHDATE_4_DIGIT_YEAR",
        "PHONE_HOME_CITY_CODE_WITH_TRUNK_PREFIX",
        "PHONE_HOME_CITY_AND_NUMBER_WITHOUT_TRUNK_PREFIX",
        "PHONE_HOME_NUMBER_PREFIX",
        "PHONE_HOME_NUMBER_SUFFIX",
        "IBAN_VALUE",
        "CREDIT_CARD_STANDALONE_VERIFICATION_CODE",
        "NUMERIC_QUANTITY",
        "ONE_TIME_CODE",
        "DELIVERY_INSTRUCTIONS",
        "ADDRESS_HOME_OVERFLOW",
        "ADDRESS_HOME_LANDMARK",
        "ADDRESS_HOME_OVERFLOW_AND_LANDMARK",
        "ADDRESS_HOME_ADMIN_LEVEL2",
        "ADDRESS_HOME_STREET_LOCATION",
        "ADDRESS_HOME_BETWEEN_STREETS",
        "ADDRESS_HOME_BETWEEN_STREETS_OR_LANDMARK",
        "ADDRESS_HOME_BETWEEN_STREETS_1",
        "ADDRESS_HOME_BETWEEN_STREETS_2",
        "SINGLE_USERNAME_FORGOT_PASSWORD",
        "MAX_VALID_FIELD_TYPE",
      ]);
      x["fieldName"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["isLongField"] = A.load.Bool(ptr + 8);
      } else {
        delete x["isLongField"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["isRequired"] = A.load.Bool(ptr + 9);
      } else {
        delete x["isRequired"];
      }
      x["placeholder"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AddressComponentRow": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["row"]);
      }
    },
    "load_AddressComponentRow": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["row"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AddressComponents": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["components"]);
        A.store.Ref(ptr + 4, x["languageCode"]);
      }
    },
    "load_AddressComponents": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["components"] = A.load.Ref(ptr + 0, undefined);
      x["languageCode"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AddressField": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(
          ptr + 0,
          [
            "NO_SERVER_DATA",
            "UNKNOWN_TYPE",
            "EMPTY_TYPE",
            "NAME_FIRST",
            "NAME_MIDDLE",
            "NAME_LAST",
            "NAME_MIDDLE_INITIAL",
            "NAME_FULL",
            "NAME_SUFFIX",
            "EMAIL_ADDRESS",
            "PHONE_HOME_NUMBER",
            "PHONE_HOME_CITY_CODE",
            "PHONE_HOME_COUNTRY_CODE",
            "PHONE_HOME_CITY_AND_NUMBER",
            "PHONE_HOME_WHOLE_NUMBER",
            "ADDRESS_HOME_LINE1",
            "ADDRESS_HOME_LINE2",
            "ADDRESS_HOME_APT_NUM",
            "ADDRESS_HOME_CITY",
            "ADDRESS_HOME_STATE",
            "ADDRESS_HOME_ZIP",
            "ADDRESS_HOME_COUNTRY",
            "CREDIT_CARD_NAME_FULL",
            "CREDIT_CARD_NUMBER",
            "CREDIT_CARD_EXP_MONTH",
            "CREDIT_CARD_EXP_2_DIGIT_YEAR",
            "CREDIT_CARD_EXP_4_DIGIT_YEAR",
            "CREDIT_CARD_EXP_DATE_2_DIGIT_YEAR",
            "CREDIT_CARD_EXP_DATE_4_DIGIT_YEAR",
            "CREDIT_CARD_TYPE",
            "CREDIT_CARD_VERIFICATION_CODE",
            "COMPANY_NAME",
            "FIELD_WITH_DEFAULT_VALUE",
            "MERCHANT_EMAIL_SIGNUP",
            "MERCHANT_PROMO_CODE",
            "PASSWORD",
            "ACCOUNT_CREATION_PASSWORD",
            "ADDRESS_HOME_STREET_ADDRESS",
            "ADDRESS_HOME_SORTING_CODE",
            "ADDRESS_HOME_DEPENDENT_LOCALITY",
            "ADDRESS_HOME_LINE3",
            "NOT_ACCOUNT_CREATION_PASSWORD",
            "USERNAME",
            "USERNAME_AND_EMAIL_ADDRESS",
            "NEW_PASSWORD",
            "PROBABLY_NEW_PASSWORD",
            "NOT_NEW_PASSWORD",
            "CREDIT_CARD_NAME_FIRST",
            "CREDIT_CARD_NAME_LAST",
            "PHONE_HOME_EXTENSION",
            "CONFIRMATION_PASSWORD",
            "AMBIGUOUS_TYPE",
            "SEARCH_TERM",
            "PRICE",
            "NOT_PASSWORD",
            "SINGLE_USERNAME",
            "NOT_USERNAME",
            "UPI_VPA",
            "ADDRESS_HOME_STREET_NAME",
            "ADDRESS_HOME_HOUSE_NUMBER",
            "ADDRESS_HOME_SUBPREMISE",
            "ADDRESS_HOME_OTHER_SUBUNIT",
            "NAME_LAST_FIRST",
            "NAME_LAST_CONJUNCTION",
            "NAME_LAST_SECOND",
            "NAME_HONORIFIC_PREFIX",
            "ADDRESS_HOME_ADDRESS",
            "ADDRESS_HOME_ADDRESS_WITH_NAME",
            "ADDRESS_HOME_FLOOR",
            "NAME_FULL_WITH_HONORIFIC_PREFIX",
            "BIRTHDATE_DAY",
            "BIRTHDATE_MONTH",
            "BIRTHDATE_4_DIGIT_YEAR",
            "PHONE_HOME_CITY_CODE_WITH_TRUNK_PREFIX",
            "PHONE_HOME_CITY_AND_NUMBER_WITHOUT_TRUNK_PREFIX",
            "PHONE_HOME_NUMBER_PREFIX",
            "PHONE_HOME_NUMBER_SUFFIX",
            "IBAN_VALUE",
            "CREDIT_CARD_STANDALONE_VERIFICATION_CODE",
            "NUMERIC_QUANTITY",
            "ONE_TIME_CODE",
            "DELIVERY_INSTRUCTIONS",
            "ADDRESS_HOME_OVERFLOW",
            "ADDRESS_HOME_LANDMARK",
            "ADDRESS_HOME_OVERFLOW_AND_LANDMARK",
            "ADDRESS_HOME_ADMIN_LEVEL2",
            "ADDRESS_HOME_STREET_LOCATION",
            "ADDRESS_HOME_BETWEEN_STREETS",
            "ADDRESS_HOME_BETWEEN_STREETS_OR_LANDMARK",
            "ADDRESS_HOME_BETWEEN_STREETS_1",
            "ADDRESS_HOME_BETWEEN_STREETS_2",
            "SINGLE_USERNAME_FORGOT_PASSWORD",
            "MAX_VALID_FIELD_TYPE",
          ].indexOf(x["type"] as string)
        );
        A.store.Ref(ptr + 4, x["value"]);
      }
    },
    "load_AddressField": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["type"] = A.load.Enum(ptr + 0, [
        "NO_SERVER_DATA",
        "UNKNOWN_TYPE",
        "EMPTY_TYPE",
        "NAME_FIRST",
        "NAME_MIDDLE",
        "NAME_LAST",
        "NAME_MIDDLE_INITIAL",
        "NAME_FULL",
        "NAME_SUFFIX",
        "EMAIL_ADDRESS",
        "PHONE_HOME_NUMBER",
        "PHONE_HOME_CITY_CODE",
        "PHONE_HOME_COUNTRY_CODE",
        "PHONE_HOME_CITY_AND_NUMBER",
        "PHONE_HOME_WHOLE_NUMBER",
        "ADDRESS_HOME_LINE1",
        "ADDRESS_HOME_LINE2",
        "ADDRESS_HOME_APT_NUM",
        "ADDRESS_HOME_CITY",
        "ADDRESS_HOME_STATE",
        "ADDRESS_HOME_ZIP",
        "ADDRESS_HOME_COUNTRY",
        "CREDIT_CARD_NAME_FULL",
        "CREDIT_CARD_NUMBER",
        "CREDIT_CARD_EXP_MONTH",
        "CREDIT_CARD_EXP_2_DIGIT_YEAR",
        "CREDIT_CARD_EXP_4_DIGIT_YEAR",
        "CREDIT_CARD_EXP_DATE_2_DIGIT_YEAR",
        "CREDIT_CARD_EXP_DATE_4_DIGIT_YEAR",
        "CREDIT_CARD_TYPE",
        "CREDIT_CARD_VERIFICATION_CODE",
        "COMPANY_NAME",
        "FIELD_WITH_DEFAULT_VALUE",
        "MERCHANT_EMAIL_SIGNUP",
        "MERCHANT_PROMO_CODE",
        "PASSWORD",
        "ACCOUNT_CREATION_PASSWORD",
        "ADDRESS_HOME_STREET_ADDRESS",
        "ADDRESS_HOME_SORTING_CODE",
        "ADDRESS_HOME_DEPENDENT_LOCALITY",
        "ADDRESS_HOME_LINE3",
        "NOT_ACCOUNT_CREATION_PASSWORD",
        "USERNAME",
        "USERNAME_AND_EMAIL_ADDRESS",
        "NEW_PASSWORD",
        "PROBABLY_NEW_PASSWORD",
        "NOT_NEW_PASSWORD",
        "CREDIT_CARD_NAME_FIRST",
        "CREDIT_CARD_NAME_LAST",
        "PHONE_HOME_EXTENSION",
        "CONFIRMATION_PASSWORD",
        "AMBIGUOUS_TYPE",
        "SEARCH_TERM",
        "PRICE",
        "NOT_PASSWORD",
        "SINGLE_USERNAME",
        "NOT_USERNAME",
        "UPI_VPA",
        "ADDRESS_HOME_STREET_NAME",
        "ADDRESS_HOME_HOUSE_NUMBER",
        "ADDRESS_HOME_SUBPREMISE",
        "ADDRESS_HOME_OTHER_SUBUNIT",
        "NAME_LAST_FIRST",
        "NAME_LAST_CONJUNCTION",
        "NAME_LAST_SECOND",
        "NAME_HONORIFIC_PREFIX",
        "ADDRESS_HOME_ADDRESS",
        "ADDRESS_HOME_ADDRESS_WITH_NAME",
        "ADDRESS_HOME_FLOOR",
        "NAME_FULL_WITH_HONORIFIC_PREFIX",
        "BIRTHDATE_DAY",
        "BIRTHDATE_MONTH",
        "BIRTHDATE_4_DIGIT_YEAR",
        "PHONE_HOME_CITY_CODE_WITH_TRUNK_PREFIX",
        "PHONE_HOME_CITY_AND_NUMBER_WITHOUT_TRUNK_PREFIX",
        "PHONE_HOME_NUMBER_PREFIX",
        "PHONE_HOME_NUMBER_SUFFIX",
        "IBAN_VALUE",
        "CREDIT_CARD_STANDALONE_VERIFICATION_CODE",
        "NUMERIC_QUANTITY",
        "ONE_TIME_CODE",
        "DELIVERY_INSTRUCTIONS",
        "ADDRESS_HOME_OVERFLOW",
        "ADDRESS_HOME_LANDMARK",
        "ADDRESS_HOME_OVERFLOW_AND_LANDMARK",
        "ADDRESS_HOME_ADMIN_LEVEL2",
        "ADDRESS_HOME_STREET_LOCATION",
        "ADDRESS_HOME_BETWEEN_STREETS",
        "ADDRESS_HOME_BETWEEN_STREETS_OR_LANDMARK",
        "ADDRESS_HOME_BETWEEN_STREETS_1",
        "ADDRESS_HOME_BETWEEN_STREETS_2",
        "SINGLE_USERNAME_FORGOT_PASSWORD",
        "MAX_VALID_FIELD_TYPE",
      ]);
      x["value"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_AddressSource": (ref: heap.Ref<string>): number => {
      const idx = ["LOCAL_OR_SYNCABLE", "ACCOUNT"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_AutofillMetadata": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 22, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 19, false);
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 15, false);
        A.store.Bool(ptr + 21, false);
        A.store.Bool(ptr + 16, false);
      } else {
        A.store.Bool(ptr + 22, true);
        A.store.Ref(ptr + 0, x["summaryLabel"]);
        A.store.Ref(ptr + 4, x["summarySublabel"]);
        A.store.Enum(ptr + 8, ["LOCAL_OR_SYNCABLE", "ACCOUNT"].indexOf(x["source"] as string));
        A.store.Bool(ptr + 17, "isLocal" in x ? true : false);
        A.store.Bool(ptr + 12, x["isLocal"] ? true : false);
        A.store.Bool(ptr + 18, "isCached" in x ? true : false);
        A.store.Bool(ptr + 13, x["isCached"] ? true : false);
        A.store.Bool(ptr + 19, "isMigratable" in x ? true : false);
        A.store.Bool(ptr + 14, x["isMigratable"] ? true : false);
        A.store.Bool(ptr + 20, "isVirtualCardEnrollmentEligible" in x ? true : false);
        A.store.Bool(ptr + 15, x["isVirtualCardEnrollmentEligible"] ? true : false);
        A.store.Bool(ptr + 21, "isVirtualCardEnrolled" in x ? true : false);
        A.store.Bool(ptr + 16, x["isVirtualCardEnrolled"] ? true : false);
      }
    },
    "load_AutofillMetadata": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["summaryLabel"] = A.load.Ref(ptr + 0, undefined);
      x["summarySublabel"] = A.load.Ref(ptr + 4, undefined);
      x["source"] = A.load.Enum(ptr + 8, ["LOCAL_OR_SYNCABLE", "ACCOUNT"]);
      if (A.load.Bool(ptr + 17)) {
        x["isLocal"] = A.load.Bool(ptr + 12);
      } else {
        delete x["isLocal"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["isCached"] = A.load.Bool(ptr + 13);
      } else {
        delete x["isCached"];
      }
      if (A.load.Bool(ptr + 19)) {
        x["isMigratable"] = A.load.Bool(ptr + 14);
      } else {
        delete x["isMigratable"];
      }
      if (A.load.Bool(ptr + 20)) {
        x["isVirtualCardEnrollmentEligible"] = A.load.Bool(ptr + 15);
      } else {
        delete x["isVirtualCardEnrollmentEligible"];
      }
      if (A.load.Bool(ptr + 21)) {
        x["isVirtualCardEnrolled"] = A.load.Bool(ptr + 16);
      } else {
        delete x["isVirtualCardEnrolled"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AddressEntry": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 35, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);

        A.store.Bool(ptr + 12 + 22, false);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 12 + 4, undefined);
        A.store.Enum(ptr + 12 + 8, -1);
        A.store.Bool(ptr + 12 + 17, false);
        A.store.Bool(ptr + 12 + 12, false);
        A.store.Bool(ptr + 12 + 18, false);
        A.store.Bool(ptr + 12 + 13, false);
        A.store.Bool(ptr + 12 + 19, false);
        A.store.Bool(ptr + 12 + 14, false);
        A.store.Bool(ptr + 12 + 20, false);
        A.store.Bool(ptr + 12 + 15, false);
        A.store.Bool(ptr + 12 + 21, false);
        A.store.Bool(ptr + 12 + 16, false);
      } else {
        A.store.Bool(ptr + 35, true);
        A.store.Ref(ptr + 0, x["guid"]);
        A.store.Ref(ptr + 4, x["fields"]);
        A.store.Ref(ptr + 8, x["languageCode"]);

        if (typeof x["metadata"] === "undefined") {
          A.store.Bool(ptr + 12 + 22, false);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Ref(ptr + 12 + 4, undefined);
          A.store.Enum(ptr + 12 + 8, -1);
          A.store.Bool(ptr + 12 + 17, false);
          A.store.Bool(ptr + 12 + 12, false);
          A.store.Bool(ptr + 12 + 18, false);
          A.store.Bool(ptr + 12 + 13, false);
          A.store.Bool(ptr + 12 + 19, false);
          A.store.Bool(ptr + 12 + 14, false);
          A.store.Bool(ptr + 12 + 20, false);
          A.store.Bool(ptr + 12 + 15, false);
          A.store.Bool(ptr + 12 + 21, false);
          A.store.Bool(ptr + 12 + 16, false);
        } else {
          A.store.Bool(ptr + 12 + 22, true);
          A.store.Ref(ptr + 12 + 0, x["metadata"]["summaryLabel"]);
          A.store.Ref(ptr + 12 + 4, x["metadata"]["summarySublabel"]);
          A.store.Enum(ptr + 12 + 8, ["LOCAL_OR_SYNCABLE", "ACCOUNT"].indexOf(x["metadata"]["source"] as string));
          A.store.Bool(ptr + 12 + 17, "isLocal" in x["metadata"] ? true : false);
          A.store.Bool(ptr + 12 + 12, x["metadata"]["isLocal"] ? true : false);
          A.store.Bool(ptr + 12 + 18, "isCached" in x["metadata"] ? true : false);
          A.store.Bool(ptr + 12 + 13, x["metadata"]["isCached"] ? true : false);
          A.store.Bool(ptr + 12 + 19, "isMigratable" in x["metadata"] ? true : false);
          A.store.Bool(ptr + 12 + 14, x["metadata"]["isMigratable"] ? true : false);
          A.store.Bool(ptr + 12 + 20, "isVirtualCardEnrollmentEligible" in x["metadata"] ? true : false);
          A.store.Bool(ptr + 12 + 15, x["metadata"]["isVirtualCardEnrollmentEligible"] ? true : false);
          A.store.Bool(ptr + 12 + 21, "isVirtualCardEnrolled" in x["metadata"] ? true : false);
          A.store.Bool(ptr + 12 + 16, x["metadata"]["isVirtualCardEnrolled"] ? true : false);
        }
      }
    },
    "load_AddressEntry": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["guid"] = A.load.Ref(ptr + 0, undefined);
      x["fields"] = A.load.Ref(ptr + 4, undefined);
      x["languageCode"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 12 + 22)) {
        x["metadata"] = {};
        x["metadata"]["summaryLabel"] = A.load.Ref(ptr + 12 + 0, undefined);
        x["metadata"]["summarySublabel"] = A.load.Ref(ptr + 12 + 4, undefined);
        x["metadata"]["source"] = A.load.Enum(ptr + 12 + 8, ["LOCAL_OR_SYNCABLE", "ACCOUNT"]);
        if (A.load.Bool(ptr + 12 + 17)) {
          x["metadata"]["isLocal"] = A.load.Bool(ptr + 12 + 12);
        } else {
          delete x["metadata"]["isLocal"];
        }
        if (A.load.Bool(ptr + 12 + 18)) {
          x["metadata"]["isCached"] = A.load.Bool(ptr + 12 + 13);
        } else {
          delete x["metadata"]["isCached"];
        }
        if (A.load.Bool(ptr + 12 + 19)) {
          x["metadata"]["isMigratable"] = A.load.Bool(ptr + 12 + 14);
        } else {
          delete x["metadata"]["isMigratable"];
        }
        if (A.load.Bool(ptr + 12 + 20)) {
          x["metadata"]["isVirtualCardEnrollmentEligible"] = A.load.Bool(ptr + 12 + 15);
        } else {
          delete x["metadata"]["isVirtualCardEnrollmentEligible"];
        }
        if (A.load.Bool(ptr + 12 + 21)) {
          x["metadata"]["isVirtualCardEnrolled"] = A.load.Bool(ptr + 12 + 16);
        } else {
          delete x["metadata"]["isVirtualCardEnrolled"];
        }
      } else {
        delete x["metadata"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CountryEntry": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["name"]);
        A.store.Ref(ptr + 4, x["countryCode"]);
      }
    },
    "load_CountryEntry": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      x["countryCode"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CreditCardEntry": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 55, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);

        A.store.Bool(ptr + 32 + 22, false);
        A.store.Ref(ptr + 32 + 0, undefined);
        A.store.Ref(ptr + 32 + 4, undefined);
        A.store.Enum(ptr + 32 + 8, -1);
        A.store.Bool(ptr + 32 + 17, false);
        A.store.Bool(ptr + 32 + 12, false);
        A.store.Bool(ptr + 32 + 18, false);
        A.store.Bool(ptr + 32 + 13, false);
        A.store.Bool(ptr + 32 + 19, false);
        A.store.Bool(ptr + 32 + 14, false);
        A.store.Bool(ptr + 32 + 20, false);
        A.store.Bool(ptr + 32 + 15, false);
        A.store.Bool(ptr + 32 + 21, false);
        A.store.Bool(ptr + 32 + 16, false);
      } else {
        A.store.Bool(ptr + 55, true);
        A.store.Ref(ptr + 0, x["guid"]);
        A.store.Ref(ptr + 4, x["name"]);
        A.store.Ref(ptr + 8, x["cardNumber"]);
        A.store.Ref(ptr + 12, x["expirationMonth"]);
        A.store.Ref(ptr + 16, x["expirationYear"]);
        A.store.Ref(ptr + 20, x["nickname"]);
        A.store.Ref(ptr + 24, x["network"]);
        A.store.Ref(ptr + 28, x["imageSrc"]);

        if (typeof x["metadata"] === "undefined") {
          A.store.Bool(ptr + 32 + 22, false);
          A.store.Ref(ptr + 32 + 0, undefined);
          A.store.Ref(ptr + 32 + 4, undefined);
          A.store.Enum(ptr + 32 + 8, -1);
          A.store.Bool(ptr + 32 + 17, false);
          A.store.Bool(ptr + 32 + 12, false);
          A.store.Bool(ptr + 32 + 18, false);
          A.store.Bool(ptr + 32 + 13, false);
          A.store.Bool(ptr + 32 + 19, false);
          A.store.Bool(ptr + 32 + 14, false);
          A.store.Bool(ptr + 32 + 20, false);
          A.store.Bool(ptr + 32 + 15, false);
          A.store.Bool(ptr + 32 + 21, false);
          A.store.Bool(ptr + 32 + 16, false);
        } else {
          A.store.Bool(ptr + 32 + 22, true);
          A.store.Ref(ptr + 32 + 0, x["metadata"]["summaryLabel"]);
          A.store.Ref(ptr + 32 + 4, x["metadata"]["summarySublabel"]);
          A.store.Enum(ptr + 32 + 8, ["LOCAL_OR_SYNCABLE", "ACCOUNT"].indexOf(x["metadata"]["source"] as string));
          A.store.Bool(ptr + 32 + 17, "isLocal" in x["metadata"] ? true : false);
          A.store.Bool(ptr + 32 + 12, x["metadata"]["isLocal"] ? true : false);
          A.store.Bool(ptr + 32 + 18, "isCached" in x["metadata"] ? true : false);
          A.store.Bool(ptr + 32 + 13, x["metadata"]["isCached"] ? true : false);
          A.store.Bool(ptr + 32 + 19, "isMigratable" in x["metadata"] ? true : false);
          A.store.Bool(ptr + 32 + 14, x["metadata"]["isMigratable"] ? true : false);
          A.store.Bool(ptr + 32 + 20, "isVirtualCardEnrollmentEligible" in x["metadata"] ? true : false);
          A.store.Bool(ptr + 32 + 15, x["metadata"]["isVirtualCardEnrollmentEligible"] ? true : false);
          A.store.Bool(ptr + 32 + 21, "isVirtualCardEnrolled" in x["metadata"] ? true : false);
          A.store.Bool(ptr + 32 + 16, x["metadata"]["isVirtualCardEnrolled"] ? true : false);
        }
      }
    },
    "load_CreditCardEntry": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["guid"] = A.load.Ref(ptr + 0, undefined);
      x["name"] = A.load.Ref(ptr + 4, undefined);
      x["cardNumber"] = A.load.Ref(ptr + 8, undefined);
      x["expirationMonth"] = A.load.Ref(ptr + 12, undefined);
      x["expirationYear"] = A.load.Ref(ptr + 16, undefined);
      x["nickname"] = A.load.Ref(ptr + 20, undefined);
      x["network"] = A.load.Ref(ptr + 24, undefined);
      x["imageSrc"] = A.load.Ref(ptr + 28, undefined);
      if (A.load.Bool(ptr + 32 + 22)) {
        x["metadata"] = {};
        x["metadata"]["summaryLabel"] = A.load.Ref(ptr + 32 + 0, undefined);
        x["metadata"]["summarySublabel"] = A.load.Ref(ptr + 32 + 4, undefined);
        x["metadata"]["source"] = A.load.Enum(ptr + 32 + 8, ["LOCAL_OR_SYNCABLE", "ACCOUNT"]);
        if (A.load.Bool(ptr + 32 + 17)) {
          x["metadata"]["isLocal"] = A.load.Bool(ptr + 32 + 12);
        } else {
          delete x["metadata"]["isLocal"];
        }
        if (A.load.Bool(ptr + 32 + 18)) {
          x["metadata"]["isCached"] = A.load.Bool(ptr + 32 + 13);
        } else {
          delete x["metadata"]["isCached"];
        }
        if (A.load.Bool(ptr + 32 + 19)) {
          x["metadata"]["isMigratable"] = A.load.Bool(ptr + 32 + 14);
        } else {
          delete x["metadata"]["isMigratable"];
        }
        if (A.load.Bool(ptr + 32 + 20)) {
          x["metadata"]["isVirtualCardEnrollmentEligible"] = A.load.Bool(ptr + 32 + 15);
        } else {
          delete x["metadata"]["isVirtualCardEnrollmentEligible"];
        }
        if (A.load.Bool(ptr + 32 + 21)) {
          x["metadata"]["isVirtualCardEnrolled"] = A.load.Bool(ptr + 32 + 16);
        } else {
          delete x["metadata"]["isVirtualCardEnrolled"];
        }
      } else {
        delete x["metadata"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_IbanEntry": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 35, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);

        A.store.Bool(ptr + 12 + 22, false);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 12 + 4, undefined);
        A.store.Enum(ptr + 12 + 8, -1);
        A.store.Bool(ptr + 12 + 17, false);
        A.store.Bool(ptr + 12 + 12, false);
        A.store.Bool(ptr + 12 + 18, false);
        A.store.Bool(ptr + 12 + 13, false);
        A.store.Bool(ptr + 12 + 19, false);
        A.store.Bool(ptr + 12 + 14, false);
        A.store.Bool(ptr + 12 + 20, false);
        A.store.Bool(ptr + 12 + 15, false);
        A.store.Bool(ptr + 12 + 21, false);
        A.store.Bool(ptr + 12 + 16, false);
      } else {
        A.store.Bool(ptr + 35, true);
        A.store.Ref(ptr + 0, x["guid"]);
        A.store.Ref(ptr + 4, x["value"]);
        A.store.Ref(ptr + 8, x["nickname"]);

        if (typeof x["metadata"] === "undefined") {
          A.store.Bool(ptr + 12 + 22, false);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Ref(ptr + 12 + 4, undefined);
          A.store.Enum(ptr + 12 + 8, -1);
          A.store.Bool(ptr + 12 + 17, false);
          A.store.Bool(ptr + 12 + 12, false);
          A.store.Bool(ptr + 12 + 18, false);
          A.store.Bool(ptr + 12 + 13, false);
          A.store.Bool(ptr + 12 + 19, false);
          A.store.Bool(ptr + 12 + 14, false);
          A.store.Bool(ptr + 12 + 20, false);
          A.store.Bool(ptr + 12 + 15, false);
          A.store.Bool(ptr + 12 + 21, false);
          A.store.Bool(ptr + 12 + 16, false);
        } else {
          A.store.Bool(ptr + 12 + 22, true);
          A.store.Ref(ptr + 12 + 0, x["metadata"]["summaryLabel"]);
          A.store.Ref(ptr + 12 + 4, x["metadata"]["summarySublabel"]);
          A.store.Enum(ptr + 12 + 8, ["LOCAL_OR_SYNCABLE", "ACCOUNT"].indexOf(x["metadata"]["source"] as string));
          A.store.Bool(ptr + 12 + 17, "isLocal" in x["metadata"] ? true : false);
          A.store.Bool(ptr + 12 + 12, x["metadata"]["isLocal"] ? true : false);
          A.store.Bool(ptr + 12 + 18, "isCached" in x["metadata"] ? true : false);
          A.store.Bool(ptr + 12 + 13, x["metadata"]["isCached"] ? true : false);
          A.store.Bool(ptr + 12 + 19, "isMigratable" in x["metadata"] ? true : false);
          A.store.Bool(ptr + 12 + 14, x["metadata"]["isMigratable"] ? true : false);
          A.store.Bool(ptr + 12 + 20, "isVirtualCardEnrollmentEligible" in x["metadata"] ? true : false);
          A.store.Bool(ptr + 12 + 15, x["metadata"]["isVirtualCardEnrollmentEligible"] ? true : false);
          A.store.Bool(ptr + 12 + 21, "isVirtualCardEnrolled" in x["metadata"] ? true : false);
          A.store.Bool(ptr + 12 + 16, x["metadata"]["isVirtualCardEnrolled"] ? true : false);
        }
      }
    },
    "load_IbanEntry": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["guid"] = A.load.Ref(ptr + 0, undefined);
      x["value"] = A.load.Ref(ptr + 4, undefined);
      x["nickname"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 12 + 22)) {
        x["metadata"] = {};
        x["metadata"]["summaryLabel"] = A.load.Ref(ptr + 12 + 0, undefined);
        x["metadata"]["summarySublabel"] = A.load.Ref(ptr + 12 + 4, undefined);
        x["metadata"]["source"] = A.load.Enum(ptr + 12 + 8, ["LOCAL_OR_SYNCABLE", "ACCOUNT"]);
        if (A.load.Bool(ptr + 12 + 17)) {
          x["metadata"]["isLocal"] = A.load.Bool(ptr + 12 + 12);
        } else {
          delete x["metadata"]["isLocal"];
        }
        if (A.load.Bool(ptr + 12 + 18)) {
          x["metadata"]["isCached"] = A.load.Bool(ptr + 12 + 13);
        } else {
          delete x["metadata"]["isCached"];
        }
        if (A.load.Bool(ptr + 12 + 19)) {
          x["metadata"]["isMigratable"] = A.load.Bool(ptr + 12 + 14);
        } else {
          delete x["metadata"]["isMigratable"];
        }
        if (A.load.Bool(ptr + 12 + 20)) {
          x["metadata"]["isVirtualCardEnrollmentEligible"] = A.load.Bool(ptr + 12 + 15);
        } else {
          delete x["metadata"]["isVirtualCardEnrollmentEligible"];
        }
        if (A.load.Bool(ptr + 12 + 21)) {
          x["metadata"]["isVirtualCardEnrolled"] = A.load.Bool(ptr + 12 + 16);
        } else {
          delete x["metadata"]["isVirtualCardEnrolled"];
        }
      } else {
        delete x["metadata"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_AddVirtualCard": (): heap.Ref<boolean> => {
      if (WEBEXT?.autofillPrivate && "addVirtualCard" in WEBEXT?.autofillPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AddVirtualCard": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autofillPrivate.addVirtualCard);
    },
    "call_AddVirtualCard": (retPtr: Pointer, cardId: heap.Ref<object>): void => {
      const _ret = WEBEXT.autofillPrivate.addVirtualCard(A.H.get<object>(cardId));
    },
    "try_AddVirtualCard": (retPtr: Pointer, errPtr: Pointer, cardId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autofillPrivate.addVirtualCard(A.H.get<object>(cardId));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AuthenticateUserAndFlipMandatoryAuthToggle": (): heap.Ref<boolean> => {
      if (WEBEXT?.autofillPrivate && "authenticateUserAndFlipMandatoryAuthToggle" in WEBEXT?.autofillPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AuthenticateUserAndFlipMandatoryAuthToggle": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autofillPrivate.authenticateUserAndFlipMandatoryAuthToggle);
    },
    "call_AuthenticateUserAndFlipMandatoryAuthToggle": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autofillPrivate.authenticateUserAndFlipMandatoryAuthToggle();
    },
    "try_AuthenticateUserAndFlipMandatoryAuthToggle": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autofillPrivate.authenticateUserAndFlipMandatoryAuthToggle();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AuthenticateUserToEditLocalCard": (): heap.Ref<boolean> => {
      if (WEBEXT?.autofillPrivate && "authenticateUserToEditLocalCard" in WEBEXT?.autofillPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AuthenticateUserToEditLocalCard": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autofillPrivate.authenticateUserToEditLocalCard);
    },
    "call_AuthenticateUserToEditLocalCard": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autofillPrivate.authenticateUserToEditLocalCard();
      A.store.Ref(retPtr, _ret);
    },
    "try_AuthenticateUserToEditLocalCard": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autofillPrivate.authenticateUserToEditLocalCard();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CheckIfDeviceAuthAvailable": (): heap.Ref<boolean> => {
      if (WEBEXT?.autofillPrivate && "checkIfDeviceAuthAvailable" in WEBEXT?.autofillPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CheckIfDeviceAuthAvailable": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autofillPrivate.checkIfDeviceAuthAvailable);
    },
    "call_CheckIfDeviceAuthAvailable": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autofillPrivate.checkIfDeviceAuthAvailable();
      A.store.Ref(retPtr, _ret);
    },
    "try_CheckIfDeviceAuthAvailable": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autofillPrivate.checkIfDeviceAuthAvailable();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAccountInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.autofillPrivate && "getAccountInfo" in WEBEXT?.autofillPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAccountInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autofillPrivate.getAccountInfo);
    },
    "call_GetAccountInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autofillPrivate.getAccountInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAccountInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autofillPrivate.getAccountInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAddressComponents": (): heap.Ref<boolean> => {
      if (WEBEXT?.autofillPrivate && "getAddressComponents" in WEBEXT?.autofillPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAddressComponents": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autofillPrivate.getAddressComponents);
    },
    "call_GetAddressComponents": (retPtr: Pointer, countryCode: heap.Ref<object>): void => {
      const _ret = WEBEXT.autofillPrivate.getAddressComponents(A.H.get<object>(countryCode));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAddressComponents": (
      retPtr: Pointer,
      errPtr: Pointer,
      countryCode: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autofillPrivate.getAddressComponents(A.H.get<object>(countryCode));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAddressList": (): heap.Ref<boolean> => {
      if (WEBEXT?.autofillPrivate && "getAddressList" in WEBEXT?.autofillPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAddressList": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autofillPrivate.getAddressList);
    },
    "call_GetAddressList": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autofillPrivate.getAddressList();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAddressList": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autofillPrivate.getAddressList();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetCountryList": (): heap.Ref<boolean> => {
      if (WEBEXT?.autofillPrivate && "getCountryList" in WEBEXT?.autofillPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetCountryList": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autofillPrivate.getCountryList);
    },
    "call_GetCountryList": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autofillPrivate.getCountryList();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetCountryList": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autofillPrivate.getCountryList();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetCreditCardList": (): heap.Ref<boolean> => {
      if (WEBEXT?.autofillPrivate && "getCreditCardList" in WEBEXT?.autofillPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetCreditCardList": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autofillPrivate.getCreditCardList);
    },
    "call_GetCreditCardList": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autofillPrivate.getCreditCardList();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetCreditCardList": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autofillPrivate.getCreditCardList();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetIbanList": (): heap.Ref<boolean> => {
      if (WEBEXT?.autofillPrivate && "getIbanList" in WEBEXT?.autofillPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetIbanList": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autofillPrivate.getIbanList);
    },
    "call_GetIbanList": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autofillPrivate.getIbanList();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetIbanList": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autofillPrivate.getIbanList();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsValidIban": (): heap.Ref<boolean> => {
      if (WEBEXT?.autofillPrivate && "isValidIban" in WEBEXT?.autofillPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsValidIban": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autofillPrivate.isValidIban);
    },
    "call_IsValidIban": (retPtr: Pointer, ibanValue: heap.Ref<object>): void => {
      const _ret = WEBEXT.autofillPrivate.isValidIban(A.H.get<object>(ibanValue));
      A.store.Ref(retPtr, _ret);
    },
    "try_IsValidIban": (retPtr: Pointer, errPtr: Pointer, ibanValue: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autofillPrivate.isValidIban(A.H.get<object>(ibanValue));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LogServerCardLinkClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.autofillPrivate && "logServerCardLinkClicked" in WEBEXT?.autofillPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LogServerCardLinkClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autofillPrivate.logServerCardLinkClicked);
    },
    "call_LogServerCardLinkClicked": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autofillPrivate.logServerCardLinkClicked();
    },
    "try_LogServerCardLinkClicked": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autofillPrivate.logServerCardLinkClicked();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_MaskCreditCard": (): heap.Ref<boolean> => {
      if (WEBEXT?.autofillPrivate && "maskCreditCard" in WEBEXT?.autofillPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_MaskCreditCard": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autofillPrivate.maskCreditCard);
    },
    "call_MaskCreditCard": (retPtr: Pointer, guid: heap.Ref<object>): void => {
      const _ret = WEBEXT.autofillPrivate.maskCreditCard(A.H.get<object>(guid));
    },
    "try_MaskCreditCard": (retPtr: Pointer, errPtr: Pointer, guid: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autofillPrivate.maskCreditCard(A.H.get<object>(guid));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_MigrateCreditCards": (): heap.Ref<boolean> => {
      if (WEBEXT?.autofillPrivate && "migrateCreditCards" in WEBEXT?.autofillPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_MigrateCreditCards": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autofillPrivate.migrateCreditCards);
    },
    "call_MigrateCreditCards": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autofillPrivate.migrateCreditCards();
    },
    "try_MigrateCreditCards": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autofillPrivate.migrateCreditCards();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPersonalDataChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.autofillPrivate?.onPersonalDataChanged &&
        "addListener" in WEBEXT?.autofillPrivate?.onPersonalDataChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPersonalDataChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autofillPrivate.onPersonalDataChanged.addListener);
    },
    "call_OnPersonalDataChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.autofillPrivate.onPersonalDataChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnPersonalDataChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autofillPrivate.onPersonalDataChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPersonalDataChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.autofillPrivate?.onPersonalDataChanged &&
        "removeListener" in WEBEXT?.autofillPrivate?.onPersonalDataChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPersonalDataChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autofillPrivate.onPersonalDataChanged.removeListener);
    },
    "call_OffPersonalDataChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.autofillPrivate.onPersonalDataChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffPersonalDataChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autofillPrivate.onPersonalDataChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPersonalDataChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.autofillPrivate?.onPersonalDataChanged &&
        "hasListener" in WEBEXT?.autofillPrivate?.onPersonalDataChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPersonalDataChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autofillPrivate.onPersonalDataChanged.hasListener);
    },
    "call_HasOnPersonalDataChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.autofillPrivate.onPersonalDataChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPersonalDataChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autofillPrivate.onPersonalDataChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveEntry": (): heap.Ref<boolean> => {
      if (WEBEXT?.autofillPrivate && "removeEntry" in WEBEXT?.autofillPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveEntry": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autofillPrivate.removeEntry);
    },
    "call_RemoveEntry": (retPtr: Pointer, guid: heap.Ref<object>): void => {
      const _ret = WEBEXT.autofillPrivate.removeEntry(A.H.get<object>(guid));
    },
    "try_RemoveEntry": (retPtr: Pointer, errPtr: Pointer, guid: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autofillPrivate.removeEntry(A.H.get<object>(guid));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveVirtualCard": (): heap.Ref<boolean> => {
      if (WEBEXT?.autofillPrivate && "removeVirtualCard" in WEBEXT?.autofillPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveVirtualCard": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autofillPrivate.removeVirtualCard);
    },
    "call_RemoveVirtualCard": (retPtr: Pointer, cardId: heap.Ref<object>): void => {
      const _ret = WEBEXT.autofillPrivate.removeVirtualCard(A.H.get<object>(cardId));
    },
    "try_RemoveVirtualCard": (retPtr: Pointer, errPtr: Pointer, cardId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autofillPrivate.removeVirtualCard(A.H.get<object>(cardId));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SaveAddress": (): heap.Ref<boolean> => {
      if (WEBEXT?.autofillPrivate && "saveAddress" in WEBEXT?.autofillPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SaveAddress": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autofillPrivate.saveAddress);
    },
    "call_SaveAddress": (retPtr: Pointer, address: Pointer): void => {
      const address_ffi = {};

      address_ffi["guid"] = A.load.Ref(address + 0, undefined);
      address_ffi["fields"] = A.load.Ref(address + 4, undefined);
      address_ffi["languageCode"] = A.load.Ref(address + 8, undefined);
      if (A.load.Bool(address + 12 + 22)) {
        address_ffi["metadata"] = {};
        address_ffi["metadata"]["summaryLabel"] = A.load.Ref(address + 12 + 0, undefined);
        address_ffi["metadata"]["summarySublabel"] = A.load.Ref(address + 12 + 4, undefined);
        address_ffi["metadata"]["source"] = A.load.Enum(address + 12 + 8, ["LOCAL_OR_SYNCABLE", "ACCOUNT"]);
        if (A.load.Bool(address + 12 + 17)) {
          address_ffi["metadata"]["isLocal"] = A.load.Bool(address + 12 + 12);
        }
        if (A.load.Bool(address + 12 + 18)) {
          address_ffi["metadata"]["isCached"] = A.load.Bool(address + 12 + 13);
        }
        if (A.load.Bool(address + 12 + 19)) {
          address_ffi["metadata"]["isMigratable"] = A.load.Bool(address + 12 + 14);
        }
        if (A.load.Bool(address + 12 + 20)) {
          address_ffi["metadata"]["isVirtualCardEnrollmentEligible"] = A.load.Bool(address + 12 + 15);
        }
        if (A.load.Bool(address + 12 + 21)) {
          address_ffi["metadata"]["isVirtualCardEnrolled"] = A.load.Bool(address + 12 + 16);
        }
      }

      const _ret = WEBEXT.autofillPrivate.saveAddress(address_ffi);
    },
    "try_SaveAddress": (retPtr: Pointer, errPtr: Pointer, address: Pointer): heap.Ref<boolean> => {
      try {
        const address_ffi = {};

        address_ffi["guid"] = A.load.Ref(address + 0, undefined);
        address_ffi["fields"] = A.load.Ref(address + 4, undefined);
        address_ffi["languageCode"] = A.load.Ref(address + 8, undefined);
        if (A.load.Bool(address + 12 + 22)) {
          address_ffi["metadata"] = {};
          address_ffi["metadata"]["summaryLabel"] = A.load.Ref(address + 12 + 0, undefined);
          address_ffi["metadata"]["summarySublabel"] = A.load.Ref(address + 12 + 4, undefined);
          address_ffi["metadata"]["source"] = A.load.Enum(address + 12 + 8, ["LOCAL_OR_SYNCABLE", "ACCOUNT"]);
          if (A.load.Bool(address + 12 + 17)) {
            address_ffi["metadata"]["isLocal"] = A.load.Bool(address + 12 + 12);
          }
          if (A.load.Bool(address + 12 + 18)) {
            address_ffi["metadata"]["isCached"] = A.load.Bool(address + 12 + 13);
          }
          if (A.load.Bool(address + 12 + 19)) {
            address_ffi["metadata"]["isMigratable"] = A.load.Bool(address + 12 + 14);
          }
          if (A.load.Bool(address + 12 + 20)) {
            address_ffi["metadata"]["isVirtualCardEnrollmentEligible"] = A.load.Bool(address + 12 + 15);
          }
          if (A.load.Bool(address + 12 + 21)) {
            address_ffi["metadata"]["isVirtualCardEnrolled"] = A.load.Bool(address + 12 + 16);
          }
        }

        const _ret = WEBEXT.autofillPrivate.saveAddress(address_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SaveCreditCard": (): heap.Ref<boolean> => {
      if (WEBEXT?.autofillPrivate && "saveCreditCard" in WEBEXT?.autofillPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SaveCreditCard": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autofillPrivate.saveCreditCard);
    },
    "call_SaveCreditCard": (retPtr: Pointer, card: Pointer): void => {
      const card_ffi = {};

      card_ffi["guid"] = A.load.Ref(card + 0, undefined);
      card_ffi["name"] = A.load.Ref(card + 4, undefined);
      card_ffi["cardNumber"] = A.load.Ref(card + 8, undefined);
      card_ffi["expirationMonth"] = A.load.Ref(card + 12, undefined);
      card_ffi["expirationYear"] = A.load.Ref(card + 16, undefined);
      card_ffi["nickname"] = A.load.Ref(card + 20, undefined);
      card_ffi["network"] = A.load.Ref(card + 24, undefined);
      card_ffi["imageSrc"] = A.load.Ref(card + 28, undefined);
      if (A.load.Bool(card + 32 + 22)) {
        card_ffi["metadata"] = {};
        card_ffi["metadata"]["summaryLabel"] = A.load.Ref(card + 32 + 0, undefined);
        card_ffi["metadata"]["summarySublabel"] = A.load.Ref(card + 32 + 4, undefined);
        card_ffi["metadata"]["source"] = A.load.Enum(card + 32 + 8, ["LOCAL_OR_SYNCABLE", "ACCOUNT"]);
        if (A.load.Bool(card + 32 + 17)) {
          card_ffi["metadata"]["isLocal"] = A.load.Bool(card + 32 + 12);
        }
        if (A.load.Bool(card + 32 + 18)) {
          card_ffi["metadata"]["isCached"] = A.load.Bool(card + 32 + 13);
        }
        if (A.load.Bool(card + 32 + 19)) {
          card_ffi["metadata"]["isMigratable"] = A.load.Bool(card + 32 + 14);
        }
        if (A.load.Bool(card + 32 + 20)) {
          card_ffi["metadata"]["isVirtualCardEnrollmentEligible"] = A.load.Bool(card + 32 + 15);
        }
        if (A.load.Bool(card + 32 + 21)) {
          card_ffi["metadata"]["isVirtualCardEnrolled"] = A.load.Bool(card + 32 + 16);
        }
      }

      const _ret = WEBEXT.autofillPrivate.saveCreditCard(card_ffi);
    },
    "try_SaveCreditCard": (retPtr: Pointer, errPtr: Pointer, card: Pointer): heap.Ref<boolean> => {
      try {
        const card_ffi = {};

        card_ffi["guid"] = A.load.Ref(card + 0, undefined);
        card_ffi["name"] = A.load.Ref(card + 4, undefined);
        card_ffi["cardNumber"] = A.load.Ref(card + 8, undefined);
        card_ffi["expirationMonth"] = A.load.Ref(card + 12, undefined);
        card_ffi["expirationYear"] = A.load.Ref(card + 16, undefined);
        card_ffi["nickname"] = A.load.Ref(card + 20, undefined);
        card_ffi["network"] = A.load.Ref(card + 24, undefined);
        card_ffi["imageSrc"] = A.load.Ref(card + 28, undefined);
        if (A.load.Bool(card + 32 + 22)) {
          card_ffi["metadata"] = {};
          card_ffi["metadata"]["summaryLabel"] = A.load.Ref(card + 32 + 0, undefined);
          card_ffi["metadata"]["summarySublabel"] = A.load.Ref(card + 32 + 4, undefined);
          card_ffi["metadata"]["source"] = A.load.Enum(card + 32 + 8, ["LOCAL_OR_SYNCABLE", "ACCOUNT"]);
          if (A.load.Bool(card + 32 + 17)) {
            card_ffi["metadata"]["isLocal"] = A.load.Bool(card + 32 + 12);
          }
          if (A.load.Bool(card + 32 + 18)) {
            card_ffi["metadata"]["isCached"] = A.load.Bool(card + 32 + 13);
          }
          if (A.load.Bool(card + 32 + 19)) {
            card_ffi["metadata"]["isMigratable"] = A.load.Bool(card + 32 + 14);
          }
          if (A.load.Bool(card + 32 + 20)) {
            card_ffi["metadata"]["isVirtualCardEnrollmentEligible"] = A.load.Bool(card + 32 + 15);
          }
          if (A.load.Bool(card + 32 + 21)) {
            card_ffi["metadata"]["isVirtualCardEnrolled"] = A.load.Bool(card + 32 + 16);
          }
        }

        const _ret = WEBEXT.autofillPrivate.saveCreditCard(card_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SaveIban": (): heap.Ref<boolean> => {
      if (WEBEXT?.autofillPrivate && "saveIban" in WEBEXT?.autofillPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SaveIban": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autofillPrivate.saveIban);
    },
    "call_SaveIban": (retPtr: Pointer, iban: Pointer): void => {
      const iban_ffi = {};

      iban_ffi["guid"] = A.load.Ref(iban + 0, undefined);
      iban_ffi["value"] = A.load.Ref(iban + 4, undefined);
      iban_ffi["nickname"] = A.load.Ref(iban + 8, undefined);
      if (A.load.Bool(iban + 12 + 22)) {
        iban_ffi["metadata"] = {};
        iban_ffi["metadata"]["summaryLabel"] = A.load.Ref(iban + 12 + 0, undefined);
        iban_ffi["metadata"]["summarySublabel"] = A.load.Ref(iban + 12 + 4, undefined);
        iban_ffi["metadata"]["source"] = A.load.Enum(iban + 12 + 8, ["LOCAL_OR_SYNCABLE", "ACCOUNT"]);
        if (A.load.Bool(iban + 12 + 17)) {
          iban_ffi["metadata"]["isLocal"] = A.load.Bool(iban + 12 + 12);
        }
        if (A.load.Bool(iban + 12 + 18)) {
          iban_ffi["metadata"]["isCached"] = A.load.Bool(iban + 12 + 13);
        }
        if (A.load.Bool(iban + 12 + 19)) {
          iban_ffi["metadata"]["isMigratable"] = A.load.Bool(iban + 12 + 14);
        }
        if (A.load.Bool(iban + 12 + 20)) {
          iban_ffi["metadata"]["isVirtualCardEnrollmentEligible"] = A.load.Bool(iban + 12 + 15);
        }
        if (A.load.Bool(iban + 12 + 21)) {
          iban_ffi["metadata"]["isVirtualCardEnrolled"] = A.load.Bool(iban + 12 + 16);
        }
      }

      const _ret = WEBEXT.autofillPrivate.saveIban(iban_ffi);
    },
    "try_SaveIban": (retPtr: Pointer, errPtr: Pointer, iban: Pointer): heap.Ref<boolean> => {
      try {
        const iban_ffi = {};

        iban_ffi["guid"] = A.load.Ref(iban + 0, undefined);
        iban_ffi["value"] = A.load.Ref(iban + 4, undefined);
        iban_ffi["nickname"] = A.load.Ref(iban + 8, undefined);
        if (A.load.Bool(iban + 12 + 22)) {
          iban_ffi["metadata"] = {};
          iban_ffi["metadata"]["summaryLabel"] = A.load.Ref(iban + 12 + 0, undefined);
          iban_ffi["metadata"]["summarySublabel"] = A.load.Ref(iban + 12 + 4, undefined);
          iban_ffi["metadata"]["source"] = A.load.Enum(iban + 12 + 8, ["LOCAL_OR_SYNCABLE", "ACCOUNT"]);
          if (A.load.Bool(iban + 12 + 17)) {
            iban_ffi["metadata"]["isLocal"] = A.load.Bool(iban + 12 + 12);
          }
          if (A.load.Bool(iban + 12 + 18)) {
            iban_ffi["metadata"]["isCached"] = A.load.Bool(iban + 12 + 13);
          }
          if (A.load.Bool(iban + 12 + 19)) {
            iban_ffi["metadata"]["isMigratable"] = A.load.Bool(iban + 12 + 14);
          }
          if (A.load.Bool(iban + 12 + 20)) {
            iban_ffi["metadata"]["isVirtualCardEnrollmentEligible"] = A.load.Bool(iban + 12 + 15);
          }
          if (A.load.Bool(iban + 12 + 21)) {
            iban_ffi["metadata"]["isVirtualCardEnrolled"] = A.load.Bool(iban + 12 + 16);
          }
        }

        const _ret = WEBEXT.autofillPrivate.saveIban(iban_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetCreditCardFIDOAuthEnabledState": (): heap.Ref<boolean> => {
      if (WEBEXT?.autofillPrivate && "setCreditCardFIDOAuthEnabledState" in WEBEXT?.autofillPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetCreditCardFIDOAuthEnabledState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autofillPrivate.setCreditCardFIDOAuthEnabledState);
    },
    "call_SetCreditCardFIDOAuthEnabledState": (retPtr: Pointer, enabled: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.autofillPrivate.setCreditCardFIDOAuthEnabledState(enabled === A.H.TRUE);
    },
    "try_SetCreditCardFIDOAuthEnabledState": (
      retPtr: Pointer,
      errPtr: Pointer,
      enabled: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autofillPrivate.setCreditCardFIDOAuthEnabledState(enabled === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
