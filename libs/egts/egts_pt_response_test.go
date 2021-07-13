package egts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	egtsPkgResp = Package{
		ProtocolVersion:  1,
		SecurityKeyID:    0,
		Prefix:           "00",
		Route:            "0",
		EncryptionAlg:    "00",
		Compression:      "0",
		Priority:         "11",
		HeaderLength:     11,
		HeaderEncoding:   0,
		FrameDataLength:  3,
		PacketIdentifier: 137,
		PacketType:       PtResponsePacket,
		HeaderCheckSum:   74,
		ServicesFrameData: &PtResponse{
			ResponsePacketID: 14357,
			ProcessingResult: egtsPcOk,
		},
		ServicesFrameDataCheckSum: 59443,
	}
	testEgtsPkgBytes = []byte{0x01, 0x00, 0x03, 0x0B, 0x00, 0x03, 0x00, 0x89, 0x00, 0x00, 0x4A, 0x15, 0x38, 0x00, 0x33, 0xE8}
)

func TestEgtsPkgResp_Encode(t *testing.T) {

	posDataBytes, err := egtsPkgResp.Encode()
	if assert.NoError(t, err) {
		assert.Equal(t, posDataBytes, testEgtsPkgBytes)
	}
}

func TestEgtsPkgResp_Decode(t *testing.T) {
	egtsPkg := Package{}

	_, err := egtsPkg.Decode(testEgtsPkgBytes)
	if assert.NoError(t, err) {
		assert.Equal(t, egtsPkg, egtsPkgResp)
	}
}
