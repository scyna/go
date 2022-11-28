/*event store template*/
DROP TABLE IF EXISTS module_name.event_store;
CREATE TABLE module_name.event_store
(
    event_id        bigint,
    entity_id       bigint,
    channel         text,
    data            blob,
    PRIMARY KEY(event_id)
);
