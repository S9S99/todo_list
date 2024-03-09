-- Create "todos" table
CREATE TABLE `todos` (`id` bigint NOT NULL AUTO_INCREMENT, `title` varchar(255) NOT NULL, `tag` varchar(255) NOT NULL, `body` mediumtext NULL, `completed` bool NULL, `progress` tinyint unsigned NOT NULL DEFAULT 0, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
