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

select * from t_auth
