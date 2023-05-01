#/bin/bash

echo "Launching server"
go run ./server/ &

sleep 1

echo "Launching client 1"
go run ./client/ &



echo "Launching client 2"
go run ./client/ &

