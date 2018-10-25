package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"time"
)

type EgtsSrPosData struct {
	NavigationTime      time.Time
	Latitude            float64
	Longitude           float64
	ALTE                string
	LOHS                string
	LAHS                string
	MV                  string
	BB                  string
	CS                  string
	FIX                 string
	VLD                 string
	DirectionHighestBit uint8
	AltitudeSign        uint8
	Speed               uint16
	Direction           byte
	Odometer            []byte
	DigitalInputs       byte
	Source              byte
	Altitude            []byte
	SourceData          int16
}

func (e *EgtsSrPosData) Decode(content []byte) error {
	var (
		err   error
		flags byte
		speed uint64
	)
	buf := bytes.NewReader(content)

	// Преобразуем время навигации к формату, который требует стандарт: количество секунд с 00:00:00 01.01.2010 UTC
	startDate := time.Date(2010, time.January, 0, 0, 0, 0, 0, time.UTC)
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
	e.Latitude = float64(int(preFieldVal) * 90 / 0xFFFFFFFF)

	// В протоколе значение хранится в виде: долгота по модулю, градусы/180*0xFFFFFFFF  и взята целая часть
	if _, err = buf.Read(tmpUint32Buf); err != nil {
		return fmt.Errorf("Не удалось получить время долгату: %v", err)
	}
	preFieldVal = binary.LittleEndian.Uint32(tmpUint32Buf)
	e.Longitude = float64(int(preFieldVal) * 180 / 0xFFFFFFFF)

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

	//TODO: разобраться с битом DirectionHighestBit
	if e.Direction, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить направление движения: %v", err)
	}

	bytesTmpBuf := make([]byte, 3)
	if _, err = buf.Read(bytesTmpBuf); err != nil {
		return fmt.Errorf("Не удалось получить пройденное расстояние (пробег) в км: %v", err)
	}
	e.Odometer = bytesTmpBuf

	if e.DigitalInputs, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить битовые флаги, определяют состояние основных дискретных входов: %v", err)
	}

	if e.Source, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить источник (событие), инициировавший посылку: %v", err)
	}

	if e.ALTE == "1" {
		if _, err = buf.Read(bytesTmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить высоту над уровнем моря: %v", err)
		}
		e.Altitude = bytesTmpBuf
	}

	//TODO: разобраться с разбором SourceData
	return err
}

func (e *EgtsSrPosData) Encode() ([]byte, error) {
	var (
		err    error
		result []byte
	)

	buf := new(bytes.Buffer)
	// Преобразуем время навигации к формату, который требует стандарт: количество секунд с 00:00:00 01.01.2010 UTC
	startDate := time.Date(2010, time.January, 0, 0, 0, 0, 0, time.UTC)
	if err := binary.Write(buf, binary.LittleEndian, uint32(e.NavigationTime.Sub(startDate).Seconds())); err != nil {
		return result, fmt.Errorf("Не удалось записать время навигации: %v", err)
	}

	// В протоколе значение хранится в виде: широта по модулю, градусы/90*0xFFFFFFFF  и взята целая часть
	if err := binary.Write(buf, binary.LittleEndian, uint32(e.Latitude/90*0xFFFFFFFF)); err != nil {
		return result, fmt.Errorf("Не удалось записать широту: %v", err)
	}


	// В протоколе значение хранится в виде: долгота по модулю, градусы/180*0xFFFFFFFF  и взята целая часть
	if err := binary.Write(buf, binary.LittleEndian, uint32(e.Longitude/180*0xFFFFFFFF)); err != nil {
		return result, fmt.Errorf("Не удалось записать долготу: %v", err)
	}

	//байт флагов
	flags, err := strconv.ParseUint(e.ALTE+e.LOHS+e.LAHS+e.MV+e.BB+e.CS+e.FIX+e.VLD, 2, 8)
	if err != nil {
		return result, fmt.Errorf("Не удалось сгенерировать байт флагов: %v", err)
	}

	if err := buf.WriteByte(uint8(flags)); err != nil {
		return result, fmt.Errorf("Не удалось записать флаги: %v", err)
	}

	// скорость
	speed := e.Speed * 10 | uint16(e.DirectionHighestBit)<<15 // 15 бит
	speed = speed | uint16(e.AltitudeSign)<<14           //14 бит
	spd := make([]byte, 2)
	binary.LittleEndian.PutUint16(spd, speed)
	if _, err = buf.Write(spd); err != nil {
		return result, fmt.Errorf("Не удалось записать скорость: %v", err)
	}

	//TODO: разобраться с битом DirectionHighestBit
	if err := binary.Write(buf, binary.LittleEndian, e.Direction); err != nil {
		return result, fmt.Errorf("Не удалось записать направление движения: %v", err)
	}

	if _, err = buf.Write(e.Odometer); err != nil {
		return result, fmt.Errorf("Не удалось запсиать пройденное расстояние (пробег) в км: %v", err)
	}

	if err := binary.Write(buf, binary.LittleEndian, e.DigitalInputs); err != nil {
		return result, fmt.Errorf("Не удалось записать битовые флаги, определяют состояние основных дискретных входов: %v", err)
	}

	if err := binary.Write(buf, binary.LittleEndian, e.Source); err != nil {
		return result, fmt.Errorf("Не удалось записать источник (событие), инициировавший посылку: %v", err)
	}

	if e.ALTE == "1" {
		if _, err = buf.Write(e.Altitude); err != nil {
			return result, fmt.Errorf("Не удалось записать высоту над уровнем моря: %v", err)
		}
	}

	//TODO: разобраться с записью SourceData
	result = buf.Bytes()
	return result, nil
}

func (e *EgtsSrPosData) Length() uint16 {
	var result uint16

	if recBytes, err := e.Encode(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
