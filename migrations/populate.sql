INSERT INTO users (id, name) VALUES (
    'testuid',
    'testname'
);

INSERT INTO account (id, user_id, total) VALUES (
    'testaid0',
    'testuid',
    0
);

INSERT INTO account (id, user_id, total) VALUES (
    'testaid1',
    'testuid',
    0
);

INSERT INTO account (id, user_id, total) VALUES (
    'testaid2',
    'testuid',
    100
);
