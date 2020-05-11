# EGTS receiver plugins

[Egts receiver](https://github.com/egts/egts-receiver) storage realization as plugins:

- [PostgreSQL](./postgresql/)
- [RabbitMQ](./rabbitmq/)
- [Tarantool](./tarantool_queue)
- [Nats](./nats)

That create a new plugin you must implementation ```Connector``` interface:

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