package main

import (
	"bytes"
	"testing"
)

func TestEGTS_SR_POS_DATA_ToBytes(t *testing.T) {
	posDataType := EGTS_SR_POS_DATA{
		NTM:  249535578,
		LAT:  3081986742,
		LONG: 951596088,
		ALTE: 0,
		LOHS: 0,
		LAHS: 0,
		MV:   0,
		BB:   0,
		CS:   0,
		FIX:  0,
		VLD:  1,
		DIRH: 0,
		ALTS: 0,
		SPD:  150,
		DIR:  246,
		ODM:  []byte{0x00, 0x00, 0x00},
		DIN:  0,
		SRC:  0,
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
	rd := RecordData{
		16,
		21,
		&EGTS_SR_POS_DATA{
			NTM:  249535578,
			LAT:  3081986742,
			LONG: 951596088,
			ALTE: 0,
			LOHS: 0,
			LAHS: 0,
			MV:   0,
			BB:   0,
			CS:   0,
			FIX:  0,
			VLD:  1,
			DIRH: 0,
			ALTS: 0,
			SPD:  150,
			DIR:  246,
			ODM:  []byte{0x00, 0x00, 0x00},
			DIN:  0,
			SRC:  0,
		},
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
