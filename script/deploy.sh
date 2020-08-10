#!/bin/bash

cd /data/go/code/src/paperstore/cmd/server/
go install paperstore.go
cp -rf /data/go/code/bin/paperstore /data/paperstore/paperstore

cd /data/paperstore
./paperstore -p 6060 &
