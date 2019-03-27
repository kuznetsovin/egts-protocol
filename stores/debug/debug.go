package main

import (
	"fmt"
)

type debugConnector struct{}

func (c debugConnector) Init(cfg map[string]string) error {
	return nil
}

func (c debugConnector) Save(msg interface{ ToBytes() ([]byte, error) }) error {
	innerPkg, err := msg.ToBytes()
	if err != nil {
		return fmt.Errorf("Ошибка сериализации  пакета: %v", err)
	}

	fmt.Println("Export packet: ", string(innerPkg))
	return nil
}

func (c debugConnector) Close() error {
	return nil
}

var Connector debugConnector
