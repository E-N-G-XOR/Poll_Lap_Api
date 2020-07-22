#!/bin/bash

go get -u github.com/go-sql-driver/mysql

#catch secrets
cd /go/src/print_env/
go run main.go > /tmp/env_dump.txt

#run web app
cd /go/src/app
go run main.go

#tail -f /dev/null