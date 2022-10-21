/*event store template*/
DROP TABLE IF EXISTS module_name.event_store;
CREATE TABLE module_name.event_store
(
    aggregate_id    bigint ,
    event           text, /*iot.account.account_created*/
    data            blob,
    version         bigint, 
    PRIMARY KEY(aggregate_id, version)
);

