package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type EgtsSrAbsCntrData struct {
	CounterNumber uint8  `json:"CN"`
	CounterValue  uint32 `json:"CNV"`
}

func (e *EgtsSrAbsCntrData) Decode(content []byte) error {
	var (
		err error
	)
	buf := bytes.NewReader(content)

	if e.CounterNumber, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить номер счетного входа: %v", err)
	}

	tmpBuf := make([]byte, 3)
	if _, err = buf.Read(tmpBuf); err != nil {
		return fmt.Errorf("Не удалось получить значение показаний счетного входа: %v", err)
	}

	counterVal := append(tmpBuf, 0x00)
	e.CounterValue = binary.LittleEndian.Uint32(counterVal)

	return err
}

func (e *EgtsSrAbsCntrData) Encode() ([]byte, error) {
	var (
		err    error
		result []byte
	)
	buf := new(bytes.Buffer)

	if err = buf.WriteByte(e.CounterNumber); err != nil {
		return result, fmt.Errorf("Не удалось записать номер счетного входа: %v", err)
	}

	counterVal := make([]byte, 4)
	binary.LittleEndian.PutUint32(counterVal, e.CounterValue)
	if _, err = buf.Write(counterVal[:3]); err != nil {
		return result, fmt.Errorf("Не удалось записать значение показаний счетного входа: %v", err)
	}

	result = buf.Bytes()
	return result, err
}

func (e *EgtsSrAbsCntrData) Length() uint16 {
	var result uint16

	if recBytes, err := e.Encode(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
