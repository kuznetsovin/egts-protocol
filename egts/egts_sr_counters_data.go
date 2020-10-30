package egts

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

//SrCountersData структура подзаписи типа EGTS_SR_COUNTERS_DATA, которая используется аппаратно-программным
//комплексом для передачи на абонентский терминал данных о значении счетных входов
type SrCountersData struct {
	CounterFieldExists1 string `json:"CFE1"`
	CounterFieldExists2 string `json:"CFE2"`
	CounterFieldExists3 string `json:"CFE3"`
	CounterFieldExists4 string `json:"CFE4"`
	CounterFieldExists5 string `json:"CFE5"`
	CounterFieldExists6 string `json:"CFE6"`
	CounterFieldExists7 string `json:"CFE7"`
	CounterFieldExists8 string `json:"CFE8"`
	Counter1            uint32 `json:"CN1"`
	Counter2            uint32 `json:"CN2"`
	Counter3            uint32 `json:"CN3"`
	Counter4            uint32 `json:"CN4"`
	Counter5            uint32 `json:"CN5"`
	Counter6            uint32 `json:"CN6"`
	Counter7            uint32 `json:"CN7"`
	Counter8            uint32 `json:"CN8"`
}

//Decode разбирает байты в структуру подзаписи
func (c *SrCountersData) Decode(content []byte) error {
	var (
		err        error
		flags      byte
		counterVal []byte
	)
	buf := bytes.NewReader(content)

	if flags, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить байт цифровых выходов sr_counters_data: %v", err)
	}
	flagBits := fmt.Sprintf("%08b", flags)

	c.CounterFieldExists8 = flagBits[:1]
	c.CounterFieldExists7 = flagBits[1:2]
	c.CounterFieldExists6 = flagBits[2:3]
	c.CounterFieldExists5 = flagBits[3:4]
	c.CounterFieldExists4 = flagBits[4:5]
	c.CounterFieldExists3 = flagBits[5:6]
	c.CounterFieldExists2 = flagBits[6:7]
	c.CounterFieldExists1 = flagBits[7:]

	tmpBuf := make([]byte, 3)
	if c.CounterFieldExists1 == "1" {
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить показания CN1: %v", err)
		}
		counterVal = append(tmpBuf, 0x00)
		c.Counter1 = binary.LittleEndian.Uint32(counterVal)
	}

	if c.CounterFieldExists2 == "1" {
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить показания CN2: %v", err)
		}
		counterVal = append(tmpBuf, 0x00)
		c.Counter1 = binary.LittleEndian.Uint32(counterVal)
	}

	if c.CounterFieldExists3 == "1" {
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить показания CN3: %v", err)
		}
		counterVal = append(tmpBuf, 0x00)
		c.Counter3 = binary.LittleEndian.Uint32(counterVal)
	}

	if c.CounterFieldExists4 == "1" {
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить показания CN4: %v", err)
		}
		counterVal = append(tmpBuf, 0x00)
		c.Counter4 = binary.LittleEndian.Uint32(counterVal)
	}

	if c.CounterFieldExists5 == "1" {
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить показания CN5: %v", err)
		}
		counterVal = append(tmpBuf, 0x00)
		c.Counter5 = binary.LittleEndian.Uint32(counterVal)
	}

	if c.CounterFieldExists6 == "1" {
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить показания CN6: %v", err)
		}
		counterVal = append(tmpBuf, 0x00)
		c.Counter6 = binary.LittleEndian.Uint32(counterVal)
	}

	if c.CounterFieldExists7 == "1" {
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить показания CN7: %v", err)
		}
		counterVal = append(tmpBuf, 0x00)
		c.Counter7 = binary.LittleEndian.Uint32(counterVal)
	}

	if c.CounterFieldExists8 == "1" {
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить показания CN8: %v", err)
		}
		counterVal = append(tmpBuf, 0x00)
		c.Counter8 = binary.LittleEndian.Uint32(counterVal)
	}
	return err
}

//Encode преобразовывает подзапись в набор байт
func (c *SrCountersData) Encode() ([]byte, error) {
	var (
		err    error
		flags  uint64
		result []byte
	)
	buf := new(bytes.Buffer)
	flagsBits := c.CounterFieldExists8 +
		c.CounterFieldExists7 +
		c.CounterFieldExists6 +
		c.CounterFieldExists5 +
		c.CounterFieldExists4 +
		c.CounterFieldExists3 +
		c.CounterFieldExists2 +
		c.CounterFieldExists1

	if flags, err = strconv.ParseUint(flagsBits, 2, 8); err != nil {
		return result, fmt.Errorf("Не удалось сгенерировать байт байт аналоговых выходов counters_data: %v", err)
	}
	if err = buf.WriteByte(uint8(flags)); err != nil {
		return result, fmt.Errorf("Не удалось записать байт байт аналоговых выходов counters_data: %v", err)
	}

	sensVal := make([]byte, 4)
	if c.CounterFieldExists1 == "1" {
		binary.LittleEndian.PutUint32(sensVal, c.Counter1)
		if _, err = buf.Write(sensVal[:3]); err != nil {
			return result, fmt.Errorf("Не удалось запистаь показания CN1: %v", err)
		}
	}

	if c.CounterFieldExists2 == "1" {
		binary.LittleEndian.PutUint32(sensVal, c.Counter2)
		if _, err = buf.Write(sensVal[:3]); err != nil {
			return result, fmt.Errorf("Не удалось запистаь показания CN2: %v", err)
		}
	}

	if c.CounterFieldExists3 == "1" {
		binary.LittleEndian.PutUint32(sensVal, c.Counter3)
		if _, err = buf.Write(sensVal[:3]); err != nil {
			return result, fmt.Errorf("Не удалось запистаь показания CN3: %v", err)
		}
	}

	if c.CounterFieldExists4 == "1" {
		binary.LittleEndian.PutUint32(sensVal, c.Counter4)
		if _, err = buf.Write(sensVal[:3]); err != nil {
			return result, fmt.Errorf("Не удалось запистаь показания CN4: %v", err)
		}
	}

	if c.CounterFieldExists5 == "1" {
		binary.LittleEndian.PutUint32(sensVal, c.Counter5)
		if _, err = buf.Write(sensVal[:3]); err != nil {
			return result, fmt.Errorf("Не удалось запистаь показания CN5: %v", err)
		}
	}

	if c.CounterFieldExists6 == "1" {
		binary.LittleEndian.PutUint32(sensVal, c.Counter6)
		if _, err = buf.Write(sensVal[:3]); err != nil {
			return result, fmt.Errorf("Не удалось запистаь показания CN6: %v", err)
		}
	}

	if c.CounterFieldExists7 == "1" {
		binary.LittleEndian.PutUint32(sensVal, c.Counter7)
		if _, err = buf.Write(sensVal[:3]); err != nil {
			return result, fmt.Errorf("Не удалось запистаь показания CN7: %v", err)
		}
	}

	if c.CounterFieldExists8 == "1" {
		binary.LittleEndian.PutUint32(sensVal, c.Counter8)
		if _, err = buf.Write(sensVal[:3]); err != nil {
			return result, fmt.Errorf("Не удалось запистаь показания CN8: %v", err)
		}
	}

	result = buf.Bytes()

	return result, err
}

//Length получает длинну закодированной подзаписи
func (c *SrCountersData) Length() uint16 {
	var result uint16

	if recBytes, err := c.Encode(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
