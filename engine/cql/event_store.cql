/*schema for EventStore*/
CREATE TABLE module_name.event_store
(
    id          bigint PRIMARY KEY,
    time        timestamp,
    subject     text,
    data        blob
);

INSERT INTO module_name.event_store(id, subject) values(0, '0');
