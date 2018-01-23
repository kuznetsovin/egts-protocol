package main

import (
	"bytes"
	"testing"
)

func TestServiceDataRecord_ToBytes(t *testing.T) {
	srd := ServiceDataRecord{
		RecordLength:             48,
		RecordNumber:             0,
		SourceServiceOnDevice:    1,
		RecipientServiceOnDevice: 0,
		Group: 0,
		RecordProcessingPriority: 0,
		TimeFieldExists:          0,
		EventIDFieldExists:       0,
		ObjectIDFieldExists:      1,
		ObjectIdentifier:         10634759,
		SourceServiceType:        2,
		RecipientServiceType:     2,
		RecordData: []RecordData{
			{
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
			},
			{
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
					SPD:  0,
					DIR:  0,
					ODM:  []byte{0x00, 0x00, 0x00},
					DIN:  0,
					SRC:  0,
				},
			},
		},
	}

	result, err := srd.ToBytes()

	if err != nil {
		t.Error("Error record data decode ", err)
	}

	correctSDR := []byte{0x30, 0x00, 0x00, 0x00, 0x81, 0x07, 0x46, 0xa2, 0x00, 0x02, 0x02, 0x10, 0x15, 0x00,
		0x5a, 0x9c, 0xdf, 0x0e, 0xb6, 0x62, 0xb3, 0xb7, 0x38, 0x34, 0xb8, 0x38, 0x01, 0x96, 0x00, 0xf6, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x10, 0x15, 0x00, 0x5a, 0x9c, 0xdf, 0x0e, 0xb6, 0x62, 0xb3, 0xb7, 0x38, 0x34, 0xb8, 0x38,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	if !bytes.Equal(result, correctSDR) {
		t.Errorf("Incorrect record data decode: %v != %v ", result, correctSDR)
	}
}

func TestEGTS_PT_APPDATA_ToBytes(t *testing.T) {
	appdata := EGTS_PT_APPDATA{
		[]ServiceDataRecord{
			{
				RecordLength:             48,
				RecordNumber:             0,
				SourceServiceOnDevice:    1,
				RecipientServiceOnDevice: 0,
				Group: 0,
				RecordProcessingPriority: 0,
				TimeFieldExists:          0,
				EventIDFieldExists:       0,
				ObjectIDFieldExists:      1,
				ObjectIdentifier:         10634759,
				SourceServiceType:        2,
				RecipientServiceType:     2,
				RecordData: []RecordData{
					{
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
					},
					{
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
							SPD:  0,
							DIR:  0,
							ODM:  []byte{0x00, 0x00, 0x00},
							DIN:  0,
							SRC:  0,
						},
					},
				},
			},
		},
	}

	result, err := appdata.ToBytes()

	if err != nil {
		t.Error("Error record data decode ", err)
	}

	correctSDR := []byte{0x30, 0x00, 0x00, 0x00, 0x81, 0x07, 0x46, 0xa2, 0x00, 0x02, 0x02, 0x10, 0x15, 0x00,
		0x5a, 0x9c, 0xdf, 0x0e, 0xb6, 0x62, 0xb3, 0xb7, 0x38, 0x34, 0xb8, 0x38, 0x01, 0x96, 0x00, 0xf6, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x10, 0x15, 0x00, 0x5a, 0x9c, 0xdf, 0x0e, 0xb6, 0x62, 0xb3, 0xb7, 0x38, 0x34, 0xb8, 0x38,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	if !bytes.Equal(result, correctSDR) {
		t.Errorf("Incorrect record data decode: %v != %v ", result, correctSDR)
	}
}
