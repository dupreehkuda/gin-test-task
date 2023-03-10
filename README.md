# Qsoft test task

## Описание

Сервис для получения кол-ва дней до запрошенного года.

Реализована "трехслойная" архитектура: `middleware` -> `handler` -> `processor`, где последнее - слой бизнес-логики.

Чтобы снизить зацепление кода, все части сервиса передаются по нужному [интерфейсу](./internal/interfaces/interfaces.go).

## Эндпоинты

У сервиса есть только один эндпоинт:

- `GET api/when/:year`
    - Хэндлер для получения кол-ва дней до запрошенного года
    - В случае ошибки выдает http-код 500
    - Ответ возвращается в формате JSON:

```json
{
    "requestedYear": 2023,
    "daysToYear": -18
}
```

## Запуск

Для сборки запуска решил использовать `Docker`, собрать контейнер можно командой
```shell
make build
```

Соответственно запустить можно командой:
```shell
make run
```

Для того чтобы запустить приложение не используя Docker, можно просто запустить приложение с флагом:
```sh
$ go run cmd/main.go -a <port>
```
Без флага `-a` запустится на порте `:8080`.
Также порт можно задать через переменнyю окружения: `SERVICE_ADDRESS`.

## Обоснование работы

По ТЗ нужно было использовать фреймворк `Gin`. Его я уже использовал, но только в одном проекте, 
так что был рад освежить в голове его основы. В своих проектах чаще отдаю предпочтение `chi`.

Так же по ТЗ нужно было отдавать строку, но я решил отдавать ответ в формате JSON. 
Посчитал, что так ответ немного наглядней, да и `Gin` предоставляет довольно удобный метод `c.JSON()`.

В качестве логгера решил попробовать `zerolog`. Раньше всегда пользовался `zap`, 
но тут решил попробовать что-то новое.

В целом по структуре и лучшим практикам – делал так, как делаю каждый свой проект, старался 
в фанатизм не уходить;)