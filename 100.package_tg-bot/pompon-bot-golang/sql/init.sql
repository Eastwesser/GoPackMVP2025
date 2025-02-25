-- Схема базы данных для проекта POMPON 🍎

-- Таблица категорий
CREATE TABLE categories
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

-- Таблица товаров
CREATE TABLE products
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(150)   NOT NULL,
    description TEXT,
    price       DECIMAL(10, 2) NOT NULL,
    category_id INT REFERENCES categories (id) ON DELETE CASCADE
);

-- Таблица подписчиков
CREATE TABLE subscribers
(
    id            SERIAL PRIMARY KEY,
    telegram_id   BIGINT NOT NULL UNIQUE,
    subscribed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Таблица заказов
CREATE TABLE orders
(
    id         SERIAL PRIMARY KEY,
    user_id    BIGINT NOT NULL,
    product_id INT    REFERENCES products (id) ON DELETE SET NULL,
    quantity   INT    NOT NULL,
    status     VARCHAR(50) DEFAULT 'pending',
    created_at TIMESTAMP   DEFAULT CURRENT_TIMESTAMP
);-- Демо-данные для тестирования

-- Добавление категорий
INSERT INTO categories (name)
VALUES ('Коробочки'),
       ('Открытки'),
       ('Обёртки');

-- Добавление товаров
INSERT INTO products (name, description, price, category_id)
VALUES ('Подарочная коробочка', 'Идеально подходит для небольших подарков', 300.00, 1),
       ('Ручная открытка', 'Уникальная открытка ручной работы', 200.00, 2),
       ('Крафтовая обёртка', 'Красивая обёрточная бумага для подарков', 150.00, 3);
