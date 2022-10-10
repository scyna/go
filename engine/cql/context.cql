/*command*/
DROP TABLE IF EXISTS scyna.command;
CREATE TABLE scyna.command
(
    context_code    text,
    name            text,
    description     text,
    url             text,
    PRIMARY KEY (context_code, name)
);
CREATE INDEX ON scyna.command (url);

/*event*/
DROP TABLE IF EXISTS scyna.event;
CREATE TABLE scyna.event
(
    context_code    text,
    name            text,
    channel         text,
    PRIMARY KEY (context_code, name)
);
CREATE INDEX ON scyna.event (channel);

/*query*/
DROP TABLE IF EXISTS scyna.query;
CREATE TABLE scyna.query
(
    context_code    text,
    name            text,
    description     text,
    url             text,
    PRIMARY KEY (context_code, name)
);
CREATE INDEX ON scyna.query (url);
