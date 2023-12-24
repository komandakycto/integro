-- Create the users table
CREATE TABLE users
(
    id         BIGSERIAL PRIMARY KEY,
    username   VARCHAR(50)  NOT NULL,
    email      VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert 3 test users
INSERT INTO users (username, email)
VALUES ('user1', 'user1@example.com'),
       ('user2', 'user2@example.com'),
       ('user3', 'user3@example.com');