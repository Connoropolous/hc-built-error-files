#line 7 "/var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/go_holochainmain.go"

#include <stdlib.h>
#include <stdint.h>
#include "seq.h"
#include "holochain.h"


#line 1 "cgo-generated-wrapper"


#line 1 "cgo-gcc-prolog"
/*
  If x and y are not equal, the type will be invalid
  (have a negative array count) and an inscrutable error will come
  out of the compiler and hopefully mention "name".
*/
#define __cgo_compile_assert_eq(x, y, name) typedef char name[(x-y)*(x-y)*-2+1];

/* Check at compile time that the sizes we use match our expectations. */
#define __cgo_size_assert(t, n) __cgo_compile_assert_eq(sizeof(t), n, _cgo_sizeof_##t##_is_not_##n)

__cgo_size_assert(char, 1)
__cgo_size_assert(short, 2)
__cgo_size_assert(int, 4)
typedef long long __cgo_long_long;
__cgo_size_assert(__cgo_long_long, 8)
__cgo_size_assert(float, 4)
__cgo_size_assert(double, 8)

extern char* _cgo_topofstack(void);

#include <errno.h>
#include <string.h>


#define CGO_NO_SANITIZE_THREAD
#define _cgo_tsan_acquire()
#define _cgo_tsan_release()

