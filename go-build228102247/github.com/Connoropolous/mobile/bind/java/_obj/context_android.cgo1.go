// Created by cgo - DO NOT EDIT

//line /Users/connor/go/src/github.com/Connoropolous/mobile/bind/java/context_android.go:5
package java

//line /Users/connor/go/src/github.com/Connoropolous/mobile/bind/java/context_android.go:12
import (
	"unsafe"

	"github.com/Connoropolous/mobile/internal/mobileinit"
)

//line /Users/connor/go/src/github.com/Connoropolous/mobile/bind/java/context_android.go:19
func setContext(vm *_Ctype_JavaVM, ctx _Ctype_jobject) {
	mobileinit.SetCurrentContext(unsafe.Pointer(vm), unsafe.Pointer(ctx))
}
