CREATE DATABASE test CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE user (
    id int NOT NULL AUTO_INCREMENT,
    username varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    nickname varchar(255),
    avatar longblob,
    PRIMARY KEY (id)
);
