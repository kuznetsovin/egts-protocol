package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

type EgtsSavePacket interface {
	ToBytes() ([]byte, error)
}

type Connector interface {
	Init() error
	Save(EgtsSavePacket) error
	Close() error
}

type RabbitMQConnector struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	config     *broker
}

func (c *RabbitMQConnector) Init() error {
	var (
		err error
	)
	if c.config == nil {
		return fmt.Errorf("Не корректная ссылка на конфигурацию")
	}

	if c.Connection, err = amqp.Dial(c.config.GetConnectionString()); err != nil {
		return fmt.Errorf("Ошибка установки соединеия RabbitMQ: %v", err)
	}

	if c.Channel, err = c.Connection.Channel(); err != nil {
		return fmt.Errorf("Ошибка открытия канала RabbitMQ: %v", err)
	}

	if err = c.Channel.ExchangeDeclare(
		c.config.Exchange,
		c.config.ExchangeType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		logger.Fatalf("Не удалось открыть exchange: %v", err)
	}

	return err
}

func (c *RabbitMQConnector) Save(msg EgtsSavePacket) error {
	if msg == nil {
		return fmt.Errorf("Не корректная ссылка на пакет")
	}

	innerPkg, err := msg.ToBytes()
	if err != nil {
		return fmt.Errorf("Ошибка сериализации  пакета: %v", err)
	}

	if err = c.Channel.Publish(
		c.config.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        innerPkg,
		}); err != nil {
		return fmt.Errorf("Ошибка отправки сырого пакета в RabbitMQ: %v", err)
	}
	return nil
}

func (c *RabbitMQConnector) Close() error {
	var err error
	if c != nil {
		if c.Channel != nil {
			if err = c.Channel.Close(); err != nil {
				return err
			}
		}
		if c.Connection != nil {
			if err = c.Connection.Close(); err != nil {
				return err
			}
		}
	}
	return err
}
