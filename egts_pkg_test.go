package main

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
)

var (
	egtsPkgPosData = EgtsPackage{
		ProtocolVersion:  1,
		SecurityKeyID:    0,
		Prefix:           "00",
		Route:            "0",
		EncryptionAlg:    "00",
		Compression:      "0",
		Priority:         "11",
		HeaderLength:     11,
		HeaderEncoding:   0,
		FrameDataLength:  35,
		PacketIdentifier: 138,
		PacketType:       1,
		HeaderCheckSum:   73,
		ServicesFrameData: &ServiceDataSet{
			ServiceDataRecord{
				RecordLength:             24,
				RecordNumber:             97,
				SourceServiceOnDevice:    "1",
				RecipientServiceOnDevice: "0",
				Group: "0",
				RecordProcessingPriority: "11",
				TimeFieldExists:          "0",
				EventIDFieldExists:       "0",
				ObjectIDFieldExists:      "1",
				ObjectIdentifier:         133552,
				SourceServiceType:        2,
				RecipientServiceType:     2,
				RecordDataSet: RecordDataSet{
					RecordData{
						SubrecordType:   16,
						SubrecordLength: 21,
						SubrecordData: &EgtsSrPosData{
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
						},
					},
				},
			},
		},
		ServicesFrameDataCheckSum: 10188, //52263
	}
)

func TestEgtsPackagePosData_Encode(t *testing.T) {
	testEgtsPkgBytes := []byte{0x01, 0x00, 0x03, 0x0B, 0x00, 0x23, 0x00, 0x8A, 0x00, 0x01, 0x49, 0x18, 0x00, 0x61,
		0x00, 0x99, 0xB0, 0x09, 0x02, 0x00, 0x02, 0x02, 0x10, 0x15, 0x00, 0xD5, 0x3F, 0x01, 0x10, 0x1b, 0xc7, 0x71, 0x9c,
		0xf4, 0x49, 0x9f, 0x34, 0x01, 0xD0, 0x87, 0x2C, 0x01, 0x00, 0x00, 0x00, 0x00, 0xAC, 0xC9}

	posDataBytes, err := egtsPkgPosData.Encode()
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(posDataBytes, testEgtsPkgBytes) {
		t.Errorf("Байтовые строки не совпадают: %v != %v ", posDataBytes, testEgtsPkgBytes)
	}
}

func TestEgtsPackagePosData_Decode(t *testing.T) {
	egtsPkgBytes := []byte{0x01, 0x00, 0x03, 0x0B, 0x00, 0x23, 0x00, 0x8A, 0x00, 0x01, 0x49, 0x18, 0x00, 0x61,
		0x00, 0x99, 0xB0, 0x09, 0x02, 0x00, 0x02, 0x02, 0x10, 0x15, 0x00, 0xD5, 0x3F, 0x01, 0x10, 0x6F, 0x1C, 0x05, 0x9E,
		0x7A, 0xB5, 0x3C, 0x35, 0x01, 0xD0, 0x87, 0x2C, 0x01, 0x00, 0x00, 0x00, 0x00, 0xCC, 0x27}

	egtsPkg := EgtsPackage{}

	if _, err := egtsPkg.Decode(egtsPkgBytes); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if diff := cmp.Diff(egtsPkg, egtsPkgPosData); diff != "" {
		t.Errorf("Записи не совпадают: (-нужно +сейчас)\n%s", diff)
	}
}
