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

echo "_cleanup.cql"
cqlsh -u ${user} -p ${password} ${host} 9042 -f _cleanup.cql
echo "application.cql"
cqlsh -u ${user} -p ${password} ${host} 9042 -f application.cql
echo "client.cql"
cqlsh -u ${user} -p ${password} ${host} 9042 -f client.cql
# echo "event_store.cql"
# cqlsh -u ${user} -p ${password} ${host} 9042 -f event_store.cql
echo "generator.cql"
cqlsh -u ${user} -p ${password} ${host} 9042 -f generator.cql
echo "module.cql"
cqlsh -u ${user} -p ${password} ${host} 9042 -f module.cql
echo "organization.cql"
cqlsh -u ${user} -p ${password} ${host} 9042 -f organization.cql
echo "session.cql"
cqlsh -u ${user} -p ${password} ${host} 9042 -f session.cql
echo "task.cql"
cqlsh -u ${user} -p ${password} ${host} 9042 -f task.cql
echo "trace.cql"
cqlsh -u ${user} -p ${password} ${host} 9042 -f trace.cql
echo "data.cql"
cqlsh -u ${user} -p ${password} ${host} 9042 -f data.cql

