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

//go:wasmimport plat/js/webext/smartcardproviderprivate constof_ConnectionState
//go:noescape
func ConstOfConnectionState(str js.Ref) uint32

//go:wasmimport plat/js/webext/smartcardproviderprivate constof_Disposition
//go:noescape
func ConstOfDisposition(str js.Ref) uint32

//go:wasmimport plat/js/webext/smartcardproviderprivate constof_Protocol
//go:noescape
func ConstOfProtocol(str js.Ref) uint32

//go:wasmimport plat/js/webext/smartcardproviderprivate store_Protocols
//go:noescape
func ProtocolsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate load_Protocols
//go:noescape
func ProtocolsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate store_ReaderStateFlags
//go:noescape
func ReaderStateFlagsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate load_ReaderStateFlags
//go:noescape
func ReaderStateFlagsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate store_ReaderStateIn
//go:noescape
func ReaderStateInJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate load_ReaderStateIn
//go:noescape
func ReaderStateInJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate store_ReaderStateOut
//go:noescape
func ReaderStateOutJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate load_ReaderStateOut
//go:noescape
func ReaderStateOutJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate constof_ResultCode
//go:noescape
func ConstOfResultCode(str js.Ref) uint32

//go:wasmimport plat/js/webext/smartcardproviderprivate constof_ShareMode
//go:noescape
func ConstOfShareMode(str js.Ref) uint32

