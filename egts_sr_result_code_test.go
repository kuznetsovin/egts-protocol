package main

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"testing"
)

/*
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
 Frame Data Length   - 11
 Packet ID           - 14357
 No route info       -
 Header Check Sum    - 0x11

EGTS Service Layer:
---------------------
 Validating result   - 0 (OK)

 Packet Type         - EGTS_PT_APPDATA
 Service Layer CS    - 0xBC3C

   Service Layer Record:
   ---------------------
   Validating Result    - 0 (OK)

   Record Length               - 4
   Record Number               - 14357
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

	  Subrecord Type      - 9 (EGTS_SR_RESULT_CODE)
	  Subrecord Length    - 1
	  Result Code            - 0 (OK)

*/

var (
	egtsPkgSrResCode = EgtsPackage{
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
		PacketType:       egtsPtAppdata,
		HeaderCheckSum:   17,
		ServicesFrameData: &ServiceDataSet{
			ServiceDataRecord{
				RecordLength:             4,
				RecordNumber:             14357,
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
						SubrecordType:   egtsSrResultCode,
						SubrecordLength: 1,
						SubrecordData: &EgtsSrResultCode{
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
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(pkgBytes, testEgtsPkgSrResCodeBytes) {
		t.Errorf("Байтовые строки не совпадают: %v != %v ", pkgBytes, testEgtsPkgSrResCodeBytes)
	}
}

func TestEgtsPkgSrResCode_Decode(t *testing.T) {
	egtsPkg := EgtsPackage{}

	if _, err := egtsPkg.Decode(testEgtsPkgSrResCodeBytes); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if diff := cmp.Diff(egtsPkg, egtsPkgSrResCode); diff != "" {
		t.Errorf("Записи не совпадают: (-нужно +сейчас)\n%s", diff)
	}
}
