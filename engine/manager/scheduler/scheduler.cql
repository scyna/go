DROP TABLE IF EXISTS scyna.todo;
CREATE TABLE scyna.todo
(
    bucket        bigint,
    task_id       bigint,
    PRIMARY KEY (bucket, task_id)
)

DROP TABLE IF EXISTS scyna.doing;
CREATE TABLE scyna.doing
(
    task_id       bigint PRIMARY KEY 
);

DROP TABLE IF EXISTS scyna.task;
CREATE TABLE scyna.task
(
    module       text,
    id           bigint,
    topic        text,
    data         blob,
    time         timestamp,
    interval     bigint, /*in minute*/
    count        bigint,
    max          bigint,
    PRIMARY KEY (module, id)
);
CREATE INDEX ON scyna.task (id);