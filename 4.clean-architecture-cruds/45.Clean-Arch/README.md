# CRUD users example

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