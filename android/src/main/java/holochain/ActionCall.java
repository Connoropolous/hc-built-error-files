// Java class holochain.ActionCall is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

public final class ActionCall implements Seq.Proxy, ArgsAction {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    ActionCall(Seq.Ref ref) { this.ref = ref; }
    
    public ActionCall() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    // skipped method ActionCall.Args with unsupported parameter or return types
    
    // skipped method ActionCall.Do with unsupported parameter or return types
    
    public native String name();
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof ActionCall)) {
            return false;
        }
        ActionCall that = (ActionCall)o;
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("ActionCall").append("{");
        return b.append("}").toString();
    }
}
