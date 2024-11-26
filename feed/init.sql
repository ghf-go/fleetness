CREATE TABLE `t_feed` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '发布人ID',
    `feed_type` TINYINT NOT NULL DEFAULT 0 COMMENT '类型，0blog,10投票,11多选投票',
    `title` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '标题',
    `imgs` VARCHAR(2048) NOT NULL DEFAULT '' COMMENT '图片',
    `content` TEXT NOT NULL COMMENT '内容',
    `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态,0 待审核，10，审核完成，20 仅自己可见，100 删除',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    KEY `idx_user` (`user_id`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '动态表';

CREATE TABLE `t_feed_vote` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `feed_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'feedid',
    `name` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '选项',
    `votes` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '投票人数',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    KEY `idx_feedid`(`feed_id`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '动态投票表';

CREATE TABLE `t_feed_vote_log` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `feed_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'feedid',
    `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '发布人ID',
    `item_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '投票选项id',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_feed_user_item` (`feed_id`,`item_id`,`user_id`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '动态投票表';