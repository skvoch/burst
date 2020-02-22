CREATE TABLE books (
id bigserial not null primary key,
name text,
description text,
review text,
rating integer,
file_path text unique,
preview_path text unique,
type integer REFERENCES types(id)
);
