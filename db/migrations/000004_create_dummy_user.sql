-- +goose Up
INSERT INTO users (username, name, email, salt)
VALUES (
    'dummy_user123',
    'Dummy User',
    'dummy@example.com',
    'randomsalt123456789'
);


-- +goose Down
DELETE FROM users WHERE username = 'dummy_user123';
