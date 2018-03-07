// Java class holochain.BSResp is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

public final class BSResp implements Seq.Proxy {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    BSResp(Seq.Ref ref) { this.ref = ref; }
    
    public BSResp() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    // skipped field BSResp.Req with unsupported type: *types.Named
    
    public final native String getRemote();
    public final native void setRemote(String v);
    
    // skipped field BSResp.LastSeen with unsupported type: *types.Named
    
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof BSResp)) {
            return false;
        }
        BSResp that = (BSResp)o;
        // skipped field BSResp.Req with unsupported type: *types.Named
        
        String thisRemote = getRemote();
        String thatRemote = that.getRemote();
        if (thisRemote == null) {
            if (thatRemote != null) {
                return false;
            }
        } else if (!thisRemote.equals(thatRemote)) {
            return false;
        }
        // skipped field BSResp.LastSeen with unsupported type: *types.Named
        
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {getRemote()});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("BSResp").append("{");
        b.append("Remote:").append(getRemote()).append(",");
        return b.append("}").toString();
    }
}
