// Java class holochain.Header is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

/**
 * Header holds chain links, type, timestamp and signature
 */
public final class Header implements Seq.Proxy {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    Header(Seq.Ref ref) { this.ref = ref; }
    
    public Header() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    public final native String getType();
    public final native void setType(String v);
    
    // skipped field Header.Time with unsupported type: *types.Named
    
    // skipped field Header.HeaderLink with unsupported type: *types.Named
    
    // skipped field Header.EntryLink with unsupported type: *types.Named
    
    // skipped field Header.TypeLink with unsupported type: *types.Named
    
    // skipped field Header.Sig with unsupported type: *types.Named
    
    // skipped field Header.Change with unsupported type: *types.Named
    
    /**
     * Marshal writes a header to bytes
     */
    public native byte[] marshal() throws Exception;
    // skipped method Header.Sum with unsupported parameter or return types
    
    /**
     * Unmarshal reads a header from bytes
     */
    public native void unmarshal(byte[] b, long hashSize) throws Exception;
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof Header)) {
            return false;
        }
        Header that = (Header)o;
        String thisType = getType();
        String thatType = that.getType();
        if (thisType == null) {
            if (thatType != null) {
                return false;
            }
        } else if (!thisType.equals(thatType)) {
            return false;
        }
        // skipped field Header.Time with unsupported type: *types.Named
        
        // skipped field Header.HeaderLink with unsupported type: *types.Named
        
        // skipped field Header.EntryLink with unsupported type: *types.Named
        
        // skipped field Header.TypeLink with unsupported type: *types.Named
        
        // skipped field Header.Sig with unsupported type: *types.Named
        
        // skipped field Header.Change with unsupported type: *types.Named
        
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {getType()});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("Header").append("{");
        b.append("Type:").append(getType()).append(",");
        return b.append("}").toString();
    }
}
