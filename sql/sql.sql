CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id varchar(36) default(uuid()) primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(50) not null unique,
    criado_em timestamp default current_timestamp
) ENGINE=INNODB

CREATE TABLE follows (
    user_id varchar(36) not null,
    follower_id varchar(36) not null,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,

    PRIMARY KEY(user_id, follower_id)
);