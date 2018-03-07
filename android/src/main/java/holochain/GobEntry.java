// Java class holochain.GobEntry is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

/**
 * GobEntry is a structure for implementing Gob encoding of Entry content
 */
public final class GobEntry implements Seq.Proxy, Entry {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    GobEntry(Seq.Ref ref) { this.ref = ref; }
    
    public GobEntry() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    // skipped field GobEntry.C with unsupported type: *types.Interface
    
    // skipped method GobEntry.Content with unsupported parameter or return types
    
    public native byte[] marshal() throws Exception;
    // skipped method GobEntry.Sum with unsupported parameter or return types
    
    public native void unmarshal(byte[] b) throws Exception;
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof GobEntry)) {
            return false;
        }
        GobEntry that = (GobEntry)o;
        // skipped field GobEntry.C with unsupported type: *types.Interface
        
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("GobEntry").append("{");
        return b.append("}").toString();
    }
}

