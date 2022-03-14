#!/bin/bash

docker exec -e "CORE_PEER_LOCALMSPID=SearchOrg" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/search.test.com/users/Admin@search.test.com/msp" -e "CORE_PEER_ADDRESS=peer0.search.test.com:7051" -it cli bash
