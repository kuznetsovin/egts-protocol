package egts

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

//SrAdSensorsData структура подзаписи типа EGTS_SR_AD_SENSORS_DATA, которая применяется абонентским
//терминалом для передачи на аппаратно-программный комплекс информации о состоянии дополнительных
//дискретных и аналоговых входов
type SrAdSensorsData struct {
	DigitalInputsOctetExists1     string `json:"DIOE1"`
	DigitalInputsOctetExists2     string `json:"DIOE2"`
	DigitalInputsOctetExists3     string `json:"DIOE3"`
	DigitalInputsOctetExists4     string `json:"DIOE4"`
	DigitalInputsOctetExists5     string `json:"DIOE5"`
	DigitalInputsOctetExists6     string `json:"DIOE6"`
	DigitalInputsOctetExists7     string `json:"DIOE7"`
	DigitalInputsOctetExists8     string `json:"DIOE8"`
	DigitalOutputs                byte   `json:"DOUT"`
	AnalogSensorFieldExists1      string `json:"ASFE1"`
	AnalogSensorFieldExists2      string `json:"ASFE2"`
	AnalogSensorFieldExists3      string `json:"ASFE3"`
	AnalogSensorFieldExists4      string `json:"ASFE4"`
	AnalogSensorFieldExists5      string `json:"ASFE5"`
	AnalogSensorFieldExists6      string `json:"ASFE6"`
	AnalogSensorFieldExists7      string `json:"ASFE7"`
	AnalogSensorFieldExists8      string `json:"ASFE8"`
	AdditionalDigitalInputsOctet1 byte   `json:"ADIO1"`
	AdditionalDigitalInputsOctet2 byte   `json:"ADIO2"`
	AdditionalDigitalInputsOctet3 byte   `json:"ADIO3"`
	AdditionalDigitalInputsOctet4 byte   `json:"ADIO4"`
	AdditionalDigitalInputsOctet5 byte   `json:"ADIO5"`
	AdditionalDigitalInputsOctet6 byte   `json:"ADIO6"`
	AdditionalDigitalInputsOctet7 byte   `json:"ADIO7"`
	AdditionalDigitalInputsOctet8 byte   `json:"ADIO8"`
	AnalogSensor1                 uint32 `json:"ANS1"`
	AnalogSensor2                 uint32 `json:"ANS2"`
	AnalogSensor3                 uint32 `json:"ANS3"`
	AnalogSensor4                 uint32 `json:"ANS4"`
	AnalogSensor5                 uint32 `json:"ANS5"`
	AnalogSensor6                 uint32 `json:"ANS6"`
	AnalogSensor7                 uint32 `json:"ANS7"`
	AnalogSensor8                 uint32 `json:"ANS8"`
}

