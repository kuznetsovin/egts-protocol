package main

import (
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
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

	conn, err := amqp.Dial(config.RabbitMQ.GetConnectionString())
	if err != nil {
		logger.Fatalf("Не удалось подключится к RabbitMq: %v", err)
	}
	defer conn.Close()

	rmqChannel, err := conn.Channel()
	if err != nil {
		logger.Fatalf("Не открыть канал RabbitMq: %v", err)
	}
	defer rmqChannel.Close()

	if err = rmqChannel.ExchangeDeclare(
		config.RabbitMQ.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		logger.Fatalf("Не удалось открыть exchange: %v", err)
	}

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
			go handleRecvPkg(conn, rmqChannel)

		}
	}
}
