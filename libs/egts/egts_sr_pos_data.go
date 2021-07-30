package egts

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"time"
)

//SrPosData структура подзаписи типа EGTS_SR_POS_DATA, которая используется абонентским
//терминалом при передаче основных данных определения местоположения
type SrPosData struct {
	NavigationTime      time.Time `json:"NTM"`
	Latitude            float64   `json:"LAT"`
	Longitude           float64   `json:"LONG"`
	ALTE                string    `json:"ALTE"`
	LOHS                string    `json:"LOHS"`
	LAHS                string    `json:"LAHS"`
	MV                  string    `json:"MV"`
	BB                  string    `json:"BB"`
	CS                  string    `json:"CS"`
	FIX                 string    `json:"FIX"`
	VLD                 string    `json:"VLD"`
	DirectionHighestBit uint8     `json:"DIRH"`
	AltitudeSign        uint8     `json:"ALTS"`
	Speed               uint16    `json:"SPD"`
	Direction           byte      `json:"DIR"`
	Odometer            uint32    `json:"ODM"`
	DigitalInputs       byte      `json:"DIN"`
	Source              byte      `json:"SRC"`
	Altitude            uint32    `json:"ALT"`
	SourceData          int16     `json:"SRCD"`
}

//Decode разбирает байты в структуру подзаписи
func (e *SrPosData) Decode(content []byte) error {
	var (
		err   error
		flags byte
		speed uint64
	)
	buf := bytes.NewReader(content)

	// Преобразуем время навигации к формату, который требует стандарт: количество секунд с 00:00:00 01.01.2010 UTC
	startDate := time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC)
	tmpUint32Buf := make([]byte, 4)
	if _, err = buf.Read(tmpUint32Buf); err != nil {
		return fmt.Errorf("Не удалось получить время навигации: %v", err)
	}
	preFieldVal := binary.LittleEndian.Uint32(tmpUint32Buf)
	e.NavigationTime = startDate.Add(time.Duration(preFieldVal) * time.Second)

	// В протоколе значение хранится в виде: широта по модулю, градусы/90*0xFFFFFFFF  и взята целая часть
	if _, err = buf.Read(tmpUint32Buf); err != nil {
		return fmt.Errorf("Не удалось получить широту: %v", err)
	}

	preFieldVal = binary.LittleEndian.Uint32(tmpUint32Buf)
	e.Latitude = float64(float64(preFieldVal) * 90 / 0xFFFFFFFF)

	// В протоколе значение хранится в виде: долгота по модулю, градусы/180*0xFFFFFFFF  и взята целая часть
	if _, err = buf.Read(tmpUint32Buf); err != nil {
		return fmt.Errorf("Не удалось получить время долгату: %v", err)
	}
	preFieldVal = binary.LittleEndian.Uint32(tmpUint32Buf)
	e.Longitude = float64(float64(preFieldVal) * 180 / 0xFFFFFFFF)

	//байт флагов
	if flags, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить байт флагов pos_data: %v", err)
	}
	flagBits := fmt.Sprintf("%08b", flags)
	e.ALTE = flagBits[:1]
	e.LOHS = flagBits[1:2]
	e.LAHS = flagBits[2:3]
	e.MV = flagBits[3:4]
	e.BB = flagBits[4:5]
	e.CS = flagBits[5:6]
	e.FIX = flagBits[6:7]
	e.VLD = flagBits[7:]

	// скорость
	tmpUint16Buf := make([]byte, 2)
	if _, err = buf.Read(tmpUint16Buf); err != nil {
		return fmt.Errorf("Не удалось получить скорость: %v", err)
	}
	spd := binary.LittleEndian.Uint16(tmpUint16Buf)
	e.DirectionHighestBit = uint8(spd >> 15 & 0x1)
	e.AltitudeSign = uint8(spd >> 14 & 0x1)

	speedBits := fmt.Sprintf("%016b", spd)
	if speed, err = strconv.ParseUint(speedBits[2:], 2, 16); err != nil {
		return fmt.Errorf("Не удалось расшифровать скорость из битов: %v", err)
	}

	// т.к. скорость с дискретностью 0,1 км
	e.Speed = uint16(speed) / 10

	if e.Direction, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить направление движения: %v", err)
	}
	e.Direction |= e.DirectionHighestBit << 7

	bytesTmpBuf := make([]byte, 3)
	if _, err = buf.Read(bytesTmpBuf); err != nil {
		return fmt.Errorf("Не удалось получить пройденное расстояние (пробег) в км: %v", err)
	}
	bytesTmpBuf = append(bytesTmpBuf, 0x00)
	e.Odometer = binary.LittleEndian.Uint32(bytesTmpBuf)

	if e.DigitalInputs, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить битовые флаги, определяют состояние основных дискретных входов: %v", err)
	}

	if e.Source, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить источник (событие), инициировавший посылку: %v", err)
	}

	if e.ALTE == "1" {
		bytesTmpBuf = []byte{0, 0, 0, 0}
		if _, err = buf.Read(bytesTmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить высоту над уровнем моря: %v", err)
		}
		e.Altitude = binary.LittleEndian.Uint32(bytesTmpBuf)
	}

	//TODO: разобраться с разбором SourceData
	return err
}

