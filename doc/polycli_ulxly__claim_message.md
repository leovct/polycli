# `polycli ulxly  claim message`

> Auto-generated documentation.

## Table of Contents

- [Description](#description)
- [Usage](#usage)
- [Flags](#flags)
- [See Also](#see-also)

## Description

Claim a message

```bash
polycli ulxly  claim message [flags]
```

## Usage

This command is used to claim a message type deposit. Here is the interface of the method that's being used:

```solidity
/**
 * @notice Verify merkle proof and execute message
 * If the receiving address is an EOA, the call will result as a success
 * Which means that the amount of ether will be transferred correctly, but the message
 * will not trigger any execution
 * @param smtProofLocalExitRoot Smt proof to proof the leaf against the exit root
 * @param smtProofRollupExitRoot Smt proof to proof the rollupLocalExitRoot against the rollups exit root
 * @param globalIndex Global index is defined as:
 * | 191 bits |    1 bit     |   32 bits   |     32 bits    |
 * |    0     |  mainnetFlag | rollupIndex | localRootIndex |
 * note that only the rollup index will be used only in case the mainnet flag is 0
 * note that global index do not assert the unused bits to 0.
 * This means that when synching the events, the globalIndex must be decoded the same way that in the Smart contract
 * to avoid possible synch attacks
 * @param mainnetExitRoot Mainnet exit root
 * @param rollupExitRoot Rollup exit root
 * @param originNetwork Origin network
 * @param originAddress Origin address
 * @param destinationNetwork Network destination
 * @param destinationAddress Address destination
 * @param amount message value
 * @param metadata Abi encoded metadata if any, empty otherwise
 */
function claimMessage(
    bytes32[_DEPOSIT_CONTRACT_TREE_DEPTH] calldata smtProofLocalExitRoot,
    bytes32[_DEPOSIT_CONTRACT_TREE_DEPTH] calldata smtProofRollupExitRoot,
    uint256 globalIndex,
    bytes32 mainnetExitRoot,
    bytes32 rollupExitRoot,
    uint32 originNetwork,
    address originAddress,
    uint32 destinationNetwork,
    address destinationAddress,
    uint256 amount,
    bytes calldata metadata
) external ifNotEmergencyState {
```

[Here](https://github.com/0xPolygonHermez/zkevm-contracts/blob/c8659e6282340de7bdb8fdbf7924a9bd2996bc98/contracts/v2/PolygonZkEVMBridgeV2.sol#L588-L623) is a link to the source code.

This command is essentially identical to `claim asset`, but it's specific to deposits that are of the message leaf type rather than assets. In order to use this command, I'm going to try to claim one of the messages that I sent while testing `polycli ulxly bridge message`.

```bash
curl -s https://bridge-api.cardona.zkevm-rpc.com/bridges/0xC92AeF5873d058a76685140F3328B0DED79733Af | jq '.'
```

This will show me the deposits that are destined for the test contract that I deployed on L2. At the moment here is the deposit I'm interested in:

```json
{
  "leaf_type": 1,
  "orig_net": 0,
  "orig_addr": "0x3878Cff9d621064d393EEF92bF1e12A944c5ba84",
  "amount": "0",
  "dest_net": 1,
  "dest_addr": "0xC92AeF5873d058a76685140F3328B0DED79733Af",
  "block_num": "7435415",
  "deposit_cnt": 67305,
  "network_id": 0,
  "tx_hash": "0x517b9d827a3a81770d608a6b997e230d992e1e0cabc0fd2797285693b1cc6a9f",
  "claim_tx_hash": "",
  "metadata": "0x40c10f190000000000000000000000003878cff9d621064d393eef92bf1e12a944c5ba84000000000000000000000000000000000000000000000000002386f26fc10000",
  "ready_for_claim": true,
  "global_index": "18446744073709618921"
}
```

I'm going to use this command to try to claim this message on L2.

```bash
polycli ulxly claim message \
    --bridge-address 0x528e26b25a34a4A5d0dbDa1d57D318153d2ED582 \
    --bridge-service-url https://bridge-api.cardona.zkevm-rpc.com \
    --private-key 0x32430699cd4f46ab2422f1df4ad6546811be20c9725544e99253a887e971f92b \
    --destination-address 0xC92AeF5873d058a76685140F3328B0DED79733Af \
    --deposit-network 0 \
    --deposit-count 67305 \
    --rpc-url https://rpc.cardona.zkevm-rpc.com
```

[Here](https://cardona-zkevm.polygonscan.com/tx/0x6df4c4e43776d703bf1996334a4e1975bb3c124192563c93e3d199d9240dd56f#eventlog) is the transaction that was generated by this command. Everything looks to have worked properly. The `MessageReceived(address,uint32,bytes)` event with signature `0xe97c9b3f13b44bc13bde4743ae654dff72f8dc2ff9ff6070efc5999f77a37716` showed up in the explorer so our contract fired properly when the claim was made.

## Flags

```bash
  -h, --help   help for message
```

The command also inherits flags from parent commands.

```bash
      --bridge-address string              the address of the lxly bridge
      --bridge-service-url string          the URL of the bridge service
      --chain-id string                    set the chain id to be used in the transaction
      --config string                      config file (default is $HOME/.polygon-cli.yaml)
      --deposit-count uint                 the deposit count of the bridge transaction
      --deposit-network uint               the rollup id of the network where the deposit was initially made
      --destination-address string         the address where the bridge will be sent to
      --dry-run                            do all of the transaction steps but do not send the transaction
      --gas-limit uint                     force a gas limit when sending a transaction
      --gas-price string                   the gas price to be used
      --global-index string                an override of the global index value
      --pretty-logs                        Should logs be in pretty format or JSON (default true)
      --private-key string                 the hex encoded private key to be used when sending the tx
      --rpc-url string                     the URL of the RPC to send the transaction
      --transaction-receipt-timeout uint   the amount of time to wait while trying to confirm a transaction receipt (default 60)
  -v, --verbosity int                      0 - Silent
                                           100 Panic
                                           200 Fatal
                                           300 Error
                                           400 Warning
                                           500 Info
                                           600 Debug
                                           700 Trace (default 500)
```

## See also

- [polycli ulxly  claim](polycli_ulxly__claim.md) - Commands for claiming deposits on a particular chain