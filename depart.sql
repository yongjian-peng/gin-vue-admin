CREATE TABLE `sys_depart` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `type` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '类型（1企业，2审核方）',
  `company_name` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '企业名称',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态 1.正常',
  `company_email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '企业、审核方邮箱',
  `create_time` bigint DEFAULT NULL COMMENT '创建时间',
  `update_time` bigint DEFAULT NULL COMMENT '更新时间',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=100000188 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='企业方/审核方的表';

INSERT INTO `gwa`.`sys_depart`(`id`, `type`, `company_name`, `status`, `company_email`, `create_time`, `update_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (100000002, '2', 'test_company', 1, 'yong@blueoceanpay.com', 1612182762, 1612182762, NULL, '2022-04-15 16:53:51.958', NULL);
INSERT INTO `gwa`.`sys_depart`(`id`, `type`, `company_name`, `status`, `company_email`, `create_time`, `update_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (100000003, '2', 'One Tech Solutions Limited', 1, 'blue@blueoceanpay.com', 1612182792, 1612182792, NULL, '2022-04-15 16:50:56.892', '2022-04-15 16:51:06.809');
INSERT INTO `gwa`.`sys_depart`(`id`, `type`, `company_name`, `status`, `company_email`, `create_time`, `update_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (100000004, '1', 'BlueOcean Tech Limited', 1, '', 1612182813, 1612182813, NULL, NULL, NULL);
INSERT INTO `gwa`.`sys_depart`(`id`, `type`, `company_name`, `status`, `company_email`, `create_time`, `update_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (100000011, '2', 'BlueOcean Tech Limited', 1, '', 1612347397, 1612347773, NULL, NULL, NULL);
INSERT INTO `gwa`.`sys_depart`(`id`, `type`, `company_name`, `status`, `company_email`, `create_time`, `update_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (100000024, '2', 'CyberPay Limited', 1, '', 1612350954, 1612350954, NULL, NULL, NULL);
INSERT INTO `gwa`.`sys_depart`(`id`, `type`, `company_name`, `status`, `company_email`, `create_time`, `update_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (100000025, '1', 'OTS - One Tech Solutions Limited', 1, '', 1612350966, 1612350966, NULL, NULL, NULL);
INSERT INTO `gwa`.`sys_depart`(`id`, `type`, `company_name`, `status`, `company_email`, `create_time`, `update_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (100000030, '2', 'Guang Rong Finance Company Limited', 1, '', 1612488222, 1612488222, NULL, NULL, NULL);
