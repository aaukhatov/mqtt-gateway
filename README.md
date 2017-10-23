# MqttService
[![Build Status](https://api.travis-ci.org/aukhatov/MqttService.png?branch=master)](https://travis-ci.org/aukhatov/MqttService)

## Описание
Это легковесный веб-сервис для работы с датчиками по протоколу MQTT.

### Структура проекта
Подробности по настройки окружения здесь: https://golang.org/doc/code.html

Кратко:
* В домашней директории пользователя должна быть папка go/
* Переменная окружения $GOPATH = '~/go' (одно окружение для всех проектов, по конвенции Go)
* Переменная окружения $GOROOT должна указывать на Go SDK

Выполнить команду go get для клонирования проекта
```bash
go get github.com/aukhatov/MqttService
```

#### Компиляция

Лучше всего компилировать следующей командой:
```bash
go install
```
Выполнять в директории где находится файл main.go

#### Запуск

Этап компиляции создаст бинарный файл ~/go/bin/MqttService
На текущий момент можно передать один аргумент, это порт на котором будет запущен веб-сервис

```bash
~/go/bin/MqttService port=8080
```
