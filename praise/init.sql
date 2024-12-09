CREATE TABLE
    `t_praise_stat` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `target_type` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '点赞类型',
        `target_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '点赞ID',
        `target_counts` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '点赞次数',
        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
        PRIMARY KEY (`id`),
        UNIQUE KEY `uniq_type_id` (`target_type`, `target_id`),
        KEY `idx_create` (`create_at`),
        KEY `idx_update` (`update_at`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '点赞统计';

CREATE TABLE
    `t_praise` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
        `target_type` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '点赞类型',
        `target_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '点赞ID',
        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
        PRIMARY KEY (`id`),
        UNIQUE KEY `uniq_uid,type_id` (`user_id`, `target_type`, `target_id`),
        KEY `idx_create` (`create_at`),
        KEY `idx_update` (`update_at`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '点赞';