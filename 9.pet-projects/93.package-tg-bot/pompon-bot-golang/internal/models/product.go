package models

// Модель товара
/*
	product.go: Модель для товаров (id, название, цена, описание, категория).
*/

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	CategoryID  int      // связь с категорией
	Category    Category // добавляем связь с моделью Category
}
