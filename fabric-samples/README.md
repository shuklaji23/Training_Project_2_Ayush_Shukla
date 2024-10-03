#CLI commands for project 1

##Create Folder and Download Fabric Samples:

    mkdir hyperledger 
    cd hyperledger 
    curl -sSL https://bit.ly/2ysbOFE | bash -s -- 2.4.9 1.5.2 

##Test if the Network is running or not:

    cd fabric-samples/test-network 
    ./network.sh up 
    docker ps 
    ./network.sh createChannel 
    ./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-go -ccl go 
    ./network.sh down 

##Stop and delete all the containers

    docker stop $(docker ps -aq) 
    docker rm $(docker ps -aq) 
    docker volume prune 

##Change the org names in the respective files:
    1. crypto-config-org1.yaml
    2. crypto-config-org2.yaml
    3. compose-couch.yaml
    4. compose-test-net.yaml
    5. docker-compose-couch.yaml
    6. docker-compose-test-net.yaml

##Generate certificate using CryptoGen tool:

    export PATH=${PWD}/../bin:${PWD}:$PATH 
    cryptogen generate --config=./organizations/cryptogen/crypto-config-org1.yaml --output="organizations" 
    cryptogen generate --config=./organizations/cryptogen/crypto-config-org2.yaml --output="organizations" 
    cryptogen generate --config=./organizations/cryptogen/crypto-config-orderer.yaml --output="organizations" 

##Start the Docker Container:

    export COMPOSE_PROJECT_NAME=net  
    export DOCKER_SOCK=/var/run/docker.sock 
    IMAGE_TAG=latest docker-compose -f compose/compose-test-net.yaml -f compose/docker/docker-compose-test-net.yaml up 

##Open a new terminal (test-network folder):

    docker ps

##Set Config Path:

    export PATH=${PWD}/../bin:${PWD}:$PATH 
    export FABRIC_CFG_PATH=${PWD}/configtx 
    export CHANNEL_NAME=mychannel 

##Change the configtx.yaml file

##Write again in CLI:

    configtxgen -profile TwoOrgsApplicationGenesis -outputBlock ./channel-artifacts/${CHANNEL_NAME}.block -channelID $CHANNEL_NAME
    configtxgen -inspectBlock ./channel-artifacts/mychannel.block > dump.json 
    cp ../config/core.yaml ./configtx/. 
    export ORDERER_CA=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem 
    export ORDERER_ADMIN_TLS_SIGN_CERT=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/server.crt 
    export ORDERER_ADMIN_TLS_PRIVATE_KEY=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/server.key 
    osnadmin channel join --channelID $CHANNEL_NAME --config-block ./channel-artifacts/${CHANNEL_NAME}.block -o localhost:7053 --ca-file "$ORDERER_CA" --client-cert "$ORDERER_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER_ADMIN_TLS_PRIVATE_KEY" 

##Update envVar.sh script file inside scripts folder with new org names 

    source ./scripts/setOrgPeerContext.sh 1 
    peer channel join -b ./channel-artifacts/${CHANNEL_NAME}.block 
    source ./scripts/setOrgPeerContext.sh 2 
    peer channel join -b ./channel-artifacts/${CHANNEL_NAME}.block 
    source ./scripts/setOrgPeerContext.sh 1 
    docker exec cli ./scripts/setAnchorPeer.sh 1 $CHANNEL_NAME 
    source ./scripts/setOrgPeerContext.sh 2 
    docker exec cli ./scripts/setAnchorPeer.sh 2 $CHANNEL_NAME 

##Go to the Chaincode Folder:

    go mod init simple-investment-application-chaincode 
    go get -u github.com/hyperledger/fabric-contract-api-go 
    <!-- chain the go version to 1.13 in go.mod -->
    go mod tidy 
    go mod vendor 
    go build 

Go back to the test-network folder

##Create a script file setOrgPeerContext.sh and write the following:

    #!/bin/bash 
    # import utils 
    source scripts/envVar.sh 
    ORG=$1  
    setGlobals $ORG 

Write in the CLI again:

    export PATH=${PWD}/../bin:${PWD}:$PATH
    export FABRIC_CFG_PATH=${PWD}/configtx
    export CHANNEL_NAME=mychannel

    source ./scripts/setChaincodeContext.sh
    source ./scripts/setOrgPeerContext.sh 1 
    peer lifecycle chaincode package iac-chaincode.tar.gz --path ${CC_SRC_PATH} --lang ${CC_RUNTIME_LANGUAGE} --label iac-chaincode_${VERSION} 
    peer lifecycle chaincode install iac-chaincode.tar.gz 
    source ./scripts/setOrgPeerContext.sh 2 
    peer lifecycle chaincode install iac-chaincode.tar.gz 
    peer lifecycle chaincode queryinstalled 2>&1 | tee outfile 

Create a file setPackageID.sh in script folder and write the following:

    #!/bin/bash
    PACK_ID=$(sed -n "/iac-chaincode_${VERSION}/{s/^Package ID: //; s/, Label:.*$//; p;}" $1) 
    export PACKAGE_ID=$PACK_ID 
    echo $PACKAGE_ID 

Write again in CLI:

    source ./scripts/setPackageID.sh outfile 
    source ./scripts/setOrgPeerContext.sh 1 
    peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name iac-chaincode --version ${VERSION} --init-required --package-id ${PACKAGE_ID} --sequence ${VERSION} 
    peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name iac-chaincode --version ${VERSION} --sequence ${VERSION} --output json --init-required 
    source ./scripts/setOrgPeerContext.sh 2
    peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name iac-chaincode --version ${VERSION} --sequence ${VERSION} --output json --init-required
    peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name iac-chaincode --version ${VERSION} --init-required --package-id ${PACKAGE_ID} --sequence ${VERSION}
    peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name iac-chaincode --version ${VERSION} --sequence ${VERSION} --output json --init-required
    source ./scripts/setOrgPeerContext.sh 1
    peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name iac-chaincode --version ${VERSION} --sequence ${VERSION} --output json --init-required

Update the envVar.sh file

Write in the CLI again:

    source ./scripts/setPeerConnectionParam.sh 1 2
    peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name iac-chaincode $PEER_CONN_PARAMS --version ${VERSION} --sequence ${VERSION} --init-required
    peer lifecycle chaincode querycommitted --channelID $CHANNEL_NAME --name iac-chaincode
    source ./scripts/setOrgPeerContext.sh 1
    peer lifecycle chaincode querycommitted --channelID $CHANNEL_NAME --name iac-chaincode


CHAINCODE EXECUTION

source ./scripts/setPeerConnectionParam.sh 1 2
source ./scripts/setOrgPeerContext.sh 1
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n fabcar $PEER_CONN_PARAMS --isInit -c '{"function":"InitLedger","Args":[]}'
