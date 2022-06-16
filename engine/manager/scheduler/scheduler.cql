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
    interval     bigint,
    count        bigint,
    max          bigint,
    running      boolean,
    PRIMARY KEY (module, id)
);
CREATE INDEX ON scyna.task (id);