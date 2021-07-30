package egts

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	egtsPkgPosDataBytes = []byte{0x01, 0x00, 0x03, 0x0B, 0x00, 0x23, 0x00, 0x8A, 0x00, 0x01, 0x49, 0x18, 0x00, 0x61,
		0x00, 0x99, 0xB0, 0x09, 0x02, 0x00, 0x02, 0x02, 0x10, 0x15, 0x00, 0xD5, 0x3F, 0x01, 0x10, 0x6F, 0x1C, 0x05, 0x9E,
		0x7A, 0xB5, 0x3C, 0x35, 0x01, 0xD0, 0x87, 0x2C, 0x01, 0x00, 0x00, 0x00, 0x00, 0xCC, 0x27}
)

func TestEgtsPackagePosData_Encode(t *testing.T) {
	egtsPkgPosData := Package{
		ProtocolVersion:  1,
		SecurityKeyID:    0,
		Prefix:           "00",
		Route:            "0",
		EncryptionAlg:    "00",
		Compression:      "0",
		Priority:         "11",
		HeaderEncoding:   0,
		FrameDataLength:  35,
		PacketIdentifier: 138,
		PacketType:       1,
		HeaderCheckSum:   73,
		ServicesFrameData: &ServiceDataSet{
			ServiceDataRecord{
				RecordNumber:             97,
				SourceServiceOnDevice:    "1",
				RecipientServiceOnDevice: "0",
				Group:                    "0",
				RecordProcessingPriority: "11",
				TimeFieldExists:          "0",
				EventIDFieldExists:       "0",
				ObjectIDFieldExists:      "1",
				ObjectIdentifier:         133552,
				SourceServiceType:        2,
				RecipientServiceType:     2,
				RecordDataSet: RecordDataSet{
					RecordData{
						SubrecordData: &SrPosData{
							NavigationTime:      time.Date(2018, time.July, 5, 20, 8, 53, 0, time.UTC),
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
							Odometer:            1,
							DigitalInputs:       0,
							Source:              0,
						},
					},
				},
			},
		},
	}

	posDataBytes, err := egtsPkgPosData.Encode()
	if assert.NoError(t, err) {
		assert.Equal(t, posDataBytes, egtsPkgPosDataBytes)
	}
}

func TestEgtsPackagePosData_Decode(t *testing.T) {
	egtsPkg := Package{}
	egtsPkgPosData := Package{
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
				Group:                    "0",
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
						SubrecordData: &SrPosData{
							NavigationTime:      time.Date(2018, time.July, 5, 20, 8, 53, 0, time.UTC),
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
							Odometer:            1,
							DigitalInputs:       0,
							Source:              0,
						},
					},
				},
			},
		},
		ServicesFrameDataCheckSum: 10188,
	}

	if _, err := egtsPkg.Decode(egtsPkgPosDataBytes); assert.NoError(t, err) {
		assert.Equal(t, egtsPkg, egtsPkgPosData)
	}
}

func TestFullCycleCoding(t *testing.T) {
	egtsPkg := Package{}
	egtsPkgPosData := Package{
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
				Group:                    "0",
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
						SubrecordData: &SrPosData{
							NavigationTime:      time.Date(2018, time.July, 5, 20, 8, 53, 0, time.UTC),
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
							Odometer:            1,
							DigitalInputs:       0,
							Source:              0,
						},
					},
				},
			},
		},
		ServicesFrameDataCheckSum: 10188,
	}

	_, err := egtsPkg.Decode(egtsPkgPosDataBytes)
	if assert.NoError(t, err) {
		assert.Equal(t, egtsPkg, egtsPkgPosData)

		posDataBytes, err := egtsPkg.Encode()
		if assert.NoError(t, err) {
			assert.Equal(t, posDataBytes, egtsPkgPosDataBytes)
		}
	}
}

