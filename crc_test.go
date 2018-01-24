package main

import (
	"fmt"
	"testing"
)

func Test_crc8(t *testing.T) {
	crc := crc8([]byte("123456789"))
	checkVal := byte(0x7F)

	if crc != checkVal {
		fmt.Errorf("Incorrect value: %s != %s\n", crc, checkVal)
	}
}

func Test_crc16(t *testing.T) {
	crc := crc16([]byte("123456789"))
	checkVal := uint16(0x29b1)

	if crc != checkVal {
		fmt.Errorf("Incorrect value: %s != %s\n", crc, checkVal)
	}
}
