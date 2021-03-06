// Java class holochain.LinkQueryResp is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

/**
 * LinkQueryResp holds response to getLinks query
 */
public final class LinkQueryResp implements Seq.Proxy {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    LinkQueryResp(Seq.Ref ref) { this.ref = ref; }
    
    public LinkQueryResp() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    // skipped field LinkQueryResp.Links with unsupported type: *types.Slice
    
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof LinkQueryResp)) {
            return false;
        }
        LinkQueryResp that = (LinkQueryResp)o;
        // skipped field LinkQueryResp.Links with unsupported type: *types.Slice
        
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("LinkQueryResp").append("{");
        return b.append("}").toString();
    }
}

