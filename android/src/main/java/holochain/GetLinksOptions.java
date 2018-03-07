// Java class holochain.GetLinksOptions is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

/**
 * GetLinksOptions options to holochain level GetLinks functions
 */
public final class GetLinksOptions implements Seq.Proxy {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    GetLinksOptions(Seq.Ref ref) { this.ref = ref; }
    
    public GetLinksOptions() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    public final native boolean getLoad();
    public final native void setLoad(boolean v);
    
    public final native long getStatusMask();
    public final native void setStatusMask(long v);
    
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof GetLinksOptions)) {
            return false;
        }
        GetLinksOptions that = (GetLinksOptions)o;
        boolean thisLoad = getLoad();
        boolean thatLoad = that.getLoad();
        if (thisLoad != thatLoad) {
            return false;
        }
        long thisStatusMask = getStatusMask();
        long thatStatusMask = that.getStatusMask();
        if (thisStatusMask != thatStatusMask) {
            return false;
        }
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {getLoad(), getStatusMask()});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("GetLinksOptions").append("{");
        b.append("Load:").append(getLoad()).append(",");
        b.append("StatusMask:").append(getStatusMask()).append(",");
        return b.append("}").toString();
    }
}
