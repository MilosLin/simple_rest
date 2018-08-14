/*
  account_db
*/

CREATE DATABASE `account_db`;

USE `account_db`;

# 會員資料表
CREATE TABLE IF NOT EXISTS `account_db`.`user` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
  `account` varchar(50) NOT NULL COMMENT '帳號',
  `password` varchar(50) NOT NULL COMMENT '密碼',
  `created_ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '建立時間',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB COMMENT='會員資料表';

INSERT INTO `account_db`.`user` (`id`, `account`, `password`) VALUES (0, "user1", "password");
