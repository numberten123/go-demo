#!/bin/sh
for i in `ls *.xlsx`
do
    python makego.py $i
done

cp -rf output/*.go ../src/test/go_config/
