package egts

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	Delimiter       = uint8(0)
	textSectionSize = 10
)

// SrModuleData структура подзапись EGTS_AUTH_SERVICE типа EGTS_SR_MODULE_DATA
type SrModuleData struct {
	ModuleType      int8   `json:"MT"`
	VendorID        uint32 `json:"VID"`
	FirmwareVersion uint16 `json:"FWV"`
	SoftwareVersion uint16 `json:"SWV"`
	Modification    byte   `json:"MD"`
	State           byte   `json:"ST"`
	SerialNumber    string `json:"SRN"`
	_               byte   `json:"-"`
	Description     string `json:"DSCR"`
	_               byte   `json:"-"`
}

// Decode разбирает байты в структуру подзаписи
//nolint:funlen
func (e *SrModuleData) Decode(content []byte) error {
	var err error
	buf := bytes.NewReader(content)

	moduleType, err := buf.ReadByte()
	if err != nil {
		return fmt.Errorf("не удалось получить тип модуля")
	}
	e.ModuleType = int8(moduleType)

	tmpBuf := make([]byte, 4)
	if _, err = buf.Read(tmpBuf); err != nil {
		return fmt.Errorf("не удалось получить ID вендора")
	}
	e.VendorID = binary.LittleEndian.Uint32(tmpBuf)

	tmpBuf = make([]byte, 2)
	if _, err = buf.Read(tmpBuf); err != nil {
		return fmt.Errorf("не удалось получить версию прошивки")
	}
	e.FirmwareVersion = binary.LittleEndian.Uint16(tmpBuf)

	tmpBuf = make([]byte, 2)
	if _, err = buf.Read(tmpBuf); err != nil {
		return fmt.Errorf("не удалось получить версию ПО")
	}
	e.SoftwareVersion = binary.LittleEndian.Uint16(tmpBuf)

	e.Modification, err = buf.ReadByte()
	if err != nil {
		return fmt.Errorf("не удалось получить модификацию")
	}

	e.State, err = buf.ReadByte()
	if err != nil {
		return fmt.Errorf("не удалось получить состояние")
	}

	serialNumber := make([]byte, 0, textSectionSize)
	for {
		b, err := buf.ReadByte()
		if err != nil {
			return fmt.Errorf("не удалось получить серийный номер")
		}
		if b == Delimiter {
			break
		}
		serialNumber = append(serialNumber, b)
	}
	e.SerialNumber = string(serialNumber)

	description := make([]byte, 0, textSectionSize)
	for {
		b, err := buf.ReadByte()
		if err != nil {
			return fmt.Errorf("не удалось получить описание")
		}
		if b == Delimiter {
			break
		}
		description = append(description, b)
	}
	e.Description = string(description)

	return err
}

// Encode преобразовывает подзапись в набор байт
func (e *SrModuleData) Encode() ([]byte, error) {
	var (
		result []byte
		err    error
	)
	buf := new(bytes.Buffer)

	if err = binary.Write(buf, binary.LittleEndian, e.ModuleType); err != nil {
		return result, fmt.Errorf("не удалось записать тип модуля")
	}
	if err = binary.Write(buf, binary.LittleEndian, e.VendorID); err != nil {
		return result, fmt.Errorf("не удалось записать ID вендора")
	}
	if err = binary.Write(buf, binary.LittleEndian, e.FirmwareVersion); err != nil {
		return result, fmt.Errorf("не удалось записать версию прошивки")
	}
	if err = binary.Write(buf, binary.LittleEndian, e.SoftwareVersion); err != nil {
		return result, fmt.Errorf("не удалось записать версию ПО")
	}
	if err = binary.Write(buf, binary.LittleEndian, e.Modification); err != nil {
		return result, fmt.Errorf("не удалось записать модификацию")
	}
	if err = binary.Write(buf, binary.LittleEndian, e.State); err != nil {
		return result, fmt.Errorf("не удалось записать состояние")
	}
	if err = binary.Write(buf, binary.LittleEndian, []byte(e.SerialNumber)); err != nil {
		return result, fmt.Errorf("не удалось записать серийный номер")
	}
	if err = binary.Write(buf, binary.LittleEndian, Delimiter); err != nil {
		return result, fmt.Errorf("не удалось записать раделитель")
	}
	if err = binary.Write(buf, binary.LittleEndian, []byte(e.Description)); err != nil {
		return result, fmt.Errorf("не удалось записать описание")
	}
	if err = binary.Write(buf, binary.LittleEndian, Delimiter); err != nil {
		return result, fmt.Errorf("не удалось записать разделитель")
	}

	return buf.Bytes(), err
}

// Length получает длину закодированной подзаписи
func (e *SrModuleData) Length() uint16 {
	var result uint16

	if recBytes, err := e.Encode(); err == nil {
		result = uint16(len(recBytes))
	}

	return result
}
