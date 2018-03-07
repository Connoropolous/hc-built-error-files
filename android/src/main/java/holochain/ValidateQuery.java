// Java class holochain.ValidateQuery is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

/**
 * ValidateQuery holds the data from a validation query on the Source protocol
 */
public final class ValidateQuery implements Seq.Proxy {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    ValidateQuery(Seq.Ref ref) { this.ref = ref; }
    
    public ValidateQuery() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    // skipped field ValidateQuery.H with unsupported type: *types.Named
    
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof ValidateQuery)) {
            return false;
        }
        ValidateQuery that = (ValidateQuery)o;
        // skipped field ValidateQuery.H with unsupported type: *types.Named
        
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("ValidateQuery").append("{");
        return b.append("}").toString();
    }
}

