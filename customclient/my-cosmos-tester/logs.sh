#!/bin/bash
# while loop
destdir=./logs.txt
while true
do
sleep 1
data=$(curl http://localhost:26670/metrics)
VAR_LENGTH=${#data}
if (( VAR_LENGTH > 100 )); then
     echo "$data" > "$destdir"
else
    echo "check latest log in logs.txt file in your current directory"
    break
fi

done
