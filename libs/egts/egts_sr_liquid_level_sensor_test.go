package egts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testSrLiquidLevelSensor = SrLiquidLevelSensor{
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

	if assert.NoError(t, err) {
		assert.Equal(t, pkgBytes, testSrLiquidLevelSensorBytes)
	}
}

func TestEgtsSrLiquidLevelSensor_Decode(t *testing.T) {
	liquidLev := SrLiquidLevelSensor{}

	if assert.NoError(t, liquidLev.Decode(testSrLiquidLevelSensorBytes)) {
		assert.Equal(t, liquidLev, testSrLiquidLevelSensor)
	}
}

// проверяем что рекордсет работает правильно с данным типом подзаписи
func TestEgtsSrLiquidLevelSensorRs(t *testing.T) {
	liquidLevelRDRDBytes := append([]byte{0x1B, 0x07, 0x00}, testSrLiquidLevelSensorBytes...)
	liquidLevelRD := RecordDataSet{
		RecordData{
			SubrecordType:   SrLiquidLevelSensorType,
			SubrecordLength: 7,
			SubrecordData:   &testSrLiquidLevelSensor,
		},
	}
	testStruct := RecordDataSet{}

	testBytes, err := liquidLevelRD.Encode()
	if assert.NoError(t, err) {
		assert.Equal(t, testBytes, liquidLevelRDRDBytes)

		if assert.NoError(t, testStruct.Decode(liquidLevelRDRDBytes)) {
			assert.Equal(t, liquidLevelRD, testStruct)
		}
	}
}
