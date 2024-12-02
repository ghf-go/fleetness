CREATE TABLE
    `t_news` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `title` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '标题',
        `sub_title` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '附标题',
        `category_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '分类',
        `img` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT '图片',
        `content` TEXT NOT NULL COMMENT '内容',
        `author` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '作者',
        `refer` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT '应用地址',
        `is_publish` TINYINT NOT NULL DEFAULT 0 COMMENT '是否发布',
        `is_del` TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除',
        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
        PRIMARY KEY (`id`),
        KEY `idx_category` (`category_id`),
        KEY `idx_create` (`create_at`),
        KEY `idx_update` (`update_at`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '新闻表';