CREATE TABLE
    `t_category` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '分类名称',
        `target_type` INT NOT NULL DEFAULT 0 COMMENT '目标类型',
        `parent_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级分类ID',
        `is_show` TINYINT NOT NULL DEFAULT 0 COMMENT '是否展示',
        `sort_index` INT NOT NULL DEFAULT 0 COMMENT '排序',
        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
        PRIMARY KEY (`id`),
        UNIQUE KEY `uniq_type_pid_name` (`target_type`, `parent_id`, `name`),
        KEY `idx_create` (`create_at`),
        KEY `idx_update` (`update_at`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '分类列表';