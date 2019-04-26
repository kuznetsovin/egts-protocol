# Протокол сбора телеметрии ЕГТС

Реализация протокола ЕГТС.

Для компиляции, приложению требуются следующие пакеты:

- [toml](github.com/BurntSushi/toml)
- [protobuf](github.com/golang/protobuf/proto)
- [gommon/log](github.com/labstack/gommon/log)
- [go.uuid](github.com/satori/go.uuid)

По умолчаню все принимаемы пакеты выводятся в лог. 

Данное поведение можно изменить подключанием плагинов, которые задаются в секции ```store``` конфигурационного файла.
На данный момент есть имеются следующие плагины:

- [PostgreSQL](plugin_stores/postgresql/README.md)
- [RabbitMQ](plugin_stores/rabbitmq/README.md)

Для написания своего плагина необходимо реализовать интерфейс:

```go
//Connector интерфейс для подключения внешних хранилищ
type Connector interface {
	// установка соединения с хранилищем
	Init(map[string]string) error
	
	// сохранение в хранилище
	Save(interface{ ToBytes() ([]byte, error) }) error
	
	//закрытие соединения с хранилищем
	Close() error
}
```