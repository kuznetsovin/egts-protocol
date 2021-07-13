package egts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	srAuthInfoPkgBytes = []byte{0x01, 0x00, 0x01, 0x0B, 0x00, 0x33, 0x00, 0x01, 0x00, 0x01, 0xEA, 0x28, 0x00, 0x00,
		0x00, 0x0C, 0x44, 0x98, 0x11, 0x11, 0x01, 0x01, 0x07, 0x25, 0x00, 0x38, 0x30, 0x30, 0x00, 0x45, 0x46, 0x32, 0x38,
		0x34, 0x45, 0x37, 0x41, 0x45, 0x33, 0x35, 0x31, 0x44, 0x36, 0x44, 0x46, 0x39, 0x32, 0x43, 0x45, 0x33, 0x32, 0x33,
		0x44, 0x37, 0x34, 0x41, 0x44, 0x32, 0x45, 0x42, 0x33, 0x00, 0x92, 0x61,
	}

	testAuthInfoPkg = Package{
		ProtocolVersion:  1,
		SecurityKeyID:    0,
		Prefix:           "00",
		Route:            "0",
		EncryptionAlg:    "00",
		Compression:      "0",
		Priority:         "01",
		HeaderLength:     11,
		HeaderEncoding:   0,
		FrameDataLength:  51,
		PacketIdentifier: 1,
		PacketType:       PtAppdataPacket,
		HeaderCheckSum:   234,
		ServicesFrameData: &ServiceDataSet{
			ServiceDataRecord{
				RecordLength:             40,
				RecordNumber:             0,
				SourceServiceOnDevice:    "0",
				RecipientServiceOnDevice: "0",
				Group:                    "0",
				RecordProcessingPriority: "01",
				TimeFieldExists:          "1",
				EventIDFieldExists:       "0",
				ObjectIDFieldExists:      "0",
				Time:                     286365764,
				SourceServiceType:        AuthService,
				RecipientServiceType:     AuthService,
				RecordDataSet: RecordDataSet{
					RecordData{
						SubrecordType:   SrAuthInfoType,
						SubrecordLength: 37,
						SubrecordData: &SrAuthInfo{
							UserName:     "800",
							UserPassword: "EF284E7AE351D6DF92CE323D74AD2EB3",
						},
					},
				},
			},
		},
		ServicesFrameDataCheckSum: 24978,
	}
)

func TestEgtsSrAuthInfo_Encode(t *testing.T) {
	authInfoPkg, err := testAuthInfoPkg.Encode()
	if assert.NoError(t, err) {
		assert.Equal(t, authInfoPkg, srAuthInfoPkgBytes)
	}
}

func TestEgtsSrAuthInfo_Decode(t *testing.T) {
	authPkg := Package{}

	if _, err := authPkg.Decode(srAuthInfoPkgBytes); assert.NoError(t, err) {
		assert.Equal(t, authPkg, testAuthInfoPkg)
	}
}
