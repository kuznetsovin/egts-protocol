package main

// BinaryData интерфейс для работы с бинарными секциями
type BinaryData interface {
	Decode([]byte) error
	Encode() ([]byte, error)
	Length() uint16
}

//Connector интерфейс для подключения внешних хранилищ
type Connector interface {
	// установка соединения с хранилищем
	Init(map[string]string) error

	// сохранение в хранилище
	Save(interface{ ToBytes() ([]byte, error) }) error

	//закрытие соединения с хранилищем
	Close() error
}
