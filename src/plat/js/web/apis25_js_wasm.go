// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/ffi/js"
	"github.com/primecitizens/std/plat/js/web/bindings"
)

func _() {
	var (
		_ abi.FuncID
	)
	assert.TODO()
}

type MIDIOptions struct {
	// Sysex is "MIDIOptions.sysex"
	//
	// Optional
	Sysex bool
	// Software is "MIDIOptions.software"
	//
	// Optional
	Software bool

	FFI_USE_Sysex    bool // for Sysex.
	FFI_USE_Software bool // for Software.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MIDIOptions with all fields set.
func (p MIDIOptions) FromRef(ref js.Ref) MIDIOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MIDIOptions MIDIOptions [// MIDIOptions] [0x140006ce960 0x140006ceaa0 0x140006cea00 0x140006ceb40] 0x14000575818 {0 0}} in the application heap.
func (p MIDIOptions) New() js.Ref {
	return bindings.MIDIOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MIDIOptions) UpdateFrom(ref js.Ref) {
	bindings.MIDIOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MIDIOptions) Update(ref js.Ref) {
	bindings.MIDIOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type HIDUnitSystem uint32

const (
	_ HIDUnitSystem = iota

	HIDUnitSystem_NONE
	HIDUnitSystem_SI_LINEAR
	HIDUnitSystem_SI_ROTATION
	HIDUnitSystem_ENGLISH_LINEAR
	HIDUnitSystem_ENGLISH_ROTATION
	HIDUnitSystem_VENDOR_DEFINED
	HIDUnitSystem_RESERVED
)

func (HIDUnitSystem) FromRef(str js.Ref) HIDUnitSystem {
	return HIDUnitSystem(bindings.ConstOfHIDUnitSystem(str))
}

func (x HIDUnitSystem) String() (string, bool) {
	switch x {
	case HIDUnitSystem_NONE:
		return "none", true
	case HIDUnitSystem_SI_LINEAR:
		return "si-linear", true
	case HIDUnitSystem_SI_ROTATION:
		return "si-rotation", true
	case HIDUnitSystem_ENGLISH_LINEAR:
		return "english-linear", true
	case HIDUnitSystem_ENGLISH_ROTATION:
		return "english-rotation", true
	case HIDUnitSystem_VENDOR_DEFINED:
		return "vendor-defined", true
	case HIDUnitSystem_RESERVED:
		return "reserved", true
	default:
		return "", false
	}
}

type HIDReportItem struct {
	// IsAbsolute is "HIDReportItem.isAbsolute"
	//
	// Optional
	IsAbsolute bool
	// IsArray is "HIDReportItem.isArray"
	//
	// Optional
	IsArray bool
	// IsBufferedBytes is "HIDReportItem.isBufferedBytes"
	//
	// Optional
	IsBufferedBytes bool
	// IsConstant is "HIDReportItem.isConstant"
	//
	// Optional
	IsConstant bool
	// IsLinear is "HIDReportItem.isLinear"
	//
	// Optional
	IsLinear bool
	// IsRange is "HIDReportItem.isRange"
	//
	// Optional
	IsRange bool
	// IsVolatile is "HIDReportItem.isVolatile"
	//
	// Optional
	IsVolatile bool
	// HasNull is "HIDReportItem.hasNull"
	//
	// Optional
	HasNull bool
	// HasPreferredState is "HIDReportItem.hasPreferredState"
	//
	// Optional
	HasPreferredState bool
	// Wrap is "HIDReportItem.wrap"
	//
	// Optional
	Wrap bool
	// Usages is "HIDReportItem.usages"
	//
	// Optional
	Usages js.Array[uint32]
	// UsageMinimum is "HIDReportItem.usageMinimum"
	//
	// Optional
	UsageMinimum uint32
	// UsageMaximum is "HIDReportItem.usageMaximum"
	//
	// Optional
	UsageMaximum uint32
	// ReportSize is "HIDReportItem.reportSize"
	//
	// Optional
	ReportSize uint16
	// ReportCount is "HIDReportItem.reportCount"
	//
	// Optional
	ReportCount uint16
	// UnitExponent is "HIDReportItem.unitExponent"
	//
	// Optional
	UnitExponent int8
	// UnitSystem is "HIDReportItem.unitSystem"
	//
	// Optional
	UnitSystem HIDUnitSystem
	// UnitFactorLengthExponent is "HIDReportItem.unitFactorLengthExponent"
	//
	// Optional
	UnitFactorLengthExponent int8
	// UnitFactorMassExponent is "HIDReportItem.unitFactorMassExponent"
	//
	// Optional
	UnitFactorMassExponent int8
	// UnitFactorTimeExponent is "HIDReportItem.unitFactorTimeExponent"
	//
	// Optional
	UnitFactorTimeExponent int8
	// UnitFactorTemperatureExponent is "HIDReportItem.unitFactorTemperatureExponent"
	//
	// Optional
	UnitFactorTemperatureExponent int8
	// UnitFactorCurrentExponent is "HIDReportItem.unitFactorCurrentExponent"
	//
	// Optional
	UnitFactorCurrentExponent int8
	// UnitFactorLuminousIntensityExponent is "HIDReportItem.unitFactorLuminousIntensityExponent"
	//
	// Optional
	UnitFactorLuminousIntensityExponent int8
	// LogicalMinimum is "HIDReportItem.logicalMinimum"
	//
	// Optional
	LogicalMinimum int32
	// LogicalMaximum is "HIDReportItem.logicalMaximum"
	//
	// Optional
	LogicalMaximum int32
	// PhysicalMinimum is "HIDReportItem.physicalMinimum"
	//
	// Optional
	PhysicalMinimum int32
	// PhysicalMaximum is "HIDReportItem.physicalMaximum"
	//
	// Optional
	PhysicalMaximum int32
	// Strings is "HIDReportItem.strings"
	//
	// Optional
	Strings js.Array[js.String]

	FFI_USE_IsAbsolute                          bool // for IsAbsolute.
	FFI_USE_IsArray                             bool // for IsArray.
	FFI_USE_IsBufferedBytes                     bool // for IsBufferedBytes.
	FFI_USE_IsConstant                          bool // for IsConstant.
	FFI_USE_IsLinear                            bool // for IsLinear.
	FFI_USE_IsRange                             bool // for IsRange.
	FFI_USE_IsVolatile                          bool // for IsVolatile.
	FFI_USE_HasNull                             bool // for HasNull.
	FFI_USE_HasPreferredState                   bool // for HasPreferredState.
	FFI_USE_Wrap                                bool // for Wrap.
	FFI_USE_UsageMinimum                        bool // for UsageMinimum.
	FFI_USE_UsageMaximum                        bool // for UsageMaximum.
	FFI_USE_ReportSize                          bool // for ReportSize.
	FFI_USE_ReportCount                         bool // for ReportCount.
	FFI_USE_UnitExponent                        bool // for UnitExponent.
	FFI_USE_UnitFactorLengthExponent            bool // for UnitFactorLengthExponent.
	FFI_USE_UnitFactorMassExponent              bool // for UnitFactorMassExponent.
	FFI_USE_UnitFactorTimeExponent              bool // for UnitFactorTimeExponent.
	FFI_USE_UnitFactorTemperatureExponent       bool // for UnitFactorTemperatureExponent.
	FFI_USE_UnitFactorCurrentExponent           bool // for UnitFactorCurrentExponent.
	FFI_USE_UnitFactorLuminousIntensityExponent bool // for UnitFactorLuminousIntensityExponent.
	FFI_USE_LogicalMinimum                      bool // for LogicalMinimum.
	FFI_USE_LogicalMaximum                      bool // for LogicalMaximum.
	FFI_USE_PhysicalMinimum                     bool // for PhysicalMinimum.
	FFI_USE_PhysicalMaximum                     bool // for PhysicalMaximum.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HIDReportItem with all fields set.
func (p HIDReportItem) FromRef(ref js.Ref) HIDReportItem {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 HIDReportItem HIDReportItem [// HIDReportItem] [0x140006cf220 0x140006cf360 0x140006cf4a0 0x140006cf5e0 0x140006cf720 0x140006cf860 0x140006cf9a0 0x140006cfae0 0x140006cfc20 0x140006cfd60 0x140006cfea0 0x140006cff40 0x140006f40a0 0x140006f41e0 0x140006f4320 0x140006f4460 0x140006f45a0 0x140006f4640 0x140006f4780 0x140006f48c0 0x140006f4a00 0x140006f4b40 0x140006f4c80 0x140006f4dc0 0x140006f4f00 0x140006f5040 0x140006f5180 0x140006f52c0 0x140006cf2c0 0x140006cf400 0x140006cf540 0x140006cf680 0x140006cf7c0 0x140006cf900 0x140006cfa40 0x140006cfb80 0x140006cfcc0 0x140006cfe00 0x140006f4000 0x140006f4140 0x140006f4280 0x140006f43c0 0x140006f4500 0x140006f46e0 0x140006f4820 0x140006f4960 0x140006f4aa0 0x140006f4be0 0x140006f4d20 0x140006f4e60 0x140006f4fa0 0x140006f50e0 0x140006f5220] 0x14000575968 {0 0}} in the application heap.
func (p HIDReportItem) New() js.Ref {
	return bindings.HIDReportItemJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p HIDReportItem) UpdateFrom(ref js.Ref) {
	bindings.HIDReportItemJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p HIDReportItem) Update(ref js.Ref) {
	bindings.HIDReportItemJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type HIDReportInfo struct {
	// ReportId is "HIDReportInfo.reportId"
	//
	// Optional
	ReportId uint8
	// Items is "HIDReportInfo.items"
	//
	// Optional
	Items js.Array[HIDReportItem]

	FFI_USE_ReportId bool // for ReportId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HIDReportInfo with all fields set.
func (p HIDReportInfo) FromRef(ref js.Ref) HIDReportInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 HIDReportInfo HIDReportInfo [// HIDReportInfo] [0x140006cf0e0 0x140006f5360 0x140006cf180] 0x14000575950 {0 0}} in the application heap.
func (p HIDReportInfo) New() js.Ref {
	return bindings.HIDReportInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p HIDReportInfo) UpdateFrom(ref js.Ref) {
	bindings.HIDReportInfoJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p HIDReportInfo) Update(ref js.Ref) {
	bindings.HIDReportInfoJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type HIDCollectionInfo struct {
	// UsagePage is "HIDCollectionInfo.usagePage"
	//
	// Optional
	UsagePage uint16
	// Usage is "HIDCollectionInfo.usage"
	//
	// Optional
	Usage uint16
	// Type is "HIDCollectionInfo.type"
	//
	// Optional
	Type uint8
	// Children is "HIDCollectionInfo.children"
	//
	// Optional
	Children js.Array[HIDCollectionInfo]
	// InputReports is "HIDCollectionInfo.inputReports"
	//
	// Optional
	InputReports js.Array[HIDReportInfo]
	// OutputReports is "HIDCollectionInfo.outputReports"
	//
	// Optional
	OutputReports js.Array[HIDReportInfo]
	// FeatureReports is "HIDCollectionInfo.featureReports"
	//
	// Optional
	FeatureReports js.Array[HIDReportInfo]

	FFI_USE_UsagePage bool // for UsagePage.
	FFI_USE_Usage     bool // for Usage.
	FFI_USE_Type      bool // for Type.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HIDCollectionInfo with all fields set.
func (p HIDCollectionInfo) FromRef(ref js.Ref) HIDCollectionInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 HIDCollectionInfo HIDCollectionInfo [// HIDCollectionInfo] [0x140006cec80 0x140006cedc0 0x140006cef00 0x140006cf040 0x140006f5400 0x140006f54a0 0x140006f5540 0x140006ced20 0x140006cee60 0x140006cefa0] 0x14000575908 {0 0}} in the application heap.
func (p HIDCollectionInfo) New() js.Ref {
	return bindings.HIDCollectionInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p HIDCollectionInfo) UpdateFrom(ref js.Ref) {
	bindings.HIDCollectionInfoJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p HIDCollectionInfo) Update(ref js.Ref) {
	bindings.HIDCollectionInfoJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type HIDDevice struct {
	EventTarget
}

func (this HIDDevice) Once() HIDDevice {
	this.Ref().Once()
	return this
}

func (this HIDDevice) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this HIDDevice) FromRef(ref js.Ref) HIDDevice {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this HIDDevice) Free() {
	this.Ref().Free()
}

// Opened returns the value of property "HIDDevice.opened".
//
// The returned bool will be false if there is no such property.
func (this HIDDevice) Opened() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHIDDeviceOpened(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// VendorId returns the value of property "HIDDevice.vendorId".
//
// The returned bool will be false if there is no such property.
func (this HIDDevice) VendorId() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetHIDDeviceVendorId(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// ProductId returns the value of property "HIDDevice.productId".
//
// The returned bool will be false if there is no such property.
func (this HIDDevice) ProductId() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetHIDDeviceProductId(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// ProductName returns the value of property "HIDDevice.productName".
//
// The returned bool will be false if there is no such property.
func (this HIDDevice) ProductName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHIDDeviceProductName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Collections returns the value of property "HIDDevice.collections".
//
// The returned bool will be false if there is no such property.
func (this HIDDevice) Collections() (js.FrozenArray[HIDCollectionInfo], bool) {
	var _ok bool
	_ret := bindings.GetHIDDeviceCollections(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[HIDCollectionInfo]{}.FromRef(_ret), _ok
}

// Open calls the method "HIDDevice.open".
//
// The returned bool will be false if there is no such method.
func (this HIDDevice) Open() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallHIDDeviceOpen(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// OpenFunc returns the method "HIDDevice.open".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HIDDevice) OpenFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.HIDDeviceOpenFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "HIDDevice.close".
//
// The returned bool will be false if there is no such method.
func (this HIDDevice) Close() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallHIDDeviceClose(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// CloseFunc returns the method "HIDDevice.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HIDDevice) CloseFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.HIDDeviceCloseFunc(
			this.Ref(),
		),
	)
}

// Forget calls the method "HIDDevice.forget".
//
// The returned bool will be false if there is no such method.
func (this HIDDevice) Forget() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallHIDDeviceForget(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ForgetFunc returns the method "HIDDevice.forget".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HIDDevice) ForgetFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.HIDDeviceForgetFunc(
			this.Ref(),
		),
	)
}

// SendReport calls the method "HIDDevice.sendReport".
//
// The returned bool will be false if there is no such method.
func (this HIDDevice) SendReport(reportId uint8, data BufferSource) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallHIDDeviceSendReport(
		this.Ref(), js.Pointer(&_ok),
		uint32(reportId),
		data.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SendReportFunc returns the method "HIDDevice.sendReport".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HIDDevice) SendReportFunc() (fn js.Func[func(reportId uint8, data BufferSource) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.HIDDeviceSendReportFunc(
			this.Ref(),
		),
	)
}

// SendFeatureReport calls the method "HIDDevice.sendFeatureReport".
//
// The returned bool will be false if there is no such method.
func (this HIDDevice) SendFeatureReport(reportId uint8, data BufferSource) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallHIDDeviceSendFeatureReport(
		this.Ref(), js.Pointer(&_ok),
		uint32(reportId),
		data.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SendFeatureReportFunc returns the method "HIDDevice.sendFeatureReport".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HIDDevice) SendFeatureReportFunc() (fn js.Func[func(reportId uint8, data BufferSource) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.HIDDeviceSendFeatureReportFunc(
			this.Ref(),
		),
	)
}

// ReceiveFeatureReport calls the method "HIDDevice.receiveFeatureReport".
//
// The returned bool will be false if there is no such method.
func (this HIDDevice) ReceiveFeatureReport(reportId uint8) (js.Promise[js.DataView], bool) {
	var _ok bool
	_ret := bindings.CallHIDDeviceReceiveFeatureReport(
		this.Ref(), js.Pointer(&_ok),
		uint32(reportId),
	)

	return js.Promise[js.DataView]{}.FromRef(_ret), _ok
}

// ReceiveFeatureReportFunc returns the method "HIDDevice.receiveFeatureReport".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HIDDevice) ReceiveFeatureReportFunc() (fn js.Func[func(reportId uint8) js.Promise[js.DataView]]) {
	return fn.FromRef(
		bindings.HIDDeviceReceiveFeatureReportFunc(
			this.Ref(),
		),
	)
}

type HIDDeviceFilter struct {
	// VendorId is "HIDDeviceFilter.vendorId"
	//
	// Optional
	VendorId uint32
	// ProductId is "HIDDeviceFilter.productId"
	//
	// Optional
	ProductId uint16
	// UsagePage is "HIDDeviceFilter.usagePage"
	//
	// Optional
	UsagePage uint16
	// Usage is "HIDDeviceFilter.usage"
	//
	// Optional
	Usage uint16

	FFI_USE_VendorId  bool // for VendorId.
	FFI_USE_ProductId bool // for ProductId.
	FFI_USE_UsagePage bool // for UsagePage.
	FFI_USE_Usage     bool // for Usage.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HIDDeviceFilter with all fields set.
func (p HIDDeviceFilter) FromRef(ref js.Ref) HIDDeviceFilter {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 HIDDeviceFilter HIDDeviceFilter [// HIDDeviceFilter] [0x140006f55e0 0x140006f5720 0x140006f5860 0x140006f59a0 0x140006f5680 0x140006f57c0 0x140006f5900 0x140006f5a40] 0x14000575c08 {0 0}} in the application heap.
func (p HIDDeviceFilter) New() js.Ref {
	return bindings.HIDDeviceFilterJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p HIDDeviceFilter) UpdateFrom(ref js.Ref) {
	bindings.HIDDeviceFilterJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p HIDDeviceFilter) Update(ref js.Ref) {
	bindings.HIDDeviceFilterJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type HIDDeviceRequestOptions struct {
	// Filters is "HIDDeviceRequestOptions.filters"
	//
	// Required
	Filters js.Array[HIDDeviceFilter]
	// ExclusionFilters is "HIDDeviceRequestOptions.exclusionFilters"
	//
	// Optional
	ExclusionFilters js.Array[HIDDeviceFilter]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HIDDeviceRequestOptions with all fields set.
func (p HIDDeviceRequestOptions) FromRef(ref js.Ref) HIDDeviceRequestOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 HIDDeviceRequestOptions HIDDeviceRequestOptions [// HIDDeviceRequestOptions] [0x140006f5ae0 0x140006f5b80] 0x14000575bd8 {0 0}} in the application heap.
func (p HIDDeviceRequestOptions) New() js.Ref {
	return bindings.HIDDeviceRequestOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p HIDDeviceRequestOptions) UpdateFrom(ref js.Ref) {
	bindings.HIDDeviceRequestOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p HIDDeviceRequestOptions) Update(ref js.Ref) {
	bindings.HIDDeviceRequestOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type HID struct {
	EventTarget
}

func (this HID) Once() HID {
	this.Ref().Once()
	return this
}

func (this HID) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this HID) FromRef(ref js.Ref) HID {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this HID) Free() {
	this.Ref().Free()
}

// GetDevices calls the method "HID.getDevices".
//
// The returned bool will be false if there is no such method.
func (this HID) GetDevices() (js.Promise[js.Array[HIDDevice]], bool) {
	var _ok bool
	_ret := bindings.CallHIDGetDevices(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[HIDDevice]]{}.FromRef(_ret), _ok
}

// GetDevicesFunc returns the method "HID.getDevices".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HID) GetDevicesFunc() (fn js.Func[func() js.Promise[js.Array[HIDDevice]]]) {
	return fn.FromRef(
		bindings.HIDGetDevicesFunc(
			this.Ref(),
		),
	)
}

// RequestDevice calls the method "HID.requestDevice".
//
// The returned bool will be false if there is no such method.
func (this HID) RequestDevice(options HIDDeviceRequestOptions) (js.Promise[js.Array[HIDDevice]], bool) {
	var _ok bool
	_ret := bindings.CallHIDRequestDevice(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[js.Array[HIDDevice]]{}.FromRef(_ret), _ok
}

// RequestDeviceFunc returns the method "HID.requestDevice".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HID) RequestDeviceFunc() (fn js.Func[func(options HIDDeviceRequestOptions) js.Promise[js.Array[HIDDevice]]]) {
	return fn.FromRef(
		bindings.HIDRequestDeviceFunc(
			this.Ref(),
		),
	)
}

type WindowControlsOverlay struct {
	EventTarget
}

func (this WindowControlsOverlay) Once() WindowControlsOverlay {
	this.Ref().Once()
	return this
}

func (this WindowControlsOverlay) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this WindowControlsOverlay) FromRef(ref js.Ref) WindowControlsOverlay {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this WindowControlsOverlay) Free() {
	this.Ref().Free()
}

// Visible returns the value of property "WindowControlsOverlay.visible".
//
// The returned bool will be false if there is no such property.
func (this WindowControlsOverlay) Visible() (bool, bool) {
	var _ok bool
	_ret := bindings.GetWindowControlsOverlayVisible(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// GetTitlebarAreaRect calls the method "WindowControlsOverlay.getTitlebarAreaRect".
//
// The returned bool will be false if there is no such method.
func (this WindowControlsOverlay) GetTitlebarAreaRect() (DOMRect, bool) {
	var _ok bool
	_ret := bindings.CallWindowControlsOverlayGetTitlebarAreaRect(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMRect{}.FromRef(_ret), _ok
}

// GetTitlebarAreaRectFunc returns the method "WindowControlsOverlay.getTitlebarAreaRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WindowControlsOverlay) GetTitlebarAreaRectFunc() (fn js.Func[func() DOMRect]) {
	return fn.FromRef(
		bindings.WindowControlsOverlayGetTitlebarAreaRectFunc(
			this.Ref(),
		),
	)
}

type Credential struct {
	ref js.Ref
}

func (this Credential) Once() Credential {
	this.Ref().Once()
	return this
}

func (this Credential) Ref() js.Ref {
	return this.ref
}

func (this Credential) FromRef(ref js.Ref) Credential {
	this.ref = ref
	return this
}

func (this Credential) Free() {
	this.Ref().Free()
}

// Id returns the value of property "Credential.id".
//
// The returned bool will be false if there is no such property.
func (this Credential) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCredentialId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Type returns the value of property "Credential.type".
//
// The returned bool will be false if there is no such property.
func (this Credential) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCredentialType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// IsConditionalMediationAvailable calls the staticmethod "Credential.isConditionalMediationAvailable".
//
// The returned bool will be false if there is no such method.
func (this Credential) IsConditionalMediationAvailable() (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallCredentialIsConditionalMediationAvailable(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// IsConditionalMediationAvailableFunc returns the staticmethod "Credential.isConditionalMediationAvailable".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Credential) IsConditionalMediationAvailableFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.CredentialIsConditionalMediationAvailableFunc(
			this.Ref(),
		),
	)
}

type CredentialMediationRequirement uint32

const (
	_ CredentialMediationRequirement = iota

	CredentialMediationRequirement_SILENT
	CredentialMediationRequirement_OPTIONAL
	CredentialMediationRequirement_CONDITIONAL
	CredentialMediationRequirement_REQUIRED
)

func (CredentialMediationRequirement) FromRef(str js.Ref) CredentialMediationRequirement {
	return CredentialMediationRequirement(bindings.ConstOfCredentialMediationRequirement(str))
}

func (x CredentialMediationRequirement) String() (string, bool) {
	switch x {
	case CredentialMediationRequirement_SILENT:
		return "silent", true
	case CredentialMediationRequirement_OPTIONAL:
		return "optional", true
	case CredentialMediationRequirement_CONDITIONAL:
		return "conditional", true
	case CredentialMediationRequirement_REQUIRED:
		return "required", true
	default:
		return "", false
	}
}

type OTPCredentialTransportType uint32

const (
	_ OTPCredentialTransportType = iota

	OTPCredentialTransportType_SMS
)

func (OTPCredentialTransportType) FromRef(str js.Ref) OTPCredentialTransportType {
	return OTPCredentialTransportType(bindings.ConstOfOTPCredentialTransportType(str))
}

func (x OTPCredentialTransportType) String() (string, bool) {
	switch x {
	case OTPCredentialTransportType_SMS:
		return "sms", true
	default:
		return "", false
	}
}

type OTPCredentialRequestOptions struct {
	// Transport is "OTPCredentialRequestOptions.transport"
	//
	// Optional, defaults to [].
	Transport js.Array[OTPCredentialTransportType]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OTPCredentialRequestOptions with all fields set.
func (p OTPCredentialRequestOptions) FromRef(ref js.Ref) OTPCredentialRequestOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 OTPCredentialRequestOptions OTPCredentialRequestOptions [// OTPCredentialRequestOptions] [0x140006f5ea0] 0x14000575cb0 {0 0}} in the application heap.
func (p OTPCredentialRequestOptions) New() js.Ref {
	return bindings.OTPCredentialRequestOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p OTPCredentialRequestOptions) UpdateFrom(ref js.Ref) {
	bindings.OTPCredentialRequestOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p OTPCredentialRequestOptions) Update(ref js.Ref) {
	bindings.OTPCredentialRequestOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PublicKeyCredentialDescriptor struct {
	// Type is "PublicKeyCredentialDescriptor.type"
	//
	// Required
	Type js.String
	// Id is "PublicKeyCredentialDescriptor.id"
	//
	// Required
	Id BufferSource
	// Transports is "PublicKeyCredentialDescriptor.transports"
	//
	// Optional
	Transports js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PublicKeyCredentialDescriptor with all fields set.
func (p PublicKeyCredentialDescriptor) FromRef(ref js.Ref) PublicKeyCredentialDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PublicKeyCredentialDescriptor PublicKeyCredentialDescriptor [// PublicKeyCredentialDescriptor] [0x140006f8280 0x140006f8320 0x140006f83c0] 0x14000575cf8 {0 0}} in the application heap.
func (p PublicKeyCredentialDescriptor) New() js.Ref {
	return bindings.PublicKeyCredentialDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PublicKeyCredentialDescriptor) UpdateFrom(ref js.Ref) {
	bindings.PublicKeyCredentialDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PublicKeyCredentialDescriptor) Update(ref js.Ref) {
	bindings.PublicKeyCredentialDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PublicKeyCredentialRequestOptions struct {
	// Challenge is "PublicKeyCredentialRequestOptions.challenge"
	//
	// Required
	Challenge BufferSource
	// Timeout is "PublicKeyCredentialRequestOptions.timeout"
	//
	// Optional
	Timeout uint32
	// RpId is "PublicKeyCredentialRequestOptions.rpId"
	//
	// Optional
	RpId js.String
	// AllowCredentials is "PublicKeyCredentialRequestOptions.allowCredentials"
	//
	// Optional, defaults to [].
	AllowCredentials js.Array[PublicKeyCredentialDescriptor]
	// UserVerification is "PublicKeyCredentialRequestOptions.userVerification"
	//
	// Optional, defaults to "preferred".
	UserVerification js.String
	// Hints is "PublicKeyCredentialRequestOptions.hints"
	//
	// Optional, defaults to [].
	Hints js.Array[js.String]
	// Attestation is "PublicKeyCredentialRequestOptions.attestation"
	//
	// Optional, defaults to "none".
	Attestation js.String
	// AttestationFormats is "PublicKeyCredentialRequestOptions.attestationFormats"
	//
	// Optional, defaults to [].
	AttestationFormats js.Array[js.String]
	// Extensions is "PublicKeyCredentialRequestOptions.extensions"
	//
	// Optional
	Extensions AuthenticationExtensionsClientInputs

	FFI_USE_Timeout bool // for Timeout.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PublicKeyCredentialRequestOptions with all fields set.
func (p PublicKeyCredentialRequestOptions) FromRef(ref js.Ref) PublicKeyCredentialRequestOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PublicKeyCredentialRequestOptions PublicKeyCredentialRequestOptions [// PublicKeyCredentialRequestOptions] [0x140006f8000 0x140006f80a0 0x140006f81e0 0x140006f8460 0x140006f8500 0x140006f85a0 0x140006f8640 0x140006f86e0 0x140006f8780 0x140006f8140] 0x14000575cc8 {0 0}} in the application heap.
func (p PublicKeyCredentialRequestOptions) New() js.Ref {
	return bindings.PublicKeyCredentialRequestOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PublicKeyCredentialRequestOptions) UpdateFrom(ref js.Ref) {
	bindings.PublicKeyCredentialRequestOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PublicKeyCredentialRequestOptions) Update(ref js.Ref) {
	bindings.PublicKeyCredentialRequestOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type IdentityProviderConfig struct {
	// ConfigURL is "IdentityProviderConfig.configURL"
	//
	// Required
	ConfigURL js.String
	// ClientId is "IdentityProviderConfig.clientId"
	//
	// Required
	ClientId js.String
	// Nonce is "IdentityProviderConfig.nonce"
	//
	// Optional
	Nonce js.String
	// LoginHint is "IdentityProviderConfig.loginHint"
	//
	// Optional
	LoginHint js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IdentityProviderConfig with all fields set.
func (p IdentityProviderConfig) FromRef(ref js.Ref) IdentityProviderConfig {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 IdentityProviderConfig IdentityProviderConfig [// IdentityProviderConfig] [0x140006f8a00 0x140006f8aa0 0x140006f8b40 0x140006f8be0] 0x14000575d58 {0 0}} in the application heap.
func (p IdentityProviderConfig) New() js.Ref {
	return bindings.IdentityProviderConfigJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p IdentityProviderConfig) UpdateFrom(ref js.Ref) {
	bindings.IdentityProviderConfigJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p IdentityProviderConfig) Update(ref js.Ref) {
	bindings.IdentityProviderConfigJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type IdentityCredentialRequestOptionsContext uint32

const (
	_ IdentityCredentialRequestOptionsContext = iota

	IdentityCredentialRequestOptionsContext_SIGNIN
	IdentityCredentialRequestOptionsContext_SIGNUP
	IdentityCredentialRequestOptionsContext_USE
	IdentityCredentialRequestOptionsContext_CONTINUE
)

func (IdentityCredentialRequestOptionsContext) FromRef(str js.Ref) IdentityCredentialRequestOptionsContext {
	return IdentityCredentialRequestOptionsContext(bindings.ConstOfIdentityCredentialRequestOptionsContext(str))
}

func (x IdentityCredentialRequestOptionsContext) String() (string, bool) {
	switch x {
	case IdentityCredentialRequestOptionsContext_SIGNIN:
		return "signin", true
	case IdentityCredentialRequestOptionsContext_SIGNUP:
		return "signup", true
	case IdentityCredentialRequestOptionsContext_USE:
		return "use", true
	case IdentityCredentialRequestOptionsContext_CONTINUE:
		return "continue", true
	default:
		return "", false
	}
}

type IdentityCredentialRequestOptions struct {
	// Providers is "IdentityCredentialRequestOptions.providers"
	//
	// Required
	Providers js.Array[IdentityProviderConfig]
	// Context is "IdentityCredentialRequestOptions.context"
	//
	// Optional, defaults to "signin".
	Context IdentityCredentialRequestOptionsContext

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IdentityCredentialRequestOptions with all fields set.
func (p IdentityCredentialRequestOptions) FromRef(ref js.Ref) IdentityCredentialRequestOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 IdentityCredentialRequestOptions IdentityCredentialRequestOptions [// IdentityCredentialRequestOptions] [0x140006f8c80 0x140006f8d20] 0x14000575d40 {0 0}} in the application heap.
func (p IdentityCredentialRequestOptions) New() js.Ref {
	return bindings.IdentityCredentialRequestOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p IdentityCredentialRequestOptions) UpdateFrom(ref js.Ref) {
	bindings.IdentityCredentialRequestOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p IdentityCredentialRequestOptions) Update(ref js.Ref) {
	bindings.IdentityCredentialRequestOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type FederatedCredentialRequestOptions struct {
	// Providers is "FederatedCredentialRequestOptions.providers"
	//
	// Optional
	Providers js.Array[js.String]
	// Protocols is "FederatedCredentialRequestOptions.protocols"
	//
	// Optional
	Protocols js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FederatedCredentialRequestOptions with all fields set.
func (p FederatedCredentialRequestOptions) FromRef(ref js.Ref) FederatedCredentialRequestOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 FederatedCredentialRequestOptions FederatedCredentialRequestOptions [// FederatedCredentialRequestOptions] [0x140006f8e60 0x140006f8f00] 0x14000575d70 {0 0}} in the application heap.
func (p FederatedCredentialRequestOptions) New() js.Ref {
	return bindings.FederatedCredentialRequestOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p FederatedCredentialRequestOptions) UpdateFrom(ref js.Ref) {
	bindings.FederatedCredentialRequestOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FederatedCredentialRequestOptions) Update(ref js.Ref) {
	bindings.FederatedCredentialRequestOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type CredentialRequestOptions struct {
	// Mediation is "CredentialRequestOptions.mediation"
	//
	// Optional, defaults to "optional".
	Mediation CredentialMediationRequirement
	// Signal is "CredentialRequestOptions.signal"
	//
	// Optional
	Signal AbortSignal
	// Otp is "CredentialRequestOptions.otp"
	//
	// Optional
	Otp OTPCredentialRequestOptions
	// PublicKey is "CredentialRequestOptions.publicKey"
	//
	// Optional
	PublicKey PublicKeyCredentialRequestOptions
	// Password is "CredentialRequestOptions.password"
	//
	// Optional, defaults to false.
	Password bool
	// Identity is "CredentialRequestOptions.identity"
	//
	// Optional
	Identity IdentityCredentialRequestOptions
	// Federated is "CredentialRequestOptions.federated"
	//
	// Optional
	Federated FederatedCredentialRequestOptions

	FFI_USE_Password bool // for Password.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CredentialRequestOptions with all fields set.
func (p CredentialRequestOptions) FromRef(ref js.Ref) CredentialRequestOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 CredentialRequestOptions CredentialRequestOptions [// CredentialRequestOptions] [0x140006f5d60 0x140006f5e00 0x140006f5f40 0x140006f8820 0x140006f88c0 0x140006f8dc0 0x140006f8fa0 0x140006f8960] 0x14000575c80 {0 0}} in the application heap.
func (p CredentialRequestOptions) New() js.Ref {
	return bindings.CredentialRequestOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CredentialRequestOptions) UpdateFrom(ref js.Ref) {
	bindings.CredentialRequestOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CredentialRequestOptions) Update(ref js.Ref) {
	bindings.CredentialRequestOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PublicKeyCredentialRpEntity struct {
	// Id is "PublicKeyCredentialRpEntity.id"
	//
	// Optional
	Id js.String
	// Name is "PublicKeyCredentialRpEntity.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PublicKeyCredentialRpEntity with all fields set.
func (p PublicKeyCredentialRpEntity) FromRef(ref js.Ref) PublicKeyCredentialRpEntity {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PublicKeyCredentialRpEntity PublicKeyCredentialRpEntity [// PublicKeyCredentialRpEntity] [0x140006f90e0 0x140006f9180] 0x14000575de8 {0 0}} in the application heap.
func (p PublicKeyCredentialRpEntity) New() js.Ref {
	return bindings.PublicKeyCredentialRpEntityJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PublicKeyCredentialRpEntity) UpdateFrom(ref js.Ref) {
	bindings.PublicKeyCredentialRpEntityJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PublicKeyCredentialRpEntity) Update(ref js.Ref) {
	bindings.PublicKeyCredentialRpEntityJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PublicKeyCredentialUserEntity struct {
	// Id is "PublicKeyCredentialUserEntity.id"
	//
	// Required
	Id BufferSource
	// DisplayName is "PublicKeyCredentialUserEntity.displayName"
	//
	// Required
	DisplayName js.String
	// Name is "PublicKeyCredentialUserEntity.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PublicKeyCredentialUserEntity with all fields set.
func (p PublicKeyCredentialUserEntity) FromRef(ref js.Ref) PublicKeyCredentialUserEntity {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PublicKeyCredentialUserEntity PublicKeyCredentialUserEntity [// PublicKeyCredentialUserEntity] [0x140006f92c0 0x140006f9360 0x140006f9400] 0x14000575e48 {0 0}} in the application heap.
func (p PublicKeyCredentialUserEntity) New() js.Ref {
	return bindings.PublicKeyCredentialUserEntityJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PublicKeyCredentialUserEntity) UpdateFrom(ref js.Ref) {
	bindings.PublicKeyCredentialUserEntityJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PublicKeyCredentialUserEntity) Update(ref js.Ref) {
	bindings.PublicKeyCredentialUserEntityJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PublicKeyCredentialParameters struct {
	// Type is "PublicKeyCredentialParameters.type"
	//
	// Required
	Type js.String
	// Alg is "PublicKeyCredentialParameters.alg"
	//
	// Required
	Alg COSEAlgorithmIdentifier

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PublicKeyCredentialParameters with all fields set.
func (p PublicKeyCredentialParameters) FromRef(ref js.Ref) PublicKeyCredentialParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PublicKeyCredentialParameters PublicKeyCredentialParameters [// PublicKeyCredentialParameters] [0x140006f95e0 0x140006f9680] 0x14000575e90 {0 0}} in the application heap.
func (p PublicKeyCredentialParameters) New() js.Ref {
	return bindings.PublicKeyCredentialParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PublicKeyCredentialParameters) UpdateFrom(ref js.Ref) {
	bindings.PublicKeyCredentialParametersJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PublicKeyCredentialParameters) Update(ref js.Ref) {
	bindings.PublicKeyCredentialParametersJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PublicKeyCredentialCreationOptions struct {
	// Rp is "PublicKeyCredentialCreationOptions.rp"
	//
	// Required
	Rp PublicKeyCredentialRpEntity
	// User is "PublicKeyCredentialCreationOptions.user"
	//
	// Required
	User PublicKeyCredentialUserEntity
	// Challenge is "PublicKeyCredentialCreationOptions.challenge"
	//
	// Required
	Challenge BufferSource
	// PubKeyCredParams is "PublicKeyCredentialCreationOptions.pubKeyCredParams"
	//
	// Required
	PubKeyCredParams js.Array[PublicKeyCredentialParameters]
	// Timeout is "PublicKeyCredentialCreationOptions.timeout"
	//
	// Optional
	Timeout uint32
	// ExcludeCredentials is "PublicKeyCredentialCreationOptions.excludeCredentials"
	//
	// Optional, defaults to [].
	ExcludeCredentials js.Array[PublicKeyCredentialDescriptor]
	// AuthenticatorSelection is "PublicKeyCredentialCreationOptions.authenticatorSelection"
	//
	// Optional
	AuthenticatorSelection AuthenticatorSelectionCriteria
	// Hints is "PublicKeyCredentialCreationOptions.hints"
	//
	// Optional, defaults to [].
	Hints js.Array[js.String]
	// Attestation is "PublicKeyCredentialCreationOptions.attestation"
	//
	// Optional, defaults to "none".
	Attestation js.String
	// AttestationFormats is "PublicKeyCredentialCreationOptions.attestationFormats"
	//
	// Optional, defaults to [].
	AttestationFormats js.Array[js.String]
	// Extensions is "PublicKeyCredentialCreationOptions.extensions"
	//
	// Optional
	Extensions AuthenticationExtensionsClientInputs

	FFI_USE_Timeout bool // for Timeout.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PublicKeyCredentialCreationOptions with all fields set.
func (p PublicKeyCredentialCreationOptions) FromRef(ref js.Ref) PublicKeyCredentialCreationOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PublicKeyCredentialCreationOptions PublicKeyCredentialCreationOptions [// PublicKeyCredentialCreationOptions] [0x140006f9220 0x140006f94a0 0x140006f9540 0x140006f9720 0x140006f97c0 0x140006f9900 0x140006f99a0 0x140006f9a40 0x140006f9ae0 0x140006f9b80 0x140006f9c20 0x140006f9860] 0x14000575db8 {0 0}} in the application heap.
func (p PublicKeyCredentialCreationOptions) New() js.Ref {
	return bindings.PublicKeyCredentialCreationOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PublicKeyCredentialCreationOptions) UpdateFrom(ref js.Ref) {
	bindings.PublicKeyCredentialCreationOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PublicKeyCredentialCreationOptions) Update(ref js.Ref) {
	bindings.PublicKeyCredentialCreationOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PasswordCredentialData struct {
	// Name is "PasswordCredentialData.name"
	//
	// Optional
	Name js.String
	// IconURL is "PasswordCredentialData.iconURL"
	//
	// Optional
	IconURL js.String
	// Origin is "PasswordCredentialData.origin"
	//
	// Required
	Origin js.String
	// Password is "PasswordCredentialData.password"
	//
	// Required
	Password js.String
	// Id is "PasswordCredentialData.id"
	//
	// Required
	Id js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PasswordCredentialData with all fields set.
func (p PasswordCredentialData) FromRef(ref js.Ref) PasswordCredentialData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PasswordCredentialData PasswordCredentialData [// PasswordCredentialData] [0x140006f9d60 0x140006f9e00 0x140006f9ea0 0x140006f9f40 0x140006fe000] 0x14000575ed8 {0 0}} in the application heap.
func (p PasswordCredentialData) New() js.Ref {
	return bindings.PasswordCredentialDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PasswordCredentialData) UpdateFrom(ref js.Ref) {
	bindings.PasswordCredentialDataJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PasswordCredentialData) Update(ref js.Ref) {
	bindings.PasswordCredentialDataJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_PasswordCredentialData_HTMLFormElement struct {
	ref js.Ref
}

func (x OneOf_PasswordCredentialData_HTMLFormElement) Ref() js.Ref {
	return x.ref
}

func (x OneOf_PasswordCredentialData_HTMLFormElement) Free() {
	x.ref.Free()
}

func (x OneOf_PasswordCredentialData_HTMLFormElement) FromRef(ref js.Ref) OneOf_PasswordCredentialData_HTMLFormElement {
	return OneOf_PasswordCredentialData_HTMLFormElement{
		ref: ref,
	}
}

func (x OneOf_PasswordCredentialData_HTMLFormElement) PasswordCredentialData() PasswordCredentialData {
	var ret PasswordCredentialData
	ret.UpdateFrom(x.ref)
	return ret
}

func (x OneOf_PasswordCredentialData_HTMLFormElement) HTMLFormElement() HTMLFormElement {
	return HTMLFormElement{}.FromRef(x.ref)
}

type PasswordCredentialInit = OneOf_PasswordCredentialData_HTMLFormElement

type FederatedCredentialInit struct {
	// Name is "FederatedCredentialInit.name"
	//
	// Optional
	Name js.String
	// IconURL is "FederatedCredentialInit.iconURL"
	//
	// Optional
	IconURL js.String
	// Origin is "FederatedCredentialInit.origin"
	//
	// Required
	Origin js.String
	// Provider is "FederatedCredentialInit.provider"
	//
	// Required
	Provider js.String
	// Protocol is "FederatedCredentialInit.protocol"
	//
	// Optional
	Protocol js.String
	// Id is "FederatedCredentialInit.id"
	//
	// Required
	Id js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FederatedCredentialInit with all fields set.
func (p FederatedCredentialInit) FromRef(ref js.Ref) FederatedCredentialInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 FederatedCredentialInit FederatedCredentialInit [// FederatedCredentialInit] [0x140006fe140 0x140006fe1e0 0x140006fe280 0x140006fe320 0x140006fe3c0 0x140006fe460] 0x14000575f20 {0 0}} in the application heap.
func (p FederatedCredentialInit) New() js.Ref {
	return bindings.FederatedCredentialInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p FederatedCredentialInit) UpdateFrom(ref js.Ref) {
	bindings.FederatedCredentialInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FederatedCredentialInit) Update(ref js.Ref) {
	bindings.FederatedCredentialInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type CredentialCreationOptions struct {
	// Signal is "CredentialCreationOptions.signal"
	//
	// Optional
	Signal AbortSignal
	// PublicKey is "CredentialCreationOptions.publicKey"
	//
	// Optional
	PublicKey PublicKeyCredentialCreationOptions
	// Password is "CredentialCreationOptions.password"
	//
	// Optional
	Password PasswordCredentialInit
	// Federated is "CredentialCreationOptions.federated"
	//
	// Optional
	Federated FederatedCredentialInit

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CredentialCreationOptions with all fields set.
func (p CredentialCreationOptions) FromRef(ref js.Ref) CredentialCreationOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 CredentialCreationOptions CredentialCreationOptions [// CredentialCreationOptions] [0x140006f9040 0x140006f9cc0 0x140006fe0a0 0x140006fe500] 0x14000575da0 {0 0}} in the application heap.
func (p CredentialCreationOptions) New() js.Ref {
	return bindings.CredentialCreationOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CredentialCreationOptions) UpdateFrom(ref js.Ref) {
	bindings.CredentialCreationOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CredentialCreationOptions) Update(ref js.Ref) {
	bindings.CredentialCreationOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type CredentialsContainer struct {
	ref js.Ref
}

func (this CredentialsContainer) Once() CredentialsContainer {
	this.Ref().Once()
	return this
}

func (this CredentialsContainer) Ref() js.Ref {
	return this.ref
}

func (this CredentialsContainer) FromRef(ref js.Ref) CredentialsContainer {
	this.ref = ref
	return this
}

func (this CredentialsContainer) Free() {
	this.Ref().Free()
}

// Get calls the method "CredentialsContainer.get".
//
// The returned bool will be false if there is no such method.
func (this CredentialsContainer) Get(options CredentialRequestOptions) (js.Promise[Credential], bool) {
	var _ok bool
	_ret := bindings.CallCredentialsContainerGet(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[Credential]{}.FromRef(_ret), _ok
}

// GetFunc returns the method "CredentialsContainer.get".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CredentialsContainer) GetFunc() (fn js.Func[func(options CredentialRequestOptions) js.Promise[Credential]]) {
	return fn.FromRef(
		bindings.CredentialsContainerGetFunc(
			this.Ref(),
		),
	)
}

// Get1 calls the method "CredentialsContainer.get".
//
// The returned bool will be false if there is no such method.
func (this CredentialsContainer) Get1() (js.Promise[Credential], bool) {
	var _ok bool
	_ret := bindings.CallCredentialsContainerGet1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[Credential]{}.FromRef(_ret), _ok
}

// Get1Func returns the method "CredentialsContainer.get".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CredentialsContainer) Get1Func() (fn js.Func[func() js.Promise[Credential]]) {
	return fn.FromRef(
		bindings.CredentialsContainerGet1Func(
			this.Ref(),
		),
	)
}

// Store calls the method "CredentialsContainer.store".
//
// The returned bool will be false if there is no such method.
func (this CredentialsContainer) Store(credential Credential) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallCredentialsContainerStore(
		this.Ref(), js.Pointer(&_ok),
		credential.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// StoreFunc returns the method "CredentialsContainer.store".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CredentialsContainer) StoreFunc() (fn js.Func[func(credential Credential) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.CredentialsContainerStoreFunc(
			this.Ref(),
		),
	)
}

// Create calls the method "CredentialsContainer.create".
//
// The returned bool will be false if there is no such method.
func (this CredentialsContainer) Create(options CredentialCreationOptions) (js.Promise[Credential], bool) {
	var _ok bool
	_ret := bindings.CallCredentialsContainerCreate(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[Credential]{}.FromRef(_ret), _ok
}

// CreateFunc returns the method "CredentialsContainer.create".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CredentialsContainer) CreateFunc() (fn js.Func[func(options CredentialCreationOptions) js.Promise[Credential]]) {
	return fn.FromRef(
		bindings.CredentialsContainerCreateFunc(
			this.Ref(),
		),
	)
}

// Create1 calls the method "CredentialsContainer.create".
//
// The returned bool will be false if there is no such method.
func (this CredentialsContainer) Create1() (js.Promise[Credential], bool) {
	var _ok bool
	_ret := bindings.CallCredentialsContainerCreate1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[Credential]{}.FromRef(_ret), _ok
}

// Create1Func returns the method "CredentialsContainer.create".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CredentialsContainer) Create1Func() (fn js.Func[func() js.Promise[Credential]]) {
	return fn.FromRef(
		bindings.CredentialsContainerCreate1Func(
			this.Ref(),
		),
	)
}

// PreventSilentAccess calls the method "CredentialsContainer.preventSilentAccess".
//
// The returned bool will be false if there is no such method.
func (this CredentialsContainer) PreventSilentAccess() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallCredentialsContainerPreventSilentAccess(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// PreventSilentAccessFunc returns the method "CredentialsContainer.preventSilentAccess".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CredentialsContainer) PreventSilentAccessFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.CredentialsContainerPreventSilentAccessFunc(
			this.Ref(),
		),
	)
}

type PositionCallbackFunc func(this js.Ref, position GeolocationPosition) js.Ref

func (fn PositionCallbackFunc) Register() js.Func[func(position GeolocationPosition)] {
	return js.RegisterCallback[func(position GeolocationPosition)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn PositionCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		GeolocationPosition{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PositionCallback[T any] struct {
	Fn  func(arg T, this js.Ref, position GeolocationPosition) js.Ref
	Arg T
}

func (cb *PositionCallback[T]) Register() js.Func[func(position GeolocationPosition)] {
	return js.RegisterCallback[func(position GeolocationPosition)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *PositionCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		GeolocationPosition{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GeolocationCoordinates struct {
	ref js.Ref
}

func (this GeolocationCoordinates) Once() GeolocationCoordinates {
	this.Ref().Once()
	return this
}

func (this GeolocationCoordinates) Ref() js.Ref {
	return this.ref
}

func (this GeolocationCoordinates) FromRef(ref js.Ref) GeolocationCoordinates {
	this.ref = ref
	return this
}

func (this GeolocationCoordinates) Free() {
	this.Ref().Free()
}

// Accuracy returns the value of property "GeolocationCoordinates.accuracy".
//
// The returned bool will be false if there is no such property.
func (this GeolocationCoordinates) Accuracy() (float64, bool) {
	var _ok bool
	_ret := bindings.GetGeolocationCoordinatesAccuracy(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Latitude returns the value of property "GeolocationCoordinates.latitude".
//
// The returned bool will be false if there is no such property.
func (this GeolocationCoordinates) Latitude() (float64, bool) {
	var _ok bool
	_ret := bindings.GetGeolocationCoordinatesLatitude(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Longitude returns the value of property "GeolocationCoordinates.longitude".
//
// The returned bool will be false if there is no such property.
func (this GeolocationCoordinates) Longitude() (float64, bool) {
	var _ok bool
	_ret := bindings.GetGeolocationCoordinatesLongitude(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Altitude returns the value of property "GeolocationCoordinates.altitude".
//
// The returned bool will be false if there is no such property.
func (this GeolocationCoordinates) Altitude() (float64, bool) {
	var _ok bool
	_ret := bindings.GetGeolocationCoordinatesAltitude(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// AltitudeAccuracy returns the value of property "GeolocationCoordinates.altitudeAccuracy".
//
// The returned bool will be false if there is no such property.
func (this GeolocationCoordinates) AltitudeAccuracy() (float64, bool) {
	var _ok bool
	_ret := bindings.GetGeolocationCoordinatesAltitudeAccuracy(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Heading returns the value of property "GeolocationCoordinates.heading".
//
// The returned bool will be false if there is no such property.
func (this GeolocationCoordinates) Heading() (float64, bool) {
	var _ok bool
	_ret := bindings.GetGeolocationCoordinatesHeading(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Speed returns the value of property "GeolocationCoordinates.speed".
//
// The returned bool will be false if there is no such property.
func (this GeolocationCoordinates) Speed() (float64, bool) {
	var _ok bool
	_ret := bindings.GetGeolocationCoordinatesSpeed(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

type EpochTimeStamp uint64

type GeolocationPosition struct {
	ref js.Ref
}

func (this GeolocationPosition) Once() GeolocationPosition {
	this.Ref().Once()
	return this
}

func (this GeolocationPosition) Ref() js.Ref {
	return this.ref
}

func (this GeolocationPosition) FromRef(ref js.Ref) GeolocationPosition {
	this.ref = ref
	return this
}

func (this GeolocationPosition) Free() {
	this.Ref().Free()
}

// Coords returns the value of property "GeolocationPosition.coords".
//
// The returned bool will be false if there is no such property.
func (this GeolocationPosition) Coords() (GeolocationCoordinates, bool) {
	var _ok bool
	_ret := bindings.GetGeolocationPositionCoords(
		this.Ref(), js.Pointer(&_ok),
	)
	return GeolocationCoordinates{}.FromRef(_ret), _ok
}

// Timestamp returns the value of property "GeolocationPosition.timestamp".
//
// The returned bool will be false if there is no such property.
func (this GeolocationPosition) Timestamp() (EpochTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetGeolocationPositionTimestamp(
		this.Ref(), js.Pointer(&_ok),
	)
	return EpochTimeStamp(_ret), _ok
}

type PositionErrorCallbackFunc func(this js.Ref, positionError GeolocationPositionError) js.Ref

func (fn PositionErrorCallbackFunc) Register() js.Func[func(positionError GeolocationPositionError)] {
	return js.RegisterCallback[func(positionError GeolocationPositionError)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn PositionErrorCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		GeolocationPositionError{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PositionErrorCallback[T any] struct {
	Fn  func(arg T, this js.Ref, positionError GeolocationPositionError) js.Ref
	Arg T
}

func (cb *PositionErrorCallback[T]) Register() js.Func[func(positionError GeolocationPositionError)] {
	return js.RegisterCallback[func(positionError GeolocationPositionError)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *PositionErrorCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		GeolocationPositionError{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

const (
	GeolocationPositionError_PERMISSION_DENIED    uint16 = 1
	GeolocationPositionError_POSITION_UNAVAILABLE uint16 = 2
	GeolocationPositionError_TIMEOUT              uint16 = 3
)

type GeolocationPositionError struct {
	ref js.Ref
}

func (this GeolocationPositionError) Once() GeolocationPositionError {
	this.Ref().Once()
	return this
}

func (this GeolocationPositionError) Ref() js.Ref {
	return this.ref
}

func (this GeolocationPositionError) FromRef(ref js.Ref) GeolocationPositionError {
	this.ref = ref
	return this
}

func (this GeolocationPositionError) Free() {
	this.Ref().Free()
}

// Code returns the value of property "GeolocationPositionError.code".
//
// The returned bool will be false if there is no such property.
func (this GeolocationPositionError) Code() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetGeolocationPositionErrorCode(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// Message returns the value of property "GeolocationPositionError.message".
//
// The returned bool will be false if there is no such property.
func (this GeolocationPositionError) Message() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGeolocationPositionErrorMessage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type PositionOptions struct {
	// EnableHighAccuracy is "PositionOptions.enableHighAccuracy"
	//
	// Optional, defaults to false.
	EnableHighAccuracy bool
	// Timeout is "PositionOptions.timeout"
	//
	// Optional, defaults to 0xFFFFFFFF.
	Timeout uint32
	// MaximumAge is "PositionOptions.maximumAge"
	//
	// Optional, defaults to 0.
	MaximumAge uint32

	FFI_USE_EnableHighAccuracy bool // for EnableHighAccuracy.
	FFI_USE_Timeout            bool // for Timeout.
	FFI_USE_MaximumAge         bool // for MaximumAge.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PositionOptions with all fields set.
func (p PositionOptions) FromRef(ref js.Ref) PositionOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PositionOptions PositionOptions [// PositionOptions] [0x140006fe780 0x140006fe8c0 0x140006fea00 0x140006fe820 0x140006fe960 0x140006feaa0] 0x14000575fb0 {0 0}} in the application heap.
func (p PositionOptions) New() js.Ref {
	return bindings.PositionOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PositionOptions) UpdateFrom(ref js.Ref) {
	bindings.PositionOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PositionOptions) Update(ref js.Ref) {
	bindings.PositionOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type Geolocation struct {
	ref js.Ref
}

func (this Geolocation) Once() Geolocation {
	this.Ref().Once()
	return this
}

func (this Geolocation) Ref() js.Ref {
	return this.ref
}

func (this Geolocation) FromRef(ref js.Ref) Geolocation {
	this.ref = ref
	return this
}

func (this Geolocation) Free() {
	this.Ref().Free()
}

// GetCurrentPosition calls the method "Geolocation.getCurrentPosition".
//
// The returned bool will be false if there is no such method.
func (this Geolocation) GetCurrentPosition(successCallback js.Func[func(position GeolocationPosition)], errorCallback js.Func[func(positionError GeolocationPositionError)], options PositionOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGeolocationGetCurrentPosition(
		this.Ref(), js.Pointer(&_ok),
		successCallback.Ref(),
		errorCallback.Ref(),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetCurrentPositionFunc returns the method "Geolocation.getCurrentPosition".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Geolocation) GetCurrentPositionFunc() (fn js.Func[func(successCallback js.Func[func(position GeolocationPosition)], errorCallback js.Func[func(positionError GeolocationPositionError)], options PositionOptions)]) {
	return fn.FromRef(
		bindings.GeolocationGetCurrentPositionFunc(
			this.Ref(),
		),
	)
}

// GetCurrentPosition1 calls the method "Geolocation.getCurrentPosition".
//
// The returned bool will be false if there is no such method.
func (this Geolocation) GetCurrentPosition1(successCallback js.Func[func(position GeolocationPosition)], errorCallback js.Func[func(positionError GeolocationPositionError)]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGeolocationGetCurrentPosition1(
		this.Ref(), js.Pointer(&_ok),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetCurrentPosition1Func returns the method "Geolocation.getCurrentPosition".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Geolocation) GetCurrentPosition1Func() (fn js.Func[func(successCallback js.Func[func(position GeolocationPosition)], errorCallback js.Func[func(positionError GeolocationPositionError)])]) {
	return fn.FromRef(
		bindings.GeolocationGetCurrentPosition1Func(
			this.Ref(),
		),
	)
}

// GetCurrentPosition2 calls the method "Geolocation.getCurrentPosition".
//
// The returned bool will be false if there is no such method.
func (this Geolocation) GetCurrentPosition2(successCallback js.Func[func(position GeolocationPosition)]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGeolocationGetCurrentPosition2(
		this.Ref(), js.Pointer(&_ok),
		successCallback.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetCurrentPosition2Func returns the method "Geolocation.getCurrentPosition".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Geolocation) GetCurrentPosition2Func() (fn js.Func[func(successCallback js.Func[func(position GeolocationPosition)])]) {
	return fn.FromRef(
		bindings.GeolocationGetCurrentPosition2Func(
			this.Ref(),
		),
	)
}

// WatchPosition calls the method "Geolocation.watchPosition".
//
// The returned bool will be false if there is no such method.
func (this Geolocation) WatchPosition(successCallback js.Func[func(position GeolocationPosition)], errorCallback js.Func[func(positionError GeolocationPositionError)], options PositionOptions) (int32, bool) {
	var _ok bool
	_ret := bindings.CallGeolocationWatchPosition(
		this.Ref(), js.Pointer(&_ok),
		successCallback.Ref(),
		errorCallback.Ref(),
		js.Pointer(&options),
	)

	return int32(_ret), _ok
}

// WatchPositionFunc returns the method "Geolocation.watchPosition".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Geolocation) WatchPositionFunc() (fn js.Func[func(successCallback js.Func[func(position GeolocationPosition)], errorCallback js.Func[func(positionError GeolocationPositionError)], options PositionOptions) int32]) {
	return fn.FromRef(
		bindings.GeolocationWatchPositionFunc(
			this.Ref(),
		),
	)
}

// WatchPosition1 calls the method "Geolocation.watchPosition".
//
// The returned bool will be false if there is no such method.
func (this Geolocation) WatchPosition1(successCallback js.Func[func(position GeolocationPosition)], errorCallback js.Func[func(positionError GeolocationPositionError)]) (int32, bool) {
	var _ok bool
	_ret := bindings.CallGeolocationWatchPosition1(
		this.Ref(), js.Pointer(&_ok),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return int32(_ret), _ok
}

// WatchPosition1Func returns the method "Geolocation.watchPosition".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Geolocation) WatchPosition1Func() (fn js.Func[func(successCallback js.Func[func(position GeolocationPosition)], errorCallback js.Func[func(positionError GeolocationPositionError)]) int32]) {
	return fn.FromRef(
		bindings.GeolocationWatchPosition1Func(
			this.Ref(),
		),
	)
}

// WatchPosition2 calls the method "Geolocation.watchPosition".
//
// The returned bool will be false if there is no such method.
func (this Geolocation) WatchPosition2(successCallback js.Func[func(position GeolocationPosition)]) (int32, bool) {
	var _ok bool
	_ret := bindings.CallGeolocationWatchPosition2(
		this.Ref(), js.Pointer(&_ok),
		successCallback.Ref(),
	)

	return int32(_ret), _ok
}

// WatchPosition2Func returns the method "Geolocation.watchPosition".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Geolocation) WatchPosition2Func() (fn js.Func[func(successCallback js.Func[func(position GeolocationPosition)]) int32]) {
	return fn.FromRef(
		bindings.GeolocationWatchPosition2Func(
			this.Ref(),
		),
	)
}

// ClearWatch calls the method "Geolocation.clearWatch".
//
// The returned bool will be false if there is no such method.
func (this Geolocation) ClearWatch(watchId int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGeolocationClearWatch(
		this.Ref(), js.Pointer(&_ok),
		int32(watchId),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearWatchFunc returns the method "Geolocation.clearWatch".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Geolocation) ClearWatchFunc() (fn js.Func[func(watchId int32)]) {
	return fn.FromRef(
		bindings.GeolocationClearWatchFunc(
			this.Ref(),
		),
	)
}

type USBTransferStatus uint32

const (
	_ USBTransferStatus = iota

	USBTransferStatus_OK
	USBTransferStatus_STALL
	USBTransferStatus_BABBLE
)

func (USBTransferStatus) FromRef(str js.Ref) USBTransferStatus {
	return USBTransferStatus(bindings.ConstOfUSBTransferStatus(str))
}

func (x USBTransferStatus) String() (string, bool) {
	switch x {
	case USBTransferStatus_OK:
		return "ok", true
	case USBTransferStatus_STALL:
		return "stall", true
	case USBTransferStatus_BABBLE:
		return "babble", true
	default:
		return "", false
	}
}

func NewUSBInTransferResult(status USBTransferStatus, data js.DataView) USBInTransferResult {
	return USBInTransferResult{}.FromRef(
		bindings.NewUSBInTransferResultByUSBInTransferResult(
			uint32(status),
			data.Ref()),
	)
}

func NewUSBInTransferResultByUSBInTransferResult1(status USBTransferStatus) USBInTransferResult {
	return USBInTransferResult{}.FromRef(
		bindings.NewUSBInTransferResultByUSBInTransferResult1(
			uint32(status)),
	)
}

type USBInTransferResult struct {
	ref js.Ref
}

func (this USBInTransferResult) Once() USBInTransferResult {
	this.Ref().Once()
	return this
}

func (this USBInTransferResult) Ref() js.Ref {
	return this.ref
}

func (this USBInTransferResult) FromRef(ref js.Ref) USBInTransferResult {
	this.ref = ref
	return this
}

func (this USBInTransferResult) Free() {
	this.Ref().Free()
}

// Data returns the value of property "USBInTransferResult.data".
//
// The returned bool will be false if there is no such property.
func (this USBInTransferResult) Data() (js.DataView, bool) {
	var _ok bool
	_ret := bindings.GetUSBInTransferResultData(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.DataView{}.FromRef(_ret), _ok
}

// Status returns the value of property "USBInTransferResult.status".
//
// The returned bool will be false if there is no such property.
func (this USBInTransferResult) Status() (USBTransferStatus, bool) {
	var _ok bool
	_ret := bindings.GetUSBInTransferResultStatus(
		this.Ref(), js.Pointer(&_ok),
	)
	return USBTransferStatus(_ret), _ok
}

type USBRequestType uint32

const (
	_ USBRequestType = iota

	USBRequestType_STANDARD
	USBRequestType_CLASS
	USBRequestType_VENDOR
)

func (USBRequestType) FromRef(str js.Ref) USBRequestType {
	return USBRequestType(bindings.ConstOfUSBRequestType(str))
}

func (x USBRequestType) String() (string, bool) {
	switch x {
	case USBRequestType_STANDARD:
		return "standard", true
	case USBRequestType_CLASS:
		return "class", true
	case USBRequestType_VENDOR:
		return "vendor", true
	default:
		return "", false
	}
}

type USBRecipient uint32
