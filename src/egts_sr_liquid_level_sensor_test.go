package main

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"testing"
)

var (
	testSrLiquidLevelSensor = EgtsSrLiquidLevelSensor{
		LiquidLevelSensorErrorFlag: "0",
		LiquidLevelSensorValueUnit: "00",
		RawDataFlag:                "0",
		LiquidLevelSensorNumber:    3,
		ModuleAddress:              1,
		LiquidLevelSensorData:      0,
	}
	testSrLiquidLevelSensorBytes = []byte{0x03, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00}
)

func TestEgtsSrLiquidLevelSensor_Encode(t *testing.T) {

	pkgBytes, err := testSrLiquidLevelSensor.Encode()
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(pkgBytes, testSrLiquidLevelSensorBytes) {
		t.Errorf("Байтовые строки не совпадают: %v != %v ", pkgBytes, testSrLiquidLevelSensorBytes)
	}
}

func TestEgtsSrLiquidLevelSensor_Decode(t *testing.T) {
	liquidLev := EgtsSrLiquidLevelSensor{}

	if err := liquidLev.Decode(testSrLiquidLevelSensorBytes); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if diff := cmp.Diff(liquidLev, testSrLiquidLevelSensor); diff != "" {
		t.Errorf("Записи не совпадают: (-нужно +сейчас)\n%s", diff)
	}
}

// проверяем что рекордсет работает правильно с данным типом подзаписи
func TestEgtsSrLiquidLevelSensorRs(t *testing.T) {
	liquidLevelRDRDBytes := append([]byte{0x1B, 0x07, 0x00}, testSrLiquidLevelSensorBytes...)
	liquidLevelRD := RecordDataSet{
		RecordData{
			SubrecordType:   egtsSrLiquidLevelSensor,
			SubrecordLength: 7,
			SubrecordData:   &testSrLiquidLevelSensor,
		},
	}
	testStruct := RecordDataSet{}

	testBytes, err := liquidLevelRD.Encode()
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(testBytes, liquidLevelRDRDBytes) {
		t.Errorf("Байтовые строки не совпадают: %v != %v ", testBytes, liquidLevelRDRDBytes)
	}

	if err = testStruct.Decode(liquidLevelRDRDBytes); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if diff := cmp.Diff(liquidLevelRD, testStruct); diff != "" {
		t.Errorf("Записи не совпадают: (-нужно +сейчас)\n%s", diff)
	}
}
