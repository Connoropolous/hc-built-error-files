// Java class holochain.DNA is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

public final class DNA implements Seq.Proxy {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    DNA(Seq.Ref ref) { this.ref = ref; }
    
    public DNA() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    public final native long getVersion();
    public final native void setVersion(long v);
    
    // skipped field DNA.UUID with unsupported type: *types.Named
    
    public final native String getName();
    public final native void setName(String v);
    
    // skipped field DNA.Properties with unsupported type: *types.Map
    
    public final native String getPropertiesSchema();
    public final native void setPropertiesSchema(String v);
    
    public final native String getAgentIdentitySchema();
    public final native void setAgentIdentitySchema(String v);
    
    // skipped field DNA.BasedOn with unsupported type: *types.Named
    
    public final native long getRequiresVersion();
    public final native void setRequiresVersion(long v);
    
    // skipped field DNA.DHTConfig with unsupported type: *types.Named
    
    // skipped field DNA.Progenitor with unsupported type: *types.Named
    
    // skipped field DNA.Zomes with unsupported type: *types.Slice
    
    /**
     * NewUUID generates a new UUID for the DNA
     */
    public native void newUUID() throws Exception;
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof DNA)) {
            return false;
        }
        DNA that = (DNA)o;
        long thisVersion = getVersion();
        long thatVersion = that.getVersion();
        if (thisVersion != thatVersion) {
            return false;
        }
        // skipped field DNA.UUID with unsupported type: *types.Named
        
        String thisName = getName();
        String thatName = that.getName();
        if (thisName == null) {
            if (thatName != null) {
                return false;
            }
        } else if (!thisName.equals(thatName)) {
            return false;
        }
        // skipped field DNA.Properties with unsupported type: *types.Map
        
        String thisPropertiesSchema = getPropertiesSchema();
        String thatPropertiesSchema = that.getPropertiesSchema();
        if (thisPropertiesSchema == null) {
            if (thatPropertiesSchema != null) {
                return false;
            }
        } else if (!thisPropertiesSchema.equals(thatPropertiesSchema)) {
            return false;
        }
        String thisAgentIdentitySchema = getAgentIdentitySchema();
        String thatAgentIdentitySchema = that.getAgentIdentitySchema();
        if (thisAgentIdentitySchema == null) {
            if (thatAgentIdentitySchema != null) {
                return false;
            }
        } else if (!thisAgentIdentitySchema.equals(thatAgentIdentitySchema)) {
            return false;
        }
        // skipped field DNA.BasedOn with unsupported type: *types.Named
        
        long thisRequiresVersion = getRequiresVersion();
        long thatRequiresVersion = that.getRequiresVersion();
        if (thisRequiresVersion != thatRequiresVersion) {
            return false;
        }
        // skipped field DNA.DHTConfig with unsupported type: *types.Named
        
        // skipped field DNA.Progenitor with unsupported type: *types.Named
        
        // skipped field DNA.Zomes with unsupported type: *types.Slice
        
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {getVersion(), getName(), getPropertiesSchema(), getAgentIdentitySchema(), getRequiresVersion()});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("DNA").append("{");
        b.append("Version:").append(getVersion()).append(",");
        b.append("Name:").append(getName()).append(",");
        b.append("PropertiesSchema:").append(getPropertiesSchema()).append(",");
        b.append("AgentIdentitySchema:").append(getAgentIdentitySchema()).append(",");
        b.append("RequiresVersion:").append(getRequiresVersion()).append(",");
        return b.append("}").toString();
    }
}

