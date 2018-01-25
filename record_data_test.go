package main

import (
	"bytes"
	"testing"
)

func TestEGTS_SR_POS_DATA_ToBytes(t *testing.T) {
	posDataType := EGTS_SR_POS_DATA{
		NavigationTime:      249535578,
		Latitude:            3081986742,
		Longitude:           951596088,
		ALTE:                0,
		LOHS:                0,
		LAHS:                0,
		MV:                  0,
		BB:                  0,
		CS:                  0,
		FIX:                 0,
		VLD:                 1,
		DirectionHighestBit: 0,
		AltitudeSign:        0,
		Speed:               150,
		Direction:           246,
		Odometer:            []byte{0x00, 0x00, 0x00},
		DigitalInputs:       0,
		Source:              0,
	}

	correctRd := []byte{0x5a, 0x9c, 0xdf, 0x0e, 0xb6, 0x62, 0xb3, 0xb7, 0x38, 0x34, 0xb8, 0x38, 0x01,
		0x96, 0x00, 0xf6, 0x00, 0x00, 0x00, 0x00, 0x00}

	resultBytes, err := posDataType.ToBytes()

	if err != nil {
		t.Error("Error EGTS_SR_POS_DATA decode ", err)
	}

	if !bytes.Equal(resultBytes, correctRd) {
		t.Errorf("Incorrect EGTS_SR_POS_DATA decode: %v != %v ", resultBytes, correctRd)
	}
}

func TestRecordData_ToBytes(t *testing.T) {
	esrpd := EGTS_SR_POS_DATA{
		NavigationTime:      249535578,
		Latitude:            3081986742,
		Longitude:           951596088,
		ALTE:                0,
		LOHS:                0,
		LAHS:                0,
		MV:                  0,
		BB:                  0,
		CS:                  0,
		FIX:                 0,
		VLD:                 1,
		DirectionHighestBit: 0,
		AltitudeSign:        0,
		Speed:               150,
		Direction:           246,
		Odometer:            []byte{0x00, 0x00, 0x00},
		DigitalInputs:       0,
		Source:              0,
	}
	rd := RecordData{
		16,
		esrpd.Length(),
		&esrpd,
	}

	resultBytes, err := rd.ToBytes()

	if err != nil {
		t.Error("Error record data decode ", err)
	}

	correctRd := []byte{0x10, 0x15, 0x00, 0x5a, 0x9c, 0xdf, 0x0e, 0xb6, 0x62, 0xb3, 0xb7, 0x38, 0x34, 0xb8, 0x38, 0x01,
		0x96, 0x00, 0xf6, 0x00, 0x00, 0x00, 0x00, 0x00}

	if !bytes.Equal(resultBytes, correctRd) {
		t.Errorf("Incorrect record data decode: %v != %v ", resultBytes, correctRd)
	}
}
