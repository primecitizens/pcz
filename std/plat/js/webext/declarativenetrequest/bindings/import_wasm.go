// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/ffi/js"
)

type (
	_ unsafe.Pointer
	_ js.Ref
)

//go:wasmimport plat/js/webext/declarativenetrequest store_Ruleset
//go:noescape
func RulesetJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_Ruleset
//go:noescape
func RulesetJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest store_DNRInfo
//go:noescape
func DNRInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_DNRInfo
//go:noescape
func DNRInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest constof_DomainType
//go:noescape
func ConstOfDomainType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativenetrequest store_TabActionCountUpdate
//go:noescape
func TabActionCountUpdateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_TabActionCountUpdate
//go:noescape
func TabActionCountUpdateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest store_ExtensionActionOptions
//go:noescape
func ExtensionActionOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_ExtensionActionOptions
//go:noescape
func ExtensionActionOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest store_GetDisabledRuleIdsOptions
//go:noescape
func GetDisabledRuleIdsOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_GetDisabledRuleIdsOptions
//go:noescape
func GetDisabledRuleIdsOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest store_MatchedRule
//go:noescape
func MatchedRuleJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_MatchedRule
//go:noescape
func MatchedRuleJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest store_MatchedRuleInfo
//go:noescape
func MatchedRuleInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_MatchedRuleInfo
//go:noescape
func MatchedRuleInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest store_RulesMatchedDetails
//go:noescape
func RulesMatchedDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_RulesMatchedDetails
//go:noescape
func RulesMatchedDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest constof_ResourceType
//go:noescape
func ConstOfResourceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativenetrequest constof_RequestMethod
//go:noescape
func ConstOfRequestMethod(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativenetrequest store_RuleCondition
//go:noescape
func RuleConditionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_RuleCondition
//go:noescape
func RuleConditionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest constof_RuleActionType
//go:noescape
func ConstOfRuleActionType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativenetrequest store_QueryKeyValue
//go:noescape
func QueryKeyValueJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_QueryKeyValue
//go:noescape
func QueryKeyValueJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest store_QueryTransform
//go:noescape
func QueryTransformJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_QueryTransform
//go:noescape
func QueryTransformJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest store_URLTransform
//go:noescape
func URLTransformJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_URLTransform
//go:noescape
func URLTransformJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest store_Redirect
//go:noescape
func RedirectJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_Redirect
//go:noescape
func RedirectJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest constof_HeaderOperation
//go:noescape
func ConstOfHeaderOperation(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativenetrequest store_ModifyHeaderInfo
//go:noescape
func ModifyHeaderInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_ModifyHeaderInfo
//go:noescape
func ModifyHeaderInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest store_RuleAction
//go:noescape
func RuleActionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_RuleAction
//go:noescape
func RuleActionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest store_Rule
//go:noescape
func RuleJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_Rule
//go:noescape
func RuleJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest store_GetRulesFilter
//go:noescape
func GetRulesFilterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_GetRulesFilter
//go:noescape
func GetRulesFilterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest constof_UnsupportedRegexReason
//go:noescape
func ConstOfUnsupportedRegexReason(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativenetrequest store_IsRegexSupportedResult
//go:noescape
func IsRegexSupportedResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_IsRegexSupportedResult
//go:noescape
func IsRegexSupportedResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest store_ManifestKeys
//go:noescape
func ManifestKeysJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_ManifestKeys
//go:noescape
func ManifestKeysJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest store_RequestDetails
//go:noescape
func RequestDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_RequestDetails
//go:noescape
func RequestDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest store_MatchedRuleInfoDebug
//go:noescape
func MatchedRuleInfoDebugJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_MatchedRuleInfoDebug
//go:noescape
func MatchedRuleInfoDebugJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest store_MatchedRulesFilter
//go:noescape
func MatchedRulesFilterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_MatchedRulesFilter
//go:noescape
func MatchedRulesFilterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest has_Properties_GUARANTEED_MINIMUM_STATIC_RULES
//go:noescape
func HasFuncPropertiesGUARANTEED_MINIMUM_STATIC_RULES(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_Properties_GUARANTEED_MINIMUM_STATIC_RULES
//go:noescape
func FuncPropertiesGUARANTEED_MINIMUM_STATIC_RULES(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_Properties_GUARANTEED_MINIMUM_STATIC_RULES
//go:noescape
func CallPropertiesGUARANTEED_MINIMUM_STATIC_RULES(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest try_Properties_GUARANTEED_MINIMUM_STATIC_RULES
//go:noescape
func TryPropertiesGUARANTEED_MINIMUM_STATIC_RULES(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_Properties_MAX_NUMBER_OF_DYNAMIC_RULES
//go:noescape
func HasFuncPropertiesMAX_NUMBER_OF_DYNAMIC_RULES(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_Properties_MAX_NUMBER_OF_DYNAMIC_RULES
//go:noescape
func FuncPropertiesMAX_NUMBER_OF_DYNAMIC_RULES(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_Properties_MAX_NUMBER_OF_DYNAMIC_RULES
//go:noescape
func CallPropertiesMAX_NUMBER_OF_DYNAMIC_RULES(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest try_Properties_MAX_NUMBER_OF_DYNAMIC_RULES
//go:noescape
func TryPropertiesMAX_NUMBER_OF_DYNAMIC_RULES(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_Properties_MAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES
//go:noescape
func HasFuncPropertiesMAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_Properties_MAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES
//go:noescape
func FuncPropertiesMAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_Properties_MAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES
//go:noescape
func CallPropertiesMAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest try_Properties_MAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES
//go:noescape
func TryPropertiesMAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_Properties_GETMATCHEDRULES_QUOTA_INTERVAL
//go:noescape
func HasFuncPropertiesGETMATCHEDRULES_QUOTA_INTERVAL(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_Properties_GETMATCHEDRULES_QUOTA_INTERVAL
//go:noescape
func FuncPropertiesGETMATCHEDRULES_QUOTA_INTERVAL(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_Properties_GETMATCHEDRULES_QUOTA_INTERVAL
//go:noescape
func CallPropertiesGETMATCHEDRULES_QUOTA_INTERVAL(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest try_Properties_GETMATCHEDRULES_QUOTA_INTERVAL
//go:noescape
func TryPropertiesGETMATCHEDRULES_QUOTA_INTERVAL(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_Properties_MAX_GETMATCHEDRULES_CALLS_PER_INTERVAL
//go:noescape
func HasFuncPropertiesMAX_GETMATCHEDRULES_CALLS_PER_INTERVAL(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_Properties_MAX_GETMATCHEDRULES_CALLS_PER_INTERVAL
//go:noescape
func FuncPropertiesMAX_GETMATCHEDRULES_CALLS_PER_INTERVAL(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_Properties_MAX_GETMATCHEDRULES_CALLS_PER_INTERVAL
//go:noescape
func CallPropertiesMAX_GETMATCHEDRULES_CALLS_PER_INTERVAL(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest try_Properties_MAX_GETMATCHEDRULES_CALLS_PER_INTERVAL
//go:noescape
func TryPropertiesMAX_GETMATCHEDRULES_CALLS_PER_INTERVAL(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_Properties_MAX_NUMBER_OF_REGEX_RULES
//go:noescape
func HasFuncPropertiesMAX_NUMBER_OF_REGEX_RULES(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_Properties_MAX_NUMBER_OF_REGEX_RULES
//go:noescape
func FuncPropertiesMAX_NUMBER_OF_REGEX_RULES(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_Properties_MAX_NUMBER_OF_REGEX_RULES
//go:noescape
func CallPropertiesMAX_NUMBER_OF_REGEX_RULES(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest try_Properties_MAX_NUMBER_OF_REGEX_RULES
//go:noescape
func TryPropertiesMAX_NUMBER_OF_REGEX_RULES(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_Properties_MAX_NUMBER_OF_STATIC_RULESETS
//go:noescape
func HasFuncPropertiesMAX_NUMBER_OF_STATIC_RULESETS(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_Properties_MAX_NUMBER_OF_STATIC_RULESETS
//go:noescape
func FuncPropertiesMAX_NUMBER_OF_STATIC_RULESETS(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_Properties_MAX_NUMBER_OF_STATIC_RULESETS
//go:noescape
func CallPropertiesMAX_NUMBER_OF_STATIC_RULESETS(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest try_Properties_MAX_NUMBER_OF_STATIC_RULESETS
//go:noescape
func TryPropertiesMAX_NUMBER_OF_STATIC_RULESETS(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_Properties_MAX_NUMBER_OF_ENABLED_STATIC_RULESETS
//go:noescape
func HasFuncPropertiesMAX_NUMBER_OF_ENABLED_STATIC_RULESETS(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_Properties_MAX_NUMBER_OF_ENABLED_STATIC_RULESETS
//go:noescape
func FuncPropertiesMAX_NUMBER_OF_ENABLED_STATIC_RULESETS(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_Properties_MAX_NUMBER_OF_ENABLED_STATIC_RULESETS
//go:noescape
func CallPropertiesMAX_NUMBER_OF_ENABLED_STATIC_RULESETS(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest try_Properties_MAX_NUMBER_OF_ENABLED_STATIC_RULESETS
//go:noescape
func TryPropertiesMAX_NUMBER_OF_ENABLED_STATIC_RULESETS(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_Properties_DYNAMIC_RULESET_ID
//go:noescape
func HasFuncPropertiesDYNAMIC_RULESET_ID(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_Properties_DYNAMIC_RULESET_ID
//go:noescape
func FuncPropertiesDYNAMIC_RULESET_ID(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_Properties_DYNAMIC_RULESET_ID
//go:noescape
func CallPropertiesDYNAMIC_RULESET_ID(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest try_Properties_DYNAMIC_RULESET_ID
//go:noescape
func TryPropertiesDYNAMIC_RULESET_ID(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_Properties_SESSION_RULESET_ID
//go:noescape
func HasFuncPropertiesSESSION_RULESET_ID(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_Properties_SESSION_RULESET_ID
//go:noescape
func FuncPropertiesSESSION_RULESET_ID(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_Properties_SESSION_RULESET_ID
//go:noescape
func CallPropertiesSESSION_RULESET_ID(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest try_Properties_SESSION_RULESET_ID
//go:noescape
func TryPropertiesSESSION_RULESET_ID(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest store_RegexOptions
//go:noescape
func RegexOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_RegexOptions
//go:noescape
func RegexOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest store_TestMatchOutcomeResult
//go:noescape
func TestMatchOutcomeResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_TestMatchOutcomeResult
//go:noescape
func TestMatchOutcomeResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest store_TestMatchRequestDetails
//go:noescape
func TestMatchRequestDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_TestMatchRequestDetails
//go:noescape
func TestMatchRequestDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest store_UpdateRuleOptions
//go:noescape
func UpdateRuleOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_UpdateRuleOptions
//go:noescape
func UpdateRuleOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest store_UpdateRulesetOptions
//go:noescape
func UpdateRulesetOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_UpdateRulesetOptions
//go:noescape
func UpdateRulesetOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest store_UpdateStaticRulesOptions
//go:noescape
func UpdateStaticRulesOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest load_UpdateStaticRulesOptions
//go:noescape
func UpdateStaticRulesOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest has_GetAvailableStaticRuleCount
//go:noescape
func HasFuncGetAvailableStaticRuleCount() js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_GetAvailableStaticRuleCount
//go:noescape
func FuncGetAvailableStaticRuleCount(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_GetAvailableStaticRuleCount
//go:noescape
func CallGetAvailableStaticRuleCount(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest try_GetAvailableStaticRuleCount
//go:noescape
func TryGetAvailableStaticRuleCount(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_GetDisabledRuleIds
//go:noescape
func HasFuncGetDisabledRuleIds() js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_GetDisabledRuleIds
//go:noescape
func FuncGetDisabledRuleIds(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_GetDisabledRuleIds
//go:noescape
func CallGetDisabledRuleIds(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest try_GetDisabledRuleIds
//go:noescape
func TryGetDisabledRuleIds(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_GetDynamicRules
//go:noescape
func HasFuncGetDynamicRules() js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_GetDynamicRules
//go:noescape
func FuncGetDynamicRules(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_GetDynamicRules
//go:noescape
func CallGetDynamicRules(
	retPtr unsafe.Pointer,
	filter unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest try_GetDynamicRules
//go:noescape
func TryGetDynamicRules(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	filter unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_GetEnabledRulesets
//go:noescape
func HasFuncGetEnabledRulesets() js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_GetEnabledRulesets
//go:noescape
func FuncGetEnabledRulesets(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_GetEnabledRulesets
//go:noescape
func CallGetEnabledRulesets(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest try_GetEnabledRulesets
//go:noescape
func TryGetEnabledRulesets(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_GetMatchedRules
//go:noescape
func HasFuncGetMatchedRules() js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_GetMatchedRules
//go:noescape
func FuncGetMatchedRules(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_GetMatchedRules
//go:noescape
func CallGetMatchedRules(
	retPtr unsafe.Pointer,
	filter unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest try_GetMatchedRules
//go:noescape
func TryGetMatchedRules(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	filter unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_GetSessionRules
//go:noescape
func HasFuncGetSessionRules() js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_GetSessionRules
//go:noescape
func FuncGetSessionRules(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_GetSessionRules
//go:noescape
func CallGetSessionRules(
	retPtr unsafe.Pointer,
	filter unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest try_GetSessionRules
//go:noescape
func TryGetSessionRules(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	filter unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_IsRegexSupported
//go:noescape
func HasFuncIsRegexSupported() js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_IsRegexSupported
//go:noescape
func FuncIsRegexSupported(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_IsRegexSupported
//go:noescape
func CallIsRegexSupported(
	retPtr unsafe.Pointer,
	regexOptions unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest try_IsRegexSupported
//go:noescape
func TryIsRegexSupported(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	regexOptions unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_OnRuleMatchedDebug
//go:noescape
func HasFuncOnRuleMatchedDebug() js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_OnRuleMatchedDebug
//go:noescape
func FuncOnRuleMatchedDebug(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_OnRuleMatchedDebug
//go:noescape
func CallOnRuleMatchedDebug(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest try_OnRuleMatchedDebug
//go:noescape
func TryOnRuleMatchedDebug(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_OffRuleMatchedDebug
//go:noescape
func HasFuncOffRuleMatchedDebug() js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_OffRuleMatchedDebug
//go:noescape
func FuncOffRuleMatchedDebug(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_OffRuleMatchedDebug
//go:noescape
func CallOffRuleMatchedDebug(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest try_OffRuleMatchedDebug
//go:noescape
func TryOffRuleMatchedDebug(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_HasOnRuleMatchedDebug
//go:noescape
func HasFuncHasOnRuleMatchedDebug() js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_HasOnRuleMatchedDebug
//go:noescape
func FuncHasOnRuleMatchedDebug(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_HasOnRuleMatchedDebug
//go:noescape
func CallHasOnRuleMatchedDebug(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest try_HasOnRuleMatchedDebug
//go:noescape
func TryHasOnRuleMatchedDebug(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_SetExtensionActionOptions
//go:noescape
func HasFuncSetExtensionActionOptions() js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_SetExtensionActionOptions
//go:noescape
func FuncSetExtensionActionOptions(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_SetExtensionActionOptions
//go:noescape
func CallSetExtensionActionOptions(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest try_SetExtensionActionOptions
//go:noescape
func TrySetExtensionActionOptions(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_TestMatchOutcome
//go:noescape
func HasFuncTestMatchOutcome() js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_TestMatchOutcome
//go:noescape
func FuncTestMatchOutcome(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_TestMatchOutcome
//go:noescape
func CallTestMatchOutcome(
	retPtr unsafe.Pointer,
	request unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest try_TestMatchOutcome
//go:noescape
func TryTestMatchOutcome(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_UpdateDynamicRules
//go:noescape
func HasFuncUpdateDynamicRules() js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_UpdateDynamicRules
//go:noescape
func FuncUpdateDynamicRules(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_UpdateDynamicRules
//go:noescape
func CallUpdateDynamicRules(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest try_UpdateDynamicRules
//go:noescape
func TryUpdateDynamicRules(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_UpdateEnabledRulesets
//go:noescape
func HasFuncUpdateEnabledRulesets() js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_UpdateEnabledRulesets
//go:noescape
func FuncUpdateEnabledRulesets(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_UpdateEnabledRulesets
//go:noescape
func CallUpdateEnabledRulesets(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest try_UpdateEnabledRulesets
//go:noescape
func TryUpdateEnabledRulesets(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_UpdateSessionRules
//go:noescape
func HasFuncUpdateSessionRules() js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_UpdateSessionRules
//go:noescape
func FuncUpdateSessionRules(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_UpdateSessionRules
//go:noescape
func CallUpdateSessionRules(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest try_UpdateSessionRules
//go:noescape
func TryUpdateSessionRules(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativenetrequest has_UpdateStaticRules
//go:noescape
func HasFuncUpdateStaticRules() js.Ref

//go:wasmimport plat/js/webext/declarativenetrequest func_UpdateStaticRules
//go:noescape
func FuncUpdateStaticRules(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest call_UpdateStaticRules
//go:noescape
func CallUpdateStaticRules(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativenetrequest try_UpdateStaticRules
//go:noescape
func TryUpdateStaticRules(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)
