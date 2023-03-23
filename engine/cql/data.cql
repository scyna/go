USE scyna;

/*scyna organization*/
insert into scyna.organization(code, name, password) values ('scyna', 'Scyna Platform', '123456');
insert into scyna.organization(code, name, password, secret) values ('org_test', 'Organization test', '123456', '123456');

/* scyna client */
insert into scyna.client(org_code, id, secret, state, type) values ('org_test', 'DEV', 'RAVT47Kz6Szw', 1, 'test');  

/*core modules*/
insert into scyna.module(org_code, code, description, type, secret) values ('scyna', 'scyna.engine', 'Scyna Engine', 2, '123456');
insert into scyna.module(org_code, code, description, type, secret) values ('scyna', 'scyna.admin', 'Scyna Admin', 2, '123456'); 
insert into scyna.module(org_code, code, description, type, secret) values ('scyna', 'scyna.test', 'Scyna Testing', 2, '123456');
insert into scyna.module(org_code, code, description, type, secret) values ('scyna', 'scyna.data', 'Scyna Data', 2, '123456');
insert into scyna.module(org_code, code, description, type, secret) values ('scyna', 'scyna.manager', 'Scyna Data', 2, '123456');

/* service */
insert into scyna.service(module_code, url, active, description) VALUES ('scyna.engine', '/scyna/task/add', true, 'Add Task');
insert into scyna.service(module_code, url, active, description) VALUES ('scyna.engine', '/scyna/task/add-recurring', true, 'Add Task');

insert into scyna.client_use_service(client_id, service_url) VALUES ('DEV', '/scyna/task/add');
insert into scyna.client_use_service(client_id, service_url) VALUES ('DEV', '/scyna/task/add-recurring');