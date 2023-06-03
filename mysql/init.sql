DROP TABLE IF EXISTS `users`;
create table `users` (
    `user_id`         BIGINT(20) AUTO_INCREMENT,
    `user_name`       VARCHAR(36) NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO users (user_name, created_at) VALUES ('ユーザー1', now());