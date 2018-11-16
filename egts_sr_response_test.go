package main

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"testing"
)

/*
File: 'egts.bin'
Packet data:
0100000B0010008600001886000006005F002001010003005F00001373


EGTS Transport Layer:
---------------------
 Validating result   - 0 (OK)

 Protocol Version    - 1
 Security Key ID     - 0
 Flags               - 00000000b (0x00)
	  Prefix         - 00
	  Route          -   0
	  Encryption Alg -    00
	  Compression    -      0
	  Priority       -       00 (the highest)
 Header Length       - 11
 Header Encoding     - 0
 Frame Data Length   - 16
 Packet ID           - 134
 No route info       -
 Header Check Sum    - 0x18

EGTS Service Layer:
---------------------
 Validating result   - 0 (OK)

 Packet Type         - EGTS_PT_RESPONSE
 Service Layer CS    - 0x7313
 Responded Packet ID - 134
 Processing Result   - 0 (OK)

   Service Layer Record:
   ---------------------
   Validating Result    - 0 (OK)

   Record Length               - 6
   Record Number               - 95
   Record flags                -     00100000b (0x20)
	   Sourse Service On Device    - 0
	   Recipient Service On Device -  0
	   Group Flag                  -   1
	   Record Processing Priority  -    00 (the highest)
	   Time Field Exists           -      0
	   Event ID Field Exists       -       0
	   Object ID Field Exists      -        0
   Source Service Type         - 1 (EGTS_AUTH_SERVICE)
   Recipient Service Type      - 1 (EGTS_AUTH_SERVICE)

	  Subrecord Data:
	  ------------------
	  Validating Result   - 0 (OK)

	  Subrecord Type      - 0 (EGTS_SR_RESPONSE)
	  Subrecord Length    - 3
	  Confirmed Record Number- 95
	  Record Status          - 0 (OK)
*/

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
		PacketType:       0,
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
					Group: "1",
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
