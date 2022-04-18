CREATE DATABASE IF NOT EXISTS `bookstore_server` DEFAULT CHARACTER SET utf8;

USE `bookstore_server`;

DROP TABLE IF EXISTS `tb_users`;

CREATE TABLE `tb_users` (
                            `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                            `username` varchar(255) NOT NULL,
                            `password` varchar(255) NOT NULL,
                            `Role`     varchar(255) NOT NULL,
                            `createdAt` timestamp NULL DEFAULT NULL,
                            `updatedAt` timestamp NULL DEFAULT NULL,
                            `deletedAt` timestamp NULL DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `username` (`username`),
                            KEY `idx_tb_users_deletedAt` (`deletedAt`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

LOCK TABLES `tb_users` WRITE;

INSERT INTO `tb_users` VALUES (0,'admin','$2a$10$Fv9BWzqsiQ.JuuGdcXdvN.Fx3ml.dVR47W22GoJMWQAlm9wHQIMVe','admin','2021-04-18 15:40:33','2021-04-18 15:40:33',NULL);

UNLOCK TABLES;


