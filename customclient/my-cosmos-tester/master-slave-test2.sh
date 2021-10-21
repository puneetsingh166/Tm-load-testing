#!/bin/bash
IP=$1
my-cosmos-tester master --expect-slaves 1 --bind localhost:26670 -c 1 -T 15 -r 200 -s 2500 --broadcast-tx-method async --endpoints ws://$IP:26657/websocket &
my-cosmos-tester slave --master ws://localhost:26670 &
destdir=./logsTest2.txt
while true
do
sleep 1
data=$(curl http://localhost:26670/metrics)
VAR_LENGTH=${#data}
if (( VAR_LENGTH > 100 )); then
     echo "$data" > "$destdir"
else
    echo "check latest log in logsTest2.txt file in your current directory"
    break
fi

done
