// Created by cgo - DO NOT EDIT

//line /Users/connor/go/src/github.com/Connoropolous/mobile/internal/mobileinit/mobileinit_android.go:5
package mobileinit

//line /Users/connor/go/src/github.com/Connoropolous/mobile/internal/mobileinit/mobileinit_android.go:5
import _cgo_unsafe "unsafe"

//line /Users/connor/go/src/github.com/Connoropolous/mobile/internal/mobileinit/mobileinit_android.go:27
import (
	"bufio"
	"log"
	"os"
	"syscall"
	"unsafe"
)

var (
														ctag	= _Cfunc_CString("GoLog")

//line /Users/connor/go/src/github.com/Connoropolous/mobile/internal/mobileinit/mobileinit_android.go:39
	stderr, stdout	*os.File
)

type infoWriter struct{}

func (infoWriter) Write(p []byte) (n int, err error) {
														cstr := _Cfunc_CString(string(p))
														_Cfunc___android_log_write(_Ciconst_ANDROID_LOG_INFO, ctag, cstr)
//line /Users/connor/go/src/github.com/Connoropolous/mobile/internal/mobileinit/mobileinit_android.go:46
	func(_cgo0 _cgo_unsafe.Pointer) {
//line /Users/connor/go/src/github.com/Connoropolous/mobile/internal/mobileinit/mobileinit_android.go:46
		_cgoCheckPointer(_cgo0)
															_Cfunc_free(_cgo0)
//line /Users/connor/go/src/github.com/Connoropolous/mobile/internal/mobileinit/mobileinit_android.go:47
	}(unsafe.Pointer(cstr))
														return len(p), nil
}

func lineLog(f *os.File, priority _Ctype_int) {
	const logSize = 1024
	r := bufio.NewReaderSize(f, logSize)
	for {
		line, _, err := r.ReadLine()
		str := string(line)
		if err != nil {
			str += " " + err.Error()
		}
															cstr := _Cfunc_CString(str)
															_Cfunc___android_log_write(priority, ctag, cstr)
//line /Users/connor/go/src/github.com/Connoropolous/mobile/internal/mobileinit/mobileinit_android.go:61
		func(_cgo0 _cgo_unsafe.Pointer) {
//line /Users/connor/go/src/github.com/Connoropolous/mobile/internal/mobileinit/mobileinit_android.go:61
			_cgoCheckPointer(_cgo0)
																_Cfunc_free(_cgo0)
//line /Users/connor/go/src/github.com/Connoropolous/mobile/internal/mobileinit/mobileinit_android.go:62
		}(unsafe.Pointer(cstr))
															if err != nil {
			break
		}
	}
}

func init() {
	log.SetOutput(infoWriter{})

	log.SetFlags(log.Flags() &^ log.LstdFlags)

	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stderr = w
	if err := syscall.Dup3(int(w.Fd()), int(os.Stderr.Fd()), 0); err != nil {
		panic(err)
	}
	go lineLog(r, _Ciconst_ANDROID_LOG_ERROR)

	r, w, err = os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout = w
	if err := syscall.Dup3(int(w.Fd()), int(os.Stdout.Fd()), 0); err != nil {
		panic(err)
	}
	go lineLog(r, _Ciconst_ANDROID_LOG_INFO)
}
