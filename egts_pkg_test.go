package main

import (
	"bytes"
	"testing"
)

func TestEgtsPkgHeader_ToBytes(t *testing.T) {
	etgsHeader := EgtsPkgHeader{
		PRV:  1,
		SKID: 0,
		PRF:  0,
		RTE:  1,
		ENA:  0,
		CMP:  0,
		PR:   0,
		HL:   16,
		HE:   0,
		FDL:  59,
		PID:  3612,
		PT:   1,
		PRA:  200,
		RCA:  201,
		TTL:  0,
	}

	correctHeader := []byte{0x01, 0x00, 0x20, 0x10, 0x00, 0x3b, 0x00, 0x1c, 0x0e, 0x01,
		0xc8, 0x00, 0xc9, 0x00, 0x00, 0xca}
	resultBytes, err := etgsHeader.ToBytes()

	if err != nil {
		t.Error("Error etgs header decode ", err)
	}

	if !bytes.Equal(resultBytes, correctHeader) {
		t.Errorf("Incorrect etgs header decode: %v != %v ", resultBytes, correctHeader)
	}
}
