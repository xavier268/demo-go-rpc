#/bin/bash

echo "Launching server"
go run ./server/ &

sleep 1

echo "Launching client"
go run ./client/

echo "Client closed"
