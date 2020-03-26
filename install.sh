#!/bin/bash
root=$(cd $(dirname $0);pwd)
cd ${root}

install(){
    cd ${root}/$1
	rm -rf bin src/github.com
	mkdir -p bin/proto-3
	cp ../packages/protoc-3.11.4-linux-x86_64.zip bin/proto-3
	cp -a ../packages/github.com src/
	sed -i 's/\r//g' install.sh
}

case $1 in
    auth | service)
		install $1
		cd ${root}/$1
		bash install.sh
        ;;
    database)
		install $1
		cd ${root}/$1
		rm -rf src/golang.org   src/google.golang.org
		cp -a ../packages/golang.org src/
		cp -a ../packages/google.golang.org src/
		cd ${root}/$1
		bash install.sh
		;;
	clean)
		cd ${root}
		fileList=(auth client database demo service testRpc)
		for i in "${!fileList[@]}";
		do
			dir=${fileList[$i]}
			if [ -d "${dir}" ];then
				rm -rf ${dir}/bin  ${dir}/pkg  ${dir}/src/github.com  ${dir}/src/google.golang.org ${dir}/src/golang.org
			fi

		done
		;;
	*)
		echo "Unknow"
esac




