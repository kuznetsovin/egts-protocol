package main

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
)

var (
	testRecordDataBytes = []byte{0x10, 0x15, 0x00, 0xD5, 0x3F, 0x01, 0x10, 0x6F, 0x1C, 0x05, 0x9E, 0x7A, 0xB5,
		0x3C, 0x35, 0x01, 0xD0, 0x87, 0x2C, 0x01, 0x00, 0x00, 0x00, 0x00}
	testRecordDataSet = RecordDataSet{
		RecordData{
			SubrecordType:   16,
			SubrecordLength: 21,
			SubrecordData: &EgtsSrPosData{
				NavigationTime:      time.Date(2018, time.July, 4, 20, 8, 53, 0, time.UTC),
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
			},
		},
	}
)

func TestRecordDataSet_Encode(t *testing.T) {
	rdBytes, err := testRecordDataSet.Encode()
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(rdBytes, testRecordDataBytes) {
		t.Errorf("Байтовые строки не совпадают: %v != %v ", rdBytes, testRecordDataBytes)
	}
}

func TestRecordDataSet_Decode(t *testing.T) {
	rds := RecordDataSet{}

	if err := rds.Decode(testRecordDataBytes); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if diff := cmp.Diff(rds, testRecordDataSet); diff != "" {
		t.Errorf("Записи не совпадают: (-нужно +сейчас)\n%s", diff)
	}

}
