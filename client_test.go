package main

import (
	"log"
	"net"
	"testing"
)

func testServer(port string) {
	server, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("couldn't start listening: " + err.Error())
	}
	defer server.Close()

	for {
		cnt, err := server.Accept()
		if err != nil {
			log.Fatal("Connect accept error: ", err.Error())
		}

		go func(conn net.Conn) {
			defer conn.Close()

			buf := make([]byte, 1024)
			size, err := conn.Read(buf)
			if err != nil {
				log.Fatal("Read error: ", err.Error())
			}

			data := append(buf[:size], []byte("Echo")...)

			conn.Write(data)
		}(cnt)
	}
}

func TestEgtsClient(t *testing.T) {
	port := "3000"

	// запускаем сервер
	go testServer(port)

	ec := EgtsClient{}
	ec.Start("127.0.0.1", port)

	testPkg := []byte("Test")

	r, err := ec.SendPackage(testPkg)
	if err != nil {
		t.Fatal("Send error: ", err.Error())
	}

	ec.Stop()

	if string(r) != "TestEcho" {
		t.Errorf("Result not equals: %v != TestEcho ", string(r))
	}
}
