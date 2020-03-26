#!/bin/bash
root=$(cd $(dirname $0);pwd)
cd ${root}
export GOPATH=${root}
export PATH=${PATH}:${root}/bin

go get -v github.com/go-sql-driver/mysql
go get -v github.com/go-redis/redis
go get -v github.com/yuin/gopher-lua
go get -v github.com/satori/go.uuid

# go clean
go install main
