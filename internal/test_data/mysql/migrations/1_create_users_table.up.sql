CREATE TABLE public.users
(
    id         integer AUTO_INCREMENT PRIMARY KEY,
    username   VARCHAR(50)  NOT NULL,
    email      VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert test users
INSERT INTO public.users (username, email)
VALUES ('user1', 'user1@example.com'),
       ('user2', 'user2@example.com'),
       ('user3', 'user3@example.com');