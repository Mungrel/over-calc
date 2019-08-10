CREATE TABLE IF NOT EXISTS `user` (
    `id` VARBINARY(36) NOT NULL,
    `username` VARCHAR(127) NOT NULL,
    `password` BINARY(60) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `user_uniq_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
