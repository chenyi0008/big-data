/*
 Navicat Premium Data Transfer

 Source Server         : ymq.xqzyyds.top
 Source Server Type    : MySQL
 Source Server Version : 50740
 Source Host           : ymq.xqzyyds.top:3306
 Source Schema         : bigdata

 Target Server Type    : MySQL
 Target Server Version : 50740
 File Encoding         : 65001

 Date: 01/12/2023 15:10:02
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for country_price_month
-- ----------------------------
DROP TABLE IF EXISTS `country_price_month`;
CREATE TABLE `country_price_month`  (
                                        `id` int(11) NOT NULL AUTO_INCREMENT,
                                        `day` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                        `prediction_price` double NULL DEFAULT NULL,
                                        `avg_price` double NULL DEFAULT NULL,
                                        PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for country_province_price
-- ----------------------------
DROP TABLE IF EXISTS `country_province_price`;
CREATE TABLE `country_province_price`  (
                                           `id` int(11) NOT NULL AUTO_INCREMENT,
                                           `province` varchar(5) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                           `avg_price` double NULL DEFAULT NULL,
                                           PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 29 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for data_source
-- ----------------------------
DROP TABLE IF EXISTS `data_source`;
CREATE TABLE `data_source`  (
                                `id` int(11) NOT NULL AUTO_INCREMENT,
                                `day` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                `province` varchar(5) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                `address` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                `product_name` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                `category` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                `price` double NULL DEFAULT NULL,
                                PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 67500 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for province_price_month
-- ----------------------------
DROP TABLE IF EXISTS `province_price_month`;
CREATE TABLE `province_price_month`  (
                                         `id` int(11) NOT NULL AUTO_INCREMENT,
                                         `province` varchar(5) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                         `day` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                         `prediction_price` double NULL DEFAULT NULL,
                                         `avg_price` double NULL DEFAULT NULL,
                                         PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 892 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
