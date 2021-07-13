package egts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_crc8(t *testing.T) {
	crc := crc8([]byte("123456789"))
	checkVal := byte(0xF7)

	assert.Equal(t, crc, checkVal)
}

func Test_crc16(t *testing.T) {
	crc := crc16([]byte("123456789"))
	checkVal := uint16(0x29b1)

	assert.Equal(t, crc, checkVal)
}
