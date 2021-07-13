package egts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	egtsPkgSrResCode = Package{
		ProtocolVersion:  1,
		SecurityKeyID:    0,
		Prefix:           "00",
		Route:            "0",
		EncryptionAlg:    "00",
		Compression:      "0",
		Priority:         "00",
		HeaderLength:     11,
		HeaderEncoding:   0,
		FrameDataLength:  11,
		PacketIdentifier: 14357,
		PacketType:       PtAppdataPacket,
		HeaderCheckSum:   17,
		ServicesFrameData: &ServiceDataSet{
			ServiceDataRecord{
				RecordLength:             4,
				RecordNumber:             14357,
				SourceServiceOnDevice:    "0",
				RecipientServiceOnDevice: "0",
				Group:                    "1",
				RecordProcessingPriority: "00",
				TimeFieldExists:          "0",
				EventIDFieldExists:       "0",
				ObjectIDFieldExists:      "0",
				SourceServiceType:        AuthService,
				RecipientServiceType:     AuthService,
				RecordDataSet: RecordDataSet{
					RecordData{
						SubrecordType:   SrResultCodeType,
						SubrecordLength: 1,
						SubrecordData: &SrResultCode{
							ResultCode: egtsPcOk,
						},
					},
				},
			},
		},
		ServicesFrameDataCheckSum: 48188,
	}

	testEgtsPkgSrResCodeBytes = []byte{0x01, 0x00, 0x00, 0x0B, 0x00, 0x0B, 0x00, 0x15, 0x38, 0x01, 0x11, 0x04, 0x00,
		0x15, 0x38, 0x20, 0x01, 0x01, 0x09, 0x01, 0x00, 0x00, 0x3C, 0xBC}
)

func TestEgtsPkgSrResCode_Encode(t *testing.T) {
	pkgBytes, err := egtsPkgSrResCode.Encode()

	if assert.NoError(t, err) {
		assert.Equal(t, pkgBytes, testEgtsPkgSrResCodeBytes)
	}
}

func TestEgtsPkgSrResCode_Decode(t *testing.T) {
	egtsPkg := Package{}

	_, err := egtsPkg.Decode(testEgtsPkgSrResCodeBytes)
	if assert.NoError(t, err) {
		assert.Equal(t, egtsPkg, egtsPkgSrResCode)
	}
}
