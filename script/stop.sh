#!/bin/sh

 for i in `ps aux | grep ./test |awk '{print $2}'`
 do
 	kill -9 $i
 done
