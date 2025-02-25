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
);