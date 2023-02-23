DROP TABLE IF EXISTS scyna.api_report_by_day;
CREATE TABLE scyna.api_report_by_day
(
    day              date PRIMARY KEY,
    total_error      int,
    total_success    int,
    total_permission int,
    avg_latency      bigint, //nano second
    min_latency      bigint, //nano second
    max_latency      bigint, //nano second
    data             text
);

DROP TABLE IF EXISTS scyna.api_report_by_error;
CREATE TABLE scyna.api_report_by_error
(
    day       date PRIMARY KEY,
    trace_id  bigint,
    path      text,
    client_id text
);

DROP TABLE IF EXISTS scyna.api_report_by_permission;
CREATE TABLE scyna.api_report_by_permission
(
    day       date PRIMARY KEY,
    trace_id  bigint,
    path      text,
    client_id text
);