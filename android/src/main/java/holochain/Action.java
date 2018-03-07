// Java class holochain.Action is a proxy for talking to a Go program.
//   gobind -lang=java github.com/Holochain/holochain-proto
//
// File is generated by gobind. Do not edit.
package holochain;

import go.Seq;

/**
 * Action provides an abstraction for grouping all the aspects of a nucleus function, i.e.
the initiating actions, receiving them, validation, ribosome generation etc
 */
public interface Action extends ArgsAction {
    // skipped method Action.Args with unsupported parameter or return types
    
    // skipped method Action.Do with unsupported parameter or return types
    
    public String name();
    // skipped method Action.Receive with unsupported parameter or return types
    
    
}

