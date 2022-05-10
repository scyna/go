/*organization*/
DROP TABLE IF EXISTS scyna.organization;
CREATE TABLE scyna.organization
(
    code     text PRIMARY KEY,
    name     text,
    password text,
    secret   text,
);

DROP TABLE IF EXISTS scyna.org_has_module;
CREATE TABLE scyna.org_has_module
(
    org_code text,
    module_code  text,
    PRIMARY KEY (org_code, module_code)
);

DROP TABLE IF EXISTS scyna.org_has_app;
CREATE TABLE scyna.org_has_app
(
    org_code text,
    app_code  text,
    PRIMARY KEY (org_code, app_code)
);

DROP TABLE IF EXISTS scyna.org_has_client;
CREATE TABLE scyna.org_has_client
(
    org_code text,
    client_id  text,
    PRIMARY KEY (org_code, client_id)
);