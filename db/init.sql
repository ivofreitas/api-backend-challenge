CREATE DATABASE IF NOT EXISTS sword;

CREATE TABLE sword.roles (
    id          VARCHAR(40) NOT NULL,
    position    VARCHAR(80) UNIQUE NULL,
    created_at  TIMESTAMP   NULL,
    updated_at  TIMESTAMP   NULL,
    PRIMARY KEY (id)
);

CREATE TABLE sword.users (
    id          VARCHAR(40)  NOT NULL,
    username    VARCHAR(80)  UNIQUE NOT NULL,
    password    VARCHAR(255) NOT NULL,
    role_id     VARCHAR(40)  NULL,
    created_at  TIMESTAMP    NULL,
    updated_at  TIMESTAMP    NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (role_id)   REFERENCES sword.roles(id)
);

CREATE TABLE sword.tasks (
    id VARCHAR(40)              NOT NULL,
    summary         TEXT        NOT NULL,
    performed_by    VARCHAR(40) NOT NULL,
    performed_at    TIMESTAMP   NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (performed_by)   REFERENCES sword.users(id)
);

INSERT INTO sword.roles
VALUES ('9c06de8d-865a-4ed4-88d0-a445f2bd4f11', 'admin', NOW(), NOW()),
       ('59a2cc28-a219-47c2-945a-e520827bbfcd', 'manager', NOW(), NOW()),
       ('ca99ac9f-f26d-4aeb-8966-234c277ff7b4', 'technician', NOW(), NOW());

