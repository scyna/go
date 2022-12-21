#!/bin/bash

host=localhost
if [ "$1" != '' ]
then
    host=$1
fi

echo "Host: ${host}"
user=cassandra

if [ "$2" != '' ]
then
    user=$2
fi

echo "User: ${user}"
password=cassandra

if [ "$3" != '' ]
then
    password=$3
fi

echo "Password: ${password}"

echo "cleanup"
cqlsh -u ${user} -p ${password} ${host} 9042 -f cleanup.cql
echo "init"
cqlsh -u ${user} -p ${password} ${host} 9042 -f init.cql
echo "template.cql"
cqlsh -u ${user} -p ${password} ${host} 9042 -f template.cql

