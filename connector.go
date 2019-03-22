package main

type EgtsSavePacket interface {
	ToBytes() ([]byte, error)
}

type Connector interface {
	Init() error
	Save(EgtsSavePacket, string) error
	Close() error
}
