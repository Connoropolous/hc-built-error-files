# hc-built-error-files

Current status: `gomobile bind -target android github.com/Holochain/holochain-proto` command is now successfully running and generating files, but there are lots of methods and data types that it would seem that it can't implement (yet). It depends on my holochain branch too: https://github.com/Holochain/holochain-proto/tree/for-mobile-build. The hope is to fix the unimplemented ones, but it may be too much work and not worth it, especially since we really don't know that it would for sure work at the end of all that.

My research to get this far is here: https://paper.dropbox.com/doc/Holochain-on-Android-research-qqoos392BFMmvjGg5vfZn

Interesting files are

- [go_holochainmain.go](https://github.com/Connoropolous/hc-built-error-files/blob/master/gomobile_bind/go_holochainmain.go)
- [holochain.h](https://github.com/Connoropolous/hc-built-error-files/blob/master/gomobile_bind/holochain.h)
- [java_holochain.c](https://github.com/Connoropolous/hc-built-error-files/blob/master/gomobile_bind/java_holochain.c)
- [src/main/java/holochain](https://github.com/Connoropolous/hc-built-error-files/tree/master/android/src/main/java/holochain)
