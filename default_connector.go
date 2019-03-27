package main

import (
	"fmt"
)

type defaultConnector struct{}

func (c defaultConnector) Init(cfg map[string]string) error {
	return nil
}

func (c defaultConnector) Save(msg interface{ ToBytes() ([]byte, error) }) error {
	innerPkg, err := msg.ToBytes()
	if err != nil {
		return fmt.Errorf("Ошибка сериализации  пакета: %v", err)
	}

	fmt.Println("Export packet: ", string(innerPkg))
	return nil
}

func (c defaultConnector) Close() error {
	return nil
}
