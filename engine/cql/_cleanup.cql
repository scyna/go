DROP KEYSPACE IF EXISTS scyna;
CREATE KEYSPACE scyna WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };