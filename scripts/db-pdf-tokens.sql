CREATE TABLE pdf_tokens (
uid text not null primary key,
bookID integer REFERENCES books(id)
);