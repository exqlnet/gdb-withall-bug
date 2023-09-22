-- auto-generated definition
create table post
(
    id         bigint auto_increment primary key,
    content    text null,
    user_id    bigint null,
    created_at datetime default (now()) null,
    updated_at datetime default (now()) null,
    deleted_at datetime null
);


create table user
(
    id         bigint auto_increment primary key,
    username   varchar(255) null,
    created_at datetime default (now()) null,
    updated_at datetime default (now()) null,
    deleted_at datetime null
);

insert into user (id, username, created_at, updated_at, deleted_at)
values (1, 'user1', '2023-09-22 12:00:00', '2023-09-22 12:00:00', null);

insert into post (id, content, user_id, created_at, updated_at, deleted_at)
VALUES (1, 'hello', 1, '2023-09-22 12:00:00', '2023-09-22 12:00:00', null);