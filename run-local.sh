#!/usr/bin/env bash

go get github.com/githubnemo/CompileDaemon
pwd

go mod download
go mod tidy

swag init

export PORT=8000
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASS=postgres
export DB_NAME=user
export ENABLE_DB=true

dt=$(date '+%d/%m/%Y %H:%M:%S')
echo "Run dev $dt"

CompileDaemon -log-prefix=false -build="go build -i -x main.go" -command="./main" -exclude-dir=".git"  -exclude-dir=".idea" -exclude-dir="vendor" -color
