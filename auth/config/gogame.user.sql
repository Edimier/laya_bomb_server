CREATE TABLE `user` (
  `numid` int(10) unsigned NOT NULL DEFAULT '0',
  `mac` varchar(64) NOT NULL,
  `name` varchar(255) NOT NULL,
  `pwd` varchar(255) NOT NULL,
  `registTime` datetime DEFAULT NULL,
  PRIMARY KEY (`numid`,`mac`),
  UNIQUE KEY `getuser2` (`numid`,`mac`) USING HASH,
  KEY `getuser1` (`numid`,`pwd`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;