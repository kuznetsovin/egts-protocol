package egts

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//SrResponse структура подзаписи типа EGTS_SR_RESPONSE, которая применяется для подтверждения
//приема результатов обработки поддержки услуг
type SrResponse struct {
	ConfirmedRecordNumber uint16 `json:"CRN"`
	RecordStatus          uint8  `json:"RST"`
}

//Decode разбирает байты в структуру подзаписи
func (s *SrResponse) Decode(content []byte) error {
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

//Encode преобразовывает подзапись в набор байт
func (s *SrResponse) Encode() ([]byte, error) {
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

//Length получает длинну закодированной подзаписи
func (s *SrResponse) Length() uint16 {
	var result uint16

	if recBytes, err := s.Encode(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
