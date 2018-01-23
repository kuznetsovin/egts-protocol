package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
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
	NTM uint32

	// широта по модулю, градусы/90*0xFFFFFFFF  и взята целая часть;
	LAT uint32

	// долгота по модулю, градусы/180*0xFFFFFFFF  и взята целая часть;
	LONG uint32

	// битовый флаг определяет наличие поля ALT в подзаписи: 1 - поле ALT передается; 0 - не передается;
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

	// старший бит (8) параметра DIR;
	DIRH uint8

	// битовый флаг, определяет высоту относительно уровня моря и имеет  смысл только при установленном флаге ALTE:
	// 0 - точка выше уровня моря;
	// 1 - ниже уровня моря;
	ALTS uint8

	// скорость в км/ч с дискретностью 0,1 км/ч;
	SPD uint16

	// направление движения.
	DIR byte

	// пройденное расстояние (пробег) в км, с дискретностью 0,1 км;
	ODM []byte

	// битовые флаги, определяют состояние основных дискретных входов 1 .. 8 (если бит равен 1, то соответствующий
	// вход активен, если 0, то неактивен). Данное поле включено для удобства использования и экономии трафика при
	// работе в системах мониторинга транспорта базового уровня;
	DIN byte

	// определяет источник (событие), инициировавший посылку данной навигационной информации
	SRC byte

	// высота над уровнем моря, м (опциональный параметр, наличие которого определяется битовым флагом ALTE);
	ALT [3]byte

	// данные, характеризующие источник (событие) из поля SRC. Наличие и интерпретация значения данного поля
	// определяется полем SRC.
	SRCD int16
}

func (rd *EGTS_SR_POS_DATA) ToBytes() ([]byte, error) {
	result := []byte{}

	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, rd.NTM); err != nil {
		return result, err
	}

	if err := binary.Write(buf, binary.LittleEndian, rd.LAT); err != nil {
		return result, err
	}

	if err := binary.Write(buf, binary.LittleEndian, rd.LONG); err != nil {
		return result, err
	}

	//байт флагов
	flagsByte := fmt.Sprintf("%b%b%b%b%b%b%b%b", rd.ALTE, rd.LOHS, rd.LAHS, rd.MV, rd.BB,
		rd.CS, rd.FIX, rd.VLD)

	flagByte, err := bitsToByte(flagsByte)
	if err != nil {
		return result, err
	}
	buf.WriteByte(flagByte)

	// скорость
	bitSPD := strings.Replace(fmt.Sprintf("%b%b%14b", rd.DIRH, rd.ALTS, rd.SPD), " ", "0", -1)
	spd, err := bitsToBytes(bitSPD, 2)
	if err != nil {
		return result, err
	}
	buf.Write(spd)

	if err := binary.Write(buf, binary.LittleEndian, rd.DIR); err != nil {
		return result, err
	}

	buf.Write(rd.ODM)

	if err := binary.Write(buf, binary.LittleEndian, rd.DIN); err != nil {
		return result, err
	}

	if err := binary.Write(buf, binary.LittleEndian, rd.SRC); err != nil {
		return result, err
	}

	result = buf.Bytes()
	return result, nil
}
