#!/bin/bash
./build.sh

./lbd  storage-status -t /dev/sdb1 | jq .
