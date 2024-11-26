CREATE TABLE `t_lottery` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '抽奖活动名称',
    `logo` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'logo',
    `kind` TINYINT NOT NULL DEFAULT 10 COMMENT '类型，10轮盘抽奖',
    `day_limit` INT NOT NULL DEFAULT 0 COMMENT '每日限制次数，0无限制',
    `is_open` TINYINT NOT NULL DEFAULT 0 COMMENT '是否开启',
    `content` TEXT COMMENT '描述',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '抽奖信息';

CREATE TABLE `t_lottery_item` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `lottery_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '抽奖ID',
    `name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '选项名称',
    `logo` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '图片',
    `content` VARCHAR(500) NOT NULL DEFAULT '' COMMENT '描述',
    `rate` INT NOT NULL DEFAULT 0 COMMENT '中奖概率',
    `is_online` TINYINT NOT NULL DEFAULT 0 COMMENT '是否开启',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    KEY `idx_lottery_id` (`lottery_id`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '抽奖配置';

CREATE TABLE `t_lottery_log` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `lottery_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '抽奖ID',
    `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
    `item_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '中奖选项ID',
    `day` DATE NOT NULL COMMENT '日期',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    KEY `idx_lottery_user_day`(`lottery_id`,`user_id`,`day`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '抽奖日志';

