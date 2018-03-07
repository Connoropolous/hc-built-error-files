//go:cgo_ldflag "-target"
//go:cgo_ldflag "armv7a-none-linux-androideabi"
//go:cgo_ldflag "-gcc-toolchain"
//go:cgo_ldflag "/Users/connor/Library/Android/sdk/ndk-bundle/toolchains/arm-linux-androideabi-4.9/prebuilt/darwin-x86_64"
//go:cgo_ldflag "--sysroot"
//go:cgo_ldflag "/Users/connor/Library/Android/sdk/ndk-bundle/platforms/android-15/arch-arm"
//go:cgo_ldflag "-L/Users/connor/go/pkg/gomobile/lib/arm"
//go:cgo_ldflag "-llog"
// Created by cgo - DO NOT EDIT

package gomobile_bind

import "unsafe"

import _ "runtime/cgo"

import "syscall"

var _ syscall.Errno
func _Cgo_ptr(ptr unsafe.Pointer) unsafe.Pointer { return ptr }

//go:linkname _Cgo_always_false runtime.cgoAlwaysFalse
var _Cgo_always_false bool
//go:linkname _Cgo_use runtime.cgoUse
func _Cgo_use(interface{})
type _Ctype___int32_t _Ctype_int

type _Ctype___int64_t _Ctype_longlong

type _Ctype___int8_t _Ctype_schar

type _Ctype_char uint8

type _Ctype_int int32

type _Ctype_int32_t _Ctype___int32_t

type _Ctype_int64_t _Ctype___int64_t

type _Ctype_int8_t _Ctype___int8_t

type _Ctype_jbyte _Ctype_int8_t

type _Ctype_jint _Ctype_int32_t

type _Ctype_jlong _Ctype_int64_t

type _Ctype_jsize _Ctype_jint

type _Ctype_longlong int64

type _Ctype_nbyteslice _Ctype_struct_nbyteslice

type _Ctype_nint _Ctype_jlong

type _Ctype_nstring _Ctype_struct_nstring

type _Ctype_schar int8

type _Ctype_size_t _Ctype_uint

type _Ctype_struct_cproxyholochain_Entry_Marshal_return struct {
	r0	_Ctype_struct_nbyteslice
	r1	_Ctype_int32_t
}

type _Ctype_struct_cproxyholochain_Revocation_Marshal_return struct {
	r0	_Ctype_struct_nbyteslice
	r1	_Ctype_int32_t
}

type _Ctype_struct_cproxyholochain_Ribosome_Receive_return struct {
	r0	_Ctype_struct_nstring
	r1	_Ctype_int32_t
}

type _Ctype_struct_cproxyholochain_Warrant_Encode_return struct {
	r0	_Ctype_struct_nbyteslice
	r1	_Ctype_int32_t
}

type _Ctype_struct_nbyteslice struct {
	ptr	unsafe.Pointer
	len	_Ctype_jsize
}

type _Ctype_struct_nstring struct {
	chars	unsafe.Pointer
	len	_Ctype_jsize
}

type _Ctype_uint uint32

type _Ctype_void [0]byte

//go:linkname _cgo_runtime_cgocall runtime.cgocall
func _cgo_runtime_cgocall(unsafe.Pointer, uintptr) int32

//go:linkname _cgo_runtime_cgocallback runtime.cgocallback
func _cgo_runtime_cgocallback(unsafe.Pointer, unsafe.Pointer, uintptr, uintptr)

//go:linkname _cgoCheckPointer runtime.cgoCheckPointer
func _cgoCheckPointer(interface{}, ...interface{})

//go:linkname _cgoCheckResult runtime.cgoCheckResult
func _cgoCheckResult(interface{})


//go:linkname _cgo_runtime_gobytes runtime.gobytes
func _cgo_runtime_gobytes(unsafe.Pointer, int) []byte

func _Cfunc_GoBytes(p unsafe.Pointer, l _Ctype_int) []byte {
	return _cgo_runtime_gobytes(p, int(l))
}

func _Cfunc__CMalloc(n _Ctype_size_t) unsafe.Pointer {
	return _cgo_cmalloc(uint64(n))
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_cproxy_error_Error
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_cproxy_error_Error _cgo_ebd397278953_Cfunc_cproxy_error_Error
var __cgofn__cgo_ebd397278953_Cfunc_cproxy_error_Error byte
var _cgo_ebd397278953_Cfunc_cproxy_error_Error = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_cproxy_error_Error)

//go:cgo_unsafe_args
func _Cfunc_cproxy_error_Error(p0 _Ctype_int32_t) (r1 _Ctype_struct_nstring) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_cproxy_error_Error, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_cproxyholochain_Action_Name
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Action_Name _cgo_ebd397278953_Cfunc_cproxyholochain_Action_Name
var __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Action_Name byte
var _cgo_ebd397278953_Cfunc_cproxyholochain_Action_Name = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Action_Name)

//go:cgo_unsafe_args
func _Cfunc_cproxyholochain_Action_Name(p0 _Ctype_int32_t) (r1 _Ctype_struct_nstring) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_cproxyholochain_Action_Name, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_CheckValidationRequest
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_CheckValidationRequest _cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_CheckValidationRequest
var __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_CheckValidationRequest byte
var _cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_CheckValidationRequest = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_CheckValidationRequest)

//go:cgo_unsafe_args
func _Cfunc_cproxyholochain_CommittingAction_CheckValidationRequest(p0 _Ctype_int32_t, p1 _Ctype_int32_t) (r1 _Ctype_int32_t) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_CheckValidationRequest, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_Entry
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_Entry _cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_Entry
var __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_Entry byte
var _cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_Entry = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_Entry)

//go:cgo_unsafe_args
func _Cfunc_cproxyholochain_CommittingAction_Entry(p0 _Ctype_int32_t) (r1 _Ctype_int32_t) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_Entry, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_EntryType
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_EntryType _cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_EntryType
var __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_EntryType byte
var _cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_EntryType = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_EntryType)

//go:cgo_unsafe_args
func _Cfunc_cproxyholochain_CommittingAction_EntryType(p0 _Ctype_int32_t) (r1 _Ctype_struct_nstring) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_EntryType, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_Name
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_Name _cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_Name
var __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_Name byte
var _cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_Name = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_Name)

//go:cgo_unsafe_args
func _Cfunc_cproxyholochain_CommittingAction_Name(p0 _Ctype_int32_t) (r1 _Ctype_struct_nstring) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_Name, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_SetHeader
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_SetHeader _cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_SetHeader
var __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_SetHeader byte
var _cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_SetHeader = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_SetHeader)

//go:cgo_unsafe_args
func _Cfunc_cproxyholochain_CommittingAction_SetHeader(p0 _Ctype_int32_t, p1 _Ctype_int32_t) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_SetHeader, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_cproxyholochain_Entry_Marshal
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Entry_Marshal _cgo_ebd397278953_Cfunc_cproxyholochain_Entry_Marshal
var __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Entry_Marshal byte
var _cgo_ebd397278953_Cfunc_cproxyholochain_Entry_Marshal = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Entry_Marshal)

//go:cgo_unsafe_args
func _Cfunc_cproxyholochain_Entry_Marshal(p0 _Ctype_int32_t) (r1 _Ctype_struct_cproxyholochain_Entry_Marshal_return) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_cproxyholochain_Entry_Marshal, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_cproxyholochain_Entry_Unmarshal
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Entry_Unmarshal _cgo_ebd397278953_Cfunc_cproxyholochain_Entry_Unmarshal
var __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Entry_Unmarshal byte
var _cgo_ebd397278953_Cfunc_cproxyholochain_Entry_Unmarshal = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Entry_Unmarshal)

//go:cgo_unsafe_args
func _Cfunc_cproxyholochain_Entry_Unmarshal(p0 _Ctype_int32_t, p1 _Ctype_struct_nbyteslice) (r1 _Ctype_int32_t) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_cproxyholochain_Entry_Unmarshal, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Marshal
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Marshal _cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Marshal
var __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Marshal byte
var _cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Marshal = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Marshal)

//go:cgo_unsafe_args
func _Cfunc_cproxyholochain_Revocation_Marshal(p0 _Ctype_int32_t) (r1 _Ctype_struct_cproxyholochain_Revocation_Marshal_return) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Marshal, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Unmarshal
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Unmarshal _cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Unmarshal
var __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Unmarshal byte
var _cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Unmarshal = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Unmarshal)

//go:cgo_unsafe_args
func _Cfunc_cproxyholochain_Revocation_Unmarshal(p0 _Ctype_int32_t, p1 _Ctype_struct_nbyteslice) (r1 _Ctype_int32_t) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Unmarshal, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Verify
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Verify _cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Verify
var __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Verify byte
var _cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Verify = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Verify)

//go:cgo_unsafe_args
func _Cfunc_cproxyholochain_Revocation_Verify(p0 _Ctype_int32_t) (r1 _Ctype_int32_t) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Verify, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_ChainGenesis
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_ChainGenesis _cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_ChainGenesis
var __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_ChainGenesis byte
var _cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_ChainGenesis = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_ChainGenesis)

//go:cgo_unsafe_args
func _Cfunc_cproxyholochain_Ribosome_ChainGenesis(p0 _Ctype_int32_t) (r1 _Ctype_int32_t) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_ChainGenesis, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_Receive
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_Receive _cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_Receive
var __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_Receive byte
var _cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_Receive = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_Receive)

//go:cgo_unsafe_args
func _Cfunc_cproxyholochain_Ribosome_Receive(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring, p2 _Ctype_struct_nstring) (r1 _Ctype_struct_cproxyholochain_Ribosome_Receive_return) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_Receive, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
		_Cgo_use(p2)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_Type
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_Type _cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_Type
var __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_Type byte
var _cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_Type = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_Type)

//go:cgo_unsafe_args
func _Cfunc_cproxyholochain_Ribosome_Type(p0 _Ctype_int32_t) (r1 _Ctype_struct_nstring) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_Type, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_cproxyholochain_ValidatingAction_CheckValidationRequest
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_ValidatingAction_CheckValidationRequest _cgo_ebd397278953_Cfunc_cproxyholochain_ValidatingAction_CheckValidationRequest
var __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_ValidatingAction_CheckValidationRequest byte
var _cgo_ebd397278953_Cfunc_cproxyholochain_ValidatingAction_CheckValidationRequest = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_ValidatingAction_CheckValidationRequest)

//go:cgo_unsafe_args
func _Cfunc_cproxyholochain_ValidatingAction_CheckValidationRequest(p0 _Ctype_int32_t, p1 _Ctype_int32_t) (r1 _Ctype_int32_t) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_cproxyholochain_ValidatingAction_CheckValidationRequest, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_cproxyholochain_ValidatingAction_Name
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_ValidatingAction_Name _cgo_ebd397278953_Cfunc_cproxyholochain_ValidatingAction_Name
var __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_ValidatingAction_Name byte
var _cgo_ebd397278953_Cfunc_cproxyholochain_ValidatingAction_Name = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_ValidatingAction_Name)

//go:cgo_unsafe_args
func _Cfunc_cproxyholochain_ValidatingAction_Name(p0 _Ctype_int32_t) (r1 _Ctype_struct_nstring) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_cproxyholochain_ValidatingAction_Name, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Decode
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Decode _cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Decode
var __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Decode byte
var _cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Decode = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Decode)

//go:cgo_unsafe_args
func _Cfunc_cproxyholochain_Warrant_Decode(p0 _Ctype_int32_t, p1 _Ctype_struct_nbyteslice) (r1 _Ctype_int32_t) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Decode, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Encode
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Encode _cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Encode
var __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Encode byte
var _cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Encode = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Encode)

//go:cgo_unsafe_args
func _Cfunc_cproxyholochain_Warrant_Encode(p0 _Ctype_int32_t) (r1 _Ctype_struct_cproxyholochain_Warrant_Encode_return) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Encode, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Type
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Type _cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Type
var __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Type byte
var _cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Type = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Type)

//go:cgo_unsafe_args
func _Cfunc_cproxyholochain_Warrant_Type(p0 _Ctype_int32_t) (r1 _Ctype_nint) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Type, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Verify
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Verify _cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Verify
var __cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Verify byte
var _cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Verify = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Verify)

//go:cgo_unsafe_args
func _Cfunc_cproxyholochain_Warrant_Verify(p0 _Ctype_int32_t, p1 _Ctype_int32_t) (r1 _Ctype_int32_t) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Verify, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_free
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_free _cgo_ebd397278953_Cfunc_free
var __cgofn__cgo_ebd397278953_Cfunc_free byte
var _cgo_ebd397278953_Cfunc_free = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_free)

//go:cgo_unsafe_args
func _Cfunc_free(p0 unsafe.Pointer) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_free, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_go_seq_dec_ref
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_go_seq_dec_ref _cgo_ebd397278953_Cfunc_go_seq_dec_ref
var __cgofn__cgo_ebd397278953_Cfunc_go_seq_dec_ref byte
var _cgo_ebd397278953_Cfunc_go_seq_dec_ref = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_go_seq_dec_ref)

//go:cgo_unsafe_args
func _Cfunc_go_seq_dec_ref(p0 _Ctype_int32_t) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_go_seq_dec_ref, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_go_seq_inc_ref
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_go_seq_inc_ref _cgo_ebd397278953_Cfunc_go_seq_inc_ref
var __cgofn__cgo_ebd397278953_Cfunc_go_seq_inc_ref byte
var _cgo_ebd397278953_Cfunc_go_seq_inc_ref = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_go_seq_inc_ref)

//go:cgo_unsafe_args
func _Cfunc_go_seq_inc_ref(p0 _Ctype_int32_t) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_go_seq_inc_ref, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_ebd397278953_Cfunc_init_proxies
//go:linkname __cgofn__cgo_ebd397278953_Cfunc_init_proxies _cgo_ebd397278953_Cfunc_init_proxies
var __cgofn__cgo_ebd397278953_Cfunc_init_proxies byte
var _cgo_ebd397278953_Cfunc_init_proxies = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc_init_proxies)

