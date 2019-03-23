CREATE TABLE `user_name` (
  `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `varchar` VARCHAR(255)  NOT NULL DEFAULT '' COMMENT '字符串示例',
  `bigint` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '大数值示例',
  `is_deleted` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '软删标识',
  `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
  PRIMARY KEY (`id`),
) ENGINE=INNODB DEFAULT CHARSET=utf8 COMMENT='建表参考示例';
