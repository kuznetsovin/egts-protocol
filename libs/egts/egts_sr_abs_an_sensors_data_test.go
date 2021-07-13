package egts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEgtsSrAbsAnSensData_Encode(t *testing.T) {
	a := SrAbsAnSensData{
		SensorNumber: 0x98,
		Value:        0x123456,
	}
	data, err := a.Encode()
	if assert.NoError(t, err) {
		assert.False(t, data[0] != 0x98 || data[1] != 0x56 || data[2] != 0x34 || data[3] != 0x12)

		b := SrAbsAnSensData{}

		if err := b.Decode(data); assert.NoError(t, err) {
			assert.False(t, a.Value != b.Value || a.SensorNumber != b.SensorNumber)
		}
	}
}
func TestEgtsSrAbsAnSensData_Decode(t *testing.T) {
	data := []byte{0x98, 0x56, 0x34, 0x12}
	a := SrAbsAnSensData{}
	err := a.Decode(data)
	if assert.NoError(t, err) {
		assert.False(t, a.SensorNumber != 0x98 || a.Value != 0x123456)
		data2, err := a.Encode()
		if assert.NoError(t, err) {
			assert.Equal(t, data, data2)
		}
	}
}
