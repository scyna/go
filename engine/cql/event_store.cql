/*schema for EventStore*/
CREATE TABLE module_name.event_store
(
    bucket      bigint,
    id          bigint,
    time        timestamp,
    subject     text,
    data        blob,
    PRIMARY(bucket, id)
);

INSERT INTO module_name.event_store(id, data) values(0, bigintAsBlob(0));


CREATE TABLE scyna.event_store
(
    id          bigint PRIMARY KEY,
    time        timestamp,
    subject     text,
    data        blob
);