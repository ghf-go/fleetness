CREATE TABLE `t_group` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `target_type` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '分组类型',
    `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
    `group_name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '分组名称',
    `items` INT NOT NULL DEFAULT 0 COMMENT '分组内条数',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    KEY `idx_type_id`(`target_type`, `user_id`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '分组表';

CREATE TABLE `t_group_item` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `group_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '分组id',
    `target_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '目标ID',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_type_id`(`group_id`, `target_id`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '分组表数据表';