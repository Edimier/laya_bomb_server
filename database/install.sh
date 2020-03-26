#!/bin/bash
root=$(cd $(dirname $0);pwd)
cd ${root}
export GOPATH=${root}
export PATH=${PATH}:${root}/bin

go get -v github.com/go-sql-driver/mysql
go get -v github.com/golang/protobuf/protoc-gen-go
go get -v github.com/golang/protobuf/proto
go get -v github.com/yuin/gopher-lua

proto3=${root}/bin/proto-3
mkdir -p ${proto3}
protoc=${root}/bin/protoc
if [ ! -f "${protoc}" ];then
	cd ${proto3}
	if [ ! -f "${proto3}/protoc-3.11.4-linux-x86_64.zip" ];then
		wget -c -O protoc-3.11.4-linux-x86_64.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip
	fi
	unzip protoc-3.11.4-linux-x86_64.zip
	cp bin/protoc ${root}/bin
	[ ! -f "${protoc}" ] && echo "${protoc}not exist, please download" && exit 1
fi

protocgengo=${root}/bin/protoc-gen-go
if [ ! -f "${protocgengo}" ];then
	echo "${protocgengo} not exist, please excute:go get -v github.com/golang/protobuf/protoc-gen-go"
	exit 2
fi

# protodir
# protodir=${root}/src/proto
# cd ${protodir}
# for dir in `ls`
# do
# 	if [ -d "${protodir}/${dir}" ];then
# 		cd ${protodir}/${dir}
# 		for p in `ls | grep -E '*.proto'`
# 		do
# 			${protoc} --go_out=. ${p}
# 		done
# 	fi
# done

# bash ${root}/parses.sh ${protodir}

protodir=${root}/src/proto/rpcdatabase
cd ${protodir}

protoc --go_out=plugins=grpc:. rpcdatabase.proto

cd ${root}

# go clean
go install main
