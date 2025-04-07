// cmd/main.go
package main

import (
	"GoPackMVP2025/42.TAXI-DADATA/dadata-taxi-api/geo-api/internal/config"
	"GoPackMVP2025/42.TAXI-DADATA/dadata-taxi-api/geo-api/internal/controller"
	"GoPackMVP2025/42.TAXI-DADATA/dadata-taxi-api/geo-api/internal/entity"
	"GoPackMVP2025/42.TAXI-DADATA/dadata-taxi-api/geo-api/internal/repository"
	"GoPackMVP2025/42.TAXI-DADATA/dadata-taxi-api/geo-api/internal/service"
	"GoPackMVP2025/42.TAXI-DADATA/dadata-taxi-api/geo-api/pkg/auth"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	// Инициализация БД
	dsn := "host=" + cfg.DB.Host + " user=" + cfg.DB.User + " password=" + cfg.DB.Password +
		" dbname=" + cfg.DB.Name + " port=" + cfg.DB.Port + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Миграции
	if err := db.AutoMigrate(&entity.Address{}); err != nil {
		panic(err)
	}

	// Инициализация слоев
	addressRepo := repository.NewAddressRepository(db)
	geocoderSvc := service.NewGeocoderService(addressRepo, cfg.DaData.APIKey, cfg.DaData.SecretKey)
	geocoderCtrl := controller.NewGeocoderController(geocoderSvc)

	// Настройка роутера
	r := gin.Default()

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API группа с JWT middleware
	api := r.Group("/api", auth.JWTMiddleware(cfg.JWT.Secret))
	{
		api.POST("/geocode", geocoderCtrl.Geocode)
		api.GET("/history", geocoderCtrl.GetHistory)
		api.DELETE("/history/:id", geocoderCtrl.DeleteFromHistory)
	}

	// Запуск сервера
	r.Run(":" + cfg.Server.Port)
}
