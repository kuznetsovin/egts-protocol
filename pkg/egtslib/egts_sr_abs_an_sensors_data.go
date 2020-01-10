package egts

import (
	"encoding/binary"
	"errors"
)

//SrAbsAnSensData структура подзаписи типа EGTS_SR_ABS_AN_SENS_DATA, которая применяется абонентским
//терминалом для передачи данных о состоянии одного аналогового входа
type SrAbsAnSensData struct {
	SensorNumber uint8  `json:"SensorNumber"`
	Value        uint32 `json:"Value"`
}

//Decode разбирает байты в структуру подзаписи
func (e *SrAbsAnSensData) Decode(content []byte) error {
	if len(content) < int(e.Length()) {
		return errors.New("Некорректный размер данных")
	}
	e.SensorNumber = uint8(content[0])
	e.Value = uint32(binary.LittleEndian.Uint32(content) >> 8)
	return nil
}

//Encode преобразовывает подзапись в набор байт
func (e *SrAbsAnSensData) Encode() ([]byte, error) {
	return []byte{
		byte(e.SensorNumber),
		byte(e.Value),
		byte(e.Value >> 8),
		byte(e.Value >> 16),
	}, nil
}

//Length получает длинну закодированной подзаписи
func (e *SrAbsAnSensData) Length() uint16 {
	return 4
}
