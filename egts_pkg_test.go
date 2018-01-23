package main

import (
	"bytes"
	"testing"
)

func TestEgtsPkgHeader_ToBytes(t *testing.T) {
	etgsHeader := EgtsPkgHeader{
		PRV:  1,
		SKID: 0,
		PRF:  0,
		RTE:  1,
		ENA:  0,
		CMP:  0,
		PR:   0,
		HL:   16,
		HE:   0,
		FDL:  59,
		PID:  3612,
		PT:   1,
		PRA:  200,
		RCA:  201,
		TTL:  0,
	}

	correctHeader := []byte{0x01, 0x00, 0x20, 0x10, 0x00, 0x3b, 0x00, 0x1c, 0x0e, 0x01,
		0xc8, 0x00, 0xc9, 0x00, 0x00, 0xca}
	resultBytes, err := etgsHeader.ToBytes()

	if err != nil {
		t.Error("Error etgs header decode ", err)
	}

	if !bytes.Equal(resultBytes, correctHeader) {
		t.Errorf("Incorrect etgs header decode: %v != %v ", resultBytes, correctHeader)
	}
}

func TestEgtsPkg_ToBytes(t *testing.T) {
	egtsPackage := EgtsPkg{
		EgtsPkgHeader: EgtsPkgHeader{
			PRV:  1,
			SKID: 0,
			PRF:  0,
			RTE:  1,
			ENA:  0,
			CMP:  0,
			PR:   0,
			HL:   16,
			HE:   0,
			FDL:  59,
			PID:  3612,
			PT:   1,
			PRA:  200,
			RCA:  201,
			TTL:  0,
		},
		SFRD: &EGTS_PT_APPDATA{
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
		},
	}

	correctPackage := []byte{0x01, 0x00, 0x20, 0x10, 0x00, 0x3b, 0x00, 0x1c, 0x0e, 0x01, 0xc8, 0x00, 0xc9,
		0x00, 0x00, 0xca, 0x30, 0x00, 0x00, 0x00, 0x81, 0x07, 0x46, 0xa2, 0x00, 0x02, 0x02, 0x10, 0x15, 0x00,
		0x5a, 0x9c, 0xdf, 0x0e, 0xb6, 0x62, 0xb3, 0xb7, 0x38, 0x34, 0xb8, 0x38, 0x01, 0x96, 0x00, 0xf6, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x10, 0x15, 0x00, 0x5a, 0x9c, 0xdf, 0x0e, 0xb6, 0x62, 0xb3, 0xb7, 0x38, 0x34, 0xb8, 0x38,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2a, 0x28}
	result, err := egtsPackage.ToBytes()

	if err != nil {
		t.Error("Error etgs header decode ", err)
	}

	if !bytes.Equal(result, correctPackage) {
		t.Errorf("Incorrect etgs header decode: %v != %v ", result, correctPackage)
	}
}
