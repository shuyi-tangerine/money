


docker exec -it mysql env LANG=C.UTF-8 /bin/bash
mysql -uroot -p
CREATE DATABASE shuyi CHARACTER SET utf8mb4;
use shuyi;
drop table if exists finance_detail;
CREATE TABLE `finance_detail` (
   `id` bigint NOT NULL AUTO_INCREMENT,
   `finance_detail_id` bigint NOT NULL DEFAULT 0 COMMENT '资金明细ID',
   `app_id` int NOT NULL DEFAULT 0 COMMENT '应用ID',
   `amount` bigint NOT NULL DEFAULT 0 COMMENT '金钱数额，单位为分',
   `operated_type` tinyint NOT NULL DEFAULT 0 COMMENT '操作类型，1-收入，2-支出',
   `operated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间，秒级时间戳',
   `operated_by` varchar(128) NOT NULL DEFAULT '' COMMENT '操作人',
   `extra` text COMMENT '一些额外的信息',
   `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   `created_by` varchar(128) NOT NULL DEFAULT '' COMMENT '创建人',
   `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
   `updated_by` varchar(128) NOT NULL DEFAULT '' COMMENT '最后更新人',
   PRIMARY KEY (`id`),
   UNIQUE KEY `uniq_fdid` (`finance_detail_id`)
) ENGINE = InnoDB CHARSET = utf8mb4 AUTO_INCREMENT = 2020 COMMENT '资金明细';

insert into finance_detail(`finance_detail_id`, `created_by`) values(1, 'chenshuyi');

