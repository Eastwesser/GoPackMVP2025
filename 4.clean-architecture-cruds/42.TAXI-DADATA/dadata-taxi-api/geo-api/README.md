# API для геокодирования адресов (Россия)

API для определения координат по адресу с использованием чистой архитектуры, CRUD операциями, JWT аутентификацией и Swagger документацией.

```bash
    geo-api/
    ├── cmd/               # Точка входа
    │   └── main.go
    ├── internal/          # Основная логика приложения
    │   ├── config/        # Конфигурация
    │   ├── controller/    # HTTP обработчики
    │   ├── entity/        # Сущности БД
    │   ├── repository/    # Работа с БД
    │   ├── service/       # Бизнес-логика
    │   ├── dadata/        # Интеграция с DaData
    │   └── middleware/    # Промежуточное ПО
    ├── migrations/        # Миграции БД
    ├── pkg/               # Вспомогательные пакеты
    │   ├── auth/          # JWT аутентификация
    │   └── logger/        # Логирование
    ├── docs/              # Swagger документация
    └── go.mod             # Зависимости
```

Download it!
```bash
    go get github.com/gin-gonic/gin@v1.9.1
    go get github.com/golang-jwt/jwt/v5@v5.0.0
    go get github.com/swaggo/swag@v1.16.2
    go get github.com/swaggo/gin-swagger@v1.6.0
    go get gorm.io/gorm@v1.25.5
    go get github.com/ekomobile/dadata/v2@v2.10.0
```

1. Геокодирование адреса
Запрос:
```http
    POST /api/geocode
    Content-Type: application/json
    Authorization: Bearer <JWT_TOKEN>
    
    {
    "address": "москва сухонская 11"
}
```
Ответ:
```json
    {
    "ID": 1,
    "CreatedAt": "2023-10-01T12:00:00Z",
    "UpdatedAt": "2023-10-01T12:00:00Z",
    "DeletedAt": null,
    "UserID": 1,
    "RawAddress": "москва сухонская 11",
    "Result": "г Москва, ул Сухонская, д 11",
    "PostalCode": "127642",
    "Country": "Россия",
    "Region": "Москва",
    "CityArea": "Северо-восточный",
    "CityDistrict": "Северное Медведково",
    "Street": "Сухонская",
    "House": "11",
    "GeoLat": 55.8782557,
    "GeoLon": 37.65372,
    "QCGeo": 0
    }
```
2. Получение истории запросов
Запрос:
```http
    GET /api/history
    Authorization: Bearer <JWT_TOKEN>
```
Ответ:
```json
    [
        {
            "ID": 1,
            "CreatedAt": "2023-10-01T12:00:00Z",
            "RawAddress": "москва сухонская 11",
            "Result": "г Москва, ул Сухонская, д 11",
            "GeoLat": 55.8782557,
            "GeoLon": 37.65372
        },
        {
            "ID": 2,
            "CreatedAt": "2023-10-01T11:30:00Z",
            "RawAddress": "санкт-петербург невский 10",
            "Result": "г Санкт-Петербург, Невский пр-кт, д 10",
            "GeoLat": 59.935846,
            "GeoLon": 30.325894
        }
    ]
```
