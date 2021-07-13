package egts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testEgtsSrCountersData = SrCountersData{
		CounterFieldExists1: "0",
		CounterFieldExists2: "0",
		CounterFieldExists3: "0",
		CounterFieldExists4: "0",
		CounterFieldExists5: "0",
		CounterFieldExists6: "0",
		CounterFieldExists7: "1",
		CounterFieldExists8: "1",
		Counter1:            0,
		Counter2:            0,
		Counter3:            0,
		Counter4:            0,
		Counter5:            0,
		Counter6:            0,
		Counter7:            0,
		Counter8:            3,
	}
	testSrCountersDataBytes = []byte{0xC0, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00}
)

func TestEgtsSrCountersData_Encode(t *testing.T) {
	countersBytes, err := testEgtsSrCountersData.Encode()
	if assert.NoError(t, err) {
		assert.Equal(t, countersBytes, testSrCountersDataBytes)
	}
}

func TestEgtsSrCountersData_Decode(t *testing.T) {
	countersData := SrCountersData{}

	if assert.NoError(t, countersData.Decode(testSrCountersDataBytes)) {
		assert.Equal(t, countersData, testEgtsSrCountersData)
	}
}

// проверяем что рекордсет работает правильно с данным типом подзаписи
func TestEgtsSrCountersDataRs(t *testing.T) {
	countersDataRDBytes := append([]byte{0x13, 0x07, 0x00}, testSrCountersDataBytes...)
	countersDataRD := RecordDataSet{
		RecordData{
			SubrecordType:   SrCountersDataType,
			SubrecordLength: testEgtsSrCountersData.Length(),
			SubrecordData:   &testEgtsSrCountersData,
		},
	}
	testStruct := RecordDataSet{}

	testBytes, err := countersDataRD.Encode()
	if assert.NoError(t, err) {
		assert.Equal(t, testBytes, countersDataRDBytes)

		if assert.NoError(t, testStruct.Decode(countersDataRDBytes)) {
			assert.Equal(t, countersDataRD, testStruct)
		}
	}
}
