CREATE TABLE
    `t_novel_info` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `name` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '名称',
        `logo` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'LOGO',
        `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '作者',
        `subscribe` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '订阅人数',
        `words` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '字数',
        `section_price` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '章节价格，单位分',
        `free_section` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '免费章节',
        `last_section_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '最新章节',
        `is_over` TINYINT NOT NULL DEFAULT 0 COMMENT '是否完结',
        `is_free` TINYINT NOT NULL DEFAULT 0 COMMENT '是否免费',
        `is_publish` TINYINT NOT NULL DEFAULT 0 COMMENT '是否上架',
        `is_audit` TINYINT NOT NULL DEFAULT 0 COMMENT '是否已经审核',
        `total_income` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '总收入',
        `today_income` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '今日收入',
        `content` TEXT NOT NULL COMMENT '描述',
        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
        PRIMARY KEY (`id`),
        KEY `idx_user` (`user_id`),
        UNIQUE KEY `uniq_name` (`name`),
        KEY `idx_create` (`create_at`),
        KEY `idx_update` (`update_at`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '小说列表';

CREATE TABLE
    `t_novel_section` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `novel_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '小说ID',
        `parent_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '父章节ID',
        `name` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '名称',
        `section_index` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '章节排序',
        `words` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '字数',
        `is_publish` TINYINT NOT NULL DEFAULT 0 COMMENT '是否发布',
        `is_audit` TINYINT NOT NULL DEFAULT 0 COMMENT '是否已经审核',
        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
        PRIMARY KEY (`id`),
        KEY `idx_novel_index` (`novel_id`, `section_index`),
        KEY `idx_create` (`create_at`),
        KEY `idx_update` (`update_at`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '小说列表';

CREATE TABLE
    `t_novel_section_content` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `section_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '章节ID',
        `content` TEXT NOT NULL COMMENT '章节内容',
        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
        PRIMARY KEY (`id`),
        UNIQUE KEY `uniq_section` (`section_id`),
        KEY `idx_create` (`create_at`),
        KEY `idx_update` (`update_at`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '小说章节内容';

CREATE TABLE
    `t_novel_subscribe` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户',
        `novel_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '小说ID',
        `read_section` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '阅读进度',
        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
        PRIMARY KEY (`id`),
        KEY `idx_user` (`user_id`),
        KEY `idx_create` (`create_at`),
        KEY `idx_update` (`update_at`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '订阅用户';

CREATE TABLE
    `t_novel_read_history` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户',
        `novel_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '小说ID',
        `section_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '章节ID',
        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
        PRIMARY KEY (`id`),
        KEY `idx_user` (`user_id`),
        KEY `idx_create` (`create_at`),
        KEY `idx_update` (`update_at`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '阅读历史';

CREATE TABLE
    `t_novel_buy_log` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户',
        `author_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '作者ID',
        `novel_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '小说ID',
        `section_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '章节ID',
        `section_price` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '章节价格，单位分',
        `status` TINYINT NOT NULL DEFAULT 0 COMMENT '0待提现，1申请中，2已提现',
        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
        PRIMARY KEY (`id`),
        UNIQUE KEY `idx_user_novel_section` (`user_id`, `novel_id`, `section_id`),
        KEY `idx_author_novel_section` (`author_id`, `novel_id`, `section_id`),
        KEY `idx_create` (`create_at`),
        KEY `idx_update` (`update_at`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '购买记录';

CREATE TABLE
    `t_novel_tx_log` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户',
        `date` DATE NOT NULL COMMENT '申请日期',
        `status` TINYINT NOT NULL DEFAULT 0 COMMENT '0申请中，10已审核，20完成',
        `pay_at` DATETIME DEFAULT NULL COMMENT '支付日期',
        `amount` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '申请提现金额',
        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
        PRIMARY KEY (`id`),
        UNIQUE KEY `uniq_user_date` (`user_id`, `date`),
        KEY `idx_create` (`create_at`),
        KEY `idx_update` (`update_at`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '提现记录';