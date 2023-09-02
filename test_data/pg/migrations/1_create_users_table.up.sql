CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(50)  NOT NULL,
    email      VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Insert test users
INSERT INTO users (username, email)
VALUES ('user1', 'user1@example.com'),
       ('user2', 'user2@example.com'),
       ('user3', 'user3@example.com');