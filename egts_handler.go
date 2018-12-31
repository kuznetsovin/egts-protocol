package main

import (
	"log"
	"net"
	"time"
)

func handleRecvPkg(conn net.Conn) {
	buf := make([]byte, 1024)

	for {
		if _, err := conn.Read(buf); err != nil {
			log.Printf("Ошибка чтения из сетевого буфера: %v\n", err)
			time.Sleep(1 * time.Second)
		} else {
			pkg := EgtsPackage{}
			resultCode, err := pkg.Decode(buf)
			if err != nil {
				log.Printf("Не удалось расшифровать пакет: %v\n", err)
			}

			if pkg.ServicesFrameData != nil {
				switch dataFrame := pkg.ServicesFrameData.(type) {
				case *ServiceDataSet:
					for _, rec := range *dataFrame {
						for _, subRec := range rec.RecordDataSet {
							switch subRec.SubrecordType {
							case egtsSrTermIdentity:

								// TODO: доделать полноценную авторизацию (нужен эмулятор)

								log.Printf("Начало авторизации принят EGTS_SR_TERM_IDENTITY")

								authResp, err := createSrRecordResponse(pkg.PacketIdentifier, rec.RecordNumber)
								if err != nil {
									log.Printf("Ошибка сборки пакета EGTS_SR_RECORD_RESPONSE при авторизации %v\n", err)
								}

								conn.Write(authResp)
								log.Printf("Отправлен пакет подтверждения EGTS_SR_RECORD_RESPONSE")

								srCodePkg, err := createSrResultCode(pkg.PacketIdentifier)
								if err != nil {
									log.Printf("Ошибка сборки пакета EGTS_SR_RECORD_RESPONSE при авторизации: %v\n", err)
								}
								conn.Write(srCodePkg)
								log.Printf("Отправлен пакет подтверждения EGTS_SR_RESULT_CODE")

							case egtsSrRecordResponse:
								log.Printf("Принят пакет подтверждения EGTS_SR_RECORD_RESPONSE")
							}
						}
					}
				case *EgtsPtResponse:
					log.Printf("Подтверждение получения EGTS_PT_RESPONSE")
				}

			}
			resp, err := createPtResponse(uint8(resultCode), pkg.PacketIdentifier)
			if err != nil {
				log.Printf("Ошибка сборки ответа: %v\n", err)
				continue
			}
			// посылаем ответ в случае удачи
			conn.Write(resp)

		}

	}
}
