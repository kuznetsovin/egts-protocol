package main

import (
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

func Test_byteToBits(t *testing.T) {
	b := byte(0x20)

	result := byteToBits(b)
	correctVal := "00100000"

	if result != correctVal {
		fmt.Errorf("Incorrect value: %s != %s\n", result, correctVal)
	}
}
