DROP TABLE IF EXISTS scyna.task;
CREATE TABLE scyna.task
(
    bucket              bigint,
    id                  bigint,
    recurring_task_id   bigint,
    send_to             text,
    type                text,
    time                timestamp,
    data                blob,
    PRIMARY KEY (bucket, id)
) WITH CLUSTERING ORDER BY (time ASC);

DROP TABLE IF EXISTS scyna.assigned_task;
CREATE TABLE scyna.assigned_task
(
    bucket        bigint,
    task_id       bigint, 
    PRIMARY KEY (bucket, task_id),
);

DROP TABLE IF EXISTS scyna.recurring_task;
CREATE TABLE scyna.recurring_task
(
    id              bigint PRIMARY KEY,
    time            timestamp,
    interval        bigint,
    count           bigint,
    send_to         text,
    type            text,
    data            blob,
);