// Created by cgo - DO NOT EDIT

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5
package gomobile_bind

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:16
import (
	_seq "github.com/Connoropolous/mobile/bind/seq"
	"github.com/Holochain/holochain-proto"
)

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:22
var _ = _seq.FromRefNum

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:29
func proxyholochain_ActionBridge_Name(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionBridge)
	res_0 := v.Name()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:38
func new_holochain_ActionBridge() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ActionBridge)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:47
func proxyholochain_ActionCall_Name(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionCall)
	res_0 := v.Name()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:56
func new_holochain_ActionCall() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ActionCall)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:63
func proxyholochain_ActionCommit_CheckValidationRequest(refnum _Ctype_int32_t, param_def _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionCommit)

	var _param_def *holochain.EntryDef
	if _param_def_ref := _seq.FromRefNum(int32(param_def)); _param_def_ref != nil {
		_param_def = _param_def_ref.Get().(*holochain.EntryDef)
	}
	res_0 := v.CheckValidationRequest(_param_def)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:82
func proxyholochain_ActionCommit_Entry(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionCommit)
	res_0 := v.Entry()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:94
func proxyholochain_ActionCommit_EntryType(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionCommit)
	res_0 := v.EntryType()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:103
func proxyholochain_ActionCommit_Name(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionCommit)
	res_0 := v.Name()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:114
func proxyholochain_ActionCommit_SetHeader(refnum _Ctype_int32_t, param_header _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionCommit)

	var _param_header *holochain.Header
	if _param_header_ref := _seq.FromRefNum(int32(param_header)); _param_header_ref != nil {
		_param_header = _param_header_ref.Get().(*holochain.Header)
	}
	v.SetHeader(_param_header)
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:128
func new_holochain_ActionCommit() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ActionCommit)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:137
func proxyholochain_ActionDebug_Name(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionDebug)
	res_0 := v.Name()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:146
func new_holochain_ActionDebug() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ActionDebug)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:153
func proxyholochain_ActionDel_CheckValidationRequest(refnum _Ctype_int32_t, param_def _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionDel)

	var _param_def *holochain.EntryDef
	if _param_def_ref := _seq.FromRefNum(int32(param_def)); _param_def_ref != nil {
		_param_def = _param_def_ref.Get().(*holochain.EntryDef)
	}
	res_0 := v.CheckValidationRequest(_param_def)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:172
func proxyholochain_ActionDel_Entry(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionDel)
	res_0 := v.Entry()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:184
func proxyholochain_ActionDel_EntryType(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionDel)
	res_0 := v.EntryType()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:193
func proxyholochain_ActionDel_Name(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionDel)
	res_0 := v.Name()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:204
func proxyholochain_ActionDel_SetHeader(refnum _Ctype_int32_t, param_header _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionDel)

	var _param_header *holochain.Header
	if _param_header_ref := _seq.FromRefNum(int32(param_header)); _param_header_ref != nil {
		_param_header = _param_header_ref.Get().(*holochain.Header)
	}
	v.SetHeader(_param_header)
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:218
func new_holochain_ActionDel() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ActionDel)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:227
func proxyholochain_ActionGet_Name(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionGet)
	res_0 := v.Name()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:240
func new_holochain_ActionGet() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ActionGet)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:249
func proxyholochain_ActionGetBridges_Name(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionGetBridges)
	res_0 := v.Name()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:258
func new_holochain_ActionGetBridges() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ActionGetBridges)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:267
func proxyholochain_ActionGetLinks_Name(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionGetLinks)
	res_0 := v.Name()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:280
func new_holochain_ActionGetLinks() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ActionGetLinks)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:287
func proxyholochain_ActionLink_CheckValidationRequest(refnum _Ctype_int32_t, param_def _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionLink)

	var _param_def *holochain.EntryDef
	if _param_def_ref := _seq.FromRefNum(int32(param_def)); _param_def_ref != nil {
		_param_def = _param_def_ref.Get().(*holochain.EntryDef)
	}
	res_0 := v.CheckValidationRequest(_param_def)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:306
func proxyholochain_ActionLink_Name(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionLink)
	res_0 := v.Name()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:319
func new_holochain_ActionLink() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ActionLink)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:328
func proxyholochain_ActionListAdd_Name(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionListAdd)
	res_0 := v.Name()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:339
func new_holochain_ActionListAdd() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ActionListAdd)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:348
func proxyholochain_ActionMakeHash_Name(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionMakeHash)
	res_0 := v.Name()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:357
func new_holochain_ActionMakeHash() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ActionMakeHash)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:364
func proxyholochain_ActionMod_CheckValidationRequest(refnum _Ctype_int32_t, param_def _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionMod)

	var _param_def *holochain.EntryDef
	if _param_def_ref := _seq.FromRefNum(int32(param_def)); _param_def_ref != nil {
		_param_def = _param_def_ref.Get().(*holochain.EntryDef)
	}
	res_0 := v.CheckValidationRequest(_param_def)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:383
func proxyholochain_ActionMod_Entry(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionMod)
	res_0 := v.Entry()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:395
func proxyholochain_ActionMod_EntryType(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionMod)
	res_0 := v.EntryType()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:404
func proxyholochain_ActionMod_Name(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionMod)
	res_0 := v.Name()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:415
func proxyholochain_ActionMod_SetHeader(refnum _Ctype_int32_t, param_header _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionMod)

	var _param_header *holochain.Header
	if _param_header_ref := _seq.FromRefNum(int32(param_header)); _param_header_ref != nil {
		_param_header = _param_header_ref.Get().(*holochain.Header)
	}
	v.SetHeader(_param_header)
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:429
func new_holochain_ActionMod() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ActionMod)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:436
func proxyholochain_ActionModAgent_Revocation_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.ActionModAgent).Revocation = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:443
func proxyholochain_ActionModAgent_Revocation_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionModAgent).Revocation
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:455
func proxyholochain_ActionModAgent_Name(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionModAgent)
	res_0 := v.Name()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:464
func new_holochain_ActionModAgent() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ActionModAgent)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:473
func proxyholochain_ActionProperty_Name(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionProperty)
	res_0 := v.Name()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:482
func new_holochain_ActionProperty() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ActionProperty)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:489
func proxyholochain_ActionPut_CheckValidationRequest(refnum _Ctype_int32_t, param_def _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionPut)

	var _param_def *holochain.EntryDef
	if _param_def_ref := _seq.FromRefNum(int32(param_def)); _param_def_ref != nil {
		_param_def = _param_def_ref.Get().(*holochain.EntryDef)
	}
	res_0 := v.CheckValidationRequest(_param_def)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:508
func proxyholochain_ActionPut_Name(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionPut)
	res_0 := v.Name()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:521
func new_holochain_ActionPut() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ActionPut)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:530
func proxyholochain_ActionQuery_Name(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionQuery)
	res_0 := v.Name()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:539
func new_holochain_ActionQuery() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ActionQuery)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:548
func proxyholochain_ActionSend_Name(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionSend)
	res_0 := v.Name()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:559
func new_holochain_ActionSend() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ActionSend)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:568
func proxyholochain_ActionSign_Name(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionSign)
	res_0 := v.Name()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:577
func new_holochain_ActionSign() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ActionSign)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:586
func proxyholochain_ActionVerifySignature_Name(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ActionVerifySignature)
	res_0 := v.Name()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:595
func new_holochain_ActionVerifySignature() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ActionVerifySignature)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:602
func proxyholochain_AgentEntry_Revocation_Set(refnum _Ctype_int32_t, v _Ctype_struct_nbyteslice) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := toSlice(v, true)
	ref.Get().(*holochain.AgentEntry).Revocation = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:609
func proxyholochain_AgentEntry_Revocation_Get(refnum _Ctype_int32_t) _Ctype_struct_nbyteslice {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.AgentEntry).Revocation
	_v := fromSlice(v, true)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:617
func proxyholochain_AgentEntry_PublicKey_Set(refnum _Ctype_int32_t, v _Ctype_struct_nbyteslice) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := toSlice(v, true)
	ref.Get().(*holochain.AgentEntry).PublicKey = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:624
func proxyholochain_AgentEntry_PublicKey_Get(refnum _Ctype_int32_t) _Ctype_struct_nbyteslice {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.AgentEntry).PublicKey
	_v := fromSlice(v, true)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:632
func new_holochain_AgentEntry() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.AgentEntry)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:637
func proxyholochain_AgentFixture_Hash_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.AgentFixture).Hash = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:644
func proxyholochain_AgentFixture_Hash_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.AgentFixture).Hash
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:652
func proxyholochain_AgentFixture_Identity_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.AgentFixture).Identity = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:659
func proxyholochain_AgentFixture_Identity_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.AgentFixture).Identity
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:667
func new_holochain_AgentFixture() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.AgentFixture)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:672
func proxyholochain_AppMsg_ZomeType_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.AppMsg).ZomeType = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:679
func proxyholochain_AppMsg_ZomeType_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.AppMsg).ZomeType
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:687
func proxyholochain_AppMsg_Body_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.AppMsg).Body = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:694
func proxyholochain_AppMsg_Body_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.AppMsg).Body
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:702
func new_holochain_AppMsg() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.AppMsg)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:707
func proxyholochain_AppPackage_Version_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.AppPackage).Version = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:714
func proxyholochain_AppPackage_Version_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.AppPackage).Version
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:722
func proxyholochain_AppPackage_Generator_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.AppPackage).Generator = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:729
func proxyholochain_AppPackage_Generator_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.AppPackage).Generator
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:745
func new_holochain_AppPackage() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.AppPackage)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:750
func proxyholochain_AppPackageScenario_Name_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.AppPackageScenario).Name = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:757
func proxyholochain_AppPackageScenario_Name_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.AppPackageScenario).Name
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:769
func new_holochain_AppPackageScenario() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.AppPackageScenario)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:774
func proxyholochain_AppPackageTests_Name_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.AppPackageTests).Name = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:781
func proxyholochain_AppPackageTests_Name_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.AppPackageTests).Name
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:791
func new_holochain_AppPackageTests() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.AppPackageTests)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:796
func proxyholochain_AppPackageUIFile_FileName_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.AppPackageUIFile).FileName = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:803
func proxyholochain_AppPackageUIFile_FileName_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.AppPackageUIFile).FileName
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:811
func proxyholochain_AppPackageUIFile_Data_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.AppPackageUIFile).Data = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:818
func proxyholochain_AppPackageUIFile_Data_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.AppPackageUIFile).Data
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:826
func proxyholochain_AppPackageUIFile_Encoding_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.AppPackageUIFile).Encoding = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:833
func proxyholochain_AppPackageUIFile_Encoding_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.AppPackageUIFile).Encoding
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:841
func new_holochain_AppPackageUIFile() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.AppPackageUIFile)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:846
func proxyholochain_Arg_Name_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.Arg).Name = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:853
func proxyholochain_Arg_Name_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Arg).Name
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:863
func proxyholochain_Arg_Optional_Set(refnum _Ctype_int32_t, v _Ctype_char) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := v != 0
	ref.Get().(*holochain.Arg).Optional = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:870
func proxyholochain_Arg_Optional_Get(refnum _Ctype_int32_t) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Arg).Optional
	var _v _Ctype_char = 0
	if v {
		_v = 1
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:883
func new_holochain_Arg() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Arg)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:888
func proxyholochain_BSReq_Version_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.BSReq).Version = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:895
func proxyholochain_BSReq_Version_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.BSReq).Version
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:903
func proxyholochain_BSReq_NodeID_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.BSReq).NodeID = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:910
func proxyholochain_BSReq_NodeID_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.BSReq).NodeID
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:918
func proxyholochain_BSReq_NodeAddr_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.BSReq).NodeAddr = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:925
func proxyholochain_BSReq_NodeAddr_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.BSReq).NodeAddr
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:933
func new_holochain_BSReq() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.BSReq)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:940
func proxyholochain_BSResp_Remote_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.BSResp).Remote = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:947
func proxyholochain_BSResp_Remote_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.BSResp).Remote
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:957
func new_holochain_BSResp() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.BSResp)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:964
func proxyholochain_Bridge_Token_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.Bridge).Token = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:971
func proxyholochain_Bridge_Token_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Bridge).Token
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:979
func proxyholochain_Bridge_Side_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.Bridge).Side = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:986
func proxyholochain_Bridge_Side_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Bridge).Side
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:994
func new_holochain_Bridge() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Bridge)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:999
func proxyholochain_BridgeApp_H_Set(refnum _Ctype_int32_t, v _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))

	var _v *holochain.Holochain
	if _v_ref := _seq.FromRefNum(int32(v)); _v_ref != nil {
		_v = _v_ref.Get().(*holochain.Holochain)
	}
	ref.Get().(*holochain.BridgeApp).H = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1010
func proxyholochain_BridgeApp_H_Get(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.BridgeApp).H
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1021
func proxyholochain_BridgeApp_Side_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.BridgeApp).Side = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1028
func proxyholochain_BridgeApp_Side_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.BridgeApp).Side
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1036
func proxyholochain_BridgeApp_BridgeGenesisDataFrom_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.BridgeApp).BridgeGenesisDataFrom = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1043
func proxyholochain_BridgeApp_BridgeGenesisDataFrom_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.BridgeApp).BridgeGenesisDataFrom
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1051
func proxyholochain_BridgeApp_BridgeGenesisDataTo_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.BridgeApp).BridgeGenesisDataTo = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1058
func proxyholochain_BridgeApp_BridgeGenesisDataTo_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.BridgeApp).BridgeGenesisDataTo
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1066
func proxyholochain_BridgeApp_Port_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.BridgeApp).Port = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1073
func proxyholochain_BridgeApp_Port_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.BridgeApp).Port
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1081
func new_holochain_BridgeApp() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.BridgeApp)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1088
func proxyholochain_Bucket_Len(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Bucket)
	res_0 := v.Len()
	_res_0 := _Ctype_nint(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1109
func new_holochain_Bucket() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Bucket)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1114
func proxyholochain_BytesSent_Bytes_Set(refnum _Ctype_int32_t, v _Ctype_int64_t) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int64(v)
	ref.Get().(*holochain.BytesSent).Bytes = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1121
