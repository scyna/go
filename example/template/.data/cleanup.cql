DROP KEYSPACE IF EXISTS template;

CREATE KEYSPACE template
WITH replication = {'class': 'NetworkTopologyStrategy', 'vinfast_cdp_dc' : 3} 
AND durable_writes = true;

-- CREATE KEYSPACE template WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
-- CREATE KEYSPACE template WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 3 };
