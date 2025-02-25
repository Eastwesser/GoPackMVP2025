package models

// Модель заказа

/*
	Модели базы данных для работы с PostgreSQL:
		order.go: Модель для заказов (id, пользователь, товар, количество, статус).
*/

type Order struct {
	ID       int
	UserID   int64
	Product  Product
	Quantity int
	Status   string
}
