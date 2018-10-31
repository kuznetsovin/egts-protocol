package main

import (
	"bytes"
	"reflect"
	"testing"
)

var (
	egtsPkgResp = EgtsPackage{
		ProtocolVersion:  1,
		SecurityKeyID:    0,
		Prefix:           "00",
		Route:            "0",
		EncryptionAlg:    "00",
		Compression:      "0",
		Priority:         "11",
		HeaderLength:     11,
		HeaderEncoding:   0,
		FrameDataLength:  3,
		PacketIdentifier: 137,
		PacketType:       0,
		HeaderCheckSum:   74,
		ServicesFrameData: &EgtsPtResponse{
			ResponsePacketID: 14357,
			ProcessingResult: egtsPcOk,
		},
		ServicesFrameDataCheckSum: 59443,
	}
)


func TestEgtsPkgResp_Encode(t *testing.T) {
	testEgtsPkgBytes := []byte{0x01, 0x00, 0x03, 0x0B, 0x00, 0x03, 0x00, 0x89, 0x00, 0x00, 0x4A, 0x15, 0x38, 0x00, 0x33, 0xE8}

	posDataBytes, err := egtsPkgResp.Encode()
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(posDataBytes, testEgtsPkgBytes) {
		t.Errorf("Байтовые строки не совпадают: %v != %v ", posDataBytes, testEgtsPkgBytes)
	}
}

func TestEgtsPkgResp_Decode(t *testing.T) {
	egtsPkgBytes := []byte{0x01, 0x00, 0x03, 0x0B, 0x00, 0x03, 0x00, 0x89, 0x00, 0x00, 0x4A, 0x15, 0x38, 0x00, 0x33, 0xE8}

	egtsPkg := EgtsPackage{}

	if _, err := egtsPkg.Decode(egtsPkgBytes); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if !reflect.DeepEqual(egtsPkg, egtsPkgResp) {
		t.Errorf("Пакеты не совпадают")
	}
}

