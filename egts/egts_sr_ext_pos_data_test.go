package egts

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"testing"
)

var (
	extPosDataBytes      = []byte{0x0E, 0x32, 0x00, 0x00, 0x00, 0x0C}
	testEgtsSrExtPosData = SrExtPosData{
		NavigationSystemFieldExists:   "0",
		SatellitesFieldExists:         "1",
		PdopFieldExists:               "1",
		HdopFieldExists:               "1",
		VdopFieldExists:               "0",
		HorizontalDilutionOfPrecision: 50,
		PositionDilutionOfPrecision:   0,
		Satellites:                    12,
	}
)

func TestEgtsSrExtPosData_Encode(t *testing.T) {
	posDataBytes, err := testEgtsSrExtPosData.Encode()
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(posDataBytes, extPosDataBytes) {
		t.Errorf("Байтовые строки не совпадают: %v != %v ", posDataBytes, extPosDataBytes)
	}
}

func TestEgtsSrExtPosData_Decode(t *testing.T) {
	extPosData := SrExtPosData{}

	if err := extPosData.Decode(extPosDataBytes); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if diff := cmp.Diff(extPosData, testEgtsSrExtPosData); diff != "" {
		t.Errorf("Записи не совпадают: (-нужно +сейчас)\n%s", diff)
	}
}

// проверяем что рекордсет работает правильно с данным типом подзаписи
func TestEgtsSrExtPosDataRs(t *testing.T) {
	extPosDataRDBytes := append([]byte{0x11, 0x06, 0x00}, extPosDataBytes...)
	extPosDataRD := RecordDataSet{
		RecordData{
			SubrecordType:   SrExtPosDataType,
			SubrecordLength: 6,
			SubrecordData:   &testEgtsSrExtPosData,
		},
	}
	testStruct := RecordDataSet{}

	testBytes, err := extPosDataRD.Encode()
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(testBytes, extPosDataRDBytes) {
		t.Errorf("Байтовые строки не совпадают: %v != %v ", testBytes, extPosDataRDBytes)
	}

	if err = testStruct.Decode(extPosDataRDBytes); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if diff := cmp.Diff(extPosDataRD, testStruct); diff != "" {
		t.Errorf("Записи не совпадают: (-нужно +сейчас)\n%s", diff)
	}
}
