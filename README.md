# hc-built-error-files

The end goal of this research is to compile the holochain code using `gomobile` to have a java package, and an .aar android archive which could be included as libraries into a native app.

Current status: `gomobile bind -target android github.com/Holochain/holochain-proto` command is now successfully running and generating files, but there are lots of methods and data types that it would seem that it can't implement (yet). It depends on my holochain branch too: https://github.com/Holochain/holochain-proto/tree/for-mobile-build. The hope is to fix the unimplemented ones, but it may be too much work and not worth it, especially since we really don't know that it would for sure work at the end of all that.

My research to get this far is here: https://paper.dropbox.com/doc/Holochain-on-Android-research-qqoos392BFMmvjGg5vfZn

Interesting files are

- [go_holochainmain.go](https://github.com/Connoropolous/hc-built-error-files/blob/master/gomobile_bind/go_holochainmain.go)
- [holochain.h](https://github.com/Connoropolous/hc-built-error-files/blob/master/gomobile_bind/holochain.h)
- [java_holochain.c](https://github.com/Connoropolous/hc-built-error-files/blob/master/gomobile_bind/java_holochain.c)
- [src/main/java/holochain](https://github.com/Connoropolous/hc-built-error-files/tree/master/android/src/main/java/holochain)

There are a lot of lines like this in the code:

`// skipped method ActionBridge.Args with unsupported parameter or return types`

In both the GO files, and the java files.

Currently getting these errors
```
/var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-376003743/gomobile_bind/go_holochainmain.go:1332: cannot use (*proxyholochain_Entry)(_v_ref) (type *proxyholochain_Entry) as type holochain.Entry in assignment:
	*proxyholochain_Entry does not implement holochain.Entry (missing Content method)
/var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-376003743/gomobile_bind/go_holochainmain.go:3996: cannot use (*proxyholochain_Entry)(_v_ref) (type *proxyholochain_Entry) as type holochain.Entry in assignment:
	*proxyholochain_Entry does not implement holochain.Entry (missing Content method)
/var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-376003743/gomobile_bind/go_holochainmain.go:4332: cannot use (*proxyholochain_Agent)(_v_ref) (type *proxyholochain_Agent) as type holochain.Agent in assignment:
	*proxyholochain_Agent does not implement holochain.Agent (missing AgentEntry method)
/var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-376003743/gomobile_bind/go_holochainmain.go:4376: cannot use (*proxyholochain_Agent)(_param_agent_ref) (type *proxyholochain_Agent) as type holochain.Agent in assignment:
	*proxyholochain_Agent does not implement holochain.Agent (missing AgentEntry method)
/var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-376003743/gomobile_bind/go_holochainmain.go:4498: cannot use (*proxyholochain_Agent)(_param_agent_ref) (type *proxyholochain_Agent) as type holochain.Agent in assignment:
	*proxyholochain_Agent does not implement holochain.Agent (missing AgentEntry method)
/var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-376003743/gomobile_bind/go_holochainmain.go:5482: cannot use (*proxyholochain_Entry)(_res_ref) (type *proxyholochain_Entry) as type holochain.Entry in assignment:
	*proxyholochain_Entry does not implement holochain.Entry (missing Content method)
/var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-376003743/gomobile_bind/go_holochainmain.go:7065: cannot use (*proxyholochain_Entry)(_param_entry_ref) (type *proxyholochain_Entry) as type holochain.Entry in assignment:
	*proxyholochain_Entry does not implement holochain.Entry (missing Content method)
/var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-376003743/gomobile_bind/go_holochainmain.go:7171: cannot use (*proxyholochain_Entry)(_param_entry_ref) (type *proxyholochain_Entry) as type holochain.Entry in assignment:
	*proxyholochain_Entry does not implement holochain.Entry (missing Content method)
/var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-376003743/gomobile_bind/go_holochainmain.go:7251: cannot use (*proxyholochain_Entry)(_param_entry_ref) (type *proxyholochain_Entry) as type holochain.Entry in assignment:
	*proxyholochain_Entry does not implement holochain.Entry (missing Content method)
/var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-376003743/gomobile_bind/go_holochainmain.go:7401: cannot use (*proxyholochain_Agent)(_param_agent_ref) (type *proxyholochain_Agent) as type holochain.Agent in assignment:
	*proxyholochain_Agent does not implement holochain.Agent (missing AgentEntry method)
/var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-376003743/gomobile_bind/go_holochainmain.go:7401: too many errors
```
