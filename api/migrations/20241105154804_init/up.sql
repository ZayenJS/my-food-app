CREATE TABLE IF NOT EXISTS `brands` (
  `brand_id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL UNIQUE,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `foods` (
  `food_id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL,
  `brand_id` INT NOT NULL REFERENCES `brands`(`brand_id`) ON DELETE CASCADE,
  `calories` DECIMAL NOT NULL,
  `protein` DECIMAL NOT NULL,
  `carbs` DECIMAL NOT NULL,
  `sugar` DECIMAL NOT NULL,
  `fat` DECIMAL NOT NULL,
  `saturated_fat` DECIMAL NULL,
  `fiber` DECIMAL NULL,
  `sodium` DECIMAL NULL,
  `image_url` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `recipes` (
  `recipe_id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL UNIQUE,
  `description` TEXT NOT NULL,
  `difficulty` INT NOT NULL,
  `prep_time` INT NOT NULL,
  `cook_time` INT NOT NULL,
  `rest_time` INT NOT NULL,
  `servings` INT NOT NULL,
  `rating` INT DEFAULT NULL,
  `image_url` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `recipe_ingredients` (
  `recipe_id` INT NOT NULL REFERENCES `recipes`(`recipe_id`) ON DELETE CASCADE,
  `food_id` INT NOT NULL REFERENCES `foods`(`food_id`) ON DELETE CASCADE,
  `quantity` INT NOT NULL,
  `unit` VARCHAR(50) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`recipe_id`, `food_id`)
);

CREATE TABLE IF NOT EXISTS `recipe_steps` (
  `recipe_step_id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `recipe_id` INT NOT NULL REFERENCES `recipes`(`recipe_id`) ON DELETE CASCADE,
  `order` INT NOT NULL,
  `text` TEXT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `tags` (
  `tag_id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL UNIQUE,
  `color` VARCHAR(7) DEFAULT '#fff',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `recipe_tags` (
  `recipe_id` INT NOT NULL REFERENCES `recipes`(`recipe_id`) ON DELETE CASCADE,
  `tag_id` INT NOT NULL REFERENCES `tags`(`tag_id`) ON DELETE CASCADE,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`recipe_id`, `tag_id`)
);