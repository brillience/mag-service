#!/bin/bash
SET NAMES utf8mb4;
   SET FOREIGN_KEY_CHECKS = 0;

   -- ----------------------------
   -- Table structure for nlpTags
   -- ----------------------------
   DROP TABLE IF EXISTS `nlpTags`;
   CREATE TABLE `nlpTags` (
   `doc_id` varchar(255) NOT NULL DEFAULT '',
   `nlp_tags` text NOT NULL DEFAULT '' COMMENT 'json字符串',
   PRIMARY KEY (`doc_id`),
   INDEX `doc_id`(`doc_id`(255))
   ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

   CREATE TABLE `user` (
   `id` bigint NOT NULL AUTO_INCREMENT,
   `username` varchar(255) NOT NULL DEFAULT '' COMMENT '账户',
   `nick` varchar(255)  NOT NULL DEFAULT '' COMMENT '昵称',
   `password` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户密码',
   `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
   `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   PRIMARY KEY (`id`),
   KEY `number_unique` (`username`)
   ) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;

   SET FOREIGN_KEY_CHECKS = 1;