func TestRebuildCycleCoding(t *testing.T) {
	egtsPkg := Package{
		ProtocolVersion:  1,
		SecurityKeyID:    0,
		Prefix:           "00",
		Route:            "0",
		EncryptionAlg:    "00",
		Compression:      "0",
		Priority:         "10",
		HeaderLength:     11,
		HeaderEncoding:   0,
		FrameDataLength:  48,
		PacketIdentifier: 4608,
		PacketType:       1,
		HeaderCheckSum:   0x1b,
		ServicesFrameData: &ServiceDataSet{
			ServiceDataRecord{
				RecordLength:             37,
				RecordNumber:             134,
				SourceServiceOnDevice:    "0",
				RecipientServiceOnDevice: "0",
				Group:                    "0",
				RecordProcessingPriority: "10",
				TimeFieldExists:          "0",
				EventIDFieldExists:       "1",
				ObjectIDFieldExists:      "0",
				EventIdentifier:          3436,
				SourceServiceType:        2,
				RecipientServiceType:     2,
				RecordDataSet: RecordDataSet{
					RecordData{
						SubrecordType:   16,
						SubrecordLength: 24,
						SubrecordData: &SrPosData{
							NavigationTime:      time.Date(2021, time.February, 20, 0, 30, 40, 0, time.UTC),
							Latitude:            46.9429406935682,
							Longitude:           142.732571163851,
							ALTE:                "1",
							LOHS:                "0",
							LAHS:                "0",
							MV:                  "1",
							BB:                  "1",
							CS:                  "0",
							FIX:                 "1",
							VLD:                 "1",
							DirectionHighestBit: 0,
							AltitudeSign:        0,
							Speed:               34,
							Direction:           172,
							Odometer:            191,
							DigitalInputs:       144,
							Source:              0,
							Altitude:            30,
						},
					},
					RecordData{
						SubrecordType:   27,
						SubrecordLength: 7,
						SubrecordData: &SrLiquidLevelSensor{
							LiquidLevelSensorErrorFlag: "1",
							LiquidLevelSensorValueUnit: "00",
							RawDataFlag:                "0",
							LiquidLevelSensorNumber:    1,
							ModuleAddress:              uint16(1),
							LiquidLevelSensorData:      uint32(0),
						},
					},
				},
			},
		},
	}
	p := []byte{0x01, 0x00, 0x02, 0x0b, 0x00, 0x30, 0x00, 0x00, 0x12, 0x01, 0x84, 0x25, 0x00, 0x86, 0x00,
		0x12, 0x6c, 0x0d, 0x00, 0x00, 0x02, 0x02, 0x10, 0x18, 0x00, 0x30, 0x1d, 0xf3, 0x14, 0x65, 0xce,
		0x86, 0x85, 0xde, 0x57, 0xff, 0xca, 0x9b, 0x54, 0x01, 0xac, 0xbf, 0x00, 0x00, 0x90, 0x00, 0x1e,
		0x00, 0x00, 0x1b, 0x07, 0x00, 0x41, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0xe8}

	encodePkg, err := egtsPkg.Encode()
	if assert.NoError(t, err) {
		assert.Equal(t, encodePkg, p)

		dp := Package{}
		_, err = dp.Decode(encodePkg)
		assert.NoError(t, err)
	}
}

func TestRebuildOID(t *testing.T) {
	egtsPkg := Package{
		ProtocolVersion:  1,
		SecurityKeyID:    0,
		Prefix:           "00",
		Route:            "0",
		EncryptionAlg:    "00",
		Compression:      "0",
		Priority:         "10",
		HeaderLength:     11,
		HeaderEncoding:   0,
		FrameDataLength:  48,
		PacketIdentifier: 4608,
		PacketType:       1,
		HeaderCheckSum:   0x1b,
		ServicesFrameData: &ServiceDataSet{
			ServiceDataRecord{
				RecordLength:             37,
				RecordNumber:             134,
				SourceServiceOnDevice:    "0",
				RecipientServiceOnDevice: "0",
				Group:                    "0",
				RecordProcessingPriority: "10",
				TimeFieldExists:          "0",
				EventIDFieldExists:       "1",
				ObjectIDFieldExists:      "1",
				EventIdentifier:          3436,
				ObjectIdentifier:         326009033,
				SourceServiceType:        2,
				RecipientServiceType:     2,
				RecordDataSet: RecordDataSet{
					RecordData{
						SubrecordType:   16,
						SubrecordLength: 24,
						SubrecordData: &SrPosData{
							NavigationTime:      time.Date(2021, time.February, 20, 0, 30, 40, 0, time.UTC),
							Latitude:            46.9429406935682,
							Longitude:           142.732571163851,
							ALTE:                "1",
							LOHS:                "0",
							LAHS:                "0",
							MV:                  "1",
							BB:                  "1",
							CS:                  "0",
							FIX:                 "1",
							VLD:                 "1",
							DirectionHighestBit: 0,
							AltitudeSign:        0,
							Speed:               34,
							Direction:           172,
							Odometer:            191,
							DigitalInputs:       144,
							Source:              0,
							Altitude:            30,
						},
					},
					RecordData{
						SubrecordType:   27,
						SubrecordLength: 7,
						SubrecordData: &SrLiquidLevelSensor{
							LiquidLevelSensorErrorFlag: "1",
							LiquidLevelSensorValueUnit: "00",
							RawDataFlag:                "0",
							LiquidLevelSensorNumber:    1,
							ModuleAddress:              uint16(1),
							LiquidLevelSensorData:      uint32(0),
						},
					},
				},
			},
		},
	}

	encodePkg, err := egtsPkg.Encode()
	if assert.NoError(t, err) {
		p := Package{}

		_, err = p.Decode(encodePkg)
		assert.NoError(t, err)
	}
}
