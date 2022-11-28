DROP KEYSPACE IF EXISTS ex;
CREATE KEYSPACE ex WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

DROP TABLE IF EXISTS ex.user;
CREATE TABLE ex.user
(
    id           bigint PRIMARY KEY,
    email        text ,
    name         text,
    password     text,
);
CREATE INDEX ON ex.user (email);

DROP TABLE IF EXISTS ex.has_friend;
CREATE TABLE ex.has_friend
(
    user_id        bigint,
    friend_id      bigint,
    PRIMARY KEY(user_id, friend_id)
);

DROP TABLE IF EXISTS ex.event_store;
CREATE TABLE ex.event_store
(
    event_id        bigint PRIMARY KEY,
    entity_id       bigint,
    channel         text,
    data            blob,
    time            time,
);
