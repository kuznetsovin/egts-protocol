package main

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_bitsToByte(t *testing.T) {
	val := "00100000"

	result, err := bitsToByte(val)
	if err != nil {
		t.Error("Error convert: ", err)
	}
	correctVal := byte(32)

	if result != correctVal {
		fmt.Errorf("Incorrect value: %s != %s\n", result, correctVal)
	}
}

func Test_bitsToBytes(t *testing.T) {
	val := "0000000010010110"

	resultBytes, err := bitsToBytes(val, 2)
	if err != nil {
		t.Error("Error convert: ", err)
	}
	correctVal := []byte{0x96, 0x00}

	if !bytes.Equal(resultBytes, correctVal) {
		fmt.Errorf("Incorrect value: %s != %s\n", resultBytes, correctVal)
	}

	val = "1100000010010110"

	resultBytes, err = bitsToBytes(val, 2)
	if err != nil {
		t.Error("Error convert: ", err)
	}
	correctVal = []byte{0x96, 0xc0}

	if !bytes.Equal(resultBytes, correctVal) {
		fmt.Errorf("Incorrect value: %s != %s\n", resultBytes, correctVal)
	}
}

func Test_byteToBits(t *testing.T) {
	b := byte(0x20)

	result := byteToBits(b)
	correctVal := "00100000"

	if result != correctVal {
		fmt.Errorf("Incorrect value: %s != %s\n", result, correctVal)
	}
}
