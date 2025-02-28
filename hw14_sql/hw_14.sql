-- Создаем схему БД
CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE Orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    order_date TIMESTAMPTZ DEFAULT NOW(),
    total_amount NUMERIC(10, 2) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE
);

CREATE TABLE Products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price NUMERIC(10, 2) NOT NULL
);

CREATE TABLE OrderProducts (
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL DEFAULT 1,
    PRIMARY KEY (order_id, product_id),
    FOREIGN KEY (order_id) REFERENCES Orders(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES Products(id) ON DELETE CASCADE
);

-- Вставка пользователей
INSERT INTO Users (name, email, password) VALUES ('Денис', 'denis@play.ru', 'Paaaswoord');
INSERT INTO Users (name, email, password) VALUES ('Виктор', 'viktor@play.ru', 'Paaaswoord1');

-- Редактирование пользователя
UPDATE Users SET name = 'Денис Б' WHERE id = 1;

-- Вставка продуктов
INSERT INTO Products (name, price) VALUES ('Магник', 29213.21);
INSERT INTO Products (name, price) VALUES ('КубикРубик', 2313.99);

-- Редактирование товара
UPDATE Products SET price = 11111.99 WHERE id = 1;

-- Сохранение заказа
INSERT INTO Orders (user_id, total_amount) VALUES (1, 59999.98);
INSERT INTO Orders (user_id, total_amount) VALUES (1, 59999.98);

-- Удаление
DELETE FROM Users WHERE id = 2;
DELETE FROM Products WHERE id = 2;
DELETE FROM Orders WHERE id = 2;

-- Выборка всех пользователей
SELECT * FROM Users;

-- Выборка всех продуктов
SELECT * FROM Products;

-- Выборка заказов по пользователю
SELECT * FROM Orders WHERE user_id = 1;

-- Выборка статистики по пользователю (общая сумма заказов, средняя цена товара)
SELECT 
    u.id AS user_id, 
    u.name, 
    COALESCE(SUM(o.total_amount), 0) AS total_spent, 
    COALESCE(AVG(p.price), 0) AS avg_product_price
FROM Users u
LEFT JOIN Orders o ON u.id = o.user_id
LEFT JOIN OrderProducts op ON o.id = op.order_id
LEFT JOIN Products p ON op.product_id = p.id
GROUP BY u.id;

-- Индексы для ускорения выборки
CREATE INDEX idx_orders_user_id ON Orders(user_id);
CREATE INDEX idx_orderproducts_order_id ON OrderProducts(order_id);
CREATE INDEX idx_orderproducts_product_id ON OrderProducts(product_id);
CREATE INDEX idx_products_name ON Products(name);