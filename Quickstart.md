# Quickstart

## Prerequisites

To create the environment install the following components first:
 - [docker](https://www.docker.com/community-edition)
 - [docker-compose](https://docs.docker.com/compose/install/)
 - [git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

## Running factomd

Before running the quickstart, make sure no factomd instance is being run on your machine.

The quickstart commands.

```
# Pull the repo into a local directory
cd ~
mkdir factom && cd factom
git clone https://github.com/FactomProject/communitytestnet.git
cd communitytestnet

# Start docker containers
docker pull emyrk/factomd_testnet_community:v1
docker-compose up -d

# Setup volume
cp factomd.conf.EXAMPLE factomd.conf
docker run --rm -v ${PWD}/factomd.conf:/source -v communitytestnet_factomd_volume:/destination busybox /bin/cp /source /destination/factomd.conf

# Start factomd
docker exec factomd_node bash /root/bin/start.sh
```

You are now running on the community testnet. To check if your docker containers are indeed running:
```
docker ps
```

You now visit:
* localhost:8090 for the control panel
* localhost:3001 for Grafana (user: admin | pass: admin)
* localhost:9090 for Prometheus

## Cleanup/stopping

To stop factomd

```
docker exec factomd_node bash /root/bin/stop.sh
```

To stop all the containers

```
docker-compose down
```

# Authority Nodes

Authority nodes need to create an identity in order to be eligble to run a Federated/Audit server. Authority nodes also need to open certain ports to the outside world.

## Ports

Authority nodes need to have 2 ports exposed to the public:

- 8110 : Peer discovery port
- 8220 : Allow ssh access for Factom Developers

Firewall rules specific to the docker container may be implmented to control network topology for security or test reasons.

### Testing open ports

To test your ports are exposed, first ensure factomd is running, then with the netcat tool try the following:

```
nc PUBLIC_IP 8110
```

You should see the following:

```
,��Parcel��Header��Payload
����
NetworkVersionTypeLength
...
```

This means your peer discovery port is open. You can also try to ssh into the machine (it will reject you)

```
ssh root@PUBLIC_IP -p 8220
```

### Additional Ports

Additional ports that your host can use and expose in whatever way you wish.

- 8090 : Control Panel
- 3001 : Grafana

## Creating an Identity

### Prerequisites

* Quickstart guide above complete
* An entry credit address with ECs

### Creating the Identity

This will take up 2 windows. First, you need to have an entry credit address with ECs in it. This will have to be provided for you. For demonstration, this EC address will be used:

Sec : Es37BZSs7jUpyn3HosZa79kENWfvj1AVUdZioWykTTqqvA2MRi9h  
Pub : EC2KnJQN86MYq4pQyeSGTHSiVdkhRCPXS3udzD4im6BXRBjZFMmR

In the first window we will launch the wallet
```
docker exec -it factomd_node factom-walletd
```

Everything from here forward will be in the second window. We need to import the EC address given above
```
docker exec factomd_node factom-cli importaddress Es37BZSs7jUpyn3HosZa79kENWfvj1AVUdZioWykTTqqvA2MRi9h
```

An identity needs to be brute forced, this program will brute force a valid identity and print out the relevent information you need to save. It will also provide you with a script to publish your identity to the blockchain.

```
docker exec factomd_node serveridentity full elements Es37BZSs7jUpyn3HosZa79kENWfvj1AVUdZioWykTTqqvA2MRi9h -n create -f
```

The output will look something like this...

```
Key Generation Complete

Private keys and their corresponding levels. Copy these down, this program will 
not save them for you.
Level 1: sk12z9Cqv5ByQjvTiLjF6gNNQkAdMS9XcazheL9wqgBGw7E9VBPsW
Level 2: sk23oicLcAFX16qKZUVG5TtnJbFCVgjStHH4hTtwcRGr361wJXmKG
Level 3: sk32YLCrDtUPMRCJkMk6TYFta1suu5WgEmhWtVqbARWmTYS5LtH1S
Level 4: sk43SqpqT3LgQYvbpcSfznywcxS6hAfV5epNcCTVpL7PViBXqE9h3

This is the entry key that will be used to add entries/chains to the Factom network. 
Copy these keys down, this program will not save them for you.
Note: Entry credits must be in the wallet before commands are executed.
Private Key: Es37BZSs7jUpyn3HosZa79kENWfvj1AVUdZioWykTTqqvA2MRi9h
Public Key : EC2KnJQN86MYq4pQyeSGTHSiVdkhRCPXS3udzD4im6BXRBjZFMmR

***********************************************************************
**                   Creating Identity Chains/Keys                   **
***********************************************************************
................
...................................
Root Chain: 888888c0754c3218a32d12004fd9d30590ccbc22954f8688e1a9d628f943be80
Management Chain: 888888c0754c3218a32d12004fd9d30590ccbc22954f8688e1a9d628f943be80

block signing private key: 63e16a00f22a38e8a007e9a73c7e532ec7c008fbbb51dfd33b698380e4c3de42
block signing public key: aaeb58ed8fa921b8139da4c5a9d2bf9a5ee7d307551e69eb1d9be5389761e17b

BTC Key: 06e737e489ef4d345818f81891bc5c20ea55fe7f
MHash Seed (hex): 65323834386364376438393530623639613433646331333461663365313465373332643566613332
***********************************************************************
**                        Factom-cli commands                        **
***********************************************************************
```

You will want to copy down the Level 1 to 4 keys, the Entry credit address and the "Creating Identity/Keys" section. All this information will be needed to claim your identity in the future. The script that is created resides in the docker container called 'create.sh'. Let's change some permissions on it and execute it.

```
docker exec factomd_node chmod 766 /root/create.sh
docker exec factomd_node bash /root/create.sh
```

Now your identity is in the blockchain, to verify we can grab your root chain information: (To the human eye it's mostly gibberish)

```
# Replace the ChainID with your root chain id seen above
docker exec factomd_node factom-cli get allentries 888888c0754c3218a32d12004fd9d30590ccbc22954f8688e1a9d628f943be80
```

Your node is not yet configured to use this new identity. To do so, we need to modify your factomd.conf file. Let's change it inside of the docker container. (This will replace the lines necessary in factomd.conf with the identity created)

```
docker exec factomd_node  bash -c "sed -i '/Node Identity Information/q' /root/.factom/m2/factomd.conf && grep Identity -A 2 create.conf >> /root/.factom/m2/factomd.conf"
```

Now your node is configured, we can reboot the node to use the identity.

```
docker exec factomd_node bash /root/bin/stop.sh
docker exec factomd_node bash /root/bin/start.sh
```

### Identity Creation Recap

Here is all the commands to create an identity recapped

```
docker exec -it factomd_node factom-walletd

# Different terminal window than the command above
docker exec factomd_node factom-cli importaddress Es37BZSs7jUpyn3HosZa79kENWfvj1AVUdZioWykTTqqvA2MRi9h
docker exec factomd_node serveridentity full elements Es37BZSs7jUpyn3HosZa79kENWfvj1AVUdZioWykTTqqvA2MRi9h -n create -f
docker exec factomd_node chmod 766 /root/create.sh
docker exec factomd_node bash /root/create.sh
docker exec factomd_node  bash -c "sed -i '/Node Identity Information/q' /root/.factom/m2/factomd.conf && grep Identity -A 2 create.conf >> /root/.factom/m2/factomd.conf"
docker exec factomd_node bash /root/bin/stop.sh
docker exec factomd_node bash /root/bin/start.sh
```


# Monitoring Tools

Factomd comes with a few ways to monitor your node's health. The most obvious tool is the control panel found at localhost:8080. Information about the control panel can be found here https://docs.factom.com/#factoid-live.

Factomd also comes with more monitoring tools that are included in this docker setup called Prometheus and Grafana. We will focus on Grafana, as this is the visualization tool that is most usful. To see Grafana, visit http://localhost:3001.

## Setting up Grafana

1. First ensure you have Grafana open by visiting http://localhost:3001
2. The username is "admin" and password "admin". Be sure not to open this port to the world or anyone can log in!
3. You will be greeted by a page that has a button saying "Add data source". Click that
    - Name: `Prometheus`
    - Type: `Prometheus`
    - Ensure `Default` is checked
    - URL: `http://prometheus:9090`
4. Once all the fields are put in, click "Save and Test". You should be greeted by "Data source is working". If you encounter an error, here are some debug steps before asking for assistance.
    - Ensure Prometheus is running `docker ps | grep prometheus`
        - If there are no results, then you did not successfully run `docker-compose up -d`. Try running `docker-compose down` then `docker-compose up -d`
5. Now we have Prometheus as a datasource, let's get some graphs up. In the top left is the Grafana logo, click it, then dashboard, then import
    - Top Left > Dashboard > Import
6. A preconfigured dashboard can be found here: https://grafana.com/dashboards/4482 
	- Input `4482` into the first input, then click `Load`
	- Make sure to select your prometheus source we added earlier from the dropdown menu next to `Prometheus`
	- Click `import` and your dashboard is now viewable.
7. Feel free to mess around and change things to your liking.
