// Created by cgo - DO NOT EDIT

//line /Users/connor/go/src/github.com/Connoropolous/mobile/internal/mobileinit/ctx_android.go:5
package mobileinit

//line /Users/connor/go/src/github.com/Connoropolous/mobile/internal/mobileinit/ctx_android.go:5
import _cgo_unsafe "unsafe"

//line /Users/connor/go/src/github.com/Connoropolous/mobile/internal/mobileinit/ctx_android.go:72
import (
	"errors"
	"runtime"
	"unsafe"
)

//line /Users/connor/go/src/github.com/Connoropolous/mobile/internal/mobileinit/ctx_android.go:81
func SetCurrentContext(vm, ctx unsafe.Pointer) {
	*_Cvar_current_vm = (*_Ctype_JavaVM)(vm)
	*_Cvar_current_ctx = (_Ctype_jobject)(ctx)
}

//line /Users/connor/go/src/github.com/Connoropolous/mobile/internal/mobileinit/ctx_android.go:94
func RunOnJVM(fn func(vm, env, ctx uintptr) error) error {
	errch := make(chan error)
	go func() {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()

		env := _Ctype_uintptr_t(0)
		attached := _Ctype_int(0)
		if errStr := _Cfunc_lockJNI(&env, &attached); errStr != nil {
			errch <- errors.New(_Cfunc_GoString(errStr))
			return
		}
		if attached != 0 {
			defer _Cfunc_unlockJNI()
		}

		vm := uintptr(unsafe.Pointer(*_Cvar_current_vm))
		if err := fn(vm, uintptr(env), uintptr(*_Cvar_current_ctx)); err != nil {
			errch <- err
			return
		}

		if exc := _Cfunc_checkException(env); exc != nil {
															errch <- errors.New(_Cfunc_GoString(exc))
//line /Users/connor/go/src/github.com/Connoropolous/mobile/internal/mobileinit/ctx_android.go:117
			func(_cgo0 _cgo_unsafe.Pointer) {
//line /Users/connor/go/src/github.com/Connoropolous/mobile/internal/mobileinit/ctx_android.go:117
				_cgoCheckPointer(_cgo0)
																_Cfunc_free(_cgo0)
//line /Users/connor/go/src/github.com/Connoropolous/mobile/internal/mobileinit/ctx_android.go:118
			}(unsafe.Pointer(exc))
															return
		}
		errch <- nil
	}()
	return <-errch
}
