CREATE TABLE
    `t_friend_links` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '名称',
        `link_type` VARCHAR(10) NOT NULL DEFAULT 'WEB' COMMENT '类型：网站，app等',
        `logo` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'logo',
        `url` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '地址',
        `ios` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'ios下载地址',
        `google_pay` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'googlepaly 地址',
        `bg_img` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '背景地址',
        `content` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT '描述',
        `is_show` TINYINT NOT NULL DEFAULT 0 COMMENT '是否展示',
        `sort_index` INT NOT NULL DEFAULT 0 COMMENT '排序',
        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
        PRIMARY KEY (`id`),
        KEY `idx_create` (`create_at`),
        KEY `idx_update` (`update_at`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '友情链接表';