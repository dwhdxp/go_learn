DROP TABLE IF EXISTS user;
CREATE TABLE `user` (
    `id` int unsigned auto_increment,
    `name` varchar(64) not null,
    `age` tinyint unsigned,
    `email` varchar(256) unique,
    `create_time` timestamp default current_timestamp,
    `update_time` timestamp default current_timestamp on update current_timestamp,
    primary key(id)
);