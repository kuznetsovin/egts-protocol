package main

import (
	"encoding/binary"
	"github.com/streadway/amqp"
	"net"
	"time"
)

func handleRecvPkg(conn net.Conn, rmqChanel *amqp.Channel) {
	var (
		isPkgSave         bool
		srResultCodePkg   []byte
		serviceType       uint8
		srResponsesRecord RecordDataSet
		err               error
	)
	buf := make([]byte, 1024)

	for {
	Received:
		serviceType = 0
		isPkgSave = false
		srResponsesRecord = nil
		srResultCodePkg = nil

		if _, err = conn.Read(buf); err != nil {
			logger.Errorf("Ошибка чтения из сетевого буфера: %v", err)
			time.Sleep(1 * time.Second)
			goto Received
		}
		logger.Debugf("Принят пакет: %X\v", buf)
		//printDecodePackage(buf)

		pkg := EgtsPackage{}
		resultCode, err := pkg.Decode(buf)
		if err != nil {
			logger.Errorf("Не удалось расшифровать пакет: %v", err)

			resp, err := pkg.CreatePtResponse(resultCode, serviceType, nil)
			if err != nil {
				logger.Errorf("Ошибка сборки ответа EGTS_PT_RESPONSE с ошибкой: %v", err)
				goto Received
			}
			conn.Write(resp)

			//printDecodePackage("Отправлен пакет EGTS_PT_RESPONSE", resp)
			goto Received
		}

		exportPacket := EgtsExportPacket{
			PacketID: uint32(pkg.PacketIdentifier),
		}
		switch pkg.PacketType {
		case egtsPtAppdata:
			logger.Info("Тип пакета EGTS_PT_APPDATA")

			for _, rec := range *pkg.ServicesFrameData.(*ServiceDataSet) {
				packetIdBytes := make([]byte, 4)

				srResponsesRecord = append(srResponsesRecord, RecordData{
					SubrecordType:   egtsSrRecordResponse,
					SubrecordLength: 3,
					SubrecordData: &EgtsSrResponse{
						ConfirmedRecordNumber: rec.RecordNumber,
						RecordStatus:          egtsPcOk,
					},
				})
				serviceType = rec.SourceServiceType
				logger.Info("Тип сервиса ", serviceType)

				exportPacket.Client = rec.ObjectIdentifier

				// проверяем не пришел ли пакет авторизации
				for _, subRec := range rec.RecordDataSet {
					switch subRecData := subRec.SubrecordData.(type) {
					case *EgtsSrTermIdentity:
						logger.Info("Разбор подзаписи EGTS_SR_TERM_IDENTITY")
						if srResultCodePkg, err = pkg.CreateSrResultCode(egtsPcOk); err != nil {
							logger.Printf("Ошибка сборки EGTS_SR_RESULT_CODE: %v", err)
						}
					case *EgtsSrAuthInfo:
						logger.Info("Разбор подзаписи EGTS_SR_AUTH_INFO")
						if srResultCodePkg, err = pkg.CreateSrResultCode(egtsPcOk); err != nil {
							logger.Printf("Ошибка сборки EGTS_SR_RESULT_CODE: %v", err)
						}
					case *EgtsSrResponse:
						logger.Info("Разбор подзаписи EGTS_SR_RESPONSE")
						goto Received
					case *EgtsSrPosData:
						logger.Info("Разбор подзаписи EGTS_SR_POS_DATA")
						isPkgSave = true

						exportPacket.NavigationTime = subRecData.NavigationTime
						exportPacket.Latitude = subRecData.Latitude
						exportPacket.Longitude = subRecData.Longitude
					case *EgtsSrAbsCntrData:
						logger.Info("Разбор подзаписи EGTS_SR_ABS_CNTR_DATA")

						switch subRecData.CounterNumber {
						case 110:
							// Три младших байта номера передаваемой записи (идет вместе с каждой POS_DATA).
							binary.BigEndian.PutUint32(packetIdBytes, subRecData.CounterValue)
							exportPacket.PacketID = subRecData.CounterValue
						case 111:
							// один старший байт номера передаваемой записи (идет вместе с каждой POS_DATA).
							tmpBuf := make([]byte, 4)
							binary.BigEndian.PutUint32(tmpBuf, subRecData.CounterValue)

							if len(packetIdBytes) == 4 {
								packetIdBytes[3] = tmpBuf[3]
							} else {
								packetIdBytes = tmpBuf
							}

							exportPacket.PacketID = binary.LittleEndian.Uint32(packetIdBytes)
						}
					case *EgtsSrLiquidLevelSensor:
						logger.Info("Разбор подзаписи EGTS_SR_LIQUID_LEVEL_SENSOR")
						sensorData := LiquidSensor{
							SensorNumber: subRecData.LiquidLevelSensorNumber,
							ErrorFlag:    subRecData.LiquidLevelSensorErrorFlag,
						}

						switch subRecData.LiquidLevelSensorValueUnit {
						case "00", "01":
							sensorData.ValueMm = subRecData.LiquidLevelSensorData
						case "10":
							sensorData.ValueL = subRecData.LiquidLevelSensorData * 10
						}

						exportPacket.LiquidSensors = append(exportPacket.LiquidSensors, sensorData)
					}
				}
			}

			if isPkgSave {
				if err := pkg.Save(rmqChanel); err != nil {
					logger.Error(err)
				}
				if err := exportPacket.Save(rmqChanel); err != nil {
					logger.Error(err)
				}
			}

			resp, err := pkg.CreatePtResponse(resultCode, serviceType, srResponsesRecord)
			if err != nil {
				logger.Errorf("Ошибка сборки ответа: %v", err)
				goto Received
			}
			conn.Write(resp)

			logger.Debugf("Отправлен пакет EGTS_PT_RESPONSE: %X", resp)
			//logger.Debug(printDecodePackage(resp))

			if len(srResultCodePkg) > 0 {
				conn.Write(srResultCodePkg)
				logger.Debugf("Отправлен пакет EGTS_SR_RESULT_CODE: %X", resp)
				//logger.Debug(printDecodePackage(srResultCodePkg))
			}
		case egtsPtResponse:
			logger.Printf("Тип пакета EGTS_PT_RESPONSE")
		}

	}
}
