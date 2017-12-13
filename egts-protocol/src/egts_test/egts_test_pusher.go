package main

import (
	"bytes"
	"egts_package"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

const (
	message       = "0100c800c90000ca30000000810746a20002021015005a9cdf0eb662b3b73834b838019600f600000000001015005a9cdf0eb662b3b73834b8380100000000000000002a28"
	StopCharacter = "\r\n\r\n"
)

func SocketClient(ip string, port int) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)

	defer conn.Close()

	if err != nil {
		log.Fatalln(err)
	}

	conn.Write([]byte(message))
	conn.Write([]byte(StopCharacter))
	log.Printf("Send: %s", message)

	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	log.Printf("Receive: %s", buff[:n])

}

func main() {

	var (
		ip   = "127.0.0.1"
		port = 6000
	)
	pkgg := egts_package.EgtsPkg{
		PRV:   1,
		SKID:  0,
		PRF:   0,
		RTE:   0,
		ENA:   0,
		CMP:   0,
		PR:    0,
		HL:    0,
		HE:    0,
		FDL:   0,
		PI:    0,
		PT:    0,
		PRA:   0,
		RCA:   0,
		TTL:   0,
		HCS:   0,
		SFRD:  []byte("1"),
		SFRCS: 0,
	}

	// fmt.Println(reflect.TypeOf(pkgg))
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.LittleEndian, pkgg)
	fmt.Printf("%v %x", err, buf.Bytes())

	SocketClient(ip, port)

}
