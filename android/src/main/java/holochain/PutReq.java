// Java class holochain.PutReq is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

/**
 * PutReq holds the data of a put request
 */
public final class PutReq implements Seq.Proxy {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    PutReq(Seq.Ref ref) { this.ref = ref; }
    
    public PutReq() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    // skipped field PutReq.H with unsupported type: *types.Named
    
    public final native long getS();
    public final native void setS(long v);
    
    // skipped field PutReq.D with unsupported type: *types.Interface
    
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof PutReq)) {
            return false;
        }
        PutReq that = (PutReq)o;
        // skipped field PutReq.H with unsupported type: *types.Named
        
        long thisS = getS();
        long thatS = that.getS();
        if (thisS != thatS) {
            return false;
        }
        // skipped field PutReq.D with unsupported type: *types.Interface
        
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {getS()});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("PutReq").append("{");
        b.append("S:").append(getS()).append(",");
        return b.append("}").toString();
    }
}

