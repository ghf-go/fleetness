CREATE TABLE `t_coupon` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '优惠券名字',
    `type` TINYINT NOT NULL DEFAULT 0 COMMENT '优惠券类型：1代金券，2打折券',
    `val` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '代金券价格或者折扣',
    `rule` TINYINT NOT NULL DEFAULT 0 COMMENT '规则:1满减，2指定分类，3指定商品',
    `expire` INT NOT NULL DEFAULT 0 COMMENT '有效期：0一直邮箱',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '优惠券';



CREATE TABLE `t_coupon_user` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '用户优惠券';