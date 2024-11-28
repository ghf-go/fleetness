CREATE TABLE `t_app_ver` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `app_ver` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '版本',
    `apk_url` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'android 下载地址',
    `wgt_url` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '热更新下载地址',
    `ver_content` TEXT NOT NULL COMMENT '版本描述',
    `is_online` TINYINT NOT NULL DEFAULT 0 COMMENT '是否上架',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_ver`(`app_ver`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '应用版本';
