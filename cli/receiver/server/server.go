package server

import (
	"encoding/binary"
	"github.com/kuznetsovin/egts-protocol/cli/receiver/storage"
	"github.com/kuznetsovin/egts-protocol/libs/egts"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"time"
)

const (
	egtsPcOk  = 0
	headerLen = 10
)

type Server struct {
	addr  string
	ttl   time.Duration
	store storage.Connector
	l     net.Listener
}

func (s *Server) Run() {
	var err error

	s.l, err = net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("Не удалось открыть соединение: %v", err)
	}
	defer s.l.Close()

	log.Infof("Запущен сервер %s...", s.addr)
	for {
		conn, err := s.l.Accept()
		if err != nil {
			log.WithField("err", err).Errorf("Ошибка соединения")
		} else {
			go s.handleConn(conn)
		}
	}
}

func (s *Server) Stop() error {
	if s.l != nil {
		return s.l.Close()
	}

	return nil
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()

	var (
		isPkgSave         bool
		srResultCodePkg   []byte
		serviceType       uint8
		srResponsesRecord egts.RecordDataSet
		recvPacket        []byte
	)

	if s.store == nil {
		log.Error("Не корректная ссылка на объект хранилища")
		return
	}
	log.WithField("ip", conn.RemoteAddr()).Info("Установлено соединение")

	for {
	Received:
		serviceType = 0
		srResponsesRecord = nil
		srResultCodePkg = nil
		recvPacket = nil

		connTimer := time.NewTimer(s.ttl)

		// считываем заголовок пакета
		headerBuf := make([]byte, headerLen)
		_, err := io.ReadFull(conn, headerBuf)

		switch err {
		case nil:
			connTimer.Reset(s.ttl)

			// если пакет не егтс формата закрываем соединение
			if headerBuf[0] != 0x01 {
				log.WithField("ip", conn.RemoteAddr()).Warn("Пакет не соответствует формату ЕГТС. Закрыто соединение")
				return
			}

			// вычисляем длину пакета, равную длине заголовка (HL) + длина тела (FDL) + CRC пакета 2 байта если есть FDL из приказа минтранса №285
			bodyLen := binary.LittleEndian.Uint16(headerBuf[5:7])
			pkgLen := uint16(headerBuf[3])
			if bodyLen > 0 {
				pkgLen += bodyLen + 2
			}
			// получаем концовку ЕГТС пакета
			buf := make([]byte, pkgLen-headerLen)
			if _, err := io.ReadFull(conn, buf); err != nil {
				log.WithField("err", err).Error("Ошибка при получении тела пакета")
				return
			}

			// формируем полный пакет
			recvPacket = append(headerBuf, buf...)
		case io.EOF:
			<-connTimer.C
			log.WithField("ip", conn.RemoteAddr()).Warnf("Соединение закрыто по таймауту")
			return
		default:
			log.WithField("err", err).Error("Ошибка при получении")
			return
		}

		log.WithField("packet", recvPacket).Debug("Принят пакет")
		pkg := egts.Package{}
		receivedTimestamp := time.Now().UTC().Unix()
		resultCode, err := pkg.Decode(recvPacket)
		if err != nil {
			log.WithField("err", err).Error("Ошибка расшифровки пакета")

			resp, err := createPtResponse(pkg.PacketIdentifier, resultCode, serviceType, nil)
			if err != nil {
				log.WithField("err", err).Error("Ошибка сборки ответа EGTS_PT_RESPONSE с ошибкой")
				goto Received
			}
			_, _ = conn.Write(resp)

			goto Received
		}

		switch pkg.PacketType {
		case egts.PtAppdataPacket:
			log.Debug("Тип пакета EGTS_PT_APPDATA")

			for _, rec := range *pkg.ServicesFrameData.(*egts.ServiceDataSet) {
				exportPacket := storage.NavRecord{
					PacketID: uint32(pkg.PacketIdentifier),
				}

				isPkgSave = false
				packetIDBytes := make([]byte, 4)

				srResponsesRecord = append(srResponsesRecord, egts.RecordData{
					SubrecordType:   egts.SrRecordResponseType,
					SubrecordLength: 3,
					SubrecordData: &egts.SrResponse{
						ConfirmedRecordNumber: rec.RecordNumber,
						RecordStatus:          egtsPcOk,
					},
				})
				serviceType = rec.SourceServiceType
				log.Info("Тип сервиса ", serviceType)

				exportPacket.Client = rec.ObjectIdentifier

				for _, subRec := range rec.RecordDataSet {
					switch subRecData := subRec.SubrecordData.(type) {
					case *egts.SrTermIdentity:
						log.Debug("Разбор подзаписи EGTS_SR_TERM_IDENTITY")
						if srResultCodePkg, err = createSrResultCode(pkg.PacketIdentifier, egtsPcOk); err != nil {
							log.Errorf("Ошибка сборки EGTS_SR_RESULT_CODE: %v", err)
						}
					case *egts.SrAuthInfo:
						log.Debug("Разбор подзаписи EGTS_SR_AUTH_INFO")
						if srResultCodePkg, err = createSrResultCode(pkg.PacketIdentifier, egtsPcOk); err != nil {
							log.Errorf("Ошибка сборки EGTS_SR_RESULT_CODE: %v", err)
						}
					case *egts.SrResponse:
						log.Debugf("Разбор подзаписи EGTS_SR_RESPONSE")
						goto Received
					case *egts.SrPosData:
						log.Debugf("Разбор подзаписи EGTS_SR_POS_DATA")
						isPkgSave = true

						exportPacket.NavigationTimestamp = subRecData.NavigationTime.Unix()
						exportPacket.ReceivedTimestamp = receivedTimestamp
						exportPacket.Latitude = subRecData.Latitude
						exportPacket.Longitude = subRecData.Longitude
						exportPacket.Speed = subRecData.Speed
						exportPacket.Course = subRecData.Direction
					case *egts.SrExtPosData:
						log.Debug("Разбор подзаписи EGTS_SR_EXT_POS_DATA")
						exportPacket.Nsat = subRecData.Satellites
						exportPacket.Pdop = subRecData.PositionDilutionOfPrecision
						exportPacket.Hdop = subRecData.HorizontalDilutionOfPrecision
						exportPacket.Vdop = subRecData.VerticalDilutionOfPrecision
						exportPacket.Ns = subRecData.NavigationSystem

					case *egts.SrAdSensorsData:
						log.Debug("Разбор подзаписи EGTS_SR_AD_SENSORS_DATA")
						if subRecData.AnalogSensorFieldExists1 == "1" {
							exportPacket.AnSensors = append(exportPacket.AnSensors, storage.AnSensor{SensorNumber: 1, Value: subRecData.AnalogSensor1})
						}

						if subRecData.AnalogSensorFieldExists2 == "1" {
							exportPacket.AnSensors = append(exportPacket.AnSensors, storage.AnSensor{SensorNumber: 2, Value: subRecData.AnalogSensor2})
						}

						if subRecData.AnalogSensorFieldExists3 == "1" {
							exportPacket.AnSensors = append(exportPacket.AnSensors, storage.AnSensor{SensorNumber: 3, Value: subRecData.AnalogSensor3})
						}
						if subRecData.AnalogSensorFieldExists4 == "1" {
							exportPacket.AnSensors = append(exportPacket.AnSensors, storage.AnSensor{SensorNumber: 4, Value: subRecData.AnalogSensor4})
						}
						if subRecData.AnalogSensorFieldExists5 == "1" {
							exportPacket.AnSensors = append(exportPacket.AnSensors, storage.AnSensor{SensorNumber: 5, Value: subRecData.AnalogSensor5})
						}
						if subRecData.AnalogSensorFieldExists6 == "1" {
							exportPacket.AnSensors = append(exportPacket.AnSensors, storage.AnSensor{SensorNumber: 6, Value: subRecData.AnalogSensor6})
						}
						if subRecData.AnalogSensorFieldExists7 == "1" {
							exportPacket.AnSensors = append(exportPacket.AnSensors, storage.AnSensor{SensorNumber: 7, Value: subRecData.AnalogSensor7})
						}
						if subRecData.AnalogSensorFieldExists8 == "1" {
							exportPacket.AnSensors = append(exportPacket.AnSensors, storage.AnSensor{SensorNumber: 8, Value: subRecData.AnalogSensor8})
						}
					case *egts.SrAbsAnSensData:
						log.Debug("Разбор подзаписи EGTS_SR_ABS_AN_SENS_DATA")
						exportPacket.AnSensors = append(exportPacket.AnSensors, storage.AnSensor{SensorNumber: subRecData.SensorNumber, Value: subRecData.Value})

					case *egts.SrAbsCntrData:
						log.Debug("Разбор подзаписи EGTS_SR_ABS_CNTR_DATA")

						switch subRecData.CounterNumber {
						case 110:
							// Три младших байта номера передаваемой записи (идет вместе с каждой POS_DATA).
							binary.BigEndian.PutUint32(packetIDBytes, subRecData.CounterValue)
							exportPacket.PacketID = subRecData.CounterValue
						case 111:
							// один старший байт номера передаваемой записи (идет вместе с каждой POS_DATA).
							tmpBuf := make([]byte, 4)
							binary.BigEndian.PutUint32(tmpBuf, subRecData.CounterValue)

							if len(packetIDBytes) == 4 {
								packetIDBytes[3] = tmpBuf[3]
							} else {
								packetIDBytes = tmpBuf
							}

							exportPacket.PacketID = binary.LittleEndian.Uint32(packetIDBytes)
						}
					case *egts.SrLiquidLevelSensor:
						log.Debug("Разбор подзаписи EGTS_SR_LIQUID_LEVEL_SENSOR")
						sensorData := storage.LiquidSensor{
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

				if isPkgSave {
					if err := s.store.Save(&exportPacket); err != nil {
						log.WithField("err", err).Error("Ошибка сохранения телеметрии")
					}
				}
			}

			resp, err := createPtResponse(pkg.PacketIdentifier, resultCode, serviceType, srResponsesRecord)
			if err != nil {
				log.WithField("err", err).Error("Ошибка сборки ответа")
				goto Received
			}
			_, _ = conn.Write(resp)

			log.WithField("packet", resp).Debug("Отправлен пакет EGTS_PT_RESPONSE")

			if len(srResultCodePkg) > 0 {
				_, _ = conn.Write(srResultCodePkg)
				log.WithField("packet", resp).Debug("Отправлен пакет EGTS_SR_RESULT_CODE")
			}
		case egts.PtResponsePacket:
			log.Debug("Тип пакета EGTS_PT_RESPONSE")
		}

	}
}

func New(srvAddress string, ttl time.Duration, s storage.Connector) Server {
	return Server{
		addr:  srvAddress,
		ttl:   ttl,
		store: s,
	}
}

func createPtResponse(pid uint16, resultCode, serviceType uint8, srResponses egts.RecordDataSet) ([]byte, error) {
	respSection := egts.PtResponse{
		ResponsePacketID: pid,
		ProcessingResult: resultCode,
	}

	if srResponses != nil {
		respSection.SDR = &egts.ServiceDataSet{
			egts.ServiceDataRecord{
				RecordLength:             srResponses.Length(),
				RecordNumber:             1,
				SourceServiceOnDevice:    "0",
				RecipientServiceOnDevice: "0",
				Group:                    "1",
				RecordProcessingPriority: "00",
				TimeFieldExists:          "0",
				EventIDFieldExists:       "0",
				ObjectIDFieldExists:      "0",
				SourceServiceType:        serviceType,
				RecipientServiceType:     serviceType,
				RecordDataSet:            srResponses,
			},
		}
	}

	respPkg := egts.Package{
		ProtocolVersion:   1,
		SecurityKeyID:     0,
		Prefix:            "00",
		Route:             "0",
		EncryptionAlg:     "00",
		Compression:       "0",
		Priority:          "00",
		HeaderLength:      11,
		HeaderEncoding:    0,
		FrameDataLength:   respSection.Length(),
		PacketIdentifier:  pid + 1,
		PacketType:        egts.PtResponsePacket,
		ServicesFrameData: &respSection,
	}

	return respPkg.Encode()
}

func createSrResultCode(pid uint16, resultCode uint8) ([]byte, error) {
	rds := egts.RecordDataSet{
		egts.RecordData{
			SubrecordType:   egts.SrResultCodeType,
			SubrecordLength: uint16(1),
			SubrecordData: &egts.SrResultCode{
				ResultCode: resultCode,
			},
		},
	}

	sfd := egts.ServiceDataSet{
		egts.ServiceDataRecord{
			RecordLength:             rds.Length(),
			RecordNumber:             1,
			SourceServiceOnDevice:    "0",
			RecipientServiceOnDevice: "0",
			Group:                    "1",
			RecordProcessingPriority: "00",
			TimeFieldExists:          "0",
			EventIDFieldExists:       "0",
			ObjectIDFieldExists:      "0",
			SourceServiceType:        egts.AuthService,
			RecipientServiceType:     egts.AuthService,
			RecordDataSet:            rds,
		},
	}

	respPkg := egts.Package{
		ProtocolVersion:   1,
		SecurityKeyID:     0,
		Prefix:            "00",
		Route:             "0",
		EncryptionAlg:     "00",
		Compression:       "0",
		Priority:          "00",
		HeaderLength:      11,
		HeaderEncoding:    0,
		FrameDataLength:   sfd.Length(),
		PacketIdentifier:  pid + 1,
		PacketType:        egts.PtResponsePacket,
		ServicesFrameData: &sfd,
	}

	return respPkg.Encode()
}
