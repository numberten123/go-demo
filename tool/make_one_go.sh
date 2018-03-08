#!/bin/sh
name=${1%%.*}
python makego.py $1
cp -rf output/*$name*.go ../src/test/go_config/