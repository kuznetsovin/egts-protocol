## Плагин для работы с NATS

Для работы плагина нужна библиотека для работы с [nats](github.com/nats-io/nats.go)

Секция настроек выглядит следующим образом:

```
[store]
plugin = "nats.so"
servers = "nats://localhost:1222, nats://localhost:1223, nats://localhost:1224"
topic = "egts"

```

Описание парамеров:

- *plugin* - путь до библиотеки
- *servers* - адреса серверов
- *topic* - топик для отправки сообщений

