#/bin/bash

echo "Launching server"
PID= go run ./server/ &

sleep 1

echo "Launching client"
go run ./client/

echo "Killing server"
kill $PID

netstat -plt