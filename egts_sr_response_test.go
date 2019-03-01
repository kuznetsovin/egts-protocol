package main

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"testing"
)

var (
	egtsPkgSrResp = EgtsPackage{
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
		PacketType:       egtsPtResponse,
		HeaderCheckSum:   24,
		ServicesFrameData: &EgtsPtResponse{
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
					SourceServiceType:        egtsAuthService,
					RecipientServiceType:     egtsAuthService,
					RecordDataSet: RecordDataSet{
						RecordData{
							SubrecordType:   egtsSrRecordResponse,
							SubrecordLength: 3,
							SubrecordData: &EgtsSrResponse{
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
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(pkgBytes, testEgtsPkgSrRespBytes) {
		t.Errorf("Байтовые строки не совпадают: %v != %v ", pkgBytes, testEgtsPkgSrRespBytes)
	}
}

func TestEgtsPkgSrResp_Decode(t *testing.T) {
	egtsPkg := EgtsPackage{}

	if _, err := egtsPkg.Decode(testEgtsPkgSrRespBytes); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if diff := cmp.Diff(egtsPkg, egtsPkgSrResp); diff != "" {
		t.Errorf("Записи не совпадают: (-нужно +сейчас)\n%s", diff)
	}
}
