# Ubiot Network

Hyperledger Fabric network for rental machine UbiotThis is a network for purchase and rent of pay per use machines

## Turn on the virtual machine "twinbiot"

* Select the virtual machine an run

https://console.cloud.google.com/compute/instances?project=guminator&folder&organizationId=148075794671&instancessize=50

## Connect to the virtual machine 

* Copy the following command and paste it in your terminal

```
gcloud beta compute ssh --zone "us-central1-a" "rtorrealba@twinbiot" --project "guminator" 
```

## Run network Ubiot

* Navigate to the ubiot network directory with the following command:

```
cd /home/rtorrealba/ubiot-network
```

* Run the startup network script

```
./start-network
```

* Now, copy the printed link in your terminal and paste it into your browser

## Hyperledger explorer
Open a second terminal and connect through shh to the twinbiot virtual machine

```
gcloud beta compute ssh --zone "us-central1-a" "rtorrealba@twinbiot" --project "guminator"
```

* You will need to give permission to the organizations folder

```
sudo chmod 777 -R ubiot-network/test-network/organizations/
```

* Copy the organization folder in the explorer folder

```
cp -r ubiot-network/test-network/organizations/ ubiot-network/explorer/
```

* Then, copy the private key of the Org1 and paste it in the connection file

```
cd ubiot-network/test-network/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/keystore/
ls
```

* Copy the name of the file printed in the terminal and go to the connection/profile directory 

```
cd ubiot-network/explorer/connection-profile/
```

* now edit the first-network.json

```
nano first-network.json
```

Paste the name in the organizations section, specifically in the path, next to the keystore
Save the file, go to the explorer's directory and execute the following command:

```
docker-compose up
```

* Open your browser and go to the virtual machine IP address on 8081 port*

## Update seed and key index in the smart contract
In order to update seed, key index fields in the smart contract and run the network use:
```
./updateSeed.sh -s "SEED" -k KEYINDEX
```
Example:
```
./updateSeed.sh -s "ABCDEFGHIJKLMNOPQRSTUVWXYZ9WOREIURETSEYNDSFNDSFHDS9AHFASFHSFHSANGBFASFHASFHASFH9" -k 9
```
