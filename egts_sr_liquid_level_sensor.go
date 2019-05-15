package egts

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

//SrLiquidLevelSensor структура подзаписи типа EGTS_SR_LIQUID_LEVEL_SENSOR, которая применяется
//абонентским терминалом для передачи на аппаратно-программный комплекс данных о показаниях ДУЖ
type SrLiquidLevelSensor struct {
	LiquidLevelSensorErrorFlag string `json:"LLSEF"`
	LiquidLevelSensorValueUnit string `json:"LLSVU"`
	RawDataFlag                string `json:"RDF"`
	LiquidLevelSensorNumber    uint8  `json:"LLSN"`
	ModuleAddress              uint16 `json:"MADDR"`
	LiquidLevelSensorData      uint32 `json:"LLSD"`
}

// Decode разбирает байты в структуру подзаписи
func (e *SrLiquidLevelSensor) Decode(content []byte) error {
	var (
		err     error
		flags   byte
		sensNum uint64
	)
	buf := bytes.NewReader(content)

	if flags, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить байт флагов liquid_level: %v", err)
	}
	flagBits := fmt.Sprintf("%08b", flags)

	e.LiquidLevelSensorErrorFlag = flagBits[1:2]
	e.LiquidLevelSensorValueUnit = flagBits[2:4]
	e.RawDataFlag = flagBits[4:5]

	if sensNum, err = strconv.ParseUint(flagBits[5:], 2, 8); err != nil {
		return fmt.Errorf("Не удалось получить номер датчика ДУЖ: %v", err)
	}
	e.LiquidLevelSensorNumber = uint8(sensNum)

	bytesTmpBuf := make([]byte, 2)
	if _, err = buf.Read(bytesTmpBuf); err != nil {
		return fmt.Errorf("Не удалось получить адрес модуля ДУЖ: %v", err)
	}
	e.ModuleAddress = binary.LittleEndian.Uint16(bytesTmpBuf)

	bytesTmpBuf = make([]byte, 4)
	if _, err = buf.Read(bytesTmpBuf); err != nil {
		return fmt.Errorf("Не удалось получить показания ДУЖ: %v", err)
	}
	e.LiquidLevelSensorData = binary.LittleEndian.Uint32(bytesTmpBuf)

	return err
}

// Encode преобразовывает подзапись в набор байт
func (e *SrLiquidLevelSensor) Encode() ([]byte, error) {
	var (
		err    error
		flags  uint64
		result []byte
	)
	buf := new(bytes.Buffer)

	flagsBits := "0" + e.LiquidLevelSensorErrorFlag + e.LiquidLevelSensorValueUnit +
		e.RawDataFlag + fmt.Sprintf("%03b", e.LiquidLevelSensorNumber)
	if flags, err = strconv.ParseUint(flagsBits, 2, 8); err != nil {
		return result, fmt.Errorf("Не удалось сгенерировать байт флагов ext_pos_data: %v", err)
	}

	if err = buf.WriteByte(uint8(flags)); err != nil {
		return result, fmt.Errorf("Не удалось записать байт флагов ext_pos_data: %v", err)
	}

	if err = binary.Write(buf, binary.LittleEndian, e.ModuleAddress); err != nil {
		return result, fmt.Errorf("Не удалось записать адрес модуля ДУЖ: %v", err)
	}

	if err = binary.Write(buf, binary.LittleEndian, e.LiquidLevelSensorData); err != nil {
		return result, fmt.Errorf("Не удалось записать показания ДУЖ: %v", err)
	}

	result = buf.Bytes()

	return result, err
}

//Length получает длинну закодированной подзаписи
func (e *SrLiquidLevelSensor) Length() uint16 {
	var result uint16

	if recBytes, err := e.Encode(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
