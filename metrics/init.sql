CREATE TABLE `t_metrics_conf` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(30) NOT NULL DEFAULT '' COMMENT '名称',
    `conf_key` VARCHAR(60) NOT NULL DEFAULT '' COMMENT '配置的key',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_key`(`conf_key`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '打点配置';

CREATE TABLE `t_metrics_stat` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `conf_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '打点的ID',
    `platform` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '平台',
    `week` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '上线后的周',
    `year` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '年',
    `month` TINYINT NOT NULL DEFAULT 0 COMMENT '月',
    `day` TINYINT NOT NULL DEFAULT 0 COMMENT '日',
    `date` DATE NOT NULL COMMENT '日期',
    `views` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '展示次数',
    `clicks` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '点击次数',
    `user_views` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '展示用户数',
    `user_clicks` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '点击用户数',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_id_date_platform` (`conf_id`,`date`,`platform`),
    KEY `idx_id_week_platform` (`conf_id`,`week`,`platform`),
    KEY `idx_id_year_month_platform` (`conf_id`,`year`, `month`, `platform`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '打点统计数据';