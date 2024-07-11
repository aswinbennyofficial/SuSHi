-- +goose Up
INSERT INTO users (username, name, email, salt)
VALUES (
    '123456',
    'Dummy User',
    'dummy@example.com',
    'randomsalt123456789'
);


-- +goose Down
DELETE FROM users WHERE username = '123456';
