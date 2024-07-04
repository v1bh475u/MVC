CREATE TABLE IF NOT EXISTS `borrowing_history`(
    `ID` INT AUTO_INCREMENT PRIMARY KEY,
    `BookID` INT,
     FOREIGN KEY (`BookID`) REFERENCES books(BookID),
    `Title` VARCHAR(255) NOT NULL,
    `Username` VARCHAR(255) NOT NULL,
    `Borrowed_date` DATE NOT NULL,
    `Returned_date` DATE
);