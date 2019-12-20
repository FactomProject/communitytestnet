# Brainswap

Brainswapping is the idea of having 2 nodes switch identities at the same time. We call is "brainswapping" because a node's identity dictates how it behaves.

Note: The procedure doesn't actually have to be a "swap"; a "brain-transfer" is also an alternative. 

It was implemented as a way to update the network without having to bring it down, as federated servers can "brainswap" with standby nodes that have already been updated with the new code. 

The network will not perceive this as a node going offline, as the identity (and thus the associated federated server) is still online.

After transferring the __Authority identity__ the node operater now shutdown their original node, perform necessary updates, bring it back online and finally brainswap the identity back into the original server.

The procedure can also be used for migrating the __Authority identity__ to a new physical server, by not performing the brainswap a second time to reverse the first swap.

## Definitions

- __Federated Node__ is the node that holds your authority identity, and needs to be updated.
- __Standby Node__ is a follower node on the network that you control.

## Prepping the Brain Swap

To perform the brain swap you will need 1 standby node that is _ready_. The standby node should be running the most recent Factomd software version.

### Determining if you Standby Node is _ready_

If your standby node is not in sync with the network, performing the brainswap will result in your authority server going offline, so it is crucial to first check the health of the standby node.

1. Check if the DBHeight matches that of the network 
2. Check if the minutes are following that of the network
3. Check the process list for `<nil>`, that indicates some network instability
_(Procedures for the above described at the bottom of this document)_

## Performing the Brain Swap (Read through before performing)

Once your standby node is ready, we can prep for the swap. The swap requires both config files (located on Federated Node and Standby Node) to be modified.

Open both config files in parallel in a text editor of your choice. 
_(Procedures for editing is described at the bottom of this document)_

Swap the following lines in the two config files:

```
IdentityChainID	                      = FA1E000000000000000000000000000000000000000000000000000000000000
LocalServerPrivKey                = 4c38c72fc5cdad68f13b74674d3ffb1f3d63a112710868c9b08946553448d26d
LocalServerPublicKey            = cc1985cdfae4e32b5a454dfda8ce5e1361558482684f3367649c3ad852c8e31a
```

(If you prefer to do a __brain-transfer__ instead of a swap, you can just comment out the lines in your Federated node by placing a `;` in front of these lines.)

Then add an additional line:

```
ChangeAcksHeight                      = 0
```

This additional line is the brainswap logic. You will want to set the `ChangeAcksHeight` to some block height in the future (remember the block height on the control panel is the last saved height, not the current working height!). 

The safe height is the one you see in the control panel +3. (localhost:8090).
If you know how to read the more detailed page, you can get away with a closer number.

Once you set the `ChangeAcksHeight` to __DBHeight+3__, save both files.

At the block height `ChangeAcksHeight` you should see both nodes change identities. If none of your nodes crash, and the identities change, the swap was successful.




## Detailed instructions for some of the above aspects

__Check if the DBHeight matches that of the network__
-- This is done by comparing DBHeight in your control panel (localhost:8090) with the dbheight of your federated server.

__Check if the minutes are following that of the network__
-- This is done by comparing the control panel of the Stanby node to that of the Federated. On the summary tab of the `More Detailed Node Information` you will see this:

```
===SummaryStart===
   FNode04[f0b7e3] L___vm01  0/ 0  0.0%  0.000   165[e0b9f8] 163/166/167  7/ 7         0/0/0/0                43400/0/0/0      0     0     2/40/100           0/0/0   0.07/0.00 0/0 - 309415
```

The `7/7` means you are on minute 7. You will want to make sure this number is the same on the Standby (`_/7`) and the Federated (`7/7`)

__Check the process list for `<nil>`, that indicates some network instability__
-- Process list is located in the control panel (localhost:8090) -> "more detailed node information". If any entries show <nil> you should not move on with the brainswap.
  
__Modify config file__ -- You must edit the factomd config file while the server is running and wait for the changes to be picked up. Don't reboot the server. The config file can be found in the factom_keys docker volume which on debian based systems will be found in `/var/lib/docker/volumes/factom_keys/_data/factomd.conf`

__If you want to use vim__
```
vi /var/lib/docker/volumes/factom_keys/_data/factomd.conf
```

Edit the files per the instructions above.  
Save & exit

__If you want to use nano__
```
nano /var/lib/docker/volumes/factom_keys/_data/factomd.conf
```

Edit the files per the instructions above.  
Save the file: ```ctrl+O```  
Exit nano: ```ctrl+x```
  
  
