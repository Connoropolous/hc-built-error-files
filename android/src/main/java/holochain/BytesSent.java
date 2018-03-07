// Java class holochain.BytesSent is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

public final class BytesSent implements Seq.Proxy {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    BytesSent(Seq.Ref ref) { this.ref = ref; }
    
    public BytesSent() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    public final native long getBytes();
    public final native void setBytes(long v);
    
    // skipped field BytesSent.MsgType with unsupported type: *types.Named
    
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof BytesSent)) {
            return false;
        }
        BytesSent that = (BytesSent)o;
        long thisBytes = getBytes();
        long thatBytes = that.getBytes();
        if (thisBytes != thatBytes) {
            return false;
        }
        // skipped field BytesSent.MsgType with unsupported type: *types.Named
        
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {getBytes()});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("BytesSent").append("{");
        b.append("Bytes:").append(getBytes()).append(",");
        return b.append("}").toString();
    }
}

