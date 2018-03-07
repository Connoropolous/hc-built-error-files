// Java class holochain.DNAFile is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

public final class DNAFile implements Seq.Proxy {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    DNAFile(Seq.Ref ref) { this.ref = ref; }
    
    public DNAFile() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    public final native long getVersion();
    public final native void setVersion(long v);
    
    // skipped field DNAFile.UUID with unsupported type: *types.Named
    
    public final native String getName();
    public final native void setName(String v);
    
    // skipped field DNAFile.Properties with unsupported type: *types.Map
    
    public final native String getPropertiesSchemaFile();
    public final native void setPropertiesSchemaFile(String v);
    
    // skipped field DNAFile.BasedOn with unsupported type: *types.Named
    
    public final native long getRequiresVersion();
    public final native void setRequiresVersion(long v);
    
    // skipped field DNAFile.DHTConfig with unsupported type: *types.Named
    
    // skipped field DNAFile.Progenitor with unsupported type: *types.Named
    
    // skipped field DNAFile.Zomes with unsupported type: *types.Slice
    
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof DNAFile)) {
            return false;
        }
        DNAFile that = (DNAFile)o;
        long thisVersion = getVersion();
        long thatVersion = that.getVersion();
        if (thisVersion != thatVersion) {
            return false;
        }
        // skipped field DNAFile.UUID with unsupported type: *types.Named
        
        String thisName = getName();
        String thatName = that.getName();
        if (thisName == null) {
            if (thatName != null) {
                return false;
            }
        } else if (!thisName.equals(thatName)) {
            return false;
        }
        // skipped field DNAFile.Properties with unsupported type: *types.Map
        
        String thisPropertiesSchemaFile = getPropertiesSchemaFile();
        String thatPropertiesSchemaFile = that.getPropertiesSchemaFile();
        if (thisPropertiesSchemaFile == null) {
            if (thatPropertiesSchemaFile != null) {
                return false;
            }
        } else if (!thisPropertiesSchemaFile.equals(thatPropertiesSchemaFile)) {
            return false;
        }
        // skipped field DNAFile.BasedOn with unsupported type: *types.Named
        
        long thisRequiresVersion = getRequiresVersion();
        long thatRequiresVersion = that.getRequiresVersion();
        if (thisRequiresVersion != thatRequiresVersion) {
            return false;
        }
        // skipped field DNAFile.DHTConfig with unsupported type: *types.Named
        
        // skipped field DNAFile.Progenitor with unsupported type: *types.Named
        
        // skipped field DNAFile.Zomes with unsupported type: *types.Slice
        
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {getVersion(), getName(), getPropertiesSchemaFile(), getRequiresVersion()});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("DNAFile").append("{");
        b.append("Version:").append(getVersion()).append(",");
        b.append("Name:").append(getName()).append(",");
        b.append("PropertiesSchemaFile:").append(getPropertiesSchemaFile()).append(",");
        b.append("RequiresVersion:").append(getRequiresVersion()).append(",");
        return b.append("}").toString();
    }
}
