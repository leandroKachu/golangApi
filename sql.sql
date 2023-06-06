CREATE DATABASE leandro;
\c leandro;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id serial primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(100) not null unique,
    created_at timestamp default current_timestamp
);