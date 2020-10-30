package egts

import (
	"testing"
)

func Test_crc8(t *testing.T) {
	crc := crc8([]byte("123456789"))
	checkVal := byte(0xF7)

	if crc != checkVal {
		t.Errorf("Incorrect value: %x != %x\n", crc, checkVal)
	}
}

func Test_crc16(t *testing.T) {
	crc := crc16([]byte("123456789"))
	checkVal := uint16(0x29b1)

	if crc != checkVal {
		t.Errorf("Incorrect value: %x != %x\n", crc, checkVal)
	}
}
