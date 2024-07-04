CREATE TABLE IF NOT EXISTS `requests`(
    `ID` INT AUTO_INCREMENT PRIMARY KEY,
    `Username` VARCHAR(255) NOT NULL,
    `BookID` INT,
    FOREIGN KEY(`BookID`) REFERENCES books(BookID),
    `Title` VARCHAR(255) NULL,
    `Request` ENUM('checkout','checkin','adminPrivs') NOT NULL,
    `Status` ENUM('pending','approved','disapproved') NOT NULL,
    `User_status` ENUM('seen','unseen','pending') NOT NULL,
    `Date` Date NOT NULL
);