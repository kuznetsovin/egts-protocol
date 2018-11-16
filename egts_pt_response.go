package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type EgtsPtResponse struct {
	ResponsePacketID uint16
	ProcessingResult uint8
	SDR              BinaryData
}

func (s *EgtsPtResponse) Decode(content []byte) error {
	var (
		err error
	)
	buf := bytes.NewBuffer(content)

	tmpIntBuf := make([]byte, 2)
	if _, err = buf.Read(tmpIntBuf); err != nil {
		return fmt.Errorf("Не удалось получить идентификатор пакета из ответа: %v", err)
	}
	s.ResponsePacketID = binary.LittleEndian.Uint16(tmpIntBuf)

	if s.ProcessingResult, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить код обработки: %v", err)
	}

	sfd := ServiceDataSet{}
	if err = sfd.Decode(buf.Bytes()); err != nil {
		return err
	}
	return err
}

func (s *EgtsPtResponse) Encode() ([]byte, error) {
	var (
		result   []byte
		sdrBytes []byte
		err      error
	)
	buf := new(bytes.Buffer)

	if err = binary.Write(buf, binary.LittleEndian, s.ResponsePacketID); err != nil {
		return result, fmt.Errorf("Не удалось записать индентификатор пакета в ответ: %v", err)
	}

	if err = buf.WriteByte(s.ProcessingResult); err != nil {
		return result, fmt.Errorf("Не удалось записать результат обработки в пакет: %v", err)
	}

	if s.SDR != nil {
		if sdrBytes, err = s.SDR.Encode(); err != nil {
			return result, err
		}
		buf.Write(sdrBytes)
	}

	result = buf.Bytes()
	return result, err
}

func (s *EgtsPtResponse) Length() uint16 {
	var result uint16

	if recBytes, err := s.Encode(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
