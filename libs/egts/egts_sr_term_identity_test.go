package egts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testEgtsSrTermIdentityBin = []byte{0xB0, 0x09, 0x02, 0x00, 0x10}
	testEgtsSrTermIdentity    = SrTermIdentity{
		TerminalIdentifier: 133552,
		MNE:                "0",
		BSE:                "0",
		NIDE:               "0",
		SSRA:               "1",
		LNGCE:              "0",
		IMSIE:              "0",
		IMEIE:              "0",
		HDIDE:              "0",
	}
	testEgtsSrTermIdentityPkgBin = []byte{0x01, 0x00, 0x03, 0x0B, 0x00, 0x13, 0x00, 0x86, 0x00, 0x01, 0xB6, 0x08, 0x00,
		0x5F, 0x00, 0x99, 0x02, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01, 0x05, 0x00, 0xB0, 0x09, 0x02, 0x00, 0x10, 0x0D, 0xCE}
	testEgtsSrTermIdentityPkg = Package{
		ProtocolVersion:  1,
		SecurityKeyID:    0,
		Prefix:           "00",
		Route:            "0",
		EncryptionAlg:    "00",
		Compression:      "0",
		Priority:         "11",
		HeaderLength:     11,
		HeaderEncoding:   0,
		FrameDataLength:  19,
		PacketIdentifier: 134,
		PacketType:       PtAppdataPacket,
		HeaderCheckSum:   182,
		ServicesFrameData: &ServiceDataSet{
			ServiceDataRecord{
				RecordLength:             8,
				RecordNumber:             95,
				SourceServiceOnDevice:    "1",
				RecipientServiceOnDevice: "0",
				Group:                    "0",
				RecordProcessingPriority: "11",
				TimeFieldExists:          "0",
				EventIDFieldExists:       "0",
				ObjectIDFieldExists:      "1",
				ObjectIdentifier:         2,
				SourceServiceType:        AuthService,
				RecipientServiceType:     AuthService,
				RecordDataSet: RecordDataSet{
					RecordData{
						SubrecordType:   SrTermIdentityType,
						SubrecordLength: 5,
						SubrecordData:   &testEgtsSrTermIdentity,
					},
				},
			},
		},
		ServicesFrameDataCheckSum: 52749,
	}
)

func TestEgtsSrTermIdentity_Encode(t *testing.T) {
	sti, err := testEgtsSrTermIdentity.Encode()

	if assert.NoError(t, err) {
		assert.Equal(t, sti, testEgtsSrTermIdentityBin)
	}
}

func TestEgtsSrTermIdentity_Decode(t *testing.T) {
	srTermIdent := SrTermIdentity{}

	if assert.NoError(t, srTermIdent.Decode(testEgtsSrTermIdentityBin)) {
		assert.Equal(t, srTermIdent, testEgtsSrTermIdentity)
	}
}

func TestEgtsSrTermIdentityPkg_Encode(t *testing.T) {
	pkg, err := testEgtsSrTermIdentityPkg.Encode()

	if assert.NoError(t, err) {
		assert.Equal(t, pkg, testEgtsSrTermIdentityPkgBin)
	}
}

func TestEgtsSrTermIdentityPkg_Decode(t *testing.T) {
	srTermIdentPkg := Package{}

	_, err := srTermIdentPkg.Decode(testEgtsSrTermIdentityPkgBin)
	if assert.NoError(t, err) {
		assert.Equal(t, srTermIdentPkg, testEgtsSrTermIdentityPkg)
	}
}
