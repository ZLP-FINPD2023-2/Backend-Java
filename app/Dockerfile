# Скачивание образа
FROM golang:alpine as builder

# Установка рабочей директории
WORKDIR /usr/src/app

# Установка зависимостей
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Копирование проекта
COPY . ./

# Сборка
RUN go build -v -o /usr/local/bin/app ./cmd/main.go

# Копирование entrypoint.sh
COPY entrypoint.sh /usr/local/bin
RUN chmod +x /usr/local/bin/entrypoint.sh

# Запуск
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD ["app"]
