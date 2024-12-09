CREATE TABLE
    `t_tags` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `target_type` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型',
        `name` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '名称',
        `logo` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'LOGO',
        `bg_color` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '背景颜色',
        `font_color` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '字体颜色',
        `sum_times` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '次数',
        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
        PRIMARY KEY (`id`),
        UNIQUE KEY `uniq_type_name` (`target_type`, `name`),
        KEY `idx_create` (`create_at`),
        KEY `idx_update` (`update_at`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '标签';

CREATE TABLE
    `t_tags_ids` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `tag_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'tagID',
        `target_type` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型',
        `target_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '目标ID',
        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
        PRIMARY KEY (`id`),
        UNIQUE KEY `uniq_tag_id` (`tag_id`, `target_id`),
        KEY `idx_target` (`target_type`, `target_id`),
        KEY `idx_create` (`create_at`),
        KEY `idx_update` (`update_at`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '标签汇总';