func proxyholochain_BytesSent_Bytes_Get(refnum _Ctype_int32_t) _Ctype_int64_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.BytesSent).Bytes
	_v := _Ctype_int64_t(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1131
func new_holochain_BytesSent() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.BytesSent)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1136
func proxyholochain_Callback_Function_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.Callback).Function = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1143
func proxyholochain_Callback_Function_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Callback).Function
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1151
func proxyholochain_Callback_ID_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.Callback).ID = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1158
func proxyholochain_Callback_ID_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Callback).ID
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1166
func new_holochain_Callback() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Callback)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1171
func proxyholochain_Capability_Token_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.Capability).Token = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1178
func proxyholochain_Capability_Token_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Capability).Token
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1190
func new_holochain_Capability() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Capability)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1209
func proxyholochain_Chain_Close(refnum _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Chain)
	v.Close()
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1222
func proxyholochain_Chain_JSON(refnum _Ctype_int32_t) (_Ctype_struct_nstring, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Chain)
	res_0, res_1 := v.JSON()
	_res_0 := encodeString(res_0)
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1235
func proxyholochain_Chain_Length(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Chain)
	res_0 := v.Length()
	_res_0 := _Ctype_nint(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1246
func proxyholochain_Chain_Nth(refnum _Ctype_int32_t, param_n _Ctype_nint) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Chain)
	_param_n := int(param_n)
	res_0 := v.Nth(_param_n)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1259
func proxyholochain_Chain_String(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Chain)
	res_0 := v.String()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1268
func proxyholochain_Chain_Top(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Chain)
	res_0 := v.Top()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1282
func proxyholochain_Chain_Validate(refnum _Ctype_int32_t, param_skipEntries _Ctype_char) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Chain)
	_param_skipEntries := param_skipEntries != 0
	res_0 := v.Validate(_param_skipEntries)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1297
func new_holochain_Chain() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Chain)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1302
func proxyholochain_ChainPair_Header_Set(refnum _Ctype_int32_t, v _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))

	var _v *holochain.Header
	if _v_ref := _seq.FromRefNum(int32(v)); _v_ref != nil {
		_v = _v_ref.Get().(*holochain.Header)
	}
	ref.Get().(*holochain.ChainPair).Header = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1313
func proxyholochain_ChainPair_Header_Get(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ChainPair).Header
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1324
func proxyholochain_ChainPair_Entry_Set(refnum _Ctype_int32_t, v _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	var _v holochain.Entry
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(holochain.Entry)
		} else {
			_v = (*proxyholochain_Entry)(_v_ref)
		}
	}
	ref.Get().(*holochain.ChainPair).Entry = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1339
func proxyholochain_ChainPair_Entry_Get(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ChainPair).Entry
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1350
func new_holochain_ChainPair() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ChainPair)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1357
func proxyholochain_Change_Message_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.Change).Message = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1364
func proxyholochain_Change_Message_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Change).Message
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1372
func proxyholochain_Change_AsOf_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.Change).AsOf = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1379
func proxyholochain_Change_AsOf_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Change).AsOf
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1387
func proxyholochain_Change_Log(refnum _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Change)
	v.Log()
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1394
func new_holochain_Change() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Change)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1399
func proxyholochain_CloneSpec_Role_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.CloneSpec).Role = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1406
func proxyholochain_CloneSpec_Role_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.CloneSpec).Role
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1414
func proxyholochain_CloneSpec_Number_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.CloneSpec).Number = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1421
func proxyholochain_CloneSpec_Number_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.CloneSpec).Number
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1429
func new_holochain_CloneSpec() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.CloneSpec)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1436
func new_holochain_CloserPeersResp() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.CloserPeersResp)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1441
func proxyholochain_Config_Port_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.Config).Port = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1448
func proxyholochain_Config_Port_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Config).Port
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1456
func proxyholochain_Config_EnableMDNS_Set(refnum _Ctype_int32_t, v _Ctype_char) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := v != 0
	ref.Get().(*holochain.Config).EnableMDNS = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1463
func proxyholochain_Config_EnableMDNS_Get(refnum _Ctype_int32_t) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Config).EnableMDNS
	var _v _Ctype_char = 0
	if v {
		_v = 1
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1474
func proxyholochain_Config_PeerModeAuthor_Set(refnum _Ctype_int32_t, v _Ctype_char) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := v != 0
	ref.Get().(*holochain.Config).PeerModeAuthor = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1481
func proxyholochain_Config_PeerModeAuthor_Get(refnum _Ctype_int32_t) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Config).PeerModeAuthor
	var _v _Ctype_char = 0
	if v {
		_v = 1
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1492
func proxyholochain_Config_PeerModeDHTNode_Set(refnum _Ctype_int32_t, v _Ctype_char) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := v != 0
	ref.Get().(*holochain.Config).PeerModeDHTNode = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1499
func proxyholochain_Config_PeerModeDHTNode_Get(refnum _Ctype_int32_t) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Config).PeerModeDHTNode
	var _v _Ctype_char = 0
	if v {
		_v = 1
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1510
func proxyholochain_Config_EnableNATUPnP_Set(refnum _Ctype_int32_t, v _Ctype_char) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := v != 0
	ref.Get().(*holochain.Config).EnableNATUPnP = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1517
func proxyholochain_Config_EnableNATUPnP_Get(refnum _Ctype_int32_t) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Config).EnableNATUPnP
	var _v _Ctype_char = 0
	if v {
		_v = 1
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1528
func proxyholochain_Config_BootstrapServer_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.Config).BootstrapServer = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1535
func proxyholochain_Config_BootstrapServer_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Config).BootstrapServer
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1547
func proxyholochain_Config_Setup(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Config)
	res_0 := v.Setup()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1559
func proxyholochain_Config_SetupLogging(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Config)
	res_0 := v.SetupLogging()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1571
func new_holochain_Config() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Config)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1580
func proxyholochain_DHT_Close(refnum _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.DHT)
	v.Close()
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1589
func proxyholochain_DHT_DumpIdx(refnum _Ctype_int32_t, param_idx _Ctype_nint) (_Ctype_struct_nstring, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.DHT)
	_param_idx := int(param_idx)
	res_0, res_1 := v.DumpIdx(_param_idx)
	_res_0 := encodeString(res_0)
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1603
func proxyholochain_DHT_DumpIdxJSON(refnum _Ctype_int32_t, param_idx _Ctype_nint) (_Ctype_struct_nstring, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.DHT)
	_param_idx := int(param_idx)
	res_0, res_1 := v.DumpIdxJSON(_param_idx)
	_res_0 := encodeString(res_0)
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1623
func proxyholochain_DHT_GetIdx(refnum _Ctype_int32_t) (_Ctype_nint, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.DHT)
	res_0, res_1 := v.GetIdx()
	_res_0 := _Ctype_nint(res_0)
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1640
func proxyholochain_DHT_HandleGossipPuts(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.DHT)
	res_0 := v.HandleGossipPuts()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1652
func proxyholochain_DHT_HandleGossipWiths(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.DHT)
	res_0 := v.HandleGossipWiths()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1666
func proxyholochain_DHT_JSON(refnum _Ctype_int32_t) (_Ctype_struct_nstring, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.DHT)
	res_0, res_1 := v.JSON()
	_res_0 := encodeString(res_0)
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1681
func proxyholochain_DHT_SetupDHT(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.DHT)
	res_0 := v.SetupDHT()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1693
func proxyholochain_DHT_Start(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.DHT)
	res_0 := v.Start()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1705
func proxyholochain_DHT_String(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.DHT)
	res_0 := v.String()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1716
func new_holochain_DHT() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.DHT)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1721
func proxyholochain_DHTConfig_HashType_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.DHTConfig).HashType = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1728
func proxyholochain_DHTConfig_HashType_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.DHTConfig).HashType
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1736
func proxyholochain_DHTConfig_NeighborhoodSize_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.DHTConfig).NeighborhoodSize = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1743
func proxyholochain_DHTConfig_NeighborhoodSize_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.DHTConfig).NeighborhoodSize
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1751
func new_holochain_DHTConfig() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.DHTConfig)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1756
func proxyholochain_DNA_Version_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.DNA).Version = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1763
func proxyholochain_DNA_Version_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.DNA).Version
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1773
func proxyholochain_DNA_Name_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.DNA).Name = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1780
func proxyholochain_DNA_Name_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.DNA).Name
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1790
func proxyholochain_DNA_PropertiesSchema_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.DNA).PropertiesSchema = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1797
func proxyholochain_DNA_PropertiesSchema_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.DNA).PropertiesSchema
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1805
func proxyholochain_DNA_AgentIdentitySchema_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.DNA).AgentIdentitySchema = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1812
func proxyholochain_DNA_AgentIdentitySchema_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.DNA).AgentIdentitySchema
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1822
func proxyholochain_DNA_RequiresVersion_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.DNA).RequiresVersion = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1829
func proxyholochain_DNA_RequiresVersion_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.DNA).RequiresVersion
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1843
func proxyholochain_DNA_NewUUID(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.DNA)
	res_0 := v.NewUUID()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1855
func new_holochain_DNA() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.DNA)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1860
func proxyholochain_DNAFile_Version_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.DNAFile).Version = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1867
func proxyholochain_DNAFile_Version_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.DNAFile).Version
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1877
func proxyholochain_DNAFile_Name_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.DNAFile).Name = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1884
func proxyholochain_DNAFile_Name_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.DNAFile).Name
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1894
func proxyholochain_DNAFile_PropertiesSchemaFile_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.DNAFile).PropertiesSchemaFile = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1901
func proxyholochain_DNAFile_PropertiesSchemaFile_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.DNAFile).PropertiesSchemaFile
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1911
func proxyholochain_DNAFile_RequiresVersion_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.DNAFile).RequiresVersion = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1918
func proxyholochain_DNAFile_RequiresVersion_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.DNAFile).RequiresVersion
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1932
func new_holochain_DNAFile() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.DNAFile)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1939
func proxyholochain_DelEntry_Message_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.DelEntry).Message = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1946
func proxyholochain_DelEntry_Message_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.DelEntry).Message
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1954
func new_holochain_DelEntry() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.DelEntry)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1963
func new_holochain_DelReq() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.DelReq)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1968
func proxyholochain_EntryDef_Name_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.EntryDef).Name = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1975
func proxyholochain_EntryDef_Name_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.EntryDef).Name
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1983
func proxyholochain_EntryDef_DataFormat_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.EntryDef).DataFormat = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1990
func proxyholochain_EntryDef_DataFormat_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.EntryDef).DataFormat
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:1998
func proxyholochain_EntryDef_Sharing_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.EntryDef).Sharing = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2005
func proxyholochain_EntryDef_Sharing_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.EntryDef).Sharing
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2013
func proxyholochain_EntryDef_Schema_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.EntryDef).Schema = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2020
func proxyholochain_EntryDef_Schema_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.EntryDef).Schema
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2028
func proxyholochain_EntryDef_BuildJSONSchemaValidator(refnum _Ctype_int32_t, param_path _Ctype_struct_nstring) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.EntryDef)
	_param_path := decodeString(param_path)
	res_0 := v.BuildJSONSchemaValidator(_param_path)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2041
func proxyholochain_EntryDef_BuildJSONSchemaValidatorFromString(refnum _Ctype_int32_t, param_schema _Ctype_struct_nstring) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.EntryDef)
	_param_schema := decodeString(param_schema)
	res_0 := v.BuildJSONSchemaValidatorFromString(_param_schema)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2054
func proxyholochain_EntryDef_IsSysEntry(refnum _Ctype_int32_t) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.EntryDef)
	res_0 := v.IsSysEntry()
	var _res_0 _Ctype_char = 0
	if res_0 {
		_res_0 = 1
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2066
func proxyholochain_EntryDef_IsVirtualEntry(refnum _Ctype_int32_t) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.EntryDef)
	res_0 := v.IsVirtualEntry()
	var _res_0 _Ctype_char = 0
	if res_0 {
		_res_0 = 1
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2078
func new_holochain_EntryDef() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.EntryDef)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2083
func proxyholochain_EntryDefFile_Name_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.EntryDefFile).Name = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2090
func proxyholochain_EntryDefFile_Name_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.EntryDefFile).Name
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2098
func proxyholochain_EntryDefFile_DataFormat_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.EntryDefFile).DataFormat = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2105
func proxyholochain_EntryDefFile_DataFormat_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.EntryDefFile).DataFormat
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2113
func proxyholochain_EntryDefFile_Schema_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.EntryDefFile).Schema = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2120
func proxyholochain_EntryDefFile_Schema_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.EntryDefFile).Schema
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2128
func proxyholochain_EntryDefFile_SchemaFile_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.EntryDefFile).SchemaFile = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2135
func proxyholochain_EntryDefFile_SchemaFile_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.EntryDefFile).SchemaFile
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2143
func proxyholochain_EntryDefFile_Sharing_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.EntryDefFile).Sharing = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2150
func proxyholochain_EntryDefFile_Sharing_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.EntryDefFile).Sharing
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2158
func new_holochain_EntryDefFile() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.EntryDefFile)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2163
func proxyholochain_ErrorResponse_Code_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.ErrorResponse).Code = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2170
func proxyholochain_ErrorResponse_Code_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ErrorResponse).Code
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2178
func proxyholochain_ErrorResponse_Message_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.ErrorResponse).Message = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2185
func proxyholochain_ErrorResponse_Message_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ErrorResponse).Message
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2195
func proxyholochain_ErrorResponse_DecodeResponseError(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ErrorResponse)
	res_0 := v.DecodeResponseError()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2207
func new_holochain_ErrorResponse() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ErrorResponse)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2214
func new_holochain_FindNodeReq() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.FindNodeReq)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2219
func proxyholochain_FunctionDef_Name_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.FunctionDef).Name = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2226
func proxyholochain_FunctionDef_Name_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.FunctionDef).Name
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2234
func proxyholochain_FunctionDef_CallingType_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.FunctionDef).CallingType = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2241
func proxyholochain_FunctionDef_CallingType_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.FunctionDef).CallingType
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2249
func proxyholochain_FunctionDef_Exposure_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.FunctionDef).Exposure = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2256
func proxyholochain_FunctionDef_Exposure_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.FunctionDef).Exposure
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2264
func proxyholochain_FunctionDef_ValidExposure(refnum _Ctype_int32_t, param_context _Ctype_struct_nstring) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.FunctionDef)
	_param_context := decodeString(param_context)
	res_0 := v.ValidExposure(_param_context)
	var _res_0 _Ctype_char = 0
	if res_0 {
		_res_0 = 1
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2277
func new_holochain_FunctionDef() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.FunctionDef)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2282
func proxyholochain_GetLinksOptions_Load_Set(refnum _Ctype_int32_t, v _Ctype_char) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := v != 0
	ref.Get().(*holochain.GetLinksOptions).Load = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2289
