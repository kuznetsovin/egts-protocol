package egts

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	testEgtsSrStateData = SrStateData{
		State:                  2,
		MainPowerSourceVoltage: 127,
		BackUpBatteryVoltage:   0,
		InternalBatteryVoltage: 41,
		NMS:                    "1",
		IBU:                    "0",
		BBU:                    "0",
	}
	testSrStateDataBytes = []byte{0x02, 0x7F, 0x00, 0x29, 0x04}
)

func TestEgtsPkgSrStateData_Encode(t *testing.T) {

	pkgBytes, err := testEgtsSrStateData.Encode()
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(pkgBytes, testSrStateDataBytes) {
		t.Errorf("Байтовые строки не совпадают: %v != %v ", pkgBytes, testSrStateDataBytes)
	}
}

func TestEgtsPkgSrStateData_Decode(t *testing.T) {
	stStateData := SrStateData{}

	if err := stStateData.Decode(testSrStateDataBytes); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if diff := cmp.Diff(stStateData, testEgtsSrStateData); diff != "" {
		t.Errorf("Записи не совпадают: (-нужно +сейчас)\n%s", diff)
	}
}

// проверяем что рекордсет работает правильно с данным типом подзаписи
func TestEgtsSrStateDataRs(t *testing.T) {
	stateDataRDBytes := append([]byte{0x14, 0x05, 0x00}, testSrStateDataBytes...)
	stateDataRD := RecordDataSet{
		RecordData{
			SubrecordType:   SrType20,
			SubrecordLength: 5,
			SubrecordData:   &testEgtsSrStateData,
		},
	}
	testStruct := RecordDataSet{}

	testBytes, err := stateDataRD.Encode()
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(testBytes, stateDataRDBytes) {
		t.Errorf("Байтовые строки не совпадают: %v != %v ", testBytes, stateDataRDBytes)
	}

	if err = testStruct.Decode(stateDataRDBytes); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if diff := cmp.Diff(testStruct, stateDataRD); diff != "" {
		t.Errorf("Записи не совпадают: (-нужно +сейчас)\n%s", diff)
	}
}
