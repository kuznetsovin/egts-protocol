package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type EgtsSrResponse struct {
	ConfirmedRecordNumber uint16
	RecordStatus          uint8
}

func (s *EgtsSrResponse) Decode(content []byte) error {
	var (
		err error
	)
	buf := bytes.NewBuffer(content)

	tmpIntBuf := make([]byte, 2)
	if _, err = buf.Read(tmpIntBuf); err != nil {
		return fmt.Errorf("Не удалось получить номер подтверждаемой записи: %v", err)
	}
	s.ConfirmedRecordNumber = binary.LittleEndian.Uint16(tmpIntBuf)

	if s.RecordStatus, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить статус обработки записи: %v", err)
	}

	sfd := ServiceDataSet{}
	if err = sfd.Decode(buf.Bytes()); err != nil {
		return err
	}
	return err
}

func (s *EgtsSrResponse) Encode() ([]byte, error) {
	var (
		result []byte
		err    error
	)
	buf := new(bytes.Buffer)

	if err = binary.Write(buf, binary.LittleEndian, s.ConfirmedRecordNumber); err != nil {
		return result, fmt.Errorf("Не удалось записать номер подтверждаемой записи: %v", err)
	}

	if err = buf.WriteByte(s.RecordStatus); err != nil {
		return result, fmt.Errorf("Не удалось записать статус обработки записи: %v", err)
	}

	result = buf.Bytes()
	return result, err
}

func (s *EgtsSrResponse) Length() uint16 {
	var result uint16

	if recBytes, err := s.Encode(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}

func createSrRecordResponse(pkgNum, recordNum uint16) ([]byte, error) {
	respSection := EgtsSrResponse{
		ConfirmedRecordNumber: recordNum,
		RecordStatus:          egtsPcOk,
	}
	respPkg := EgtsPackage{
		ProtocolVersion:   1,
		SecurityKeyID:     0,
		Prefix:            "00",
		Route:             "0",
		EncryptionAlg:     "00",
		Compression:       "0",
		Priority:          "00",
		HeaderLength:      11,
		HeaderEncoding:    0,
		FrameDataLength:   respSection.Length(),
		PacketIdentifier:  pkgNum,
		PacketType:        egtsPtAppdata,
		ServicesFrameData: &respSection,
	}

	return respPkg.Encode()
}
