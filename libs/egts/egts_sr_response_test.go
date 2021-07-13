package egts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	egtsPkgSrResp = Package{
		ProtocolVersion:  1,
		SecurityKeyID:    0,
		Prefix:           "00",
		Route:            "0",
		EncryptionAlg:    "00",
		Compression:      "0",
		Priority:         "00",
		HeaderLength:     11,
		HeaderEncoding:   0,
		FrameDataLength:  16,
		PacketIdentifier: 134,
		PacketType:       PtResponsePacket,
		HeaderCheckSum:   24,
		ServicesFrameData: &PtResponse{
			ResponsePacketID: 134,
			ProcessingResult: 0,
			SDR: &ServiceDataSet{
				ServiceDataRecord{
					RecordLength:             6,
					RecordNumber:             95,
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
							SubrecordType:   SrRecordResponseType,
							SubrecordLength: 3,
							SubrecordData: &SrResponse{
								ConfirmedRecordNumber: 95,
								RecordStatus:          egtsPcOk,
							},
						},
					},
				},
			},
		},
		ServicesFrameDataCheckSum: 29459,
	}
	testEgtsPkgSrRespBytes = []byte{0x01, 0x00, 0x00, 0x0B, 0x00, 0x10, 0x00, 0x86, 0x00, 0x00, 0x18, 0x86, 0x00, 0x00,
		0x06, 0x00, 0x5F, 0x00, 0x20, 0x01, 0x01, 0x00, 0x03, 0x00, 0x5F, 0x00, 0x00, 0x13, 0x73}
)

func TestEgtsPkgSrResp_Encode(t *testing.T) {
	pkgBytes, err := egtsPkgSrResp.Encode()

	if assert.NoError(t, err) {
		assert.Equal(t, pkgBytes, testEgtsPkgSrRespBytes)
	}
}

func TestEgtsPkgSrResp_Decode(t *testing.T) {
	egtsPkg := Package{}

	_, err := egtsPkg.Decode(testEgtsPkgSrRespBytes)
	if assert.NoError(t, err) {
		assert.Equal(t, egtsPkg, egtsPkgSrResp)
	}
}
