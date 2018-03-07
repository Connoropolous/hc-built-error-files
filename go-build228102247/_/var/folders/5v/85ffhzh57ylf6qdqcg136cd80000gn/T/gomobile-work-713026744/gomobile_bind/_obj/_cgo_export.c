/* Created by cgo - DO NOT EDIT. */
#include <stdlib.h>
#include "_cgo_export.h"

extern void crosscall2(void (*fn)(void *, int, __SIZE_TYPE__), void *, int, __SIZE_TYPE__);
extern __SIZE_TYPE__ _cgo_wait_runtime_init_done();
extern void _cgo_release_context(__SIZE_TYPE__);

extern char* _cgo_topofstack(void);
#define CGO_NO_SANITIZE_THREAD
#define _cgo_tsan_acquire()
#define _cgo_tsan_release()

extern void _cgoexp_ebd397278953_initClasses(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void initClasses()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		char unused;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_initClasses, &a, 0, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionBridge_Name(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ActionBridge_Name(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionBridge_Name, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ActionBridge(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ActionBridge()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ActionBridge, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionCall_Name(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ActionCall_Name(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionCall_Name, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ActionCall(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ActionCall()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ActionCall, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionCommit_CheckValidationRequest(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_ActionCommit_CheckValidationRequest(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionCommit_CheckValidationRequest, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionCommit_Entry(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_ActionCommit_Entry(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionCommit_Entry, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionCommit_EntryType(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ActionCommit_EntryType(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionCommit_EntryType, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionCommit_Name(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ActionCommit_Name(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionCommit_Name, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionCommit_SetHeader(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ActionCommit_SetHeader(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionCommit_SetHeader, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_new_holochain_ActionCommit(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ActionCommit()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ActionCommit, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionDebug_Name(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ActionDebug_Name(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionDebug_Name, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ActionDebug(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ActionDebug()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ActionDebug, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionDel_CheckValidationRequest(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_ActionDel_CheckValidationRequest(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionDel_CheckValidationRequest, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionDel_Entry(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_ActionDel_Entry(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionDel_Entry, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionDel_EntryType(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ActionDel_EntryType(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionDel_EntryType, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionDel_Name(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ActionDel_Name(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionDel_Name, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionDel_SetHeader(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ActionDel_SetHeader(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionDel_SetHeader, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_new_holochain_ActionDel(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ActionDel()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ActionDel, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionGet_Name(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ActionGet_Name(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionGet_Name, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ActionGet(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ActionGet()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ActionGet, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionGetBridges_Name(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ActionGetBridges_Name(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionGetBridges_Name, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ActionGetBridges(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ActionGetBridges()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ActionGetBridges, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionGetLinks_Name(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ActionGetLinks_Name(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionGetLinks_Name, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ActionGetLinks(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ActionGetLinks()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ActionGetLinks, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionLink_CheckValidationRequest(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_ActionLink_CheckValidationRequest(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionLink_CheckValidationRequest, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionLink_Name(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ActionLink_Name(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionLink_Name, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ActionLink(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ActionLink()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ActionLink, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionListAdd_Name(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ActionListAdd_Name(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionListAdd_Name, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ActionListAdd(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ActionListAdd()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ActionListAdd, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionMakeHash_Name(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ActionMakeHash_Name(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionMakeHash_Name, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ActionMakeHash(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ActionMakeHash()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ActionMakeHash, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionMod_CheckValidationRequest(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_ActionMod_CheckValidationRequest(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionMod_CheckValidationRequest, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionMod_Entry(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_ActionMod_Entry(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionMod_Entry, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionMod_EntryType(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ActionMod_EntryType(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionMod_EntryType, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionMod_Name(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ActionMod_Name(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionMod_Name, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionMod_SetHeader(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ActionMod_SetHeader(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionMod_SetHeader, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_new_holochain_ActionMod(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ActionMod()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ActionMod, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionModAgent_Revocation_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ActionModAgent_Revocation_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionModAgent_Revocation_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionModAgent_Revocation_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ActionModAgent_Revocation_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionModAgent_Revocation_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionModAgent_Name(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ActionModAgent_Name(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionModAgent_Name, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ActionModAgent(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ActionModAgent()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ActionModAgent, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionProperty_Name(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ActionProperty_Name(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionProperty_Name, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ActionProperty(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ActionProperty()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ActionProperty, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionPut_CheckValidationRequest(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_ActionPut_CheckValidationRequest(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionPut_CheckValidationRequest, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionPut_Name(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ActionPut_Name(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionPut_Name, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ActionPut(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ActionPut()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ActionPut, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionQuery_Name(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ActionQuery_Name(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionQuery_Name, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ActionQuery(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ActionQuery()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ActionQuery, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionSend_Name(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ActionSend_Name(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionSend_Name, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ActionSend(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ActionSend()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ActionSend, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionSign_Name(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ActionSign_Name(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionSign_Name, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ActionSign(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ActionSign()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ActionSign, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ActionVerifySignature_Name(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ActionVerifySignature_Name(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ActionVerifySignature_Name, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ActionVerifySignature(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ActionVerifySignature()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ActionVerifySignature, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_AgentEntry_Revocation_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_AgentEntry_Revocation_Set(int32_t p0, nbyteslice p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AgentEntry_Revocation_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_AgentEntry_Revocation_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nbyteslice proxyholochain_AgentEntry_Revocation_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AgentEntry_Revocation_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_AgentEntry_PublicKey_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_AgentEntry_PublicKey_Set(int32_t p0, nbyteslice p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AgentEntry_PublicKey_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_AgentEntry_PublicKey_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nbyteslice proxyholochain_AgentEntry_PublicKey_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AgentEntry_PublicKey_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_AgentEntry(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_AgentEntry()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_AgentEntry, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_AgentFixture_Hash_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_AgentFixture_Hash_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AgentFixture_Hash_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_AgentFixture_Hash_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_AgentFixture_Hash_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AgentFixture_Hash_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_AgentFixture_Identity_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_AgentFixture_Identity_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AgentFixture_Identity_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_AgentFixture_Identity_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_AgentFixture_Identity_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AgentFixture_Identity_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_AgentFixture(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_AgentFixture()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_AgentFixture, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_AppMsg_ZomeType_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_AppMsg_ZomeType_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AppMsg_ZomeType_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_AppMsg_ZomeType_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_AppMsg_ZomeType_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AppMsg_ZomeType_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_AppMsg_Body_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_AppMsg_Body_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AppMsg_Body_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_AppMsg_Body_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_AppMsg_Body_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AppMsg_Body_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_AppMsg(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_AppMsg()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_AppMsg, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_AppPackage_Version_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_AppPackage_Version_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AppPackage_Version_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_AppPackage_Version_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_AppPackage_Version_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AppPackage_Version_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_AppPackage_Generator_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_AppPackage_Generator_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AppPackage_Generator_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_AppPackage_Generator_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_AppPackage_Generator_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AppPackage_Generator_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_AppPackage(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_AppPackage()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_AppPackage, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_AppPackageScenario_Name_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_AppPackageScenario_Name_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AppPackageScenario_Name_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_AppPackageScenario_Name_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_AppPackageScenario_Name_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AppPackageScenario_Name_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_AppPackageScenario(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_AppPackageScenario()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_AppPackageScenario, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_AppPackageTests_Name_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_AppPackageTests_Name_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AppPackageTests_Name_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_AppPackageTests_Name_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_AppPackageTests_Name_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AppPackageTests_Name_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_AppPackageTests(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_AppPackageTests()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_AppPackageTests, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_FileName_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_AppPackageUIFile_FileName_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_FileName_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_FileName_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_AppPackageUIFile_FileName_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_FileName_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Data_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_AppPackageUIFile_Data_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Data_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Data_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_AppPackageUIFile_Data_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Data_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Encoding_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_AppPackageUIFile_Encoding_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Encoding_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Encoding_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_AppPackageUIFile_Encoding_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_AppPackageUIFile_Encoding_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_AppPackageUIFile(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_AppPackageUIFile()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_AppPackageUIFile, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Arg_Name_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Arg_Name_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Arg_Name_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Arg_Name_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Arg_Name_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Arg_Name_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Arg_Optional_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Arg_Optional_Set(int32_t p0, char p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char p1;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Arg_Optional_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Arg_Optional_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_Arg_Optional_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Arg_Optional_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_Arg(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Arg()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Arg, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_BSReq_Version_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_BSReq_Version_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_BSReq_Version_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_BSReq_Version_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_BSReq_Version_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_BSReq_Version_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_BSReq_NodeID_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_BSReq_NodeID_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_BSReq_NodeID_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_BSReq_NodeID_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_BSReq_NodeID_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_BSReq_NodeID_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_BSReq_NodeAddr_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_BSReq_NodeAddr_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_BSReq_NodeAddr_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_BSReq_NodeAddr_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_BSReq_NodeAddr_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_BSReq_NodeAddr_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_BSReq(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_BSReq()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_BSReq, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_BSResp_Remote_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_BSResp_Remote_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_BSResp_Remote_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_BSResp_Remote_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_BSResp_Remote_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_BSResp_Remote_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_BSResp(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_BSResp()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_BSResp, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Bridge_Token_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Bridge_Token_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Bridge_Token_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Bridge_Token_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Bridge_Token_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Bridge_Token_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Bridge_Side_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Bridge_Side_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Bridge_Side_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Bridge_Side_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_Bridge_Side_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Bridge_Side_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_Bridge(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Bridge()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Bridge, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_BridgeApp_H_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_BridgeApp_H_Set(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_BridgeApp_H_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_BridgeApp_H_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_BridgeApp_H_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_BridgeApp_H_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_BridgeApp_Side_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_BridgeApp_Side_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_BridgeApp_Side_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_BridgeApp_Side_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_BridgeApp_Side_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_BridgeApp_Side_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataFrom_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_BridgeApp_BridgeGenesisDataFrom_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataFrom_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataFrom_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_BridgeApp_BridgeGenesisDataFrom_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataFrom_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataTo_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_BridgeApp_BridgeGenesisDataTo_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataTo_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataTo_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_BridgeApp_BridgeGenesisDataTo_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_BridgeApp_BridgeGenesisDataTo_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_BridgeApp_Port_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_BridgeApp_Port_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_BridgeApp_Port_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_BridgeApp_Port_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_BridgeApp_Port_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_BridgeApp_Port_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_BridgeApp(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_BridgeApp()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_BridgeApp, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Bucket_Len(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_Bucket_Len(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Bucket_Len, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_Bucket(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Bucket()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Bucket, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_BytesSent_Bytes_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_BytesSent_Bytes_Set(int32_t p0, int64_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int64_t p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_BytesSent_Bytes_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_BytesSent_Bytes_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int64_t proxyholochain_BytesSent_Bytes_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int64_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_BytesSent_Bytes_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_BytesSent(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_BytesSent()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_BytesSent, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Callback_Function_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Callback_Function_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Callback_Function_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Callback_Function_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Callback_Function_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Callback_Function_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Callback_ID_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Callback_ID_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Callback_ID_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Callback_ID_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Callback_ID_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Callback_ID_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_Callback(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Callback()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Callback, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Capability_Token_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Capability_Token_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Capability_Token_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Capability_Token_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Capability_Token_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Capability_Token_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_Capability(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Capability()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Capability, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Chain_Close(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Chain_Close(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Chain_Close, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Chain_JSON(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_Chain_JSON_return proxyholochain_Chain_JSON(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_Chain_JSON_return r;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Chain_JSON, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_Chain_Length(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_Chain_Length(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Chain_Length, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Chain_Nth(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Chain_Nth(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Chain_Nth, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Chain_String(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Chain_String(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Chain_String, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Chain_Top(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Chain_Top(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Chain_Top, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Chain_Validate(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Chain_Validate(int32_t p0, char p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char p1;
		char __pad0[3];
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Chain_Validate, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_Chain(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Chain()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Chain, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ChainPair_Header_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ChainPair_Header_Set(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ChainPair_Header_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_ChainPair_Header_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_ChainPair_Header_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ChainPair_Header_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ChainPair_Entry_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ChainPair_Entry_Set(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ChainPair_Entry_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_ChainPair_Entry_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_ChainPair_Entry_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ChainPair_Entry_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ChainPair(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ChainPair()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ChainPair, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Change_Message_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Change_Message_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Change_Message_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Change_Message_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Change_Message_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Change_Message_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Change_AsOf_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Change_AsOf_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Change_AsOf_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Change_AsOf_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_Change_AsOf_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Change_AsOf_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Change_Log(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Change_Log(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Change_Log, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_new_holochain_Change(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Change()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Change, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_CloneSpec_Role_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_CloneSpec_Role_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_CloneSpec_Role_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_CloneSpec_Role_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_CloneSpec_Role_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_CloneSpec_Role_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_CloneSpec_Number_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_CloneSpec_Number_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_CloneSpec_Number_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_CloneSpec_Number_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_CloneSpec_Number_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_CloneSpec_Number_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_CloneSpec(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_CloneSpec()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_CloneSpec, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_CloserPeersResp(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_CloserPeersResp()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_CloserPeersResp, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Config_Port_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Config_Port_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Config_Port_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Config_Port_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_Config_Port_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Config_Port_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Config_EnableMDNS_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Config_EnableMDNS_Set(int32_t p0, char p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char p1;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Config_EnableMDNS_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Config_EnableMDNS_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_Config_EnableMDNS_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Config_EnableMDNS_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Config_PeerModeAuthor_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Config_PeerModeAuthor_Set(int32_t p0, char p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char p1;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Config_PeerModeAuthor_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Config_PeerModeAuthor_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_Config_PeerModeAuthor_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Config_PeerModeAuthor_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Config_PeerModeDHTNode_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Config_PeerModeDHTNode_Set(int32_t p0, char p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char p1;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Config_PeerModeDHTNode_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Config_PeerModeDHTNode_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_Config_PeerModeDHTNode_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Config_PeerModeDHTNode_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Config_EnableNATUPnP_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Config_EnableNATUPnP_Set(int32_t p0, char p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char p1;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Config_EnableNATUPnP_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Config_EnableNATUPnP_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_Config_EnableNATUPnP_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Config_EnableNATUPnP_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Config_BootstrapServer_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Config_BootstrapServer_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Config_BootstrapServer_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Config_BootstrapServer_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Config_BootstrapServer_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Config_BootstrapServer_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Config_Setup(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Config_Setup(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Config_Setup, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Config_SetupLogging(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Config_SetupLogging(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Config_SetupLogging, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_Config(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Config()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Config, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_DHT_Close(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_DHT_Close(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DHT_Close, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_DHT_DumpIdx(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_DHT_DumpIdx_return proxyholochain_DHT_DumpIdx(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
		nstring r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_DHT_DumpIdx_return r;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DHT_DumpIdx, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_DHT_DumpIdxJSON(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_DHT_DumpIdxJSON_return proxyholochain_DHT_DumpIdxJSON(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
		nstring r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_DHT_DumpIdxJSON_return r;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DHT_DumpIdxJSON, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_DHT_GetIdx(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_DHT_GetIdx_return proxyholochain_DHT_GetIdx(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_DHT_GetIdx_return r;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DHT_GetIdx, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_DHT_HandleGossipPuts(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_DHT_HandleGossipPuts(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DHT_HandleGossipPuts, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_DHT_HandleGossipWiths(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_DHT_HandleGossipWiths(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DHT_HandleGossipWiths, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_DHT_JSON(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_DHT_JSON_return proxyholochain_DHT_JSON(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_DHT_JSON_return r;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DHT_JSON, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_DHT_SetupDHT(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_DHT_SetupDHT(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DHT_SetupDHT, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_DHT_Start(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_DHT_Start(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DHT_Start, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_DHT_String(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_DHT_String(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DHT_String, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_DHT(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_DHT()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_DHT, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_DHTConfig_HashType_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_DHTConfig_HashType_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DHTConfig_HashType_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_DHTConfig_HashType_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_DHTConfig_HashType_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DHTConfig_HashType_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_DHTConfig_NeighborhoodSize_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_DHTConfig_NeighborhoodSize_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DHTConfig_NeighborhoodSize_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_DHTConfig_NeighborhoodSize_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_DHTConfig_NeighborhoodSize_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DHTConfig_NeighborhoodSize_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_DHTConfig(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_DHTConfig()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_DHTConfig, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_DNA_Version_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_DNA_Version_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DNA_Version_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_DNA_Version_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_DNA_Version_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DNA_Version_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_DNA_Name_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_DNA_Name_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DNA_Name_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_DNA_Name_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_DNA_Name_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DNA_Name_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_DNA_PropertiesSchema_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_DNA_PropertiesSchema_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DNA_PropertiesSchema_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_DNA_PropertiesSchema_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_DNA_PropertiesSchema_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DNA_PropertiesSchema_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_DNA_AgentIdentitySchema_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_DNA_AgentIdentitySchema_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DNA_AgentIdentitySchema_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_DNA_AgentIdentitySchema_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_DNA_AgentIdentitySchema_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DNA_AgentIdentitySchema_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_DNA_RequiresVersion_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_DNA_RequiresVersion_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DNA_RequiresVersion_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_DNA_RequiresVersion_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_DNA_RequiresVersion_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DNA_RequiresVersion_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_DNA_NewUUID(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_DNA_NewUUID(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DNA_NewUUID, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_DNA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_DNA()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_DNA, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_DNAFile_Version_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_DNAFile_Version_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DNAFile_Version_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_DNAFile_Version_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_DNAFile_Version_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DNAFile_Version_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_DNAFile_Name_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_DNAFile_Name_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DNAFile_Name_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_DNAFile_Name_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_DNAFile_Name_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DNAFile_Name_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_DNAFile_PropertiesSchemaFile_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_DNAFile_PropertiesSchemaFile_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DNAFile_PropertiesSchemaFile_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_DNAFile_PropertiesSchemaFile_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_DNAFile_PropertiesSchemaFile_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DNAFile_PropertiesSchemaFile_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_DNAFile_RequiresVersion_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_DNAFile_RequiresVersion_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DNAFile_RequiresVersion_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_DNAFile_RequiresVersion_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_DNAFile_RequiresVersion_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DNAFile_RequiresVersion_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_DNAFile(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_DNAFile()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_DNAFile, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_DelEntry_Message_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_DelEntry_Message_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DelEntry_Message_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_DelEntry_Message_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_DelEntry_Message_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_DelEntry_Message_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_DelEntry(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_DelEntry()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_DelEntry, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_DelReq(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_DelReq()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_DelReq, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_EntryDef_Name_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_EntryDef_Name_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_EntryDef_Name_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_EntryDef_Name_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_EntryDef_Name_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_EntryDef_Name_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_EntryDef_DataFormat_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_EntryDef_DataFormat_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_EntryDef_DataFormat_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_EntryDef_DataFormat_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_EntryDef_DataFormat_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_EntryDef_DataFormat_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_EntryDef_Sharing_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_EntryDef_Sharing_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_EntryDef_Sharing_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_EntryDef_Sharing_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_EntryDef_Sharing_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_EntryDef_Sharing_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_EntryDef_Schema_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_EntryDef_Schema_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_EntryDef_Schema_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_EntryDef_Schema_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_EntryDef_Schema_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_EntryDef_Schema_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_EntryDef_BuildJSONSchemaValidator(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_EntryDef_BuildJSONSchemaValidator(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_EntryDef_BuildJSONSchemaValidator, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_EntryDef_BuildJSONSchemaValidatorFromString(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_EntryDef_BuildJSONSchemaValidatorFromString(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_EntryDef_BuildJSONSchemaValidatorFromString, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_EntryDef_IsSysEntry(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_EntryDef_IsSysEntry(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_EntryDef_IsSysEntry, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_EntryDef_IsVirtualEntry(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_EntryDef_IsVirtualEntry(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_EntryDef_IsVirtualEntry, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_EntryDef(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_EntryDef()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_EntryDef, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Name_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_EntryDefFile_Name_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_EntryDefFile_Name_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Name_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_EntryDefFile_Name_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_EntryDefFile_Name_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_EntryDefFile_DataFormat_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_EntryDefFile_DataFormat_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_EntryDefFile_DataFormat_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_EntryDefFile_DataFormat_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_EntryDefFile_DataFormat_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_EntryDefFile_DataFormat_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Schema_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_EntryDefFile_Schema_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_EntryDefFile_Schema_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Schema_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_EntryDefFile_Schema_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_EntryDefFile_Schema_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_EntryDefFile_SchemaFile_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_EntryDefFile_SchemaFile_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_EntryDefFile_SchemaFile_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_EntryDefFile_SchemaFile_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_EntryDefFile_SchemaFile_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_EntryDefFile_SchemaFile_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Sharing_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_EntryDefFile_Sharing_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_EntryDefFile_Sharing_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_EntryDefFile_Sharing_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_EntryDefFile_Sharing_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_EntryDefFile_Sharing_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_EntryDefFile(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_EntryDefFile()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_EntryDefFile, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ErrorResponse_Code_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ErrorResponse_Code_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ErrorResponse_Code_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_ErrorResponse_Code_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_ErrorResponse_Code_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ErrorResponse_Code_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ErrorResponse_Message_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ErrorResponse_Message_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ErrorResponse_Message_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_ErrorResponse_Message_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ErrorResponse_Message_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ErrorResponse_Message_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ErrorResponse_DecodeResponseError(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_ErrorResponse_DecodeResponseError(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ErrorResponse_DecodeResponseError, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ErrorResponse(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ErrorResponse()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ErrorResponse, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_FindNodeReq(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_FindNodeReq()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_FindNodeReq, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_FunctionDef_Name_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_FunctionDef_Name_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_FunctionDef_Name_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_FunctionDef_Name_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_FunctionDef_Name_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_FunctionDef_Name_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_FunctionDef_CallingType_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_FunctionDef_CallingType_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_FunctionDef_CallingType_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_FunctionDef_CallingType_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_FunctionDef_CallingType_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_FunctionDef_CallingType_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_FunctionDef_Exposure_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_FunctionDef_Exposure_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_FunctionDef_Exposure_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_FunctionDef_Exposure_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_FunctionDef_Exposure_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_FunctionDef_Exposure_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_FunctionDef_ValidExposure(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_FunctionDef_ValidExposure(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_FunctionDef_ValidExposure, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_FunctionDef(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_FunctionDef()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_FunctionDef, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_GetLinksOptions_Load_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_GetLinksOptions_Load_Set(int32_t p0, char p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char p1;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GetLinksOptions_Load_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_GetLinksOptions_Load_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_GetLinksOptions_Load_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GetLinksOptions_Load_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_GetLinksOptions_StatusMask_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_GetLinksOptions_StatusMask_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GetLinksOptions_StatusMask_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_GetLinksOptions_StatusMask_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_GetLinksOptions_StatusMask_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GetLinksOptions_StatusMask_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_GetLinksOptions(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_GetLinksOptions()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_GetLinksOptions, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_GetOptions_StatusMask_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_GetOptions_StatusMask_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GetOptions_StatusMask_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_GetOptions_StatusMask_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_GetOptions_StatusMask_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GetOptions_StatusMask_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_GetOptions_GetMask_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_GetOptions_GetMask_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GetOptions_GetMask_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_GetOptions_GetMask_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_GetOptions_GetMask_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GetOptions_GetMask_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_GetOptions_Local_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_GetOptions_Local_Set(int32_t p0, char p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char p1;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GetOptions_Local_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_GetOptions_Local_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_GetOptions_Local_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GetOptions_Local_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_GetOptions(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_GetOptions()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_GetOptions, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_GetReq_StatusMask_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_GetReq_StatusMask_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GetReq_StatusMask_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_GetReq_StatusMask_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_GetReq_StatusMask_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GetReq_StatusMask_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_GetReq_GetMask_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_GetReq_GetMask_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GetReq_GetMask_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_GetReq_GetMask_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_GetReq_GetMask_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GetReq_GetMask_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_GetReq(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_GetReq()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_GetReq, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_GetResp_EntryType_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_GetResp_EntryType_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GetResp_EntryType_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_GetResp_EntryType_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_GetResp_EntryType_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GetResp_EntryType_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_GetResp_FollowHash_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_GetResp_FollowHash_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GetResp_FollowHash_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_GetResp_FollowHash_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_GetResp_FollowHash_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GetResp_FollowHash_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_GetResp(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_GetResp()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_GetResp, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_GobEntry_Marshal(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_GobEntry_Marshal_return proxyholochain_GobEntry_Marshal(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_GobEntry_Marshal_return r;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GobEntry_Marshal, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_GobEntry_Unmarshal(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_GobEntry_Unmarshal(int32_t p0, nbyteslice p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GobEntry_Unmarshal, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_GobEntry(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_GobEntry()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_GobEntry, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_Gossip(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Gossip()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Gossip, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_GossipReq_MyIdx_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_GossipReq_MyIdx_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GossipReq_MyIdx_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_GossipReq_MyIdx_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_GossipReq_MyIdx_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GossipReq_MyIdx_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_GossipReq_YourIdx_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_GossipReq_YourIdx_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GossipReq_YourIdx_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_GossipReq_YourIdx_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_GossipReq_YourIdx_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_GossipReq_YourIdx_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_GossipReq(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_GossipReq()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_GossipReq, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Header_Type_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Header_Type_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Header_Type_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Header_Type_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Header_Type_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Header_Type_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Header_Marshal(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_Header_Marshal_return proxyholochain_Header_Marshal(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_Header_Marshal_return r;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Header_Marshal, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_Header_Unmarshal(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Header_Unmarshal(int32_t p0, nbyteslice p1, nint p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice p1;
		nint p2;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Header_Unmarshal, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_Header(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Header()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Header, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_Activate(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Holochain_Activate(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_Activate, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_Agent(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Holochain_Agent(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_Agent, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_BSget(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Holochain_BSget(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_BSget, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_BSpost(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Holochain_BSpost(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_BSpost, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_BuildBridge(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Holochain_BuildBridge(int32_t p0, int32_t p1, nstring p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
		nstring p2;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_BuildBridge, &a, 20, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_Chain(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Holochain_Chain(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_Chain, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_Close(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Holochain_Close(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_Close, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_DBPath(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Holochain_DBPath(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_DBPath, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_DHT(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Holochain_DHT(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_DHT, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_DNAPath(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Holochain_DNAPath(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_DNAPath, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_Debug(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Holochain_Debug(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_Debug, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_GetEntryDef(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_Holochain_GetEntryDef_return proxyholochain_Holochain_GetEntryDef(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_Holochain_GetEntryDef_return r;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_GetEntryDef, &a, 20, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_GetProperty(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_Holochain_GetProperty_return proxyholochain_Holochain_GetProperty(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
		nstring r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_Holochain_GetProperty_return r;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_GetProperty, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_GetZome(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_Holochain_GetZome_return proxyholochain_Holochain_GetZome(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_Holochain_GetZome_return r;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_GetZome, &a, 20, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_GetZomeForEntryType(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_Holochain_GetZomeForEntryType_return proxyholochain_Holochain_GetZomeForEntryType(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_Holochain_GetZomeForEntryType_return r;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_GetZomeForEntryType, &a, 20, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_HandleAsyncSends(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Holochain_HandleAsyncSends(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_HandleAsyncSends, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_MakeRibosome(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_Holochain_MakeRibosome_return proxyholochain_Holochain_MakeRibosome(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_Holochain_MakeRibosome_return r;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_MakeRibosome, &a, 20, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_Name(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Holochain_Name(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_Name, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_Node(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Holochain_Node(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_Node, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_NodeIDStr(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Holochain_NodeIDStr(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_NodeIDStr, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_Nucleus(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Holochain_Nucleus(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_Nucleus, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_Prepare(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Holochain_Prepare(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_Prepare, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_PrepareHashType(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Holochain_PrepareHashType(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_PrepareHashType, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_Reset(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Holochain_Reset(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_Reset, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_RootPath(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Holochain_RootPath(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_RootPath, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_Sign(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_Holochain_Sign_return proxyholochain_Holochain_Sign(int32_t p0, nbyteslice p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice p1;
		nbyteslice r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_Holochain_Sign_return r;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_Sign, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_StartBackgroundTasks(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Holochain_StartBackgroundTasks(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_StartBackgroundTasks, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_Started(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_Holochain_Started(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_Started, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_TestPath(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Holochain_TestPath(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_TestPath, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_UIPath(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Holochain_UIPath(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_UIPath, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Holochain_ZomePath(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Holochain_ZomePath(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Holochain_ZomePath, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_Holochain(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Holochain()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Holochain, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_JSONEntry_Marshal(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_JSONEntry_Marshal_return proxyholochain_JSONEntry_Marshal(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_JSONEntry_Marshal_return r;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_JSONEntry_Marshal, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_JSONEntry_Unmarshal(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_JSONEntry_Unmarshal(int32_t p0, nbyteslice p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_JSONEntry_Unmarshal, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_JSONEntry(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_JSONEntry()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_JSONEntry, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_JSONSchemaValidator(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_JSONSchemaValidator()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_JSONSchemaValidator, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_JSRibosome_ChainGenesis(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_JSRibosome_ChainGenesis(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_JSRibosome_ChainGenesis, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_JSRibosome_Receive(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_JSRibosome_Receive_return proxyholochain_JSRibosome_Receive(int32_t p0, nstring p1, nstring p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
		nstring p2;
		nstring r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_JSRibosome_Receive_return r;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_JSRibosome_Receive, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_JSRibosome_Type(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_JSRibosome_Type(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_JSRibosome_Type, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_JSRibosome(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_JSRibosome()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_JSRibosome, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_LibP2PAgent(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_LibP2PAgent()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_LibP2PAgent, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Link_LinkAction_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Link_LinkAction_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Link_LinkAction_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Link_LinkAction_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Link_LinkAction_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Link_LinkAction_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Link_Base_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Link_Base_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Link_Base_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Link_Base_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Link_Base_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Link_Base_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Link_Link_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Link_Link_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Link_Link_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Link_Link_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Link_Link_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Link_Link_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Link_Tag_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Link_Tag_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Link_Tag_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Link_Tag_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Link_Tag_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Link_Tag_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_Link(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Link()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Link, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_LinkEvent_Status_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_LinkEvent_Status_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_LinkEvent_Status_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_LinkEvent_Status_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_LinkEvent_Status_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_LinkEvent_Status_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_LinkEvent_Source_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_LinkEvent_Source_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_LinkEvent_Source_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_LinkEvent_Source_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_LinkEvent_Source_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_LinkEvent_Source_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_LinkEvent_LinksEntry_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_LinkEvent_LinksEntry_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_LinkEvent_LinksEntry_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_LinkEvent_LinksEntry_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_LinkEvent_LinksEntry_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_LinkEvent_LinksEntry_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_LinkEvent(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_LinkEvent()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_LinkEvent, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_LinkQuery_T_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_LinkQuery_T_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_LinkQuery_T_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_LinkQuery_T_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_LinkQuery_T_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_LinkQuery_T_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_LinkQuery_StatusMask_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_LinkQuery_StatusMask_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_LinkQuery_StatusMask_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_LinkQuery_StatusMask_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_LinkQuery_StatusMask_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_LinkQuery_StatusMask_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_LinkQuery(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_LinkQuery()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_LinkQuery, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_LinkQueryResp(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_LinkQueryResp()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_LinkQueryResp, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_LinkReq(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_LinkReq()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_LinkReq, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_LinksEntry(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_LinksEntry()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_LinksEntry, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ListAddReq_ListType_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ListAddReq_ListType_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ListAddReq_ListType_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_ListAddReq_ListType_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ListAddReq_ListType_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ListAddReq_ListType_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ListAddReq_WarrantType_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ListAddReq_WarrantType_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ListAddReq_WarrantType_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_ListAddReq_WarrantType_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_ListAddReq_WarrantType_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ListAddReq_WarrantType_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ListAddReq_Warrant_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ListAddReq_Warrant_Set(int32_t p0, nbyteslice p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ListAddReq_Warrant_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_ListAddReq_Warrant_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nbyteslice proxyholochain_ListAddReq_Warrant_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ListAddReq_Warrant_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ListAddReq(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ListAddReq()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ListAddReq, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Logger_Name_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Logger_Name_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Logger_Name_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Logger_Name_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Logger_Name_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Logger_Name_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Logger_Enabled_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Logger_Enabled_Set(int32_t p0, char p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char p1;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Logger_Enabled_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Logger_Enabled_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_Logger_Enabled_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Logger_Enabled_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Logger_Format_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Logger_Format_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Logger_Format_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Logger_Format_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Logger_Format_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Logger_Format_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Logger_Prefix_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Logger_Prefix_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Logger_Prefix_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Logger_Prefix_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Logger_Prefix_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Logger_Prefix_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Logger_SetPrefix(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Logger_SetPrefix(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Logger_SetPrefix, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_new_holochain_Logger(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Logger()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Logger, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_Loggers(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Loggers()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Loggers, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Message_Encode(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_Message_Encode_return proxyholochain_Message_Encode(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_Message_Encode_return r;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Message_Encode, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_Message_String(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Message_String(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Message_String, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_Message(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Message()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Message, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Meta_T_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Meta_T_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Meta_T_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Meta_T_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Meta_T_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Meta_T_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Meta_V_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Meta_V_Set(int32_t p0, nbyteslice p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Meta_V_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Meta_V_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nbyteslice proxyholochain_Meta_V_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Meta_V_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_Meta(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Meta()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Meta, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Identity_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ModAgentOptions_Identity_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Identity_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Identity_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ModAgentOptions_Identity_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Identity_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Revocation_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ModAgentOptions_Revocation_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Revocation_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Revocation_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ModAgentOptions_Revocation_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ModAgentOptions_Revocation_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ModAgentOptions(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ModAgentOptions()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ModAgentOptions, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ModReq(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ModReq()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ModReq, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Node_Close(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Node_Close(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Node_Close, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Node_StartProtocol(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Node_StartProtocol(int32_t p0, int32_t p1, nint p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
		nint p2;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Node_StartProtocol, &a, 20, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_Node(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Node()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Node, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Nucleus_DNA(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Nucleus_DNA(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Nucleus_DNA, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Nucleus_RunGenesis(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Nucleus_RunGenesis(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Nucleus_RunGenesis, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Nucleus_Start(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Nucleus_Start(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Nucleus_Start, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_Nucleus(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Nucleus()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Nucleus, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Package_Chain_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Package_Chain_Set(int32_t p0, nbyteslice p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Package_Chain_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Package_Chain_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nbyteslice proxyholochain_Package_Chain_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Package_Chain_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_Package(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Package()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Package, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_PeerInfo_ID_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_PeerInfo_ID_Set(int32_t p0, nbyteslice p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_PeerInfo_ID_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_PeerInfo_ID_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nbyteslice proxyholochain_PeerInfo_ID_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_PeerInfo_ID_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_PeerInfo(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_PeerInfo()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_PeerInfo, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_PeerList(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_PeerList()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_PeerList, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_PeerRecord_Warrant_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_PeerRecord_Warrant_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_PeerRecord_Warrant_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_PeerRecord_Warrant_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_PeerRecord_Warrant_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_PeerRecord_Warrant_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_PeerRecord(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_PeerRecord()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_PeerRecord, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Progenitor_Identity_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Progenitor_Identity_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Progenitor_Identity_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Progenitor_Identity_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Progenitor_Identity_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Progenitor_Identity_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Progenitor_PubKey_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Progenitor_PubKey_Set(int32_t p0, nbyteslice p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Progenitor_PubKey_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Progenitor_PubKey_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nbyteslice proxyholochain_Progenitor_PubKey_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Progenitor_PubKey_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_Progenitor(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Progenitor()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Progenitor, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_Protocol(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Protocol()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Protocol, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Put_Idx_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Put_Idx_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Put_Idx_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Put_Idx_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_Put_Idx_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Put_Idx_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_Put(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Put()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Put, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_PutReq_S_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_PutReq_S_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_PutReq_S_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_PutReq_S_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_PutReq_S_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_PutReq_S_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_PutReq(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_PutReq()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_PutReq, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Contains_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_QueryConstrain_Contains_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_QueryConstrain_Contains_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Contains_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_QueryConstrain_Contains_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_QueryConstrain_Contains_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Equals_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_QueryConstrain_Equals_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_QueryConstrain_Equals_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Equals_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_QueryConstrain_Equals_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_QueryConstrain_Equals_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Matches_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_QueryConstrain_Matches_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_QueryConstrain_Matches_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Matches_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_QueryConstrain_Matches_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_QueryConstrain_Matches_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Count_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_QueryConstrain_Count_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_QueryConstrain_Count_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Count_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_QueryConstrain_Count_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_QueryConstrain_Count_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Page_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_QueryConstrain_Page_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_QueryConstrain_Page_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_QueryConstrain_Page_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_QueryConstrain_Page_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_QueryConstrain_Page_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_QueryConstrain(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_QueryConstrain()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_QueryConstrain, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_QueryOptions(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_QueryOptions()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_QueryOptions, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_QueryOrder_Ascending_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_QueryOrder_Ascending_Set(int32_t p0, char p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char p1;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_QueryOrder_Ascending_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_QueryOrder_Ascending_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_QueryOrder_Ascending_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_QueryOrder_Ascending_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_QueryOrder(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_QueryOrder()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_QueryOrder, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_QueryResult_Header_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_QueryResult_Header_Set(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_QueryResult_Header_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_QueryResult_Header_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_QueryResult_Header_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_QueryResult_Header_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_QueryResult_Entry_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_QueryResult_Entry_Set(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_QueryResult_Entry_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_QueryResult_Entry_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_QueryResult_Entry_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_QueryResult_Entry_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_QueryResult(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_QueryResult()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_QueryResult, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_QueryReturn_Hashes_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_QueryReturn_Hashes_Set(int32_t p0, char p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char p1;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_QueryReturn_Hashes_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_QueryReturn_Hashes_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_QueryReturn_Hashes_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_QueryReturn_Hashes_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_QueryReturn_Entries_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_QueryReturn_Entries_Set(int32_t p0, char p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char p1;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_QueryReturn_Entries_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_QueryReturn_Entries_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_QueryReturn_Entries_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_QueryReturn_Entries_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_QueryReturn_Headers_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_QueryReturn_Headers_Set(int32_t p0, char p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char p1;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_QueryReturn_Headers_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_QueryReturn_Headers_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_QueryReturn_Headers_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_QueryReturn_Headers_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_QueryReturn(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_QueryReturn()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_QueryReturn, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_RoutingTable_IsEmpty(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_RoutingTable_IsEmpty(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_RoutingTable_IsEmpty, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_RoutingTable_Print(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_RoutingTable_Print(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_RoutingTable_Print, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_RoutingTable_Size(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_RoutingTable_Size(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_RoutingTable_Size, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_RoutingTable(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_RoutingTable()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_RoutingTable, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Data_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_SelfRevocation_Data_Set(int32_t p0, nbyteslice p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_SelfRevocation_Data_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Data_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nbyteslice proxyholochain_SelfRevocation_Data_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_SelfRevocation_Data_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_SelfRevocation_OldSig_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_SelfRevocation_OldSig_Set(int32_t p0, nbyteslice p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_SelfRevocation_OldSig_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_SelfRevocation_OldSig_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nbyteslice proxyholochain_SelfRevocation_OldSig_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_SelfRevocation_OldSig_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_SelfRevocation_NewSig_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_SelfRevocation_NewSig_Set(int32_t p0, nbyteslice p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_SelfRevocation_NewSig_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_SelfRevocation_NewSig_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nbyteslice proxyholochain_SelfRevocation_NewSig_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_SelfRevocation_NewSig_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Marshal(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_SelfRevocation_Marshal_return proxyholochain_SelfRevocation_Marshal(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_SelfRevocation_Marshal_return r;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_SelfRevocation_Marshal, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Unmarshal(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_SelfRevocation_Unmarshal(int32_t p0, nbyteslice p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_SelfRevocation_Unmarshal, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_SelfRevocation_Verify(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_SelfRevocation_Verify(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_SelfRevocation_Verify, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_SelfRevocation(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_SelfRevocation()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_SelfRevocation, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Decode(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_SelfRevocationWarrant_Decode(int32_t p0, nbyteslice p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Decode, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Encode(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_SelfRevocationWarrant_Encode_return proxyholochain_SelfRevocationWarrant_Encode(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_SelfRevocationWarrant_Encode_return r;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Encode, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Type(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_SelfRevocationWarrant_Type(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Type, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Verify(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_SelfRevocationWarrant_Verify(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_SelfRevocationWarrant_Verify, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_SelfRevocationWarrant(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_SelfRevocationWarrant()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_SelfRevocationWarrant, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_SendOptions_Callback_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_SendOptions_Callback_Set(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_SendOptions_Callback_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_SendOptions_Callback_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_SendOptions_Callback_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_SendOptions_Callback_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_SendOptions_Timeout_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_SendOptions_Timeout_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_SendOptions_Timeout_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_SendOptions_Timeout_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_SendOptions_Timeout_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_SendOptions_Timeout_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_SendOptions(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_SendOptions()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_SendOptions, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Service_DefaultAgent_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Service_DefaultAgent_Set(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Service_DefaultAgent_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Service_DefaultAgent_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Service_DefaultAgent_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Service_DefaultAgent_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Service_Path_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Service_Path_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Service_Path_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Service_Path_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Service_Path_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Service_Path_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Service_Clone(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_Service_Clone_return proxyholochain_Service_Clone(int32_t p0, nstring p1, nstring p2, int32_t p3, char p4, char p5)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
		nstring p2;
		int32_t p3;
		char p4;
		char p5;
		char __pad0[2];
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_Service_Clone_return r;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	a.p5 = p5;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Service_Clone, &a, 36, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_Service_GenChain(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_Service_GenChain_return proxyholochain_Service_GenChain(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_Service_GenChain_return r;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Service_GenChain, &a, 20, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_Service_InitAppDir(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Service_InitAppDir(int32_t p0, nstring p1, nstring p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
		nstring p2;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Service_InitAppDir, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Service_IsConfigured(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_Service_IsConfigured_return proxyholochain_Service_IsConfigured(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
		nstring r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_Service_IsConfigured_return r;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Service_IsConfigured, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_Service_ListChains(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Service_ListChains(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Service_ListChains, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Service_Load(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_Service_Load_return proxyholochain_Service_Load(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_Service_Load_return r;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Service_Load, &a, 20, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_Service_MakeAppPackage(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_Service_MakeAppPackage_return proxyholochain_Service_MakeAppPackage(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
		nbyteslice r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_Service_MakeAppPackage_return r;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Service_MakeAppPackage, &a, 20, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_Service_MakeTestingApp(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_Service_MakeTestingApp_return proxyholochain_Service_MakeTestingApp(int32_t p0, nstring p1, nstring p2, char p3, char p4, int32_t p5)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
		nstring p2;
		char p3;
		char p4;
		char __pad0[2];
		int32_t p5;
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_Service_MakeTestingApp_return r;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	a.p5 = p5;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Service_MakeTestingApp, &a, 36, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_new_holochain_Service(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Service()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Service, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Set(int32_t p0, char p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char p1;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeAuthor_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Set(int32_t p0, char p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char p1;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultPeerModeDHTNode_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultBootstrapServer_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ServiceConfig_DefaultBootstrapServer_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultBootstrapServer_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultBootstrapServer_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ServiceConfig_DefaultBootstrapServer_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultBootstrapServer_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableMDNS_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ServiceConfig_DefaultEnableMDNS_Set(int32_t p0, char p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char p1;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableMDNS_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableMDNS_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_ServiceConfig_DefaultEnableMDNS_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableMDNS_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Set(int32_t p0, char p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char p1;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ServiceConfig_DefaultEnableNATUPnP_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ServiceConfig_Validate(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_ServiceConfig_Validate(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ServiceConfig_Validate, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ServiceConfig(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ServiceConfig()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ServiceConfig, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Signature_S_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Signature_S_Set(int32_t p0, nbyteslice p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Signature_S_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Signature_S_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nbyteslice proxyholochain_Signature_S_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Signature_S_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_Signature(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Signature()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Signature, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_StatusChange_Action_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_StatusChange_Action_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_StatusChange_Action_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_StatusChange_Action_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_StatusChange_Action_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_StatusChange_Action_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_StatusChange(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_StatusChange()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_StatusChange, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_TaggedHash_H_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_TaggedHash_H_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TaggedHash_H_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_TaggedHash_H_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_TaggedHash_H_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TaggedHash_H_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_TaggedHash_E_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_TaggedHash_E_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TaggedHash_E_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_TaggedHash_E_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_TaggedHash_E_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TaggedHash_E_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_TaggedHash_EntryType_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_TaggedHash_EntryType_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TaggedHash_EntryType_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_TaggedHash_EntryType_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_TaggedHash_EntryType_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TaggedHash_EntryType_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_TaggedHash_T_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_TaggedHash_T_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TaggedHash_T_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_TaggedHash_T_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_TaggedHash_T_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TaggedHash_T_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_TaggedHash_Source_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_TaggedHash_Source_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TaggedHash_Source_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_TaggedHash_Source_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_TaggedHash_Source_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TaggedHash_Source_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_TaggedHash(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_TaggedHash()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_TaggedHash, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_TestConfig_GossipInterval_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_TestConfig_GossipInterval_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestConfig_GossipInterval_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_TestConfig_GossipInterval_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_TestConfig_GossipInterval_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestConfig_GossipInterval_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_TestConfig_Duration_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_TestConfig_Duration_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestConfig_Duration_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_TestConfig_Duration_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_TestConfig_Duration_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestConfig_Duration_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_TestConfig(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_TestConfig()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_TestConfig, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_TestData_Convey_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_TestData_Convey_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestData_Convey_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_TestData_Convey_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_TestData_Convey_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestData_Convey_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_TestData_Zome_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_TestData_Zome_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestData_Zome_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_TestData_Zome_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_TestData_Zome_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestData_Zome_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_TestData_FnName_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_TestData_FnName_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestData_FnName_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_TestData_FnName_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_TestData_FnName_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestData_FnName_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_TestData_Regexp_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_TestData_Regexp_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestData_Regexp_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_TestData_Regexp_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_TestData_Regexp_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestData_Regexp_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_TestData_Exposure_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_TestData_Exposure_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestData_Exposure_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_TestData_Exposure_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_TestData_Exposure_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestData_Exposure_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_TestData_Raw_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_TestData_Raw_Set(int32_t p0, char p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char p1;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestData_Raw_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_TestData_Raw_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_TestData_Raw_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestData_Raw_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_TestData_Repeat_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_TestData_Repeat_Set(int32_t p0, nint p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestData_Repeat_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_TestData_Repeat_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_TestData_Repeat_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestData_Repeat_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_TestData_Benchmark_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_TestData_Benchmark_Set(int32_t p0, char p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char p1;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestData_Benchmark_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_TestData_Benchmark_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_TestData_Benchmark_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestData_Benchmark_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_TestData(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_TestData()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_TestData, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_TestFixtures(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_TestFixtures()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_TestFixtures, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_TestSet_Identity_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_TestSet_Identity_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestSet_Identity_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_TestSet_Identity_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_TestSet_Identity_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestSet_Identity_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_TestSet_Benchmark_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_TestSet_Benchmark_Set(int32_t p0, char p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char p1;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestSet_Benchmark_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_TestSet_Benchmark_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain_TestSet_Benchmark_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_TestSet_Benchmark_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_TestSet(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_TestSet()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_TestSet, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ValidateQuery(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ValidateQuery()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ValidateQuery, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ValidateResponse_Type_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ValidateResponse_Type_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ValidateResponse_Type_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_ValidateResponse_Type_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ValidateResponse_Type_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ValidateResponse_Type_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ValidateResponse(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ValidateResponse()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ValidateResponse, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ValidationPackage_Chain_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ValidationPackage_Chain_Set(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ValidationPackage_Chain_Set, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_ValidationPackage_Chain_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_ValidationPackage_Chain_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ValidationPackage_Chain_Get, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ValidationPackage(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ValidationPackage()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ValidationPackage, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Zome_Name_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Zome_Name_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Zome_Name_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Zome_Name_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Zome_Name_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Zome_Name_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Zome_Description_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Zome_Description_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Zome_Description_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Zome_Description_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Zome_Description_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Zome_Description_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Zome_Code_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Zome_Code_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Zome_Code_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Zome_Code_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Zome_Code_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Zome_Code_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Zome_RibosomeType_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_Zome_RibosomeType_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Zome_RibosomeType_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Zome_RibosomeType_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Zome_RibosomeType_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Zome_RibosomeType_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Zome_CodeFileName(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Zome_CodeFileName(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Zome_CodeFileName, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Zome_GetEntryDef(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_Zome_GetEntryDef_return proxyholochain_Zome_GetEntryDef(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_Zome_GetEntryDef_return r;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Zome_GetEntryDef, &a, 20, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_Zome_GetFunctionDef(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_Zome_GetFunctionDef_return proxyholochain_Zome_GetFunctionDef(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_Zome_GetFunctionDef_return r;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Zome_GetFunctionDef, &a, 20, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_Zome_MakeRibosome(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_Zome_MakeRibosome_return proxyholochain_Zome_MakeRibosome(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_Zome_MakeRibosome_return r;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Zome_MakeRibosome, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_new_holochain_Zome(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_Zome()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_Zome, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ZomeFile_Name_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ZomeFile_Name_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ZomeFile_Name_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_ZomeFile_Name_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ZomeFile_Name_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ZomeFile_Name_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ZomeFile_Description_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ZomeFile_Description_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ZomeFile_Description_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_ZomeFile_Description_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ZomeFile_Description_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ZomeFile_Description_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ZomeFile_CodeFile_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ZomeFile_CodeFile_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ZomeFile_CodeFile_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_ZomeFile_CodeFile_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ZomeFile_CodeFile_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ZomeFile_CodeFile_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ZomeFile_RibosomeType_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ZomeFile_RibosomeType_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ZomeFile_RibosomeType_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_ZomeFile_RibosomeType_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ZomeFile_RibosomeType_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ZomeFile_RibosomeType_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ZomeFile_BridgeTo_Set(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_ZomeFile_BridgeTo_Set(int32_t p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ZomeFile_BridgeTo_Set, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_ZomeFile_BridgeTo_Get(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ZomeFile_BridgeTo_Get(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ZomeFile_BridgeTo_Get, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ZomeFile(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ZomeFile()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ZomeFile, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ZygoRibosome_ChainGenesis(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_ZygoRibosome_ChainGenesis(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ZygoRibosome_ChainGenesis, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ZygoRibosome_Receive(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_ZygoRibosome_Receive_return proxyholochain_ZygoRibosome_Receive(int32_t p0, nstring p1, nstring p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
		nstring p2;
		nstring r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_ZygoRibosome_Receive_return r;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ZygoRibosome_Receive, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_ZygoRibosome_Type(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ZygoRibosome_Type(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ZygoRibosome_Type, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_new_holochain_ZygoRibosome(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t new_holochain_ZygoRibosome()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_new_holochain_ZygoRibosome, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Action_Name(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Action_Name(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Action_Name, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_CommittingAction_CheckValidationRequest(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_CommittingAction_CheckValidationRequest(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_CommittingAction_CheckValidationRequest, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_CommittingAction_Entry(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_CommittingAction_Entry(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_CommittingAction_Entry, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_CommittingAction_EntryType(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_CommittingAction_EntryType(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_CommittingAction_EntryType, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_CommittingAction_Name(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_CommittingAction_Name(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_CommittingAction_Name, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_CommittingAction_SetHeader(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain_CommittingAction_SetHeader(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_CommittingAction_SetHeader, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain_Entry_Marshal(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_Entry_Marshal_return proxyholochain_Entry_Marshal(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_Entry_Marshal_return r;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Entry_Marshal, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_Entry_Unmarshal(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Entry_Unmarshal(int32_t p0, nbyteslice p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Entry_Unmarshal, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Revocation_Marshal(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_Revocation_Marshal_return proxyholochain_Revocation_Marshal(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_Revocation_Marshal_return r;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Revocation_Marshal, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_Revocation_Unmarshal(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Revocation_Unmarshal(int32_t p0, nbyteslice p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Revocation_Unmarshal, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Revocation_Verify(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Revocation_Verify(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Revocation_Verify, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Ribosome_ChainGenesis(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Ribosome_ChainGenesis(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Ribosome_ChainGenesis, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Ribosome_Receive(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_Ribosome_Receive_return proxyholochain_Ribosome_Receive(int32_t p0, nstring p1, nstring p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring p1;
		nstring p2;
		nstring r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_Ribosome_Receive_return r;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Ribosome_Receive, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_Ribosome_Type(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_Ribosome_Type(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Ribosome_Type, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ValidatingAction_CheckValidationRequest(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_ValidatingAction_CheckValidationRequest(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ValidatingAction_CheckValidationRequest, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_ValidatingAction_Name(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain_ValidatingAction_Name(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_ValidatingAction_Name, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Warrant_Decode(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Warrant_Decode(int32_t p0, nbyteslice p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Warrant_Decode, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Warrant_Encode(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain_Warrant_Encode_return proxyholochain_Warrant_Encode(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nbyteslice r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain_Warrant_Encode_return r;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Warrant_Encode, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain_Warrant_Type(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint proxyholochain_Warrant_Type(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nint r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Warrant_Type, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain_Warrant_Verify(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain_Warrant_Verify(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain_Warrant_Verify, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_AgentEntryDef(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_AgentEntryDef(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_AgentEntryDef, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_AgentEntryDef(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_AgentEntryDef()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_AgentEntryDef, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_AlphaValue(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_AlphaValue(nint p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nint p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_AlphaValue, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_AlphaValue(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint var_getholochain_AlphaValue()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nint r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_AlphaValue, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_BasicTemplateAppPackage(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_BasicTemplateAppPackage(nstring p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_BasicTemplateAppPackage, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_BasicTemplateAppPackage(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring var_getholochain_BasicTemplateAppPackage()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_BasicTemplateAppPackage, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_BridgeAppNotFoundErr(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_BridgeAppNotFoundErr(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_BridgeAppNotFoundErr, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_BridgeAppNotFoundErr(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_BridgeAppNotFoundErr()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_BridgeAppNotFoundErr, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_CapabilityInvalidErr(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_CapabilityInvalidErr(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_CapabilityInvalidErr, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_CapabilityInvalidErr(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_CapabilityInvalidErr()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_CapabilityInvalidErr, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_CloserPeerCount(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_CloserPeerCount(nint p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nint p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_CloserPeerCount, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_CloserPeerCount(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint var_getholochain_CloserPeerCount()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nint r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_CloserPeerCount, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_DNAEntryDef(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_DNAEntryDef(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_DNAEntryDef, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_DNAEntryDef(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_DNAEntryDef()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_DNAEntryDef, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_EnableAllLoggersEnv(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_EnableAllLoggersEnv(nstring p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_EnableAllLoggersEnv, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_EnableAllLoggersEnv(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring var_getholochain_EnableAllLoggersEnv()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_EnableAllLoggersEnv, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_ErrBlockedListed(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_ErrBlockedListed(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_ErrBlockedListed, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_ErrBlockedListed(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_ErrBlockedListed()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_ErrBlockedListed, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_ErrDHTErrNoGossipersAvailable(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_ErrDHTErrNoGossipersAvailable(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_ErrDHTErrNoGossipersAvailable, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_ErrDHTErrNoGossipersAvailable(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_ErrDHTErrNoGossipersAvailable()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_ErrDHTErrNoGossipersAvailable, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_ErrDHTExpectedGossipReqInBody(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_ErrDHTExpectedGossipReqInBody(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_ErrDHTExpectedGossipReqInBody, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_ErrDHTExpectedGossipReqInBody(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_ErrDHTExpectedGossipReqInBody()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_ErrDHTExpectedGossipReqInBody, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_ErrDHTUnexpectedTypeInBody(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_ErrDHTUnexpectedTypeInBody(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_ErrDHTUnexpectedTypeInBody, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_ErrDHTUnexpectedTypeInBody(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_ErrDHTUnexpectedTypeInBody()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_ErrDHTUnexpectedTypeInBody, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_ErrEmptyRoutingTable(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_ErrEmptyRoutingTable(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_ErrEmptyRoutingTable, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_ErrEmptyRoutingTable(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_ErrEmptyRoutingTable()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_ErrEmptyRoutingTable, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_ErrEntryTypeMismatch(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_ErrEntryTypeMismatch(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_ErrEntryTypeMismatch, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_ErrEntryTypeMismatch(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_ErrEntryTypeMismatch()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_ErrEntryTypeMismatch, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_ErrHashDeleted(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_ErrHashDeleted(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_ErrHashDeleted, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_ErrHashDeleted(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_ErrHashDeleted()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_ErrHashDeleted, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_ErrHashModified(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_ErrHashModified(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_ErrHashModified, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_ErrHashModified(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_ErrHashModified()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_ErrHashModified, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_ErrHashNotFound(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_ErrHashNotFound(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_ErrHashNotFound, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_ErrHashNotFound(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_ErrHashNotFound()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_ErrHashNotFound, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_ErrHashRejected(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_ErrHashRejected(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_ErrHashRejected, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_ErrHashRejected(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_ErrHashRejected()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_ErrHashRejected, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_ErrIncompleteChain(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_ErrIncompleteChain(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_ErrIncompleteChain, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_ErrIncompleteChain(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_ErrIncompleteChain()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_ErrIncompleteChain, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_ErrLinkNotFound(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_ErrLinkNotFound(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_ErrLinkNotFound, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_ErrLinkNotFound(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_ErrLinkNotFound()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_ErrLinkNotFound, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_ErrNilEntryInvalid(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_ErrNilEntryInvalid(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_ErrNilEntryInvalid, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_ErrNilEntryInvalid(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_ErrNilEntryInvalid()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_ErrNilEntryInvalid, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_ErrNoSuchIdx(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_ErrNoSuchIdx(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_ErrNoSuchIdx, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_ErrNoSuchIdx(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_ErrNoSuchIdx()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_ErrNoSuchIdx, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_ErrNotValidForAgentType(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_ErrNotValidForAgentType(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_ErrNotValidForAgentType, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_ErrNotValidForAgentType(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_ErrNotValidForAgentType()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_ErrNotValidForAgentType, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_ErrNotValidForDNAType(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_ErrNotValidForDNAType(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_ErrNotValidForDNAType, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_ErrNotValidForDNAType(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_ErrNotValidForDNAType()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_ErrNotValidForDNAType, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_ErrNotValidForKeyType(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_ErrNotValidForKeyType(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_ErrNotValidForKeyType, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_ErrNotValidForKeyType(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_ErrNotValidForKeyType()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_ErrNotValidForKeyType, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_ErrPutLinkOverDeleted(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_ErrPutLinkOverDeleted(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_ErrPutLinkOverDeleted, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_ErrPutLinkOverDeleted(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_ErrPutLinkOverDeleted()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_ErrPutLinkOverDeleted, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_ErrWrongNargs(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_ErrWrongNargs(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_ErrWrongNargs, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_ErrWrongNargs(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_ErrWrongNargs()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_ErrWrongNargs, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_IsDevMode(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_IsDevMode(char p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		char p0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_IsDevMode, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_IsDevMode(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char var_getholochain_IsDevMode()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_IsDevMode, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_KValue(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_KValue(nint p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nint p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_KValue, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_KValue(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nint var_getholochain_KValue()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nint r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_KValue, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_KeyEntryDef(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_KeyEntryDef(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_KeyEntryDef, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_KeyEntryDef(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_KeyEntryDef()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_KeyEntryDef, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_NonCallableAction(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_NonCallableAction(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_NonCallableAction, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_NonCallableAction(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_NonCallableAction()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_NonCallableAction, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_NonDHTAction(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_NonDHTAction(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_NonDHTAction, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_NonDHTAction(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_NonDHTAction()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_NonDHTAction, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_SelfRevocationDoesNotVerify(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_SelfRevocationDoesNotVerify(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_SelfRevocationDoesNotVerify, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_SelfRevocationDoesNotVerify(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_SelfRevocationDoesNotVerify()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_SelfRevocationDoesNotVerify, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_SendTimeoutErr(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_SendTimeoutErr(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_SendTimeoutErr, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_SendTimeoutErr(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_SendTimeoutErr()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_SendTimeoutErr, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_UnknownWarrantTypeErr(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_UnknownWarrantTypeErr(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_UnknownWarrantTypeErr, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_UnknownWarrantTypeErr(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_UnknownWarrantTypeErr()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_UnknownWarrantTypeErr, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_ValidationFailedErr(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_ValidationFailedErr(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_ValidationFailedErr, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_ValidationFailedErr(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_ValidationFailedErr()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_ValidationFailedErr, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_var_setholochain_WarrantPropertyNotFoundErr(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void var_setholochain_WarrantPropertyNotFoundErr(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_setholochain_WarrantPropertyNotFoundErr, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_var_getholochain_WarrantPropertyNotFoundErr(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t var_getholochain_WarrantPropertyNotFoundErr()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_var_getholochain_WarrantPropertyNotFoundErr, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain__BootstrapRefreshTask(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain__BootstrapRefreshTask(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__BootstrapRefreshTask, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain__BuildJSONSchemaValidatorFromFile(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain__BuildJSONSchemaValidatorFromFile_return proxyholochain__BuildJSONSchemaValidatorFromFile(nstring p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring p0;
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain__BuildJSONSchemaValidatorFromFile_return r;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__BuildJSONSchemaValidatorFromFile, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain__BuildJSONSchemaValidatorFromString(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain__BuildJSONSchemaValidatorFromString_return proxyholochain__BuildJSONSchemaValidatorFromString(nstring p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring p0;
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain__BuildJSONSchemaValidatorFromString_return r;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__BuildJSONSchemaValidatorFromString, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain__CopyDir(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain__CopyDir(nstring p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring p0;
		nstring p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__CopyDir, &a, 20, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain__CopyFile(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain__CopyFile(nstring p0, nstring p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring p0;
		nstring p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__CopyFile, &a, 20, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain__CreateRibosome(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain__CreateRibosome_return proxyholochain__CreateRibosome(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain__CreateRibosome_return r;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__CreateRibosome, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain__Debug(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain__Debug(nstring p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__Debug, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain__DecodeWarrant(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain__DecodeWarrant_return proxyholochain__DecodeWarrant(nint p0, nbyteslice p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nint p0;
		nbyteslice p1;
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain__DecodeWarrant_return r;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__DecodeWarrant, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain__EncodingFormat(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain__EncodingFormat(nstring p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__EncodingFormat, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain__EscapeJSONValue(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain__EscapeJSONValue(nstring p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__EscapeJSONValue, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain__GossipTask(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain__GossipTask(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__GossipTask, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain__Info(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain__Info(nstring p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__Info, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain__InitializeHolochain(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain__InitializeHolochain()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		char unused;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__InitializeHolochain, &a, 0, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain__IsInitialized(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
char proxyholochain__IsInitialized(nstring p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring p0;
		char r0;
		char __pad0[3];
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__IsInitialized, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain__LoadAgent(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain__LoadAgent_return proxyholochain__LoadAgent(nstring p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring p0;
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain__LoadAgent_return r;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__LoadAgent, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain__LoadService(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain__LoadService_return proxyholochain__LoadService(nstring p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring p0;
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain__LoadService_return r;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__LoadService, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain__LoadTestConfig(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain__LoadTestConfig_return proxyholochain__LoadTestConfig(nstring p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring p0;
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain__LoadTestConfig_return r;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__LoadTestConfig, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain__MakeActionFromMessage(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain__MakeActionFromMessage_return proxyholochain__MakeActionFromMessage(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain__MakeActionFromMessage_return r;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__MakeActionFromMessage, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain__MakeDirs(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain__MakeDirs(nstring p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__MakeDirs, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain__MakeValidationPackage(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain__MakeValidationPackage_return proxyholochain__MakeValidationPackage(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain__MakeValidationPackage_return r;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__MakeValidationPackage, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain__NewCommitAction(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain__NewCommitAction(nstring p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring p0;
		int32_t p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__NewCommitAction, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain__NewDHT(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain__NewDHT(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__NewDHT, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain__NewDebugAction(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain__NewDebugAction(nstring p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__NewDebugAction, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain__NewGetBridgesAction(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain__NewGetBridgesAction(nbyteslice p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nbyteslice p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__NewGetBridgesAction, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain__NewGetLinksAction(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain__NewGetLinksAction(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__NewGetLinksAction, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain__NewJSRibosome(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain__NewJSRibosome_return proxyholochain__NewJSRibosome(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain__NewJSRibosome_return r;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__NewJSRibosome, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain__NewMakeHashAction(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain__NewMakeHashAction(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__NewMakeHashAction, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain__NewNode(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain__NewNode_return proxyholochain__NewNode(nstring p0, nstring p1, int32_t p2, char p3, int32_t p4)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring p0;
		nstring p1;
		int32_t p2;
		char p3;
		char __pad0[3];
		int32_t p4;
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain__NewNode_return r;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	a.p4 = p4;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__NewNode, &a, 36, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain__NewNucleus(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain__NewNucleus(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__NewNucleus, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain__NewPropertyAction(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain__NewPropertyAction(nstring p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__NewPropertyAction, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain__NewPutAction(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain__NewPutAction(nstring p0, int32_t p1, int32_t p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring p0;
		int32_t p1;
		int32_t p2;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__NewPutAction, &a, 20, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain__NewQueryAction(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain__NewQueryAction(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__NewQueryAction, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain__NewSelfRevocationWarrant(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain__NewSelfRevocationWarrant_return proxyholochain__NewSelfRevocationWarrant(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain__NewSelfRevocationWarrant_return r;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__NewSelfRevocationWarrant, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain__NewSignAction(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain__NewSignAction(nbyteslice p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nbyteslice p0;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__NewSignAction, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain__NewVerifySignatureAction(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain__NewVerifySignatureAction(nstring p0, nstring p1, nstring p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring p0;
		nstring p1;
		nstring p2;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__NewVerifySignatureAction, &a, 28, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain__NewZygoRibosome(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain__NewZygoRibosome_return proxyholochain__NewZygoRibosome(int32_t p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		int32_t p1;
		int32_t r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain__NewZygoRibosome_return r;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__NewZygoRibosome, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain__PrettyPrintJSON(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
struct proxyholochain__PrettyPrintJSON_return proxyholochain__PrettyPrintJSON(nbyteslice p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nbyteslice p0;
		nstring r0;
		int32_t r1;
	} __attribute__((__packed__)) a;
	struct proxyholochain__PrettyPrintJSON_return r;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__PrettyPrintJSON, &a, 20, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	r.r0 = a.r0;
	r.r1 = a.r1;
	return r;
}
extern void _cgoexp_ebd397278953_proxyholochain__RegisterBultinRibosomes(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain__RegisterBultinRibosomes()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		char unused;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__RegisterBultinRibosomes, &a, 0, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain__RetryTask(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain__RetryTask(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__RetryTask, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain__RoutingRefreshTask(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void proxyholochain__RoutingRefreshTask(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__RoutingRefreshTask, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_proxyholochain__SaveAgent(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
int32_t proxyholochain__SaveAgent(nstring p0, int32_t p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring p0;
		int32_t p1;
		int32_t r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__SaveAgent, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxyholochain__TestingAppAppPackage(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxyholochain__TestingAppAppPackage()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		nstring r0;
	} __attribute__((__packed__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxyholochain__TestingAppAppPackage, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_proxy_error_Error(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
nstring proxy_error_Error(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
		nstring r0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_proxy_error_Error, &a, 12, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_ebd397278953_IncGoRef(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void IncGoRef(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_IncGoRef, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_ebd397278953_DestroyRef(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void DestroyRef(int32_t p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		int32_t p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ebd397278953_DestroyRef, &a, 4, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}

CGO_NO_SANITIZE_THREAD
void _cgo_ebd397278953_Cfunc__Cmalloc(void *v) {
	struct {
		unsigned long long p0;
		void *r1;
	} __attribute__((__packed__)) *a = v;
	void *ret;
	_cgo_tsan_acquire();
	ret = malloc(a->p0);
	if (ret == 0 && a->p0 == 0) {
		ret = malloc(1);
	}
	a->r1 = ret;
	_cgo_tsan_release();
}
