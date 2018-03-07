// Java class holochain.Capability is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

public final class Capability implements Seq.Proxy {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    // skipped constructor Capability.NewCapability with unsupported parameter or return types
    
    Capability(Seq.Ref ref) { this.ref = ref; }
    
    public final native String getToken();
    public final native void setToken(String v);
    
    // skipped method Capability.Revoke with unsupported parameter or return types
    
    // skipped method Capability.Validate with unsupported parameter or return types
    
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof Capability)) {
            return false;
        }
        Capability that = (Capability)o;
        String thisToken = getToken();
        String thatToken = that.getToken();
        if (thisToken == null) {
            if (thatToken != null) {
                return false;
            }
        } else if (!thisToken.equals(thatToken)) {
            return false;
        }
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {getToken()});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("Capability").append("{");
        b.append("Token:").append(getToken()).append(",");
        return b.append("}").toString();
    }
}
