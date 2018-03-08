#!/bin/sh

rm -rf output/*
protoc --plugin=../bin/protoc-gen-go --go_out=./output *.proto
cp -rf output/* ../src/test/go_proto/