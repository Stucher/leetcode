#!/usr/bin/env bash
if [[ $# == 0 ]]; then
	echo "请输入要编译的程序的编号"
	exit -1
fi

NUM=${1}
if [[ ${NUM} -le 0 ]];then
	echo "要编译的文件编号非法"
	exit -2
fi

if [[ ${1} -lt 10 ]];then
	NUM="0${NUM}"
fi

CODE_FILE="leet_code_${NUM}.go"
LL=$(ls -al | grep "${CODE_FILE}")
if [[ "${LL}" == "" ]]; then
	echo "go文件找不到: ${CODE_FILE}"
	exit -3
fi

go build -o ./bin/main main.go ${CODE_FILE}
./bin/main