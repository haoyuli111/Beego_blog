/*
Navicat MySQL Data Transfer

Source Server         : liyuhao
Source Server Version : 50559
Source Host           : localhost:3306
Source Database       : db_lccblog

Target Server Type    : MYSQL
Target Server Version : 50559
File Encoding         : 65001

Date: 2020-01-11 22:30:01
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for notes
-- ----------------------------
DROP TABLE IF EXISTS `notes`;
CREATE TABLE `notes` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `n_key` varchar(255) COLLATE utf8_bin NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `title` varchar(200) COLLATE utf8_bin DEFAULT NULL,
  `summary` varchar(800) COLLATE utf8_bin DEFAULT NULL,
  `content` text COLLATE utf8_bin,
  `visit` int(11) DEFAULT '0',
  `praise` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_notes_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of notes
-- ----------------------------
INSERT INTO `notes` VALUES ('1', '2019-05-28 19:52:08', '2019-05-28 19:52:08', null, '90ce049f-7182-4b14-8491-d8df158a51c5', '3', 'lcc', '', 0x6C6363, '0', '0');
INSERT INTO `notes` VALUES ('2', '2019-05-28 19:54:31', '2019-05-29 16:25:32', null, '7ebe7959-d61b-4b2c-ac5d-bade8a0e1d64', '3', 'lcc1', '', 0x6C636331, '0', '0');
INSERT INTO `notes` VALUES ('3', '2019-05-29 12:32:26', '2019-05-29 12:32:26', null, 'dcf0cc36-c9ce-4f67-9dc0-2307158c24e0', '3', 'lcc', '', 0x6C6363, '0', '0');
INSERT INTO `notes` VALUES ('4', '2019-05-29 16:20:42', '2019-05-29 16:20:42', null, '23227350-30eb-4eef-95c9-a2ce503ee921', '0', 'lcc', '', 0x6C6363, '0', '0');
INSERT INTO `notes` VALUES ('5', '2019-05-29 16:22:14', '2019-05-29 16:22:14', null, '644c5669-413f-4fe8-944a-18bb6d42b431', '0', 'lcc', '', 0x6C6363, '0', '0');
INSERT INTO `notes` VALUES ('6', '2019-05-29 16:23:25', '2019-05-29 16:23:25', null, '18968507-dc53-4ce7-949e-4e5917bea3d8', '0', 'lcc', '', 0x6C6363, '0', '0');
INSERT INTO `notes` VALUES ('7', '2019-05-29 16:24:37', '2019-05-29 16:24:37', null, '57b84d8d-fcf2-4e10-b2f1-9d2317f2597d', '3', 'lcc', '', 0x6C63637178, '0', '0');
INSERT INTO `notes` VALUES ('8', '2019-05-29 17:52:07', '2019-05-29 17:52:07', null, '83c6f0e4-d524-4d14-8f6e-0eddd14d8667', '3', 'lcc', '', 0x6C6363, '0', '0');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `email` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `phone` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `password` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `avatar` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `role` int(11) DEFAULT NULL,
  `user_name` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  `signature` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_users_name` (`name`),
  UNIQUE KEY `uix_users_email` (`email`),
  UNIQUE KEY `uix_users_phone` (`phone`),
  UNIQUE KEY `uix_users_user_name` (`user_name`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES ('1', '2019-05-19 18:22:44', '2019-05-19 18:22:44', null, 'admin', 'admin', 'admin', 'admin', '/static/images/admin.png', '0', null, null, null);
