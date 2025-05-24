# CRUD users example (with clean architecture)

A demonstration of clean architecture implementation in Go for user management with Chi router.

## Project Structure
```text
cleanuser/
├── cmd/
│ └── main.go # Application entry point
├── internal/
│ ├── controllers/ # HTTP layer
│ │ ├── handlers/ # Request handlers
│ │ └── routers/ # Route definitions
│ ├── domain/ # Core business entities and interfaces
│ ├── repository/ # Data access implementations
│ ├── server/ # Server configuration
│ └── usecase/ # Business logic layer
```

## Key Features

- Clean Architecture implementation
- Chi router for HTTP routing
- In-memory storage (can be easily replaced with DB)
- Proper layer separation
- Ready-to-use CRUD operations

## Getting Started

### Prerequisites
- Go 1.20+
- Git

### Installation
```bash
git clone https://github.com/yourusername/cleanuser.git
cd cleanuser
go mod tidy
```

## Running the Server
```bash
go run cmd/main.go
```

## API Documentation
User Model
```json
{
  "id": 1,
  "name": "John Doe",
  "age": 30,
  "nickname": "johndoe",
  "phone": "+1234567890",
  "email": "john@example.com"
}
```

## Available Endpoints
```text
Method	Endpoint	Description	Example Request
POST	/users	        Create new user	{"name":"Alice","age":25,"nickname":"alice","email":"alice@example.com"}
GET	/users	        List all users	-
GET	/users/{id}	Get user by ID	-
PUT	/users/{id}	Update user	{"name":"Alice Updated","age":26}
DELETE	/users/{id}	Delete user	-
```

## Example Requests
```bash
# Create user
curl -X POST -H "Content-Type: application/json" -d '{
  "name": "Alice",
  "age": 25,
  "nickname": "alice",
  "email": "alice@example.com"
}' 
```

**http://localhost:8080/users**

### List users
```bash
curl http://localhost:8080/users
```

### Get user
```bash
curl http://localhost:8080/users/1
```

### Update user

```bash
curl -X PUT -H "Content-Type: application/json" -d '{
  "name": "Alice Updated",
  "age": 26
}'
```
**http://localhost:8080/users/1**

### Delete user
```bash
curl -X DELETE http://localhost:8080/users/1
```

## Development Workflow

### Define models in domain/

- Business entities

- Repository interfaces

### Implement storage in repository/

- In-memory implementation

- Can be replaced with DB implementation

### Add business logic in usecase/

- Validation rules

- Business operations

### Create handlers in controllers/handlers/

- Request/response handling

- Error handling

### Set up routes in controllers/routers/

- API endpoints

- Middleware

### Configure server in server/

- Dependency initialization

### Server setup
```text
Architecture Diagram
HTTP Requests → [Controllers] → [Use Cases] → [Repositories] → [Data Storage]
                      ↑            ↑
                 (DTOs)      (Business Entities)
```

### Future Improvements
- Add database persistence (PostgreSQL)

- Implement authentication

- Add request validation

- Add logging middleware

- Add Swagger documentation

- Implement unit tests

```bash
# Создать пользователя
curl -X POST -H "Content-Type: application/json" -d '{"name":"Alice","age":25,"nickname":"alice","email":"alice@example.com"}' http://localhost:8080/users

# Получить всех пользователей
curl http://localhost:8080/users

# Получить пользователя по ID
curl http://localhost:8080/users/1

# Обновить пользователя
curl -X PUT -H "Content-Type: application/json" -d '{"name":"Alice Updated","age":26,"nickname":"alice","email":"alice@example.com"}' http://localhost:8080/users/1

# Удалить пользователя
curl -X DELETE http://localhost:8080/users/1
```