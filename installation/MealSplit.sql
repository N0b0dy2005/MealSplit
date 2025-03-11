/*
 Navicat Premium Data Transfer

 Source Server         : test
 Source Server Type    : MySQL
 Source Server Version : 80300 (8.3.0)
 Source Host           : localhost:3306
 Source Schema         : MealSplit

 Target Server Type    : MySQL
 Target Server Version : 80300 (8.3.0)
 File Encoding         : 65001

 Date: 11/03/2025 20:18:38
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for activities
-- ----------------------------
DROP TABLE IF EXISTS `activities`;
CREATE TABLE `activities` (
  `id` int NOT NULL AUTO_INCREMENT,
  `type` enum('meal','payment','user') COLLATE utf8mb4_bs_0900_as_cs NOT NULL,
  `description` text COLLATE utf8mb4_bs_0900_as_cs,
  `from_user_id` int DEFAULT NULL,
  `to_user_id` int DEFAULT NULL,
  `user_id` int DEFAULT NULL,
  `amount` decimal(10,2) DEFAULT NULL,
  `date` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `from_user_id` (`from_user_id`),
  KEY `to_user_id` (`to_user_id`),
  KEY `payer_id` (`user_id`),
  CONSTRAINT `activities_ibfk_1` FOREIGN KEY (`from_user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL,
  CONSTRAINT `activities_ibfk_2` FOREIGN KEY (`to_user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL,
  CONSTRAINT `activities_ibfk_3` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bs_0900_as_cs;

-- ----------------------------
-- Records of activities
-- ----------------------------
BEGIN;
INSERT INTO `activities` (`id`, `type`, `description`, `from_user_id`, `to_user_id`, `user_id`, `amount`, `date`) VALUES (6, 'user', 'User created', NULL, NULL, NULL, 0.00, '2025-03-11 17:04:05');
INSERT INTO `activities` (`id`, `type`, `description`, `from_user_id`, `to_user_id`, `user_id`, `amount`, `date`) VALUES (7, 'meal', 'es war sehr lecker', NULL, NULL, 8, 23.45, '2025-03-11 00:00:00');
INSERT INTO `activities` (`id`, `type`, `description`, `from_user_id`, `to_user_id`, `user_id`, `amount`, `date`) VALUES (8, 'payment', 'Thai rotes curry ', 11, 10, 11, 12.50, '2025-03-11 17:11:51');
COMMIT;

-- ----------------------------
-- Table structure for debts
-- ----------------------------
DROP TABLE IF EXISTS `debts`;
CREATE TABLE `debts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `from_user_id` int NOT NULL,
  `to_user_id` int NOT NULL,
  `amount` decimal(10,2) NOT NULL,
  `meals_count` int DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_debt` (`from_user_id`,`to_user_id`),
  KEY `to_user_id` (`to_user_id`),
  CONSTRAINT `debts_ibfk_1` FOREIGN KEY (`from_user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `debts_ibfk_2` FOREIGN KEY (`to_user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bs_0900_as_cs;

-- ----------------------------
-- Records of debts
-- ----------------------------
BEGIN;
INSERT INTO `debts` (`id`, `from_user_id`, `to_user_id`, `amount`, `meals_count`, `created_at`, `updated_at`) VALUES (27, 9, 8, 7.89, 0, '2025-03-10 19:40:24', '2025-03-10 20:08:09');
INSERT INTO `debts` (`id`, `from_user_id`, `to_user_id`, `amount`, `meals_count`, `created_at`, `updated_at`) VALUES (29, 13, 10, 3.09, 1, '2025-03-11 12:00:42', '2025-03-11 12:00:42');
INSERT INTO `debts` (`id`, `from_user_id`, `to_user_id`, `amount`, `meals_count`, `created_at`, `updated_at`) VALUES (30, 12, 10, 3.09, 1, '2025-03-11 12:00:42', '2025-03-11 12:00:42');
INSERT INTO `debts` (`id`, `from_user_id`, `to_user_id`, `amount`, `meals_count`, `created_at`, `updated_at`) VALUES (31, 10, 8, 4.69, 1, '2025-03-11 17:11:07', '2025-03-11 17:11:07');
INSERT INTO `debts` (`id`, `from_user_id`, `to_user_id`, `amount`, `meals_count`, `created_at`, `updated_at`) VALUES (32, 11, 8, 4.69, 1, '2025-03-11 17:11:07', '2025-03-11 17:11:07');
INSERT INTO `debts` (`id`, `from_user_id`, `to_user_id`, `amount`, `meals_count`, `created_at`, `updated_at`) VALUES (33, 12, 8, 4.69, 1, '2025-03-11 17:11:07', '2025-03-11 17:11:07');
INSERT INTO `debts` (`id`, `from_user_id`, `to_user_id`, `amount`, `meals_count`, `created_at`, `updated_at`) VALUES (34, 13, 8, 4.69, 1, '2025-03-11 17:11:07', '2025-03-11 17:11:07');
INSERT INTO `debts` (`id`, `from_user_id`, `to_user_id`, `amount`, `meals_count`, `created_at`, `updated_at`) VALUES (35, 10, 11, 9.41, 0, '2025-03-11 17:11:51', '2025-03-11 17:11:51');
COMMIT;

-- ----------------------------
-- Table structure for meals
-- ----------------------------
DROP TABLE IF EXISTS `meals`;
CREATE TABLE `meals` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_bs_0900_as_cs NOT NULL,
  `date` date NOT NULL,
  `total_amount` decimal(10,2) NOT NULL,
  `user_id` int NOT NULL,
  `description` text COLLATE utf8mb4_bs_0900_as_cs,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `user_ids` varchar(255) COLLATE utf8mb4_bs_0900_as_cs DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `payer_id` (`user_id`),
  CONSTRAINT `meals_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bs_0900_as_cs;

-- ----------------------------
-- Records of meals
-- ----------------------------
BEGIN;
INSERT INTO `meals` (`id`, `name`, `date`, `total_amount`, `user_id`, `description`, `created_at`, `updated_at`, `user_ids`) VALUES (27, 'Bolognese', '2025-03-11', 12.34, 10, 'lecker', '2025-03-11 12:00:42', '2025-03-11 12:00:42', '11,13,12,10');
INSERT INTO `meals` (`id`, `name`, `date`, `total_amount`, `user_id`, `description`, `created_at`, `updated_at`, `user_ids`) VALUES (28, 'Geschnezeltes', '2025-03-11', 23.45, 8, 'es war sehr lecker', '2025-03-11 17:11:07', '2025-03-11 17:11:07', '10,11,12,13,8');
COMMIT;

-- ----------------------------
-- Table structure for payment
-- ----------------------------
DROP TABLE IF EXISTS `payment`;
CREATE TABLE `payment` (
  `id` int NOT NULL AUTO_INCREMENT,
  `from_user_id` int DEFAULT NULL,
  `to_user_id` int DEFAULT NULL,
  `amount` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bs_0900_as_cs DEFAULT NULL,
  `description` varchar(255) COLLATE utf8mb4_bs_0900_as_cs DEFAULT NULL,
  `date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bs_0900_as_cs;

-- ----------------------------
-- Records of payment
-- ----------------------------
BEGIN;
INSERT INTO `payment` (`id`, `from_user_id`, `to_user_id`, `amount`, `description`, `date`) VALUES (9, 8, 9, '12.34', 'test', '2025-03-10 19:40:24');
INSERT INTO `payment` (`id`, `from_user_id`, `to_user_id`, `amount`, `description`, `date`) VALUES (10, 9, 8, '4.45', 'test', '2025-03-10 20:08:09');
INSERT INTO `payment` (`id`, `from_user_id`, `to_user_id`, `amount`, `description`, `date`) VALUES (11, 11, 10, '12.5', 'Thai rotes curry ', '2025-03-11 17:11:51');
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_bs_0900_as_cs NOT NULL,
  `email` varchar(100) COLLATE utf8mb4_bs_0900_as_cs NOT NULL,
  `debts` decimal(10,2) DEFAULT '0.00',
  `credits` decimal(10,2) DEFAULT '0.00',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bs_0900_as_cs DEFAULT NULL,
  `phone_number` varchar(255) COLLATE utf8mb4_bs_0900_as_cs DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bs_0900_as_cs;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` (`id`, `name`, `email`, `debts`, `credits`, `created_at`, `updated_at`, `password`, `phone_number`) VALUES (8, 'Felix Eichhorn', 'felix@r2e2.de', 0.00, 26.65, '2025-03-10 16:21:54', '2025-03-11 17:11:07', '9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08', NULL);
INSERT INTO `users` (`id`, `name`, `email`, `debts`, `credits`, `created_at`, `updated_at`, `password`, `phone_number`) VALUES (9, 'Marius Eichhorn', 'marius@eichhorn-frankfurt.de', 7.89, 0.00, '2025-03-10 16:22:05', '2025-03-10 20:08:09', NULL, NULL);
INSERT INTO `users` (`id`, `name`, `email`, `debts`, `credits`, `created_at`, `updated_at`, `password`, `phone_number`) VALUES (10, 'Robert Eichhorn', 'robert@r2e2.de', 14.10, 6.18, '2025-03-10 16:22:18', '2025-03-11 17:11:51', NULL, NULL);
INSERT INTO `users` (`id`, `name`, `email`, `debts`, `credits`, `created_at`, `updated_at`, `password`, `phone_number`) VALUES (11, 'Michael Hölzl', 'Michael@r2e2.de', 4.69, 9.41, '2025-03-10 16:23:56', '2025-03-11 17:11:51', NULL, NULL);
INSERT INTO `users` (`id`, `name`, `email`, `debts`, `credits`, `created_at`, `updated_at`, `password`, `phone_number`) VALUES (12, 'Lena Glebe', 'Lena@eichhorn-frankfurt.de', 7.78, 0.00, '2025-03-10 16:24:27', '2025-03-11 17:11:07', NULL, NULL);
INSERT INTO `users` (`id`, `name`, `email`, `debts`, `credits`, `created_at`, `updated_at`, `password`, `phone_number`) VALUES (13, 'Marlene Röhling', 'Marlene@eichhorn-frankfurt.de', 7.78, 0.00, '2025-03-10 16:26:23', '2025-03-11 17:11:07', NULL, NULL);
INSERT INTO `users` (`id`, `name`, `email`, `debts`, `credits`, `created_at`, `updated_at`, `password`, `phone_number`) VALUES (14, 'Sefa Özgel', 'sefa@r2e2.de', 0.00, 0.00, '2025-03-10 16:26:39', '2025-03-10 16:26:39', NULL, NULL);
INSERT INTO `users` (`id`, `name`, `email`, `debts`, `credits`, `created_at`, `updated_at`, `password`, `phone_number`) VALUES (15, 'test user', 'user@example.de', 0.00, 0.00, '2025-03-11 17:04:05', '2025-03-11 17:04:05', NULL, NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
