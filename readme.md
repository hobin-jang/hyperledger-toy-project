simple hyperledger fabric toy project

test env

- ubuntu 18.04
- hyperledger fabric v1.4.3
- go 1.11
- docker 20.04
- DB : couchdb
- SDK : Node.js sdk

chaincode : RequestOrg send data, then HashingOrg set hash the data, and search hashed data

- HashingOrg : Only using setHash, getData 
- RequestOrg : Only using setData, getData
- SearchOrg : Only using getData

CA : ca.hashing.test.com

Anchors : peer0 in Orgs

Orderer : orderer.test.com (solo)