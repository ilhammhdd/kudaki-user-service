CREATE DATABASE IF NOT EXISTS kudaki_user_domain DEFAULT COLLATE = utf8_general_ci;

CREATE USER IF NOT EXISTS 'kudakiuser'@'localhost' IDENTIFIED BY 'kudakiuserrocks';

GRANT ALL PRIVILEGES ON kudaki_user_domain.* TO 'kudakiuser'@'localhost'
WITH GRANT OPTION;

USE kudaki_user_domain;

CREATE TABLE IF NOT EXISTS profiles(
    `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `uuid` VARCHAR(64) UNIQUE,
    `first_name` VARCHAR(255),
    `last_name` VARCHAR(255),
    `gender` ENUM('MALE','FEMALE'),
    `phone_number` VARCHAR(255),
    `address` VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS users (
    `id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `uuid` VARCHAR(64) UNIQUE,
    `username` VARCHAR(255) UNIQUE,
    `email` VARCHAR(255) UNIQUE,
    `password` VARCHAR(255),
    `token` VARCHAR(255),
    `account_type` VARCHAR(255),
    `role` ENUM('KUDAKI_TEAM','SHOP_OWNER','SHOP_KEEPER','HIKER')
);
