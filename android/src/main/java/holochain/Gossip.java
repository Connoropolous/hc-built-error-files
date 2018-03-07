// Java class holochain.Gossip is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

/**
 * Gossip holds a gossip message
 */
public final class Gossip implements Seq.Proxy {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    Gossip(Seq.Ref ref) { this.ref = ref; }
    
    public Gossip() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    // skipped field Gossip.Puts with unsupported type: *types.Slice
    
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof Gossip)) {
            return false;
        }
        Gossip that = (Gossip)o;
        // skipped field Gossip.Puts with unsupported type: *types.Slice
        
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("Gossip").append("{");
        return b.append("}").toString();
    }
}

