#!/bin/bash
host=db-oneprofile.vin3s.vn
username=cassandra
password=cassandra

files='init.cql'

for file in $files
do
    echo ${file}
    cqlsh ${host} -u ${username} -p ${password} -f ${file}
done
