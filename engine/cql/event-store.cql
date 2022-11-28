/*event store template*/
DROP TABLE IF EXISTS module_name.event_store;
CREATE TABLE module_name.event_store
(
    event_id        bigint,
    aggregate_id    bigint,
    channel         text,
    data            blob,
    PRIMARY KEY(event_id)
);


DROP TABLE IF EXISTS module_name.activity;
CREATE TABLE module_name.activity
(
    aggregate_id    bigint,
    event_id        bigint, 
    PRIMARY KEY(aggregate_id, event_id)
);

