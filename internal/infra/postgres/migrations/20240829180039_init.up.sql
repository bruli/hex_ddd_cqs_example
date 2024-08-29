create table if not exists users(
    id uuid primary key ,
    username varchar not null,
    phone varchar
);