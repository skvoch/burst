CREATE TABLE preview_tokens (
uid text not null primary key,
bookID REFERENCES books(id)
);