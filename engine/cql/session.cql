USE scyna;

DROP TABLE IF EXISTS session;
CREATE TABLE session
(
    id          bigint,
    context_code text,
    start       timestamp,
    last_update timestamp,
    end         timestamp,
    exit_code   int, /*normal, crash, ....*/
    PRIMARY KEY (module_code, id) 
);
CREATE INDEX ON scyna.session (id);

DROP TABLE IF EXISTS session_log;
CREATE TABLE session_log
(
    session_id bigint,
    day        int,
    time       timestamp,
    seq        bigint,
    message    text,
    level      smallint,
    PRIMARY KEY ((session_id, day), seq) /*sesion_has_log*/
);

