![Static Badge](https://img.shields.io/badge/%D1%81%D1%82%D0%B0%D1%82%D1%83%D1%81-%D0%B3%D0%BE%D1%82%D0%BE%D0%B2%D0%BE-blue)
![Static Badge](https://img.shields.io/badge/GO-1.23+-blue)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/zagart47/rsssf)
![GitHub last commit (by committer)](https://img.shields.io/github/last-commit/zagart47/rsssf)
![GitHub forks](https://img.shields.io/github/forks/zagart47/rsssf)

# Итоговый проект SF
Микросервисы для работы с RSS и комментариями

## Содержание
- [Технологии](#технологии)
- [Использование](#использование)
- [Разработка](#разработка)
- [Contributing](#contributing)
- [FAQ](#faq)
- [To do](#to-do)
- [Команда проекта](#команда-проекта)

## Технологии
- [Golang](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)

## Использование
Внести настройки бд в ```config/config.json```.

Склонировать репозитории
```powershell
go run cmd/<repo_name>/main.go
```

Gateway поддерживает следующие эндпоинты

```
GET /news/filter
```
параметры
```
page - страница новостей (по умолчанию "1", лимит новостей на страницу - "10")
latest - сортирует по дате с конца
s - поиск по фразе в заголовке новости
```

```
GET "/news/{id:[0-9]+}"

получение новости по id
```

## Разработка

### Требования
Для установки и запуска проекта необходимы golang и прямые руки.

## Contributing
Если у вас есть предложения или идеи по дополнению проекта или вы нашли ошибку, то пишите мне в tg: @zagart47

## FAQ
### Зачем вы разработали этот проект?
Это моя итоговая работа.


## Команда проекта
- [Артур Загиров](https://t.me/zagart47) — Golang Developer

