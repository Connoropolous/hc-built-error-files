package mobileinit
//go:cgo_import_dynamic __libc_init __libc_init#LIBC "libc.so"
//go:cgo_import_dynamic __cxa_atexit __cxa_atexit#LIBC "libc.so"
//go:cgo_import_dynamic malloc malloc#LIBC "libc.so"
//go:cgo_import_dynamic free free#LIBC "libc.so"
//go:cgo_import_dynamic __android_log_write __android_log_write ""
//go:cgo_import_dynamic _ _ "libandroid.so"
//go:cgo_import_dynamic _ _ "liblog.so"
//go:cgo_import_dynamic _ _ "libdl.so"
//go:cgo_import_dynamic _ _ "libc.so"