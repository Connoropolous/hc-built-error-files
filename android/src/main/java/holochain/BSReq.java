// Java class holochain.BSReq is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

public final class BSReq implements Seq.Proxy {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    BSReq(Seq.Ref ref) { this.ref = ref; }
    
    public BSReq() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    public final native long getVersion();
    public final native void setVersion(long v);
    
    public final native String getNodeID();
    public final native void setNodeID(String v);
    
    public final native String getNodeAddr();
    public final native void setNodeAddr(String v);
    
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof BSReq)) {
            return false;
        }
        BSReq that = (BSReq)o;
        long thisVersion = getVersion();
        long thatVersion = that.getVersion();
        if (thisVersion != thatVersion) {
            return false;
        }
        String thisNodeID = getNodeID();
        String thatNodeID = that.getNodeID();
        if (thisNodeID == null) {
            if (thatNodeID != null) {
                return false;
            }
        } else if (!thisNodeID.equals(thatNodeID)) {
            return false;
        }
        String thisNodeAddr = getNodeAddr();
        String thatNodeAddr = that.getNodeAddr();
        if (thisNodeAddr == null) {
            if (thatNodeAddr != null) {
                return false;
            }
        } else if (!thisNodeAddr.equals(thatNodeAddr)) {
            return false;
        }
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {getVersion(), getNodeID(), getNodeAddr()});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("BSReq").append("{");
        b.append("Version:").append(getVersion()).append(",");
        b.append("NodeID:").append(getNodeID()).append(",");
        b.append("NodeAddr:").append(getNodeAddr()).append(",");
        return b.append("}").toString();
    }
}

