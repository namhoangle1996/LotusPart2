#!/usr/bin/env bash

go get github.com/githubnemo/CompileDaemon
pwd

go mod download
go mod tidy

swag init

dt=$(date '+%d/%m/%Y %H:%M:%S')
echo "Run dev $dt"

CompileDaemon -log-prefix=false -build="go build  -x main.go" -command="./main" -exclude-dir=".git"  -exclude-dir=".idea" -exclude-dir="vendor" -color
