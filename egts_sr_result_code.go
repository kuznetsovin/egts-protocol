package main

import (
	"bytes"
	"fmt"
)

type EgtsSrResultCode struct {
	ResultCode uint8
}

func (s *EgtsSrResultCode) Decode(content []byte) error {
	var (
		err error
	)
	buf := bytes.NewBuffer(content)

	if s.ResultCode, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить код результата: %v", err)
	}

	return err
}

func (s *EgtsSrResultCode) Encode() ([]byte, error) {
	var (
		result []byte
		err    error
	)
	buf := new(bytes.Buffer)

	if err = buf.WriteByte(s.ResultCode); err != nil {
		return result, fmt.Errorf("Не удалось записать код результата: %v", err)
	}

	result = buf.Bytes()
	return result, err
}

func (s *EgtsSrResultCode) Length() uint16 {
	var result uint16

	if recBytes, err := s.Encode(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
