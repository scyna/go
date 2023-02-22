DROP KEYSPACE IF EXISTS template;

CREATE KEYSPACE template
WITH replication = {'class': 'NetworkTopologyStrategy', 'replication_factor' : 3}
AND durable_writes = true;
