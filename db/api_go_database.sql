-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Tempo de geração: 15/02/2024 às 02:27
-- Versão do servidor: 10.4.32-MariaDB
-- Versão do PHP: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Banco de dados: `api_go_database`
--

-- --------------------------------------------------------

--
-- Estrutura para tabela `categories`
--

CREATE TABLE `categories` (
  `cat_id` int(11) NOT NULL,
  `cat_description` varchar(60) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Despejando dados para a tabela `categories`
--

INSERT INTO `categories` (`cat_id`, `cat_description`) VALUES
(1, 'Teste');

-- --------------------------------------------------------

--
-- Estrutura para tabela `posts`
--

CREATE TABLE `posts` (
  `pst_id` int(11) NOT NULL,
  `pst_title` varchar(60) NOT NULL,
  `pst_date_creation` date NOT NULL DEFAULT current_timestamp(),
  `pst_date_edit` date DEFAULT NULL,
  `pst_content` longtext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Despejando dados para a tabela `posts`
--

INSERT INTO `posts` (`pst_id`, `pst_title`, `pst_date_creation`, `pst_date_edit`, `pst_content`) VALUES
(1, 'Post Teste', '2024-02-04', NULL, 'Post teste com o propósito de verificar a funcionalidade de CRUD da API.');

-- --------------------------------------------------------

--
-- Estrutura para tabela `post_categories`
--

CREATE TABLE `post_categories` (
  `pca_id` int(11) NOT NULL,
  `pca_id_post` int(11) NOT NULL,
  `pca_id_categorie` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Despejando dados para a tabela `post_categories`
--

INSERT INTO `post_categories` (`pca_id`, `pca_id_post`, `pca_id_categorie`) VALUES
(1, 1, 1);

--
-- Índices para tabelas despejadas
--

--
-- Índices de tabela `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`cat_id`);

--
-- Índices de tabela `posts`
--
ALTER TABLE `posts`
  ADD PRIMARY KEY (`pst_id`);

--
-- Índices de tabela `post_categories`
--
ALTER TABLE `post_categories`
  ADD PRIMARY KEY (`pca_id`),
  ADD KEY `pca_to_pst_constraint` (`pca_id_post`),
  ADD KEY `pca_to_cat_constraint` (`pca_id_categorie`);

--
-- AUTO_INCREMENT para tabelas despejadas
--

--
-- AUTO_INCREMENT de tabela `categories`
--
ALTER TABLE `categories`
  MODIFY `cat_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT de tabela `posts`
--
ALTER TABLE `posts`
  MODIFY `pst_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT de tabela `post_categories`
--
ALTER TABLE `post_categories`
  MODIFY `pca_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- Restrições para tabelas despejadas
--

--
-- Restrições para tabelas `post_categories`
--
ALTER TABLE `post_categories`
  ADD CONSTRAINT `pca_to_cat_constraint` FOREIGN KEY (`pca_id_categorie`) REFERENCES `categories` (`cat_id`),
  ADD CONSTRAINT `pca_to_pst_constraint` FOREIGN KEY (`pca_id_post`) REFERENCES `posts` (`pst_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
