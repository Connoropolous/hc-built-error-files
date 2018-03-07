// Java class holochain.TestConfig is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

/**
 * TestConfig holds the configuration options for a test
 */
public final class TestConfig implements Seq.Proxy {
    static { Holochain.touch(); }
    
    private final Seq.Ref ref;
    
    @Override public final int incRefnum() {
          int refnum = ref.refnum;
          Seq.incGoRef(refnum);
          return refnum;
    }
    
    TestConfig(Seq.Ref ref) { this.ref = ref; }
    
    public TestConfig() { this.ref = __New(); }
    
    private static native Seq.Ref __New();
    
    public final native long getGossipInterval();
    public final native void setGossipInterval(long v);
    
    public final native long getDuration();
    public final native void setDuration(long v);
    
    // skipped field TestConfig.Clone with unsupported type: *types.Slice
    
    @Override public boolean equals(Object o) {
        if (o == null || !(o instanceof TestConfig)) {
            return false;
        }
        TestConfig that = (TestConfig)o;
        long thisGossipInterval = getGossipInterval();
        long thatGossipInterval = that.getGossipInterval();
        if (thisGossipInterval != thatGossipInterval) {
            return false;
        }
        long thisDuration = getDuration();
        long thatDuration = that.getDuration();
        if (thisDuration != thatDuration) {
            return false;
        }
        // skipped field TestConfig.Clone with unsupported type: *types.Slice
        
        return true;
    }
    
    @Override public int hashCode() {
        return java.util.Arrays.hashCode(new Object[] {getGossipInterval(), getDuration()});
    }
    
    @Override public String toString() {
        StringBuilder b = new StringBuilder();
        b.append("TestConfig").append("{");
        b.append("GossipInterval:").append(getGossipInterval()).append(",");
        b.append("Duration:").append(getDuration()).append(",");
        return b.append("}").toString();
    }
}
