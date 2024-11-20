CREATE TABLE `t_mall_category` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '分类名称',
    `parent_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '父分类ID',
    `logo` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '图标',
    `is_show` TINYINT NOT NULL DEFAULT 0 COMMENT '是否显示',
    `sort_index` INT NOT NULL DEFAULT 0 COMMENT '排序',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '商城分类表';

CREATE TABLE `t_mall_brand` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '品牌名称',
    `logo` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '品牌LOGO',
    `more` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT '品牌描述',
    `first_char` CHAR(1) NOT NULL DEFAULT '' COMMENT '首字符',
    `is_show` TINYINT NOT NULL DEFAULT 0 COMMENT '是否显示',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '商城品牌';

CREATE TABLE `t_mall_category_brand` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `cate_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '分类ID',
    `brand_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '品牌ID',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`),
    UNIQUE KEY `uniq_cate_brand`(`cate_id`,`brand_id`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '商城分类品牌';

CREATE TABLE `t_mall_freight` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '运费规则名称',
    `is_free` TINYINT NOT NULL DEFAULT 0 COMMENT '是否包邮',
    `default_price` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '默认运费',
    `default_num` INT UNSIGNED NOT NULL DEFAULT 1 COMMENT '默认数量',
    `over_price` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '超出运费',
    `over_num` INT UNSIGNED NOT NULL DEFAULT 1 COMMENT '超出数量',
    `forbid` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT '禁运地区',
    `excluding` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT '例外运费',
    `exclud_price` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '默认运费',
    `exclud_num` INT UNSIGNED NOT NULL DEFAULT 1 COMMENT '默认数量',
    `exclud_over_price` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '超出运费',
    `exclud_over_num` INT UNSIGNED NOT NULL DEFAULT 1 COMMENT '超出数量',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '商城运费规则表';

CREATE TABLE `t_mall_spec` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '名称',
    `root_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '根ID',
    `parent_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '父ID',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`),
    KEY `idx_root` (`root_id`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '商城规格表';

CREATE TABLE `t_mall_goods` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '商品名字',
    `brand_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '品牌ID',
    `cate_id1` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '一级分类ID',
    `cate_id2` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '二级分类ID',
    `cate_id3` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '三级分类ID',
    `freight_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '运费规则ID',
    `spec_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '规格ID',
    `sales` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '销量',
    `market_price` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '市场价',
    `sell_price` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '售价',
    `cost_price` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '成本价',
    `stock` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '库存',
    `is_onlone` TINYINT NOT NULL DEFAULT 0 COMMENT '是否上架',
    `imgs` VARCHAR(2048) NOT NULL DEFAULT '' COMMENT '图片',
    `detail` TEXT NOT NULL COMMENT '商品描述',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '商城商品表';

CREATE TABLE `t_mall_goods_spec` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `goods_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品ID',
    `logo` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '规格图片',
    `item_key` VARCHAR(66) NOT NULL DEFAULT '' COMMENT '规格ID,_分隔',
    `item_name` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '规格名称',
    `market_price` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '市场价',
    `sell_price` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '售价',
    `cost_price` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '成本价',
    `stock` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '库存',
    `sales` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '销量',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '创建IP',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `update_ip` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '更新IP',
    PRIMARY KEY (`id`),
    KEY `idx_create` (`create_at`),
    KEY `idx_update` (`update_at`)
) ENGINE = innodb DEFAULT CHARSET = utf8mb4 COMMENT = '商城商品表';

