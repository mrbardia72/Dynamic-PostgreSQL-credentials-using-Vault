-- +migrate Up
ALTER DEFAULT PRIVILEGES FOR ROLE bardia GRANT ALL ON TABLES TO PUBLIC;
ALTER DEFAULT PRIVILEGES FOR ROLE bardia GRANT ALL ON SEQUENCES TO PUBLIC;

CREATE TABLE authors
(
    id      SERIAL PRIMARY KEY,
    name       VARCHAR(256) NOT NULL
);

-- +migrate Down

DROP TABLE authors;
