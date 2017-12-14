# mqtt-gateway
[![Build Status](https://travis-ci.org/aukhatov/mqtt-gateway.svg?branch=test)](https://travis-ci.org/aukhatov/mqtt-gateway)

## Описание
Это легковесный веб-сервис для работы с датчиками по протоколу MQTT.

## API

* POST /ESP/{chipId}/CONTROL/LED data: ON / OFF

### Структура проекта
Подробности по настройки окружения здесь: https://golang.org/doc/code.html

Кратко:
* В домашней директории пользователя должна быть папка go/
* Переменная окружения $GOPATH = '~/go' (одно окружение для всех проектов, по конвенции Go)
* Переменная окружения $GOROOT должна указывать на Go SDK

Выполнить команду go get для клонирования проекта
```bash
go get github.com/aukhatov/mqtt-gateway
```

#### Компиляция

Лучше всего компилировать следующей командой:
```bash
go install
```
Выполнять в директории где находится файл main.go

#### Запуск

Этап компиляции создаст бинарный файл ~/go/bin/mqtt-gateway
На текущий момент можно передать один аргумент, это порт на котором будет запущен веб-сервис.
Если не указан, порт поумолчанию равен 80.

```bash
~/go/bin/mqtt-gateway port=8080
```
