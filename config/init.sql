CREATE TABLE `t_config` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `conf_key` VARCHAR(30) NOT NULL DEFAULT '' COMMENT '分组KEY',
    `group_key` VARCHAR(30) NOT NULL DEFAULT '' COMMENT '中间key',
    `item_key` VARCHAR(30) NOT NULL DEFAULT '' COMMENT 'KEY',
    `val` TEXT NOT NULL COMMENT '具体的值',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_conf_group_item` (`conf_key`, `group_key`,`item_key`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '配置信息表';