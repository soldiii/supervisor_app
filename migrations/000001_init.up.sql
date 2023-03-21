CREATE TABLE users
(
    id serial not null primary key,
    email varchar(30) not null unique,
    name varchar(30) not null,
    surname varchar(30) not null,
    patronymic varchar(30),
    reg_date_time timestamp not null,
    password varchar(30) not null
);