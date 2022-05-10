DROP TABLE IF EXISTS scyna.call;
CREATE TABLE scyna.call
(
    id         bigint,
    day        int,
    time       timestamp,
    duration   int, /*in nano sec*/
    source     text, /*service url, event subject, signal subject, push subject*/
    request    blob, /*JSON/PROTOBUF*/
    response   blob, /*JSON/PROTOBUF*/
    status     int,
    session_id bigint,
    caller_id  text, /*client or app, optional*/
    PRIMARY KEY ((source, day), id) /*source_has_call*/
);

DROP TABLE IF EXISTS scyna.call_log;
CREATE TABLE scyna.call_log
(
    call_id     bigint,
    seq         int,
    time        timestamp,
    message     text,
    level       smallint,
    PRIMARY KEY (call_id, seq) /*call_has_log*/
);

