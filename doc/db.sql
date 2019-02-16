DROP TABLE IF EXISTS `t_auth`;
CREATE TABLE `t_auth`(
  `id` BIGINT auto_increment ,
  `user_name` VARCHAR(64) NOT NULL,
  `password` VARCHAR(200) NOT NULL,
  `created` INT    NOT NULL,
  `updated` INT,
  primary key (`id`),
  unique key(`user_name`)
);

DROP TABLE IF EXISTS `t_slat`;
CREATE TABLE `t_slat`(
  `id` BIGINT ,
  `slat` varchar(6) NOT NULL,
  `created` INT NOT NULL,
  `updated` INT,
  primary key (`id`)
);

insert into t_auth(`user_name`,`password`, `created`) values('jerry',  md5('123456123456'),1550297154);
insert into t_slat(`id`,`slat`,`created`)values (1,'123456',1550297154);


DROP TABLE IF EXISTS `t_category`;
CREATE TABLE `t_category`(
    `id` BIGINT  auto_increment not null,
    `web_name` varchar(64) not null,
    `sys_name` varchar(64) not null,
    `parent_id` BIGINT not null default 0,
    `priority` tinyint not null default 1 comment '展示顺序',
    `operator_type` tinyint not null comment '0 不可以修改,1 可以修改',
    `state` tinyint not null default 1 comment '0 为激活，1激活',
    `remark` varchar(200),
    `created` BIGINT not null,
    `updated` BIGINT not null,
    `operator_id` bigint not null,
    primary key (`id`)
);

DROP TABLE IF EXISTS `t_case`;
CREATE TABLE `t_case`(
	  `id` BIGINT auto_increment not null,
    `name` varchar(64) not null,
    `price` int not null default 0,
    `type` int not null default 0 comment ''0 个人，1 公司'',
    `owner_name` varchar(64) default null,
    `phone_number` varchar(20) default null,
    `add` varchar(200) default null,
    `created` bigint not null,
    `updated` bigint default null,
    primary key (`id`)
);

drop table if exists `t_case_category`;
create table `t_case_category`(
	  `id` bigint auto_increment not null,
    `case_id` bigint not null,
    `category_id` bigint not null,
    `created` bigint not null,
    `updated` bigint default 0,
    `state` tinyint not null default 1 comment ''0 失效, 1 有效'',
    primary key (`id`)
);

drop table if exists `t_pic_category`;
create table `t_pic_category`(
	  `id` bigint auto_increment not null,
    `pic_id` bigint not null,
    `category_id` bigint not null,
    `created` bigint not null,
    `updated` bigint not null,
	  `state` tinyint not null comment '0 失效，1 有效',
    primary key (`id`),
    unique key(`pic_id`,`category_id`)
);

drop table if exists `t_picture`;
create table `t_picture`(
	  `id` bigint auto_increment not null,
    `case_id` bigint not null comment '归属案例ID',
    `name` varchar(64) default null,
    `addr` varchar(1024) not null,
    primary key (`id`)
);
