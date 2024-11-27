/*
 Navicat Premium Data Transfer

 Source Server         : local-mysql
 Source Server Type    : MySQL
 Source Server Version : 80039 (8.0.39)
 Source Host           : localhost:3306
 Source Schema         : chgame_admin

 Target Server Type    : MySQL
 Target Server Version : 80039 (8.0.39)
 File Encoding         : 65001

 Date: 27/11/2024 10:12:01
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_admin_info
-- ----------------------------
DROP TABLE IF EXISTS `t_admin_info`;
CREATE TABLE `t_admin_info` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `account` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '账号',
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '账号名字',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
  `password_sign` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '密码盐值',
  `role_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  `status` int unsigned NOT NULL DEFAULT '1' COMMENT '1=正常；2=异常',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_account` (`account`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of t_admin_info
-- ----------------------------
BEGIN;
INSERT INTO `t_admin_info` (`id`, `account`, `name`, `password`, `password_sign`, `role_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'admin', '', 'fc225818e1f2e2204e850a4777ec40ecbfccba29fa9990435acc74ee5a34f5a7', '453sr5rt', 1, 1, NULL, '2024-11-26 10:00:40', NULL);
COMMIT;

-- ----------------------------
-- Table structure for t_admin_login_token
-- ----------------------------
DROP TABLE IF EXISTS `t_admin_login_token`;
CREATE TABLE `t_admin_login_token` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `admin_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '管理员ID',
  `token_sign` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '登录token签名',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of t_admin_login_token
-- ----------------------------
BEGIN;
INSERT INTO `t_admin_login_token` (`id`, `admin_id`, `token_sign`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 1, 'myvJZNhv', '2024-11-26 02:10:09', '2024-11-26 10:13:05', NULL);
COMMIT;

-- ----------------------------
-- Table structure for t_menu
-- ----------------------------
DROP TABLE IF EXISTS `t_menu`;
CREATE TABLE `t_menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `menu_pid` bigint unsigned NOT NULL DEFAULT '0',
  `menu_name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `menu_type` int unsigned NOT NULL DEFAULT '0' COMMENT '菜单类型 1=页面，2=按钮',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '请求路径，接口路径',
  `status` int unsigned NOT NULL DEFAULT '0' COMMENT '1=启用；2=禁用',
  `sort` int NOT NULL DEFAULT '1' COMMENT '排序，越小越前',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of t_menu
-- ----------------------------
BEGIN;
INSERT INTO `t_menu` (`id`, `menu_pid`, `menu_name`, `menu_type`, `path`, `status`, `sort`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 0, '1111', 1, '', 1, 1, NULL, '2024-11-25 16:56:08', NULL);
INSERT INTO `t_menu` (`id`, `menu_pid`, `menu_name`, `menu_type`, `path`, `status`, `sort`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 0, '222', 1, '', 1, 1, NULL, '2024-11-25 16:56:11', NULL);
INSERT INTO `t_menu` (`id`, `menu_pid`, `menu_name`, `menu_type`, `path`, `status`, `sort`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 1, '1111-333', 1, '', 1, 1, NULL, '2024-11-25 16:56:48', NULL);
INSERT INTO `t_menu` (`id`, `menu_pid`, `menu_name`, `menu_type`, `path`, `status`, `sort`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 2, '222-444', 1, '', 1, 1, NULL, '2024-11-25 16:57:07', NULL);
COMMIT;

-- ----------------------------
-- Table structure for t_permission_menu
-- ----------------------------
DROP TABLE IF EXISTS `t_permission_menu`;
CREATE TABLE `t_permission_menu` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `permission_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `menu_ids` json DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of t_permission_menu
-- ----------------------------
BEGIN;
INSERT INTO `t_permission_menu` (`id`, `permission_name`, `menu_ids`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, '组1', '[1, 2]', NULL, NULL, NULL);
INSERT INTO `t_permission_menu` (`id`, `permission_name`, `menu_ids`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, '组2', '[3, 4]', NULL, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for t_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `t_role_permission`;
CREATE TABLE `t_role_permission` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `role_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `permission_ids` json NOT NULL COMMENT '权限ID',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of t_role_permission
-- ----------------------------
BEGIN;
INSERT INTO `t_role_permission` (`id`, `role_name`, `permission_ids`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, '管理员', '[1, 2]', NULL, '2024-11-26 10:23:22', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