func proxyholochain_GetLinksOptions_Load_Get(refnum _Ctype_int32_t) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.GetLinksOptions).Load
	var _v _Ctype_char = 0
	if v {
		_v = 1
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2300
func proxyholochain_GetLinksOptions_StatusMask_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.GetLinksOptions).StatusMask = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2307
func proxyholochain_GetLinksOptions_StatusMask_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.GetLinksOptions).StatusMask
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2315
func new_holochain_GetLinksOptions() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.GetLinksOptions)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2320
func proxyholochain_GetOptions_StatusMask_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.GetOptions).StatusMask = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2327
func proxyholochain_GetOptions_StatusMask_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.GetOptions).StatusMask
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2335
func proxyholochain_GetOptions_GetMask_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.GetOptions).GetMask = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2342
func proxyholochain_GetOptions_GetMask_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.GetOptions).GetMask
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2350
func proxyholochain_GetOptions_Local_Set(refnum _Ctype_int32_t, v _Ctype_char) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := v != 0
	ref.Get().(*holochain.GetOptions).Local = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2357
func proxyholochain_GetOptions_Local_Get(refnum _Ctype_int32_t) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.GetOptions).Local
	var _v _Ctype_char = 0
	if v {
		_v = 1
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2368
func new_holochain_GetOptions() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.GetOptions)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2375
func proxyholochain_GetReq_StatusMask_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.GetReq).StatusMask = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2382
func proxyholochain_GetReq_StatusMask_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.GetReq).StatusMask
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2390
func proxyholochain_GetReq_GetMask_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.GetReq).GetMask = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2397
func proxyholochain_GetReq_GetMask_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.GetReq).GetMask
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2405
func new_holochain_GetReq() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.GetReq)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2412
func proxyholochain_GetResp_EntryType_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.GetResp).EntryType = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2419
func proxyholochain_GetResp_EntryType_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.GetResp).EntryType
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2429
func proxyholochain_GetResp_FollowHash_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.GetResp).FollowHash = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2436
func proxyholochain_GetResp_FollowHash_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.GetResp).FollowHash
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2444
func new_holochain_GetResp() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.GetResp)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2453
func proxyholochain_GobEntry_Marshal(refnum _Ctype_int32_t) (_Ctype_struct_nbyteslice, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.GobEntry)
	res_0, res_1 := v.Marshal()
	_res_0 := fromSlice(res_0, true)
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2468
func proxyholochain_GobEntry_Unmarshal(refnum _Ctype_int32_t, param_b _Ctype_struct_nbyteslice) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.GobEntry)
	_param_b := toSlice(param_b, false)
	res_0 := v.Unmarshal(_param_b)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2481
func new_holochain_GobEntry() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.GobEntry)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2488
func new_holochain_Gossip() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Gossip)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2493
func proxyholochain_GossipReq_MyIdx_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.GossipReq).MyIdx = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2500
func proxyholochain_GossipReq_MyIdx_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.GossipReq).MyIdx
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2508
func proxyholochain_GossipReq_YourIdx_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.GossipReq).YourIdx = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2515
func proxyholochain_GossipReq_YourIdx_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.GossipReq).YourIdx
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2523
func new_holochain_GossipReq() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.GossipReq)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2528
func proxyholochain_Header_Type_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.Header).Type = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2535
func proxyholochain_Header_Type_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Header).Type
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2555
func proxyholochain_Header_Marshal(refnum _Ctype_int32_t) (_Ctype_struct_nbyteslice, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Header)
	res_0, res_1 := v.Marshal()
	_res_0 := fromSlice(res_0, true)
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2570
func proxyholochain_Header_Unmarshal(refnum _Ctype_int32_t, param_b _Ctype_struct_nbyteslice, param_hashSize _Ctype_nint) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Header)
	_param_b := toSlice(param_b, false)
	_param_hashSize := int(param_hashSize)
	res_0 := v.Unmarshal(_param_b, _param_hashSize)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2584
func new_holochain_Header() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Header)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2591
func proxyholochain_Holochain_Activate(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	res_0 := v.Activate()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2611
func proxyholochain_Holochain_Agent(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	res_0 := v.Agent()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2627
func proxyholochain_Holochain_BSget(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	res_0 := v.BSget()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2639
func proxyholochain_Holochain_BSpost(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	res_0 := v.BSpost()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2653
func proxyholochain_Holochain_BuildBridge(refnum _Ctype_int32_t, param_app _Ctype_int32_t, param_port _Ctype_struct_nstring) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)

	var _param_app *holochain.BridgeApp
	if _param_app_ref := _seq.FromRefNum(int32(param_app)); _param_app_ref != nil {
		_param_app = _param_app_ref.Get().(*holochain.BridgeApp)
	}
	_param_port := decodeString(param_port)
	res_0 := v.BuildBridge(_param_app, _param_port)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2673
func proxyholochain_Holochain_Chain(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	res_0 := v.Chain()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2685
func proxyholochain_Holochain_Close(refnum _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	v.Close()
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2692
func proxyholochain_Holochain_DBPath(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	res_0 := v.DBPath()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2701
func proxyholochain_Holochain_DHT(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	res_0 := v.DHT()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2715
func proxyholochain_Holochain_DNAPath(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	res_0 := v.DNAPath()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2724
func proxyholochain_Holochain_Debug(refnum _Ctype_int32_t, param_m _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	_param_m := decodeString(param_m)
	v.Debug(_param_m)
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2742
func proxyholochain_Holochain_GetEntryDef(refnum _Ctype_int32_t, param_t _Ctype_struct_nstring) (_Ctype_int32_t, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	_param_t := decodeString(param_t)
	res_0, res_1 := v.GetEntryDef(_param_t)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2761
func proxyholochain_Holochain_GetProperty(refnum _Ctype_int32_t, param_prop _Ctype_struct_nstring) (_Ctype_struct_nstring, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	_param_prop := decodeString(param_prop)
	res_0, res_1 := v.GetProperty(_param_prop)
	_res_0 := encodeString(res_0)
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2777
func proxyholochain_Holochain_GetZome(refnum _Ctype_int32_t, param_zName _Ctype_struct_nstring) (_Ctype_int32_t, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	_param_zName := decodeString(param_zName)
	res_0, res_1 := v.GetZome(_param_zName)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2794
func proxyholochain_Holochain_GetZomeForEntryType(refnum _Ctype_int32_t, param_t _Ctype_struct_nstring) (_Ctype_int32_t, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	_param_t := decodeString(param_t)
	res_0, res_1 := v.GetZomeForEntryType(_param_t)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2811
func proxyholochain_Holochain_HandleAsyncSends(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	res_0 := v.HandleAsyncSends()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2827
func proxyholochain_Holochain_MakeRibosome(refnum _Ctype_int32_t, param_t _Ctype_struct_nstring) (_Ctype_int32_t, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	_param_t := decodeString(param_t)
	res_0, res_1 := v.MakeRibosome(_param_t)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2844
func proxyholochain_Holochain_Name(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	res_0 := v.Name()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2855
func proxyholochain_Holochain_Node(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	res_0 := v.Node()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2867
func proxyholochain_Holochain_NodeIDStr(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	res_0 := v.NodeIDStr()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2876
func proxyholochain_Holochain_Nucleus(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	res_0 := v.Nucleus()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2888
func proxyholochain_Holochain_Prepare(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	res_0 := v.Prepare()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2900
func proxyholochain_Holochain_PrepareHashType(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	res_0 := v.PrepareHashType()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2914
func proxyholochain_Holochain_Reset(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	res_0 := v.Reset()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2926
func proxyholochain_Holochain_RootPath(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	res_0 := v.RootPath()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2939
func proxyholochain_Holochain_Sign(refnum _Ctype_int32_t, param_doc _Ctype_struct_nbyteslice) (_Ctype_struct_nbyteslice, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	_param_doc := toSlice(param_doc, false)
	res_0, res_1 := v.Sign(_param_doc)
	_res_0 := fromSlice(res_0, true)
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2953
func proxyholochain_Holochain_StartBackgroundTasks(refnum _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	v.StartBackgroundTasks()
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2960
func proxyholochain_Holochain_Started(refnum _Ctype_int32_t) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	res_0 := v.Started()
	var _res_0 _Ctype_char = 0
	if res_0 {
		_res_0 = 1
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2974
func proxyholochain_Holochain_TestPath(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	res_0 := v.TestPath()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:2985
func proxyholochain_Holochain_UIPath(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)
	res_0 := v.UIPath()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3000
func proxyholochain_Holochain_ZomePath(refnum _Ctype_int32_t, param_z _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Holochain)

	var _param_z *holochain.Zome
	if _param_z_ref := _seq.FromRefNum(int32(param_z)); _param_z_ref != nil {
		_param_z = _param_z_ref.Get().(*holochain.Zome)
	}
	res_0 := v.ZomePath(_param_z)
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3014
func new_holochain_Holochain() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Holochain)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3023
func proxyholochain_JSONEntry_Marshal(refnum _Ctype_int32_t) (_Ctype_struct_nbyteslice, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.JSONEntry)
	res_0, res_1 := v.Marshal()
	_res_0 := fromSlice(res_0, true)
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3036
func proxyholochain_JSONEntry_Unmarshal(refnum _Ctype_int32_t, param_b _Ctype_struct_nbyteslice) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.JSONEntry)
	_param_b := toSlice(param_b, false)
	res_0 := v.Unmarshal(_param_b)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3049
func new_holochain_JSONEntry() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.JSONEntry)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3056
func new_holochain_JSONSchemaValidator() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.JSONSchemaValidator)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3065
func proxyholochain_JSRibosome_ChainGenesis(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.JSRibosome)
	res_0 := v.ChainGenesis()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3077
func proxyholochain_JSRibosome_Receive(refnum _Ctype_int32_t, param_from _Ctype_struct_nstring, param_msg _Ctype_struct_nstring) (_Ctype_struct_nstring, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.JSRibosome)
	_param_from := decodeString(param_from)
	_param_msg := decodeString(param_msg)
	res_0, res_1 := v.Receive(_param_from, _param_msg)
	_res_0 := encodeString(res_0)
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3096
func proxyholochain_JSRibosome_Type(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.JSRibosome)
	res_0 := v.Type()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3109
func new_holochain_JSRibosome() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.JSRibosome)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3130
func new_holochain_LibP2PAgent() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.LibP2PAgent)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3135
func proxyholochain_Link_LinkAction_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.Link).LinkAction = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3142
func proxyholochain_Link_LinkAction_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Link).LinkAction
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3150
func proxyholochain_Link_Base_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.Link).Base = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3157
func proxyholochain_Link_Base_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Link).Base
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3165
func proxyholochain_Link_Link_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.Link).Link = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3172
func proxyholochain_Link_Link_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Link).Link
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3180
func proxyholochain_Link_Tag_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.Link).Tag = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3187
func proxyholochain_Link_Tag_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Link).Tag
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3195
func new_holochain_Link() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Link)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3200
func proxyholochain_LinkEvent_Status_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.LinkEvent).Status = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3207
func proxyholochain_LinkEvent_Status_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.LinkEvent).Status
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3215
func proxyholochain_LinkEvent_Source_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.LinkEvent).Source = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3222
func proxyholochain_LinkEvent_Source_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.LinkEvent).Source
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3230
func proxyholochain_LinkEvent_LinksEntry_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.LinkEvent).LinksEntry = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3237
func proxyholochain_LinkEvent_LinksEntry_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.LinkEvent).LinksEntry
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3245
func new_holochain_LinkEvent() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.LinkEvent)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3252
func proxyholochain_LinkQuery_T_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.LinkQuery).T = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3259
func proxyholochain_LinkQuery_T_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.LinkQuery).T
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3267
func proxyholochain_LinkQuery_StatusMask_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.LinkQuery).StatusMask = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3274
func proxyholochain_LinkQuery_StatusMask_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.LinkQuery).StatusMask
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3282
func new_holochain_LinkQuery() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.LinkQuery)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3289
func new_holochain_LinkQueryResp() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.LinkQueryResp)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3298
func new_holochain_LinkReq() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.LinkReq)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3305
func new_holochain_LinksEntry() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.LinksEntry)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3310
func proxyholochain_ListAddReq_ListType_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.ListAddReq).ListType = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3317
func proxyholochain_ListAddReq_ListType_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ListAddReq).ListType
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3327
func proxyholochain_ListAddReq_WarrantType_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.ListAddReq).WarrantType = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3334
func proxyholochain_ListAddReq_WarrantType_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ListAddReq).WarrantType
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3342
func proxyholochain_ListAddReq_Warrant_Set(refnum _Ctype_int32_t, v _Ctype_struct_nbyteslice) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := toSlice(v, true)
	ref.Get().(*holochain.ListAddReq).Warrant = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3349
func proxyholochain_ListAddReq_Warrant_Get(refnum _Ctype_int32_t) _Ctype_struct_nbyteslice {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ListAddReq).Warrant
	_v := fromSlice(v, true)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3357
func new_holochain_ListAddReq() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ListAddReq)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3362
func proxyholochain_Logger_Name_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.Logger).Name = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3369
func proxyholochain_Logger_Name_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Logger).Name
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3377
func proxyholochain_Logger_Enabled_Set(refnum _Ctype_int32_t, v _Ctype_char) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := v != 0
	ref.Get().(*holochain.Logger).Enabled = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3384
func proxyholochain_Logger_Enabled_Get(refnum _Ctype_int32_t) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Logger).Enabled
	var _v _Ctype_char = 0
	if v {
		_v = 1
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3395
func proxyholochain_Logger_Format_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.Logger).Format = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3402
func proxyholochain_Logger_Format_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Logger).Format
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3410
func proxyholochain_Logger_Prefix_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.Logger).Prefix = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3417
func proxyholochain_Logger_Prefix_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Logger).Prefix
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3435
func proxyholochain_Logger_SetPrefix(refnum _Ctype_int32_t, param_prefixFormat _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Logger)
	_param_prefixFormat := decodeString(param_prefixFormat)
	v.SetPrefix(_param_prefixFormat)
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3443
func new_holochain_Logger() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Logger)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3462
func new_holochain_Loggers() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Loggers)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3477
func proxyholochain_Message_Encode(refnum _Ctype_int32_t) (_Ctype_struct_nbyteslice, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Message)
	res_0, res_1 := v.Encode()
	_res_0 := fromSlice(res_0, true)
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3492
func proxyholochain_Message_String(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Message)
	res_0 := v.String()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3501
func new_holochain_Message() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Message)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3508
func proxyholochain_Meta_T_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.Meta).T = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3515
func proxyholochain_Meta_T_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Meta).T
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3523
func proxyholochain_Meta_V_Set(refnum _Ctype_int32_t, v _Ctype_struct_nbyteslice) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := toSlice(v, true)
	ref.Get().(*holochain.Meta).V = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3530
