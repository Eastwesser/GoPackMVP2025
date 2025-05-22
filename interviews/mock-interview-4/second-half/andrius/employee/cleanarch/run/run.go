package run

import (
	"cleanarch/internal/controller"
	"cleanarch/internal/repository"
	"cleanarch/internal/usecase"
	"log"
	"net/http"
	"time"
)

func Run() {
	// dependency injection

	// 1. Создаем базовый репозиторий (in-memory реализация)
	baseRepo := repository.NewEmpRepo() // инфраструктура

	// 2. Оборачиваем в прокси с кэшированием (TTL = 5 минут)
	cachedProxyRepo := repository.NewCacheProxyRepo(baseRepo, 5*time.Minute)

	//uc := usecase.NewEmpUseCase(baseRepo) // old (without proxy)

	// 3. Создаем use case, который будет работать с репозиторием
	//    Он не знает, что это прокси - для него это просто IEmpRepo
	uc := usecase.NewEmpUseCase(cachedProxyRepo) // бизнес-логика

	// 4. Настраиваем роутер и HTTP обработчики
	router := controller.NewEmpRouter(uc) // доставка

	// 5. Запускаем сервер
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
