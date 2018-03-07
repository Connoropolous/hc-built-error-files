// Java class holochain.ModReq is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

/**
 * ModReq holds the data of a mod request
 */
public final class ModReq implements Seq.Proxy {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    ModReq(Seq.Ref ref) { this.ref = ref; }
    
    public ModReq() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    // skipped field ModReq.H with unsupported type: *types.Named
    
    // skipped field ModReq.N with unsupported type: *types.Named
    
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof ModReq)) {
            return false;
        }
        ModReq that = (ModReq)o;
        // skipped field ModReq.H with unsupported type: *types.Named
        
        // skipped field ModReq.N with unsupported type: *types.Named
        
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("ModReq").append("{");
        return b.append("}").toString();
    }
}

