DROP TABLE IF EXISTS scyna.task;
CREATE TABLE scyna.task
(
    period      bigint,
    time        bigint,
    type        text,
    context     blob,
    PRIMARY KEY (period, time)
) WITH CLUSTERING ORDER BY (done ASC, time ASC);

DROP TABLE IF EXISTS scyna.task_assigned;
CREATE TABLE scyna.task_assigned
(
    time        bigint PRIMARY KEY,
    createdat   timestamp,    
);

DROP TABLE IF EXISTS scyna.repeat_task;
CREATE TABLE scyna.repeat_task
(
    time            bigint,
    interval        bigint,
    context         blob,
);