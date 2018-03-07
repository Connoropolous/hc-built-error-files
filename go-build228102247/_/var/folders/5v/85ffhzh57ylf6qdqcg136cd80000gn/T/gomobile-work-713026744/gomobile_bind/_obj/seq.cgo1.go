// Created by cgo - DO NOT EDIT

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/seq.go:5
package gomobile_bind

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/seq.go:15
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_seq "github.com/Connoropolous/mobile/bind/seq"
)

func init() {
	_seq.FinalizeRef = func(ref *_seq.Ref) {
		refnum := ref.Bind_Num
		if refnum < 0 {
			panic(fmt.Sprintf("not a foreign ref: %d", refnum))
		}
		_Cfunc_go_seq_dec_ref(_Ctype_int32_t(refnum))
	}
	_seq.IncForeignRef = func(refnum int32) {
		if refnum < 0 {
			panic(fmt.Sprintf("not a foreign ref: %d", refnum))
		}
		_Cfunc_go_seq_inc_ref(_Ctype_int32_t(refnum))
	}

	signal.Notify(make(chan os.Signal), syscall.SIGPIPE)
}

//line /var/folders/5v/85ffhzh57ylf6qdqcg136cd80000gn/T/gomobile-work-713026744/gomobile_bind/seq.go:45
func IncGoRef(refnum _Ctype_int32_t) {
	_seq.Inc(int32(refnum))
}
