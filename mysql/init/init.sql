CREATE DATABASE IF NOT EXISTS sample-db;

create table if not exists todos (
    id integer auto_increment primary key,
    item varchar(255),
    is_done tinyint(1) DEFAULT '0',
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
