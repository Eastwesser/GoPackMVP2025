# Используем официальный образ Go
FROM golang:1.21-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum (если есть)
COPY go.mod .

# Скачиваем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go build -o server .

# Открываем порт 8080
EXPOSE 8080

# Запускаем сервер
CMD ["./server"]