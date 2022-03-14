#!/bin/bash

docker-compose stop \
orderer.test.com peer0.hashing.test.com peer1.hashing.test.com \
peer0.search.test.com peer1.search.test.com \
peer0.request.test.com peer1.request.test.com cli
