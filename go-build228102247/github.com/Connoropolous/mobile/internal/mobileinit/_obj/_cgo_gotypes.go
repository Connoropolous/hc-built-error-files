//go:cgo_ldflag "-target"
//go:cgo_ldflag "armv7a-none-linux-androideabi"
//go:cgo_ldflag "-gcc-toolchain"
//go:cgo_ldflag "/Users/connor/Library/Android/sdk/ndk-bundle/toolchains/arm-linux-androideabi-4.9/prebuilt/darwin-x86_64"
//go:cgo_ldflag "--sysroot"
//go:cgo_ldflag "/Users/connor/Library/Android/sdk/ndk-bundle/platforms/android-15/arch-arm"
//go:cgo_ldflag "-L/Users/connor/go/pkg/gomobile/lib/arm"
//go:cgo_ldflag "-landroid"
//go:cgo_ldflag "-llog"
// Created by cgo - DO NOT EDIT

package mobileinit

import "unsafe"

import _ "runtime/cgo"

import "syscall"

var _ syscall.Errno
func _Cgo_ptr(ptr unsafe.Pointer) unsafe.Pointer { return ptr }

//go:linkname _Cgo_always_false runtime.cgoAlwaysFalse
var _Cgo_always_false bool
//go:linkname _Cgo_use runtime.cgoUse
func _Cgo_use(interface{})
type _Ctype_JavaVM *_Ctype_struct_JNIInvokeInterface

type _Ctype___uintptr_t _Ctype_uint

type _Ctype_char uint8

type _Ctype_int int32

type _Ctype_jobject unsafe.Pointer

type _Ctype_struct_JNIInvokeInterface struct {
	reserved0			unsafe.Pointer
	reserved1			unsafe.Pointer
	reserved2			unsafe.Pointer
	DestroyJavaVM			*[0]byte
	AttachCurrentThread		*[0]byte
	DetachCurrentThread		*[0]byte
	GetEnv				*[0]byte
	AttachCurrentThreadAsDaemon	*[0]byte
}

type _Ctype_uint uint32

type _Ctype_uintptr_t _Ctype___uintptr_t

type _Ctype_void [0]byte

//go:linkname _cgo_runtime_cgocall runtime.cgocall
func _cgo_runtime_cgocall(unsafe.Pointer, uintptr) int32

//go:linkname _cgo_runtime_cgocallback runtime.cgocallback
func _cgo_runtime_cgocallback(unsafe.Pointer, unsafe.Pointer, uintptr, uintptr)

//go:linkname _cgoCheckPointer runtime.cgoCheckPointer
func _cgoCheckPointer(interface{}, ...interface{})

//go:linkname _cgoCheckResult runtime.cgoCheckResult
func _cgoCheckResult(interface{})
//go:linkname __cgo_current_ctx current_ctx
//go:cgo_import_static current_ctx
var __cgo_current_ctx byte
var _Cvar_current_ctx *_Ctype_jobject = (*_Ctype_jobject)(unsafe.Pointer(&__cgo_current_ctx))
//go:linkname __cgo_current_vm current_vm
//go:cgo_import_static current_vm
var __cgo_current_vm byte
var _Cvar_current_vm **_Ctype_JavaVM = (**_Ctype_JavaVM)(unsafe.Pointer(&__cgo_current_vm))
const _Ciconst_ANDROID_LOG_ERROR = 0x6
const _Ciconst_ANDROID_LOG_INFO = 0x4


func _Cfunc_CString(s string) *_Ctype_char {
	p := _cgo_cmalloc(uint64(len(s)+1))
	pp := (*[1<<30]byte)(p)
	copy(pp[:], s)
	pp[len(s)] = 0
	return (*_Ctype_char)(p)
}

//go:linkname _cgo_runtime_gostring runtime.gostring
func _cgo_runtime_gostring(*_Ctype_char) string

func _Cfunc_GoString(p *_Ctype_char) string {
	return _cgo_runtime_gostring(p)
}
//go:cgo_import_static _cgo_191360d7b6db_Cfunc___android_log_write
//go:linkname __cgofn__cgo_191360d7b6db_Cfunc___android_log_write _cgo_191360d7b6db_Cfunc___android_log_write
var __cgofn__cgo_191360d7b6db_Cfunc___android_log_write byte
var _cgo_191360d7b6db_Cfunc___android_log_write = unsafe.Pointer(&__cgofn__cgo_191360d7b6db_Cfunc___android_log_write)

