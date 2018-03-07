// Java class holochain.Ribosome is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

/**
 * Ribosome type abstracts the functions of code execution environments
 */
public interface Ribosome {
    // skipped method Ribosome.BridgeGenesis with unsupported parameter or return types
    
    // skipped method Ribosome.Call with unsupported parameter or return types
    
    public void chainGenesis() throws Exception;
    public String receive(String from, String msg) throws Exception;
    // skipped method Ribosome.Run with unsupported parameter or return types
    
    // skipped method Ribosome.RunAsyncSendResponse with unsupported parameter or return types
    
    public String type();
    // skipped method Ribosome.ValidateAction with unsupported parameter or return types
    
    // skipped method Ribosome.ValidatePackagingRequest with unsupported parameter or return types
    
    
}