func proxyholochain_Meta_V_Get(refnum _Ctype_int32_t) _Ctype_struct_nbyteslice {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Meta).V
	_v := fromSlice(v, true)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3538
func new_holochain_Meta() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Meta)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3543
func proxyholochain_ModAgentOptions_Identity_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.ModAgentOptions).Identity = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3550
func proxyholochain_ModAgentOptions_Identity_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ModAgentOptions).Identity
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3558
func proxyholochain_ModAgentOptions_Revocation_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.ModAgentOptions).Revocation = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3565
func proxyholochain_ModAgentOptions_Revocation_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ModAgentOptions).Revocation
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3573
func new_holochain_ModAgentOptions() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ModAgentOptions)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3582
func new_holochain_ModReq() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ModReq)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3593
func proxyholochain_Node_Close(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Node)
	res_0 := v.Close()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3629
func proxyholochain_Node_StartProtocol(refnum _Ctype_int32_t, param_h _Ctype_int32_t, param_proto _Ctype_nint) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Node)

	var _param_h *holochain.Holochain
	if _param_h_ref := _seq.FromRefNum(int32(param_h)); _param_h_ref != nil {
		_param_h = _param_h_ref.Get().(*holochain.Holochain)
	}
	_param_proto := int(param_proto)
	res_0 := v.StartProtocol(_param_h, _param_proto)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3649
func new_holochain_Node() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Node)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3654
func proxyholochain_Nucleus_DNA(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Nucleus)
	res_0 := v.DNA()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3666
func proxyholochain_Nucleus_RunGenesis(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Nucleus)
	res_0 := v.RunGenesis()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3678
func proxyholochain_Nucleus_Start(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Nucleus)
	res_0 := v.Start()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3690
func new_holochain_Nucleus() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Nucleus)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3695
func proxyholochain_Package_Chain_Set(refnum _Ctype_int32_t, v _Ctype_struct_nbyteslice) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := toSlice(v, true)
	ref.Get().(*holochain.Package).Chain = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3702
func proxyholochain_Package_Chain_Get(refnum _Ctype_int32_t) _Ctype_struct_nbyteslice {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Package).Chain
	_v := fromSlice(v, true)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3710
func new_holochain_Package() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Package)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3715
func proxyholochain_PeerInfo_ID_Set(refnum _Ctype_int32_t, v _Ctype_struct_nbyteslice) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := toSlice(v, true)
	ref.Get().(*holochain.PeerInfo).ID = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3722
func proxyholochain_PeerInfo_ID_Get(refnum _Ctype_int32_t) _Ctype_struct_nbyteslice {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.PeerInfo).ID
	_v := fromSlice(v, true)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3732
func new_holochain_PeerInfo() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.PeerInfo)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3741
func new_holochain_PeerList() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.PeerList)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3748
func proxyholochain_PeerRecord_Warrant_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.PeerRecord).Warrant = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3755
func proxyholochain_PeerRecord_Warrant_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.PeerRecord).Warrant
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3763
func new_holochain_PeerRecord() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.PeerRecord)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3768
func proxyholochain_Progenitor_Identity_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.Progenitor).Identity = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3775
func proxyholochain_Progenitor_Identity_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Progenitor).Identity
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3783
func proxyholochain_Progenitor_PubKey_Set(refnum _Ctype_int32_t, v _Ctype_struct_nbyteslice) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := toSlice(v, true)
	ref.Get().(*holochain.Progenitor).PubKey = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3790
func proxyholochain_Progenitor_PubKey_Get(refnum _Ctype_int32_t) _Ctype_struct_nbyteslice {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Progenitor).PubKey
	_v := fromSlice(v, true)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3798
func new_holochain_Progenitor() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Progenitor)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3807
func new_holochain_Protocol() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Protocol)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3812
func proxyholochain_Put_Idx_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.Put).Idx = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3819
func proxyholochain_Put_Idx_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Put).Idx
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3829
func new_holochain_Put() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Put)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3836
func proxyholochain_PutReq_S_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.PutReq).S = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3843
func proxyholochain_PutReq_S_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.PutReq).S
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3853
func new_holochain_PutReq() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.PutReq)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3860
func proxyholochain_QueryConstrain_Contains_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.QueryConstrain).Contains = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3867
func proxyholochain_QueryConstrain_Contains_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.QueryConstrain).Contains
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3875
func proxyholochain_QueryConstrain_Equals_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.QueryConstrain).Equals = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3882
func proxyholochain_QueryConstrain_Equals_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.QueryConstrain).Equals
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3890
func proxyholochain_QueryConstrain_Matches_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.QueryConstrain).Matches = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3897
func proxyholochain_QueryConstrain_Matches_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.QueryConstrain).Matches
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3905
func proxyholochain_QueryConstrain_Count_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.QueryConstrain).Count = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3912
func proxyholochain_QueryConstrain_Count_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.QueryConstrain).Count
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3920
func proxyholochain_QueryConstrain_Page_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.QueryConstrain).Page = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3927
func proxyholochain_QueryConstrain_Page_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.QueryConstrain).Page
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3935
func new_holochain_QueryConstrain() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.QueryConstrain)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3946
func new_holochain_QueryOptions() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.QueryOptions)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3951
func proxyholochain_QueryOrder_Ascending_Set(refnum _Ctype_int32_t, v _Ctype_char) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := v != 0
	ref.Get().(*holochain.QueryOrder).Ascending = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3958
func proxyholochain_QueryOrder_Ascending_Get(refnum _Ctype_int32_t) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.QueryOrder).Ascending
	var _v _Ctype_char = 0
	if v {
		_v = 1
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3969
func new_holochain_QueryOrder() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.QueryOrder)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3974
func proxyholochain_QueryResult_Header_Set(refnum _Ctype_int32_t, v _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))

	var _v *holochain.Header
	if _v_ref := _seq.FromRefNum(int32(v)); _v_ref != nil {
		_v = _v_ref.Get().(*holochain.Header)
	}
	ref.Get().(*holochain.QueryResult).Header = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3985
func proxyholochain_QueryResult_Header_Get(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.QueryResult).Header
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:3996
func proxyholochain_QueryResult_Entry_Set(refnum _Ctype_int32_t, v _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	var _v holochain.Entry
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(holochain.Entry)
		} else {
			_v = (*proxyholochain_Entry)(_v_ref)
		}
	}
	ref.Get().(*holochain.QueryResult).Entry = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4011
func proxyholochain_QueryResult_Entry_Get(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.QueryResult).Entry
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4022
func new_holochain_QueryResult() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.QueryResult)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4027
func proxyholochain_QueryReturn_Hashes_Set(refnum _Ctype_int32_t, v _Ctype_char) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := v != 0
	ref.Get().(*holochain.QueryReturn).Hashes = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4034
func proxyholochain_QueryReturn_Hashes_Get(refnum _Ctype_int32_t) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.QueryReturn).Hashes
	var _v _Ctype_char = 0
	if v {
		_v = 1
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4045
func proxyholochain_QueryReturn_Entries_Set(refnum _Ctype_int32_t, v _Ctype_char) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := v != 0
	ref.Get().(*holochain.QueryReturn).Entries = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4052
func proxyholochain_QueryReturn_Entries_Get(refnum _Ctype_int32_t) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.QueryReturn).Entries
	var _v _Ctype_char = 0
	if v {
		_v = 1
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4063
func proxyholochain_QueryReturn_Headers_Set(refnum _Ctype_int32_t, v _Ctype_char) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := v != 0
	ref.Get().(*holochain.QueryReturn).Headers = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4070
func proxyholochain_QueryReturn_Headers_Get(refnum _Ctype_int32_t) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.QueryReturn).Headers
	var _v _Ctype_char = 0
	if v {
		_v = 1
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4081
func new_holochain_QueryReturn() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.QueryReturn)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4094
func proxyholochain_RoutingTable_IsEmpty(refnum _Ctype_int32_t) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.RoutingTable)
	res_0 := v.IsEmpty()
	var _res_0 _Ctype_char = 0
	if res_0 {
		_res_0 = 1
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4112
func proxyholochain_RoutingTable_Print(refnum _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.RoutingTable)
	v.Print()
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4121
func proxyholochain_RoutingTable_Size(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.RoutingTable)
	res_0 := v.Size()
	_res_0 := _Ctype_nint(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4132
func new_holochain_RoutingTable() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.RoutingTable)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4137
func proxyholochain_SelfRevocation_Data_Set(refnum _Ctype_int32_t, v _Ctype_struct_nbyteslice) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := toSlice(v, true)
	ref.Get().(*holochain.SelfRevocation).Data = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4144
func proxyholochain_SelfRevocation_Data_Get(refnum _Ctype_int32_t) _Ctype_struct_nbyteslice {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.SelfRevocation).Data
	_v := fromSlice(v, true)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4152
func proxyholochain_SelfRevocation_OldSig_Set(refnum _Ctype_int32_t, v _Ctype_struct_nbyteslice) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := toSlice(v, true)
	ref.Get().(*holochain.SelfRevocation).OldSig = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4159
func proxyholochain_SelfRevocation_OldSig_Get(refnum _Ctype_int32_t) _Ctype_struct_nbyteslice {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.SelfRevocation).OldSig
	_v := fromSlice(v, true)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4167
func proxyholochain_SelfRevocation_NewSig_Set(refnum _Ctype_int32_t, v _Ctype_struct_nbyteslice) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := toSlice(v, true)
	ref.Get().(*holochain.SelfRevocation).NewSig = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4174
func proxyholochain_SelfRevocation_NewSig_Get(refnum _Ctype_int32_t) _Ctype_struct_nbyteslice {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.SelfRevocation).NewSig
	_v := fromSlice(v, true)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4182
func proxyholochain_SelfRevocation_Marshal(refnum _Ctype_int32_t) (_Ctype_struct_nbyteslice, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.SelfRevocation)
	res_0, res_1 := v.Marshal()
	_res_0 := fromSlice(res_0, true)
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4195
func proxyholochain_SelfRevocation_Unmarshal(refnum _Ctype_int32_t, param_data _Ctype_struct_nbyteslice) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.SelfRevocation)
	_param_data := toSlice(param_data, false)
	res_0 := v.Unmarshal(_param_data)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4208
func proxyholochain_SelfRevocation_Verify(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.SelfRevocation)
	res_0 := v.Verify()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4220
func new_holochain_SelfRevocation() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.SelfRevocation)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4227
func proxyholochain_SelfRevocationWarrant_Decode(refnum _Ctype_int32_t, param_data _Ctype_struct_nbyteslice) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.SelfRevocationWarrant)
	_param_data := toSlice(param_data, false)
	res_0 := v.Decode(_param_data)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4240
func proxyholochain_SelfRevocationWarrant_Encode(refnum _Ctype_int32_t) (_Ctype_struct_nbyteslice, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.SelfRevocationWarrant)
	res_0, res_1 := v.Encode()
	_res_0 := fromSlice(res_0, true)
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4257
func proxyholochain_SelfRevocationWarrant_Type(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.SelfRevocationWarrant)
	res_0 := v.Type()
	_res_0 := _Ctype_nint(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4266
func proxyholochain_SelfRevocationWarrant_Verify(refnum _Ctype_int32_t, param_h _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.SelfRevocationWarrant)

	var _param_h *holochain.Holochain
	if _param_h_ref := _seq.FromRefNum(int32(param_h)); _param_h_ref != nil {
		_param_h = _param_h_ref.Get().(*holochain.Holochain)
	}
	res_0 := v.Verify(_param_h)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4283
func new_holochain_SelfRevocationWarrant() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.SelfRevocationWarrant)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4288
func proxyholochain_SendOptions_Callback_Set(refnum _Ctype_int32_t, v _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))

	var _v *holochain.Callback
	if _v_ref := _seq.FromRefNum(int32(v)); _v_ref != nil {
		_v = _v_ref.Get().(*holochain.Callback)
	}
	ref.Get().(*holochain.SendOptions).Callback = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4299
func proxyholochain_SendOptions_Callback_Get(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.SendOptions).Callback
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4310
func proxyholochain_SendOptions_Timeout_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.SendOptions).Timeout = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4317
func proxyholochain_SendOptions_Timeout_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.SendOptions).Timeout
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4325
func new_holochain_SendOptions() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.SendOptions)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4332
func proxyholochain_Service_DefaultAgent_Set(refnum _Ctype_int32_t, v _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	var _v holochain.Agent
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(holochain.Agent)
		} else {
			_v = (*proxyholochain_Agent)(_v_ref)
		}
	}
	ref.Get().(*holochain.Service).DefaultAgent = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4347
func proxyholochain_Service_DefaultAgent_Get(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Service).DefaultAgent
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4358
func proxyholochain_Service_Path_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.Service).Path = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4365
func proxyholochain_Service_Path_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Service).Path
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4373
func proxyholochain_Service_Clone(refnum _Ctype_int32_t, param_srcPath _Ctype_struct_nstring, param_root _Ctype_struct_nstring, param_agent _Ctype_int32_t, param_new _Ctype_char, param_initDB _Ctype_char) (_Ctype_int32_t, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Service)
	_param_srcPath := decodeString(param_srcPath)
	_param_root := decodeString(param_root)
	var _param_agent holochain.Agent
	_param_agent_ref := _seq.FromRefNum(int32(param_agent))
	if _param_agent_ref != nil {
		if param_agent < 0 {
			_param_agent = _param_agent_ref.Get().(holochain.Agent)
		} else {
			_param_agent = (*proxyholochain_Agent)(_param_agent_ref)
		}
	}
	_param_new := param_new != 0
	_param_initDB := param_initDB != 0
	res_0, res_1 := v.Clone(_param_srcPath, _param_root, _param_agent, _param_new, _param_initDB)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4404
func proxyholochain_Service_GenChain(refnum _Ctype_int32_t, param_name _Ctype_struct_nstring) (_Ctype_int32_t, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Service)
	_param_name := decodeString(param_name)
	res_0, res_1 := v.GenChain(_param_name)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4421
func proxyholochain_Service_InitAppDir(refnum _Ctype_int32_t, param_root _Ctype_struct_nstring, param_encodingFormat _Ctype_struct_nstring) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Service)
	_param_root := decodeString(param_root)
	_param_encodingFormat := decodeString(param_encodingFormat)
	res_0 := v.InitAppDir(_param_root, _param_encodingFormat)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4435
