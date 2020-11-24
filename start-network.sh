cd /home/alejandro/ubiot-network/test-network

./network.sh down

docker rm -f $(docker ps -aq)

docker volume prune -f

./network.sh down

docker ps

docker ps -a

docker volume ls

# Deploying a smart contract to a channel
./network.sh down

./network.sh up createChannel -ca

#./monitordocker.sh net_test


# Vendor 
#GO111MODULE=on go mod vendor


export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
peer version

#Package chaincode
peer lifecycle chaincode package basic.tar.gz --path ../smart-contract/contract-tutorial/ --lang golang --label basic_1.0

#Org1
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

peer lifecycle chaincode install basic.tar.gz

#Org2
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_ADDRESS=localhost:9051

peer lifecycle chaincode install basic.tar.gz

out=$(peer \
lifecycle \
chaincode \
queryinstalled \
| \
cut \
-d \
':' \
-f \
2,3 \
| \
cut \
-d \
',' \
-f \
1 )


export CC_PACKAGE_ID=$out


peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID mychannel --name basic --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem


#Org1
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_ADDRESS=localhost:7051

peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID mychannel --name basic --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem


peer lifecycle chaincode checkcommitreadiness --channelID mychannel --name basic --version 1.0 --sequence 1 --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --output json


peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID mychannel --name basic --version 1.0 --sequence 1 --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt


peer lifecycle chaincode querycommitted --channelID mychannel --name basic --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

# sudo chmod 777 -R /home/alejandro/ubiot-network/organizations
# cp -r /home/alejandro/ubiot-network/organizations explorer/


echo 'httpOptions:
      verify: false

channels:
  mychannel:
    orderers:
      - orderer.example.com
    peers:
      peer0.org1.example.com:
        endorsingPeer: true
        chaincodeQuery: true
        ledgerQuery: true
        eventSource: true

orderers:
  orderer.example.com:
    url: grpcs://localhost:7050
    tlsCACerts:
      path: "/home/alejandro/ubiot-network/test-network/organizations/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem"

    grpcOptions:
      ssl-target-name-override: orderer.example.com' >> /home/alejandro/ubiot-network/test-network/organizations/peerOrganizations/org1.example.com/connection-org1.yaml

echo 'httpOptions:
      verify: false

channels:
  mychannel:
    orderers:
      - orderer.example.com
    peers:
      peer0.org1.example.com:
        endorsingPeer: true
        chaincodeQuery: true
        ledgerQuery: true
        eventSource: true

orderers:
  orderer.example.com:
    url: grpcs://localhost:7050
    tlsCACerts:
      path: "/home/alejandro/ubiot-network/test-network/organizations/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem"

    grpcOptions:
      ssl-target-name-override: orderer.example.com' >> /home/alejandro/ubiot-network/test-network/organizations/peerOrganizations/org2.example.com/connection-org2.yaml

sudo rm -rf /home/alejandro/ubiot-network/app/keystore
sudo rm -rf /home/alejandro/ubiot-network/app/wallet

MYIP=`curl checkip.amazonaws.com`
PORT=":8081"
echo "https://$MYIP$PORT"

cd /home/alejandro/ubiot-network/app
go run server.go
