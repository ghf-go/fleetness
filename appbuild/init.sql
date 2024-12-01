CREATE TABLE
    `t_app_build_module` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '模块名称',
        `module_desc` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT '描述',
        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
        PRIMARY KEY (`id`),
        KEY `idx_create` (`create_at`),
        KEY `idx_update` (`update_at`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '生成配置表';

CREATE TABLE
    `t_app_build_template` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `platform` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '平台名称',
        `module_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '模块ID',
        `fname` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '文件名称',
        `fdesc` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '描述',
        `ftemplate` TEXT NOT NULL COMMENT '文件模版',
        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
        PRIMARY KEY (`id`),
        KEY `idx_platform_module` (`platform`, `module_id`),
        KEY `idx_create` (`create_at`),
        KEY `idx_update` (`update_at`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '代码模版表';