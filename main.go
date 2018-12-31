package main

import (
	"log"
	"net"
	"os"
)

var (
	config Config
)

func main() {
	if len(os.Args) == 2 {
		if err := config.Load(os.Args[1]); err != nil {
			log.Fatalf("Ошибка парсинга конфига: %v\n", err)
		}
	} else {
		log.Fatalf("Не задан путь до конфига")
	}

	l, err := net.Listen("tcp", config.GetListenAddress())
	if err != nil {
		log.Fatalf("Не удалось открыть соединение: %v\n", err)
	}
	defer l.Close()

	log.Printf("Запущен сервер %s...\n", config.GetListenAddress())
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Ошибка соединения: %v\n", err)
		} else {
			go handleRecvPkg(conn)

		}
	}
}
