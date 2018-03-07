// Java class holochain.DelEntry is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

/**
 * DelEntry struct holds the record of an entry&#39;s deletion
 */
public final class DelEntry implements Seq.Proxy {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    DelEntry(Seq.Ref ref) { this.ref = ref; }
    
    public DelEntry() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    // skipped field DelEntry.Hash with unsupported type: *types.Named
    
    public final native String getMessage();
    public final native void setMessage(String v);
    
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof DelEntry)) {
            return false;
        }
        DelEntry that = (DelEntry)o;
        // skipped field DelEntry.Hash with unsupported type: *types.Named
        
        String thisMessage = getMessage();
        String thatMessage = that.getMessage();
        if (thisMessage == null) {
            if (thatMessage != null) {
                return false;
            }
        } else if (!thisMessage.equals(thatMessage)) {
            return false;
        }
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {getMessage()});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("DelEntry").append("{");
        b.append("Message:").append(getMessage()).append(",");
        return b.append("}").toString();
    }
}

