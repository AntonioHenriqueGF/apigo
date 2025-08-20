CREATE TABLE `users` (
    `usu_id` int(11) NOT NULL AUTO_INCREMENT,
    `usu_name` varchar(60) NOT NULL,
    `usu_email` varchar(100) NOT NULL,
    `usu_password_hash` varchar(255) NOT NULL,
    `usu_created_at` timestamp NOT NULL DEFAULT current_timestamp(),
    PRIMARY KEY (`usu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;