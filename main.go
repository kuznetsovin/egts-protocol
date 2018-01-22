package main

import (
	"log"
)

func main() {
	msg := []byte{}
	log.Printf("Send: %s", msg)

	//addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	//conn, err := net.Dial("tcp", addr)
	//
	//defer conn.Close()
	//
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//// conn.Write([]byte(message))
	//conn.Write([]byte(msg))
	//conn.Write([]byte(StopCharacter))
	//log.Printf("Send: %s", msg)
	//
	//buff := make([]byte, 1024)
	//n, _ := conn.Read(buff)
	//log.Printf("Receive: %s", buff[:n])

}
