DROP TABLE IF EXISTS scyna.trace;
CREATE TABLE scyna.trace
(
    path        text,
    day         int,
    id          bigint,
    time        timestamp,
    type        smallint,
    session_id  bigint,
    duration    bigint,
    parent_id   bigint,
    source      text,
    status      int,
    PRIMARY KEY ((path, day), id)
);
CREATE INDEX ON scyna.trace(id);

DROP TABLE IF EXISTS scyna.app_has_trace;
CREATE TABLE scyna.app_has_trace
(
    app_code     text,
    trace_id     bigint,
    day         int,
    PRIMARY KEY ((app_code,day), trace_id)
);

DROP TABLE IF EXISTS scyna.client_has_trace;
CREATE TABLE scyna.client_has_trace
(
    client_id     text,
    trace_id     bigint,
    day         int,
    PRIMARY KEY ((client_id,day), trace_id)
);

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
    value        text,
    PRIMARY KEY (trace_id, key)
);

DROP TABLE IF EXISTS scyna.log;
CREATE TABLE scyna.log
(
    trace_id     bigint,
    seq          bigint,
    time         timestamp,
    message      text,
    level       smallint,
    PRIMARY KEY (trace_id, seq)
);
