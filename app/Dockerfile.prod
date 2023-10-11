###########
# BUILDER #
###########

# Скачивание образа
FROM golang as builder

# Установка рабочей директории
WORKDIR /usr/src/app

# Установка переменных окружения
ENV CGO_ENABLED 0
ENV GOOS linux

# Установка зависимостей
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Копирование проекта
COPY . ./

# Сборка
RUN go build -v -o /usr/local/bin/app ./cmd/main.go

#########
# FINAL #
#########

# Скачивание образа
FROM alpine

# Создание директории для пользователя app
RUN mkdir -p /home/app

# Создание пользователя app
RUN addgroup -S app && adduser -S app -G app

# Создание каталогов
ENV HOME=/home/app
ENV APP_HOME=/home/app/web
RUN mkdir $APP_HOME
WORKDIR $APP_HOME

# Копирование entrypoint.sh
COPY entrypoint.sh $APP_HOME
RUN sed -i 's/\r$//g' $APP_HOME/entrypoint.sh
RUN chmod +x $APP_HOME/entrypoint.sh

# Установка приложения
COPY --from=builder /usr/local/bin/app $APP_HOME
RUN chmod +x $APP_HOME/app

# Передача владения файлами пользовалелю app
RUN chown -R app:app $APP_HOME

# Смена пользователя
USER app

# Запуск
ENTRYPOINT ["/home/app/web/entrypoint.sh"]
