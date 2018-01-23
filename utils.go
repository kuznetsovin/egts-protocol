package main

import (
	"encoding/binary"
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

// Функция преобразования битов в число
func bitsToBytes(Bits string, countBytes int) ([]byte, error) {
	result := make([]byte, countBytes)

	lenBits := countBytes * 8
	val, err := strconv.ParseUint(Bits, 2, lenBits)
	if err != nil {
		return result, err
	}

	switch lenBits {
	case 16:
		binary.LittleEndian.PutUint16(result, uint16(val))
	case 32:
		binary.LittleEndian.PutUint32(result, uint32(val))
	default:
		binary.LittleEndian.PutUint32(result, uint32(val))
	}

	return result, err
}

// Функция преобразования байтов и биты
func byteToBits(b byte) string {
	return fmt.Sprintf("%08b", b)
}
