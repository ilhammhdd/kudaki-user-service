CREATE DATABASE IF NOT EXISTS kudaki_user_domain DEFAULT COLLATE = utf8_general_ci;

CREATE USER IF NOT EXISTS 'kudakiuser'@'localhost' IDENTIFIED BY 'kudakiuserrocks';

GRANT ALL PRIVILEGES ON kudaki_user_domain.* TO 'kudakiuser'@'localhost'
WITH GRANT OPTION;

USE kudaki_user_domain;

CREATE TABLE IF NOT EXISTS users (
    `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `uuid` VARCHAR(64) UNIQUE,
    `email` VARCHAR(255) UNIQUE,
    `password` VARCHAR(255),
    `token` VARCHAR(255),
    `role` ENUM('USER','KUDAKI_TEAM','ORGANIZER'),
    `phone_number` VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS profiles(
    `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `user_uuid` VARCHAR(64) UNIQUE,
    `uuid` VARCHAR(64) UNIQUE,
    `full_name` VARCHAR(255),
    `photo` VARCHAR(255),
    `reputation` INT(20) UNSIGNED,

    FOREIGN KEY(user_uuid)
    REFERENCES users(uuid)
    ON DELETE CASCADE
);