//go:cgo_unsafe_args
func _Cfunc___android_log_write(p0 _Ctype_int, p1 *_Ctype_char, p2 *_Ctype_char) (r1 _Ctype_int) {
	_cgo_runtime_cgocall(_cgo_191360d7b6db_Cfunc___android_log_write, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
		_Cgo_use(p2)
	}
	return
}
//go:cgo_import_static _cgo_191360d7b6db_Cfunc_checkException
//go:linkname __cgofn__cgo_191360d7b6db_Cfunc_checkException _cgo_191360d7b6db_Cfunc_checkException
var __cgofn__cgo_191360d7b6db_Cfunc_checkException byte
var _cgo_191360d7b6db_Cfunc_checkException = unsafe.Pointer(&__cgofn__cgo_191360d7b6db_Cfunc_checkException)

//go:cgo_unsafe_args
func _Cfunc_checkException(p0 _Ctype_uintptr_t) (r1 *_Ctype_char) {
	_cgo_runtime_cgocall(_cgo_191360d7b6db_Cfunc_checkException, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_191360d7b6db_Cfunc_free
//go:linkname __cgofn__cgo_191360d7b6db_Cfunc_free _cgo_191360d7b6db_Cfunc_free
var __cgofn__cgo_191360d7b6db_Cfunc_free byte
var _cgo_191360d7b6db_Cfunc_free = unsafe.Pointer(&__cgofn__cgo_191360d7b6db_Cfunc_free)

//go:cgo_unsafe_args
func _Cfunc_free(p0 unsafe.Pointer) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_191360d7b6db_Cfunc_free, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_191360d7b6db_Cfunc_lockJNI
//go:linkname __cgofn__cgo_191360d7b6db_Cfunc_lockJNI _cgo_191360d7b6db_Cfunc_lockJNI
var __cgofn__cgo_191360d7b6db_Cfunc_lockJNI byte
var _cgo_191360d7b6db_Cfunc_lockJNI = unsafe.Pointer(&__cgofn__cgo_191360d7b6db_Cfunc_lockJNI)

//go:cgo_unsafe_args
func _Cfunc_lockJNI(p0 *_Ctype_uintptr_t, p1 *_Ctype_int) (r1 *_Ctype_char) {
	_cgo_runtime_cgocall(_cgo_191360d7b6db_Cfunc_lockJNI, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
	}
	return
}
//go:cgo_import_static _cgo_191360d7b6db_Cfunc_unlockJNI
//go:linkname __cgofn__cgo_191360d7b6db_Cfunc_unlockJNI _cgo_191360d7b6db_Cfunc_unlockJNI
var __cgofn__cgo_191360d7b6db_Cfunc_unlockJNI byte
var _cgo_191360d7b6db_Cfunc_unlockJNI = unsafe.Pointer(&__cgofn__cgo_191360d7b6db_Cfunc_unlockJNI)

//go:cgo_unsafe_args
func _Cfunc_unlockJNI() (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_191360d7b6db_Cfunc_unlockJNI, uintptr(unsafe.Pointer(&r1)))
	if _Cgo_always_false {
	}
	return
}

//go:cgo_import_static _cgo_191360d7b6db_Cfunc__Cmalloc
//go:linkname __cgofn__cgo_191360d7b6db_Cfunc__Cmalloc _cgo_191360d7b6db_Cfunc__Cmalloc
var __cgofn__cgo_191360d7b6db_Cfunc__Cmalloc byte
var _cgo_191360d7b6db_Cfunc__Cmalloc = unsafe.Pointer(&__cgofn__cgo_191360d7b6db_Cfunc__Cmalloc)

//go:linkname runtime_throw runtime.throw
func runtime_throw(string)

//go:cgo_unsafe_args
func _cgo_cmalloc(p0 uint64) (r1 unsafe.Pointer) {
	_cgo_runtime_cgocall(_cgo_191360d7b6db_Cfunc__Cmalloc, uintptr(unsafe.Pointer(&p0)))
	if r1 == nil {
		runtime_throw("runtime: C malloc failed")
	}
	return
}
