# EGTS receiver plugins

Хранилища для выходных записей реализованы в форме плагинов:

- [PostgreSQL](./postgresql/)
- [RabbitMQ](./rabbitmq/)
- [Tarantool](./tarantool_queue)
- [Nats](./nats)

Есл необходим новый плагин, то он реализуется четез определение интерфейса ```Connector```:

```go
type Connector interface {
	// setup store connection
	Init(map[string]string) error
	
	// save to store method
	Save(interface{ ToBytes() ([]byte, error) }) error
	
	// close connection with store
	Close() error
}
```
