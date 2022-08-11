#!/bin/bash
host=dbstg-op.vin3s.vn
username=scylla
password=scylla

files='_cleanup.cql generator.cql module.cql session.cql trace.cql client.cql organization.cql application.cql task.cql data.cql'

for file in $files
do
    echo ${file}
    cqlsh ${host} -u ${username} -p ${password} -f ${file}
done