func proxyholochain_Service_IsConfigured(refnum _Ctype_int32_t, param_name _Ctype_struct_nstring) (_Ctype_struct_nstring, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Service)
	_param_name := decodeString(param_name)
	res_0, res_1 := v.IsConfigured(_param_name)
	_res_0 := encodeString(res_0)
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4449
func proxyholochain_Service_ListChains(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Service)
	res_0 := v.ListChains()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4458
func proxyholochain_Service_Load(refnum _Ctype_int32_t, param_name _Ctype_struct_nstring) (_Ctype_int32_t, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Service)
	_param_name := decodeString(param_name)
	res_0, res_1 := v.Load(_param_name)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4475
func proxyholochain_Service_MakeAppPackage(refnum _Ctype_int32_t, param_h _Ctype_int32_t) (_Ctype_struct_nbyteslice, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Service)

	var _param_h *holochain.Holochain
	if _param_h_ref := _seq.FromRefNum(int32(param_h)); _param_h_ref != nil {
		_param_h = _param_h_ref.Get().(*holochain.Holochain)
	}
	res_0, res_1 := v.MakeAppPackage(_param_h)
	_res_0 := fromSlice(res_0, true)
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4493
func proxyholochain_Service_MakeTestingApp(refnum _Ctype_int32_t, param_root _Ctype_struct_nstring, param_encodingFormat _Ctype_struct_nstring, param_initDB _Ctype_char, param_newUUID _Ctype_char, param_agent _Ctype_int32_t) (_Ctype_int32_t, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Service)
	_param_root := decodeString(param_root)
	_param_encodingFormat := decodeString(param_encodingFormat)
	_param_initDB := param_initDB != 0
	_param_newUUID := param_newUUID != 0
	var _param_agent holochain.Agent
	_param_agent_ref := _seq.FromRefNum(int32(param_agent))
	if _param_agent_ref != nil {
		if param_agent < 0 {
			_param_agent = _param_agent_ref.Get().(holochain.Agent)
		} else {
			_param_agent = (*proxyholochain_Agent)(_param_agent_ref)
		}
	}
	res_0, res_1 := v.MakeTestingApp(_param_root, _param_encodingFormat, _param_initDB, _param_newUUID, _param_agent)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4524
func new_holochain_Service() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Service)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4529
func proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Set(refnum _Ctype_int32_t, v _Ctype_char) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := v != 0
	ref.Get().(*holochain.ServiceConfig).DefaultPeerModeAuthor = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4536
func proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Get(refnum _Ctype_int32_t) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ServiceConfig).DefaultPeerModeAuthor
	var _v _Ctype_char = 0
	if v {
		_v = 1
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4547
func proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Set(refnum _Ctype_int32_t, v _Ctype_char) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := v != 0
	ref.Get().(*holochain.ServiceConfig).DefaultPeerModeDHTNode = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4554
func proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Get(refnum _Ctype_int32_t) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ServiceConfig).DefaultPeerModeDHTNode
	var _v _Ctype_char = 0
	if v {
		_v = 1
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4565
func proxyholochain_ServiceConfig_DefaultBootstrapServer_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.ServiceConfig).DefaultBootstrapServer = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4572
func proxyholochain_ServiceConfig_DefaultBootstrapServer_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ServiceConfig).DefaultBootstrapServer
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4580
func proxyholochain_ServiceConfig_DefaultEnableMDNS_Set(refnum _Ctype_int32_t, v _Ctype_char) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := v != 0
	ref.Get().(*holochain.ServiceConfig).DefaultEnableMDNS = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4587
func proxyholochain_ServiceConfig_DefaultEnableMDNS_Get(refnum _Ctype_int32_t) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ServiceConfig).DefaultEnableMDNS
	var _v _Ctype_char = 0
	if v {
		_v = 1
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4598
func proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Set(refnum _Ctype_int32_t, v _Ctype_char) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := v != 0
	ref.Get().(*holochain.ServiceConfig).DefaultEnableNATUPnP = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4605
func proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Get(refnum _Ctype_int32_t) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ServiceConfig).DefaultEnableNATUPnP
	var _v _Ctype_char = 0
	if v {
		_v = 1
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4616
func proxyholochain_ServiceConfig_Validate(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ServiceConfig)
	res_0 := v.Validate()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4628
func new_holochain_ServiceConfig() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ServiceConfig)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4633
func proxyholochain_Signature_S_Set(refnum _Ctype_int32_t, v _Ctype_struct_nbyteslice) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := toSlice(v, true)
	ref.Get().(*holochain.Signature).S = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4640
func proxyholochain_Signature_S_Get(refnum _Ctype_int32_t) _Ctype_struct_nbyteslice {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Signature).S
	_v := fromSlice(v, true)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4648
func new_holochain_Signature() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Signature)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4653
func proxyholochain_StatusChange_Action_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.StatusChange).Action = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4660
func proxyholochain_StatusChange_Action_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.StatusChange).Action
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4670
func new_holochain_StatusChange() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.StatusChange)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4675
func proxyholochain_TaggedHash_H_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.TaggedHash).H = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4682
func proxyholochain_TaggedHash_H_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.TaggedHash).H
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4690
func proxyholochain_TaggedHash_E_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.TaggedHash).E = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4697
func proxyholochain_TaggedHash_E_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.TaggedHash).E
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4705
func proxyholochain_TaggedHash_EntryType_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.TaggedHash).EntryType = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4712
func proxyholochain_TaggedHash_EntryType_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.TaggedHash).EntryType
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4720
func proxyholochain_TaggedHash_T_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.TaggedHash).T = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4727
func proxyholochain_TaggedHash_T_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.TaggedHash).T
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4735
func proxyholochain_TaggedHash_Source_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.TaggedHash).Source = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4742
func proxyholochain_TaggedHash_Source_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.TaggedHash).Source
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4750
func new_holochain_TaggedHash() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.TaggedHash)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4755
func proxyholochain_TestConfig_GossipInterval_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.TestConfig).GossipInterval = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4762
func proxyholochain_TestConfig_GossipInterval_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.TestConfig).GossipInterval
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4770
func proxyholochain_TestConfig_Duration_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.TestConfig).Duration = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4777
func proxyholochain_TestConfig_Duration_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.TestConfig).Duration
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4787
func new_holochain_TestConfig() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.TestConfig)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4792
func proxyholochain_TestData_Convey_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.TestData).Convey = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4799
func proxyholochain_TestData_Convey_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.TestData).Convey
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4807
func proxyholochain_TestData_Zome_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.TestData).Zome = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4814
func proxyholochain_TestData_Zome_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.TestData).Zome
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4822
func proxyholochain_TestData_FnName_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.TestData).FnName = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4829
func proxyholochain_TestData_FnName_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.TestData).FnName
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4843
func proxyholochain_TestData_Regexp_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.TestData).Regexp = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4850
func proxyholochain_TestData_Regexp_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.TestData).Regexp
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4862
func proxyholochain_TestData_Exposure_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.TestData).Exposure = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4869
func proxyholochain_TestData_Exposure_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.TestData).Exposure
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4877
func proxyholochain_TestData_Raw_Set(refnum _Ctype_int32_t, v _Ctype_char) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := v != 0
	ref.Get().(*holochain.TestData).Raw = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4884
func proxyholochain_TestData_Raw_Get(refnum _Ctype_int32_t) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.TestData).Raw
	var _v _Ctype_char = 0
	if v {
		_v = 1
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4895
func proxyholochain_TestData_Repeat_Set(refnum _Ctype_int32_t, v _Ctype_nint) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := int(v)
	ref.Get().(*holochain.TestData).Repeat = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4902
func proxyholochain_TestData_Repeat_Get(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.TestData).Repeat
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4910
func proxyholochain_TestData_Benchmark_Set(refnum _Ctype_int32_t, v _Ctype_char) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := v != 0
	ref.Get().(*holochain.TestData).Benchmark = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4917
func proxyholochain_TestData_Benchmark_Get(refnum _Ctype_int32_t) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.TestData).Benchmark
	var _v _Ctype_char = 0
	if v {
		_v = 1
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4928
func new_holochain_TestData() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.TestData)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4935
func new_holochain_TestFixtures() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.TestFixtures)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4942
func proxyholochain_TestSet_Identity_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.TestSet).Identity = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4949
func proxyholochain_TestSet_Identity_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.TestSet).Identity
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4959
func proxyholochain_TestSet_Benchmark_Set(refnum _Ctype_int32_t, v _Ctype_char) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := v != 0
	ref.Get().(*holochain.TestSet).Benchmark = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4966
func proxyholochain_TestSet_Benchmark_Get(refnum _Ctype_int32_t) _Ctype_char {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.TestSet).Benchmark
	var _v _Ctype_char = 0
	if v {
		_v = 1
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4977
func new_holochain_TestSet() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.TestSet)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4984
func new_holochain_ValidateQuery() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ValidateQuery)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4989
func proxyholochain_ValidateResponse_Type_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.ValidateResponse).Type = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:4996
func proxyholochain_ValidateResponse_Type_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ValidateResponse).Type
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5010
func new_holochain_ValidateResponse() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ValidateResponse)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5015
func proxyholochain_ValidationPackage_Chain_Set(refnum _Ctype_int32_t, v _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))

	var _v *holochain.Chain
	if _v_ref := _seq.FromRefNum(int32(v)); _v_ref != nil {
		_v = _v_ref.Get().(*holochain.Chain)
	}
	ref.Get().(*holochain.ValidationPackage).Chain = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5026
func proxyholochain_ValidationPackage_Chain_Get(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ValidationPackage).Chain
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5037
func new_holochain_ValidationPackage() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ValidationPackage)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5042
func proxyholochain_Zome_Name_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.Zome).Name = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5049
func proxyholochain_Zome_Name_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Zome).Name
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5057
func proxyholochain_Zome_Description_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.Zome).Description = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5064
func proxyholochain_Zome_Description_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Zome).Description
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5072
func proxyholochain_Zome_Code_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.Zome).Code = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5079
func proxyholochain_Zome_Code_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Zome).Code
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5089
func proxyholochain_Zome_RibosomeType_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.Zome).RibosomeType = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5096
func proxyholochain_Zome_RibosomeType_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Zome).RibosomeType
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5112
func proxyholochain_Zome_CodeFileName(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Zome)
	res_0 := v.CodeFileName()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5121
func proxyholochain_Zome_GetEntryDef(refnum _Ctype_int32_t, param_entryName _Ctype_struct_nstring) (_Ctype_int32_t, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Zome)
	_param_entryName := decodeString(param_entryName)
	res_0, res_1 := v.GetEntryDef(_param_entryName)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5138
func proxyholochain_Zome_GetFunctionDef(refnum _Ctype_int32_t, param_fnName _Ctype_struct_nstring) (_Ctype_int32_t, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Zome)
	_param_fnName := decodeString(param_fnName)
	res_0, res_1 := v.GetFunctionDef(_param_fnName)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5157
func proxyholochain_Zome_MakeRibosome(refnum _Ctype_int32_t, param_h _Ctype_int32_t) (_Ctype_int32_t, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.Zome)

	var _param_h *holochain.Holochain
	if _param_h_ref := _seq.FromRefNum(int32(param_h)); _param_h_ref != nil {
		_param_h = _param_h_ref.Get().(*holochain.Holochain)
	}
	res_0, res_1 := v.MakeRibosome(_param_h)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5178
func new_holochain_Zome() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.Zome)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5183
func proxyholochain_ZomeFile_Name_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.ZomeFile).Name = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5190
func proxyholochain_ZomeFile_Name_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ZomeFile).Name
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5198
func proxyholochain_ZomeFile_Description_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.ZomeFile).Description = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5205
func proxyholochain_ZomeFile_Description_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ZomeFile).Description
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5213
func proxyholochain_ZomeFile_CodeFile_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.ZomeFile).CodeFile = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5220
func proxyholochain_ZomeFile_CodeFile_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ZomeFile).CodeFile
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5228
func proxyholochain_ZomeFile_RibosomeType_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.ZomeFile).RibosomeType = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5235
func proxyholochain_ZomeFile_RibosomeType_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ZomeFile).RibosomeType
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5245
func proxyholochain_ZomeFile_BridgeTo_Set(refnum _Ctype_int32_t, v _Ctype_struct_nstring) {
	ref := _seq.FromRefNum(int32(refnum))
	_v := decodeString(v)
	ref.Get().(*holochain.ZomeFile).BridgeTo = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5252
func proxyholochain_ZomeFile_BridgeTo_Get(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ZomeFile).BridgeTo
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5266
func new_holochain_ZomeFile() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ZomeFile)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5275
func proxyholochain_ZygoRibosome_ChainGenesis(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ZygoRibosome)
	res_0 := v.ChainGenesis()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5287
