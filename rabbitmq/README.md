## Плагин для работы с RabbitMQ

Для работы плагина нужна библиотека для работы с [amqp](github.com/streadway/amqp)

Секция настроек выглядит следующим образом:

```
[store]
plugin = "rabbitmq.so"
host = "localhost"
port = "5672"
user = "guest"
password = "guest"
key = ""
```

Описание парамеров:

- *plugin* - путь до библиотеки
- *host* - адрес сервера
- *port* - порт
- *user* - пользователь
- *password* - пароль
- *exchange* - exchange для запасиси
- *key* - имя конкретной очереди (если нужно)

