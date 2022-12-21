DROP KEYSPACE IF EXISTS scyna;
CREATE KEYSPACE scyna WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

-- CREATE KEYSPACE scyna 
-- WITH replication = {'class': 'NetworkTopologyStrategy', 'vinfast_cdp_dc' : 3} 
-- AND durable_writes = true;