CREATE TABLE
    `t_upload_file` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `file_key` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '文件MD5',
        `file_name` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '文件名称',
        `file_size` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '文件大小',
        `url` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '文件地址',
        `upload_times` INT UNSIGNED NOT NULL DEFAULT 1 COMMENT '上传次数',
        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
        PRIMARY KEY (`id`),
        UNIQUE KEY `uniq_key` (`file_key`),
        KEY `idx_create` (`create_at`),
        KEY `idx_update` (`update_at`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '上传文件';