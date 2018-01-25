echo "Updating the Image"
docker pull emyrk/factomd_testnet_community:v1

echo "Removing the old container"
docker rm -f factomd_testnet_community_node

echo "Running the new container"
docker run -d -p 8109:22 -p 8090:8090 -p 9876:9876 -p 8110:8110 -v factomd_volume:/root/.factom/m2 --name factomd_testnet_community_node emyrk/factomd_testnet_community:v1