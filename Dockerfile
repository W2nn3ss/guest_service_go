# Задаем базовый образ
FROM golang:1.20-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы проекта в контейнер
COPY . .

# Загружаем зависимости
RUN go mod tidy

# Собираем приложение
RUN go build -o main .

# Определяем команду запуска контейнера
CMD ["/app/main"]