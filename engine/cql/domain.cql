/*domain*/
DROP TABLE IF EXISTS scyna.domain;
CREATE TABLE scyna.domain
(
    code     text PRIMARY KEY, /*iot*/
    name     text,             /*IoT Platform*/ 
    password text, 
);

/*context*/
DROP TABLE IF EXISTS scyna.context;
CREATE TABLE scyna.context
(
    domain      text,  /*iot*/
    code        text,  /*iot.control*/
    name        text,  /*IoT Control*/
    secret      text,
    PRIMARY KEY(domain, code)
);
CREATE INDEX ON scyna.context (code);

/*client*/
DROP TABLE IF EXISTS scyna.client;
CREATE TABLE scyna.client
(
    domain      text,  /*iot*/
    id          text,  /*datalake*/
    secret      text, 
    PRIMARY KEY(domain, id)
);
CREATE INDEX ON scyna.client (id);

/*application*/
DROP TABLE IF EXISTS scyna.application;
CREATE TABLE scyna.application
(
    domain      text, /* iot */
    code        text, /* fota */
    name        text, /* FOTA Manager */
    PRIMARY KEY(domain, code)
);
CREATE INDEX ON scyna.application (code);

/*service: Command or Query*/
DROP TABLE IF EXISTS scyna.service;
CREATE TABLE scyna.service
(
    context     text, /* iot.control */
    url         text, /* iot/control/poweron */  
    name        text, /* PowerOn */
    description text,
    PRIMARY KEY (context, url)
);
CREATE INDEX ON scyna.service (url);

/*event*/
DROP TABLE IF EXISTS scyna.event;
CREATE TABLE scyna.event
(
    context     text, /* iot.control */
    channel     text, /* iot.control.poweredon */
    name        text, /* PoweredOn */
    PRIMARY KEY (context, channel)
);
CREATE INDEX ON scyna.event (channel);

/*setting*/
DROP TABLE IF EXISTS scyna.setting;
CREATE TABLE scyna.setting
(
    context     text,
    key         text,
    value       text,
    PRIMARY KEY (context, key)
);

/*relations*/
DROP TABLE IF EXISTS scyna.client_use_service;
CREATE TABLE scyna.client_use_service
(
    client      text,   /*client id*/  
    service     text,   /*service path*/
    PRIMARY KEY (client, service)
);

DROP TABLE IF EXISTS scyna.application_use_service;
CREATE TABLE scyna.application_use_service
(
    application     text,   /*application code*/  
    service         text,   /*service path*/
    PRIMARY KEY (application, service)
);