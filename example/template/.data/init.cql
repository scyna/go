/* Init module */
INSERT INTO scyna.organization(code, name, password) VALUES('org', 'Organization', '123456789aA@#') IF NOT EXISTS;
INSERT INTO scyna.module(org_code ,code, type, secret) VALUES ('org' ,'template', 1, '123456789aA@#') IF NOT EXISTS;

/* Init setting */
INSERT INTO scyna.setting(module_code, key, value) VALUES ('template', 'scyna.module.config', '{
    "NatsUrl": "localhost",
    "NatsUsername": "",
    "NatsPassword": "",
    "DBHost": "localhost",
    "DBUsername": "",
    "DBPassword": "",
    "DBLocation": ""
}') IF NOT EXISTS ;

DROP TABLE IF EXISTS template.activity;
CREATE TABLE template.activity
(
    entity_id      bigint,
    event_id       bigint,
    PRIMARY KEY (entity_id, event_id)
) WITH CLUSTERING ORDER BY (event_id DESC);

DROP TABLE IF EXISTS template.event_store;
CREATE TABLE template.event_store
(
    bucket      bigint,
    id          bigint,
    entity_id   list<bigint>,
    time        timestamp,
    subject     text,
    data        blob,
    PRIMARY KEY (bucket, id)
) WITH CLUSTERING ORDER BY (id DESC);
INSERT INTO template.event_store(bucket, id) VALUES (0,1);
