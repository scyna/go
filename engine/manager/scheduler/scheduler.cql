DROP TABLE IF EXISTS scyna.todo;
CREATE TABLE scyna.todo
(
    bucket        bigint,
    task_id       bigint,
    PRIMARY KEY (bucket, task_id)
);

DROP TABLE IF EXISTS scyna.doing;
CREATE TABLE scyna.doing
(
    task_id       bigint PRIMARY KEY 
);

DROP TABLE IF EXISTS scyna.task;
CREATE TABLE scyna.task
(
    id           bigint PRIMARY KEY ,
    topic        text,
    data         blob,
    start        timestamp,
    next         timestamp,
    interval     bigint,
    loop_count   bigint,
    loop_max     bigint,
    active       boolean,
);

DROP TABLE IF EXISTS scyna.module_has_task;
CREATE TABLE scyna.module_has_task
(
    module       text,
    task_id      bigint,
    PRIMARY KEY (module, task_id)
);