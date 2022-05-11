/*client*/
DROP TABLE IF EXISTS scyna.client;
CREATE TABLE scyna.client
(
    id          text,
    secret      text,
    state       smallint, /*active, inactive, ...*/
    type        text,
    org_code    text,
    PRIMARY KEY (org_code, id)
);
CREATE INDEX ON scyna.client (id);

DROP TABLE IF EXISTS scyna.client_use_service;
CREATE TABLE scyna.client_use_service
(
    client_id text,
    service_url   text,
    PRIMARY KEY (client_id, service_url)
);
