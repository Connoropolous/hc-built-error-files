// Created by cgo - DO NOT EDIT

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_main.go:5
package gomobile_bind

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_main.go:16
import (
	_seq "github.com/Connoropolous/mobile/bind/seq"
)

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_main.go:21
var _ = _seq.FromRefNum

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_main.go:24
func proxy_error_Error(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(error)
	res_0 := v.Error()
	_res_0 := encodeString(res_0)
	return _res_0
}

type proxy_error _seq.Ref

func (p *proxy_error) Bind_proxy_refnum__() int32	{ return (*_seq.Ref)(p).Bind_IncNum() }

func (p *proxy_error) Error() string {
	res := _Cfunc_cproxy_error_Error(_Ctype_int32_t(p.Bind_proxy_refnum__()))
	_res := decodeString(res)
	return _res
}
