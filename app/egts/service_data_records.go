package egts

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

//ServiceDataRecord запись содержащая мониторинговую информацию
type ServiceDataRecord struct {
	RecordLength             uint16 `json:"RL"`
	RecordNumber             uint16 `json:"RN"`
	SourceServiceOnDevice    string `json:"SSOD"`
	RecipientServiceOnDevice string `json:"RSOD"`
	Group                    string `json:"GRP"`
	RecordProcessingPriority string `json:"RPP"`
	TimeFieldExists          string `json:"TMFE"`
	EventIDFieldExists       string `json:"EVFE"`
	ObjectIDFieldExists      string `json:"OBFE"`
	ObjectIdentifier         uint32 `json:"OID"`
	EventIdentifier          uint32 `json:"EVID"`
	Time                     uint32 `json:"TM"` //TODO: сделать time также как в pos_data
	SourceServiceType        byte   `json:"SST"`
	RecipientServiceType     byte   `json:"RST"`
	RecordDataSet            `json:"RD"`
}

//ServiceDataSet набор последовательных записей с информаций
type ServiceDataSet []ServiceDataRecord

//Decode разбирает байты в структуру подзаписи
func (s *ServiceDataSet) Decode(serviceDS []byte) error {
	var (
		err   error
		flags byte
	)
	buf := bytes.NewReader(serviceDS)

	for buf.Len() > 0 {
		sdr := ServiceDataRecord{}
		tmpIntBuf := make([]byte, 2)
		if _, err = buf.Read(tmpIntBuf); err != nil {
			return fmt.Errorf("Не удалось получить длину записи SDR: %v", err)
		}
		sdr.RecordLength = binary.LittleEndian.Uint16(tmpIntBuf)

		if _, err = buf.Read(tmpIntBuf); err != nil {
			return fmt.Errorf("Не удалось получить номер записи SDR: %v", err)
		}
		sdr.RecordNumber = binary.LittleEndian.Uint16(tmpIntBuf)

		if flags, err = buf.ReadByte(); err != nil {
			return fmt.Errorf("Не удалось считать байт флагов SDR: %v", err)
		}
		flagBits := fmt.Sprintf("%08b", flags)
		sdr.SourceServiceOnDevice = flagBits[:1]
		sdr.RecipientServiceOnDevice = flagBits[1:2]
		sdr.Group = flagBits[2:3]
		sdr.RecordProcessingPriority = flagBits[3:5]
		sdr.TimeFieldExists = flagBits[5:6]
		sdr.EventIDFieldExists = flagBits[6:7]
		sdr.ObjectIDFieldExists = flagBits[7:]

		if sdr.ObjectIDFieldExists == "1" {
			oid := make([]byte, 4)
			if _, err := buf.Read(oid); err != nil {
				return fmt.Errorf("Не удалось получить идентификатор объекта SDR: %v", err)
			}
			sdr.ObjectIdentifier = binary.LittleEndian.Uint32(oid)
		}

		if sdr.EventIDFieldExists == "1" {
			event := make([]byte, 4)
			if _, err := buf.Read(event); err != nil {
				return fmt.Errorf("Не удалось получить идентификатор события SDR: %v", err)
			}
			sdr.EventIdentifier = binary.LittleEndian.Uint32(event)
		}

		if sdr.TimeFieldExists == "1" {
			tm := make([]byte, 4)
			if _, err := buf.Read(tm); err != nil {
				return fmt.Errorf("Не удалось получить время формирования записи на стороне отправителя SDR: %v", err)
			}
			sdr.Time = binary.LittleEndian.Uint32(tm)
		}

		if sdr.SourceServiceType, err = buf.ReadByte(); err != nil {
			return fmt.Errorf("Не удалось считать идентификатор тип сервиса-отправителя SDR: %v", err)
		}

		if sdr.RecipientServiceType, err = buf.ReadByte(); err != nil {
			return fmt.Errorf("Не удалось считать идентификатор тип сервиса-получателя SDR: %v", err)
		}

		if buf.Len() != 0 {
			rds := RecordDataSet{}
			rdsBytes := make([]byte, sdr.RecordLength)
			if _, err = buf.Read(rdsBytes); err != nil {
				return err
			}

			if err = rds.Decode(rdsBytes); err != nil {
				return err
			}
			sdr.RecordDataSet = rds
		}

		*s = append(*s, sdr)
	}
	return err
}

//Encode кодирование структуры в байты
func (s *ServiceDataSet) Encode() ([]byte, error) {
	var (
		result []byte
		flags  uint64
	)

	buf := new(bytes.Buffer)

	for _, sdr := range *s {
		rd, err := sdr.RecordDataSet.Encode()
		if err != nil {
			return result, err
		}

		if sdr.RecordLength == 0 {
			sdr.RecordLength = uint16(len(rd))
		}
		if err = binary.Write(buf, binary.LittleEndian, sdr.RecordLength); err != nil {
			return result, fmt.Errorf("Не удалось записать длину записи SDR: %v", err)
		}

		if err = binary.Write(buf, binary.LittleEndian, sdr.RecordNumber); err != nil {
			return result, fmt.Errorf("Не удалось записать номер записи SDR: %v", err)
		}

		//составной байт
		flagsBits := sdr.SourceServiceOnDevice + sdr.RecipientServiceOnDevice + sdr.Group + sdr.RecordProcessingPriority +
			sdr.TimeFieldExists + sdr.EventIDFieldExists + sdr.ObjectIDFieldExists
		if flags, err = strconv.ParseUint(flagsBits, 2, 8); err != nil {
			return result, fmt.Errorf("Не удалось сгенерировать байт флагов SDR: %v", err)
		}
		if err = buf.WriteByte(uint8(flags)); err != nil {
			return result, fmt.Errorf("Не удалось записать флаги SDR: %v", err)
		}

		if sdr.ObjectIDFieldExists == "1" {
			if err = binary.Write(buf, binary.LittleEndian, sdr.ObjectIdentifier); err != nil {
				return result, fmt.Errorf("Не удалось записать идентификатор объекта SDR: %v", err)
			}
		}

		if sdr.EventIDFieldExists == "1" {
			if err = binary.Write(buf, binary.LittleEndian, sdr.EventIdentifier); err != nil {
				return result, fmt.Errorf("Не удалось записать идентификатор события SDR: %v", err)
			}
		}

		if sdr.TimeFieldExists == "1" {
			if err := binary.Write(buf, binary.LittleEndian, sdr.Time); err != nil {
				return result, fmt.Errorf("Не удалось записать время формирования записи на стороне отправителя SDR: %v", err)
			}
		}

		if err := buf.WriteByte(sdr.SourceServiceType); err != nil {
			return result, fmt.Errorf("Не удалось записать идентификатор тип сервиса-отправителя SDR: %v", err)
		}

		if err := buf.WriteByte(sdr.RecipientServiceType); err != nil {
			return result, fmt.Errorf("Не удалось записать идентификатор тип сервиса-получателя SDR: %v", err)
		}

		buf.Write(rd)
	}

	result = buf.Bytes()

	return result, nil
}

//Length получает длину массива записей
func (s *ServiceDataSet) Length() uint16 {
	var result uint16

	if recBytes, err := s.Encode(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
