package main

import (
	"fmt"
	"strconv"
)

// Функция преобразования битов в число
func bitsToByte(Bits string) (byte, error) {
	s, err := strconv.ParseUint(Bits, 2, 8)
	if err != nil {
		return uint8(s), err
	}

	return uint8(s), err

}

// Функция преобразования байтов и биты
func byteToBits(b byte) string {
	return fmt.Sprintf("%08b", b)
}

