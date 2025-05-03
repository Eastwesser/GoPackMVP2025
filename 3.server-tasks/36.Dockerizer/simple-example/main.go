package main

import (
	"fmt"
	"log"
	"net/http"
)

// Определяем тип функции для обработки запросов
type RequestHandler func(w http.ResponseWriter, r *http.Request)

// Реализуем метод ServeHTTP для типа RequestHandler
func (rh RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rh(w, r)
}

func main() {
	// Определяем обработчик
	handler := RequestHandler(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Привет, мир!")
	})

	// Регистрируем обработчик
	http.Handle("/", handler)

	// Запускаем сервер
	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
