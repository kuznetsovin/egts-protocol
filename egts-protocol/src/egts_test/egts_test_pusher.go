package main

import (
	"bytes"
	"egts_package"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"reflect"
	"service_data_records"
	"strconv"
	"strings"
)

const (
	message       = "0100c800c90000ca30000000810746a20002021015005a9cdf0eb662b3b73834b838019600f600000000001015005a9cdf0eb662b3b73834b8380100000000000000002a28"
	StopCharacter = "\r\n\r\n"
)

func SocketClient(ip string, port int, pkgg egts_package.EgtsPkg) {
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
			binary.Write(bytes, binary.LittleEndian, field.Interface())
		case reflect.Uint8:
			binary.Write(bytes, binary.LittleEndian, uint8(field.Uint()))
		case reflect.Uint16:
			binary.Write(bytes, binary.LittleEndian, uint16(field.Uint()))

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

	sdr := service_data_records.ServiceDataRecord{
		RL:   uint(136),
		RN:   uint(35267),
		RFL:  byte(0),
		SSOD: 0,
		RSOD: 0,
		GRP:  0,
		RPP:  0,
		TMFE: 0,
		EVFE: 0,
		OBFE: 0,
		OID:  0,
		EVID: 0,
		TM:   0,
		SST:  1,
		RST:  1,
		RD:   []byte("2010003b001c0e01c38800c3890000c38a30000000c2810746c2a20002021015005ac29cc39f0ec2b662c2b3c2b73834c2b83801c29600c3b600000000001015005ac29cc39f"),
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

	pkgg := egts_package.EgtsPkg{
		PRV:  byte(1),
		SKID: byte(0),
		// PRF:   strconv.ParseInt("00100000", 2, 8),
		PRF:   uint8(0),
		RTE:   uint8(0),
		ENA:   uint8(0),
		CMP:   uint8(0),
		PR:    uint8(0),
		HL:    byte(16),
		HE:    byte(0),
		FDL:   uint16(59),
		PI:    uint16(3612),
		PT:    byte(1),
		PRA:   uint16(72),
		RCA:   uint16(11),
		TTL:   byte(0),
		HCS:   byte(195),
		SFRD:  StructToBytes(sdr),
		SFRCS: uint16(64431),
	}

	SocketClient(ip, port, pkgg)

}
