package egts

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// SrDispatcherIdentity структура подзаписи типа EGTS_SR_DISPATCHER_IDENTITY, которая используется
//только авторизуемой ТП при запросе авторизации на авторизующей ТП и содержит учетные данные
//авторизуемой АСН
type SrDispatcherIdentity struct {
	DispatcherType uint8  `json:"DT"`
	DispatcherID   uint32 `json:"DID"`
	Description    string `json:"DSCR"`
}

// Decode разбирает байты в структуру подзаписи
func (d *SrDispatcherIdentity) Decode(content []byte) error {
	var err error

	buf := bytes.NewBuffer(content)

	if d.DispatcherType, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить тип диспетчера: %v", err)
	}

	tmpIntBuf := make([]byte, 4)
	if _, err = buf.Read(tmpIntBuf); err != nil {
		return fmt.Errorf("Не удалось получить уникальный идентификатор диспетчера: %v", err)
	}
	d.DispatcherID = binary.LittleEndian.Uint32(tmpIntBuf)

	d.Description = buf.String()

	return err
}

// Encode преобразовывает подзапись в набор байт
func (d *SrDispatcherIdentity) Encode() ([]byte, error) {
	var (
		result []byte
		err    error
	)

	buf := new(bytes.Buffer)

	if err = buf.WriteByte(d.DispatcherType); err != nil {
		return result, fmt.Errorf("Не удалось записать тип диспетчера: %v", err)
	}

	if err = binary.Write(buf, binary.LittleEndian, d.DispatcherID); err != nil {
		return result, fmt.Errorf("Не удалось записать уникальный идентификатор диспетчера: %v", err)
	}

	if _, err = buf.WriteString(d.Description); err != nil {
		return result, fmt.Errorf("Не удалось записать уникальный краткое описание: %v", err)
	}

	return buf.Bytes(), err
}

//Length получает длинну закодированной подзаписи
func (d *SrDispatcherIdentity) Length() uint16 {
	var result uint16

	if recBytes, err := d.Encode(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
