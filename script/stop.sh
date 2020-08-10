#!/bin/bash

PORT=6060

PID=$(lsof -t -i:$PORT)

if [ ! -z "$PID" ]
then
  kill $PID
fi
