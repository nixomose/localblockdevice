#!/bin/bash

find .. -name "go.mod" |xargs -n1 sed -i 's-// replace -replace -g'
