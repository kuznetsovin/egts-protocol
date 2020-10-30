## RabbitMQ plugin

Plugin depends on [amqp](http://github.com/streadway/amqp) package

Config section:

```
[store]
plugin = "rabbitmq.so"
host = "localhost"
port = "5672"
user = "guest"
password = "guest"
key = ""
```

Parameters description:

- *plugin* - path to *.so library file
- *host* - rabbitmq server address
- *port* - rabbitmq port
- *user* - rabbitmq user
- *password* - user password
- *exchange* - exchange for publish
- *key* - routing key (or queue name)

