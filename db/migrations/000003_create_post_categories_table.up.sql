CREATE TABLE `post_categories` (
  `pca_id` int(11) NOT NULL AUTO_INCREMENT,
  `pca_id_post` int(11) NOT NULL,
  `pca_id_categorie` int(11) NOT NULL,
  PRIMARY KEY (`pca_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;