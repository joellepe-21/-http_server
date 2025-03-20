create table users (
id serial primary key,
login varchar(50) not null unique,
password text not null unique
);
create table articles (
id serial primary key,
name varchar(255) not null unique,
article text not null
);
