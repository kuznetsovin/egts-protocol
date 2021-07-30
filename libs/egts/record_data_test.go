package egts

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	testRecordDataBytes = []byte{0x10, 0x15, 0x00, 0xD5, 0x3F, 0x01, 0x10, 0x6F, 0x1C, 0x05, 0x9E, 0x7A, 0xB5,
		0x3C, 0x35, 0x01, 0xD0, 0x87, 0x2C, 0x01, 0x00, 0x00, 0x00, 0x00}
)

func TestRecordDataSet_Encode(t *testing.T) {
	testRecordDataSet := RecordDataSet{
		RecordData{
			SubrecordData: &SrPosData{
				NavigationTime:      time.Date(2018, time.July, 5, 20, 8, 53, 0, time.UTC),
				Latitude:            55.55389399769574,
				Longitude:           37.43236696287812,
				ALTE:                "0",
				LOHS:                "0",
				LAHS:                "0",
				MV:                  "0",
				BB:                  "0",
				CS:                  "0",
				FIX:                 "0",
				VLD:                 "1",
				DirectionHighestBit: 1,
				AltitudeSign:        0,
				Speed:               200,
				Direction:           172,
				Odometer:            1,
				DigitalInputs:       0,
				Source:              0,
			},
		},
	}

	rdBytes, err := testRecordDataSet.Encode()
	if assert.NoError(t, err) {
		assert.Equal(t, rdBytes, testRecordDataBytes)
	}
}

func TestRecordDataSet_Decode(t *testing.T) {
	rds := RecordDataSet{}
	testRecordDataSet := RecordDataSet{
		RecordData{
			SubrecordType:   SrPosDataType,
			SubrecordLength: 21,
			SubrecordData: &SrPosData{
				NavigationTime:      time.Date(2018, time.July, 5, 20, 8, 53, 0, time.UTC),
				Latitude:            55.55389399769574,
				Longitude:           37.43236696287812,
				ALTE:                "0",
				LOHS:                "0",
				LAHS:                "0",
				MV:                  "0",
				BB:                  "0",
				CS:                  "0",
				FIX:                 "0",
				VLD:                 "1",
				DirectionHighestBit: 1,
				AltitudeSign:        0,
				Speed:               200,
				Direction:           172,
				Odometer:            1,
				DigitalInputs:       0,
				Source:              0,
			},
		},
	}
	if assert.NoError(t, rds.Decode(testRecordDataBytes)) {
		assert.Equal(t, rds, testRecordDataSet)
	}
}
