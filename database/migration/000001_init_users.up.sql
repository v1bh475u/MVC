CREATE TABLE  IF NOT EXISTS `users`(
    `ID` INT AUTO_INCREMENT PRIMARY KEY,
    `Username` VARCHAR(255) NOT NULL,
    `Password` VARCHAR(255) NOT NULL,
    `Role` ENUM('user','admin') NOT NULL
);