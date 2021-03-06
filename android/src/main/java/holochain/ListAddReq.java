// Java class holochain.ListAddReq is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

public final class ListAddReq implements Seq.Proxy {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    ListAddReq(Seq.Ref ref) { this.ref = ref; }
    
    public ListAddReq() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    public final native String getListType();
    public final native void setListType(String v);
    
    // skipped field ListAddReq.Peers with unsupported type: *types.Slice
    
    public final native long getWarrantType();
    public final native void setWarrantType(long v);
    
    public final native byte[] getWarrant();
    public final native void setWarrant(byte[] v);
    
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof ListAddReq)) {
            return false;
        }
        ListAddReq that = (ListAddReq)o;
        String thisListType = getListType();
        String thatListType = that.getListType();
        if (thisListType == null) {
            if (thatListType != null) {
                return false;
            }
        } else if (!thisListType.equals(thatListType)) {
            return false;
        }
        // skipped field ListAddReq.Peers with unsupported type: *types.Slice
        
        long thisWarrantType = getWarrantType();
        long thatWarrantType = that.getWarrantType();
        if (thisWarrantType != thatWarrantType) {
            return false;
        }
        byte[] thisWarrant = getWarrant();
        byte[] thatWarrant = that.getWarrant();
        if (thisWarrant == null) {
            if (thatWarrant != null) {
                return false;
            }
        } else if (!thisWarrant.equals(thatWarrant)) {
            return false;
        }
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {getListType(), getWarrantType(), getWarrant()});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("ListAddReq").append("{");
        b.append("ListType:").append(getListType()).append(",");
        b.append("WarrantType:").append(getWarrantType()).append(",");
        b.append("Warrant:").append(getWarrant()).append(",");
        return b.append("}").toString();
    }
}

