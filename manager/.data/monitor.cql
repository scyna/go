DROP TABLE IF EXISTS scyna.api_report_by_day;
CREATE TABLE scyna.api_report_by_day
(
    day           date PRIMARY KEY,
    total_error   int,
    total_success int,
    avg_latency   bigint, //nano second
    min_latency   bigint, //nano second
    max_latency   bigint, //nano second
    data          text
);
