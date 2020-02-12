CREATE TABLE books (
id bigserial not null primary key,
name text,
description text,
review text,
rating integer,
type integer REFERENCES types(id)
);
