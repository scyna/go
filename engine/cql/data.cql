USE scyna;

/*scyna domain*/
insert into scyna.domain(code, name, password) values ('scyna', 'Scyna Platform', '123456');

/*core context*/
insert into scyna.context(domain, code, name, secret) values ('scyna', 'scyna.engine', 'Scyna Engine', '123456');
insert into scyna.context(domain, code, name, secret) values ('scyna', 'scyna.manager', 'Scyna Manager', '123456');
insert into scyna.context(domain, code, name, secret) values ('scyna', 'scyna.test', 'Scyna Test', '123456');
