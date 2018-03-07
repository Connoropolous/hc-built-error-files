// Java class holochain.Warrant is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

/**
 * Warrant abstracts the notion of a multi-party cryptographically verifiable signed claim
the meaning of the warrant is understood by the warrant name an/or by properties contained in it
 */
public interface Warrant {
    /**
     * Decode unmarshals a warrant from bytes
     */
    public void decode(byte[] data) throws Exception;
    /**
     * Encode marshals the warrant into bytes for sending over the wire
     */
    public byte[] encode() throws Exception;
    // skipped method Warrant.Parties with unsupported parameter or return types
    
    // skipped method Warrant.Property with unsupported parameter or return types
    
    /**
     * Int returns the warrant type
     */
    public long type();
    /**
     * Verify confirms that the content of a warrant is valid and has been signed by the
    the parties in it.  Requires a Holochain object for context, returns nil if it
    verfies or an error
     */
    public void verify(Holochain_ h) throws Exception;
    
}

