-- USE `authservice`;

create table `oauth_user` (
    `oauth_user_id` integer not null auto_increment,
    `full_name` varchar(255),
    `login_count` integer not null,
    `oauth_id` varchar(255),
    primary key (`oauth_user_id`)
) engine=InnoDB DEFAULT CHARSET=utf8;
