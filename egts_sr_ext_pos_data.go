package egts

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

//SrExtPosData структура подзаписи типа EGTS_SR_EXT_POS_DATA, которая используется абонентским
//терминалом при передаче дополнительных данных определения местоположения
type SrExtPosData struct {
	NavigationSystemFieldExists   string `json:"NSFE"`
	SatellitesFieldExists         string `json:"SFE"`
	PdopFieldExists               string `json:"PFE"`
	HdopFieldExists               string `json:"HFE"`
	VdopFieldExists               string `json:"VFE"`
	VerticalDilutionOfPrecision   uint16 `json:"VDOP"`
	HorizontalDilutionOfPrecision uint16 `json:"HDOP"`
	PositionDilutionOfPrecision   uint16 `json:"PDOP"`
	Satellites                    uint8  `json:"SAT"`
	NavigationSystem              uint16 `json:"NS"`
}

//Decode разбирает байты в структуру подзаписи
func (e *SrExtPosData) Decode(content []byte) error {
	var (
		err   error
		flags byte
	)
	tmpBuf := make([]byte, 2)
	buf := bytes.NewReader(content)

	//байт флагов
	if flags, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить байт флагов ext_pos_data: %v", err)
	}
	flagBits := fmt.Sprintf("%08b", flags)

	e.NavigationSystemFieldExists = flagBits[3:4]
	e.SatellitesFieldExists = flagBits[4:5]
	e.PdopFieldExists = flagBits[5:6]
	e.HdopFieldExists = flagBits[6:7]
	e.VdopFieldExists = flagBits[7:]

	if e.VdopFieldExists == "1" {
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить снижение точности в вертикальной плоскости: %v", err)
		}
		e.VerticalDilutionOfPrecision = binary.LittleEndian.Uint16(tmpBuf)
	}

	if e.HdopFieldExists == "1" {
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить снижение точности в горизонтальной плоскости: %v", err)
		}
		e.HorizontalDilutionOfPrecision = binary.LittleEndian.Uint16(tmpBuf)
	}

	if e.PdopFieldExists == "1" {
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить снижение точности по местоположению: %v", err)
		}
		e.PositionDilutionOfPrecision = binary.LittleEndian.Uint16(tmpBuf)
	}

	if e.SatellitesFieldExists == "1" {
		if e.Satellites, err = buf.ReadByte(); err != nil {
			return fmt.Errorf("Не удалось получить количество видимых спутников: %v", err)
		}
	}

	if e.NavigationSystemFieldExists == "1" {
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить битовые флаги спутниковых систем: %v", err)
		}
		e.NavigationSystem = binary.LittleEndian.Uint16(tmpBuf)
	}

	return err
}

//Encode преобразовывает подзапись в набор байт
func (e *SrExtPosData) Encode() ([]byte, error) {
	var (
		err    error
		flags  uint64
		result []byte
	)

	buf := new(bytes.Buffer)

	//байт флагов
	flagsBits := "000" + e.NavigationSystemFieldExists + e.SatellitesFieldExists +
		e.PdopFieldExists + e.HdopFieldExists + e.VdopFieldExists
	if flags, err = strconv.ParseUint(flagsBits, 2, 8); err != nil {
		return result, fmt.Errorf("Не удалось сгенерировать байт флагов ext_pos_data: %v", err)
	}

	if err = buf.WriteByte(uint8(flags)); err != nil {
		return result, fmt.Errorf("Не удалось записать байт флагов ext_pos_data: %v", err)
	}

	if e.VdopFieldExists == "1" {
		if err = binary.Write(buf, binary.LittleEndian, e.VerticalDilutionOfPrecision); err != nil {
			return result, fmt.Errorf("Не удалось записать снижение точности в вертикальной плоскости: %v", err)
		}
	}

	if e.HdopFieldExists == "1" {
		if err = binary.Write(buf, binary.LittleEndian, e.HorizontalDilutionOfPrecision); err != nil {
			return result, fmt.Errorf("Не удалось записать снижение точности в горизонтальной плоскости: %v", err)
		}
	}

	if e.PdopFieldExists == "1" {
		if err = binary.Write(buf, binary.LittleEndian, e.PositionDilutionOfPrecision); err != nil {
			return result, fmt.Errorf("Не удалось записать снижение точности по местоположению: %v", err)
		}
	}

	if e.SatellitesFieldExists == "1" {
		if err = buf.WriteByte(e.Satellites); err != nil {
			return result, fmt.Errorf("Не удалось записать количество видимых спутников: %v", err)
		}
	}

	if e.NavigationSystemFieldExists == "1" {
		if err = binary.Write(buf, binary.LittleEndian, e.NavigationSystem); err != nil {
			return result, fmt.Errorf("Не удалось записать битовые флаги спутниковых систем: %v", err)
		}
	}

	result = buf.Bytes()
	return result, err
}

//Length получает длинну закодированной подзаписи
func (e *SrExtPosData) Length() uint16 {
	var result uint16

	if recBytes, err := e.Encode(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
