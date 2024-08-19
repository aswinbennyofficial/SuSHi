-- +goose Up
INSERT INTO users (username, name, email, salt)
VALUES (
    '6b338158-32ca-4b53-a273-f54e3244697e',
    'Dummy User',
    'dummy@example.com',
    'randomsalt123456789'
);



-- +goose Down
DELETE FROM users WHERE email = 'dummy@example.com';
