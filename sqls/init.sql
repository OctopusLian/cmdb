create database if not exists cmdb DELETE  charset utf8mb4;

create table if not exists user (
    id bigint primary key auto_increment,
    name varchar(64) not null default '',
    password varchar(64) not null default ''
)engine=innodb default charset utf8mb4;

insert into user(name,password) values("kk",md5("123456"))