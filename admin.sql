-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2024-08-21 14:02:47
-- 服务器版本： 5.7.26
-- PHP 版本： 7.3.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `admin`
--

-- --------------------------------------------------------

--
-- 表的结构 `system_account`
--

CREATE TABLE `system_account` (
  `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键',
  `user_name` longtext COLLATE utf8_unicode_ci NOT NULL COMMENT '用户名',
  `nick_name` longtext COLLATE utf8_unicode_ci COMMENT '昵称',
  `avatar` varchar(191) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '头像',
  `password` longtext COLLATE utf8_unicode_ci NOT NULL COMMENT '密码',
  `salt` longtext COLLATE utf8_unicode_ci NOT NULL COMMENT '密码盐',
  `email` longtext COLLATE utf8_unicode_ci COMMENT '邮箱',
  `status` bigint(20) DEFAULT '1' COMMENT '状态',
  `dept_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT '部门',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间'
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `system_account`
--

INSERT INTO `system_account` (`id`, `user_name`, `nick_name`, `avatar`, `password`, `salt`, `email`, `status`, `dept_id`, `deleted_at`, `created_at`, `updated_at`) VALUES
(1, 'admin', 'admin', NULL, 'c55f9fc7e453519664c82e512bd6c37d', 's6w5cqqsky4913g', 'ocink@qq.com', 1, 1, NULL, '2024-08-13 11:21:45', '2024-08-13 11:21:45');

-- --------------------------------------------------------

--
-- 表的结构 `system_account_role`
--

CREATE TABLE `system_account_role` (
  `system_account_id` bigint(20) UNSIGNED NOT NULL COMMENT '主键',
  `system_role_id` bigint(20) UNSIGNED NOT NULL COMMENT '主键'
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `system_account_role`
--

INSERT INTO `system_account_role` (`system_account_id`, `system_role_id`) VALUES
(1, 1),
(1, 2);

-- --------------------------------------------------------

--
-- 表的结构 `system_dept`
--

CREATE TABLE `system_dept` (
  `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键',
  `parent_id` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '上级部门',
  `name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '部门名称',
  `sort` bigint(20) DEFAULT NULL COMMENT '排序',
  `status` bigint(20) DEFAULT NULL COMMENT '状态',
  `description` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '描述',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间'
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `system_dept`
--

INSERT INTO `system_dept` (`id`, `parent_id`, `name`, `sort`, `status`, `description`, `deleted_at`, `created_at`, `updated_at`) VALUES
(1, NULL, 'XXX网络科技有限公司', 1, 1, '公司总部', NULL, '2024-08-13 11:49:08', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `system_dict`
--

CREATE TABLE `system_dict` (
  `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键',
  `name` longtext COLLATE utf8_unicode_ci NOT NULL COMMENT '字典名称',
  `desc` longtext COLLATE utf8_unicode_ci COMMENT '描述',
  `code` longtext COLLATE utf8_unicode_ci COMMENT '字典编码',
  `status` bigint(20) DEFAULT '1' COMMENT '字典状态',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间'
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `system_dict`
--

INSERT INTO `system_dict` (`id`, `name`, `desc`, `code`, `status`, `deleted_at`, `created_at`, `updated_at`) VALUES
(1, '性别', '所有性别', 'sex', 1, NULL, '2024-08-13 16:04:38', '2024-08-13 16:15:19'),
(2, '状态', '系统通用状态', 'status', 1, NULL, '2024-08-14 17:35:23', '2024-08-14 17:35:24');

-- --------------------------------------------------------

--
-- 表的结构 `system_dict_data`
--

CREATE TABLE `system_dict_data` (
  `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键',
  `label` longtext COLLATE utf8_unicode_ci NOT NULL COMMENT '字典名',
  `value` longtext COLLATE utf8_unicode_ci COMMENT '字典值',
  `status` bigint(20) DEFAULT '1' COMMENT '字典状态',
  `dict_id` bigint(20) UNSIGNED NOT NULL COMMENT '所属字典ID',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间'
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `system_dict_data`
--

INSERT INTO `system_dict_data` (`id`, `label`, `value`, `status`, `dict_id`, `deleted_at`, `created_at`, `updated_at`) VALUES
(1, '男', '1', 1, 1, NULL, '2024-08-13 16:05:01', NULL),
(2, '女', '2', 1, 1, NULL, '2024-08-13 16:05:11', NULL),
(3, '正常', '1', 1, 2, NULL, '2024-08-14 17:35:40', NULL),
(4, '禁用', '0', 1, 2, NULL, '2024-08-14 17:35:53', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `system_menu`
--

CREATE TABLE `system_menu` (
  `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键',
  `parent_id` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '父级菜单ID',
  `path` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'URL路径',
  `component` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '前端组件路径',
  `redirect` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '重定向路径',
  `type` bigint(20) DEFAULT NULL COMMENT '菜单类型',
  `permission` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '权限标识',
  `title` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '标题',
  `svg_icon` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'SVG图标路径',
  `icon` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '字体图标名称',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '是否缓存保留',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否隐藏',
  `sort` bigint(20) DEFAULT NULL COMMENT '排序',
  `active_menu` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '菜单激活引用',
  `breadcrumb` tinyint(1) DEFAULT NULL COMMENT '面包屑显示',
  `status` bigint(20) DEFAULT '1' COMMENT '状态',
  `show_in_tabs` tinyint(1) DEFAULT NULL COMMENT '标签页显示',
  `always_show` tinyint(1) DEFAULT NULL COMMENT '是否总显示',
  `affix` tinyint(1) DEFAULT NULL COMMENT '是否固定',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间'
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `system_menu`
--

INSERT INTO `system_menu` (`id`, `parent_id`, `path`, `component`, `redirect`, `type`, `permission`, `title`, `svg_icon`, `icon`, `keep_alive`, `hidden`, `sort`, `active_menu`, `breadcrumb`, `status`, `show_in_tabs`, `always_show`, `affix`, `deleted_at`, `created_at`, `updated_at`) VALUES
(1, NULL, '/analyse', 'Layout', NULL, 1, NULL, '分析页', NULL, NULL, 0, 0, 1, NULL, 1, 1, 1, 0, 0, NULL, '2024-08-13 11:56:57', NULL),
(3, NULL, '/system', 'Layout', NULL, 1, NULL, '系统设置', 'menu-example', 'IconSettings', 0, 0, 2, NULL, 1, 1, 1, 0, 0, NULL, '2024-08-14 15:30:23', '2024-08-17 20:26:33'),
(2, '1', '/analyse/index', 'analyse/index', NULL, 2, NULL, '分析页', 'menu-chart', 'IconDesktop', 0, 0, 1, NULL, 1, 1, 1, 0, 0, NULL, '2024-08-13 20:01:54', '2024-08-17 20:26:26'),
(4, '3', '/system/user', 'system/user/index', NULL, 2, NULL, '用户管理', NULL, 'IconUser', 0, 0, 1, NULL, 1, 1, 1, 0, 0, NULL, '2024-08-14 15:31:51', '2024-08-17 22:47:45'),
(5, '3', '/system/menu', 'system/menu/index', NULL, 2, NULL, '菜单管理', NULL, 'IconList', 0, 0, 1, NULL, 1, 1, 1, 0, 0, NULL, '2024-08-14 15:31:51', '2024-08-17 22:44:20'),
(6, '3', '/system/role', 'system/role/index', NULL, 2, NULL, '角色管理', NULL, 'IconUserGroup', 0, 0, 1, NULL, 1, 1, 1, 0, 0, NULL, '2024-08-14 15:31:51', '2024-08-17 22:46:08'),
(7, '3', '/system/dict', 'system/dict/index', NULL, 2, NULL, '字典管理', NULL, 'IconBookmark', 0, 0, 1, NULL, 1, 1, 1, 0, 0, NULL, '2024-08-14 15:31:51', '2024-08-17 22:47:33'),
(8, '3', '/system/dept', 'system/dept/index', NULL, 2, NULL, '部门管理', NULL, 'IconMindMapping', 0, 0, 1, NULL, 1, 1, 1, 0, 0, NULL, '2024-08-14 15:31:51', '2024-08-17 22:46:23'),
(9, '0', '/error', '', '/error/403', 1, '', '错误页', 'menu-error', 'IconCloseCircle', 0, 0, 3, '', 1, 1, 1, 0, 0, NULL, '2024-08-14 23:10:20', '2024-08-17 20:26:46'),
(10, '9', '/error/403', 'error/403', '', 2, '', '403页', '', 'IconUnorderedList', 0, 0, 1, '', 1, 1, 1, 0, 0, NULL, '2024-08-14 23:12:31', '2024-08-17 22:48:14'),
(11, '9', '/error/404', 'error/404', '', 2, '', '404页', '', 'IconUnorderedList', 0, 0, 2, '', 1, 1, 1, 0, 0, NULL, '2024-08-14 23:17:47', '2024-08-17 22:48:08');

-- --------------------------------------------------------

--
-- 表的结构 `system_role`
--

CREATE TABLE `system_role` (
  `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键',
  `name` longtext COLLATE utf8_unicode_ci NOT NULL COMMENT '角色名称',
  `description` longtext COLLATE utf8_unicode_ci COMMENT '描述',
  `code` longtext COLLATE utf8_unicode_ci COMMENT '角色编码',
  `status` bigint(20) DEFAULT '1' COMMENT '角色状态',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间'
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `system_role`
--

INSERT INTO `system_role` (`id`, `name`, `description`, `code`, `status`, `deleted_at`, `created_at`, `updated_at`) VALUES
(1, '管理员', '系统管理员', 'admin', 1, NULL, '2024-08-13 11:48:31.000', NULL),
(2, '测试', '测试', 'test', 1, NULL, '2024-08-13 15:07:42.000', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `system_role_menu`
--

CREATE TABLE `system_role_menu` (
  `system_role_id` bigint(20) UNSIGNED NOT NULL COMMENT '主键',
  `system_menu_id` bigint(20) UNSIGNED NOT NULL COMMENT '主键'
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `system_role_menu`
--

INSERT INTO `system_role_menu` (`system_role_id`, `system_menu_id`) VALUES
(1, 1),
(1, 2),
(1, 3),
(1, 4),
(1, 5),
(1, 6),
(1, 7),
(1, 8),
(1, 9),
(1, 10),
(1, 11),
(2, 1);

--
-- 转储表的索引
--

--
-- 表的索引 `system_account`
--
ALTER TABLE `system_account`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_system_account_deleted_at` (`deleted_at`);

--
-- 表的索引 `system_account_role`
--
ALTER TABLE `system_account_role`
  ADD PRIMARY KEY (`system_account_id`,`system_role_id`);

--
-- 表的索引 `system_dept`
--
ALTER TABLE `system_dept`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_system_dept_deleted_at` (`deleted_at`);

--
-- 表的索引 `system_dict`
--
ALTER TABLE `system_dict`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_system_dict_deleted_at` (`deleted_at`);

--
-- 表的索引 `system_dict_data`
--
ALTER TABLE `system_dict_data`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_system_dict_data_deleted_at` (`deleted_at`);

--
-- 表的索引 `system_menu`
--
ALTER TABLE `system_menu`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_system_menu_deleted_at` (`deleted_at`);

--
-- 表的索引 `system_role`
--
ALTER TABLE `system_role`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_system_role_deleted_at` (`deleted_at`);

--
-- 表的索引 `system_role_menu`
--
ALTER TABLE `system_role_menu`
  ADD PRIMARY KEY (`system_role_id`,`system_menu_id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `system_account`
--
ALTER TABLE `system_account`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `system_dept`
--
ALTER TABLE `system_dept`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `system_dict`
--
ALTER TABLE `system_dict`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `system_dict_data`
--
ALTER TABLE `system_dict_data`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=5;

--
-- 使用表AUTO_INCREMENT `system_menu`
--
ALTER TABLE `system_menu`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=12;

--
-- 使用表AUTO_INCREMENT `system_role`
--
ALTER TABLE `system_role`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
