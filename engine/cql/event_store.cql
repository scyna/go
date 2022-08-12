/*schema for EventStore*/
CREATE TABLE module_name.event_store
(
    bucket      bigint,
    id          bigint,
    entity_id   list<bigint>,
    time        timestamp,
    subject     text,
    data        blob,
    PRIMARY KEY (bucket, id)
) WITH CLUSTERING ORDER BY (id DESC);

CREATE TABLE module_name.activity
(
    entity_id      bigint,
    event_id       bigint,
    PRIMARY KEY (entity_id, event_id)
) WITH CLUSTERING ORDER BY (event_id DESC);

INSERT INTO module_name.event_store(bucket, id) VALUES (0,1);

CREATE TABLE module_name.event_store
(
    aggregate     bigint,
    version       bigint,
    time          timestamp,
    event         text,
    data          blob,
    PRIMARY KEY (aggregate, version)
) WITH CLUSTERING ORDER BY (version DESC);