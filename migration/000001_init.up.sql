USE engineerpro;

CREATE TABLE `users` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `password` VARCHAR(50) NOT NULL,
  `salt` VARCHAR(20) NOT NULL,
  `first_name` VARCHAR(50) NOT NULL,
  `last_name` VARCHAR(50) NOT NULL,
  -- `date_of_birth` TIMESTAMP NOT NULL,
  `email` VARCHAR(50) NOT NULL,
  `status` BOOL NOT NULL DEFAULT 1,
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` TIMESTAMP NULL,
  INDEX idx_username (email)
);

CREATE TABLE `posts` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `content_text` VARCHAR(500),
  `content_image_path` VARCHAR(255),
  `user_id` INT NOT NULL,
  `status` BOOL NOT NULL DEFAULT 1,
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` TIMESTAMP NULL,
  FOREIGN KEY (`user_id`) REFERENCES users(`id`)
);

CREATE TABLE `following` (
  `user_id` INT NOT NULL,
  `friend_id` INT NOT NULL,
  FOREIGN KEY (`user_id`) REFERENCES users(`id`)
  FOREIGN KEY (`friend_id`) REFERENCES users(`id`)
  PRIMARY KEY (user_id, friend_id)
);

CREATE TABLE `comments` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `content` VARCHAR(255),
  `user_id` INT NOT NULL,
  `post_id` INT NOT NULL,
  `status` BOOL NOT NULL DEFAULT 1,
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` TIMESTAMP NULL,
  FOREIGN KEY (`user_id`) REFERENCES users(`id`),
  FOREIGN KEY (`post_id`) REFENRECES posts(`id`)
);

CREATE TABLE `likes` (
  `user_id` INT NOT NULL,
  `post_id` INT NOT NULL,
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (post_id) REFERENCES posts(id),
  PRIMARY KEY (post_id, user_id)
);