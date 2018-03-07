// Java class holochain.Signature is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

public final class Signature implements Seq.Proxy {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    Signature(Seq.Ref ref) { this.ref = ref; }
    
    public Signature() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    public final native byte[] getS();
    public final native void setS(byte[] v);
    
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof Signature)) {
            return false;
        }
        Signature that = (Signature)o;
        byte[] thisS = getS();
        byte[] thatS = that.getS();
        if (thisS == null) {
            if (thatS != null) {
                return false;
            }
        } else if (!thisS.equals(thatS)) {
            return false;
        }
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {getS()});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("Signature").append("{");
        b.append("S:").append(getS()).append(",");
        return b.append("}").toString();
    }
}

