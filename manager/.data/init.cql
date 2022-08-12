INSERT INTO scyna.module (org_code, code, description, secret, type) VALUES ('scyna', 'scyna.manager', 'Scyna Manager', '123456', 2);

/* Init setting */
INSERT INTO scyna.setting(module_code, key, value) VALUES ('scyna.manager', 'scyna.module.config', '{
    "NatsUrl": "localhost",
    "NatsUsername": "",
    "NatsPassword": "",
    "DBHost": "localhost",
    "DBUsername": "",
    "DBPassword": "",
    "DBLocation": ""
}') IF NOT EXISTS ;

INSERT INTO scyna.client_use_service (client_id, service_url) VALUES ('DEV', '/scyna/manager/organization/create');
INSERT INTO scyna.client_use_service (client_id, service_url) VALUES ('DEV', '/scyna/manager/application/create');
INSERT INTO scyna.client_use_service (client_id, service_url) VALUES ('DEV', '/scyna/manager/module/create');
INSERT INTO scyna.client_use_service (client_id, service_url) VALUES ('DEV', '/scyna/manager/service/create');
INSERT INTO scyna.client_use_service (client_id, service_url) VALUES ('DEV', '/scyna/manager/client/create');
INSERT INTO scyna.client_use_service (client_id, service_url) VALUES ('DEV', '/scyna/manager/client/add-service');
INSERT INTO scyna.client_use_service (client_id, service_url) VALUES ('DEV', '/scyna/manager/client/remove-service');
INSERT INTO scyna.client_use_service (client_id, service_url) VALUES ('DEV', '/scyna/manager/event/create');
INSERT INTO scyna.client_use_service (client_id, service_url) VALUES ('DEV', '/scyna/manager/sync/create');