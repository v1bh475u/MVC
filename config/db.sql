CREATE DATABASE IF NOT EXISTS `lib_db`;
USE `lib_db`;

CREATE TABLE `users` (
    `ID` INT AUTO_INCREMENT PRIMARY KEY,
    `Username` VARCHAR(255) NOT NULL,
    `Password` VARCHAR(255) NOT NULL,
    `Role` ENUM('user','admin') NOT NULL
);

CREATE TABLE `books`(
    `BookID` INT AUTO_INCREMENT PRIMARY KEY,
    `Title` VARCHAR(255) NOT NULL,
    `Author` VARCHAR(255) NOT NULL,
    `Genre` VARCHAR(255) NOT NULL,
    `Quantity` INT
);

CREATE TABLE `requests`(
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



CREATE TABLE `borrowing_history`(
    `ID` INT AUTO_INCREMENT PRIMARY KEY,
    `BookID` INT,
     FOREIGN KEY (`BookID`) REFERENCES books(BookID),
    `Title` VARCHAR(255) NOT NULL,
    `Username` VARCHAR(255) NOT NULL,
    `Borrowed_date` DATE NOT NULL,
    `Returned_date` DATE
);
