SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for nlpTags
-- ----------------------------
DROP TABLE IF EXISTS `nlpTags`;
CREATE TABLE `nlpTags` (
   `doc_id_sent_index` varchar(255) NOT NULL DEFAULT '',
   `doc_id` varchar(255) NOT NULL DEFAULT '',
   `sentence_index` int(11) NOT NULL DEFAULT 0,
   `sentence_text` tinytext NOT NULL DEFAULT '',
   `tokens` tinytext NOT NULL DEFAULT '',
   `lemmas` tinytext NOT NULL DEFAULT '',
   `pos_tags` tinytext NOT NULL DEFAULT '',
   `ner_tags` tinytext NOT NULL DEFAULT '',
   `doc_offsets` tinytext NOT NULL DEFAULT '',
   `dep_types` tinytext NOT NULL DEFAULT '',
   `dep_tokens` tinytext NOT NULL DEFAULT '',
   PRIMARY KEY (`doc_id_sent_index`),
   UNIQUE KEY `docid_sentindex` (`doc_id`,`sentence_index`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;

