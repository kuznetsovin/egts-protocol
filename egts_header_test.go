package main

import (
	"testing"
	"reflect"
	"bytes"
)

var (
	testHeaderBytes = []byte{0x01, 0x00, 0x00, 0x0B, 0x00, 0x10, 0x00, 0x86, 0x00, 0x00, 0x18}
	testHeader      = EgtsHeader{
		ProtocolVersion:  1,
		SecurityKeyID:    0,
		Prefix:           "00",
		Route:            "0",
		EncryptionAlg:    "00",
		Compression:      "0",
		Priority:         "00",
		HeaderLength:     11,
		HeaderEncoding:   0,
		FrameDataLength:  16,
		PacketIdentifier: 134,
		HeaderCheckSum:   0x18,
	}
)


func TestEgtsHeader_Encode(t *testing.T) {
	hb, err := testHeader.Encode()
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(hb, testHeaderBytes) {
		t.Errorf("Байтовые строки не совпадают: %v != %v ", hb, testHeaderBytes)
	}
}

func TestEgtsHeader_Decode(t *testing.T) {
	h := EgtsHeader{}

	if err := h.Decode(testHeaderBytes); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if !reflect.DeepEqual(h, testHeader) {
		t.Errorf("Заголовки не совпадают: %v != %v ", h, testHeader)
	}


}




