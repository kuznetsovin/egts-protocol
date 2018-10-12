package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type RecordData struct {
	SubrecordType   byte
	SubrecordLength uint16
	SubrecordData   BinaryData
}

func (rd *RecordData) Decode(recordBytes []byte) error {
	var (
		err   error
	)

	buf := bytes.NewBuffer(recordBytes)
	if rd.SubrecordType, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить тип записи rd: %v", err)
	}

	tmpIntBuf := make([]byte, 2)
	if _, err = buf.Read(tmpIntBuf); err != nil {
		return fmt.Errorf("Не удалось получить длину записи rd: %v", err)
	}
	rd.SubrecordLength = binary.LittleEndian.Uint16(tmpIntBuf)

	subRecordBytes := buf.Next(int(rd.SubrecordLength))

	switch rd.SubrecordType {
	case EGTS_SR_POS_DATA:
		rd.SubrecordData = &EgtsSrPosData{}
	default:
		return fmt.Errorf("Не известный тип подзаписи: %d", rd.SubrecordType)
	}

	if err = rd.SubrecordData.Decode(subRecordBytes); err != nil {
		return err
	}

	return err
}

func (rd *RecordData) Encode() ([]byte, error) {
	var result []byte

	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, rd.SubrecordType); err != nil {
		return result, err
	}

	if err := binary.Write(buf, binary.LittleEndian, rd.SubrecordLength); err != nil {
		return result, err
	}

	srd, err := rd.SubrecordData.Encode()
	if err != nil {
		return result, err
	}

	buf.Write(srd)

	result = buf.Bytes()
	return result, nil
}

