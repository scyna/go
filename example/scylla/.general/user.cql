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

insert into scyna.module(org_code, code, description, type, secret)
values ('scyna', 'scyna.example', 'Scyna Example', 2, '123456');
insert into scyna.application(org_code, code, auth, name) VALUES ('scyna', 'scyna.example', '', '');