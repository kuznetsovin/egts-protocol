package egts

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestServiceDataRecord_Encode(t *testing.T) {
	testServiceDataRecord := ServiceDataSet{
		ServiceDataRecord{
			RecordLength:             0,
			RecordNumber:             97,
			SourceServiceOnDevice:    "1",
			RecipientServiceOnDevice: "0",
			Group:                    "0",
			RecordProcessingPriority: "11",
			TimeFieldExists:          "0",
			EventIDFieldExists:       "0",
			ObjectIDFieldExists:      "1",
			ObjectIdentifier:         133552,
			SourceServiceType:        2,
			RecipientServiceType:     2,
		},
	}
	testServiceDataRecordBytes := []byte{0x00, 0x00, 0x61, 0x00, 0x99, 0xB0, 0x09, 0x02, 0x00, 0x02, 0x02}

	sdr, err := testServiceDataRecord.Encode()
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(sdr, testServiceDataRecordBytes) {
		t.Errorf("Байтовые строки не совпадают: %v != %v ", sdr, testServiceDataRecordBytes)
	}
}

func TestServiceDataRecord_Decode(t *testing.T) {
	sdr := ServiceDataSet{}
	testServiceDataRecord := ServiceDataSet{
		ServiceDataRecord{
			RecordLength:             24,
			RecordNumber:             97,
			SourceServiceOnDevice:    "1",
			RecipientServiceOnDevice: "0",
			Group:                    "0",
			RecordProcessingPriority: "11",
			TimeFieldExists:          "0",
			EventIDFieldExists:       "0",
			ObjectIDFieldExists:      "1",
			ObjectIdentifier:         133552,
			SourceServiceType:        2,
			RecipientServiceType:     2,
		},
	}
	testServiceDataRecordBytes := []byte{0x18, 0x00, 0x61, 0x00, 0x99, 0xB0, 0x09, 0x02, 0x00, 0x02, 0x02}

	if err := sdr.Decode(testServiceDataRecordBytes); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if diff := cmp.Diff(sdr, testServiceDataRecord); diff != "" {
		t.Errorf("Записи не совпадают: (-нужно +сейчас)\n%s", diff)
	}
}
