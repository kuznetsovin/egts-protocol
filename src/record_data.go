package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type RecordData struct {
	SubrecordType   byte       `json:"SRT"`
	SubrecordLength uint16     `json:"SRL"`
	SubrecordData   BinaryData `json:"SRD"`
}

//RecordDataSet описывает массив с подзаписями протокола ЕГТС
type RecordDataSet []RecordData

func (rds *RecordDataSet) Decode(recDS []byte) error {
	var (
		err error
	)
	buf := bytes.NewBuffer(recDS)
	for buf.Len() > 0 {
		rd := RecordData{}
		if rd.SubrecordType, err = buf.ReadByte(); err != nil {
			return fmt.Errorf("Не удалось получить тип записи subrecord data: %v", err)
		}

		tmpIntBuf := make([]byte, 2)
		if _, err = buf.Read(tmpIntBuf); err != nil {
			return fmt.Errorf("Не удалось получить длину записи subrecord data: %v", err)
		}
		rd.SubrecordLength = binary.LittleEndian.Uint16(tmpIntBuf)

		subRecordBytes := buf.Next(int(rd.SubrecordLength))

		switch rd.SubrecordType {
		case egtsSrPosData:
			rd.SubrecordData = &EgtsSrPosData{}
		case egtsSrTermIdentity:
			rd.SubrecordData = &EgtsSrTermIdentity{}
		case egtsSrRecordResponse:
			rd.SubrecordData = &EgtsSrResponse{}
		case egtsSrResultCode:
			rd.SubrecordData = &EgtsSrResultCode{}
		case egtsSrExtPosData:
			rd.SubrecordData = &EgtsSrExtPosData{}
		case egtsSrAdSensorsData:
			rd.SubrecordData = &EgtsSrAdSensorsData{}
		case egtsSrStateData:
			rd.SubrecordData = &EgtsSrStateData{}
		case egtsSrLiquidLevelSensor:
			rd.SubrecordData = &EgtsSrLiquidLevelSensor{}
		case egtsSrAbsCntrData:
			rd.SubrecordData = &EgtsSrAbsCntrData{}
		case egtsSrAuthInfo:
			rd.SubrecordData = &EgtsSrAuthInfo{}
		case egtsSrEgtsPlusData:
			rd.SubrecordData = &StorageRecord{}
		default:
			return fmt.Errorf("Не известный пакета: %d", rd.SubrecordType)
		}

		if err = rd.SubrecordData.Decode(subRecordBytes); err != nil {
			return err
		}
		*rds = append(*rds, rd)
	}

	return err
}

func (rds *RecordDataSet) Encode() ([]byte, error) {
	var (
		result []byte
		err    error
	)
	buf := new(bytes.Buffer)

	for _, rd := range *rds {
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
	}

	result = buf.Bytes()

	return result, err
}
