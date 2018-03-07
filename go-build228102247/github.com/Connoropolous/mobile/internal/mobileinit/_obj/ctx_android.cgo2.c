#line 7 "/Users/connor/go/src/github.com/Connoropolous/mobile/internal/mobileinit/ctx_android.go"

#include <jni.h>
#include <stdlib.h>

// current_vm is stored to initialize other cgo packages.
//
// As all the Go packages in a program form a single shared library,
// there can only be one JNI_OnLoad function for initialization. In
// OpenJDK there is JNI_GetCreatedJavaVMs, but this is not available
// on android.
JavaVM* current_vm;

// current_ctx is Android's android.context.Context. May be NULL.
jobject current_ctx;

char* lockJNI(uintptr_t* envp, int* attachedp) {
	JNIEnv* env;

	if (current_vm == NULL) {
		return "no current JVM";
	}

	*attachedp = 0;
	switch ((*current_vm)->GetEnv(current_vm, (void**)&env, JNI_VERSION_1_6)) {
	case JNI_OK:
		break;
	case JNI_EDETACHED:
		if ((*current_vm)->AttachCurrentThread(current_vm, &env, 0) != 0) {
			return "cannot attach to JVM";
		}
		*attachedp = 1;
		break;
	case JNI_EVERSION:
		return "bad JNI version";
	default:
		return "unknown JNI error from GetEnv";
	}

	*envp = (uintptr_t)env;
	return NULL;
}

char* checkException(uintptr_t jnienv) {
	jthrowable exc;
	JNIEnv* env = (JNIEnv*)jnienv;

	if (!(*env)->ExceptionCheck(env)) {
		return NULL;
	}

	exc = (*env)->ExceptionOccurred(env);
	(*env)->ExceptionClear(env);

	jclass clazz = (*env)->FindClass(env, "java/lang/Throwable");
	jmethodID toString = (*env)->GetMethodID(env, clazz, "toString", "()Ljava/lang/String;");
	jobject msgStr = (*env)->CallObjectMethod(env, exc, toString);
	return (char*)(*env)->GetStringUTFChars(env, msgStr, 0);
}

void unlockJNI() {
	(*current_vm)->DetachCurrentThread(current_vm);
}

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
_cgo_191360d7b6db_Cfunc_checkException(void *v)
{
	struct {
		uintptr_t p0;
		char* r;
	} __attribute__((__packed__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r;
	_cgo_tsan_acquire();
	r = (__typeof__(a->r)) checkException(a->p0);
	_cgo_tsan_release();
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

CGO_NO_SANITIZE_THREAD
void
_cgo_191360d7b6db_Cfunc_free(void *v)
{
	struct {
		void* p0;
	} __attribute__((__packed__)) *a = v;
	_cgo_tsan_acquire();
	free(a->p0);
	_cgo_tsan_release();
}

CGO_NO_SANITIZE_THREAD
void
_cgo_191360d7b6db_Cfunc_lockJNI(void *v)
{
	struct {
		uintptr_t* p0;
		int* p1;
		char* r;
	} __attribute__((__packed__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r;
	_cgo_tsan_acquire();
	r = (__typeof__(a->r)) lockJNI(a->p0, a->p1);
	_cgo_tsan_release();
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

CGO_NO_SANITIZE_THREAD
void
_cgo_191360d7b6db_Cfunc_unlockJNI(void *v)
{
	struct {
		char unused;
	} __attribute__((__packed__)) *a = v;
	_cgo_tsan_acquire();
	unlockJNI();
	_cgo_tsan_release();
}

