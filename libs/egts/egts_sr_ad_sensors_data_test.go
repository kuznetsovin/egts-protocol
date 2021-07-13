package egts

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
	if assert.NoError(t, err) {
		assert.Equal(t, sensDataBytes, srAdSensorsDataBytes)
	}
}

func TestEgtsSrAdSensorsData_Decode(t *testing.T) {
	adSensData := SrAdSensorsData{}

	if assert.NoError(t, adSensData.Decode(srAdSensorsDataBytes)) {
		assert.Equal(t, adSensData, testEgtsSrAdSensorsData)
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
	if assert.NoError(t, err) {
		assert.Equal(t, testBytes, adSensDataRDBytes)

		if assert.NoError(t, testStruct.Decode(adSensDataRDBytes)) {
			assert.Equal(t, adSensDataRD, testStruct)
		}
	}
}
