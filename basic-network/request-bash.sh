#!/bin/bash

docker exec -e "CORE_PEER_LOCALMSPID=RequestOrg" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/request.test.com/users/Admin@request.test.com/msp" -e "CORE_PEER_ADDRESS=peer0.request.test.com:7051" -it cli bash
