#!/bin/bash

cd /data/go/code/src/pagestore/cmd/server/
go install pagestore.go
cp -rf /data/go/code/bin/pagestore /data/pagestore/pagestore

cd /data/pagestore
./pagestore -p 6060 &
