# EGTS receiver

Simple example EGTS receiver realization. For parsing protocol binary message used library [libegts](github.com/kuznetsovin/libegts).

For compile required libraries:

- [toml](github.com/BurntSushi/toml)
- [protobuf](github.com/golang/protobuf/proto)
- [gommon/log](github.com/labstack/gommon/log)
- [go.uuid](github.com/satori/go.uuid)
- [libegts](github.com/kuznetsovin/libegts)

By default received message output to console if it includes gps coordinates section (EGTS_SR_POS_DATA). 

Besides you can connect different plugins for working with parsing data. This plugins setup in ```store``` section 
in config file. Receiver have several plugins out of the box:

- [PostgreSQL](plugin_stores/postgresql/README.md)
- [RabbitMQ](plugin_stores/rabbitmq/README.md)

If you want create another plugin, then you must implementation ```Connector``` interface in you code:

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