//go:wasmimport plat/js/webext/smartcardproviderprivate store_Timeout
//go:noescape
func TimeoutJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate load_Timeout
//go:noescape
func TimeoutJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OnBeginTransactionRequested
//go:noescape
func HasFuncOnBeginTransactionRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OnBeginTransactionRequested
//go:noescape
func FuncOnBeginTransactionRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OnBeginTransactionRequested
//go:noescape
func CallOnBeginTransactionRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OnBeginTransactionRequested
//go:noescape
func TryOnBeginTransactionRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OffBeginTransactionRequested
//go:noescape
func HasFuncOffBeginTransactionRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OffBeginTransactionRequested
//go:noescape
func FuncOffBeginTransactionRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OffBeginTransactionRequested
//go:noescape
func CallOffBeginTransactionRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OffBeginTransactionRequested
//go:noescape
func TryOffBeginTransactionRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_HasOnBeginTransactionRequested
//go:noescape
func HasFuncHasOnBeginTransactionRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_HasOnBeginTransactionRequested
//go:noescape
func FuncHasOnBeginTransactionRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_HasOnBeginTransactionRequested
//go:noescape
func CallHasOnBeginTransactionRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_HasOnBeginTransactionRequested
//go:noescape
func TryHasOnBeginTransactionRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OnCancelRequested
//go:noescape
func HasFuncOnCancelRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OnCancelRequested
//go:noescape
func FuncOnCancelRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OnCancelRequested
//go:noescape
func CallOnCancelRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OnCancelRequested
//go:noescape
func TryOnCancelRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OffCancelRequested
//go:noescape
func HasFuncOffCancelRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OffCancelRequested
//go:noescape
func FuncOffCancelRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OffCancelRequested
//go:noescape
func CallOffCancelRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OffCancelRequested
//go:noescape
func TryOffCancelRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_HasOnCancelRequested
//go:noescape
func HasFuncHasOnCancelRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_HasOnCancelRequested
//go:noescape
func FuncHasOnCancelRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_HasOnCancelRequested
//go:noescape
func CallHasOnCancelRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_HasOnCancelRequested
//go:noescape
func TryHasOnCancelRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OnConnectRequested
//go:noescape
func HasFuncOnConnectRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OnConnectRequested
//go:noescape
func FuncOnConnectRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OnConnectRequested
//go:noescape
func CallOnConnectRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OnConnectRequested
//go:noescape
func TryOnConnectRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OffConnectRequested
//go:noescape
func HasFuncOffConnectRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OffConnectRequested
//go:noescape
func FuncOffConnectRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OffConnectRequested
//go:noescape
func CallOffConnectRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OffConnectRequested
//go:noescape
func TryOffConnectRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_HasOnConnectRequested
//go:noescape
func HasFuncHasOnConnectRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_HasOnConnectRequested
//go:noescape
func FuncHasOnConnectRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_HasOnConnectRequested
//go:noescape
func CallHasOnConnectRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_HasOnConnectRequested
//go:noescape
func TryHasOnConnectRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OnControlRequested
//go:noescape
func HasFuncOnControlRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OnControlRequested
//go:noescape
func FuncOnControlRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OnControlRequested
//go:noescape
func CallOnControlRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OnControlRequested
//go:noescape
func TryOnControlRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OffControlRequested
//go:noescape
func HasFuncOffControlRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OffControlRequested
//go:noescape
func FuncOffControlRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OffControlRequested
//go:noescape
func CallOffControlRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OffControlRequested
//go:noescape
func TryOffControlRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_HasOnControlRequested
//go:noescape
func HasFuncHasOnControlRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_HasOnControlRequested
//go:noescape
func FuncHasOnControlRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_HasOnControlRequested
//go:noescape
func CallHasOnControlRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_HasOnControlRequested
//go:noescape
func TryHasOnControlRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OnDisconnectRequested
//go:noescape
func HasFuncOnDisconnectRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OnDisconnectRequested
//go:noescape
func FuncOnDisconnectRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OnDisconnectRequested
//go:noescape
func CallOnDisconnectRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OnDisconnectRequested
//go:noescape
func TryOnDisconnectRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OffDisconnectRequested
//go:noescape
func HasFuncOffDisconnectRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OffDisconnectRequested
//go:noescape
func FuncOffDisconnectRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OffDisconnectRequested
//go:noescape
func CallOffDisconnectRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OffDisconnectRequested
//go:noescape
func TryOffDisconnectRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_HasOnDisconnectRequested
//go:noescape
func HasFuncHasOnDisconnectRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_HasOnDisconnectRequested
//go:noescape
func FuncHasOnDisconnectRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_HasOnDisconnectRequested
//go:noescape
func CallHasOnDisconnectRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_HasOnDisconnectRequested
//go:noescape
func TryHasOnDisconnectRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OnEndTransactionRequested
//go:noescape
func HasFuncOnEndTransactionRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OnEndTransactionRequested
//go:noescape
func FuncOnEndTransactionRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OnEndTransactionRequested
//go:noescape
func CallOnEndTransactionRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OnEndTransactionRequested
//go:noescape
func TryOnEndTransactionRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OffEndTransactionRequested
//go:noescape
func HasFuncOffEndTransactionRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OffEndTransactionRequested
//go:noescape
func FuncOffEndTransactionRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OffEndTransactionRequested
//go:noescape
func CallOffEndTransactionRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OffEndTransactionRequested
//go:noescape
func TryOffEndTransactionRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_HasOnEndTransactionRequested
//go:noescape
func HasFuncHasOnEndTransactionRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_HasOnEndTransactionRequested
//go:noescape
func FuncHasOnEndTransactionRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_HasOnEndTransactionRequested
//go:noescape
func CallHasOnEndTransactionRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_HasOnEndTransactionRequested
//go:noescape
func TryHasOnEndTransactionRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OnEstablishContextRequested
//go:noescape
func HasFuncOnEstablishContextRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OnEstablishContextRequested
//go:noescape
func FuncOnEstablishContextRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OnEstablishContextRequested
//go:noescape
func CallOnEstablishContextRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OnEstablishContextRequested
//go:noescape
func TryOnEstablishContextRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OffEstablishContextRequested
//go:noescape
func HasFuncOffEstablishContextRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OffEstablishContextRequested
//go:noescape
func FuncOffEstablishContextRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OffEstablishContextRequested
//go:noescape
func CallOffEstablishContextRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OffEstablishContextRequested
//go:noescape
func TryOffEstablishContextRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_HasOnEstablishContextRequested
//go:noescape
func HasFuncHasOnEstablishContextRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_HasOnEstablishContextRequested
//go:noescape
func FuncHasOnEstablishContextRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_HasOnEstablishContextRequested
//go:noescape
func CallHasOnEstablishContextRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_HasOnEstablishContextRequested
//go:noescape
func TryHasOnEstablishContextRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OnGetAttribRequested
//go:noescape
func HasFuncOnGetAttribRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OnGetAttribRequested
//go:noescape
func FuncOnGetAttribRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OnGetAttribRequested
//go:noescape
func CallOnGetAttribRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OnGetAttribRequested
//go:noescape
func TryOnGetAttribRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OffGetAttribRequested
//go:noescape
func HasFuncOffGetAttribRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OffGetAttribRequested
//go:noescape
func FuncOffGetAttribRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OffGetAttribRequested
//go:noescape
func CallOffGetAttribRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OffGetAttribRequested
//go:noescape
func TryOffGetAttribRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_HasOnGetAttribRequested
//go:noescape
func HasFuncHasOnGetAttribRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_HasOnGetAttribRequested
//go:noescape
func FuncHasOnGetAttribRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_HasOnGetAttribRequested
//go:noescape
func CallHasOnGetAttribRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_HasOnGetAttribRequested
//go:noescape
func TryHasOnGetAttribRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OnGetStatusChangeRequested
//go:noescape
func HasFuncOnGetStatusChangeRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OnGetStatusChangeRequested
//go:noescape
func FuncOnGetStatusChangeRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OnGetStatusChangeRequested
//go:noescape
func CallOnGetStatusChangeRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OnGetStatusChangeRequested
//go:noescape
func TryOnGetStatusChangeRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OffGetStatusChangeRequested
//go:noescape
func HasFuncOffGetStatusChangeRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OffGetStatusChangeRequested
//go:noescape
func FuncOffGetStatusChangeRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OffGetStatusChangeRequested
//go:noescape
func CallOffGetStatusChangeRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OffGetStatusChangeRequested
//go:noescape
func TryOffGetStatusChangeRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_HasOnGetStatusChangeRequested
//go:noescape
func HasFuncHasOnGetStatusChangeRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_HasOnGetStatusChangeRequested
//go:noescape
func FuncHasOnGetStatusChangeRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_HasOnGetStatusChangeRequested
//go:noescape
func CallHasOnGetStatusChangeRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_HasOnGetStatusChangeRequested
//go:noescape
func TryHasOnGetStatusChangeRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OnListReadersRequested
//go:noescape
func HasFuncOnListReadersRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OnListReadersRequested
//go:noescape
func FuncOnListReadersRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OnListReadersRequested
//go:noescape
func CallOnListReadersRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OnListReadersRequested
//go:noescape
func TryOnListReadersRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OffListReadersRequested
//go:noescape
func HasFuncOffListReadersRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OffListReadersRequested
//go:noescape
func FuncOffListReadersRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OffListReadersRequested
//go:noescape
func CallOffListReadersRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OffListReadersRequested
//go:noescape
func TryOffListReadersRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_HasOnListReadersRequested
//go:noescape
func HasFuncHasOnListReadersRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_HasOnListReadersRequested
//go:noescape
func FuncHasOnListReadersRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_HasOnListReadersRequested
//go:noescape
func CallHasOnListReadersRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_HasOnListReadersRequested
//go:noescape
func TryHasOnListReadersRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OnReleaseContextRequested
//go:noescape
func HasFuncOnReleaseContextRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OnReleaseContextRequested
//go:noescape
func FuncOnReleaseContextRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OnReleaseContextRequested
//go:noescape
func CallOnReleaseContextRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OnReleaseContextRequested
//go:noescape
func TryOnReleaseContextRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OffReleaseContextRequested
//go:noescape
func HasFuncOffReleaseContextRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OffReleaseContextRequested
//go:noescape
func FuncOffReleaseContextRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OffReleaseContextRequested
//go:noescape
func CallOffReleaseContextRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OffReleaseContextRequested
//go:noescape
func TryOffReleaseContextRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_HasOnReleaseContextRequested
//go:noescape
func HasFuncHasOnReleaseContextRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_HasOnReleaseContextRequested
//go:noescape
func FuncHasOnReleaseContextRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_HasOnReleaseContextRequested
//go:noescape
func CallHasOnReleaseContextRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_HasOnReleaseContextRequested
//go:noescape
func TryHasOnReleaseContextRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OnSetAttribRequested
//go:noescape
func HasFuncOnSetAttribRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OnSetAttribRequested
//go:noescape
func FuncOnSetAttribRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OnSetAttribRequested
//go:noescape
func CallOnSetAttribRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OnSetAttribRequested
//go:noescape
func TryOnSetAttribRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OffSetAttribRequested
//go:noescape
func HasFuncOffSetAttribRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OffSetAttribRequested
//go:noescape
func FuncOffSetAttribRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OffSetAttribRequested
//go:noescape
func CallOffSetAttribRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OffSetAttribRequested
//go:noescape
func TryOffSetAttribRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_HasOnSetAttribRequested
//go:noescape
func HasFuncHasOnSetAttribRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_HasOnSetAttribRequested
//go:noescape
func FuncHasOnSetAttribRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_HasOnSetAttribRequested
//go:noescape
func CallHasOnSetAttribRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_HasOnSetAttribRequested
//go:noescape
func TryHasOnSetAttribRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OnStatusRequested
//go:noescape
func HasFuncOnStatusRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OnStatusRequested
//go:noescape
func FuncOnStatusRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OnStatusRequested
//go:noescape
func CallOnStatusRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OnStatusRequested
//go:noescape
func TryOnStatusRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OffStatusRequested
//go:noescape
func HasFuncOffStatusRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OffStatusRequested
//go:noescape
func FuncOffStatusRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OffStatusRequested
//go:noescape
func CallOffStatusRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OffStatusRequested
//go:noescape
func TryOffStatusRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_HasOnStatusRequested
//go:noescape
func HasFuncHasOnStatusRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_HasOnStatusRequested
//go:noescape
func FuncHasOnStatusRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_HasOnStatusRequested
//go:noescape
func CallHasOnStatusRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_HasOnStatusRequested
//go:noescape
func TryHasOnStatusRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OnTransmitRequested
//go:noescape
func HasFuncOnTransmitRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OnTransmitRequested
//go:noescape
func FuncOnTransmitRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OnTransmitRequested
//go:noescape
func CallOnTransmitRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OnTransmitRequested
//go:noescape
func TryOnTransmitRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_OffTransmitRequested
//go:noescape
func HasFuncOffTransmitRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_OffTransmitRequested
//go:noescape
func FuncOffTransmitRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_OffTransmitRequested
//go:noescape
func CallOffTransmitRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_OffTransmitRequested
//go:noescape
func TryOffTransmitRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_HasOnTransmitRequested
//go:noescape
func HasFuncHasOnTransmitRequested() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_HasOnTransmitRequested
//go:noescape
func FuncHasOnTransmitRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_HasOnTransmitRequested
//go:noescape
func CallHasOnTransmitRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_HasOnTransmitRequested
//go:noescape
func TryHasOnTransmitRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_ReportConnectResult
//go:noescape
func HasFuncReportConnectResult() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_ReportConnectResult
//go:noescape
func FuncReportConnectResult(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_ReportConnectResult
//go:noescape
func CallReportConnectResult(
	retPtr unsafe.Pointer,
	requestId int32,
	scardHandle int32,
	activeProtocol uint32,
	resultCode uint32)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_ReportConnectResult
//go:noescape
func TryReportConnectResult(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	requestId int32,
	scardHandle int32,
	activeProtocol uint32,
	resultCode uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_ReportDataResult
//go:noescape
func HasFuncReportDataResult() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_ReportDataResult
//go:noescape
func FuncReportDataResult(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_ReportDataResult
//go:noescape
func CallReportDataResult(
	retPtr unsafe.Pointer,
	requestId int32,
	data js.Ref,
	resultCode uint32)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_ReportDataResult
//go:noescape
func TryReportDataResult(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	requestId int32,
	data js.Ref,
	resultCode uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_ReportEstablishContextResult
//go:noescape
func HasFuncReportEstablishContextResult() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_ReportEstablishContextResult
//go:noescape
func FuncReportEstablishContextResult(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_ReportEstablishContextResult
//go:noescape
func CallReportEstablishContextResult(
	retPtr unsafe.Pointer,
	requestId int32,
	scardContext int32,
	resultCode uint32)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_ReportEstablishContextResult
//go:noescape
func TryReportEstablishContextResult(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	requestId int32,
	scardContext int32,
	resultCode uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_ReportGetStatusChangeResult
//go:noescape
func HasFuncReportGetStatusChangeResult() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_ReportGetStatusChangeResult
//go:noescape
func FuncReportGetStatusChangeResult(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_ReportGetStatusChangeResult
//go:noescape
func CallReportGetStatusChangeResult(
	retPtr unsafe.Pointer,
	requestId int32,
	readerStates js.Ref,
	resultCode uint32)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_ReportGetStatusChangeResult
//go:noescape
func TryReportGetStatusChangeResult(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	requestId int32,
	readerStates js.Ref,
	resultCode uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_ReportListReadersResult
//go:noescape
func HasFuncReportListReadersResult() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_ReportListReadersResult
//go:noescape
func FuncReportListReadersResult(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_ReportListReadersResult
//go:noescape
func CallReportListReadersResult(
	retPtr unsafe.Pointer,
	requestId int32,
	readers js.Ref,
	resultCode uint32)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_ReportListReadersResult
//go:noescape
func TryReportListReadersResult(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	requestId int32,
	readers js.Ref,
	resultCode uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_ReportPlainResult
//go:noescape
func HasFuncReportPlainResult() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_ReportPlainResult
//go:noescape
func FuncReportPlainResult(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_ReportPlainResult
//go:noescape
func CallReportPlainResult(
	retPtr unsafe.Pointer,
	requestId int32,
	resultCode uint32)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_ReportPlainResult
//go:noescape
func TryReportPlainResult(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	requestId int32,
	resultCode uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_ReportReleaseContextResult
//go:noescape
func HasFuncReportReleaseContextResult() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_ReportReleaseContextResult
//go:noescape
func FuncReportReleaseContextResult(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_ReportReleaseContextResult
//go:noescape
func CallReportReleaseContextResult(
	retPtr unsafe.Pointer,
	requestId int32,
	resultCode uint32)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_ReportReleaseContextResult
//go:noescape
func TryReportReleaseContextResult(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	requestId int32,
	resultCode uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/smartcardproviderprivate has_ReportStatusResult
//go:noescape
func HasFuncReportStatusResult() js.Ref

//go:wasmimport plat/js/webext/smartcardproviderprivate func_ReportStatusResult
//go:noescape
func FuncReportStatusResult(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/smartcardproviderprivate call_ReportStatusResult
//go:noescape
func CallReportStatusResult(
	retPtr unsafe.Pointer,
	requestId int32,
	readerName js.Ref,
	state uint32,
	protocol uint32,
	atr js.Ref,
	resultCode uint32)

//go:wasmimport plat/js/webext/smartcardproviderprivate try_ReportStatusResult
//go:noescape
func TryReportStatusResult(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	requestId int32,
	readerName js.Ref,
	state uint32,
	protocol uint32,
	atr js.Ref,
	resultCode uint32) (ok js.Ref)
