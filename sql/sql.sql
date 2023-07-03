CREATE DATABASE IF NOT EXISTS `socialmedia`;
USE `socialmedia`;

DROP TABLE IF EXISTS `usuarios`;
DROP TABLE IF EXISTS `seguidores`;

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

CREATE TABLE `seguidores` (
	`usuario_id` INT NOT NULL,
	`seguidor_id` INT NOT NULL,
	PRIMARY KEY (`usuario_id`, `seguidor_id`),
	CONSTRAINT `FK__usuarios` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT `FK__usuarios_2` FOREIGN KEY (`seguidor_id`) REFERENCES `usuarios` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
)

CREATE TABLE publicacoes (
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	titulo VARCHAR(50) NOT NULL,
	conteudo VARCHAR(300) NOT NULL,
	autor_id INT NOT NULL,
	FOREIGN KEY (autor_id) REFERENCES usuarios(id) ON DELETE CASCADE,
	curtidas INT DEFAULT 0,
	criadoEm TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=INNODB;