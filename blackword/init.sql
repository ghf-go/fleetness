CREATE TABLE
    `t_black_word` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `word` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '名称',
        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
        PRIMARY KEY (`id`),
        UNIQUE KEY `uniq_word` (`word`),
        KEY `idx_create` (`create_at`),
        KEY `idx_update` (`update_at`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '敏感词';