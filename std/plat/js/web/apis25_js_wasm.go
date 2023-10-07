// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

type MIDIOptions struct {
	// Sysex is "MIDIOptions.sysex"
	//
	// Optional
	//
	// NOTE: FFI_USE_Sysex MUST be set to true to make this field effective.
	Sysex bool
	// Software is "MIDIOptions.software"
	//
	// Optional
	//
	// NOTE: FFI_USE_Software MUST be set to true to make this field effective.
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

// New creates a new MIDIOptions in the application heap.
func (p MIDIOptions) New() js.Ref {
	return bindings.MIDIOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MIDIOptions) UpdateFrom(ref js.Ref) {
	bindings.MIDIOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MIDIOptions) Update(ref js.Ref) {
	bindings.MIDIOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MIDIOptions) FreeMembers(recursive bool) {
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
	//
	// NOTE: FFI_USE_IsAbsolute MUST be set to true to make this field effective.
	IsAbsolute bool
	// IsArray is "HIDReportItem.isArray"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsArray MUST be set to true to make this field effective.
	IsArray bool
	// IsBufferedBytes is "HIDReportItem.isBufferedBytes"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsBufferedBytes MUST be set to true to make this field effective.
	IsBufferedBytes bool
	// IsConstant is "HIDReportItem.isConstant"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsConstant MUST be set to true to make this field effective.
	IsConstant bool
	// IsLinear is "HIDReportItem.isLinear"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsLinear MUST be set to true to make this field effective.
	IsLinear bool
	// IsRange is "HIDReportItem.isRange"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsRange MUST be set to true to make this field effective.
	IsRange bool
	// IsVolatile is "HIDReportItem.isVolatile"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsVolatile MUST be set to true to make this field effective.
	IsVolatile bool
	// HasNull is "HIDReportItem.hasNull"
	//
	// Optional
	//
	// NOTE: FFI_USE_HasNull MUST be set to true to make this field effective.
	HasNull bool
	// HasPreferredState is "HIDReportItem.hasPreferredState"
	//
	// Optional
	//
	// NOTE: FFI_USE_HasPreferredState MUST be set to true to make this field effective.
	HasPreferredState bool
	// Wrap is "HIDReportItem.wrap"
	//
	// Optional
	//
	// NOTE: FFI_USE_Wrap MUST be set to true to make this field effective.
	Wrap bool
	// Usages is "HIDReportItem.usages"
	//
	// Optional
	Usages js.Array[uint32]
	// UsageMinimum is "HIDReportItem.usageMinimum"
	//
	// Optional
	//
	// NOTE: FFI_USE_UsageMinimum MUST be set to true to make this field effective.
	UsageMinimum uint32
	// UsageMaximum is "HIDReportItem.usageMaximum"
	//
	// Optional
	//
	// NOTE: FFI_USE_UsageMaximum MUST be set to true to make this field effective.
	UsageMaximum uint32
	// ReportSize is "HIDReportItem.reportSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_ReportSize MUST be set to true to make this field effective.
	ReportSize uint16
	// ReportCount is "HIDReportItem.reportCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_ReportCount MUST be set to true to make this field effective.
	ReportCount uint16
	// UnitExponent is "HIDReportItem.unitExponent"
	//
	// Optional
	//
	// NOTE: FFI_USE_UnitExponent MUST be set to true to make this field effective.
	UnitExponent int8
	// UnitSystem is "HIDReportItem.unitSystem"
	//
	// Optional
	UnitSystem HIDUnitSystem
	// UnitFactorLengthExponent is "HIDReportItem.unitFactorLengthExponent"
	//
	// Optional
	//
	// NOTE: FFI_USE_UnitFactorLengthExponent MUST be set to true to make this field effective.
	UnitFactorLengthExponent int8
	// UnitFactorMassExponent is "HIDReportItem.unitFactorMassExponent"
	//
	// Optional
	//
	// NOTE: FFI_USE_UnitFactorMassExponent MUST be set to true to make this field effective.
	UnitFactorMassExponent int8
	// UnitFactorTimeExponent is "HIDReportItem.unitFactorTimeExponent"
	//
	// Optional
	//
	// NOTE: FFI_USE_UnitFactorTimeExponent MUST be set to true to make this field effective.
	UnitFactorTimeExponent int8
	// UnitFactorTemperatureExponent is "HIDReportItem.unitFactorTemperatureExponent"
	//
	// Optional
	//
	// NOTE: FFI_USE_UnitFactorTemperatureExponent MUST be set to true to make this field effective.
	UnitFactorTemperatureExponent int8
	// UnitFactorCurrentExponent is "HIDReportItem.unitFactorCurrentExponent"
	//
	// Optional
	//
	// NOTE: FFI_USE_UnitFactorCurrentExponent MUST be set to true to make this field effective.
	UnitFactorCurrentExponent int8
	// UnitFactorLuminousIntensityExponent is "HIDReportItem.unitFactorLuminousIntensityExponent"
	//
	// Optional
	//
	// NOTE: FFI_USE_UnitFactorLuminousIntensityExponent MUST be set to true to make this field effective.
	UnitFactorLuminousIntensityExponent int8
	// LogicalMinimum is "HIDReportItem.logicalMinimum"
	//
	// Optional
	//
	// NOTE: FFI_USE_LogicalMinimum MUST be set to true to make this field effective.
	LogicalMinimum int32
	// LogicalMaximum is "HIDReportItem.logicalMaximum"
	//
	// Optional
	//
	// NOTE: FFI_USE_LogicalMaximum MUST be set to true to make this field effective.
	LogicalMaximum int32
	// PhysicalMinimum is "HIDReportItem.physicalMinimum"
	//
	// Optional
	//
	// NOTE: FFI_USE_PhysicalMinimum MUST be set to true to make this field effective.
	PhysicalMinimum int32
	// PhysicalMaximum is "HIDReportItem.physicalMaximum"
	//
	// Optional
	//
	// NOTE: FFI_USE_PhysicalMaximum MUST be set to true to make this field effective.
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

// New creates a new HIDReportItem in the application heap.
func (p HIDReportItem) New() js.Ref {
	return bindings.HIDReportItemJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HIDReportItem) UpdateFrom(ref js.Ref) {
	bindings.HIDReportItemJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HIDReportItem) Update(ref js.Ref) {
	bindings.HIDReportItemJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HIDReportItem) FreeMembers(recursive bool) {
	js.Free(
		p.Usages.Ref(),
		p.Strings.Ref(),
	)
	p.Usages = p.Usages.FromRef(js.Undefined)
	p.Strings = p.Strings.FromRef(js.Undefined)
}

type HIDReportInfo struct {
	// ReportId is "HIDReportInfo.reportId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ReportId MUST be set to true to make this field effective.
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

// New creates a new HIDReportInfo in the application heap.
func (p HIDReportInfo) New() js.Ref {
	return bindings.HIDReportInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HIDReportInfo) UpdateFrom(ref js.Ref) {
	bindings.HIDReportInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HIDReportInfo) Update(ref js.Ref) {
	bindings.HIDReportInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HIDReportInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Items.Ref(),
	)
	p.Items = p.Items.FromRef(js.Undefined)
}

type HIDCollectionInfo struct {
	// UsagePage is "HIDCollectionInfo.usagePage"
	//
	// Optional
	//
	// NOTE: FFI_USE_UsagePage MUST be set to true to make this field effective.
	UsagePage uint16
	// Usage is "HIDCollectionInfo.usage"
	//
	// Optional
	//
	// NOTE: FFI_USE_Usage MUST be set to true to make this field effective.
	Usage uint16
	// Type is "HIDCollectionInfo.type"
	//
	// Optional
	//
	// NOTE: FFI_USE_Type MUST be set to true to make this field effective.
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

// New creates a new HIDCollectionInfo in the application heap.
func (p HIDCollectionInfo) New() js.Ref {
	return bindings.HIDCollectionInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HIDCollectionInfo) UpdateFrom(ref js.Ref) {
	bindings.HIDCollectionInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HIDCollectionInfo) Update(ref js.Ref) {
	bindings.HIDCollectionInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HIDCollectionInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Children.Ref(),
		p.InputReports.Ref(),
		p.OutputReports.Ref(),
		p.FeatureReports.Ref(),
	)
	p.Children = p.Children.FromRef(js.Undefined)
	p.InputReports = p.InputReports.FromRef(js.Undefined)
	p.OutputReports = p.OutputReports.FromRef(js.Undefined)
	p.FeatureReports = p.FeatureReports.FromRef(js.Undefined)
}

type HIDDevice struct {
	EventTarget
}

func (this HIDDevice) Once() HIDDevice {
	this.ref.Once()
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
	this.ref.Free()
}

// Opened returns the value of property "HIDDevice.opened".
//
// It returns ok=false if there is no such property.
func (this HIDDevice) Opened() (ret bool, ok bool) {
	ok = js.True == bindings.GetHIDDeviceOpened(
		this.ref, js.Pointer(&ret),
	)
	return
}

// VendorId returns the value of property "HIDDevice.vendorId".
//
// It returns ok=false if there is no such property.
func (this HIDDevice) VendorId() (ret uint16, ok bool) {
	ok = js.True == bindings.GetHIDDeviceVendorId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ProductId returns the value of property "HIDDevice.productId".
//
// It returns ok=false if there is no such property.
func (this HIDDevice) ProductId() (ret uint16, ok bool) {
	ok = js.True == bindings.GetHIDDeviceProductId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ProductName returns the value of property "HIDDevice.productName".
//
// It returns ok=false if there is no such property.
func (this HIDDevice) ProductName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHIDDeviceProductName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Collections returns the value of property "HIDDevice.collections".
//
// It returns ok=false if there is no such property.
func (this HIDDevice) Collections() (ret js.FrozenArray[HIDCollectionInfo], ok bool) {
	ok = js.True == bindings.GetHIDDeviceCollections(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncOpen returns true if the method "HIDDevice.open" exists.
func (this HIDDevice) HasFuncOpen() bool {
	return js.True == bindings.HasFuncHIDDeviceOpen(
		this.ref,
	)
}

// FuncOpen returns the method "HIDDevice.open".
func (this HIDDevice) FuncOpen() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncHIDDeviceOpen(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Open calls the method "HIDDevice.open".
func (this HIDDevice) Open() (ret js.Promise[js.Void]) {
	bindings.CallHIDDeviceOpen(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryOpen calls the method "HIDDevice.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HIDDevice) TryOpen() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHIDDeviceOpen(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncClose returns true if the method "HIDDevice.close" exists.
func (this HIDDevice) HasFuncClose() bool {
	return js.True == bindings.HasFuncHIDDeviceClose(
		this.ref,
	)
}

// FuncClose returns the method "HIDDevice.close".
func (this HIDDevice) FuncClose() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncHIDDeviceClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "HIDDevice.close".
func (this HIDDevice) Close() (ret js.Promise[js.Void]) {
	bindings.CallHIDDeviceClose(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "HIDDevice.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HIDDevice) TryClose() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHIDDeviceClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncForget returns true if the method "HIDDevice.forget" exists.
func (this HIDDevice) HasFuncForget() bool {
	return js.True == bindings.HasFuncHIDDeviceForget(
		this.ref,
	)
}

// FuncForget returns the method "HIDDevice.forget".
func (this HIDDevice) FuncForget() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncHIDDeviceForget(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Forget calls the method "HIDDevice.forget".
func (this HIDDevice) Forget() (ret js.Promise[js.Void]) {
	bindings.CallHIDDeviceForget(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryForget calls the method "HIDDevice.forget"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HIDDevice) TryForget() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHIDDeviceForget(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSendReport returns true if the method "HIDDevice.sendReport" exists.
func (this HIDDevice) HasFuncSendReport() bool {
	return js.True == bindings.HasFuncHIDDeviceSendReport(
		this.ref,
	)
}

// FuncSendReport returns the method "HIDDevice.sendReport".
func (this HIDDevice) FuncSendReport() (fn js.Func[func(reportId uint8, data BufferSource) js.Promise[js.Void]]) {
	bindings.FuncHIDDeviceSendReport(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SendReport calls the method "HIDDevice.sendReport".
func (this HIDDevice) SendReport(reportId uint8, data BufferSource) (ret js.Promise[js.Void]) {
	bindings.CallHIDDeviceSendReport(
		this.ref, js.Pointer(&ret),
		uint32(reportId),
		data.Ref(),
	)

	return
}

// TrySendReport calls the method "HIDDevice.sendReport"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HIDDevice) TrySendReport(reportId uint8, data BufferSource) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHIDDeviceSendReport(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(reportId),
		data.Ref(),
	)

	return
}

// HasFuncSendFeatureReport returns true if the method "HIDDevice.sendFeatureReport" exists.
func (this HIDDevice) HasFuncSendFeatureReport() bool {
	return js.True == bindings.HasFuncHIDDeviceSendFeatureReport(
		this.ref,
	)
}

// FuncSendFeatureReport returns the method "HIDDevice.sendFeatureReport".
func (this HIDDevice) FuncSendFeatureReport() (fn js.Func[func(reportId uint8, data BufferSource) js.Promise[js.Void]]) {
	bindings.FuncHIDDeviceSendFeatureReport(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SendFeatureReport calls the method "HIDDevice.sendFeatureReport".
func (this HIDDevice) SendFeatureReport(reportId uint8, data BufferSource) (ret js.Promise[js.Void]) {
	bindings.CallHIDDeviceSendFeatureReport(
		this.ref, js.Pointer(&ret),
		uint32(reportId),
		data.Ref(),
	)

	return
}

// TrySendFeatureReport calls the method "HIDDevice.sendFeatureReport"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HIDDevice) TrySendFeatureReport(reportId uint8, data BufferSource) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHIDDeviceSendFeatureReport(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(reportId),
		data.Ref(),
	)

	return
}

// HasFuncReceiveFeatureReport returns true if the method "HIDDevice.receiveFeatureReport" exists.
func (this HIDDevice) HasFuncReceiveFeatureReport() bool {
	return js.True == bindings.HasFuncHIDDeviceReceiveFeatureReport(
		this.ref,
	)
}

// FuncReceiveFeatureReport returns the method "HIDDevice.receiveFeatureReport".
func (this HIDDevice) FuncReceiveFeatureReport() (fn js.Func[func(reportId uint8) js.Promise[js.DataView]]) {
	bindings.FuncHIDDeviceReceiveFeatureReport(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReceiveFeatureReport calls the method "HIDDevice.receiveFeatureReport".
func (this HIDDevice) ReceiveFeatureReport(reportId uint8) (ret js.Promise[js.DataView]) {
	bindings.CallHIDDeviceReceiveFeatureReport(
		this.ref, js.Pointer(&ret),
		uint32(reportId),
	)

	return
}

// TryReceiveFeatureReport calls the method "HIDDevice.receiveFeatureReport"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HIDDevice) TryReceiveFeatureReport(reportId uint8) (ret js.Promise[js.DataView], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHIDDeviceReceiveFeatureReport(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(reportId),
	)

	return
}

type HIDDeviceFilter struct {
	// VendorId is "HIDDeviceFilter.vendorId"
	//
	// Optional
	//
	// NOTE: FFI_USE_VendorId MUST be set to true to make this field effective.
	VendorId uint32
	// ProductId is "HIDDeviceFilter.productId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ProductId MUST be set to true to make this field effective.
	ProductId uint16
	// UsagePage is "HIDDeviceFilter.usagePage"
	//
	// Optional
	//
	// NOTE: FFI_USE_UsagePage MUST be set to true to make this field effective.
	UsagePage uint16
	// Usage is "HIDDeviceFilter.usage"
	//
	// Optional
	//
	// NOTE: FFI_USE_Usage MUST be set to true to make this field effective.
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

// New creates a new HIDDeviceFilter in the application heap.
func (p HIDDeviceFilter) New() js.Ref {
	return bindings.HIDDeviceFilterJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HIDDeviceFilter) UpdateFrom(ref js.Ref) {
	bindings.HIDDeviceFilterJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HIDDeviceFilter) Update(ref js.Ref) {
	bindings.HIDDeviceFilterJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HIDDeviceFilter) FreeMembers(recursive bool) {
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

// New creates a new HIDDeviceRequestOptions in the application heap.
func (p HIDDeviceRequestOptions) New() js.Ref {
	return bindings.HIDDeviceRequestOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HIDDeviceRequestOptions) UpdateFrom(ref js.Ref) {
	bindings.HIDDeviceRequestOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HIDDeviceRequestOptions) Update(ref js.Ref) {
	bindings.HIDDeviceRequestOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HIDDeviceRequestOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Filters.Ref(),
		p.ExclusionFilters.Ref(),
	)
	p.Filters = p.Filters.FromRef(js.Undefined)
	p.ExclusionFilters = p.ExclusionFilters.FromRef(js.Undefined)
}

type HID struct {
	EventTarget
}

func (this HID) Once() HID {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncGetDevices returns true if the method "HID.getDevices" exists.
func (this HID) HasFuncGetDevices() bool {
	return js.True == bindings.HasFuncHIDGetDevices(
		this.ref,
	)
}

// FuncGetDevices returns the method "HID.getDevices".
func (this HID) FuncGetDevices() (fn js.Func[func() js.Promise[js.Array[HIDDevice]]]) {
	bindings.FuncHIDGetDevices(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetDevices calls the method "HID.getDevices".
func (this HID) GetDevices() (ret js.Promise[js.Array[HIDDevice]]) {
	bindings.CallHIDGetDevices(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetDevices calls the method "HID.getDevices"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HID) TryGetDevices() (ret js.Promise[js.Array[HIDDevice]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHIDGetDevices(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRequestDevice returns true if the method "HID.requestDevice" exists.
func (this HID) HasFuncRequestDevice() bool {
	return js.True == bindings.HasFuncHIDRequestDevice(
		this.ref,
	)
}

// FuncRequestDevice returns the method "HID.requestDevice".
func (this HID) FuncRequestDevice() (fn js.Func[func(options HIDDeviceRequestOptions) js.Promise[js.Array[HIDDevice]]]) {
	bindings.FuncHIDRequestDevice(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestDevice calls the method "HID.requestDevice".
func (this HID) RequestDevice(options HIDDeviceRequestOptions) (ret js.Promise[js.Array[HIDDevice]]) {
	bindings.CallHIDRequestDevice(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRequestDevice calls the method "HID.requestDevice"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HID) TryRequestDevice(options HIDDeviceRequestOptions) (ret js.Promise[js.Array[HIDDevice]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHIDRequestDevice(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

type WindowControlsOverlay struct {
	EventTarget
}

func (this WindowControlsOverlay) Once() WindowControlsOverlay {
	this.ref.Once()
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
	this.ref.Free()
}

// Visible returns the value of property "WindowControlsOverlay.visible".
//
// It returns ok=false if there is no such property.
func (this WindowControlsOverlay) Visible() (ret bool, ok bool) {
	ok = js.True == bindings.GetWindowControlsOverlayVisible(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetTitlebarAreaRect returns true if the method "WindowControlsOverlay.getTitlebarAreaRect" exists.
func (this WindowControlsOverlay) HasFuncGetTitlebarAreaRect() bool {
	return js.True == bindings.HasFuncWindowControlsOverlayGetTitlebarAreaRect(
		this.ref,
	)
}

// FuncGetTitlebarAreaRect returns the method "WindowControlsOverlay.getTitlebarAreaRect".
func (this WindowControlsOverlay) FuncGetTitlebarAreaRect() (fn js.Func[func() DOMRect]) {
	bindings.FuncWindowControlsOverlayGetTitlebarAreaRect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetTitlebarAreaRect calls the method "WindowControlsOverlay.getTitlebarAreaRect".
func (this WindowControlsOverlay) GetTitlebarAreaRect() (ret DOMRect) {
	bindings.CallWindowControlsOverlayGetTitlebarAreaRect(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetTitlebarAreaRect calls the method "WindowControlsOverlay.getTitlebarAreaRect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WindowControlsOverlay) TryGetTitlebarAreaRect() (ret DOMRect, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowControlsOverlayGetTitlebarAreaRect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type Credential struct {
	ref js.Ref
}

func (this Credential) Once() Credential {
	this.ref.Once()
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
	this.ref.Free()
}

// Id returns the value of property "Credential.id".
//
// It returns ok=false if there is no such property.
func (this Credential) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCredentialId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Type returns the value of property "Credential.type".
//
// It returns ok=false if there is no such property.
func (this Credential) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCredentialType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncIsConditionalMediationAvailable returns true if the static method "Credential.isConditionalMediationAvailable" exists.
func (this Credential) HasFuncIsConditionalMediationAvailable() bool {
	return js.True == bindings.HasFuncCredentialIsConditionalMediationAvailable(
		this.ref,
	)
}

// FuncIsConditionalMediationAvailable returns the static method "Credential.isConditionalMediationAvailable".
func (this Credential) FuncIsConditionalMediationAvailable() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncCredentialIsConditionalMediationAvailable(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsConditionalMediationAvailable calls the static method "Credential.isConditionalMediationAvailable".
func (this Credential) IsConditionalMediationAvailable() (ret js.Promise[js.Boolean]) {
	bindings.CallCredentialIsConditionalMediationAvailable(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryIsConditionalMediationAvailable calls the static method "Credential.isConditionalMediationAvailable"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Credential) TryIsConditionalMediationAvailable() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCredentialIsConditionalMediationAvailable(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

// New creates a new OTPCredentialRequestOptions in the application heap.
func (p OTPCredentialRequestOptions) New() js.Ref {
	return bindings.OTPCredentialRequestOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OTPCredentialRequestOptions) UpdateFrom(ref js.Ref) {
	bindings.OTPCredentialRequestOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OTPCredentialRequestOptions) Update(ref js.Ref) {
	bindings.OTPCredentialRequestOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OTPCredentialRequestOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Transport.Ref(),
	)
	p.Transport = p.Transport.FromRef(js.Undefined)
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

// New creates a new PublicKeyCredentialDescriptor in the application heap.
func (p PublicKeyCredentialDescriptor) New() js.Ref {
	return bindings.PublicKeyCredentialDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PublicKeyCredentialDescriptor) UpdateFrom(ref js.Ref) {
	bindings.PublicKeyCredentialDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PublicKeyCredentialDescriptor) Update(ref js.Ref) {
	bindings.PublicKeyCredentialDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PublicKeyCredentialDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Type.Ref(),
		p.Id.Ref(),
		p.Transports.Ref(),
	)
	p.Type = p.Type.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Transports = p.Transports.FromRef(js.Undefined)
}

type PublicKeyCredentialRequestOptions struct {
	// Challenge is "PublicKeyCredentialRequestOptions.challenge"
	//
	// Required
	Challenge BufferSource
	// Timeout is "PublicKeyCredentialRequestOptions.timeout"
	//
	// Optional
	//
	// NOTE: FFI_USE_Timeout MUST be set to true to make this field effective.
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
	//
	// NOTE: Extensions.FFI_USE MUST be set to true to get Extensions used.
	Extensions AuthenticationExtensionsClientInputs

	FFI_USE_Timeout bool // for Timeout.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PublicKeyCredentialRequestOptions with all fields set.
func (p PublicKeyCredentialRequestOptions) FromRef(ref js.Ref) PublicKeyCredentialRequestOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PublicKeyCredentialRequestOptions in the application heap.
func (p PublicKeyCredentialRequestOptions) New() js.Ref {
	return bindings.PublicKeyCredentialRequestOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PublicKeyCredentialRequestOptions) UpdateFrom(ref js.Ref) {
	bindings.PublicKeyCredentialRequestOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PublicKeyCredentialRequestOptions) Update(ref js.Ref) {
	bindings.PublicKeyCredentialRequestOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PublicKeyCredentialRequestOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Challenge.Ref(),
		p.RpId.Ref(),
		p.AllowCredentials.Ref(),
		p.UserVerification.Ref(),
		p.Hints.Ref(),
		p.Attestation.Ref(),
		p.AttestationFormats.Ref(),
	)
	p.Challenge = p.Challenge.FromRef(js.Undefined)
	p.RpId = p.RpId.FromRef(js.Undefined)
	p.AllowCredentials = p.AllowCredentials.FromRef(js.Undefined)
	p.UserVerification = p.UserVerification.FromRef(js.Undefined)
	p.Hints = p.Hints.FromRef(js.Undefined)
	p.Attestation = p.Attestation.FromRef(js.Undefined)
	p.AttestationFormats = p.AttestationFormats.FromRef(js.Undefined)
	if recursive {
		p.Extensions.FreeMembers(true)
	}
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

// New creates a new IdentityProviderConfig in the application heap.
func (p IdentityProviderConfig) New() js.Ref {
	return bindings.IdentityProviderConfigJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *IdentityProviderConfig) UpdateFrom(ref js.Ref) {
	bindings.IdentityProviderConfigJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IdentityProviderConfig) Update(ref js.Ref) {
	bindings.IdentityProviderConfigJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IdentityProviderConfig) FreeMembers(recursive bool) {
	js.Free(
		p.ConfigURL.Ref(),
		p.ClientId.Ref(),
		p.Nonce.Ref(),
		p.LoginHint.Ref(),
	)
	p.ConfigURL = p.ConfigURL.FromRef(js.Undefined)
	p.ClientId = p.ClientId.FromRef(js.Undefined)
	p.Nonce = p.Nonce.FromRef(js.Undefined)
	p.LoginHint = p.LoginHint.FromRef(js.Undefined)
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

// New creates a new IdentityCredentialRequestOptions in the application heap.
func (p IdentityCredentialRequestOptions) New() js.Ref {
	return bindings.IdentityCredentialRequestOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *IdentityCredentialRequestOptions) UpdateFrom(ref js.Ref) {
	bindings.IdentityCredentialRequestOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IdentityCredentialRequestOptions) Update(ref js.Ref) {
	bindings.IdentityCredentialRequestOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IdentityCredentialRequestOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Providers.Ref(),
	)
	p.Providers = p.Providers.FromRef(js.Undefined)
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

// New creates a new FederatedCredentialRequestOptions in the application heap.
func (p FederatedCredentialRequestOptions) New() js.Ref {
	return bindings.FederatedCredentialRequestOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FederatedCredentialRequestOptions) UpdateFrom(ref js.Ref) {
	bindings.FederatedCredentialRequestOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FederatedCredentialRequestOptions) Update(ref js.Ref) {
	bindings.FederatedCredentialRequestOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FederatedCredentialRequestOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Providers.Ref(),
		p.Protocols.Ref(),
	)
	p.Providers = p.Providers.FromRef(js.Undefined)
	p.Protocols = p.Protocols.FromRef(js.Undefined)
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
	//
	// NOTE: Otp.FFI_USE MUST be set to true to get Otp used.
	Otp OTPCredentialRequestOptions
	// PublicKey is "CredentialRequestOptions.publicKey"
	//
	// Optional
	//
	// NOTE: PublicKey.FFI_USE MUST be set to true to get PublicKey used.
	PublicKey PublicKeyCredentialRequestOptions
	// Password is "CredentialRequestOptions.password"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Password MUST be set to true to make this field effective.
	Password bool
	// Identity is "CredentialRequestOptions.identity"
	//
	// Optional
	//
	// NOTE: Identity.FFI_USE MUST be set to true to get Identity used.
	Identity IdentityCredentialRequestOptions
	// Federated is "CredentialRequestOptions.federated"
	//
	// Optional
	//
	// NOTE: Federated.FFI_USE MUST be set to true to get Federated used.
	Federated FederatedCredentialRequestOptions

	FFI_USE_Password bool // for Password.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CredentialRequestOptions with all fields set.
func (p CredentialRequestOptions) FromRef(ref js.Ref) CredentialRequestOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CredentialRequestOptions in the application heap.
func (p CredentialRequestOptions) New() js.Ref {
	return bindings.CredentialRequestOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CredentialRequestOptions) UpdateFrom(ref js.Ref) {
	bindings.CredentialRequestOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CredentialRequestOptions) Update(ref js.Ref) {
	bindings.CredentialRequestOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CredentialRequestOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Signal.Ref(),
	)
	p.Signal = p.Signal.FromRef(js.Undefined)
	if recursive {
		p.Otp.FreeMembers(true)
		p.PublicKey.FreeMembers(true)
		p.Identity.FreeMembers(true)
		p.Federated.FreeMembers(true)
	}
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

// New creates a new PublicKeyCredentialRpEntity in the application heap.
func (p PublicKeyCredentialRpEntity) New() js.Ref {
	return bindings.PublicKeyCredentialRpEntityJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PublicKeyCredentialRpEntity) UpdateFrom(ref js.Ref) {
	bindings.PublicKeyCredentialRpEntityJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PublicKeyCredentialRpEntity) Update(ref js.Ref) {
	bindings.PublicKeyCredentialRpEntityJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PublicKeyCredentialRpEntity) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.Name.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
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

// New creates a new PublicKeyCredentialUserEntity in the application heap.
func (p PublicKeyCredentialUserEntity) New() js.Ref {
	return bindings.PublicKeyCredentialUserEntityJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PublicKeyCredentialUserEntity) UpdateFrom(ref js.Ref) {
	bindings.PublicKeyCredentialUserEntityJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PublicKeyCredentialUserEntity) Update(ref js.Ref) {
	bindings.PublicKeyCredentialUserEntityJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PublicKeyCredentialUserEntity) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.DisplayName.Ref(),
		p.Name.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.DisplayName = p.DisplayName.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
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

// New creates a new PublicKeyCredentialParameters in the application heap.
func (p PublicKeyCredentialParameters) New() js.Ref {
	return bindings.PublicKeyCredentialParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PublicKeyCredentialParameters) UpdateFrom(ref js.Ref) {
	bindings.PublicKeyCredentialParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PublicKeyCredentialParameters) Update(ref js.Ref) {
	bindings.PublicKeyCredentialParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PublicKeyCredentialParameters) FreeMembers(recursive bool) {
	js.Free(
		p.Type.Ref(),
	)
	p.Type = p.Type.FromRef(js.Undefined)
}

type PublicKeyCredentialCreationOptions struct {
	// Rp is "PublicKeyCredentialCreationOptions.rp"
	//
	// Required
	//
	// NOTE: Rp.FFI_USE MUST be set to true to get Rp used.
	Rp PublicKeyCredentialRpEntity
	// User is "PublicKeyCredentialCreationOptions.user"
	//
	// Required
	//
	// NOTE: User.FFI_USE MUST be set to true to get User used.
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
	//
	// NOTE: FFI_USE_Timeout MUST be set to true to make this field effective.
	Timeout uint32
	// ExcludeCredentials is "PublicKeyCredentialCreationOptions.excludeCredentials"
	//
	// Optional, defaults to [].
	ExcludeCredentials js.Array[PublicKeyCredentialDescriptor]
	// AuthenticatorSelection is "PublicKeyCredentialCreationOptions.authenticatorSelection"
	//
	// Optional
	//
	// NOTE: AuthenticatorSelection.FFI_USE MUST be set to true to get AuthenticatorSelection used.
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
	//
	// NOTE: Extensions.FFI_USE MUST be set to true to get Extensions used.
	Extensions AuthenticationExtensionsClientInputs

	FFI_USE_Timeout bool // for Timeout.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PublicKeyCredentialCreationOptions with all fields set.
func (p PublicKeyCredentialCreationOptions) FromRef(ref js.Ref) PublicKeyCredentialCreationOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PublicKeyCredentialCreationOptions in the application heap.
func (p PublicKeyCredentialCreationOptions) New() js.Ref {
	return bindings.PublicKeyCredentialCreationOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PublicKeyCredentialCreationOptions) UpdateFrom(ref js.Ref) {
	bindings.PublicKeyCredentialCreationOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PublicKeyCredentialCreationOptions) Update(ref js.Ref) {
	bindings.PublicKeyCredentialCreationOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PublicKeyCredentialCreationOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Challenge.Ref(),
		p.PubKeyCredParams.Ref(),
		p.ExcludeCredentials.Ref(),
		p.Hints.Ref(),
		p.Attestation.Ref(),
		p.AttestationFormats.Ref(),
	)
	p.Challenge = p.Challenge.FromRef(js.Undefined)
	p.PubKeyCredParams = p.PubKeyCredParams.FromRef(js.Undefined)
	p.ExcludeCredentials = p.ExcludeCredentials.FromRef(js.Undefined)
	p.Hints = p.Hints.FromRef(js.Undefined)
	p.Attestation = p.Attestation.FromRef(js.Undefined)
	p.AttestationFormats = p.AttestationFormats.FromRef(js.Undefined)
	if recursive {
		p.Rp.FreeMembers(true)
		p.User.FreeMembers(true)
		p.AuthenticatorSelection.FreeMembers(true)
		p.Extensions.FreeMembers(true)
	}
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

// New creates a new PasswordCredentialData in the application heap.
func (p PasswordCredentialData) New() js.Ref {
	return bindings.PasswordCredentialDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PasswordCredentialData) UpdateFrom(ref js.Ref) {
	bindings.PasswordCredentialDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PasswordCredentialData) Update(ref js.Ref) {
	bindings.PasswordCredentialDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PasswordCredentialData) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.IconURL.Ref(),
		p.Origin.Ref(),
		p.Password.Ref(),
		p.Id.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.IconURL = p.IconURL.FromRef(js.Undefined)
	p.Origin = p.Origin.FromRef(js.Undefined)
	p.Password = p.Password.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
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

// New creates a new FederatedCredentialInit in the application heap.
func (p FederatedCredentialInit) New() js.Ref {
	return bindings.FederatedCredentialInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FederatedCredentialInit) UpdateFrom(ref js.Ref) {
	bindings.FederatedCredentialInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FederatedCredentialInit) Update(ref js.Ref) {
	bindings.FederatedCredentialInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FederatedCredentialInit) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.IconURL.Ref(),
		p.Origin.Ref(),
		p.Provider.Ref(),
		p.Protocol.Ref(),
		p.Id.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.IconURL = p.IconURL.FromRef(js.Undefined)
	p.Origin = p.Origin.FromRef(js.Undefined)
	p.Provider = p.Provider.FromRef(js.Undefined)
	p.Protocol = p.Protocol.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
}

type CredentialCreationOptions struct {
	// Signal is "CredentialCreationOptions.signal"
	//
	// Optional
	Signal AbortSignal
	// PublicKey is "CredentialCreationOptions.publicKey"
	//
	// Optional
	//
	// NOTE: PublicKey.FFI_USE MUST be set to true to get PublicKey used.
	PublicKey PublicKeyCredentialCreationOptions
	// Password is "CredentialCreationOptions.password"
	//
	// Optional
	Password PasswordCredentialInit
	// Federated is "CredentialCreationOptions.federated"
	//
	// Optional
	//
	// NOTE: Federated.FFI_USE MUST be set to true to get Federated used.
	Federated FederatedCredentialInit

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CredentialCreationOptions with all fields set.
func (p CredentialCreationOptions) FromRef(ref js.Ref) CredentialCreationOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CredentialCreationOptions in the application heap.
func (p CredentialCreationOptions) New() js.Ref {
	return bindings.CredentialCreationOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CredentialCreationOptions) UpdateFrom(ref js.Ref) {
	bindings.CredentialCreationOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CredentialCreationOptions) Update(ref js.Ref) {
	bindings.CredentialCreationOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CredentialCreationOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Signal.Ref(),
		p.Password.Ref(),
	)
	p.Signal = p.Signal.FromRef(js.Undefined)
	p.Password = p.Password.FromRef(js.Undefined)
	if recursive {
		p.PublicKey.FreeMembers(true)
		p.Federated.FreeMembers(true)
	}
}

type CredentialsContainer struct {
	ref js.Ref
}

func (this CredentialsContainer) Once() CredentialsContainer {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncGet returns true if the method "CredentialsContainer.get" exists.
func (this CredentialsContainer) HasFuncGet() bool {
	return js.True == bindings.HasFuncCredentialsContainerGet(
		this.ref,
	)
}

// FuncGet returns the method "CredentialsContainer.get".
func (this CredentialsContainer) FuncGet() (fn js.Func[func(options CredentialRequestOptions) js.Promise[Credential]]) {
	bindings.FuncCredentialsContainerGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "CredentialsContainer.get".
func (this CredentialsContainer) Get(options CredentialRequestOptions) (ret js.Promise[Credential]) {
	bindings.CallCredentialsContainerGet(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGet calls the method "CredentialsContainer.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CredentialsContainer) TryGet(options CredentialRequestOptions) (ret js.Promise[Credential], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCredentialsContainerGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncGet1 returns true if the method "CredentialsContainer.get" exists.
func (this CredentialsContainer) HasFuncGet1() bool {
	return js.True == bindings.HasFuncCredentialsContainerGet1(
		this.ref,
	)
}

// FuncGet1 returns the method "CredentialsContainer.get".
func (this CredentialsContainer) FuncGet1() (fn js.Func[func() js.Promise[Credential]]) {
	bindings.FuncCredentialsContainerGet1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get1 calls the method "CredentialsContainer.get".
func (this CredentialsContainer) Get1() (ret js.Promise[Credential]) {
	bindings.CallCredentialsContainerGet1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGet1 calls the method "CredentialsContainer.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CredentialsContainer) TryGet1() (ret js.Promise[Credential], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCredentialsContainerGet1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStore returns true if the method "CredentialsContainer.store" exists.
func (this CredentialsContainer) HasFuncStore() bool {
	return js.True == bindings.HasFuncCredentialsContainerStore(
		this.ref,
	)
}

// FuncStore returns the method "CredentialsContainer.store".
func (this CredentialsContainer) FuncStore() (fn js.Func[func(credential Credential) js.Promise[js.Void]]) {
	bindings.FuncCredentialsContainerStore(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Store calls the method "CredentialsContainer.store".
func (this CredentialsContainer) Store(credential Credential) (ret js.Promise[js.Void]) {
	bindings.CallCredentialsContainerStore(
		this.ref, js.Pointer(&ret),
		credential.Ref(),
	)

	return
}

// TryStore calls the method "CredentialsContainer.store"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CredentialsContainer) TryStore(credential Credential) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCredentialsContainerStore(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		credential.Ref(),
	)

	return
}

// HasFuncCreate returns true if the method "CredentialsContainer.create" exists.
func (this CredentialsContainer) HasFuncCreate() bool {
	return js.True == bindings.HasFuncCredentialsContainerCreate(
		this.ref,
	)
}

// FuncCreate returns the method "CredentialsContainer.create".
func (this CredentialsContainer) FuncCreate() (fn js.Func[func(options CredentialCreationOptions) js.Promise[Credential]]) {
	bindings.FuncCredentialsContainerCreate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Create calls the method "CredentialsContainer.create".
func (this CredentialsContainer) Create(options CredentialCreationOptions) (ret js.Promise[Credential]) {
	bindings.CallCredentialsContainerCreate(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryCreate calls the method "CredentialsContainer.create"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CredentialsContainer) TryCreate(options CredentialCreationOptions) (ret js.Promise[Credential], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCredentialsContainerCreate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncCreate1 returns true if the method "CredentialsContainer.create" exists.
func (this CredentialsContainer) HasFuncCreate1() bool {
	return js.True == bindings.HasFuncCredentialsContainerCreate1(
		this.ref,
	)
}

// FuncCreate1 returns the method "CredentialsContainer.create".
func (this CredentialsContainer) FuncCreate1() (fn js.Func[func() js.Promise[Credential]]) {
	bindings.FuncCredentialsContainerCreate1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Create1 calls the method "CredentialsContainer.create".
func (this CredentialsContainer) Create1() (ret js.Promise[Credential]) {
	bindings.CallCredentialsContainerCreate1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreate1 calls the method "CredentialsContainer.create"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CredentialsContainer) TryCreate1() (ret js.Promise[Credential], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCredentialsContainerCreate1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPreventSilentAccess returns true if the method "CredentialsContainer.preventSilentAccess" exists.
func (this CredentialsContainer) HasFuncPreventSilentAccess() bool {
	return js.True == bindings.HasFuncCredentialsContainerPreventSilentAccess(
		this.ref,
	)
}

// FuncPreventSilentAccess returns the method "CredentialsContainer.preventSilentAccess".
func (this CredentialsContainer) FuncPreventSilentAccess() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncCredentialsContainerPreventSilentAccess(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PreventSilentAccess calls the method "CredentialsContainer.preventSilentAccess".
func (this CredentialsContainer) PreventSilentAccess() (ret js.Promise[js.Void]) {
	bindings.CallCredentialsContainerPreventSilentAccess(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPreventSilentAccess calls the method "CredentialsContainer.preventSilentAccess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CredentialsContainer) TryPreventSilentAccess() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCredentialsContainerPreventSilentAccess(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
		js.ThrowInvalidCallbackInvocation()
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
	this.ref.Once()
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
	this.ref.Free()
}

// Accuracy returns the value of property "GeolocationCoordinates.accuracy".
//
// It returns ok=false if there is no such property.
func (this GeolocationCoordinates) Accuracy() (ret float64, ok bool) {
	ok = js.True == bindings.GetGeolocationCoordinatesAccuracy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Latitude returns the value of property "GeolocationCoordinates.latitude".
//
// It returns ok=false if there is no such property.
func (this GeolocationCoordinates) Latitude() (ret float64, ok bool) {
	ok = js.True == bindings.GetGeolocationCoordinatesLatitude(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Longitude returns the value of property "GeolocationCoordinates.longitude".
//
// It returns ok=false if there is no such property.
func (this GeolocationCoordinates) Longitude() (ret float64, ok bool) {
	ok = js.True == bindings.GetGeolocationCoordinatesLongitude(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Altitude returns the value of property "GeolocationCoordinates.altitude".
//
// It returns ok=false if there is no such property.
func (this GeolocationCoordinates) Altitude() (ret float64, ok bool) {
	ok = js.True == bindings.GetGeolocationCoordinatesAltitude(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AltitudeAccuracy returns the value of property "GeolocationCoordinates.altitudeAccuracy".
//
// It returns ok=false if there is no such property.
func (this GeolocationCoordinates) AltitudeAccuracy() (ret float64, ok bool) {
	ok = js.True == bindings.GetGeolocationCoordinatesAltitudeAccuracy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Heading returns the value of property "GeolocationCoordinates.heading".
//
// It returns ok=false if there is no such property.
func (this GeolocationCoordinates) Heading() (ret float64, ok bool) {
	ok = js.True == bindings.GetGeolocationCoordinatesHeading(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Speed returns the value of property "GeolocationCoordinates.speed".
//
// It returns ok=false if there is no such property.
func (this GeolocationCoordinates) Speed() (ret float64, ok bool) {
	ok = js.True == bindings.GetGeolocationCoordinatesSpeed(
		this.ref, js.Pointer(&ret),
	)
	return
}

type EpochTimeStamp uint64

type GeolocationPosition struct {
	ref js.Ref
}

func (this GeolocationPosition) Once() GeolocationPosition {
	this.ref.Once()
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
	this.ref.Free()
}

// Coords returns the value of property "GeolocationPosition.coords".
//
// It returns ok=false if there is no such property.
func (this GeolocationPosition) Coords() (ret GeolocationCoordinates, ok bool) {
	ok = js.True == bindings.GetGeolocationPositionCoords(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Timestamp returns the value of property "GeolocationPosition.timestamp".
//
// It returns ok=false if there is no such property.
func (this GeolocationPosition) Timestamp() (ret EpochTimeStamp, ok bool) {
	ok = js.True == bindings.GetGeolocationPositionTimestamp(
		this.ref, js.Pointer(&ret),
	)
	return
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
		js.ThrowInvalidCallbackInvocation()
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
	this.ref.Once()
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
	this.ref.Free()
}

// Code returns the value of property "GeolocationPositionError.code".
//
// It returns ok=false if there is no such property.
func (this GeolocationPositionError) Code() (ret uint16, ok bool) {
	ok = js.True == bindings.GetGeolocationPositionErrorCode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Message returns the value of property "GeolocationPositionError.message".
//
// It returns ok=false if there is no such property.
func (this GeolocationPositionError) Message() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGeolocationPositionErrorMessage(
		this.ref, js.Pointer(&ret),
	)
	return
}

type PositionOptions struct {
	// EnableHighAccuracy is "PositionOptions.enableHighAccuracy"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_EnableHighAccuracy MUST be set to true to make this field effective.
	EnableHighAccuracy bool
	// Timeout is "PositionOptions.timeout"
	//
	// Optional, defaults to 0xFFFFFFFF.
	//
	// NOTE: FFI_USE_Timeout MUST be set to true to make this field effective.
	Timeout uint32
	// MaximumAge is "PositionOptions.maximumAge"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_MaximumAge MUST be set to true to make this field effective.
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

// New creates a new PositionOptions in the application heap.
func (p PositionOptions) New() js.Ref {
	return bindings.PositionOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PositionOptions) UpdateFrom(ref js.Ref) {
	bindings.PositionOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PositionOptions) Update(ref js.Ref) {
	bindings.PositionOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PositionOptions) FreeMembers(recursive bool) {
}

type Geolocation struct {
	ref js.Ref
}

func (this Geolocation) Once() Geolocation {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncGetCurrentPosition returns true if the method "Geolocation.getCurrentPosition" exists.
func (this Geolocation) HasFuncGetCurrentPosition() bool {
	return js.True == bindings.HasFuncGeolocationGetCurrentPosition(
		this.ref,
	)
}

// FuncGetCurrentPosition returns the method "Geolocation.getCurrentPosition".
func (this Geolocation) FuncGetCurrentPosition() (fn js.Func[func(successCallback js.Func[func(position GeolocationPosition)], errorCallback js.Func[func(positionError GeolocationPositionError)], options PositionOptions)]) {
	bindings.FuncGeolocationGetCurrentPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetCurrentPosition calls the method "Geolocation.getCurrentPosition".
func (this Geolocation) GetCurrentPosition(successCallback js.Func[func(position GeolocationPosition)], errorCallback js.Func[func(positionError GeolocationPositionError)], options PositionOptions) (ret js.Void) {
	bindings.CallGeolocationGetCurrentPosition(
		this.ref, js.Pointer(&ret),
		successCallback.Ref(),
		errorCallback.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryGetCurrentPosition calls the method "Geolocation.getCurrentPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Geolocation) TryGetCurrentPosition(successCallback js.Func[func(position GeolocationPosition)], errorCallback js.Func[func(positionError GeolocationPositionError)], options PositionOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGeolocationGetCurrentPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		successCallback.Ref(),
		errorCallback.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncGetCurrentPosition1 returns true if the method "Geolocation.getCurrentPosition" exists.
func (this Geolocation) HasFuncGetCurrentPosition1() bool {
	return js.True == bindings.HasFuncGeolocationGetCurrentPosition1(
		this.ref,
	)
}

// FuncGetCurrentPosition1 returns the method "Geolocation.getCurrentPosition".
func (this Geolocation) FuncGetCurrentPosition1() (fn js.Func[func(successCallback js.Func[func(position GeolocationPosition)], errorCallback js.Func[func(positionError GeolocationPositionError)])]) {
	bindings.FuncGeolocationGetCurrentPosition1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetCurrentPosition1 calls the method "Geolocation.getCurrentPosition".
func (this Geolocation) GetCurrentPosition1(successCallback js.Func[func(position GeolocationPosition)], errorCallback js.Func[func(positionError GeolocationPositionError)]) (ret js.Void) {
	bindings.CallGeolocationGetCurrentPosition1(
		this.ref, js.Pointer(&ret),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// TryGetCurrentPosition1 calls the method "Geolocation.getCurrentPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Geolocation) TryGetCurrentPosition1(successCallback js.Func[func(position GeolocationPosition)], errorCallback js.Func[func(positionError GeolocationPositionError)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGeolocationGetCurrentPosition1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// HasFuncGetCurrentPosition2 returns true if the method "Geolocation.getCurrentPosition" exists.
func (this Geolocation) HasFuncGetCurrentPosition2() bool {
	return js.True == bindings.HasFuncGeolocationGetCurrentPosition2(
		this.ref,
	)
}

// FuncGetCurrentPosition2 returns the method "Geolocation.getCurrentPosition".
func (this Geolocation) FuncGetCurrentPosition2() (fn js.Func[func(successCallback js.Func[func(position GeolocationPosition)])]) {
	bindings.FuncGeolocationGetCurrentPosition2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetCurrentPosition2 calls the method "Geolocation.getCurrentPosition".
func (this Geolocation) GetCurrentPosition2(successCallback js.Func[func(position GeolocationPosition)]) (ret js.Void) {
	bindings.CallGeolocationGetCurrentPosition2(
		this.ref, js.Pointer(&ret),
		successCallback.Ref(),
	)

	return
}

// TryGetCurrentPosition2 calls the method "Geolocation.getCurrentPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Geolocation) TryGetCurrentPosition2(successCallback js.Func[func(position GeolocationPosition)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGeolocationGetCurrentPosition2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		successCallback.Ref(),
	)

	return
}

// HasFuncWatchPosition returns true if the method "Geolocation.watchPosition" exists.
func (this Geolocation) HasFuncWatchPosition() bool {
	return js.True == bindings.HasFuncGeolocationWatchPosition(
		this.ref,
	)
}

// FuncWatchPosition returns the method "Geolocation.watchPosition".
func (this Geolocation) FuncWatchPosition() (fn js.Func[func(successCallback js.Func[func(position GeolocationPosition)], errorCallback js.Func[func(positionError GeolocationPositionError)], options PositionOptions) int32]) {
	bindings.FuncGeolocationWatchPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// WatchPosition calls the method "Geolocation.watchPosition".
func (this Geolocation) WatchPosition(successCallback js.Func[func(position GeolocationPosition)], errorCallback js.Func[func(positionError GeolocationPositionError)], options PositionOptions) (ret int32) {
	bindings.CallGeolocationWatchPosition(
		this.ref, js.Pointer(&ret),
		successCallback.Ref(),
		errorCallback.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryWatchPosition calls the method "Geolocation.watchPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Geolocation) TryWatchPosition(successCallback js.Func[func(position GeolocationPosition)], errorCallback js.Func[func(positionError GeolocationPositionError)], options PositionOptions) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGeolocationWatchPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		successCallback.Ref(),
		errorCallback.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncWatchPosition1 returns true if the method "Geolocation.watchPosition" exists.
func (this Geolocation) HasFuncWatchPosition1() bool {
	return js.True == bindings.HasFuncGeolocationWatchPosition1(
		this.ref,
	)
}

// FuncWatchPosition1 returns the method "Geolocation.watchPosition".
func (this Geolocation) FuncWatchPosition1() (fn js.Func[func(successCallback js.Func[func(position GeolocationPosition)], errorCallback js.Func[func(positionError GeolocationPositionError)]) int32]) {
	bindings.FuncGeolocationWatchPosition1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// WatchPosition1 calls the method "Geolocation.watchPosition".
func (this Geolocation) WatchPosition1(successCallback js.Func[func(position GeolocationPosition)], errorCallback js.Func[func(positionError GeolocationPositionError)]) (ret int32) {
	bindings.CallGeolocationWatchPosition1(
		this.ref, js.Pointer(&ret),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// TryWatchPosition1 calls the method "Geolocation.watchPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Geolocation) TryWatchPosition1(successCallback js.Func[func(position GeolocationPosition)], errorCallback js.Func[func(positionError GeolocationPositionError)]) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGeolocationWatchPosition1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// HasFuncWatchPosition2 returns true if the method "Geolocation.watchPosition" exists.
func (this Geolocation) HasFuncWatchPosition2() bool {
	return js.True == bindings.HasFuncGeolocationWatchPosition2(
		this.ref,
	)
}

// FuncWatchPosition2 returns the method "Geolocation.watchPosition".
func (this Geolocation) FuncWatchPosition2() (fn js.Func[func(successCallback js.Func[func(position GeolocationPosition)]) int32]) {
	bindings.FuncGeolocationWatchPosition2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// WatchPosition2 calls the method "Geolocation.watchPosition".
func (this Geolocation) WatchPosition2(successCallback js.Func[func(position GeolocationPosition)]) (ret int32) {
	bindings.CallGeolocationWatchPosition2(
		this.ref, js.Pointer(&ret),
		successCallback.Ref(),
	)

	return
}

// TryWatchPosition2 calls the method "Geolocation.watchPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Geolocation) TryWatchPosition2(successCallback js.Func[func(position GeolocationPosition)]) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGeolocationWatchPosition2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		successCallback.Ref(),
	)

	return
}

// HasFuncClearWatch returns true if the method "Geolocation.clearWatch" exists.
func (this Geolocation) HasFuncClearWatch() bool {
	return js.True == bindings.HasFuncGeolocationClearWatch(
		this.ref,
	)
}

// FuncClearWatch returns the method "Geolocation.clearWatch".
func (this Geolocation) FuncClearWatch() (fn js.Func[func(watchId int32)]) {
	bindings.FuncGeolocationClearWatch(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearWatch calls the method "Geolocation.clearWatch".
func (this Geolocation) ClearWatch(watchId int32) (ret js.Void) {
	bindings.CallGeolocationClearWatch(
		this.ref, js.Pointer(&ret),
		int32(watchId),
	)

	return
}

// TryClearWatch calls the method "Geolocation.clearWatch"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Geolocation) TryClearWatch(watchId int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGeolocationClearWatch(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(watchId),
	)

	return
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

func NewUSBInTransferResult(status USBTransferStatus, data js.DataView) (ret USBInTransferResult) {
	ret.ref = bindings.NewUSBInTransferResultByUSBInTransferResult(
		uint32(status),
		data.Ref())
	return
}

func NewUSBInTransferResultByUSBInTransferResult1(status USBTransferStatus) (ret USBInTransferResult) {
	ret.ref = bindings.NewUSBInTransferResultByUSBInTransferResult1(
		uint32(status))
	return
}

type USBInTransferResult struct {
	ref js.Ref
}

func (this USBInTransferResult) Once() USBInTransferResult {
	this.ref.Once()
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
	this.ref.Free()
}

// Data returns the value of property "USBInTransferResult.data".
//
// It returns ok=false if there is no such property.
func (this USBInTransferResult) Data() (ret js.DataView, ok bool) {
	ok = js.True == bindings.GetUSBInTransferResultData(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Status returns the value of property "USBInTransferResult.status".
//
// It returns ok=false if there is no such property.
func (this USBInTransferResult) Status() (ret USBTransferStatus, ok bool) {
	ok = js.True == bindings.GetUSBInTransferResultStatus(
		this.ref, js.Pointer(&ret),
	)
	return
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
