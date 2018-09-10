package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"time"
)

type RecordData struct {
	SubrecordType   byte
	SubrecordLength uint16
	SubrecordData   BinaryData
}

func (rd *RecordData) Encode() ([]byte, error) {
	var result []byte

	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, rd.SubrecordType); err != nil {
		return result, err
	}

	if err := binary.Write(buf, binary.LittleEndian, rd.SubrecordLength); err != nil {
		return result, err
	}

	srd, err := rd.SubrecordData.Encode()
	if err != nil {
		return result, err
	}

	buf.Write(srd)

	result = buf.Bytes()
	return result, nil
}

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

func (rd *EgtsSrPosData) Encode() ([]byte, error) {
	var (
		result []byte
	)

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
	flags, err := strconv.ParseUint(rd.VLD+rd.FIX+rd.CS+rd.BB+rd.MV+rd.LAHS+rd.LOHS+rd.ALTE, 2, 8)
	if err != nil {
		return result, fmt.Errorf("Не удалось сгенерировать байт флагов: %v", err)
	}

	if err := buf.WriteByte(uint8(flags)); err != nil {
		return result, fmt.Errorf("Не удалось записать флаги: %v", err)
	}

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

func (rd *EgtsSrPosData) Length() uint16 {
	var result uint16

	if recBytes, err := rd.Encode(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
