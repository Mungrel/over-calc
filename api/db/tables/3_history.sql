CREATE TABLE IF NOT EXISTS `history` (
    `id` VARBINARY(36) NOT NULL,
    `user_id` VARBINARY(36) NOT NULL,
    `left_operand_id` VARBINARY(36) NOT NULL,
    `right_operand_id` VARBINARY(36) NOT NULL,
    `operation` ENUM('add', 'sub', 'mult', 'div') NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `history_user_id_user_id`
        FOREIGN KEY `fk_history_user_id` (`user_id`)
        REFERENCES `user` (`id`)
        ON DELETE CASCADE,
    CONSTRAINT `left_operand_id_operand_id`
        FOREIGN KEY `fk_history_left_operand_id` (`left_operand_id`)
        REFERENCES `operand` (`id`)
        ON DELETE CASCADE,
    CONSTRAINT `right_operand_id_operand_id`
        FOREIGN KEY `fk_history_right_operand_id` (`right_operand_id`)
        REFERENCES `operand` (`id`)
        ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
