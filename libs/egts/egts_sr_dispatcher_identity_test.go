package egts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	srDispatcherIdentityPkgBytes = []byte{0x01, 0x00, 0x00, 0x0b, 0x00, 0x0f, 0x00, 0x01, 0x00,
		0x01, 0x06, 0x08, 0x00, 0x00, 0x00, 0x98, 0x01, 0x01, 0x05, 0x05, 0x00, 0x00, 0x47, 0x00,
		0x00, 0x00, 0x51, 0x9d}

	testDispatcherIdentityPkg = Package{
		ProtocolVersion:  1,
		SecurityKeyID:    0,
		Prefix:           "00",
		Route:            "0",
		EncryptionAlg:    "00",
		Compression:      "0",
		Priority:         "00",
		HeaderLength:     11,
		HeaderEncoding:   0,
		FrameDataLength:  15,
		PacketIdentifier: 1,
		PacketType:       PtAppdataPacket,
		HeaderCheckSum:   6,
		ServicesFrameData: &ServiceDataSet{
			{
				RecordLength:             0x08,
				SourceServiceOnDevice:    "1",
				RecipientServiceOnDevice: "0",
				Group:                    "0",
				RecordProcessingPriority: "11",
				TimeFieldExists:          "0",
				EventIDFieldExists:       "0",
				ObjectIDFieldExists:      "0",
				SourceServiceType:        0x01,
				RecipientServiceType:     0x01,
				RecordDataSet: RecordDataSet{
					{
						SubrecordType:   0x05,
						SubrecordLength: 0x05,
						SubrecordData: &SrDispatcherIdentity{
							DispatcherType: 0,
							DispatcherID:   71,
						},
					},
				},
			},
		},
		ServicesFrameDataCheckSum: 40273,
	}
)

func TestEgtsSrDispatcherIdentity_Encode(t *testing.T) {
	authInfoPkg, err := testAuthInfoPkg.Encode()
	if assert.NoError(t, err) {
		assert.Equal(t, authInfoPkg, srAuthInfoPkgBytes)
	}
}

func TestEgtsSrDispatcherIdentity_Decode(t *testing.T) {
	authPkg := Package{}

	if _, err := authPkg.Decode(srDispatcherIdentityPkgBytes); assert.NoError(t, err) {
		assert.Equal(t, authPkg, testDispatcherIdentityPkg)
	}
}
