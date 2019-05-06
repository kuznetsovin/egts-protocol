package egts_receiver

//Connector интерфейс для подключения внешних хранилищ
type Connector interface {
	// установка соединения с хранилищем
	Init(map[string]string) error

	// сохранение в хранилище
	Save(interface{ ToBytes() ([]byte, error) }) error

	//закрытие соединения с хранилищем
	Close() error
}
