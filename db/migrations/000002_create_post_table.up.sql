CREATE TABLE `posts` (
  `pst_id` int(11) NOT NULL AUTO_INCREMENT,
  `pst_title` varchar(60) NOT NULL,
  `pst_date_creation` timestamp NOT NULL DEFAULT current_timestamp(),
  `pst_date_edit` timestamp DEFAULT NULL,
  `pst_content` longtext NOT NULL,
  PRIMARY KEY (`pst_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;