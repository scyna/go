DROP KEYSPACE IF EXISTS ex_account;
CREATE KEYSPACE ex_account WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

DROP TABLE IF EXISTS ex_account.account;
CREATE TABLE ex_account.account
(
    id           bigint PRIMARY KEY,
    email        text,
    name         text,
    password     text,
    avatar       text,
    dob          text,
    gender       smallint,
    tel          text,
);
CREATE INDEX ON ex_account.account (email);

DROP TABLE IF EXISTS ex_account.event_store;
CREATE TABLE ex_account.event_store
(
    event_id        bigint PRIMARY KEY,
    entity_id       bigint,
    channel         text,
    data            blob,
    time            timestamp,
);