CGO_NO_SANITIZE_THREAD
void
_cgo_ebd397278953_Cfunc_cproxyholochain_Action_Name(void *v)
{
	struct {
		int32_t p0;
		nstring r;
	} __attribute__((__packed__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r;
	_cgo_tsan_acquire();
	r = cproxyholochain_Action_Name(a->p0);
	_cgo_tsan_release();
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

CGO_NO_SANITIZE_THREAD
void
_cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_CheckValidationRequest(void *v)
{
	struct {
		int32_t p0;
		int32_t p1;
		int32_t r;
	} __attribute__((__packed__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r;
	_cgo_tsan_acquire();
	r = cproxyholochain_CommittingAction_CheckValidationRequest(a->p0, a->p1);
	_cgo_tsan_release();
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

CGO_NO_SANITIZE_THREAD
void
_cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_Entry(void *v)
{
	struct {
		int32_t p0;
		int32_t r;
	} __attribute__((__packed__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r;
	_cgo_tsan_acquire();
	r = cproxyholochain_CommittingAction_Entry(a->p0);
	_cgo_tsan_release();
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

CGO_NO_SANITIZE_THREAD
void
_cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_EntryType(void *v)
{
	struct {
		int32_t p0;
		nstring r;
	} __attribute__((__packed__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r;
	_cgo_tsan_acquire();
	r = cproxyholochain_CommittingAction_EntryType(a->p0);
	_cgo_tsan_release();
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

CGO_NO_SANITIZE_THREAD
void
_cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_Name(void *v)
{
	struct {
		int32_t p0;
		nstring r;
	} __attribute__((__packed__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r;
	_cgo_tsan_acquire();
	r = cproxyholochain_CommittingAction_Name(a->p0);
	_cgo_tsan_release();
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

CGO_NO_SANITIZE_THREAD
void
_cgo_ebd397278953_Cfunc_cproxyholochain_CommittingAction_SetHeader(void *v)
{
	struct {
		int32_t p0;
		int32_t p1;
	} __attribute__((__packed__)) *a = v;
	_cgo_tsan_acquire();
	cproxyholochain_CommittingAction_SetHeader(a->p0, a->p1);
	_cgo_tsan_release();
}

CGO_NO_SANITIZE_THREAD
void
_cgo_ebd397278953_Cfunc_cproxyholochain_Entry_Marshal(void *v)
{
	struct {
		int32_t p0;
		struct cproxyholochain_Entry_Marshal_return r;
	} __attribute__((__packed__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r;
	_cgo_tsan_acquire();
	r = cproxyholochain_Entry_Marshal(a->p0);
	_cgo_tsan_release();
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

CGO_NO_SANITIZE_THREAD
void
_cgo_ebd397278953_Cfunc_cproxyholochain_Entry_Unmarshal(void *v)
{
	struct {
		int32_t p0;
		nbyteslice p1;
		int32_t r;
	} __attribute__((__packed__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r;
	_cgo_tsan_acquire();
	r = cproxyholochain_Entry_Unmarshal(a->p0, a->p1);
	_cgo_tsan_release();
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

CGO_NO_SANITIZE_THREAD
void
_cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Marshal(void *v)
{
	struct {
		int32_t p0;
		struct cproxyholochain_Revocation_Marshal_return r;
	} __attribute__((__packed__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r;
	_cgo_tsan_acquire();
	r = cproxyholochain_Revocation_Marshal(a->p0);
	_cgo_tsan_release();
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

CGO_NO_SANITIZE_THREAD
void
_cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Unmarshal(void *v)
{
	struct {
		int32_t p0;
		nbyteslice p1;
		int32_t r;
	} __attribute__((__packed__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r;
	_cgo_tsan_acquire();
	r = cproxyholochain_Revocation_Unmarshal(a->p0, a->p1);
	_cgo_tsan_release();
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

CGO_NO_SANITIZE_THREAD
void
_cgo_ebd397278953_Cfunc_cproxyholochain_Revocation_Verify(void *v)
{
	struct {
		int32_t p0;
		int32_t r;
	} __attribute__((__packed__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r;
	_cgo_tsan_acquire();
	r = cproxyholochain_Revocation_Verify(a->p0);
	_cgo_tsan_release();
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

CGO_NO_SANITIZE_THREAD
void
_cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_ChainGenesis(void *v)
{
	struct {
		int32_t p0;
		int32_t r;
	} __attribute__((__packed__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r;
	_cgo_tsan_acquire();
	r = cproxyholochain_Ribosome_ChainGenesis(a->p0);
	_cgo_tsan_release();
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

CGO_NO_SANITIZE_THREAD
void
_cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_Receive(void *v)
{
	struct {
		int32_t p0;
		nstring p1;
		nstring p2;
		struct cproxyholochain_Ribosome_Receive_return r;
	} __attribute__((__packed__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r;
	_cgo_tsan_acquire();
	r = cproxyholochain_Ribosome_Receive(a->p0, a->p1, a->p2);
	_cgo_tsan_release();
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

CGO_NO_SANITIZE_THREAD
void
_cgo_ebd397278953_Cfunc_cproxyholochain_Ribosome_Type(void *v)
{
	struct {
		int32_t p0;
		nstring r;
	} __attribute__((__packed__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r;
	_cgo_tsan_acquire();
	r = cproxyholochain_Ribosome_Type(a->p0);
	_cgo_tsan_release();
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

CGO_NO_SANITIZE_THREAD
void
_cgo_ebd397278953_Cfunc_cproxyholochain_ValidatingAction_CheckValidationRequest(void *v)
{
	struct {
		int32_t p0;
		int32_t p1;
		int32_t r;
	} __attribute__((__packed__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r;
	_cgo_tsan_acquire();
	r = cproxyholochain_ValidatingAction_CheckValidationRequest(a->p0, a->p1);
	_cgo_tsan_release();
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

CGO_NO_SANITIZE_THREAD
void
_cgo_ebd397278953_Cfunc_cproxyholochain_ValidatingAction_Name(void *v)
{
	struct {
		int32_t p0;
		nstring r;
	} __attribute__((__packed__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r;
	_cgo_tsan_acquire();
	r = cproxyholochain_ValidatingAction_Name(a->p0);
	_cgo_tsan_release();
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

CGO_NO_SANITIZE_THREAD
void
_cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Decode(void *v)
{
	struct {
		int32_t p0;
		nbyteslice p1;
		int32_t r;
	} __attribute__((__packed__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r;
	_cgo_tsan_acquire();
	r = cproxyholochain_Warrant_Decode(a->p0, a->p1);
	_cgo_tsan_release();
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

CGO_NO_SANITIZE_THREAD
void
_cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Encode(void *v)
{
	struct {
		int32_t p0;
		struct cproxyholochain_Warrant_Encode_return r;
	} __attribute__((__packed__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r;
	_cgo_tsan_acquire();
	r = cproxyholochain_Warrant_Encode(a->p0);
	_cgo_tsan_release();
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

CGO_NO_SANITIZE_THREAD
void
_cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Type(void *v)
{
	struct {
		int32_t p0;
		nint r;
	} __attribute__((__packed__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r;
	_cgo_tsan_acquire();
	r = cproxyholochain_Warrant_Type(a->p0);
	_cgo_tsan_release();
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

CGO_NO_SANITIZE_THREAD
void
_cgo_ebd397278953_Cfunc_cproxyholochain_Warrant_Verify(void *v)
{
	struct {
		int32_t p0;
		int32_t p1;
		int32_t r;
	} __attribute__((__packed__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r;
	_cgo_tsan_acquire();
	r = cproxyholochain_Warrant_Verify(a->p0, a->p1);
	_cgo_tsan_release();
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

