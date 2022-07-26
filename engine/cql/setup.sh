#!/bin/bash
echo "cleanup"
cqlsh -f _cleanup.cql

echo "generator"
cqlsh -f generator.cql

echo "module"
cqlsh -f module.cql

echo "session"
cqlsh -f session.cql

echo "trace"
cqlsh -f trace.cql

echo "client"
cqlsh -f client.cql

echo "organization"
cqlsh -f organization.cql

echo "application"
cqlsh -f application.cql

echo "schedule"
cqlsh -f schedule.cql

echo "data"
cqlsh -f data.cql
#go run data.go