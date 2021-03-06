// Java class go.Universe is a proxy for talking to a Go program.
//   gobind -lang=java 
//
// File is generated by gobind. Do not edit.
package go;

import go.Seq;

public abstract class Universe {
    static {
        Seq.touch(); // for loading the native library
        _init();
    }
    
    private Universe() {} // uninstantiable
    
    // touch is called from other bound packages to initialize this package
    public static void touch() {}
    
    private static native void _init();
    
    private static final class proxyerror extends Exception implements Seq.Proxy, error {
        private final Seq.Ref ref;
        
        @Override public final int incRefnum() {
              int refnum = ref.refnum;
              Seq.incGoRef(refnum);
              return refnum;
        }
        
        proxyerror(Seq.Ref ref) { this.ref = ref; }
        
        @Override public String getMessage() { return error(); }
        
        public native String error();
    }
    
    
}
