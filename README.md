# Приемщик EGTS

Реаализация сервиса приема данных по протоколу ЕГТС. Разбор пакета с данными делается с помощью 
библиотеки [egtslib](pkg/egtslib/README.md).

Приемщик сохраняет все записи из пакета, которые содержат позапись местонахождения (EGTS_SR_POS_DATA). 

Хранилища для выходных записей реализованы в форме плагинов:

- [PostgreSQL](pkg/store-plugins/postgresql/README.md)
- [RabbitMQ](pkg/store-plugins/rabbitmq/README.md)

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