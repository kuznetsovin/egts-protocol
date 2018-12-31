package main

import (
	"encoding/binary"
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

			resp, err := createPtResponse(uint8(resultCode), pkg.PacketIdentifier)
			if err != nil {
				log.Printf("Ошибка сборки ответа: %v\n", err)
				continue
			}
			// посылаем ответ
			conn.Write(resp)

			outData := EgtsExportPacket{
				PacketID: uint32(pkg.PacketIdentifier),
			}

			if pkg.ServicesFrameData != nil {
				switch dataFrame := pkg.ServicesFrameData.(type) {
				case *ServiceDataSet:
					savePacket := false
					packetIdBytes := make([]byte, 4)

					log.Println("Тип пакета EGTS_PT_APPDATA")
					for _, rec := range *dataFrame {
						for _, subRec := range rec.RecordDataSet {
							switch subRecData := subRec.SubrecordData.(type) {
							case *EgtsSrTermIdentity:
								log.Printf("Тип подзаписи EGTS_SR_TERM_IDENTITY (EGTS_SR_AUTH_INFO)")

								srCodePkg, err := createSrResultCode(pkg.PacketIdentifier)
								if err != nil {
									log.Printf("Ошибка сборки пакета EGTS_SR_RECORD_RESPONSE при авторизации: %v\n", err)
								}
								conn.Write(srCodePkg)
								log.Printf("Отправлен пакет подтверждения EGTS_SR_RESULT_CODE")
							case *EgtsSrResponse:
								log.Printf("Тип подзаписи EGTS_SR_RECORD_RESPONSE")
							case *EgtsSrPosData:
								log.Printf("Тип подзаписи EGTS_SR_POS_DATA")
								savePacket = true

								outData.NavigationTime = subRecData.NavigationTime
								outData.Latitude = subRecData.Latitude
								outData.Longitude = subRecData.Longitude
							case *EgtsSrAbsCntrData:
								log.Printf("Тип подзаписи EGTS_SR_ABS_CNTR_DATA")

								switch subRecData.CounterNumber {
								case 100:
									// Количество всех подтвержденных записей, переданных на сервер на момент формирования пакета.
									continue
								case 101:
									// Количество записей, которые так и не получилось отправить на момент формирования пакета.
									// Это количество всех потерянных (непереданных) записей, которые были затерты.
									continue
								case 102:
									// Количество соединений с сервером на момент формирования пакета.
									continue
								case 103:
									// Номер самой новой записи в истории на момент формирования пакета.
									continue
								case 104:
									// Номер самой старой записи в истории на момент формирования пакета.
									continue
								case 105:
									//Три малдших байта дата/время (POSIX) создания самой старой записи в истории на момент формирования пакета.
									continue
								case 106:
									//Три старших байта дата/время (POSIX) создания самой старой записи в истории на момент формирования пакета.
									continue
								case 110:
									// Три младших байта номера передаваемой записи (идет вместе с каждой POS_DATA).
									binary.BigEndian.PutUint32(packetIdBytes, subRecData.CounterValue)
									outData.PacketID = subRecData.CounterValue
								case 111:
									// один старший байт номера передаваемой записи (идет вместе с каждой POS_DATA).
									tmpBuf := make([]byte, 4)
									binary.BigEndian.PutUint32(tmpBuf, subRecData.CounterValue)

									if len(packetIdBytes) == 4 {
										packetIdBytes[3] = tmpBuf[3]
									} else {
										packetIdBytes = tmpBuf
									}

									outData.PacketID = binary.LittleEndian.Uint32(packetIdBytes)

									continue
								}
							case *EgtsSrLiquidLevelSensor:
								log.Printf("Тип подзаписи EGTS_SR_LIQUID_LEVEL_SENSOR")
								sensorData := Sensor{
									SensorNumber: subRecData.LiquidLevelSensorNumber,
									ErrorFlag:    subRecData.LiquidLevelSensorErrorFlag,
								}

								switch subRecData.LiquidLevelSensorValueUnit {
								case "00", "01":
									sensorData.ValueMm = subRecData.LiquidLevelSensorData
								case "10":
									sensorData.ValueL = subRecData.LiquidLevelSensorData * 10
								}

								outData.LiquidSensor = append(outData.LiquidSensor, sensorData)
							}
						}

						if savePacket {
							if err = outData.Save(); err != nil {
								log.Printf("Ошибка сохраения экспортного пакета: %v\n", err)
							}
						}

						srResp, err := createSrRecordResponse(uint16(outData.PacketID), rec.RecordNumber)
						if err != nil {
							log.Printf("Ошибка сборки EGTS_SR_RECORD_RESPONSE")
							continue
						}
						conn.Write(srResp)
					}
				case *EgtsPtResponse:
					log.Printf("Тип пакета EGTS_PT_RESPONSE")
				}
			}
		}
	}
}
