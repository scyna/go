/*application*/
DROP TABLE IF EXISTS scyna.application;
CREATE TABLE scyna.application
(
    code     text,
    name     text,
    auth     text, /*authentication URL*/
    org_code text,
    PRIMARY KEY (org_code, code)
);
CREATE INDEX ON scyna.application (code);

DROP TABLE IF EXISTS scyna.app_use_service;
CREATE TABLE scyna.app_use_service
(
    app_code    text,
    service_url text,
    PRIMARY KEY (app_code, service_url)
);

DROP TABLE IF EXISTS scyna.authentication;
CREATE TABLE scyna.authentication
(
    id      text PRIMARY KEY,
    time    timestamp,
    expired timestamp,
    apps    list<text>,
    uid     text,
    data    blob,
);

DROP TABLE IF EXISTS scyna.app_has_auth;
CREATE TABLE scyna.app_has_auth
(
    app_code text,
    auth_id  text,
    user_id  text,
    PRIMARY KEY ( (app_code, user_id), auth_id )
);