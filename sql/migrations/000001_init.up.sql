create table categories
(
    id          varchar(36) not null primary key,
    name        text        not null,
    description text
);

create table courses
(
    id          varchar(36)    not null primary key,
    category_id varchar(36)    not null,
    name        text           not null,
    description text,
    price       decimal(10, 2) not null,
    foreign key (category_id) references categories (id)
)