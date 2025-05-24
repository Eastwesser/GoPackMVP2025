package run

import (
	"cleanuser/internal/controllers"
	"log"
	"net/http"

	"cleanuser/internal/repository"
	"cleanuser/internal/usecase"
)

func Run() {
	// Инициализация зависимостей
	repo := repository.NewMemoryUserRepository()

	uc := usecase.NewUserUseCase(repo)

	router := controllers.SetupUserRouter(uc)

	// CONTROLLER -> UC -> REPO
	// REPO -> UC -> CONTROLLER

	// Запуск сервера
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
