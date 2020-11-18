# Ubiot Network
Hyperledger Fabric network for rental machine UbiotThis is a network for purchase and rent of pay per use machines

## Turn on the virtual machine "twinbiot"

Select the virtual machine an run
https://console.cloud.google.com/compute/instancesproject=guminator&instancessize=50

##Connect to the virtual machine 

Copy the following command and paste it in your terminal
gcloud beta compute ssh --zone "us-central1-a" "rtorrealba@twinbiot" --project "guminator" 

## Run network Ubiot
Navigate to the ubiot network directory with the following command:
cd /home/rtorrealba/ubiot-network

Run the startup network script
./start-network

Now, copy the printed link in your terminal and paste it into your browser