#!/bin/bash

grep "^replace" go.mod > /dev/null
if [ "$?" == "0" ]; then
  ./commentoutreplace.sh 
else
  ./uncommentreplace.sh
fi

