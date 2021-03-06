// Java class holochain.ActionSend is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

public final class ActionSend implements Seq.Proxy, Action, ArgsAction {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    ActionSend(Seq.Ref ref) { this.ref = ref; }
    
    public ActionSend() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    // skipped method ActionSend.Args with unsupported parameter or return types
    
    // skipped method ActionSend.Do with unsupported parameter or return types
    
    public native String name();
    // skipped method ActionSend.Receive with unsupported parameter or return types
    
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof ActionSend)) {
            return false;
        }
        ActionSend that = (ActionSend)o;
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("ActionSend").append("{");
        return b.append("}").toString();
    }
}

