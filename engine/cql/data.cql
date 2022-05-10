USE scyna;

/*scyna organization*/
insert into scyna.organization(code, name, password) values('scyna', 'Scyna Platform', '123456');

/*core modules*/
insert into scyna.module(org_code, code, description, type, secret) values ('scyna', 'scyna.engine', 'Scyna Engine', 2, '123456');
insert into scyna.module(org_code, code, description, type, secret) values ('scyna', 'scyna.admin', 'Scyna Admin', 2, '123456'); 
insert into scyna.module(org_code, code, description, type, secret) values ('scyna', 'scyna.test', 'Scyna Testing', 2, '123456');
insert into scyna.module(org_code, code, description, type, secret) values ('scyna', 'scyna.data', 'Scyna Data', 2, '123456');

