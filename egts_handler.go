package main

import (
	"log"
	"net"
	"time"
)

func handleRecvPkg(conn net.Conn) {
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	for {
		if _, err := conn.Read(buf); err != nil {
			log.Printf("Ошибка чтения из сетевого буфера: %v\n", err)
			time.Sleep(1*time.Second)
		} else {
			// TODO: добавить авторизацию
			pkg := EgtsPackage{}
			resultCode, err := pkg.Decode(buf)
			if err != nil {
				log.Printf("Не удалось расшифровать пакет: %v\n", err)
			}

			resp, err := createPtResponse(uint8(resultCode), pkg.PacketIdentifier)
			if err != nil {
				log.Printf("Не создать ответ: %v\n", err)
			} else {
				// посылаем ответ в случае удачи
				conn.Write(resp)
			}
		}

	}
}

func createPtResponse(resultCode uint8, pkgNum uint16) ([]byte, error) {
	respPkg := EgtsPackage{
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
		PacketIdentifier: pkgNum+1,
		PacketType:       0,
		ServicesFrameData: &EgtsPtResponse{
			ResponsePacketID: pkgNum,
			ProcessingResult: resultCode,
		},
	}

	return respPkg.Encode()
}
