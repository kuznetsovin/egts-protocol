package main

import (
	"bytes"
	"testing"
)

func TestEgtsPkgHeader_ToBytes(t *testing.T) {
	etgsHeader := EgtsPkgHeader{
		ProtocolVersion:  1,
		SecurityKeyID:    0,
		PRF:              0,
		RTE:              1,
		ENA:              0,
		CMP:              0,
		PR:               0,
		HeaderLength:     16,
		HeaderEncoding:   0,
		FrameDataLength:  59,
		PacketIdentifier: 3612,
		PacketType:       1,
		PeerAddress:      200,
		RecipientAddress: 201,
		TimeToLive:       0,
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
			ProtocolVersion:  1,
			SecurityKeyID:    0,
			PRF:              0,
			RTE:              1,
			ENA:              0,
			CMP:              0,
			PR:               0,
			HeaderLength:     16,
			HeaderEncoding:   0,
			FrameDataLength:  59,
			PacketIdentifier: 3612,
			PacketType:       1,
			PeerAddress:      200,
			RecipientAddress: 201,
			TimeToLive:       0,
		},
		ServicesFrameData: &EGTS_PT_APPDATA{
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
					RecordDataSet: RecordDataSet{
						{
							16,
							21,
							&EGTS_SR_POS_DATA{
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
							},
						},
						{
							16,
							21,
							&EGTS_SR_POS_DATA{
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
								Speed:               0,
								Direction:           0,
								Odometer:            []byte{0x00, 0x00, 0x00},
								DigitalInputs:       0,
								Source:              0,
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
