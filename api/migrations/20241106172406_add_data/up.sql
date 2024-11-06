INSERT INTO `brand` (`id`, `name`)
VALUES
(1,'Nestle'),
(2, 'Carrefour'),
(3, 'Danone'),
(4, 'Coca Cola'),
(5, 'Pepsi'),
(6, 'Lays'),
(7, 'Pringles'),
(8, 'Kinder'),
(9, 'Ferrero'),
(10, 'Oreo'),
(11, 'Kikkoman'),
(12, 'Maille');

INSERT INTO `food` (`name`, `brand_id`, `calories`, `protein`, `carbs`, `sugar`, `fat`, `saturated_fat`, `fiber`, `sodium`, `image_url`, `category`)
VALUES
('Nesquik', 1, 100, 3, 20, 15, 1, 0, 0, 0, '', 'drink'),
('Actimel', 3, 50, 3, 10, 5, 1, 0, 0, 0, '', 'drink'),
('Coca Cola', 4, 150, 0, 40, 40, 0, 0, 0, 0, '', 'drink'),
('Pepsi', 5, 150, 0, 40, 40, 0, 0, 0, 0, '', 'drink'),
('Lays', 6, 200, 2, 20, 0, 10, 2, 0, 0, '', 'snack'),
('Pringles', 7, 200, 2, 20, 0, 10, 2, 0, 0, '', 'snack'),
('Kinder Bueno', 8, 200, 2, 20, 0, 10, 2, 0, 0, '', 'snack'),
('Ferrero Rocher', 9, 200, 2, 20, 0, 10, 2, 0, 0, '', 'snack'),
('Oreo', 10, 200, 2, 20, 0, 10, 2, 0, 0, '', 'snack'),
('Soy Sauce', 11, 0, 0, 0, 0, 0, 0, 0, 0, '', 'condiment'),
('Mustard', 12, 0, 0, 0, 0, 0, 0, 0, 0, '', 'condiment');