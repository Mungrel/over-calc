CREATE TABLE IF NOT EXISTS `operand` (
    `id` VARBINARY(36) NOT NULL,
    `user_id` VARBINARY(36) NOT NULL,
    `value` DECIMAL(15, 5) NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `operand_user_id_user_id`
        FOREIGN KEY `fk_operand_user_id` (`user_id`)
        REFERENCES `user` (`id`)
        ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
