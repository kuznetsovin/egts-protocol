package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

var (
	testEgtsSrPosData = EgtsSrPosData{
		NavigationTime:      time.Date(2018, time.July, 4, 20, 8, 53, 0, time.UTC),
		Latitude:            55,
		Longitude:           37,
		ALTE:                "0",
		LOHS:                "0",
		LAHS:                "0",
		MV:                  "0",
		BB:                  "0",
		CS:                  "0",
		FIX:                 "0",
		VLD:                 "1",
		DirectionHighestBit: 1,
		AltitudeSign:        0,
		Speed:               200,
		Direction:           44,
		Odometer:            []byte{0x01, 0x00, 0x00},
		DigitalInputs:       0,
		Source:              0,
	}
)

func TestEgtsSrPosData_Encode(t *testing.T) {
	//из-за преобразования float64 в формат пакета просхоит погрешность поэтому тестовый пакет немного изменяется
	testEgtsSrPosDataBytes := []byte{0xD5, 0x3F, 0x01, 0x10, 0x1b, 0xc7, 0x71, 0x9c, 0xf4, 0x49, 0x9f, 0x34,
		0x01, 0xD0, 0x87, 0x2C, 0x01, 0x00, 0x00, 0x00, 0x00}

	posDataBytes, err := testEgtsSrPosData.Encode()
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(posDataBytes, testEgtsSrPosDataBytes) {
		t.Errorf("Байтовые строки не совпадают: %v != %v ", posDataBytes, testEgtsSrPosDataBytes)
	}
}

func TestEgtsSrPosData_Decode(t *testing.T) {
	testEgtsSrPosDataBytes := []byte{0xD5, 0x3F, 0x01, 0x10, 0x6F, 0x1C, 0x05, 0x9E, 0x7A, 0xB5, 0x3C, 0x35,
		0x01, 0xD0, 0x87, 0x2C, 0x01, 0x00, 0x00, 0x00, 0x00}

	posData := EgtsSrPosData{}

	if err := posData.Decode(testEgtsSrPosDataBytes); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if !reflect.DeepEqual(posData, testEgtsSrPosData) {
		t.Errorf("Запись не совпадают: %v != %v ", posData, testEgtsSrPosData)
	}

}
