# Используем официальный образ Go
FROM golang:1.20

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum и устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальные файлы
COPY . .

# Собираем приложение
RUN go build -o appointments-service

# Открываем порт (по умолчанию Gin использует 8080)
EXPOSE 8080

# Команда запуска
CMD ["./appointments-service"]
