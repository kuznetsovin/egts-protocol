package egts

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	srAdSensorsDataBytes = []byte{0x01, 0x0F, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	testEgtsSrAdSensorsData = SrAdSensorsData{
		DigitalInputsOctetExists1:     "1",
		DigitalInputsOctetExists2:     "0",
		DigitalInputsOctetExists3:     "0",
		DigitalInputsOctetExists4:     "0",
		DigitalInputsOctetExists5:     "0",
		DigitalInputsOctetExists6:     "0",
		DigitalInputsOctetExists7:     "0",
		DigitalInputsOctetExists8:     "0",
		DigitalOutputs:                15,
		AnalogSensorFieldExists1:      "1",
		AnalogSensorFieldExists2:      "1",
		AnalogSensorFieldExists3:      "1",
		AnalogSensorFieldExists4:      "1",
		AnalogSensorFieldExists5:      "1",
		AnalogSensorFieldExists6:      "1",
		AnalogSensorFieldExists7:      "1",
		AnalogSensorFieldExists8:      "1",
		AdditionalDigitalInputsOctet1: 0,
		AnalogSensor1:                 0,
		AnalogSensor2:                 0,
		AnalogSensor3:                 0,
		AnalogSensor4:                 0,
		AnalogSensor5:                 0,
		AnalogSensor6:                 0,
		AnalogSensor7:                 0,
		AnalogSensor8:                 0,
	}
)

func TestEgtsSrAdSensorsData_Encode(t *testing.T) {
	sensDataBytes, err := testEgtsSrAdSensorsData.Encode()
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(sensDataBytes, srAdSensorsDataBytes) {
		t.Errorf("Байтовые строки не совпадают: %v != %v ", sensDataBytes, srAdSensorsDataBytes)
	}
}

func TestEgtsSrAdSensorsData_Decode(t *testing.T) {
	adSensData := SrAdSensorsData{}

	if err := adSensData.Decode(srAdSensorsDataBytes); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if diff := cmp.Diff(adSensData, testEgtsSrAdSensorsData); diff != "" {
		t.Errorf("Записи не совпадают: (-нужно +сейчас)\n%s", diff)
	}
}

// проверяем что рекордсет работает правильно с данным типом подзаписи
func TestEgtsSrAdSensorsDataaRs(t *testing.T) {
	adSensDataRDBytes := append([]byte{0x12, 0x1C, 0x00}, srAdSensorsDataBytes...)
	adSensDataRD := RecordDataSet{
		RecordData{
			SubrecordType:   SrAdSensorsDataType,
			SubrecordLength: testEgtsSrAdSensorsData.Length(),
			SubrecordData:   &testEgtsSrAdSensorsData,
		},
	}
	testStruct := RecordDataSet{}

	testBytes, err := adSensDataRD.Encode()
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(testBytes, adSensDataRDBytes) {
		t.Errorf("Байтовые строки не совпадают: %v != %v ", testBytes, adSensDataRDBytes)
	}

	if err = testStruct.Decode(adSensDataRDBytes); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if diff := cmp.Diff(adSensDataRD, testStruct); diff != "" {
		t.Errorf("Записи не совпадают: (-нужно +сейчас)\n%s", diff)
	}
}
