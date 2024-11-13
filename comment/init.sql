CREATE TABLE `t_comment` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
    `target_type` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论类型',
    `target_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论类型ID',
    `reply_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '回复评论ID',
    `content` TEXT COMMENT '评论内容',
    `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态,0 待审核，10，审核完成，20 仅自己可见，100 删除',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    KEY `idx_type_id_rid`(`target_type`,`target_id`,`reply_id`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '评论表';

CREATE TABLE `t_comment_stat` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `target_type` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '收藏类型',
    `target_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '收藏ID',
    `target_counts` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '收藏次数',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_type_id`(`target_type`, `target_id`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '评论统计';