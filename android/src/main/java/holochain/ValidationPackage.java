// Java class holochain.ValidationPackage is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

/**
 * ValidationPackage holds app specified data needed for validation. This version
holds the package with any chain data un-marshaled after validation for passing
into the app for app level validation
 */
public final class ValidationPackage implements Seq.Proxy {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    ValidationPackage(Seq.Ref ref) { this.ref = ref; }
    
    public ValidationPackage() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    public final native Chain getChain();
    public final native void setChain(Chain v);
    
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof ValidationPackage)) {
            return false;
        }
        ValidationPackage that = (ValidationPackage)o;
        Chain thisChain = getChain();
        Chain thatChain = that.getChain();
        if (thisChain == null) {
            if (thatChain != null) {
                return false;
            }
        } else if (!thisChain.equals(thatChain)) {
            return false;
        }
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {getChain()});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("ValidationPackage").append("{");
        b.append("Chain:").append(getChain()).append(",");
        return b.append("}").toString();
    }
}

