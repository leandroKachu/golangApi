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



CREATE TABLE followers (
    user_id int not null,
    follower_id int not null,
    PRIMARY KEY (user_id, follower_id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (follower_id) REFERENCES users (id) ON DELETE CASCADE
);