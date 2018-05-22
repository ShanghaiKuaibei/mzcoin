# Overview

bp - blockchain parser
bp analyzes blockchain database and retrieve information required

## Steps

- Talk to a Suncoin node, not write code to read blockchain database directly!
- Get block from 0 to head
- Collect unique addresss from blocks 
- Get uxouts for each address 
- Remove uxout after a specific seq
- Aggregate balance for each address 
- Write to a file 
- Send Suncoin2 












- call `NewVisor`
- `visor.GetUnspentOutputs`
> GetUnspentOutputs returns ReadableOutputs
> Before we can call GetUnspentOutputs, we need to prepare visor for that.

## Packages to be imported

- "github.com/skycoin/skycoin/src/visor"