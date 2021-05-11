#!/bin/bash
processName="biudis"
trap "rm ${processName};kill 0" EXIT
lsof -i:8001|grep ${processName} |awk '{print $2}' |xargs kill -9
lsof -i:8002|grep ${processName} |awk '{print $2}' |xargs kill -9
lsof -i:8003|grep ${processName} |awk '{print $2}' |xargs kill -9

go build -o biudis
./${processName} -port=8001 &
./${processName} -port=8002 &
./${processName} -port=8003 -enableApi=1 &
# todo fixed mulil reqeust send
sleep 2
echo ">>> start test"
curl "http://localhost:9999/api?key=Tom" &
curl "http://localhost:9999/api?key=xxx" &
curl "http://localhost:9999/api?key=Tom" &
curl "http://localhost:9999/api?key=Tom" &
curl "http://localhost:9999/api?key=Tom" &
curl "http://localhost:9999/api?key=Tom" &
curl "http://localhost:9999/api?key=Tom" &
curl "http://localhost:9999/api?key=Tom" &
curl "http://localhost:9999/api?key=Tom" &
curl "http://localhost:9999/api?key=Tom" &
curl "http://localhost:9999/api?key=Tom" &
curl "http://localhost:9999/api?key=Jack" &
curl "http://localhost:9999/api?key=Sam" &
curl "http://localhost:9999/api?key=Sam" &
echo ">>> end test"
wait