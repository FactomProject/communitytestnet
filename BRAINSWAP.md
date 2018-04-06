# Brainswap

Brainswapping is the idea of having 2 nodes switch identities at the same time. We call is "brainswapping" because a node's identity dictates how it behaves.

It was implemented as a way to update the network without having to bring it down. All federated servers could "brainswap" with standby nodes that have already been updated. To the network, the federated server never went down, but to the node operator, they can now shutdown their federated node, perform updates, then bring it back online and brainswap the identity back into the original server.

## Definitions

- __Original Node__ is the node that holds your authority identity, and needs to be updated
- __Standby Node__ is a follower node on the network that you control

## Prepping the Brain Swap

To perform the brain swap you will need 1 standby node that is _ready_. It might be a good idea to have the standby node on the most recent update.

### Determining if you Standby Node is _ready_

If your standby node is not in sync with the network, performing the brainswap will result in your authority node going offline, so it is crucial to first check the health of the standby node.

1. Check if the DBHeight matches that of the network
2. Check if the minutes are following that of the network
3. Check the process list for `<nil>`, that indicates some network instability

## Performing the Brain Swap (Read through before performing)

Once your standby node is ready, we can prep for the swap. The swap requires both config files (original and stanby node files) to be modified.

Open both config files up in a text editor of your choice, also it is wise to have a control panel open to watch the block heights.

Swap the following lines: (Original --> Standby)

```
IdentityChainID	                      = FA1E000000000000000000000000000000000000000000000000000000000000
LocalServerPrivKey                = 4c38c72fc5cdad68f13b74674d3ffb1f3d63a112710868c9b08946553448d26d
LocalServerPublicKey            = cc1985cdfae4e32b5a454dfda8ce5e1361558482684f3367649c3ad852c8e31a
```

(It is not needed to swap Standby --> Original, so you can just comment out the lines in your original node by placing a `;` in front of these lines)

And add an additional line:

```
ChangeAcksHeight                      = 0
```

This additional line is the brainswap logic. You will want to set the `ChangeAcksHeight` to some block height in the future (remember the block height on the control panel is the last saved height, not the current working height!). The safe height is the one you see in the control panel +3. If you know how to read the more detailed page, you can get away with a closer number.

Once you set the `ChangeAcksHeight` to DBHeight+3, save both files.

At the block height `ChangeAcksHeight` you should see both nodes change identities. If none of your nodes crash, and the identities change, the swap was successful.