#!/usr/bin/env bash

for file in $@;
do
#	if [[ ${file} != 'main.go' ]]; then
#	mv ${file} "leet_code_00"`echo ${file} | awk -F"_." '{print $3}'`"go"
#	fi
	mv ${file} `echo ${file} | sed -e 's/00.\.go00/00/g'`
done