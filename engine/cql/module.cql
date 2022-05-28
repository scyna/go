/*module*/
DROP TABLE IF EXISTS scyna.module;
CREATE TABLE scyna.module
(
    code        text,
    description text,
    type        smallint, /*PRIVATE=1, PUBLIC=2*/
    org_code    text,
    secret      text,
    PRIMARY KEY (org_code, code)
);
CREATE INDEX ON scyna.module (code);

/*setting*/
DROP TABLE IF EXISTS scyna.setting;
CREATE TABLE scyna.setting
(
    module_code text,
    key         text,
    value       text,
    PRIMARY KEY (module_code, key) /*module_has_setting*/
);

DROP TABLE IF EXISTS scyna.service;
CREATE TABLE scyna.service
(
    module_code text,
    url         text,
    active      boolean,
    description text,
    PRIMARY KEY (module_code, url)
);
CREATE INDEX ON scyna.service (url);

DROP TABLE IF EXISTS scyna.module_has_service;
CREATE TABLE scyna.module_has_service
(
    module_code text,
    service_url      text,
    PRIMARY KEY (module_code, service_url)
);