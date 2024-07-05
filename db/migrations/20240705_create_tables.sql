-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE organizations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE machines (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    hostname VARCHAR(255) NOT NULL,
    port INTEGER NOT NULL DEFAULT 22,
    username VARCHAR(50) NOT NULL,
    auth_type VARCHAR(20) NOT NULL CHECK (auth_type IN ('password', 'private_key')),
    password_encrypted TEXT,
    private_key_encrypted TEXT,
    passphrase_encrypted TEXT, -- passphrase that encrypts private_key
    owner_id INTEGER NOT NULL,
    owner_type VARCHAR(20) NOT NULL CHECK (owner_type IN ('user', 'organization')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    -- Note: The complex foreign key constraint commented out because it is not supported directly in SQL.
);

-- +goose Down
DROP TABLE IF EXISTS machines;
DROP TABLE IF EXISTS organizations;
DROP TABLE IF EXISTS users;
