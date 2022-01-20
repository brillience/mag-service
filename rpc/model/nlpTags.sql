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

SET FOREIGN_KEY_CHECKS = 1;

