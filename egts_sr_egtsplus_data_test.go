package main

import (
	"bytes"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

var (
	srEgtsPlusBytes = []byte{0x08, 0xFB, 0xD4, 0x03, 0x15, 0x3B, 0x46, 0x5F, 0x5C, 0x25, 0x00, 0x00, 0x00, 0x00,
		0x82, 0x01, 0x04, 0x08, 0x01, 0x30, 0x4F, 0x8A, 0x01, 0x02, 0x08, 0x01}

	rn     = uint32(60027)
	ts     = uint32(1549747771)
	sf     = uint32(0)
	sn     = uint32(1)
	t      = int32(-40)
	scltde = SensCanLogTmpDataExt{
		SensNum: &sn,
	}
	scld = SensCanLogData{
		SensNum:           &sn,
		EngineTemperature: &t,
	}
	testEgtsPlusData = StorageRecord{
		RecordNumber:         &rn,
		TimeStamp:            &ts,
		StatusFlags:          &sf,
		SensCanLogData:       []*SensCanLogData{&scld},
		SensCanLogTmpDataExt: []*SensCanLogTmpDataExt{&scltde},
	}
)

func TestStorageRecord_Encode(t *testing.T) {
	egtsPlusBytes, err := testEgtsPlusData.Encode()
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(egtsPlusBytes, srEgtsPlusBytes) {
		t.Errorf("Байтовые строки не совпадают: %v != %v ", egtsPlusBytes, srEgtsPlusBytes)
	}
}

func TestStorageRecord_Decode(t *testing.T) {
	egtsPlus := StorageRecord{}

	if err := egtsPlus.Decode(srEgtsPlusBytes); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	fmt.Println(*egtsPlus.RecordNumber)
	fmt.Println(*egtsPlus.TimeStamp)
	fmt.Println(*egtsPlus.StatusFlags)

	if diff := cmp.Diff(egtsPlus, testEgtsPlusData); diff != "" {
		t.Errorf("Записи не совпадают: (-нужно +сейчас)\n%s", diff)
	}
}

// проверяем что рекордсет работает правильно с данным типом подзаписи
func TestStorageRecordRs(t *testing.T) {
	egtsPlusRDBytes := append([]byte{0x0F, 0x1A, 0x00}, srEgtsPlusBytes...)
	egtsPlusDataRD := RecordDataSet{
		RecordData{
			SubrecordType:   egtsSrEgtsPlusData,
			SubrecordLength: testEgtsPlusData.Length(),
			SubrecordData:   &testEgtsPlusData,
		},
	}
	testStruct := RecordDataSet{}

	testBytes, err := egtsPlusDataRD.Encode()
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(testBytes, egtsPlusRDBytes) {
		t.Errorf("Байтовые строки не совпадают: %v != %v ", testBytes, egtsPlusRDBytes)
	}

	if err = testStruct.Decode(egtsPlusRDBytes); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if diff := cmp.Diff(egtsPlusDataRD, testStruct); diff != "" {
		t.Errorf("Записи не совпадают: (-нужно +сейчас)\n%s", diff)
	}
}
