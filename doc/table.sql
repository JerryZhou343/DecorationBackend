DROP TABLE IF EXISTS `t_auth`;
CREATE TABLE `t_auth`(
  `id` BIGINT auto_increment ,
  `user_name` VARCHAR(64) NOT NULL,
  `password` VARCHAR(200) NOT NULL,
  `created` INT    NOT NULL,
  `updated` INT default  0,
  primary key (`id`),
  unique key(`user_name`)
);



DROP TABLE IF EXISTS `t_slat`;
CREATE TABLE `t_slat`(
  `id` BIGINT ,
  `slat` varchar(6) NOT NULL,
  `created` INT NOT NULL,
  `updated` INT default  0,
  primary key (`id`)
);



DROP TABLE IF EXISTS `t_category`;
CREATE TABLE `t_category`(
    `id` BIGINT  auto_increment not null,
    `name` varchar(64) not null,
    `parent_id` BIGINT not null default 0,
    `priority` tinyint not null default 1 comment '展示顺序',
    `state` tinyint not null default 1 comment '0 为激活，1激活',
    `remark` varchar(200),
    `created` INT not null,
    `updated` INT default  0,
    `operator_id` bigint not null,
    primary key (`id`)
);



DROP TABLE IF EXISTS `t_case`;
CREATE TABLE `t_case`(
	  `id` BIGINT auto_increment not null,
    `name` varchar(64) not null,
    `price` int not null default 0,
    `type` int not null default 0 comment '0 个人，1 公司',
    `owner_name` varchar(64) default null,
    `phone_number` varchar(20) default null,
    `addr` varchar(200) default null,
    `state` tinyint not null default 1 comment '1 激活，0 未激活'
    `created` int not null,
    `updated` int default 0,
    primary key (`id`)
);

drop table if exists `t_case_category`;
create table `t_case_category`(
	  `id` bigint auto_increment not null,
    `case_id` bigint not null,
    `category_id` bigint not null,
    `created` int not null,
    `updated` int default 0,
    `state` tinyint not null default 1 comment '0 失效, 1 有效',
    primary key (`id`)
);

drop table if exists `t_pic_category`;
create table `t_pic_category`(
	  `id` bigint auto_increment not null,
    `pic_id` bigint not null,
    `category_id` bigint not null,
    `created` int not null,
    `updated` int default 0,
	  `state` tinyint default 1 not null comment '0 失效，1 有效',
    primary key (`id`),
    unique key(`pic_id`,`category_id`)
);

drop table if exists `t_picture`;
create table `t_picture`(
	  `id` bigint auto_increment not null,
    `case_id` bigint not null comment '归属案例ID',
    `name` varchar(64) default null,
    `addr` varchar(1024) not null,
    `state` tinyint default 1 not null comment '0失效,1有效'
    `remark` varchar(128) default  null,
    `created` int default 0,
    primary key (`id`)
);
