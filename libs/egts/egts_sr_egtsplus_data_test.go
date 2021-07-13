package egts

import (
	"github.com/stretchr/testify/assert"
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
	if assert.NoError(t, err) {
		assert.Equal(t, egtsPlusBytes, srEgtsPlusBytes)
	}
}

func TestStorageRecord_Decode(t *testing.T) {
	egtsPlus := StorageRecord{}

	if assert.NoError(t, egtsPlus.Decode(srEgtsPlusBytes)) {
		assert.Equal(t, egtsPlus, testEgtsPlusData)
	}
}

// проверяем что рекордсет работает правильно с данным типом подзаписи
func TestStorageRecordRs(t *testing.T) {
	egtsPlusRDBytes := append([]byte{0x0F, 0x1A, 0x00}, srEgtsPlusBytes...)
	egtsPlusDataRD := RecordDataSet{
		RecordData{
			SubrecordType:   SrEgtsPlusDataType,
			SubrecordLength: testEgtsPlusData.Length(),
			SubrecordData:   &testEgtsPlusData,
		},
	}
	testStruct := RecordDataSet{}

	testBytes, err := egtsPlusDataRD.Encode()
	if assert.NoError(t, err) {
		assert.Equal(t, testBytes, egtsPlusRDBytes)

		if assert.NoError(t, testStruct.Decode(egtsPlusRDBytes)) {
			assert.Equal(t, egtsPlusDataRD, testStruct)
		}
	}
}
