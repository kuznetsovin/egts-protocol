package egts

import (
	"encoding/binary"
)

//SrAbsAnSensData структура подзаписи типа EGTS_SR_ABS_AN_SENS_DATA, которая применяется абонентским
//терминалом для передачи данных о состоянии одного аналогового входа
type SrAbsAnSensData struct {
	ASN int `json:"ASN"`
	ASV int `json:"ASV"`
}

//Decode разбирает байты в структуру подзаписи
func (e *SrAbsAnSensData) Decode(content []byte) error {
	data := binary.BigEndian.Uint32(content)
	e.ASN = int(data >> 24)
	e.ASV = int(data & 0x00ffffff)
	return nil
}

//Encode преобразовывает подзапись в набор байт
func (e *SrAbsAnSensData) Encode() ([]byte, error) {
	content := make([]byte, 4)
	binary.BigEndian.PutUint32(content, uint32(e.ASN<<24|e.ASV))
	return content, nil
}

//Length получает длинну закодированной подзаписи
func (e *SrAbsAnSensData) Length() uint16 {
	return 4
}
