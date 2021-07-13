package egts

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
	if assert.NoError(t, err) {
		assert.Equal(t, pkgBytes, testSrStateDataBytes)
	}
}

func TestEgtsPkgSrStateData_Decode(t *testing.T) {
	stStateData := SrStateData{}

	if assert.NoError(t, stStateData.Decode(testSrStateDataBytes)) {
		assert.Equal(t, stStateData, testEgtsSrStateData)
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
	if assert.NoError(t, err) {
		assert.Equal(t, testBytes, stateDataRDBytes)

		if assert.NoError(t, testStruct.Decode(stateDataRDBytes)) {
			assert.Equal(t, stateDataRD, testStruct)
		}
	}
}
