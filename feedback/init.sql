CREATE TABLE `t_feedback` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
    `imgs` VARCHAR(2048) NOT NULL DEFAULT '' COMMENT '图片',
    `content` VARCHAR(2048) NOT NULL DEFAULT '' COMMENT '发布的内容',
    `replay_content` TEXT COMMENT '后台回复的内容',
    `is_replay` TINYINT NOT NULL DEFAULT 0 COMMENT '是否已回复',
    `replay_at` DATETIME DEFAULT NULL COMMENT '回复时间',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    KEY `idx_user`(`user_id`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '意见反馈';