func proxyholochain_ZygoRibosome_Receive(refnum _Ctype_int32_t, param_from _Ctype_struct_nstring, param_msg _Ctype_struct_nstring) (_Ctype_struct_nstring, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ZygoRibosome)
	_param_from := decodeString(param_from)
	_param_msg := decodeString(param_msg)
	res_0, res_1 := v.Receive(_param_from, _param_msg)
	_res_0 := encodeString(res_0)
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5306
func proxyholochain_ZygoRibosome_Type(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(*holochain.ZygoRibosome)
	res_0 := v.Type()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5319
func new_holochain_ZygoRibosome() _Ctype_int32_t {
	return _Ctype_int32_t(_seq.ToRefNum(new(holochain.ZygoRibosome)))
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5328
func proxyholochain_Action_Name(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(holochain.Action)
	res_0 := v.Name()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5338
type proxyholochain_Action _seq.Ref

func (p *proxyholochain_Action) Bind_proxy_refnum__() int32	{ return (*_seq.Ref)(p).Bind_IncNum() }

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5344
func (p *proxyholochain_Action) Name() string {
	res := _Cfunc_cproxyholochain_Action_Name(_Ctype_int32_t(p.Bind_proxy_refnum__()))
	_res := decodeString(res)
	return _res
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5367
type proxyholochain_Agent _seq.Ref

func (p *proxyholochain_Agent) Bind_proxy_refnum__() int32	{ return (*_seq.Ref)(p).Bind_IncNum() }

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5383
type proxyholochain_ArgsAction _seq.Ref

func (p *proxyholochain_ArgsAction) Bind_proxy_refnum__() int32	{ return (*_seq.Ref)(p).Bind_IncNum() }

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5392
func proxyholochain_CommittingAction_CheckValidationRequest(refnum _Ctype_int32_t, param_def _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(holochain.CommittingAction)

	var _param_def *holochain.EntryDef
	if _param_def_ref := _seq.FromRefNum(int32(param_def)); _param_def_ref != nil {
		_param_def = _param_def_ref.Get().(*holochain.EntryDef)
	}
	res_0 := v.CheckValidationRequest(_param_def)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5411
func proxyholochain_CommittingAction_Entry(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(holochain.CommittingAction)
	res_0 := v.Entry()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5423
func proxyholochain_CommittingAction_EntryType(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(holochain.CommittingAction)
	res_0 := v.EntryType()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5432
func proxyholochain_CommittingAction_Name(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(holochain.CommittingAction)
	res_0 := v.Name()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5443
func proxyholochain_CommittingAction_SetHeader(refnum _Ctype_int32_t, param_header _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(holochain.CommittingAction)

	var _param_header *holochain.Header
	if _param_header_ref := _seq.FromRefNum(int32(param_header)); _param_header_ref != nil {
		_param_header = _param_header_ref.Get().(*holochain.Header)
	}
	v.SetHeader(_param_header)
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5456
type proxyholochain_CommittingAction _seq.Ref

func (p *proxyholochain_CommittingAction) Bind_proxy_refnum__() int32 {
	return (*_seq.Ref)(p).Bind_IncNum()
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5463
func (p *proxyholochain_CommittingAction) CheckValidationRequest(param_def *holochain.EntryDef) error {
	var _param_def _Ctype_int32_t = _seq.NullRefNum
	if param_def != nil {
		_param_def = _Ctype_int32_t(_seq.ToRefNum(param_def))
	}
	res := _Cfunc_cproxyholochain_CommittingAction_CheckValidationRequest(_Ctype_int32_t(p.Bind_proxy_refnum__()), _param_def)
	var _res error
	_res_ref := _seq.FromRefNum(int32(res))
	if _res_ref != nil {
		if res < 0 {
			_res = _res_ref.Get().(error)
		} else {
			_res = (*proxy_error)(_res_ref)
		}
	}
	return _res
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5482
func (p *proxyholochain_CommittingAction) Entry() holochain.Entry {
	res := _Cfunc_cproxyholochain_CommittingAction_Entry(_Ctype_int32_t(p.Bind_proxy_refnum__()))
	var _res holochain.Entry
	_res_ref := _seq.FromRefNum(int32(res))
	if _res_ref != nil {
		if res < 0 {
			_res = _res_ref.Get().(holochain.Entry)
		} else {
			_res = (*proxyholochain_Entry)(_res_ref)
		}
	}
	return _res
}

func (p *proxyholochain_CommittingAction) EntryType() string {
	res := _Cfunc_cproxyholochain_CommittingAction_EntryType(_Ctype_int32_t(p.Bind_proxy_refnum__()))
	_res := decodeString(res)
	return _res
}

func (p *proxyholochain_CommittingAction) Name() string {
	res := _Cfunc_cproxyholochain_CommittingAction_Name(_Ctype_int32_t(p.Bind_proxy_refnum__()))
	_res := decodeString(res)
	return _res
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5509
func (p *proxyholochain_CommittingAction) SetHeader(param_header *holochain.Header) {
	var _param_header _Ctype_int32_t = _seq.NullRefNum
	if param_header != nil {
		_param_header = _Ctype_int32_t(_seq.ToRefNum(param_header))
	}
	_Cfunc_cproxyholochain_CommittingAction_SetHeader(_Ctype_int32_t(p.Bind_proxy_refnum__()), _param_header)
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5521
func proxyholochain_Entry_Marshal(refnum _Ctype_int32_t) (_Ctype_struct_nbyteslice, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(holochain.Entry)
	res_0, res_1 := v.Marshal()
	_res_0 := fromSlice(res_0, true)
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5536
func proxyholochain_Entry_Unmarshal(refnum _Ctype_int32_t, param_p0 _Ctype_struct_nbyteslice) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(holochain.Entry)
	_param_p0 := toSlice(param_p0, false)
	res_0 := v.Unmarshal(_param_p0)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

type proxyholochain_Entry _seq.Ref

func (p *proxyholochain_Entry) Bind_proxy_refnum__() int32	{ return (*_seq.Ref)(p).Bind_IncNum() }

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5553
func (p *proxyholochain_Entry) Marshal() ([]byte, error) {
	res := _Cfunc_cproxyholochain_Entry_Marshal(_Ctype_int32_t(p.Bind_proxy_refnum__()))
	res_0 := toSlice(res.r0, true)
	var res_1 error
	res_1_ref := _seq.FromRefNum(int32(res.r1))
	if res_1_ref != nil {
		if res.r1 < 0 {
			res_1 = res_1_ref.Get().(error)
		} else {
			res_1 = (*proxy_error)(res_1_ref)
		}
	}
	return res_0, res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5569
func (p *proxyholochain_Entry) Unmarshal(param_p0 []byte) error {
	_param_p0 := fromSlice(param_p0, false)
	res := func(_cgo0 _Ctype_int32_t, _cgo1 _Ctype_struct_nbyteslice) _Ctype_int32_t {
//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5571
		_cgoCheckPointer(_cgo1)
//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5571
		return _Cfunc_cproxyholochain_Entry_Unmarshal(_cgo0, _cgo1)
//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5571
	}(_Ctype_int32_t(p.Bind_proxy_refnum__()), _param_p0)
															var _res error
															_res_ref := _seq.FromRefNum(int32(res))
															if _res_ref != nil {
		if res < 0 {
			_res = _res_ref.Get().(error)
		} else {
			_res = (*proxy_error)(_res_ref)
		}
	}
	return _res
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5585
func proxyholochain_Revocation_Marshal(refnum _Ctype_int32_t) (_Ctype_struct_nbyteslice, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(holochain.Revocation)
	res_0, res_1 := v.Marshal()
	_res_0 := fromSlice(res_0, true)
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5598
func proxyholochain_Revocation_Unmarshal(refnum _Ctype_int32_t, param_p0 _Ctype_struct_nbyteslice) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(holochain.Revocation)
	_param_p0 := toSlice(param_p0, false)
	res_0 := v.Unmarshal(_param_p0)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5611
func proxyholochain_Revocation_Verify(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(holochain.Revocation)
	res_0 := v.Verify()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

type proxyholochain_Revocation _seq.Ref

func (p *proxyholochain_Revocation) Bind_proxy_refnum__() int32	{ return (*_seq.Ref)(p).Bind_IncNum() }

func (p *proxyholochain_Revocation) Marshal() ([]byte, error) {
	res := _Cfunc_cproxyholochain_Revocation_Marshal(_Ctype_int32_t(p.Bind_proxy_refnum__()))
	res_0 := toSlice(res.r0, true)
	var res_1 error
	res_1_ref := _seq.FromRefNum(int32(res.r1))
	if res_1_ref != nil {
		if res.r1 < 0 {
			res_1 = res_1_ref.Get().(error)
		} else {
			res_1 = (*proxy_error)(res_1_ref)
		}
	}
	return res_0, res_1
}

func (p *proxyholochain_Revocation) Unmarshal(param_p0 []byte) error {
	_param_p0 := fromSlice(param_p0, false)
	res := func(_cgo0 _Ctype_int32_t, _cgo1 _Ctype_struct_nbyteslice) _Ctype_int32_t {
//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5643
		_cgoCheckPointer(_cgo1)
//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5643
		return _Cfunc_cproxyholochain_Revocation_Unmarshal(_cgo0, _cgo1)
//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5643
	}(_Ctype_int32_t(p.Bind_proxy_refnum__()), _param_p0)
															var _res error
															_res_ref := _seq.FromRefNum(int32(res))
															if _res_ref != nil {
		if res < 0 {
			_res = _res_ref.Get().(error)
		} else {
			_res = (*proxy_error)(_res_ref)
		}
	}
	return _res
}

func (p *proxyholochain_Revocation) Verify() error {
	res := _Cfunc_cproxyholochain_Revocation_Verify(_Ctype_int32_t(p.Bind_proxy_refnum__()))
	var _res error
	_res_ref := _seq.FromRefNum(int32(res))
	if _res_ref != nil {
		if res < 0 {
			_res = _res_ref.Get().(error)
		} else {
			_res = (*proxy_error)(_res_ref)
		}
	}
	return _res
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5675
func proxyholochain_Ribosome_ChainGenesis(refnum _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(holochain.Ribosome)
	res_0 := v.ChainGenesis()
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5687
func proxyholochain_Ribosome_Receive(refnum _Ctype_int32_t, param_from _Ctype_struct_nstring, param_msg _Ctype_struct_nstring) (_Ctype_struct_nstring, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(holochain.Ribosome)
	_param_from := decodeString(param_from)
	_param_msg := decodeString(param_msg)
	res_0, res_1 := v.Receive(_param_from, _param_msg)
	_res_0 := encodeString(res_0)
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5706
func proxyholochain_Ribosome_Type(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(holochain.Ribosome)
	res_0 := v.Type()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5718
type proxyholochain_Ribosome _seq.Ref

func (p *proxyholochain_Ribosome) Bind_proxy_refnum__() int32	{ return (*_seq.Ref)(p).Bind_IncNum() }

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5724
func (p *proxyholochain_Ribosome) ChainGenesis() error {
	res := _Cfunc_cproxyholochain_Ribosome_ChainGenesis(_Ctype_int32_t(p.Bind_proxy_refnum__()))
	var _res error
	_res_ref := _seq.FromRefNum(int32(res))
	if _res_ref != nil {
		if res < 0 {
			_res = _res_ref.Get().(error)
		} else {
			_res = (*proxy_error)(_res_ref)
		}
	}
	return _res
}

func (p *proxyholochain_Ribosome) Receive(param_from string, param_msg string) (string, error) {
	_param_from := encodeString(param_from)
	_param_msg := encodeString(param_msg)
	res := func(_cgo0 _Ctype_int32_t, _cgo1 _Ctype_struct_nstring, _cgo2 _Ctype_struct_nstring) _Ctype_struct_cproxyholochain_Ribosome_Receive_return {
//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5741
		_cgoCheckPointer(_cgo1)
//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5741
		_cgoCheckPointer(_cgo2)
//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5741
		return _Cfunc_cproxyholochain_Ribosome_Receive(_cgo0, _cgo1, _cgo2)
//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5741
	}(_Ctype_int32_t(p.Bind_proxy_refnum__()), _param_from, _param_msg)
															res_0 := decodeString(res.r0)
															var res_1 error
															res_1_ref := _seq.FromRefNum(int32(res.r1))
															if res_1_ref != nil {
		if res.r1 < 0 {
			res_1 = res_1_ref.Get().(error)
		} else {
			res_1 = (*proxy_error)(res_1_ref)
		}
	}
	return res_0, res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5757
func (p *proxyholochain_Ribosome) Type() string {
	res := _Cfunc_cproxyholochain_Ribosome_Type(_Ctype_int32_t(p.Bind_proxy_refnum__()))
	_res := decodeString(res)
	return _res
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5767
type proxyholochain_SchemaValidator _seq.Ref

func (p *proxyholochain_SchemaValidator) Bind_proxy_refnum__() int32 {
	return (*_seq.Ref)(p).Bind_IncNum()
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5777
func proxyholochain_ValidatingAction_CheckValidationRequest(refnum _Ctype_int32_t, param_def _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(holochain.ValidatingAction)

	var _param_def *holochain.EntryDef
	if _param_def_ref := _seq.FromRefNum(int32(param_def)); _param_def_ref != nil {
		_param_def = _param_def_ref.Get().(*holochain.EntryDef)
	}
	res_0 := v.CheckValidationRequest(_param_def)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5796
func proxyholochain_ValidatingAction_Name(refnum _Ctype_int32_t) _Ctype_struct_nstring {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(holochain.ValidatingAction)
	res_0 := v.Name()
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5808
type proxyholochain_ValidatingAction _seq.Ref

func (p *proxyholochain_ValidatingAction) Bind_proxy_refnum__() int32 {
	return (*_seq.Ref)(p).Bind_IncNum()
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5815
func (p *proxyholochain_ValidatingAction) CheckValidationRequest(param_def *holochain.EntryDef) error {
	var _param_def _Ctype_int32_t = _seq.NullRefNum
	if param_def != nil {
		_param_def = _Ctype_int32_t(_seq.ToRefNum(param_def))
	}
	res := _Cfunc_cproxyholochain_ValidatingAction_CheckValidationRequest(_Ctype_int32_t(p.Bind_proxy_refnum__()), _param_def)
	var _res error
	_res_ref := _seq.FromRefNum(int32(res))
	if _res_ref != nil {
		if res < 0 {
			_res = _res_ref.Get().(error)
		} else {
			_res = (*proxy_error)(_res_ref)
		}
	}
	return _res
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5834
func (p *proxyholochain_ValidatingAction) Name() string {
	res := _Cfunc_cproxyholochain_ValidatingAction_Name(_Ctype_int32_t(p.Bind_proxy_refnum__()))
	_res := decodeString(res)
	return _res
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5843
func proxyholochain_Warrant_Decode(refnum _Ctype_int32_t, param_data _Ctype_struct_nbyteslice) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(holochain.Warrant)
	_param_data := toSlice(param_data, false)
	res_0 := v.Decode(_param_data)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5856
func proxyholochain_Warrant_Encode(refnum _Ctype_int32_t) (_Ctype_struct_nbyteslice, _Ctype_int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(holochain.Warrant)
	res_0, res_1 := v.Encode()
	_res_0 := fromSlice(res_0, true)
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5873
func proxyholochain_Warrant_Type(refnum _Ctype_int32_t) _Ctype_nint {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(holochain.Warrant)
	res_0 := v.Type()
	_res_0 := _Ctype_nint(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5882
func proxyholochain_Warrant_Verify(refnum _Ctype_int32_t, param_h _Ctype_int32_t) _Ctype_int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get().(holochain.Warrant)

	var _param_h *holochain.Holochain
	if _param_h_ref := _seq.FromRefNum(int32(param_h)); _param_h_ref != nil {
		_param_h = _param_h_ref.Get().(*holochain.Holochain)
	}
	res_0 := v.Verify(_param_h)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

type proxyholochain_Warrant _seq.Ref

func (p *proxyholochain_Warrant) Bind_proxy_refnum__() int32	{ return (*_seq.Ref)(p).Bind_IncNum() }

func (p *proxyholochain_Warrant) Decode(param_data []byte) error {
	_param_data := fromSlice(param_data, false)
	res := func(_cgo0 _Ctype_int32_t, _cgo1 _Ctype_struct_nbyteslice) _Ctype_int32_t {
//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5904
		_cgoCheckPointer(_cgo1)
//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5904
		return _Cfunc_cproxyholochain_Warrant_Decode(_cgo0, _cgo1)
//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5904
	}(_Ctype_int32_t(p.Bind_proxy_refnum__()), _param_data)
															var _res error
															_res_ref := _seq.FromRefNum(int32(res))
															if _res_ref != nil {
		if res < 0 {
			_res = _res_ref.Get().(error)
		} else {
			_res = (*proxy_error)(_res_ref)
		}
	}
	return _res
}

func (p *proxyholochain_Warrant) Encode() ([]byte, error) {
	res := _Cfunc_cproxyholochain_Warrant_Encode(_Ctype_int32_t(p.Bind_proxy_refnum__()))
	res_0 := toSlice(res.r0, true)
	var res_1 error
	res_1_ref := _seq.FromRefNum(int32(res.r1))
	if res_1_ref != nil {
		if res.r1 < 0 {
			res_1 = res_1_ref.Get().(error)
		} else {
			res_1 = (*proxy_error)(res_1_ref)
		}
	}
	return res_0, res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5934
func (p *proxyholochain_Warrant) Type() int {
	res := _Cfunc_cproxyholochain_Warrant_Type(_Ctype_int32_t(p.Bind_proxy_refnum__()))
	_res := int(res)
	return _res
}

func (p *proxyholochain_Warrant) Verify(param_h *holochain.Holochain) error {
	var _param_h _Ctype_int32_t = _seq.NullRefNum
	if param_h != nil {
		_param_h = _Ctype_int32_t(_seq.ToRefNum(param_h))
	}
	res := _Cfunc_cproxyholochain_Warrant_Verify(_Ctype_int32_t(p.Bind_proxy_refnum__()), _param_h)
	var _res error
	_res_ref := _seq.FromRefNum(int32(res))
	if _res_ref != nil {
		if res < 0 {
			_res = _res_ref.Get().(error)
		} else {
			_res = (*proxy_error)(_res_ref)
		}
	}
	return _res
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5959
func var_setholochain_AgentEntryDef(v _Ctype_int32_t) {

	var _v *holochain.EntryDef
	if _v_ref := _seq.FromRefNum(int32(v)); _v_ref != nil {
		_v = _v_ref.Get().(*holochain.EntryDef)
	}
	holochain.AgentEntryDef = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5969
func var_getholochain_AgentEntryDef() _Ctype_int32_t {
	v := holochain.AgentEntryDef
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5979
func var_setholochain_AlphaValue(v _Ctype_nint) {
	_v := int(v)
	holochain.AlphaValue = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5985
func var_getholochain_AlphaValue() _Ctype_nint {
	v := holochain.AlphaValue
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5992
func var_setholochain_BasicTemplateAppPackage(v _Ctype_struct_nstring) {
	_v := decodeString(v)
	holochain.BasicTemplateAppPackage = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:5998
func var_getholochain_BasicTemplateAppPackage() _Ctype_struct_nstring {
	v := holochain.BasicTemplateAppPackage
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6005
func var_setholochain_BridgeAppNotFoundErr(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.BridgeAppNotFoundErr = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6019
func var_getholochain_BridgeAppNotFoundErr() _Ctype_int32_t {
	v := holochain.BridgeAppNotFoundErr
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6031
func var_setholochain_CapabilityInvalidErr(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.CapabilityInvalidErr = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6045
func var_getholochain_CapabilityInvalidErr() _Ctype_int32_t {
	v := holochain.CapabilityInvalidErr
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6057
func var_setholochain_CloserPeerCount(v _Ctype_nint) {
	_v := int(v)
	holochain.CloserPeerCount = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6063
func var_getholochain_CloserPeerCount() _Ctype_nint {
	v := holochain.CloserPeerCount
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6070
func var_setholochain_DNAEntryDef(v _Ctype_int32_t) {

	var _v *holochain.EntryDef
	if _v_ref := _seq.FromRefNum(int32(v)); _v_ref != nil {
		_v = _v_ref.Get().(*holochain.EntryDef)
	}
	holochain.DNAEntryDef = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6080
func var_getholochain_DNAEntryDef() _Ctype_int32_t {
	v := holochain.DNAEntryDef
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6092
func var_setholochain_EnableAllLoggersEnv(v _Ctype_struct_nstring) {
	_v := decodeString(v)
	holochain.EnableAllLoggersEnv = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6098
func var_getholochain_EnableAllLoggersEnv() _Ctype_struct_nstring {
	v := holochain.EnableAllLoggersEnv
	_v := encodeString(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6105
func var_setholochain_ErrBlockedListed(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.ErrBlockedListed = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6119
func var_getholochain_ErrBlockedListed() _Ctype_int32_t {
	v := holochain.ErrBlockedListed
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6129
func var_setholochain_ErrDHTErrNoGossipersAvailable(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.ErrDHTErrNoGossipersAvailable = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6143
func var_getholochain_ErrDHTErrNoGossipersAvailable() _Ctype_int32_t {
	v := holochain.ErrDHTErrNoGossipersAvailable
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6153
func var_setholochain_ErrDHTExpectedGossipReqInBody(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.ErrDHTExpectedGossipReqInBody = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6167
func var_getholochain_ErrDHTExpectedGossipReqInBody() _Ctype_int32_t {
	v := holochain.ErrDHTExpectedGossipReqInBody
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6177
func var_setholochain_ErrDHTUnexpectedTypeInBody(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.ErrDHTUnexpectedTypeInBody = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6191
func var_getholochain_ErrDHTUnexpectedTypeInBody() _Ctype_int32_t {
	v := holochain.ErrDHTUnexpectedTypeInBody
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6201
func var_setholochain_ErrEmptyRoutingTable(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.ErrEmptyRoutingTable = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6215
func var_getholochain_ErrEmptyRoutingTable() _Ctype_int32_t {
	v := holochain.ErrEmptyRoutingTable
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6225
func var_setholochain_ErrEntryTypeMismatch(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.ErrEntryTypeMismatch = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6239
func var_getholochain_ErrEntryTypeMismatch() _Ctype_int32_t {
	v := holochain.ErrEntryTypeMismatch
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6249
func var_setholochain_ErrHashDeleted(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.ErrHashDeleted = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6263
func var_getholochain_ErrHashDeleted() _Ctype_int32_t {
	v := holochain.ErrHashDeleted
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6273
func var_setholochain_ErrHashModified(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.ErrHashModified = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6287
func var_getholochain_ErrHashModified() _Ctype_int32_t {
	v := holochain.ErrHashModified
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6297
func var_setholochain_ErrHashNotFound(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.ErrHashNotFound = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6311
func var_getholochain_ErrHashNotFound() _Ctype_int32_t {
	v := holochain.ErrHashNotFound
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6321
func var_setholochain_ErrHashRejected(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.ErrHashRejected = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6335
func var_getholochain_ErrHashRejected() _Ctype_int32_t {
	v := holochain.ErrHashRejected
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6345
func var_setholochain_ErrIncompleteChain(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.ErrIncompleteChain = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6359
func var_getholochain_ErrIncompleteChain() _Ctype_int32_t {
	v := holochain.ErrIncompleteChain
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6369
func var_setholochain_ErrLinkNotFound(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.ErrLinkNotFound = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6383
func var_getholochain_ErrLinkNotFound() _Ctype_int32_t {
	v := holochain.ErrLinkNotFound
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6393
func var_setholochain_ErrNilEntryInvalid(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.ErrNilEntryInvalid = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6407
func var_getholochain_ErrNilEntryInvalid() _Ctype_int32_t {
	v := holochain.ErrNilEntryInvalid
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6417
func var_setholochain_ErrNoSuchIdx(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.ErrNoSuchIdx = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6431
func var_getholochain_ErrNoSuchIdx() _Ctype_int32_t {
	v := holochain.ErrNoSuchIdx
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6441
func var_setholochain_ErrNotValidForAgentType(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.ErrNotValidForAgentType = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6455
func var_getholochain_ErrNotValidForAgentType() _Ctype_int32_t {
	v := holochain.ErrNotValidForAgentType
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6465
func var_setholochain_ErrNotValidForDNAType(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.ErrNotValidForDNAType = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6479
func var_getholochain_ErrNotValidForDNAType() _Ctype_int32_t {
	v := holochain.ErrNotValidForDNAType
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6489
func var_setholochain_ErrNotValidForKeyType(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.ErrNotValidForKeyType = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6503
func var_getholochain_ErrNotValidForKeyType() _Ctype_int32_t {
	v := holochain.ErrNotValidForKeyType
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6513
func var_setholochain_ErrPutLinkOverDeleted(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.ErrPutLinkOverDeleted = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6527
func var_getholochain_ErrPutLinkOverDeleted() _Ctype_int32_t {
	v := holochain.ErrPutLinkOverDeleted
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6537
func var_setholochain_ErrWrongNargs(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.ErrWrongNargs = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6551
func var_getholochain_ErrWrongNargs() _Ctype_int32_t {
	v := holochain.ErrWrongNargs
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6561
func var_setholochain_IsDevMode(v _Ctype_char) {
	_v := v != 0
	holochain.IsDevMode = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6567
func var_getholochain_IsDevMode() _Ctype_char {
	v := holochain.IsDevMode
	var _v _Ctype_char = 0
	if v {
		_v = 1
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6577
func var_setholochain_KValue(v _Ctype_nint) {
	_v := int(v)
	holochain.KValue = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6583
func var_getholochain_KValue() _Ctype_nint {
	v := holochain.KValue
	_v := _Ctype_nint(v)
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6590
func var_setholochain_KeyEntryDef(v _Ctype_int32_t) {

	var _v *holochain.EntryDef
	if _v_ref := _seq.FromRefNum(int32(v)); _v_ref != nil {
		_v = _v_ref.Get().(*holochain.EntryDef)
	}
	holochain.KeyEntryDef = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6600
func var_getholochain_KeyEntryDef() _Ctype_int32_t {
	v := holochain.KeyEntryDef
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6610
func var_setholochain_NonCallableAction(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.NonCallableAction = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6624
func var_getholochain_NonCallableAction() _Ctype_int32_t {
	v := holochain.NonCallableAction
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6634
func var_setholochain_NonDHTAction(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.NonDHTAction = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6648
func var_getholochain_NonDHTAction() _Ctype_int32_t {
	v := holochain.NonDHTAction
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6658
func var_setholochain_SelfRevocationDoesNotVerify(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.SelfRevocationDoesNotVerify = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6672
func var_getholochain_SelfRevocationDoesNotVerify() _Ctype_int32_t {
	v := holochain.SelfRevocationDoesNotVerify
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6682
func var_setholochain_SendTimeoutErr(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.SendTimeoutErr = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6696
func var_getholochain_SendTimeoutErr() _Ctype_int32_t {
	v := holochain.SendTimeoutErr
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6706
func var_setholochain_UnknownWarrantTypeErr(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.UnknownWarrantTypeErr = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6720
func var_getholochain_UnknownWarrantTypeErr() _Ctype_int32_t {
	v := holochain.UnknownWarrantTypeErr
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6730
func var_setholochain_ValidationFailedErr(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.ValidationFailedErr = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6744
func var_getholochain_ValidationFailedErr() _Ctype_int32_t {
	v := holochain.ValidationFailedErr
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6754
func var_setholochain_WarrantPropertyNotFoundErr(v _Ctype_int32_t) {
	var _v error
	_v_ref := _seq.FromRefNum(int32(v))
	if _v_ref != nil {
		if v < 0 {
			_v = _v_ref.Get().(error)
		} else {
			_v = (*proxy_error)(_v_ref)
		}
	}
	holochain.WarrantPropertyNotFoundErr = _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6768
func var_getholochain_WarrantPropertyNotFoundErr() _Ctype_int32_t {
	v := holochain.WarrantPropertyNotFoundErr
	var _v _Ctype_int32_t = _seq.NullRefNum
	if v != nil {
		_v = _Ctype_int32_t(_seq.ToRefNum(v))
	}
	return _v
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6779
func proxyholochain__BootstrapRefreshTask(param_h _Ctype_int32_t) {

	var _param_h *holochain.Holochain
	if _param_h_ref := _seq.FromRefNum(int32(param_h)); _param_h_ref != nil {
		_param_h = _param_h_ref.Get().(*holochain.Holochain)
	}
	holochain.BootstrapRefreshTask(_param_h)
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6789
func proxyholochain__BuildJSONSchemaValidatorFromFile(param_path _Ctype_struct_nstring) (_Ctype_int32_t, _Ctype_int32_t) {
	_param_path := decodeString(param_path)
	res_0, res_1 := holochain.BuildJSONSchemaValidatorFromFile(_param_path)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6804
func proxyholochain__BuildJSONSchemaValidatorFromString(param_input _Ctype_struct_nstring) (_Ctype_int32_t, _Ctype_int32_t) {
	_param_input := decodeString(param_input)
	res_0, res_1 := holochain.BuildJSONSchemaValidatorFromString(_param_input)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6821
func proxyholochain__CopyDir(param_source _Ctype_struct_nstring, param_dest _Ctype_struct_nstring) _Ctype_int32_t {
	_param_source := decodeString(param_source)
	_param_dest := decodeString(param_dest)
	res_0 := holochain.CopyDir(_param_source, _param_dest)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6833
func proxyholochain__CopyFile(param_source _Ctype_struct_nstring, param_dest _Ctype_struct_nstring) _Ctype_int32_t {
	_param_source := decodeString(param_source)
	_param_dest := decodeString(param_dest)
	res_0 := holochain.CopyFile(_param_source, _param_dest)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6845
func proxyholochain__CreateRibosome(param_h _Ctype_int32_t, param_zome _Ctype_int32_t) (_Ctype_int32_t, _Ctype_int32_t) {

	var _param_h *holochain.Holochain
	if _param_h_ref := _seq.FromRefNum(int32(param_h)); _param_h_ref != nil {
		_param_h = _param_h_ref.Get().(*holochain.Holochain)
	}

	var _param_zome *holochain.Zome
	if _param_zome_ref := _seq.FromRefNum(int32(param_zome)); _param_zome_ref != nil {
		_param_zome = _param_zome_ref.Get().(*holochain.Zome)
	}
	res_0, res_1 := holochain.CreateRibosome(_param_h, _param_zome)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6870
func proxyholochain__Debug(param_m _Ctype_struct_nstring) {
	_param_m := decodeString(param_m)
	holochain.Debug(_param_m)
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6878
func proxyholochain__DecodeWarrant(param_warrantType _Ctype_nint, param_data _Ctype_struct_nbyteslice) (_Ctype_int32_t, _Ctype_int32_t) {
	_param_warrantType := int(param_warrantType)
	_param_data := toSlice(param_data, false)
	res_0, res_1 := holochain.DecodeWarrant(_param_warrantType, _param_data)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6896
func proxyholochain__EncodingFormat(param_file _Ctype_struct_nstring) _Ctype_struct_nstring {
	_param_file := decodeString(param_file)
	res_0 := holochain.EncodingFormat(_param_file)
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6904
func proxyholochain__EscapeJSONValue(param_json _Ctype_struct_nstring) _Ctype_struct_nstring {
	_param_json := decodeString(param_json)
	res_0 := holochain.EscapeJSONValue(_param_json)
	_res_0 := encodeString(res_0)
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6918
func proxyholochain__GossipTask(param_h _Ctype_int32_t) {

	var _param_h *holochain.Holochain
	if _param_h_ref := _seq.FromRefNum(int32(param_h)); _param_h_ref != nil {
		_param_h = _param_h_ref.Get().(*holochain.Holochain)
	}
	holochain.GossipTask(_param_h)
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6928
func proxyholochain__Info(param_m _Ctype_struct_nstring) {
	_param_m := decodeString(param_m)
	holochain.Info(_param_m)
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6936
func proxyholochain__InitializeHolochain() {
	holochain.InitializeHolochain()
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6941
func proxyholochain__IsInitialized(param_root _Ctype_struct_nstring) _Ctype_char {
	_param_root := decodeString(param_root)
	res_0 := holochain.IsInitialized(_param_root)
	var _res_0 _Ctype_char = 0
	if res_0 {
		_res_0 = 1
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6953
func proxyholochain__LoadAgent(param_path _Ctype_struct_nstring) (_Ctype_int32_t, _Ctype_int32_t) {
	_param_path := decodeString(param_path)
	res_0, res_1 := holochain.LoadAgent(_param_path)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6969
func proxyholochain__LoadService(param_path _Ctype_struct_nstring) (_Ctype_int32_t, _Ctype_int32_t) {
	_param_path := decodeString(param_path)
	res_0, res_1 := holochain.LoadService(_param_path)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:6984
func proxyholochain__LoadTestConfig(param_dir _Ctype_struct_nstring) (_Ctype_int32_t, _Ctype_int32_t) {
	_param_dir := decodeString(param_dir)
	res_0, res_1 := holochain.LoadTestConfig(_param_dir)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7001
func proxyholochain__MakeActionFromMessage(param_msg _Ctype_int32_t) (_Ctype_int32_t, _Ctype_int32_t) {

	var _param_msg *holochain.Message
	if _param_msg_ref := _seq.FromRefNum(int32(param_msg)); _param_msg_ref != nil {
		_param_msg = _param_msg_ref.Get().(*holochain.Message)
	}
	res_0, res_1 := holochain.MakeActionFromMessage(_param_msg)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7020
func proxyholochain__MakeDirs(param_devPath _Ctype_struct_nstring) _Ctype_int32_t {
	_param_devPath := decodeString(param_devPath)
	res_0 := holochain.MakeDirs(_param_devPath)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7032
func proxyholochain__MakeValidationPackage(param_h _Ctype_int32_t, param_pkg _Ctype_int32_t) (_Ctype_int32_t, _Ctype_int32_t) {

	var _param_h *holochain.Holochain
	if _param_h_ref := _seq.FromRefNum(int32(param_h)); _param_h_ref != nil {
		_param_h = _param_h_ref.Get().(*holochain.Holochain)
	}

	var _param_pkg *holochain.Package
	if _param_pkg_ref := _seq.FromRefNum(int32(param_pkg)); _param_pkg_ref != nil {
		_param_pkg = _param_pkg_ref.Get().(*holochain.Package)
	}
	res_0, res_1 := holochain.MakeValidationPackage(_param_h, _param_pkg)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7065
func proxyholochain__NewCommitAction(param_entryType _Ctype_struct_nstring, param_entry _Ctype_int32_t) _Ctype_int32_t {
	_param_entryType := decodeString(param_entryType)
	var _param_entry holochain.Entry
	_param_entry_ref := _seq.FromRefNum(int32(param_entry))
	if _param_entry_ref != nil {
		if param_entry < 0 {
			_param_entry = _param_entry_ref.Get().(holochain.Entry)
		} else {
			_param_entry = (*proxyholochain_Entry)(_param_entry_ref)
		}
	}
	res_0 := holochain.NewCommitAction(_param_entryType, _param_entry)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7085
func proxyholochain__NewDHT(param_h _Ctype_int32_t) _Ctype_int32_t {

	var _param_h *holochain.Holochain
	if _param_h_ref := _seq.FromRefNum(int32(param_h)); _param_h_ref != nil {
		_param_h = _param_h_ref.Get().(*holochain.Holochain)
	}
	res_0 := holochain.NewDHT(_param_h)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7100
func proxyholochain__NewDebugAction(param_msg _Ctype_struct_nstring) _Ctype_int32_t {
	_param_msg := decodeString(param_msg)
	res_0 := holochain.NewDebugAction(_param_msg)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7114
func proxyholochain__NewGetBridgesAction(param_doc _Ctype_struct_nbyteslice) _Ctype_int32_t {
	_param_doc := toSlice(param_doc, false)
	res_0 := holochain.NewGetBridgesAction(_param_doc)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7125
func proxyholochain__NewGetLinksAction(param_linkQuery _Ctype_int32_t, param_options _Ctype_int32_t) _Ctype_int32_t {

	var _param_linkQuery *holochain.LinkQuery
	if _param_linkQuery_ref := _seq.FromRefNum(int32(param_linkQuery)); _param_linkQuery_ref != nil {
		_param_linkQuery = _param_linkQuery_ref.Get().(*holochain.LinkQuery)
	}

	var _param_options *holochain.GetLinksOptions
	if _param_options_ref := _seq.FromRefNum(int32(param_options)); _param_options_ref != nil {
		_param_options = _param_options_ref.Get().(*holochain.GetLinksOptions)
	}
	res_0 := holochain.NewGetLinksAction(_param_linkQuery, _param_options)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7146
func proxyholochain__NewJSRibosome(param_h _Ctype_int32_t, param_zome _Ctype_int32_t) (_Ctype_int32_t, _Ctype_int32_t) {

	var _param_h *holochain.Holochain
	if _param_h_ref := _seq.FromRefNum(int32(param_h)); _param_h_ref != nil {
		_param_h = _param_h_ref.Get().(*holochain.Holochain)
	}

	var _param_zome *holochain.Zome
	if _param_zome_ref := _seq.FromRefNum(int32(param_zome)); _param_zome_ref != nil {
		_param_zome = _param_zome_ref.Get().(*holochain.Zome)
	}
	res_0, res_1 := holochain.NewJSRibosome(_param_h, _param_zome)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7172
func proxyholochain__NewMakeHashAction(param_entry _Ctype_int32_t) _Ctype_int32_t {
	var _param_entry holochain.Entry
	_param_entry_ref := _seq.FromRefNum(int32(param_entry))
	if _param_entry_ref != nil {
		if param_entry < 0 {
			_param_entry = _param_entry_ref.Get().(holochain.Entry)
		} else {
			_param_entry = (*proxyholochain_Entry)(_param_entry_ref)
		}
	}
	res_0 := holochain.NewMakeHashAction(_param_entry)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7193
func proxyholochain__NewNode(param_listenAddr _Ctype_struct_nstring, param_protoMux _Ctype_struct_nstring, param_agent _Ctype_int32_t, param_enableNATUPnP _Ctype_char, param_log _Ctype_int32_t) (_Ctype_int32_t, _Ctype_int32_t) {
	_param_listenAddr := decodeString(param_listenAddr)
	_param_protoMux := decodeString(param_protoMux)

	var _param_agent *holochain.LibP2PAgent
	if _param_agent_ref := _seq.FromRefNum(int32(param_agent)); _param_agent_ref != nil {
		_param_agent = _param_agent_ref.Get().(*holochain.LibP2PAgent)
	}
	_param_enableNATUPnP := param_enableNATUPnP != 0

	var _param_log *holochain.Logger
	if _param_log_ref := _seq.FromRefNum(int32(param_log)); _param_log_ref != nil {
		_param_log = _param_log_ref.Get().(*holochain.Logger)
	}
	res_0, res_1 := holochain.NewNode(_param_listenAddr, _param_protoMux, _param_agent, _param_enableNATUPnP, _param_log)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7220
func proxyholochain__NewNucleus(param_h _Ctype_int32_t, param_dna _Ctype_int32_t) _Ctype_int32_t {

	var _param_h *holochain.Holochain
	if _param_h_ref := _seq.FromRefNum(int32(param_h)); _param_h_ref != nil {
		_param_h = _param_h_ref.Get().(*holochain.Holochain)
	}

	var _param_dna *holochain.DNA
	if _param_dna_ref := _seq.FromRefNum(int32(param_dna)); _param_dna_ref != nil {
		_param_dna = _param_dna_ref.Get().(*holochain.DNA)
	}
	res_0 := holochain.NewNucleus(_param_h, _param_dna)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7240
func proxyholochain__NewPropertyAction(param_prop _Ctype_struct_nstring) _Ctype_int32_t {
	_param_prop := decodeString(param_prop)
	res_0 := holochain.NewPropertyAction(_param_prop)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7251
func proxyholochain__NewPutAction(param_entryType _Ctype_struct_nstring, param_entry _Ctype_int32_t, param_header _Ctype_int32_t) _Ctype_int32_t {
	_param_entryType := decodeString(param_entryType)
	var _param_entry holochain.Entry
	_param_entry_ref := _seq.FromRefNum(int32(param_entry))
	if _param_entry_ref != nil {
		if param_entry < 0 {
			_param_entry = _param_entry_ref.Get().(holochain.Entry)
		} else {
			_param_entry = (*proxyholochain_Entry)(_param_entry_ref)
		}
	}

	var _param_header *holochain.Header
	if _param_header_ref := _seq.FromRefNum(int32(param_header)); _param_header_ref != nil {
		_param_header = _param_header_ref.Get().(*holochain.Header)
	}
	res_0 := holochain.NewPutAction(_param_entryType, _param_entry, _param_header)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7276
func proxyholochain__NewQueryAction(param_options _Ctype_int32_t) _Ctype_int32_t {

	var _param_options *holochain.QueryOptions
	if _param_options_ref := _seq.FromRefNum(int32(param_options)); _param_options_ref != nil {
		_param_options = _param_options_ref.Get().(*holochain.QueryOptions)
	}
	res_0 := holochain.NewQueryAction(_param_options)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7293
func proxyholochain__NewSelfRevocationWarrant(param_revocation _Ctype_int32_t) (_Ctype_int32_t, _Ctype_int32_t) {

	var _param_revocation *holochain.SelfRevocation
	if _param_revocation_ref := _seq.FromRefNum(int32(param_revocation)); _param_revocation_ref != nil {
		_param_revocation = _param_revocation_ref.Get().(*holochain.SelfRevocation)
	}
	res_0, res_1 := holochain.NewSelfRevocationWarrant(_param_revocation)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7313
func proxyholochain__NewSignAction(param_doc _Ctype_struct_nbyteslice) _Ctype_int32_t {
	_param_doc := toSlice(param_doc, false)
	res_0 := holochain.NewSignAction(_param_doc)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7324
func proxyholochain__NewVerifySignatureAction(param_signature _Ctype_struct_nstring, param_data _Ctype_struct_nstring, param_pubKey _Ctype_struct_nstring) _Ctype_int32_t {
	_param_signature := decodeString(param_signature)
	_param_data := decodeString(param_data)
	_param_pubKey := decodeString(param_pubKey)
	res_0 := holochain.NewVerifySignatureAction(_param_signature, _param_data, _param_pubKey)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7337
func proxyholochain__NewZygoRibosome(param_h _Ctype_int32_t, param_zome _Ctype_int32_t) (_Ctype_int32_t, _Ctype_int32_t) {

	var _param_h *holochain.Holochain
	if _param_h_ref := _seq.FromRefNum(int32(param_h)); _param_h_ref != nil {
		_param_h = _param_h_ref.Get().(*holochain.Holochain)
	}

	var _param_zome *holochain.Zome
	if _param_zome_ref := _seq.FromRefNum(int32(param_zome)); _param_zome_ref != nil {
		_param_zome = _param_zome_ref.Get().(*holochain.Zome)
	}
	res_0, res_1 := holochain.NewZygoRibosome(_param_h, _param_zome)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7361
func proxyholochain__PrettyPrintJSON(param_b _Ctype_struct_nbyteslice) (_Ctype_struct_nstring, _Ctype_int32_t) {
	_param_b := toSlice(param_b, false)
	res_0, res_1 := holochain.PrettyPrintJSON(_param_b)
	_res_0 := encodeString(res_0)
	var _res_1 _Ctype_int32_t = _seq.NullRefNum
	if res_1 != nil {
		_res_1 = _Ctype_int32_t(_seq.ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7374
func proxyholochain__RegisterBultinRibosomes() {
	holochain.RegisterBultinRibosomes()
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7380
func proxyholochain__RetryTask(param_h _Ctype_int32_t) {

	var _param_h *holochain.Holochain
	if _param_h_ref := _seq.FromRefNum(int32(param_h)); _param_h_ref != nil {
		_param_h = _param_h_ref.Get().(*holochain.Holochain)
	}
	holochain.RetryTask(_param_h)
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7390
func proxyholochain__RoutingRefreshTask(param_h _Ctype_int32_t) {

	var _param_h *holochain.Holochain
	if _param_h_ref := _seq.FromRefNum(int32(param_h)); _param_h_ref != nil {
		_param_h = _param_h_ref.Get().(*holochain.Holochain)
	}
	holochain.RoutingRefreshTask(_param_h)
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7401
func proxyholochain__SaveAgent(param_path _Ctype_struct_nstring, param_agent _Ctype_int32_t) _Ctype_int32_t {
	_param_path := decodeString(param_path)
	var _param_agent holochain.Agent
	_param_agent_ref := _seq.FromRefNum(int32(param_agent))
	if _param_agent_ref != nil {
		if param_agent < 0 {
			_param_agent = _param_agent_ref.Get().(holochain.Agent)
		} else {
			_param_agent = (*proxyholochain_Agent)(_param_agent_ref)
		}
	}
	res_0 := holochain.SaveAgent(_param_path, _param_agent)
	var _res_0 _Ctype_int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = _Ctype_int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go:7422
func proxyholochain__TestingAppAppPackage() _Ctype_struct_nstring {
	res_0 := holochain.TestingAppAppPackage()
	_res_0 := encodeString(res_0)
	return _res_0
}
