CREATE TABLE `t_message_user` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
    `to_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '消息接收人ID',
    `mkey` VARCHAR(45) NOT NULL DEFAULT '' COMMENT '消息KEY',
    `un_reads` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '未读消息数',
    `last_time` BIGINT NOT NULL DEFAULT 0 COMMENT '最后消息时间',
    `last_uid` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后发送人',
    `last_msg` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '最后消息摘要',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_uid_to`(`user_id`,`to_id`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '消息表';

CREATE TABLE `t_message_content` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `mkey` VARCHAR(45) NOT NULL DEFAULT '' COMMENT '消息KEY',
    `from_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '发送人',
    `recv_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '接收人',
    `is_read` TINYINT NOT NULL DEFAULT 0 COMMENT '是否已读',
    `content` TEXT NOT NULL COMMENT '消息内容',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    KEY `idx_mkey`(`mkey`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '消息内容表';