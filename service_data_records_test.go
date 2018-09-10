package main

import (
	"bytes"
	"reflect"
	"testing"
)

/*
Packet data:
 0100030B0023008A0001491800610099B00902000202101500D53F01106F1C059E7AB53C3501D0872C0100000000CC27

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
  Frame Data Length   - 35
  Packet ID           - 138
  No route info       -
  Header Check Sum    - 0x49

EGTS Service Layer:
---------------------
  Validating result   - 0 (OK)

  Packet Type         - EGTS_PT_APPDATA
  Service Layer CS    - 0x27CC

    Service Layer Record:
    ---------------------
    Validating Result    - 0 (OK)

    Record Length               - 24
    Record Number               - 97
    Record flags                -     10011001b (0x99)
        Sourse Service On Device    - 1
        Recipient Service On Device -  0
        Group Flag                  -   0
        Record Processing Priority  -    11 (low)
        Time Field Exists           -      0
        Event ID Field Exists       -       0
        Object ID Field Exists      -        1
    Object Identifier           - 133552
    Source Service Type         - 2 (EGTS_TELEDATA_SERVICE) from ST
    Recipient Service Type      - 2 (EGTS_TELEDATA_SERVICE)

       Subrecord Data:
       ------------------
       Validating Result   - 150 (Unknown service)

       Subrecord Type      - 16 (unspecified)
       Subrecord Length    - 21
*/

var (
	testServiceDataRecordBytes = []byte{0x18, 0x00, 0x61, 0x00, 0x99, 0xB0, 0x09, 0x02, 0x00, 0x02, 0x02}
	testServiceDataRecord      = ServiceDataRecord{
		RecordLength:             24,
		RecordNumber:             97,
		SourceServiceOnDevice:    "1",
		RecipientServiceOnDevice: "0",
		Group: "0",
		RecordProcessingPriority: "11",
		TimeFieldExists:          "0",
		EventIDFieldExists:       "0",
		ObjectIDFieldExists:      "1",
		ObjectIdentifier:         133552,
		SourceServiceType:        2,
		RecipientServiceType:     2,
	}
)

func TestServiceDataRecord_Encode(t *testing.T) {
	sdr, err := testServiceDataRecord.Encode()
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(sdr, testServiceDataRecordBytes) {
		t.Errorf("Байтовые строки не совпадают: %v != %v ", sdr, testServiceDataRecordBytes)
	}
}

func TestServiceDataRecord_Decode(t *testing.T) {
	sdr := ServiceDataRecord{}

	if err := sdr.Decode(testServiceDataRecordBytes); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if !reflect.DeepEqual(sdr, testServiceDataRecord) {
		t.Errorf("Запись не совпадают: %v != %v ", sdr, testServiceDataRecord)
	}

}
