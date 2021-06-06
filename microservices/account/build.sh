#!/usr/bin/env bash
cd ../../
RootDir=`pwd`
echo $RootDir
scp -r ${RootDir}/microservices op@172.16.20.96:~/develop/lab
scp  ${RootDir}/go.mod op@172.16.20.96:~/develop/lab/go.mod

#kubectl apply -f kubernetes.yml
#kubectl delete -f kubernetes.yml