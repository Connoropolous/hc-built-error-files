// Created by cgo - DO NOT EDIT

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/seq_android.go:5
package gomobile_bind

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/seq_android.go:5
import _cgo_unsafe "unsafe"

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/seq_android.go:18
import (
	"unsafe"

	"github.com/Connoropolous/mobile/bind/seq"
)

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/seq_android.go:26
func DestroyRef(refnum _Ctype_int32_t) {
	seq.Delete(int32(refnum))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/seq_android.go:35
func encodeString(s string) _Ctype_struct_nstring {
	n := _Ctype_int(len(s))
	if n == 0 {
		return _Ctype_struct_nstring{}
	}

	worstCaseLen := 4 * len(s)
	utf16buf := _Cfunc__CMalloc(_Ctype_size_t(worstCaseLen))
	if utf16buf == nil {
		panic("encodeString: malloc failed")
	}
	chars := (*[1<<30 - 1]uint16)(unsafe.Pointer(utf16buf))[:worstCaseLen/2 : worstCaseLen/2]
	nchars := seq.UTF16Encode(s, chars)
	return _Ctype_struct_nstring{chars: unsafe.Pointer(utf16buf), len: _Ctype_jsize(nchars * 2)}
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/seq_android.go:53
func decodeString(str _Ctype_struct_nstring) string {
	if str.chars == nil {
		return ""
	}
															chars := (*[1<<31 - 1]byte)(str.chars)[:str.len]
															s := string(chars)
//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/seq_android.go:58
	func(_cgo0 _cgo_unsafe.Pointer) {
//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/seq_android.go:58
		_cgoCheckPointer(_cgo0)
																_Cfunc_free(_cgo0)
//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/seq_android.go:59
	}(str.chars)
															return s
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/seq_android.go:65
func fromSlice(s []byte, cpy bool) _Ctype_struct_nbyteslice {
	if s == nil || len(s) == 0 {
		return _Ctype_struct_nbyteslice{}
	}
	var ptr *_Ctype_jbyte
	n := _Ctype_jsize(len(s))
	if cpy {
		ptr = (*_Ctype_jbyte)(_Cfunc__CMalloc(_Ctype_size_t(n)))
		if ptr == nil {
			panic("fromSlice: malloc failed")
		}
		copy((*[1<<31 - 1]byte)(unsafe.Pointer(ptr))[:n], s)
	} else {
		ptr = (*_Ctype_jbyte)(unsafe.Pointer(&s[0]))
	}
	return _Ctype_struct_nbyteslice{ptr: unsafe.Pointer(ptr), len: n}
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/seq_android.go:86
func toSlice(s _Ctype_struct_nbyteslice, cpy bool) []byte {
	if s.ptr == nil || s.len == 0 {
		return nil
	}
	var b []byte
	if cpy {
		b = func(_cgo0 _cgo_unsafe.Pointer, _cgo1 _Ctype_int) []byte {
//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/seq_android.go:92
			_cgoCheckPointer(_cgo0)
//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/seq_android.go:92
			return _Cfunc_GoBytes(_cgo0, _cgo1)
//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/seq_android.go:92
		}(s.ptr, _Ctype_int(s.len))
//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/seq_android.go:92
		func(_cgo0 _cgo_unsafe.Pointer) {
//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/seq_android.go:92
			_cgoCheckPointer(_cgo0)
																	_Cfunc_free(_cgo0)
//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/seq_android.go:93
		}(s.ptr)
	} else {
		b = (*[1<<31 - 1]byte)(unsafe.Pointer(s.ptr))[:s.len:s.len]
	}
	return b
}
