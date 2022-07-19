/*schema for EventStore*/
CREATE TABLE module_name.event_store
(
    bucket      bigint,
    id          bigint,
    time        timestamp,
    subject     text,
    data        blob,
    PRIMARY(bucket, id)
) WITH CLUSTERING ORDER BY (id DESC);
