CREATE TABLE `station_status_history` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `station_sn` varchar(32) NOT NULL DEFAULT '',
  `station_model` varchar(32) NOT NULL DEFAULT '',
  `email` varchar(64) NOT NULL DEFAULT '' COMMENT 'email',
  `user_id` varchar(64) NOT NULL COMMENT 'user_id',
  `ab_code` varchar(32) NOT NULL DEFAULT '',
  `conn_type` varchar(32) NOT NULL DEFAULT '' COMMENT 'server: EU US',
  `did` varchar(32) NOT NULL DEFAULT '',
  `env` varchar(32) NOT NULL DEFAULT '' COMMENT 'server: EU US',
  `is_inner` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0 not 1 yes',
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `create_time` int(11) NOT NULL DEFAULT '0',
  `update_time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `station_sn` (`station_sn`)
) ENGINE=InnoDB AUTO_INCREMENT=360731 DEFAULT CHARSET=utf8;