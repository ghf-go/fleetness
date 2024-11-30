CREATE TABLE
    `t_user` (
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

CREATE TABLE
    `t_user_bind` (
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
        UNIQUE KEY `uniq_val` (`bind_val`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '用户绑定账号信息';

CREATE TABLE
    `t_user_info` (
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
        UNIQUE KEY `uniq_uid_key` (`user_id`, `ukey`),
        KEY `idx_create` (`create_at`),
        KEY `idx_update` (`update_at`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '用户信息表';

CREATE TABLE
    `t_user_cash` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
        `ukey` VARCHAR(50) NOT NULL DEFAULT '' COMMENT 'key',
        `val` INT NOT NULL DEFAULT 0 COMMENT 'val',
        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
        PRIMARY KEY (`id`),
        UNIQUE KEY `uniq_uid_key` (`user_id`, `ukey`),
        KEY `idx_create` (`create_at`),
        KEY `idx_update` (`update_at`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '用户资金表';

CREATE TABLE
    `t_user_addr` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
        `mobile` varchar(60) NOT NULL DEFAULT '' COMMENT '手机',
        `consignee` varchar(60) NOT NULL DEFAULT '' COMMENT '收货人',
        `province` varchar(100) NOT NULL DEFAULT '0' COMMENT '省份',
        `city` varchar(100) NOT NULL DEFAULT '0' COMMENT '城市',
        `district` varchar(100) NOT NULL DEFAULT '0' COMMENT '地区',
        `address` varchar(120) NOT NULL DEFAULT '' COMMENT '地址',
        `is_default` tinyint (1) DEFAULT '0' COMMENT '默认收货地址',
        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
        PRIMARY KEY (`id`),
        KEY `idx_user` (`user_id`),
        KEY `idx_create` (`create_at`),
        KEY `idx_update` (`update_at`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '用户收货地址';

CREATE TABLE
    `t_user_cash_log` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
        `ukey` VARCHAR(50) NOT NULL DEFAULT '' COMMENT 'key',
        `val` INT NOT NULL DEFAULT 0 COMMENT 'val',
        `content` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '描述',
        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
        PRIMARY KEY (`id`),
        key `idx_user_key` (`user_id`, `ukey`),
        KEY `idx_create` (`create_at`),
        KEY `idx_update` (`update_at`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '用户资金日志表';

CREATE TABLE
    `t_admin_user` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `login_name` VARCHAR(30) NOT NULL DEFAULT '' COMMENT '登录名称',
        `nick_name` VARCHAR(30) NOT NULL DEFAULT '' COMMENT '昵称',
        `passwd` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '密码',
        `pass_sign` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '密码加盐',
        `tfa_key` VARCHAR(16) NOT NULL DEFAULT '' COMMENT '二次验证码key',
        `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态：100删除，90禁止登录',
        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
        PRIMARY KEY (`id`),
        UNIQUE KEY `uniq_loginname` (`login_name`),
        KEY `idx_create` (`create_at`),
        KEY `idx_update` (`update_at`)
    ) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '管理员用户';