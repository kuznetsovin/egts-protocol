## NATS plugin

Plugin depends on [nats](http://github.com/nats-io/nats.go) package

Config section:

```
[store]
plugin = "nats.so"
servers = "nats://localhost:1222, nats://localhost:1223, nats://localhost:1224"
topic = "egts"
user = "guest"
password = "guest"

```

Parameters description:

- *plugin* - path to *.so library file  
- *servers* - nats server address
- *topic* - nats topic, where publish message
- *user* - nats user
- *password* - user password

