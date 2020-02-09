CREATE TABLE books (
id integer PRIMARY KEY,
name text,
description text,
review text,
rating integer,
type integer REFERENCES types(id)
);
