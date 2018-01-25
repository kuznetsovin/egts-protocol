package main

import (
	"bytes"
	"encoding/binary"
	"time"
)

type RecordData struct {
	// тип подзаписи
	SubrecordType byte

	// длина подзаписи в поле SubrecordData
	SubrecordLength uint16

	// данные позаписи
	SubrecordData BinaryData
}

func (rd *RecordData) ToBytes() ([]byte, error) {
	var result []byte

	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, rd.SubrecordType); err != nil {
		return result, err
	}

	if err := binary.Write(buf, binary.LittleEndian, rd.SubrecordLength); err != nil {
		return result, err
	}

	srd, err := rd.SubrecordData.ToBytes()
	if err != nil {
		return result, err
	}

	buf.Write(srd)

	result = buf.Bytes()
	return result, nil
}

type EGTS_SR_POS_DATA struct {
	//время навигации (количество секунд с 00:00:00 01.01.2010 UTC);
	NavigationTime time.Time

	// Значение широты в градусах WGS84
	Latitude float64

	// Значение долготы в градусах WGS84
	Longitude float64

	// битовый флаг определяет наличие поля Altitude в подзаписи: 1 - поле Altitude передается; 0 - не передается;
	ALTE uint8

	// битовый флаг определяет полушарие долготы: 0 - восточная долгота; 1 - западная долгота;
	LOHS uint8

	//  битовый флаг определяет полушарие широты: 0 - северная широта; 1 - южная широта;
	LAHS uint8

	// битовый флаг, признак движения: 1 - движение; 0 - транспортное средство находится в режиме стоянки;
	MV uint8

	// битовый флаг, признак отправки данных из памяти ("черный ящик"): 0 - актуальные данные;
	// 1 - данные из памяти ("черного ящика");
	BB uint8

	// битовое поле, тип определения координат: 0 - 2D fix; 1 - 3D fix;
	CS uint8

	// битовое поле, тип используемой системы:
	// 0 - система координат WGS-84;
	// 1 - государственная геоцентрическая система координат (ПЗ-90.02);
	FIX uint8

	// битовый флаг, признак "валидности" координатных данных: 1 - данные "валидны"; 0 - "невалидные" данные;
	VLD uint8

	// старший бит (8) параметра Direction;
	DirectionHighestBit uint8

	// битовый флаг, определяет высоту относительно уровня моря и имеет  смысл только при установленном флаге ALTE:
	// 0 - точка выше уровня моря;
	// 1 - ниже уровня моря;
	AltitudeSign uint8

	// скорость в км/ч с дискретностью 0,1 км/ч;
	Speed uint16

	// направление движения.
	Direction byte

	// пройденное расстояние (пробег) в км, с дискретностью 0,1 км;
	Odometer []byte

	// битовые флаги, определяют состояние основных дискретных входов 1 .. 8 (если бит равен 1, то соответствующий
	// вход активен, если 0, то неактивен). Данное поле включено для удобства использования и экономии трафика при
	// работе в системах мониторинга транспорта базового уровня;
	DigitalInputs byte

	// определяет источник (событие), инициировавший посылку данной навигационной информации
	Source byte

	// высота над уровнем моря, м (опциональный параметр, наличие которого определяется битовым флагом ALTE);
	Altitude []byte

	// данные, характеризующие источник (событие) из поля Source. Наличие и интерпретация значения данного поля
	// определяется полем Source.
	SourceData int16
}

func (rd *EGTS_SR_POS_DATA) ToBytes() ([]byte, error) {
	result := []byte{}

	buf := new(bytes.Buffer)
	// Преобразуем время навигации к формату, который требует стандарт: количество секунд с 00:00:00 01.01.2010 UTC
	startDate := time.Date(2010, time.January, 0, 0, 0, 0, 0, time.UTC)
	if err := binary.Write(buf, binary.LittleEndian, uint32(rd.NavigationTime.Sub(startDate).Seconds())); err != nil {
		return result, err
	}

	// В протоколе значение хранится в виде: широта по модулю, градусы/90*0xFFFFFFFF  и взята целая часть
	if err := binary.Write(buf, binary.LittleEndian, uint32(rd.Latitude/90*0xFFFFFFFF)); err != nil {
		return result, err
	}

	// В протоколе значение хранится в виде: долгота по модулю, градусы/180*0xFFFFFFFF  и взята целая часть
	if err := binary.Write(buf, binary.LittleEndian, uint32(rd.Longitude/180*0xFFFFFFFF)); err != nil {
		return result, err
	}

	//байт флагов
	flagsByte := byte(0) | rd.VLD      // 0 бит
	flagsByte = flagsByte | rd.FIX<<1  // 1 бит
	flagsByte = flagsByte | rd.CS<<2   // 2 бит
	flagsByte = flagsByte | rd.BB<<3   // 3 бит
	flagsByte = flagsByte | rd.MV<<4   // 4 бит
	flagsByte = flagsByte | rd.LAHS<<5 // 5 бит
	flagsByte = flagsByte | rd.LOHS<<6 // 6 бит
	flagsByte = flagsByte | rd.ALTE<<7 // 7 бит
	buf.WriteByte(flagsByte)

	// скорость
	speed := rd.Speed | uint16(rd.DirectionHighestBit)<<15 // 15 бит
	speed = speed | uint16(rd.AltitudeSign)<<14            //14 бит
	spd := make([]byte, 2)
	binary.LittleEndian.PutUint16(spd, speed)
	buf.Write(spd)

	if err := binary.Write(buf, binary.LittleEndian, rd.Direction); err != nil {
		return result, err
	}

	buf.Write(rd.Odometer)

	if err := binary.Write(buf, binary.LittleEndian, rd.DigitalInputs); err != nil {
		return result, err
	}

	if err := binary.Write(buf, binary.LittleEndian, rd.Source); err != nil {
		return result, err
	}

	result = buf.Bytes()
	return result, nil
}

func (rd *EGTS_SR_POS_DATA) Length() uint16 {
	var result uint16

	if recBytes, err := rd.ToBytes(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
