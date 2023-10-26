# Скачивание образа
FROM golang:alpine

# Установка рабочей директории
WORKDIR /usr/src/finapp

# Установка зависимостей
RUN go install github.com/swaggo/swag/cmd/swag@latest
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Копирование проекта
COPY . ./

# Сборка
RUN swag init -g server.go
RUN go build -v -o /usr/local/bin/finapp ./server.go

# Копирование entrypoint.sh
COPY docker/entrypoint.sh /usr/local/bin
RUN chmod +x /usr/local/bin/entrypoint.sh

# Запуск
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD ["finapp", "app:serve"]
