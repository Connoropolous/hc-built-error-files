//go:cgo_ldflag "-target"
//go:cgo_ldflag "armv7a-none-linux-androideabi"
//go:cgo_ldflag "-gcc-toolchain"
//go:cgo_ldflag "/Users/connor/Library/Android/sdk/ndk-bundle/toolchains/arm-linux-androideabi-4.9/prebuilt/darwin-x86_64"
//go:cgo_ldflag "--sysroot"
//go:cgo_ldflag "/Users/connor/Library/Android/sdk/ndk-bundle/platforms/android-15/arch-arm"
//go:cgo_ldflag "-L/Users/connor/go/pkg/gomobile/lib/arm"
//go:cgo_ldflag "-llog"
// Created by cgo - DO NOT EDIT

package java

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

type _Ctype_void [0]byte

//go:linkname _cgo_runtime_cgocall runtime.cgocall
func _cgo_runtime_cgocall(unsafe.Pointer, uintptr) int32

//go:linkname _cgo_runtime_cgocallback runtime.cgocallback
func _cgo_runtime_cgocallback(unsafe.Pointer, unsafe.Pointer, uintptr, uintptr)

//go:linkname _cgoCheckPointer runtime.cgoCheckPointer
func _cgoCheckPointer(interface{}, ...interface{})

//go:linkname _cgoCheckResult runtime.cgoCheckResult
func _cgoCheckResult(interface{})

//go:cgo_export_dynamic setContext
//go:linkname _cgoexp_4cfac081e983_setContext _cgoexp_4cfac081e983_setContext
//go:cgo_export_static _cgoexp_4cfac081e983_setContext
//go:nosplit
//go:norace
func _cgoexp_4cfac081e983_setContext(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_4cfac081e983_setContext
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_4cfac081e983_setContext(p0 *_Ctype_JavaVM, p1 _Ctype_jobject) {
	setContext(p0, p1)
}
