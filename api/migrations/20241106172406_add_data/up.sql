INSERT INTO `brands` (`brand_id`, `name`)
VALUES
(1,'nestle'),
(2, 'carrefour');

INSERT INTO `foods` (`food_id`,`name`, `brand_id`, `calories`, `protein`, `carbs`, `sugar`, `fat`, `saturated_fat`, `fiber`, `sodium`, `image_url`)
VALUES
(1, 'Farine T150', 2, 350, 10, 70, 0, 1, 0, 3, 0, ''),
(2, 'Sucre', 2, 400, 0, 100, 100, 0, 0, 0, 0, ''),
(3, 'Levure', 2, 0, 0, 0, 0, 0, 0, 0, 0, ''),
(4, 'Lait', 2, 50, 5, 5, 5, 5, 5, 5, 5, ''),
(5, 'Beurre', 2, 500, 0, 0, 0, 50, 50, 0, 0, ''),
(6, 'Oeuf', 2, 100, 10, 0, 0, 10, 10, 0, 0, ''),
(7, 'Chocolat', 2, 500, 10, 50, 50, 50, 50, 0, 0, ''),
(8, "Saint Môret", 2, 200, 10, 10, 10, 10, 10, 10, 10, ''),
(9, "Erythritol", 2, 0, 0, 0, 0, 0, 0, 0, 0, ''),
(10, "Farine T110", 2, 300, 10, 60, 0, 1, 0, 3, 0, '');

INSERT INTO `tags` (`tag_id`, `name`, `color`)
VALUES
(1, "Dessert", "#F1B"),
(2, "Epicerie sucrée", "#F94");