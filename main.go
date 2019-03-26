package main

import (
	"github.com/labstack/gommon/log"
	"net"
	"os"
)

var (
	config Config
	logger *log.Logger
)

func main() {
	logger = log.New("-")

	if len(os.Args) == 2 {
		if err := config.Load(os.Args[1]); err != nil {
			logger.Fatalf("Ошибка парсинга конфига: %v", err)
		}
	} else {
		logger.Fatalf("Не задан путь до конфига")
	}
	logger.SetLevel(config.GetLogLevel())

	store := RabbitMQConnector{config: &config.RabbitMQ}
	//store := DebugConnector{}
	if err := store.Init(); err != nil {
		logger.Fatal(err)
	}
	defer store.Close()

	l, err := net.Listen("tcp", config.GetListenAddress())
	if err != nil {
		logger.Fatalf("Не удалось открыть соединение: %v", err)
	}
	defer l.Close()

	logger.Infof("Запущен сервер %s...", config.GetListenAddress())
	for {
		conn, err := l.Accept()
		if err != nil {
			logger.Errorf("Ошибка соединения: %v", err)
		} else {
			go handleRecvPkg(conn, &store)

		}
	}
}
