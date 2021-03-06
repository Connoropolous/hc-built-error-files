// Java class holochain.EntryDef is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

/**
 * EntryDef struct holds an entry definition
 */
public final class EntryDef implements Seq.Proxy {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    EntryDef(Seq.Ref ref) { this.ref = ref; }
    
    public EntryDef() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    public final native String getName();
    public final native void setName(String v);
    
    public final native String getDataFormat();
    public final native void setDataFormat(String v);
    
    public final native String getSharing();
    public final native void setSharing(String v);
    
    public final native String getSchema();
    public final native void setSchema(String v);
    
    /**
     * BuildJSONSchemaValidator builds a validator in an EntryDef
     */
    public native void buildJSONSchemaValidator(String path) throws Exception;
    public native void buildJSONSchemaValidatorFromString(String schema) throws Exception;
    /**
     * IsSysEntry returns true if the entry type is system defined
     */
    public native boolean isSysEntry();
    /**
     * IsVirtualEntry returns true if the entry type is virtual
     */
    public native boolean isVirtualEntry();
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof EntryDef)) {
            return false;
        }
        EntryDef that = (EntryDef)o;
        String thisName = getName();
        String thatName = that.getName();
        if (thisName == null) {
            if (thatName != null) {
                return false;
            }
        } else if (!thisName.equals(thatName)) {
            return false;
        }
        String thisDataFormat = getDataFormat();
        String thatDataFormat = that.getDataFormat();
        if (thisDataFormat == null) {
            if (thatDataFormat != null) {
                return false;
            }
        } else if (!thisDataFormat.equals(thatDataFormat)) {
            return false;
        }
        String thisSharing = getSharing();
        String thatSharing = that.getSharing();
        if (thisSharing == null) {
            if (thatSharing != null) {
                return false;
            }
        } else if (!thisSharing.equals(thatSharing)) {
            return false;
        }
        String thisSchema = getSchema();
        String thatSchema = that.getSchema();
        if (thisSchema == null) {
            if (thatSchema != null) {
                return false;
            }
        } else if (!thisSchema.equals(thatSchema)) {
            return false;
        }
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {getName(), getDataFormat(), getSharing(), getSchema()});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("EntryDef").append("{");
        b.append("Name:").append(getName()).append(",");
        b.append("DataFormat:").append(getDataFormat()).append(",");
        b.append("Sharing:").append(getSharing()).append(",");
        b.append("Schema:").append(getSchema()).append(",");
        return b.append("}").toString();
    }
}

