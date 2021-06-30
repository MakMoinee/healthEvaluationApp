/*
 Navicat Premium Data Transfer

 Source Server         : LOCAL_MYSQL
 Source Server Type    : MySQL
 Source Server Version : 80023
 Source Host           : localhost:3306
 Source Schema         : he

 Target Server Type    : MySQL
 Target Server Version : 80023
 File Encoding         : 65001

 Date: 30/06/2021 00:46:42
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tokens
-- ----------------------------
DROP TABLE IF EXISTS `tokens`;
CREATE TABLE `tokens`  (
  `TokenID` int(0) NOT NULL AUTO_INCREMENT,
  `TokenString` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `UserID` int(0) NOT NULL,
  `DateCreated` datetime(0) NOT NULL,
  PRIMARY KEY (`TokenID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 56 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tokens
-- ----------------------------
INSERT INTO `tokens` VALUES (22, '$2a$10$qP01zzdpx5Hu1EEQFtuuBuQNpe20AZQFezS.ii.w1lZY.jOv5zo8G', 1, '2021-06-05 18:49:10');
INSERT INTO `tokens` VALUES (23, '$2a$10$ujVDiozRb0Mw5yhMH0oH5.9dpAJH.RCr/7/NhAeze4bZJOjtMoMmG', 1, '2021-06-05 18:49:10');
INSERT INTO `tokens` VALUES (24, '$2a$10$wkxUyeEdRRG8x1HL9dd7F.Clf3iz5zGPLSfTHhmwAf2Dbt8F/XAYq', 1, '2021-06-05 18:57:41');
INSERT INTO `tokens` VALUES (25, '$2a$10$MA0oRDekhznDf1/aJk.fqO3DAfz9LY3conodtYnIcNLPq1iLby.Ki', 1, '2021-06-05 18:59:06');
INSERT INTO `tokens` VALUES (26, '$2a$10$E7FPdrBItbz5A6OXR1xuCuyIclqilhXDvf3h28vsSWBTX3ddFPtdq', 1, '2021-06-05 19:13:58');
INSERT INTO `tokens` VALUES (27, '$2a$10$ZH4Ir3AoYYlV2wN4PC7fReHS3wXJzX8G3hrS6NOMWgdCuU8TJkvWi', 1, '2021-06-05 21:10:19');
INSERT INTO `tokens` VALUES (28, '$2a$10$h4JjFSFgGZCz2faFQVAebOtZnF7c6qsXClOneRzh2CSNG04yB9xyW', 1, '2021-06-05 21:10:19');
INSERT INTO `tokens` VALUES (29, '$2a$10$PpBf5sS2DBmtLIAf/pZe1OPLAm2Op.lrjpYxjMP/fafSWB77YnnFm', 1, '2021-06-05 21:27:17');
INSERT INTO `tokens` VALUES (30, '$2a$10$io0bZKvA7JsSpTkZGrmxG.CXwYPFkPxiNUnF.Z663By6oAt4918wK', 1, '2021-06-05 22:07:50');
INSERT INTO `tokens` VALUES (31, '$2a$10$hqHDMadBL8P9UjxiyqPC8.9k3gjUusuh6Yca9G2wwgeytXaAiTJTO', 1, '2021-06-05 22:33:10');
INSERT INTO `tokens` VALUES (32, '$2a$10$hEFfGlaLRjWTwm5jDNLPc.VZDAQbnjI9Z9mDbpiKWysU.7F5erTbq', 1, '2021-06-05 23:06:03');
INSERT INTO `tokens` VALUES (33, '$2a$10$ENVlR09tehpRwf2Wi/PuOuCpxnrBtrifPuSVtsXecjsXcPPMkmisO', 1, '2021-06-05 23:06:03');
INSERT INTO `tokens` VALUES (34, '$2a$10$ETBEEkJNSGseGiVIrmUba.3DIikXnsWdGHuphMBy0P0ssM5IgV1UK', 1, '2021-06-05 23:06:12');
INSERT INTO `tokens` VALUES (35, '$2a$10$nAYgUH/4mbPQe.FLMDRhJ.1cg6IopiQvzecmhCZEHY27P4LQkirDC', 1, '2021-06-05 23:26:49');
INSERT INTO `tokens` VALUES (36, '$2a$10$wmLse4EqWtJ7mzbJ2mcx.e9B0HuMoo9PF7KG6etqM2NedKvBklLAC', 1, '2021-06-05 23:51:55');
INSERT INTO `tokens` VALUES (37, '$2a$10$uRMO9CbXpHczyVTLM7My1ePCJu9lkp5gHJLdtLlZuquaBx7fJMcZS', 1, '2021-06-05 23:52:03');
INSERT INTO `tokens` VALUES (38, '$2a$10$/bOzNsFOY/KD7UfAI6Ys5upDWT9P4.pLnCjPA/4dTYtx9JUECeLwi', 1, '2021-06-07 00:44:43');
INSERT INTO `tokens` VALUES (39, '$2a$10$zpeWaMElvfHnwFq3GiskeeE2tVqjdQgij3FvRl0FAPc1SQY9trSqm', 1, '2021-06-07 00:58:49');
INSERT INTO `tokens` VALUES (40, '$2a$10$EmdwDqrxQkdQf5/avO36e.gdvs6ADvHAcvPEdalZUQEGkL2y94tr.', 1, '2021-06-07 00:58:50');
INSERT INTO `tokens` VALUES (41, '$2a$10$Z1uBoAAvfAZR24MrMCQCjedjjcpqTTJ4vUtf/Q1p0dSO3DZiNbDWi', 1, '2021-06-07 01:21:02');
INSERT INTO `tokens` VALUES (42, '$2a$10$Y8Kz5QCt.ZHwzm59s/NYfuCBEI9bdjYnVSCSRBTpDjrdN9iSW8n9C', 1, '2021-06-07 01:53:00');
INSERT INTO `tokens` VALUES (43, '$2a$10$YJu46cv3toYeQxSRUclxY.yyFX79SP3n8TpbEAdzLSCqAHsUintM2', 1, '2021-06-07 02:04:26');
INSERT INTO `tokens` VALUES (44, '$2a$10$MGSxdKQbO2bRlR3NoIlilOINjJdjxEHMwwF2ayZEN0hGZqPTcnsVS', 1, '2021-06-07 02:18:19');
INSERT INTO `tokens` VALUES (45, '$2a$10$tK6m0U80oChbyXkgo7MXGO/LA4FDBzDYI5YkNaHaXLdYaCeiywCOO', 1, '2021-06-07 02:39:51');
INSERT INTO `tokens` VALUES (46, '$2a$10$JQVqPJ9fP8QF9mLLvSKxLefmp24aWukeI1GoA6ssLVxfrRs6czrIa', 1, '2021-06-13 01:11:53');
INSERT INTO `tokens` VALUES (47, '$2a$10$Dk4jdCEgo3ArlgvCRNdw2OPDPAPjB3HzjBw2ILA29gABddY5YlABy', 1, '2021-06-18 10:59:24');
INSERT INTO `tokens` VALUES (48, '$2a$10$kNDCaOBPsCcpDAMJG53uy.AdwJHa1hlAdrgR/K4t2u5VCQXNZBDgO', 1, '2021-06-18 10:59:24');
INSERT INTO `tokens` VALUES (49, '$2a$10$VGJr9qySj20Iew2OdEoEH.r1d4AofuQYaVwNHxLJO34vzVNJRcjw6', 1, '2021-06-18 10:59:50');
INSERT INTO `tokens` VALUES (50, '$2a$10$VGJUdBvG1dpX/582mNPv/.ooef6H37d7.KHnwQZRtjGAzfx8B4Hf6', 1, '2021-06-18 11:19:14');
INSERT INTO `tokens` VALUES (51, '$2a$10$FMwA4L1cQCr7Uao5u2BYcOV5A2PnnUkFvIxpFpBq5z/TuNCsO7siW', 1, '2021-06-18 11:33:13');
INSERT INTO `tokens` VALUES (52, '$2a$10$pAN5LocAEXLTHrMnQmxrjOWADRF3rYmH9PK5E3MEpXMUCOdiHiubC', 1, '2021-06-18 11:50:15');
INSERT INTO `tokens` VALUES (53, '$2a$10$eSwWUSQJdEZe9S8EasFgzecaceKEnyGsv3o4.dhQTO/6kWzfwTNqO', 1, '2021-06-18 12:34:09');
INSERT INTO `tokens` VALUES (54, '$2a$10$SosfnfuxbeMHYL5JrgeogO6Sd0FgBd8jxwUa4rTpUlKbTdxBryfba', 1, '2021-06-18 12:34:13');
INSERT INTO `tokens` VALUES (55, '$2a$10$8IuD.dZfdEJg/oFjx6BKpuOvg2n/NmM0iTRhk06./d3MmDtZOFC2m', 1, '2021-06-18 12:34:17');
INSERT INTO `tokens` VALUES (56, '$2a$10$g5nq.mXeIvCX/ivJERDg0O12a1Y3H6AsOYKP1e4KxNmUeNS/H4cHm', 1, '2021-06-18 12:34:43');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `UserID` int(0) NOT NULL AUTO_INCREMENT,
  `Username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `Password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `UserType` int(0) NOT NULL,
  `LastModifiedDate` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`UserID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'admin@aclcbukidnon.com', '1234', 1, '2021-06-01 18:46:59');
INSERT INTO `users` VALUES (7, 'mak', '1234', 2, '2021-05-30 18:22:40');

SET FOREIGN_KEY_CHECKS = 1;
