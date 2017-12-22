package main

import (
	"./egts-protocol"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"reflect"
	"strconv"
	"strings"
)

const (
	message       = "0100c800c90000ca30000000810746a20002021015005a9cdf0eb662b3b73834b838019600f600000000001015005a9cdf0eb662b3b73834b8380100000000000000002a28"
	StopCharacter = "\r\n\r\n"
)

func SocketClient(ip string, port int, pkgg egts_protocol.EgtsPkg) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)

	defer conn.Close()

	if err != nil {
		log.Fatalln(err)
	}

	// conn.Write([]byte(message))
	msg := StructToBytes(pkgg)
	conn.Write([]byte(msg))
	conn.Write([]byte(StopCharacter))
	log.Printf("Send: %s", msg)

	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	log.Printf("Receive: %s", buff[:n])

}

// func (s *A) BinSeralise() byte {
// 	var buffer bytes.Buffer
// 	for i := 0; i < A.NumField(); i++ {
// 		buffer.WriteString(field)
// 		return buffer.Bytes()
// 	}
// }

// func (b *B) BinSeralise() []byte {
// 	// тут будет сериализация в байты основной структуры
// 	biteFromBits := b.A.BinSeralise()
// 	return []byte{biteFromBits, 1}
// }

func StructToBytes(data interface{}) []byte {
	structType := reflect.ValueOf(data)
	structKind := structType.Kind()
	if structKind == reflect.Ptr {
		structType = reflect.ValueOf(data).Elem()
		structKind = structType.Kind()
	}

	if structKind != reflect.Struct {
		panic("data must of type struct or struct ptr, got: " + structKind.String())
	}

	bytes := new(bytes.Buffer)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)

		switch kind := field.Kind(); kind {
		case reflect.Struct:
			binary.Write(bytes, binary.LittleEndian, StructToBytes(field.Interface()))
		case reflect.Array, reflect.Slice:
			for _, item := range field.Value() {
				binary.Write(bytes, binary.LittleEndian, StructToBytes(item.Interface()))
			}
		case reflect.Uint8:
			binary.Write(bytes, binary.LittleEndian, uint8(field.Uint()))
		case reflect.Uint16:
			binary.Write(bytes, binary.LittleEndian, uint16(field.Uint()))
		case reflect.Uint32:
			binary.Write(bytes, binary.LittleEndian, uint32(field.Uint()))
		case reflect.String:
			binary.Write(bytes, binary.LittleEndian, string(field.Uint()))
		}
	}
	fmt.Printf("%+x", bytes.Bytes())
	return bytes.Bytes()
}