//Decode разбирает байты в структуру подзаписи
func (e *SrAdSensorsData) Decode(content []byte) error {
	var (
		err           error
		flags         byte
		analogSensVal []byte
	)
	buf := bytes.NewReader(content)

	//байт флагов
	if flags, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить байт цифровых выходов ad_sesor_data: %v", err)
	}
	flagBits := fmt.Sprintf("%08b", flags)

	e.DigitalInputsOctetExists8 = flagBits[:1]
	e.DigitalInputsOctetExists7 = flagBits[1:2]
	e.DigitalInputsOctetExists6 = flagBits[2:3]
	e.DigitalInputsOctetExists5 = flagBits[3:4]
	e.DigitalInputsOctetExists4 = flagBits[4:5]
	e.DigitalInputsOctetExists3 = flagBits[5:6]
	e.DigitalInputsOctetExists2 = flagBits[6:7]
	e.DigitalInputsOctetExists1 = flagBits[7:]

	if e.DigitalOutputs, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить битовые флаги дискретных выходов: %v", err)
	}

	if flags, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить байт аналоговых выходов ad_sesor_data: %v", err)
	}
	flagBits = fmt.Sprintf("%08b", flags)

	e.AnalogSensorFieldExists8 = flagBits[:1]
	e.AnalogSensorFieldExists7 = flagBits[1:2]
	e.AnalogSensorFieldExists6 = flagBits[2:3]
	e.AnalogSensorFieldExists5 = flagBits[3:4]
	e.AnalogSensorFieldExists4 = flagBits[4:5]
	e.AnalogSensorFieldExists3 = flagBits[5:6]
	e.AnalogSensorFieldExists2 = flagBits[6:7]
	e.AnalogSensorFieldExists1 = flagBits[7:]

	if e.DigitalInputsOctetExists1 == "1" {
		if e.AdditionalDigitalInputsOctet1, err = buf.ReadByte(); err != nil {
			return fmt.Errorf("Не удалось получить байт показания ADIO1: %v", err)
		}
	}

	if e.DigitalInputsOctetExists2 == "1" {
		if e.AdditionalDigitalInputsOctet2, err = buf.ReadByte(); err != nil {
			return fmt.Errorf("Не удалось получить байт показания ADIO2: %v", err)
		}
	}

	if e.DigitalInputsOctetExists3 == "1" {
		if e.AdditionalDigitalInputsOctet3, err = buf.ReadByte(); err != nil {
			return fmt.Errorf("Не удалось получить байт показания ADIO3: %v", err)
		}
	}

	if e.DigitalInputsOctetExists4 == "1" {
		if e.AdditionalDigitalInputsOctet4, err = buf.ReadByte(); err != nil {
			return fmt.Errorf("Не удалось получить байт показания ADIO4: %v", err)
		}
	}

	if e.DigitalInputsOctetExists5 == "1" {
		if e.AdditionalDigitalInputsOctet5, err = buf.ReadByte(); err != nil {
			return fmt.Errorf("Не удалось получить байт показания ADIO5: %v", err)
		}
	}

	if e.DigitalInputsOctetExists6 == "1" {
		if e.AdditionalDigitalInputsOctet6, err = buf.ReadByte(); err != nil {
			return fmt.Errorf("Не удалось получить байт показания ADIO6: %v", err)
		}
	}

	if e.DigitalInputsOctetExists7 == "1" {
		if e.AdditionalDigitalInputsOctet7, err = buf.ReadByte(); err != nil {
			return fmt.Errorf("Не удалось получить байт показания ADIO7: %v", err)
		}
	}

	if e.DigitalInputsOctetExists8 == "1" {
		if e.AdditionalDigitalInputsOctet8, err = buf.ReadByte(); err != nil {
			return fmt.Errorf("Не удалось получить байт показания ADIO8: %v", err)
		}
	}

	tmpBuf := make([]byte, 3)
	if e.AnalogSensorFieldExists1 == "1" {
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить показания ANS1: %v", err)
		}
		analogSensVal = append(tmpBuf, 0x00)
		e.AnalogSensor1 = binary.LittleEndian.Uint32(analogSensVal)
	}

	if e.AnalogSensorFieldExists2 == "1" {
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить показания ANS2: %v", err)
		}
		analogSensVal = append(tmpBuf, 0x00)
		e.AnalogSensor2 = binary.LittleEndian.Uint32(analogSensVal)
	}

	if e.AnalogSensorFieldExists3 == "1" {
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить показания ANS3: %v", err)
		}
		analogSensVal = append(tmpBuf, 0x00)
		e.AnalogSensor3 = binary.LittleEndian.Uint32(analogSensVal)
	}

	if e.AnalogSensorFieldExists4 == "1" {
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить показания ANS4: %v", err)
		}
		analogSensVal = append(tmpBuf, 0x00)
		e.AnalogSensor4 = binary.LittleEndian.Uint32(analogSensVal)
	}

	if e.AnalogSensorFieldExists5 == "1" {
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить показания ANS5: %v", err)
		}
		analogSensVal = append(tmpBuf, 0x00)
		e.AnalogSensor5 = binary.LittleEndian.Uint32(analogSensVal)
	}

	if e.AnalogSensorFieldExists6 == "1" {
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить показания ANS6: %v", err)
		}
		analogSensVal = append(tmpBuf, 0x00)
		e.AnalogSensor6 = binary.LittleEndian.Uint32(analogSensVal)
	}

	if e.AnalogSensorFieldExists7 == "1" {
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить показания ANS7: %v", err)
		}
		analogSensVal = append(tmpBuf, 0x00)
		e.AnalogSensor7 = binary.LittleEndian.Uint32(analogSensVal)
	}

	if e.AnalogSensorFieldExists8 == "1" {
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить показания ANS8: %v", err)
		}
		analogSensVal = append(tmpBuf, 0x00)
		e.AnalogSensor8 = binary.LittleEndian.Uint32(analogSensVal)
	}
	return err
}

