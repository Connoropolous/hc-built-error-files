// Java class holochain.DelReq is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

/**
 * DelReq holds the data of a del request
 */
public final class DelReq implements Seq.Proxy {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    DelReq(Seq.Ref ref) { this.ref = ref; }
    
    public DelReq() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    // skipped field DelReq.H with unsupported type: *types.Named
    
    // skipped field DelReq.By with unsupported type: *types.Named
    
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof DelReq)) {
            return false;
        }
        DelReq that = (DelReq)o;
        // skipped field DelReq.H with unsupported type: *types.Named
        
        // skipped field DelReq.By with unsupported type: *types.Named
        
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("DelReq").append("{");
        return b.append("}").toString();
    }
}

