CREATE TABLE `user`
(
    `id`       int(10) unsigned NOT NULL AUTO_INCREMENT,
    `mobile`   char(11)     NOT NULL DEFAULT '' COMMENT '手机号',
    `nickname` varchar(64)  NOT NULL COMMENT '昵称',
    `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
    `sex`      int(10) unsigned DEFAULT NULL COMMENT '性别',
    `avatar`   varchar(255)          DEFAULT NULL COMMENT '头像',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uni_mobile` (`mobile`) USING BTREE,
    KEY        `idx_mobile_pwd` (`mobile`,`password`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;