//Encode преобразовывает подзапись в набор байт
func (e *SrAdSensorsData) Encode() ([]byte, error) {
	var (
		err    error
		flags  uint64
		result []byte
	)

	buf := new(bytes.Buffer)

	flagsBits := e.DigitalInputsOctetExists8 +
		e.DigitalInputsOctetExists7 +
		e.DigitalInputsOctetExists6 +
		e.DigitalInputsOctetExists5 +
		e.DigitalInputsOctetExists4 +
		e.DigitalInputsOctetExists3 +
		e.DigitalInputsOctetExists2 +
		e.DigitalInputsOctetExists1

	if flags, err = strconv.ParseUint(flagsBits, 2, 8); err != nil {
		return result, fmt.Errorf("Не удалось сгенерировать байт байт цифровых выходов ad_sesor_data: %v", err)
	}

	if err = buf.WriteByte(uint8(flags)); err != nil {
		return result, fmt.Errorf("Не удалось записать байт флагов ext_pos_data: %v", err)
	}

	if err = buf.WriteByte(e.DigitalOutputs); err != nil {
		return result, fmt.Errorf("Не удалось записать битовые флаги дискретных выходов: %v", err)
	}

	flagsBits = e.AnalogSensorFieldExists1 +
		e.AnalogSensorFieldExists2 +
		e.AnalogSensorFieldExists3 +
		e.AnalogSensorFieldExists4 +
		e.AnalogSensorFieldExists5 +
		e.AnalogSensorFieldExists6 +
		e.AnalogSensorFieldExists7 +
		e.AnalogSensorFieldExists8

	if flags, err = strconv.ParseUint(flagsBits, 2, 8); err != nil {
		return result, fmt.Errorf("Не удалось сгенерировать байт байт аналоговых выходов ad_sesor_data: %v", err)
	}
	if err = buf.WriteByte(uint8(flags)); err != nil {
		return result, fmt.Errorf("Не удалось записать байт байт аналоговых выходов ad_sesor_data: %v", err)
	}

	if e.DigitalInputsOctetExists1 == "1" {
		if err = buf.WriteByte(e.AdditionalDigitalInputsOctet1); err != nil {
			return result, fmt.Errorf("Не удалось записать байт показания ADIO1: %v", err)
		}
	}

	if e.DigitalInputsOctetExists2 == "1" {
		if err = buf.WriteByte(e.AdditionalDigitalInputsOctet2); err != nil {
			return result, fmt.Errorf("Не удалось записать байт показания ADIO2: %v", err)
		}
	}

	if e.DigitalInputsOctetExists3 == "1" {
		if err = buf.WriteByte(e.AdditionalDigitalInputsOctet3); err != nil {
			return result, fmt.Errorf("Не удалось записать байт показания ADIO3: %v", err)
		}
	}

	if e.DigitalInputsOctetExists4 == "1" {
		if err = buf.WriteByte(e.AdditionalDigitalInputsOctet4); err != nil {
			return result, fmt.Errorf("Не удалось записать байт показания ADIO4: %v", err)
		}
	}

	if e.DigitalInputsOctetExists5 == "1" {
		if err = buf.WriteByte(e.AdditionalDigitalInputsOctet5); err != nil {
			return result, fmt.Errorf("Не удалось записать байт показания ADIO5: %v", err)
		}
	}

	if e.DigitalInputsOctetExists6 == "1" {
		if err = buf.WriteByte(e.AdditionalDigitalInputsOctet6); err != nil {
			return result, fmt.Errorf("Не удалось записать байт показания ADIO6: %v", err)
		}
	}

	if e.DigitalInputsOctetExists7 == "1" {
		if err = buf.WriteByte(e.AdditionalDigitalInputsOctet7); err != nil {
			return result, fmt.Errorf("Не удалось записать байт показания ADIO7: %v", err)
		}
	}

	if e.DigitalInputsOctetExists8 == "1" {
		if err = buf.WriteByte(e.AdditionalDigitalInputsOctet8); err != nil {
			return result, fmt.Errorf("Не удалось записать байт показания ADIO8: %v", err)
		}
	}

	sensVal := make([]byte, 4)
	if e.AnalogSensorFieldExists1 == "1" {
		binary.LittleEndian.PutUint32(sensVal, e.AnalogSensor1)
		if _, err = buf.Write(sensVal[:3]); err != nil {
			return result, fmt.Errorf("Не удалось запистаь показания ANS1: %v", err)
		}
	}

	if e.AnalogSensorFieldExists2 == "1" {
		binary.LittleEndian.PutUint32(sensVal, e.AnalogSensor2)
		if _, err = buf.Write(sensVal[:3]); err != nil {
			return result, fmt.Errorf("Не удалось запистаь показания ANS2: %v", err)
		}
	}

	if e.AnalogSensorFieldExists3 == "1" {
		binary.LittleEndian.PutUint32(sensVal, e.AnalogSensor3)
		if _, err = buf.Write(sensVal[:3]); err != nil {
			return result, fmt.Errorf("Не удалось запистаь показания ANS3: %v", err)
		}
	}

	if e.AnalogSensorFieldExists4 == "1" {
		binary.LittleEndian.PutUint32(sensVal, e.AnalogSensor4)
		if _, err = buf.Write(sensVal[:3]); err != nil {
			return result, fmt.Errorf("Не удалось запистаь показания ANS4: %v", err)
		}
	}

	if e.AnalogSensorFieldExists5 == "1" {
		binary.LittleEndian.PutUint32(sensVal, e.AnalogSensor5)
		if _, err = buf.Write(sensVal[:3]); err != nil {
			return result, fmt.Errorf("Не удалось запистаь показания ANS5: %v", err)
		}
	}

	if e.AnalogSensorFieldExists6 == "1" {
		binary.LittleEndian.PutUint32(sensVal, e.AnalogSensor6)
		if _, err = buf.Write(sensVal[:3]); err != nil {
			return result, fmt.Errorf("Не удалось запистаь показания ANS6: %v", err)
		}
	}

	if e.AnalogSensorFieldExists7 == "1" {
		binary.LittleEndian.PutUint32(sensVal, e.AnalogSensor7)
		if _, err = buf.Write(sensVal[:3]); err != nil {
			return result, fmt.Errorf("Не удалось запистаь показания ANS7: %v", err)
		}
	}

	if e.AnalogSensorFieldExists8 == "1" {
		binary.LittleEndian.PutUint32(sensVal, e.AnalogSensor8)
		if _, err = buf.Write(sensVal[:3]); err != nil {
			return result, fmt.Errorf("Не удалось запистаь показания ANS8: %v", err)
		}
	}

	result = buf.Bytes()
	return result, err
}

//Length получает длинну закодированной подзаписи
func (e *SrAdSensorsData) Length() uint16 {
	var result uint16

	if recBytes, err := e.Encode(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
