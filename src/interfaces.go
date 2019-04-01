package main

// BinaryData интерфейс для работы с бинарными секциями
type BinaryData interface {
	Decode([]byte) error
	Encode() ([]byte, error)
	Length() uint16
}

//Connector интерфейс для подключения внешних хранилищ
type Connector interface {
	Init(map[string]string) error
	Save(interface{ ToBytes() ([]byte, error) }) error
	Close() error
}
