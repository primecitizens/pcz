// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package browsingdata

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/browsingdata/bindings"
)

type DataTypeSet struct {
	// Appcache is "DataTypeSet.appcache"
	//
	// Optional
	//
	// NOTE: FFI_USE_Appcache MUST be set to true to make this field effective.
	Appcache bool
	// Cache is "DataTypeSet.cache"
	//
	// Optional
	//
	// NOTE: FFI_USE_Cache MUST be set to true to make this field effective.
	Cache bool
	// CacheStorage is "DataTypeSet.cacheStorage"
	//
	// Optional
	//
	// NOTE: FFI_USE_CacheStorage MUST be set to true to make this field effective.
	CacheStorage bool
	// Cookies is "DataTypeSet.cookies"
	//
	// Optional
	//
	// NOTE: FFI_USE_Cookies MUST be set to true to make this field effective.
	Cookies bool
	// Downloads is "DataTypeSet.downloads"
	//
	// Optional
	//
	// NOTE: FFI_USE_Downloads MUST be set to true to make this field effective.
	Downloads bool
	// FileSystems is "DataTypeSet.fileSystems"
	//
	// Optional
	//
	// NOTE: FFI_USE_FileSystems MUST be set to true to make this field effective.
	FileSystems bool
	// FormData is "DataTypeSet.formData"
	//
	// Optional
	//
	// NOTE: FFI_USE_FormData MUST be set to true to make this field effective.
	FormData bool
	// History is "DataTypeSet.history"
	//
	// Optional
	//
	// NOTE: FFI_USE_History MUST be set to true to make this field effective.
	History bool
	// IndexedDB is "DataTypeSet.indexedDB"
	//
	// Optional
	//
	// NOTE: FFI_USE_IndexedDB MUST be set to true to make this field effective.
	IndexedDB bool
	// LocalStorage is "DataTypeSet.localStorage"
	//
	// Optional
	//
	// NOTE: FFI_USE_LocalStorage MUST be set to true to make this field effective.
	LocalStorage bool
	// Passwords is "DataTypeSet.passwords"
	//
	// Optional
	//
	// NOTE: FFI_USE_Passwords MUST be set to true to make this field effective.
	Passwords bool
	// PluginData is "DataTypeSet.pluginData"
	//
	// Optional
	//
	// NOTE: FFI_USE_PluginData MUST be set to true to make this field effective.
	PluginData bool
	// ServerBoundCertificates is "DataTypeSet.serverBoundCertificates"
	//
	// Optional
	//
	// NOTE: FFI_USE_ServerBoundCertificates MUST be set to true to make this field effective.
	ServerBoundCertificates bool
	// ServiceWorkers is "DataTypeSet.serviceWorkers"
	//
	// Optional
	//
	// NOTE: FFI_USE_ServiceWorkers MUST be set to true to make this field effective.
	ServiceWorkers bool
	// WebSQL is "DataTypeSet.webSQL"
	//
	// Optional
	//
	// NOTE: FFI_USE_WebSQL MUST be set to true to make this field effective.
	WebSQL bool

	FFI_USE_Appcache                bool // for Appcache.
	FFI_USE_Cache                   bool // for Cache.
	FFI_USE_CacheStorage            bool // for CacheStorage.
	FFI_USE_Cookies                 bool // for Cookies.
	FFI_USE_Downloads               bool // for Downloads.
	FFI_USE_FileSystems             bool // for FileSystems.
	FFI_USE_FormData                bool // for FormData.
	FFI_USE_History                 bool // for History.
	FFI_USE_IndexedDB               bool // for IndexedDB.
	FFI_USE_LocalStorage            bool // for LocalStorage.
	FFI_USE_Passwords               bool // for Passwords.
	FFI_USE_PluginData              bool // for PluginData.
	FFI_USE_ServerBoundCertificates bool // for ServerBoundCertificates.
	FFI_USE_ServiceWorkers          bool // for ServiceWorkers.
	FFI_USE_WebSQL                  bool // for WebSQL.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DataTypeSet with all fields set.
func (p DataTypeSet) FromRef(ref js.Ref) DataTypeSet {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DataTypeSet in the application heap.
func (p DataTypeSet) New() js.Ref {
	return bindings.DataTypeSetJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DataTypeSet) UpdateFrom(ref js.Ref) {
	bindings.DataTypeSetJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DataTypeSet) Update(ref js.Ref) {
	bindings.DataTypeSetJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DataTypeSet) FreeMembers(recursive bool) {
}

type RemovalOptionsFieldOriginTypes struct {
	// Extension is "RemovalOptionsFieldOriginTypes.extension"
	//
	// Optional
	//
	// NOTE: FFI_USE_Extension MUST be set to true to make this field effective.
	Extension bool
	// ProtectedWeb is "RemovalOptionsFieldOriginTypes.protectedWeb"
	//
	// Optional
	//
	// NOTE: FFI_USE_ProtectedWeb MUST be set to true to make this field effective.
	ProtectedWeb bool
	// UnprotectedWeb is "RemovalOptionsFieldOriginTypes.unprotectedWeb"
	//
	// Optional
	//
	// NOTE: FFI_USE_UnprotectedWeb MUST be set to true to make this field effective.
	UnprotectedWeb bool

	FFI_USE_Extension      bool // for Extension.
	FFI_USE_ProtectedWeb   bool // for ProtectedWeb.
	FFI_USE_UnprotectedWeb bool // for UnprotectedWeb.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RemovalOptionsFieldOriginTypes with all fields set.
func (p RemovalOptionsFieldOriginTypes) FromRef(ref js.Ref) RemovalOptionsFieldOriginTypes {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RemovalOptionsFieldOriginTypes in the application heap.
func (p RemovalOptionsFieldOriginTypes) New() js.Ref {
	return bindings.RemovalOptionsFieldOriginTypesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RemovalOptionsFieldOriginTypes) UpdateFrom(ref js.Ref) {
	bindings.RemovalOptionsFieldOriginTypesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RemovalOptionsFieldOriginTypes) Update(ref js.Ref) {
	bindings.RemovalOptionsFieldOriginTypesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RemovalOptionsFieldOriginTypes) FreeMembers(recursive bool) {
}

type RemovalOptions struct {
	// ExcludeOrigins is "RemovalOptions.excludeOrigins"
	//
	// Optional
	ExcludeOrigins js.Array[js.String]
	// OriginTypes is "RemovalOptions.originTypes"
	//
	// Optional
	//
	// NOTE: OriginTypes.FFI_USE MUST be set to true to get OriginTypes used.
	OriginTypes RemovalOptionsFieldOriginTypes
	// Origins is "RemovalOptions.origins"
	//
	// Optional
	Origins js.Array[js.String]
	// Since is "RemovalOptions.since"
	//
	// Optional
	//
	// NOTE: FFI_USE_Since MUST be set to true to make this field effective.
	Since float64

	FFI_USE_Since bool // for Since.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RemovalOptions with all fields set.
func (p RemovalOptions) FromRef(ref js.Ref) RemovalOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RemovalOptions in the application heap.
func (p RemovalOptions) New() js.Ref {
	return bindings.RemovalOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RemovalOptions) UpdateFrom(ref js.Ref) {
	bindings.RemovalOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RemovalOptions) Update(ref js.Ref) {
	bindings.RemovalOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RemovalOptions) FreeMembers(recursive bool) {
	js.Free(
		p.ExcludeOrigins.Ref(),
		p.Origins.Ref(),
	)
	p.ExcludeOrigins = p.ExcludeOrigins.FromRef(js.Undefined)
	p.Origins = p.Origins.FromRef(js.Undefined)
	if recursive {
		p.OriginTypes.FreeMembers(true)
	}
}

type SettingsReturnType struct {
	// DataRemovalPermitted is "SettingsReturnType.dataRemovalPermitted"
	//
	// Required
	//
	// NOTE: DataRemovalPermitted.FFI_USE MUST be set to true to get DataRemovalPermitted used.
	DataRemovalPermitted DataTypeSet
	// DataToRemove is "SettingsReturnType.dataToRemove"
	//
	// Required
	//
	// NOTE: DataToRemove.FFI_USE MUST be set to true to get DataToRemove used.
	DataToRemove DataTypeSet
	// Options is "SettingsReturnType.options"
	//
	// Required
	//
	// NOTE: Options.FFI_USE MUST be set to true to get Options used.
	Options RemovalOptions

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SettingsReturnType with all fields set.
func (p SettingsReturnType) FromRef(ref js.Ref) SettingsReturnType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SettingsReturnType in the application heap.
func (p SettingsReturnType) New() js.Ref {
	return bindings.SettingsReturnTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SettingsReturnType) UpdateFrom(ref js.Ref) {
	bindings.SettingsReturnTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SettingsReturnType) Update(ref js.Ref) {
	bindings.SettingsReturnTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SettingsReturnType) FreeMembers(recursive bool) {
	if recursive {
		p.DataRemovalPermitted.FreeMembers(true)
		p.DataToRemove.FreeMembers(true)
		p.Options.FreeMembers(true)
	}
}

// HasFuncRemove returns true if the function "WEBEXT.browsingData.remove" exists.
func HasFuncRemove() bool {
	return js.True == bindings.HasFuncRemove()
}

// FuncRemove returns the function "WEBEXT.browsingData.remove".
func FuncRemove() (fn js.Func[func(options RemovalOptions, dataToRemove DataTypeSet) js.Promise[js.Void]]) {
	bindings.FuncRemove(
		js.Pointer(&fn),
	)
	return
}

// Remove calls the function "WEBEXT.browsingData.remove" directly.
func Remove(options RemovalOptions, dataToRemove DataTypeSet) (ret js.Promise[js.Void]) {
	bindings.CallRemove(
		js.Pointer(&ret),
		js.Pointer(&options),
		js.Pointer(&dataToRemove),
	)

	return
}

// TryRemove calls the function "WEBEXT.browsingData.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemove(options RemovalOptions, dataToRemove DataTypeSet) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemove(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
		js.Pointer(&dataToRemove),
	)

	return
}

// HasFuncRemoveAppcache returns true if the function "WEBEXT.browsingData.removeAppcache" exists.
func HasFuncRemoveAppcache() bool {
	return js.True == bindings.HasFuncRemoveAppcache()
}

// FuncRemoveAppcache returns the function "WEBEXT.browsingData.removeAppcache".
func FuncRemoveAppcache() (fn js.Func[func(options RemovalOptions) js.Promise[js.Void]]) {
	bindings.FuncRemoveAppcache(
		js.Pointer(&fn),
	)
	return
}

// RemoveAppcache calls the function "WEBEXT.browsingData.removeAppcache" directly.
func RemoveAppcache(options RemovalOptions) (ret js.Promise[js.Void]) {
	bindings.CallRemoveAppcache(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRemoveAppcache calls the function "WEBEXT.browsingData.removeAppcache"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveAppcache(options RemovalOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveAppcache(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncRemoveCache returns true if the function "WEBEXT.browsingData.removeCache" exists.
func HasFuncRemoveCache() bool {
	return js.True == bindings.HasFuncRemoveCache()
}

// FuncRemoveCache returns the function "WEBEXT.browsingData.removeCache".
func FuncRemoveCache() (fn js.Func[func(options RemovalOptions) js.Promise[js.Void]]) {
	bindings.FuncRemoveCache(
		js.Pointer(&fn),
	)
	return
}

// RemoveCache calls the function "WEBEXT.browsingData.removeCache" directly.
func RemoveCache(options RemovalOptions) (ret js.Promise[js.Void]) {
	bindings.CallRemoveCache(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRemoveCache calls the function "WEBEXT.browsingData.removeCache"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveCache(options RemovalOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveCache(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncRemoveCacheStorage returns true if the function "WEBEXT.browsingData.removeCacheStorage" exists.
func HasFuncRemoveCacheStorage() bool {
	return js.True == bindings.HasFuncRemoveCacheStorage()
}

// FuncRemoveCacheStorage returns the function "WEBEXT.browsingData.removeCacheStorage".
func FuncRemoveCacheStorage() (fn js.Func[func(options RemovalOptions) js.Promise[js.Void]]) {
	bindings.FuncRemoveCacheStorage(
		js.Pointer(&fn),
	)
	return
}

// RemoveCacheStorage calls the function "WEBEXT.browsingData.removeCacheStorage" directly.
func RemoveCacheStorage(options RemovalOptions) (ret js.Promise[js.Void]) {
	bindings.CallRemoveCacheStorage(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRemoveCacheStorage calls the function "WEBEXT.browsingData.removeCacheStorage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveCacheStorage(options RemovalOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveCacheStorage(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncRemoveCookies returns true if the function "WEBEXT.browsingData.removeCookies" exists.
func HasFuncRemoveCookies() bool {
	return js.True == bindings.HasFuncRemoveCookies()
}

// FuncRemoveCookies returns the function "WEBEXT.browsingData.removeCookies".
func FuncRemoveCookies() (fn js.Func[func(options RemovalOptions) js.Promise[js.Void]]) {
	bindings.FuncRemoveCookies(
		js.Pointer(&fn),
	)
	return
}

// RemoveCookies calls the function "WEBEXT.browsingData.removeCookies" directly.
func RemoveCookies(options RemovalOptions) (ret js.Promise[js.Void]) {
	bindings.CallRemoveCookies(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRemoveCookies calls the function "WEBEXT.browsingData.removeCookies"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveCookies(options RemovalOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveCookies(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncRemoveDownloads returns true if the function "WEBEXT.browsingData.removeDownloads" exists.
func HasFuncRemoveDownloads() bool {
	return js.True == bindings.HasFuncRemoveDownloads()
}

// FuncRemoveDownloads returns the function "WEBEXT.browsingData.removeDownloads".
func FuncRemoveDownloads() (fn js.Func[func(options RemovalOptions) js.Promise[js.Void]]) {
	bindings.FuncRemoveDownloads(
		js.Pointer(&fn),
	)
	return
}

// RemoveDownloads calls the function "WEBEXT.browsingData.removeDownloads" directly.
func RemoveDownloads(options RemovalOptions) (ret js.Promise[js.Void]) {
	bindings.CallRemoveDownloads(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRemoveDownloads calls the function "WEBEXT.browsingData.removeDownloads"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveDownloads(options RemovalOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveDownloads(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncRemoveFileSystems returns true if the function "WEBEXT.browsingData.removeFileSystems" exists.
func HasFuncRemoveFileSystems() bool {
	return js.True == bindings.HasFuncRemoveFileSystems()
}

// FuncRemoveFileSystems returns the function "WEBEXT.browsingData.removeFileSystems".
func FuncRemoveFileSystems() (fn js.Func[func(options RemovalOptions) js.Promise[js.Void]]) {
	bindings.FuncRemoveFileSystems(
		js.Pointer(&fn),
	)
	return
}

// RemoveFileSystems calls the function "WEBEXT.browsingData.removeFileSystems" directly.
func RemoveFileSystems(options RemovalOptions) (ret js.Promise[js.Void]) {
	bindings.CallRemoveFileSystems(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRemoveFileSystems calls the function "WEBEXT.browsingData.removeFileSystems"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveFileSystems(options RemovalOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveFileSystems(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncRemoveFormData returns true if the function "WEBEXT.browsingData.removeFormData" exists.
func HasFuncRemoveFormData() bool {
	return js.True == bindings.HasFuncRemoveFormData()
}

// FuncRemoveFormData returns the function "WEBEXT.browsingData.removeFormData".
func FuncRemoveFormData() (fn js.Func[func(options RemovalOptions) js.Promise[js.Void]]) {
	bindings.FuncRemoveFormData(
		js.Pointer(&fn),
	)
	return
}

// RemoveFormData calls the function "WEBEXT.browsingData.removeFormData" directly.
func RemoveFormData(options RemovalOptions) (ret js.Promise[js.Void]) {
	bindings.CallRemoveFormData(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRemoveFormData calls the function "WEBEXT.browsingData.removeFormData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveFormData(options RemovalOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveFormData(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncRemoveHistory returns true if the function "WEBEXT.browsingData.removeHistory" exists.
func HasFuncRemoveHistory() bool {
	return js.True == bindings.HasFuncRemoveHistory()
}

// FuncRemoveHistory returns the function "WEBEXT.browsingData.removeHistory".
func FuncRemoveHistory() (fn js.Func[func(options RemovalOptions) js.Promise[js.Void]]) {
	bindings.FuncRemoveHistory(
		js.Pointer(&fn),
	)
	return
}

// RemoveHistory calls the function "WEBEXT.browsingData.removeHistory" directly.
func RemoveHistory(options RemovalOptions) (ret js.Promise[js.Void]) {
	bindings.CallRemoveHistory(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRemoveHistory calls the function "WEBEXT.browsingData.removeHistory"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveHistory(options RemovalOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveHistory(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncRemoveIndexedDB returns true if the function "WEBEXT.browsingData.removeIndexedDB" exists.
func HasFuncRemoveIndexedDB() bool {
	return js.True == bindings.HasFuncRemoveIndexedDB()
}

// FuncRemoveIndexedDB returns the function "WEBEXT.browsingData.removeIndexedDB".
func FuncRemoveIndexedDB() (fn js.Func[func(options RemovalOptions) js.Promise[js.Void]]) {
	bindings.FuncRemoveIndexedDB(
		js.Pointer(&fn),
	)
	return
}

// RemoveIndexedDB calls the function "WEBEXT.browsingData.removeIndexedDB" directly.
func RemoveIndexedDB(options RemovalOptions) (ret js.Promise[js.Void]) {
	bindings.CallRemoveIndexedDB(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRemoveIndexedDB calls the function "WEBEXT.browsingData.removeIndexedDB"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveIndexedDB(options RemovalOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveIndexedDB(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncRemoveLocalStorage returns true if the function "WEBEXT.browsingData.removeLocalStorage" exists.
func HasFuncRemoveLocalStorage() bool {
	return js.True == bindings.HasFuncRemoveLocalStorage()
}

// FuncRemoveLocalStorage returns the function "WEBEXT.browsingData.removeLocalStorage".
func FuncRemoveLocalStorage() (fn js.Func[func(options RemovalOptions) js.Promise[js.Void]]) {
	bindings.FuncRemoveLocalStorage(
		js.Pointer(&fn),
	)
	return
}

// RemoveLocalStorage calls the function "WEBEXT.browsingData.removeLocalStorage" directly.
func RemoveLocalStorage(options RemovalOptions) (ret js.Promise[js.Void]) {
	bindings.CallRemoveLocalStorage(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRemoveLocalStorage calls the function "WEBEXT.browsingData.removeLocalStorage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveLocalStorage(options RemovalOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveLocalStorage(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncRemovePasswords returns true if the function "WEBEXT.browsingData.removePasswords" exists.
func HasFuncRemovePasswords() bool {
	return js.True == bindings.HasFuncRemovePasswords()
}

// FuncRemovePasswords returns the function "WEBEXT.browsingData.removePasswords".
func FuncRemovePasswords() (fn js.Func[func(options RemovalOptions) js.Promise[js.Void]]) {
	bindings.FuncRemovePasswords(
		js.Pointer(&fn),
	)
	return
}

// RemovePasswords calls the function "WEBEXT.browsingData.removePasswords" directly.
func RemovePasswords(options RemovalOptions) (ret js.Promise[js.Void]) {
	bindings.CallRemovePasswords(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRemovePasswords calls the function "WEBEXT.browsingData.removePasswords"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemovePasswords(options RemovalOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemovePasswords(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncRemovePluginData returns true if the function "WEBEXT.browsingData.removePluginData" exists.
func HasFuncRemovePluginData() bool {
	return js.True == bindings.HasFuncRemovePluginData()
}

// FuncRemovePluginData returns the function "WEBEXT.browsingData.removePluginData".
func FuncRemovePluginData() (fn js.Func[func(options RemovalOptions) js.Promise[js.Void]]) {
	bindings.FuncRemovePluginData(
		js.Pointer(&fn),
	)
	return
}

// RemovePluginData calls the function "WEBEXT.browsingData.removePluginData" directly.
func RemovePluginData(options RemovalOptions) (ret js.Promise[js.Void]) {
	bindings.CallRemovePluginData(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRemovePluginData calls the function "WEBEXT.browsingData.removePluginData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemovePluginData(options RemovalOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemovePluginData(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncRemoveServiceWorkers returns true if the function "WEBEXT.browsingData.removeServiceWorkers" exists.
func HasFuncRemoveServiceWorkers() bool {
	return js.True == bindings.HasFuncRemoveServiceWorkers()
}

// FuncRemoveServiceWorkers returns the function "WEBEXT.browsingData.removeServiceWorkers".
func FuncRemoveServiceWorkers() (fn js.Func[func(options RemovalOptions) js.Promise[js.Void]]) {
	bindings.FuncRemoveServiceWorkers(
		js.Pointer(&fn),
	)
	return
}

// RemoveServiceWorkers calls the function "WEBEXT.browsingData.removeServiceWorkers" directly.
func RemoveServiceWorkers(options RemovalOptions) (ret js.Promise[js.Void]) {
	bindings.CallRemoveServiceWorkers(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRemoveServiceWorkers calls the function "WEBEXT.browsingData.removeServiceWorkers"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveServiceWorkers(options RemovalOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveServiceWorkers(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncRemoveWebSQL returns true if the function "WEBEXT.browsingData.removeWebSQL" exists.
func HasFuncRemoveWebSQL() bool {
	return js.True == bindings.HasFuncRemoveWebSQL()
}

// FuncRemoveWebSQL returns the function "WEBEXT.browsingData.removeWebSQL".
func FuncRemoveWebSQL() (fn js.Func[func(options RemovalOptions) js.Promise[js.Void]]) {
	bindings.FuncRemoveWebSQL(
		js.Pointer(&fn),
	)
	return
}

// RemoveWebSQL calls the function "WEBEXT.browsingData.removeWebSQL" directly.
func RemoveWebSQL(options RemovalOptions) (ret js.Promise[js.Void]) {
	bindings.CallRemoveWebSQL(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRemoveWebSQL calls the function "WEBEXT.browsingData.removeWebSQL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveWebSQL(options RemovalOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveWebSQL(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncSettings returns true if the function "WEBEXT.browsingData.settings" exists.
func HasFuncSettings() bool {
	return js.True == bindings.HasFuncSettings()
}

// FuncSettings returns the function "WEBEXT.browsingData.settings".
func FuncSettings() (fn js.Func[func() js.Promise[SettingsReturnType]]) {
	bindings.FuncSettings(
		js.Pointer(&fn),
	)
	return
}

// Settings calls the function "WEBEXT.browsingData.settings" directly.
func Settings() (ret js.Promise[SettingsReturnType]) {
	bindings.CallSettings(
		js.Pointer(&ret),
	)

	return
}

// TrySettings calls the function "WEBEXT.browsingData.settings"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySettings() (ret js.Promise[SettingsReturnType], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySettings(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}
