package main

import (
	"bytes"
	"reflect"
	"testing"
)

/*
File: 'egts.bin'
Packet data:
 0100030B001300860001B608005F0099020000000101010500B0090200100DCE


EGTS Transport Layer:
---------------------
  Validating result   - 0 (OK)

  Protocol Version    - 1
  Security Key ID     - 0
  Flags               - 00000011b (0x03)
       Prefix         - 00
       Route          -   0
       Encryption Alg -    00
       Compression    -      0
       Priority       -       11 (low)
  Header Length       - 11
  Header Encoding     - 0
  Frame Data Length   - 19
  Packet ID           - 134
  No route info       -
  Header Check Sum    - 0xB6

EGTS Service Layer:
---------------------
  Validating result   - 0 (OK)

  Packet Type         - EGTS_PT_APPDATA
  Service Layer CS    - 0xCE0D

    Service Layer Record:
    ---------------------
    Validating Result    - 0 (OK)

    Record Length               - 8
    Record Number               - 95
    Record flags                -     10011001b (0x99)
        Sourse Service On Device    - 1
        Recipient Service On Device -  0
        Group Flag                  -   0
        Record Processing Priority  -    11 (low)
        Time Field Exists           -      0
        Event ID Field Exists       -       0
        Object ID Field Exists      -        1
    Object Identifier           - 2
    Source Service Type         - 1 (EGTS_AUTH_SERVICE) from ST
    Recipient Service Type      - 1 (EGTS_AUTH_SERVICE)

       Subrecord Data:
       ------------------
       Validating Result   - 165 (Unknown service subrecord type)

       Subrecord Type      - 1 (EGTS_SR_TERM_IDENTITY)
       Subrecord Length    - 5

 */

var (
	testEgtsSrTermIdentityBin = []byte{0xB0, 0x09, 0x02, 0x00, 0x10}
	testEgtsSrTermIdentity    = EgtsSrTermIdentity{
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
	testEgtsSrTermIdentityPkg = EgtsPackage{
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
		PacketType:       egtsPtAppdata,
		HeaderCheckSum:   182,
		ServicesFrameData: &ServiceDataSet{
			ServiceDataRecord{
				RecordLength:             8,
				RecordNumber:             95,
				SourceServiceOnDevice:    "1",
				RecipientServiceOnDevice: "0",
				Group: "0",
				RecordProcessingPriority: "11",
				TimeFieldExists:          "0",
				EventIDFieldExists:       "0",
				ObjectIDFieldExists:      "1",
				ObjectIdentifier:         2,
				SourceServiceType:        egtsAuthService,
				RecipientServiceType:     egtsAuthService,
				RecordDataSet: RecordDataSet{
					RecordData{
						SubrecordType:   egtsSrTermIdentity,
						SubrecordLength: 5,
						SubrecordData: &testEgtsSrTermIdentity,
					},
				},
			},
		},
		ServicesFrameDataCheckSum: 52749,
	}
)

func TestEgtsSrTermIdentity_Encode(t *testing.T) {
	sti, err := testEgtsSrTermIdentity.Encode()
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(sti, testEgtsSrTermIdentityBin) {
		t.Errorf("Байтовые строки не совпадают: %v != %v ", sti, testEgtsSrTermIdentityBin)
	}
}

func TestEgtsSrTermIdentity_Decode(t *testing.T) {
	srTermIdent := EgtsSrTermIdentity{}

	if err := srTermIdent.Decode(testEgtsSrTermIdentityBin); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if !reflect.DeepEqual(srTermIdent, testEgtsSrTermIdentity) {
		t.Errorf("Секция не совпадает %v != %v\n", srTermIdent, testEgtsSrTermIdentity)
	}

}

func TestEgtsSrTermIdentityPkg_Encode(t *testing.T) {
	pkg, err := testEgtsSrTermIdentityPkg.Encode()
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(pkg, testEgtsSrTermIdentityPkgBin) {
		t.Errorf("Байтовые строки не совпадают: %v != %v ", pkg, testEgtsSrTermIdentityPkgBin)
	}
}

func TestEgtsSrTermIdentityPkg_Decode(t *testing.T) {
	srTermIdentPkg := EgtsPackage{}

	if _, err := srTermIdentPkg.Decode(testEgtsSrTermIdentityPkgBin); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if !reflect.DeepEqual(srTermIdentPkg, testEgtsSrTermIdentityPkg) {

	}

}

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

/*
File: 'egts.bin'
Packet data:
0100000B000B001538011104001538200101090100003CBC


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

//[1 0 3 11 0 19 0 134 0 1 182 8 0 95 0 153 2 0 0 0 1 1 1 5 0 0 176 9 2 0 55 102]
//[1 0 3 11 0 19 0 134 0 1 182 8 0 95 0 153 2 0 0 0 1 1 1 5 0 176 9 2 0 55 102]