// Java class holochain.ActionSign is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

public final class ActionSign implements Seq.Proxy, ArgsAction {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    ActionSign(Seq.Ref ref) { this.ref = ref; }
    
    public ActionSign() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    // skipped method ActionSign.Args with unsupported parameter or return types
    
    // skipped method ActionSign.Do with unsupported parameter or return types
    
    public native String name();
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof ActionSign)) {
            return false;
        }
        ActionSign that = (ActionSign)o;
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("ActionSign").append("{");
        return b.append("}").toString();
    }
}

