#!/bin/bash
root=$(cd $(dirname $0);pwd)
cd ${root}
export GOPATH=${root}
export PATH=${PATH}:${root}/bin

go get -v github.com/golang/protobuf/protoc-gen-go
go get -v github.com/golang/protobuf/proto
go get -v github.com/yuin/gopher-lua
go get -v github.com/gorilla/websocket

proto3=${root}/bin/proto-3
mkdir -p ${proto3}
protoc=${root}/bin/protoc
if [ ! -f "${protoc}" ];then
	cd ${proto3}
	if [ ! -f "${proto3}/protoc-3.11.4-linux-x86_64.zip" ];then
		wget https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip
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
protodir=${root}/src/proto
cd ${protodir}
for dir in `ls`
do
	if [ -d "${protodir}/${dir}" ];then
		cd ${protodir}/${dir}
		for p in `ls | grep '.proto'`
		do
			${protoc} --go_out=. ${p}
		done
	fi
done

# 将proto协议转换为json配置，对协议名做一个映射，方便网络收发的时候做解析
protoDir=${protodir}
mkdir -p ${protoDir}/protomp
msgprotocol=${protoDir}/protomp/protomp.go
echo "package protomp
var protoMap map[string]string
func init(){
	protoMap = map[string]string{" > ${msgprotocol}

fileList=(`find ${protoDir} | grep -E '*\.proto' | tr '\n' ' '`)
fileNum=${#fileList[@]}
let fileNum=fileNum-1
# echo $fileNum
for i in "${!fileList[@]}";   
do  
	StartIndex=$i
	endFlag=0
	if [ "${fileNum}" -eq "${i}" ];then
		endFlag=1
	fi
	let StartIndex=(StartIndex+1)*1000
	file=${fileList[$i]} 
	protoName=`echo $file | awk '{
		lena=split($0,a,"/");
		lenb=split(a[lena],b,".");
		print b[1];
	}'`
	num=`grep -v "//" $file | grep "message.*{" |wc -l`
 	grep -v "//" $file | grep "message.*{" | tr '{' ' ' | awk 'BEGIN{
		i="'"${StartIndex}"'";
		flag="'"${endFlag}"'";
		n="'"${num}"'";
	}
	{
		print "\t\t\""i"\":\"'$protoName'."$2"\",";
		# if (flag==1 && NR==n)
		# 	print "\"""'"$protoName"'""."$2"\":\""i"\"";
		# else
			print "\t\t\"""'"$protoName"'""."$2"\":\""i"\",";
		i++;
	}' >> ${msgprotocol}
done 

echo '	}
}
func GetProto(key string) (string, bool){
	if value, ok := protoMap[key];ok{
		return value, true
	} else {
		return "", false
	}
}' >> ${msgprotocol}

cd ${root}

# go clean
go install main
