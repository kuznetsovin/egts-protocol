package egts

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
)

var (
	testEgtsSrPosDataBytes = []byte{0x55, 0x91, 0x02, 0x10, 0x6F, 0x1C, 0x05, 0x9E, 0x7A, 0xB5, 0x3C, 0x35,
		0x01, 0xD0, 0x87, 0x2C, 0x01, 0x00, 0x00, 0x00, 0x00}
	testEgtsSrPosData = SrPosData{
		NavigationTime:      time.Date(2018, time.July, 6, 20, 8, 53, 0, time.UTC),
		Latitude:            55.55389399769574,
		Longitude:           37.43236696287812,
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
		Direction:           172,
		Odometer:            []byte{0x01, 0x00, 0x00},
		DigitalInputs:       0,
		Source:              0,
	}
)

func TestEgtsSrPosData_Encode(t *testing.T) {
	posDataBytes, err := testEgtsSrPosData.Encode()
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(posDataBytes, testEgtsSrPosDataBytes) {
		t.Errorf("Байтовые строки не совпадают: %#X != %#X ", posDataBytes, testEgtsSrPosDataBytes)
	}
}

func TestEgtsSrPosData_Decode(t *testing.T) {
	posData := SrPosData{}

	if err := posData.Decode(testEgtsSrPosDataBytes); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if diff := cmp.Diff(posData, testEgtsSrPosData); diff != "" {
		t.Errorf("Записи не совпадают: (-нужно +сейчас)\n%s", diff)
	}

}
