#!/bin/bash

docker-compose up -d \
orderer.test.com peer0.hashing.test.com peer1.hashing.test.com \
peer0.search.test.com peer1.search.test.com \
peer0.request.test.com peer1.request.test.com \
cli \
couchdb1 couchdb2 couchdb3 couchdb4 couchdb5 couchdb6

docker-compose -f docker-compose-ca.yaml up -d ca.hashing.test.com
