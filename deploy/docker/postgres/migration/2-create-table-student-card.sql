CREATE SCHEMA IF NOT EXISTS student;
CREATE TABLE IF NOT EXISTS student.card (
    id bigserial PRIMARY KEY,
    data jsonb NOT NULL,
    createdDateTime TIMESTAMP NOT NULL
);
