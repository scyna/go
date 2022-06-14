DROP TABLE IF EXISTS scyna.task;
CREATE TABLE scyna.task
(
    period      bigint,
    time        bigint,
    type        text,
    context     blob,
    PRIMARY KEY (period, time)
) WITH CLUSTERING ORDER BY (time ASC);

DROP TABLE IF EXISTS scyna.task_assigned;
CREATE TABLE scyna.task_assigned
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
    context         blob,
    PRIMARY KEY (period, time),
) WITH CLUSTERING ORDER BY (time ASC);