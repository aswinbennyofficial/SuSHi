-- +goose Up
INSERT INTO users (username, auth_provider_id, name, email, salt)
VALUES (
    '6b338158-32ca-4b53-a273-f54e3244697e',
    'auth_provider|dummy123456',
    'Dummy User',
    'dummy@example.com',
    'randomsalt123456789'
);



-- +goose Down
DELETE FROM users WHERE email = 'dummy@example.com';
