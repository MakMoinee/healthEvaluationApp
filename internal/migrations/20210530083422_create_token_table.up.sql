CREATE TABLE `he`.`token`  (
  `TokenID` int NOT NULL AUTO_INCREMENT,
  `TokenString` varchar(255) NOT NULL,
  `UserID` int NOT NULL,
  `DateCreated` datetime NOT NULL,
  PRIMARY KEY (`TokenID`)
);