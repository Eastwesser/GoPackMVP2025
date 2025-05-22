package run

import (
	"github.com/gin-gonic/gin"
	"internal/__controller"
	"internal/__repository"
	"internal/__usecase"
	"time"
)

func Run() {
	// Инициализация слоёв
	mockRepo := &__repository.MockUserRepo{}
	cachedRepo := __repository.NewCachedUserRepo(mockRepo, 5*time.Minute)
	uc := __usecase.NewUserUsecase(cachedRepo)
	handler := __controller.NewUserHandler(uc)

	// Настройка Gin
	r := gin.Default()
	__controller.SetupRoutes(r, handler)

	// Запуск
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
