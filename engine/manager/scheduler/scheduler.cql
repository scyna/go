DROP TABLE IF EXISTS scyna.todo;
CREATE TABLE scyna.todo
(
    bucket              bigint,
    task_id             bigint,
    repeater_id         bigint,
    time                timestamp,
    topic               text,
    data                blob,
    PRIMARY KEY (bucket, task_id)
)

DROP TABLE IF EXISTS scyna.doing;
CREATE TABLE scyna.doing
(
    bucket        bigint,
    task_id       bigint, 
    PRIMARY KEY (bucket, task_id),
);

DROP TABLE IF EXISTS scyna.repeater;
CREATE TABLE scyna.repeater
(
    id              bigint PRIMARY KEY,
    start           timestamp, /*null: inactive*/
    interval        bigint,
    count           bigint, /*0 = infinitive*/
    topic           text,
    data            blob,
);

DROP TABLE IF EXISTS scyna.task;
CREATE TABLE scyna.task
(
    module              text,
    id                  bigint,
    repeater_id         bigint,
    time                timestamp,
    topic               text,
    data                blob,
    PRIMARY KEY (bucket, task_id)
);
