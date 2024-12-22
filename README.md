# Yandex_Lyceum
___
Итоговая задача модуля 1 курса Яндекс Лицей. 
Данный проект реализует веб-сервис, принимающий 
http-запрос с выражением, состоящем из _цифр_ и возвращающий ответ в JSON-формате
___
## Инструкция по запуску 

1. [Установите](https://www.youtube.com/watch?v=xoz-Y9T8gRc&t=319s) Golang на свой компьютер (версию 1.23 и выше)
2. Откройте терминал и впишите следующую команду:
```bash
go run ./cmd/main.go
```
3. Откройте другой терминал и впишите свой cURL запрос с портом 8080:

```bash
curl --location 'localhost:8080/api/v1/calculate' \         
--header 'Content-Type: application/json' \         
--data '
{
  "expression": "Введите своё арифметическое выражение"
}'
```
> Если вы пользуетесь Windows, то запросы нужно делать через Git Bash или WSL
___
## Примеры запросов
___
#### Успешный запрос:

```bash
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```

#### Ответ:

```
{
  "result": "6"
}
```
___
#### Ошибка 422 (невалидный запрос):

```bash
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+a"
}'
```

#### Ответ:

```
{
  "error": "invalid expression"
}
```
___
#### Ошибка 500 (внутренняя ошибка сервера):

```bash
curl --location 'localhost:8080/api/v1/calculate'
```

#### Ответ:

```
{
  "error": "internal server error"
}
```