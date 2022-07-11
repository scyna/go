INSERT INTO scyna.module (org_code, code, description, secret, type) VALUES ('scyna', 'scyna.manager', 'Scyna Manager', '123456', 2);
INSERT INTO scyna.client_use_service (client_id, service_url) VALUES ('DEV', '/scyna/manager/organization/create');
INSERT INTO scyna.client_use_service (client_id, service_url) VALUES ('DEV', '/scyna/manager/application/create');