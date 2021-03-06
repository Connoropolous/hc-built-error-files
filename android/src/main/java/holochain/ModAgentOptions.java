// Java class holochain.ModAgentOptions is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

public final class ModAgentOptions implements Seq.Proxy {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    ModAgentOptions(Seq.Ref ref) { this.ref = ref; }
    
    public ModAgentOptions() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    public final native String getIdentity();
    public final native void setIdentity(String v);
    
    public final native String getRevocation();
    public final native void setRevocation(String v);
    
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof ModAgentOptions)) {
            return false;
        }
        ModAgentOptions that = (ModAgentOptions)o;
        String thisIdentity = getIdentity();
        String thatIdentity = that.getIdentity();
        if (thisIdentity == null) {
            if (thatIdentity != null) {
                return false;
            }
        } else if (!thisIdentity.equals(thatIdentity)) {
            return false;
        }
        String thisRevocation = getRevocation();
        String thatRevocation = that.getRevocation();
        if (thisRevocation == null) {
            if (thatRevocation != null) {
                return false;
            }
        } else if (!thisRevocation.equals(thatRevocation)) {
            return false;
        }
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {getIdentity(), getRevocation()});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("ModAgentOptions").append("{");
        b.append("Identity:").append(getIdentity()).append(",");
        b.append("Revocation:").append(getRevocation()).append(",");
        return b.append("}").toString();
    }
}