func main() {

	var (
		ip   = "127.0.0.1"
		port = 6000
	)
	// Record Length = 136
	// Record Number = 35267
	// RFL = 00000000
	// OBFE = 0
	// EVFE = 0
	// TMFE = 0
	// RPP = 0
	// GRP = 0
	// RSOD = 0
	// SSOD = 0

	egtsSRServiceResponse := egts_protocol.EGTS_SR_RECORD_RESPONSE{
		CRN: int(136),
		RST: byte(1),
	}

	egtsSRTermIdentity := egts_protocol.EGTS_SR_TERM_IDENTITY{
		TID:    int(1),
		HDIDE:  uint8(0),
		IMEIE:  uint8(1),
		IMSIE:  uint8(1),
		LNGCE:  uint8(0),
		SSRA:   uint8(1),
		NIDE:   uint8(0),
		BSE:    uint8(0),
		MNE:    uint8(1),
		HDID:   int(0),
		IMEI:   string("020000000000010"),
		IMSI:   string("0200000000000103"),
		LNGC:   string(""),
		NID:    uint8(0),
		BS:     int(0),
		MSISDN: string(""),
	}

	egtsSRModuleData := egts_protocol.EGTS_SR_MODULE_DATA{
		MT:   byte(1),
		VID:  uint(0),
		FWV:  uint8(0),
		SWV:  uint8(0),
		MD:   byte(0),
		ST:   byte(1),
		SRN:  string("123"),
		D:    byte(0),
		DSCR: string("description sr module data"),
	}

	egtsSRVehicleData := egts_protocol.EGTS_SR_VEHICLE_DATA{
		VIN:  string("0123"),
		VHT:  uint(9),
		VPST: uint(9),
	}

	egtsSRAuthParams := egts_protocol.EGTS_SR_AUTH_PARAMS{
		EXE:  uint8(0),
		SSE:  uint8(0),
		MSE:  uint8(0),
		ISLE: uint8(0),
		PKE:  uint8(0),
		ENA:  uint8(0),
		PKL:  int(0),
		PBK:  int8(0),
		ISL:  int(0),
		MSZ:  int(0),
		SS:   string(0),
		D:    byte(0),
		EXP:  string(0),
	}

	egtsSRAuthInfo := egts_protocol.EGTS_SR_AUTH_INFO{
		UNM:  string("fak"),
		D:    byte(0),
		UPSW: string("sssss"),
		SS:   string(""),
	}

	egtsSRServiceInfo := egts_protocol.EGTS_SR_SERVICE_INFO{
		ST:    uint8(1),
		SST:   uint8(0),
		SRVP:  uint8(0),
		SRVA:  uint8(0),
		SRVRP: uint8(0),
	}

	egtsSRResultCode := egts_protocol.EGTS_SR_RESULT_CODE{
		RCD: byte(0),
	}

	// 	Transport level --------------------
	// Protocol Version = 1
	// Security Key ID = 0
	// PRF_full = 00100000
	// RTE = 0
	// Header Length = 16
	// Header Encoding = 0
	// Frame Data Length = 59
	// Packet Identifier = 3612
	// Packet Type = 1
	// Header Check Sum = 195
	// Calculated Header Check Sum = 155
	// SFRCS (crc) = 40899
	// Calculated SFRCS(crc) = 64431
	// b'01002010003b001c0e01c38800c3890000c38a30000000c2810746c2a20002021015005ac29cc39f0ec2b662c2b3c2b73834c2b83801c29600c3b600000000001015005ac29cc39f'

	egtsAuthService := egts_protocol.EGTS_AUTH_SERVICE{
		EGTS_SR_RECORD_RESPONSE: []egts_protocol.EGTS_SR_RECORD_RESPONSE{egtsSRServiceResponse},
		EGTS_SR_TERM_IDENTITY:   []egts_protocol.EGTS_SR_TERM_IDENTITY{egtsSRTermIdentity},
		EGTS_SR_MODULE_DATA:     []egts_protocol.EGTS_SR_MODULE_DATA{egtsSRModuleData},
		EGTS_SR_VEHICLE_DATA:    []egts_protocol.EGTS_SR_VEHICLE_DATA{egtsSRVehicleData},
		EGTS_SR_AUTH_PARAMS:     []egts_protocol.EGTS_SR_AUTH_PARAMS{egtsSRAuthParams},
		EGTS_SR_AUTH_INFO:       []egts_protocol.EGTS_SR_AUTH_INFO{egtsSRAuthInfo},
		EGTS_SR_SERVICE_INFO:    []egts_protocol.EGTS_SR_SERVICE_INFO{egtsSRServiceInfo},
		EGTS_SR_RESULT_CODE:     []egts_protocol.EGTS_SR_RESULT_CODE{egtsSRResultCode},
	}

	serviceDataSubrecord := egts_protocol.ServiceDataSubrecord{
		SRT: byte(1),
		SRL: int(160),
		SRD: StructToBytes(egtsAuthService),
	}

	egtsServiceDataRecord := egts_protocol.ServiceDataRecord{
		RL:   uint(136),
		RN:   uint(35267),
		RFL:  byte(0),
		SSOD: 1,
		RSOD: 1,
		GRP:  1,
		RPP:  1,
		TMFE: 0,
		EVFE: 0,
		OBFE: 0,
		OID:  1,
		EVID: 1,
		TM:   0,
		SST:  1,
		RST:  1,
		RD:   StructToBytes(serviceDataSubrecord),
	}
	pkgg := egts_protocol.EgtsPkg{
		PRV:   byte(1),
		SKID:  byte(0),
		PRF:   uint8(7),
		HL:    byte(11),
		HE:    byte(0),
		FDL:   uint16(59),
		PI:    uint16(3612),
		PT:    byte(0),
		PRA:   uint16(1),
		RCA:   uint16(1),
		TTL:   byte(1),
		HCS:   byte(0),
		SFRD:  StructToBytes(egtsServiceDataRecord),
		SFRCS: uint16(0),
	}

	SocketClient(ip, port, pkgg)

}
