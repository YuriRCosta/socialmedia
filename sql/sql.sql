CREATE DATABASE IF NOT EXISTS `socialmedia`;
USE `socialmedia`;

DROP TABLE IF EXISTS `usuarios`;

CREATE TABLE `usuarios` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`nome` VARCHAR(50) NOT NULL,
	`nick` VARCHAR(50) NOT NULL,
	`email` VARCHAR(50) NOT NULL,
	`senha` VARCHAR(50) NOT NULL,
	`criadoEm` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
	PRIMARY KEY (`id`),
	UNIQUE INDEX `nick` (`nick`),
	UNIQUE INDEX `email` (`email`),
	UNIQUE INDEX `senha` (`senha`)
)