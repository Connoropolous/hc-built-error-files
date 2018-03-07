// Java class holochain.Protocol is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

/**
 * Protocol encapsulates data for our different protocols
 */
public final class Protocol implements Seq.Proxy {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    Protocol(Seq.Ref ref) { this.ref = ref; }
    
    public Protocol() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    // skipped field Protocol.ID with unsupported type: *types.Named
    
    // skipped field Protocol.Receiver with unsupported type: *types.Named
    
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof Protocol)) {
            return false;
        }
        Protocol that = (Protocol)o;
        // skipped field Protocol.ID with unsupported type: *types.Named
        
        // skipped field Protocol.Receiver with unsupported type: *types.Named
        
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("Protocol").append("{");
        return b.append("}").toString();
    }
}
