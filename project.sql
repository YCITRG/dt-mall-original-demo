--- web

DROP SCHEMA IF EXISTS yecao_mall_web_bff;
CREATE SCHEMA yecao_mall_web_bff DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci;


-- 库存
DROP SCHEMA IF EXISTS yecao_mall_inventory;
CREATE SCHEMA yecao_mall_inventory DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci;

USE
yecao_mall_inventory;

DROP TABLE IF EXISTS tb_inventory;
CREATE TABLE tb_inventory
(
    id                BIGINT PRIMARY KEY AUTO_INCREMENT,
    product_id        BIGINT   NOT NULL COMMENT '商品ID',
    stock_quantity    INT      NOT NULL COMMENT '库存数量',
    reserved_quantity INT      NOT NULL DEFAULT 0 COMMENT '预留库存数量',
    created_time      datetime NOT NULL DEFAULT current_timestamp,
    updated_time      datetime NOT NULL DEFAULT current_timestamp,
    unique key (product_id)
);

INSERT INTO tb_inventory (product_id, stock_quantity, reserved_quantity)
VALUES (1, 9999, 1000),
       (2, 200, 100),
       (3, 20, 10);

-- 订单

DROP SCHEMA IF EXISTS yecao_mall_order;
CREATE SCHEMA yecao_mall_order DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci;

USE
yecao_mall_order;
DROP TABLE IF EXISTS tb_order;
CREATE TABLE tb_order
(
    id           BIGINT PRIMARY KEY AUTO_INCREMENT,
    order_id     BIGINT         NOT NULL COMMENT '订单ID',
    user_id      BIGINT         NOT NULL COMMENT '用户ID',
    order_status ENUM ('CREATED', 'ENSURE', 'PAID', 'SHIPPED', 'DELIVERED', 'CANCELLED') NOT NULL DEFAULT 'CREATED' COMMENT '订单状态',
    total_amount DECIMAL(10, 2) NOT NULL COMMENT '订单总金额',
    created_time datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_time datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS tb_order_item;
CREATE TABLE tb_order_item
(
    id           BIGINT PRIMARY KEY AUTO_INCREMENT,
--     item_id      BIGINT         NOT NULL COMMENT '订单项ID',
    order_id     BIGINT         NOT NULL COMMENT '订单ID',
    product_id   BIGINT         NOT NULL COMMENT '商品ID',
    quantity     INT            NOT NULL COMMENT '商品数量',
    unit_price   DECIMAL(10, 2) NOT NULL COMMENT '商品单价',
    created_time datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_time datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 优惠券
DROP SCHEMA IF EXISTS yecao_mall_coupon;
CREATE SCHEMA yecao_mall_coupon DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci;

USE
yecao_mall_coupon;
DROP TABLE IF EXISTS tb_coupon;
CREATE TABLE tb_coupon
(
    id              BIGINT PRIMARY KEY AUTO_INCREMENT,
    coupon_id       BIGINT         NOT NULL COMMENT '优惠券ID',
    user_id         BIGINT         NOT NULL COMMENT '用户ID',
    discount_amount DECIMAL(10, 2) NOT NULL COMMENT '优惠金额',
    expiration_date DATE           NOT NULL COMMENT '过期日期',
    is_used         BOOLEAN        NOT NULL DEFAULT FALSE COMMENT '是否已使用',
    created_time    datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_time    datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE key (coupon_id),
    UNIQUE key (user_id, coupon_id)
);

INSERT INTO tb_coupon (coupon_id, user_id, discount_amount, expiration_date, is_used)
VALUES (1, 1, 10.00, '2025-06-30', FALSE),
       (2, 2, 5.00, '2025-05-31', FALSE),
       (3, 1, 20.00, '2025-07-15', FALSE);

-- 支付单
DROP SCHEMA IF EXISTS yecao_mall_payment;
CREATE SCHEMA yecao_mall_payment DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci;

USE
yecao_mall_payment;
DROP TABLE IF EXISTS tb_payment_order;
CREATE TABLE tb_payment_order
(
    id             BIGINT PRIMARY KEY AUTO_INCREMENT,
    payment_id     BIGINT         NOT NULL COMMENT '支付单ID',
    user_id        BIGINT         not null COMMENT '用户ID',
    order_id       BIGINT         NOT NULL COMMENT '订单ID',
    payment_amount DECIMAL(10, 2) NOT NULL COMMENT '支付金额',
    payment_status ENUM ('PENDING', 'SUCCESS', 'FAILED') NOT NULL DEFAULT 'PENDING' COMMENT '支付状态',
    payment_method ENUM ('ALIPAY', 'WECHAT_PAY', 'CREDIT_CARD') COMMENT '支付方式',
    created_time   datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_time   datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE key (payment_id)
);