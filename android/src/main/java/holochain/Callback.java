// Java class holochain.Callback is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

public final class Callback implements Seq.Proxy {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    Callback(Seq.Ref ref) { this.ref = ref; }
    
    public Callback() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    public final native String getFunction();
    public final native void setFunction(String v);
    
    public final native String getID();
    public final native void setID(String v);
    
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof Callback)) {
            return false;
        }
        Callback that = (Callback)o;
        String thisFunction = getFunction();
        String thatFunction = that.getFunction();
        if (thisFunction == null) {
            if (thatFunction != null) {
                return false;
            }
        } else if (!thisFunction.equals(thatFunction)) {
            return false;
        }
        String thisID = getID();
        String thatID = that.getID();
        if (thisID == null) {
            if (thatID != null) {
                return false;
            }
        } else if (!thisID.equals(thatID)) {
            return false;
        }
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {getFunction(), getID()});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("Callback").append("{");
        b.append("Function:").append(getFunction()).append(",");
        b.append("ID:").append(getID()).append(",");
        return b.append("}").toString();
    }
}
