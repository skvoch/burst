CREATE TABLE books (
id bigserial not null primary key,
name text,
description text,
review text,
rating integer,
file_path text,
preview_path text,
type integer REFERENCES types(id)
);
