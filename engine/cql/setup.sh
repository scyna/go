#!/bin/bash
host=localhost
username=scylla
password=Scyll@2022#

files='_cleanup.cql generator.cql domain.cql session.cql trace.cql data.cql'


for file in $files
do
    echo ${file}
    #cqlsh ${host} -u ${username} -p ${password} -f ${file}
    cqlsh -f ${file}
done