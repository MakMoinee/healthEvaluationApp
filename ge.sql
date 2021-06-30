-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               8.0.17 - MySQL Community Server - GPL
-- Server OS:                    Win64
-- HeidiSQL Version:             10.3.0.5771
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- Dumping structure for table he.assessment
CREATE TABLE IF NOT EXISTS `assessment` (
  `AssessmentID` int(11) NOT NULL AUTO_INCREMENT,
  `Sequence` int(11) NOT NULL,
  `Question` varchar(255) NOT NULL,
  `answerA` varchar(255) NOT NULL,
  `answerB` varchar(255) NOT NULL,
  `answerC` varchar(255) NOT NULL,
  `answerD` varchar(255) NOT NULL,
  `answerAPoints` int(11) NOT NULL,
  `answerBPoints` int(11) NOT NULL,
  `answerCPoints` int(11) NOT NULL,
  `answerDPoints` int(11) NOT NULL,
  `Status` int(11) NOT NULL,
  `Category` int(11) NOT NULL,
  PRIMARY KEY (`AssessmentID`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table he.assessment: ~0 rows (approximately)
/*!40000 ALTER TABLE `assessment` DISABLE KEYS */;
INSERT INTO `assessment` (`AssessmentID`, `Sequence`, `Question`, `answerA`, `answerB`, `answerC`, `answerD`, `answerAPoints`, `answerBPoints`, `answerCPoints`, `answerDPoints`, `Status`, `Category`) VALUES
	(1, 1, ' Feeling nervous, anxious or on edge', '0', '1', '2', '3', 0, 1, 2, 3, 1, 1),
	(2, 2, 'Not being able to stop or control worrying', '0', '1', '2', '3', 0, 1, 2, 3, 1, 1),
	(3, 3, 'Worrying too much about different things', '0', '1', '2', '3', 0, 1, 2, 3, 1, 1),
	(4, 4, 'Trouble relaxing', '0', '1', '2', '3', 0, 1, 2, 3, 1, 1);
/*!40000 ALTER TABLE `assessment` ENABLE KEYS */;

-- Dumping structure for table he.assessment_details
CREATE TABLE IF NOT EXISTS `assessment_details` (
  `AssessmentDetailID` int(11) NOT NULL AUTO_INCREMENT,
  `AssessmentID` int(11) NOT NULL,
  `UserID` int(11) NOT NULL,
  `Answer` varchar(255) NOT NULL,
  `DateAnswered` varchar(255) NOT NULL,
  PRIMARY KEY (`AssessmentDetailID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table he.assessment_details: ~0 rows (approximately)
/*!40000 ALTER TABLE `assessment_details` DISABLE KEYS */;
/*!40000 ALTER TABLE `assessment_details` ENABLE KEYS */;

-- Dumping structure for table he.habit
CREATE TABLE IF NOT EXISTS `habit` (
  `HabitID` int(11) NOT NULL AUTO_INCREMENT,
  `UserID` int(11) NOT NULL,
  `HabitName` varchar(255) NOT NULL,
  `HabitDate` varchar(255) NOT NULL,
  `HabitTime` varchar(45) NOT NULL,
  `HabitStatus` int(11) NOT NULL,
  PRIMARY KEY (`HabitID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table he.habit: ~0 rows (approximately)
/*!40000 ALTER TABLE `habit` DISABLE KEYS */;
/*!40000 ALTER TABLE `habit` ENABLE KEYS */;

-- Dumping structure for table he.mood
CREATE TABLE IF NOT EXISTS `mood` (
  `MoodID` int(11) NOT NULL AUTO_INCREMENT,
  `UserID` int(11) NOT NULL,
  `Mood` varchar(255) NOT NULL,
  `MoodDescription` varchar(255) NOT NULL,
  `created_at` date NOT NULL,
  `updated_at` date DEFAULT NULL,
  PRIMARY KEY (`MoodID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table he.mood: ~0 rows (approximately)
/*!40000 ALTER TABLE `mood` DISABLE KEYS */;
/*!40000 ALTER TABLE `mood` ENABLE KEYS */;

-- Dumping structure for table he.todo
CREATE TABLE IF NOT EXISTS `todo` (
  `TodoID` int(11) NOT NULL AUTO_INCREMENT,
  `UserID` int(11) NOT NULL,
  `Note` varchar(255) NOT NULL,
  `StartDate` datetime NOT NULL,
  `EndDate` datetime NOT NULL,
  PRIMARY KEY (`TodoID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table he.todo: ~0 rows (approximately)
/*!40000 ALTER TABLE `todo` DISABLE KEYS */;
/*!40000 ALTER TABLE `todo` ENABLE KEYS */;

-- Dumping structure for table he.tokens
CREATE TABLE IF NOT EXISTS `tokens` (
  `TokenID` int(11) NOT NULL AUTO_INCREMENT,
  `TokenString` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `UserID` int(11) NOT NULL,
  `DateCreated` datetime NOT NULL,
  PRIMARY KEY (`TokenID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- Dumping data for table he.tokens: ~35 rows (approximately)
/*!40000 ALTER TABLE `tokens` DISABLE KEYS */;
INSERT INTO `tokens` (`TokenID`, `TokenString`, `UserID`, `DateCreated`) VALUES
	(1, '$2a$10$SFvp8I6JNupkHxOsS/GesuZERLO4CXohx6d84culSDi4eSyqNcDXq', 1, '2021-06-30 02:08:46'),
	(2, '$2a$10$AQq4VRERDOFJOd5BAc74Weq.OjMB3LBpImmZg4uuTKMb0/nDwOl2C', 1, '2021-06-30 03:12:01'),
	(3, '$2a$10$Mu/tqYVchw3viVTVOS10l.08Kju63teYPpBOnBL5Pj6tYe.VqbXF2', 1, '2021-06-30 03:38:26'),
	(4, '$2a$10$Is3DuW4niB1uXn1sAXew7.ornpz/BQXVfjDKegfBTYCS8K9FeAiTe', 1, '2021-06-30 04:02:06'),
	(5, '$2a$10$Cro7nSRU2/gS2I9qT34QlO4qZLvDI4C5H9KSGM/ati0X83KYoVSWe', 1, '2021-06-30 04:17:03'),
	(6, '$2a$10$17DktR0JAvrjzeLdXBz2WexuBwZSsPze4T7wG4opX0/bV3INYD2Uq', 1, '2021-06-30 05:39:03'),
	(7, '$2a$10$1O2OAIMLByMNiM1/xjF17uEKAfeGhXE5GdYGORa.dvhH3SRGOEqEi', 1, '2021-06-30 05:39:21'),
	(8, '$2a$10$0V4El9l6OFJ78rrkLUeRR.GjLSosH8waCpuJhNFuKVreHP2EZdcTi', 1, '2021-06-30 06:37:10'),
	(9, '$2a$10$KuVXWjIjwvMPlapaLgdMI.z64XoHmG4Zx5DsjFvhiuWfQv.26CeiO', 1, '2021-06-30 06:37:20'),
	(10, '$2a$10$DJJ6gxtNwhkWjWfqUf7oBe5TNffXVTn7slgcYbZYhUSCc11qfzHfe', 1, '2021-06-30 07:51:47'),
	(11, '$2a$10$GVtDoivf/RovfTxaSdggIO3.mdIhUBoJnIqODAwbKYAlpSLntbKk2', 1, '2021-06-30 07:52:07'),
	(12, '$2a$10$fLRm/QSaKYi6S.H04EbNtOgejpFZz5QaYS2nvkOQ22UTYwaNHixR6', 1, '2021-06-30 07:52:19'),
	(13, '$2a$10$9HiVuZIv68IBu61WmnISbesDLMw1Y5D6ypsUJfN28fAA/.wGp0Bm6', 1, '2021-06-30 07:54:55'),
	(14, '$2a$10$NzQBcKAu2wo4Uhmlx1lFR.5u2/E.D7NtiCzdYUBrDcjK0eP.BgutS', 1, '2021-06-30 07:56:20'),
	(15, '$2a$10$gDaTXEGodR5g3cP3VViakuCcoMTs1wLJjqphbJnCxioOcruoCRmvW', 1, '2021-06-30 08:07:12'),
	(16, '$2a$10$coIqHkpnFOIi0.5/SCSAju2rGkKQcSa96oU1B9uygE9T7SyaFvmjy', 1, '2021-06-30 08:23:48'),
	(17, '$2a$10$dyBGC2ktRrChqC3KtjxSrOt0m5iKYq6UxvaIGeG5UHjY3c/YP6Nb6', 1, '2021-06-30 08:26:41'),
	(18, '$2a$10$ujqY6PyUzfr0CJ8O9WeW1e1yYwHrEuh8GHiKD7MDude1GsSe8.jSS', 1, '2021-06-30 08:53:30'),
	(19, '$2a$10$7vNU58m4z0zWyZ2Rp62NMupKUXai4N2KsGpXOZqPAMYbQV3XPadaa', 1, '2021-06-30 08:53:39'),
	(20, '$2a$10$p6lzSfTaiFR2KYYWucpJSOMqhdsibG6c/R.AWzdbbig53t9SK/7nS', 1, '2021-06-30 09:20:27'),
	(21, '$2a$10$U8Vu4D183IotnS8qLUu7TurbGING8Kw66nkx0odOVQYoCF.7g88ze', 1, '2021-06-30 10:49:12'),
	(22, '$2a$10$4MKW2QDuJck0A4pxM/vtc.MIc96jjA9Po9/lQBLDtYEX/8LIBO3.a', 1, '2021-06-30 10:49:43');
/*!40000 ALTER TABLE `tokens` ENABLE KEYS */;

-- Dumping structure for table he.users
CREATE TABLE IF NOT EXISTS `users` (
  `UserID` int(11) NOT NULL AUTO_INCREMENT,
  `Username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `Password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `UserType` int(11) NOT NULL,
  `LastModifiedDate` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`UserID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- Dumping data for table he.users: ~2 rows (approximately)
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` (`UserID`, `Username`, `Password`, `UserType`, `LastModifiedDate`) VALUES
	(1, 'admin@aclcbukidnon.com', '1234', 1, '2021-06-01 18:46:59'),
	(7, 'mak', '1234', 2, '2021-05-30 18:22:40');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
