#!/bin/bash

docker-compose start \
orderer.test.com peer0.hashing.test.com peer1.hashing.test.com \
peer0.search.test.com peer1.search.test.com \
peer0.request.test.com peer1.request.test.com \
cli \
couchdb1 couchdb2 couchdb3 couchdb4 couchdb5 couchdb6

