-- Auto-generated at Fri Apr 13 16:21:27 CEST 2018
-- Please do not change the name attributes

-- name: up

CREATE TABLE IF NOT EXISTS migrations (
 id          TEXT      NOT NULL PRIMARY KEY,
 description TEXT      NOT NULL,
 created_at  TIMESTAMP NOT NULL
);
-- name: down

DROP TABLE IF EXISTS migrations;