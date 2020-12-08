package egts

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"testing"
)

var (
	testServiceDataRecordBytes = []byte{0x18, 0x00, 0x61, 0x00, 0x99, 0xB0, 0x09, 0x02, 0x00, 0x02, 0x02}
	testServiceDataRecord      = ServiceDataSet{
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
	sdr := ServiceDataSet{}

	if err := sdr.Decode(testServiceDataRecordBytes); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if diff := cmp.Diff(sdr, testServiceDataRecord); diff != "" {
		t.Errorf("Записи не совпадают: (-нужно +сейчас)\n%s", diff)
	}
}