//go:cgo_unsafe_args
func _Cfunc_init_proxies() (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc_init_proxies, uintptr(unsafe.Pointer(&r1)))
	if _Cgo_always_false {
	}
	return
}
//go:cgo_export_dynamic initClasses
//go:linkname _cgoexp_ebd397278953_initClasses _cgoexp_ebd397278953_initClasses
//go:cgo_export_static _cgoexp_ebd397278953_initClasses
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_initClasses(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_initClasses
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_initClasses() {
	initClasses()
}
//go:cgo_export_dynamic proxyholochain_ActionBridge_Name
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionBridge_Name _cgoexp_ebd397278953_proxyholochain_ActionBridge_Name
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionBridge_Name
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionBridge_Name(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionBridge_Name
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionBridge_Name(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ActionBridge_Name(p0)
}
//go:cgo_export_dynamic new_holochain_ActionBridge
//go:linkname _cgoexp_ebd397278953_new_holochain_ActionBridge _cgoexp_ebd397278953_new_holochain_ActionBridge
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ActionBridge
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ActionBridge(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ActionBridge
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ActionBridge() (r0 _Ctype_int32_t) {
	return new_holochain_ActionBridge()
}
//go:cgo_export_dynamic proxyholochain_ActionCall_Name
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionCall_Name _cgoexp_ebd397278953_proxyholochain_ActionCall_Name
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionCall_Name
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionCall_Name(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionCall_Name
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionCall_Name(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ActionCall_Name(p0)
}
//go:cgo_export_dynamic new_holochain_ActionCall
//go:linkname _cgoexp_ebd397278953_new_holochain_ActionCall _cgoexp_ebd397278953_new_holochain_ActionCall
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ActionCall
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ActionCall(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ActionCall
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ActionCall() (r0 _Ctype_int32_t) {
	return new_holochain_ActionCall()
}
//go:cgo_export_dynamic proxyholochain_ActionCommit_CheckValidationRequest
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionCommit_CheckValidationRequest _cgoexp_ebd397278953_proxyholochain_ActionCommit_CheckValidationRequest
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionCommit_CheckValidationRequest
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionCommit_CheckValidationRequest(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionCommit_CheckValidationRequest
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionCommit_CheckValidationRequest(p0 _Ctype_int32_t, p1 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_ActionCommit_CheckValidationRequest(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ActionCommit_Entry
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionCommit_Entry _cgoexp_ebd397278953_proxyholochain_ActionCommit_Entry
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionCommit_Entry
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionCommit_Entry(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionCommit_Entry
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionCommit_Entry(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_ActionCommit_Entry(p0)
}
//go:cgo_export_dynamic proxyholochain_ActionCommit_EntryType
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionCommit_EntryType _cgoexp_ebd397278953_proxyholochain_ActionCommit_EntryType
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionCommit_EntryType
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionCommit_EntryType(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionCommit_EntryType
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionCommit_EntryType(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ActionCommit_EntryType(p0)
}
//go:cgo_export_dynamic proxyholochain_ActionCommit_Name
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionCommit_Name _cgoexp_ebd397278953_proxyholochain_ActionCommit_Name
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionCommit_Name
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionCommit_Name(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionCommit_Name
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionCommit_Name(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ActionCommit_Name(p0)
}
//go:cgo_export_dynamic proxyholochain_ActionCommit_SetHeader
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionCommit_SetHeader _cgoexp_ebd397278953_proxyholochain_ActionCommit_SetHeader
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionCommit_SetHeader
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionCommit_SetHeader(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionCommit_SetHeader
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionCommit_SetHeader(p0 _Ctype_int32_t, p1 _Ctype_int32_t) {
	proxyholochain_ActionCommit_SetHeader(p0, p1)
}
//go:cgo_export_dynamic new_holochain_ActionCommit
//go:linkname _cgoexp_ebd397278953_new_holochain_ActionCommit _cgoexp_ebd397278953_new_holochain_ActionCommit
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ActionCommit
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ActionCommit(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ActionCommit
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ActionCommit() (r0 _Ctype_int32_t) {
	return new_holochain_ActionCommit()
}
//go:cgo_export_dynamic proxyholochain_ActionDebug_Name
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionDebug_Name _cgoexp_ebd397278953_proxyholochain_ActionDebug_Name
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionDebug_Name
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionDebug_Name(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionDebug_Name
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionDebug_Name(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ActionDebug_Name(p0)
}
//go:cgo_export_dynamic new_holochain_ActionDebug
//go:linkname _cgoexp_ebd397278953_new_holochain_ActionDebug _cgoexp_ebd397278953_new_holochain_ActionDebug
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ActionDebug
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ActionDebug(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ActionDebug
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ActionDebug() (r0 _Ctype_int32_t) {
	return new_holochain_ActionDebug()
}
//go:cgo_export_dynamic proxyholochain_ActionDel_CheckValidationRequest
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionDel_CheckValidationRequest _cgoexp_ebd397278953_proxyholochain_ActionDel_CheckValidationRequest
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionDel_CheckValidationRequest
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionDel_CheckValidationRequest(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionDel_CheckValidationRequest
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionDel_CheckValidationRequest(p0 _Ctype_int32_t, p1 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_ActionDel_CheckValidationRequest(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ActionDel_Entry
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionDel_Entry _cgoexp_ebd397278953_proxyholochain_ActionDel_Entry
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionDel_Entry
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionDel_Entry(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionDel_Entry
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionDel_Entry(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_ActionDel_Entry(p0)
}
//go:cgo_export_dynamic proxyholochain_ActionDel_EntryType
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionDel_EntryType _cgoexp_ebd397278953_proxyholochain_ActionDel_EntryType
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionDel_EntryType
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionDel_EntryType(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionDel_EntryType
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionDel_EntryType(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ActionDel_EntryType(p0)
}
//go:cgo_export_dynamic proxyholochain_ActionDel_Name
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionDel_Name _cgoexp_ebd397278953_proxyholochain_ActionDel_Name
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionDel_Name
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionDel_Name(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionDel_Name
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionDel_Name(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ActionDel_Name(p0)
}
//go:cgo_export_dynamic proxyholochain_ActionDel_SetHeader
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionDel_SetHeader _cgoexp_ebd397278953_proxyholochain_ActionDel_SetHeader
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionDel_SetHeader
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionDel_SetHeader(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionDel_SetHeader
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionDel_SetHeader(p0 _Ctype_int32_t, p1 _Ctype_int32_t) {
	proxyholochain_ActionDel_SetHeader(p0, p1)
}
//go:cgo_export_dynamic new_holochain_ActionDel
//go:linkname _cgoexp_ebd397278953_new_holochain_ActionDel _cgoexp_ebd397278953_new_holochain_ActionDel
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ActionDel
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ActionDel(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ActionDel
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ActionDel() (r0 _Ctype_int32_t) {
	return new_holochain_ActionDel()
}
//go:cgo_export_dynamic proxyholochain_ActionGet_Name
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionGet_Name _cgoexp_ebd397278953_proxyholochain_ActionGet_Name
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionGet_Name
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionGet_Name(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionGet_Name
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionGet_Name(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ActionGet_Name(p0)
}
//go:cgo_export_dynamic new_holochain_ActionGet
//go:linkname _cgoexp_ebd397278953_new_holochain_ActionGet _cgoexp_ebd397278953_new_holochain_ActionGet
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ActionGet
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ActionGet(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ActionGet
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ActionGet() (r0 _Ctype_int32_t) {
	return new_holochain_ActionGet()
}
//go:cgo_export_dynamic proxyholochain_ActionGetBridges_Name
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionGetBridges_Name _cgoexp_ebd397278953_proxyholochain_ActionGetBridges_Name
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionGetBridges_Name
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionGetBridges_Name(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionGetBridges_Name
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionGetBridges_Name(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ActionGetBridges_Name(p0)
}
//go:cgo_export_dynamic new_holochain_ActionGetBridges
//go:linkname _cgoexp_ebd397278953_new_holochain_ActionGetBridges _cgoexp_ebd397278953_new_holochain_ActionGetBridges
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ActionGetBridges
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ActionGetBridges(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ActionGetBridges
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ActionGetBridges() (r0 _Ctype_int32_t) {
	return new_holochain_ActionGetBridges()
}
//go:cgo_export_dynamic proxyholochain_ActionGetLinks_Name
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionGetLinks_Name _cgoexp_ebd397278953_proxyholochain_ActionGetLinks_Name
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionGetLinks_Name
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionGetLinks_Name(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionGetLinks_Name
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionGetLinks_Name(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ActionGetLinks_Name(p0)
}
//go:cgo_export_dynamic new_holochain_ActionGetLinks
//go:linkname _cgoexp_ebd397278953_new_holochain_ActionGetLinks _cgoexp_ebd397278953_new_holochain_ActionGetLinks
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ActionGetLinks
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ActionGetLinks(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ActionGetLinks
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ActionGetLinks() (r0 _Ctype_int32_t) {
	return new_holochain_ActionGetLinks()
}
//go:cgo_export_dynamic proxyholochain_ActionLink_CheckValidationRequest
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionLink_CheckValidationRequest _cgoexp_ebd397278953_proxyholochain_ActionLink_CheckValidationRequest
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionLink_CheckValidationRequest
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionLink_CheckValidationRequest(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionLink_CheckValidationRequest
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionLink_CheckValidationRequest(p0 _Ctype_int32_t, p1 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_ActionLink_CheckValidationRequest(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ActionLink_Name
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionLink_Name _cgoexp_ebd397278953_proxyholochain_ActionLink_Name
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionLink_Name
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionLink_Name(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionLink_Name
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionLink_Name(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ActionLink_Name(p0)
}
//go:cgo_export_dynamic new_holochain_ActionLink
//go:linkname _cgoexp_ebd397278953_new_holochain_ActionLink _cgoexp_ebd397278953_new_holochain_ActionLink
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ActionLink
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ActionLink(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ActionLink
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ActionLink() (r0 _Ctype_int32_t) {
	return new_holochain_ActionLink()
}
//go:cgo_export_dynamic proxyholochain_ActionListAdd_Name
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionListAdd_Name _cgoexp_ebd397278953_proxyholochain_ActionListAdd_Name
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionListAdd_Name
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionListAdd_Name(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionListAdd_Name
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionListAdd_Name(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ActionListAdd_Name(p0)
}
//go:cgo_export_dynamic new_holochain_ActionListAdd
//go:linkname _cgoexp_ebd397278953_new_holochain_ActionListAdd _cgoexp_ebd397278953_new_holochain_ActionListAdd
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ActionListAdd
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ActionListAdd(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ActionListAdd
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ActionListAdd() (r0 _Ctype_int32_t) {
	return new_holochain_ActionListAdd()
}
//go:cgo_export_dynamic proxyholochain_ActionMakeHash_Name
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionMakeHash_Name _cgoexp_ebd397278953_proxyholochain_ActionMakeHash_Name
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionMakeHash_Name
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionMakeHash_Name(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionMakeHash_Name
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionMakeHash_Name(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ActionMakeHash_Name(p0)
}
//go:cgo_export_dynamic new_holochain_ActionMakeHash
//go:linkname _cgoexp_ebd397278953_new_holochain_ActionMakeHash _cgoexp_ebd397278953_new_holochain_ActionMakeHash
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ActionMakeHash
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ActionMakeHash(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ActionMakeHash
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ActionMakeHash() (r0 _Ctype_int32_t) {
	return new_holochain_ActionMakeHash()
}
//go:cgo_export_dynamic proxyholochain_ActionMod_CheckValidationRequest
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionMod_CheckValidationRequest _cgoexp_ebd397278953_proxyholochain_ActionMod_CheckValidationRequest
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionMod_CheckValidationRequest
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionMod_CheckValidationRequest(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionMod_CheckValidationRequest
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionMod_CheckValidationRequest(p0 _Ctype_int32_t, p1 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_ActionMod_CheckValidationRequest(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ActionMod_Entry
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionMod_Entry _cgoexp_ebd397278953_proxyholochain_ActionMod_Entry
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionMod_Entry
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionMod_Entry(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionMod_Entry
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionMod_Entry(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_ActionMod_Entry(p0)
}
//go:cgo_export_dynamic proxyholochain_ActionMod_EntryType
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionMod_EntryType _cgoexp_ebd397278953_proxyholochain_ActionMod_EntryType
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionMod_EntryType
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionMod_EntryType(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionMod_EntryType
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionMod_EntryType(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ActionMod_EntryType(p0)
}
//go:cgo_export_dynamic proxyholochain_ActionMod_Name
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionMod_Name _cgoexp_ebd397278953_proxyholochain_ActionMod_Name
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionMod_Name
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionMod_Name(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionMod_Name
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionMod_Name(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ActionMod_Name(p0)
}
//go:cgo_export_dynamic proxyholochain_ActionMod_SetHeader
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionMod_SetHeader _cgoexp_ebd397278953_proxyholochain_ActionMod_SetHeader
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionMod_SetHeader
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionMod_SetHeader(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionMod_SetHeader
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionMod_SetHeader(p0 _Ctype_int32_t, p1 _Ctype_int32_t) {
	proxyholochain_ActionMod_SetHeader(p0, p1)
}
//go:cgo_export_dynamic new_holochain_ActionMod
//go:linkname _cgoexp_ebd397278953_new_holochain_ActionMod _cgoexp_ebd397278953_new_holochain_ActionMod
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ActionMod
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ActionMod(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ActionMod
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ActionMod() (r0 _Ctype_int32_t) {
	return new_holochain_ActionMod()
}
//go:cgo_export_dynamic proxyholochain_ActionModAgent_Revocation_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionModAgent_Revocation_Set _cgoexp_ebd397278953_proxyholochain_ActionModAgent_Revocation_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionModAgent_Revocation_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionModAgent_Revocation_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionModAgent_Revocation_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionModAgent_Revocation_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_ActionModAgent_Revocation_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ActionModAgent_Revocation_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionModAgent_Revocation_Get _cgoexp_ebd397278953_proxyholochain_ActionModAgent_Revocation_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionModAgent_Revocation_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionModAgent_Revocation_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionModAgent_Revocation_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionModAgent_Revocation_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ActionModAgent_Revocation_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_ActionModAgent_Name
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionModAgent_Name _cgoexp_ebd397278953_proxyholochain_ActionModAgent_Name
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionModAgent_Name
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionModAgent_Name(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionModAgent_Name
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionModAgent_Name(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ActionModAgent_Name(p0)
}
//go:cgo_export_dynamic new_holochain_ActionModAgent
//go:linkname _cgoexp_ebd397278953_new_holochain_ActionModAgent _cgoexp_ebd397278953_new_holochain_ActionModAgent
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ActionModAgent
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ActionModAgent(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ActionModAgent
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ActionModAgent() (r0 _Ctype_int32_t) {
	return new_holochain_ActionModAgent()
}
//go:cgo_export_dynamic proxyholochain_ActionProperty_Name
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionProperty_Name _cgoexp_ebd397278953_proxyholochain_ActionProperty_Name
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionProperty_Name
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionProperty_Name(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionProperty_Name
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionProperty_Name(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ActionProperty_Name(p0)
}
//go:cgo_export_dynamic new_holochain_ActionProperty
//go:linkname _cgoexp_ebd397278953_new_holochain_ActionProperty _cgoexp_ebd397278953_new_holochain_ActionProperty
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ActionProperty
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ActionProperty(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ActionProperty
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ActionProperty() (r0 _Ctype_int32_t) {
	return new_holochain_ActionProperty()
}
//go:cgo_export_dynamic proxyholochain_ActionPut_CheckValidationRequest
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionPut_CheckValidationRequest _cgoexp_ebd397278953_proxyholochain_ActionPut_CheckValidationRequest
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionPut_CheckValidationRequest
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionPut_CheckValidationRequest(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionPut_CheckValidationRequest
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionPut_CheckValidationRequest(p0 _Ctype_int32_t, p1 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_ActionPut_CheckValidationRequest(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ActionPut_Name
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionPut_Name _cgoexp_ebd397278953_proxyholochain_ActionPut_Name
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionPut_Name
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionPut_Name(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionPut_Name
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionPut_Name(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ActionPut_Name(p0)
}
//go:cgo_export_dynamic new_holochain_ActionPut
//go:linkname _cgoexp_ebd397278953_new_holochain_ActionPut _cgoexp_ebd397278953_new_holochain_ActionPut
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ActionPut
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ActionPut(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ActionPut
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ActionPut() (r0 _Ctype_int32_t) {
	return new_holochain_ActionPut()
}
//go:cgo_export_dynamic proxyholochain_ActionQuery_Name
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionQuery_Name _cgoexp_ebd397278953_proxyholochain_ActionQuery_Name
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionQuery_Name
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionQuery_Name(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionQuery_Name
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionQuery_Name(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ActionQuery_Name(p0)
}
//go:cgo_export_dynamic new_holochain_ActionQuery
//go:linkname _cgoexp_ebd397278953_new_holochain_ActionQuery _cgoexp_ebd397278953_new_holochain_ActionQuery
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ActionQuery
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ActionQuery(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ActionQuery
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ActionQuery() (r0 _Ctype_int32_t) {
	return new_holochain_ActionQuery()
}
//go:cgo_export_dynamic proxyholochain_ActionSend_Name
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionSend_Name _cgoexp_ebd397278953_proxyholochain_ActionSend_Name
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionSend_Name
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionSend_Name(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionSend_Name
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionSend_Name(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ActionSend_Name(p0)
}
//go:cgo_export_dynamic new_holochain_ActionSend
//go:linkname _cgoexp_ebd397278953_new_holochain_ActionSend _cgoexp_ebd397278953_new_holochain_ActionSend
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ActionSend
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ActionSend(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ActionSend
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ActionSend() (r0 _Ctype_int32_t) {
	return new_holochain_ActionSend()
}
//go:cgo_export_dynamic proxyholochain_ActionSign_Name
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionSign_Name _cgoexp_ebd397278953_proxyholochain_ActionSign_Name
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionSign_Name
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionSign_Name(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionSign_Name
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionSign_Name(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ActionSign_Name(p0)
}
//go:cgo_export_dynamic new_holochain_ActionSign
//go:linkname _cgoexp_ebd397278953_new_holochain_ActionSign _cgoexp_ebd397278953_new_holochain_ActionSign
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ActionSign
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ActionSign(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ActionSign
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ActionSign() (r0 _Ctype_int32_t) {
	return new_holochain_ActionSign()
}
//go:cgo_export_dynamic proxyholochain_ActionVerifySignature_Name
//go:linkname _cgoexp_ebd397278953_proxyholochain_ActionVerifySignature_Name _cgoexp_ebd397278953_proxyholochain_ActionVerifySignature_Name
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ActionVerifySignature_Name
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ActionVerifySignature_Name(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ActionVerifySignature_Name
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ActionVerifySignature_Name(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ActionVerifySignature_Name(p0)
}
//go:cgo_export_dynamic new_holochain_ActionVerifySignature
//go:linkname _cgoexp_ebd397278953_new_holochain_ActionVerifySignature _cgoexp_ebd397278953_new_holochain_ActionVerifySignature
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ActionVerifySignature
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ActionVerifySignature(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ActionVerifySignature
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ActionVerifySignature() (r0 _Ctype_int32_t) {
	return new_holochain_ActionVerifySignature()
}
//go:cgo_export_dynamic proxyholochain_AgentEntry_Revocation_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_AgentEntry_Revocation_Set _cgoexp_ebd397278953_proxyholochain_AgentEntry_Revocation_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AgentEntry_Revocation_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AgentEntry_Revocation_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AgentEntry_Revocation_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AgentEntry_Revocation_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nbyteslice) {
	proxyholochain_AgentEntry_Revocation_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_AgentEntry_Revocation_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_AgentEntry_Revocation_Get _cgoexp_ebd397278953_proxyholochain_AgentEntry_Revocation_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AgentEntry_Revocation_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AgentEntry_Revocation_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AgentEntry_Revocation_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AgentEntry_Revocation_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nbyteslice) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_AgentEntry_Revocation_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_AgentEntry_PublicKey_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_AgentEntry_PublicKey_Set _cgoexp_ebd397278953_proxyholochain_AgentEntry_PublicKey_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AgentEntry_PublicKey_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AgentEntry_PublicKey_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AgentEntry_PublicKey_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AgentEntry_PublicKey_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nbyteslice) {
	proxyholochain_AgentEntry_PublicKey_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_AgentEntry_PublicKey_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_AgentEntry_PublicKey_Get _cgoexp_ebd397278953_proxyholochain_AgentEntry_PublicKey_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AgentEntry_PublicKey_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AgentEntry_PublicKey_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AgentEntry_PublicKey_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AgentEntry_PublicKey_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nbyteslice) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_AgentEntry_PublicKey_Get(p0)
}
//go:cgo_export_dynamic new_holochain_AgentEntry
//go:linkname _cgoexp_ebd397278953_new_holochain_AgentEntry _cgoexp_ebd397278953_new_holochain_AgentEntry
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_AgentEntry
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_AgentEntry(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_AgentEntry
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_AgentEntry() (r0 _Ctype_int32_t) {
	return new_holochain_AgentEntry()
}
//go:cgo_export_dynamic proxyholochain_AgentFixture_Hash_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_AgentFixture_Hash_Set _cgoexp_ebd397278953_proxyholochain_AgentFixture_Hash_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AgentFixture_Hash_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AgentFixture_Hash_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AgentFixture_Hash_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AgentFixture_Hash_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_AgentFixture_Hash_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_AgentFixture_Hash_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_AgentFixture_Hash_Get _cgoexp_ebd397278953_proxyholochain_AgentFixture_Hash_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AgentFixture_Hash_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AgentFixture_Hash_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AgentFixture_Hash_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AgentFixture_Hash_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_AgentFixture_Hash_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_AgentFixture_Identity_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_AgentFixture_Identity_Set _cgoexp_ebd397278953_proxyholochain_AgentFixture_Identity_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AgentFixture_Identity_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AgentFixture_Identity_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AgentFixture_Identity_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AgentFixture_Identity_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_AgentFixture_Identity_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_AgentFixture_Identity_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_AgentFixture_Identity_Get _cgoexp_ebd397278953_proxyholochain_AgentFixture_Identity_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AgentFixture_Identity_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AgentFixture_Identity_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AgentFixture_Identity_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AgentFixture_Identity_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_AgentFixture_Identity_Get(p0)
}
//go:cgo_export_dynamic new_holochain_AgentFixture
//go:linkname _cgoexp_ebd397278953_new_holochain_AgentFixture _cgoexp_ebd397278953_new_holochain_AgentFixture
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_AgentFixture
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_AgentFixture(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_AgentFixture
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_AgentFixture() (r0 _Ctype_int32_t) {
	return new_holochain_AgentFixture()
}
//go:cgo_export_dynamic proxyholochain_AppMsg_ZomeType_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_AppMsg_ZomeType_Set _cgoexp_ebd397278953_proxyholochain_AppMsg_ZomeType_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AppMsg_ZomeType_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AppMsg_ZomeType_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AppMsg_ZomeType_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AppMsg_ZomeType_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_AppMsg_ZomeType_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_AppMsg_ZomeType_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_AppMsg_ZomeType_Get _cgoexp_ebd397278953_proxyholochain_AppMsg_ZomeType_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AppMsg_ZomeType_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AppMsg_ZomeType_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AppMsg_ZomeType_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AppMsg_ZomeType_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_AppMsg_ZomeType_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_AppMsg_Body_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_AppMsg_Body_Set _cgoexp_ebd397278953_proxyholochain_AppMsg_Body_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AppMsg_Body_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AppMsg_Body_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AppMsg_Body_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AppMsg_Body_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_AppMsg_Body_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_AppMsg_Body_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_AppMsg_Body_Get _cgoexp_ebd397278953_proxyholochain_AppMsg_Body_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AppMsg_Body_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AppMsg_Body_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AppMsg_Body_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AppMsg_Body_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_AppMsg_Body_Get(p0)
}
//go:cgo_export_dynamic new_holochain_AppMsg
//go:linkname _cgoexp_ebd397278953_new_holochain_AppMsg _cgoexp_ebd397278953_new_holochain_AppMsg
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_AppMsg
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_AppMsg(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_AppMsg
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_AppMsg() (r0 _Ctype_int32_t) {
	return new_holochain_AppMsg()
}
//go:cgo_export_dynamic proxyholochain_AppPackage_Version_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_AppPackage_Version_Set _cgoexp_ebd397278953_proxyholochain_AppPackage_Version_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AppPackage_Version_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AppPackage_Version_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AppPackage_Version_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AppPackage_Version_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_AppPackage_Version_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_AppPackage_Version_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_AppPackage_Version_Get _cgoexp_ebd397278953_proxyholochain_AppPackage_Version_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AppPackage_Version_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AppPackage_Version_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AppPackage_Version_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AppPackage_Version_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_AppPackage_Version_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_AppPackage_Generator_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_AppPackage_Generator_Set _cgoexp_ebd397278953_proxyholochain_AppPackage_Generator_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AppPackage_Generator_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AppPackage_Generator_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AppPackage_Generator_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AppPackage_Generator_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_AppPackage_Generator_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_AppPackage_Generator_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_AppPackage_Generator_Get _cgoexp_ebd397278953_proxyholochain_AppPackage_Generator_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AppPackage_Generator_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AppPackage_Generator_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AppPackage_Generator_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AppPackage_Generator_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_AppPackage_Generator_Get(p0)
}
//go:cgo_export_dynamic new_holochain_AppPackage
//go:linkname _cgoexp_ebd397278953_new_holochain_AppPackage _cgoexp_ebd397278953_new_holochain_AppPackage
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_AppPackage
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_AppPackage(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_AppPackage
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_AppPackage() (r0 _Ctype_int32_t) {
	return new_holochain_AppPackage()
}
//go:cgo_export_dynamic proxyholochain_AppPackageScenario_Name_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_AppPackageScenario_Name_Set _cgoexp_ebd397278953_proxyholochain_AppPackageScenario_Name_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AppPackageScenario_Name_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AppPackageScenario_Name_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AppPackageScenario_Name_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AppPackageScenario_Name_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_AppPackageScenario_Name_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_AppPackageScenario_Name_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_AppPackageScenario_Name_Get _cgoexp_ebd397278953_proxyholochain_AppPackageScenario_Name_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AppPackageScenario_Name_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AppPackageScenario_Name_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AppPackageScenario_Name_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AppPackageScenario_Name_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_AppPackageScenario_Name_Get(p0)
}
//go:cgo_export_dynamic new_holochain_AppPackageScenario
//go:linkname _cgoexp_ebd397278953_new_holochain_AppPackageScenario _cgoexp_ebd397278953_new_holochain_AppPackageScenario
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_AppPackageScenario
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_AppPackageScenario(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_AppPackageScenario
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_AppPackageScenario() (r0 _Ctype_int32_t) {
	return new_holochain_AppPackageScenario()
}
//go:cgo_export_dynamic proxyholochain_AppPackageTests_Name_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_AppPackageTests_Name_Set _cgoexp_ebd397278953_proxyholochain_AppPackageTests_Name_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AppPackageTests_Name_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AppPackageTests_Name_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AppPackageTests_Name_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AppPackageTests_Name_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_AppPackageTests_Name_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_AppPackageTests_Name_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_AppPackageTests_Name_Get _cgoexp_ebd397278953_proxyholochain_AppPackageTests_Name_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AppPackageTests_Name_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AppPackageTests_Name_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AppPackageTests_Name_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AppPackageTests_Name_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_AppPackageTests_Name_Get(p0)
}
//go:cgo_export_dynamic new_holochain_AppPackageTests
//go:linkname _cgoexp_ebd397278953_new_holochain_AppPackageTests _cgoexp_ebd397278953_new_holochain_AppPackageTests
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_AppPackageTests
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_AppPackageTests(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_AppPackageTests
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_AppPackageTests() (r0 _Ctype_int32_t) {
	return new_holochain_AppPackageTests()
}
//go:cgo_export_dynamic proxyholochain_AppPackageUIFile_FileName_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_FileName_Set _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_FileName_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_FileName_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_FileName_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AppPackageUIFile_FileName_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AppPackageUIFile_FileName_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_AppPackageUIFile_FileName_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_AppPackageUIFile_FileName_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_FileName_Get _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_FileName_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_FileName_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_FileName_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AppPackageUIFile_FileName_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AppPackageUIFile_FileName_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_AppPackageUIFile_FileName_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_AppPackageUIFile_Data_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Data_Set _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Data_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Data_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Data_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AppPackageUIFile_Data_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AppPackageUIFile_Data_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_AppPackageUIFile_Data_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_AppPackageUIFile_Data_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Data_Get _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Data_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Data_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Data_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AppPackageUIFile_Data_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AppPackageUIFile_Data_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_AppPackageUIFile_Data_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_AppPackageUIFile_Encoding_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Encoding_Set _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Encoding_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Encoding_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Encoding_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AppPackageUIFile_Encoding_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AppPackageUIFile_Encoding_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_AppPackageUIFile_Encoding_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_AppPackageUIFile_Encoding_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Encoding_Get _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Encoding_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Encoding_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Encoding_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_AppPackageUIFile_Encoding_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_AppPackageUIFile_Encoding_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_AppPackageUIFile_Encoding_Get(p0)
}
//go:cgo_export_dynamic new_holochain_AppPackageUIFile
//go:linkname _cgoexp_ebd397278953_new_holochain_AppPackageUIFile _cgoexp_ebd397278953_new_holochain_AppPackageUIFile
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_AppPackageUIFile
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_AppPackageUIFile(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_AppPackageUIFile
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_AppPackageUIFile() (r0 _Ctype_int32_t) {
	return new_holochain_AppPackageUIFile()
}
//go:cgo_export_dynamic proxyholochain_Arg_Name_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Arg_Name_Set _cgoexp_ebd397278953_proxyholochain_Arg_Name_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Arg_Name_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Arg_Name_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Arg_Name_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Arg_Name_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Arg_Name_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Arg_Name_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Arg_Name_Get _cgoexp_ebd397278953_proxyholochain_Arg_Name_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Arg_Name_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Arg_Name_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Arg_Name_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Arg_Name_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Arg_Name_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Arg_Optional_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Arg_Optional_Set _cgoexp_ebd397278953_proxyholochain_Arg_Optional_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Arg_Optional_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Arg_Optional_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Arg_Optional_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Arg_Optional_Set(p0 _Ctype_int32_t, p1 _Ctype_char) {
	proxyholochain_Arg_Optional_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Arg_Optional_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Arg_Optional_Get _cgoexp_ebd397278953_proxyholochain_Arg_Optional_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Arg_Optional_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Arg_Optional_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Arg_Optional_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Arg_Optional_Get(p0 _Ctype_int32_t) (r0 _Ctype_char) {
	return proxyholochain_Arg_Optional_Get(p0)
}
//go:cgo_export_dynamic new_holochain_Arg
//go:linkname _cgoexp_ebd397278953_new_holochain_Arg _cgoexp_ebd397278953_new_holochain_Arg
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Arg
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Arg(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Arg
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Arg() (r0 _Ctype_int32_t) {
	return new_holochain_Arg()
}
//go:cgo_export_dynamic proxyholochain_BSReq_Version_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_BSReq_Version_Set _cgoexp_ebd397278953_proxyholochain_BSReq_Version_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_BSReq_Version_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_BSReq_Version_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_BSReq_Version_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_BSReq_Version_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_BSReq_Version_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_BSReq_Version_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_BSReq_Version_Get _cgoexp_ebd397278953_proxyholochain_BSReq_Version_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_BSReq_Version_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_BSReq_Version_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_BSReq_Version_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_BSReq_Version_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_BSReq_Version_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_BSReq_NodeID_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_BSReq_NodeID_Set _cgoexp_ebd397278953_proxyholochain_BSReq_NodeID_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_BSReq_NodeID_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_BSReq_NodeID_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_BSReq_NodeID_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_BSReq_NodeID_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_BSReq_NodeID_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_BSReq_NodeID_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_BSReq_NodeID_Get _cgoexp_ebd397278953_proxyholochain_BSReq_NodeID_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_BSReq_NodeID_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_BSReq_NodeID_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_BSReq_NodeID_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_BSReq_NodeID_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_BSReq_NodeID_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_BSReq_NodeAddr_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_BSReq_NodeAddr_Set _cgoexp_ebd397278953_proxyholochain_BSReq_NodeAddr_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_BSReq_NodeAddr_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_BSReq_NodeAddr_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_BSReq_NodeAddr_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_BSReq_NodeAddr_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_BSReq_NodeAddr_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_BSReq_NodeAddr_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_BSReq_NodeAddr_Get _cgoexp_ebd397278953_proxyholochain_BSReq_NodeAddr_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_BSReq_NodeAddr_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_BSReq_NodeAddr_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_BSReq_NodeAddr_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_BSReq_NodeAddr_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_BSReq_NodeAddr_Get(p0)
}
//go:cgo_export_dynamic new_holochain_BSReq
//go:linkname _cgoexp_ebd397278953_new_holochain_BSReq _cgoexp_ebd397278953_new_holochain_BSReq
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_BSReq
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_BSReq(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_BSReq
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_BSReq() (r0 _Ctype_int32_t) {
	return new_holochain_BSReq()
}
//go:cgo_export_dynamic proxyholochain_BSResp_Remote_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_BSResp_Remote_Set _cgoexp_ebd397278953_proxyholochain_BSResp_Remote_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_BSResp_Remote_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_BSResp_Remote_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_BSResp_Remote_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_BSResp_Remote_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_BSResp_Remote_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_BSResp_Remote_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_BSResp_Remote_Get _cgoexp_ebd397278953_proxyholochain_BSResp_Remote_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_BSResp_Remote_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_BSResp_Remote_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_BSResp_Remote_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_BSResp_Remote_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_BSResp_Remote_Get(p0)
}
//go:cgo_export_dynamic new_holochain_BSResp
//go:linkname _cgoexp_ebd397278953_new_holochain_BSResp _cgoexp_ebd397278953_new_holochain_BSResp
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_BSResp
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_BSResp(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_BSResp
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_BSResp() (r0 _Ctype_int32_t) {
	return new_holochain_BSResp()
}
//go:cgo_export_dynamic proxyholochain_Bridge_Token_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Bridge_Token_Set _cgoexp_ebd397278953_proxyholochain_Bridge_Token_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Bridge_Token_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Bridge_Token_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Bridge_Token_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Bridge_Token_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Bridge_Token_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Bridge_Token_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Bridge_Token_Get _cgoexp_ebd397278953_proxyholochain_Bridge_Token_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Bridge_Token_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Bridge_Token_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Bridge_Token_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Bridge_Token_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Bridge_Token_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Bridge_Side_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Bridge_Side_Set _cgoexp_ebd397278953_proxyholochain_Bridge_Side_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Bridge_Side_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Bridge_Side_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Bridge_Side_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Bridge_Side_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_Bridge_Side_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Bridge_Side_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Bridge_Side_Get _cgoexp_ebd397278953_proxyholochain_Bridge_Side_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Bridge_Side_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Bridge_Side_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Bridge_Side_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Bridge_Side_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_Bridge_Side_Get(p0)
}
//go:cgo_export_dynamic new_holochain_Bridge
//go:linkname _cgoexp_ebd397278953_new_holochain_Bridge _cgoexp_ebd397278953_new_holochain_Bridge
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Bridge
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Bridge(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Bridge
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Bridge() (r0 _Ctype_int32_t) {
	return new_holochain_Bridge()
}
//go:cgo_export_dynamic proxyholochain_BridgeApp_H_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_BridgeApp_H_Set _cgoexp_ebd397278953_proxyholochain_BridgeApp_H_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_BridgeApp_H_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_BridgeApp_H_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_BridgeApp_H_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_BridgeApp_H_Set(p0 _Ctype_int32_t, p1 _Ctype_int32_t) {
	proxyholochain_BridgeApp_H_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_BridgeApp_H_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_BridgeApp_H_Get _cgoexp_ebd397278953_proxyholochain_BridgeApp_H_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_BridgeApp_H_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_BridgeApp_H_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_BridgeApp_H_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_BridgeApp_H_Get(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_BridgeApp_H_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_BridgeApp_Side_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_BridgeApp_Side_Set _cgoexp_ebd397278953_proxyholochain_BridgeApp_Side_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_BridgeApp_Side_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_BridgeApp_Side_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_BridgeApp_Side_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_BridgeApp_Side_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_BridgeApp_Side_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_BridgeApp_Side_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_BridgeApp_Side_Get _cgoexp_ebd397278953_proxyholochain_BridgeApp_Side_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_BridgeApp_Side_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_BridgeApp_Side_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_BridgeApp_Side_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_BridgeApp_Side_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_BridgeApp_Side_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_BridgeApp_BridgeGenesisDataFrom_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataFrom_Set _cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataFrom_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataFrom_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataFrom_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataFrom_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataFrom_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_BridgeApp_BridgeGenesisDataFrom_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_BridgeApp_BridgeGenesisDataFrom_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataFrom_Get _cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataFrom_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataFrom_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataFrom_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataFrom_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataFrom_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_BridgeApp_BridgeGenesisDataFrom_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_BridgeApp_BridgeGenesisDataTo_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataTo_Set _cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataTo_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataTo_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataTo_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataTo_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataTo_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_BridgeApp_BridgeGenesisDataTo_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_BridgeApp_BridgeGenesisDataTo_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataTo_Get _cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataTo_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataTo_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataTo_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataTo_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataTo_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_BridgeApp_BridgeGenesisDataTo_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_BridgeApp_Port_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_BridgeApp_Port_Set _cgoexp_ebd397278953_proxyholochain_BridgeApp_Port_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_BridgeApp_Port_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_BridgeApp_Port_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_BridgeApp_Port_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_BridgeApp_Port_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_BridgeApp_Port_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_BridgeApp_Port_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_BridgeApp_Port_Get _cgoexp_ebd397278953_proxyholochain_BridgeApp_Port_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_BridgeApp_Port_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_BridgeApp_Port_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_BridgeApp_Port_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_BridgeApp_Port_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_BridgeApp_Port_Get(p0)
}
//go:cgo_export_dynamic new_holochain_BridgeApp
//go:linkname _cgoexp_ebd397278953_new_holochain_BridgeApp _cgoexp_ebd397278953_new_holochain_BridgeApp
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_BridgeApp
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_BridgeApp(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_BridgeApp
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_BridgeApp() (r0 _Ctype_int32_t) {
	return new_holochain_BridgeApp()
}
//go:cgo_export_dynamic proxyholochain_Bucket_Len
//go:linkname _cgoexp_ebd397278953_proxyholochain_Bucket_Len _cgoexp_ebd397278953_proxyholochain_Bucket_Len
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Bucket_Len
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Bucket_Len(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Bucket_Len
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Bucket_Len(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_Bucket_Len(p0)
}
//go:cgo_export_dynamic new_holochain_Bucket
//go:linkname _cgoexp_ebd397278953_new_holochain_Bucket _cgoexp_ebd397278953_new_holochain_Bucket
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Bucket
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Bucket(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Bucket
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Bucket() (r0 _Ctype_int32_t) {
	return new_holochain_Bucket()
}
//go:cgo_export_dynamic proxyholochain_BytesSent_Bytes_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_BytesSent_Bytes_Set _cgoexp_ebd397278953_proxyholochain_BytesSent_Bytes_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_BytesSent_Bytes_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_BytesSent_Bytes_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_BytesSent_Bytes_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_BytesSent_Bytes_Set(p0 _Ctype_int32_t, p1 _Ctype_int64_t) {
	proxyholochain_BytesSent_Bytes_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_BytesSent_Bytes_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_BytesSent_Bytes_Get _cgoexp_ebd397278953_proxyholochain_BytesSent_Bytes_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_BytesSent_Bytes_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_BytesSent_Bytes_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_BytesSent_Bytes_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_BytesSent_Bytes_Get(p0 _Ctype_int32_t) (r0 _Ctype_int64_t) {
	return proxyholochain_BytesSent_Bytes_Get(p0)
}
//go:cgo_export_dynamic new_holochain_BytesSent
//go:linkname _cgoexp_ebd397278953_new_holochain_BytesSent _cgoexp_ebd397278953_new_holochain_BytesSent
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_BytesSent
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_BytesSent(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_BytesSent
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_BytesSent() (r0 _Ctype_int32_t) {
	return new_holochain_BytesSent()
}
//go:cgo_export_dynamic proxyholochain_Callback_Function_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Callback_Function_Set _cgoexp_ebd397278953_proxyholochain_Callback_Function_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Callback_Function_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Callback_Function_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Callback_Function_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Callback_Function_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Callback_Function_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Callback_Function_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Callback_Function_Get _cgoexp_ebd397278953_proxyholochain_Callback_Function_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Callback_Function_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Callback_Function_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Callback_Function_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Callback_Function_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Callback_Function_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Callback_ID_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Callback_ID_Set _cgoexp_ebd397278953_proxyholochain_Callback_ID_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Callback_ID_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Callback_ID_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Callback_ID_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Callback_ID_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Callback_ID_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Callback_ID_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Callback_ID_Get _cgoexp_ebd397278953_proxyholochain_Callback_ID_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Callback_ID_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Callback_ID_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Callback_ID_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Callback_ID_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Callback_ID_Get(p0)
}
//go:cgo_export_dynamic new_holochain_Callback
//go:linkname _cgoexp_ebd397278953_new_holochain_Callback _cgoexp_ebd397278953_new_holochain_Callback
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Callback
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Callback(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Callback
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Callback() (r0 _Ctype_int32_t) {
	return new_holochain_Callback()
}
//go:cgo_export_dynamic proxyholochain_Capability_Token_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Capability_Token_Set _cgoexp_ebd397278953_proxyholochain_Capability_Token_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Capability_Token_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Capability_Token_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Capability_Token_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Capability_Token_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Capability_Token_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Capability_Token_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Capability_Token_Get _cgoexp_ebd397278953_proxyholochain_Capability_Token_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Capability_Token_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Capability_Token_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Capability_Token_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Capability_Token_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Capability_Token_Get(p0)
}
//go:cgo_export_dynamic new_holochain_Capability
//go:linkname _cgoexp_ebd397278953_new_holochain_Capability _cgoexp_ebd397278953_new_holochain_Capability
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Capability
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Capability(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Capability
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Capability() (r0 _Ctype_int32_t) {
	return new_holochain_Capability()
}
//go:cgo_export_dynamic proxyholochain_Chain_Close
//go:linkname _cgoexp_ebd397278953_proxyholochain_Chain_Close _cgoexp_ebd397278953_proxyholochain_Chain_Close
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Chain_Close
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Chain_Close(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Chain_Close
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Chain_Close(p0 _Ctype_int32_t) {
	proxyholochain_Chain_Close(p0)
}
//go:cgo_export_dynamic proxyholochain_Chain_JSON
//go:linkname _cgoexp_ebd397278953_proxyholochain_Chain_JSON _cgoexp_ebd397278953_proxyholochain_Chain_JSON
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Chain_JSON
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Chain_JSON(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Chain_JSON
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Chain_JSON(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring, r1 _Ctype_int32_t) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Chain_JSON(p0)
}
//go:cgo_export_dynamic proxyholochain_Chain_Length
//go:linkname _cgoexp_ebd397278953_proxyholochain_Chain_Length _cgoexp_ebd397278953_proxyholochain_Chain_Length
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Chain_Length
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Chain_Length(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Chain_Length
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Chain_Length(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_Chain_Length(p0)
}
//go:cgo_export_dynamic proxyholochain_Chain_Nth
//go:linkname _cgoexp_ebd397278953_proxyholochain_Chain_Nth _cgoexp_ebd397278953_proxyholochain_Chain_Nth
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Chain_Nth
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Chain_Nth(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Chain_Nth
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Chain_Nth(p0 _Ctype_int32_t, p1 _Ctype_nint) (r0 _Ctype_int32_t) {
	return proxyholochain_Chain_Nth(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Chain_String
//go:linkname _cgoexp_ebd397278953_proxyholochain_Chain_String _cgoexp_ebd397278953_proxyholochain_Chain_String
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Chain_String
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Chain_String(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Chain_String
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Chain_String(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Chain_String(p0)
}
//go:cgo_export_dynamic proxyholochain_Chain_Top
//go:linkname _cgoexp_ebd397278953_proxyholochain_Chain_Top _cgoexp_ebd397278953_proxyholochain_Chain_Top
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Chain_Top
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Chain_Top(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Chain_Top
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Chain_Top(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_Chain_Top(p0)
}
//go:cgo_export_dynamic proxyholochain_Chain_Validate
//go:linkname _cgoexp_ebd397278953_proxyholochain_Chain_Validate _cgoexp_ebd397278953_proxyholochain_Chain_Validate
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Chain_Validate
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Chain_Validate(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Chain_Validate
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Chain_Validate(p0 _Ctype_int32_t, p1 _Ctype_char) (r0 _Ctype_int32_t) {
	return proxyholochain_Chain_Validate(p0, p1)
}
//go:cgo_export_dynamic new_holochain_Chain
//go:linkname _cgoexp_ebd397278953_new_holochain_Chain _cgoexp_ebd397278953_new_holochain_Chain
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Chain
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Chain(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Chain
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Chain() (r0 _Ctype_int32_t) {
	return new_holochain_Chain()
}
//go:cgo_export_dynamic proxyholochain_ChainPair_Header_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_ChainPair_Header_Set _cgoexp_ebd397278953_proxyholochain_ChainPair_Header_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ChainPair_Header_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ChainPair_Header_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ChainPair_Header_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ChainPair_Header_Set(p0 _Ctype_int32_t, p1 _Ctype_int32_t) {
	proxyholochain_ChainPair_Header_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ChainPair_Header_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_ChainPair_Header_Get _cgoexp_ebd397278953_proxyholochain_ChainPair_Header_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ChainPair_Header_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ChainPair_Header_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ChainPair_Header_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ChainPair_Header_Get(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_ChainPair_Header_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_ChainPair_Entry_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_ChainPair_Entry_Set _cgoexp_ebd397278953_proxyholochain_ChainPair_Entry_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ChainPair_Entry_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ChainPair_Entry_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ChainPair_Entry_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ChainPair_Entry_Set(p0 _Ctype_int32_t, p1 _Ctype_int32_t) {
	proxyholochain_ChainPair_Entry_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ChainPair_Entry_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_ChainPair_Entry_Get _cgoexp_ebd397278953_proxyholochain_ChainPair_Entry_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ChainPair_Entry_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ChainPair_Entry_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ChainPair_Entry_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ChainPair_Entry_Get(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_ChainPair_Entry_Get(p0)
}
//go:cgo_export_dynamic new_holochain_ChainPair
//go:linkname _cgoexp_ebd397278953_new_holochain_ChainPair _cgoexp_ebd397278953_new_holochain_ChainPair
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ChainPair
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ChainPair(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ChainPair
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ChainPair() (r0 _Ctype_int32_t) {
	return new_holochain_ChainPair()
}
//go:cgo_export_dynamic proxyholochain_Change_Message_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Change_Message_Set _cgoexp_ebd397278953_proxyholochain_Change_Message_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Change_Message_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Change_Message_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Change_Message_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Change_Message_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Change_Message_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Change_Message_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Change_Message_Get _cgoexp_ebd397278953_proxyholochain_Change_Message_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Change_Message_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Change_Message_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Change_Message_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Change_Message_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Change_Message_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Change_AsOf_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Change_AsOf_Set _cgoexp_ebd397278953_proxyholochain_Change_AsOf_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Change_AsOf_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Change_AsOf_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Change_AsOf_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Change_AsOf_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_Change_AsOf_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Change_AsOf_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Change_AsOf_Get _cgoexp_ebd397278953_proxyholochain_Change_AsOf_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Change_AsOf_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Change_AsOf_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Change_AsOf_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Change_AsOf_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_Change_AsOf_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Change_Log
//go:linkname _cgoexp_ebd397278953_proxyholochain_Change_Log _cgoexp_ebd397278953_proxyholochain_Change_Log
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Change_Log
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Change_Log(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Change_Log
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Change_Log(p0 _Ctype_int32_t) {
	proxyholochain_Change_Log(p0)
}
//go:cgo_export_dynamic new_holochain_Change
//go:linkname _cgoexp_ebd397278953_new_holochain_Change _cgoexp_ebd397278953_new_holochain_Change
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Change
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Change(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Change
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Change() (r0 _Ctype_int32_t) {
	return new_holochain_Change()
}
//go:cgo_export_dynamic proxyholochain_CloneSpec_Role_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_CloneSpec_Role_Set _cgoexp_ebd397278953_proxyholochain_CloneSpec_Role_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_CloneSpec_Role_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_CloneSpec_Role_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_CloneSpec_Role_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_CloneSpec_Role_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_CloneSpec_Role_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_CloneSpec_Role_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_CloneSpec_Role_Get _cgoexp_ebd397278953_proxyholochain_CloneSpec_Role_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_CloneSpec_Role_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_CloneSpec_Role_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_CloneSpec_Role_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_CloneSpec_Role_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_CloneSpec_Role_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_CloneSpec_Number_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_CloneSpec_Number_Set _cgoexp_ebd397278953_proxyholochain_CloneSpec_Number_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_CloneSpec_Number_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_CloneSpec_Number_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_CloneSpec_Number_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_CloneSpec_Number_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_CloneSpec_Number_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_CloneSpec_Number_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_CloneSpec_Number_Get _cgoexp_ebd397278953_proxyholochain_CloneSpec_Number_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_CloneSpec_Number_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_CloneSpec_Number_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_CloneSpec_Number_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_CloneSpec_Number_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_CloneSpec_Number_Get(p0)
}
//go:cgo_export_dynamic new_holochain_CloneSpec
//go:linkname _cgoexp_ebd397278953_new_holochain_CloneSpec _cgoexp_ebd397278953_new_holochain_CloneSpec
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_CloneSpec
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_CloneSpec(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_CloneSpec
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_CloneSpec() (r0 _Ctype_int32_t) {
	return new_holochain_CloneSpec()
}
//go:cgo_export_dynamic new_holochain_CloserPeersResp
//go:linkname _cgoexp_ebd397278953_new_holochain_CloserPeersResp _cgoexp_ebd397278953_new_holochain_CloserPeersResp
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_CloserPeersResp
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_CloserPeersResp(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_CloserPeersResp
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_CloserPeersResp() (r0 _Ctype_int32_t) {
	return new_holochain_CloserPeersResp()
}
//go:cgo_export_dynamic proxyholochain_Config_Port_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Config_Port_Set _cgoexp_ebd397278953_proxyholochain_Config_Port_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Config_Port_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Config_Port_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Config_Port_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Config_Port_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_Config_Port_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Config_Port_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Config_Port_Get _cgoexp_ebd397278953_proxyholochain_Config_Port_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Config_Port_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Config_Port_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Config_Port_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Config_Port_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_Config_Port_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Config_EnableMDNS_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Config_EnableMDNS_Set _cgoexp_ebd397278953_proxyholochain_Config_EnableMDNS_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Config_EnableMDNS_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Config_EnableMDNS_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Config_EnableMDNS_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Config_EnableMDNS_Set(p0 _Ctype_int32_t, p1 _Ctype_char) {
	proxyholochain_Config_EnableMDNS_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Config_EnableMDNS_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Config_EnableMDNS_Get _cgoexp_ebd397278953_proxyholochain_Config_EnableMDNS_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Config_EnableMDNS_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Config_EnableMDNS_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Config_EnableMDNS_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Config_EnableMDNS_Get(p0 _Ctype_int32_t) (r0 _Ctype_char) {
	return proxyholochain_Config_EnableMDNS_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Config_PeerModeAuthor_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Config_PeerModeAuthor_Set _cgoexp_ebd397278953_proxyholochain_Config_PeerModeAuthor_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Config_PeerModeAuthor_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Config_PeerModeAuthor_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Config_PeerModeAuthor_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Config_PeerModeAuthor_Set(p0 _Ctype_int32_t, p1 _Ctype_char) {
	proxyholochain_Config_PeerModeAuthor_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Config_PeerModeAuthor_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Config_PeerModeAuthor_Get _cgoexp_ebd397278953_proxyholochain_Config_PeerModeAuthor_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Config_PeerModeAuthor_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Config_PeerModeAuthor_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Config_PeerModeAuthor_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Config_PeerModeAuthor_Get(p0 _Ctype_int32_t) (r0 _Ctype_char) {
	return proxyholochain_Config_PeerModeAuthor_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Config_PeerModeDHTNode_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Config_PeerModeDHTNode_Set _cgoexp_ebd397278953_proxyholochain_Config_PeerModeDHTNode_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Config_PeerModeDHTNode_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Config_PeerModeDHTNode_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Config_PeerModeDHTNode_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Config_PeerModeDHTNode_Set(p0 _Ctype_int32_t, p1 _Ctype_char) {
	proxyholochain_Config_PeerModeDHTNode_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Config_PeerModeDHTNode_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Config_PeerModeDHTNode_Get _cgoexp_ebd397278953_proxyholochain_Config_PeerModeDHTNode_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Config_PeerModeDHTNode_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Config_PeerModeDHTNode_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Config_PeerModeDHTNode_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Config_PeerModeDHTNode_Get(p0 _Ctype_int32_t) (r0 _Ctype_char) {
	return proxyholochain_Config_PeerModeDHTNode_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Config_EnableNATUPnP_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Config_EnableNATUPnP_Set _cgoexp_ebd397278953_proxyholochain_Config_EnableNATUPnP_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Config_EnableNATUPnP_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Config_EnableNATUPnP_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Config_EnableNATUPnP_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Config_EnableNATUPnP_Set(p0 _Ctype_int32_t, p1 _Ctype_char) {
	proxyholochain_Config_EnableNATUPnP_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Config_EnableNATUPnP_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Config_EnableNATUPnP_Get _cgoexp_ebd397278953_proxyholochain_Config_EnableNATUPnP_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Config_EnableNATUPnP_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Config_EnableNATUPnP_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Config_EnableNATUPnP_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Config_EnableNATUPnP_Get(p0 _Ctype_int32_t) (r0 _Ctype_char) {
	return proxyholochain_Config_EnableNATUPnP_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Config_BootstrapServer_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Config_BootstrapServer_Set _cgoexp_ebd397278953_proxyholochain_Config_BootstrapServer_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Config_BootstrapServer_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Config_BootstrapServer_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Config_BootstrapServer_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Config_BootstrapServer_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Config_BootstrapServer_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Config_BootstrapServer_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Config_BootstrapServer_Get _cgoexp_ebd397278953_proxyholochain_Config_BootstrapServer_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Config_BootstrapServer_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Config_BootstrapServer_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Config_BootstrapServer_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Config_BootstrapServer_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Config_BootstrapServer_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Config_Setup
//go:linkname _cgoexp_ebd397278953_proxyholochain_Config_Setup _cgoexp_ebd397278953_proxyholochain_Config_Setup
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Config_Setup
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Config_Setup(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Config_Setup
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Config_Setup(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_Config_Setup(p0)
}
//go:cgo_export_dynamic proxyholochain_Config_SetupLogging
//go:linkname _cgoexp_ebd397278953_proxyholochain_Config_SetupLogging _cgoexp_ebd397278953_proxyholochain_Config_SetupLogging
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Config_SetupLogging
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Config_SetupLogging(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Config_SetupLogging
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Config_SetupLogging(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_Config_SetupLogging(p0)
}
//go:cgo_export_dynamic new_holochain_Config
//go:linkname _cgoexp_ebd397278953_new_holochain_Config _cgoexp_ebd397278953_new_holochain_Config
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Config
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Config(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Config
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Config() (r0 _Ctype_int32_t) {
	return new_holochain_Config()
}
//go:cgo_export_dynamic proxyholochain_DHT_Close
//go:linkname _cgoexp_ebd397278953_proxyholochain_DHT_Close _cgoexp_ebd397278953_proxyholochain_DHT_Close
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DHT_Close
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DHT_Close(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DHT_Close
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DHT_Close(p0 _Ctype_int32_t) {
	proxyholochain_DHT_Close(p0)
}
//go:cgo_export_dynamic proxyholochain_DHT_DumpIdx
//go:linkname _cgoexp_ebd397278953_proxyholochain_DHT_DumpIdx _cgoexp_ebd397278953_proxyholochain_DHT_DumpIdx
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DHT_DumpIdx
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DHT_DumpIdx(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DHT_DumpIdx
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DHT_DumpIdx(p0 _Ctype_int32_t, p1 _Ctype_nint) (r0 _Ctype_struct_nstring, r1 _Ctype_int32_t) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_DHT_DumpIdx(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_DHT_DumpIdxJSON
//go:linkname _cgoexp_ebd397278953_proxyholochain_DHT_DumpIdxJSON _cgoexp_ebd397278953_proxyholochain_DHT_DumpIdxJSON
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DHT_DumpIdxJSON
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DHT_DumpIdxJSON(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DHT_DumpIdxJSON
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DHT_DumpIdxJSON(p0 _Ctype_int32_t, p1 _Ctype_nint) (r0 _Ctype_struct_nstring, r1 _Ctype_int32_t) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_DHT_DumpIdxJSON(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_DHT_GetIdx
//go:linkname _cgoexp_ebd397278953_proxyholochain_DHT_GetIdx _cgoexp_ebd397278953_proxyholochain_DHT_GetIdx
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DHT_GetIdx
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DHT_GetIdx(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DHT_GetIdx
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DHT_GetIdx(p0 _Ctype_int32_t) (r0 _Ctype_nint, r1 _Ctype_int32_t) {
	return proxyholochain_DHT_GetIdx(p0)
}
//go:cgo_export_dynamic proxyholochain_DHT_HandleGossipPuts
//go:linkname _cgoexp_ebd397278953_proxyholochain_DHT_HandleGossipPuts _cgoexp_ebd397278953_proxyholochain_DHT_HandleGossipPuts
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DHT_HandleGossipPuts
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DHT_HandleGossipPuts(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DHT_HandleGossipPuts
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DHT_HandleGossipPuts(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_DHT_HandleGossipPuts(p0)
}
//go:cgo_export_dynamic proxyholochain_DHT_HandleGossipWiths
//go:linkname _cgoexp_ebd397278953_proxyholochain_DHT_HandleGossipWiths _cgoexp_ebd397278953_proxyholochain_DHT_HandleGossipWiths
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DHT_HandleGossipWiths
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DHT_HandleGossipWiths(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DHT_HandleGossipWiths
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DHT_HandleGossipWiths(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_DHT_HandleGossipWiths(p0)
}
//go:cgo_export_dynamic proxyholochain_DHT_JSON
//go:linkname _cgoexp_ebd397278953_proxyholochain_DHT_JSON _cgoexp_ebd397278953_proxyholochain_DHT_JSON
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DHT_JSON
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DHT_JSON(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DHT_JSON
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DHT_JSON(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring, r1 _Ctype_int32_t) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_DHT_JSON(p0)
}
//go:cgo_export_dynamic proxyholochain_DHT_SetupDHT
//go:linkname _cgoexp_ebd397278953_proxyholochain_DHT_SetupDHT _cgoexp_ebd397278953_proxyholochain_DHT_SetupDHT
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DHT_SetupDHT
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DHT_SetupDHT(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DHT_SetupDHT
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DHT_SetupDHT(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_DHT_SetupDHT(p0)
}
//go:cgo_export_dynamic proxyholochain_DHT_Start
//go:linkname _cgoexp_ebd397278953_proxyholochain_DHT_Start _cgoexp_ebd397278953_proxyholochain_DHT_Start
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DHT_Start
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DHT_Start(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DHT_Start
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DHT_Start(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_DHT_Start(p0)
}
//go:cgo_export_dynamic proxyholochain_DHT_String
//go:linkname _cgoexp_ebd397278953_proxyholochain_DHT_String _cgoexp_ebd397278953_proxyholochain_DHT_String
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DHT_String
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DHT_String(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DHT_String
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DHT_String(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_DHT_String(p0)
}
//go:cgo_export_dynamic new_holochain_DHT
//go:linkname _cgoexp_ebd397278953_new_holochain_DHT _cgoexp_ebd397278953_new_holochain_DHT
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_DHT
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_DHT(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_DHT
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_DHT() (r0 _Ctype_int32_t) {
	return new_holochain_DHT()
}
//go:cgo_export_dynamic proxyholochain_DHTConfig_HashType_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_DHTConfig_HashType_Set _cgoexp_ebd397278953_proxyholochain_DHTConfig_HashType_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DHTConfig_HashType_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DHTConfig_HashType_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DHTConfig_HashType_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DHTConfig_HashType_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_DHTConfig_HashType_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_DHTConfig_HashType_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_DHTConfig_HashType_Get _cgoexp_ebd397278953_proxyholochain_DHTConfig_HashType_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DHTConfig_HashType_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DHTConfig_HashType_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DHTConfig_HashType_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DHTConfig_HashType_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_DHTConfig_HashType_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_DHTConfig_NeighborhoodSize_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_DHTConfig_NeighborhoodSize_Set _cgoexp_ebd397278953_proxyholochain_DHTConfig_NeighborhoodSize_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DHTConfig_NeighborhoodSize_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DHTConfig_NeighborhoodSize_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DHTConfig_NeighborhoodSize_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DHTConfig_NeighborhoodSize_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_DHTConfig_NeighborhoodSize_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_DHTConfig_NeighborhoodSize_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_DHTConfig_NeighborhoodSize_Get _cgoexp_ebd397278953_proxyholochain_DHTConfig_NeighborhoodSize_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DHTConfig_NeighborhoodSize_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DHTConfig_NeighborhoodSize_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DHTConfig_NeighborhoodSize_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DHTConfig_NeighborhoodSize_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_DHTConfig_NeighborhoodSize_Get(p0)
}
//go:cgo_export_dynamic new_holochain_DHTConfig
//go:linkname _cgoexp_ebd397278953_new_holochain_DHTConfig _cgoexp_ebd397278953_new_holochain_DHTConfig
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_DHTConfig
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_DHTConfig(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_DHTConfig
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_DHTConfig() (r0 _Ctype_int32_t) {
	return new_holochain_DHTConfig()
}
//go:cgo_export_dynamic proxyholochain_DNA_Version_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_DNA_Version_Set _cgoexp_ebd397278953_proxyholochain_DNA_Version_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DNA_Version_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DNA_Version_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DNA_Version_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DNA_Version_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_DNA_Version_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_DNA_Version_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_DNA_Version_Get _cgoexp_ebd397278953_proxyholochain_DNA_Version_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DNA_Version_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DNA_Version_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DNA_Version_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DNA_Version_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_DNA_Version_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_DNA_Name_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_DNA_Name_Set _cgoexp_ebd397278953_proxyholochain_DNA_Name_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DNA_Name_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DNA_Name_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DNA_Name_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DNA_Name_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_DNA_Name_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_DNA_Name_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_DNA_Name_Get _cgoexp_ebd397278953_proxyholochain_DNA_Name_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DNA_Name_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DNA_Name_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DNA_Name_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DNA_Name_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_DNA_Name_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_DNA_PropertiesSchema_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_DNA_PropertiesSchema_Set _cgoexp_ebd397278953_proxyholochain_DNA_PropertiesSchema_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DNA_PropertiesSchema_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DNA_PropertiesSchema_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DNA_PropertiesSchema_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DNA_PropertiesSchema_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_DNA_PropertiesSchema_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_DNA_PropertiesSchema_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_DNA_PropertiesSchema_Get _cgoexp_ebd397278953_proxyholochain_DNA_PropertiesSchema_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DNA_PropertiesSchema_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DNA_PropertiesSchema_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DNA_PropertiesSchema_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DNA_PropertiesSchema_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_DNA_PropertiesSchema_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_DNA_AgentIdentitySchema_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_DNA_AgentIdentitySchema_Set _cgoexp_ebd397278953_proxyholochain_DNA_AgentIdentitySchema_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DNA_AgentIdentitySchema_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DNA_AgentIdentitySchema_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DNA_AgentIdentitySchema_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DNA_AgentIdentitySchema_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_DNA_AgentIdentitySchema_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_DNA_AgentIdentitySchema_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_DNA_AgentIdentitySchema_Get _cgoexp_ebd397278953_proxyholochain_DNA_AgentIdentitySchema_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DNA_AgentIdentitySchema_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DNA_AgentIdentitySchema_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DNA_AgentIdentitySchema_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DNA_AgentIdentitySchema_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_DNA_AgentIdentitySchema_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_DNA_RequiresVersion_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_DNA_RequiresVersion_Set _cgoexp_ebd397278953_proxyholochain_DNA_RequiresVersion_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DNA_RequiresVersion_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DNA_RequiresVersion_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DNA_RequiresVersion_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DNA_RequiresVersion_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_DNA_RequiresVersion_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_DNA_RequiresVersion_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_DNA_RequiresVersion_Get _cgoexp_ebd397278953_proxyholochain_DNA_RequiresVersion_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DNA_RequiresVersion_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DNA_RequiresVersion_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DNA_RequiresVersion_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DNA_RequiresVersion_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_DNA_RequiresVersion_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_DNA_NewUUID
//go:linkname _cgoexp_ebd397278953_proxyholochain_DNA_NewUUID _cgoexp_ebd397278953_proxyholochain_DNA_NewUUID
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DNA_NewUUID
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DNA_NewUUID(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DNA_NewUUID
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DNA_NewUUID(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_DNA_NewUUID(p0)
}
//go:cgo_export_dynamic new_holochain_DNA
//go:linkname _cgoexp_ebd397278953_new_holochain_DNA _cgoexp_ebd397278953_new_holochain_DNA
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_DNA
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_DNA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_DNA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_DNA() (r0 _Ctype_int32_t) {
	return new_holochain_DNA()
}
//go:cgo_export_dynamic proxyholochain_DNAFile_Version_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_DNAFile_Version_Set _cgoexp_ebd397278953_proxyholochain_DNAFile_Version_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DNAFile_Version_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DNAFile_Version_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DNAFile_Version_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DNAFile_Version_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_DNAFile_Version_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_DNAFile_Version_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_DNAFile_Version_Get _cgoexp_ebd397278953_proxyholochain_DNAFile_Version_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DNAFile_Version_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DNAFile_Version_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DNAFile_Version_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DNAFile_Version_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_DNAFile_Version_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_DNAFile_Name_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_DNAFile_Name_Set _cgoexp_ebd397278953_proxyholochain_DNAFile_Name_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DNAFile_Name_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DNAFile_Name_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DNAFile_Name_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DNAFile_Name_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_DNAFile_Name_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_DNAFile_Name_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_DNAFile_Name_Get _cgoexp_ebd397278953_proxyholochain_DNAFile_Name_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DNAFile_Name_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DNAFile_Name_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DNAFile_Name_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DNAFile_Name_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_DNAFile_Name_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_DNAFile_PropertiesSchemaFile_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_DNAFile_PropertiesSchemaFile_Set _cgoexp_ebd397278953_proxyholochain_DNAFile_PropertiesSchemaFile_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DNAFile_PropertiesSchemaFile_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DNAFile_PropertiesSchemaFile_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DNAFile_PropertiesSchemaFile_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DNAFile_PropertiesSchemaFile_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_DNAFile_PropertiesSchemaFile_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_DNAFile_PropertiesSchemaFile_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_DNAFile_PropertiesSchemaFile_Get _cgoexp_ebd397278953_proxyholochain_DNAFile_PropertiesSchemaFile_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DNAFile_PropertiesSchemaFile_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DNAFile_PropertiesSchemaFile_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DNAFile_PropertiesSchemaFile_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DNAFile_PropertiesSchemaFile_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_DNAFile_PropertiesSchemaFile_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_DNAFile_RequiresVersion_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_DNAFile_RequiresVersion_Set _cgoexp_ebd397278953_proxyholochain_DNAFile_RequiresVersion_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DNAFile_RequiresVersion_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DNAFile_RequiresVersion_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DNAFile_RequiresVersion_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DNAFile_RequiresVersion_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_DNAFile_RequiresVersion_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_DNAFile_RequiresVersion_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_DNAFile_RequiresVersion_Get _cgoexp_ebd397278953_proxyholochain_DNAFile_RequiresVersion_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DNAFile_RequiresVersion_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DNAFile_RequiresVersion_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DNAFile_RequiresVersion_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DNAFile_RequiresVersion_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_DNAFile_RequiresVersion_Get(p0)
}
//go:cgo_export_dynamic new_holochain_DNAFile
//go:linkname _cgoexp_ebd397278953_new_holochain_DNAFile _cgoexp_ebd397278953_new_holochain_DNAFile
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_DNAFile
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_DNAFile(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_DNAFile
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_DNAFile() (r0 _Ctype_int32_t) {
	return new_holochain_DNAFile()
}
//go:cgo_export_dynamic proxyholochain_DelEntry_Message_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_DelEntry_Message_Set _cgoexp_ebd397278953_proxyholochain_DelEntry_Message_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DelEntry_Message_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DelEntry_Message_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DelEntry_Message_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DelEntry_Message_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_DelEntry_Message_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_DelEntry_Message_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_DelEntry_Message_Get _cgoexp_ebd397278953_proxyholochain_DelEntry_Message_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_DelEntry_Message_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_DelEntry_Message_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_DelEntry_Message_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_DelEntry_Message_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_DelEntry_Message_Get(p0)
}
//go:cgo_export_dynamic new_holochain_DelEntry
//go:linkname _cgoexp_ebd397278953_new_holochain_DelEntry _cgoexp_ebd397278953_new_holochain_DelEntry
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_DelEntry
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_DelEntry(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_DelEntry
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_DelEntry() (r0 _Ctype_int32_t) {
	return new_holochain_DelEntry()
}
//go:cgo_export_dynamic new_holochain_DelReq
//go:linkname _cgoexp_ebd397278953_new_holochain_DelReq _cgoexp_ebd397278953_new_holochain_DelReq
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_DelReq
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_DelReq(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_DelReq
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_DelReq() (r0 _Ctype_int32_t) {
	return new_holochain_DelReq()
}
//go:cgo_export_dynamic proxyholochain_EntryDef_Name_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_EntryDef_Name_Set _cgoexp_ebd397278953_proxyholochain_EntryDef_Name_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_EntryDef_Name_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_EntryDef_Name_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_Name_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_Name_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_EntryDef_Name_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_EntryDef_Name_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_EntryDef_Name_Get _cgoexp_ebd397278953_proxyholochain_EntryDef_Name_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_EntryDef_Name_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_EntryDef_Name_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_Name_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_Name_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_EntryDef_Name_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_EntryDef_DataFormat_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_EntryDef_DataFormat_Set _cgoexp_ebd397278953_proxyholochain_EntryDef_DataFormat_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_EntryDef_DataFormat_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_EntryDef_DataFormat_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_DataFormat_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_DataFormat_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_EntryDef_DataFormat_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_EntryDef_DataFormat_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_EntryDef_DataFormat_Get _cgoexp_ebd397278953_proxyholochain_EntryDef_DataFormat_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_EntryDef_DataFormat_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_EntryDef_DataFormat_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_DataFormat_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_DataFormat_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_EntryDef_DataFormat_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_EntryDef_Sharing_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_EntryDef_Sharing_Set _cgoexp_ebd397278953_proxyholochain_EntryDef_Sharing_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_EntryDef_Sharing_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_EntryDef_Sharing_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_Sharing_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_Sharing_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_EntryDef_Sharing_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_EntryDef_Sharing_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_EntryDef_Sharing_Get _cgoexp_ebd397278953_proxyholochain_EntryDef_Sharing_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_EntryDef_Sharing_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_EntryDef_Sharing_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_Sharing_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_Sharing_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_EntryDef_Sharing_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_EntryDef_Schema_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_EntryDef_Schema_Set _cgoexp_ebd397278953_proxyholochain_EntryDef_Schema_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_EntryDef_Schema_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_EntryDef_Schema_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_Schema_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_Schema_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_EntryDef_Schema_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_EntryDef_Schema_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_EntryDef_Schema_Get _cgoexp_ebd397278953_proxyholochain_EntryDef_Schema_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_EntryDef_Schema_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_EntryDef_Schema_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_Schema_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_Schema_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_EntryDef_Schema_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_EntryDef_BuildJSONSchemaValidator
//go:linkname _cgoexp_ebd397278953_proxyholochain_EntryDef_BuildJSONSchemaValidator _cgoexp_ebd397278953_proxyholochain_EntryDef_BuildJSONSchemaValidator
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_EntryDef_BuildJSONSchemaValidator
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_EntryDef_BuildJSONSchemaValidator(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_BuildJSONSchemaValidator
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_BuildJSONSchemaValidator(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) (r0 _Ctype_int32_t) {
	return proxyholochain_EntryDef_BuildJSONSchemaValidator(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_EntryDef_BuildJSONSchemaValidatorFromString
//go:linkname _cgoexp_ebd397278953_proxyholochain_EntryDef_BuildJSONSchemaValidatorFromString _cgoexp_ebd397278953_proxyholochain_EntryDef_BuildJSONSchemaValidatorFromString
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_EntryDef_BuildJSONSchemaValidatorFromString
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_EntryDef_BuildJSONSchemaValidatorFromString(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_BuildJSONSchemaValidatorFromString
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_BuildJSONSchemaValidatorFromString(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) (r0 _Ctype_int32_t) {
	return proxyholochain_EntryDef_BuildJSONSchemaValidatorFromString(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_EntryDef_IsSysEntry
//go:linkname _cgoexp_ebd397278953_proxyholochain_EntryDef_IsSysEntry _cgoexp_ebd397278953_proxyholochain_EntryDef_IsSysEntry
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_EntryDef_IsSysEntry
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_EntryDef_IsSysEntry(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_IsSysEntry
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_IsSysEntry(p0 _Ctype_int32_t) (r0 _Ctype_char) {
	return proxyholochain_EntryDef_IsSysEntry(p0)
}
//go:cgo_export_dynamic proxyholochain_EntryDef_IsVirtualEntry
//go:linkname _cgoexp_ebd397278953_proxyholochain_EntryDef_IsVirtualEntry _cgoexp_ebd397278953_proxyholochain_EntryDef_IsVirtualEntry
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_EntryDef_IsVirtualEntry
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_EntryDef_IsVirtualEntry(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_IsVirtualEntry
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_EntryDef_IsVirtualEntry(p0 _Ctype_int32_t) (r0 _Ctype_char) {
	return proxyholochain_EntryDef_IsVirtualEntry(p0)
}
//go:cgo_export_dynamic new_holochain_EntryDef
//go:linkname _cgoexp_ebd397278953_new_holochain_EntryDef _cgoexp_ebd397278953_new_holochain_EntryDef
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_EntryDef
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_EntryDef(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_EntryDef
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_EntryDef() (r0 _Ctype_int32_t) {
	return new_holochain_EntryDef()
}
//go:cgo_export_dynamic proxyholochain_EntryDefFile_Name_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Name_Set _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Name_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Name_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Name_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_EntryDefFile_Name_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_EntryDefFile_Name_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_EntryDefFile_Name_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_EntryDefFile_Name_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Name_Get _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Name_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Name_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Name_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_EntryDefFile_Name_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_EntryDefFile_Name_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_EntryDefFile_Name_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_EntryDefFile_DataFormat_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_EntryDefFile_DataFormat_Set _cgoexp_ebd397278953_proxyholochain_EntryDefFile_DataFormat_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_EntryDefFile_DataFormat_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_EntryDefFile_DataFormat_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_EntryDefFile_DataFormat_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_EntryDefFile_DataFormat_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_EntryDefFile_DataFormat_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_EntryDefFile_DataFormat_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_EntryDefFile_DataFormat_Get _cgoexp_ebd397278953_proxyholochain_EntryDefFile_DataFormat_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_EntryDefFile_DataFormat_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_EntryDefFile_DataFormat_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_EntryDefFile_DataFormat_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_EntryDefFile_DataFormat_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_EntryDefFile_DataFormat_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_EntryDefFile_Schema_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Schema_Set _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Schema_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Schema_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Schema_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_EntryDefFile_Schema_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_EntryDefFile_Schema_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_EntryDefFile_Schema_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_EntryDefFile_Schema_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Schema_Get _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Schema_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Schema_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Schema_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_EntryDefFile_Schema_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_EntryDefFile_Schema_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_EntryDefFile_Schema_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_EntryDefFile_SchemaFile_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_EntryDefFile_SchemaFile_Set _cgoexp_ebd397278953_proxyholochain_EntryDefFile_SchemaFile_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_EntryDefFile_SchemaFile_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_EntryDefFile_SchemaFile_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_EntryDefFile_SchemaFile_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_EntryDefFile_SchemaFile_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_EntryDefFile_SchemaFile_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_EntryDefFile_SchemaFile_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_EntryDefFile_SchemaFile_Get _cgoexp_ebd397278953_proxyholochain_EntryDefFile_SchemaFile_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_EntryDefFile_SchemaFile_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_EntryDefFile_SchemaFile_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_EntryDefFile_SchemaFile_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_EntryDefFile_SchemaFile_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_EntryDefFile_SchemaFile_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_EntryDefFile_Sharing_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Sharing_Set _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Sharing_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Sharing_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Sharing_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_EntryDefFile_Sharing_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_EntryDefFile_Sharing_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_EntryDefFile_Sharing_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_EntryDefFile_Sharing_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Sharing_Get _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Sharing_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Sharing_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Sharing_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_EntryDefFile_Sharing_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_EntryDefFile_Sharing_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_EntryDefFile_Sharing_Get(p0)
}
//go:cgo_export_dynamic new_holochain_EntryDefFile
//go:linkname _cgoexp_ebd397278953_new_holochain_EntryDefFile _cgoexp_ebd397278953_new_holochain_EntryDefFile
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_EntryDefFile
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_EntryDefFile(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_EntryDefFile
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_EntryDefFile() (r0 _Ctype_int32_t) {
	return new_holochain_EntryDefFile()
}
//go:cgo_export_dynamic proxyholochain_ErrorResponse_Code_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_ErrorResponse_Code_Set _cgoexp_ebd397278953_proxyholochain_ErrorResponse_Code_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ErrorResponse_Code_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ErrorResponse_Code_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ErrorResponse_Code_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ErrorResponse_Code_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_ErrorResponse_Code_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ErrorResponse_Code_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_ErrorResponse_Code_Get _cgoexp_ebd397278953_proxyholochain_ErrorResponse_Code_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ErrorResponse_Code_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ErrorResponse_Code_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ErrorResponse_Code_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ErrorResponse_Code_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_ErrorResponse_Code_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_ErrorResponse_Message_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_ErrorResponse_Message_Set _cgoexp_ebd397278953_proxyholochain_ErrorResponse_Message_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ErrorResponse_Message_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ErrorResponse_Message_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ErrorResponse_Message_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ErrorResponse_Message_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_ErrorResponse_Message_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ErrorResponse_Message_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_ErrorResponse_Message_Get _cgoexp_ebd397278953_proxyholochain_ErrorResponse_Message_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ErrorResponse_Message_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ErrorResponse_Message_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ErrorResponse_Message_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ErrorResponse_Message_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ErrorResponse_Message_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_ErrorResponse_DecodeResponseError
//go:linkname _cgoexp_ebd397278953_proxyholochain_ErrorResponse_DecodeResponseError _cgoexp_ebd397278953_proxyholochain_ErrorResponse_DecodeResponseError
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ErrorResponse_DecodeResponseError
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ErrorResponse_DecodeResponseError(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ErrorResponse_DecodeResponseError
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ErrorResponse_DecodeResponseError(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_ErrorResponse_DecodeResponseError(p0)
}
//go:cgo_export_dynamic new_holochain_ErrorResponse
//go:linkname _cgoexp_ebd397278953_new_holochain_ErrorResponse _cgoexp_ebd397278953_new_holochain_ErrorResponse
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ErrorResponse
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ErrorResponse(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ErrorResponse
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ErrorResponse() (r0 _Ctype_int32_t) {
	return new_holochain_ErrorResponse()
}
//go:cgo_export_dynamic new_holochain_FindNodeReq
//go:linkname _cgoexp_ebd397278953_new_holochain_FindNodeReq _cgoexp_ebd397278953_new_holochain_FindNodeReq
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_FindNodeReq
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_FindNodeReq(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_FindNodeReq
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_FindNodeReq() (r0 _Ctype_int32_t) {
	return new_holochain_FindNodeReq()
}
//go:cgo_export_dynamic proxyholochain_FunctionDef_Name_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_FunctionDef_Name_Set _cgoexp_ebd397278953_proxyholochain_FunctionDef_Name_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_FunctionDef_Name_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_FunctionDef_Name_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_FunctionDef_Name_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_FunctionDef_Name_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_FunctionDef_Name_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_FunctionDef_Name_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_FunctionDef_Name_Get _cgoexp_ebd397278953_proxyholochain_FunctionDef_Name_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_FunctionDef_Name_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_FunctionDef_Name_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_FunctionDef_Name_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_FunctionDef_Name_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_FunctionDef_Name_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_FunctionDef_CallingType_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_FunctionDef_CallingType_Set _cgoexp_ebd397278953_proxyholochain_FunctionDef_CallingType_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_FunctionDef_CallingType_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_FunctionDef_CallingType_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_FunctionDef_CallingType_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_FunctionDef_CallingType_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_FunctionDef_CallingType_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_FunctionDef_CallingType_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_FunctionDef_CallingType_Get _cgoexp_ebd397278953_proxyholochain_FunctionDef_CallingType_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_FunctionDef_CallingType_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_FunctionDef_CallingType_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_FunctionDef_CallingType_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_FunctionDef_CallingType_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_FunctionDef_CallingType_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_FunctionDef_Exposure_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_FunctionDef_Exposure_Set _cgoexp_ebd397278953_proxyholochain_FunctionDef_Exposure_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_FunctionDef_Exposure_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_FunctionDef_Exposure_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_FunctionDef_Exposure_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_FunctionDef_Exposure_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_FunctionDef_Exposure_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_FunctionDef_Exposure_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_FunctionDef_Exposure_Get _cgoexp_ebd397278953_proxyholochain_FunctionDef_Exposure_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_FunctionDef_Exposure_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_FunctionDef_Exposure_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_FunctionDef_Exposure_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_FunctionDef_Exposure_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_FunctionDef_Exposure_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_FunctionDef_ValidExposure
//go:linkname _cgoexp_ebd397278953_proxyholochain_FunctionDef_ValidExposure _cgoexp_ebd397278953_proxyholochain_FunctionDef_ValidExposure
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_FunctionDef_ValidExposure
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_FunctionDef_ValidExposure(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_FunctionDef_ValidExposure
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_FunctionDef_ValidExposure(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) (r0 _Ctype_char) {
	return proxyholochain_FunctionDef_ValidExposure(p0, p1)
}
//go:cgo_export_dynamic new_holochain_FunctionDef
//go:linkname _cgoexp_ebd397278953_new_holochain_FunctionDef _cgoexp_ebd397278953_new_holochain_FunctionDef
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_FunctionDef
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_FunctionDef(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_FunctionDef
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_FunctionDef() (r0 _Ctype_int32_t) {
	return new_holochain_FunctionDef()
}
//go:cgo_export_dynamic proxyholochain_GetLinksOptions_Load_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_GetLinksOptions_Load_Set _cgoexp_ebd397278953_proxyholochain_GetLinksOptions_Load_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GetLinksOptions_Load_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GetLinksOptions_Load_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GetLinksOptions_Load_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GetLinksOptions_Load_Set(p0 _Ctype_int32_t, p1 _Ctype_char) {
	proxyholochain_GetLinksOptions_Load_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_GetLinksOptions_Load_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_GetLinksOptions_Load_Get _cgoexp_ebd397278953_proxyholochain_GetLinksOptions_Load_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GetLinksOptions_Load_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GetLinksOptions_Load_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GetLinksOptions_Load_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GetLinksOptions_Load_Get(p0 _Ctype_int32_t) (r0 _Ctype_char) {
	return proxyholochain_GetLinksOptions_Load_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_GetLinksOptions_StatusMask_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_GetLinksOptions_StatusMask_Set _cgoexp_ebd397278953_proxyholochain_GetLinksOptions_StatusMask_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GetLinksOptions_StatusMask_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GetLinksOptions_StatusMask_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GetLinksOptions_StatusMask_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GetLinksOptions_StatusMask_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_GetLinksOptions_StatusMask_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_GetLinksOptions_StatusMask_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_GetLinksOptions_StatusMask_Get _cgoexp_ebd397278953_proxyholochain_GetLinksOptions_StatusMask_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GetLinksOptions_StatusMask_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GetLinksOptions_StatusMask_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GetLinksOptions_StatusMask_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GetLinksOptions_StatusMask_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_GetLinksOptions_StatusMask_Get(p0)
}
//go:cgo_export_dynamic new_holochain_GetLinksOptions
//go:linkname _cgoexp_ebd397278953_new_holochain_GetLinksOptions _cgoexp_ebd397278953_new_holochain_GetLinksOptions
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_GetLinksOptions
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_GetLinksOptions(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_GetLinksOptions
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_GetLinksOptions() (r0 _Ctype_int32_t) {
	return new_holochain_GetLinksOptions()
}
//go:cgo_export_dynamic proxyholochain_GetOptions_StatusMask_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_GetOptions_StatusMask_Set _cgoexp_ebd397278953_proxyholochain_GetOptions_StatusMask_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GetOptions_StatusMask_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GetOptions_StatusMask_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GetOptions_StatusMask_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GetOptions_StatusMask_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_GetOptions_StatusMask_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_GetOptions_StatusMask_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_GetOptions_StatusMask_Get _cgoexp_ebd397278953_proxyholochain_GetOptions_StatusMask_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GetOptions_StatusMask_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GetOptions_StatusMask_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GetOptions_StatusMask_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GetOptions_StatusMask_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_GetOptions_StatusMask_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_GetOptions_GetMask_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_GetOptions_GetMask_Set _cgoexp_ebd397278953_proxyholochain_GetOptions_GetMask_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GetOptions_GetMask_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GetOptions_GetMask_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GetOptions_GetMask_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GetOptions_GetMask_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_GetOptions_GetMask_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_GetOptions_GetMask_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_GetOptions_GetMask_Get _cgoexp_ebd397278953_proxyholochain_GetOptions_GetMask_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GetOptions_GetMask_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GetOptions_GetMask_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GetOptions_GetMask_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GetOptions_GetMask_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_GetOptions_GetMask_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_GetOptions_Local_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_GetOptions_Local_Set _cgoexp_ebd397278953_proxyholochain_GetOptions_Local_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GetOptions_Local_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GetOptions_Local_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GetOptions_Local_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GetOptions_Local_Set(p0 _Ctype_int32_t, p1 _Ctype_char) {
	proxyholochain_GetOptions_Local_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_GetOptions_Local_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_GetOptions_Local_Get _cgoexp_ebd397278953_proxyholochain_GetOptions_Local_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GetOptions_Local_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GetOptions_Local_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GetOptions_Local_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GetOptions_Local_Get(p0 _Ctype_int32_t) (r0 _Ctype_char) {
	return proxyholochain_GetOptions_Local_Get(p0)
}
//go:cgo_export_dynamic new_holochain_GetOptions
//go:linkname _cgoexp_ebd397278953_new_holochain_GetOptions _cgoexp_ebd397278953_new_holochain_GetOptions
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_GetOptions
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_GetOptions(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_GetOptions
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_GetOptions() (r0 _Ctype_int32_t) {
	return new_holochain_GetOptions()
}
//go:cgo_export_dynamic proxyholochain_GetReq_StatusMask_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_GetReq_StatusMask_Set _cgoexp_ebd397278953_proxyholochain_GetReq_StatusMask_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GetReq_StatusMask_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GetReq_StatusMask_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GetReq_StatusMask_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GetReq_StatusMask_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_GetReq_StatusMask_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_GetReq_StatusMask_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_GetReq_StatusMask_Get _cgoexp_ebd397278953_proxyholochain_GetReq_StatusMask_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GetReq_StatusMask_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GetReq_StatusMask_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GetReq_StatusMask_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GetReq_StatusMask_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_GetReq_StatusMask_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_GetReq_GetMask_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_GetReq_GetMask_Set _cgoexp_ebd397278953_proxyholochain_GetReq_GetMask_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GetReq_GetMask_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GetReq_GetMask_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GetReq_GetMask_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GetReq_GetMask_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_GetReq_GetMask_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_GetReq_GetMask_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_GetReq_GetMask_Get _cgoexp_ebd397278953_proxyholochain_GetReq_GetMask_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GetReq_GetMask_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GetReq_GetMask_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GetReq_GetMask_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GetReq_GetMask_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_GetReq_GetMask_Get(p0)
}
//go:cgo_export_dynamic new_holochain_GetReq
//go:linkname _cgoexp_ebd397278953_new_holochain_GetReq _cgoexp_ebd397278953_new_holochain_GetReq
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_GetReq
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_GetReq(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_GetReq
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_GetReq() (r0 _Ctype_int32_t) {
	return new_holochain_GetReq()
}
//go:cgo_export_dynamic proxyholochain_GetResp_EntryType_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_GetResp_EntryType_Set _cgoexp_ebd397278953_proxyholochain_GetResp_EntryType_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GetResp_EntryType_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GetResp_EntryType_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GetResp_EntryType_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GetResp_EntryType_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_GetResp_EntryType_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_GetResp_EntryType_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_GetResp_EntryType_Get _cgoexp_ebd397278953_proxyholochain_GetResp_EntryType_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GetResp_EntryType_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GetResp_EntryType_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GetResp_EntryType_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GetResp_EntryType_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_GetResp_EntryType_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_GetResp_FollowHash_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_GetResp_FollowHash_Set _cgoexp_ebd397278953_proxyholochain_GetResp_FollowHash_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GetResp_FollowHash_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GetResp_FollowHash_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GetResp_FollowHash_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GetResp_FollowHash_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_GetResp_FollowHash_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_GetResp_FollowHash_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_GetResp_FollowHash_Get _cgoexp_ebd397278953_proxyholochain_GetResp_FollowHash_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GetResp_FollowHash_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GetResp_FollowHash_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GetResp_FollowHash_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GetResp_FollowHash_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_GetResp_FollowHash_Get(p0)
}
//go:cgo_export_dynamic new_holochain_GetResp
//go:linkname _cgoexp_ebd397278953_new_holochain_GetResp _cgoexp_ebd397278953_new_holochain_GetResp
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_GetResp
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_GetResp(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_GetResp
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_GetResp() (r0 _Ctype_int32_t) {
	return new_holochain_GetResp()
}
//go:cgo_export_dynamic proxyholochain_GobEntry_Marshal
//go:linkname _cgoexp_ebd397278953_proxyholochain_GobEntry_Marshal _cgoexp_ebd397278953_proxyholochain_GobEntry_Marshal
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GobEntry_Marshal
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GobEntry_Marshal(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GobEntry_Marshal
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GobEntry_Marshal(p0 _Ctype_int32_t) (r0 _Ctype_struct_nbyteslice, r1 _Ctype_int32_t) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_GobEntry_Marshal(p0)
}
//go:cgo_export_dynamic proxyholochain_GobEntry_Unmarshal
//go:linkname _cgoexp_ebd397278953_proxyholochain_GobEntry_Unmarshal _cgoexp_ebd397278953_proxyholochain_GobEntry_Unmarshal
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GobEntry_Unmarshal
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GobEntry_Unmarshal(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GobEntry_Unmarshal
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GobEntry_Unmarshal(p0 _Ctype_int32_t, p1 _Ctype_struct_nbyteslice) (r0 _Ctype_int32_t) {
	return proxyholochain_GobEntry_Unmarshal(p0, p1)
}
//go:cgo_export_dynamic new_holochain_GobEntry
//go:linkname _cgoexp_ebd397278953_new_holochain_GobEntry _cgoexp_ebd397278953_new_holochain_GobEntry
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_GobEntry
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_GobEntry(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_GobEntry
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_GobEntry() (r0 _Ctype_int32_t) {
	return new_holochain_GobEntry()
}
//go:cgo_export_dynamic new_holochain_Gossip
//go:linkname _cgoexp_ebd397278953_new_holochain_Gossip _cgoexp_ebd397278953_new_holochain_Gossip
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Gossip
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Gossip(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Gossip
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Gossip() (r0 _Ctype_int32_t) {
	return new_holochain_Gossip()
}
//go:cgo_export_dynamic proxyholochain_GossipReq_MyIdx_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_GossipReq_MyIdx_Set _cgoexp_ebd397278953_proxyholochain_GossipReq_MyIdx_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GossipReq_MyIdx_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GossipReq_MyIdx_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GossipReq_MyIdx_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GossipReq_MyIdx_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_GossipReq_MyIdx_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_GossipReq_MyIdx_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_GossipReq_MyIdx_Get _cgoexp_ebd397278953_proxyholochain_GossipReq_MyIdx_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GossipReq_MyIdx_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GossipReq_MyIdx_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GossipReq_MyIdx_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GossipReq_MyIdx_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_GossipReq_MyIdx_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_GossipReq_YourIdx_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_GossipReq_YourIdx_Set _cgoexp_ebd397278953_proxyholochain_GossipReq_YourIdx_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GossipReq_YourIdx_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GossipReq_YourIdx_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GossipReq_YourIdx_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GossipReq_YourIdx_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_GossipReq_YourIdx_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_GossipReq_YourIdx_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_GossipReq_YourIdx_Get _cgoexp_ebd397278953_proxyholochain_GossipReq_YourIdx_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_GossipReq_YourIdx_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_GossipReq_YourIdx_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_GossipReq_YourIdx_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_GossipReq_YourIdx_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_GossipReq_YourIdx_Get(p0)
}
//go:cgo_export_dynamic new_holochain_GossipReq
//go:linkname _cgoexp_ebd397278953_new_holochain_GossipReq _cgoexp_ebd397278953_new_holochain_GossipReq
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_GossipReq
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_GossipReq(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_GossipReq
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_GossipReq() (r0 _Ctype_int32_t) {
	return new_holochain_GossipReq()
}
//go:cgo_export_dynamic proxyholochain_Header_Type_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Header_Type_Set _cgoexp_ebd397278953_proxyholochain_Header_Type_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Header_Type_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Header_Type_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Header_Type_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Header_Type_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Header_Type_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Header_Type_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Header_Type_Get _cgoexp_ebd397278953_proxyholochain_Header_Type_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Header_Type_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Header_Type_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Header_Type_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Header_Type_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Header_Type_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Header_Marshal
//go:linkname _cgoexp_ebd397278953_proxyholochain_Header_Marshal _cgoexp_ebd397278953_proxyholochain_Header_Marshal
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Header_Marshal
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Header_Marshal(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Header_Marshal
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Header_Marshal(p0 _Ctype_int32_t) (r0 _Ctype_struct_nbyteslice, r1 _Ctype_int32_t) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Header_Marshal(p0)
}
//go:cgo_export_dynamic proxyholochain_Header_Unmarshal
//go:linkname _cgoexp_ebd397278953_proxyholochain_Header_Unmarshal _cgoexp_ebd397278953_proxyholochain_Header_Unmarshal
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Header_Unmarshal
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Header_Unmarshal(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Header_Unmarshal
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Header_Unmarshal(p0 _Ctype_int32_t, p1 _Ctype_struct_nbyteslice, p2 _Ctype_nint) (r0 _Ctype_int32_t) {
	return proxyholochain_Header_Unmarshal(p0, p1, p2)
}
//go:cgo_export_dynamic new_holochain_Header
//go:linkname _cgoexp_ebd397278953_new_holochain_Header _cgoexp_ebd397278953_new_holochain_Header
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Header
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Header(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Header
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Header() (r0 _Ctype_int32_t) {
	return new_holochain_Header()
}
//go:cgo_export_dynamic proxyholochain_Holochain_Activate
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_Activate _cgoexp_ebd397278953_proxyholochain_Holochain_Activate
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_Activate
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_Activate(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Activate
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Activate(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_Holochain_Activate(p0)
}
//go:cgo_export_dynamic proxyholochain_Holochain_Agent
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_Agent _cgoexp_ebd397278953_proxyholochain_Holochain_Agent
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_Agent
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_Agent(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Agent
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Agent(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_Holochain_Agent(p0)
}
//go:cgo_export_dynamic proxyholochain_Holochain_BSget
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_BSget _cgoexp_ebd397278953_proxyholochain_Holochain_BSget
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_BSget
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_BSget(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_BSget
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_BSget(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_Holochain_BSget(p0)
}
//go:cgo_export_dynamic proxyholochain_Holochain_BSpost
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_BSpost _cgoexp_ebd397278953_proxyholochain_Holochain_BSpost
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_BSpost
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_BSpost(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_BSpost
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_BSpost(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_Holochain_BSpost(p0)
}
//go:cgo_export_dynamic proxyholochain_Holochain_BuildBridge
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_BuildBridge _cgoexp_ebd397278953_proxyholochain_Holochain_BuildBridge
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_BuildBridge
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_BuildBridge(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_BuildBridge
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_BuildBridge(p0 _Ctype_int32_t, p1 _Ctype_int32_t, p2 _Ctype_struct_nstring) (r0 _Ctype_int32_t) {
	return proxyholochain_Holochain_BuildBridge(p0, p1, p2)
}
//go:cgo_export_dynamic proxyholochain_Holochain_Chain
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_Chain _cgoexp_ebd397278953_proxyholochain_Holochain_Chain
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_Chain
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_Chain(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Chain
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Chain(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_Holochain_Chain(p0)
}
//go:cgo_export_dynamic proxyholochain_Holochain_Close
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_Close _cgoexp_ebd397278953_proxyholochain_Holochain_Close
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_Close
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_Close(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Close
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Close(p0 _Ctype_int32_t) {
	proxyholochain_Holochain_Close(p0)
}
//go:cgo_export_dynamic proxyholochain_Holochain_DBPath
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_DBPath _cgoexp_ebd397278953_proxyholochain_Holochain_DBPath
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_DBPath
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_DBPath(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_DBPath
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_DBPath(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Holochain_DBPath(p0)
}
//go:cgo_export_dynamic proxyholochain_Holochain_DHT
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_DHT _cgoexp_ebd397278953_proxyholochain_Holochain_DHT
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_DHT
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_DHT(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_DHT
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_DHT(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_Holochain_DHT(p0)
}
//go:cgo_export_dynamic proxyholochain_Holochain_DNAPath
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_DNAPath _cgoexp_ebd397278953_proxyholochain_Holochain_DNAPath
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_DNAPath
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_DNAPath(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_DNAPath
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_DNAPath(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Holochain_DNAPath(p0)
}
//go:cgo_export_dynamic proxyholochain_Holochain_Debug
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_Debug _cgoexp_ebd397278953_proxyholochain_Holochain_Debug
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_Debug
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_Debug(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Debug
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Debug(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Holochain_Debug(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Holochain_GetEntryDef
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_GetEntryDef _cgoexp_ebd397278953_proxyholochain_Holochain_GetEntryDef
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_GetEntryDef
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_GetEntryDef(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_GetEntryDef
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_GetEntryDef(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain_Holochain_GetEntryDef(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Holochain_GetProperty
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_GetProperty _cgoexp_ebd397278953_proxyholochain_Holochain_GetProperty
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_GetProperty
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_GetProperty(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_GetProperty
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_GetProperty(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) (r0 _Ctype_struct_nstring, r1 _Ctype_int32_t) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Holochain_GetProperty(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Holochain_GetZome
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_GetZome _cgoexp_ebd397278953_proxyholochain_Holochain_GetZome
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_GetZome
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_GetZome(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_GetZome
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_GetZome(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain_Holochain_GetZome(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Holochain_GetZomeForEntryType
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_GetZomeForEntryType _cgoexp_ebd397278953_proxyholochain_Holochain_GetZomeForEntryType
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_GetZomeForEntryType
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_GetZomeForEntryType(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_GetZomeForEntryType
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_GetZomeForEntryType(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain_Holochain_GetZomeForEntryType(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Holochain_HandleAsyncSends
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_HandleAsyncSends _cgoexp_ebd397278953_proxyholochain_Holochain_HandleAsyncSends
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_HandleAsyncSends
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_HandleAsyncSends(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_HandleAsyncSends
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_HandleAsyncSends(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_Holochain_HandleAsyncSends(p0)
}
//go:cgo_export_dynamic proxyholochain_Holochain_MakeRibosome
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_MakeRibosome _cgoexp_ebd397278953_proxyholochain_Holochain_MakeRibosome
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_MakeRibosome
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_MakeRibosome(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_MakeRibosome
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_MakeRibosome(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain_Holochain_MakeRibosome(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Holochain_Name
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_Name _cgoexp_ebd397278953_proxyholochain_Holochain_Name
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_Name
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_Name(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Name
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Name(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Holochain_Name(p0)
}
//go:cgo_export_dynamic proxyholochain_Holochain_Node
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_Node _cgoexp_ebd397278953_proxyholochain_Holochain_Node
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_Node
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_Node(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Node
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Node(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_Holochain_Node(p0)
}
//go:cgo_export_dynamic proxyholochain_Holochain_NodeIDStr
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_NodeIDStr _cgoexp_ebd397278953_proxyholochain_Holochain_NodeIDStr
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_NodeIDStr
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_NodeIDStr(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_NodeIDStr
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_NodeIDStr(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Holochain_NodeIDStr(p0)
}
//go:cgo_export_dynamic proxyholochain_Holochain_Nucleus
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_Nucleus _cgoexp_ebd397278953_proxyholochain_Holochain_Nucleus
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_Nucleus
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_Nucleus(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Nucleus
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Nucleus(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_Holochain_Nucleus(p0)
}
//go:cgo_export_dynamic proxyholochain_Holochain_Prepare
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_Prepare _cgoexp_ebd397278953_proxyholochain_Holochain_Prepare
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_Prepare
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_Prepare(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Prepare
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Prepare(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_Holochain_Prepare(p0)
}
//go:cgo_export_dynamic proxyholochain_Holochain_PrepareHashType
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_PrepareHashType _cgoexp_ebd397278953_proxyholochain_Holochain_PrepareHashType
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_PrepareHashType
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_PrepareHashType(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_PrepareHashType
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_PrepareHashType(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_Holochain_PrepareHashType(p0)
}
//go:cgo_export_dynamic proxyholochain_Holochain_Reset
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_Reset _cgoexp_ebd397278953_proxyholochain_Holochain_Reset
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_Reset
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_Reset(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Reset
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Reset(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_Holochain_Reset(p0)
}
//go:cgo_export_dynamic proxyholochain_Holochain_RootPath
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_RootPath _cgoexp_ebd397278953_proxyholochain_Holochain_RootPath
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_RootPath
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_RootPath(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_RootPath
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_RootPath(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Holochain_RootPath(p0)
}
//go:cgo_export_dynamic proxyholochain_Holochain_Sign
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_Sign _cgoexp_ebd397278953_proxyholochain_Holochain_Sign
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_Sign
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_Sign(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Sign
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Sign(p0 _Ctype_int32_t, p1 _Ctype_struct_nbyteslice) (r0 _Ctype_struct_nbyteslice, r1 _Ctype_int32_t) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Holochain_Sign(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Holochain_StartBackgroundTasks
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_StartBackgroundTasks _cgoexp_ebd397278953_proxyholochain_Holochain_StartBackgroundTasks
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_StartBackgroundTasks
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_StartBackgroundTasks(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_StartBackgroundTasks
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_StartBackgroundTasks(p0 _Ctype_int32_t) {
	proxyholochain_Holochain_StartBackgroundTasks(p0)
}
//go:cgo_export_dynamic proxyholochain_Holochain_Started
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_Started _cgoexp_ebd397278953_proxyholochain_Holochain_Started
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_Started
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_Started(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Started
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_Started(p0 _Ctype_int32_t) (r0 _Ctype_char) {
	return proxyholochain_Holochain_Started(p0)
}
//go:cgo_export_dynamic proxyholochain_Holochain_TestPath
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_TestPath _cgoexp_ebd397278953_proxyholochain_Holochain_TestPath
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_TestPath
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_TestPath(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_TestPath
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_TestPath(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Holochain_TestPath(p0)
}
//go:cgo_export_dynamic proxyholochain_Holochain_UIPath
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_UIPath _cgoexp_ebd397278953_proxyholochain_Holochain_UIPath
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_UIPath
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_UIPath(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_UIPath
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_UIPath(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Holochain_UIPath(p0)
}
//go:cgo_export_dynamic proxyholochain_Holochain_ZomePath
//go:linkname _cgoexp_ebd397278953_proxyholochain_Holochain_ZomePath _cgoexp_ebd397278953_proxyholochain_Holochain_ZomePath
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Holochain_ZomePath
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Holochain_ZomePath(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Holochain_ZomePath
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Holochain_ZomePath(p0 _Ctype_int32_t, p1 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Holochain_ZomePath(p0, p1)
}
//go:cgo_export_dynamic new_holochain_Holochain
//go:linkname _cgoexp_ebd397278953_new_holochain_Holochain _cgoexp_ebd397278953_new_holochain_Holochain
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Holochain
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Holochain(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Holochain
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Holochain() (r0 _Ctype_int32_t) {
	return new_holochain_Holochain()
}
//go:cgo_export_dynamic proxyholochain_JSONEntry_Marshal
//go:linkname _cgoexp_ebd397278953_proxyholochain_JSONEntry_Marshal _cgoexp_ebd397278953_proxyholochain_JSONEntry_Marshal
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_JSONEntry_Marshal
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_JSONEntry_Marshal(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_JSONEntry_Marshal
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_JSONEntry_Marshal(p0 _Ctype_int32_t) (r0 _Ctype_struct_nbyteslice, r1 _Ctype_int32_t) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_JSONEntry_Marshal(p0)
}
//go:cgo_export_dynamic proxyholochain_JSONEntry_Unmarshal
//go:linkname _cgoexp_ebd397278953_proxyholochain_JSONEntry_Unmarshal _cgoexp_ebd397278953_proxyholochain_JSONEntry_Unmarshal
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_JSONEntry_Unmarshal
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_JSONEntry_Unmarshal(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_JSONEntry_Unmarshal
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_JSONEntry_Unmarshal(p0 _Ctype_int32_t, p1 _Ctype_struct_nbyteslice) (r0 _Ctype_int32_t) {
	return proxyholochain_JSONEntry_Unmarshal(p0, p1)
}
//go:cgo_export_dynamic new_holochain_JSONEntry
//go:linkname _cgoexp_ebd397278953_new_holochain_JSONEntry _cgoexp_ebd397278953_new_holochain_JSONEntry
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_JSONEntry
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_JSONEntry(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_JSONEntry
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_JSONEntry() (r0 _Ctype_int32_t) {
	return new_holochain_JSONEntry()
}
//go:cgo_export_dynamic new_holochain_JSONSchemaValidator
//go:linkname _cgoexp_ebd397278953_new_holochain_JSONSchemaValidator _cgoexp_ebd397278953_new_holochain_JSONSchemaValidator
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_JSONSchemaValidator
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_JSONSchemaValidator(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_JSONSchemaValidator
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_JSONSchemaValidator() (r0 _Ctype_int32_t) {
	return new_holochain_JSONSchemaValidator()
}
//go:cgo_export_dynamic proxyholochain_JSRibosome_ChainGenesis
//go:linkname _cgoexp_ebd397278953_proxyholochain_JSRibosome_ChainGenesis _cgoexp_ebd397278953_proxyholochain_JSRibosome_ChainGenesis
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_JSRibosome_ChainGenesis
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_JSRibosome_ChainGenesis(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_JSRibosome_ChainGenesis
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_JSRibosome_ChainGenesis(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_JSRibosome_ChainGenesis(p0)
}
//go:cgo_export_dynamic proxyholochain_JSRibosome_Receive
//go:linkname _cgoexp_ebd397278953_proxyholochain_JSRibosome_Receive _cgoexp_ebd397278953_proxyholochain_JSRibosome_Receive
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_JSRibosome_Receive
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_JSRibosome_Receive(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_JSRibosome_Receive
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_JSRibosome_Receive(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring, p2 _Ctype_struct_nstring) (r0 _Ctype_struct_nstring, r1 _Ctype_int32_t) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_JSRibosome_Receive(p0, p1, p2)
}
//go:cgo_export_dynamic proxyholochain_JSRibosome_Type
//go:linkname _cgoexp_ebd397278953_proxyholochain_JSRibosome_Type _cgoexp_ebd397278953_proxyholochain_JSRibosome_Type
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_JSRibosome_Type
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_JSRibosome_Type(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_JSRibosome_Type
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_JSRibosome_Type(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_JSRibosome_Type(p0)
}
//go:cgo_export_dynamic new_holochain_JSRibosome
//go:linkname _cgoexp_ebd397278953_new_holochain_JSRibosome _cgoexp_ebd397278953_new_holochain_JSRibosome
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_JSRibosome
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_JSRibosome(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_JSRibosome
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_JSRibosome() (r0 _Ctype_int32_t) {
	return new_holochain_JSRibosome()
}
//go:cgo_export_dynamic new_holochain_LibP2PAgent
//go:linkname _cgoexp_ebd397278953_new_holochain_LibP2PAgent _cgoexp_ebd397278953_new_holochain_LibP2PAgent
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_LibP2PAgent
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_LibP2PAgent(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_LibP2PAgent
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_LibP2PAgent() (r0 _Ctype_int32_t) {
	return new_holochain_LibP2PAgent()
}
//go:cgo_export_dynamic proxyholochain_Link_LinkAction_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Link_LinkAction_Set _cgoexp_ebd397278953_proxyholochain_Link_LinkAction_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Link_LinkAction_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Link_LinkAction_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Link_LinkAction_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Link_LinkAction_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Link_LinkAction_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Link_LinkAction_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Link_LinkAction_Get _cgoexp_ebd397278953_proxyholochain_Link_LinkAction_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Link_LinkAction_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Link_LinkAction_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Link_LinkAction_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Link_LinkAction_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Link_LinkAction_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Link_Base_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Link_Base_Set _cgoexp_ebd397278953_proxyholochain_Link_Base_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Link_Base_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Link_Base_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Link_Base_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Link_Base_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Link_Base_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Link_Base_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Link_Base_Get _cgoexp_ebd397278953_proxyholochain_Link_Base_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Link_Base_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Link_Base_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Link_Base_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Link_Base_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Link_Base_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Link_Link_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Link_Link_Set _cgoexp_ebd397278953_proxyholochain_Link_Link_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Link_Link_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Link_Link_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Link_Link_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Link_Link_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Link_Link_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Link_Link_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Link_Link_Get _cgoexp_ebd397278953_proxyholochain_Link_Link_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Link_Link_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Link_Link_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Link_Link_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Link_Link_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Link_Link_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Link_Tag_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Link_Tag_Set _cgoexp_ebd397278953_proxyholochain_Link_Tag_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Link_Tag_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Link_Tag_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Link_Tag_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Link_Tag_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Link_Tag_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Link_Tag_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Link_Tag_Get _cgoexp_ebd397278953_proxyholochain_Link_Tag_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Link_Tag_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Link_Tag_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Link_Tag_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Link_Tag_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Link_Tag_Get(p0)
}
//go:cgo_export_dynamic new_holochain_Link
//go:linkname _cgoexp_ebd397278953_new_holochain_Link _cgoexp_ebd397278953_new_holochain_Link
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Link
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Link(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Link
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Link() (r0 _Ctype_int32_t) {
	return new_holochain_Link()
}
//go:cgo_export_dynamic proxyholochain_LinkEvent_Status_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_LinkEvent_Status_Set _cgoexp_ebd397278953_proxyholochain_LinkEvent_Status_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_LinkEvent_Status_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_LinkEvent_Status_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_LinkEvent_Status_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_LinkEvent_Status_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_LinkEvent_Status_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_LinkEvent_Status_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_LinkEvent_Status_Get _cgoexp_ebd397278953_proxyholochain_LinkEvent_Status_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_LinkEvent_Status_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_LinkEvent_Status_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_LinkEvent_Status_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_LinkEvent_Status_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_LinkEvent_Status_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_LinkEvent_Source_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_LinkEvent_Source_Set _cgoexp_ebd397278953_proxyholochain_LinkEvent_Source_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_LinkEvent_Source_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_LinkEvent_Source_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_LinkEvent_Source_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_LinkEvent_Source_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_LinkEvent_Source_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_LinkEvent_Source_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_LinkEvent_Source_Get _cgoexp_ebd397278953_proxyholochain_LinkEvent_Source_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_LinkEvent_Source_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_LinkEvent_Source_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_LinkEvent_Source_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_LinkEvent_Source_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_LinkEvent_Source_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_LinkEvent_LinksEntry_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_LinkEvent_LinksEntry_Set _cgoexp_ebd397278953_proxyholochain_LinkEvent_LinksEntry_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_LinkEvent_LinksEntry_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_LinkEvent_LinksEntry_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_LinkEvent_LinksEntry_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_LinkEvent_LinksEntry_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_LinkEvent_LinksEntry_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_LinkEvent_LinksEntry_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_LinkEvent_LinksEntry_Get _cgoexp_ebd397278953_proxyholochain_LinkEvent_LinksEntry_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_LinkEvent_LinksEntry_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_LinkEvent_LinksEntry_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_LinkEvent_LinksEntry_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_LinkEvent_LinksEntry_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_LinkEvent_LinksEntry_Get(p0)
}
//go:cgo_export_dynamic new_holochain_LinkEvent
//go:linkname _cgoexp_ebd397278953_new_holochain_LinkEvent _cgoexp_ebd397278953_new_holochain_LinkEvent
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_LinkEvent
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_LinkEvent(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_LinkEvent
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_LinkEvent() (r0 _Ctype_int32_t) {
	return new_holochain_LinkEvent()
}
//go:cgo_export_dynamic proxyholochain_LinkQuery_T_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_LinkQuery_T_Set _cgoexp_ebd397278953_proxyholochain_LinkQuery_T_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_LinkQuery_T_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_LinkQuery_T_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_LinkQuery_T_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_LinkQuery_T_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_LinkQuery_T_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_LinkQuery_T_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_LinkQuery_T_Get _cgoexp_ebd397278953_proxyholochain_LinkQuery_T_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_LinkQuery_T_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_LinkQuery_T_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_LinkQuery_T_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_LinkQuery_T_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_LinkQuery_T_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_LinkQuery_StatusMask_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_LinkQuery_StatusMask_Set _cgoexp_ebd397278953_proxyholochain_LinkQuery_StatusMask_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_LinkQuery_StatusMask_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_LinkQuery_StatusMask_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_LinkQuery_StatusMask_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_LinkQuery_StatusMask_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_LinkQuery_StatusMask_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_LinkQuery_StatusMask_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_LinkQuery_StatusMask_Get _cgoexp_ebd397278953_proxyholochain_LinkQuery_StatusMask_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_LinkQuery_StatusMask_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_LinkQuery_StatusMask_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_LinkQuery_StatusMask_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_LinkQuery_StatusMask_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_LinkQuery_StatusMask_Get(p0)
}
//go:cgo_export_dynamic new_holochain_LinkQuery
//go:linkname _cgoexp_ebd397278953_new_holochain_LinkQuery _cgoexp_ebd397278953_new_holochain_LinkQuery
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_LinkQuery
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_LinkQuery(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_LinkQuery
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_LinkQuery() (r0 _Ctype_int32_t) {
	return new_holochain_LinkQuery()
}
//go:cgo_export_dynamic new_holochain_LinkQueryResp
//go:linkname _cgoexp_ebd397278953_new_holochain_LinkQueryResp _cgoexp_ebd397278953_new_holochain_LinkQueryResp
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_LinkQueryResp
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_LinkQueryResp(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_LinkQueryResp
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_LinkQueryResp() (r0 _Ctype_int32_t) {
	return new_holochain_LinkQueryResp()
}
//go:cgo_export_dynamic new_holochain_LinkReq
//go:linkname _cgoexp_ebd397278953_new_holochain_LinkReq _cgoexp_ebd397278953_new_holochain_LinkReq
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_LinkReq
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_LinkReq(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_LinkReq
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_LinkReq() (r0 _Ctype_int32_t) {
	return new_holochain_LinkReq()
}
//go:cgo_export_dynamic new_holochain_LinksEntry
//go:linkname _cgoexp_ebd397278953_new_holochain_LinksEntry _cgoexp_ebd397278953_new_holochain_LinksEntry
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_LinksEntry
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_LinksEntry(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_LinksEntry
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_LinksEntry() (r0 _Ctype_int32_t) {
	return new_holochain_LinksEntry()
}
//go:cgo_export_dynamic proxyholochain_ListAddReq_ListType_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_ListAddReq_ListType_Set _cgoexp_ebd397278953_proxyholochain_ListAddReq_ListType_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ListAddReq_ListType_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ListAddReq_ListType_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ListAddReq_ListType_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ListAddReq_ListType_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_ListAddReq_ListType_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ListAddReq_ListType_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_ListAddReq_ListType_Get _cgoexp_ebd397278953_proxyholochain_ListAddReq_ListType_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ListAddReq_ListType_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ListAddReq_ListType_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ListAddReq_ListType_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ListAddReq_ListType_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ListAddReq_ListType_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_ListAddReq_WarrantType_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_ListAddReq_WarrantType_Set _cgoexp_ebd397278953_proxyholochain_ListAddReq_WarrantType_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ListAddReq_WarrantType_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ListAddReq_WarrantType_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ListAddReq_WarrantType_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ListAddReq_WarrantType_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_ListAddReq_WarrantType_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ListAddReq_WarrantType_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_ListAddReq_WarrantType_Get _cgoexp_ebd397278953_proxyholochain_ListAddReq_WarrantType_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ListAddReq_WarrantType_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ListAddReq_WarrantType_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ListAddReq_WarrantType_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ListAddReq_WarrantType_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_ListAddReq_WarrantType_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_ListAddReq_Warrant_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_ListAddReq_Warrant_Set _cgoexp_ebd397278953_proxyholochain_ListAddReq_Warrant_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ListAddReq_Warrant_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ListAddReq_Warrant_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ListAddReq_Warrant_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ListAddReq_Warrant_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nbyteslice) {
	proxyholochain_ListAddReq_Warrant_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ListAddReq_Warrant_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_ListAddReq_Warrant_Get _cgoexp_ebd397278953_proxyholochain_ListAddReq_Warrant_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ListAddReq_Warrant_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ListAddReq_Warrant_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ListAddReq_Warrant_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ListAddReq_Warrant_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nbyteslice) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ListAddReq_Warrant_Get(p0)
}
//go:cgo_export_dynamic new_holochain_ListAddReq
//go:linkname _cgoexp_ebd397278953_new_holochain_ListAddReq _cgoexp_ebd397278953_new_holochain_ListAddReq
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ListAddReq
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ListAddReq(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ListAddReq
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ListAddReq() (r0 _Ctype_int32_t) {
	return new_holochain_ListAddReq()
}
//go:cgo_export_dynamic proxyholochain_Logger_Name_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Logger_Name_Set _cgoexp_ebd397278953_proxyholochain_Logger_Name_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Logger_Name_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Logger_Name_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Logger_Name_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Logger_Name_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Logger_Name_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Logger_Name_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Logger_Name_Get _cgoexp_ebd397278953_proxyholochain_Logger_Name_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Logger_Name_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Logger_Name_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Logger_Name_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Logger_Name_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Logger_Name_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Logger_Enabled_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Logger_Enabled_Set _cgoexp_ebd397278953_proxyholochain_Logger_Enabled_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Logger_Enabled_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Logger_Enabled_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Logger_Enabled_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Logger_Enabled_Set(p0 _Ctype_int32_t, p1 _Ctype_char) {
	proxyholochain_Logger_Enabled_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Logger_Enabled_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Logger_Enabled_Get _cgoexp_ebd397278953_proxyholochain_Logger_Enabled_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Logger_Enabled_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Logger_Enabled_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Logger_Enabled_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Logger_Enabled_Get(p0 _Ctype_int32_t) (r0 _Ctype_char) {
	return proxyholochain_Logger_Enabled_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Logger_Format_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Logger_Format_Set _cgoexp_ebd397278953_proxyholochain_Logger_Format_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Logger_Format_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Logger_Format_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Logger_Format_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Logger_Format_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Logger_Format_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Logger_Format_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Logger_Format_Get _cgoexp_ebd397278953_proxyholochain_Logger_Format_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Logger_Format_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Logger_Format_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Logger_Format_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Logger_Format_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Logger_Format_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Logger_Prefix_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Logger_Prefix_Set _cgoexp_ebd397278953_proxyholochain_Logger_Prefix_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Logger_Prefix_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Logger_Prefix_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Logger_Prefix_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Logger_Prefix_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Logger_Prefix_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Logger_Prefix_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Logger_Prefix_Get _cgoexp_ebd397278953_proxyholochain_Logger_Prefix_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Logger_Prefix_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Logger_Prefix_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Logger_Prefix_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Logger_Prefix_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Logger_Prefix_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Logger_SetPrefix
//go:linkname _cgoexp_ebd397278953_proxyholochain_Logger_SetPrefix _cgoexp_ebd397278953_proxyholochain_Logger_SetPrefix
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Logger_SetPrefix
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Logger_SetPrefix(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Logger_SetPrefix
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Logger_SetPrefix(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Logger_SetPrefix(p0, p1)
}
//go:cgo_export_dynamic new_holochain_Logger
//go:linkname _cgoexp_ebd397278953_new_holochain_Logger _cgoexp_ebd397278953_new_holochain_Logger
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Logger
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Logger(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Logger
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Logger() (r0 _Ctype_int32_t) {
	return new_holochain_Logger()
}
//go:cgo_export_dynamic new_holochain_Loggers
//go:linkname _cgoexp_ebd397278953_new_holochain_Loggers _cgoexp_ebd397278953_new_holochain_Loggers
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Loggers
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Loggers(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Loggers
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Loggers() (r0 _Ctype_int32_t) {
	return new_holochain_Loggers()
}
//go:cgo_export_dynamic proxyholochain_Message_Encode
//go:linkname _cgoexp_ebd397278953_proxyholochain_Message_Encode _cgoexp_ebd397278953_proxyholochain_Message_Encode
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Message_Encode
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Message_Encode(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Message_Encode
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Message_Encode(p0 _Ctype_int32_t) (r0 _Ctype_struct_nbyteslice, r1 _Ctype_int32_t) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Message_Encode(p0)
}
//go:cgo_export_dynamic proxyholochain_Message_String
//go:linkname _cgoexp_ebd397278953_proxyholochain_Message_String _cgoexp_ebd397278953_proxyholochain_Message_String
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Message_String
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Message_String(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Message_String
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Message_String(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Message_String(p0)
}
//go:cgo_export_dynamic new_holochain_Message
//go:linkname _cgoexp_ebd397278953_new_holochain_Message _cgoexp_ebd397278953_new_holochain_Message
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Message
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Message(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Message
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Message() (r0 _Ctype_int32_t) {
	return new_holochain_Message()
}
//go:cgo_export_dynamic proxyholochain_Meta_T_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Meta_T_Set _cgoexp_ebd397278953_proxyholochain_Meta_T_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Meta_T_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Meta_T_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Meta_T_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Meta_T_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Meta_T_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Meta_T_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Meta_T_Get _cgoexp_ebd397278953_proxyholochain_Meta_T_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Meta_T_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Meta_T_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Meta_T_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Meta_T_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Meta_T_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Meta_V_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Meta_V_Set _cgoexp_ebd397278953_proxyholochain_Meta_V_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Meta_V_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Meta_V_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Meta_V_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Meta_V_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nbyteslice) {
	proxyholochain_Meta_V_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Meta_V_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Meta_V_Get _cgoexp_ebd397278953_proxyholochain_Meta_V_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Meta_V_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Meta_V_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Meta_V_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Meta_V_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nbyteslice) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Meta_V_Get(p0)
}
//go:cgo_export_dynamic new_holochain_Meta
//go:linkname _cgoexp_ebd397278953_new_holochain_Meta _cgoexp_ebd397278953_new_holochain_Meta
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Meta
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Meta(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Meta
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Meta() (r0 _Ctype_int32_t) {
	return new_holochain_Meta()
}
//go:cgo_export_dynamic proxyholochain_ModAgentOptions_Identity_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Identity_Set _cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Identity_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Identity_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Identity_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ModAgentOptions_Identity_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ModAgentOptions_Identity_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_ModAgentOptions_Identity_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ModAgentOptions_Identity_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Identity_Get _cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Identity_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Identity_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Identity_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ModAgentOptions_Identity_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ModAgentOptions_Identity_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ModAgentOptions_Identity_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_ModAgentOptions_Revocation_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Revocation_Set _cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Revocation_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Revocation_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Revocation_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ModAgentOptions_Revocation_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ModAgentOptions_Revocation_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_ModAgentOptions_Revocation_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ModAgentOptions_Revocation_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Revocation_Get _cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Revocation_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Revocation_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Revocation_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ModAgentOptions_Revocation_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ModAgentOptions_Revocation_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ModAgentOptions_Revocation_Get(p0)
}
//go:cgo_export_dynamic new_holochain_ModAgentOptions
//go:linkname _cgoexp_ebd397278953_new_holochain_ModAgentOptions _cgoexp_ebd397278953_new_holochain_ModAgentOptions
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ModAgentOptions
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ModAgentOptions(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ModAgentOptions
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ModAgentOptions() (r0 _Ctype_int32_t) {
	return new_holochain_ModAgentOptions()
}
//go:cgo_export_dynamic new_holochain_ModReq
//go:linkname _cgoexp_ebd397278953_new_holochain_ModReq _cgoexp_ebd397278953_new_holochain_ModReq
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ModReq
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ModReq(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ModReq
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ModReq() (r0 _Ctype_int32_t) {
	return new_holochain_ModReq()
}
//go:cgo_export_dynamic proxyholochain_Node_Close
//go:linkname _cgoexp_ebd397278953_proxyholochain_Node_Close _cgoexp_ebd397278953_proxyholochain_Node_Close
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Node_Close
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Node_Close(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Node_Close
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Node_Close(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_Node_Close(p0)
}
//go:cgo_export_dynamic proxyholochain_Node_StartProtocol
//go:linkname _cgoexp_ebd397278953_proxyholochain_Node_StartProtocol _cgoexp_ebd397278953_proxyholochain_Node_StartProtocol
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Node_StartProtocol
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Node_StartProtocol(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Node_StartProtocol
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Node_StartProtocol(p0 _Ctype_int32_t, p1 _Ctype_int32_t, p2 _Ctype_nint) (r0 _Ctype_int32_t) {
	return proxyholochain_Node_StartProtocol(p0, p1, p2)
}
//go:cgo_export_dynamic new_holochain_Node
//go:linkname _cgoexp_ebd397278953_new_holochain_Node _cgoexp_ebd397278953_new_holochain_Node
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Node
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Node(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Node
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Node() (r0 _Ctype_int32_t) {
	return new_holochain_Node()
}
//go:cgo_export_dynamic proxyholochain_Nucleus_DNA
//go:linkname _cgoexp_ebd397278953_proxyholochain_Nucleus_DNA _cgoexp_ebd397278953_proxyholochain_Nucleus_DNA
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Nucleus_DNA
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Nucleus_DNA(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Nucleus_DNA
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Nucleus_DNA(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_Nucleus_DNA(p0)
}
//go:cgo_export_dynamic proxyholochain_Nucleus_RunGenesis
//go:linkname _cgoexp_ebd397278953_proxyholochain_Nucleus_RunGenesis _cgoexp_ebd397278953_proxyholochain_Nucleus_RunGenesis
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Nucleus_RunGenesis
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Nucleus_RunGenesis(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Nucleus_RunGenesis
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Nucleus_RunGenesis(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_Nucleus_RunGenesis(p0)
}
//go:cgo_export_dynamic proxyholochain_Nucleus_Start
//go:linkname _cgoexp_ebd397278953_proxyholochain_Nucleus_Start _cgoexp_ebd397278953_proxyholochain_Nucleus_Start
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Nucleus_Start
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Nucleus_Start(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Nucleus_Start
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Nucleus_Start(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_Nucleus_Start(p0)
}
//go:cgo_export_dynamic new_holochain_Nucleus
//go:linkname _cgoexp_ebd397278953_new_holochain_Nucleus _cgoexp_ebd397278953_new_holochain_Nucleus
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Nucleus
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Nucleus(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Nucleus
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Nucleus() (r0 _Ctype_int32_t) {
	return new_holochain_Nucleus()
}
//go:cgo_export_dynamic proxyholochain_Package_Chain_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Package_Chain_Set _cgoexp_ebd397278953_proxyholochain_Package_Chain_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Package_Chain_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Package_Chain_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Package_Chain_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Package_Chain_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nbyteslice) {
	proxyholochain_Package_Chain_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Package_Chain_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Package_Chain_Get _cgoexp_ebd397278953_proxyholochain_Package_Chain_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Package_Chain_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Package_Chain_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Package_Chain_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Package_Chain_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nbyteslice) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Package_Chain_Get(p0)
}
//go:cgo_export_dynamic new_holochain_Package
//go:linkname _cgoexp_ebd397278953_new_holochain_Package _cgoexp_ebd397278953_new_holochain_Package
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Package
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Package(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Package
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Package() (r0 _Ctype_int32_t) {
	return new_holochain_Package()
}
//go:cgo_export_dynamic proxyholochain_PeerInfo_ID_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_PeerInfo_ID_Set _cgoexp_ebd397278953_proxyholochain_PeerInfo_ID_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_PeerInfo_ID_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_PeerInfo_ID_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_PeerInfo_ID_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_PeerInfo_ID_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nbyteslice) {
	proxyholochain_PeerInfo_ID_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_PeerInfo_ID_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_PeerInfo_ID_Get _cgoexp_ebd397278953_proxyholochain_PeerInfo_ID_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_PeerInfo_ID_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_PeerInfo_ID_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_PeerInfo_ID_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_PeerInfo_ID_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nbyteslice) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_PeerInfo_ID_Get(p0)
}
//go:cgo_export_dynamic new_holochain_PeerInfo
//go:linkname _cgoexp_ebd397278953_new_holochain_PeerInfo _cgoexp_ebd397278953_new_holochain_PeerInfo
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_PeerInfo
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_PeerInfo(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_PeerInfo
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_PeerInfo() (r0 _Ctype_int32_t) {
	return new_holochain_PeerInfo()
}
//go:cgo_export_dynamic new_holochain_PeerList
//go:linkname _cgoexp_ebd397278953_new_holochain_PeerList _cgoexp_ebd397278953_new_holochain_PeerList
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_PeerList
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_PeerList(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_PeerList
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_PeerList() (r0 _Ctype_int32_t) {
	return new_holochain_PeerList()
}
//go:cgo_export_dynamic proxyholochain_PeerRecord_Warrant_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_PeerRecord_Warrant_Set _cgoexp_ebd397278953_proxyholochain_PeerRecord_Warrant_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_PeerRecord_Warrant_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_PeerRecord_Warrant_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_PeerRecord_Warrant_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_PeerRecord_Warrant_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_PeerRecord_Warrant_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_PeerRecord_Warrant_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_PeerRecord_Warrant_Get _cgoexp_ebd397278953_proxyholochain_PeerRecord_Warrant_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_PeerRecord_Warrant_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_PeerRecord_Warrant_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_PeerRecord_Warrant_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_PeerRecord_Warrant_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_PeerRecord_Warrant_Get(p0)
}
//go:cgo_export_dynamic new_holochain_PeerRecord
//go:linkname _cgoexp_ebd397278953_new_holochain_PeerRecord _cgoexp_ebd397278953_new_holochain_PeerRecord
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_PeerRecord
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_PeerRecord(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_PeerRecord
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_PeerRecord() (r0 _Ctype_int32_t) {
	return new_holochain_PeerRecord()
}
//go:cgo_export_dynamic proxyholochain_Progenitor_Identity_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Progenitor_Identity_Set _cgoexp_ebd397278953_proxyholochain_Progenitor_Identity_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Progenitor_Identity_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Progenitor_Identity_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Progenitor_Identity_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Progenitor_Identity_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Progenitor_Identity_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Progenitor_Identity_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Progenitor_Identity_Get _cgoexp_ebd397278953_proxyholochain_Progenitor_Identity_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Progenitor_Identity_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Progenitor_Identity_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Progenitor_Identity_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Progenitor_Identity_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Progenitor_Identity_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Progenitor_PubKey_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Progenitor_PubKey_Set _cgoexp_ebd397278953_proxyholochain_Progenitor_PubKey_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Progenitor_PubKey_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Progenitor_PubKey_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Progenitor_PubKey_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Progenitor_PubKey_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nbyteslice) {
	proxyholochain_Progenitor_PubKey_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Progenitor_PubKey_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Progenitor_PubKey_Get _cgoexp_ebd397278953_proxyholochain_Progenitor_PubKey_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Progenitor_PubKey_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Progenitor_PubKey_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Progenitor_PubKey_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Progenitor_PubKey_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nbyteslice) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Progenitor_PubKey_Get(p0)
}
//go:cgo_export_dynamic new_holochain_Progenitor
//go:linkname _cgoexp_ebd397278953_new_holochain_Progenitor _cgoexp_ebd397278953_new_holochain_Progenitor
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Progenitor
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Progenitor(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Progenitor
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Progenitor() (r0 _Ctype_int32_t) {
	return new_holochain_Progenitor()
}
//go:cgo_export_dynamic new_holochain_Protocol
//go:linkname _cgoexp_ebd397278953_new_holochain_Protocol _cgoexp_ebd397278953_new_holochain_Protocol
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Protocol
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Protocol(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Protocol
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Protocol() (r0 _Ctype_int32_t) {
	return new_holochain_Protocol()
}
//go:cgo_export_dynamic proxyholochain_Put_Idx_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Put_Idx_Set _cgoexp_ebd397278953_proxyholochain_Put_Idx_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Put_Idx_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Put_Idx_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Put_Idx_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Put_Idx_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_Put_Idx_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Put_Idx_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Put_Idx_Get _cgoexp_ebd397278953_proxyholochain_Put_Idx_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Put_Idx_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Put_Idx_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Put_Idx_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Put_Idx_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_Put_Idx_Get(p0)
}
//go:cgo_export_dynamic new_holochain_Put
//go:linkname _cgoexp_ebd397278953_new_holochain_Put _cgoexp_ebd397278953_new_holochain_Put
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Put
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Put(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Put
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Put() (r0 _Ctype_int32_t) {
	return new_holochain_Put()
}
//go:cgo_export_dynamic proxyholochain_PutReq_S_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_PutReq_S_Set _cgoexp_ebd397278953_proxyholochain_PutReq_S_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_PutReq_S_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_PutReq_S_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_PutReq_S_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_PutReq_S_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_PutReq_S_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_PutReq_S_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_PutReq_S_Get _cgoexp_ebd397278953_proxyholochain_PutReq_S_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_PutReq_S_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_PutReq_S_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_PutReq_S_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_PutReq_S_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_PutReq_S_Get(p0)
}
//go:cgo_export_dynamic new_holochain_PutReq
//go:linkname _cgoexp_ebd397278953_new_holochain_PutReq _cgoexp_ebd397278953_new_holochain_PutReq
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_PutReq
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_PutReq(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_PutReq
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_PutReq() (r0 _Ctype_int32_t) {
	return new_holochain_PutReq()
}
//go:cgo_export_dynamic proxyholochain_QueryConstrain_Contains_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Contains_Set _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Contains_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Contains_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Contains_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_QueryConstrain_Contains_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_QueryConstrain_Contains_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_QueryConstrain_Contains_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_QueryConstrain_Contains_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Contains_Get _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Contains_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Contains_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Contains_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_QueryConstrain_Contains_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_QueryConstrain_Contains_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_QueryConstrain_Contains_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_QueryConstrain_Equals_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Equals_Set _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Equals_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Equals_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Equals_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_QueryConstrain_Equals_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_QueryConstrain_Equals_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_QueryConstrain_Equals_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_QueryConstrain_Equals_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Equals_Get _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Equals_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Equals_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Equals_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_QueryConstrain_Equals_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_QueryConstrain_Equals_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_QueryConstrain_Equals_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_QueryConstrain_Matches_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Matches_Set _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Matches_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Matches_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Matches_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_QueryConstrain_Matches_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_QueryConstrain_Matches_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_QueryConstrain_Matches_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_QueryConstrain_Matches_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Matches_Get _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Matches_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Matches_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Matches_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_QueryConstrain_Matches_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_QueryConstrain_Matches_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_QueryConstrain_Matches_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_QueryConstrain_Count_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Count_Set _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Count_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Count_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Count_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_QueryConstrain_Count_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_QueryConstrain_Count_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_QueryConstrain_Count_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_QueryConstrain_Count_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Count_Get _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Count_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Count_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Count_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_QueryConstrain_Count_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_QueryConstrain_Count_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_QueryConstrain_Count_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_QueryConstrain_Page_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Page_Set _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Page_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Page_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Page_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_QueryConstrain_Page_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_QueryConstrain_Page_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_QueryConstrain_Page_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_QueryConstrain_Page_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Page_Get _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Page_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Page_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Page_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_QueryConstrain_Page_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_QueryConstrain_Page_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_QueryConstrain_Page_Get(p0)
}
//go:cgo_export_dynamic new_holochain_QueryConstrain
//go:linkname _cgoexp_ebd397278953_new_holochain_QueryConstrain _cgoexp_ebd397278953_new_holochain_QueryConstrain
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_QueryConstrain
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_QueryConstrain(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_QueryConstrain
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_QueryConstrain() (r0 _Ctype_int32_t) {
	return new_holochain_QueryConstrain()
}
//go:cgo_export_dynamic new_holochain_QueryOptions
//go:linkname _cgoexp_ebd397278953_new_holochain_QueryOptions _cgoexp_ebd397278953_new_holochain_QueryOptions
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_QueryOptions
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_QueryOptions(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_QueryOptions
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_QueryOptions() (r0 _Ctype_int32_t) {
	return new_holochain_QueryOptions()
}
//go:cgo_export_dynamic proxyholochain_QueryOrder_Ascending_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_QueryOrder_Ascending_Set _cgoexp_ebd397278953_proxyholochain_QueryOrder_Ascending_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_QueryOrder_Ascending_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_QueryOrder_Ascending_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_QueryOrder_Ascending_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_QueryOrder_Ascending_Set(p0 _Ctype_int32_t, p1 _Ctype_char) {
	proxyholochain_QueryOrder_Ascending_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_QueryOrder_Ascending_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_QueryOrder_Ascending_Get _cgoexp_ebd397278953_proxyholochain_QueryOrder_Ascending_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_QueryOrder_Ascending_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_QueryOrder_Ascending_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_QueryOrder_Ascending_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_QueryOrder_Ascending_Get(p0 _Ctype_int32_t) (r0 _Ctype_char) {
	return proxyholochain_QueryOrder_Ascending_Get(p0)
}
//go:cgo_export_dynamic new_holochain_QueryOrder
//go:linkname _cgoexp_ebd397278953_new_holochain_QueryOrder _cgoexp_ebd397278953_new_holochain_QueryOrder
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_QueryOrder
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_QueryOrder(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_QueryOrder
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_QueryOrder() (r0 _Ctype_int32_t) {
	return new_holochain_QueryOrder()
}
//go:cgo_export_dynamic proxyholochain_QueryResult_Header_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_QueryResult_Header_Set _cgoexp_ebd397278953_proxyholochain_QueryResult_Header_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_QueryResult_Header_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_QueryResult_Header_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_QueryResult_Header_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_QueryResult_Header_Set(p0 _Ctype_int32_t, p1 _Ctype_int32_t) {
	proxyholochain_QueryResult_Header_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_QueryResult_Header_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_QueryResult_Header_Get _cgoexp_ebd397278953_proxyholochain_QueryResult_Header_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_QueryResult_Header_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_QueryResult_Header_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_QueryResult_Header_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_QueryResult_Header_Get(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_QueryResult_Header_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_QueryResult_Entry_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_QueryResult_Entry_Set _cgoexp_ebd397278953_proxyholochain_QueryResult_Entry_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_QueryResult_Entry_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_QueryResult_Entry_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_QueryResult_Entry_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_QueryResult_Entry_Set(p0 _Ctype_int32_t, p1 _Ctype_int32_t) {
	proxyholochain_QueryResult_Entry_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_QueryResult_Entry_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_QueryResult_Entry_Get _cgoexp_ebd397278953_proxyholochain_QueryResult_Entry_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_QueryResult_Entry_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_QueryResult_Entry_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_QueryResult_Entry_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_QueryResult_Entry_Get(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_QueryResult_Entry_Get(p0)
}
//go:cgo_export_dynamic new_holochain_QueryResult
//go:linkname _cgoexp_ebd397278953_new_holochain_QueryResult _cgoexp_ebd397278953_new_holochain_QueryResult
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_QueryResult
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_QueryResult(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_QueryResult
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_QueryResult() (r0 _Ctype_int32_t) {
	return new_holochain_QueryResult()
}
//go:cgo_export_dynamic proxyholochain_QueryReturn_Hashes_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_QueryReturn_Hashes_Set _cgoexp_ebd397278953_proxyholochain_QueryReturn_Hashes_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_QueryReturn_Hashes_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_QueryReturn_Hashes_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_QueryReturn_Hashes_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_QueryReturn_Hashes_Set(p0 _Ctype_int32_t, p1 _Ctype_char) {
	proxyholochain_QueryReturn_Hashes_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_QueryReturn_Hashes_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_QueryReturn_Hashes_Get _cgoexp_ebd397278953_proxyholochain_QueryReturn_Hashes_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_QueryReturn_Hashes_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_QueryReturn_Hashes_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_QueryReturn_Hashes_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_QueryReturn_Hashes_Get(p0 _Ctype_int32_t) (r0 _Ctype_char) {
	return proxyholochain_QueryReturn_Hashes_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_QueryReturn_Entries_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_QueryReturn_Entries_Set _cgoexp_ebd397278953_proxyholochain_QueryReturn_Entries_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_QueryReturn_Entries_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_QueryReturn_Entries_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_QueryReturn_Entries_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_QueryReturn_Entries_Set(p0 _Ctype_int32_t, p1 _Ctype_char) {
	proxyholochain_QueryReturn_Entries_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_QueryReturn_Entries_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_QueryReturn_Entries_Get _cgoexp_ebd397278953_proxyholochain_QueryReturn_Entries_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_QueryReturn_Entries_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_QueryReturn_Entries_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_QueryReturn_Entries_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_QueryReturn_Entries_Get(p0 _Ctype_int32_t) (r0 _Ctype_char) {
	return proxyholochain_QueryReturn_Entries_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_QueryReturn_Headers_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_QueryReturn_Headers_Set _cgoexp_ebd397278953_proxyholochain_QueryReturn_Headers_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_QueryReturn_Headers_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_QueryReturn_Headers_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_QueryReturn_Headers_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_QueryReturn_Headers_Set(p0 _Ctype_int32_t, p1 _Ctype_char) {
	proxyholochain_QueryReturn_Headers_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_QueryReturn_Headers_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_QueryReturn_Headers_Get _cgoexp_ebd397278953_proxyholochain_QueryReturn_Headers_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_QueryReturn_Headers_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_QueryReturn_Headers_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_QueryReturn_Headers_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_QueryReturn_Headers_Get(p0 _Ctype_int32_t) (r0 _Ctype_char) {
	return proxyholochain_QueryReturn_Headers_Get(p0)
}
//go:cgo_export_dynamic new_holochain_QueryReturn
//go:linkname _cgoexp_ebd397278953_new_holochain_QueryReturn _cgoexp_ebd397278953_new_holochain_QueryReturn
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_QueryReturn
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_QueryReturn(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_QueryReturn
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_QueryReturn() (r0 _Ctype_int32_t) {
	return new_holochain_QueryReturn()
}
//go:cgo_export_dynamic proxyholochain_RoutingTable_IsEmpty
//go:linkname _cgoexp_ebd397278953_proxyholochain_RoutingTable_IsEmpty _cgoexp_ebd397278953_proxyholochain_RoutingTable_IsEmpty
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_RoutingTable_IsEmpty
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_RoutingTable_IsEmpty(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_RoutingTable_IsEmpty
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_RoutingTable_IsEmpty(p0 _Ctype_int32_t) (r0 _Ctype_char) {
	return proxyholochain_RoutingTable_IsEmpty(p0)
}
//go:cgo_export_dynamic proxyholochain_RoutingTable_Print
//go:linkname _cgoexp_ebd397278953_proxyholochain_RoutingTable_Print _cgoexp_ebd397278953_proxyholochain_RoutingTable_Print
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_RoutingTable_Print
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_RoutingTable_Print(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_RoutingTable_Print
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_RoutingTable_Print(p0 _Ctype_int32_t) {
	proxyholochain_RoutingTable_Print(p0)
}
//go:cgo_export_dynamic proxyholochain_RoutingTable_Size
//go:linkname _cgoexp_ebd397278953_proxyholochain_RoutingTable_Size _cgoexp_ebd397278953_proxyholochain_RoutingTable_Size
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_RoutingTable_Size
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_RoutingTable_Size(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_RoutingTable_Size
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_RoutingTable_Size(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_RoutingTable_Size(p0)
}
//go:cgo_export_dynamic new_holochain_RoutingTable
//go:linkname _cgoexp_ebd397278953_new_holochain_RoutingTable _cgoexp_ebd397278953_new_holochain_RoutingTable
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_RoutingTable
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_RoutingTable(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_RoutingTable
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_RoutingTable() (r0 _Ctype_int32_t) {
	return new_holochain_RoutingTable()
}
//go:cgo_export_dynamic proxyholochain_SelfRevocation_Data_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Data_Set _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Data_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Data_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Data_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocation_Data_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocation_Data_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nbyteslice) {
	proxyholochain_SelfRevocation_Data_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_SelfRevocation_Data_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Data_Get _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Data_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Data_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Data_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocation_Data_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocation_Data_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nbyteslice) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_SelfRevocation_Data_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_SelfRevocation_OldSig_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_SelfRevocation_OldSig_Set _cgoexp_ebd397278953_proxyholochain_SelfRevocation_OldSig_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_SelfRevocation_OldSig_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_SelfRevocation_OldSig_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocation_OldSig_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocation_OldSig_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nbyteslice) {
	proxyholochain_SelfRevocation_OldSig_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_SelfRevocation_OldSig_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_SelfRevocation_OldSig_Get _cgoexp_ebd397278953_proxyholochain_SelfRevocation_OldSig_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_SelfRevocation_OldSig_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_SelfRevocation_OldSig_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocation_OldSig_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocation_OldSig_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nbyteslice) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_SelfRevocation_OldSig_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_SelfRevocation_NewSig_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_SelfRevocation_NewSig_Set _cgoexp_ebd397278953_proxyholochain_SelfRevocation_NewSig_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_SelfRevocation_NewSig_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_SelfRevocation_NewSig_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocation_NewSig_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocation_NewSig_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nbyteslice) {
	proxyholochain_SelfRevocation_NewSig_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_SelfRevocation_NewSig_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_SelfRevocation_NewSig_Get _cgoexp_ebd397278953_proxyholochain_SelfRevocation_NewSig_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_SelfRevocation_NewSig_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_SelfRevocation_NewSig_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocation_NewSig_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocation_NewSig_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nbyteslice) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_SelfRevocation_NewSig_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_SelfRevocation_Marshal
//go:linkname _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Marshal _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Marshal
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Marshal
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Marshal(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocation_Marshal
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocation_Marshal(p0 _Ctype_int32_t) (r0 _Ctype_struct_nbyteslice, r1 _Ctype_int32_t) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_SelfRevocation_Marshal(p0)
}
//go:cgo_export_dynamic proxyholochain_SelfRevocation_Unmarshal
//go:linkname _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Unmarshal _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Unmarshal
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Unmarshal
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Unmarshal(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocation_Unmarshal
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocation_Unmarshal(p0 _Ctype_int32_t, p1 _Ctype_struct_nbyteslice) (r0 _Ctype_int32_t) {
	return proxyholochain_SelfRevocation_Unmarshal(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_SelfRevocation_Verify
//go:linkname _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Verify _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Verify
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Verify
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Verify(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocation_Verify
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocation_Verify(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_SelfRevocation_Verify(p0)
}
//go:cgo_export_dynamic new_holochain_SelfRevocation
//go:linkname _cgoexp_ebd397278953_new_holochain_SelfRevocation _cgoexp_ebd397278953_new_holochain_SelfRevocation
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_SelfRevocation
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_SelfRevocation(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_SelfRevocation
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_SelfRevocation() (r0 _Ctype_int32_t) {
	return new_holochain_SelfRevocation()
}
//go:cgo_export_dynamic proxyholochain_SelfRevocationWarrant_Decode
//go:linkname _cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Decode _cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Decode
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Decode
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Decode(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocationWarrant_Decode
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocationWarrant_Decode(p0 _Ctype_int32_t, p1 _Ctype_struct_nbyteslice) (r0 _Ctype_int32_t) {
	return proxyholochain_SelfRevocationWarrant_Decode(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_SelfRevocationWarrant_Encode
//go:linkname _cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Encode _cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Encode
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Encode
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Encode(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocationWarrant_Encode
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocationWarrant_Encode(p0 _Ctype_int32_t) (r0 _Ctype_struct_nbyteslice, r1 _Ctype_int32_t) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_SelfRevocationWarrant_Encode(p0)
}
//go:cgo_export_dynamic proxyholochain_SelfRevocationWarrant_Type
//go:linkname _cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Type _cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Type
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Type
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Type(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocationWarrant_Type
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocationWarrant_Type(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_SelfRevocationWarrant_Type(p0)
}
//go:cgo_export_dynamic proxyholochain_SelfRevocationWarrant_Verify
//go:linkname _cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Verify _cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Verify
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Verify
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Verify(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocationWarrant_Verify
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_SelfRevocationWarrant_Verify(p0 _Ctype_int32_t, p1 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_SelfRevocationWarrant_Verify(p0, p1)
}
//go:cgo_export_dynamic new_holochain_SelfRevocationWarrant
//go:linkname _cgoexp_ebd397278953_new_holochain_SelfRevocationWarrant _cgoexp_ebd397278953_new_holochain_SelfRevocationWarrant
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_SelfRevocationWarrant
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_SelfRevocationWarrant(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_SelfRevocationWarrant
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_SelfRevocationWarrant() (r0 _Ctype_int32_t) {
	return new_holochain_SelfRevocationWarrant()
}
//go:cgo_export_dynamic proxyholochain_SendOptions_Callback_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_SendOptions_Callback_Set _cgoexp_ebd397278953_proxyholochain_SendOptions_Callback_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_SendOptions_Callback_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_SendOptions_Callback_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_SendOptions_Callback_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_SendOptions_Callback_Set(p0 _Ctype_int32_t, p1 _Ctype_int32_t) {
	proxyholochain_SendOptions_Callback_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_SendOptions_Callback_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_SendOptions_Callback_Get _cgoexp_ebd397278953_proxyholochain_SendOptions_Callback_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_SendOptions_Callback_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_SendOptions_Callback_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_SendOptions_Callback_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_SendOptions_Callback_Get(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_SendOptions_Callback_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_SendOptions_Timeout_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_SendOptions_Timeout_Set _cgoexp_ebd397278953_proxyholochain_SendOptions_Timeout_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_SendOptions_Timeout_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_SendOptions_Timeout_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_SendOptions_Timeout_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_SendOptions_Timeout_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_SendOptions_Timeout_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_SendOptions_Timeout_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_SendOptions_Timeout_Get _cgoexp_ebd397278953_proxyholochain_SendOptions_Timeout_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_SendOptions_Timeout_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_SendOptions_Timeout_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_SendOptions_Timeout_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_SendOptions_Timeout_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_SendOptions_Timeout_Get(p0)
}
//go:cgo_export_dynamic new_holochain_SendOptions
//go:linkname _cgoexp_ebd397278953_new_holochain_SendOptions _cgoexp_ebd397278953_new_holochain_SendOptions
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_SendOptions
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_SendOptions(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_SendOptions
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_SendOptions() (r0 _Ctype_int32_t) {
	return new_holochain_SendOptions()
}
//go:cgo_export_dynamic proxyholochain_Service_DefaultAgent_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Service_DefaultAgent_Set _cgoexp_ebd397278953_proxyholochain_Service_DefaultAgent_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Service_DefaultAgent_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Service_DefaultAgent_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Service_DefaultAgent_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Service_DefaultAgent_Set(p0 _Ctype_int32_t, p1 _Ctype_int32_t) {
	proxyholochain_Service_DefaultAgent_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Service_DefaultAgent_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Service_DefaultAgent_Get _cgoexp_ebd397278953_proxyholochain_Service_DefaultAgent_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Service_DefaultAgent_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Service_DefaultAgent_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Service_DefaultAgent_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Service_DefaultAgent_Get(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_Service_DefaultAgent_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Service_Path_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Service_Path_Set _cgoexp_ebd397278953_proxyholochain_Service_Path_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Service_Path_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Service_Path_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Service_Path_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Service_Path_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Service_Path_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Service_Path_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Service_Path_Get _cgoexp_ebd397278953_proxyholochain_Service_Path_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Service_Path_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Service_Path_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Service_Path_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Service_Path_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Service_Path_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Service_Clone
//go:linkname _cgoexp_ebd397278953_proxyholochain_Service_Clone _cgoexp_ebd397278953_proxyholochain_Service_Clone
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Service_Clone
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Service_Clone(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Service_Clone
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Service_Clone(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring, p2 _Ctype_struct_nstring, p3 _Ctype_int32_t, p4 _Ctype_char, p5 _Ctype_char) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain_Service_Clone(p0, p1, p2, p3, p4, p5)
}
//go:cgo_export_dynamic proxyholochain_Service_GenChain
//go:linkname _cgoexp_ebd397278953_proxyholochain_Service_GenChain _cgoexp_ebd397278953_proxyholochain_Service_GenChain
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Service_GenChain
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Service_GenChain(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Service_GenChain
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Service_GenChain(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain_Service_GenChain(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Service_InitAppDir
//go:linkname _cgoexp_ebd397278953_proxyholochain_Service_InitAppDir _cgoexp_ebd397278953_proxyholochain_Service_InitAppDir
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Service_InitAppDir
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Service_InitAppDir(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Service_InitAppDir
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Service_InitAppDir(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring, p2 _Ctype_struct_nstring) (r0 _Ctype_int32_t) {
	return proxyholochain_Service_InitAppDir(p0, p1, p2)
}
//go:cgo_export_dynamic proxyholochain_Service_IsConfigured
//go:linkname _cgoexp_ebd397278953_proxyholochain_Service_IsConfigured _cgoexp_ebd397278953_proxyholochain_Service_IsConfigured
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Service_IsConfigured
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Service_IsConfigured(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Service_IsConfigured
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Service_IsConfigured(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) (r0 _Ctype_struct_nstring, r1 _Ctype_int32_t) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Service_IsConfigured(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Service_ListChains
//go:linkname _cgoexp_ebd397278953_proxyholochain_Service_ListChains _cgoexp_ebd397278953_proxyholochain_Service_ListChains
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Service_ListChains
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Service_ListChains(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Service_ListChains
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Service_ListChains(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Service_ListChains(p0)
}
//go:cgo_export_dynamic proxyholochain_Service_Load
//go:linkname _cgoexp_ebd397278953_proxyholochain_Service_Load _cgoexp_ebd397278953_proxyholochain_Service_Load
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Service_Load
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Service_Load(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Service_Load
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Service_Load(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain_Service_Load(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Service_MakeAppPackage
//go:linkname _cgoexp_ebd397278953_proxyholochain_Service_MakeAppPackage _cgoexp_ebd397278953_proxyholochain_Service_MakeAppPackage
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Service_MakeAppPackage
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Service_MakeAppPackage(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Service_MakeAppPackage
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Service_MakeAppPackage(p0 _Ctype_int32_t, p1 _Ctype_int32_t) (r0 _Ctype_struct_nbyteslice, r1 _Ctype_int32_t) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Service_MakeAppPackage(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Service_MakeTestingApp
//go:linkname _cgoexp_ebd397278953_proxyholochain_Service_MakeTestingApp _cgoexp_ebd397278953_proxyholochain_Service_MakeTestingApp
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Service_MakeTestingApp
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Service_MakeTestingApp(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Service_MakeTestingApp
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Service_MakeTestingApp(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring, p2 _Ctype_struct_nstring, p3 _Ctype_char, p4 _Ctype_char, p5 _Ctype_int32_t) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain_Service_MakeTestingApp(p0, p1, p2, p3, p4, p5)
}
//go:cgo_export_dynamic new_holochain_Service
//go:linkname _cgoexp_ebd397278953_new_holochain_Service _cgoexp_ebd397278953_new_holochain_Service
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Service
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Service(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Service
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Service() (r0 _Ctype_int32_t) {
	return new_holochain_Service()
}
//go:cgo_export_dynamic proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Set _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Set(p0 _Ctype_int32_t, p1 _Ctype_char) {
	proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Get _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Get(p0 _Ctype_int32_t) (r0 _Ctype_char) {
	return proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Set _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Set(p0 _Ctype_int32_t, p1 _Ctype_char) {
	proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Get _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Get(p0 _Ctype_int32_t) (r0 _Ctype_char) {
	return proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_ServiceConfig_DefaultBootstrapServer_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultBootstrapServer_Set _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultBootstrapServer_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultBootstrapServer_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultBootstrapServer_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ServiceConfig_DefaultBootstrapServer_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ServiceConfig_DefaultBootstrapServer_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_ServiceConfig_DefaultBootstrapServer_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ServiceConfig_DefaultBootstrapServer_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultBootstrapServer_Get _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultBootstrapServer_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultBootstrapServer_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultBootstrapServer_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ServiceConfig_DefaultBootstrapServer_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ServiceConfig_DefaultBootstrapServer_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ServiceConfig_DefaultBootstrapServer_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_ServiceConfig_DefaultEnableMDNS_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableMDNS_Set _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableMDNS_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableMDNS_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableMDNS_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableMDNS_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableMDNS_Set(p0 _Ctype_int32_t, p1 _Ctype_char) {
	proxyholochain_ServiceConfig_DefaultEnableMDNS_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ServiceConfig_DefaultEnableMDNS_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableMDNS_Get _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableMDNS_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableMDNS_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableMDNS_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableMDNS_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableMDNS_Get(p0 _Ctype_int32_t) (r0 _Ctype_char) {
	return proxyholochain_ServiceConfig_DefaultEnableMDNS_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Set _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Set(p0 _Ctype_int32_t, p1 _Ctype_char) {
	proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Get _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Get(p0 _Ctype_int32_t) (r0 _Ctype_char) {
	return proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_ServiceConfig_Validate
//go:linkname _cgoexp_ebd397278953_proxyholochain_ServiceConfig_Validate _cgoexp_ebd397278953_proxyholochain_ServiceConfig_Validate
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ServiceConfig_Validate
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ServiceConfig_Validate(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ServiceConfig_Validate
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ServiceConfig_Validate(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_ServiceConfig_Validate(p0)
}
//go:cgo_export_dynamic new_holochain_ServiceConfig
//go:linkname _cgoexp_ebd397278953_new_holochain_ServiceConfig _cgoexp_ebd397278953_new_holochain_ServiceConfig
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ServiceConfig
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ServiceConfig(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ServiceConfig
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ServiceConfig() (r0 _Ctype_int32_t) {
	return new_holochain_ServiceConfig()
}
//go:cgo_export_dynamic proxyholochain_Signature_S_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Signature_S_Set _cgoexp_ebd397278953_proxyholochain_Signature_S_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Signature_S_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Signature_S_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Signature_S_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Signature_S_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nbyteslice) {
	proxyholochain_Signature_S_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Signature_S_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Signature_S_Get _cgoexp_ebd397278953_proxyholochain_Signature_S_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Signature_S_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Signature_S_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Signature_S_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Signature_S_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nbyteslice) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Signature_S_Get(p0)
}
//go:cgo_export_dynamic new_holochain_Signature
//go:linkname _cgoexp_ebd397278953_new_holochain_Signature _cgoexp_ebd397278953_new_holochain_Signature
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Signature
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Signature(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Signature
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Signature() (r0 _Ctype_int32_t) {
	return new_holochain_Signature()
}
//go:cgo_export_dynamic proxyholochain_StatusChange_Action_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_StatusChange_Action_Set _cgoexp_ebd397278953_proxyholochain_StatusChange_Action_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_StatusChange_Action_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_StatusChange_Action_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_StatusChange_Action_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_StatusChange_Action_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_StatusChange_Action_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_StatusChange_Action_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_StatusChange_Action_Get _cgoexp_ebd397278953_proxyholochain_StatusChange_Action_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_StatusChange_Action_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_StatusChange_Action_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_StatusChange_Action_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_StatusChange_Action_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_StatusChange_Action_Get(p0)
}
//go:cgo_export_dynamic new_holochain_StatusChange
//go:linkname _cgoexp_ebd397278953_new_holochain_StatusChange _cgoexp_ebd397278953_new_holochain_StatusChange
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_StatusChange
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_StatusChange(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_StatusChange
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_StatusChange() (r0 _Ctype_int32_t) {
	return new_holochain_StatusChange()
}
//go:cgo_export_dynamic proxyholochain_TaggedHash_H_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_TaggedHash_H_Set _cgoexp_ebd397278953_proxyholochain_TaggedHash_H_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TaggedHash_H_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TaggedHash_H_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TaggedHash_H_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TaggedHash_H_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_TaggedHash_H_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_TaggedHash_H_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_TaggedHash_H_Get _cgoexp_ebd397278953_proxyholochain_TaggedHash_H_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TaggedHash_H_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TaggedHash_H_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TaggedHash_H_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TaggedHash_H_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_TaggedHash_H_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_TaggedHash_E_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_TaggedHash_E_Set _cgoexp_ebd397278953_proxyholochain_TaggedHash_E_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TaggedHash_E_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TaggedHash_E_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TaggedHash_E_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TaggedHash_E_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_TaggedHash_E_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_TaggedHash_E_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_TaggedHash_E_Get _cgoexp_ebd397278953_proxyholochain_TaggedHash_E_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TaggedHash_E_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TaggedHash_E_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TaggedHash_E_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TaggedHash_E_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_TaggedHash_E_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_TaggedHash_EntryType_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_TaggedHash_EntryType_Set _cgoexp_ebd397278953_proxyholochain_TaggedHash_EntryType_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TaggedHash_EntryType_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TaggedHash_EntryType_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TaggedHash_EntryType_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TaggedHash_EntryType_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_TaggedHash_EntryType_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_TaggedHash_EntryType_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_TaggedHash_EntryType_Get _cgoexp_ebd397278953_proxyholochain_TaggedHash_EntryType_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TaggedHash_EntryType_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TaggedHash_EntryType_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TaggedHash_EntryType_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TaggedHash_EntryType_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_TaggedHash_EntryType_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_TaggedHash_T_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_TaggedHash_T_Set _cgoexp_ebd397278953_proxyholochain_TaggedHash_T_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TaggedHash_T_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TaggedHash_T_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TaggedHash_T_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TaggedHash_T_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_TaggedHash_T_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_TaggedHash_T_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_TaggedHash_T_Get _cgoexp_ebd397278953_proxyholochain_TaggedHash_T_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TaggedHash_T_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TaggedHash_T_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TaggedHash_T_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TaggedHash_T_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_TaggedHash_T_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_TaggedHash_Source_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_TaggedHash_Source_Set _cgoexp_ebd397278953_proxyholochain_TaggedHash_Source_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TaggedHash_Source_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TaggedHash_Source_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TaggedHash_Source_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TaggedHash_Source_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_TaggedHash_Source_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_TaggedHash_Source_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_TaggedHash_Source_Get _cgoexp_ebd397278953_proxyholochain_TaggedHash_Source_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TaggedHash_Source_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TaggedHash_Source_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TaggedHash_Source_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TaggedHash_Source_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_TaggedHash_Source_Get(p0)
}
//go:cgo_export_dynamic new_holochain_TaggedHash
//go:linkname _cgoexp_ebd397278953_new_holochain_TaggedHash _cgoexp_ebd397278953_new_holochain_TaggedHash
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_TaggedHash
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_TaggedHash(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_TaggedHash
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_TaggedHash() (r0 _Ctype_int32_t) {
	return new_holochain_TaggedHash()
}
//go:cgo_export_dynamic proxyholochain_TestConfig_GossipInterval_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestConfig_GossipInterval_Set _cgoexp_ebd397278953_proxyholochain_TestConfig_GossipInterval_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestConfig_GossipInterval_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestConfig_GossipInterval_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestConfig_GossipInterval_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestConfig_GossipInterval_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_TestConfig_GossipInterval_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_TestConfig_GossipInterval_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestConfig_GossipInterval_Get _cgoexp_ebd397278953_proxyholochain_TestConfig_GossipInterval_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestConfig_GossipInterval_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestConfig_GossipInterval_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestConfig_GossipInterval_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestConfig_GossipInterval_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_TestConfig_GossipInterval_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_TestConfig_Duration_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestConfig_Duration_Set _cgoexp_ebd397278953_proxyholochain_TestConfig_Duration_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestConfig_Duration_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestConfig_Duration_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestConfig_Duration_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestConfig_Duration_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_TestConfig_Duration_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_TestConfig_Duration_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestConfig_Duration_Get _cgoexp_ebd397278953_proxyholochain_TestConfig_Duration_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestConfig_Duration_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestConfig_Duration_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestConfig_Duration_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestConfig_Duration_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_TestConfig_Duration_Get(p0)
}
//go:cgo_export_dynamic new_holochain_TestConfig
//go:linkname _cgoexp_ebd397278953_new_holochain_TestConfig _cgoexp_ebd397278953_new_holochain_TestConfig
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_TestConfig
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_TestConfig(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_TestConfig
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_TestConfig() (r0 _Ctype_int32_t) {
	return new_holochain_TestConfig()
}
//go:cgo_export_dynamic proxyholochain_TestData_Convey_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestData_Convey_Set _cgoexp_ebd397278953_proxyholochain_TestData_Convey_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestData_Convey_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestData_Convey_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestData_Convey_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestData_Convey_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_TestData_Convey_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_TestData_Convey_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestData_Convey_Get _cgoexp_ebd397278953_proxyholochain_TestData_Convey_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestData_Convey_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestData_Convey_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestData_Convey_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestData_Convey_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_TestData_Convey_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_TestData_Zome_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestData_Zome_Set _cgoexp_ebd397278953_proxyholochain_TestData_Zome_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestData_Zome_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestData_Zome_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestData_Zome_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestData_Zome_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_TestData_Zome_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_TestData_Zome_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestData_Zome_Get _cgoexp_ebd397278953_proxyholochain_TestData_Zome_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestData_Zome_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestData_Zome_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestData_Zome_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestData_Zome_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_TestData_Zome_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_TestData_FnName_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestData_FnName_Set _cgoexp_ebd397278953_proxyholochain_TestData_FnName_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestData_FnName_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestData_FnName_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestData_FnName_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestData_FnName_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_TestData_FnName_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_TestData_FnName_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestData_FnName_Get _cgoexp_ebd397278953_proxyholochain_TestData_FnName_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestData_FnName_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestData_FnName_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestData_FnName_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestData_FnName_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_TestData_FnName_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_TestData_Regexp_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestData_Regexp_Set _cgoexp_ebd397278953_proxyholochain_TestData_Regexp_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestData_Regexp_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestData_Regexp_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestData_Regexp_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestData_Regexp_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_TestData_Regexp_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_TestData_Regexp_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestData_Regexp_Get _cgoexp_ebd397278953_proxyholochain_TestData_Regexp_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestData_Regexp_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestData_Regexp_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestData_Regexp_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestData_Regexp_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_TestData_Regexp_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_TestData_Exposure_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestData_Exposure_Set _cgoexp_ebd397278953_proxyholochain_TestData_Exposure_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestData_Exposure_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestData_Exposure_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestData_Exposure_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestData_Exposure_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_TestData_Exposure_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_TestData_Exposure_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestData_Exposure_Get _cgoexp_ebd397278953_proxyholochain_TestData_Exposure_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestData_Exposure_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestData_Exposure_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestData_Exposure_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestData_Exposure_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_TestData_Exposure_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_TestData_Raw_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestData_Raw_Set _cgoexp_ebd397278953_proxyholochain_TestData_Raw_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestData_Raw_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestData_Raw_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestData_Raw_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestData_Raw_Set(p0 _Ctype_int32_t, p1 _Ctype_char) {
	proxyholochain_TestData_Raw_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_TestData_Raw_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestData_Raw_Get _cgoexp_ebd397278953_proxyholochain_TestData_Raw_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestData_Raw_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestData_Raw_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestData_Raw_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestData_Raw_Get(p0 _Ctype_int32_t) (r0 _Ctype_char) {
	return proxyholochain_TestData_Raw_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_TestData_Repeat_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestData_Repeat_Set _cgoexp_ebd397278953_proxyholochain_TestData_Repeat_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestData_Repeat_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestData_Repeat_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestData_Repeat_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestData_Repeat_Set(p0 _Ctype_int32_t, p1 _Ctype_nint) {
	proxyholochain_TestData_Repeat_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_TestData_Repeat_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestData_Repeat_Get _cgoexp_ebd397278953_proxyholochain_TestData_Repeat_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestData_Repeat_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestData_Repeat_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestData_Repeat_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestData_Repeat_Get(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_TestData_Repeat_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_TestData_Benchmark_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestData_Benchmark_Set _cgoexp_ebd397278953_proxyholochain_TestData_Benchmark_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestData_Benchmark_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestData_Benchmark_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestData_Benchmark_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestData_Benchmark_Set(p0 _Ctype_int32_t, p1 _Ctype_char) {
	proxyholochain_TestData_Benchmark_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_TestData_Benchmark_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestData_Benchmark_Get _cgoexp_ebd397278953_proxyholochain_TestData_Benchmark_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestData_Benchmark_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestData_Benchmark_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestData_Benchmark_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestData_Benchmark_Get(p0 _Ctype_int32_t) (r0 _Ctype_char) {
	return proxyholochain_TestData_Benchmark_Get(p0)
}
//go:cgo_export_dynamic new_holochain_TestData
//go:linkname _cgoexp_ebd397278953_new_holochain_TestData _cgoexp_ebd397278953_new_holochain_TestData
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_TestData
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_TestData(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_TestData
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_TestData() (r0 _Ctype_int32_t) {
	return new_holochain_TestData()
}
//go:cgo_export_dynamic new_holochain_TestFixtures
//go:linkname _cgoexp_ebd397278953_new_holochain_TestFixtures _cgoexp_ebd397278953_new_holochain_TestFixtures
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_TestFixtures
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_TestFixtures(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_TestFixtures
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_TestFixtures() (r0 _Ctype_int32_t) {
	return new_holochain_TestFixtures()
}
//go:cgo_export_dynamic proxyholochain_TestSet_Identity_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestSet_Identity_Set _cgoexp_ebd397278953_proxyholochain_TestSet_Identity_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestSet_Identity_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestSet_Identity_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestSet_Identity_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestSet_Identity_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_TestSet_Identity_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_TestSet_Identity_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestSet_Identity_Get _cgoexp_ebd397278953_proxyholochain_TestSet_Identity_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestSet_Identity_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestSet_Identity_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestSet_Identity_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestSet_Identity_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_TestSet_Identity_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_TestSet_Benchmark_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestSet_Benchmark_Set _cgoexp_ebd397278953_proxyholochain_TestSet_Benchmark_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestSet_Benchmark_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestSet_Benchmark_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestSet_Benchmark_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestSet_Benchmark_Set(p0 _Ctype_int32_t, p1 _Ctype_char) {
	proxyholochain_TestSet_Benchmark_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_TestSet_Benchmark_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_TestSet_Benchmark_Get _cgoexp_ebd397278953_proxyholochain_TestSet_Benchmark_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_TestSet_Benchmark_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_TestSet_Benchmark_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_TestSet_Benchmark_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_TestSet_Benchmark_Get(p0 _Ctype_int32_t) (r0 _Ctype_char) {
	return proxyholochain_TestSet_Benchmark_Get(p0)
}
//go:cgo_export_dynamic new_holochain_TestSet
//go:linkname _cgoexp_ebd397278953_new_holochain_TestSet _cgoexp_ebd397278953_new_holochain_TestSet
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_TestSet
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_TestSet(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_TestSet
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_TestSet() (r0 _Ctype_int32_t) {
	return new_holochain_TestSet()
}
//go:cgo_export_dynamic new_holochain_ValidateQuery
//go:linkname _cgoexp_ebd397278953_new_holochain_ValidateQuery _cgoexp_ebd397278953_new_holochain_ValidateQuery
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ValidateQuery
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ValidateQuery(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ValidateQuery
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ValidateQuery() (r0 _Ctype_int32_t) {
	return new_holochain_ValidateQuery()
}
//go:cgo_export_dynamic proxyholochain_ValidateResponse_Type_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_ValidateResponse_Type_Set _cgoexp_ebd397278953_proxyholochain_ValidateResponse_Type_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ValidateResponse_Type_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ValidateResponse_Type_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ValidateResponse_Type_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ValidateResponse_Type_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_ValidateResponse_Type_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ValidateResponse_Type_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_ValidateResponse_Type_Get _cgoexp_ebd397278953_proxyholochain_ValidateResponse_Type_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ValidateResponse_Type_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ValidateResponse_Type_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ValidateResponse_Type_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ValidateResponse_Type_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ValidateResponse_Type_Get(p0)
}
//go:cgo_export_dynamic new_holochain_ValidateResponse
//go:linkname _cgoexp_ebd397278953_new_holochain_ValidateResponse _cgoexp_ebd397278953_new_holochain_ValidateResponse
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ValidateResponse
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ValidateResponse(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ValidateResponse
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ValidateResponse() (r0 _Ctype_int32_t) {
	return new_holochain_ValidateResponse()
}
//go:cgo_export_dynamic proxyholochain_ValidationPackage_Chain_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_ValidationPackage_Chain_Set _cgoexp_ebd397278953_proxyholochain_ValidationPackage_Chain_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ValidationPackage_Chain_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ValidationPackage_Chain_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ValidationPackage_Chain_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ValidationPackage_Chain_Set(p0 _Ctype_int32_t, p1 _Ctype_int32_t) {
	proxyholochain_ValidationPackage_Chain_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ValidationPackage_Chain_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_ValidationPackage_Chain_Get _cgoexp_ebd397278953_proxyholochain_ValidationPackage_Chain_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ValidationPackage_Chain_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ValidationPackage_Chain_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ValidationPackage_Chain_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ValidationPackage_Chain_Get(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_ValidationPackage_Chain_Get(p0)
}
//go:cgo_export_dynamic new_holochain_ValidationPackage
//go:linkname _cgoexp_ebd397278953_new_holochain_ValidationPackage _cgoexp_ebd397278953_new_holochain_ValidationPackage
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ValidationPackage
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ValidationPackage(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ValidationPackage
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ValidationPackage() (r0 _Ctype_int32_t) {
	return new_holochain_ValidationPackage()
}
//go:cgo_export_dynamic proxyholochain_Zome_Name_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Zome_Name_Set _cgoexp_ebd397278953_proxyholochain_Zome_Name_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Zome_Name_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Zome_Name_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Zome_Name_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Zome_Name_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Zome_Name_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Zome_Name_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Zome_Name_Get _cgoexp_ebd397278953_proxyholochain_Zome_Name_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Zome_Name_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Zome_Name_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Zome_Name_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Zome_Name_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Zome_Name_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Zome_Description_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Zome_Description_Set _cgoexp_ebd397278953_proxyholochain_Zome_Description_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Zome_Description_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Zome_Description_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Zome_Description_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Zome_Description_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Zome_Description_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Zome_Description_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Zome_Description_Get _cgoexp_ebd397278953_proxyholochain_Zome_Description_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Zome_Description_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Zome_Description_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Zome_Description_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Zome_Description_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Zome_Description_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Zome_Code_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Zome_Code_Set _cgoexp_ebd397278953_proxyholochain_Zome_Code_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Zome_Code_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Zome_Code_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Zome_Code_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Zome_Code_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Zome_Code_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Zome_Code_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Zome_Code_Get _cgoexp_ebd397278953_proxyholochain_Zome_Code_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Zome_Code_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Zome_Code_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Zome_Code_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Zome_Code_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Zome_Code_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Zome_RibosomeType_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_Zome_RibosomeType_Set _cgoexp_ebd397278953_proxyholochain_Zome_RibosomeType_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Zome_RibosomeType_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Zome_RibosomeType_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Zome_RibosomeType_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Zome_RibosomeType_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_Zome_RibosomeType_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Zome_RibosomeType_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_Zome_RibosomeType_Get _cgoexp_ebd397278953_proxyholochain_Zome_RibosomeType_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Zome_RibosomeType_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Zome_RibosomeType_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Zome_RibosomeType_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Zome_RibosomeType_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Zome_RibosomeType_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_Zome_CodeFileName
//go:linkname _cgoexp_ebd397278953_proxyholochain_Zome_CodeFileName _cgoexp_ebd397278953_proxyholochain_Zome_CodeFileName
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Zome_CodeFileName
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Zome_CodeFileName(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Zome_CodeFileName
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Zome_CodeFileName(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Zome_CodeFileName(p0)
}
//go:cgo_export_dynamic proxyholochain_Zome_GetEntryDef
//go:linkname _cgoexp_ebd397278953_proxyholochain_Zome_GetEntryDef _cgoexp_ebd397278953_proxyholochain_Zome_GetEntryDef
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Zome_GetEntryDef
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Zome_GetEntryDef(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Zome_GetEntryDef
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Zome_GetEntryDef(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain_Zome_GetEntryDef(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Zome_GetFunctionDef
//go:linkname _cgoexp_ebd397278953_proxyholochain_Zome_GetFunctionDef _cgoexp_ebd397278953_proxyholochain_Zome_GetFunctionDef
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Zome_GetFunctionDef
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Zome_GetFunctionDef(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Zome_GetFunctionDef
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Zome_GetFunctionDef(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain_Zome_GetFunctionDef(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Zome_MakeRibosome
//go:linkname _cgoexp_ebd397278953_proxyholochain_Zome_MakeRibosome _cgoexp_ebd397278953_proxyholochain_Zome_MakeRibosome
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Zome_MakeRibosome
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Zome_MakeRibosome(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Zome_MakeRibosome
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Zome_MakeRibosome(p0 _Ctype_int32_t, p1 _Ctype_int32_t) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain_Zome_MakeRibosome(p0, p1)
}
//go:cgo_export_dynamic new_holochain_Zome
//go:linkname _cgoexp_ebd397278953_new_holochain_Zome _cgoexp_ebd397278953_new_holochain_Zome
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_Zome
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_Zome(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_Zome
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_Zome() (r0 _Ctype_int32_t) {
	return new_holochain_Zome()
}
//go:cgo_export_dynamic proxyholochain_ZomeFile_Name_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_ZomeFile_Name_Set _cgoexp_ebd397278953_proxyholochain_ZomeFile_Name_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ZomeFile_Name_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ZomeFile_Name_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ZomeFile_Name_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ZomeFile_Name_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_ZomeFile_Name_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ZomeFile_Name_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_ZomeFile_Name_Get _cgoexp_ebd397278953_proxyholochain_ZomeFile_Name_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ZomeFile_Name_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ZomeFile_Name_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ZomeFile_Name_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ZomeFile_Name_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ZomeFile_Name_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_ZomeFile_Description_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_ZomeFile_Description_Set _cgoexp_ebd397278953_proxyholochain_ZomeFile_Description_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ZomeFile_Description_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ZomeFile_Description_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ZomeFile_Description_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ZomeFile_Description_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_ZomeFile_Description_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ZomeFile_Description_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_ZomeFile_Description_Get _cgoexp_ebd397278953_proxyholochain_ZomeFile_Description_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ZomeFile_Description_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ZomeFile_Description_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ZomeFile_Description_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ZomeFile_Description_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ZomeFile_Description_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_ZomeFile_CodeFile_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_ZomeFile_CodeFile_Set _cgoexp_ebd397278953_proxyholochain_ZomeFile_CodeFile_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ZomeFile_CodeFile_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ZomeFile_CodeFile_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ZomeFile_CodeFile_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ZomeFile_CodeFile_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_ZomeFile_CodeFile_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ZomeFile_CodeFile_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_ZomeFile_CodeFile_Get _cgoexp_ebd397278953_proxyholochain_ZomeFile_CodeFile_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ZomeFile_CodeFile_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ZomeFile_CodeFile_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ZomeFile_CodeFile_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ZomeFile_CodeFile_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ZomeFile_CodeFile_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_ZomeFile_RibosomeType_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_ZomeFile_RibosomeType_Set _cgoexp_ebd397278953_proxyholochain_ZomeFile_RibosomeType_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ZomeFile_RibosomeType_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ZomeFile_RibosomeType_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ZomeFile_RibosomeType_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ZomeFile_RibosomeType_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_ZomeFile_RibosomeType_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ZomeFile_RibosomeType_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_ZomeFile_RibosomeType_Get _cgoexp_ebd397278953_proxyholochain_ZomeFile_RibosomeType_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ZomeFile_RibosomeType_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ZomeFile_RibosomeType_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ZomeFile_RibosomeType_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ZomeFile_RibosomeType_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ZomeFile_RibosomeType_Get(p0)
}
//go:cgo_export_dynamic proxyholochain_ZomeFile_BridgeTo_Set
//go:linkname _cgoexp_ebd397278953_proxyholochain_ZomeFile_BridgeTo_Set _cgoexp_ebd397278953_proxyholochain_ZomeFile_BridgeTo_Set
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ZomeFile_BridgeTo_Set
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ZomeFile_BridgeTo_Set(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ZomeFile_BridgeTo_Set
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ZomeFile_BridgeTo_Set(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring) {
	proxyholochain_ZomeFile_BridgeTo_Set(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ZomeFile_BridgeTo_Get
//go:linkname _cgoexp_ebd397278953_proxyholochain_ZomeFile_BridgeTo_Get _cgoexp_ebd397278953_proxyholochain_ZomeFile_BridgeTo_Get
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ZomeFile_BridgeTo_Get
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ZomeFile_BridgeTo_Get(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ZomeFile_BridgeTo_Get
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ZomeFile_BridgeTo_Get(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ZomeFile_BridgeTo_Get(p0)
}
//go:cgo_export_dynamic new_holochain_ZomeFile
//go:linkname _cgoexp_ebd397278953_new_holochain_ZomeFile _cgoexp_ebd397278953_new_holochain_ZomeFile
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ZomeFile
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ZomeFile(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ZomeFile
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ZomeFile() (r0 _Ctype_int32_t) {
	return new_holochain_ZomeFile()
}
//go:cgo_export_dynamic proxyholochain_ZygoRibosome_ChainGenesis
//go:linkname _cgoexp_ebd397278953_proxyholochain_ZygoRibosome_ChainGenesis _cgoexp_ebd397278953_proxyholochain_ZygoRibosome_ChainGenesis
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ZygoRibosome_ChainGenesis
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ZygoRibosome_ChainGenesis(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ZygoRibosome_ChainGenesis
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ZygoRibosome_ChainGenesis(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_ZygoRibosome_ChainGenesis(p0)
}
//go:cgo_export_dynamic proxyholochain_ZygoRibosome_Receive
//go:linkname _cgoexp_ebd397278953_proxyholochain_ZygoRibosome_Receive _cgoexp_ebd397278953_proxyholochain_ZygoRibosome_Receive
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ZygoRibosome_Receive
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ZygoRibosome_Receive(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ZygoRibosome_Receive
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ZygoRibosome_Receive(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring, p2 _Ctype_struct_nstring) (r0 _Ctype_struct_nstring, r1 _Ctype_int32_t) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ZygoRibosome_Receive(p0, p1, p2)
}
//go:cgo_export_dynamic proxyholochain_ZygoRibosome_Type
//go:linkname _cgoexp_ebd397278953_proxyholochain_ZygoRibosome_Type _cgoexp_ebd397278953_proxyholochain_ZygoRibosome_Type
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ZygoRibosome_Type
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ZygoRibosome_Type(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ZygoRibosome_Type
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ZygoRibosome_Type(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ZygoRibosome_Type(p0)
}
//go:cgo_export_dynamic new_holochain_ZygoRibosome
//go:linkname _cgoexp_ebd397278953_new_holochain_ZygoRibosome _cgoexp_ebd397278953_new_holochain_ZygoRibosome
//go:cgo_export_static _cgoexp_ebd397278953_new_holochain_ZygoRibosome
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_new_holochain_ZygoRibosome(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_new_holochain_ZygoRibosome
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_new_holochain_ZygoRibosome() (r0 _Ctype_int32_t) {
	return new_holochain_ZygoRibosome()
}
//go:cgo_export_dynamic proxyholochain_Action_Name
//go:linkname _cgoexp_ebd397278953_proxyholochain_Action_Name _cgoexp_ebd397278953_proxyholochain_Action_Name
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Action_Name
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Action_Name(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Action_Name
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Action_Name(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Action_Name(p0)
}
//go:cgo_export_dynamic proxyholochain_CommittingAction_CheckValidationRequest
//go:linkname _cgoexp_ebd397278953_proxyholochain_CommittingAction_CheckValidationRequest _cgoexp_ebd397278953_proxyholochain_CommittingAction_CheckValidationRequest
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_CommittingAction_CheckValidationRequest
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_CommittingAction_CheckValidationRequest(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_CommittingAction_CheckValidationRequest
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_CommittingAction_CheckValidationRequest(p0 _Ctype_int32_t, p1 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_CommittingAction_CheckValidationRequest(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_CommittingAction_Entry
//go:linkname _cgoexp_ebd397278953_proxyholochain_CommittingAction_Entry _cgoexp_ebd397278953_proxyholochain_CommittingAction_Entry
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_CommittingAction_Entry
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_CommittingAction_Entry(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_CommittingAction_Entry
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_CommittingAction_Entry(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_CommittingAction_Entry(p0)
}
//go:cgo_export_dynamic proxyholochain_CommittingAction_EntryType
//go:linkname _cgoexp_ebd397278953_proxyholochain_CommittingAction_EntryType _cgoexp_ebd397278953_proxyholochain_CommittingAction_EntryType
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_CommittingAction_EntryType
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_CommittingAction_EntryType(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_CommittingAction_EntryType
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_CommittingAction_EntryType(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_CommittingAction_EntryType(p0)
}
//go:cgo_export_dynamic proxyholochain_CommittingAction_Name
//go:linkname _cgoexp_ebd397278953_proxyholochain_CommittingAction_Name _cgoexp_ebd397278953_proxyholochain_CommittingAction_Name
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_CommittingAction_Name
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_CommittingAction_Name(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_CommittingAction_Name
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_CommittingAction_Name(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_CommittingAction_Name(p0)
}
//go:cgo_export_dynamic proxyholochain_CommittingAction_SetHeader
//go:linkname _cgoexp_ebd397278953_proxyholochain_CommittingAction_SetHeader _cgoexp_ebd397278953_proxyholochain_CommittingAction_SetHeader
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_CommittingAction_SetHeader
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_CommittingAction_SetHeader(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_CommittingAction_SetHeader
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_CommittingAction_SetHeader(p0 _Ctype_int32_t, p1 _Ctype_int32_t) {
	proxyholochain_CommittingAction_SetHeader(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Entry_Marshal
//go:linkname _cgoexp_ebd397278953_proxyholochain_Entry_Marshal _cgoexp_ebd397278953_proxyholochain_Entry_Marshal
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Entry_Marshal
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Entry_Marshal(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Entry_Marshal
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Entry_Marshal(p0 _Ctype_int32_t) (r0 _Ctype_struct_nbyteslice, r1 _Ctype_int32_t) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Entry_Marshal(p0)
}
//go:cgo_export_dynamic proxyholochain_Entry_Unmarshal
//go:linkname _cgoexp_ebd397278953_proxyholochain_Entry_Unmarshal _cgoexp_ebd397278953_proxyholochain_Entry_Unmarshal
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Entry_Unmarshal
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Entry_Unmarshal(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Entry_Unmarshal
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Entry_Unmarshal(p0 _Ctype_int32_t, p1 _Ctype_struct_nbyteslice) (r0 _Ctype_int32_t) {
	return proxyholochain_Entry_Unmarshal(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Revocation_Marshal
//go:linkname _cgoexp_ebd397278953_proxyholochain_Revocation_Marshal _cgoexp_ebd397278953_proxyholochain_Revocation_Marshal
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Revocation_Marshal
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Revocation_Marshal(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Revocation_Marshal
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Revocation_Marshal(p0 _Ctype_int32_t) (r0 _Ctype_struct_nbyteslice, r1 _Ctype_int32_t) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Revocation_Marshal(p0)
}
//go:cgo_export_dynamic proxyholochain_Revocation_Unmarshal
//go:linkname _cgoexp_ebd397278953_proxyholochain_Revocation_Unmarshal _cgoexp_ebd397278953_proxyholochain_Revocation_Unmarshal
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Revocation_Unmarshal
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Revocation_Unmarshal(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Revocation_Unmarshal
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Revocation_Unmarshal(p0 _Ctype_int32_t, p1 _Ctype_struct_nbyteslice) (r0 _Ctype_int32_t) {
	return proxyholochain_Revocation_Unmarshal(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Revocation_Verify
//go:linkname _cgoexp_ebd397278953_proxyholochain_Revocation_Verify _cgoexp_ebd397278953_proxyholochain_Revocation_Verify
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Revocation_Verify
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Revocation_Verify(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Revocation_Verify
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Revocation_Verify(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_Revocation_Verify(p0)
}
//go:cgo_export_dynamic proxyholochain_Ribosome_ChainGenesis
//go:linkname _cgoexp_ebd397278953_proxyholochain_Ribosome_ChainGenesis _cgoexp_ebd397278953_proxyholochain_Ribosome_ChainGenesis
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Ribosome_ChainGenesis
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Ribosome_ChainGenesis(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Ribosome_ChainGenesis
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Ribosome_ChainGenesis(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_Ribosome_ChainGenesis(p0)
}
//go:cgo_export_dynamic proxyholochain_Ribosome_Receive
//go:linkname _cgoexp_ebd397278953_proxyholochain_Ribosome_Receive _cgoexp_ebd397278953_proxyholochain_Ribosome_Receive
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Ribosome_Receive
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Ribosome_Receive(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Ribosome_Receive
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Ribosome_Receive(p0 _Ctype_int32_t, p1 _Ctype_struct_nstring, p2 _Ctype_struct_nstring) (r0 _Ctype_struct_nstring, r1 _Ctype_int32_t) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Ribosome_Receive(p0, p1, p2)
}
//go:cgo_export_dynamic proxyholochain_Ribosome_Type
//go:linkname _cgoexp_ebd397278953_proxyholochain_Ribosome_Type _cgoexp_ebd397278953_proxyholochain_Ribosome_Type
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Ribosome_Type
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Ribosome_Type(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Ribosome_Type
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Ribosome_Type(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Ribosome_Type(p0)
}
//go:cgo_export_dynamic proxyholochain_ValidatingAction_CheckValidationRequest
//go:linkname _cgoexp_ebd397278953_proxyholochain_ValidatingAction_CheckValidationRequest _cgoexp_ebd397278953_proxyholochain_ValidatingAction_CheckValidationRequest
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ValidatingAction_CheckValidationRequest
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ValidatingAction_CheckValidationRequest(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ValidatingAction_CheckValidationRequest
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ValidatingAction_CheckValidationRequest(p0 _Ctype_int32_t, p1 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_ValidatingAction_CheckValidationRequest(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_ValidatingAction_Name
//go:linkname _cgoexp_ebd397278953_proxyholochain_ValidatingAction_Name _cgoexp_ebd397278953_proxyholochain_ValidatingAction_Name
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_ValidatingAction_Name
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_ValidatingAction_Name(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_ValidatingAction_Name
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_ValidatingAction_Name(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_ValidatingAction_Name(p0)
}
//go:cgo_export_dynamic proxyholochain_Warrant_Decode
//go:linkname _cgoexp_ebd397278953_proxyholochain_Warrant_Decode _cgoexp_ebd397278953_proxyholochain_Warrant_Decode
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Warrant_Decode
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Warrant_Decode(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Warrant_Decode
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Warrant_Decode(p0 _Ctype_int32_t, p1 _Ctype_struct_nbyteslice) (r0 _Ctype_int32_t) {
	return proxyholochain_Warrant_Decode(p0, p1)
}
//go:cgo_export_dynamic proxyholochain_Warrant_Encode
//go:linkname _cgoexp_ebd397278953_proxyholochain_Warrant_Encode _cgoexp_ebd397278953_proxyholochain_Warrant_Encode
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Warrant_Encode
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Warrant_Encode(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Warrant_Encode
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Warrant_Encode(p0 _Ctype_int32_t) (r0 _Ctype_struct_nbyteslice, r1 _Ctype_int32_t) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain_Warrant_Encode(p0)
}
//go:cgo_export_dynamic proxyholochain_Warrant_Type
//go:linkname _cgoexp_ebd397278953_proxyholochain_Warrant_Type _cgoexp_ebd397278953_proxyholochain_Warrant_Type
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Warrant_Type
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Warrant_Type(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Warrant_Type
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Warrant_Type(p0 _Ctype_int32_t) (r0 _Ctype_nint) {
	return proxyholochain_Warrant_Type(p0)
}
//go:cgo_export_dynamic proxyholochain_Warrant_Verify
//go:linkname _cgoexp_ebd397278953_proxyholochain_Warrant_Verify _cgoexp_ebd397278953_proxyholochain_Warrant_Verify
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain_Warrant_Verify
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain_Warrant_Verify(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain_Warrant_Verify
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain_Warrant_Verify(p0 _Ctype_int32_t, p1 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain_Warrant_Verify(p0, p1)
}
//go:cgo_export_dynamic var_setholochain_AgentEntryDef
//go:linkname _cgoexp_ebd397278953_var_setholochain_AgentEntryDef _cgoexp_ebd397278953_var_setholochain_AgentEntryDef
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_AgentEntryDef
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_AgentEntryDef(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_AgentEntryDef
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_AgentEntryDef(p0 _Ctype_int32_t) {
	var_setholochain_AgentEntryDef(p0)
}
//go:cgo_export_dynamic var_getholochain_AgentEntryDef
//go:linkname _cgoexp_ebd397278953_var_getholochain_AgentEntryDef _cgoexp_ebd397278953_var_getholochain_AgentEntryDef
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_AgentEntryDef
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_AgentEntryDef(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_AgentEntryDef
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_AgentEntryDef() (r0 _Ctype_int32_t) {
	return var_getholochain_AgentEntryDef()
}
//go:cgo_export_dynamic var_setholochain_AlphaValue
//go:linkname _cgoexp_ebd397278953_var_setholochain_AlphaValue _cgoexp_ebd397278953_var_setholochain_AlphaValue
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_AlphaValue
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_AlphaValue(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_AlphaValue
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_AlphaValue(p0 _Ctype_nint) {
	var_setholochain_AlphaValue(p0)
}
//go:cgo_export_dynamic var_getholochain_AlphaValue
//go:linkname _cgoexp_ebd397278953_var_getholochain_AlphaValue _cgoexp_ebd397278953_var_getholochain_AlphaValue
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_AlphaValue
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_AlphaValue(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_AlphaValue
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_AlphaValue() (r0 _Ctype_nint) {
	return var_getholochain_AlphaValue()
}
//go:cgo_export_dynamic var_setholochain_BasicTemplateAppPackage
//go:linkname _cgoexp_ebd397278953_var_setholochain_BasicTemplateAppPackage _cgoexp_ebd397278953_var_setholochain_BasicTemplateAppPackage
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_BasicTemplateAppPackage
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_BasicTemplateAppPackage(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_BasicTemplateAppPackage
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_BasicTemplateAppPackage(p0 _Ctype_struct_nstring) {
	var_setholochain_BasicTemplateAppPackage(p0)
}
//go:cgo_export_dynamic var_getholochain_BasicTemplateAppPackage
//go:linkname _cgoexp_ebd397278953_var_getholochain_BasicTemplateAppPackage _cgoexp_ebd397278953_var_getholochain_BasicTemplateAppPackage
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_BasicTemplateAppPackage
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_BasicTemplateAppPackage(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_BasicTemplateAppPackage
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_BasicTemplateAppPackage() (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return var_getholochain_BasicTemplateAppPackage()
}
//go:cgo_export_dynamic var_setholochain_BridgeAppNotFoundErr
//go:linkname _cgoexp_ebd397278953_var_setholochain_BridgeAppNotFoundErr _cgoexp_ebd397278953_var_setholochain_BridgeAppNotFoundErr
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_BridgeAppNotFoundErr
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_BridgeAppNotFoundErr(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_BridgeAppNotFoundErr
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_BridgeAppNotFoundErr(p0 _Ctype_int32_t) {
	var_setholochain_BridgeAppNotFoundErr(p0)
}
//go:cgo_export_dynamic var_getholochain_BridgeAppNotFoundErr
//go:linkname _cgoexp_ebd397278953_var_getholochain_BridgeAppNotFoundErr _cgoexp_ebd397278953_var_getholochain_BridgeAppNotFoundErr
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_BridgeAppNotFoundErr
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_BridgeAppNotFoundErr(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_BridgeAppNotFoundErr
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_BridgeAppNotFoundErr() (r0 _Ctype_int32_t) {
	return var_getholochain_BridgeAppNotFoundErr()
}
//go:cgo_export_dynamic var_setholochain_CapabilityInvalidErr
//go:linkname _cgoexp_ebd397278953_var_setholochain_CapabilityInvalidErr _cgoexp_ebd397278953_var_setholochain_CapabilityInvalidErr
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_CapabilityInvalidErr
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_CapabilityInvalidErr(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_CapabilityInvalidErr
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_CapabilityInvalidErr(p0 _Ctype_int32_t) {
	var_setholochain_CapabilityInvalidErr(p0)
}
//go:cgo_export_dynamic var_getholochain_CapabilityInvalidErr
//go:linkname _cgoexp_ebd397278953_var_getholochain_CapabilityInvalidErr _cgoexp_ebd397278953_var_getholochain_CapabilityInvalidErr
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_CapabilityInvalidErr
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_CapabilityInvalidErr(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_CapabilityInvalidErr
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_CapabilityInvalidErr() (r0 _Ctype_int32_t) {
	return var_getholochain_CapabilityInvalidErr()
}
//go:cgo_export_dynamic var_setholochain_CloserPeerCount
//go:linkname _cgoexp_ebd397278953_var_setholochain_CloserPeerCount _cgoexp_ebd397278953_var_setholochain_CloserPeerCount
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_CloserPeerCount
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_CloserPeerCount(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_CloserPeerCount
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_CloserPeerCount(p0 _Ctype_nint) {
	var_setholochain_CloserPeerCount(p0)
}
//go:cgo_export_dynamic var_getholochain_CloserPeerCount
//go:linkname _cgoexp_ebd397278953_var_getholochain_CloserPeerCount _cgoexp_ebd397278953_var_getholochain_CloserPeerCount
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_CloserPeerCount
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_CloserPeerCount(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_CloserPeerCount
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_CloserPeerCount() (r0 _Ctype_nint) {
	return var_getholochain_CloserPeerCount()
}
//go:cgo_export_dynamic var_setholochain_DNAEntryDef
//go:linkname _cgoexp_ebd397278953_var_setholochain_DNAEntryDef _cgoexp_ebd397278953_var_setholochain_DNAEntryDef
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_DNAEntryDef
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_DNAEntryDef(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_DNAEntryDef
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_DNAEntryDef(p0 _Ctype_int32_t) {
	var_setholochain_DNAEntryDef(p0)
}
//go:cgo_export_dynamic var_getholochain_DNAEntryDef
//go:linkname _cgoexp_ebd397278953_var_getholochain_DNAEntryDef _cgoexp_ebd397278953_var_getholochain_DNAEntryDef
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_DNAEntryDef
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_DNAEntryDef(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_DNAEntryDef
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_DNAEntryDef() (r0 _Ctype_int32_t) {
	return var_getholochain_DNAEntryDef()
}
//go:cgo_export_dynamic var_setholochain_EnableAllLoggersEnv
//go:linkname _cgoexp_ebd397278953_var_setholochain_EnableAllLoggersEnv _cgoexp_ebd397278953_var_setholochain_EnableAllLoggersEnv
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_EnableAllLoggersEnv
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_EnableAllLoggersEnv(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_EnableAllLoggersEnv
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_EnableAllLoggersEnv(p0 _Ctype_struct_nstring) {
	var_setholochain_EnableAllLoggersEnv(p0)
}
//go:cgo_export_dynamic var_getholochain_EnableAllLoggersEnv
//go:linkname _cgoexp_ebd397278953_var_getholochain_EnableAllLoggersEnv _cgoexp_ebd397278953_var_getholochain_EnableAllLoggersEnv
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_EnableAllLoggersEnv
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_EnableAllLoggersEnv(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_EnableAllLoggersEnv
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_EnableAllLoggersEnv() (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return var_getholochain_EnableAllLoggersEnv()
}
//go:cgo_export_dynamic var_setholochain_ErrBlockedListed
//go:linkname _cgoexp_ebd397278953_var_setholochain_ErrBlockedListed _cgoexp_ebd397278953_var_setholochain_ErrBlockedListed
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_ErrBlockedListed
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_ErrBlockedListed(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_ErrBlockedListed
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_ErrBlockedListed(p0 _Ctype_int32_t) {
	var_setholochain_ErrBlockedListed(p0)
}
//go:cgo_export_dynamic var_getholochain_ErrBlockedListed
//go:linkname _cgoexp_ebd397278953_var_getholochain_ErrBlockedListed _cgoexp_ebd397278953_var_getholochain_ErrBlockedListed
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_ErrBlockedListed
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_ErrBlockedListed(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_ErrBlockedListed
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_ErrBlockedListed() (r0 _Ctype_int32_t) {
	return var_getholochain_ErrBlockedListed()
}
//go:cgo_export_dynamic var_setholochain_ErrDHTErrNoGossipersAvailable
//go:linkname _cgoexp_ebd397278953_var_setholochain_ErrDHTErrNoGossipersAvailable _cgoexp_ebd397278953_var_setholochain_ErrDHTErrNoGossipersAvailable
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_ErrDHTErrNoGossipersAvailable
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_ErrDHTErrNoGossipersAvailable(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_ErrDHTErrNoGossipersAvailable
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_ErrDHTErrNoGossipersAvailable(p0 _Ctype_int32_t) {
	var_setholochain_ErrDHTErrNoGossipersAvailable(p0)
}
//go:cgo_export_dynamic var_getholochain_ErrDHTErrNoGossipersAvailable
//go:linkname _cgoexp_ebd397278953_var_getholochain_ErrDHTErrNoGossipersAvailable _cgoexp_ebd397278953_var_getholochain_ErrDHTErrNoGossipersAvailable
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_ErrDHTErrNoGossipersAvailable
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_ErrDHTErrNoGossipersAvailable(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_ErrDHTErrNoGossipersAvailable
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_ErrDHTErrNoGossipersAvailable() (r0 _Ctype_int32_t) {
	return var_getholochain_ErrDHTErrNoGossipersAvailable()
}
//go:cgo_export_dynamic var_setholochain_ErrDHTExpectedGossipReqInBody
//go:linkname _cgoexp_ebd397278953_var_setholochain_ErrDHTExpectedGossipReqInBody _cgoexp_ebd397278953_var_setholochain_ErrDHTExpectedGossipReqInBody
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_ErrDHTExpectedGossipReqInBody
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_ErrDHTExpectedGossipReqInBody(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_ErrDHTExpectedGossipReqInBody
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_ErrDHTExpectedGossipReqInBody(p0 _Ctype_int32_t) {
	var_setholochain_ErrDHTExpectedGossipReqInBody(p0)
}
//go:cgo_export_dynamic var_getholochain_ErrDHTExpectedGossipReqInBody
//go:linkname _cgoexp_ebd397278953_var_getholochain_ErrDHTExpectedGossipReqInBody _cgoexp_ebd397278953_var_getholochain_ErrDHTExpectedGossipReqInBody
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_ErrDHTExpectedGossipReqInBody
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_ErrDHTExpectedGossipReqInBody(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_ErrDHTExpectedGossipReqInBody
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_ErrDHTExpectedGossipReqInBody() (r0 _Ctype_int32_t) {
	return var_getholochain_ErrDHTExpectedGossipReqInBody()
}
//go:cgo_export_dynamic var_setholochain_ErrDHTUnexpectedTypeInBody
//go:linkname _cgoexp_ebd397278953_var_setholochain_ErrDHTUnexpectedTypeInBody _cgoexp_ebd397278953_var_setholochain_ErrDHTUnexpectedTypeInBody
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_ErrDHTUnexpectedTypeInBody
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_ErrDHTUnexpectedTypeInBody(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_ErrDHTUnexpectedTypeInBody
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_ErrDHTUnexpectedTypeInBody(p0 _Ctype_int32_t) {
	var_setholochain_ErrDHTUnexpectedTypeInBody(p0)
}
//go:cgo_export_dynamic var_getholochain_ErrDHTUnexpectedTypeInBody
//go:linkname _cgoexp_ebd397278953_var_getholochain_ErrDHTUnexpectedTypeInBody _cgoexp_ebd397278953_var_getholochain_ErrDHTUnexpectedTypeInBody
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_ErrDHTUnexpectedTypeInBody
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_ErrDHTUnexpectedTypeInBody(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_ErrDHTUnexpectedTypeInBody
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_ErrDHTUnexpectedTypeInBody() (r0 _Ctype_int32_t) {
	return var_getholochain_ErrDHTUnexpectedTypeInBody()
}
//go:cgo_export_dynamic var_setholochain_ErrEmptyRoutingTable
//go:linkname _cgoexp_ebd397278953_var_setholochain_ErrEmptyRoutingTable _cgoexp_ebd397278953_var_setholochain_ErrEmptyRoutingTable
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_ErrEmptyRoutingTable
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_ErrEmptyRoutingTable(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_ErrEmptyRoutingTable
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_ErrEmptyRoutingTable(p0 _Ctype_int32_t) {
	var_setholochain_ErrEmptyRoutingTable(p0)
}
//go:cgo_export_dynamic var_getholochain_ErrEmptyRoutingTable
//go:linkname _cgoexp_ebd397278953_var_getholochain_ErrEmptyRoutingTable _cgoexp_ebd397278953_var_getholochain_ErrEmptyRoutingTable
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_ErrEmptyRoutingTable
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_ErrEmptyRoutingTable(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_ErrEmptyRoutingTable
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_ErrEmptyRoutingTable() (r0 _Ctype_int32_t) {
	return var_getholochain_ErrEmptyRoutingTable()
}
//go:cgo_export_dynamic var_setholochain_ErrEntryTypeMismatch
//go:linkname _cgoexp_ebd397278953_var_setholochain_ErrEntryTypeMismatch _cgoexp_ebd397278953_var_setholochain_ErrEntryTypeMismatch
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_ErrEntryTypeMismatch
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_ErrEntryTypeMismatch(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_ErrEntryTypeMismatch
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_ErrEntryTypeMismatch(p0 _Ctype_int32_t) {
	var_setholochain_ErrEntryTypeMismatch(p0)
}
//go:cgo_export_dynamic var_getholochain_ErrEntryTypeMismatch
//go:linkname _cgoexp_ebd397278953_var_getholochain_ErrEntryTypeMismatch _cgoexp_ebd397278953_var_getholochain_ErrEntryTypeMismatch
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_ErrEntryTypeMismatch
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_ErrEntryTypeMismatch(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_ErrEntryTypeMismatch
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_ErrEntryTypeMismatch() (r0 _Ctype_int32_t) {
	return var_getholochain_ErrEntryTypeMismatch()
}
//go:cgo_export_dynamic var_setholochain_ErrHashDeleted
//go:linkname _cgoexp_ebd397278953_var_setholochain_ErrHashDeleted _cgoexp_ebd397278953_var_setholochain_ErrHashDeleted
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_ErrHashDeleted
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_ErrHashDeleted(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_ErrHashDeleted
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_ErrHashDeleted(p0 _Ctype_int32_t) {
	var_setholochain_ErrHashDeleted(p0)
}
//go:cgo_export_dynamic var_getholochain_ErrHashDeleted
//go:linkname _cgoexp_ebd397278953_var_getholochain_ErrHashDeleted _cgoexp_ebd397278953_var_getholochain_ErrHashDeleted
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_ErrHashDeleted
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_ErrHashDeleted(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_ErrHashDeleted
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_ErrHashDeleted() (r0 _Ctype_int32_t) {
	return var_getholochain_ErrHashDeleted()
}
//go:cgo_export_dynamic var_setholochain_ErrHashModified
//go:linkname _cgoexp_ebd397278953_var_setholochain_ErrHashModified _cgoexp_ebd397278953_var_setholochain_ErrHashModified
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_ErrHashModified
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_ErrHashModified(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_ErrHashModified
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_ErrHashModified(p0 _Ctype_int32_t) {
	var_setholochain_ErrHashModified(p0)
}
//go:cgo_export_dynamic var_getholochain_ErrHashModified
//go:linkname _cgoexp_ebd397278953_var_getholochain_ErrHashModified _cgoexp_ebd397278953_var_getholochain_ErrHashModified
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_ErrHashModified
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_ErrHashModified(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_ErrHashModified
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_ErrHashModified() (r0 _Ctype_int32_t) {
	return var_getholochain_ErrHashModified()
}
//go:cgo_export_dynamic var_setholochain_ErrHashNotFound
//go:linkname _cgoexp_ebd397278953_var_setholochain_ErrHashNotFound _cgoexp_ebd397278953_var_setholochain_ErrHashNotFound
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_ErrHashNotFound
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_ErrHashNotFound(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_ErrHashNotFound
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_ErrHashNotFound(p0 _Ctype_int32_t) {
	var_setholochain_ErrHashNotFound(p0)
}
//go:cgo_export_dynamic var_getholochain_ErrHashNotFound
//go:linkname _cgoexp_ebd397278953_var_getholochain_ErrHashNotFound _cgoexp_ebd397278953_var_getholochain_ErrHashNotFound
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_ErrHashNotFound
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_ErrHashNotFound(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_ErrHashNotFound
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_ErrHashNotFound() (r0 _Ctype_int32_t) {
	return var_getholochain_ErrHashNotFound()
}
//go:cgo_export_dynamic var_setholochain_ErrHashRejected
//go:linkname _cgoexp_ebd397278953_var_setholochain_ErrHashRejected _cgoexp_ebd397278953_var_setholochain_ErrHashRejected
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_ErrHashRejected
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_ErrHashRejected(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_ErrHashRejected
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_ErrHashRejected(p0 _Ctype_int32_t) {
	var_setholochain_ErrHashRejected(p0)
}
//go:cgo_export_dynamic var_getholochain_ErrHashRejected
//go:linkname _cgoexp_ebd397278953_var_getholochain_ErrHashRejected _cgoexp_ebd397278953_var_getholochain_ErrHashRejected
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_ErrHashRejected
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_ErrHashRejected(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_ErrHashRejected
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_ErrHashRejected() (r0 _Ctype_int32_t) {
	return var_getholochain_ErrHashRejected()
}
//go:cgo_export_dynamic var_setholochain_ErrIncompleteChain
//go:linkname _cgoexp_ebd397278953_var_setholochain_ErrIncompleteChain _cgoexp_ebd397278953_var_setholochain_ErrIncompleteChain
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_ErrIncompleteChain
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_ErrIncompleteChain(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_ErrIncompleteChain
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_ErrIncompleteChain(p0 _Ctype_int32_t) {
	var_setholochain_ErrIncompleteChain(p0)
}
//go:cgo_export_dynamic var_getholochain_ErrIncompleteChain
//go:linkname _cgoexp_ebd397278953_var_getholochain_ErrIncompleteChain _cgoexp_ebd397278953_var_getholochain_ErrIncompleteChain
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_ErrIncompleteChain
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_ErrIncompleteChain(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_ErrIncompleteChain
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_ErrIncompleteChain() (r0 _Ctype_int32_t) {
	return var_getholochain_ErrIncompleteChain()
}
//go:cgo_export_dynamic var_setholochain_ErrLinkNotFound
//go:linkname _cgoexp_ebd397278953_var_setholochain_ErrLinkNotFound _cgoexp_ebd397278953_var_setholochain_ErrLinkNotFound
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_ErrLinkNotFound
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_ErrLinkNotFound(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_ErrLinkNotFound
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_ErrLinkNotFound(p0 _Ctype_int32_t) {
	var_setholochain_ErrLinkNotFound(p0)
}
//go:cgo_export_dynamic var_getholochain_ErrLinkNotFound
//go:linkname _cgoexp_ebd397278953_var_getholochain_ErrLinkNotFound _cgoexp_ebd397278953_var_getholochain_ErrLinkNotFound
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_ErrLinkNotFound
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_ErrLinkNotFound(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_ErrLinkNotFound
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_ErrLinkNotFound() (r0 _Ctype_int32_t) {
	return var_getholochain_ErrLinkNotFound()
}
//go:cgo_export_dynamic var_setholochain_ErrNilEntryInvalid
//go:linkname _cgoexp_ebd397278953_var_setholochain_ErrNilEntryInvalid _cgoexp_ebd397278953_var_setholochain_ErrNilEntryInvalid
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_ErrNilEntryInvalid
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_ErrNilEntryInvalid(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_ErrNilEntryInvalid
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_ErrNilEntryInvalid(p0 _Ctype_int32_t) {
	var_setholochain_ErrNilEntryInvalid(p0)
}
//go:cgo_export_dynamic var_getholochain_ErrNilEntryInvalid
//go:linkname _cgoexp_ebd397278953_var_getholochain_ErrNilEntryInvalid _cgoexp_ebd397278953_var_getholochain_ErrNilEntryInvalid
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_ErrNilEntryInvalid
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_ErrNilEntryInvalid(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_ErrNilEntryInvalid
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_ErrNilEntryInvalid() (r0 _Ctype_int32_t) {
	return var_getholochain_ErrNilEntryInvalid()
}
//go:cgo_export_dynamic var_setholochain_ErrNoSuchIdx
//go:linkname _cgoexp_ebd397278953_var_setholochain_ErrNoSuchIdx _cgoexp_ebd397278953_var_setholochain_ErrNoSuchIdx
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_ErrNoSuchIdx
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_ErrNoSuchIdx(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_ErrNoSuchIdx
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_ErrNoSuchIdx(p0 _Ctype_int32_t) {
	var_setholochain_ErrNoSuchIdx(p0)
}
//go:cgo_export_dynamic var_getholochain_ErrNoSuchIdx
//go:linkname _cgoexp_ebd397278953_var_getholochain_ErrNoSuchIdx _cgoexp_ebd397278953_var_getholochain_ErrNoSuchIdx
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_ErrNoSuchIdx
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_ErrNoSuchIdx(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_ErrNoSuchIdx
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_ErrNoSuchIdx() (r0 _Ctype_int32_t) {
	return var_getholochain_ErrNoSuchIdx()
}
//go:cgo_export_dynamic var_setholochain_ErrNotValidForAgentType
//go:linkname _cgoexp_ebd397278953_var_setholochain_ErrNotValidForAgentType _cgoexp_ebd397278953_var_setholochain_ErrNotValidForAgentType
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_ErrNotValidForAgentType
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_ErrNotValidForAgentType(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_ErrNotValidForAgentType
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_ErrNotValidForAgentType(p0 _Ctype_int32_t) {
	var_setholochain_ErrNotValidForAgentType(p0)
}
//go:cgo_export_dynamic var_getholochain_ErrNotValidForAgentType
//go:linkname _cgoexp_ebd397278953_var_getholochain_ErrNotValidForAgentType _cgoexp_ebd397278953_var_getholochain_ErrNotValidForAgentType
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_ErrNotValidForAgentType
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_ErrNotValidForAgentType(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_ErrNotValidForAgentType
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_ErrNotValidForAgentType() (r0 _Ctype_int32_t) {
	return var_getholochain_ErrNotValidForAgentType()
}
//go:cgo_export_dynamic var_setholochain_ErrNotValidForDNAType
//go:linkname _cgoexp_ebd397278953_var_setholochain_ErrNotValidForDNAType _cgoexp_ebd397278953_var_setholochain_ErrNotValidForDNAType
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_ErrNotValidForDNAType
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_ErrNotValidForDNAType(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_ErrNotValidForDNAType
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_ErrNotValidForDNAType(p0 _Ctype_int32_t) {
	var_setholochain_ErrNotValidForDNAType(p0)
}
//go:cgo_export_dynamic var_getholochain_ErrNotValidForDNAType
//go:linkname _cgoexp_ebd397278953_var_getholochain_ErrNotValidForDNAType _cgoexp_ebd397278953_var_getholochain_ErrNotValidForDNAType
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_ErrNotValidForDNAType
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_ErrNotValidForDNAType(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_ErrNotValidForDNAType
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_ErrNotValidForDNAType() (r0 _Ctype_int32_t) {
	return var_getholochain_ErrNotValidForDNAType()
}
//go:cgo_export_dynamic var_setholochain_ErrNotValidForKeyType
//go:linkname _cgoexp_ebd397278953_var_setholochain_ErrNotValidForKeyType _cgoexp_ebd397278953_var_setholochain_ErrNotValidForKeyType
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_ErrNotValidForKeyType
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_ErrNotValidForKeyType(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_ErrNotValidForKeyType
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_ErrNotValidForKeyType(p0 _Ctype_int32_t) {
	var_setholochain_ErrNotValidForKeyType(p0)
}
//go:cgo_export_dynamic var_getholochain_ErrNotValidForKeyType
//go:linkname _cgoexp_ebd397278953_var_getholochain_ErrNotValidForKeyType _cgoexp_ebd397278953_var_getholochain_ErrNotValidForKeyType
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_ErrNotValidForKeyType
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_ErrNotValidForKeyType(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_ErrNotValidForKeyType
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_ErrNotValidForKeyType() (r0 _Ctype_int32_t) {
	return var_getholochain_ErrNotValidForKeyType()
}
//go:cgo_export_dynamic var_setholochain_ErrPutLinkOverDeleted
//go:linkname _cgoexp_ebd397278953_var_setholochain_ErrPutLinkOverDeleted _cgoexp_ebd397278953_var_setholochain_ErrPutLinkOverDeleted
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_ErrPutLinkOverDeleted
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_ErrPutLinkOverDeleted(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_ErrPutLinkOverDeleted
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_ErrPutLinkOverDeleted(p0 _Ctype_int32_t) {
	var_setholochain_ErrPutLinkOverDeleted(p0)
}
//go:cgo_export_dynamic var_getholochain_ErrPutLinkOverDeleted
//go:linkname _cgoexp_ebd397278953_var_getholochain_ErrPutLinkOverDeleted _cgoexp_ebd397278953_var_getholochain_ErrPutLinkOverDeleted
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_ErrPutLinkOverDeleted
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_ErrPutLinkOverDeleted(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_ErrPutLinkOverDeleted
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_ErrPutLinkOverDeleted() (r0 _Ctype_int32_t) {
	return var_getholochain_ErrPutLinkOverDeleted()
}
//go:cgo_export_dynamic var_setholochain_ErrWrongNargs
//go:linkname _cgoexp_ebd397278953_var_setholochain_ErrWrongNargs _cgoexp_ebd397278953_var_setholochain_ErrWrongNargs
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_ErrWrongNargs
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_ErrWrongNargs(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_ErrWrongNargs
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_ErrWrongNargs(p0 _Ctype_int32_t) {
	var_setholochain_ErrWrongNargs(p0)
}
//go:cgo_export_dynamic var_getholochain_ErrWrongNargs
//go:linkname _cgoexp_ebd397278953_var_getholochain_ErrWrongNargs _cgoexp_ebd397278953_var_getholochain_ErrWrongNargs
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_ErrWrongNargs
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_ErrWrongNargs(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_ErrWrongNargs
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_ErrWrongNargs() (r0 _Ctype_int32_t) {
	return var_getholochain_ErrWrongNargs()
}
//go:cgo_export_dynamic var_setholochain_IsDevMode
//go:linkname _cgoexp_ebd397278953_var_setholochain_IsDevMode _cgoexp_ebd397278953_var_setholochain_IsDevMode
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_IsDevMode
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_IsDevMode(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_IsDevMode
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_IsDevMode(p0 _Ctype_char) {
	var_setholochain_IsDevMode(p0)
}
//go:cgo_export_dynamic var_getholochain_IsDevMode
//go:linkname _cgoexp_ebd397278953_var_getholochain_IsDevMode _cgoexp_ebd397278953_var_getholochain_IsDevMode
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_IsDevMode
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_IsDevMode(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_IsDevMode
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_IsDevMode() (r0 _Ctype_char) {
	return var_getholochain_IsDevMode()
}
//go:cgo_export_dynamic var_setholochain_KValue
//go:linkname _cgoexp_ebd397278953_var_setholochain_KValue _cgoexp_ebd397278953_var_setholochain_KValue
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_KValue
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_KValue(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_KValue
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_KValue(p0 _Ctype_nint) {
	var_setholochain_KValue(p0)
}
//go:cgo_export_dynamic var_getholochain_KValue
//go:linkname _cgoexp_ebd397278953_var_getholochain_KValue _cgoexp_ebd397278953_var_getholochain_KValue
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_KValue
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_KValue(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_KValue
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_KValue() (r0 _Ctype_nint) {
	return var_getholochain_KValue()
}
//go:cgo_export_dynamic var_setholochain_KeyEntryDef
//go:linkname _cgoexp_ebd397278953_var_setholochain_KeyEntryDef _cgoexp_ebd397278953_var_setholochain_KeyEntryDef
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_KeyEntryDef
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_KeyEntryDef(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_KeyEntryDef
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_KeyEntryDef(p0 _Ctype_int32_t) {
	var_setholochain_KeyEntryDef(p0)
}
//go:cgo_export_dynamic var_getholochain_KeyEntryDef
//go:linkname _cgoexp_ebd397278953_var_getholochain_KeyEntryDef _cgoexp_ebd397278953_var_getholochain_KeyEntryDef
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_KeyEntryDef
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_KeyEntryDef(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_KeyEntryDef
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_KeyEntryDef() (r0 _Ctype_int32_t) {
	return var_getholochain_KeyEntryDef()
}
//go:cgo_export_dynamic var_setholochain_NonCallableAction
//go:linkname _cgoexp_ebd397278953_var_setholochain_NonCallableAction _cgoexp_ebd397278953_var_setholochain_NonCallableAction
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_NonCallableAction
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_NonCallableAction(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_NonCallableAction
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_NonCallableAction(p0 _Ctype_int32_t) {
	var_setholochain_NonCallableAction(p0)
}
//go:cgo_export_dynamic var_getholochain_NonCallableAction
//go:linkname _cgoexp_ebd397278953_var_getholochain_NonCallableAction _cgoexp_ebd397278953_var_getholochain_NonCallableAction
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_NonCallableAction
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_NonCallableAction(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_NonCallableAction
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_NonCallableAction() (r0 _Ctype_int32_t) {
	return var_getholochain_NonCallableAction()
}
//go:cgo_export_dynamic var_setholochain_NonDHTAction
//go:linkname _cgoexp_ebd397278953_var_setholochain_NonDHTAction _cgoexp_ebd397278953_var_setholochain_NonDHTAction
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_NonDHTAction
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_NonDHTAction(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_NonDHTAction
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_NonDHTAction(p0 _Ctype_int32_t) {
	var_setholochain_NonDHTAction(p0)
}
//go:cgo_export_dynamic var_getholochain_NonDHTAction
//go:linkname _cgoexp_ebd397278953_var_getholochain_NonDHTAction _cgoexp_ebd397278953_var_getholochain_NonDHTAction
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_NonDHTAction
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_NonDHTAction(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_NonDHTAction
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_NonDHTAction() (r0 _Ctype_int32_t) {
	return var_getholochain_NonDHTAction()
}
//go:cgo_export_dynamic var_setholochain_SelfRevocationDoesNotVerify
//go:linkname _cgoexp_ebd397278953_var_setholochain_SelfRevocationDoesNotVerify _cgoexp_ebd397278953_var_setholochain_SelfRevocationDoesNotVerify
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_SelfRevocationDoesNotVerify
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_SelfRevocationDoesNotVerify(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_SelfRevocationDoesNotVerify
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_SelfRevocationDoesNotVerify(p0 _Ctype_int32_t) {
	var_setholochain_SelfRevocationDoesNotVerify(p0)
}
//go:cgo_export_dynamic var_getholochain_SelfRevocationDoesNotVerify
//go:linkname _cgoexp_ebd397278953_var_getholochain_SelfRevocationDoesNotVerify _cgoexp_ebd397278953_var_getholochain_SelfRevocationDoesNotVerify
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_SelfRevocationDoesNotVerify
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_SelfRevocationDoesNotVerify(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_SelfRevocationDoesNotVerify
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_SelfRevocationDoesNotVerify() (r0 _Ctype_int32_t) {
	return var_getholochain_SelfRevocationDoesNotVerify()
}
//go:cgo_export_dynamic var_setholochain_SendTimeoutErr
//go:linkname _cgoexp_ebd397278953_var_setholochain_SendTimeoutErr _cgoexp_ebd397278953_var_setholochain_SendTimeoutErr
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_SendTimeoutErr
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_SendTimeoutErr(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_SendTimeoutErr
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_SendTimeoutErr(p0 _Ctype_int32_t) {
	var_setholochain_SendTimeoutErr(p0)
}
//go:cgo_export_dynamic var_getholochain_SendTimeoutErr
//go:linkname _cgoexp_ebd397278953_var_getholochain_SendTimeoutErr _cgoexp_ebd397278953_var_getholochain_SendTimeoutErr
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_SendTimeoutErr
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_SendTimeoutErr(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_SendTimeoutErr
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_SendTimeoutErr() (r0 _Ctype_int32_t) {
	return var_getholochain_SendTimeoutErr()
}
//go:cgo_export_dynamic var_setholochain_UnknownWarrantTypeErr
//go:linkname _cgoexp_ebd397278953_var_setholochain_UnknownWarrantTypeErr _cgoexp_ebd397278953_var_setholochain_UnknownWarrantTypeErr
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_UnknownWarrantTypeErr
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_UnknownWarrantTypeErr(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_UnknownWarrantTypeErr
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_UnknownWarrantTypeErr(p0 _Ctype_int32_t) {
	var_setholochain_UnknownWarrantTypeErr(p0)
}
//go:cgo_export_dynamic var_getholochain_UnknownWarrantTypeErr
//go:linkname _cgoexp_ebd397278953_var_getholochain_UnknownWarrantTypeErr _cgoexp_ebd397278953_var_getholochain_UnknownWarrantTypeErr
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_UnknownWarrantTypeErr
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_UnknownWarrantTypeErr(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_UnknownWarrantTypeErr
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_UnknownWarrantTypeErr() (r0 _Ctype_int32_t) {
	return var_getholochain_UnknownWarrantTypeErr()
}
//go:cgo_export_dynamic var_setholochain_ValidationFailedErr
//go:linkname _cgoexp_ebd397278953_var_setholochain_ValidationFailedErr _cgoexp_ebd397278953_var_setholochain_ValidationFailedErr
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_ValidationFailedErr
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_ValidationFailedErr(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_ValidationFailedErr
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_ValidationFailedErr(p0 _Ctype_int32_t) {
	var_setholochain_ValidationFailedErr(p0)
}
//go:cgo_export_dynamic var_getholochain_ValidationFailedErr
//go:linkname _cgoexp_ebd397278953_var_getholochain_ValidationFailedErr _cgoexp_ebd397278953_var_getholochain_ValidationFailedErr
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_ValidationFailedErr
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_ValidationFailedErr(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_ValidationFailedErr
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_ValidationFailedErr() (r0 _Ctype_int32_t) {
	return var_getholochain_ValidationFailedErr()
}
//go:cgo_export_dynamic var_setholochain_WarrantPropertyNotFoundErr
//go:linkname _cgoexp_ebd397278953_var_setholochain_WarrantPropertyNotFoundErr _cgoexp_ebd397278953_var_setholochain_WarrantPropertyNotFoundErr
//go:cgo_export_static _cgoexp_ebd397278953_var_setholochain_WarrantPropertyNotFoundErr
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_setholochain_WarrantPropertyNotFoundErr(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_setholochain_WarrantPropertyNotFoundErr
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_setholochain_WarrantPropertyNotFoundErr(p0 _Ctype_int32_t) {
	var_setholochain_WarrantPropertyNotFoundErr(p0)
}
//go:cgo_export_dynamic var_getholochain_WarrantPropertyNotFoundErr
//go:linkname _cgoexp_ebd397278953_var_getholochain_WarrantPropertyNotFoundErr _cgoexp_ebd397278953_var_getholochain_WarrantPropertyNotFoundErr
//go:cgo_export_static _cgoexp_ebd397278953_var_getholochain_WarrantPropertyNotFoundErr
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_var_getholochain_WarrantPropertyNotFoundErr(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_var_getholochain_WarrantPropertyNotFoundErr
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_var_getholochain_WarrantPropertyNotFoundErr() (r0 _Ctype_int32_t) {
	return var_getholochain_WarrantPropertyNotFoundErr()
}
//go:cgo_export_dynamic proxyholochain__BootstrapRefreshTask
//go:linkname _cgoexp_ebd397278953_proxyholochain__BootstrapRefreshTask _cgoexp_ebd397278953_proxyholochain__BootstrapRefreshTask
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__BootstrapRefreshTask
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__BootstrapRefreshTask(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__BootstrapRefreshTask
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__BootstrapRefreshTask(p0 _Ctype_int32_t) {
	proxyholochain__BootstrapRefreshTask(p0)
}
//go:cgo_export_dynamic proxyholochain__BuildJSONSchemaValidatorFromFile
//go:linkname _cgoexp_ebd397278953_proxyholochain__BuildJSONSchemaValidatorFromFile _cgoexp_ebd397278953_proxyholochain__BuildJSONSchemaValidatorFromFile
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__BuildJSONSchemaValidatorFromFile
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__BuildJSONSchemaValidatorFromFile(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__BuildJSONSchemaValidatorFromFile
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__BuildJSONSchemaValidatorFromFile(p0 _Ctype_struct_nstring) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain__BuildJSONSchemaValidatorFromFile(p0)
}
//go:cgo_export_dynamic proxyholochain__BuildJSONSchemaValidatorFromString
//go:linkname _cgoexp_ebd397278953_proxyholochain__BuildJSONSchemaValidatorFromString _cgoexp_ebd397278953_proxyholochain__BuildJSONSchemaValidatorFromString
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__BuildJSONSchemaValidatorFromString
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__BuildJSONSchemaValidatorFromString(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__BuildJSONSchemaValidatorFromString
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__BuildJSONSchemaValidatorFromString(p0 _Ctype_struct_nstring) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain__BuildJSONSchemaValidatorFromString(p0)
}
//go:cgo_export_dynamic proxyholochain__CopyDir
//go:linkname _cgoexp_ebd397278953_proxyholochain__CopyDir _cgoexp_ebd397278953_proxyholochain__CopyDir
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__CopyDir
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__CopyDir(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__CopyDir
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__CopyDir(p0 _Ctype_struct_nstring, p1 _Ctype_struct_nstring) (r0 _Ctype_int32_t) {
	return proxyholochain__CopyDir(p0, p1)
}
//go:cgo_export_dynamic proxyholochain__CopyFile
//go:linkname _cgoexp_ebd397278953_proxyholochain__CopyFile _cgoexp_ebd397278953_proxyholochain__CopyFile
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__CopyFile
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__CopyFile(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__CopyFile
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__CopyFile(p0 _Ctype_struct_nstring, p1 _Ctype_struct_nstring) (r0 _Ctype_int32_t) {
	return proxyholochain__CopyFile(p0, p1)
}
//go:cgo_export_dynamic proxyholochain__CreateRibosome
//go:linkname _cgoexp_ebd397278953_proxyholochain__CreateRibosome _cgoexp_ebd397278953_proxyholochain__CreateRibosome
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__CreateRibosome
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__CreateRibosome(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__CreateRibosome
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__CreateRibosome(p0 _Ctype_int32_t, p1 _Ctype_int32_t) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain__CreateRibosome(p0, p1)
}
//go:cgo_export_dynamic proxyholochain__Debug
//go:linkname _cgoexp_ebd397278953_proxyholochain__Debug _cgoexp_ebd397278953_proxyholochain__Debug
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__Debug
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__Debug(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__Debug
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__Debug(p0 _Ctype_struct_nstring) {
	proxyholochain__Debug(p0)
}
//go:cgo_export_dynamic proxyholochain__DecodeWarrant
//go:linkname _cgoexp_ebd397278953_proxyholochain__DecodeWarrant _cgoexp_ebd397278953_proxyholochain__DecodeWarrant
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__DecodeWarrant
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__DecodeWarrant(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__DecodeWarrant
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__DecodeWarrant(p0 _Ctype_nint, p1 _Ctype_struct_nbyteslice) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain__DecodeWarrant(p0, p1)
}
//go:cgo_export_dynamic proxyholochain__EncodingFormat
//go:linkname _cgoexp_ebd397278953_proxyholochain__EncodingFormat _cgoexp_ebd397278953_proxyholochain__EncodingFormat
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__EncodingFormat
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__EncodingFormat(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__EncodingFormat
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__EncodingFormat(p0 _Ctype_struct_nstring) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain__EncodingFormat(p0)
}
//go:cgo_export_dynamic proxyholochain__EscapeJSONValue
//go:linkname _cgoexp_ebd397278953_proxyholochain__EscapeJSONValue _cgoexp_ebd397278953_proxyholochain__EscapeJSONValue
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__EscapeJSONValue
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__EscapeJSONValue(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__EscapeJSONValue
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__EscapeJSONValue(p0 _Ctype_struct_nstring) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain__EscapeJSONValue(p0)
}
//go:cgo_export_dynamic proxyholochain__GossipTask
//go:linkname _cgoexp_ebd397278953_proxyholochain__GossipTask _cgoexp_ebd397278953_proxyholochain__GossipTask
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__GossipTask
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__GossipTask(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__GossipTask
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__GossipTask(p0 _Ctype_int32_t) {
	proxyholochain__GossipTask(p0)
}
//go:cgo_export_dynamic proxyholochain__Info
//go:linkname _cgoexp_ebd397278953_proxyholochain__Info _cgoexp_ebd397278953_proxyholochain__Info
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__Info
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__Info(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__Info
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__Info(p0 _Ctype_struct_nstring) {
	proxyholochain__Info(p0)
}
//go:cgo_export_dynamic proxyholochain__InitializeHolochain
//go:linkname _cgoexp_ebd397278953_proxyholochain__InitializeHolochain _cgoexp_ebd397278953_proxyholochain__InitializeHolochain
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__InitializeHolochain
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__InitializeHolochain(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__InitializeHolochain
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__InitializeHolochain() {
	proxyholochain__InitializeHolochain()
}
//go:cgo_export_dynamic proxyholochain__IsInitialized
//go:linkname _cgoexp_ebd397278953_proxyholochain__IsInitialized _cgoexp_ebd397278953_proxyholochain__IsInitialized
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__IsInitialized
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__IsInitialized(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__IsInitialized
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__IsInitialized(p0 _Ctype_struct_nstring) (r0 _Ctype_char) {
	return proxyholochain__IsInitialized(p0)
}
//go:cgo_export_dynamic proxyholochain__LoadAgent
//go:linkname _cgoexp_ebd397278953_proxyholochain__LoadAgent _cgoexp_ebd397278953_proxyholochain__LoadAgent
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__LoadAgent
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__LoadAgent(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__LoadAgent
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__LoadAgent(p0 _Ctype_struct_nstring) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain__LoadAgent(p0)
}
//go:cgo_export_dynamic proxyholochain__LoadService
//go:linkname _cgoexp_ebd397278953_proxyholochain__LoadService _cgoexp_ebd397278953_proxyholochain__LoadService
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__LoadService
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__LoadService(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__LoadService
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__LoadService(p0 _Ctype_struct_nstring) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain__LoadService(p0)
}
//go:cgo_export_dynamic proxyholochain__LoadTestConfig
//go:linkname _cgoexp_ebd397278953_proxyholochain__LoadTestConfig _cgoexp_ebd397278953_proxyholochain__LoadTestConfig
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__LoadTestConfig
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__LoadTestConfig(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__LoadTestConfig
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__LoadTestConfig(p0 _Ctype_struct_nstring) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain__LoadTestConfig(p0)
}
//go:cgo_export_dynamic proxyholochain__MakeActionFromMessage
//go:linkname _cgoexp_ebd397278953_proxyholochain__MakeActionFromMessage _cgoexp_ebd397278953_proxyholochain__MakeActionFromMessage
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__MakeActionFromMessage
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__MakeActionFromMessage(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__MakeActionFromMessage
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__MakeActionFromMessage(p0 _Ctype_int32_t) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain__MakeActionFromMessage(p0)
}
//go:cgo_export_dynamic proxyholochain__MakeDirs
//go:linkname _cgoexp_ebd397278953_proxyholochain__MakeDirs _cgoexp_ebd397278953_proxyholochain__MakeDirs
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__MakeDirs
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__MakeDirs(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__MakeDirs
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__MakeDirs(p0 _Ctype_struct_nstring) (r0 _Ctype_int32_t) {
	return proxyholochain__MakeDirs(p0)
}
//go:cgo_export_dynamic proxyholochain__MakeValidationPackage
//go:linkname _cgoexp_ebd397278953_proxyholochain__MakeValidationPackage _cgoexp_ebd397278953_proxyholochain__MakeValidationPackage
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__MakeValidationPackage
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__MakeValidationPackage(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__MakeValidationPackage
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__MakeValidationPackage(p0 _Ctype_int32_t, p1 _Ctype_int32_t) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain__MakeValidationPackage(p0, p1)
}
//go:cgo_export_dynamic proxyholochain__NewCommitAction
//go:linkname _cgoexp_ebd397278953_proxyholochain__NewCommitAction _cgoexp_ebd397278953_proxyholochain__NewCommitAction
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__NewCommitAction
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__NewCommitAction(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__NewCommitAction
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__NewCommitAction(p0 _Ctype_struct_nstring, p1 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain__NewCommitAction(p0, p1)
}
//go:cgo_export_dynamic proxyholochain__NewDHT
//go:linkname _cgoexp_ebd397278953_proxyholochain__NewDHT _cgoexp_ebd397278953_proxyholochain__NewDHT
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__NewDHT
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__NewDHT(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__NewDHT
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__NewDHT(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain__NewDHT(p0)
}
//go:cgo_export_dynamic proxyholochain__NewDebugAction
//go:linkname _cgoexp_ebd397278953_proxyholochain__NewDebugAction _cgoexp_ebd397278953_proxyholochain__NewDebugAction
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__NewDebugAction
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__NewDebugAction(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__NewDebugAction
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__NewDebugAction(p0 _Ctype_struct_nstring) (r0 _Ctype_int32_t) {
	return proxyholochain__NewDebugAction(p0)
}
//go:cgo_export_dynamic proxyholochain__NewGetBridgesAction
//go:linkname _cgoexp_ebd397278953_proxyholochain__NewGetBridgesAction _cgoexp_ebd397278953_proxyholochain__NewGetBridgesAction
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__NewGetBridgesAction
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__NewGetBridgesAction(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__NewGetBridgesAction
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__NewGetBridgesAction(p0 _Ctype_struct_nbyteslice) (r0 _Ctype_int32_t) {
	return proxyholochain__NewGetBridgesAction(p0)
}
//go:cgo_export_dynamic proxyholochain__NewGetLinksAction
//go:linkname _cgoexp_ebd397278953_proxyholochain__NewGetLinksAction _cgoexp_ebd397278953_proxyholochain__NewGetLinksAction
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__NewGetLinksAction
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__NewGetLinksAction(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__NewGetLinksAction
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__NewGetLinksAction(p0 _Ctype_int32_t, p1 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain__NewGetLinksAction(p0, p1)
}
//go:cgo_export_dynamic proxyholochain__NewJSRibosome
//go:linkname _cgoexp_ebd397278953_proxyholochain__NewJSRibosome _cgoexp_ebd397278953_proxyholochain__NewJSRibosome
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__NewJSRibosome
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__NewJSRibosome(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__NewJSRibosome
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__NewJSRibosome(p0 _Ctype_int32_t, p1 _Ctype_int32_t) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain__NewJSRibosome(p0, p1)
}
//go:cgo_export_dynamic proxyholochain__NewMakeHashAction
//go:linkname _cgoexp_ebd397278953_proxyholochain__NewMakeHashAction _cgoexp_ebd397278953_proxyholochain__NewMakeHashAction
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__NewMakeHashAction
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__NewMakeHashAction(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__NewMakeHashAction
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__NewMakeHashAction(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain__NewMakeHashAction(p0)
}
//go:cgo_export_dynamic proxyholochain__NewNode
//go:linkname _cgoexp_ebd397278953_proxyholochain__NewNode _cgoexp_ebd397278953_proxyholochain__NewNode
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__NewNode
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__NewNode(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__NewNode
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__NewNode(p0 _Ctype_struct_nstring, p1 _Ctype_struct_nstring, p2 _Ctype_int32_t, p3 _Ctype_char, p4 _Ctype_int32_t) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain__NewNode(p0, p1, p2, p3, p4)
}
//go:cgo_export_dynamic proxyholochain__NewNucleus
//go:linkname _cgoexp_ebd397278953_proxyholochain__NewNucleus _cgoexp_ebd397278953_proxyholochain__NewNucleus
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__NewNucleus
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__NewNucleus(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__NewNucleus
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__NewNucleus(p0 _Ctype_int32_t, p1 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain__NewNucleus(p0, p1)
}
//go:cgo_export_dynamic proxyholochain__NewPropertyAction
//go:linkname _cgoexp_ebd397278953_proxyholochain__NewPropertyAction _cgoexp_ebd397278953_proxyholochain__NewPropertyAction
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__NewPropertyAction
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__NewPropertyAction(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__NewPropertyAction
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__NewPropertyAction(p0 _Ctype_struct_nstring) (r0 _Ctype_int32_t) {
	return proxyholochain__NewPropertyAction(p0)
}
//go:cgo_export_dynamic proxyholochain__NewPutAction
//go:linkname _cgoexp_ebd397278953_proxyholochain__NewPutAction _cgoexp_ebd397278953_proxyholochain__NewPutAction
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__NewPutAction
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__NewPutAction(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__NewPutAction
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__NewPutAction(p0 _Ctype_struct_nstring, p1 _Ctype_int32_t, p2 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain__NewPutAction(p0, p1, p2)
}
//go:cgo_export_dynamic proxyholochain__NewQueryAction
//go:linkname _cgoexp_ebd397278953_proxyholochain__NewQueryAction _cgoexp_ebd397278953_proxyholochain__NewQueryAction
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__NewQueryAction
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__NewQueryAction(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__NewQueryAction
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__NewQueryAction(p0 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain__NewQueryAction(p0)
}
//go:cgo_export_dynamic proxyholochain__NewSelfRevocationWarrant
//go:linkname _cgoexp_ebd397278953_proxyholochain__NewSelfRevocationWarrant _cgoexp_ebd397278953_proxyholochain__NewSelfRevocationWarrant
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__NewSelfRevocationWarrant
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__NewSelfRevocationWarrant(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__NewSelfRevocationWarrant
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__NewSelfRevocationWarrant(p0 _Ctype_int32_t) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain__NewSelfRevocationWarrant(p0)
}
//go:cgo_export_dynamic proxyholochain__NewSignAction
//go:linkname _cgoexp_ebd397278953_proxyholochain__NewSignAction _cgoexp_ebd397278953_proxyholochain__NewSignAction
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__NewSignAction
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__NewSignAction(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__NewSignAction
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__NewSignAction(p0 _Ctype_struct_nbyteslice) (r0 _Ctype_int32_t) {
	return proxyholochain__NewSignAction(p0)
}
//go:cgo_export_dynamic proxyholochain__NewVerifySignatureAction
//go:linkname _cgoexp_ebd397278953_proxyholochain__NewVerifySignatureAction _cgoexp_ebd397278953_proxyholochain__NewVerifySignatureAction
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__NewVerifySignatureAction
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__NewVerifySignatureAction(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__NewVerifySignatureAction
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__NewVerifySignatureAction(p0 _Ctype_struct_nstring, p1 _Ctype_struct_nstring, p2 _Ctype_struct_nstring) (r0 _Ctype_int32_t) {
	return proxyholochain__NewVerifySignatureAction(p0, p1, p2)
}
//go:cgo_export_dynamic proxyholochain__NewZygoRibosome
//go:linkname _cgoexp_ebd397278953_proxyholochain__NewZygoRibosome _cgoexp_ebd397278953_proxyholochain__NewZygoRibosome
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__NewZygoRibosome
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__NewZygoRibosome(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__NewZygoRibosome
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__NewZygoRibosome(p0 _Ctype_int32_t, p1 _Ctype_int32_t) (r0 _Ctype_int32_t, r1 _Ctype_int32_t) {
	return proxyholochain__NewZygoRibosome(p0, p1)
}
//go:cgo_export_dynamic proxyholochain__PrettyPrintJSON
//go:linkname _cgoexp_ebd397278953_proxyholochain__PrettyPrintJSON _cgoexp_ebd397278953_proxyholochain__PrettyPrintJSON
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__PrettyPrintJSON
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__PrettyPrintJSON(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__PrettyPrintJSON
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__PrettyPrintJSON(p0 _Ctype_struct_nbyteslice) (r0 _Ctype_struct_nstring, r1 _Ctype_int32_t) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain__PrettyPrintJSON(p0)
}
//go:cgo_export_dynamic proxyholochain__RegisterBultinRibosomes
//go:linkname _cgoexp_ebd397278953_proxyholochain__RegisterBultinRibosomes _cgoexp_ebd397278953_proxyholochain__RegisterBultinRibosomes
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__RegisterBultinRibosomes
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__RegisterBultinRibosomes(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__RegisterBultinRibosomes
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__RegisterBultinRibosomes() {
	proxyholochain__RegisterBultinRibosomes()
}
//go:cgo_export_dynamic proxyholochain__RetryTask
//go:linkname _cgoexp_ebd397278953_proxyholochain__RetryTask _cgoexp_ebd397278953_proxyholochain__RetryTask
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__RetryTask
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__RetryTask(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__RetryTask
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__RetryTask(p0 _Ctype_int32_t) {
	proxyholochain__RetryTask(p0)
}
//go:cgo_export_dynamic proxyholochain__RoutingRefreshTask
//go:linkname _cgoexp_ebd397278953_proxyholochain__RoutingRefreshTask _cgoexp_ebd397278953_proxyholochain__RoutingRefreshTask
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__RoutingRefreshTask
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__RoutingRefreshTask(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__RoutingRefreshTask
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__RoutingRefreshTask(p0 _Ctype_int32_t) {
	proxyholochain__RoutingRefreshTask(p0)
}
//go:cgo_export_dynamic proxyholochain__SaveAgent
//go:linkname _cgoexp_ebd397278953_proxyholochain__SaveAgent _cgoexp_ebd397278953_proxyholochain__SaveAgent
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__SaveAgent
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__SaveAgent(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__SaveAgent
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__SaveAgent(p0 _Ctype_struct_nstring, p1 _Ctype_int32_t) (r0 _Ctype_int32_t) {
	return proxyholochain__SaveAgent(p0, p1)
}
//go:cgo_export_dynamic proxyholochain__TestingAppAppPackage
//go:linkname _cgoexp_ebd397278953_proxyholochain__TestingAppAppPackage _cgoexp_ebd397278953_proxyholochain__TestingAppAppPackage
//go:cgo_export_static _cgoexp_ebd397278953_proxyholochain__TestingAppAppPackage
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxyholochain__TestingAppAppPackage(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxyholochain__TestingAppAppPackage
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxyholochain__TestingAppAppPackage() (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxyholochain__TestingAppAppPackage()
}
//go:cgo_export_dynamic proxy_error_Error
//go:linkname _cgoexp_ebd397278953_proxy_error_Error _cgoexp_ebd397278953_proxy_error_Error
//go:cgo_export_static _cgoexp_ebd397278953_proxy_error_Error
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_proxy_error_Error(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_proxy_error_Error
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_proxy_error_Error(p0 _Ctype_int32_t) (r0 _Ctype_struct_nstring) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return proxy_error_Error(p0)
}
//go:cgo_export_dynamic IncGoRef
//go:linkname _cgoexp_ebd397278953_IncGoRef _cgoexp_ebd397278953_IncGoRef
//go:cgo_export_static _cgoexp_ebd397278953_IncGoRef
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_IncGoRef(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_IncGoRef
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_IncGoRef(p0 _Ctype_int32_t) {
	IncGoRef(p0)
}
//go:cgo_export_dynamic DestroyRef
//go:linkname _cgoexp_ebd397278953_DestroyRef _cgoexp_ebd397278953_DestroyRef
//go:cgo_export_static _cgoexp_ebd397278953_DestroyRef
//go:nosplit
//go:norace
func _cgoexp_ebd397278953_DestroyRef(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_ebd397278953_DestroyRef
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_ebd397278953_DestroyRef(p0 _Ctype_int32_t) {
	DestroyRef(p0)
}

//go:cgo_import_static _cgo_ebd397278953_Cfunc__Cmalloc
//go:linkname __cgofn__cgo_ebd397278953_Cfunc__Cmalloc _cgo_ebd397278953_Cfunc__Cmalloc
var __cgofn__cgo_ebd397278953_Cfunc__Cmalloc byte
var _cgo_ebd397278953_Cfunc__Cmalloc = unsafe.Pointer(&__cgofn__cgo_ebd397278953_Cfunc__Cmalloc)

//go:linkname runtime_throw runtime.throw
func runtime_throw(string)

//go:cgo_unsafe_args
func _cgo_cmalloc(p0 uint64) (r1 unsafe.Pointer) {
	_cgo_runtime_cgocall(_cgo_ebd397278953_Cfunc__Cmalloc, uintptr(unsafe.Pointer(&p0)))
	if r1 == nil {
		runtime_throw("runtime: C malloc failed")
	}
	return
}
