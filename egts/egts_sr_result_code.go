package egts

import (
	"bytes"
	"fmt"
)

//SrResultCode структура подзаписи типа EGTS_SR_RESULT_CODE, которая применяется телематической
//платформой для информирования АС о результатах процедуры аутентификации АС.
type SrResultCode struct {
	ResultCode uint8 `json:"RCD"`
}

//Decode разбирает байты в структуру подзаписи
func (s *SrResultCode) Decode(content []byte) error {
	var (
		err error
	)
	buf := bytes.NewBuffer(content)

	if s.ResultCode, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить код результата: %v", err)
	}

	return err
}

//Encode преобразовывает подзапись в набор байт
func (s *SrResultCode) Encode() ([]byte, error) {
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

//Length получает длинну закодированной подзаписи
func (s *SrResultCode) Length() uint16 {
	var result uint16

	if recBytes, err := s.Encode(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
