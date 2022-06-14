DROP TABLE IF EXISTS scyna.task;
CREATE TABLE scyna.task
(
    period      bigint,
    time        bigint,
    repeat      boolean,
    send_to     text,
    type        text,
    data        blob,
    PRIMARY KEY (period, time)
) WITH CLUSTERING ORDER BY (time ASC);

DROP TABLE IF EXISTS scyna.assigned_task;
CREATE TABLE scyna.assigned_task
(
    time        bigint PRIMARY KEY,
    createdat   timestamp, 
);

DROP TABLE IF EXISTS scyna.recurring_task;
CREATE TABLE scyna.recurring_task
(
    period          bigint,
    time            bigint,
    interval        bigint,
    send_to         text,
    type            text,
    data            blob,
    PRIMARY KEY (period, time),
) WITH CLUSTERING ORDER BY (time ASC);