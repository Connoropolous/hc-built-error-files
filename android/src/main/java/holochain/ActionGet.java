// Java class holochain.ActionGet is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

/**
 * ------------------------------------------------------------
Get
 */
public final class ActionGet implements Seq.Proxy, Action, ArgsAction {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    ActionGet(Seq.Ref ref) { this.ref = ref; }
    
    public ActionGet() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    // skipped method ActionGet.Args with unsupported parameter or return types
    
    // skipped method ActionGet.Do with unsupported parameter or return types
    
    public native String name();
    // skipped method ActionGet.Receive with unsupported parameter or return types
    
    // skipped method ActionGet.SysValidation with unsupported parameter or return types
    
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof ActionGet)) {
            return false;
        }
        ActionGet that = (ActionGet)o;
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("ActionGet").append("{");
        return b.append("}").toString();
    }
}
