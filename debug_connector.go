package main

import (
	"fmt"
)

type DebugConnector struct{}

func (c *DebugConnector) Init() error {
	return nil
}

func (c *DebugConnector) Save(msg EgtsSavePacket) error {
	innerPkg, err := msg.ToBytes()
	if err != nil {
		return fmt.Errorf("Ошибка сериализации  пакета: %v", err)
	}

	fmt.Println("Export packet: ", string(innerPkg))
	return nil
}

func (c *DebugConnector) Close() error {
	return nil
}
