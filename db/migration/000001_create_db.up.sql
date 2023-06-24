CREATE DATABASE IF NOT EXISTS sword;

CREATE TABLE sword.tasks (
  id              VARCHAR(40) NOT NULL,
  summary         TEXT        NULL,
  performed_at    TIMESTAMP   NULL,
  PRIMARY KEY (id)
);

CREATE TABLE sword.roles (
  id          VARCHAR(40) NOT NULL,
  position    VARCHAR(80) NULL,
  created_at  TIMESTAMP   NULL,
  updated_at  TIMESTAMP   NULL,
  PRIMARY KEY (id)
);

CREATE TABLE sword.users (
   id          VARCHAR(40) NOT NULL,
   name        VARCHAR(80) NULL,
   role_id     VARCHAR(40) NULL,
   created_at  TIMESTAMP   NULL,
   updated_at  TIMESTAMP   NULL,
   PRIMARY KEY (id),
   FOREIGN KEY (role_id) REFERENCES sword.roles(id)
);