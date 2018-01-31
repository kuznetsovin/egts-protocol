package main

import (
	"log"
	"net"
)

type EgtsClient struct {
	conn *net.TCPConn
}

func (c *EgtsClient) Start(server, port string) {

	srvAddr := server + ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp", srvAddr)
	if err != nil {
		log.Fatal("ResolveTCPAddr failed:", err.Error())
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal("Dial failed:", err.Error())
	}

	c.conn = conn
}

func (c *EgtsClient) SendPackage(pkg []byte) ([]byte, error) {
	_, err := c.conn.Write([]byte(pkg))

	reply := make([]byte, 1024)

	size, err := c.conn.Read(reply)
	if err != nil {
		log.Println("Can't get response")
		return reply, err
	}
	return reply[:size], err
}

func (c *EgtsClient) Stop() {
	c.conn.Close()
}
