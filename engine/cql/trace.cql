DROP TABLE IF EXISTS scyna.trace;
CREATE TABLE scyna.trace
(
    path        text,
    day         int,
    id          bigint,
    time        timestamp,
    session_id  bigint,
    parent_id   bigint,
    PRIMARY KEY ((path, day), id)
);
CREATE INDEX ON scyna.trace(id);

DROP TABLE IF EXISTS scyna.span;
CREATE TABLE scyna.span
(
    parent_id     bigint,
    child_id     bigint,
    PRIMARY KEY (parent_id, child_id)
);

DROP TABLE IF EXISTS scyna.tag;
CREATE TABLE scyna.tag
(
    trace_id     bigint,
    key          text,
    value        blob,
    PRIMARY KEY (trace_id, tag)
);

DROP TABLE IF EXISTS scyna.log;
CREATE TABLE scyna.log
(
    trace_id     bigint,
    seq          bigint,
    time         timestamp,
    message      text,
    PRIMARY KEY (trace_id, seq)
);
