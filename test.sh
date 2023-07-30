#!/bin/bash


go build -o main main.go

echored(){
	echo -e "\033[31m$1\033[0m"
}
echogreen(){
	echo -e "\033[32m$1\033[0m"
}

testfunc(){
	echo "1: $1"
	output=$1
	expected=$2
	if [ "$output" == "$expected" ]; then
		echogreen "Test passed"
	else
		echored "Test failed"
		echored "	Expected: $expected"
		echored "	Got: $output"
	fi
}

testfunc "`./main -f testfile/testcase1.rb -s '#'`" "Comment1 is here.

And comment2 is here."