//Encode преобразовывает подзапись в набор байт
func (e *SrPosData) Encode() ([]byte, error) {
	var (
		err    error
		flags  uint64
		result []byte
	)

	buf := new(bytes.Buffer)
	// Преобразуем время навигации к формату, который требует стандарт: количество секунд с 00:00:00 01.01.2010 UTC
	startDate := time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC)
	if err = binary.Write(buf, binary.LittleEndian, uint32(e.NavigationTime.Sub(startDate).Seconds())); err != nil {
		return result, fmt.Errorf("Не удалось записать время навигации: %v", err)
	}

	// В протоколе значение хранится в виде: широта по модулю, градусы/90*0xFFFFFFFF  и взята целая часть
	if err = binary.Write(buf, binary.LittleEndian, uint32(e.Latitude/90*0xFFFFFFFF)); err != nil {
		return result, fmt.Errorf("Не удалось записать широту: %v", err)
	}

	// В протоколе значение хранится в виде: долгота по модулю, градусы/180*0xFFFFFFFF  и взята целая часть
	if err = binary.Write(buf, binary.LittleEndian, uint32(e.Longitude/180*0xFFFFFFFF)); err != nil {
		return result, fmt.Errorf("Не удалось записать долготу: %v", err)
	}

	//байт флагов
	flags, err = strconv.ParseUint(e.ALTE+e.LOHS+e.LAHS+e.MV+e.BB+e.CS+e.FIX+e.VLD, 2, 8)
	if err != nil {
		return result, fmt.Errorf("Не удалось сгенерировать байт флагов pos_data: %v", err)
	}

	if err = buf.WriteByte(uint8(flags)); err != nil {
		return result, fmt.Errorf("Не удалось записать флаги: %v", err)
	}

	// скорость
	speed := e.Speed*10 | uint16(e.DirectionHighestBit)<<15 // 15 бит
	speed = speed | uint16(e.AltitudeSign)<<14              //14 бит
	spd := make([]byte, 2)
	binary.LittleEndian.PutUint16(spd, speed)
	if _, err = buf.Write(spd); err != nil {
		return result, fmt.Errorf("Не удалось записать скорость: %v", err)
	}

	dir := e.Direction &^ (e.DirectionHighestBit << 7)
	if err = binary.Write(buf, binary.LittleEndian, dir); err != nil {
		return result, fmt.Errorf("Не удалось записать направление движения: %v", err)
	}

	bytesTmpBuf := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytesTmpBuf, e.Odometer)
	if _, err = buf.Write(bytesTmpBuf[:3]); err != nil {
		return result, fmt.Errorf("Не удалось запсиать пройденное расстояние (пробег) в км: %v", err)
	}

	if err = binary.Write(buf, binary.LittleEndian, e.DigitalInputs); err != nil {
		return result, fmt.Errorf("Не удалось записать битовые флаги, определяют состояние основных дискретных входов: %v", err)
	}

	if err = binary.Write(buf, binary.LittleEndian, e.Source); err != nil {
		return result, fmt.Errorf("Не удалось записать источник (событие), инициировавший посылку: %v", err)
	}

	if e.ALTE == "1" {
		bytesTmpBuf = []byte{0, 0, 0, 0}
		binary.LittleEndian.PutUint32(bytesTmpBuf, e.Altitude)
		if _, err = buf.Write(bytesTmpBuf[:3]); err != nil {
			return result, fmt.Errorf("Не удалось записать высоту над уровнем моря: %v", err)
		}
	}

	//TODO: разобраться с записью SourceData
	result = buf.Bytes()
	return result, nil
}

//Length получает длинну закодированной подзаписи
func (e *SrPosData) Length() uint16 {
	var result uint16

	if recBytes, err := e.Encode(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
