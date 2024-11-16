CREATE TABLE `t_user` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `nick_name` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '昵称',
    `avatar` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '头像',
    `passwd` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '密码',
    `pass_sign` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '密码加盐',
    `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态，0正常',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '用户';

CREATE TABLE `t_user_bind` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
    `bind_type` TINYINT NOT NULL DEFAULT 0 COMMENT '类型：1 手机号，2邮箱，2微信，4微博',
    `bind_val` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '绑定账号，例如手机号或者微信 openid',
    `bind_display` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '绑定账号的名称，例如微信昵称',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`),
    KEY `idx_uid` (`user_id`),
    UNIQUE KEY `uniq_val`(`bind_val`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '用户绑定账号信息';

CREATE TABLE `t_user_info` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
    `ukey` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '',
    `uval` TEXT NOT NULL COMMENT '内容',
    `newval` TEXT NOT NULL COMMENT '新内容',
    `is_audit` TINYINT NOT NULL DEFAULT 0 COMMENT '是否已审核',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_uid_key` (`user_id`,`ukey`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '用户信息表';