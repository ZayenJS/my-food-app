INSERT INTO `brands` (`id`, `name`)
VALUES
(1,'Nestle'),
(2, 'Carrefour');

INSERT INTO `foods` (`id`,`name`, `brand_id`, `calories`, `protein`, `carbs`, `sugar`, `fat`, `saturated_fat`, `fiber`, `sodium`, `image_url`, `category`)
VALUES
(1, 'Farine T150', 2, 350, 10, 70, 0, 1, 0, 3, 0, '', 'Pâtisserie'),
(2, 'Sucre', 2, 400, 0, 100, 100, 0, 0, 0, 0, '', 'Pâtisserie'),
(3, 'Levure', 2, 0, 0, 0, 0, 0, 0, 0, 0, '', 'Pâtisserie'),
(4, 'Lait', 2, 50, 5, 5, 5, 5, 5, 5, 5, '', 'Produit laitier'),
(5, 'Beurre', 2, 500, 0, 0, 0, 50, 50, 0, 0, '', 'Produit laitier'),
(6, 'Oeuf', 2, 100, 10, 0, 0, 10, 10, 0, 0, '', 'Produit laitier'),
(7, 'Chocolat', 2, 500, 10, 50, 50, 50, 50, 0, 0, '', 'Pâtisserie');

INSERT INTO `tags` (`id`, `name`, `color`)
VALUES
(1, "Dessert", "#F1B"),
(2, "Epicerie sucrée", "#F94");