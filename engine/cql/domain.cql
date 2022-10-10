/* domain */
DROP TABLE IF EXISTS scyna.domain;
CREATE TABLE scyna.domain
(
    code     text PRIMARY KEY,
    name     text,
    password text,
);

/* context */
DROP TABLE IF EXISTS scyna.context;
CREATE TABLE scyna.context
(
    domain_code     text,
    code            text,
    name            text,
    secret          text,
    PRIMARY KEY (domain_code, code)
);
CREATE INDEX ON scyna.context (code);

/* application */
DROP TABLE IF EXISTS scyna.application;
CREATE TABLE scyna.application
(
    domain_code     text,
    code            text,
    name            text,
    PRIMARY KEY (domain_code, code)
);
CREATE INDEX ON scyna.application (code);

/* client */
DROP TABLE IF EXISTS scyna.client;
CREATE TABLE scyna.client
(
    domain_code     text,
    id              text,
    name            text,
    secret          text,
    PRIMARY KEY (domain_code, id)
);
CREATE INDEX ON scyna.client (id);

/*relations*/
DROP TABLE IF EXISTS scyna.client_use_service;
CREATE TABLE scyna.client_use_service
(
    client_id   text,
    url         text,
    PRIMARY KEY (client_id, url)
);

DROP TABLE IF EXISTS scyna.application_use_service;
CREATE TABLE scyna.application_use_service
(
    application_code    text,
    url                 text,
    PRIMARY KEY (application_code, url)
);
